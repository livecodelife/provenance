package lint

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/livecodelife/provenance/migrate"
	"github.com/livecodelife/provenance/record"
)

func newRecord(id string, status record.Status) *record.Record {
	return &record.Record{
		Schema: 2, ID: id, Title: "T", Type: record.TypeBlueprint, Status: status, Version: 1,
		CreatedAt: "2026-10-15", Author: "a@b", Intent: "Do it.", Constraints: []string{"C"},
	}
}

func rules(res *Result, id string) map[string]int {
	out := map[string]int{}
	for _, i := range res.Issues {
		if id == "" || i.RecordID == id {
			out[i.Rule]++
		}
	}
	return out
}

func lintStore(t *testing.T, s *record.FilesStore) *Result {
	t.Helper()
	res, err := Lint(s, Options{Root: s.Root, Enforcement: "strict"})
	if err != nil {
		t.Fatal(err)
	}
	return res
}

func TestVersionRules(t *testing.T) {
	s := &record.FilesStore{Root: t.TempDir()}
	draft := newRecord("prov-2026-0000d001", record.StatusDraft)
	if err := s.Save(draft, nil); err != nil {
		t.Fatal(err)
	}
	if n := rules(lintStore(t, s), draft.ID)[RuleVersionSnapshots]; n != 0 {
		t.Fatalf("draft without snapshots flagged %d times", n)
	}

	r := newRecord("prov-2026-0000a001", record.StatusOpen)
	_ = s.Save(r, nil)
	if n := rules(lintStore(t, s), r.ID)[RuleVersionSnapshots]; n != 1 {
		t.Fatalf("open record without snapshot: PROV026 x%d, want 1", n)
	}
	if _, err := record.Commit(s, r, nil, record.Bump{Author: "a@b"}); err != nil {
		t.Fatal(err)
	}
	if got := rules(lintStore(t, s), r.ID); got[RuleUnbumpedChange]+got[RuleVersionSnapshots]+got[RuleBumpWithoutReason] != 0 {
		t.Fatalf("clean v1 flagged: %v", got)
	}

	// Hand edit without a bump.
	r, doc, _ := s.Load(r.ID)
	r.Constraints = append(r.Constraints, "New")
	_ = s.Save(r, doc)
	if n := rules(lintStore(t, s), r.ID)[RuleUnbumpedChange]; n != 1 {
		t.Fatalf("unbumped change: PROV024 x%d, want 1", n)
	}
	// A label edit is not a content change.
	r, doc, _ = s.Load(r.ID)
	r.Constraints = r.Constraints[:1]
	r.Title = "Renamed"
	_ = s.Save(r, doc)
	if n := rules(lintStore(t, s), r.ID)[RuleUnbumpedChange]; n != 0 {
		t.Fatal("title edit flagged as a content change")
	}

	// Revise properly, then strip the reason.
	r, doc, _ = s.Load(r.ID)
	r.Constraints = append(r.Constraints, "New")
	if _, err := record.Commit(s, r, doc, record.Bump{Reason: "Add New"}); err != nil {
		t.Fatal(err)
	}
	if got := rules(lintStore(t, s), r.ID); got[RuleUnbumpedChange]+got[RuleVersionSnapshots]+got[RuleBumpWithoutReason] != 0 {
		t.Fatalf("clean v2 flagged: %v", got)
	}
	r, doc, _ = s.Load(r.ID)
	r.Reason = ""
	_ = s.Save(r, doc)
	if n := rules(lintStore(t, s), r.ID)[RuleBumpWithoutReason]; n != 1 {
		t.Fatalf("PROV025 x%d, want 1", n)
	}

	// Version ahead of the snapshots.
	r, doc, _ = s.Load(r.ID)
	r.Reason, r.Version = "x", 3
	_ = s.Save(r, doc)
	if n := rules(lintStore(t, s), r.ID)[RuleVersionSnapshots]; n != 1 {
		t.Fatalf("version ahead: PROV026 x%d, want 1", n)
	}
}

func TestImplementsPin(t *testing.T) {
	s := &record.FilesStore{Root: t.TempDir()}
	bp := newRecord("prov-2026-0000b001", record.StatusDraft)
	bp.Version = 2
	bp.Reason = "r"
	im := newRecord("prov-2026-0000c001", record.StatusDraft)
	im.Type, im.Implements, im.ImplementsVersion = record.TypeImprint, bp.ID, 3
	stray := newRecord("prov-2026-0000c002", record.StatusDraft)
	stray.Type, stray.Implements, stray.ImplementsRepo = record.TypeImprint, bp.ID, "nowhere"
	for _, r := range []*record.Record{bp, im, stray} {
		_ = s.Save(r, nil)
	}
	res := lintStore(t, s)
	if rules(res, im.ID)[RuleImplementsPin] != 1 || rules(res, stray.ID)[RuleImplementsPin] != 1 {
		t.Fatalf("PROV027: %v / %v", rules(res, im.ID), rules(res, stray.ID))
	}
	im.ImplementsVersion = 2
	_ = s.Save(im, nil)
	res, _ = Lint(s, Options{Root: s.Root, SharedRepos: []string{"nowhere"}})
	if rules(res, im.ID)[RuleImplementsPin]+rules(res, stray.ID)[RuleImplementsPin] != 0 {
		t.Fatal("valid pins flagged")
	}
}

func TestV1AndOverlap(t *testing.T) {
	root := t.TempDir()
	s := &record.FilesStore{Root: root}
	os.MkdirAll(filepath.Join(root, "provenance"), 0o755)
	os.WriteFile(filepath.Join(root, "provenance", "prov-2026-0000e001.yml"), []byte("id: prov-2026-0000e001\ntitle: Old\nstatus: draft\ncreated_at: \"2026-01-01\"\nauthor: a\nintent: x\nconstraints: [c]\nsupersedes: \"\"\n"), 0o644)
	a := newRecord("prov-2026-0000a002", record.StatusOpen)
	a.AffectedScope = []string{"src/**"}
	b := newRecord("prov-2026-0000a003", record.StatusOpen)
	b.AffectedScope = []string{"src/x.go"}
	_ = s.Save(a, nil)
	_ = s.Save(b, nil)
	res := lintStore(t, s)
	if rules(res, "prov-2026-0000e001")[RuleNotMigrated] != 1 {
		t.Fatal("v1 record not flagged PROV028")
	}
	var overlap []Issue
	for _, i := range res.Issues {
		if i.Rule == RuleScopeOverlap {
			overlap = append(overlap, i)
		}
	}
	if len(overlap) != 1 || overlap[0].Severity != SeverityWarning || overlap[0].RecordID != b.ID {
		t.Fatalf("PROV015: %+v", overlap)
	}
}

func TestParseCommitTags(t *testing.T) {
	tags := ParseCommitTags("Fix [prov-2026-1a2b3c4d@3] and [prov-2026-001] but not [prov-2026-1a2b3c4d@0] or [prov-26-x] [prov-2026-1a2b3c4d-api@12]")
	got := make([]string, len(tags))
	for i, tg := range tags {
		got[i] = tg.String()
	}
	if strings.Join(got, " ") != "[prov-2026-1a2b3c4d@3] [prov-2026-001] [prov-2026-1a2b3c4d-api@12]" {
		t.Fatalf("tags %v", got)
	}
}

func TestCheckTransition(t *testing.T) {
	ok := [][2]record.Status{{record.StatusDraft, record.StatusOpen}, {record.StatusOpen, record.StatusImplemented}, {record.StatusImplemented, record.StatusDeprecated}}
	for _, p := range ok {
		if err := CheckTransition(p[0], p[1], false); err != nil {
			t.Errorf("%s->%s: %v", p[0], p[1], err)
		}
	}
	if CheckTransition(record.StatusImplemented, record.StatusOpen, false) == nil {
		t.Error("reopening without a bump allowed")
	}
	if CheckTransition(record.StatusImplemented, record.StatusOpen, true) != nil {
		t.Error("reopening with a bump refused")
	}
	if err := CheckTransition(record.StatusDeprecated, record.StatusOpen, true); err == nil || !strings.Contains(err.Error(), "terminal") {
		t.Errorf("deprecated not terminal: %v", err)
	}
	if CheckTransition(record.StatusDraft, record.StatusImplemented, false) == nil {
		t.Error("draft -> implemented allowed")
	}
}

// TestMigratedOutputIsClean: on migrate.Apply output (temp copies only), the
// v2 version rules report nothing.
func TestMigratedOutputIsClean(t *testing.T) {
	ws := os.Getenv("PROVENANCE_REAL_REPOS")
	if ws == "" {
		t.Skip("set PROVENANCE_REAL_REPOS to run")
	}
	for _, name := range []string{"linespec", "sedum", "scion", "linespec-provenance-ui"} {
		t.Run(name, func(t *testing.T) {
			src := filepath.Join(ws, name)
			repo, err := migrate.Load(src, migrate.LoadOptions{Workspace: ws})
			if err != nil {
				t.Fatal(err)
			}
			p, err := migrate.MakePlan(repo)
			if err != nil {
				t.Fatal(err)
			}
			tmp := t.TempDir()
			entries, _ := os.ReadDir(filepath.Join(src, repo.Dir))
			os.MkdirAll(filepath.Join(tmp, repo.Dir), 0o755)
			for _, e := range entries {
				if !e.IsDir() && strings.HasSuffix(e.Name(), ".yml") {
					data, _ := os.ReadFile(filepath.Join(src, repo.Dir, e.Name()))
					os.WriteFile(filepath.Join(tmp, repo.Dir, e.Name()), data, 0o644)
				}
			}
			store := &record.FilesStore{Root: tmp, Dir: repo.Dir, SpecRoot: src}
			if _, err := migrate.Apply(p, migrate.Options{}, store); err != nil {
				t.Fatal(err)
			}
			res, err := Lint(store, Options{Root: src, Enforcement: "strict", SharedRepos: sharedNames(repo)})
			if err != nil {
				t.Fatal(err)
			}
			got := rules(res, "")
			for _, r := range []string{RuleUnbumpedChange, RuleBumpWithoutReason, RuleVersionSnapshots, RuleNotMigrated} {
				if got[r] != 0 {
					for _, i := range res.Issues {
						if i.Rule == r {
							t.Errorf("%s %s: %s", i.Rule, i.RecordID, i.Message)
							break
						}
					}
				}
			}
			t.Logf("%d records linted; PROV027 %d, PROV015 %d", res.Records, got[RuleImplementsPin], got[RuleScopeOverlap])
		})
	}
}

func sharedNames(r *migrate.Repo) []string {
	var out []string
	for _, s := range r.Config.SharedRepos {
		out = append(out, s.Name)
	}
	return out
}
