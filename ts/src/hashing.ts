// Canonical record hashing per docs/HASHING.md, mirroring the Go packages
// record (gen_hash, content_hash, spec_hash) and evidence (code_hash). Input
// is the JSON hash input model (§1): hashed keys hold strings, lists of
// strings (null items allowed) or null; every other key is ignored.

import { sha256Hex } from "./sha256.ts";

const encoder = new TextEncoder();

/** A record in the hash input model, with the file's snake_case keys. */
export interface HashInput {
  intent?: string | null;
  constraints?: (string | null)[] | null;
  affected_scope?: (string | null)[] | null;
  forbidden_scope?: (string | null)[] | null;
  associated_specs?: ({ path?: string | null } | null)[] | null;
  [other: string]: unknown;
}

/** A sedum region; only owned regions' gen and pkg feed the code hash. */
export interface Region {
  gen?: string;
  pkg?: string;
  tier?: string;
}

/** Spec hash for each associated spec, or null when unresolved. */
export type SpecResolver = (path: string) => string | null;

/** Thrown by codeHash when an owned region lacks gen or pkg (§7). */
export class UnstampedError extends Error {
  constructor() {
    super("a region has no gen or pkg; run sedum grow --stamp");
    this.name = "UnstampedError";
  }
}

// The Unicode White_Space property, listed explicitly (§2). JavaScript's \s
// also matches U+FEFF, which is not whitespace here.
const WHITE_SPACE = /[\u0009-\u000D\u0020\u0085\u00A0\u1680\u2000-\u200A\u2028\u2029\u202F\u205F\u3000]+/g;

/** §2: NFC, collapse each White_Space run to one space, trim. */
export function normalize(s: string): string {
  // Lone surrogates become U+FFFD, matching Go's handling of invalid UTF-8.
  return s.toWellFormed().normalize("NFC").replace(WHITE_SPACE, " ").replace(/^ | $/g, "");
}

type Canon = string | boolean | number | bigint | null | Canon[] | { [key: string]: Canon };

/** §5 canonical JSON: byte-sorted keys, no whitespace, minimal escapes. Booleans and
 * integers (bigint, or a safe-integer number) are written as bare literals;
 * floats and unsafe numbers throw. */
export function canonical(v: Canon): string {
  if (v === null) return "null";
  if (typeof v === "string") return quote(v);
  if (typeof v === "boolean") return v ? "true" : "false";
  if (typeof v === "bigint") return integer(v);
  if (typeof v === "number") {
    if (!Number.isSafeInteger(v)) throw new TypeError(`${v} is not a safe integer`);
    return integer(BigInt(v));
  }
  if (typeof v !== "object") throw new TypeError(`a ${typeof v} cannot be hashed`);
  if (Array.isArray(v)) return "[" + v.map(canonical).join(",") + "]";
  const keys = Object.keys(v).sort(compareBytes);
  return "{" + keys.map((k) => quote(k) + ":" + canonical(v[k])).join(",") + "}";
}

// §5: integers are bare base-10 digits within int64 min .. uint64 max; BigInt
// has no negative zero or leading zeros, so toString is already canonical.
function integer(n: bigint): string {
  if (n < -(2n ** 63n) || n > 2n ** 64n - 1n) throw new RangeError(`${n} is outside int64..uint64`);
  return n.toString();
}

function quote(s: string): string {
  let out = '"';
  for (const ch of s) {
    const c = ch.codePointAt(0)!;
    if (ch === '"') out += '\\"';
    else if (ch === "\\") out += "\\\\";
    else if (c < 0x20) out += "\\u00" + c.toString(16).padStart(2, "0");
    else out += ch;
  }
  return out + '"';
}

// UTF-8 byte order equals code point order, which is not UTF-16 order.
function compareBytes(a: string, b: string): number {
  const x = encoder.encode(a), y = encoder.encode(b);
  for (let i = 0; i < Math.min(x.length, y.length); i++) {
    if (x[i] !== y[i]) return x[i] - y[i];
  }
  return x.length - y.length;
}

function hashString(s: string): string {
  return "sha256:" + sha256Hex(encoder.encode(s));
}

function str(v: unknown, key: string): string {
  if (v === undefined || v === null) return "";
  if (typeof v !== "string") throw new TypeError(`${key} must be a string`);
  return normalize(v);
}

function list(v: unknown, key: string): string[] {
  if (v === undefined || v === null) return [];
  if (!Array.isArray(v)) throw new TypeError(`${key} must be a list`);
  return v.map((x) => (x === null ? "" : str(x, key)));
}

function genObject(r: HashInput): Canon {
  return {
    intent: str(r.intent, "intent"),
    constraints: list(r.constraints, "constraints"),
    affected_scope: list(r.affected_scope, "affected_scope"),
    forbidden_scope: list(r.forbidden_scope, "forbidden_scope"),
  };
}

/** The exact bytes gen_hash covers (§3). */
export function genCanonical(r: HashInput): string {
  return canonical(genObject(r));
}

/** gen_hash (§3): "sha256:<64 lowercase hex>". */
export function genHash(r: HashInput): string {
  return hashString(genCanonical(r));
}

/** spec hash (§4): SHA-256 of the file with CRLF pairs folded to LF. */
export function specHash(content: Uint8Array | string): string {
  const bytes = typeof content === "string" ? encoder.encode(content) : content;
  const out: number[] = [];
  for (let i = 0; i < bytes.length; i++) {
    if (bytes[i] === 0x0d && bytes[i + 1] === 0x0a) continue;
    out.push(bytes[i]);
  }
  return "sha256:" + sha256Hex(Uint8Array.from(out));
}

/** What a spec that is a record file contributes to content_hash instead of its bytes (§4). */
export const RECORD_FILE_MARKER = "record-file/v1";

const recordFileName = /^prov-[0-9]{4}-[0-9a-f]{8}(-[A-Za-z0-9._-]+)?\.yml$/;

/** True when the base name of path (after the last "/") is prov-YYYY-<hex8>[-suffix].yml. */
export function isRecordFilePath(path: string): boolean {
  return recordFileName.test(path.slice(path.lastIndexOf("/") + 1));
}

/** Resolves each associated spec in order; null when it has no path or no file. */
export function specHashes(r: HashInput, resolve: SpecResolver): (string | null)[] {
  return (r.associated_specs ?? []).map((s) => {
    const path = s?.path;
    if (typeof path === "string" && isRecordFilePath(path)) return RECORD_FILE_MARKER;
    return typeof path === "string" && path !== "" ? resolve(path) : null;
  });
}

/** The exact bytes content_hash covers (§4), given resolved spec hashes. */
export function contentCanonical(r: HashInput, hashes: (string | null)[]): string {
  return canonical({ ...(genObject(r) as { [key: string]: Canon }), associated_specs: hashes });
}

/** content_hash (§4): the behavioral fields plus each spec's content hash. */
export function contentHash(r: HashInput, specs: SpecResolver | (string | null)[]): string {
  const hashes = typeof specs === "function" ? specHashes(r, specs) : specs;
  return hashString(contentCanonical(r, hashes));
}

function owned(regions: Region[]): Region[] {
  return regions.filter((r) => r.tier === undefined || r.tier === "" || r.tier === "owned");
}

/** The exact bytes code_hash covers (§7). */
export function codeCanonical(regions: Region[]): string {
  const enc = owned(regions).map((r) => canonical({ gen: r.gen ?? "", pkg: r.pkg ?? "" }));
  enc.sort(compareBytes);
  return "[" + enc.join(",") + "]";
}

/** code_hash (§7) over owned regions; throws UnstampedError on a missing gen or pkg. */
export function codeHash(regions: Region[]): string {
  for (const r of owned(regions)) {
    if (!r.gen || !r.pkg) throw new UnstampedError();
  }
  return hashString(codeCanonical(regions));
}
