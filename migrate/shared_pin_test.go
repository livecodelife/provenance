package migrate

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/livecodelife/provenance/record"
	"github.com/livecodelife/provenance/stale"
)

// Proof for pinning shared parents at migration: shared records are read
// from the sync cache dirs (LoadOptions.SharedDirs), implements_repo is
// resolved or kept, each shared pin is reported, and an unsynced cache fails.

const (
	spA  = "prov-2026-0000d001" // in shared repo "alpha" only, at version 3
	spB  = "prov-2026-0000d002" // in alpha and beta: ambiguous
	spC  = "prov-2026-0000d003" // named by implements_repo but absent from the cache
	cRep = "prov-2026-0000c001" // implements spA with implements_repo: alpha
	cId  = "prov-2026-0000c002" // implements spA, no implements_repo
	cAmb = "prov-2026-0000c003" // implements spB, no implements_repo
	cGap = "prov-2026-0000c004" // implements spC with implements_repo: alpha
)

func sharedFixture(t *testing.T) (root string, dirs map[string]string) {
	t.Helper()
	root = t.TempDir()
	writeFile(t, filepath.Join(root, ".linespec.yml"), "provenance:\n  dir: provenance\n  shared_repos:\n"+
		"    - name: alpha\n      url: https://example.com/org/alpha.git\n      dir: provenance\n"+
		"    - name: beta\n      url: https://example.com/org/beta.git\n      dir: provenance\n")
	d := filepath.Join(root, "provenance")
	writeFile(t, filepath.Join(d, cRep+".yml"), v1(cRep, "imprint", "implemented", "implements: "+spA+"\nimplements_repo: alpha\n"))
	writeFile(t, filepath.Join(d, cId+".yml"), v1(cId, "imprint", "implemented", "implements: "+spA+"\n"))
	writeFile(t, filepath.Join(d, cAmb+".yml"), v1(cAmb, "imprint", "implemented", "implements: "+spB+"\n"))
	writeFile(t, filepath.Join(d, cGap+".yml"), v1(cGap, "imprint", "implemented", "implements: "+spC+"\nimplements_repo: alpha\n"))
	alpha, beta := t.TempDir(), t.TempDir()
	writeFile(t, filepath.Join(alpha, spA+".yml"), v1(spA, "blueprint", "implemented", "")+"version: 3\nreason: \"\"\n")
	writeFile(t, filepath.Join(alpha, spB+".yml"), v1(spB, "blueprint", "implemented", ""))
	writeFile(t, filepath.Join(beta, spB+".yml"), v1(spB, "blueprint", "implemented", ""))
	return root, map[string]string{"alpha": alpha, "beta": beta}
}

func migrateShared(t *testing.T, root string, dirs map[string]string, opts Options) (*record.FilesStore, *Result, error) {
	t.Helper()
	repo, err := Load(root, LoadOptions{SharedDirs: dirs, Git: fakeGit{
		cRep + ".yml": 1, cId + ".yml": 2, cAmb + ".yml": 3, cGap + ".yml": 4}})
	if err != nil {
		t.Fatal(err)
	}
	p, err := MakePlan(repo)
	if err != nil {
		t.Fatal(err)
	}
	store := &record.FilesStore{Root: root}
	res, err := Apply(p, opts, store)
	return store, res, err
}

func TestSharedParentsPinnedFromCache(t *testing.T) {
	root, dirs := sharedFixture(t)
	store, res, err := migrateShared(t, root, dirs, Options{})
	if err != nil {
		t.Fatal(err)
	}
	want := map[string][3]string{ // repo, parent, version
		cRep: {"alpha", spA, "3"},
		cId:  {"alpha", spA, "3"},
		cAmb: {"", spB, "0"},
		cGap: {"alpha", spC, "0"},
	}
	for id, w := range want {
		r, _, err := store.Load(id)
		if err != nil {
			t.Fatal(err)
		}
		got := [3]string{r.ImplementsRepo, r.Implements, strconv.Itoa(r.ImplementsVersion)}
		if got != w {
			t.Errorf("%s: repo/parent/version = %v, want %v", id, got, w)
		}
	}
	pins := map[string]SharedPin{}
	for _, sp := range res.SharedPins {
		pins[sp.ID] = sp
	}
	if len(pins) != 2 || pins[cRep] != (SharedPin{ID: cRep, Repo: "alpha", Parent: spA, Version: 3}) || pins[cId].Version != 3 {
		t.Errorf("SharedPins = %+v", res.SharedPins)
	}
	for i := 1; i < len(res.SharedPins); i++ {
		if res.SharedPins[i-1].ID >= res.SharedPins[i].ID {
			t.Errorf("SharedPins not sorted: %+v", res.SharedPins)
		}
	}
}

func TestAmbiguousSharedIdStaysAWarning(t *testing.T) {
	root, dirs := sharedFixture(t)
	repo, err := Load(root, LoadOptions{SharedDirs: dirs})
	if err != nil {
		t.Fatal(err)
	}
	p, err := MakePlan(repo)
	if err != nil {
		t.Fatal(err)
	}
	var dangling []string
	for _, w := range warningsOf(p, "dangling_ref") {
		dangling = append(dangling, fmt.Sprint(w.Fields["src"])+">"+fmt.Sprint(w.Fields["target"]))
	}
	// The ambiguous id and the id named by implements_repo but absent from the
	// cache warn; ids that resolve (spA) do not.
	if got := strings.Join(dangling, ","); got != cAmb+">"+spB+","+cGap+">"+spC {
		t.Errorf("dangling_ref warnings = %q", got)
	}
}

func TestUnsyncedSharedCacheFails(t *testing.T) {
	root, dirs := sharedFixture(t)
	dirs["alpha"] = filepath.Join(t.TempDir(), "missing")
	_, _, err := migrateShared(t, root, dirs, Options{})
	if err == nil || !strings.Contains(err.Error(), "provenance sync") || !strings.Contains(err.Error(), "alpha") {
		t.Fatalf("err = %v, want one naming alpha and 'provenance sync'", err)
	}
	if _, err := os.Stat(filepath.Join(root, "provenance", ".versions")); err == nil {
		t.Error("a failed migrate wrote snapshots")
	}
	store, res, err := migrateShared(t, root, dirs, Options{AllowUnsyncedShared: true})
	if err != nil {
		t.Fatal(err)
	}
	r, _, _ := store.Load(cRep)
	if r.ImplementsRepo != "alpha" || r.Implements != spA || r.ImplementsVersion != 0 {
		t.Errorf("unsynced pin %s %s@%d", r.ImplementsRepo, r.Implements, r.ImplementsVersion)
	}
	// With alpha unsynced, spB is only in beta, so cAmb resolves there.
	if len(res.SharedPins) != 1 || res.SharedPins[0] != (SharedPin{ID: cAmb, Repo: "beta", Parent: spB, Version: 1}) {
		t.Errorf("SharedPins = %+v", res.SharedPins)
	}
}

func TestStaleCleanAfterSharedPins(t *testing.T) {
	root, dirs := sharedFixture(t)
	store, _, err := migrateShared(t, root, dirs, Options{})
	if err != nil {
		t.Fatal(err)
	}
	shared := map[string]*record.Record{}
	for _, d := range dirs {
		ents, _ := os.ReadDir(d)
		for _, e := range ents {
			data, _ := os.ReadFile(filepath.Join(d, e.Name()))
			if rec, err := record.Decode(data); err == nil {
				shared[rec.ID] = rec
			}
		}
	}
	found, err := stale.FindShared(store, shared)
	if err != nil {
		t.Fatal(err)
	}
	for _, f := range found {
		if f.Kind != stale.Unknown {
			t.Errorf("behind after migrate: %+v", f)
		}
		if f.ID == cRep || f.ID == cId {
			t.Errorf("shared child %s still unjudged: %+v", f.ID, f)
		}
	}
}
