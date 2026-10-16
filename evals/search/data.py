"""Corpus and ground truth for the search eval.

Ground truth comes from the repositories' own history, not from the record
text: a tagged commit's subject is a query, the record in its tag is the
answer; a file path is a query, the records whose scope covers it are the
answer. Nothing here calls a model.
"""
import fnmatch
import glob
import os
import random
import re
import subprocess

import yaml

TAG = re.compile(r"\[(prov-\d{4}-[0-9a-f]+)(?:@\d+)?\]")
ANY_ID = re.compile(r"prov-\d{4}-[0-9a-f]+(?:@\d+)?")
BOILERPLATE = re.compile(
    r"^(open provenance record|complete provenance record|refresh|reopen|merge|revise|imprint intent|"
    r"widen scope|mark |pin )", re.I)
WORD = re.compile(r"[a-z0-9]{2,}")


def terms(text):
    return WORD.findall(text.lower())


def load_records(repo):
    """Records of a repo by id, plus an alias -> id map and each record's
    snapshot reasons (the search index covers them)."""
    cfg = {}
    for name in (".provenance.yml", ".linespec.yml"):
        p = os.path.join(repo, name)
        if os.path.exists(p):
            cfg = yaml.safe_load(open(p)) or {}
            cfg = cfg.get("provenance", cfg)
            break
    d = os.path.join(repo, cfg.get("dir", "provenance"))
    recs, alias = {}, {}
    for f in sorted(glob.glob(os.path.join(d, "prov-*.yml"))):
        r = yaml.safe_load(open(f))
        reasons = []
        for s in sorted(glob.glob(os.path.join(d, ".versions", r["id"], "*.yml"))):
            reason = (yaml.safe_load(open(s)) or {}).get("reason") or ""
            if reason:
                reasons.append(reason)
        r["_reasons"] = reasons
        recs[r["id"]] = r
        for a in r.get("aliases") or []:
            alias[a] = r["id"]
    return recs, alias


def doc_text(r, limit=4000):
    """What a record is indexed as: the fields `provenance search` weighs, with
    the short ones first so truncation only ever cuts the intent."""
    parts = [r.get("title") or "", " ".join(r.get("tags") or []),
             " ".join(r.get("affected_scope") or []),
             "\n".join(r.get("constraints") or []),
             "\n".join(r.get("_reasons") or []),
             r.get("intent") or ""]
    return "\n".join(p for p in parts if p)[:limit]


def commit_queries(repo, recs, alias, per_record=2, want=150, seed=7):
    """One query per tagged commit subject, tag and ids stripped."""
    out = subprocess.run(["git", "-C", repo, "log", "--all", "--format=%s"], capture_output=True,
                         text=True, check=True).stdout.splitlines()
    seen, by_rec = set(), {}
    for subj in out:
        m = TAG.findall(subj)
        if not m:
            continue
        rid = alias.get(m[-1], m[-1])
        text = ANY_ID.sub("", subj).replace("[]", "").strip(" -:;,.")
        if rid not in recs or BOILERPLATE.match(text) or len(text.split()) < 4 or text in seen:
            continue
        seen.add(text)
        by_rec.setdefault(rid, []).append(text)
    rng = random.Random(seed)
    qs = []
    for rid in sorted(by_rec):
        texts = by_rec[rid]
        rng.shuffle(texts)
        qs += [{"kind": "commit", "query": t, "truth": [rid]} for t in texts[:per_record]]
    rng.shuffle(qs)
    return qs[:want]


def covers(pattern, path):
    if pattern.endswith("/**"):
        return path.startswith(pattern[:-2])
    if "*" in pattern:
        return fnmatch.fnmatch(path, pattern)
    return pattern == path


def path_queries(recs, want=80, seed=7):
    """One query per file named in some record's scope; the answer is every
    record whose scope covers the file."""
    files = sorted({p for r in recs.values() for p in (r.get("affected_scope") or [])
                    if "*" not in p and os.path.splitext(p)[1]})
    rng = random.Random(seed)
    rng.shuffle(files)
    qs = []
    for f in files:
        truth = sorted(i for i, r in recs.items() if any(covers(p, f) for p in (r.get("affected_scope") or [])))
        if truth:
            qs.append({"kind": "path", "query": f, "truth": truth})
    return qs[:want]


def overlap(query, text):
    """Share of the query's terms that occur in the text: how much the
    wording is shared, whatever the model."""
    q = set(terms(query))
    return len(q & set(terms(text))) / len(q) if q else 0.0
