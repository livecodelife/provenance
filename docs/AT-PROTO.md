# Connecting provenance to AT Proto

How a provenance record gets from a git repository onto AT Proto, who can write it, who can read it, and how
records in different accounts point at each other. It joins what DESIGN.md ("Identity", "AT Proto storage and
publishing", "Lineage"), PLAN-IDENTITY-UI-FIRST.md (sessions and the CLI) and LEXICONS.md (field names) say in
pieces. Where those documents decide something, this one links to it rather than deciding again.

Status words used below: **built** (on main), **in flight** (a branch or PR exists), **not started**, and
**proposal** (a design for Caleb to confirm, not a decision).

## 1. The picture

```
 git repo (canonical)                         AT Proto
 ───────────────────                          ────────
 provenance/*.yml        provenance publish   author's PDS repo
 provenance/.versions/ ───────────────────▶   dev.provisio.provenance.version   (append-only, one per version)
 .provenance/versions.jsonl  (needs a         dev.provisio.provenance.record    (mutable head, key = record id)
                              Provisio        dev.provisio.project              (proposal, section 3)
                              session)                    │
        ▲                                                 │ Jetstream
        │ provenance import                               ▼
        │ (UI-made versions)                  AppView / indexer (Postgres)
        └─────────────────────────────────    lists, history, backlinks, search; the UI and CLI read it
```

- **Git is canonical.** Publishing copies versions out of git history; nothing is edited on the PDS first, except
  versions made in the Provisio UI, which `import` brings back (section 4).
- **One DID ties it together.** The account that publishes owns the records. A record's address is
  `at://<did>/dev.provisio.provenance.record/<record id>`, so the same record id in two accounts is two records.
- **The PDS keeps only the latest value of a record,** so history is kept by writing each version as its own
  immutable record rather than by updating one.
- **Local commands never touch the network.** Everything that reads or edits a record in a working tree works
  logged out and offline (section 6).

## 2. Identity and auth providers

Every Provisio user is an AT Proto account, identified by a DID. Full reasoning is in DESIGN.md, "Identity".

| Provider | How it signs in | Account lives on |
| --- | --- | --- |
| Provisio handle | AT Proto OAuth against the Provisio PDS (the default) | Provisio PDS |
| Bluesky or any AT Proto account | AT Proto OAuth against the account's own server | their own server |
| GitHub or GitLab | login option on the Provisio sign-in page, handled by the auth service | Provisio PDS (created on first login) |

**The auth service** (`auth.provisio.dev`, in provisio-backend) is the one AT Proto OAuth client for the product.
It keeps the PDS sessions (DPoP-bound); no other component sees PDS tokens. After sign-in it issues a Provisio
session whose subject is the DID. The record server and the CLI use only that session.

**Tokens** (PLAN-IDENTITY-UI-FIRST §3):

| Token | Lifetime | Used by |
| --- | --- | --- |
| Access token | 15 minutes, signed | UI cookie; CLI and API as `Authorization: Bearer` |
| Refresh token | 30 days, rotating | UI cookie; CLI credential store |
| Personal token (`pvt_…`) | until revoked or the chosen expiry | CI, as `PROVISIO_TOKEN`; scoped (`records:read`, `records:write`); cannot mint tokens |

**CLI sign-in.** `provenance login` runs a device-authorization flow (RFC 8628): it prints a verification URL and a
short code, opens the browser unless `--no-browser`, and polls until the code is approved, denied or expired.
`--with-token` stores a personal token read from stdin. `--server` overrides the auth URL
(`$PROVISIO_AUTH_URL`, default `https://auth.provisio.dev`). Credentials come from `PROVISIO_TOKEN` first, then the
OS keychain, then `~/.provisio/credentials` (mode 0600). `logout` and `whoami` complete the set.
Status: `login`, `--with-token` and the token store are **built**. A `--provider provisio|bluesky` flag
(default `provisio`) is **not started** (prov-cli).

**Where the browser pages come from.** The pages a user sees during `provenance login` (enter code, sign in,
approve or deny, finished) are rendered by the auth service in provisio-backend
(`renderApprovalPage` in `packages/auth/src/approval-page.ts`). They are not in this repository, and the CLI serves
no page of its own: it only opens the verification URL. They are **being restyled** with the Provisio palette and
wordmark (prov-web, provisio-backend); this section should be updated when that lands.

**Scopes.** Provisio asks only for its own collections (`repo:dev.provisio.provenance.record`-style grants, or a
published permission set bundling `dev.provisio.*`), so consent reads "write your provenance records", not "full
account access". Older servers may offer only broad access; check before relying on it.

**Writes from a non-AT-Proto login.** A GitHub or GitLab user's PDS account has no OAuth session, so the auth
service writes with an app password it generated at account creation and stores encrypted. See DESIGN.md.

## 3. What gets published

Namespace `dev.provisio.*`. Field-by-field spelling is in LEXICONS.md; this is the shape.

| Collection | Key | Mutable | Holds | Status |
| --- | --- | --- | --- | --- |
| `dev.provisio.provenance.record` | the record id (`prov-2026-1a2b3c4d`) | yes (head) | current `content`, `status`, `version`, `contentHash`, `genHash`, `latest` (strong ref to the newest version), optional `source` and `project` | built |
| `dev.provisio.provenance.version` | TID | never | `record` (head URI), `recordId`, `version`, `contentHash`, `parentHash`, `parent` (strong ref), `reason`, `author`, `commit`, `specHashes`, `content`, optional `origin` / `baseVersion` / `baseContentHash` | built |
| `dev.provisio.project` | TID or slug | yes | a codebase and the records that belong to it | **proposal** below; not in `lexicons/` yet |
| `dev.provisio.lineage.copy` | TID | never | a copy's pinned source project and record versions | not started (DESIGN.md, "Lineage") |

**Two chains, and why the hashes are not AT CIDs.** `parentHash` chains each version to the previous version's
`contentHash` (the provenance chain). `parent` chains it to the previous version record's `{uri, cid}` (the AT
Proto chain). A verifier checks both. An AT CID covers the whole AT record, status and envelope included, so it
cannot stand in for `contentHash`, which covers content alone (HASHING.md §4). `contentHash` can be recomputed
from `content` and `specHashes` alone: `Snapshot.Verify` in Go, `verifyVersion` in TypeScript, both pinned by the
conformance corpus.

**Not in the content hash:** lifecycle fields and lineage fields (`implementsUri`, `implementsRepo`). A parent
link can therefore be added or corrected without creating a version.

**Not published:** `pins.jsonl`. A pin appears only as the head's current `content.implementsVersion`.

### Project and source records (proposal, for Caleb to confirm)

The head already has two optional fields that point at these: `source` (`{repo, path, commit}`, where the record
sits in git) and `project` (an at-uri of a `dev.provisio.project`). The publish workflow already passes a project
URI (`PROVISIO_PROJECT`, e.g. `at://did:plc:…/dev.provisio.project/provenance`). What the project record itself
holds is not yet written down in a lexicon. Proposed fields:

| Field | Type | Notes |
| --- | --- | --- |
| `name` | string | required; display name |
| `description` | string | optional |
| `repo` | string | the canonical git URL, e.g. `https://github.com/livecodelife/provenance` |
| `defaultBranch` | string | optional; where `commit` values on the project's records are reachable from |
| `createdAt` | datetime | required |

Proposed rules:

- **A project does not list its records.** The head's `project` field is the link, so the project record stays small
  and does not need a write every time a record is published. "Which records belong" is a backlink query on the
  indexer.
- **`source` stays on the head, not the project.** A record's path and commit differ per record; the repo is the
  project's. If a head has `source.repo` and a `project`, the AppView reports a mismatch with the project's `repo`
  rather than rejecting the write.
- **One project per repository,** created by the first publish with `--project`, and referenced by every later
  publish. A monorepo can use one project per root (each `.provenance.yml` root).
- **Questions for Caleb:** (a) Should `dev.provisio.project` be mutable with a TID key, or keyed by a slug so the
  at-uri is readable? The workflow's variable assumes a slug (`…/dev.provisio.project/provenance`). (b) Is a
  private repo allowed to have a project record at all, given everything published is public? (c) Does a copy
  (lineage) create a new project in the copier's repo, as DESIGN.md says, and should it carry an `origin`
  at-uri back to the source project?

## 4. The publish flow

`provenance publish [id…] [--all] [--plan] [--project <at-uri>] [--yes]`. Needs a Provisio session.

1. **Opt in.** Publishing is per project and explicit; nothing is public until someone runs it (or CI is switched
   on with `PROVISIO_PUBLISH=true`).
2. **Plan.** For each record the CLI reads its local snapshots and compares them with what is already on the PDS.
   `--plan` prints exactly which versions would become public and writes nothing.
3. **Confirm.** It shows the plan and asks, unless `--yes`.
4. **Write versions oldest first.** Each missing version becomes a `dev.provisio.provenance.version` record whose
   `parent` is the previous version's `{uri, cid}`. Long histories go in batches under the `applyWrites` cap.
5. **Update the head** with `swapRecord` on the head's CID (compare and swap). If another publisher moved the head,
   the write is refused with the current head and the CLI re-plans; a lost race is a skip, not an overwrite.
6. **Re-running is a no-op** when nothing changed.

`provenance repin` moves only the head's `content.implementsVersion`: it writes no version, so the newest
version's `content.implementsVersion` can be older than the head's.

**CI.** `.github/workflows/publish.yml` publishes on a push to main using a scoped personal token
(`PROVISIO_TOKEN`, `records:write`), in a protected GitHub Environment, never on forks. While `PROVISIO_PUBLISH` is
not `true` it only prints the plan.

**Coming back: `import`.** A version made in the Provisio UI is written to the PDS with `origin: ui` plus the
version and content hash it was based on. `provenance import` brings such versions into the git project as new
versions with their reason and author. It refuses when the local record has moved past that base, and says to
`revise` first; it does not merge. Records edited in a git repository keep the pull-request flow.

**`clone <DIR> [id…]`** creates a local project from published records. Public records need no login.

**Conflict cases**

| Situation | Result |
| --- | --- |
| Head moved since the plan | write refused, re-plan |
| Local version exists whose hash differs from the published version with the same number | stop and report; never overwrite a version record |
| UI version based on an older local version | `import` refuses; `revise` locally, then publish |
| Same record id published from two repositories | two heads would collide; the second publisher is refused by the swap and must resolve the id clash in git |

## 5. Lineage across accounts

A record can name a parent in another account. Four fields, none of them in the content hash:

| Field | Meaning |
| --- | --- |
| `implements` | the parent's record id |
| `implementsRepo` | the repo that holds it, when not this one |
| `implementsUri` | `at://<did>/dev.provisio.provenance.record/<record id>`, the parent's head in another account; requires `implements`, whose value must be the same record id |
| `implementsVersion` | the parent version last confirmed; 0 or absent means unpinned |

- **Built:** the lexicon field, the model, lint (the at-uri must parse and name the record in `implements`),
  carrying it through `publish` and `import`, and `create --implements <id|at-uri>`. A cross-account parent is
  created unpinned and the CLI says so.
- **Not started (P3):** resolving the parent. `stale`, `check` and `lint` should look the parent up through the
  AppView, falling back to the owning account's PDS (`getRecord` on the DID's repo), to learn its current version
  and compare it with `implementsVersion`. Waiting on decisions about the AppView.
- **Not started (P4b):** `publish` printing a note when a record's parent is in another account, since the pin
  cannot be checked locally.
- **Copies keep record ids; the DID tells them apart.** A copy's heads live in the copier's repo with an `origin`
  pin (the source's address plus the CID copied). Forks compare by content hash, never by version number. See
  DESIGN.md, "Lineage".

## 6. Reading

| What | Source | Needs |
| --- | --- | --- |
| Your own records, heads and versions | your PDS (`listRecords`, `getRecord`) | session for private-by-session API; the data itself is public |
| Another account's records | the AppView, or that account's PDS directly | nothing for public reads |
| Backlinks (what implements, pins or copied this), search, per-user lists | the AppView, fed by Jetstream for `dev.provisio.*` | nothing for public reads |
| Anything in a working tree | git | nothing |

**Works logged out and offline:** `init`, `create`, `open`, `edit`, `revise`, `complete`, `deprecate`, `repin`,
`status`, `history`, `diff`, `lint`, `check`, `stale`, `next`, `govern`, `graph`, `context`, `search`, `reconcile`,
`compile`, `migrate`, `hooks`, `evidence`, `run-specs`, `add-scope`. These never read the credential store.
`publish` and `import` need a session; `clone` needs one only for private reads. A command that needs a session and
has none exits with status 3 and says to run `provenance login` (or set `PROVISIO_TOKEN` in CI); it never prompts
in a non-interactive shell.

## 7. Privacy, scopes and limits

- **Everything published is public,** and copies of it may be cached or indexed. The publish plan is the
  confirmation step; `--yes` is for scripts.
- **No deletion guarantee.** A PDS delete leaves no tombstone for downstream indexers, and mirrors may keep old
  values. Do not publish a record that should not be public.
- **Private records:** AT Proto Spaces (alpha since August 2026) give access control, not encryption, and promise
  no stability. Revisit when it ships.
- **Limits:** the PDS record size cap (believed 1 MiB) and the 200-writes-per-`applyWrites` cap are the real
  ceilings; the largest real `intent` is about 28 KB, so one version fits. A repository with hundreds of versions
  needs batched publishing.
- **Scopes:** a personal token carries `records:read` / `records:write`, by project.

## 8. Status and open questions

| Piece | Status | Owner |
| --- | --- | --- |
| Record and version lexicons, Go and TypeScript types, conformance corpus | built | prov-lib |
| `publish`, `import`, `clone`, `login`, token store | built | prov-cli |
| CI publish workflow | built; switched on by Caleb | prov-cli / Caleb |
| `implementsUri` field, lint, `create --implements <at-uri>`, publish and import carry | built | prov-lib |
| `login --provider provisio\|bluesky` | not started | prov-cli |
| Login pages restyled to the Provisio palette | in flight | prov-web |
| `dev.provisio.project` lexicon and project rules | proposal (section 3) | Caleb to confirm |
| Parent resolution through the AppView (P3) | not started; waits on AppView decisions | prov-lib |
| Publish note for cross-account parents (P4b) | not started; after P3 | prov-lib |
| `dev.provisio.lineage.copy`, origin pins, pull-forward merge | not started | – |
| Own PDS, auth service and indexer on the Miren cluster | per DESIGN.md, "Operations" | provisio-mgr |

Open questions for Caleb: the project record (section 3); whether a cross-account `implementsUri` should ever be
resolved by the CLI without the AppView; whether to publish `pins.jsonl`; and when to move off broad PDS scopes
once the target servers support permission sets.
