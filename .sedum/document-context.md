# Context for growing the document library (package document, directory document/)

Pass with: sedum-haiku grow --generators generators --records provenance --prompt-file .sedum/document-context.md --only <id>
Advisory only: behaviour belongs in the records.

- Package document lives in the directory document/ of the module github.com/livecodelife/provenance. Each imprint owns the file named in its scope (render.go, parse.go, plan.go, apply.go).
- Where a record states a message as "The message of X is: text." the final period ends the sentence; the message is the text without it (for example `a reason is required`).
- Where a record states a heading, anchor or marker line, write it with the exact characters stated; do not add or drop spaces.
- Generated Go is gofmt-clean. Each import is its own addImport, owned by exactly one record per file; every other record in the file adds no imports.
- Other records of this package export the types Section, Change, Problem, Plan, RenderOptions, Revised, ApplyOptions and the functions Render, Parse, BuildPlan, Apply as their own records state; call them, do not redefine them.
- The record package is github.com/livecodelife/provenance/record; its Store interface has List() ([]string, error), Load(id string) (*Record, *Document, error), Save(*Record, *Document) error and Versions(id string) ([]int, error); the scope package is github.com/livecodelife/provenance/scope with Match(filePath, pattern string) (bool, error).
