# Plan: repos, build, deploy and governance for the TypeScript services and UI

Proposal only, for provisio-mgr and Caleb. Nothing here is created, built, deployed or contacted. It
answers the first task in `.herd/brief-prov-web.md` and fills in the TypeScript side of
`docs/PLAN-IDENTITY-UI-FIRST.md` (milestones M1 auth service, M3 record server, M5 UI sign-in and
`RecordSource`, M8 Provisio-handle sign-in). Facts about third-party software (Miren's Node support,
`@atproto/*` packages, `node:sqlite`) are from memory and are listed in section 8 to verify; none is
relied on before its check passes.

## 1. Repo layout

**Recommendation: one new repo, `provisio-backend`, a small npm-workspaces monorepo. The UI stays in
`linespec-provenance-ui`. Lexicons and the hashing package stay in `provenance`.**

```
provisio-backend/
  .provenance.yml          provenance v2 config (dir: provenance)
  .miren/app.toml          one Miren app: provisio-api
  package.json             workspaces: packages/*
  packages/
    shared/                session tokens (sign/verify, JWKS), DID/handle types, error shapes, SQLite
                           open + migrate + online backup, config loading from env
    auth/                  AT Proto OAuth client, user table, sessions, device flow, personal tokens
    records/               record server: list/get/putVersion/head CAS/hash verification
    api/                   the one process: mounts auth + records routes, health, backup timer
  test/                    cross-package tests (fake PDS, end-to-end sign-in and write)
  provenance/              records (v2), .versions/, .evidence/
```

Why one repo and one deployable: auth and records share the session verifier, the SQLite store, the
fake-PDS test double and the release cadence, and `docs/PLAN-IDENTITY-UI-FIRST.md` already says
`provisio-api` is one process behind two hostnames (`auth.` and `api.`), split later if needed. Packages
keep the seam clean so the split costs a new `app.toml` and not a rewrite. Two repos now would mean
publishing `shared` before there is anything to share it with.

Why the UI is not in it: it deploys to Vercel on its own cadence, has its own records (migrate-dry is
mid-way through UI v2 there), and its `RecordSource` change is small. M5 is a branch of
`linespec-provenance-ui`.

Why lexicons stay in `provenance`: the Go CLI and sedum need the same JSON. `provisio-backend` consumes
them (and the generated TS types) at a pinned SHA, the way it consumes hashing below.

**The hashing package.** `provenance/ts` (`@livecodelife/provenance-hashing`, private, Node >= 23.6,
runs `.ts` directly) is the only TypeScript implementation of `gen_hash` and `content_hash` and the
record server must not fork it. npm cannot install a subdirectory of a git repo, so options are:
(a) publish it to GitHub Packages as private; (b) a vendored copy at `packages/shared/vendor/hashing/`
with a test that fails when the files differ from upstream at the pinned SHA; (c) move it to its own
tiny repo. **Recommend (b) to start, (a) when a second consumer (the UI v2 parser, per migrate-dry's
analysis) needs it.** The conformance corpus (`conformance/hashing/cases.json`) runs against whichever
copy the backend uses, in CI, on every change.

**Names.** `provisio-backend` matches the existing plan. Package scope `@provisio/*`, all `private`.
(`provisio-ops` is for PDS deploy config and the backup wrapper, as already planned; not mine.)

## 2. Language and tooling

| Choice | Pick | Reason |
| --- | --- | --- |
| Runtime | Node 24 LTS, ESM only | Vercel's current default; `provenance/ts` already needs >= 23.6 |
| TypeScript | strict, `tsc --noEmit` as the type gate; Node's type stripping runs the source (no emit step) if Miren's Node build allows it, otherwise `tsc` emits to `dist/` | fewest moving parts; mirrors `provenance/ts` |
| HTTP | Hono on `@hono/node-server` | small, typed, no framework state; handlers are plain functions of `Request` |
| AT Proto | `@atproto/oauth-client-node` (confidential client, DPoP, PAR, JWKS), `@atproto/api` for repo calls, `@atproto/lex-cli` for lexicon types | the plan's rule: do not hand-roll OAuth |
| Storage | SQLite on a Miren disk via the built-in `node:sqlite` (no native addon, so the buildpack needs no compiler) | one instance, one disk, in-app backup per the proof |
| Tests | `node:test` + `node:assert` for everything, `tsx`-free | same runner as `provenance/ts`; no extra framework |
| Lint/format | eslint + prettier, config files handwritten (exempt per COMMON) | |

**Test setup.** Three layers, all run by `npm test`:
1. Unit tests per package (token sign/verify and rotation, device-code state machine, scope checks,
   canonical encoding via the corpus).
2. A **fake PDS and fake authorization server** in `test/fakes/` (XRPC `getRecord`, `listRecords`,
   `putRecord` with `swapRecord`, `applyWrites`; handle and DID resolution; OAuth endpoints with DPoP
   nonce). M1 and M3 done-checks run against it in CI with no network.
3. A **live suite** (`npm run test:live`, never in CI by default) against a real Bluesky test account
   whose credentials are Caleb's; it only runs when `PROVISIO_LIVE_*` variables are set, and it writes
   only to a scratch collection prefix. This is the "real Bluesky test account" done-check from M1 and
   M3.

## 3. Build and deploy to Miren

One app, `provisio-api`, on the `miren-provisio` cluster. Everything below follows the proof's findings
(`MIREN_PDS_PROOF.md`) and is a plan; Caleb and provisio-mgr approve each cluster action.

- **Build.** `.miren/app.toml` with the start command pointing at the built entry. The proof showed the Go
  buildpack's binary is `/bin/app`; the Node buildpack's behaviour is unverified, so the first deploy is a
  placeholder (a `/healthz` server) to learn it. If Miren's Node support doesn't fit, fall back to a
  Dockerfile (handwritten, exempt).
- **Order.** Placeholder deploy first and `miren route unset-default` at once (finding 1), then real
  routes `auth.` and `api.` on the `provisio.dev` zone, then the real image. Raise `miren route timeout`
  only if a streaming route appears (not in M1/M3).
- **State.** One fixed instance, one disk for the SQLite file. Backups run **inside the process**, not as
  a wrapper: every interval `node:sqlite`'s online backup (or `VACUUM INTO`) to a temp file, encrypt to
  the public `age` recipient, SigV4 PUT to R2. This is the pattern the proof validated; the `age` identity
  stays off the cluster. A restore drill (into a new disk) is a done-check before any real user.
- **Secrets** (names only, in `miren secret set`, bound with `miren env set -e KEY --ref`): OAuth client
  private JWKS, session-signing key, cookie secret, token-hash pepper, R2 key id and secret. Values are
  Caleb's; agents never see or post them. A rotation needs a new `env set --ref` (finding 3).
- **CI deploy.** GitHub Actions with Miren's OIDC binding, which can deploy and read logs but not back up
  (finding 6). Workflow: `npm ci`, typecheck, `npm test`, `provenance lint --enforcement strict`,
  `provenance stale`, `provenance check --range` on PRs (the pattern sedum and the UI repo now use, CLI
  pinned to a full SHA), then deploy on main. Handwritten, exempt (workflow file). The workflow doesn't
  run until Caleb clears it; before then I only deploy placeholders by hand when approved.
- **UI (M5).** Stays on Vercel. For the first-party cookie it needs a `provisio.dev` host (custom domain);
  that is Caleb's DNS/Vercel step, listed here so it's asked for early.

## 4. Governance by provenance records (v2)

`provisio-backend` starts as a v2 repo (`provenance init`, `.provenance.yml`, `commit_tag_required: true`,
strict lint, hooks installed per developer). Following COMMON's "revise, don't open new records":
**few, broad records that get revised as the code matures.**

| Level | Record | Notes |
| --- | --- | --- |
| Brief | one: "Provisio identity backend" | links to DESIGN Identity and the identity plan |
| Blueprint | `auth`, `records`, `shared`, `ops` (build, deploy, backup) | specs below; blueprints own integration specs |
| Imprints | per blueprint, one per area, not per file (e.g. auth: oauth+sessions; device flow+tokens; records: reads; writes+CAS+hash verify) | revised in place when the area changes |

Specs (the proof the blueprints carry): `node --test` over the package's tests and the conformance corpus
for `shared`/`records`, the fake-PDS end-to-end suite for `auth` and `records`, and for `ops` a script that
builds the image and runs the backup-and-restore round trip locally. `exclude_paths`: lockfile,
`.miren/**`, `.github/**`, tool configs, `provenance/**`. `affected_scope` per imprint lists the package
directories it owns. Tests and fixtures are handwritten and covered by the same records.

M5 in the UI repo is governed by that repo's records: a revise of the UI v2 reads imprint (migrate-dry's
`101d36d0`) if the files it governs change, otherwise one new blueprint "Provisio session and record-server
RecordSource" with two imprints (sign-in/session context; RecordSource + edit via CAS). I'd check with
migrate-dry before touching `src/lib/recordSources.ts`.

## 5. Generated (sedum) vs handwritten

Sedum has Go generator packages only (`cobra`, `go`; `docs/GENERATOR_BACKLOG.md` already lists the ts
package as missing). **Every TypeScript line in this plan is handwritten today.** I will post each
handwritten file in #provisio with a one-line reason. What is not handwritten:

- Lexicon types for TS and Go: generated by `lex-cli` from the lexicon JSON (a tool, not sedum).
- Config and build files (`package.json`, tsconfig, `.miren/app.toml`, Dockerfile, workflows): exempt.

Where a sedum TypeScript package would pay off (feeds the backlog, not a dependency of this plan):
route-handler + schema scaffolds (`addRoute{method,path,handler}`), SQLite table + migration + typed
accessor from a column list, and the broad `setFunction`/`setType` pair with literal-code kwargs, the same
shape as the Go package. The repeated shapes in `auth` and `records` (CRUD over a table, request
validation, typed errors) are exactly that. **Ask:** should I write the TypeScript generator package
proposal now (for sedum-grow), or build M1 handwritten and file the backlog items from what repeats? I
recommend the latter: the shapes should come from real code, and M1 must not wait.

## 6. Order: M1 -> M3 -> M5 (then M8)

| Step | Scope | Done-check (from the identity plan) |
| --- | --- | --- |
| 0 | After approval: create `provisio-backend` locally (no remote), `provenance init`, brief + blueprints, placeholder `/healthz`, fake PDS skeleton | `npm test` green, lint 0 errors |
| M1 | `shared` (tokens, JWKS, SQLite, config) then `auth` (OAuth client, callback, `GET /session`, sign-out, user table, cookie session). Device flow endpoints are part of this service but arrive with prov-cli's M2 | fake-PDS suite passes; live: a Bluesky test account signs in and `GET /session` returns its DID and handle; a second handle on another PDS also works |
| M3 | `records`: list, get, put version, head CAS (409 with the current head), hash verification via the corpus, `applyWrites` batching | CLI-token `curl` round trip; stale head -> 409; wrong `content_hash` rejected; live: scratch collection on the test account; size/rate check with linespec's 526 versions |
| M5 | UI: Provisio session context beside `ClerkProvider` behind a flag, `RecordSource` over the record server, edit through CAS | Cypress as in the plan; the GitHub-token path still works with the flag on |
| M8 | Provisio-handle sign-in in `auth`, after M7 (PDS) | as in the plan |

M1 first because everything else needs a session. M3 needs only M1's token verifier, so records and the
CLI's M4 can start once M1's `shared` package is in. M5 needs M1 and M3 and the `provisio.dev` host for
the cookie. M2 (prov-cli) must be tested against the real M1 service, so I'll publish the device-flow
contract (request/response shapes) as soon as M1 starts.

## 7. What I need decided

1. Approve one repo `provisio-backend` (workspaces: `shared`, `auth`, `records`, `api`), created local-only,
   remote created by Caleb.
2. Hashing reuse: vendored copy with a drift test (recommended), or GitHub Packages.
3. Handwritten TypeScript now, backlog items later (recommended), or a TS generator proposal first.
4. Node's built-in `node:sqlite` plus Hono plus `node:test` as the stack (recommended), subject to the
   checks below.
5. Who creates the Bluesky test account for the live suite (Caleb), and the `provisio.dev` host for the UI
   cookie.

## 8. Assumptions to verify (each before it is relied on)

| Assumption | Check | When |
| --- | --- | --- |
| Miren's Node buildpack runs a Node 24 app (and `node:sqlite`, `--experimental-strip-types` or a Dockerfile fallback) | placeholder deploy | step 0, needs approval |
| `node:sqlite` has an online backup API on the Node version used | read the Node docs, test locally | step 0 |
| `@atproto/oauth-client-node` supports confidential clients with `private_key_jwt` and the scopes we need | read its README/tests; Bluesky test account | M1 |
| Bluesky-hosted accounts accept `dev.provisio.*` writes via OAuth | scratch repo | M1 |
| `age` encryption is available as a maintained JS package | check before M1 ops step | M1 |
| `applyWrites` limits fit 526 versions of up to 27 KB | scratch account | M3 |
| Vercel custom domain on `provisio.dev` gives the UI a first-party cookie | try on a preview | M5 |
