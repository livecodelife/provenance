# Analysis: version snapshot storage

Read-only analysis for provisio-mgr. No code changed. Numbers are from linespec's `v2-migration` branch
(a6e65615) and the `main` branches of provenance (09c87ac) and sedum. The readers were found by grepping
`Store.Snapshot`, `Store.Versions` and `versions.jsonl` on provenance `main`.

## Today

One YAML file per version at `provenance/.versions/<id>/<version>-<hash>.yml`: an envelope (`id`, `version`,
`content_hash`, `parent_hash`, `reason`, `author`, `created_at`, `spec_hashes`) plus `content`, the record
without its lifecycle fields.

| | files | bytes | gzip (all files together) | median / p90 / max file |
|---|---|---|---|---|
| linespec | 527 (470 dirs) | 1,488,421 (1.4 MB) | 406 KB | 2.1 KB / 5.3 KB / 27 KB |
| sedum | 291 | 1,669,870 | 338 KB | |
| provenance | 184 | 1,155,488 | 171 KB | |

linespec's PR #239 adds 26,232 lines for these files (about 50 per file). The git pack of linespec is 127 MiB, so
the snapshots are about 0.3% of what is stored once compressed. The cost is review noise in the one migration PR
and a larger working tree, not repository size. The `.gitattributes` lines already on #239 address the noise.

In native v2 repos the average file is larger (provenance 6.3 KB against linespec's 2.8 KB) because
literal-code intents are copied whole into every version. Growth follows versions times intent size.

## What reads snapshots

| Reader | Needs | Notes |
|---|---|---|
| `lint` PROV026 | content, spec_hashes, parent_hash, version, file names | `Verify()` recomputes content_hash from content and spec_hashes for every snapshot; checks the chain, consecutive numbers, record version = latest, and the manifest against the latest hash |
| `lint` PROV024 | latest content_hash only | compares the record's hash with the latest snapshot's |
| `Commit` / `Append` (every `open`, `revise`, `edit`, `complete`) | latest snapshot | bump rules, no write over an existing version |
| manifest (`.provenance/versions.jsonl`) | latest hash per record | 470 lines, 60 KB in linespec |
| `check --staged/--range` | file names and content of version N-1; base-commit file list | `snapshotOwner` is path based (`.versions/<id>/`); a rewritten snapshot is STATE001; `deletedFromPreviousScope` reads version N-1's `affected_scope` |
| `diff` | content of two versions | the only user-facing content reader |
| `history` | envelope | reason, author, created_at, hash; commits come from `git log --grep` on the tag |
| `graph`, `context`, `search` | reason only | search indexes reasons as a field |
| tag resolution in lint/commit | envelope of version N | |
| `stale` | nothing | its `pins.jsonl` lives in the same directory and must survive any repack |
| UI (`linespec-provenance-ui/api`) | nothing found | no references to `.versions` |
| publish (DESIGN, not built) | full snapshots | "publishing copies snapshots" |

So content is read by Verify, `diff`, and one check path; everything else reads the envelope. Lint on
linespec (470 records, 527 snapshots) takes 0.36 s, so snapshot reading is not a speed problem today.

## What gives tamper evidence

1. `Verify()` binds content and spec hashes to `content_hash`. Editing a snapshot's content alone fails.
2. `parent_hash` chains version N to N-1. Editing a middle version breaks the next link.
3. The manifest pins each record's latest hash. A forger must rewrite the whole chain and the manifest together.
4. `check --range` compares snapshot files to the base commit and reports a rewrite (STATE001). This is the real
   immutability guard, and it needs git.

Steps 1 to 3 detect partial edits only. A consistent rewrite of a whole chain plus manifest passes them; steps
1 to 3 do not stop it, git history and `check` do. That limits what any option may give up: it must keep
the content-to-hash binding (1), the chain (2) and add-only against a base (4).

## Options

**A. Keep the format (status quo plus `.gitattributes`).** Done on #239. Per-file immutability stays visible to
`check`. Cost: file count and PR lines on migration, nothing else.

**B. One packed file per record.** An append-only `<id>.snapshots.yml`. linespec: 527 files become 470; bytes
and lines unchanged (the content is the same, one `---` apart). Saves nothing that matters. It loses path-based
immutability (`snapshotOwner`, STATE001 would have to parse and diff inside a file), and a revise then edits an
existing file instead of adding one. It changes DESIGN's "immutable snapshot named by its content hash".

**C. Commit SHA plus hash for migrated v1 versions (stub instead of a copy).** The stub keeps the envelope and
`spec_hashes` and a `source: {commit, path}`; content is re-derived from git. Estimate for linespec: a stub is
about 0.45 KB, so 527 stubs are about 0.24 MB against 1.49 MB raw (-84%) and about 6,500 lines against 26,232
(-75%). In pack terms it saves well under 0.4 MB of a 127 MiB pack. Problems:

- The snapshot is not the v1 file's bytes. Migrate rewrites it into the v2 shape (type, `aliases`, folded
  unions of constraints and scope from bug records, lifecycle fields stripped) and resolves `spec_hashes`
  at migration time, against files that have since changed. "The same bytes exist in git history" is true of
  the v1 file, not of the snapshot. Re-deriving means keeping migrate's transform in the verifier forever.
- It needs full history. A shallow CI clone (`fetch-depth: 1`) cannot verify or `diff`. Verify stops being
  self-contained, and publishing needs the content materialised anyway.
- It applies to migrated versions only (all 527 of linespec's today, none of new ones), so it is a one-off
  shrink with a permanent second code path.

**D. Envelope only for every version, content from git.** Same as C for everything. Rejected for C's reasons,
and because `diff` and the check path would depend on git being present.

**E. Do nothing in the store, cut the migration PR's noise instead.** Put the snapshots and manifest in their own
commit (the PR's reviewers can skip it), with the `.gitattributes` lines. Costs nothing in the format.

## Recommendation

**A plus E: keep the storage model as it is.** The numbers do not justify a change: 1.4 MB raw and 0.4 MB
compressed for 527 versions, 0.36 s to verify all of them, in a 127 MiB repository. Every compact option
(B, C, D) trades something the checks rely on, which is the per-file immutability and a self-contained verifier,
for a saving of about 1 MB. The only real problem is the migration PR's first impression, and `.gitattributes`
(done) plus a separate commit fixes it.

Revisit if a repo's native snapshots grow into the tens of megabytes. The driver there is literal-code intents
copied whole into each version (provenance averages 6.3 KB per file). The honest fix for that is delta
storage inside the snapshot (store the content as a patch against the previous version, verified by
applying it), not a different location; it keeps the chain and Verify, and it is a design change to
bring to Caleb first.

No change to DESIGN's storage model is recommended, so nothing needs Caleb's decision.
