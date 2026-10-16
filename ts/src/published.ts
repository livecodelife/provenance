// The published form of a provenance version (docs/LEXICONS.md): the lexicon
// dev.provisio.provenance.version in camelCase. Mirrors record/published.go.
// Only the fields the content hash covers are typed here; the generated
// lexicon types describe the rest.

import { contentCanonical, contentHash } from "./hashing.ts";
import type { HashInput } from "./hashing.ts";

export interface PublishedSpec {
  path: string;
  type?: string;
}

export interface PublishedContent {
  title: string;
  intent: string;
  constraints?: string[];
  affectedScope?: string[];
  forbiddenScope?: string[];
  associatedSpecs?: PublishedSpec[];
  [other: string]: unknown;
}

export interface PublishedVersion {
  contentHash: string;
  specHashes: string[];
  content: PublishedContent;
  [other: string]: unknown;
}

/** Thrown when a version's content does not hash to the contentHash it states. */
export class VersionHashError extends Error {
  constructor(message: string) {
    super(message);
    this.name = "VersionHashError";
  }
}

/** The hash input model (snake_case) of a published content. */
export function hashInputOf(c: PublishedContent): HashInput {
  return {
    intent: c.intent,
    constraints: c.constraints ?? [],
    affected_scope: c.affectedScope ?? [],
    forbidden_scope: c.forbiddenScope ?? [],
    associated_specs: (c.associatedSpecs ?? []).map((s) => ({ path: s.path })),
  };
}

/** specHashes with "" (no hash) as null, one per associated spec. */
export function specHashesOf(v: PublishedVersion): (string | null)[] {
  const n = (v.content.associatedSpecs ?? []).length;
  if (v.specHashes.length !== n) {
    throw new VersionHashError(`${v.specHashes.length} spec hashes for ${n} specs`);
  }
  return v.specHashes.map((h) => (h === "" ? null : h));
}

/** The exact bytes the version's content hash covers. */
export function versionCanonical(v: PublishedVersion): string {
  return contentCanonical(hashInputOf(v.content), specHashesOf(v));
}

/** The content hash recomputed from content and specHashes. */
export function versionContentHash(v: PublishedVersion): string {
  return contentHash(hashInputOf(v.content), specHashesOf(v));
}

/** Throws unless the stated contentHash is the recomputed one. */
export function verifyVersion(v: PublishedVersion): void {
  const got = versionContentHash(v);
  if (got !== v.contentHash) {
    throw new VersionHashError(`content hashes to ${got}, the version says ${v.contentHash}`);
  }
}
