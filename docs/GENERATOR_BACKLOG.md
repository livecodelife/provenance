# Generator backlog

Hand-written library code that predates COMMON.md's "Generate first" rule.
provisio-mgr grandfathered it (prov-2026-7961a897): it is proven by the hashing
corpus, linespec lint parity and the migration oracle, and converting it now
would be risk without payoff. New library code goes through sedum's `go`
package. This list is for later, when a package changes enough that regrowing
it is cheaper than editing it.

The `go` package's actions (provenance-cli `generators/go`) are `addImport`,
`setConst`, `setVar`, `setType`, `setFunction` and `setMethod`, each taking
literal Go code. Any of the files below could be expressed as one record per
file with those actions today. The "repeated shape" column marks where a
dedicated action would buy more than literal code, because the same pattern
recurs and a template could enforce it.

**Regrowing must keep the proofs green.** For each package, the gate is its
existing tests, unchanged: `conformance/hashing/cases.json`, the
`TestParityWithLinespec` and `TestOracle` comparisons, and the real-repo checks.

## record

| File | Shape | Candidate actions | Repeated shape worth an action |
| --- | --- | --- | --- |
| `types.go` | 5 types (Status, Type, Record, AssociatedSpec, KeyNode), status/type constants, Validate | setType, setConst, setMethod | **enum type**: a string type, its constants and a `Valid(schema)` method recur for Status and Type |
| `decode.go` | YAML node walk into the hash input model; per-field kind table; string/list/int/bool/spec decoders | setFunction, setVar (knownFields table) | **field table**: knownFields, stringField and listField are three views of one field list; one table-driven action could emit all three |
| `codec.go` | byte-range splicing save, key orders, insertion anchors, Marshal | setFunction, setVar (v1Order, v2Order, insertAfter) | the key-order tables derive from the same field list as decode.go |
| `hash.go` | normalize, White_Space set, canonical JSON writer, GenHash/ContentHash/SpecHash | setFunction, setConst | no; HASHING.md and the corpus are the spec, so keep it as literal code |
| `snapshot.go` | Snapshot type, NewSnapshot, Verify, (un)marshal | setType, setFunction, setMethod | no |
| `store.go` | Store interface, Commit, FilesStore (List/Load/Save/Remove/Versions/Snapshot/Append) | setType, setFunction, setMethod | **store backend**: a Provisio API backend will implement the same 8-method interface; a store-backend action with per-method bodies fits that |
| `manifest.go` | ManifestLine, Build/Marshal/ParseManifest | setType, setFunction | no |

## migrate

| File | Shape | Candidate actions | Repeated shape worth an action |
| --- | --- | --- | --- |
| `load.go` | Config, Source, Repo, Load, loadDir, shared-repo resolution | setType, setFunction, setMethod | no |
| `git.go` | Git interface, GitCLI, gitCommand (GIT_* strip) | setType, setMethod, setFunction | **safe git exec**: the GIT_* strip, `-C` and toplevel check now recur in migrate, lint tests and the CLI; one action would make the COMMON.md git rule structural |
| `plan.go` | rules 1-4 and C3, survivor walk, version order (heap), keep-separate suggestions, status warnings | setFunction, setMethod, setType | migration rules as data (rule, from-type, to-type, direction) could replace classifyEdge's switch, but the oracle parity matters more than the shape |
| `apply.go` | version construction, pins, kept-bug imprints, v2 status mapping | setFunction, setMethod | no |

## scope

| File | Shape | Candidate actions | Repeated shape worth an action |
| --- | --- | --- | --- |
| `scope.go` | Match, GlobToRegex, InScope, Excluded, Overlap (a verbatim port of linespec) | setFunction | no; it must stay byte-for-byte equivalent to linespec until linespec drops its copy |

## lint

| File | Shape | Candidate actions | Repeated shape worth an action |
| --- | --- | --- | --- |
| `lint.go` | ~30 rule methods on a linter, rule id constants, Issue/Result/Options, the manifest check | setMethod, setConst, setType | **lint rule**: every rule is "id, field, severity, message template, predicate". A `lintRule` action taking those as kwargs plus a predicate body would generate the 34 methods and the rule constants together, and keep messages identical to linespec's |
| `commit.go` | ParseCommitTags, ResolveTag, Transitions table, CheckTransition | setFunction, setVar (Transitions) | **transition table**: Transitions is data; the CLI and the UI also need it |

## stale

| File | Shape | Candidate actions | Repeated shape worth an action |
| --- | --- | --- | --- |
| `stale.go` | Find (direct/transitive/unknown), Pin, PinLog, FilePinLog, Repin | setType, setFunction, setMethod | **append-only JSONL log**: FilePinLog repeats the pattern of the evidence store and the version manifest (append or rewrite lines, strict parse) |

## evidence

| File | Shape | Candidate actions | Repeated shape worth an action |
| --- | --- | --- | --- |
| `evidence.go` | Region, CodeHash, Evidence, Current, Check, Store (append-only files) | setType, setFunction, setMethod | shares the append-only file-store shape with stale and snapshots |

## ts (TypeScript)

| File | Shape | Candidate actions | Repeated shape worth an action |
| --- | --- | --- | --- |
| `src/hashing.ts` | normalize, canonical, gen/content/spec/code hashes | none yet: no TypeScript generator package exists | a TypeScript generator package (setFunction/setType equivalents) would be the prerequisite |
| `src/sha256.ts` | pure SHA-256 | none | no; a fixed algorithm, keep it literal |

## Suggested order

1. **lint rule action.** It has the most repetition (34 methods), and parity is a strict gate.
2. **safe git exec action.** It turns the COMMON.md git rule from a convention into structure.
3. **field table action** for record (decode, codec key orders).
4. **store backend action**, when the Provisio API backend is built.
