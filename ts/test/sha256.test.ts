// The hand-rolled SHA-256 is this package's one risk, so it is checked
// against the NIST vectors (FIPS 180-4 examples) and node:crypto.
import { test } from "node:test";
import assert from "node:assert/strict";
import { createHash, randomBytes } from "node:crypto";
import { sha256Hex } from "../src/sha256.ts";

const enc = (s: string) => new TextEncoder().encode(s);

test("NIST vectors", () => {
  assert.equal(sha256Hex(enc("")), "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855");
  assert.equal(sha256Hex(enc("abc")), "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad");
  assert.equal(
    sha256Hex(enc("abcdbcdecdefdefgefghfghighijhijkijkljklmklmnlmnomnopnopq")),
    "248d6a61d20638b8e5c026930c3e6039a33ce45964ff2167f6ecedd419db06c1",
  );
  assert.equal(
    sha256Hex(enc("abcdefghbcdefghicdefghijdefghijkefghijklfghijklmghijklmnhijklmnoijklmnopjklmnopqklmnopqrlmnopqrsmnopqrstnopqrstu")),
    "cf5b16a778af8380036ce59e7b0492370b249b11e8f07a51afac45037afee9d1",
  );
  assert.equal(sha256Hex(new Uint8Array(1_000_000).fill(0x61)), "cdc76e5c9914fb9281a1c7e284d73e67f1809a48a497200e046d39ccc7112cd0");
});

test("matches node:crypto on random inputs of every length 0-1000", () => {
  for (let n = 0; n <= 1000; n++) {
    const data = randomBytes(n);
    const want = createHash("sha256").update(data).digest("hex");
    assert.equal(sha256Hex(new Uint8Array(data)), want, `length ${n}`);
  }
});
