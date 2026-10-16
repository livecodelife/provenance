# @livecodelife/provenance-hashing

TypeScript implementation of [docs/HASHING.md](../docs/HASHING.md): the
canonical encoding behind `gen_hash`, `content_hash`, spec hashes and
`code_hash`. It mirrors the Go `record` and `evidence` packages, and both pass
[`conformance/hashing/cases.json`](../conformance/hashing/cases.json), read in
place.

- No dependencies. NFC is `String.prototype.normalize`, and SHA-256 is a small
  synchronous implementation (`src/sha256.ts`), checked against the NIST
  vectors and `node:crypto`. So it runs unchanged in Node, Deno, browsers and
  workers.
- Input is the JSON hash input model with the record file's snake_case keys,
  which is how Provisio receives records. YAML decoding is not included, so
  the corpus's `decode` cases are an explicit skip.
- Erasable TypeScript: Node 23.6+ runs it directly.

```sh
cd ts && node --test
```

```ts
import { genHash, contentHash, specHash, codeHash } from "./src/hashing.ts";

genHash({ intent: "Users can sign in.", constraints: ["No passwords"] });
contentHash(record, (path) => specHash(readSpec(path))); // null when unreadable
codeHash(regions); // throws UnstampedError on an owned region without gen/pkg
```
