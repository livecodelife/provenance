# Context for sedum grow (the cmd/provenance CLI and the record package)

Pass with: sedum-haiku grow --generators generators --records provenance --prompt-file .sedum/context.md --only <id>
Advisory only: behaviour belongs in the record. Editing this file does not regenerate anything.

Conventions: cmd/provenance (package main)
- Package main, cmd/provenance. Every run function is func run<Cmd>(out io.Writer, cmd *cobra.Command, args []string) error and prints through out, never os.Stdout. User-facing errors go through r.fail(format, args...), then return an error.
- Each import is its own addImport. Write exactly `_, skip := ...` when a binding would otherwise be unused.
- Generated Go is gofmt-clean and calls only helpers listed here or in the record.

Helpers that exist in package main (call, do not redefine)
- openRepo(cmd *cobra.Command) (*repo, error); repo has Root string, Config config, Store *record.FilesStore.
- (r *repo) resolve(id string) (*record.Record, *record.Document, error) prints its own message.
- (r *repo) recordPath(id) string and versionsDir(id) string (relative to r.Root); const manifestPath; (r *repo) writeManifest() error; author() string; fail(format string, args ...any).
- (r *repo) git(args ...string) (string, error); commit(message string, paths ...string) error; unstage(paths ...string).
- snapshotFiles(paths ...string) []fileState; restoreFiles(states []fileState); tag(id string, version int) string gives "[<id>@<version>]".
- recordArg(cmd, args) (string, error) takes the id from --record or the first argument.

Record library (github.com/livecodelife/provenance/record)
- Record fields ID, Schema, Status, Version, Title, Intent, Constraints, AffectedScope, ForbiddenScope, Reason, ReasonKind, SealedAtSHA, AssociatedSpecs, Implements, ImplementsVersion.
- (*Record).ContentHash(specs SpecResolver) (string, []string); (*FilesStore).Specs(), .Versions(id), .Snapshot(id, v), .Save(r, doc), .Load(id), .List().
- record.Commit(s Store, r *Record, doc *Document, b Bump{Reason, ReasonKind, Author string}) (*Snapshot, error).

Conventions: record package
- Package record (module github.com/livecodelife/provenance). Generated Go is gofmt-clean and calls only helpers listed here or in the record.
- Canonical encoding is docs/HASHING.md section 5: bytes are written to a *bytes.Buffer; no encoding/json.
- Each import is its own addImport. Write exactly `_, skip := ...` when a binding would otherwise be unused.

Helpers that exist in package record (call, do not redefine)
- writeCanonical(b *bytes.Buffer, v any): the encoder switch in hash.go. Its default case panics "record: non-canonical value type".
- writeString(b *bytes.Buffer, s string): a JSON string with only the section 5 escapes.
- Canonical(v any) string; hashBytes(data []byte) string.

Document command (package main; files document_cmd.go, document_run.go, document_print.go)
- Where a record states a message as "The message of X is: text." the final period ends the sentence; the message is the text without it. Write headings, line templates and usage texts with the exact characters stated.
- The library package is github.com/livecodelife/provenance/document with Render(store, RenderOptions) (string, error), Parse(text) []Section, BuildPlan(store, text) (*Plan, error), Apply(store, *Plan, ApplyOptions) ([]Revised, error), and the types Section, Change, Problem, Plan, RenderOptions, Revised, ApplyOptions and the errors ErrNoReason, ErrConflicts, ErrStale; use the exact signatures the records state; call them, do not redefine them.
- printPlan already exists in package main (migrate); the new function is printDocumentPlan.
- Steer (the cobra flag kinds matter): register the scope flag with cmd.Flags().StringSlice (the string slice kind), never StringArray, because runDocument reads it with GetStringSlice.
- Steer (printDocumentPlan): when adding a conflict or problem line pass that Problem's own Line field as the sort line (for conflicts pass c.Line, not 0); only change lines use 0.
