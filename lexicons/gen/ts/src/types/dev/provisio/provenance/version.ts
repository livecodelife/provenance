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
import type * as DevProvisioProvenanceDefs from './defs.js'
import type * as ComAtprotoRepoStrongRef from '../../../com/atproto/repo/strongRef.js'

const is$typed = _is$typed,
  validate = _validate
const id = 'dev.provisio.provenance.version'

export interface Main {
  $type: 'dev.provisio.provenance.version'
  /** The head record this is a version of. */
  record: string
  /** The provenance id; the key of the head. */
  recordId: string
  version: number
  contentHash: DevProvisioProvenanceDefs.Hash
  parentHash?: DevProvisioProvenanceDefs.Hash
  parent?: ComAtprotoRepoStrongRef.Main
  /** Why this version exists; required after version 1. */
  reason?: string
  reasonKind?: 'fix' | (string & {})
  author: string
  createdAt: string
  /** Git commit the version was published from. */
  commit?: string
  /** One entry per content.associatedSpecs, in the same order. */
  specHashes: DevProvisioProvenanceDefs.SpecHash[]
  content: DevProvisioProvenanceDefs.Content
  origin?: 'git' | 'ui' | (string & {})
  /** For origin ui: the version the editor started from. */
  baseVersion?: number
  baseContentHash?: DevProvisioProvenanceDefs.Hash
  [k: string]: unknown
}

const hashMain = 'main'

export function isMain<V>(v: V) {
  return is$typed(v, id, hashMain)
}

export function validateMain<V>(v: V) {
  return validate<Main & V>(v, id, hashMain, true)
}

export {
  type Main as Record,
  isMain as isRecord,
  validateMain as validateRecord,
}
