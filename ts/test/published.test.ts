// Runs the version_hash cases of conformance/hashing/cases.json, read in
// place, so Go and TypeScript agree on every published version.
import { test } from "node:test";
import assert from "node:assert/strict";
import { readFileSync, readdirSync } from "node:fs";
import {
  VersionHashError, verifyVersion, versionCanonical, versionContentHash,
} from "../src/published.ts";
import type { PublishedVersion } from "../src/published.ts";

interface Case {
  id: string;
  kind: string;
  version: PublishedVersion;
  canonical?: string;
  hash?: string;
  valid?: boolean;
  error?: boolean;
}

const corpus = JSON.parse(readFileSync(new URL("../../conformance/hashing/cases.json", import.meta.url), "utf8"));
const cases = (corpus.cases as Case[]).filter((c) => c.kind === "version_hash");

test("the corpus has version_hash cases", () => assert.ok(cases.length > 0));

for (const c of cases) {
  test(c.id, () => {
    if (c.error) {
      assert.throws(() => versionContentHash(c.version), VersionHashError);
      return;
    }
    assert.equal(versionCanonical(c.version), c.canonical);
    assert.equal(versionContentHash(c.version), c.hash);
    if (c.valid) verifyVersion(c.version);
    else assert.throws(() => verifyVersion(c.version), VersionHashError);
  });
}

// The sample versions the Go mapping writes verify here too.
const samples = new URL("../../lexicons/samples/", import.meta.url);
for (const name of readdirSync(samples).filter((n) => n.startsWith("version-"))) {
  test(`sample ${name}`, () => {
    verifyVersion(JSON.parse(readFileSync(new URL(name, samples), "utf8")));
  });
}
