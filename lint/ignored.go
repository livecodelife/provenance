package lint

import (
	"bytes"
	"os/exec"
	"strings"
)

// gitIgnored returns the subset of paths (relative to root) that git ignores.
// It reports nothing when root is not a git work tree or git is missing, so
// lint still runs outside a repository. Tracked files are never reported:
// check-ignore only matches untracked paths.
func gitIgnored(root string, paths []string) map[string]bool {
	if len(paths) == 0 {
		return nil
	}
	cmd := exec.Command("git", "-C", root, "check-ignore", "-z", "--stdin")
	cmd.Stdin = strings.NewReader(strings.Join(paths, "\x00") + "\x00")
	var out bytes.Buffer
	cmd.Stdout = &out
	// Exit 1 means nothing matched; 128 is not a repository. Either way the
	// output is the answer.
	_ = cmd.Run()
	found := map[string]bool{}
	for _, p := range strings.Split(out.String(), "\x00") {
		if p != "" {
			found[p] = true
		}
	}
	return found
}
