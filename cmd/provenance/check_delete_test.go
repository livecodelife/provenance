package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// moveRepo is checkRepo with the record revised to version 2: a.go left its
// scope and c.go joined it (one version, not three).
func moveRepo(t *testing.T) (root, cfg string) {
	t.Helper()
	root, cfg = checkRepo(t)
	mustRun(t, "-c", cfg, "edit", chkID, "--remove-scope", "a.go", "--add-scope", "c.go", "--reason", "move a.go to c.go")
	return root, cfg
}

func TestCheckAllowsDeletingAFileThePreviousVersionCovered(t *testing.T) {
	root, cfg := moveRepo(t)
	if err := os.Remove(filepath.Join(root, "a.go")); err != nil {
		t.Fatal(err)
	}
	writeFile(t, filepath.Join(root, "c.go"), "package a\n")
	gitIn(t, root, "add", "-A")
	if out, ok := check(t, cfg, "Move a.go to c.go ["+chkID+"@2]"); !ok {
		t.Fatalf("a move in one version must pass:\n%s", out)
	}
	gitIn(t, root, "commit", "-q", "-m", "Move a.go to c.go ["+chkID+"@2]")
	if out := mustRun(t, "-c", cfg, "check", "--commit", "HEAD"); !strings.Contains(out, "✓") {
		t.Fatalf("--commit agrees:\n%s", out)
	}
	if out, err := run(t, "-c", cfg, "check", "--range", "HEAD~1..HEAD"); err != nil {
		t.Fatalf("--range agrees:\n%s", out)
	}
}

func TestCheckStillRefusesOtherChangesToPathsOutsideTheScope(t *testing.T) {
	root, cfg := moveRepo(t)
	// Deleting a file no version of the record covered is still a violation.
	if err := os.Remove(filepath.Join(root, "b.go")); err != nil {
		t.Fatal(err)
	}
	gitIn(t, root, "add", "-A")
	if out, ok := check(t, cfg, "Drop b.go ["+chkID+"@2]"); ok || !strings.Contains(out, "b.go") {
		t.Fatalf("deleting a never-covered file passed:\n%s", out)
	}
	gitIn(t, root, "reset", "-q")
	if err := os.WriteFile(filepath.Join(root, "b.go"), []byte("package b\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	// Modifying (not deleting) a file that left the scope is still a violation. (The write
	// lock already made a.go read-only when it left the scope; the check must refuse the
	// change even when that lock is bypassed.)
	if err := os.Chmod(filepath.Join(root, "a.go"), 0o644); err != nil {
		t.Fatal(err)
	}
	writeFile(t, filepath.Join(root, "a.go"), "package a // edited after leaving the scope\n")
	gitIn(t, root, "add", "-A")
	if out, ok := check(t, cfg, "Touch a.go ["+chkID+"@2]"); ok || !strings.Contains(out, "a.go") {
		t.Fatalf("modifying a file that left the scope passed:\n%s", out)
	}
}
