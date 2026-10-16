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
