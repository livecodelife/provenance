package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// A regeneration commit refused because sedum cannot be found must say so
// rather than only advising a revise.

const missingSedumMessage = "regeneration check could not run sedum: not found on PATH; set SEDUM_BIN or install sedum"

func TestRegenerationCommitSaysWhenSedumIsMissingFromSEDUM_BIN(t *testing.T) {
	root, cfg, _, _ := regenRepo(t)
	t.Setenv("SEDUM_BIN", filepath.Join(t.TempDir(), "no-such-sedum"))
	p := baseParts()
	p.pkg = "sha256:p2"
	stage(t, root, map[string]string{"a.go": p.render()})
	out, ok := check(t, cfg, regenMsg())
	if ok || !strings.Contains(out, chkID+" "+missingSedumMessage) {
		t.Fatalf("a missing sedum must be named in the refusal:\n%s", out)
	}
	if !strings.Contains(out, "provenance revise "+chkID) {
		t.Fatalf("the revise advice must stay for a commit that is not a regeneration:\n%s", out)
	}
}

func TestRegenerationCommitSaysWhenSedumIsMissingFromPATH(t *testing.T) {
	root, cfg, _, _ := regenRepo(t)
	t.Setenv("SEDUM_BIN", "")
	git, err := exec.LookPath("git")
	if err != nil {
		t.Skip("git not found")
	}
	bin := t.TempDir()
	if err := os.Symlink(git, filepath.Join(bin, "git")); err != nil {
		t.Fatal(err)
	}
	t.Setenv("PATH", bin)
	p := baseParts()
	p.pkg = "sha256:p2"
	stage(t, root, map[string]string{"a.go": p.render()})
	if out, ok := check(t, cfg, regenMsg()); ok || !strings.Contains(out, missingSedumMessage) {
		t.Fatalf("a sedum absent from PATH must be named in the refusal:\n%s", out)
	}
}

func TestRegenerationCommitDoesNotBlameSedumWhenItSaysStale(t *testing.T) {
	root, cfg, state, _ := regenRepo(t)
	writeFile(t, state, "stale\n")
	p := baseParts()
	p.pkg = "sha256:p2"
	stage(t, root, map[string]string{"a.go": p.render()})
	out, ok := check(t, cfg, regenMsg())
	if ok || strings.Contains(out, "could not run sedum") || !strings.Contains(out, "provenance revise") {
		t.Fatalf("a stale answer keeps the plain revise advice:\n%s", out)
	}
}
