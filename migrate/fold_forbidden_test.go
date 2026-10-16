package migrate

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/livecodelife/provenance/record"
)

// A blueprint forbids a file; a later fix to it names that file as affected
// (the fix exists because the blueprint kept the file out of reach). Folding
// the chain must not leave the file in both lists: the later version wins.
func TestFoldDropsForbiddenPathTheFixAffects(t *testing.T) {
	root := t.TempDir()
	writeFile(t, filepath.Join(root, ".linespec.yml"), "provenance:\n  dir: provenance\n")
	d := filepath.Join(root, "provenance")
	bp := strings.Replace(v1("prov-2026-0000c001", "blueprint", "implemented", ""), "forbidden_scope: []", "forbidden_scope:\n  - ui/Graph.tsx\n  - keep/out.go", 1)
	fix := strings.Replace(v1("prov-2026-0000c002", "bug", "implemented", "extends: prov-2026-0000c001\n"), "  - a.go\n", "  - ui/Graph.tsx\n", 1)
	writeFile(t, filepath.Join(d, "prov-2026-0000c001.yml"), bp)
	writeFile(t, filepath.Join(d, "prov-2026-0000c002.yml"), fix)

	repo, err := Load(root, LoadOptions{Git: fakeGit{"prov-2026-0000c001.yml": 100, "prov-2026-0000c002.yml": 200}})
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
	if res.Records != 1 {
		t.Fatalf("fix did not fold: %+v", res)
	}
	got, _, err := store.Load("prov-2026-0000c001")
	if err != nil {
		t.Fatal(err)
	}
	if strings.Join(got.AffectedScope, ",") != "a.go,ui/Graph.tsx" {
		t.Fatalf("affected_scope %v", got.AffectedScope)
	}
	if strings.Join(got.ForbiddenScope, ",") != "keep/out.go" {
		t.Fatalf("forbidden_scope %v: want the fix path dropped and the unrelated one kept", got.ForbiddenScope)
	}
	// The first version still says what it said.
	v1snap, err := store.Snapshot("prov-2026-0000c001", 1)
	if err != nil || strings.Join(v1snap.Content.ForbiddenScope, ",") != "ui/Graph.tsx,keep/out.go" {
		t.Fatalf("version 1 changed: %+v (%v)", v1snap, err)
	}
	checkMigrated(t, root, "provenance", store)
}
