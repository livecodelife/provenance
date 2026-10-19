# Lexicons for published provenance records

The one place the published field names are decided (M0). Everything else is derived: the lexicon JSON in
`lexicons/dev/provisio/provenance/`, the generated types in `lexicons/gen/`, the Go mapping in
`record/published.go`, the TypeScript twin in `ts/src/published.ts`, and `docs/RECORDS-API.md` in the backend.
Where DESIGN's "AT Proto storage and publishing" says what is published, this says how it is spelled.

## Conventions

- **Names.** The record's snake_case field names in camelCase (`content_hash` is `contentHash`). The hash *input*
  keeps its snake_case canonical JSON (HASHING.md §5); the mapping happens before hashing and changes no hash.
- **Absent equals empty.** An empty string, list or zero is omitted, as in HASHING.md. `title` and `intent` are
  always present. A lexicon has no null and no float.
- **Hashes** are `"sha256:"` plus 64 lowercase hex digits (71 characters).
- **Dates.** An envelope `createdAt` is an RFC 3339 UTC `datetime`. `content.createdAt` is the record's creation
  day, plain `YYYY-MM-DD` (not a datetime), kept as the record has it.
- **Open vocabularies.** `reasonKind`, `origin`, `type` and `status` use `knownValues`, so a new value is not a
  breaking change. `author` is a plain string (a git author's email today), not a DID.

## `dev.provisio.provenance.defs#content`

The record minus its lifecycle and envelope fields. Shared by the head and every version.

| Field | Type | Notes |
|---|---|---|
| `title` | string | required |
| `type` | string | brief, blueprint or imprint |
| `createdAt` | string | `YYYY-MM-DD` |
| `author` | string | |
| `intent` | string | required, up to 200,000 bytes (the largest real intent is about 28 KB) |
| `constraints`, `affectedScope`, `forbiddenScope` | string[] | |
| `related`, `tags`, `aliases` | string[] | `aliases` are old ids folded in by migration |
| `implements`, `implementsRepo` | string | parent id, and the repo that holds it when not local |
| `implementsUri` | at-uri | the parent's head, `at://<did>/dev.provisio.provenance.record/<recordId>`, when it is published in another account. Optional; requires `implements`, whose value is that record id. Lineage only: not in the content hash. |
| `implementsVersion` | integer | the parent version last confirmed; 0 or absent is unpinned |
| `associatedSpecs` | `{path, type?}`[] | `path` required (empty for a spec without one); `runCommand` is never published |
| `associatedTraces`, `monitors` | string[] | |
| `locked` | boolean | |

Not in content: `status`, `sealedAtSha`, `version`, `reason`, `supersedes`, `supersededBy`, `extends`. The record
id is the head's key (`recordId` in a version).

## `dev.provisio.provenance.version` (key: TID, never updated)

| Field | Type | Notes |
|---|---|---|
| `record` | at-uri | the head this belongs to |
| `recordId` | string | the provenance id, the head's key |
| `version` | integer ≥ 1 | |
| `contentHash` | hash | |
| `parentHash` | hash | the previous version's `contentHash`; absent on version 1 |
| `parent` | strongRef | the previous version record; absent on version 1 |
| `reason`, `reasonKind` | string | `reason` is required after version 1; `reasonKind` is `fix` today |
| `author`, `createdAt` | string, datetime | required |
| `commit` | string | the git commit published from, optional |
| `specHashes` | string[] | one per `content.associatedSpecs`, same order, required |
| `content` | defs#content | required |
| `origin`, `baseVersion`, `baseContentHash` | | UI-made versions: `git` or `ui`, and what the editor started from |

**Two chains.** `parentHash` is the provenance chain (the hash of the previous version's content); `parent` is the
AT Proto chain (the previous record's `{uri, cid}`). A verifier checks both; they cover different things, because an
AT CID covers the whole AT record and `contentHash` covers only content (HASHING.md §4).

**`specHashes`.** Each entry is `sha256:<hex>`, the constant `record-file/v1` for a spec that is a record file, or the
empty string when the spec could not be read (the snapshot's null; a real hash can never be empty).

**Verifying.** `contentHash` must equal the hash recomputed from `content` and `specHashes` alone; this is
`Snapshot.Verify` in Go, `verifyVersion` in TypeScript, and the corpus kind `version_hash` (HASHING.md §6) pins
both to the same answer. A different number of spec hashes than specs is an error.

## `dev.provisio.provenance.record` (key: the provenance id, mutable head)

| Field | Type | Notes |
|---|---|---|
| `status` | string | draft, open, implemented, deprecated |
| `version`, `contentHash`, `genHash` | | the current version and hashes, required |
| `latest` | strongRef | the newest version record, required |
| `content` | defs#content | the *current* content, including the parent pin |
| `source` | `{repo, path, commit}` | where it lives in git, optional |
| `project` | at-uri | the `dev.provisio.project` it belongs to, optional, so a head without one validates |
| `updatedAt` | datetime | required |

The head is updated with `swapRecord` on its CID. `provenance repin` moves only `content.implementsVersion` on the
head; it writes no version, so the newest version's `content.implementsVersion` can be older than the head's.

## Not published in v1

`pins.jsonl` (the repin log, `stale/`) is **not** published. A pin appears only as the head's current
`content.implementsVersion`. Publishing the log would need its own record type; open question for later.

## Limits

String and array `maxLength` values are set in the lexicons. The PDS record size cap (believed 1 MiB) and the
200-writes-per-`applyWrites` cap are the real limits; prov-web verifies both in the live suite.

## Generated code

`lexicons/gen/regen.sh` rebuilds both outputs; running it twice changes nothing (checked).

| | Tool and pin | Output |
|---|---|---|
| TypeScript | `@atproto/lex-cli` 0.10.14 (2026-09-29) | `lexicons/gen/ts/src`: types, validators, `lexicons.ts`; the client class is dropped |
| Go | indigo `cmd/lexgen` at commit `d8478c47f0ce6ba53f214e1294c219732795e03b` (2026-05-15), then `cbor-gen` | `lexicons/gen/go`: its own module, so the CLI does not depend on indigo until M4 imports it |

Runtime dependencies (exact): `@atproto/lexicon` 0.7.15, `@atproto/xrpc` 0.8.14, `multiformats` 14.0.5; dev:
`typescript` 7.0.2, `@types/node` 26.6.4. Chosen as the newest versions published before 2026-10-02.
`lexicons/gen/ts/package-lock.json` sha256 `fbc297869d67dd945eaf12c61eddf6082c042259acb6e87ad5631ac53a4d481e`;
`lexicons/gen/go/go.sum` is committed (sha256 `7c3f2f7e5361366a758b76e0520c94744cf85526344eae763b5c70c4ece215f1`).
indigo requires Go 1.26, so `go` downloads its 1.26.9 toolchain for that module.
`lexicons/vendor/com/atproto/repo/strongRef.json` is copied from bluesky-social/atproto at
`a7c42cfe3981e8f1491f8d7a09fde041a9e83afe`, because lex-cli and lexgen need it to resolve `strongRef`.

The generated Go types carry JSON and CBOR (DAG-CBOR, what a PDS stores) methods. lexgen leaves the CBOR methods to
`cbor-gen`, and the types register themselves with indigo's `lexutil` only once those methods exist, which is why
`regen.sh` holds the registration back for one step.

## Checks

- `go test ./record` (`TestPublished*`, `TestVersionHashCorpus`): the mapping, the samples, every snapshot of this
  repository round-tripped, tampering rejected.
- `node --test` in `ts/`: the `version_hash` cases and the samples against `ts/src/published.ts`.
- `npm test` in `lexicons/gen/ts` (after `npm ci`): the sample head and three versions through the generated
  validators, plus negative cases.
- `go test ./...` in `lexicons/gen/go`: the same samples through the generated types, as JSON and as CBOR.

`lexicons/samples/` holds the sample head and three versions. The Go mapping writes them (`UPDATE_SAMPLES=1 go test
./record -run PublishedSamples`) and all three languages read them.
