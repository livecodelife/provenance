# /provenance-update: bring records back in line after a hand edit

When you edit a blueprint or imprint by hand, the records that implement it fall behind: they pin the parent's old
version, and the code Sedum grew from them may no longer match. `/provenance-update` is a Claude Code slash
command that walks you from "I edited a record" to "everything that followed it is current, and I know which
Sedum commands to run". It changes records only after you approve a plan, and it never runs a model grow.

## Install

The command is one file. In this repository it is already there: `.claude/commands/provenance-update.md`. To use it
in another repository, copy that file to `.claude/commands/provenance-update.md` in that repository (or into
`~/.claude/commands/` to have it in every repository), then restart Claude Code. It needs the `provenance` CLI on
your path, and it uses two sub-agents if you have them: `provenance-records` (writes record text) and
`provenance-record-reviewer` (checks it). Without them it stops at the first record that needs rewriting.

## Use

```
/provenance-update prov-2026-1a2b3c4d
/provenance-update prov-2026-1a2b3c4d prov-2026-5e6f7a8b
```

Name the records you edited. Then it goes in six steps:

1. **Show what changed and lint.** It shows the `git diff` of each edited record and says in plain words what
   changed. It notes whether the version moved: a hand edit that skipped `provenance edit` or `revise` leaves the
   version where it was, and it will ask you for a reason to record the new version. Then it runs
   `provenance lint` and stops if the edited records have errors.
2. **Find what is behind.** It runs `provenance stale --json` and keeps the records whose chain leads to what you
   edited, directly or through other records. Records whose parent cannot be read (for example a parent in
   another repository) are listed apart for a person to look at.
3. **Classify and wait.** For each stale record it proposes *unaffected* (still correct: repin with a reason) or
   *affected* (needs a revision, with what must change). It shows a table and waits for your answer. Nothing
   changes before you approve.
4. **Apply.** It repins the unaffected ones and has the sub-agents revise the affected ones, nearest generation
   first, then re-runs lint and stale until they are clean.
5. **Sedum commands.** It runs only the offline check, `sedum grow --check --generators generators --records
   provenance`, summarises it, and prints the `sedum grow ... --only <id> --only <id>` command for you to run
   with your own model setup. It points out stale records whose scope is a wide pattern, which cannot be
   reconciled one record at a time.
6. **Reminders.** Commit tags (`[<id>@<version>]`, one record per commit), commit regenerated files after the
   records they come from, and run `check --range`, `lint` and `stale` before pushing.

## Permissions

`provenance repin`, `edit` and `revise` change records, so Claude Code asks for your permission the first time in a
session. Approving the plan in step 3 is separate from allowing the tool: if you run Claude Code in a mode that
cannot prompt (for example `claude -p`), allow `Bash(provenance:*)` for the run, or the command will stop and tell
you which commands it could not run.

## Tested

Tried on a scratch clone: one constraint added to a parent with six children left all six stale. The command listed
them, proposed five repins and (when told to) one revision, waited for approval, applied them through
the sub-agents, and finished with lint clean, stale empty and the Sedum command printed. Revising a record whose
older text breaks the reviewer's rules (a large intent that carries code) still fails review on that older text;
the command reports it and leaves the fix to you.

## What it will not do

- Change a record before you approve the plan, or any record that is not on the approved list.
- Call a model. Running `sedum grow` against a model is always left to you.
- Use a force flag or skip hooks.
