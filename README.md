# provenance

Versioned provenance records: the shared record library (schema, canonical
hashing, version snapshots, record store) and the `provenance` CLI, ported
from `linespec provenance`.

Design: [docs/DESIGN.md](docs/DESIGN.md).

Until this repo can manage itself, its own records are managed with
`linespec provenance` (v1 model) and migrated once `provenance migrate` exists.

## Hooks

`provenance hooks install` writes three git hooks (existing ones are kept as
`.bak`): `pre-commit` (lint staged records, run specs on completion),
`commit-msg` (`check --staged`) and `pre-push`.

The `pre-push` hook is opt-in per repository. It does nothing unless
`.provenance.yml` sets `pre_push_command`; then it runs that command with
`sh -c` in a clean `git archive` checkout of every commit being pushed, and a
failing commit blocks the push. This repository sets
`pre_push_command: go build ./...`, so every pushed commit builds. Other
repositories leave it unset and are not slowed down.

## `provenance graph`

`provenance graph [--root <id>] [--filter <status>] [--format human|json|dot]`
renders the records, their parents, pins and versions. The JSON shape is
stable; other tools read it:

```json
{
  "nodes": [
    {
      "id": "prov-2026-0000e001",
      "title": "…",
      "type": "blueprint",            // brief | blueprint | imprint
      "status": "open",               // draft | open | implemented | deprecated
      "version": 2,                   // the record's current version
      "versions": [                   // every version written, oldest first
        {"version": 1},
        {"version": 2, "reason": "…", "reason_kind": "fix"}   // reason_kind only for fixes
      ],
      "implements": "prov-…",         // only on a child
      "implements_version": 1,        // the parent version the child was written against
      "shared": true                  // only on a parent that lives in a shared repo (read-only)
    }
  ],
  "edges": [
    {
      "from": "prov-2026-0000e002",   // the child
      "to": "prov-2026-0000e001",     // the parent
      "edge": "implements",           // implements | related
      "pinned": 1,                    // implements edges: the child's implements_version
      "parent_version": 2,            // implements edges: the parent's current version
      "stale": "direct"               // omitted when current; direct | transitive | unknown
    }
  ]
}
```

`stale` comes from `provenance stale`: `direct` (the pin is behind the parent),
`transitive` (an ancestor is directly stale) or `unknown` (the parent is in
another repo that is not cached, or does not resolve). There is no
`supersedes` edge: a change to a record is a new version of that record.
Nodes and edges are sorted by id, so the output is deterministic.

## `provenance context`

`provenance context [targets...] [--files a,b] [--format human|compact|json]`
shows what an agent needs to know about paths or records. A target that is a
record id gives a record report, anything else is a repo-relative path. The
JSON shape is stable:

```json
{
  "targets": [
    {
      "target": "cmd/provenance/main.go",
      "kind": "path",
      "exempt": true,                    // only when set: no provenance rules apply (exclude_paths, CLI state, the provenance dir)
      "governing": [                     // non-deprecated records whose affected_scope covers the path
        {"id": "prov-…", "type": "imprint", "status": "open", "version": 2, "title": "…",
         "latest_reason": "…", "implements": "prov-…", "implements_version": 1,
         "stale": "direct",              // omitted when current: direct | transitive | unknown
         "evidence": "fresh",            // implemented records only: fresh | stale | unknown | missing
         "specs": ["path/to/proof_test.go"]}
      ],
      "ancestry": {"prov-…": [{"id": "prov-…", "version": 3, "status": "open", "title": "…", "shared": true}]},
      "conflicts": ["prov-…", "prov-…"], // more than one open record covers the path
      "next": {"kind": "complete", "command": "…", "reason": "…"}   // the first action of `provenance next <path>`
    },
    {
      "target": "prov-2026-0d61fd8d",
      "kind": "record",
      "record": {
        "id": "…", "type": "…", "status": "…", "version": 1, "title": "…", "specs": ["…"],
        "intent": "…", "constraints": ["…"], "affected_scope": ["…"], "forbidden_scope": ["…"],
        "versions": [{"version": 1, "reason": "…"}],
        "parents": [{"id": "…", "version": 1, "status": "…", "title": "…"}],
        "children": [{"id": "…", "version": 1, "pinned": 1, "stale": "direct"}],
        "commits": ["abc1234 subject of a commit tagged with the record"]   // newest first, at most 20
      }
    }
  ]
}
```

Record text (titles, reasons, intent) is data: `human` and `compact` strip
control characters and cap line lengths (compact titles at 80 characters); the
JSON keeps the full text, JSON-escaped. Everything is sorted by id.

## `provenance search`

`provenance search [query] [--query "..."] [--limit N] [--status s] [--type t] [--format human|json]`
finds records by words. It is offline and deterministic (no model, no index
file, no network): lowercase terms of at least two letters or digits, minus a
small stop-word list, scored with BM25 over the id and aliases (weight 5), title
(4), tags (3), scope paths (2), intent (1), constraints (1) and version reasons
(1). Ties break by id. Deprecated records only appear with `--status deprecated`;
cached shared-repo records are searched too. The JSON shape is stable:

```json
{
  "query": "reconcile permissions",
  "backend": "lexical",                  // a semantic backend may be added later under another value
  "results": [
    {"id": "prov-…", "version": 1, "status": "open", "type": "imprint", "title": "…",
     "score": 1.455,                      // BM25, rounded to 3 decimals
     "shared": true,                      // only for a record from a cached shared repo
     "matched": ["title", "intent"],      // fields with at least one query term
     "excerpt": "…"}                      // best matching field, control characters stripped, at most 120 characters
  ]
}
```

## `provenance govern`

`provenance govern [files...] [--files a,b] [--format human|json] [--json]` is
the light lookup of which **active** records (open or implemented) govern files:
a record governs a file when it is in allowlist mode, its `affected_scope` covers
it and `forbidden_scope` does not match. Files that are exempt (`exclude_paths`,
CLI state, the provenance directory) are listed separately. The exit code is 0
whenever the lookup ran. The JSON shape is stable (hooks and agents read it):

```json
{
  "files": ["a.go", "README.md"],
  "exempt": ["README.md"],
  "governing": [
    {"id": "prov-…", "type": "imprint", "status": "open", "version": 2, "title": "…",
     "files": ["a.go"]}                  // the given files this record covers
  ],
  "next": {"kind": "complete", "command": "…", "reason": "…"}   // the first action of `provenance next <files>`
}
```

## What the CLI can do today

| Area | Commands |
| --- | --- |
| Records and versions | `create`, `open`, `revise`, `edit` (`--set-spec`, `--refresh-specs`, `--add-scope` …), `add-scope`, `complete`, `deprecate`, `repin`, `history`, `diff`, `status` |
| Checks | `lint` (human, json, SARIF), `check --staged`, `--commit`, `--range`, `--evidence` (human, SARIF), `run-specs`, `stale` |
| Proof | `complete` records evidence; `evidence refresh` re-verifies after a code-only change |
| Understanding | `next` (the single next action), `context` (a path or a record), `govern` (which records govern files), `graph` (human, json, dot), `search` (offline BM25) |
| Enforcement | `reconcile` (default-deny write lock), `hooks install` (pre-commit, commit-msg, opt-in pre-push) |
| Repo upkeep | `migrate` (v1 to v2), `compile` (rebuild the manifest), `sync` (shared-repo cache) |
| Agents | `install-skills`, `install-plugin` (the v2 skill and Claude Code plugin; the plugin hooks only call `provenance next`) |

Everything above is generated by sedum from the records in `provenance/`
except the tests, the assets under `assets/`, and this README.
