package lint

import (
	"os"
	"path/filepath"
	"testing"
)

// migratedToV2 reports whether a repository has been migrated to v2 (it has a
// .provenance.yml). `linespec provenance lint` no longer reads such a repo's
// records, so there is nothing to compare: the v1 parity of the carried rules
// was proven on these repos before they migrated (prov-2026-e135b1f5).
func migratedToV2(root string) bool {
	_, err := os.Stat(filepath.Join(root, ".provenance.yml"))
	return err == nil
}

func TestMigratedToV2(t *testing.T) {
	dir := t.TempDir()
	if migratedToV2(dir) {
		t.Fatal("an empty directory is not migrated")
	}
	if err := os.WriteFile(filepath.Join(dir, ".linespec.yml"), nil, 0o644); err != nil {
		t.Fatal(err)
	}
	if migratedToV2(dir) {
		t.Fatal(".linespec.yml alone is a v1 repo")
	}
	if err := os.WriteFile(filepath.Join(dir, ".provenance.yml"), nil, 0o644); err != nil {
		t.Fatal(err)
	}
	if !migratedToV2(dir) {
		t.Fatal(".provenance.yml marks a v2 repo")
	}
}
