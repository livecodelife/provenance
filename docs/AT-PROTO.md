# Publish your provenance records to Provisio

This guide takes a git repository that already keeps provenance records and publishes them to Provisio, where
they are stored on AT Proto under your account. It covers sign-in, a first publish, keeping records current, CI,
and what to do when something goes wrong. How it works inside is in the [appendix](#appendix-how-it-works).

## 1. What you get, and what becomes public

Publishing copies your records, version by version, from git to your Provisio account. Then:

- the Provisio UI shows them, with their history, outside the repository;
- other people can read them, link to them, and (later) see who built on them;
- versions made in the UI can be brought back into git with `provenance import`.

**Everything you publish is public.** Anyone can read it, and it may be copied or indexed; deleting it later
does not guarantee it disappears. `provenance publish` always lists exactly what will become public and asks
before it writes. Git stays the source of truth: publishing never changes your repository.

Local commands (`create`, `revise`, `lint`, `check`, `stale` and the rest) never use the network and never need
an account. You only need one for `publish`, `import` and `clone`.

## 2. Before you start

- `provenance` installed, and a repository with provenance records (`provenance init`, then `create`/`open`).
  Install: `curl -fsSL https://raw.githubusercontent.com/livecodelife/provenance/main/install.sh | sh`.
- A Provisio account. You can sign in with a Provisio handle, or with a Bluesky or other AT Proto account
  (see section 3).
- Records worth publishing. A draft that has never been opened has no version yet and is skipped; run
  `provenance open --record <id>` first.
- A git remote named `origin`. Published records carry where they live in git; without a remote `publish` works
  but says `no usable origin remote: the records are published without a source` and leaves the source out.

## 3. Sign in

```
provenance login
```

It prints a verification URL and a short code, opens your browser, and waits until you approve the sign-in in
the browser. On a server or over SSH, use `provenance login --no-browser`, open the URL on any device, and enter
the code there. The approval page names the CLI and the machine, so check them before approving.

**Which account to use.** On the Provisio sign-in page choose "Continue with Provisio" for a Provisio handle, or
enter a Bluesky or other AT Proto handle to sign in with that account; the records are then stored under that
account. (A `provenance login --provider provisio|bluesky` shortcut is planned; until it lands, choose the
provider on the page.)

Check it worked:

```
provenance whoami          # handle, DID, where the credential comes from
provenance logout          # signs out and deletes the local credential
```

Your credential is kept in the OS keychain where there is one, else in `~/.provisio/credentials` (readable by
you only). Set `PROVISIO_TOKEN` to use a token instead; it wins over any stored credential (section 7).

If you cannot open a browser at all, create a personal token (section 7) and run
`echo "$TOKEN" | provenance login --with-token`.

## 4. Choose a project

A project groups a repository's records under one address in your account. Pick a name and build the address
from your DID (shown by `provenance whoami`):

```
at://<your DID>/dev.provisio.project/<name>
# for example
at://did:plc:abc123/dev.provisio.project/my-service
```

The name uses letters, digits and `. _ ~ : -`. Use one project per repository. Pass it to `publish` with
`--project`; the CLI checks that it is a well-formed AT URI and rejects placeholders such as `<your DID>`.

The service does not create the project for you and does not look it up: it only checks that the address is
a well-formed AT URI. So you can publish with an address that has no project record behind it, and the records
are filed under it. `provenance project create` and `provenance project list` are planned (they wait on a
backend route); until they exist, no project record is made for you.

## 5. First publish

Look before you write. `--plan` lists what would be published and writes nothing:

```
provenance publish --all --plan --project at://did:plc:abc123/dev.provisio.project/my-service
```

The summary line reads `N unpublished, N local ahead, N up to date, N server ahead, N diverged, N not
publishable, N skipped`. Anything in `diverged`, `server ahead` or `skipped` needs a look before you continue
(section 9). When the plan is what you expect:

```
provenance publish --all --project at://did:plc:abc123/dev.provisio.project/my-service
```

It prints the records and versions that will become public and asks `Publish these? [y/N]`. Answer `y`. It then
writes each record's versions oldest first and prints `published <id> v<N>` for each, ending with
`✓ Published N write(s).` To publish only some records, name them instead of `--all`:
`provenance publish prov-2026-1a2b3c4d`. In scripts, add `--yes` to skip the question.

Run it again and it does nothing: `Nothing to publish: the server already has every version and status.`

## 6. Keep it current

Publish again whenever records change. It sends only what is new:

- a `revise` or `edit` adds a version, which is published;
- a status change (open, implemented, deprecated) updates the record's status without a new version;
- `repin` updates which parent version the record was confirmed against, without a new version.

Each published version points at the git commit tagged `[<id>@<version>]`, so keep the full history (a shallow
clone cannot find them). To bring changes made in the UI back into git:

```
provenance import --all
```

Imported versions become new versions in your repository, with their reason and author. `import` refuses a
record whose history differs from the server's at a version both sides have, and says to `revise` it locally;
it never merges.

## 7. Publish from CI

Use a personal token, not your own login. A token is scoped, expires, and can be revoked without touching your
account.

1. **Create a token** with scope `records:write` and a 90-day expiry (the maximum is 365 days). Tokens start
   with `pvt_` and are shown once. Creating one needs a browser or device sign-in session; a token cannot create
   another token. There is no tokens screen in the UI yet. `provenance token create`, `list` and `revoke` are
   coming with the token PR (defaults: name `ci-<repo>`, scope `records:write`, 90 days; `--secret-only` prints
   just the token, for piping to `gh secret set PROVISIO_TOKEN`). Until it merges, mint one with
   `POST /tokens` on the auth service while signed in.
2. **Store it** as a secret named `PROVISIO_TOKEN`, preferably in a protected GitHub Environment limited to your
   main branch, and store the project address as a variable `PROVISIO_PROJECT`.
3. **Add the workflow.** The one this repository uses is a good template: `.github/workflows/publish.yml`.
   It runs only on pushes to main (never on forks or pull requests), installs a pinned `provenance` release,
   and has two steps:

   ```yaml
   - uses: actions/checkout@v4
     with: { fetch-depth: 0, persist-credentials: false }   # full history: versions point at commits
   - name: Plan
     env: { PROVISIO_TOKEN: "${{ secrets.PROVISIO_TOKEN }}", PROVISIO_PROJECT: "${{ vars.PROVISIO_PROJECT }}" }
     run: provenance publish --all --plan --project "$PROVISIO_PROJECT"
   - name: Publish
     if: vars.PROVISIO_PUBLISH == 'true'
     env: { PROVISIO_TOKEN: "${{ secrets.PROVISIO_TOKEN }}", PROVISIO_PROJECT: "${{ vars.PROVISIO_PROJECT }}" }
     run: provenance publish --all --yes --project "$PROVISIO_PROJECT"
   ```

4. **Switch it on in two steps.** Merge the workflow with `PROVISIO_PUBLISH` unset, so every push only runs the
   plan and you can read it in the log. When the plan looks right, set the repository variable
   `PROVISIO_PUBLISH` to `true`; from then on each push to main publishes.

A red run is the signal: the plan step fails on any problem (no session, a server error, a record skipped as
diverged, server ahead or invalid), so nothing is silently dropped. Concurrent runs queue rather than race.

## 8. What shows up in the UI

Signed in at Provisio, your published records appear with their full version history, reasons, status, and
what each one implements. Versions you create there are marked as made in the UI and come back with
`provenance import`. Records published under a project are grouped by it.

## 9. Troubleshooting

| You see | Why, and what to do |
| --- | --- |
| `This command needs a Provisio session. Run provenance login (or set PROVISIO_TOKEN in CI).` (exit 3) | You are not signed in, or `PROVISIO_TOKEN` is unset. Run `provenance login`, or set the token in CI. |
| `The personal token was rejected (revoked or expired). Replace PROVISIO_TOKEN, or run provenance login.` (exit 3) | The token was revoked, has expired, or was copied wrongly. Create a new one and update the secret. |
| `Not signed in.` from `whoami` (exit 1) | No stored credential. Run `provenance login`. |
| `--project "…": must start with at://…` or `looks like a placeholder` | The project address is malformed. Use your real DID: `at://did:plc:…/dev.provisio.project/<name>`. |
| `skipped … a draft with no version yet (open it first)` | The record is still a draft. `provenance open --record <id>`, then publish. |
| `note: no usable origin remote…` | The repo has no `origin`. Add it with `git remote add origin <url>`; otherwise the record is published without its source. |
| `not confirmed, nothing was published (run interactively, or pass --yes)` | `publish` asked for confirmation and there was no terminal. Pass `--yes` in scripts and CI. |
| A 409 conflict from `publish` | Someone else changed the same record's head since your plan. Run `publish --plan` again and review; nothing was overwritten. |
| `diverged` or `server ahead` in the plan | The published history differs from yours at the same version, or the UI added versions. Run `provenance import` for UI versions; for a real divergence, `revise` locally so the histories line up. |
| `import` says a version differs between the server and this repository | Both sides have the same version number with different content. `revise` the record locally, then publish. |
| Versions published but no commit shown | The repository was cloned shallow. Use `fetch-depth: 0` in CI or `git fetch --unshallow`. |
| The CI Plan step passes but nothing is published | `PROVISIO_PUBLISH` is not `true`. That is the intended first stage. |

## Open points

Planned or unconfirmed, so this guide will change:

1. **`provenance project create/list`** wait on a backend route.
2. **`provenance token create/list/revoke`** are coming with the token PR; the UI has no tokens screen yet.
3. **The UI address** for browsing published records is not stated here yet.
4. **`login --provider provisio|bluesky`** is the next CLI change after the token PR.
5. **Public `clone`.** `clone` and `import` need a session even for public records; there is no anonymous read in
   the CLI today.

---

# Appendix: how it works

For people who want the model behind the commands. Field-by-field spelling is in
[LEXICONS.md](LEXICONS.md); design reasoning is in [DESIGN.md](DESIGN.md) ("Identity", "AT Proto storage and
publishing", "Lineage") and [PLAN-IDENTITY-UI-FIRST.md](PLAN-IDENTITY-UI-FIRST.md).

**Where things live.** Git is canonical. Publishing writes to your account's repository on its AT Proto server:

| Collection | Key | Mutable | Holds |
| --- | --- | --- | --- |
| `dev.provisio.provenance.record` | the record id | yes (the "head") | current content, status, version, hashes, a reference to the newest version, optional `source` (repo, path, commit) and `project` |
| `dev.provisio.provenance.version` | TID | never | one version: content, hashes, reason, author, commit, a reference to the previous version |
| `dev.provisio.project` | name | yes | a codebase; fields are still a proposal |

A server keeps only the latest value of a record, so history lives in the immutable version records, each
chained to the previous one two ways: `parentHash` (the previous version's content hash) and `parent` (the
previous record's AT reference). A record's address includes your account's DID, so the same id in two accounts
is two different records.

**How a publish talks to the service.** The CLI does not write to your AT Proto server directly. It sends each
missing version to the Provisio record server (`PUT /records/<id>/versions`) and head-only changes such as status
(`PUT /records/<id>/head`), each naming the head it expects (`expectedHeadCid`). If the head has moved, the
server answers 409 and the CLI reports a conflict; nothing is overwritten.

**Sign-in.** The Provisio auth service is the single AT Proto OAuth client. It keeps your server sessions and
issues Provisio tokens: a 15-minute access token, a 30-day rotating refresh token, and personal tokens (`pvt_…`)
for CI. The pages you see during `provenance login` (enter code, sign in, approve or deny, finished) are served
by that service, not by the CLI, and are being restyled with the Provisio brand.

**Parents in another account.** A record can name a parent published by someone else: `provenance create
--implements at://<did>/dev.provisio.provenance.record/<id>`. The address is lineage only (it is not in the
content hash) and the pin starts unpinned, because the CLI cannot yet look the parent up. Resolving it for
`stale` and `check` is planned.

**Status of the pieces.** Built: the lexicons and generated types, the conformance hash corpus, `publish`,
`import`, `clone`, `login`/`logout`/`whoami`, the CI workflow template, `implementsUri`. Planned: `login
--provider`, parent resolution through an AppView, a publish note for cross-account parents, and the copy/lineage
record (`dev.provisio.lineage.copy`).
