# Analysis: how many records could collapse

Read-only analysis for provisio-mgr (proposal A). No code changed. Counts come from the record files on
provenance `main` (09c87ac), sedum `main`, and linespec's `v2-migration` branch (a6e65615).
The scripts parse every `provenance/prov-*.yml` and group imprints by `implements` and `affected_scope`.
They are heuristics over scope, not a reading of each record, so every number below is a *candidate* count.

## What is in the three repos

| | records | blueprints | imprints | imprints with literal code | intent text | of it, literal code |
|---|---|---|---|---|---|---|
| provenance | 87 | 39 | 48 | 30 | 338 KB | 222 KB (66%) |
| sedum | 171 | 90 (+2 briefs) | 79 | 16 | 495 KB | 71 KB (14%) |
| linespec | 470 | 286 | 184 | 4 | 373 KB | 6 KB (2%) |

"Literal code" means the intent carries `setFunction`/`setType`/`setMethod` blocks or "code exactly" text,
the form sedum-grown regions take. Multi-version records: provenance 40 of 87, sedum 50 of 171, linespec 34 of 470.

## Three kinds of collapsible imprint

**1. Versions of an earlier imprint (strict).** Same parent blueprint, identical `affected_scope`, no glob.
This is a later change to the same code that got its own record instead of a `revise`.

| | groups | imprints that would fold |
|---|---|---|
| provenance | 0 | 0 |
| sedum | 0 | 0 (6 glob-scope groups, 11 imprints, are separate design decisions; see below) |
| linespec | 9 | 14 |

provenance and sedum are already clean here, because the herd used `revise` for follow-up changes. linespec's
14 are real: for example three `release prep for vX` records on the same files, and `Guard Lint --save-scope
against remote records` followed by two later fixes to the same guards. This is the case rule B now prevents.

**2. Literal-code imprints, one per file.** The `go:` records that replay one Go file each. A package-level
imprint (the `sedum adopt` shape: one imprint per package) would hold them.

| | literal imprints | would fold into | survivors |
|---|---|---|---|
| provenance | 30 | 26 into the package imprints (cmd/provenance 26 records, record 2) | 4 |
| sedum | 16 | 7 | 9 |
| linespec | 4 | 0 | 4 |

The cost: a package record carries all its files' code in one intent. cmd/provenance would carry most of the
222 KB of literal code, and every `revise` rewrites that intent and its snapshot. That is the trade against fewer records.

**3. Same parent, same directory (upper bound, needs judgement).** Imprints under one parent whose scope
sits in one directory.

| | would fold | note |
|---|---|---|
| provenance | 26 (cmd/provenance: 9, 8, 6, 3 per parent) | different commands, not versions of each other |
| sedum | 17 (evals: 8, internal/resolve: 3, ...) | separate decisions that happen to share an area |
| linespec | 21 (pkg/provenance: 6 and 3, cmd/linespec: 4, ...) | mixed |

Folding these loses information: they are different commands and decisions. It is a ceiling, not a recommendation.

## What `provenance consolidate` would reduce counts to

Imprints only; blueprints are untouched and not counted down.

| | records now | kind 1 only | kind 1 + 2 | kind 1 + 2 + 3 (ceiling) |
|---|---|---|---|---|
| provenance | 87 | 87 | 61 | 51 |
| sedum | 171 | 171 | 164 | 151 |
| linespec | 470 | 456 | 456 | 449 |

(The last column is the union of kinds 2 and 3, so the groups are not double counted.)

Reading: the savings are modest where `revise` was already used (sedum: 0 to 20 of 171) and concentrated in
literal-code files for provenance (26 of 87). Blueprints are the larger share everywhere (provenance 39 of 87,
sedum 90 of 171, linespec 286 of 470). In the pair workflow every change adds a blueprint and an imprint, so
the lever that matters for growth is rule B (revise instead of a new pair), not consolidation after the fact.

## How a consolidate would work

It is the same operation `migrate` does for folded parents, so the machinery exists:

1. The survivor is the earliest imprint of the group. Each later imprint becomes a version of it, in
   `created_at` order, with its title as the reason. Its id becomes an alias.
2. Children that pinned a folded imprint are re-pinned to the survivor's final version (migrate's `pinTable`,
   reported as "Re-pinned to the final version of a folded parent").
3. Snapshots are written in order, so the chain verifies. The `superseded -> implemented` problem `check`
   had with the migration commit applies here too, and for the same reason: it rewrites record files that
   already exist at the base commit.

Open questions for provisio-mgr:

- Which tier is the rule? Kind 1 is mechanical and safe. Kind 2 depends on the `sedum adopt` design (C).
- A consolidate rewrites history of governed records. Should it need a flag, as `migrate` does?
- Aliases keep old tags (`[prov-2026-xxxx@1]`) resolvable; old commit tags on folded ids must keep working.
