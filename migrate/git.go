package migrate

import (
	"bufio"
	"bytes"
	"os/exec"
	"path"
	"strconv"
	"strings"
)

// GitTime is the commit that first added a record file.
type GitTime struct {
	Unix int64
	SHA  string
}

// Git reads record first-add times. Version order depends on them, so tests
// can substitute a fixed history.
type Git interface {
	// FirstAdded maps each file basename under dir to the oldest commit that
	// added it (log --diff-filter=A --reverse).
	FirstAdded(root, dir string) (map[string]GitTime, error)
	// FollowFirst returns the oldest commit touching relpath across renames
	// and merges: the fallback when FirstAdded does not know a file.
	FollowFirst(root, relpath string) (GitTime, bool)
}

// GitCLI implements Git with the git command. A repository without git
// history yields no times, never an error, so every record falls back to
// created_at.
type GitCLI struct{}

func (GitCLI) FirstAdded(root, dir string) (map[string]GitTime, error) {
	out, err := exec.Command("git", "-C", root, "log", "--diff-filter=A", "--reverse",
		"--format=C %ct %H", "--name-only", "--", dir).Output()
	first := map[string]GitTime{}
	if err != nil {
		return first, nil
	}
	var cur *GitTime
	sc := bufio.NewScanner(bytes.NewReader(out))
	sc.Buffer(make([]byte, 1<<20), 1<<20)
	for sc.Scan() {
		line := sc.Text()
		if strings.HasPrefix(line, "C ") {
			f := strings.Fields(line)
			if len(f) == 3 {
				ts, _ := strconv.ParseInt(f[1], 10, 64)
				cur = &GitTime{Unix: ts, SHA: f[2]}
			}
			continue
		}
		if name := strings.TrimSpace(line); name != "" && cur != nil {
			base := path.Base(name)
			if _, seen := first[base]; !seen {
				first[base] = *cur
			}
		}
	}
	return first, nil
}

func (GitCLI) FollowFirst(root, relpath string) (GitTime, bool) {
	out, err := exec.Command("git", "-C", root, "log", "--follow", "--format=%ct %H", "--", relpath).Output()
	if err != nil {
		return GitTime{}, false
	}
	lines := strings.Fields(strings.TrimSpace(string(out)))
	if len(lines) < 2 {
		return GitTime{}, false
	}
	ts, err := strconv.ParseInt(lines[len(lines)-2], 10, 64)
	if err != nil {
		return GitTime{}, false
	}
	return GitTime{Unix: ts, SHA: lines[len(lines)-1]}, true
}
