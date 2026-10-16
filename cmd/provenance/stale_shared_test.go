package main

import (
	"path/filepath"
	"strings"
	"testing"
)

// Proof for the stale command judging parents in shared repos from the
// synced cache.

func TestStaleJudgesSharedParent(t *testing.T) {
	remote := t.TempDir()
	writeFile(t, filepath.Join(remote, "provenance", sharedParent+".yml"),
		strings.Replace(v2Draft(sharedParent, "open"), "version: 1\nreason: \"\"", "version: 2\nreason: moved", 1))
	gitRepo(t, remote)
	_, cfg := sharedConsumer(t, remote)

	out := mustRun(t, "-c", cfg, "stale")
	if !strings.Contains(out, "cannot be judged: parent is in shared repo shared") {
		t.Fatalf("before sync the shared parent cannot be judged:\n%s", out)
	}

	mustRun(t, "-c", cfg, "sync")
	out = mustRun(t, "-c", cfg, "stale")
	want := sharedChild + "@1  pinned to " + sharedParent + "@1, parent is at 2  (direct)"
	if !strings.Contains(out, want) || strings.Contains(out, "cannot be judged") {
		t.Fatalf("after sync the shared parent is judged:\n%s", out)
	}
	if _, err := run(t, "-c", cfg, "stale", "--exit-code"); err == nil {
		t.Fatal("--exit-code must fail for a child behind a shared parent")
	}
}
