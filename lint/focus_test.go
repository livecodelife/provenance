package lint

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/livecodelife/provenance/record"
)

// focusRepo builds a repository that trips most rules, several of them across
// records: overlapping open scopes, a locked layer, an imprint wider than its
// blueprint, a tampered snapshot, an unbumped edit, a v1 record, a stale
// manifest, a shared record.
func focusRepo(t *testing.T) (*record.FilesStore, Options) {
	t.Helper()
	root := t.TempDir()
	s := &record.FilesStore{Root: root}
	write := func(rel, body string) {
		p := filepath.Join(root, rel)
		if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(p, []byte(body), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	for _, f := range []string{"src/x.go", "src/y.go", "src/sub/z.go", "tests/x_test.sh", "docs/a.md"} {
		write(f, "x\n")
	}
	commit := func(r *record.Record) {
		t.Helper()
		if _, err := record.Commit(s, r, nil, record.Bump{Author: "a@b", Reason: "r"}); err != nil {
			t.Fatal(err)
		}
	}
	mk := func(n string, status record.Status, scope ...string) *record.Record {
		r := newRecord("prov-2026-0000f"+n, status)
		r.AffectedScope = scope
		return r
	}

	broad := mk("001", record.StatusOpen, "src/**")
	narrow := mk("002", record.StatusOpen, "src/x.go", "src/*.go")
	globs := mk("003", record.StatusOpen, "src/**/*.go", "re:^src/.*\\.go$")
	badRe := mk("004", record.StatusOpen, "re:(", "src/y.go")
	badRe.ForbiddenScope = []string{"src/y.go"}
	locked := mk("005", record.StatusImplemented, "docs/**")
	locked.Locked = true
	intrudes := mk("006", record.StatusOpen, "docs/a.md")
	missing := mk("007", record.StatusOpen, "nowhere/file.go", "nowhere/**", "re:^nowhere/")
	gone := mk("008", record.StatusImplemented, "nowhere/old.go")
	spec := mk("009", record.StatusOpen, "tests/x_test.sh")
	spec.AssociatedSpecs = []record.AssociatedSpec{{Path: "tests/x_test.sh"}, {Path: "tests/absent_test.sh"}}
	blueprint := mk("010", record.StatusOpen, "src/sub/**")
	blueprint.Type = record.TypeBlueprint
	imprint := mk("011", record.StatusOpen, "src/sub/z.go", "src/x.go", "re:^src/y")
	imprint.Type, imprint.Implements, imprint.ImplementsVersion = record.TypeImprint, blueprint.ID, 1
	orphan := mk("012", record.StatusOpen, "src/y.go")
	orphan.Type, orphan.Implements = record.TypeImprint, "prov-2026-0000fabc"
	ahead := mk("013", record.StatusOpen, "docs/**")
	ahead.Type, ahead.Implements, ahead.ImplementsVersion = record.TypeImprint, blueprint.ID, 9
	draft := mk("014", record.StatusDraft, "src/**")
	draft.Title = strings.Repeat("long ", 30)
	tampered := mk("015", record.StatusOpen, "src/sub/z.go")
	_ = tampered
	unbumped := mk("016", record.StatusOpen, "docs/a.md")
	later := mk("017", record.StatusOpen, "src/sub/**")
	later.Constraints = append(later.Constraints, "second")
	deprecated := mk("018", record.StatusSuperseded, "nowhere/**")

	all := []*record.Record{broad, narrow, globs, badRe, locked, intrudes, missing, gone, spec, blueprint, imprint,
		orphan, ahead, tampered, unbumped, later, deprecated}
	for _, r := range all {
		commit(r)
	}
	// A second version for one record; a draft has no snapshot.
	r2, doc, _ := s.Load(later.ID)
	r2.Constraints = append(r2.Constraints, "third")
	if _, err := record.Commit(s, r2, doc, record.Bump{Reason: "more"}); err != nil {
		t.Fatal(err)
	}
	if err := s.Save(draft, nil); err != nil {
		t.Fatal(err)
	}
	// Hand edit without a bump.
	r3, doc, _ := s.Load(unbumped.ID)
	r3.Constraints = append(r3.Constraints, "edited")
	_ = s.Save(r3, doc)
	// A gap in a record's snapshot chain.
	gap, _ := filepath.Glob(filepath.Join(root, "provenance", ".versions", later.ID, "1-*.yml"))
	if err := os.Remove(gap[0]); err != nil {
		t.Fatal(err)
	}
	// A v1 record.
	write("provenance/prov-2026-0000f0e1.yml", "id: prov-2026-0000f0e1\ntitle: Old\nstatus: draft\ncreated_at: \"2026-01-01\"\nauthor: a\nintent: x\nconstraints: [c]\nsupersedes: \"\"\n")

	// A stale manifest: wrong version for one, no line for another, an extra one.
	lines, err := record.BuildManifest(s)
	if err != nil {
		t.Fatal(err)
	}
	var stale []record.ManifestLine
	for i, l := range lines {
		switch i {
		case 1:
			continue
		case 2:
			l.Version = 7
		}
		stale = append(stale, l)
	}
	stale = append(stale, record.ManifestLine{ID: "prov-2026-0000fffe", Version: 1, ContentHash: "sha256:00"})
	sort.Slice(stale, func(i, j int) bool { return stale[i].ID < stale[j].ID })
	write(record.ManifestPath, string(record.MarshalManifest(stale)))

	shared := newRecord("prov-2026-0000e0e0", record.StatusOpen)
	shared.AffectedScope = []string{"src/**"}
	opts := Options{Root: root, Enforcement: "strict", ManifestPath: record.ManifestPath,
		ExcludePaths: []string{"docs"}, Shared: map[string]*record.Record{shared.ID: shared}}
	return s, opts
}

// keep is the issues a full run reports for the focused ids and the repo.
func keep(res *Result, focus map[string]bool) []Issue {
	var out []Issue
	for _, i := range res.Issues {
		if i.RecordID == "" || focus[i.RecordID] {
			out = append(out, i)
		}
	}
	return out
}

func TestFocusMatchesFullRun(t *testing.T) {
	s, opts := focusRepo(t)
	full, err := Lint(s, opts)
	if err != nil {
		t.Fatal(err)
	}
	seen := map[string]bool{}
	for _, i := range full.Issues {
		seen[i.Rule] = true
	}
	for _, rule := range []string{RuleScopeOverlap, RuleLockedScope, RuleInvalidPattern, RuleScopePath, RuleSpecFile,
		RuleImprintContainment, RuleImplementsUnknown, RuleImplementsPin, RuleUnbumpedChange, RuleVersionSnapshots,
		RuleNotMigrated, RuleScopeConflict, RuleTitleLength, RuleDeadRecord} {
		if !seen[rule] {
			t.Errorf("the fixture never trips %s", rule)
		}
	}

	ids, _ := s.List()
	ids = append(ids, "prov-2026-0000e0e0", "prov-2026-0000fffe", "prov-2026-0000ffff")
	var focuses []map[string]bool
	focuses = append(focuses, map[string]bool{}) // repo-level issues only
	for i, id := range ids {
		focuses = append(focuses, map[string]bool{id: true})
		focuses = append(focuses, map[string]bool{id: true, ids[(i+5)%len(ids)]: true})
	}
	all := map[string]bool{}
	for _, id := range ids {
		all[id] = true
	}
	focuses = append(focuses, all)

	for _, focus := range focuses {
		o := opts
		o.Focus = focus
		got, err := Lint(s, o)
		if err != nil {
			t.Fatalf("focus %v: %v", focus, err)
		}
		want := keep(full, focus)
		if !reflect.DeepEqual(got.Issues, want) {
			t.Errorf("focus %v:\n got  %v\n want %v", focus, got.Issues, want)
		}
		if got.ErrorCount+got.WarningCount+got.HintCount != len(got.Issues) {
			t.Errorf("focus %v: counts do not add up", focus)
		}
	}
}

// A full run and a focused run fail the same way when a record cannot be read.
func TestFocusStillFailsOnUnreadableRecord(t *testing.T) {
	s, opts := focusRepo(t)
	if err := os.WriteFile(filepath.Join(s.Root, "provenance", "prov-2026-0000fbad.yml"), []byte("id: [unclosed\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	_, full := Lint(s, opts)
	opts.Focus = map[string]bool{"prov-2026-0000f001": true}
	_, focused := Lint(s, opts)
	if full == nil || focused == nil || full.Error() != focused.Error() {
		t.Fatalf("full: %v, focused: %v", full, focused)
	}
}

// A snapshot that does not verify fails the manifest check outright, for every
// focus, as it does in a full run.
func TestFocusStillFailsOnTamperedSnapshot(t *testing.T) {
	s, opts := focusRepo(t)
	snaps, _ := filepath.Glob(filepath.Join(s.Root, "provenance", ".versions", "prov-2026-0000f015", "*.yml"))
	data, _ := os.ReadFile(snaps[0])
	_ = os.Chmod(snaps[0], 0o644)
	if err := os.WriteFile(snaps[0], []byte(strings.Replace(string(data), "Do it.", "Did it.", 1)), 0o644); err != nil {
		t.Fatal(err)
	}
	_, full := Lint(s, opts)
	opts.Focus = map[string]bool{"prov-2026-0000f001": true}
	_, focused := Lint(s, opts)
	if full == nil || focused == nil || full.Error() != focused.Error() {
		t.Fatalf("full: %v, focused: %v", full, focused)
	}
	// Without a manifest check the same snapshot is an issue on its record.
	opts.ManifestPath, opts.Focus = "", nil
	res, err := Lint(s, opts)
	if err != nil || rules(res, "prov-2026-0000f015")[RuleVersionSnapshots] == 0 {
		t.Fatalf("tampered snapshot not reported: %v %v", err, res)
	}
}

// The parallel reads must not change what a run reports.
func TestLintIsDeterministic(t *testing.T) {
	s, opts := focusRepo(t)
	first, err := Lint(s, opts)
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 5; i++ {
		again, err := Lint(s, opts)
		if err != nil || !reflect.DeepEqual(first, again) {
			t.Fatalf("run %d differs: %v\n%s", i, err, fmt.Sprint(again.Issues))
		}
	}
}
