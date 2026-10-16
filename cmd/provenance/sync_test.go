package main

import (
	"path/filepath"
	"strings"
	"testing"
)

const (
	sharedParent = "prov-2026-0000f001"
	sharedChild  = "prov-2026-0000f002"
)

// sharedRemote is a git repo whose provenance dir holds one v2 blueprint.
func sharedRemote(t *testing.T) string {
	t.Helper()
	remote := t.TempDir()
	writeFile(t, filepath.Join(remote, "provenance", sharedParent+".yml"), v2Draft(sharedParent, "open"))
	gitRepo(t, remote)
	return remote
}

// sharedConsumer is a repo with one record implementing the remote's blueprint.
func sharedConsumer(t *testing.T, url string) (root, cfg string) {
	t.Helper()
	child := strings.Replace(v2Draft(sharedChild, "draft"), "type: blueprint", "type: imprint\nimplements: "+sharedParent+"\nimplements_repo: shared\nimplements_version: 1", 1)
	root, cfg = v2Repo(t, false, map[string]string{sharedChild: child})
	writeFile(t, cfg, readFile(t, cfg)+"shared_repos:\n  - name: shared\n    url: "+url+"\n    ref: HEAD\n")
	return root, cfg
}

func TestSyncAndSharedLint(t *testing.T) {
	remote := sharedRemote(t)
	_, cfg := sharedConsumer(t, remote)

	out, err := run(t, "-c", cfg, "lint")
	if err == nil || !strings.Contains(out, sharedParent) || !strings.Contains(out, "no cache") {
		t.Fatalf("before sync lint must fail on the unresolved parent and say why (err %v):\n%s", err, out)
	}
	if out := mustRun(t, "-c", cfg, "sync"); !strings.Contains(out, "Synced shared") || !strings.Contains(out, "1 records") {
		t.Fatalf("sync:\n%s", out)
	}
	if out := mustRun(t, "-c", cfg, "lint"); strings.Contains(out, "unknown record") || strings.Contains(out, "stale") {
		t.Fatalf("after sync the shared parent resolves:\n%s", out)
	}
	if out := mustRun(t, "-c", cfg, "sync"); !strings.Contains(out, "cache is fresh, skipped") {
		t.Fatalf("second sync:\n%s", out)
	}
	if out := mustRun(t, "-c", cfg, "sync", "--force"); !strings.Contains(out, "Synced shared") {
		t.Fatalf("forced sync:\n%s", out)
	}
}

func TestSyncWithoutSharedRepos(t *testing.T) {
	_, cfg := v2Repo(t, false, nil)
	if out := mustRun(t, "-c", cfg, "sync"); !strings.Contains(out, "No shared_repos") {
		t.Fatalf("output:\n%s", out)
	}
}
