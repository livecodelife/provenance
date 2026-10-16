package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// TestHooksPrePush pushes through the installed pre-push hook: a commit that
// lacks the file pre_push_command checks for blocks the push, and the hook
// is a no-op when the repository sets no command.
func TestHooksPrePush(t *testing.T) {
	bin := filepath.Join(t.TempDir(), "provenance")
	build := exec.Command("go", "build", "-o", bin, ".")
	if out, err := build.CombinedOutput(); err != nil {
		t.Fatalf("build: %v\n%s", err, out)
	}
	t.Setenv("PROVENANCE", bin)
	root, cfg, _ := hooksRepo(t)
	mustRun(t, "-c", cfg, "hooks", "install")
	if _, err := os.Stat(filepath.Join(root, ".git", "hooks", "pre-push")); err != nil {
		t.Fatalf("pre-push not installed: %v", err)
	}
	remote := t.TempDir()
	gitPlain(t, remote, "init", "-q", "--bare")
	gitIn(t, root, "remote", "add", "origin", remote)
	branch := gitIn(t, root, "rev-parse", "--abbrev-ref", "HEAD")

	// No pre_push_command: the push goes through.
	gitPlain(t, root, "push", "-q", "origin", branch)

	writeFile(t, cfg, readFile(t, cfg)+"pre_push_command: test -f ok.txt\n")
	gitIn(t, root, "add", "-A")
	gitIn(t, root, "commit", "-q", "--no-verify", "-m", "config")
	writeFile(t, filepath.Join(root, "ok.txt"), "ok\n")
	gitIn(t, root, "add", "-A")
	gitIn(t, root, "commit", "-q", "--no-verify", "-m", "adds ok.txt")
	push := exec.Command("git", "-C", root, "push", "-q", "origin", branch)
	push.Env = gitPlainEnv()
	out, err := push.CombinedOutput()
	if err == nil || !strings.Contains(string(out), "does not pass pre_push_command") {
		t.Fatalf("the config commit has no ok.txt, so the push must fail (err %v):\n%s", err, out)
	}
	if strings.Count(string(out), "does not pass") != 1 {
		t.Fatalf("only the commit without ok.txt may fail:\n%s", out)
	}
}

// gitPlainEnv is the environment without the repository-selecting GIT_*
// variables (GIT_CONFIG_GLOBAL and the like stay), for commands that must
// find their own repository: a bare remote, or a push that runs hooks.
func gitPlainEnv() []string {
	var env []string
	for _, kv := range os.Environ() {
		switch {
		case strings.HasPrefix(kv, "GIT_DIR="), strings.HasPrefix(kv, "GIT_WORK_TREE="), strings.HasPrefix(kv, "GIT_INDEX_FILE="):
		default:
			env = append(env, kv)
		}
	}
	return env
}

func gitPlain(t *testing.T, dir string, args ...string) string {
	t.Helper()
	real, err := filepath.EvalSymlinks(dir)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(real+string(filepath.Separator), testTemp+string(filepath.Separator)) {
		t.Fatalf("refusing git %s outside the test temp root: %s", strings.Join(args, " "), dir)
	}
	c := exec.Command("git", append([]string{"-C", real}, args...)...)
	c.Dir = real
	c.Env = gitPlainEnv()
	return runGit(t, c)
}
