# Search eval: BM25, embeddings, and their fusion

Run for provisio-mgr and Caleb. It answers one question: is semantic search worth building for agents, or does
`provenance search` (offline BM25) already find the right record? Harness: `evals/search/` (re-runnable, no
model needed for the BM25 arm). Numbers below are from one run on 2026-10-16.

## Method

**Corpus.** Every record of three repos, as of this run: sedum (171), provenance (90) and linespec (470, the
`v2-migration` tree). A record is indexed as its title, tags, scope, constraints, snapshot reasons and intent,
the fields `provenance search` weighs, truncated to 4,000 characters (short fields first, so only the intent is cut).

**Queries and ground truth come from git history, not from the records.**

- *Commit queries:* the subject of a tagged commit with the tag and ids removed is the query; the record in its
  tag (aliases resolved to the surviving record) is the answer. Open, complete, refresh, reopen, merge and
  revise commits are skipped, as are subjects under four words. At most 2 per record, then a seeded sample of
  150 per repo: 422 queries. These are real descriptions of work written while doing it.
- *Path queries:* a file named in some record's scope is the query; the answer is every record whose scope covers
  it. 80 per repo: 240 queries. This is the "what governs this file" question, which `govern <path>` answers exactly.

**Arms.**

- `bm25`: the real `provenance search --format json --limit 50`, run in each repo (so it is what an agent gets).
- `embed`: cosine similarity over ollama `nomic-embed-text` (137M parameters, F16, 768 dimensions, local, no key),
  with its `search_document:` / `search_query:` prefixes.
- `hybrid`: reciprocal rank fusion (k = 60) of the two top-50 lists.

**Metrics.** R@k is the share of the answer found in the top k, out of at most k (with one true record: 1 if it is in
the top k). MRR@10 is the mean of 1 / rank of the first true record, 0 beyond rank 10.

**Leakage check.** The search index includes snapshot reasons, which are sometimes close to a commit subject. Only
27 of the 422 commit queries share 75% or more of their words with the answer's reasons (sedum 4, provenance 18,
linespec 5), so the lexical arm is not carried by them.

## Results

Commit queries (n = 422):

| repo | arm | R@1 | R@3 | R@5 | MRR@10 |
|---|---|---|---|---|---|
| sedum (150) | bm25 | 0.67 | 0.81 | 0.85 | 0.75 |
| | embed | 0.62 | 0.72 | 0.78 | 0.68 |
| | hybrid | 0.73 | 0.81 | 0.85 | 0.79 |
| provenance (122) | bm25 | 0.58 | 0.81 | 0.88 | 0.71 |
| | embed | 0.55 | 0.80 | 0.85 | 0.68 |
| | hybrid | 0.59 | 0.83 | 0.89 | 0.73 |
| linespec (150) | bm25 | 0.68 | 0.77 | 0.83 | 0.74 |
| | embed | 0.59 | 0.67 | 0.71 | 0.64 |
| | hybrid | 0.64 | 0.73 | 0.75 | 0.70 |
| **all** | bm25 | **0.65** | **0.80** | **0.85** | 0.73 |
| | embed | 0.59 | 0.72 | 0.78 | 0.67 |
| | hybrid | **0.66** | 0.79 | 0.82 | **0.74** |

File-path queries (n = 240; R@3 and R@5 fall because several records cover a file):

| arm | R@1 | R@3 | R@5 | MRR@10 |
|---|---|---|---|---|
| bm25 | 0.84 | 0.67 | 0.61 | 0.90 |
| embed | 0.60 | 0.44 | 0.41 | 0.70 |
| hybrid | 0.73 | 0.57 | 0.52 | 0.82 |

Commit queries split by how many of the query's words the answer's text contains (the paraphrase test):

| words shared | n | arm | R@1 | R@5 | MRR@10 |
|---|---|---|---|---|---|
| 75% or more | 260 | bm25 | 0.81 | 0.97 | 0.88 |
| | | embed | 0.73 | 0.91 | 0.81 |
| | | hybrid | 0.80 | 0.95 | 0.87 |
| 50% to 75% | 122 | bm25 | 0.46 | 0.80 | 0.60 |
| | | embed | 0.43 | 0.66 | 0.52 |
| | | hybrid | 0.55 | 0.75 | 0.64 |
| under 50% | 40 | bm25 | 0.15 | 0.28 | 0.20 |
| | | embed | 0.12 | 0.25 | 0.19 |
| | | hybrid | 0.10 | 0.23 | 0.16 |

## What the numbers say

1. **BM25 is at least as good as the embedding model in every cut except one**, and clearly better on paths
   (R@1 0.84 against 0.60) and on linespec. At 422 queries a difference under about 5 points is noise.
2. **Hybrid does not clearly beat BM25.** It gains on sedum (R@1 +6) and in the middle overlap band (R@1 +9),
   loses on linespec and on paths, and ties overall (R@1 0.66 against 0.65, R@5 0.82 against 0.85).
3. **The paraphrase case, where embeddings should win, is where nothing works.** With under half of the
   query's words in the record, all three arms find the record in the top 5 only 23% to 28% of the time, and the
   embedding model is no better than BM25. That is 40 of 422 queries (9%).
4. Most real queries share their wording with the answer: 260 of 422 (62%) share 75% or more, and BM25 gets
   R@5 0.97 there.

## Limits of this eval

- **One small model.** nomic-embed-text is the only embedding model installed. A larger one (mxbai-embed-large, bge)
  may close the gap; none was pulled.
- **Documents are cut to 4,000 characters** and many records in provenance and sedum are mostly literal Go code,
  which a text embedding represents poorly. A code-aware summary or embedding of title plus constraints only was not tried.
- **Commit subjects are one query style.** An agent describing a task in its own words, or a paraphrase written
  by a model, was not tested; that is the case most likely to favour embeddings, and the under-50% band is only a proxy.
- **One answer per commit query.** Several records can be right for a subject, so every arm's score is a lower bound.
- **n is modest** (422 and 240): plus or minus 4 to 5 points at 95% confidence.

## Recommendation

**Keep lexical search; do not build a semantic index now.** It costs a model, an index and a sync story, and in this
test it did not beat what is already there. For file questions the structural commands (`govern`, `context`, graph
neighbours) are exact and BM25 already answers path queries at R@1 0.84.

Reopen the question only with evidence of a gap, in this order:

1. Collect the real failures: log the agent's queries that returned nothing useful (the under-50% band).
2. Re-run this harness with a larger embedding model and with model-written task descriptions as a second query set
   (`run.py --model <name>` takes any ollama embedding model).
3. If a hybrid then wins by more than 5 points of R@5 on those, add it as an opt-in `search --semantic`, not the default.

**Where a vector index would live, if built:** never in git. A per-user cache under `~/.cache/provenance/embeddings/`,
keyed by `(record id, version, model digest)`, so a record is embedded once per version and a model change
invalidates only that model's file. Size: 768 floats per record (3 KB at float32), so linespec's 470 records would be
about 1.4 MB; a cold build of all three repos here took about 2.5 minutes. The eval itself keeps its cache outside the
repo (`--cache`).

## Reproduce

```
python3 -I evals/search/run.py --provenance <provenance binary> \
  --repo sedum=<path> --repo provenance=<path> --repo linespec=<path> \
  --model nomic-embed-text --cache <dir outside any repo> --out results.json
python3 -I evals/search/report.py results.json
python3 -I evals/search/test_metrics.py
```

Needs ollama running with the model pulled, and PyYAML. `--no-embed` runs the BM25 arm alone.
