// Proof for prov-2026-4234f0bf: a spec whose path is a record file contributes
// the constant marker record-file/v1 to content_hash (HASHING.md section 4).
import { test } from "node:test";
import assert from "node:assert/strict";
import { contentHash } from "../src/hashing.ts";

function hashOf(path: string, files: Record<string, string>): string {
  const resolve = (p: string) => (p in files ? `file:${files[p]}` : null);
  return contentHash({ intent: "Does the thing.", associated_specs: [{ path }] }, resolve as never);
}

const REC = "provenance/prov-2026-0000b001.yml";

test("record-file spec ignores the file bytes and existence", () => {
  const a = hashOf(REC, { [REC]: "version: 1\n" });
  assert.equal(a, hashOf(REC, { [REC]: "version: 2\n" }));
  assert.equal(a, hashOf(REC, {}));
});

test("record-file spec in any directory or with a suffix hashes the same", () => {
  const want = hashOf(REC, {});
  for (const p of ["prov-2026-0000b001.yml", "x/y/prov-2026-0000b001.yml", "provenance/prov-2026-0000b001-api.yml"]) {
    assert.equal(hashOf(p, {}), want, p);
  }
});

test("lookalike names are ordinary specs", () => {
  for (const p of ["specs/prov-2026-0000b001.txt", "specs/prov-2026-xyz.yml", "specs/my-prov-2026-0000b001.yml"]) {
    assert.notEqual(hashOf(p, { [p]: "a" }), hashOf(p, { [p]: "b" }), p);
  }
});
