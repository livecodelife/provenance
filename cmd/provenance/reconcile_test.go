package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func writableBit(t *testing.T, path string) bool {
	t.Helper()
	info, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	return info.Mode().Perm()&0o200 != 0
}

// reconcileRepo is checkRepo plus a snapshot made writable and a pins.jsonl,
// all tracked, so reconcile has to lock and unlock real files.
func reconcileRepo(t *testing.T, extraConfig string) (root, cfg, snap string) {
	t.Helper()
	root, cfg = checkRepo(t)
	if extraConfig != "" {
		writeFile(t, cfg, readFile(t, cfg)+extraConfig)
	}
	snaps, _ := filepath.Glob(filepath.Join(root, "provenance", ".versions", chkID, "1-*.yml"))
	if len(snaps) != 1 {
		t.Fatalf("snapshots %v", snaps)
	}
	snap = snaps[0]
	if err := os.Chmod(snap, 0o644); err != nil {
		t.Fatal(err)
	}
	writeFile(t, filepath.Join(root, "provenance", ".versions", chkID, "pins.jsonl"), "{}\n")
	if err := os.Chmod(filepath.Join(root, "b.go"), 0o644); err != nil {
		t.Fatal(err)
	}
	gitIn(t, root, "add", "-A")
	gitIn(t, root, "commit", "-q", "-m", "state ["+chkID+"]")
	return root, cfg, snap
}

func TestReconcileDefaultDeny(t *testing.T) {
	root, cfg, snap := reconcileRepo(t, "")
	out := mustRun(t, "-c", cfg, "reconcile")
	if !strings.Contains(out, "b.go") {
		t.Fatalf("output:\n%s", out)
	}
	for path, want := range map[string]bool{
		"a.go":      true,  // covered by the open record
		"spec.sh":   true,  // covered, and an associated spec
		"b.go":      false, // tracked, uncovered
		"README.md": true,  // exclude_paths
		filepath.Join("provenance", chkID+".yml"):                     true, // the provenance dir
		filepath.Join("provenance", ".versions", chkID, "pins.jsonl"): true, // tool state
		".provenance.yml": true,
	} {
		if got := writableBit(t, filepath.Join(root, path)); got != want {
			t.Errorf("%s writable = %v, want %v", path, got, want)
		}
	}
	if writableBit(t, snap) {
		t.Error("a snapshot file was left writable")
	}
	if again := mustRun(t, "-c", cfg, "reconcile"); strings.Contains(again, "b.go") {
		t.Errorf("reconcile is not idempotent:\n%s", again)
	}
}

func TestReconcileDryRunAndOptOut(t *testing.T) {
	root, cfg, snap := reconcileRepo(t, "")
	out := mustRun(t, "-c", cfg, "reconcile", "--dry-run")
	if !strings.Contains(out, "Would") || !writableBit(t, filepath.Join(root, "b.go")) || writableBit(t, snap) == false {
		t.Fatalf("dry run must change nothing (snapshot still as left):\n%s", out)
	}
	mustRun(t, "-c", cfg, "reconcile")
	writeFile(t, cfg, readFile(t, cfg)+"write_restriction: false\n")
	mustRun(t, "-c", cfg, "reconcile")
	if !writableBit(t, filepath.Join(root, "b.go")) {
		t.Error("write_restriction: false must unlock everything")
	}
	if writableBit(t, snap) {
		t.Error("write_restriction: false must still keep snapshots immutable")
	}
}

func TestReconcileColdStartAndGlobDirs(t *testing.T) {
	root := t.TempDir()
	cfg := filepath.Join(root, ".provenance.yml")
	writeFile(t, cfg, "dir: provenance\n")
	writeFile(t, filepath.Join(root, "x.go"), "package x\n")
	gitRepo(t, root)
	mustRun(t, "-c", cfg, "reconcile")
	if !writableBit(t, filepath.Join(root, "x.go")) {
		t.Fatal("a repo with no records is skipped")
	}

	root, cfg = checkRepo(t)
	rec := filepath.Join(root, "provenance", chkID+".yml")
	writeFile(t, rec, strings.Replace(readFile(t, rec), "  - a.go\n", "  - a.go\n  - newdir/*.go\n", 1))
	dir := filepath.Join(root, "newdir")
	if err := os.MkdirAll(dir, 0o555); err != nil {
		t.Fatal(err)
	}
	mustRun(t, "-c", cfg, "reconcile")
	if info, _ := os.Stat(dir); info.Mode().Perm()&0o300 != 0o300 {
		t.Errorf("glob base directory left %v", info.Mode().Perm())
	}
}
