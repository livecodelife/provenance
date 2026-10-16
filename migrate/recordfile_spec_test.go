package migrate

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/livecodelife/provenance/record"
)

// Proof for prov-2026-4234f0bf: a migrated record whose associated spec is a
// record file, another record's or its own, keeps a stable content_hash even
// though migrate rewrites record files.
func TestRecordFileSpecsStayStable(t *testing.T) {
	root := t.TempDir()
	writeFile(t, filepath.Join(root, ".linespec.yml"), "provenance:\n  dir: provenance\n")
	d := filepath.Join(root, "provenance")
	other := "prov-2026-0000c001"
	self := "prov-2026-0000c002"
	writeFile(t, filepath.Join(d, other+".yml"), v1(other, "blueprint", "implemented", ""))
	body := v1(self, "imprint", "implemented", "implements: "+other+"\n")
	spec := "  - path: provenance/%s.yml\n    type: linespec\n"
	body = strings.Replace(body, "associated_specs: []\n", "associated_specs:\n"+strings.ReplaceAll(spec, "%s", other)+strings.ReplaceAll(spec, "%s", self), 1)
	writeFile(t, filepath.Join(d, self+".yml"), body)

	repo, err := Load(root, LoadOptions{Git: fakeGit{other + ".yml": 100, self + ".yml": 200}})
	if err != nil {
		t.Fatal(err)
	}
	p, err := MakePlan(repo)
	if err != nil {
		t.Fatal(err)
	}
	store := &record.FilesStore{Root: root}
	if _, err := Apply(p, Options{}, store); err != nil {
		t.Fatal(err)
	}
	// checkMigrated compares each record's recomputed content_hash with its head snapshot.
	checkMigrated(t, root, "provenance", store)
}
