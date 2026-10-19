package main

import (
	"strings"
	"testing"
)

// The merge attribute only helps a team when it is committed: the first
// install says so, and a repeat install, which changes nothing, does not.
func TestHooksInstallTellsToCommitGitattributes(t *testing.T) {
	_, cfg, _ := hooksRepo(t)
	first := mustRun(t, "-c", cfg, "hooks", "install")
	if !strings.Contains(first, "Commit .gitattributes") || !strings.Contains(first, "git add .gitattributes") {
		t.Errorf("first install does not tell the user to commit .gitattributes:\n%s", first)
	}
	if again := mustRun(t, "-c", cfg, "hooks", "install"); strings.Contains(again, "Commit .gitattributes") {
		t.Errorf("a repeat install repeats the commit hint:\n%s", again)
	}
}
