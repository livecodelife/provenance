package main

import (
	"path/filepath"
	"strings"
	"testing"
)

// Proof for the CLI side of shared pins: repin against the sync cache, and
// migrate pinning shared parents from it (failing clearly when unsynced).

func TestRepinSharedChild(t *testing.T) {
	remote := t.TempDir()
	writeFile(t, filepath.Join(remote, "provenance", sharedParent+".yml"),
		strings.Replace(v2Draft(sharedParent, "open"), "version: 1\nreason: \"\"", "version: 2\nreason: moved", 1))
	gitRepo(t, remote)
	_, cfg := sharedConsumer(t, remote)

	out, err := run(t, "-c", cfg, "repin", sharedChild, "--reason", "checked")
	if err == nil || !strings.Contains(out+err.Error(), "provenance sync") {
		t.Fatalf("before sync repin must say to sync (err %v):\n%s", err, out)
	}
	mustRun(t, "-c", cfg, "sync")
	out = mustRun(t, "-c", cfg, "repin", sharedChild, "--reason", "checked")
	if !strings.Contains(out, "now pinned to "+sharedParent+"@2 (was 1)") {
		t.Fatalf("repin:\n%s", out)
	}
	if out := mustRun(t, "-c", cfg, "stale"); strings.Contains(out, sharedChild) {
		t.Fatalf("still stale:\n%s", out)
	}
}

func TestMigratePinsSharedParentFromCache(t *testing.T) {
	remote := t.TempDir()
	writeFile(t, filepath.Join(remote, "provenance", sharedParent+".yml"), v1(sharedParent, "blueprint", "implemented", ""))
	gitRepo(t, remote)
	root := t.TempDir()
	writeFile(t, filepath.Join(root, ".linespec.yml"), "provenance:\n  dir: provenance\n  shared_repos:\n    - name: shared\n      url: "+remote+"\n      ref: HEAD\n")
	writeFile(t, filepath.Join(root, "provenance", sharedChild+".yml"), v1(sharedChild, "imprint", "implemented", "implements: "+sharedParent+"\n"))
	gitRepo(t, root)
	cfg := filepath.Join(root, ".linespec.yml")

	out, err := run(t, "-c", cfg, "migrate")
	if err == nil || !strings.Contains(out+err.Error(), "provenance sync") {
		t.Fatalf("unsynced migrate must say to sync (err %v):\n%s", err, out)
	}
	mustRun(t, "-c", cfg, "sync")
	out = mustRun(t, "-c", cfg, "migrate")
	if !strings.Contains(out, "Shared parents pinned") || !strings.Contains(out, sharedChild+"  shared:"+sharedParent+"@1") {
		t.Fatalf("migrate output:\n%s", out)
	}
	if out := mustRun(t, "-c", filepath.Join(root, ".provenance.yml"), "stale"); strings.Contains(out, "cannot be judged") {
		t.Fatalf("unjudged after migrate:\n%s", out)
	}
}
