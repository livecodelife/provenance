// Package lint validates provenance records. It is ported from linespec
// pkg/provenance/linter.go: the rules that survive versioning keep their
// messages and severities, the supersession rules retire, and new rules check
// versions against their snapshots. Every issue names its rule explicitly.
package lint

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/livecodelife/provenance/record"
	"github.com/livecodelife/provenance/scope"
)

// Severity is how serious an issue is.
type Severity string

const (
	SeverityError   Severity = "error"
	SeverityWarning Severity = "warning"
	SeverityHint    Severity = "hint"
)

// Rule ids. PROV001-022 keep linespec's meanings; PROV023 is linespec's
// imprint scope containment; PROV024-028 are the v2 version rules; PROV029
// and up name carried checks that linespec's SARIF table never had an id for.
const (
	RuleRequiredField       = "PROV002"
	RuleStatus              = "PROV003"
	RuleID                  = "PROV004"
	RuleDate                = "PROV005"
	RuleRelated             = "PROV008"
	RuleScopeConflict       = "PROV009" // a pattern in both affected and forbidden scope
	RuleSpecsMissing        = "PROV010"
	RuleConstraintsNoSpecs  = "PROV011"
	RuleIntentNoConstraints = "PROV012"
	RuleTitleLength         = "PROV013"
	RuleScopeOverlap        = "PROV015" // open records sharing scope (warning)
	RuleDeadRecord          = "PROV016"
	RuleSpecFile            = "PROV018"
	RuleInvalidPattern      = "PROV019"
	RuleImplementsType      = "PROV021"
	RuleImplementsUnknown   = "PROV022"
	RuleImprintContainment  = "PROV023"
	RuleUnbumpedChange      = "PROV024"
	RuleBumpWithoutReason   = "PROV025"
	RuleVersionSnapshots    = "PROV026"
	RuleImplementsPin       = "PROV027"
	RuleNotMigrated         = "PROV028"
	RuleScopePath           = "PROV029"
	RuleType                = "PROV030"
	RuleNotApplicable       = "PROV031"
	RuleSealedAtSHA         = "PROV032"
	RuleLockedScope         = "PROV033"
	RuleSpecIgnored         = "PROV034" // a spec on a gitignored path (warning)
)

// Issue is one finding.
type Issue struct {
	Rule     string
	RecordID string
	Field    string
	Message  string
	Severity Severity
}

// Result holds every issue and the counts linespec reports.
type Result struct {
	Issues       []Issue
	Records      int
	PassedCount  int
	ErrorCount   int
	WarningCount int
	HintCount    int
	Enforcement  string
}

func (r *Result) add(i Issue) {
	r.Issues = append(r.Issues, i)
	switch i.Severity {
	case SeverityError:
		r.ErrorCount++
	case SeverityWarning:
		r.WarningCount++
	case SeverityHint:
		r.HintCount++
	}
}

// HasErrors reports whether any issue is an error.
func (r *Result) HasErrors() bool { return r.ErrorCount > 0 }

// Options configure a lint run.
type Options struct {
	// Root is the repository root that scope and spec paths resolve against.
	Root string
	// Enforcement is none, warn or strict; it decides how a missing spec on
	// an open record is reported.
	Enforcement string
	// ExcludePaths are exempt from scope path checks.
	ExcludePaths []string
	// Shared are records from shared repos, by id: they resolve references
	// and are linted as read-only remote records.
	Shared map[string]*record.Record
	// SharedRepos names the configured shared repos, for implements_repo.
	SharedRepos []string
	// ManifestPath, when set, is the version manifest to check against the
	// snapshots (PROV026); relative paths resolve against Root.
	ManifestPath string
}

type linter struct {
	store  record.Store
	opts   Options
	local  map[string]*record.Record
	all    map[string]*record.Record
	remote map[string]bool
	res    *Result
	// ignored holds the associated spec paths git ignores, filled once.
	ignored map[string]bool
}

// Lint validates every record in store, plus the shared records in opts.
func Lint(store record.Store, opts Options) (*Result, error) {
	ids, err := store.List()
	if err != nil {
		return nil, err
	}
	l := &linter{store: store, opts: opts, local: map[string]*record.Record{}, all: map[string]*record.Record{},
		remote: map[string]bool{}, res: &Result{Enforcement: opts.Enforcement}}
	var records []*record.Record
	for _, id := range ids {
		r, _, err := store.Load(id)
		if err != nil {
			return nil, err
		}
		if _, dup := l.all[r.ID]; dup {
			return nil, fmt.Errorf("duplicate record id %s", r.ID)
		}
		l.local[r.ID], l.all[r.ID] = r, r
		records = append(records, r)
	}
	for _, id := range sortedKeys(opts.Shared) {
		r := opts.Shared[id]
		if _, dup := l.all[id]; dup {
			continue
		}
		l.all[id], l.remote[id] = r, true
		records = append(records, r)
	}
	l.ignored = gitIgnored(opts.Root, specPaths(records))
	for _, r := range records {
		l.lintRecord(r)
	}
	l.checkDeadRecords(records)
	l.checkLockedScope(records)
	l.checkOpenOverlap(records)
	l.warnNotMigrated(records)
	if opts.ManifestPath != "" {
		if err := l.checkManifest(); err != nil {
			return nil, err
		}
	}
	l.res.Records = len(records)
	l.res.PassedCount = len(records) - l.res.ErrorCount
	return l.res, nil
}

func (l *linter) add(rule string, r *record.Record, field string, sev Severity, format string, args ...any) {
	msg := format
	if len(args) > 0 {
		msg = fmt.Sprintf(format, args...)
	}
	l.res.add(Issue{Rule: rule, RecordID: r.ID, Field: field, Message: msg, Severity: sev})
}

func (l *linter) lintRecord(r *record.Record) {
	l.requiredFields(r)
	l.id(r)
	l.status(r)
	l.typ(r)
	l.date(r)
	l.related(r)
	l.implements(r)
	l.notApplicable(r)
	if r.Schema == 1 && !l.remote[r.ID] {
		// One repo-level warning is added after the loop; the record's own
		// entry is only a hint, so unmigrated repos are not flooded.
		l.add(RuleNotMigrated, r, "version", SeverityHint, "Record uses the v1 schema; run `provenance migrate` to version it.")
	}
	if r.Schema == 2 && !l.remote[r.ID] {
		l.versionRules(r)
	}

	if r.Status == record.StatusDraft {
		l.titleLength(r)
		l.sealedAtSHA(r)
		return
	}
	if l.remote[r.ID] {
		l.titleLength(r)
		l.constraintsHint(r)
		return
	}
	l.scopePatterns(r)
	l.scopeSelfOverlap(r)
	l.scopePaths(r)
	l.imprintContainment(r)
	l.associatedSpecs(r)
	l.titleLength(r)
	l.constraintsHint(r)
	l.sealedAtSHA(r)
}

func terminal(s record.Status) bool {
	return s == record.StatusImplemented || s == record.StatusSuperseded || s == record.StatusDeprecated
}

func (l *linter) requiredFields(r *record.Record) {
	for _, f := range []struct{ name, value string }{
		{"id", r.ID}, {"title", r.Title}, {"status", string(r.Status)},
		{"created_at", r.CreatedAt}, {"author", r.Author}, {"intent", r.Intent},
	} {
		if strings.TrimSpace(f.value) == "" || f.value == "null" {
			l.add(RuleRequiredField, r, f.name, SeverityError, "Missing required field: %s", f.name)
		}
	}
	switch r.EffectiveType() {
	case record.TypeBrief:
		if len(r.Constraints) == 0 {
			l.add(RuleRequiredField, r, "constraints", SeverityError, "Brief records must have constraints. Briefs define the business rationale and must carry machine-verifiable constraints.")
		}
	case record.TypeImprint:
		if strings.TrimSpace(r.Implements) == "" {
			l.add(RuleRequiredField, r, "implements", SeverityError, "Imprint records must set implements pointing at the parent Blueprint.")
		}
	}
}

func (l *linter) id(r *record.Record) {
	if r.ID != "" && !record.IDPattern.MatchString(r.ID) {
		l.add(RuleID, r, "id", SeverityError, "Invalid ID format %q: must be prov-YYYY-NNN", r.ID)
	}
}

func (l *linter) status(r *record.Record) {
	if r.Status == "" || r.Status.Valid(r.Schema) {
		return
	}
	if r.Schema == 1 {
		l.add(RuleStatus, r, "status", SeverityError, "Invalid status %q: must be one of draft, open, implemented, superseded, deprecated", r.Status)
	} else {
		l.add(RuleStatus, r, "status", SeverityError, "Invalid status %q: must be one of draft, open, implemented, deprecated", r.Status)
	}
}

func (l *linter) typ(r *record.Record) {
	if r.Type == "" {
		if r.Schema == 1 {
			l.add(RuleType, r, "type", SeverityHint, "No type field set; defaulting to 'blueprint' for backward compatibility. Set type: brief|blueprint|bug|imprint to suppress this hint.")
		} else {
			l.add(RuleType, r, "type", SeverityHint, "No type field set; defaulting to 'blueprint'. Set type: brief|blueprint|imprint to suppress this hint.")
		}
		return
	}
	if !r.Type.Valid(r.Schema) {
		if r.Schema == 1 {
			l.add(RuleType, r, "type", SeverityError, "Invalid type %q: must be one of brief, blueprint, bug, imprint", r.Type)
		} else {
			l.add(RuleType, r, "type", SeverityError, "Invalid type %q: must be one of brief, blueprint, imprint", r.Type)
		}
	}
}

var datePattern = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)

func (l *linter) date(r *record.Record) {
	if r.CreatedAt != "" && !datePattern.MatchString(r.CreatedAt) {
		l.add(RuleDate, r, "created_at", SeverityError, "Invalid date format %q: must be YYYY-MM-DD", r.CreatedAt)
	}
}

func (l *linter) related(r *record.Record) {
	for _, id := range r.Related {
		if _, ok := l.all[id]; !ok {
			l.add(RuleRelated, r, "related", SeverityWarning, "related references unknown record: %s", id)
		}
	}
}

func (l *linter) implements(r *record.Record) {
	if r.Implements == "" {
		return
	}
	if !record.IDPattern.MatchString(r.Implements) {
		l.add(RuleImplementsType, r, "implements", SeverityError,
			"implements value %q is not a valid provenance record ID (expected prov-YYYY-NNN format).", r.Implements)
		return
	}
	target, ok := l.all[r.Implements]
	if !ok {
		l.add(RuleImplementsUnknown, r, "implements", SeverityError,
			"implements references unknown record: %s. The referenced record was not found locally or in any configured shared_repo cache. If this record lives in a remote repository, add it to shared_repos in .linespec.yml and run 'linespec provenance sync'.", r.Implements)
		return
	}
	rt, tt := r.EffectiveType(), target.EffectiveType()
	parent := map[record.Type]record.Type{record.TypeBlueprint: record.TypeBrief, record.TypeImprint: record.TypeBlueprint}
	want, allowed := parent[rt]
	if !allowed {
		l.add(RuleImplementsType, r, "implements", SeverityError,
			"A %s record cannot use the implements field. Only blueprint (implements brief) and imprint (implements blueprint) records may use implements.", rt)
		return
	}
	if tt != want {
		l.add(RuleImplementsType, r, "implements", SeverityError,
			"Invalid implements relationship: a %s record must implement a %s record, but %s is a %s. Allowed relationships: blueprint implements brief, imprint implements blueprint.", rt, want, r.Implements, tt)
	}
}

func (l *linter) notApplicable(r *record.Record) {
	switch r.EffectiveType() {
	case record.TypeBrief:
		if len(r.AffectedScope) > 0 {
			l.add(RuleNotApplicable, r, "affected_scope", SeverityWarning, "affected_scope is not applicable on brief records; briefs define business rationale, not implementation scope")
		}
		if len(r.ForbiddenScope) > 0 {
			l.add(RuleNotApplicable, r, "forbidden_scope", SeverityWarning, "forbidden_scope is not applicable on brief records")
		}
		if len(r.AssociatedSpecs) > 0 {
			l.add(RuleNotApplicable, r, "associated_specs", SeverityWarning, "associated_specs is not applicable on brief records; proof artifacts belong on blueprints and imprints")
		}
	case record.TypeImprint:
		if len(r.AssociatedTraces) > 0 {
			l.add(RuleNotApplicable, r, "associated_traces", SeverityWarning, "associated_traces is not applicable on imprint records; trace results are operational artifacts anchored at the blueprint tier")
		}
		if len(r.Monitors) > 0 {
			l.add(RuleNotApplicable, r, "monitors", SeverityWarning, "monitors is not applicable on imprint records; production monitoring is anchored at the blueprint tier")
		}
	}
}

func (l *linter) path(p string) string {
	if l.opts.Root == "" || filepath.IsAbs(p) {
		return p
	}
	return filepath.Join(l.opts.Root, p)
}

// walk visits every regular file under Root by root-relative path; visit
// returns true to stop.
func (l *linter) walk(visit func(rel string) bool) error {
	root := l.opts.Root
	if root == "" {
		root = "."
	}
	return filepath.Walk(root, func(p string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}
		rel := p
		if root != "." {
			if r, e := filepath.Rel(root, p); e == nil {
				rel = r
			}
		}
		if visit(rel) {
			return filepath.SkipDir
		}
		return nil
	})
}

func (l *linter) scopePatterns(r *record.Record) {
	for _, p := range append(append([]string(nil), r.AffectedScope...), r.ForbiddenScope...) {
		if scope.IsRegex(p) {
			if _, err := regexp.Compile(p[3:]); err != nil {
				l.add(RuleInvalidPattern, r, "scope", SeverityError, "Invalid regex pattern %q: %v", p, err)
			}
		}
	}
}

func (l *linter) scopeSelfOverlap(r *record.Record) {
	for _, a := range r.AffectedScope {
		for _, f := range r.ForbiddenScope {
			if scope.Overlap(a, f) {
				l.add(RuleScopeConflict, r, "scope", SeverityError, "Pattern %q appears in both affected_scope and forbidden_scope", a)
			}
		}
	}
}

// scopePaths: open records error on missing paths (affected_scope only hints,
// for declare-before-create), terminal records warn, drafts are skipped.
func (l *linter) scopePaths(r *record.Record) {
	if r.Status == record.StatusDraft {
		return
	}
	missing := SeverityError
	if terminal(r.Status) {
		missing = SeverityWarning
	}
	affectedNotFound := missing
	if r.Status == record.StatusOpen {
		affectedNotFound = SeverityHint
	}
	l.scopeList(r, r.AffectedScope, missing, affectedNotFound)
	l.scopeList(r, r.ForbiddenScope, missing, missing)
}

func (l *linter) scopeList(r *record.Record, patterns []string, invalid, notFound Severity) {
	for _, p := range patterns {
		if strings.TrimSpace(p) == "" || scope.Excluded(p, l.opts.ExcludePaths) {
			continue
		}
		switch {
		case scope.IsRegex(p):
			l.regexPattern(r, p, notFound)
		case strings.Contains(p, "*") || strings.Contains(p, "?"):
			l.globPattern(r, p, notFound)
		default:
			l.exactPath(r, p, invalid, notFound)
		}
	}
}

func (l *linter) exactPath(r *record.Record, p string, invalid, notFound Severity) {
	info, err := os.Stat(l.path(p))
	if err != nil {
		if os.IsNotExist(err) {
			if notFound == SeverityHint {
				l.add(RuleScopePath, r, "scope", notFound, "Scope path does not exist yet — declared for creation; opening unlocks its parent directory so it can be created: %s", p)
			} else {
				l.add(RuleScopePath, r, "scope", notFound, "Scope path does not exist: %s", p)
			}
		} else {
			l.add(RuleScopePath, r, "scope", invalid, "Cannot access scope path %s: %v", p, err)
		}
		return
	}
	if info.IsDir() {
		l.add(RuleScopePath, r, "scope", invalid, "Scope path is a directory, not a file (use glob pattern for directories): %s", p)
	}
}

func (l *linter) globPattern(r *record.Record, p string, notFound Severity) {
	matches, err := filepath.Glob(l.path(p))
	if err != nil {
		l.add(RuleInvalidPattern, r, "scope", SeverityError, "Invalid glob pattern %q: %v", p, err)
		return
	}
	if len(matches) == 0 {
		if notFound == SeverityHint {
			l.add(RuleScopePath, r, "scope", notFound, "Glob pattern matches no files yet — declared for creation; opening unlocks its base directory so a matching file can be created: %s", p)
		} else {
			l.add(RuleScopePath, r, "scope", notFound, "Glob pattern matches no files: %s", p)
		}
	}
}

func (l *linter) regexPattern(r *record.Record, p string, notFound Severity) {
	re, err := regexp.Compile(p[3:])
	if err != nil {
		l.add(RuleInvalidPattern, r, "scope", SeverityError, "Invalid regex pattern %q: %v", p, err)
		return
	}
	found := false
	if err := l.walk(func(rel string) bool {
		if re.MatchString(rel) {
			found = true
			return true
		}
		return false
	}); err != nil {
		l.add(RuleInvalidPattern, r, "scope", SeverityError, "Error walking filesystem for regex pattern %s: %v", p, err)
		return
	}
	if !found {
		if notFound == SeverityHint {
			l.add(RuleScopePath, r, "scope", notFound, "Regex pattern matches no files yet — declared for creation; opening unlocks its base directory so a matching file can be created: %s", p)
		} else {
			l.add(RuleScopePath, r, "scope", notFound, "Regex pattern matches no files: %s", p)
		}
	}
}

func (l *linter) associatedSpecs(r *record.Record) {
	missing := SeverityError
	if terminal(r.Status) {
		missing = SeverityWarning
	}
	for _, s := range r.AssociatedSpecs {
		info, err := os.Stat(l.path(s.Path))
		if err != nil {
			if os.IsNotExist(err) {
				l.add(RuleSpecFile, r, "associated_specs", missing, "Proof artifact does not exist: %s", s.Path)
			} else {
				l.add(RuleSpecFile, r, "associated_specs", missing, "Cannot access proof artifact %s: %v", s.Path, err)
			}
			continue
		}
		if info.IsDir() {
			l.add(RuleSpecFile, r, "associated_specs", missing, "Proof artifact path is a directory, not a file: %s", s.Path)
		} else if l.ignored[s.Path] {
			l.add(RuleSpecIgnored, r, "associated_specs", SeverityWarning, "Proof artifact %s is gitignored, so lint depends on the checkout: it hashes the file only where it has been built, and a clean clone reports the record as changed. Track the file or pick another proof.", s.Path)
		}
	}
	if r.EffectiveType() == record.TypeBrief || r.Status != record.StatusOpen || len(r.AssociatedSpecs) > 0 {
		return
	}
	switch l.opts.Enforcement {
	case "strict":
		l.add(RuleSpecsMissing, r, "associated_specs", SeverityError, "No associated specs (open) [strict]")
	case "warn":
		l.add(RuleSpecsMissing, r, "associated_specs", SeverityWarning, "No associated specs (open)")
	case "none":
		if len(r.Constraints) > 0 {
			l.add(RuleConstraintsNoSpecs, r, "associated_specs", SeverityHint, "Record has constraints but no associated specs")
		}
	}
}

func (l *linter) titleLength(r *record.Record) {
	if len(r.Title) > 120 {
		l.add(RuleTitleLength, r, "title", SeverityWarning, "Title exceeds 120 characters (%d chars)", len(r.Title))
	}
}

func (l *linter) constraintsHint(r *record.Record) {
	if r.Intent != "" && len(r.Constraints) == 0 {
		l.add(RuleIntentNoConstraints, r, "constraints", SeverityHint, "Record has intent but no constraints. Consider adding specific behavioral constraints.")
	}
}

var shaPattern = regexp.MustCompile(`^[a-f0-9]{7,40}$`)

func (l *linter) sealedAtSHA(r *record.Record) {
	if r.SealedAtSHA == "" {
		return
	}
	if !terminal(r.Status) {
		if r.Schema == 1 {
			l.add(RuleSealedAtSHA, r, "sealed_at_sha", SeverityWarning, "sealed_at_sha is set but record status is %s (should only be set for implemented, superseded, or deprecated records)", r.Status)
		} else {
			l.add(RuleSealedAtSHA, r, "sealed_at_sha", SeverityWarning, "sealed_at_sha is set but record status is %s (should only be set for implemented or deprecated records)", r.Status)
		}
	}
	if !shaPattern.MatchString(r.SealedAtSHA) {
		l.add(RuleSealedAtSHA, r, "sealed_at_sha", SeverityWarning, "Invalid sealed_at_sha format: %s (expected 7-40 hex characters)", r.SealedAtSHA)
	}
}

func (l *linter) imprintContainment(r *record.Record) {
	if r.EffectiveType() != record.TypeImprint || r.Status != record.StatusOpen || r.Implements == "" || len(r.AffectedScope) == 0 {
		return
	}
	bp, ok := l.all[r.Implements]
	if !ok || len(bp.AffectedScope) == 0 {
		return
	}
	for _, file := range l.filesMatching(r.AffectedScope) {
		covered := false
		for _, p := range bp.AffectedScope {
			if ok, _ := scope.Match(file, p); ok {
				covered = true
				break
			}
		}
		if !covered {
			l.add(RuleImprintContainment, r, "affected_scope", SeverityError,
				"PROV023: imprint scope exceeds blueprint scope — %q is governed by this imprint but not covered by any pattern in %s's affected_scope. Widen the blueprint's scope first.", file, bp.ID)
		}
	}
}

func (l *linter) filesMatching(patterns []string) []string {
	seen := map[string]bool{}
	var files []string
	for _, p := range patterns {
		if strings.TrimSpace(p) == "" {
			continue
		}
		var expr string
		switch {
		case scope.IsRegex(p):
			expr = p[3:]
		case strings.Contains(p, "*") || strings.Contains(p, "?"):
			expr = scope.GlobToRegex(p)
		default:
			if _, err := os.Stat(l.path(p)); err == nil && !seen[p] {
				seen[p] = true
				files = append(files, p)
			}
			continue
		}
		re, err := regexp.Compile(expr)
		if err != nil {
			continue
		}
		_ = l.walk(func(rel string) bool {
			if re.MatchString(rel) && !seen[rel] {
				seen[rel] = true
				files = append(files, rel)
			}
			return false
		})
	}
	return files
}

// checkDeadRecords warns when no file matches any affected_scope pattern.
func (l *linter) checkDeadRecords(records []*record.Record) {
	for _, r := range records {
		if len(r.AffectedScope) == 0 {
			continue
		}
		any := false
		for _, p := range r.AffectedScope {
			if strings.TrimSpace(p) == "" {
				continue
			}
			if l.patternHasMatches(p) {
				any = true
				break
			}
		}
		if !any {
			l.add(RuleDeadRecord, r, "scope", SeverityWarning, "Dead record: all governed files have been deleted. Consider deprecating this record.")
		}
	}
}

func (l *linter) patternHasMatches(p string) bool {
	switch {
	case scope.IsRegex(p):
		re, err := regexp.Compile(p[3:])
		if err != nil {
			return false
		}
		found := false
		_ = l.walk(func(rel string) bool {
			if re.MatchString(rel) {
				found = true
				return true
			}
			return false
		})
		return found
	case strings.Contains(p, "*") || strings.Contains(p, "?"):
		m, err := filepath.Glob(l.path(p))
		return err == nil && len(m) > 0
	default:
		_, err := os.Stat(l.path(p))
		return !os.IsNotExist(err)
	}
}

// checkLockedScope: an open record's surfaces may not overlap a locked
// record's, unless (v1) it supersedes that locked record.
func (l *linter) checkLockedScope(records []*record.Record) {
	var locked []*record.Record
	for _, r := range records {
		if r.Locked {
			locked = append(locked, r)
		}
	}
	if len(locked) == 0 {
		return
	}
	for _, open := range records {
		if open.Status != record.StatusOpen {
			continue
		}
		surfaces := append([]string(nil), open.AffectedScope...)
		for _, s := range open.AssociatedSpecs {
			surfaces = append(surfaces, s.Path)
		}
		for _, lk := range locked {
			if open.Supersedes == lk.ID {
				continue
			}
			lockedSurfaces := append([]string(nil), lk.AffectedScope...)
			for _, s := range lk.AssociatedSpecs {
				lockedSurfaces = append(lockedSurfaces, s.Path)
			}
			for _, os := range surfaces {
				for _, ls := range lockedSurfaces {
					if !scope.Overlap(os, ls) {
						continue
					}
					field := "affected_scope"
					if !contains(open.AffectedScope, os) {
						field = "associated_specs"
					}
					l.add(RuleLockedScope, open, field, SeverityError,
						"Surface %q overlaps with locked record %s surface %q. Resolving this is a governance decision: narrow this record's scope, or — only if you are revising the locked decision — supersede %s to reopen the layer.", os, lk.ID, ls, lk.ID)
				}
			}
		}
	}
}

// checkOpenOverlap warns when two open local records share scope (PROV015).
// Region ownership protects the code, so this is never an error.
func (l *linter) checkOpenOverlap(records []*record.Record) {
	var open []*record.Record
	for _, r := range records {
		if r.Status == record.StatusOpen && !l.remote[r.ID] {
			open = append(open, r)
		}
	}
	for i, a := range open {
		for _, b := range open[i+1:] {
			for _, pa := range a.AffectedScope {
				for _, pb := range b.AffectedScope {
					if scope.Overlap(pa, pb) {
						l.add(RuleScopeOverlap, b, "affected_scope", SeverityWarning, "Scope %q overlaps open record %s (%q)", pb, a.ID, pa)
					}
				}
			}
		}
	}
}

// versionRules are the v2 checks against the record's snapshots.
func (l *linter) versionRules(r *record.Record) {
	if r.Version > 1 && strings.TrimSpace(r.Reason) == "" {
		l.add(RuleBumpWithoutReason, r, "reason", SeverityError, "Version %d has no reason. Every version after the first needs one.", r.Version)
	}
	l.implementsPin(r)

	versions, err := l.store.Versions(r.ID)
	if err != nil {
		l.add(RuleVersionSnapshots, r, "version", SeverityError, "Cannot read version snapshots: %v", err)
		return
	}
	if len(versions) == 0 {
		if r.Status != record.StatusDraft {
			l.add(RuleVersionSnapshots, r, "version", SeverityError, "Record is %s but has no version snapshot. Leaving draft writes version 1.", r.Status)
		}
		return
	}
	var prev *record.Snapshot
	for i, v := range versions {
		snap, err := l.store.Snapshot(r.ID, v)
		if err != nil {
			l.add(RuleVersionSnapshots, r, "version", SeverityError, "Snapshot for version %d does not verify: %v", v, err)
			return
		}
		if v != i+1 {
			l.add(RuleVersionSnapshots, r, "version", SeverityError, "Snapshot versions are not consecutive: found %d where %d was expected", v, i+1)
			return
		}
		want := ""
		if prev != nil {
			want = prev.ContentHash
		}
		if snap.ParentHash != want {
			l.add(RuleVersionSnapshots, r, "version", SeverityError, "Snapshot for version %d does not chain to version %d", v, v-1)
			return
		}
		prev = snap
	}
	if r.Version != prev.Version {
		l.add(RuleVersionSnapshots, r, "version", SeverityError, "Record says version %d but the latest snapshot is version %d", r.Version, prev.Version)
		return
	}
	if h, _ := r.ContentHash(l.store.Specs()); h != prev.ContentHash {
		l.add(RuleUnbumpedChange, r, "version", SeverityError, "Content changed since version %d without a version bump. Run `provenance revise %s --reason \"…\"`.", r.Version, r.ID)
	}
}

func (l *linter) implementsPin(r *record.Record) {
	if r.ImplementsVersion < 0 {
		l.add(RuleImplementsPin, r, "implements_version", SeverityError, "implements_version %d is negative", r.ImplementsVersion)
	}
	if (r.ImplementsVersion != 0 || r.ImplementsRepo != "") && r.Implements == "" {
		l.add(RuleImplementsPin, r, "implements_version", SeverityError, "implements_version or implements_repo is set without implements")
		return
	}
	if r.ImplementsRepo != "" && !contains(l.opts.SharedRepos, r.ImplementsRepo) {
		l.add(RuleImplementsPin, r, "implements_repo", SeverityError, "implements_repo %q is not a configured shared repo", r.ImplementsRepo)
	}
	if parent, ok := l.all[r.Implements]; ok && parent.Schema == 2 && r.ImplementsVersion > parent.Version {
		l.add(RuleImplementsPin, r, "implements_version", SeverityError, "implements_version %d is ahead of %s, which is at version %d", r.ImplementsVersion, parent.ID, parent.Version)
	}
}

func contains(list []string, s string) bool {
	for _, x := range list {
		if x == s {
			return true
		}
	}
	return false
}

func sortedKeys[V any](m map[string]V) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

// checkManifest compares the version manifest with the snapshots (PROV026).
// Repo-level problems carry an empty record id.
func (l *linter) checkManifest() error {
	want, err := record.BuildManifest(l.store)
	if err != nil {
		return err
	}
	repo := &record.Record{}
	path := l.path(l.opts.ManifestPath)
	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		if len(want) > 0 {
			l.add(RuleVersionSnapshots, repo, "manifest", SeverityError, "Version manifest %s is missing but %d records have snapshots. Rebuild it from the snapshots.", l.opts.ManifestPath, len(want))
		}
		return nil
	}
	if err != nil {
		return err
	}
	got, err := record.ParseManifest(data)
	if err != nil {
		l.add(RuleVersionSnapshots, repo, "manifest", SeverityError, "Version manifest %s is malformed: %v", l.opts.ManifestPath, err)
		return nil
	}
	have := map[string]record.ManifestLine{}
	for _, g := range got {
		have[g.ID] = g
	}
	expected := map[string]bool{}
	for _, w := range want {
		expected[w.ID] = true
		r := l.local[w.ID]
		g, ok := have[w.ID]
		switch {
		case !ok:
			l.add(RuleVersionSnapshots, r, "manifest", SeverityError, "Version manifest has no line for this record (snapshots are at version %d).", w.Version)
		case g.Version != w.Version || g.ContentHash != w.ContentHash:
			l.add(RuleVersionSnapshots, r, "manifest", SeverityError, "Version manifest says version %d (%s) but the latest snapshot is version %d (%s).", g.Version, g.ContentHash, w.Version, w.ContentHash)
		}
	}
	for _, g := range got {
		if expected[g.ID] {
			continue
		}
		r := &record.Record{ID: g.ID}
		if lr, ok := l.local[g.ID]; ok {
			r = lr
			l.add(RuleVersionSnapshots, r, "manifest", SeverityError, "Version manifest lists version %d but this record has no snapshots.", g.Version)
		} else {
			l.add(RuleVersionSnapshots, r, "manifest", SeverityError, "Version manifest lists %s, which has no record file.", g.ID)
		}
	}
	return nil
}

// warnNotMigrated adds the single repo-level PROV028 warning for local
// records still on the v1 schema. Shared (remote) records are not counted.
func (l *linter) warnNotMigrated(records []*record.Record) {
	n := 0
	for _, r := range records {
		if r.Schema == 1 && !l.remote[r.ID] {
			n++
		}
	}
	if n > 0 {
		l.add(RuleNotMigrated, &record.Record{}, "version", SeverityWarning,
			"%d records use the v1 schema; run `provenance migrate`.", n)
	}
}

// specPaths lists every associated spec path of the records, once.
func specPaths(records []*record.Record) []string {
	seen := map[string]bool{}
	var out []string
	for _, r := range records {
		for _, s := range r.AssociatedSpecs {
			if !seen[s.Path] {
				seen[s.Path] = true
				out = append(out, s.Path)
			}
		}
	}
	return out
}
