package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const (
	pinsPath = "provenance/.versions/" + chkID + "/pins.jsonl"
	pinsOne  = `{"child":"` + chkID + `","parent":"prov-2026-0000aaaa","from":1,"to":2}` + "\n"
	pinsTwo  = `{"child":"` + chkID + `","parent":"prov-2026-0000aaaa","from":2,"to":3}` + "\n"
)

func commitPins(t *testing.T, root string) {
	t.Helper()
	stage(t, root, map[string]string{pinsPath: pinsOne})
	gitIn(t, root, "commit", "-q", "-m", "First repin ["+chkID+"@1]")
}

func TestCheckAllowsAppendingToAPinLog(t *testing.T) {
	root, cfg := checkRepo(t)
	commitPins(t, root)
	stage(t, root, map[string]string{pinsPath: pinsOne + pinsTwo})
	if out, ok := check(t, cfg, "Second repin ["+chkID+"@1]"); !ok || strings.Contains(out, "append-only") {
		t.Fatalf("a second repin only appends to pins.jsonl and must pass:\n%s", out)
	}
	gitIn(t, root, "commit", "-q", "-m", "Second repin ["+chkID+"@1]")
	if out := mustRun(t, "-c", cfg, "check", "--commit", "HEAD"); !strings.Contains(out, "✓") {
		t.Fatalf("--commit agrees:\n%s", out)
	}
}

func TestCheckRejectsRewritingAPinLog(t *testing.T) {
	for name, body := range map[string]string{
		"rewritten line":     strings.Replace(pinsOne, `"to":2`, `"to":9`, 1) + pinsTwo,
		"dropped line":       pinsTwo,
		"emptied":            "",
		"edited last line":   strings.TrimSuffix(pinsOne, "\n") + ` ` + "\n",
		"missing final line": strings.TrimSuffix(pinsOne, "\n"),
	} {
		t.Run(name, func(t *testing.T) {
			root, cfg := checkRepo(t)
			commitPins(t, root)
			stage(t, root, map[string]string{pinsPath: body})
			if out, ok := check(t, cfg, "Rewrite ["+chkID+"@1]"); ok || !strings.Contains(out, "snapshot files are append-only") {
				t.Fatalf("%s must be refused:\n%s", name, out)
			}
		})
	}
}

func TestCheckStillRejectsEditingAVersionSnapshotOrANestedPinLog(t *testing.T) {
	root, cfg := checkRepo(t)
	snaps, err := filepath.Glob(filepath.Join(root, "provenance", ".versions", chkID, "1-*.yml"))
	if err != nil || len(snaps) != 1 {
		t.Fatalf("version 1 snapshot: %v (%v)", snaps, err)
	}
	snap, err := filepath.Rel(root, snaps[0])
	if err != nil {
		t.Fatal(err)
	}
	nested := "provenance/.versions/" + chkID + "/sub/pins.jsonl"
	stage(t, root, map[string]string{nested: pinsOne})
	gitIn(t, root, "commit", "-q", "-m", "Seed ["+chkID+"@1]")

	if err := os.Chmod(snaps[0], 0o644); err != nil {
		t.Fatal(err)
	}
	stage(t, root, map[string]string{snap: readFile(t, snaps[0]) + "# edited\n"})
	if out, ok := check(t, cfg, "Edit snapshot ["+chkID+"@1]"); ok || !strings.Contains(out, "snapshot files are append-only") {
		t.Fatalf("editing a version snapshot must be refused:\n%s", out)
	}
	gitIn(t, root, "reset", "-q", "--hard")
	stage(t, root, map[string]string{nested: pinsOne + pinsTwo})
	if out, ok := check(t, cfg, "Edit nested ["+chkID+"@1]"); ok || !strings.Contains(out, "snapshot files are append-only") {
		t.Fatalf("only <id>/pins.jsonl is an append-only log:\n%s", out)
	}
}
