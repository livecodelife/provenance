package main

import (
	"path/filepath"
	"strings"
	"testing"
)

// Proof for the migrate report line that lists re-pinned children.
func TestMigrateReportsRepinnedChildren(t *testing.T) {
	root := migrateFixture(t)
	out := mustRun(t, "-c", filepath.Join(root, ".linespec.yml"), "migrate")
	for _, want := range []string{
		"Re-pinned to the final version of a folded parent (pins only, no new version): 1",
		"    " + imp + "  " + bpNew + "@1 → @2",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("migrate output lacks %q:\n%s", want, out)
		}
	}
}
