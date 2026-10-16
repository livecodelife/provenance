package migrate

import (
	"path/filepath"
	"testing"

	"github.com/livecodelife/provenance/record"
)

// Proof for the per-id fallback: a workspace checkout of a shared repo wins
// for the ids it has, and the sync cache supplies the ids it lacks.
func TestCheckoutLacksIdCacheFillsIn(t *testing.T) {
	root, dirs := sharedFixture(t)
	// A sibling checkout of alpha (named by the url's last segment) that holds
	// spB at version 1 but not spA; the cache has spA at version 3.
	ws := t.TempDir()
	writeFile(t, filepath.Join(ws, "alpha", "provenance", spB+".yml"), v1(spB, "blueprint", "implemented", ""))
	repo, err := Load(root, LoadOptions{Workspace: ws, SharedDirs: dirs, Git: fakeGit{
		cRep + ".yml": 1, cId + ".yml": 2, cAmb + ".yml": 3, cGap + ".yml": 4}})
	if err != nil {
		t.Fatal(err)
	}
	if got := repo.SharedRepoOf[spA]; got != "alpha" {
		t.Fatalf("spA (only in the cache) resolves to %q, want alpha", got)
	}
	p, err := MakePlan(repo)
	if err != nil {
		t.Fatal(err)
	}
	store := &record.FilesStore{Root: root}
	if _, err := Apply(p, Options{}, store); err != nil {
		t.Fatal(err)
	}
	for _, id := range []string{cRep, cId} {
		r, _, _ := store.Load(id)
		if r.ImplementsRepo != "alpha" || r.ImplementsVersion != 3 {
			t.Errorf("%s pins %s@%d, want alpha@3 from the cache", id, r.ImplementsRepo, r.ImplementsVersion)
		}
	}
}

func TestCheckoutWinsForIdsItHas(t *testing.T) {
	root, dirs := sharedFixture(t)
	ws := t.TempDir()
	// The checkout's spA is at version 5, the cache's at 3: the checkout wins.
	writeFile(t, filepath.Join(ws, "alpha", "provenance", spA+".yml"), v1(spA, "blueprint", "implemented", "")+"version: 5\nreason: \"\"\n")
	repo, err := Load(root, LoadOptions{Workspace: ws, SharedDirs: dirs})
	if err != nil {
		t.Fatal(err)
	}
	if v := repo.Shared[spA].Rec.Version; v != 5 {
		t.Errorf("spA version %d, want 5 from the checkout", v)
	}
	if len(repo.SharedMissing) != 0 {
		t.Errorf("SharedMissing = %v", repo.SharedMissing)
	}
}
