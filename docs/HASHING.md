# Canonical record hashing

Status: normative. Decided by prov-2026-20c3360e; §7 by prov-2026-7752337c.

A provenance record has two hashes, both computed and never stored in the
record file:

| Hash | Covers | Pinned by |
| --- | --- | --- |
| `gen_hash` | exactly what sedum reads: `intent`, `constraints`, `affected_scope`, `forbidden_scope` | sedum markers (`gen`) |
| `content_hash` | the behavioral fields (`gen_hash`'s four) plus the content of each associated spec (§4) | child records, linespecs, copies, version snapshots |

Every implementation (the Go record library, sedum, the Provisio TypeScript
backend) must produce byte-identical canonical encodings, and so identical
hashes. [`conformance/hashing/cases.json`](../conformance/hashing/cases.json)
is the check. When this document and the corpus disagree, one of them has a
bug: fix it and bump `corpus_version`. Never change behaviour silently.

The pipeline has four stages, and each one is specified separately:

1. **Decode** the YAML file into the *hash input model* (§1).
2. **Normalize** every string (§2).
3. **Project** the fields a hash covers into a fixed-shape object (§3, §4).
4. **Encode** that object as canonical JSON and hash it (§5).

## 1. Decoding: YAML to the hash input model

The hash input model has only three kinds of value: strings, lists of
strings, and (for associated specs only) null. No number, boolean, date or
mapping value ever reaches a hash. Records that arrive as JSON (for example
from AT Proto) skip YAML and enter at this model directly.

Rules for a YAML reader:

- **Top-level keys** match exactly and are case-sensitive. `affectedScope` is
  an unknown key, not an alias. A duplicate top-level key is a decode error.
- **Unknown keys** are preserved by the codec and ignored by both hashes.
  Adding a field to either hash is a spec change with a `corpus_version` bump.
- **String fields** (`intent`, and every other scalar text field such as `title`, `type`, `implements`): a scalar's value
  is its YAML *source text after scalar resolution of quoting and block
  style*, never a typed value:
  - plain `1.50` is `"1.50"` (not `1.5`), plain `true` is `"true"`, plain
    `2026-10-15` is `"2026-10-15"`;
  - quoted scalars (`'…'`, `"…"`) use their unescaped content;
  - block scalars (`|`, `>`, with any chomping or indentation indicator) use
    the content YAML defines for them. Folding and chomping differences then
    vanish in normalization (§2), so `|`, `>` and a one-line value with the
    same words hash the same;
  - plain `null`, `Null`, `NULL`, `~`, or an empty value are **null**.
  - A sequence or mapping where a string is expected is a decode error.
- **List fields** (`constraints`, `affected_scope`, `forbidden_scope`): a YAML sequence. Each item follows the string-field
  rules, with null allowed. A null or empty value in place of the whole list is
  null. A scalar or mapping where a list is expected is a decode error.
- **`associated_specs`**: a sequence of mappings. Only `path` matters to the
  hash (§4). Other keys (`type`, `run_command`, `description`, …) are
  preserved and never hashed.
- Anchors and aliases resolve as YAML defines them. Comments never matter.
- Input must be valid UTF-8. Invalid UTF-8 is a decode error. (A Go caller
  that bypasses the decoder and passes invalid UTF-8 straight to `GenHash`
  gets each invalid byte replaced by U+FFFD, Go's range semantics. This is a
  defensive fallback that the corpus does not cover.)

### Missing, null and empty are one value

| Field kind | missing | `null` / `~` / empty | `""` / `[]` | becomes |
| --- | --- | --- | --- | --- |
| string | ✓ | ✓ | ✓ | `""` |
| list | ✓ | ✓ | ✓ | `[]` |
| list item | — | ✓ | ✓ | `""` (kept in place) |

The projected object always contains every key in the hash's fixed set (§3,
§4). Leaving a key out is never a way to say "empty".

## 2. String normalization

Apply to every string (field values and list items, never keys or spec file
bytes), in this order:

1. **NFC.** Unicode Normalization Form C. Implementations should use Unicode
   15.1 or later; NFC is stable for assigned code points, so version skew can
   only affect code points unassigned in the older version.
2. **Collapse whitespace.** Replace each maximal run of whitespace code points
   with a single U+0020 SPACE.
3. **Trim.** Remove a leading and a trailing U+0020, if present.

"Whitespace" is exactly the Unicode `White_Space` property, listed here so no
implementation depends on its regex engine's idea of `\s`:

```
U+0009..U+000D  U+0020  U+0085  U+00A0  U+1680  U+2000..U+200A
U+2028  U+2029  U+202F  U+205F  U+3000
```

Not whitespace (left untouched): U+200B ZERO WIDTH SPACE, U+180E, U+FEFF BOM,
U+2060. JavaScript's `\s` includes U+FEFF and Go's `unicode.IsSpace` matches
this list exactly. Use the explicit list in both.

Consequences that are intended:

- Line breaks inside `intent` count as spaces. Re-wrapping a paragraph or
  switching between `|` and `>` changes neither hash.
- Markdown structure that lives only in whitespace (indentation, blank lines
  between paragraphs) does not count. Characters such as `-` or `*` do.
- Case, punctuation and quoting style inside a string do count.

## 3. `gen_hash`

The projected object has exactly these four keys:

| Key | Kind |
| --- | --- |
| `affected_scope` | list |
| `constraints` | list |
| `forbidden_scope` | list |
| `intent` | string |

List order is kept. Items are never sorted, deduplicated or dropped, even when
empty or repeated, because sedum reads the lists in order.

Go API (package `github.com/livecodelife/provenance/record`):

```go
type GenInput struct {
	Intent                                   string
	Constraints, AffectedScope, ForbiddenScope []string
}

func GenHash(in GenInput) string // "sha256:<64 lowercase hex>"
func (r *Record) GenHash() string
```

A nil slice and an empty slice are the same value.

## 4. `content_hash`

The rule (Caleb, #provisio; field set applied by provisio-mgr, 58864):
**only fields that change what the record makes the software do are content.** Labels, cross-references, bookkeeping and
lifecycle are not, so editing them never forces a new version.

The projected object has exactly these five keys. This table is the single
definition of the set; the Go function `contentFields` mirrors it.

| Key | Kind | Notes |
| --- | --- | --- |
| `affected_scope` | list | |
| `associated_specs` | list of spec hash or null | see below |
| `constraints` | list | |
| `forbidden_scope` | list | |
| `intent` | string | |

**Left out, and why:**

| Field | Why |
| --- | --- |
| `status`, `sealed_at_sha`, `version`, `reason` | lifecycle: a status change must not look like a content change (DESIGN.md) |
| `implements_version` | `provenance repin` moves only the pin and must not bump a version |
| `implements`, `implements_repo`, `implements_uri` | the parent link is lineage, not behaviour; staleness is tracked through the pin |
| `type` | the tier classifies the record; what it makes the software do is in intent, constraints and scope |
| `title`, `tags`, `related` | labels and cross-references, which change no behaviour |
| `id` | identity, not content. The hash says which version of an id. |
| `author`, `created_at` | metadata that the version snapshot envelope carries |
| `aliases` | migration metadata: folding an old id in is not a content change |
| `locked` | governance flag |
| `monitors`, `associated_traces` | operational attachments, mutable after implementation |
| `supersedes`, `superseded_by`, `extends` | v1-only, retired by versioning |
| unknown keys | preserved by the codec, never hashed (§1) |

**Associated specs.** Each entry of `associated_specs`, in list order,
contributes the content hash of the spec file it names:

- *spec hash* = `"sha256:" + lowercase hex(SHA-256(bytes))`, where `bytes` is
  the file content with every CRLF (`\r\n`) pair replaced by LF (`\n`).
  Nothing else changes: no NFC, no trimming, and a lone CR or a BOM stays.
- `path` is resolved relative to the repository root (the directory that holds
  the provenance config). It is taken literally: no globbing.
- An entry whose file cannot be read, or which has no `path`, contributes
  `null`. Implementations report it as a warning, not an error, so one
  missing proof file does not make a record unhashable. Creating the file
  later changes the hash, as it should.
- *Record files.* An entry whose `path` has a base name (the part after the
  last `/`) of the form `prov-YYYY-<hex8>.yml` or `prov-YYYY-<hex8>-<suffix>.yml`
  (regex `^prov-[0-9]{4}-[0-9a-f]{8}(-[A-Za-z0-9._-]+)?\.yml$`) names a record
  file. It contributes the constant string `record-file/v1` in its list slot,
  whatever the directory, whether or not the file exists, and it is never
  reported missing. The marker has no `sha256:` prefix, so it cannot equal a
  spec hash. Reason: migrate, revise and complete rewrite record files, and a
  record that names its own file would have to hash itself; a record file's
  identity is already its own `content_hash`. Any other base name is hashed by
  its bytes as above (`prov-2026-xyz.yml` and `my-prov-2026-0000b001.yml` are
  ordinary specs).
- `path`, `type`, `run_command` and any other entry keys are not hashed. Paths
  mean nothing once a record is published or copied, and `run_command` is
  local-only (DESIGN.md, Associated specs). Renaming a spec file without
  changing it does not change the hash. Reordering specs does.

Go API:

```go
type SpecResolver func(path string) (hash string, ok bool)

func FileSpecResolver(root string) SpecResolver
func SpecHash(content []byte) string
func (r *Record) ContentHash(specs SpecResolver) (hash string, missing []string)
```

## 5. Canonical JSON and the hash

Serialize the projected object as JSON with:

- object keys sorted by their UTF-8 bytes (all keys here are lowercase ASCII,
  so this is plain ASCII order);
- no whitespace anywhere outside strings;
- strings in double quotes, escaping **only** `"` as `\"`, `\` as `\\`, and
  each code point U+0000–U+001F as `\u00XX` with lowercase hex. Every other
  code point, including non-ASCII, `/`, `<`, `>`, `&`, U+2028 and U+2029, is
  written as raw UTF-8. (After §2 no U+0009–U+000D, U+2028 or U+2029 can
  remain, but the rule is stated for completeness.)
- `null` only inside `associated_specs`.
- booleans (for objects other than record projections, such as sedum's pkg
  hash) as the bare literals `true` and `false`. No other number type is
  defined: an encoder given a number must fail, never write one. No record
  projection (§3, §4, §7) contains a boolean, so every `gen_hash`,
  `content_hash` and `code_hash` is unchanged by this rule.
- integers (corpus version 5; again only for objects other than record
  projections) as bare base-10 digits: no plus sign, no leading zeros, zero as
  `0` (never `-0`), negatives as `-N`, from -9223372036854775808 to
  18446744073709551615. A float, or an integer outside that range, must make
  the encoder fail. Go accepts `int`, `int64` and `uint64`; TypeScript accepts
  `bigint`, and `number` only when `Number.isSafeInteger`.

Then: `hash = "sha256:" + lowercase hex(SHA-256(UTF-8 bytes of the JSON))`.

Note on stock encoders: Go's `encoding/json` escapes `<`, `>`, `&`, U+2028 and
U+2029 by default, and JavaScript's `JSON.stringify` writes control characters
as `\b`, `\f`, `\n`, `\r`, `\t`. Both disagree with this section on some
inputs, so use a small purpose-written encoder or check against the corpus.

### Example

```yaml
intent: |
  Users can
  sign in.
constraints: [ "No  passwords " ]
affected_scope:
forbidden_scope: ~
```

projects to

```json
{"affected_scope":[],"constraints":["No passwords"],"forbidden_scope":[],"intent":"Users can sign in."}
```

and its `gen_hash` is the SHA-256 of those bytes (see corpus case
`gen/example-from-spec`).

## 7. `code_hash` (verification evidence)

Verification evidence can record which code a spec run exercised, without
git: a hash over the `gen` and `pkg` values of the record's generated
regions (DESIGN.md, Associated specs). It is the code hash a tool that tracks
generated regions may record; provenance itself does not compute or store it.
A caller passes one `{gen, pkg}` pair per region.

0. Only owned regions count. A region whose `tier` is present and not
   `"owned"` (for example `"seeded"`) is left out, because the generator never
   regenerates it. A region with no `tier` counts, so a
   caller may filter beforehand.
1. Each region projects to `{"gen": <gen>, "pkg": <pkg>}`. The values are
   hash strings, taken exactly as given: no normalization.
   **A region with an empty or missing `gen` or `pkg` has no code hash.**
   Implementations refuse to compute one, and evidence then reports the code
   as unknown, never fresh. This matches "absent = stale" on markers; stamping
   the region is the way out.
2. Encode each region with §5's canonical JSON, then sort the regions by
   those bytes. Where regions sit in files must not matter. Duplicates are
   kept, because two regions are two regions.
3. Encode the sorted list as a canonical JSON array (no regions is `[]`) and
   hash it as in §5.

Go API: `evidence.CodeHash(regions []evidence.Region) (string, error)`;
the error is `evidence.ErrUnstamped`.

## 6. The conformance corpus

`conformance/hashing/cases.json` holds one array of cases. Every case has `id`,
`kind`, `decision` (the record that decided it) and `why`. By kind:

- **`normalize`**: `input` (string) and `output` (string).
- **`gen_hash`**: `record` (a JSON object: the hashed keys follow the hash
  input model, may be null or missing, and any other key may hold any JSON
  value and must be ignored), `canonical` (the exact JSON
  string from §5) and `hash`.
- **`content_hash`**: `record`, optional `spec_files` (a map from path to file
  content; a path not in the map is unreadable), `canonical` and `hash`.
- **`spec_hash`**: `content` (string) and `hash`.
- **`code_hash`**: `regions` (a list of `{gen, pkg, tier?}` objects), then either
  `canonical` and `hash`, or `error: true` when a region lacks gen or pkg.
- **`canonical`** (corpus version 4): `value` (any JSON value built from
  strings, booleans, null, lists and maps) and either `canonical` (the §5
  encoding of it) and `hash` (of those bytes), or `error: true` when the value
  contains a number, which an encoder must reject.
- **`canonical_int`** (corpus version 5): like `canonical`, except that an
  integer is carried as the object `{"$int": "<text>"}` so no language rounds
  a 64-bit value through a float. The text is an optional sign and digits;
  leading zeros and a plus sign are accepted on input and normalised away
  (`-0` becomes `0`). `canonical` and `hash` give the encoding, or `error: true`
  when the text is not an integer or lies outside the range in §5. Plain JSON
  numbers in `canonical` cases are floats there (the corpus has no safe way to
  carry larger integers) and an error.
- **`version_hash`** (corpus version 7): `version`, a published version record
  (docs/LEXICONS.md) of which only `content`, `specHashes` and `contentHash`
  matter. Hash it from `content` and `specHashes` alone: the five content keys
  from `content` (`affectedScope` and `forbiddenScope` are `affected_scope` and
  `forbidden_scope`), and each `specHashes` entry in its list slot, with the
  empty string as null. `canonical` and `hash` give the result; `valid` says
  whether the version's own `contentHash` equals it. Or `error: true` when
  `specHashes` and `content.associatedSpecs` differ in length.
- **`decode`**: `yaml` (a whole record file) and either `fields` (the decoded
  hash input model for the keys listed, with null kept as null) or
  `error: true` with `error_names` (a substring the error must contain).

Compare `canonical` and `hash` as bytes. There is no runner here on purpose:
load the JSON, drive your own implementation, and compare. The Go tests in
`record/` read this file directly.
