# Versioned Provenance on Provisio — Design

Oct 15, 2026 · @Caleb Cowen

## Summary

Provenance records become versioned documents with a stable id, and every piece of code generated from one remembers the version it came from. When a record changes, `sedum grow` regenerates only that record's code; when it hasn't, `grow` does nothing. Records live in git and publish, on request, to the author's AT Proto repo under a Provisio identity, which is what lets Provisio show where software came from and where it went.

**Goals**

- **Selective regeneration.** Editing a record regenerates the regions it produced and nothing else, borrowing Phoenix's idea of tracing a spec change down to its code.
- **Versions, not supersession.** A record keeps its id across changes. Its history and the reason for each change stay readable.
- **One account everywhere.** People sign in to Scion, Provisio and the CLI with the same Provisio account and manage their records in Provisio.
- **Shared-software lineage.** When an app built on Scion is copied, both sides can see the origin and how each copy evolved.

**Decided so far**

| Question | Decision |
| --- | --- |
| Where the canonical record lives | Files in the repo for now, behind a record-store interface so it can move. History lives in the store's own version snapshots, not in git. The CLI publishes versions to the user's AT Proto repo. |
| Visibility | Publishing is explicit and opt-in, because AT Proto records are public. |
| Identity | Every Provisio and Scion session comes from AT Proto sign-in through a Provisio auth service; Clerk is retired. Sign in with a Provisio handle, any AT Proto account such as Bluesky, or GitHub or GitLab. |
| Who runs regeneration | Sedum. `grow` skips unchanged records; `--force` regenerates anyway. |
| Record types | Imprints stay separate records, because sedum needs small units of implementation. |
| Commit links | Tags name the version: `[prov-…@3]`. |
| Products and CLIs | Provenance gets its own repo and provenance CLI, ported from linespec on a shared record library. linespec stays the test tool. One Provisio login covers every product CLI. |

**Not in scope yet:** linespecs as AT Proto records, and private records (AT Proto's private data feature is still alpha).

## Where things stand today

None of the five repos contains any AT Proto code, and no record or marker carries a version today. What each one gives us to build on:

| Repo | What exists | What matters for this design |
| --- | --- | --- |
| linespec | About 520 YAML records in `provenance/`, one file each. Types: brief, blueprint, imprint, bug. Lifecycle: draft, open, implemented, superseded, deprecated. | Implemented records are frozen in four places: the hash manifest, `check` rejecting tags on implemented records, `complete` rollbacks, and the UI's allowed transitions. `create --supersedes` writes `superseded_by` into the old record. |
| linespec | `.linespec/hash_manifest.json`: SHA-256 of the whole marshalled record. | The hash covers status and `sealed_at_sha`, so a status flip changes it. Unusable as a regeneration key. |
| linespec | `publish` makes content-addressed, immutable versions (v1, v2…) of the whole record set. | Versioning already exists, but for the set, not per record. |
| sedum | Markers wrap each region: `sedum:action:variant {tier, record, writer, kwargs, …}`. Unknown keys are preserved. | `record` is the last record that wrote the region. No hash is stored. Region identity is action + selecting kwargs, not the record. |
| sedum | The model picks `{action, kwargs}` from a closed catalog; templates render the code deterministically. | Two separate causes of change: the record (needs a model call) and the package templates (re-render only). |
| sedum | No prune: a region whose kwargs change leaves the old one behind (`OPEN_QUESTIONS.md` §16). | Must be solved for regeneration to be safe. |
| sedum | `TOOL_BOUNDARIES.md` puts any convergence loop above sedum. | This design moves the skip-or-regenerate decision into `grow`, so the boundary doc needs amending. |
| scion | A "Provisio account" is a user in a Clerk instance shared with the Provisio UI. | Identities are Clerk `user_…` ids. No DIDs. |
| scion | Copying an app writes `provenance/app-lineage.yml` with ancestor app, author and date. | Lineage exists but is attribution only, keyed by Clerk ids and box names. |
| linespec-provenance-ui | The Provisio web app. Reads records straight from GitHub/GitLab with OAuth tokens Clerk holds; edits go through PRs. | Fits a git-first design. Its status type lacks superseded and deprecated. |
| phoenix | Spec → clauses → canonical nodes → implementation units → code, with content-hash ids and a hash-chained journal. | Content-hash ids changed on every rewording, so it added a stable name key. Regeneration overwrites hand edits and only reports them afterwards. |

## Core principle: everything derived pins the version it came from

A record has a stable id and a content hash per version. Anything built from a record stores the hash of the version it was built from. It is stale exactly when that stored hash differs from the record's current hash. One rule then covers regeneration, review and lineage:

| Derived thing | Pins | Stale means |
| --- | --- | --- |
| Sedum region (marker) | record id + `gen_hash` | `grow` re-selects and regenerates that record's regions |
| Blueprint or imprint | its parent's version (brief or blueprint) | its parent, or any ancestor, moved since this was implemented |
| Linespec | the record version it verifies | the spec needs rerunning or review |
| Scion copy | each source record's version at copy time | upstream evolved since you copied |

Two consequences follow:

- **Ids stay put.** Phoenix found that ids derived from content churn on every rewording. Identity comes from the stable id; the hash only says which version.
- **Hash only what drives the output.** A status change must not look like a content change. Each consumer hashes the fields it actually reads (see `gen_hash` below).

## Record model v2

A record stays one YAML file at a fixed path, edited in place. An append-only snapshot store, not git, is the version log; the file itself carries the current version number and the reason for the latest change.

**Version log without git.** Every version bump writes an immutable snapshot named by its content hash: `{version, content, reason, author, created_at, parent_hash}`. The files backend keeps snapshots at `provenance/.versions/<id>/<version>-<hash>.yml`. The Provisio backend keeps the same shape as `dev.provisio.provenance.version` records. So publishing copies snapshots, and `history` and `diff` read snapshots instead of `git log`. Where git is present it only carries the files, and commit tags become an optional cross-check.

**Field changes**

| Field | Change | Why |
| --- | --- | --- |
| `version` | New. Integer, starts at 1, bumped when content fields change. | Readable version without hashing anything. |
| `reason` | New. Required on every bump. | The "why" of a change survives in the file, not only in a commit message. |
| `supersedes`, `superseded_by` | Removed. | Versions replace supersession. |
| `aliases` | New. Old ids folded into this record during migration. | Old `[prov-…]` commit tags still resolve. |
| `implements` | Keeps the target id; gains `implements_version`. | An imprint pins the blueprint version it realized. |
| `status` | Applies to the current version. Editing an implemented record bumps the version and returns it to open. | Lifecycle continues per version instead of freezing forever. |

**Two hashes, never stored in the record**

- `gen_hash`: SHA-256 over a canonical encoding of exactly what sedum reads: `intent`, `constraints`, `affected_scope`, `forbidden_scope`. Canonical means sorted keys, UTF-8 NFC, trimmed and collapsed whitespace, list order kept. This is what markers store.
- `content_hash`: the same encoding over all content fields, excluding lifecycle fields (`status`, `sealed_at_sha`, `version`, `reason`). It also folds in the content hash of each associated integration spec (see Associated specs). This is what child records, linespecs and copies pin.

Both are computed, never written into the file. A shared conformance corpus, like sedum's `conformance/markers/cases.json`, keeps linespec, sedum and the Provisio backend producing identical hashes.

**Commit tags carry the version: `[prov-2026-1a2b3c4d@3]`.** A bare `[prov-…]` tag would still resolve to the record's version at that commit, but the explicit form survives rebases, squashes and records moved between repos. Markers already link code to a record version without git (see Sedum below), so commit tags become a cross-check rather than the only link.

**Migration**

1. An imprint superseded by an imprint with the same `implements` becomes earlier versions of one record.
2. A blueprint superseded by a blueprint becomes earlier versions of the newer one; the old id goes into `aliases`.
3. A bug that supersedes a blueprint becomes a version of that blueprint, with the bug's intent as `reason`.
4. A bug that `extends` a record also becomes a new version of it, with the bug's intent as `reason`. Every bug is then a version, so the `bug` type retires; a `kind: fix` on the version's reason keeps fixes visible. A deprecated record keeps its history and gets `deprecated` as a terminal status: it takes no new versions, and its regions become orphans for `--prune-orphans`.
5. The hash manifest becomes a version manifest: `{id, version, content_hash}` per line.

**Changes to the ported checks:** relax the four immutability checks; retire PROV006, PROV007, PROV014 and PROV020; add rules that a content change must bump `version` and that a bump needs a `reason`.

**CLI commands for versioning.** Changing content goes through commands that bump the version, record the reason and reopen the record in one step. `provenance` is the Provenance product's CLI (see Products, repos and CLIs).

| Command | What it does | Bumps version / reopens |
| --- | --- | --- |
| `provenance revise <id> --reason "…"` | Opens the content fields in `$EDITOR`; saves a new version only if `content_hash` changed | yes / yes |
| `provenance edit <id> --intent … --add-constraint … --remove-constraint <n> --add-scope … --remove-scope … --reason "…"` | Field-level changes without an editor, for scripts and agents | yes / yes |
| `provenance repin <id> [--to <version>] --reason "…"` | Confirms a child is still correct against its parent's newer version; moves only `implements_version` | no / no |
| `provenance history <id>` | Lists versions with reasons and commits (read from the version snapshots) | — |
| `provenance diff <id> [<vA> <vB>]` | Shows what changed between two versions | — |
| `provenance stale` | Lists records behind a parent or ancestor, direct or transitive | — |
| `provenance complete <id>` | Seals the current version, as today | — |

Hand-editing the YAML stays possible. Lint flags a content change without a version bump and points to `revise`.

**When a parent changes.** Imprints stay separate records, because sedum needs small units of implementation. Each child pins its parent's version (`implements_version`), so a parent version bump makes its children visibly out of date:

1. A brief or blueprint gets a new version with a `reason`.
2. The provenance system flags every child whose `implements_version` is now behind, in the CLI (`status`, `check`) and in Provisio. Staleness flows down the tree: a brief change flags its blueprints directly and their imprints transitively.
3. A person, or a rewriting agent, updates each child. Either the child's content changes, which bumps its version, or it is confirmed as still correct and only its pin moves. Work goes top-down: a blueprint that is only re-pinned clears its imprints' transitive flag, while a blueprint that gets a new version makes them directly stale.
4. `sedum grow` regenerates only children whose `gen_hash` changed. A re-pin alone regenerates nothing, because `implements_version` is not part of `gen_hash`.

Step 3 is the AI product: given the parent's diff and `reason`, propose new child versions for review, then trigger regeneration. Pins make it possible, since the system always knows which children are behind and by how much.

## Sedum: hash-aware grow

`sedum grow` becomes safe to run on every record, every time: unchanged records cost nothing, changed ones regenerate only their own regions. All the state needed lives in the markers, so sedum still keeps no manifest or lockfile.

**Marker additions** (first-class attributes, since sedum itself reads them):

```
# sedum:createControllerMethod:index {"tier":"owned","record":"prov-2026-1a2b3c4d","gen":"sha256:…","pkg":"sha256:…","kwargs":{"controller":"users"}}
```

- `gen`: the record's `gen_hash` when this region was written.
- `pkg`: a hash of the action definition and its templates when this region was rendered.

**What `grow` does per record**

| Record `gen_hash` | Package `pkg` | Action | Model call |
| --- | --- | --- | --- |
| matches every region | matches | skip | no |
| matches | differs | re-render each region from the action + kwargs already in its marker | no |
| differs, or `--force` | any | re-select, then reconcile regions | yes |
| record has no regions yet | — | normal grow | yes |

Re-rendering needs no recording: a marker already holds its action and kwargs, so it is its own one-line recording.

**Reconciling after re-selection.** Sedum compares the new selections with the record's existing owned regions by identity: same identity is replaced, a new one is inserted, and an owned region of this record that was not selected again is pruned. Pruning only happens once the record's whole selection has validated; a failed run leaves every region alone. Seeded regions are never touched, but `grow` reports any seeded region whose record is stale so a person can look.

**Region ownership.** A region belongs to the record named on its marker, and only that record ever rewrites or prunes it. One record overwriting another's region would mean someone wrote a new record where they should have revised an existing one, so sedum treats it as an error, not an update.

- **Owner.** The marker's `record` attribute now means the owner, not the last writer. It is set at first write and changes only through an explicit handover.
- **Who grows.** Only imprints grow regions; blueprints and briefs describe. Ownership conflicts are therefore always between imprints.
- **Identity stays the same.** A region's identity is its action, its variant, and the kwargs that select it (the action's `identity:` list, or its required kwargs). The record stays out of identity, so two records can never silently produce two copies of one definition.

What `grow` does with each region it meets, for the record being grown:

| Region | Action |
| --- | --- |
| Owned by another record | Leave it alone; list it in the selection prompt as already present |
| Owned by this record, `gen` matches | Current: skip |
| Owned by this record, `gen` differs | Stale: regenerate |
| New identity at a shared anchor | Insert a new region owned by this record, beside the others |
| Same identity as a region another record owns | Error naming both records: revise the owner, or change this record's selection |

**Steering the model away from conflicts.** For each file a record may write to, the selection prompt lists the regions other records already own there, as action, selecting kwargs and owner, marked "already present, do not select". The hash stays out of the prompt; sedum uses it to decide current versus stale.

**Handovers** are explicit, through an `adopts` list on the receiving imprint naming region identities. They cover splits and merges from the parent-change cascade. A retired record's regions become orphans for `--prune-orphans`.

**Order at a shared anchor** is fixed by owner creation time, then id, so regenerating from scratch reproduces the same file byte for byte.

**What changes in sedum:** `CheckDuplicatePaths` (`internal/record/record.go`), which rejects any file named by two records, moves down to the region level. `inject.go` already catches two records claiming one identity within a single run; that check extends to regions already on disk. linespec's scope-overlap rule (PROV015) becomes a warning, since region ownership is what protects the code.

**New flags**

- `--force`: re-select even when hashes match.
- `--check`: report stale records and regions, write nothing, exit non-zero if any are stale. Suitable for CI and pre-commit.
- `--prune-orphans`: remove owned regions whose record no longer exists.

**Known edge cases**

- **A record that grows no regions has nowhere to keep its hash.** This case is narrow: creating files from templates (Phase 3) never calls the model, and a record with nothing sedum could inject is already skipped without a model call (`selectAll`). What remains is a record the model chose to inject nothing for, or whose regions were all pruned. Decided: accept a model call for these on each `grow`, and have `grow --check` report them as "unknown" rather than stale. Revisit if they turn out to be common.
- **Re-selection is not deterministic.** The model may change kwargs on regions that did not need it. Mitigation: pass prior selections (rebuilt from markers) as context, and show a diff with `--dry-run`.

**Boundary change.** `TOOL_BOUNDARIES.md` currently puts the convergence loop above sedum. Skip-or-regenerate per record moves into `grow`; repair loops, escalation and test-driven passes can stay above it. That amendment needs its own record, as does replacing prov-2026-36c8a99c's "the record is the last writer" rule with region ownership.

## Associated specs

Associated specs come in two kinds that fit versioning differently. Human-written integration specs are inputs, so editing one versions the record that owns it. Generated unit specs are outputs, so they are just regions. Sedum does not read `associated_specs` at all today.

**Decided for now:** spec text does not feed sedum's selection prompt, so specs stay out of `gen_hash`. Only blueprints own integration specs, so a spec edit never cascades down from a brief.

|  | Integration specs | Generated unit specs |
| --- | --- | --- |
| Written by | people (a model may draft; a person submits) | sedum, from templates |
| Lives | out of tree, attached to blueprints only, for now (`OPEN_QUESTIONS.md` §8) | in tree, inside marked regions |
| Model access | read only | owned and regenerated |
| Role in versioning | input: its content hash is part of the owning record's `content_hash` | output: covered by the region's `gen` and `pkg` hashes |
| An edit means | a new version of the owning record, with a `reason`; children are flagged and `grow` regenerates | regenerated with its region, like any other code |

**Integration specs drive the north star.** In §14 of sedum's `OPEN_QUESTIONS.md`, people change behaviour by editing specs, and the code follows. Folding each spec's content hash into its owner's `content_hash` makes that work today, without making linespecs separate records: a spec edit triggers exactly the cascade a record edit does. `provenance revise` should detect a changed spec and ask for the `reason`.

**Versioning removes the update-record problem for generated specs.** `OPEN_QUESTIONS.md` §6 says an update record describes a delta ("also include addresses"), so regenerating a spec from it would drop every existing assertion. With versions there are no delta records: the imprint is revised, its intent stays complete, and regenerating its spec from that intent is sound. This holds only if authors keep intents complete. `revise` helps by opening the full current intent rather than a blank delta.

**Verification evidence belongs to a version.** Today `complete` runs specs once and keeps no result. Following Phoenix's evidence records, each run stores:

- the record's `content_hash` (which version was checked)
- each spec's content hash (what it was checked against)
- a code hash over the `gen` and `pkg` values of the record's regions (what code ran, without needing git)
- pass or fail, and when

Evidence goes stale when any of the three hashes changes, so Provisio can show "v4 verified, v5 unverified". Evidence publishes alongside version records.

**Markers make spec runs exact.** The `overlap_specs_on_complete` setting finds affected specs by glob-matching changed files against scope patterns. After `grow`, markers say exactly which records' regions changed, so exactly those records' specs run. Failure traceback (§9) gains one more step: a failing region's `gen` hash names the record version that produced it, and that version's `reason`.

**Specs must survive publishing and copying.**

- **References:** `associated_specs` names specs by repo path, which means nothing once a record is published or copied. Published records refer to a spec by stable id plus content hash; when linespecs become records, that is a `dev.provisio.linespec.spec` reference.
- **`run_command`:** arbitrary shell. It stays in local config and the local store, is never published, and never runs automatically from a record that arrived by copy or by pulling an upstream version.

## Identity: Provisio sign-in on AT Proto, Clerk retired

Every Provisio user is an AT Proto account, identified by a DID, and every session in Provisio and Scion comes from AT Proto sign-in through one Provisio auth service. Clerk is removed once existing users have migrated. Provisio runs its own PDS (the server that hosts an account's records), so a Provisio handle also works in Bluesky and other AT Proto apps, just as a `bsky.social` handle does. Handles live on their own domain (for example `name.provisio.social`), separate from provisio.dev, so a handle can never collide with a service subdomain such as `api` or `scion`.

**Ways to sign in**

| Sign-in | How it works | Account lives on |
| --- | --- | --- |
| Provisio handle | AT Proto OAuth against the Provisio PDS. The default "Continue with Provisio" button. | Provisio PDS |
| Bluesky or any AT Proto account | AT Proto OAuth against the account's own server | their own server |
| New email sign-up | Creates a Provisio handle, then signs in with it | Provisio PDS |
| GitHub or GitLab | A login option on the Provisio sign-in page (see below); the account is created on first login | Provisio PDS |

**Why any AT Proto account works.** AT Proto OAuth is built for it. The app resolves the typed handle to a DID, the DID document names the account's server, and that server names its authorization server, where the user signs in. Provisio publishes one client-metadata document and needs no per-provider setup.

**GitHub and GitLab sign in through the Provisio auth service.** The [OAuth spec](https://atproto.com/specs/oauth) lets a server choose its own login methods, but the reference PDS sign-in (`@atproto/oauth-provider`) supports only username, password and email code, and its customization covers branding only. Adding GitHub to the PDS's own page would mean forking that provider or building a separate sign-in front end (Bluesky calls this an "entryway"). So the auth service accepts GitHub and GitLab OAuth itself, maps the account to the user's DID and issues the same Provisio session. Provisio and Scion can't tell the difference; outside AT Proto apps simply don't offer GitHub.

Provisio also stores GitHub and GitLab repo tokens itself, encrypted, replacing `api/clerk-repo-token.ts`.

**The Provisio auth service** (for example `auth.provisio.dev`) is the one AT Proto OAuth client for the whole product. It holds:

- the client metadata and a JWKS (its public signing keys)
- the DPoP-bound OAuth sessions
- the user table: DID, handle, linked GitHub and GitLab accounts, repo tokens

After sign-in it issues a signed Provisio session token whose subject is the DID. Scion and the Provisio UI verify that token instead of a Clerk JWT. Because both apps live under provisio.dev, one sign-in covers both.

**Scoped access.** Provisio asks only for its own collections, not the whole account. AT Proto's [permission scopes](https://atproto.com/specs/permission) allow `repo:dev.provisio.provenance.record`-style grants, or a published permission set (`include:dev.provisio.…`) bundling every `dev.provisio.*` collection. The consent screen then reads "write your provenance records", not "full account access". Servers running older software may only offer broad access, so check before relying on it.

**Writes.** The Provisio API writes to a user's repo through their AT Proto OAuth session in the auth service. A GitHub or GitLab user's background PDS account has no such session, so for those accounts the auth service generates an app password at account creation and stores it encrypted. Provisio never holds a password the user knows. To use their handle in Bluesky or other AT Proto apps, such a user claims the account by setting a password, since the PDS sign-in page is password-based.

**What changes in Scion and the Provisio UI**

| Area | Today (Clerk) | After |
| --- | --- | --- |
| Scion sign-in | Clerk `<SignIn>`/`<SignUp>`; `SessionSync` posts a Clerk JWT to `/api/auth/session` | Provisio sign-in; the same route verifies a Provisio token and still sets `scion_session` |
| Scion user id | Clerk `user_…` from `getUserId()`, in SQLite owner columns, box stamps, app grants, telemetry, `SCION_INTERNAL_USER_IDS` | the DID |
| Scion billing | Stripe customer metadata `scion_user_id` holds the Clerk id | metadata rewritten to the DID |
| Scion per-app gate | `/api/auth/gate` checks Clerk, then ownership | checks the Provisio session, then ownership |
| Provisio UI | `@clerk/react`; repo tokens via `api/clerk-repo-token.ts` | Provisio session; repo tokens from the auth service |

**Migrating off Clerk**

1. Stand up the PDS and the auth service. Create a Provisio account for every Clerk user, recording `clerk_id ↔ did` and each user's GitHub and GitLab account ids from Clerk's backend API.
2. Dual-run: Scion and the UI accept either a Clerk or a Provisio session, both resolving to the DID.
3. Rewrite stored ids from Clerk ids to DIDs: SQLite tables, Stripe metadata, lineage files, telemetry and config.
4. Switch the sign-in pages to Provisio. GitHub and GitLab users are matched by their provider account id; email users get a set-your-password email for their new handle.
5. Remove Clerk once no Clerk sessions remain.

**Portability.** A Provisio-hosted user can add their own rotation key to their `did:plc` and later move to another server without losing their identity or records.

**CLI sign-in.** A device-authorization flow against the auth service gives the CLI a Provisio API token. Later, the CLI could act as an AT Proto OAuth client itself.

**Scion lineage.** Lineage entries switch from Clerk ids to DIDs, and `ancestorGoldenId` becomes the `at://` URI of the published app.

**Operations: the PDS, auth service and indexer run on a Miren managed cluster.** They become three apps on one cluster. No real cluster exists yet (the `provisio-dev` and `provisio-prod` configs were a local test). Decided: a Miren managed cluster, which Miren provisions by hand (usually the same day in US business hours), upgrades and monitors, for $29 a month for six months and then $79 ([pricing](https://miren.dev/pricing)). Scion stays on boxd, which suits its per-app VMs but not a PDS: it doesn't document routing many handle subdomains to one machine, and its sleep defaults and bot protection get in the way of always-on services.

| Need | How Miren covers it |
| --- | --- |
| Handle subdomains with TLS | A wildcard route on the handle domain. Before certifying a new name, Miren calls `GET /tls-check?domain=<name>`, the same check the official PDS install uses, so it should plug straight in ([traffic routing](https://miren.md/traffic-routing)). |
| PDS data (SQLite and blobs) | A Miren disk with a single fixed instance; disks survive restarts and redeploys ([disks](https://miren.md/disks)) |
| Relay and Jetstream connections | Websockets stay open while traffic flows; raise the 60-second idle timeout with `miren route timeout` |
| Email | an SMTP provider set through environment variables |

**Prove it on the managed cluster before any real accounts exist:**

- [ ] The original Host header reaches the app. The PDS answers `/.well-known/atproto-did` per handle hostname, and Miren's docs don't say.
- [ ] A new handle gets a real certificate through the `/tls-check` flow.
- [ ] Scheduled `miren disk backup --cloud` runs. Miren disks are a developer preview with no documented replication.
- [ ] The PDS's PLC rotation key has a separate, secure off-site copy. Losing it means losing control of users' identities.

Step-by-step test, ready to run once the cluster is provisioned: PDS test brief

## AT Proto storage and publishing

Published records live in the author's AT Proto repo as a mutable head plus an append-only log of version records, because a PDS keeps only the latest version of anything. Git stays canonical; publishing copies versions out of git history.

**Lexicons** (namespace `dev.provisio.*`, under the provisio.dev domain)

| Collection | Record key | Mutable | Holds |
| --- | --- | --- | --- |
| `dev.provisio.provenance.record` | the prov id | yes (head) | current content, `version`, `content_hash`, `gen_hash`, source repo + path + commit, a strong reference to the latest version record |
| `dev.provisio.provenance.version` | TID (time-ordered id) | never | record URI, version number, content snapshot, `reason`, commit SHA, parent version's strong reference |
| `dev.provisio.project` | TID | yes | a codebase: name, git URL, which records belong to it |
| `dev.provisio.lineage.copy` (later) | TID | never | source project and record versions, copy date, the new project |
| `dev.provisio.linespec.spec` (later) | path-derived | yes | a linespec and the record version it verifies |

A strong reference is AT Proto's `{uri, cid}` pair. Chaining each version to its parent's CID makes the log tamper-evident, like Phoenix's hash-chained journal. AT Proto's own CID can't replace `gen_hash` or `content_hash`, because it covers the whole AT record, status included.

**Why not just mutate the head?** The [repository spec](https://atproto.com/specs/repository) says updates keep no prior version and deletes leave no tombstone. Rebuilding history from the event stream through an indexer would lose versions whenever an event is missed.

**Publishing**

1. The user opts a project in, through config or per command.
2. For each record, the CLI reads its local version snapshots and finds those not yet published.
3. Through the Provisio API, it writes each missing version record oldest first, then updates the head. The head update is a compare-and-swap on the head's current CID (`swapRecord`), so concurrent publishers can't silently overwrite each other.
4. Re-running publish is a no-op when nothing changed.

**Reading across users.** Provisio runs an indexer (an AppView) that consumes Jetstream, AT Proto's filtered event stream, for `dev.provisio.*` collections and stores them in Postgres. That powers per-user record lists, version history, backlinks (what implements, pins or copied this) and search. The Provisio UI reads published records from the indexer and unpublished ones from git, as it does today.

**Privacy.** Everything published is public. AT Proto's private data feature, now called [Spaces](https://atproto.com/blog/atproto-spaces-alpha), entered alpha in August 2026; it provides access control, not encryption, and promises no stability yet. Revisit it once it ships.

## Lineage: shared software you can trace

A copy becomes a public, pinned reference from the new project to the exact record versions it started from, so both authors can see where software came from and how each side has since moved.

**On copy, Scion:**

1. Publishes the source app's project and records, if the author allowed it when they shared the app.
2. Creates the new project in the copier's repo.
3. Writes a `dev.provisio.lineage.copy` record in the copier's repo pinning the source project and each source record version by strong reference.
4. Keeps `provenance/app-lineage.yml` in the app, now with DIDs and `at://` URIs instead of Clerk ids and box names.

**What Provisio can then show:**

- **Upstream drift:** the source record is at v7, and you copied v4. Same pin comparison as everywhere else.
- **Downstream reach:** everyone who copied your app, found through the indexer's backlinks. A copy record lives in the copier's repo, so only an indexer can answer "who copied me".
- **Divergence:** for a pinned record, whether the copy has since versioned it on its own.

**Pulling changes forward (later).** Because both sides version the same record ids, a copier could take an upstream record's new version, and `sedum grow` would regenerate just the affected regions. That turns lineage from attribution into a way of sharing updates.

**Copies keep record ids; the author's DID tells them apart.** On AT Proto a record's address includes the repo it lives in, so the same id under two DIDs is two records:

```
at://did:plc:original/dev.provisio.provenance.record/prov-2026-1a2b3c4d   ← original author's
at://did:plc:copier/dev.provisio.provenance.record/prov-2026-1a2b3c4d     ← the copy
```

The copier cannot write to the original's repo, so the original author's versions are preserved automatically, and ownership traces back through the id rather than the version.

- **Two kinds of ownership.** A record owns its regions (region owner); a DID owns its records (record authority). Locally there is one record store, so a marker's bare `prov-…` id always means the local record; in a copied app, that is the copier's.
- **Origin pin.** At copy time each record gets `origin`: the source record's address plus the CID of the version copied. A copy of a copy pins its direct parent, so full ancestry comes from walking `origin` pins.
- **Version numbers are per fork.** v5 in the original and v5 in the copy are different things, so forks are compared by content hash, never by number, much as git forks share history up to the fork point.
- **No regeneration on copy.** The content is identical, so every marker's `gen` hash still matches.
- **Drift reads from both sides,** for example "upstream is at v7, you forked from v4, and you've made 2 versions of your own".
- **Pulling upstream (later)** is a three-way merge of record content: the pinned origin version is the base, then yours, then theirs. After the merge, `grow` regenerates what changed.

## Products, repos and CLIs

Provenance becomes its own product, with its own repo and `provenance` CLI ported from linespec, while linespec keeps the test DSL and runner. Every product CLI shares one Provisio login, so no CLI ever has to merge into another.

| Product | Repo | CLI | Owns |
| --- | --- | --- | --- |
| Provisio (the company) | backend services | none needed | the account, auth service, PDS and the provisio.dev hub |
| Provenance | new `provenance` repo | `provenance` | records, versions, hashing, scope enforcement, hooks, publishing |
| linespec | `linespec` | `linespec` | the test DSL and runner; drops its `provenance` subcommands after a deprecation window |
| sedum | `sedum` | `sedum` | generation; imports the record library for `gen_hash` |
| scion | `scion` | none | the app builder; its in-app agents call `provenance` |

**The shared record library.** A standalone Go module inside the Provenance repo is the one definition of a record. It contains:

- the schema and a YAML codec that keeps formatting on save (today's `SaveRecord`)
- version snapshots and the record-store interface, with a files backend now and a Provisio API backend later
- canonical encoding for `gen_hash` and `content_hash`
- scope patterns and matching
- lint rules
- the conformance corpus: JSON cases that the Go module and the TypeScript backend must both pass

The `provenance` CLI and sedum import the module; sedum replaces its own lenient record decoder with it. The Provisio backend stays consistent through the corpus rather than the module.

**Port, don't rewrite.** `pkg/provenance` is about 11,800 lines and imports only `pkg/config` and `pkg/embeddings`, so most of it moves as is.

| Carries over unchanged | Changes |
| --- | --- |
| scope patterns and file write-restriction (`reconcile`) | supersession becomes versions (see Record model v2) |
| git hooks and their installer | the hash manifest becomes version snapshots |
| formatting-preserving record saves | the four immutability checks and retired lint codes |
| SARIF output, `graph`, `context`, `next`, `govern`, `search` | `publish`, `clone`, `import` (whole-set manifests) give way to AT Proto publishing |
| shared-repo cache and embeddings | the `provenance:` block in `.linespec.yml` moves to `.provenance.yml` |

Running associated specs stays a shell call to `linespec test`, so Provenance never links against linespec.

**One Provisio login for every CLI.** A small Go auth module is shared by every product CLI:

- `login`, `logout` and `whoami` in each CLI are thin wrappers over it.
- Sign-in is a device-authorization flow against the Provisio auth service.
- One credential file (for example `~/.provisio/credentials`) holds the token and refreshes it, so signing in through any CLI signs in all of them.
- CI uses a non-interactive token from an environment variable such as `PROVISIO_TOKEN`.

A separate account CLI, like `gcloud auth`, can come later if it earns its place.

**Migrating from `linespec provenance`.** The move happens in the same release as versioning, so every repo migrates once:

1. Release `provenance` v1 with the versioned model and `provenance migrate`. It converts supersession chains to versions, writes snapshots, keeps old ids as `aliases` so old commit tags still resolve, moves config to `.provenance.yml` and reinstalls hooks. Until a repo is migrated, `provenance` reads the old `.linespec.yml` block.
2. In the same release, `linespec provenance …` forwards to `provenance` with a deprecation warning, or explains how to install it.
3. Switch every dependent in that release:
   - Scion's guest provenance skill, template golden and hook install
   - the Claude Code provenance skill, plugin and SessionStart hook
   - linespec's provenance CI workflows (SARIF lint, indexing)
   - the record type mirror in the provenance UI
4. Migrate the active repos first (linespec, sedum, scion, the provenance UI), then the rest. 27 local repos carry a `.linespec.yml` today.
5. After the deprecation window, remove `pkg/provenance` and the forwarding from linespec.

**Open:** `discover` does two jobs: it drafts blueprints (Provenance's job) and skeleton `.linespec` files (linespec's). It probably splits, with `provenance discover` calling linespec to stub the specs.

## Phased plan

Phases 1–3 deliver selective regeneration with no backend at all; the AT Proto track can start alongside them and needs only the record model's schema and hashes.

&#91;embedded content: phases · local loop and backend tracks\]

Arrows show what each phase needs first. Phase 1 feeds both tracks.

1. **Record model v2.** Version fields, `reason`, `implements_version`, removal of supersession, migration of existing chains, integration-spec hashes folded into the content hash, and a shared hashing library with a conformance corpus. Done when linespec and sedum compute identical hashes for every record in the three repos.
2. **Hash-aware grow.** `gen` and `pkg` in markers, skip/re-render/re-select, prune, `--force`, `--check`, and the `TOOL_BOUNDARIES.md` amendment. Done when a second `grow` over unchanged records makes no model calls and edits no files.
3. **Stale-child alerts.** `status` and `check` report children behind their parent's version. Verification evidence is stored per version. Done when bumping a brief lists exactly its blueprints and their imprints, and their evidence shows as stale.
4. **Provisio identity.** PDS hosting, AT Proto OAuth sign-in for Provisio and other handles, the Provisio auth service, GitHub and GitLab login, and the migration off Clerk, starting from the `clerk_user_id ↔ did` mapping. Done when a Provisio handle, a Bluesky handle and a GitHub login each land on one linked Provisio account, in both Scion and Provisio, with Clerk removed.
5. **Publish.** Lexicons, the Provisio write API, CLI sign-in, and opt-in publishing of version records and heads. Done when re-publishing an unchanged project writes nothing.
6. **Indexer and Provisio UI.** A Jetstream consumer into Postgres; the UI shows published records, version history and backlinks.
7. **Scion lineage.** Copy records with DIDs and pinned record versions; Provisio shows upstream drift and downstream reach.

## Risks and open questions

**Moving away from git.** Git does five jobs today: version log, review (PRs), linking code to records (commit tags), sealing (`sealed_at_sha`), and offline/CI access. This design already takes over the first three:

| Job | Today | After this design |
| --- | --- | --- |
| Version log | git history of the YAML file | `version` + `reason` in the record, version snapshots in the record store, locally and on AT Proto |
| Code ↔ record link | `[prov-…]` commit tags | markers carry record id + `gen_hash`; commit tags become a cross-check |
| Review | PRs on the record file | could move into Provisio |
| Sealing | `sealed_at_sha` | still git |
| Offline and CI | files on disk | still git, or a local cache of the Provisio store |

So git-first can be a starting point rather than a lock-in. Build the CLI against a record-store interface from day one, with a files backend now and a Provisio API backend later; then a project can switch its canonical store without changing anything above that interface.

**Risks**

- **Key custody.** Like any PDS host, Provisio holds the repo signing keys for every hosted account, and the auth service holds everyone's OAuth sessions. A breach could rewrite users' public records. User-held rotation keys limit the damage.
- **Hash drift between implementations.** Go (linespec, sedum) and TypeScript (Provisio) must compute identical `gen_hash` and `content_hash` values, or everything looks stale. The conformance corpus is the guard.
- **Re-selection churn.** A changed record can shift kwargs on regions whose meaning didn't change. Prior selections as context and `--dry-run` diffs reduce it but don't eliminate it.
- **New infrastructure.** A PDS, an auth service and an indexer are three new services though all three run as apps on one Miren cluster, which has to be set up first.
- **Accidental publishing.** Opt-in publishing is the guard; the CLI should show exactly what will become public before the first publish.

**Open questions**

- [x] Do bugs that `extends` a record, and deprecated records, survive versioning, or become versions and a terminal status? Decided: versions and a terminal status (see Migration).
- [x] Decided (accept a model call on each grow; see Sedum): how does a record that only creates files remember its `gen_hash`?
- [x] Can GitHub and GitLab become login methods on the PDS's own sign-in page, or do they live in the auth service? Decided: the auth service, since the reference PDS sign-in is password-only.
- [x] Where does the PDS run: boxd, a VPS, or a managed host? Decided: Miren, after the checks under Operations pass on a test cluster.
- [x] Does the parent-change rewriting agent run locally, in Provisio, or both, and who pays for its model calls? We won't worry about this right now that's a later product idea.
- [x] Which domain do handles live on (for example provisio.social)? provisio.social
- [x] A Miren managed cluster, or Miren installed on a VPS? Decided: managed cluster.
