package main

import (
	"path/filepath"
	"strings"
	"testing"
)

// migratedCommit migrates the fixture and stages the result.
func migratedCommit(t *testing.T) (root, cfg string) {
	t.Helper()
	root = migrateFixture(t)
	mustRun(t, "-c", filepath.Join(root, ".linespec.yml"), "migrate", "--keep-separate", bpOld)
	cfg = filepath.Join(root, ".provenance.yml")
	gitIn(t, root, "add", "-A")
	return root, cfg
}

func TestCheckPassesTheMigrationCommit(t *testing.T) {
	root, cfg := migratedCommit(t)
	msg := "Migrate to provenance v2 [prov-2026-0000ffff]"
	if out, ok := check(t, cfg, msg); !ok {
		t.Fatalf("the v1 -> v2 conversion commit must pass --staged:\n%s", out)
	}
	gitIn(t, root, "commit", "-q", "-m", msg)
	if out := mustRun(t, "-c", cfg, "check", "--commit", "HEAD"); !strings.Contains(out, "✓") {
		t.Fatalf("--commit agrees:\n%s", out)
	}
	if out, err := run(t, "-c", cfg, "check", "--range", "HEAD~1..HEAD"); err != nil {
		t.Fatalf("--range agrees:\n%s", out)
	}
}

func TestCheckStillRejectsTransitionsOnRecordsAlreadyV2(t *testing.T) {
	root, cfg := migratedCommit(t)
	msg := "Migrate to provenance v2 [prov-2026-0000ffff]"
	gitIn(t, root, "commit", "-q", "-m", msg)
	// bpNew is a v2 record now: changing its status by hand is a lifecycle change.
	path := filepath.Join(root, "provenance", bpNew+".yml")
	body := readFile(t, path)
	writeFile(t, path, strings.Replace(body, "status: implemented", "status: draft", 1))
	gitIn(t, root, "add", "-A")
	if out, ok := check(t, cfg, "Break it ["+bpNew+"]"); ok || !strings.Contains(out, "not an allowed transition") {
		t.Fatalf("an implemented -> draft change on a v2 record must still fail on the transition:\n%s", out)
	}
}
