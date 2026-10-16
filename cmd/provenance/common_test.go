package main

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// run executes the provenance command tree with args and returns its stdout.
func run(t *testing.T, args ...string) (string, error) {
	t.Helper()
	root := newRootCmd()
	var out bytes.Buffer
	root.SetOut(&out)
	root.SetErr(&out)
	root.SetArgs(args)
	err := root.Execute()
	return out.String(), err
}

func mustRun(t *testing.T, args ...string) string {
	t.Helper()
	out, err := run(t, args...)
	if err != nil {
		t.Fatalf("provenance %s: %v\n%s", strings.Join(args, " "), err, out)
	}
	return out
}

func writeFile(t *testing.T, path, body string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
}

func readFile(t *testing.T, path string) string {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return string(data)
}

func gitIn(t *testing.T, dir string, args ...string) string {
	t.Helper()
	c := exec.Command("git", args...)
	c.Dir = dir
	c.Env = append(os.Environ(), "GIT_CONFIG_GLOBAL=/dev/null", "GIT_CONFIG_SYSTEM=/dev/null")
	out, err := c.CombinedOutput()
	if err != nil {
		t.Fatalf("git %s: %v\n%s", strings.Join(args, " "), err, out)
	}
	return strings.TrimSpace(string(out))
}

// gitRepo makes dir a git repository with an identity and no hooks, and
// commits everything in it.
func gitRepo(t *testing.T, dir string) {
	t.Helper()
	gitIn(t, dir, "init", "-q")
	gitIn(t, dir, "config", "user.email", "test@example.com")
	gitIn(t, dir, "config", "user.name", "Test")
	gitIn(t, dir, "config", "core.hooksPath", "/dev/null")
	gitIn(t, dir, "add", "-A")
	gitIn(t, dir, "commit", "-q", "-m", "fixture")
}

// v1 is a v1 record file, as linespec writes it.
func v1(id, typ, status, extra string) string {
	return "id: " + id + "\ntitle: T " + id + "\nstatus: " + status + "\ncreated_at: \"2026-01-01\"\nauthor: a@b\n" +
		"intent: Intent of " + id + "\nconstraints:\n  - Shared\naffected_scope:\n  - a.go\nforbidden_scope: []\ntype: " + typ + "\n" +
		"supersedes: \"\"\nsuperseded_by: \"\"\nrelated: []\nsealed_at_sha: \"\"\nassociated_specs: []\nassociated_traces: []\nmonitors: []\ntags: []\n" + extra
}
