package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCloneReproducesTheManifest(t *testing.T) {
	remoteEnv(t)
	root, cfg := checkRepo(t)
	signedIn(t)
	mustRun(t, "-c", cfg, "edit", chkID, "--add-constraint", "Second", "--reason", "Tighten it")
	if _, _, err := runCLI(t, "", "-c", cfg, "publish", "--all", "--yes"); err != nil {
		t.Fatal(err)
	}
	dir := filepath.Join(t.TempDir(), "copy")
	out, _, err := runCLI(t, "", "clone", dir)
	if err != nil || !strings.Contains(out, "Cloned 1 record") {
		t.Fatalf("clone: %v\n%s", err, out)
	}
	if got, want := manifestLines(t, dir), manifestLines(t, root); !strings.Contains(want, strings.TrimSpace(got)) {
		t.Fatalf("manifest differs:\nclone: %s\nsource: %s", got, want)
	}
	for _, p := range []string{"provenance/" + chkID + ".yml", ".provenance.yml", ".gitattributes"} {
		if _, err := os.Stat(filepath.Join(dir, p)); err != nil {
			t.Errorf("clone lacks %s", p)
		}
	}
	vs, _ := filepath.Glob(filepath.Join(dir, "provenance", ".versions", chkID, "*.yml"))
	if len(vs) != 2 {
		t.Fatalf("snapshots = %v", vs)
	}
	if _, _, err := runCLI(t, "", "clone", dir); err == nil {
		t.Fatal("cloning into a directory that is not empty must fail")
	}
}

func TestCloneNeedsSession(t *testing.T) {
	remoteEnv(t)
	_, _, err := runCLI(t, "", "clone", filepath.Join(t.TempDir(), "x"))
	if err == nil || exitCode(err) != 3 {
		t.Fatalf("err=%v code=%d", err, exitCode(err))
	}
}
