"""Ranking metrics for the search eval. Pure functions, no I/O."""


def recall_at(ranked, truth, k):
    """Share of the truth found in the top k, out of at most k: with one true
    record this is 1 if it is in the top k and 0 otherwise."""
    if not truth:
        return 0.0
    hit = len(set(ranked[:k]) & set(truth))
    return hit / min(k, len(truth))


def reciprocal_rank(ranked, truth, cutoff=10):
    """1 / rank of the first true record within the cutoff, else 0."""
    truth = set(truth)
    for i, r in enumerate(ranked[:cutoff], 1):
        if r in truth:
            return 1.0 / i
    return 0.0


def rrf(rankings, k=60, top=50):
    """Reciprocal rank fusion of several ranked id lists."""
    score = {}
    for ranked in rankings:
        for i, r in enumerate(ranked[:top], 1):
            score[r] = score.get(r, 0.0) + 1.0 / (k + i)
    return [r for r, _ in sorted(score.items(), key=lambda kv: (-kv[1], kv[0]))]


def summarize(rows):
    """rows: list of (ranked, truth). Returns n, R@1, R@3, R@5 and MRR@10."""
    n = len(rows)
    if n == 0:
        return {"n": 0, "R@1": 0.0, "R@3": 0.0, "R@5": 0.0, "MRR": 0.0}
    out = {"n": n}
    for k in (1, 3, 5):
        out["R@%d" % k] = sum(recall_at(r, t, k) for r, t in rows) / n
    out["MRR"] = sum(reciprocal_rank(r, t) for r, t in rows) / n
    return out
