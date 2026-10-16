package stale

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/livecodelife/provenance/evidence"
	"github.com/livecodelife/provenance/migrate"
	"github.com/livecodelife/provenance/record"
)

var now = time.Date(2026, 10, 15, 20, 0, 0, 0, time.UTC)

func rec(id string, typ record.Type, parent string) *record.Record {
	r := &record.Record{Schema: 2, ID: id, Title: id, Type: typ, Status: record.StatusOpen,
		CreatedAt: "2026-10-15", Author: "a@b", Intent: "Intent " + id, Constraints: []string{"c"}, Implements: parent}
	if parent != "" {
		r.ImplementsVersion = 1
	}
	return r
}

// tree: brief B; blueprints P1, P2 implement B; imprints I1, I2 implement P1,
// I3 implements P2. Every record is committed at version 1.
func tree(t *testing.T) *record.FilesStore {
	t.Helper()
	s := &record.FilesStore{Root: t.TempDir()}
	for _, r := range []*record.Record{
		rec("prov-2026-0000b000", record.TypeBrief, ""),
		rec("prov-2026-0000b001", record.TypeBlueprint, "prov-2026-0000b000"),
		rec("prov-2026-0000b002", record.TypeBlueprint, "prov-2026-0000b000"),
		rec("prov-2026-0000a011", record.TypeImprint, "prov-2026-0000b001"),
		rec("prov-2026-0000a012", record.TypeImprint, "prov-2026-0000b001"),
		rec("prov-2026-0000a021", record.TypeImprint, "prov-2026-0000b002"),
	} {
		if _, err := record.Commit(s, r, nil, record.Bump{Author: "a@b", Now: now}); err != nil {
			t.Fatal(err)
		}
	}
	return s
}

func bump(t *testing.T, s *record.FilesStore, id, constraint string) {
	t.Helper()
	r, doc, err := s.Load(id)
	if err != nil {
		t.Fatal(err)
	}
	r.Constraints = append(r.Constraints, constraint)
	if _, err := record.Commit(s, r, doc, record.Bump{Reason: "add " + constraint, Now: now}); err != nil {
		t.Fatal(err)
	}
}

func summary(t *testing.T, s record.Store) string {
	t.Helper()
	fs, err := Find(s)
	if err != nil {
		t.Fatal(err)
	}
	var parts []string
	for _, f := range fs {
		p := strings.TrimPrefix(f.ID, "prov-2026-0000") + ":" + string(f.Kind)
		if f.Cause != "" {
			p += "<" + strings.TrimPrefix(f.Cause, "prov-2026-0000")
		}
		parts = append(parts, p)
	}
	return strings.Join(parts, " ")
}

func TestNothingStale(t *testing.T) {
	if got := summary(t, tree(t)); got != "" {
		t.Fatalf("fresh tree: %s", got)
	}
}

// DESIGN Phase 3 done-check: bumping a brief lists exactly its blueprints
// (direct) and their imprints (transitive), and their evidence shows stale.
func TestBriefBump(t *testing.T) {
	s := tree(t)
	ev := &evidence.Store{Root: s.Root}
	regions := []evidence.Region{{Gen: "sha256:aa", Pkg: "sha256:01"}}
	ids, _ := s.List()
	for _, id := range ids {
		r, _, _ := s.Load(id)
		if err := ev.Append(evidence.New(r, evidence.CurrentFor(r, s.Specs(), regions), true, now.Format(time.RFC3339Nano), "test")); err != nil {
			t.Fatal(err)
		}
	}

	bump(t, s, "prov-2026-0000b000", "new rule")
	want := "a011:transitive<b001 a012:transitive<b001 a021:transitive<b002 b001:direct b002:direct"
	if got := summary(t, s); got != want {
		t.Fatalf("after brief bump:\n got %s\nwant %s", got, want)
	}

	findings, _ := Find(s)
	behind := map[string]bool{}
	for _, f := range findings {
		behind[f.ID] = true
	}
	for _, id := range ids {
		r, _, _ := s.Load(id)
		cur := evidence.CurrentFor(r, s.Specs(), regions)
		cur.Behind = behind[id]
		e, _ := ev.Latest(id, 1)
		v := evidence.Check(e, cur)
		switch {
		case id == "prov-2026-0000b000":
			if v.Verified() || !contains(v.Reasons, "content") {
				t.Errorf("brief's own evidence: %+v", v)
			}
		case v.Verified() || !contains(v.Reasons, "parent"):
			t.Errorf("%s evidence should be stale (parent): %+v", id, v)
		}
	}
}

func TestRepinClearsTransitive(t *testing.T) {
	s := tree(t)
	bump(t, s, "prov-2026-0000b000", "new rule")
	log := FilePinLog{Root: s.Root}

	path := filepath.Join(s.Root, "provenance", "prov-2026-0000b001.yml")
	before, _ := os.ReadFile(path)
	versionsBefore, _ := s.Versions("prov-2026-0000b001")
	r0, _, _ := s.Load("prov-2026-0000b001")
	hashBefore, _ := r0.ContentHash(s.Specs())

	p, err := Repin(s, log, "prov-2026-0000b001", 0, "still correct under the new rule", "a@b", now)
	if err != nil {
		t.Fatal(err)
	}
	if p.From != 1 || p.To != 2 {
		t.Fatalf("pin %+v", p)
	}
	after, _ := os.ReadFile(path)
	if strings.Replace(string(before), "implements_version: 1", "implements_version: 2", 1) != string(after) {
		t.Fatalf("repin changed more than implements_version:\n%s", after)
	}
	r1, _, _ := s.Load("prov-2026-0000b001")
	if h, _ := r1.ContentHash(s.Specs()); h != hashBefore || r1.Version != r0.Version {
		t.Fatal("repin changed content_hash or version")
	}
	if v, _ := s.Versions("prov-2026-0000b001"); len(v) != len(versionsBefore) {
		t.Fatal("repin wrote a snapshot")
	}
	if pins, _ := log.Read("prov-2026-0000b001"); len(pins) != 1 || pins[0].Reason != "still correct under the new rule" {
		t.Fatalf("pin log %+v", pins)
	}

	// Re-pinning b001 clears its imprints; b002's line is unchanged.
	if got := summary(t, s); got != "a021:transitive<b002 b002:direct" {
		t.Fatalf("after repin: %s", got)
	}
	// A new version of b002 makes its imprint directly stale instead.
	bump(t, s, "prov-2026-0000b002", "changed")
	if got := summary(t, s); got != "a021:direct b002:direct" {
		t.Fatalf("after blueprint version: %s", got)
	}
}

func TestRepinRefusals(t *testing.T) {
	s := tree(t)
	log := FilePinLog{Root: s.Root}
	if _, err := Repin(s, log, "prov-2026-0000b001", 0, "", "a", now); !errors.Is(err, ErrNoReason) {
		t.Errorf("no reason: %v", err)
	}
	if _, err := Repin(s, log, "prov-2026-0000b001", 0, "r", "a", now); !errors.Is(err, ErrAlreadyThat) {
		t.Errorf("already current: %v", err)
	}
	if _, err := Repin(s, log, "prov-2026-0000b001", 5, "r", "a", now); err == nil || !strings.Contains(err.Error(), "outside") {
		t.Errorf("out of range: %v", err)
	}
	if _, err := Repin(s, log, "prov-2026-0000b000", 0, "r", "a", now); !errors.Is(err, ErrNoParent) {
		t.Errorf("no parent: %v", err)
	}
	r, doc, _ := s.Load("prov-2026-0000a011")
	r.Status = record.StatusDeprecated
	_ = s.Save(r, doc)
	if _, err := Repin(s, log, "prov-2026-0000a011", 0, "r", "a", now); !errors.Is(err, ErrDeprecated) {
		t.Errorf("deprecated: %v", err)
	}
}

func TestUnknownAndDeprecated(t *testing.T) {
	s := tree(t)
	set := func(id string, f func(r *record.Record)) {
		r, doc, _ := s.Load(id)
		f(r)
		_ = s.Save(r, doc)
	}
	set("prov-2026-0000a011", func(r *record.Record) { r.ImplementsVersion = 0 })
	set("prov-2026-0000a012", func(r *record.Record) { r.ImplementsRepo = "livecodelife-cli" })
	set("prov-2026-0000a021", func(r *record.Record) { r.Implements = "prov-2026-ffffffff" })
	bump(t, s, "prov-2026-0000b000", "x")
	set("prov-2026-0000b002", func(r *record.Record) { r.Status = record.StatusDeprecated })
	if got := summary(t, s); got != "a011:unknown a012:unknown a021:unknown b001:direct" {
		t.Fatalf("got %s", got)
	}
}

func contains(l []string, s string) bool {
	for _, x := range l {
		if x == s {
			return true
		}
	}
	return false
}

// TestMigratedRepos checks Find's invariants on migrate.Apply output (temp
// copies of the 4 repositories, which are only read).
func TestMigratedRepos(t *testing.T) {
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
			os.MkdirAll(filepath.Join(tmp, repo.Dir), 0o755)
			entries, _ := os.ReadDir(filepath.Join(src, repo.Dir))
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
			fs, err := Find(store)
			if err != nil {
				t.Fatal(err)
			}
			kinds := map[Kind]int{}
			directIDs := map[string]bool{}
			for _, f := range fs {
				kinds[f.Kind]++
				if f.Kind == Direct {
					directIDs[f.ID] = true
					if f.Pinned >= f.ParentVersion {
						t.Errorf("%s: direct but pinned %d >= parent %d", f.ID, f.Pinned, f.ParentVersion)
					}
				}
			}
			for _, f := range fs {
				if f.Kind == Transitive && !directIDs[f.Cause] {
					t.Errorf("%s: transitive cause %s is not directly stale", f.ID, f.Cause)
				}
			}
			t.Logf("direct %d, transitive %d, unknown %d", kinds[Direct], kinds[Transitive], kinds[Unknown])
		})
	}
}
