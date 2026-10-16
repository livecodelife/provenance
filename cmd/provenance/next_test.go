package main

import (
	"encoding/json"
	"path/filepath"
	"strings"
	"testing"
)

func nextKind(t *testing.T, cfg string, args ...string) (string, string) {
	t.Helper()
	out := mustRun(t, append([]string{"-c", cfg, "next", "--json"}, args...)...)
	var doc struct {
		Actions []struct {
			Kind    string `json:"kind"`
			Command string `json:"command"`
		} `json:"actions"`
	}
	if err := json.Unmarshal([]byte(out), &doc); err != nil || len(doc.Actions) == 0 {
		t.Fatalf("json (%v):\n%s", err, out)
	}
	return doc.Actions[0].Kind, doc.Actions[0].Command
}

func TestNextAdvice(t *testing.T) {
	// No record governs b.go; the open record chkID governs a.go.
	root, cfg := checkRepo(t)
	if kind, _ := nextKind(t, cfg, "b.go"); kind != "create" {
		t.Errorf("ungoverned file: %s", kind)
	}
	if kind, cmd := nextKind(t, cfg, "a.go", "b.go"); kind != "add_scope" || !strings.Contains(cmd, "--add-scope b.go") {
		t.Errorf("a file the open record does not cover: %s %s", kind, cmd)
	}
	if kind, _ := nextKind(t, cfg, "a.go"); kind != "complete" {
		t.Errorf("covered file, nothing staged: %s", kind)
	}
	stage(t, root, map[string]string{"a.go": "package a // changed\n"})
	if kind, cmd := nextKind(t, cfg); kind != "commit" || !strings.Contains(cmd, "["+chkID+"@1]") {
		t.Errorf("staged work: %s %s", kind, cmd)
	}

	out := mustRun(t, "-c", cfg, "next", "a.go")
	if !strings.Contains(out, "NEXT: commit") && !strings.Contains(out, "NEXT: complete") {
		t.Errorf("human output:\n%s", out)
	}
	if _, err := run(t, "-c", cfg, "next", "--format", "xml"); err == nil {
		t.Error("unknown format accepted")
	}
}

func TestNextDraftAndImplemented(t *testing.T) {
	root, cfg := v2Repo(t, false, map[string]string{draftID: v2Draft(draftID, "draft")})
	if kind, cmd := nextKind(t, cfg, "a.go"); kind != "open" || !strings.Contains(cmd, "open --record "+draftID) {
		t.Errorf("draft governs: %s %s", kind, cmd)
	}
	_ = root

	root, cfg = completeRepo(t, "run_associated_specs_on_complete: true\n", "true\n")
	mustRun(t, "-c", cfg, "complete", "--record", chkID)
	if kind, cmd := nextKind(t, cfg, "a.go"); kind != "revise" || !strings.Contains(cmd, "revise "+chkID) {
		t.Errorf("implemented record governs: %s %s", kind, cmd)
	}
	writeFile(t, filepath.Join(root, "x.txt"), "x\n")
	if kind, _ := nextKind(t, cfg); kind != "create" {
		t.Errorf("untracked file with no governing record: %s", kind)
	}
}
