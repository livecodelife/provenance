package main

import (
	"strings"
	"testing"
)

func TestAddScopeFromTaggedCommits(t *testing.T) {
	root, cfg := checkRepo(t)
	stage(t, root, map[string]string{"b.go": "package b // two\n", "c.go": "package c\n", "README.md": "docs\n", "a.go": "package a // one\n"})
	gitIn(t, root, "commit", "-q", "-m", "Work ["+chkID+"]")

	out := mustRun(t, "-c", cfg, "add-scope", "--record", chkID, "--dry-run")
	if !strings.Contains(out, "b.go") || !strings.Contains(out, "c.go") || strings.Contains(out, "a.go") || strings.Contains(out, "README.md") {
		t.Fatalf("dry run lists the new, non-excluded, uncovered files only:\n%s", out)
	}
	if len(versions(t, root, chkID)) != 1 {
		t.Fatal("a dry run wrote a version")
	}

	out = mustRun(t, "-c", cfg, "add-scope", "--record", chkID)
	if !strings.Contains(out, chkID+" is now version 2") {
		t.Fatalf("output:\n%s", out)
	}
	rec := load(t, root, chkID)
	got := strings.Join(rec.AffectedScope, ",")
	if got != "a.go,spec.sh,b.go,c.go" || !strings.Contains(rec.Reason, "add-scope") {
		t.Fatalf("scope %q reason %q", got, rec.Reason)
	}

	if _, err := run(t, "-c", cfg, "add-scope", "--record", chkID); err == nil {
		t.Fatal("nothing new to add must be an error")
	}
	if _, err := run(t, "-c", cfg, "add-scope"); err == nil {
		t.Fatal("--record is required")
	}
}
