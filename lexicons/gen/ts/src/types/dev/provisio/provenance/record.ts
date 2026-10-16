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
const id = 'dev.provisio.provenance.record'

export interface Main {
  $type: 'dev.provisio.provenance.record'
  status: 'draft' | 'open' | 'implemented' | 'deprecated' | (string & {})
  version: number
  contentHash: DevProvisioProvenanceDefs.Hash
  genHash: DevProvisioProvenanceDefs.Hash
  latest: ComAtprotoRepoStrongRef.Main
  content: DevProvisioProvenanceDefs.Content
  source?: Source
  /** The dev.provisio.project record this record belongs to; optional, so a head without a project validates. */
  project?: string
  updatedAt: string
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

/** Where the record lives in git. */
export interface Source {
  $type?: 'dev.provisio.provenance.record#source'
  repo?: string
  path?: string
  commit?: string
}

const hashSource = 'source'

export function isSource<V>(v: V) {
  return is$typed(v, id, hashSource)
}

export function validateSource<V>(v: V) {
  return validate<Source & V>(v, id, hashSource)
}
