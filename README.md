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
