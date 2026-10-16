#!/usr/bin/env python3
"""Turn run.py's JSON into the tables of docs/EVAL-SEARCH.md."""
import json
import os
import sys

sys.path.insert(0, os.path.dirname(os.path.abspath(__file__)))
import metrics  # noqa: E402

ARMS = ["bm25", "embed", "hybrid"]


def table(title, groups):
    lines = ["**%s**" % title, "", "| group | arm | n | R@1 | R@3 | R@5 | MRR@10 |", "|---|---|---|---|---|---|---|"]
    for gname, qs in groups:
        for arm in ARMS:
            rows = [(q["ranked"][arm], q["truth"]) for q in qs if arm in q["ranked"]]
            if not rows:
                continue
            s = metrics.summarize(rows)
            lines.append("| %s | %s | %d | %.2f | %.2f | %.2f | %.2f |" % (gname, arm, s["n"], s["R@1"], s["R@3"], s["R@5"], s["MRR"]))
    return "\n".join(lines) + "\n"


def main():
    res = json.load(open(sys.argv[1]))
    print("Model: %s\n" % json.dumps(res["model"]))
    allq = []
    for name, r in res["repos"].items():
        allq += r["queries"]
    for kind in ("commit", "path"):
        groups = [(n, [q for q in r["queries"] if q["kind"] == kind]) for n, r in res["repos"].items()]
        groups.append(("all", [q for q in allq if q["kind"] == kind]))
        print(table("Queries from %s" % ("tagged commit subjects" if kind == "commit" else "file paths"), groups))
    cq = [q for q in allq if q["kind"] == "commit"]
    buckets = [("wording shared >= 0.75", [q for q in cq if q["overlap"] >= 0.75]),
               ("0.5 to 0.75", [q for q in cq if 0.5 <= q["overlap"] < 0.75]),
               ("< 0.5", [q for q in cq if q["overlap"] < 0.5])]
    print(table("Commit queries by how many of the query's words the record text contains", buckets))


if __name__ == "__main__":
    main()
