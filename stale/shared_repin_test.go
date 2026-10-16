package stale

import (
	"errors"
	"strings"
	"testing"

	"github.com/livecodelife/provenance/record"
)

// Proof for RepinShared: a child of a parent in a shared repo is pinned
// against the synced cache, and says to sync when the parent is not there.

func TestRepinSharedPinsToCachedVersion(t *testing.T) {
	s := tree(t)
	const sp = "prov-2026-0000c001"
	setChild(t, s, "prov-2026-0000a011", func(r *record.Record) { r.Implements, r.ImplementsRepo, r.ImplementsVersion = sp, "shared", 0 })
	shared := map[string]*record.Record{sp: sharedRec(sp, 3)}
	log := FilePinLog{Root: s.Root}
	p, err := RepinShared(s, shared, log, "prov-2026-0000a011", 0, "reviewed", "a@b", now)
	if err != nil {
		t.Fatal(err)
	}
	if p.Parent != sp || p.From != 0 || p.To != 3 {
		t.Fatalf("pin %+v", p)
	}
	r, _, _ := s.Load("prov-2026-0000a011")
	if r.ImplementsVersion != 3 || r.Version != 1 {
		t.Errorf("record pin %d version %d", r.ImplementsVersion, r.Version)
	}
	if got := sharedSummary(t, s, shared); got != "" {
		t.Errorf("still reported after repin: %q", got)
	}
	if _, err := RepinShared(s, shared, log, "prov-2026-0000a011", 4, "again", "a@b", now); err == nil || !strings.Contains(err.Error(), "outside") {
		t.Errorf("version beyond the shared parent's: %v", err)
	}
}

func TestRepinSharedUnsyncedSaysSync(t *testing.T) {
	s := tree(t)
	setChild(t, s, "prov-2026-0000a011", func(r *record.Record) {
		r.Implements, r.ImplementsRepo, r.ImplementsVersion = "prov-2026-0000c001", "shared", 0
	})
	for _, shared := range []map[string]*record.Record{nil, {}} {
		_, err := RepinShared(s, shared, FilePinLog{Root: s.Root}, "prov-2026-0000a011", 0, "r", "a@b", now)
		if !errors.Is(err, ErrNoParent) || !strings.Contains(err.Error(), "provenance sync") || !strings.Contains(err.Error(), "shared") {
			t.Errorf("err = %v, want ErrNoParent naming the repo and 'provenance sync'", err)
		}
	}
	// Repin without the shared map keeps refusing a shared child.
	if _, err := Repin(s, FilePinLog{Root: s.Root}, "prov-2026-0000a011", 0, "r", "a@b", now); !errors.Is(err, ErrNoParent) {
		t.Errorf("Repin on a shared child: %v", err)
	}
}
