package main

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestEnsureMergeDriverWritesAttributeAndConfig(t *testing.T) {
	root := t.TempDir()
	gitMust(t, root, "init", "-q")
	writeFile(t, attrFile(root), "*.png binary")
	status, err := ensureMergeDriver(root)
	if err != nil || status != "updated" {
		t.Fatalf("first run: %q, %v", status, err)
	}
	if got, want := readFile(t, attrFile(root)), "*.png binary\n"+manifestPath+" merge=provenance\n"; got != want {
		t.Fatalf(".gitattributes:\n%q\nwant\n%q", got, want)
	}
	if got := strings.TrimSpace(gitMust(t, root, "config", "merge.provenance.driver")); got != "provenance merge-manifest %O %A %B" {
		t.Fatalf("merge.provenance.driver = %q", got)
	}
	if status, err := ensureMergeDriver(root); err != nil || status != "unchanged" {
		t.Fatalf("second run: %q, %v", status, err)
	}
}

func TestEnsureMergeDriverOutsideGitOnlyWritesTheAttribute(t *testing.T) {
	root := t.TempDir()
	status, err := ensureMergeDriver(root)
	if err != nil || status != "created" {
		t.Fatalf("run: %q, %v", status, err)
	}
	if !strings.Contains(readFile(t, filepath.Join(root, ".gitattributes")), manifestPath+" merge=provenance\n") {
		t.Fatal("the attribute line is missing")
	}
}

func TestHooksInstallSetsUpTheMergeDriver(t *testing.T) {
	root, cfg, _ := hooksRepo(t)
	out := mustRun(t, "-c", cfg, "hooks", "install")
	if !strings.Contains(out, "merge driver (") {
		t.Errorf("install output misses the merge driver line:\n%s", out)
	}
	if !strings.Contains(readFile(t, attrFile(root)), manifestPath+" merge=provenance\n") {
		t.Error("hooks install did not write the merge attribute")
	}
	if got := strings.TrimSpace(gitMust(t, root, "config", "merge.provenance.driver")); got != "provenance merge-manifest %O %A %B" {
		t.Errorf("merge.provenance.driver = %q", got)
	}
}
