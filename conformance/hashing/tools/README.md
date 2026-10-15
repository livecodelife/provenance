# Corpus tooling

`cases.json` is generated, never edited by hand:

```sh
python3 -I conformance/hashing/tools/case_defs.py /tmp/cases.src.json
python3 -I conformance/hashing/tools/reference.py /tmp/cases.src.json conformance/hashing/cases.json
```

- `case_defs.py` holds the inputs and the `why` of every case, plus the
  hand-written expectations for `decode` cases.
- `reference.py` is an independent Python implementation of docs/HASHING.md
  that fills in `output`, `canonical` and `hash`. The Go implementation in
  `record/` is then checked against a second implementation rather than its
  own output.
