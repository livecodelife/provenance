# Migrating from v1 (linespec) to v2

v2 is the record model `provenance` manages from release 2.0.0 on. This page
says what changes and what `provenance migrate` does. The steps, flags and
rollback are in the README section
[Migrate from linespec v1](../README.md#migrate-from-linespec-v1).

## What changes in v2

| | v1 (linespec) | v2 (provenance) |
| --- | --- | --- |
| Editing a record | In place. History is only in git. | `provenance edit` or `revise` writes a new **version**. The record file always holds the latest one. |
| What a version is | Nothing in the file. | A numbered snapshot under `provenance/.versions/<id>/`, with a content hash listed in `.provenance/versions.jsonl`. |
| Tiers | One kind of record. | **Blueprints** state a subject and its constraints. **Imprints** are bounded units under a blueprint (a method, a command). |
| Parents | A parent id. | A pin: `implements` plus `implements_version`. `provenance stale` lists a child pinned to an older parent version, and `provenance repin` confirms it against the new one. |
| Commits | Any commit. | A commit touching governed files is tagged `[id@N]` with the record and version it belongs to. `provenance check --range base..HEAD` verifies that. |
| Write lock | None. | A governed file can be edited only while its record is open. |
| Config | The provenance block of `.linespec.yml`. | `.provenance.yml`. |

The ids of kept records do not change, so references to them keep working.

## What `provenance migrate` does

1. Reads every v1 record in the repository.
2. Folds chains of records that were really versions of one record into a
   single v2 record whose versions are the old ones, oldest first. The plan
   (`migrate --plan`, which writes nothing) shows every chain, and you can keep a
   record separate with `--keep-separate <id>`.
3. Writes the kept records as v2 files, writes one snapshot per v1 record,
   writes the version manifest, and re-pins children whose parent was folded.
4. Replaces the provenance block of `.linespec.yml` with `.provenance.yml`,
   removes `.linespec/hash_manifest.json`, and writes the `.gitattributes`
   lines that mark generated files.

It never touches your source code, your git history, your git hooks (run
`provenance hooks install` yourself), or the content of any record: a record's
text in v2 is the text it had in v1.

## After migrating

- `provenance lint` should report no errors and `provenance stale` nothing.
- Commit the result as one commit, so it is easy to review and to revert.
- Replace calls to linespec's `provenance` subcommand with `provenance` in CI
  and Makefiles, and pin the installer with `PROVENANCE_VERSION`.
- To go back, `git revert` the migration commit. Don't mix v1 and v2 tools in
  one repository: v1 tooling reads only v1 records.
