// Round trip of the sample project (lexicons/samples/project.json) through
// the generated type and validator, and the limits the lexicon states.
import { test } from "node:test";
import assert from "node:assert/strict";
import { readFileSync } from "node:fs";
import * as Project from "../src/types/dev/provisio/project.js";

const dir = new URL("../../../../samples/", import.meta.url);
const load = (name: string): Record<string, unknown> => JSON.parse(readFileSync(new URL(name, dir), "utf8"));

test("the sample project validates and survives a round trip", () => {
  const raw = load("project.json");
  const res = Project.validateMain({ $type: "dev.provisio.project", ...raw });
  if (!res.success) assert.fail(res.error.message);
  const back = JSON.parse(JSON.stringify(res.value));
  delete back.$type;
  assert.deepEqual(back, raw);
});

test("a project without a description validates", () => {
  const raw = load("project.json");
  delete raw.description;
  assert.ok(Project.validateMain({ $type: "dev.provisio.project", ...raw }).success);
});

test("a project without a name is rejected", () => {
  const raw = load("project.json");
  delete raw.name;
  assert.ok(!Project.validateMain({ $type: "dev.provisio.project", ...raw }).success);
});

test("a project without createdAt is rejected", () => {
  const raw = load("project.json");
  delete raw.createdAt;
  assert.ok(!Project.validateMain({ $type: "dev.provisio.project", ...raw }).success);
});

test("a name over 64 graphemes is rejected and 64 is accepted", () => {
  const raw = load("project.json");
  assert.ok(Project.validateMain({ $type: "dev.provisio.project", ...raw, name: "a".repeat(64) }).success);
  assert.ok(!Project.validateMain({ $type: "dev.provisio.project", ...raw, name: "a".repeat(65) }).success);
});

test("a description over 300 graphemes is rejected and 300 is accepted", () => {
  const raw = load("project.json");
  assert.ok(Project.validateMain({ $type: "dev.provisio.project", ...raw, description: "a".repeat(300) }).success);
  assert.ok(!Project.validateMain({ $type: "dev.provisio.project", ...raw, description: "a".repeat(301) }).success);
});
