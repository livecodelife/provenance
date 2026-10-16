// Proof for the integer extension of canonical() (HASHING.md section 5):
// bigint, and number only when a safe integer, are written as bare digits in
// the uint64/int64 range; floats and unsafe numbers throw.
import { test } from "node:test";
import assert from "node:assert/strict";
import { readFileSync } from "node:fs";
import { canonical } from "../src/hashing.ts";

const corpus = JSON.parse(readFileSync(new URL("../../conformance/hashing/cases.json", import.meta.url), "utf8"));
const cases = corpus.cases.filter((c: { kind: string }) => c.kind === "canonical_int");

// {"$int": text} carries an integer as text; everything else is itself.
function fromCorpus(v: unknown): unknown {
  if (Array.isArray(v)) return v.map(fromCorpus);
  if (v !== null && typeof v === "object") {
    const o = v as Record<string, unknown>;
    const keys = Object.keys(o);
    if (keys.length === 1 && keys[0] === "$int") {
      const text = o["$int"] as string;
      if (!/^[+-]?[0-9]+$/.test(text)) throw new SyntaxError(`${text} is not an integer`);
      return BigInt(text);
    }
    return Object.fromEntries(keys.map((k) => [k, fromCorpus(o[k])]));
  }
  return v;
}

test("corpus has the canonical_int cases", () => {
  assert.ok(corpus.corpus_version >= 5, `corpus_version ${corpus.corpus_version}`);
  assert.ok(cases.length >= 8, `${cases.length} canonical_int cases`);
});

for (const c of cases) {
  test(c.id, () => {
    if (c.error) {
      assert.throws(() => canonical(fromCorpus(c.value) as never));
    } else {
      assert.equal(canonical(fromCorpus(c.value) as never), c.canonical);
    }
  });
}

test("integers", () => {
  assert.equal(canonical(0), "0");
  assert.equal(canonical(-7), "-7");
  assert.equal(canonical(-0), "0");
  assert.equal(canonical(18446744073709551615n), "18446744073709551615");
  assert.equal(canonical(-9223372036854775808n), "-9223372036854775808");
  assert.equal(canonical({ n: 3, ok: true }), '{"n":3,"ok":true}');
});

test("rejections", () => {
  assert.throws(() => canonical(1.5));
  assert.throws(() => canonical(NaN));
  assert.throws(() => canonical(2 ** 53));
  assert.throws(() => canonical(18446744073709551616n));
  assert.throws(() => canonical(-9223372036854775809n));
});
