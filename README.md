# provenance

Versioned provenance records: the shared record library (schema, canonical
hashing, version snapshots, record store) and the `provenance` CLI, ported
from `linespec provenance`.

Design: [docs/DESIGN.md](docs/DESIGN.md).

This repository manages itself with the `provenance` CLI it contains.

## Install

```sh
curl -fsSL https://raw.githubusercontent.com/livecodelife/provenance/main/install.sh | sh
```

The script downloads the prebuilt binary of the latest release for your
platform (Linux or macOS, amd64 or arm64), checks it against the release's
`checksums.txt`, and installs it into `$HOME/.local/bin` with no sudo and no
writes anywhere else. `PROVENANCE_BIN_DIR=<dir>` installs somewhere else and
`PROVENANCE_VERSION=<x.y.z>` installs that release. `PROVENANCE_REF=<commit,
tag or branch>` builds that ref from source instead: it needs Go
(https://go.dev/doc/install) and runs `go install` into `GOBIN` (else
`$(go env GOPATH)/bin`). Until the first release is published, or on a platform
with no prebuilt binary, the script builds a pinned commit the same way. It
prints a PATH hint when the install directory is not on your `PATH`, then prints
the follow-ups without running them: `provenance init` for a new repository,
`provenance migrate --plan` then `provenance migrate` for a repository with v1
(linespec) records, and `provenance hooks install`. `provenance version` prints
the installed version. [docs/MIGRATING.md](docs/MIGRATING.md) explains what v2
changes before you migrate.
By default the commands that talk to Provisio use the live service. To point
them somewhere else, set `PROVISIO_AUTH_URL` (the sign-in service, for example
`https://auth.provisio.dev`) and `PROVISIO_API_URL` (the record service, which
must include its version path, for example `https://api.provisio.dev/v1`; without
`/v1` every request fails with `http_404`). Do this before `provenance login`.

Homebrew is coming.

To publish your repository's records to Provisio (stored on AT Proto), follow [docs/AT-PROTO.md](docs/AT-PROTO.md).

Migrating a v1 repository: run the installer, then `provenance migrate --plan`
to read the plan (it writes nothing) and `provenance migrate` to apply it.

## Migrate from linespec v1

This section moves a repository from v1 records (managed by linespec) to v2
records (managed by `provenance`). The migration is one reviewable commit, and
you can revert it.

To let Claude Code do the work, install the plugin (see
[Install the plugin and skills](#install-the-plugin-and-skills)) and run
`/provenance:migrate`.

### Before you begin

- Install `provenance`. See [Install](#install).
- Start from a clean working tree. This command should print nothing:

  ```sh
  git status --short
  ```

- Create a branch for the migration:

  ```sh
  git switch -c provenance-v2
  ```

- If your records point at parents in other repositories, list those
  repositories under `shared_repos` in your config.

### Read the plan

1. Print the plan. This command writes nothing.

   ```sh
   provenance migrate --plan
   ```

   The plan shows how many records become how many, the version chains, and any
   warnings:

   ```
   Migration plan for my-repo: 184 records → 171

   Version chains (8), oldest version first:
     prov-2026-1a142df0  ← prov-2026-1a142df0, prov-2026-d1431737
   ```

2. Read the chains. A chain is a set of v1 records that become versions of one
   v2 record.
3. Read the warnings. A `dangling_ref` warning means a record points at a parent
   that isn't in this repository. Add the parent's repository to `shared_repos`.
4. Read the suggestions, if any. Each one is a record that the migration would
   fold, but that you might want to keep as its own record.

### Apply the migration

1. If you use `shared_repos`, fetch their records:

   ```sh
   provenance sync
   ```

2. Run the migration:

   ```sh
   provenance migrate
   ```

   By default, `migrate` folds every chain and every suggestion. To change that,
   use these flags:

   | Flag | Effect |
   | --- | --- |
   | `--accept-suggestions` | Keep every suggested record separate. |
   | `--keep-separate <id>` | Keep this record separate. Repeat for more records. |
   | `--fold <id>` | Fold this record, even when the plan suggests keeping it. Repeat for more records. |
   | `--workspace <dir>` | Look for local checkouts of shared repositories in this directory. The default is the repository's parent directory. |

   The command prints a summary:

   ```
   ✓ Migrated provenance to v2: 171 records (from 184), 184 snapshots
   ```

   It also lists the children that it re-pinned, one per line, as
   `child  parent@1 → @N`. Only their pins change.

3. Review what changed on disk:

   | Path | Change |
   | --- | --- |
   | `provenance/<id>.yml` | Each kept record is rewritten as v2. Folded records are deleted. |
   | `provenance/.versions/<id>/` | One snapshot for each v1 record. |
   | `.provenance/versions.jsonl` | The version manifest. |
   | `.provenance.yml` | Replaces the provenance block of `.linespec.yml`. |
   | `.linespec/hash_manifest.json` | Removed. |

   Migration doesn't touch your git hooks.

### Verify

1. Check the records:

   ```sh
   provenance lint
   provenance stale
   ```

   `lint` should report no errors, and `stale` should list nothing.

2. If `stale` lists a pin that is behind its parent, confirm that the record is
   still correct, then repin it:

   ```sh
   provenance repin <id> --reason "Still correct against the new parent"
   ```

   To pin to a specific parent version, add `--to <n>`.

3. Replace the linespec hooks. The old hooks are kept as `.bak` files.

   ```sh
   provenance hooks install
   ```

4. Commit the migration.
5. Check the commits that you added:

   ```sh
   provenance check --range <base>..HEAD
   ```

### Update CI and Makefile targets

1. Replace each call to linespec's `provenance` subcommand with `provenance`.
2. Install the CLI in CI with the one-liner, and pin the release so that a new
   release can't change your build:

   ```sh
   curl -fsSL https://raw.githubusercontent.com/livecodelife/provenance/main/install.sh | PROVENANCE_VERSION=<x.y.z> sh
   ```

linespec's `provenance` subcommand is being replaced by a forwarder that runs
`provenance` and prints a warning. Call `provenance` directly so that nothing
depends on the forwarder.

### Roll back

1. Revert the migration commit:

   ```sh
   git revert <migration-commit>
   ```

2. Don't mix tools afterward. v1 tooling reads only v1 records.

## Install the plugin and skills

`provenance` ships a Claude Code skill and a Claude Code plugin. Both are
optional, and both install into your repository.

| Command | Installs | Contains |
| --- | --- | --- |
| `provenance install-skills` | `.claude/skills/provenance/` | The workflow skill, which tells an agent how to create, open, edit and complete records. |
| `provenance install-plugin` | `.claude/plugins/provenance/` | Hooks that show the next action at session start, before the first edit of a file, and after a rejected commit; the `/provenance-next` command; and the `/provenance:migrate` skill. |

### Install

1. Install the skill. The command prints the files it will write, then writes
   them:

   ```sh
   provenance install-skills
   ```

2. Install the plugin:

   ```sh
   provenance install-plugin
   ```

   The output looks like this:

   ```
   Will write 8 file(s) under .claude/plugins:
       create     .claude/plugins/provenance/.claude-plugin/marketplace.json
       ...
   ✓ Installed plugin: 8 written, 0 unchanged
   ```

3. In Claude Code, add the plugin from your repository:

   ```
   /plugin marketplace add ./.claude/plugins/provenance
   ```

4. Install the plugin:

   ```
   /plugin install provenance@provenance
   ```

The plugin hooks need `jq` and `provenance` on your `PATH`. Without them, the
hooks do nothing. To install somewhere else, add `--path <dir>` (relative to the
repository root) to either command.

### Confirm that it works

1. In Claude Code, run `/plugin` and check that `provenance` is installed.
2. Type `/provenance` and check that the list includes `/provenance:migrate` and
   `/provenance-next`.
3. Start a new session in the repository. The session-start hook shows the next
   provenance action.

### Update

1. Update `provenance`.
2. Reinstall the files. Without `--force`, a command refuses to overwrite a file
   that differs.

   ```sh
   provenance install-skills --force
   provenance install-plugin --force
   ```

3. In Claude Code, refresh the plugin:

   ```
   /plugin marketplace update provenance
   ```

### Uninstall

1. In Claude Code, remove the plugin:

   ```
   /plugin uninstall provenance@provenance
   ```

2. Delete the installed files:

   ```sh
   rm -r .claude/skills/provenance .claude/plugins/provenance
   ```

## Hooks

`provenance hooks install` writes three git hooks (existing ones are kept as
`.bak`): `pre-commit` (lint staged records, run specs on completion),
`commit-msg` (`check --staged`) and `pre-push`.

The `pre-push` hook is opt-in per repository. It does nothing unless
`.provenance.yml` sets `pre_push_command`; then it runs that command with
`sh -c` in a clean `git archive` checkout of every commit being pushed, and a
failing commit blocks the push. This repository sets
`pre_push_command: go build ./...`, so every pushed commit builds. Other
repositories leave it unset and are not slowed down.

## `provenance graph`

`provenance graph [--root <id>] [--filter <status>] [--format human|json|dot]`
renders the records, their parents, pins and versions. The JSON shape is
stable; other tools read it:

```json
{
  "nodes": [
    {
      "id": "prov-2026-0000e001",
      "title": "…",
      "type": "blueprint",            // brief | blueprint | imprint
      "status": "open",               // draft | open | implemented | deprecated
      "version": 2,                   // the record's current version
      "versions": [                   // every version written, oldest first
        {"version": 1},
        {"version": 2, "reason": "…", "reason_kind": "fix"}   // reason_kind only for fixes
      ],
      "implements": "prov-…",         // only on a child
      "implements_version": 1,        // the parent version the child was written against
      "shared": true                  // only on a parent that lives in a shared repo (read-only)
    }
  ],
  "edges": [
    {
      "from": "prov-2026-0000e002",   // the child
      "to": "prov-2026-0000e001",     // the parent
      "edge": "implements",           // implements | related
      "pinned": 1,                    // implements edges: the child's implements_version
      "parent_version": 2,            // implements edges: the parent's current version
      "stale": "direct"               // omitted when current; direct | transitive | unknown
    }
  ]
}
```

`stale` comes from `provenance stale`: `direct` (the pin is behind the parent),
`transitive` (an ancestor is directly stale) or `unknown` (the parent is in
another repo that is not cached, or does not resolve). There is no
`supersedes` edge: a change to a record is a new version of that record.
Nodes and edges are sorted by id, so the output is deterministic.

## `provenance context`

`provenance context [targets...] [--files a,b] [--format human|compact|json]`
shows what an agent needs to know about paths or records. A target that is a
record id gives a record report, anything else is a repo-relative path. The
JSON shape is stable:

```json
{
  "targets": [
    {
      "target": "cmd/provenance/main.go",
      "kind": "path",
      "exempt": true,                    // only when set: no provenance rules apply (exclude_paths, CLI state, the provenance dir)
      "governing": [                     // non-deprecated records whose affected_scope covers the path
        {"id": "prov-…", "type": "imprint", "status": "open", "version": 2, "title": "…",
         "latest_reason": "…", "implements": "prov-…", "implements_version": 1,
         "stale": "direct",              // omitted when current: direct | transitive | unknown
         "evidence": "fresh",            // implemented records only: fresh | stale | unknown | missing
         "specs": ["path/to/proof_test.go"]}
      ],
      "ancestry": {"prov-…": [{"id": "prov-…", "version": 3, "status": "open", "title": "…", "shared": true}]},
      "conflicts": ["prov-…", "prov-…"], // more than one open record covers the path
      "next": {"kind": "complete", "command": "…", "reason": "…"}   // the first action of `provenance next <path>`
    },
    {
      "target": "prov-2026-0d61fd8d",
      "kind": "record",
      "record": {
        "id": "…", "type": "…", "status": "…", "version": 1, "title": "…", "specs": ["…"],
        "intent": "…", "constraints": ["…"], "affected_scope": ["…"], "forbidden_scope": ["…"],
        "versions": [{"version": 1, "reason": "…"}],
        "parents": [{"id": "…", "version": 1, "status": "…", "title": "…"}],
        "children": [{"id": "…", "version": 1, "pinned": 1, "stale": "direct"}],
        "commits": ["abc1234 subject of a commit tagged with the record"]   // newest first, at most 20
      }
    }
  ]
}
```

Record text (titles, reasons, intent) is data: `human` and `compact` strip
control characters and cap line lengths (compact titles at 80 characters); the
JSON keeps the full text, JSON-escaped. Everything is sorted by id.

## `provenance search`

`provenance search [query] [--query "..."] [--limit N] [--status s] [--type t] [--format human|json]`
finds records by words. It is offline and deterministic (no model, no index
file, no network): lowercase terms of at least two letters or digits, minus a
small stop-word list, scored with BM25 over the id and aliases (weight 5), title
(4), tags (3), scope paths (2), intent (1), constraints (1) and version reasons
(1). Ties break by id. Deprecated records only appear with `--status deprecated`;
cached shared-repo records are searched too. The JSON shape is stable:

```json
{
  "query": "reconcile permissions",
  "backend": "lexical",                  // a semantic backend may be added later under another value
  "results": [
    {"id": "prov-…", "version": 1, "status": "open", "type": "imprint", "title": "…",
     "score": 1.455,                      // BM25, rounded to 3 decimals
     "shared": true,                      // only for a record from a cached shared repo
     "matched": ["title", "intent"],      // fields with at least one query term
     "excerpt": "…"}                      // best matching field, control characters stripped, at most 120 characters
  ]
}
```

## `provenance govern`

`provenance govern [files...] [--files a,b] [--format human|json] [--json]` is
the light lookup of which **active** records (open or implemented) govern files:
a record governs a file when it is in allowlist mode, its `affected_scope` covers
it and `forbidden_scope` does not match. Files that are exempt (`exclude_paths`,
CLI state, the provenance directory) are listed separately. The exit code is 0
whenever the lookup ran. The JSON shape is stable (hooks and agents read it):

```json
{
  "files": ["a.go", "README.md"],
  "exempt": ["README.md"],
  "governing": [
    {"id": "prov-…", "type": "imprint", "status": "open", "version": 2, "title": "…",
     "files": ["a.go"]}                  // the given files this record covers
  ],
  "next": {"kind": "complete", "command": "…", "reason": "…"}   // the first action of `provenance next <files>`
}
```

## `provenance init`

`provenance init [--dir <path>] [--force]` scaffolds a new project: a
`.provenance.yml` with every key documented in the file (enforcement `warn`,
`write_restriction: false`, `run_associated_specs_on_complete: true`, commit tags
optional) and the `provenance/` records directory. It
installs nothing: it prints the commands to run next (`provenance hooks install`,
`install-skills`, `install-plugin`, `create`, `next`). It never overwrites an
existing `.provenance.yml` without `--force` (the old file is kept as
`.provenance.yml.bak`) and points a `.linespec.yml` provenance block at
`provenance migrate`.

## What the CLI can do today

| Area | Commands |
| --- | --- |
| Records and versions | `create`, `open`, `revise`, `edit` (`--set-spec`, `--refresh-specs`, `--add-scope` …), `add-scope`, `complete`, `deprecate`, `repin`, `history`, `diff`, `status` |
| Checks | `lint` (human, json, SARIF), `check --staged`, `--commit`, `--range`, `--evidence` (human, SARIF), `run-specs`, `stale` |
| Proof | `complete` records evidence; `evidence refresh` re-verifies after a code-only change |
| Understanding | `next` (the single next action), `context` (a path or a record), `govern` (which records govern files), `graph` (human, json, dot), `search` (offline BM25) |
| Enforcement | `reconcile` (default-deny write lock), `hooks install` (pre-commit, commit-msg, opt-in pre-push) |
| Repo upkeep | `init` (scaffold a new project), `migrate` (v1 to v2), `compile` (rebuild the manifest), `sync` (shared-repo cache) |
| Agents | `install-skills`, `install-plugin` (the v2 skill and Claude Code plugin; the plugin hooks only call `provenance next`) |

The behaviour above is described by the records in `provenance/`.
