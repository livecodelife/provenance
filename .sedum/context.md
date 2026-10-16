# Context for sedum grow (cmd/provenance)

Pass with: sedum-haiku grow --generators generators --records provenance --prompt-file .sedum/context.md --only <id>
Advisory only: behaviour belongs in the record. Editing this file does not regenerate anything.

Conventions
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
