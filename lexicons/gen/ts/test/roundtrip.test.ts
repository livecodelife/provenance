// Round trip of the sample head and three versions through the generated
// types and validators (lexicons/samples, written by the Go mapping in
// record/published.go and checked against it by record/published_test.go).
import { test } from "node:test";
import assert from "node:assert/strict";
import { readFileSync } from "node:fs";
import * as Version from "../src/types/dev/provisio/provenance/version.js";
import * as Head from "../src/types/dev/provisio/provenance/record.js";

const dir = new URL("../../../../samples/", import.meta.url);
const load = (name: string): Record<string, unknown> => JSON.parse(readFileSync(new URL(name, dir), "utf8"));

for (const name of ["version-1.json", "version-2.json", "version-3.json"]) {
  test(`${name} validates and survives a round trip`, () => {
    const raw = load(name);
    const res = Version.validateMain({ $type: "dev.provisio.provenance.version", ...raw });
    if (!res.success) assert.fail(res.error.message);
    const typed: Version.Main = res.value;
    const back = JSON.parse(JSON.stringify(typed));
    delete back.$type;
    assert.deepEqual(back, raw);
    assert.equal(typed.specHashes.length, typed.content.associatedSpecs?.length ?? 0);
  });
}

test("head validates and survives a round trip", () => {
  const raw = load("head.json");
  const res = Head.validateMain({ $type: "dev.provisio.provenance.record", ...raw });
  if (!res.success) assert.fail(res.error.message);
  const back = JSON.parse(JSON.stringify(res.value));
  delete back.$type;
  assert.deepEqual(back, raw);
  // The head's content carries the repinned parent version; the newest version does not.
  const v3 = load("version-3.json") as unknown as Version.Main;
  assert.notEqual((raw.content as { implementsVersion: number }).implementsVersion, v3.content.implementsVersion);
});

test("a version without its content hash is rejected", () => {
  const raw = load("version-1.json");
  delete raw.contentHash;
  assert.ok(!Version.validateMain({ $type: "dev.provisio.provenance.version", ...raw }).success);
});

test("a head with a bad status field type is rejected", () => {
  const raw = load("head.json");
  raw.version = "3";
  assert.ok(!Head.validateMain({ $type: "dev.provisio.provenance.record", ...raw }).success);
});

test("a head without a project still validates", () => {
  const raw = load("head.json");
  delete raw.project;
  assert.ok(Head.validateMain({ $type: "dev.provisio.provenance.record", ...raw }).success);
});
