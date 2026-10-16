package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCompileRebuildsManifest(t *testing.T) {
	root, cfg := openedRepo(t, draftID)
	manifest := filepath.Join(root, manifestPath)
	good := readFile(t, manifest)
	if out := mustRun(t, "-c", cfg, "compile"); !strings.Contains(out, "up to date (1 records)") {
		t.Fatalf("already current:\n%s", out)
	}
	if err := os.Remove(manifest); err != nil {
		t.Fatal(err)
	}
	if out := mustRun(t, "-c", cfg, "compile"); !strings.Contains(out, "compiled (1 records)") || readFile(t, manifest) != good {
		t.Fatalf("a missing manifest is rebuilt from the snapshots:\n%s", out)
	}
	writeFile(t, manifest, good+"{\"id\":\"prov-2026-0000ffff\",\"version\":9}\n")
	if out := mustRun(t, "-c", cfg, "compile"); !strings.Contains(out, "compiled") || readFile(t, manifest) != good {
		t.Fatalf("a wrong manifest is repaired:\n%s", out)
	}
}
