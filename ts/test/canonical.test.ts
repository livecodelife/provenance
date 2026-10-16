// Proof for prov-2026-32ee2643: canonical() encodes booleans as the bare
// literals true and false, and the corpus's canonical cases pass.
import { test } from "node:test";
import assert from "node:assert/strict";
import { readFileSync } from "node:fs";
import { canonical } from "../src/hashing.ts";

const corpus = JSON.parse(readFileSync(new URL("../../conformance/hashing/cases.json", import.meta.url), "utf8"));
const cases = corpus.cases.filter((c: { kind: string }) => c.kind === "canonical");

test("corpus has the canonical cases", () => {
  assert.ok(corpus.corpus_version >= 4, `corpus_version ${corpus.corpus_version}`);
  assert.ok(cases.length >= 7, `${cases.length} canonical cases`);
});

for (const c of cases) {
  test(c.id, () => {
    if (c.error) {
      assert.throws(() => canonical(c.value));
    } else {
      assert.equal(canonical(c.value), c.canonical);
    }
  });
}

test("boolean literals", () => {
  assert.equal(canonical(true), "true");
  assert.equal(canonical(false), "false");
  assert.equal(canonical({ b: true, a: [false, "true", null] }), '{"a":[false,"true",null],"b":true}');
});
