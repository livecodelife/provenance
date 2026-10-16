package stale

import (
	"strings"
	"testing"

	"github.com/livecodelife/provenance/record"
)

// Proof for FindShared: a parent in a shared repo is judged from the synced
// cache, like a local parent (docs: DESIGN.md, When a parent changes).

func sharedRec(id string, version int) *record.Record {
	r := rec(id, record.TypeBlueprint, "")
	r.Version = version
	return r
}

func sharedSummary(t *testing.T, s record.Store, shared map[string]*record.Record) string {
	t.Helper()
	fs, err := FindShared(s, shared)
	if err != nil {
		t.Fatal(err)
	}
	var parts []string
	for _, f := range fs {
		p := strings.TrimPrefix(f.ID, "prov-2026-0000") + ":" + string(f.Kind)
		if f.Kind == Unknown {
			p += "(" + f.Why + ")"
		}
		parts = append(parts, p)
	}
	return strings.Join(parts, " ")
}

func setChild(t *testing.T, s *record.FilesStore, id string, f func(r *record.Record)) {
	t.Helper()
	r, doc, err := s.Load(id)
	if err != nil {
		t.Fatal(err)
	}
	f(r)
	if err := s.Save(r, doc); err != nil {
		t.Fatal(err)
	}
}

func TestSharedParentJudged(t *testing.T) {
	s := tree(t)
	const sp = "prov-2026-0000c001"
	setChild(t, s, "prov-2026-0000a011", func(r *record.Record) { r.Implements, r.ImplementsRepo, r.ImplementsVersion = sp, "shared", 1 })
	setChild(t, s, "prov-2026-0000a012", func(r *record.Record) { r.Implements, r.ImplementsRepo, r.ImplementsVersion = sp, "shared", 3 })
	shared := map[string]*record.Record{sp: sharedRec(sp, 3)}
	// a011 is pinned at 1 against shared v3: direct; a012 is at 3: current.
	if got := sharedSummary(t, s, shared); got != "a011:direct" {
		t.Fatalf("got %q", got)
	}
}

func TestSharedParentNotCachedStaysUnknown(t *testing.T) {
	s := tree(t)
	setChild(t, s, "prov-2026-0000a011", func(r *record.Record) { r.Implements, r.ImplementsRepo = "prov-2026-0000c001", "shared" })
	for _, shared := range []map[string]*record.Record{nil, {}} {
		got := sharedSummary(t, s, shared)
		if !strings.HasPrefix(got, "a011:unknown(parent is in shared repo shared") {
			t.Fatalf("got %q", got)
		}
	}
}

func TestSharedUnpinnedIsUnknown(t *testing.T) {
	s := tree(t)
	const sp = "prov-2026-0000c001"
	setChild(t, s, "prov-2026-0000a011", func(r *record.Record) { r.Implements, r.ImplementsRepo, r.ImplementsVersion = sp, "shared", 0 })
	if got := sharedSummary(t, s, map[string]*record.Record{sp: sharedRec(sp, 2)}); got != "a011:unknown(unpinned)" {
		t.Fatalf("got %q", got)
	}
}

func TestSharedStaleMakesLocalChildTransitive(t *testing.T) {
	s := tree(t)
	const sp = "prov-2026-0000c001"
	// b001 (a blueprint) implements the shared parent and is behind it; a011 and
	// a012 implement b001 and are current, so they are transitive via b001.
	setChild(t, s, "prov-2026-0000b001", func(r *record.Record) { r.Implements, r.ImplementsRepo, r.ImplementsVersion = sp, "shared", 1 })
	if got := sharedSummary(t, s, map[string]*record.Record{sp: sharedRec(sp, 2)}); got != "a011:transitive a012:transitive b001:direct" {
		t.Fatalf("got %q", got)
	}
}

func TestFindWithoutSharedMatchesFindShared(t *testing.T) {
	s := tree(t)
	bump(t, s, "prov-2026-0000b000", "x")
	a, err := Find(s)
	if err != nil {
		t.Fatal(err)
	}
	b, err := FindShared(s, nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(a) != len(b) {
		t.Fatalf("%d vs %d findings", len(a), len(b))
	}
	for i := range a {
		if a[i] != b[i] {
			t.Errorf("finding %d: %+v vs %+v", i, a[i], b[i])
		}
	}
}
