#!/usr/bin/env python3
"""Search eval: BM25 (`provenance search`), embeddings, and their fusion.

    python3 -I evals/search/run.py --provenance /path/to/provenance \
        --repo sedum=/path/sedum --repo provenance=/path/provenance \
        --model nomic-embed-text --cache /tmp/search-eval-cache --out results.json

The embedding index lives in --cache and is never written under a repo.
"""
import argparse
import hashlib
import json
import math
import os
import subprocess
import sys
import urllib.request

sys.path.insert(0, os.path.dirname(os.path.abspath(__file__)))
import data  # noqa: E402
import metrics  # noqa: E402

OLLAMA = os.environ.get("OLLAMA_HOST", "http://localhost:11434")


def embed(model, texts, prefix, cache):
    """Embeddings for texts, cached by (model, prefix, text)."""
    os.makedirs(cache, exist_ok=True)
    out, todo = {}, []
    for t in texts:
        key = hashlib.sha256((model + "\0" + prefix + "\0" + t).encode()).hexdigest()
        path = os.path.join(cache, key + ".json")
        if os.path.exists(path):
            out[t] = json.load(open(path))
        else:
            todo.append((t, path))
    for i in range(0, len(todo), 16):
        batch = todo[i:i + 16]
        body = json.dumps({"model": model, "input": [prefix + t for t, _ in batch], "truncate": True}).encode()
        req = urllib.request.Request(OLLAMA + "/api/embed", body, {"Content-Type": "application/json"})
        vecs = json.load(urllib.request.urlopen(req, timeout=300))["embeddings"]
        for (t, path), v in zip(batch, vecs):
            json.dump(v, open(path, "w"))
            out[t] = v
    return out


def cosine(a, b):
    return sum(x * y for x, y in zip(a, b)) / (math.sqrt(sum(x * x for x in a)) * math.sqrt(sum(y * y for y in b)))


def bm25(binary, repo, query):
    r = subprocess.run([binary, "search", "--format", "json", "--limit", "50", query], cwd=repo,
                       capture_output=True, text=True)
    if r.returncode != 0:
        raise SystemExit("search failed in %s: %s" % (repo, r.stderr.strip()))
    return [h["id"] for h in json.loads(r.stdout)["results"]]


def model_version(model):
    req = urllib.request.Request(OLLAMA + "/api/show", json.dumps({"model": model}).encode(),
                                 {"Content-Type": "application/json"})
    info = json.load(urllib.request.urlopen(req, timeout=30))
    return {"model": model, "digest": info.get("details", {}).get("family"), "params": info.get("details", {}).get("parameter_size"),
            "quantization": info.get("details", {}).get("quantization_level")}


def main():
    ap = argparse.ArgumentParser()
    ap.add_argument("--provenance", required=True, help="provenance binary for the BM25 arm")
    ap.add_argument("--repo", action="append", required=True, help="name=path")
    ap.add_argument("--model", default="nomic-embed-text")
    ap.add_argument("--doc-prefix", default="search_document: ")
    ap.add_argument("--query-prefix", default="search_query: ")
    ap.add_argument("--cache", required=True)
    ap.add_argument("--out", required=True)
    ap.add_argument("--no-embed", action="store_true")
    a = ap.parse_args()

    result = {"model": None if a.no_embed else model_version(a.model), "repos": {}}
    for spec in a.repo:
        name, repo = spec.split("=", 1)
        recs, alias = data.load_records(repo)
        qs = data.commit_queries(repo, recs, alias) + data.path_queries(recs)
        ids = sorted(recs)
        texts = {i: data.doc_text(recs[i]) for i in ids}
        if not a.no_embed:
            dvec = embed(a.model, list(texts.values()), a.doc_prefix, a.cache)
            qvec = embed(a.model, [q["query"] for q in qs], a.query_prefix, a.cache)
        rows = {}
        for q in qs:
            ranked = {"bm25": bm25(a.provenance, repo, q["query"])}
            if not a.no_embed:
                qv = qvec[q["query"]]
                ranked["embed"] = [i for _, i in sorted(((-cosine(qv, dvec[texts[i]]), i) for i in ids))][:50]
                ranked["hybrid"] = metrics.rrf([ranked["bm25"], ranked["embed"]])
            best = max(data.overlap(q["query"], texts[t]) for t in q["truth"] if t in texts)
            q["overlap"] = round(best, 2)
            q["ranked"] = {k: v[:10] for k, v in ranked.items()}
            rows[q["query"]] = ranked
        result["repos"][name] = {"records": len(recs), "queries": qs}
        print("%s: %d records, %d queries" % (name, len(recs), len(qs)), file=sys.stderr)
    json.dump(result, open(a.out, "w"), indent=1)


if __name__ == "__main__":
    main()
