// Runs conformance/hashing/cases.json, read in place and never copied, so
// the Go and TypeScript implementations are checked against the same file.
import { test } from "node:test";
import assert from "node:assert/strict";
import { readFileSync } from "node:fs";
import {
  canonical, codeCanonical, codeHash, contentCanonical, contentHash, genCanonical, genHash,
  normalize, specHash, specHashes, UnstampedError,
} from "../src/hashing.ts";
import type { HashInput, Region } from "../src/hashing.ts";

interface Case {
  id: string;
  kind: string;
  input?: string;
  output?: string;
  record?: HashInput;
  spec_files?: Record<string, string>;
  canonical?: string;
  hash?: string;
  content?: string;
  regions?: Region[];
  error?: boolean;
}

const corpus = JSON.parse(readFileSync(new URL("../../conformance/hashing/cases.json", import.meta.url), "utf8"));
assert.equal(corpus.format, "provenance.hashing");

// decode cases pin YAML scalar resolution. This package has no YAML parser
// (zero dependencies); Provisio receives records as JSON, which enters at the
// hash input model after decode (HASHING.md §1).
const SKIPPED_KINDS: Record<string, string> = {
  decode: "YAML decoding is out of scope for the zero-dependency TypeScript package",
};

for (const c of corpus.cases as Case[]) {
  const reason = SKIPPED_KINDS[c.kind];
  test(c.id, reason ? { skip: reason } : {}, () => {
    switch (c.kind) {
      case "normalize":
        assert.equal(normalize(c.input!), c.output);
        break;
      case "gen_hash":
        assert.equal(genCanonical(c.record!), c.canonical);
        assert.equal(genHash(c.record!), c.hash);
        break;
      case "content_hash": {
        const files = c.spec_files ?? {};
        const resolve = (p: string) => (p in files ? specHash(files[p]) : null);
        assert.equal(contentCanonical(c.record!, specHashes(c.record!, resolve)), c.canonical);
        assert.equal(contentHash(c.record!, resolve), c.hash);
        break;
      }
      case "spec_hash":
        assert.equal(specHash(c.content!), c.hash);
        break;
      case "code_hash":
        if (c.error) {
          assert.throws(() => codeHash(c.regions!), UnstampedError);
        } else {
          assert.equal(codeCanonical(c.regions!), c.canonical);
          assert.equal(codeHash(c.regions!), c.hash);
        }
        break;
      default:
        assert.fail(`unknown corpus kind ${c.kind}`);
    }
  });
}

test("every kind is run or explicitly skipped", () => {
  const known = new Set(["normalize", "gen_hash", "content_hash", "spec_hash", "code_hash", ...Object.keys(SKIPPED_KINDS)]);
  for (const c of corpus.cases as Case[]) assert.ok(known.has(c.kind), `${c.id}: unhandled kind ${c.kind}`);
});

test("canonical rejects nothing it should encode", () => {
  assert.equal(canonical({ b: ["x", null], a: "\u0008é" }), '{"a":"\\u0008é","b":["x",null]}');
});
