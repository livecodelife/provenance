package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCompileCheckReportsAStaleManifest(t *testing.T) {
	root, cfg := openedRepo(t, draftID)
	if out := mustRun(t, "-c", cfg, "compile", "--check"); !strings.Contains(out, "up to date (1 records)") {
		t.Fatalf("a current manifest passes --check:\n%s", out)
	}
	manifest := filepath.Join(root, manifestPath)
	good := readFile(t, manifest)
	writeFile(t, manifest, good+"{\"id\":\"prov-2026-0000ffff\",\"version\":9}\n")
	out, err := run(t, "-c", cfg, "compile", "--check")
	if err == nil {
		t.Fatalf("a stale manifest passed --check:\n%s", out)
	}
	for _, want := range []string{"provenance compile", "git add .provenance/versions.jsonl", "git commit"} {
		if !strings.Contains(out, want) {
			t.Errorf("--check output misses %q:\n%s", want, out)
		}
	}
	if readFile(t, manifest) == good {
		t.Error("--check rewrote the manifest")
	}
	if err := os.Remove(manifest); err != nil {
		t.Fatal(err)
	}
	if _, err := run(t, "-c", cfg, "compile", "--check"); err == nil {
		t.Fatal("a missing manifest passed --check")
	}
}
