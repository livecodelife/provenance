package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const (
	staleParent = "prov-2026-0000e001"
	staleChild  = "prov-2026-0000e002"
)

// staleRepo has a parent (opened, then revised to version 2) and an imprint
// pinned to its version 1.
func staleRepo(t *testing.T) (root, cfg string) {
	t.Helper()
	child := strings.Replace(v2Draft(staleChild, "draft"), "type: blueprint", "type: imprint\nimplements: "+staleParent+"\nimplements_version: 1", 1)
	root, cfg = v2Repo(t, false, map[string]string{staleParent: v2Draft(staleParent, "draft"), staleChild: child})
	mustRun(t, "-c", cfg, "open", "--record", staleParent)
	mustRun(t, "-c", cfg, "open", "--record", staleChild)
	editorScript(t, root, "s/^intent: Do the thing$/intent: Do the other thing/")
	mustRun(t, "-c", cfg, "revise", staleParent, "--reason", "parent moved")
	return root, cfg
}

func TestStaleListsAndExitCode(t *testing.T) {
	_, cfg := staleRepo(t)
	out := mustRun(t, "-c", cfg, "stale")
	want := staleChild + "@1  pinned to " + staleParent + "@1, parent is at 2  (direct)"
	if !strings.Contains(out, want) || !strings.Contains(out, "provenance repin") {
		t.Fatalf("output:\n%s", out)
	}
	js := mustRun(t, "-c", cfg, "stale", "--json")
	if !strings.Contains(js, `"kind": "direct"`) || !strings.Contains(js, `"parent_version": 2`) {
		t.Fatalf("json:\n%s", js)
	}
	if _, err := run(t, "-c", cfg, "stale", "--exit-code"); err == nil {
		t.Fatal("--exit-code must fail while a record is stale")
	}
}

func TestRepinClearsWithoutVersion(t *testing.T) {
	root, cfg := staleRepo(t)
	if _, err := run(t, "-c", cfg, "repin", staleChild); err == nil {
		t.Fatal("repin without --reason must fail")
	}
	out := mustRun(t, "-c", cfg, "repin", staleChild, "--reason", "parent change does not affect this imprint")
	if !strings.Contains(out, "now pinned to "+staleParent+"@2 (was 1)") || !strings.Contains(out, "["+staleChild+"@1]") {
		t.Fatalf("output:\n%s", out)
	}
	if got := versions(t, root, staleChild); len(got) != 1 {
		t.Fatalf("repin wrote a version: %v", got)
	}
	if rec := load(t, root, staleChild); rec.ImplementsVersion != 2 || rec.Version != 1 {
		t.Fatalf("record %+v", rec)
	}
	pins, err := os.ReadFile(filepath.Join(root, "provenance", ".versions", staleChild, "pins.jsonl"))
	if err != nil || !strings.Contains(string(pins), "does not affect") {
		t.Fatalf("pin log %q (%v)", pins, err)
	}
	if out := mustRun(t, "-c", cfg, "stale", "--exit-code"); !strings.Contains(out, "No stale records") {
		t.Fatalf("still stale:\n%s", out)
	}
	if _, err := run(t, "-c", cfg, "repin", staleChild, "--reason", "again"); err == nil {
		t.Fatal("repinning to the version it already has must fail")
	}
}
