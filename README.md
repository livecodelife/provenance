# provenance

Versioned provenance records: the shared record library (schema, canonical
hashing, version snapshots, record store) and the `provenance` CLI, ported
from `linespec provenance`.

Design: [docs/DESIGN.md](docs/DESIGN.md).

Until this repo can manage itself, its own records are managed with
`linespec provenance` (v1 model) and migrated once `provenance migrate` exists.
