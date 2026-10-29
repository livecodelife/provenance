---
description: Draft provenance records for existing code, one folder at a time, and stop for review after each folder
argument-hint: [directory ...]
---

Draft provenance records for code that has none yet. The directories to work on are: $ARGUMENTS

If no directory was given, work through the whole repository one top-level folder at a time. Everything this command writes is a DRAFT: say plainly, at the start and again at the end, that the results are drafts for the user to refine, not finished records.

For each folder:

1. Read the files in the folder and its git history (`git log` limited to the folder's path) to understand what the code does and why it changed.
2. Skip code that is already governed by a record. Find the records that govern a file with `provenance govern <files>` and see the next provenance action for it with `provenance next <files>`. Never change an existing record.
3. Use the provenance-records agent to draft blueprints for the behaviours you find, and imprints for the units that implement those behaviours. Each draft states one behaviour (blueprint) or one unit (imprint) in words, with one requirement per constraint, the files it covers as scope, and the existing tests as specs. Create each draft with `provenance create` and pass the `--no-edit` flag. Never run `provenance open` and never run `provenance complete`: the drafts stay drafts.
4. Run the provenance-record-reviewer agent over the drafts written for this folder and fix what it flags before showing them.
5. Stop. Show the drafts of this folder, with the reviewer's verdict, and ask the user before moving on to the next folder.
