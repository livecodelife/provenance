/**
 * GENERATED CODE - DO NOT MODIFY
 */
import { type ValidationResult, BlobRef } from '@atproto/lexicon'
import { CID } from 'multiformats/cid'
import { validate as _validate } from '../../../../lexicons.js'
import {
  type $Typed,
  is$typed as _is$typed,
  type OmitKey,
} from '../../../../util.js'

const is$typed = _is$typed,
  validate = _validate
const id = 'dev.provisio.provenance.defs'

/** "sha256:" followed by 64 lowercase hex digits. */
export type Hash = string
/** One associated spec's contribution to the content hash: a hash, the constant "record-file/v1" for a spec that is a record file, or the empty string when the file could not be read. */
export type SpecHash = string

/** A proof artifact. runCommand is local-only and never published. */
export interface Spec {
  $type?: 'dev.provisio.provenance.defs#spec'
  path: string
  type?: string
}

const hashSpec = 'spec'

export function isSpec<V>(v: V) {
  return is$typed(v, id, hashSpec)
}

export function validateSpec<V>(v: V) {
  return validate<Spec & V>(v, id, hashSpec)
}

/** A record without its lifecycle and envelope fields (status, version, reason, sealed commit). Absent and empty are the same value. */
export interface Content {
  $type?: 'dev.provisio.provenance.defs#content'
  title: string
  type?: 'brief' | 'blueprint' | 'imprint' | (string & {})
  /** The day the record was created, YYYY-MM-DD. */
  createdAt?: string
  author?: string
  intent: string
  constraints?: string[]
  affectedScope?: string[]
  forbiddenScope?: string[]
  related?: string[]
  tags?: string[]
  aliases?: string[]
  /** Id of the parent record. */
  implements?: string
  /** Name of the repository holding the parent when it is not local. */
  implementsRepo?: string
  /** at-uri of the parent's head record (at://<did>/dev.provisio.provenance.record/<recordId>) when it lives in another account; requires implements, whose value is that record id. */
  implementsUri?: string
  /** Parent version last confirmed against; 0 or absent means unpinned. */
  implementsVersion?: number
  associatedSpecs?: Spec[]
  associatedTraces?: string[]
  monitors?: string[]
  locked?: boolean
}

const hashContent = 'content'

export function isContent<V>(v: V) {
  return is$typed(v, id, hashContent)
}

export function validateContent<V>(v: V) {
  return validate<Content & V>(v, id, hashContent)
}
