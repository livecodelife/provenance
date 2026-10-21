---
description: After hand-editing a provenance record, check it, find the records now behind it, and get them current (never runs a model grow)
argument-hint: <record-id> [more record ids]
allowed-tools: Bash(provenance:*), Bash(git diff:*), Bash(git status:*), Bash(git log:*), Bash(sedum grow --check:*), Read, Grep, Glob, Agent
---

You are bringing a repository's provenance records back to a consistent state after a person edited the records
named in: $ARGUMENTS

The edited records are the "parents" below. Work through the steps in order. Do not skip ahead, and never change
a record before step 3 has been approved. Keep every record you write at the behaviour level: no code, no
"exactly" or "verbatim" text, one requirement per constraint. If no record id was given, stop and ask for one.

## 1. Show what changed, then lint

For each edited record:

- Show the working-tree change with `git diff -- provenance/<id>.yml`.
- Say in one or two sentences what changed in behaviour terms (intent, constraints, scope, specs). Say whether the
  version number moved. A hand edit that did not go through `provenance edit` or `provenance revise` usually did
  not: the record then differs from its latest snapshot in `provenance/.versions/<id>/`.
- If the file changed but the version did not, say so, and tell the user the record needs
  `provenance revise <id> --reason "..."` (or `provenance edit <id> --reason "..."`) before anything can follow
  it. Offer to run it once they give a reason. Do not invent one.

Then run `provenance lint`. Show only the findings about the edited records. If any of them is an error, STOP and
report it; the rest of this command is meaningless until the edited records lint clean. Errors in other records
that were already there are not yours to fix: mention them once and carry on.

## 2. Find what is behind

Run `provenance stale --json`. The output is a list of findings, each with the record `id`, the `parent` it
implements, the version it is `pinned` to, the parent's current version, and a `kind`: `direct` (its parent moved),
`transitive` (an ancestor moved) or `unknown` (its parent cannot be read, for example a parent in another
repository, so it cannot be judged).

Keep the findings whose chain leads to one of the edited records. Walk the `parent` links to decide. Group them by
generation: children first, then their children. Report records you left out only as a count, and list the
`unknown` ones separately; those need a person.

If nothing is behind, say so and go straight to step 5.

## 3. Classify each stale child and wait

For each stale record, nearest generation first, read its file and its parent's change, and propose exactly one of:

- **UNAFFECTED**: the record is still correct against the new parent. Propose `provenance repin` with a one-line
  reason that says why the parent's change does not touch this record.
- **AFFECTED**: something the record states or relies on no longer holds, or the change creates a gap it should
  cover. Say which constraint or sentence is affected and what the record must now say, in words.

When a parent's change decides a child's classification, a child that is AFFECTED makes its own children worth
re-reading: classify them after you have the child's new text in mind.

Show the whole list as a table: record id, title, UNAFFECTED or AFFECTED, and the reason or the needed change.
Then ASK the user to approve, edit, or reject the list. Change nothing until they answer. If they change a
classification, use theirs.

## 4. Apply

Only after approval, and one record at a time, nearest generation first:

- UNAFFECTED: `provenance repin <id> --reason "<the approved reason>"`.
- AFFECTED: delegate the new record text to the `provenance-records` sub-agent (revise the existing record; do not
  create a new one). Ask for the complete change at once and apply it in a single edit, because every edit costs a
  version. Apply it with `provenance edit <id> --reason "..."` or `provenance revise`, then have the
  `provenance-record-reviewer` sub-agent review the revised record. Fix what it reports and review again until it
  passes. If either sub-agent is not installed, say so and stop at that record.

After each batch run `provenance lint` and `provenance stale --json` again. Repeat until lint has no new errors and
`stale` lists nothing behind the edited records. If something stays stale, report why.

Commands that change records may need the user's permission in the session; ask for it rather than working around
a refusal. Never use a force flag, never skip hooks, and never edit a record that is not on the approved list.

## 5. Print the Sedum commands. Do not run a model grow.

Do not call a model. You may run the offline check, and only that:

`sedum grow --check --generators generators --records provenance`

If the check prints no per-record status (only warnings about files no generator claims, for example), say so
plainly instead of guessing, and still print the grow command for the records you changed. Otherwise summarise
its output in one short table: records that are stale or need re-rendering, records that are unknown
(never grown), and records that are current. Then print, for the user to run themselves, one command for the
stale or re-render records (the repeatable `--only` flag names one record each):

`sedum grow --generators generators --records provenance --only <id> --only <id> ...`

Use the project's own wrapper for the model call if it has one; say that the user chooses the model and key. In
the Provisio herd that wrapper is `.herd/sedum-haiku grow ...` with the same arguments, and no other model is
allowed. Point out any stale record whose scope is a wide pattern: a record that authorises a whole directory
tree cannot be reconciled one record at a time, and the user should expect to review its regions by hand.

## 6. Remind about commit tags

Finish with a short checklist:

- Each commit that follows a record change carries that record's tag, `[<id>@<version>]`, using the version
  `edit`, `revise` or `repin` printed. One record tag per commit.
- Governed files regenerated by Sedum are committed after the records they come from.
- Run `provenance check --range origin/main..HEAD`, `provenance lint` and `provenance stale` before pushing.
