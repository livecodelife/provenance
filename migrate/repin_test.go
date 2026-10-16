package migrate

import (
	"path/filepath"
	"testing"

	"github.com/livecodelife/provenance/record"
	"github.com/livecodelife/provenance/stale"
)

// Proof for the pin rule: a child that named a folded version is pinned to
// the survivor's final version, and Apply lists it in Result.Repinned.
func TestChildPinsFinalVersionOfFoldedParent(t *testing.T) {
	root := fixtureRepo(t)
	// a002 implements bp2 itself (position 2 of the chain bp1, bp2, f001).
	writeFile(t, filepath.Join(root, "provenance", "prov-2026-0000a002.yml"), v1("prov-2026-0000a002", "imprint", "implemented", "implements: prov-2026-0000b002\n"))
	repo, err := Load(root, LoadOptions{Git: fakeGit{
		"prov-2026-0000b001.yml": 100, "prov-2026-0000b002.yml": 200, "prov-2026-0000f001.yml": 300,
		"prov-2026-0000f002.yml": 400, "prov-2026-0000a001.yml": 150, "prov-2026-0000a002.yml": 160,
	}})
	if err != nil {
		t.Fatal(err)
	}
	p, err := MakePlan(repo)
	if err != nil {
		t.Fatal(err)
	}
	store := &record.FilesStore{Root: root}
	res, err := Apply(p, Options{}, store)
	if err != nil {
		t.Fatal(err)
	}
	const bp = "prov-2026-0000b002"
	for _, id := range []string{"prov-2026-0000a001", "prov-2026-0000a002", "prov-2026-0000f002"} {
		r, _, err := store.Load(id)
		if err != nil {
			t.Fatal(err)
		}
		if r.Implements != bp || r.ImplementsVersion != 3 {
			t.Errorf("%s pins %s@%d, want %s@3 (final version of the chain)", id, r.Implements, r.ImplementsVersion, bp)
		}
		if r.Version != 1 {
			t.Errorf("%s: re-pinning must not add a version, got %d", id, r.Version)
		}
	}
	want := map[string]Repin{
		"prov-2026-0000a001": {ID: "prov-2026-0000a001", Parent: bp, From: 1, To: 3},
		"prov-2026-0000a002": {ID: "prov-2026-0000a002", Parent: bp, From: 2, To: 3},
		"prov-2026-0000f002": {ID: "prov-2026-0000f002", Parent: bp, From: 2, To: 3},
	}
	if len(res.Repinned) != len(want) {
		t.Fatalf("Repinned = %+v", res.Repinned)
	}
	for i, got := range res.Repinned {
		if want[got.ID] != got {
			t.Errorf("Repinned[%d] = %+v, want %+v", i, got, want[got.ID])
		}
		if i > 0 && res.Repinned[i-1].ID >= got.ID {
			t.Errorf("Repinned not sorted by id: %+v", res.Repinned)
		}
	}
	found, err := stale.Find(store)
	if err != nil {
		t.Fatal(err)
	}
	for _, f := range found {
		if f.Kind != stale.Unknown {
			t.Errorf("stale after migrate: %+v", f)
		}
	}
}

// A child of a parent that folded nothing keeps its pin and is not reported.
func TestUnfoldedParentChildIsNotRepinned(t *testing.T) {
	root := t.TempDir()
	writeFile(t, filepath.Join(root, ".linespec.yml"), "provenance:\n  dir: provenance\n")
	d := filepath.Join(root, "provenance")
	writeFile(t, filepath.Join(d, "prov-2026-0000b001.yml"), v1("prov-2026-0000b001", "blueprint", "implemented", ""))
	writeFile(t, filepath.Join(d, "prov-2026-0000a001.yml"), v1("prov-2026-0000a001", "imprint", "implemented", "implements: prov-2026-0000b001\n"))
	repo, err := Load(root, LoadOptions{Git: fakeGit{"prov-2026-0000b001.yml": 1, "prov-2026-0000a001.yml": 2}})
	if err != nil {
		t.Fatal(err)
	}
	p, err := MakePlan(repo)
	if err != nil {
		t.Fatal(err)
	}
	store := &record.FilesStore{Root: root}
	res, err := Apply(p, Options{}, store)
	if err != nil {
		t.Fatal(err)
	}
	r, _, _ := store.Load("prov-2026-0000a001")
	if r.ImplementsVersion != 1 || len(res.Repinned) != 0 {
		t.Fatalf("pin %d, repinned %+v", r.ImplementsVersion, res.Repinned)
	}
}
