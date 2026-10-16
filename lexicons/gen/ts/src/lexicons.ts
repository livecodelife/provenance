/**
 * GENERATED CODE - DO NOT MODIFY
 */
import {
  type LexiconDoc,
  Lexicons,
  ValidationError,
  type ValidationResult,
} from '@atproto/lexicon'
import { type $Typed, is$typed, maybe$typed } from './util.js'

export const schemaDict = {
  DevProvisioProvenanceDefs: {
    lexicon: 1,
    id: 'dev.provisio.provenance.defs',
    description:
      'Shared shapes of provenance records. Field names are the snake_case names of the YAML record in camelCase.',
    defs: {
      hash: {
        type: 'string',
        description: '"sha256:" followed by 64 lowercase hex digits.',
        maxLength: 71,
      },
      specHash: {
        type: 'string',
        description:
          'One associated spec\'s contribution to the content hash: a hash, the constant "record-file/v1" for a spec that is a record file, or the empty string when the file could not be read.',
        maxLength: 71,
      },
      spec: {
        type: 'object',
        description:
          'A proof artifact. runCommand is local-only and never published.',
        required: ['path'],
        properties: {
          path: {
            type: 'string',
            maxLength: 1024,
          },
          type: {
            type: 'string',
            maxLength: 64,
          },
        },
      },
      content: {
        type: 'object',
        description:
          'A record without its lifecycle and envelope fields (status, version, reason, sealed commit). Absent and empty are the same value.',
        required: ['title', 'intent'],
        properties: {
          title: {
            type: 'string',
            maxLength: 1024,
          },
          type: {
            type: 'string',
            knownValues: ['brief', 'blueprint', 'imprint'],
          },
          createdAt: {
            type: 'string',
            description: 'The day the record was created, YYYY-MM-DD.',
            maxLength: 10,
          },
          author: {
            type: 'string',
            maxLength: 512,
          },
          intent: {
            type: 'string',
            maxLength: 200000,
          },
          constraints: {
            type: 'array',
            maxLength: 2000,
            items: {
              type: 'string',
              maxLength: 20000,
            },
          },
          affectedScope: {
            type: 'array',
            maxLength: 2000,
            items: {
              type: 'string',
              maxLength: 1024,
            },
          },
          forbiddenScope: {
            type: 'array',
            maxLength: 2000,
            items: {
              type: 'string',
              maxLength: 1024,
            },
          },
          related: {
            type: 'array',
            maxLength: 2000,
            items: {
              type: 'string',
              maxLength: 128,
            },
          },
          tags: {
            type: 'array',
            maxLength: 200,
            items: {
              type: 'string',
              maxLength: 128,
            },
          },
          aliases: {
            type: 'array',
            maxLength: 2000,
            items: {
              type: 'string',
              maxLength: 128,
            },
          },
          implements: {
            type: 'string',
            description: 'Id of the parent record.',
            maxLength: 128,
          },
          implementsRepo: {
            type: 'string',
            description:
              'Name of the repository holding the parent when it is not local.',
            maxLength: 256,
          },
          implementsVersion: {
            type: 'integer',
            minimum: 0,
            description:
              'Parent version last confirmed against; 0 or absent means unpinned.',
          },
          associatedSpecs: {
            type: 'array',
            maxLength: 500,
            items: {
              type: 'ref',
              ref: 'lex:dev.provisio.provenance.defs#spec',
            },
          },
          associatedTraces: {
            type: 'array',
            maxLength: 500,
            items: {
              type: 'string',
              maxLength: 1024,
            },
          },
          monitors: {
            type: 'array',
            maxLength: 500,
            items: {
              type: 'string',
              maxLength: 1024,
            },
          },
          locked: {
            type: 'boolean',
          },
        },
      },
    },
  },
  DevProvisioProvenanceRecord: {
    lexicon: 1,
    id: 'dev.provisio.provenance.record',
    defs: {
      main: {
        type: 'record',
        description:
          'The mutable head of a provenance record; the key is the provenance id. Updated with swapRecord on its CID.',
        key: 'any',
        record: {
          type: 'object',
          required: [
            'status',
            'version',
            'contentHash',
            'genHash',
            'latest',
            'content',
            'updatedAt',
          ],
          properties: {
            status: {
              type: 'string',
              knownValues: ['draft', 'open', 'implemented', 'deprecated'],
              maxLength: 32,
            },
            version: {
              type: 'integer',
              minimum: 1,
            },
            contentHash: {
              type: 'ref',
              ref: 'lex:dev.provisio.provenance.defs#hash',
            },
            genHash: {
              type: 'ref',
              ref: 'lex:dev.provisio.provenance.defs#hash',
            },
            latest: {
              type: 'ref',
              ref: 'lex:com.atproto.repo.strongRef',
              description: 'The newest version record.',
            },
            content: {
              type: 'ref',
              ref: 'lex:dev.provisio.provenance.defs#content',
              description:
                'The current content, including the parent pin, which moves without a new version.',
            },
            source: {
              type: 'ref',
              ref: 'lex:dev.provisio.provenance.record#source',
            },
            project: {
              type: 'string',
              format: 'at-uri',
              description:
                'The dev.provisio.project record this record belongs to; optional, so a head without a project validates.',
            },
            updatedAt: {
              type: 'string',
              format: 'datetime',
            },
          },
        },
      },
      source: {
        type: 'object',
        description: 'Where the record lives in git.',
        properties: {
          repo: {
            type: 'string',
            format: 'uri',
          },
          path: {
            type: 'string',
            maxLength: 1024,
          },
          commit: {
            type: 'string',
            maxLength: 64,
          },
        },
      },
    },
  },
  DevProvisioProvenanceVersion: {
    lexicon: 1,
    id: 'dev.provisio.provenance.version',
    defs: {
      main: {
        type: 'record',
        description:
          'One immutable version of a provenance record. Never updated; chained to its parent twice, by contentHash and by strong reference.',
        key: 'tid',
        record: {
          type: 'object',
          required: [
            'record',
            'recordId',
            'version',
            'contentHash',
            'author',
            'createdAt',
            'specHashes',
            'content',
          ],
          properties: {
            record: {
              type: 'string',
              format: 'at-uri',
              description: 'The head record this is a version of.',
            },
            recordId: {
              type: 'string',
              maxLength: 128,
              description: 'The provenance id; the key of the head.',
            },
            version: {
              type: 'integer',
              minimum: 1,
            },
            contentHash: {
              type: 'ref',
              ref: 'lex:dev.provisio.provenance.defs#hash',
            },
            parentHash: {
              type: 'ref',
              ref: 'lex:dev.provisio.provenance.defs#hash',
              description:
                "The previous version's contentHash; absent on version 1.",
            },
            parent: {
              type: 'ref',
              ref: 'lex:com.atproto.repo.strongRef',
              description: 'The previous version record; absent on version 1.',
            },
            reason: {
              type: 'string',
              maxLength: 20000,
              description: 'Why this version exists; required after version 1.',
            },
            reasonKind: {
              type: 'string',
              knownValues: ['fix'],
              maxLength: 32,
            },
            author: {
              type: 'string',
              maxLength: 512,
            },
            createdAt: {
              type: 'string',
              format: 'datetime',
            },
            commit: {
              type: 'string',
              maxLength: 64,
              description: 'Git commit the version was published from.',
            },
            specHashes: {
              type: 'array',
              maxLength: 500,
              description:
                'One entry per content.associatedSpecs, in the same order.',
              items: {
                type: 'ref',
                ref: 'lex:dev.provisio.provenance.defs#specHash',
              },
            },
            content: {
              type: 'ref',
              ref: 'lex:dev.provisio.provenance.defs#content',
            },
            origin: {
              type: 'string',
              knownValues: ['git', 'ui'],
              maxLength: 16,
            },
            baseVersion: {
              type: 'integer',
              minimum: 1,
              description:
                'For origin ui: the version the editor started from.',
            },
            baseContentHash: {
              type: 'ref',
              ref: 'lex:dev.provisio.provenance.defs#hash',
              description: "For origin ui: that version's contentHash.",
            },
          },
        },
      },
    },
  },
  ComAtprotoRepoStrongRef: {
    lexicon: 1,
    id: 'com.atproto.repo.strongRef',
    description: 'A URI with a content-hash fingerprint.',
    defs: {
      main: {
        type: 'object',
        required: ['uri', 'cid'],
        properties: {
          uri: {
            type: 'string',
            format: 'at-uri',
          },
          cid: {
            type: 'string',
            format: 'cid',
          },
        },
      },
    },
  },
} as const satisfies Record<string, LexiconDoc>
export const schemas = Object.values(schemaDict) satisfies LexiconDoc[]
export const lexicons: Lexicons = new Lexicons(schemas)

export function validate<T extends { $type: string }>(
  v: unknown,
  id: string,
  hash: string,
  requiredType: true,
): ValidationResult<T>
export function validate<T extends { $type?: string }>(
  v: unknown,
  id: string,
  hash: string,
  requiredType?: false,
): ValidationResult<T>
export function validate(
  v: unknown,
  id: string,
  hash: string,
  requiredType?: boolean,
): ValidationResult {
  return (requiredType ? is$typed : maybe$typed)(v, id, hash)
    ? lexicons.validate(`${id}#${hash}`, v)
    : {
        success: false,
        error: new ValidationError(
          `Must be an object with "${hash === 'main' ? id : `${id}#${hash}`}" $type property`,
        ),
      }
}

export const ids = {
  DevProvisioProvenanceDefs: 'dev.provisio.provenance.defs',
  DevProvisioProvenanceRecord: 'dev.provisio.provenance.record',
  DevProvisioProvenanceVersion: 'dev.provisio.provenance.version',
  ComAtprotoRepoStrongRef: 'com.atproto.repo.strongRef',
} as const
