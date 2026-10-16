package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCheckExemptsNonRecordFilesUnderTheProvenanceDir(t *testing.T) {
	root, cfg := checkRepo(t)
	stage(t, root, map[string]string{"provenance/.gitkeep": ""})
	if out, ok := check(t, cfg, "Keep the dir ["+chkID+"@1]"); !ok {
		t.Fatalf(".gitkeep alone under a tag must pass:\n%s", out)
	}
	stage(t, root, map[string]string{"a.go": "package a // changed\n", "provenance/notes.txt": "n\n"})
	if out, ok := check(t, cfg, "Change a ["+chkID+"@1]"); !ok {
		t.Fatalf(".gitkeep and a note next to an in-scope file must pass:\n%s", out)
	}
	gitIn(t, root, "commit", "-q", "-m", "Change a ["+chkID+"@1]")
	if out := mustRun(t, "-c", cfg, "check", "--commit", "HEAD"); !strings.Contains(out, "✓") {
		t.Fatalf("--commit agrees:\n%s", out)
	}
}

func TestCheckStillEnforcesRecordFilesAndPathsOutsideTheDir(t *testing.T) {
	root, cfg := checkRepo(t)
	stage(t, root, map[string]string{"provenance/.gitkeep": "", "b.go": "package b // changed\n"})
	if out, ok := check(t, cfg, "Touch b ["+chkID+"@1]"); ok || !strings.Contains(out, "b.go") || strings.Contains(out, ".gitkeep") {
		t.Fatalf("an out-of-scope file outside the dir must fail, and only it:\n%s", out)
	}
	gitIn(t, root, "reset", "-q")
	mustRun(t, "-c", cfg, "complete", "--record", chkID)
	gitIn(t, root, "add", "-A")
	gitIn(t, root, "commit", "-q", "-m", "Complete ["+chkID+"@1]")
	rec := filepath.Join(root, "provenance", chkID+".yml")
	data, err := os.ReadFile(rec)
	if err != nil {
		t.Fatal(err)
	}
	stage(t, root, map[string]string{"provenance/" + chkID + ".yml": string(data) + "# edited\n"})
	if out, ok := check(t, cfg, "Edit the record ["+chkID+"@1]"); ok || !strings.Contains(out, "revise") {
		t.Fatalf("changing an implemented record's file must still demand a revise:\n%s", out)
	}
}
