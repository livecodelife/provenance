package migrate

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// isolatedGit runs `git -C dir <args>` with no inherited GIT_* variables and
// with HOME and global/system config pointed into the test's temp dir. Before
// any command but init, it checks that dir is the repository's toplevel, so
// test setup can never write to a real repository (COMMON.md git rule).
func isolatedGit(t *testing.T, home, dir string, args ...string) {
	t.Helper()
	env := append(withoutGitEnv(os.Environ()),
		"HOME="+home, "GIT_CONFIG_NOSYSTEM=1", "GIT_CONFIG_GLOBAL="+filepath.Join(home, "gitconfig"),
		"GIT_AUTHOR_NAME=t", "GIT_AUTHOR_EMAIL=t@t", "GIT_COMMITTER_NAME=t", "GIT_COMMITTER_EMAIL=t@t",
		"GIT_AUTHOR_DATE=2026-01-02T00:00:00Z", "GIT_COMMITTER_DATE=2026-01-02T00:00:00Z")
	run := func(args ...string) string {
		cmd := exec.Command("git", append([]string{"-C", dir}, args...)...)
		cmd.Env = env
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("git -C %s %v: %v\n%s", dir, args, err, out)
		}
		return strings.TrimSpace(string(out))
	}
	if args[0] != "init" {
		top, err := filepath.EvalSymlinks(run("rev-parse", "--show-toplevel"))
		want, err2 := filepath.EvalSymlinks(dir)
		if err != nil || err2 != nil || top != want {
			t.Fatalf("refusing git %v: toplevel is %q, not the test repo %q", args, top, want)
		}
	}
	run(args...)
}

func newRepo(t *testing.T, home, dir, file string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Join(dir, "provenance"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "provenance", file), []byte("id: x\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	isolatedGit(t, home, dir, "init", "-q")
	isolatedGit(t, home, dir, "add", "-A")
	isolatedGit(t, home, dir, "commit", "-q", "-m", "add "+file)
}

// TestGitCLIIgnoresHookEnv simulates a git hook: GIT_DIR and GIT_INDEX_FILE
// name a decoy repository. GitCLI must still read the repository it is given
// and leave the decoy exactly as it was. Every repository is in t.TempDir().
func TestGitCLIIgnoresHookEnv(t *testing.T) {
	tmp := t.TempDir()
	home := filepath.Join(tmp, "home")
	target, decoy := filepath.Join(tmp, "target"), filepath.Join(tmp, "decoy")
	os.MkdirAll(home, 0o755)
	newRepo(t, home, target, "prov-2026-0000a001.yml")
	newRepo(t, home, decoy, "prov-2026-0000d001.yml")
	before, _ := os.ReadFile(filepath.Join(decoy, ".git", "config"))
	head, _ := os.ReadFile(filepath.Join(decoy, ".git", "HEAD"))

	t.Setenv("GIT_DIR", filepath.Join(decoy, ".git"))
	t.Setenv("GIT_INDEX_FILE", filepath.Join(decoy, ".git", "index"))
	t.Setenv("GIT_WORK_TREE", decoy)

	added, err := GitCLI{}.FirstAdded(target, "provenance")
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := added["prov-2026-0000a001.yml"]; !ok || len(added) != 1 {
		t.Fatalf("FirstAdded read the wrong repository: %v", added)
	}
	if _, ok := (GitCLI{}).FollowFirst(target, "provenance/prov-2026-0000a001.yml"); !ok {
		t.Fatal("FollowFirst did not find the target's file")
	}
	after, _ := os.ReadFile(filepath.Join(decoy, ".git", "config"))
	head2, _ := os.ReadFile(filepath.Join(decoy, ".git", "HEAD"))
	if string(before) != string(after) || string(head) != string(head2) {
		t.Fatal("decoy repository changed")
	}
}

func TestWithoutGitEnv(t *testing.T) {
	got := withoutGitEnv([]string{"GIT_DIR=/x", "HOME=/h", "GIT_INDEX_FILE=/i", "PATH=/p", "XGIT_KEEP=1"})
	if len(got) != 3 || got[0] != "HOME=/h" || got[1] != "PATH=/p" || got[2] != "XGIT_KEEP=1" {
		t.Fatalf("got %v", got)
	}
}
