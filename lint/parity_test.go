package lint

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"gopkg.in/yaml.v3"

	"github.com/livecodelife/provenance/record"
)

// retiredFields are the fields of linespec issues from rules this package
// retires: supersession (PROV006/007/014/020), the bug/extends checks, and
// hash-manifest immutability (PROV017).
// "cache" is linespec's shared-repo cache freshness warning: CLI state with no
// record, which this library does not manage.
var retiredFields = map[string]bool{"supersedes": true, "superseded_by": true, "extends": true, "integrity": true, "cache": true}

// newRules have no linespec counterpart.
var newRules = map[string]bool{RuleScopeOverlap: true, RuleUnbumpedChange: true, RuleBumpWithoutReason: true,
	RuleVersionSnapshots: true, RuleImplementsPin: true, RuleNotMigrated: true}

type lsConfig struct {
	Provenance struct {
		Dir          string   `yaml:"dir"`
		ExcludePaths []string `yaml:"exclude_paths"`
		SharedRepos  []struct {
			Name string `yaml:"name"`
			URL  string `yaml:"url"`
		} `yaml:"shared_repos"`
	} `yaml:"provenance"`
}

// TestParityWithLinespec runs `linespec provenance lint --format json --all`
// and this package on the same git clone of each repository (never on the
// original) and requires identical issues for the carried-over rules.
func TestParityWithLinespec(t *testing.T) {
	ws := os.Getenv("PROVENANCE_REAL_REPOS")
	if ws == "" {
		t.Skip("set PROVENANCE_REAL_REPOS to run")
	}
	if _, err := exec.LookPath("linespec"); err != nil {
		t.Skip("linespec not installed")
	}
	for _, name := range []string{"linespec", "sedum", "scion", "linespec-provenance-ui"} {
		t.Run(name, func(t *testing.T) {
			clone := filepath.Join(t.TempDir(), name)
			// GIT_* stripped: under a hook, git clone would honour GIT_DIR.
			gc := exec.Command("git", "clone", "--quiet", "--local", filepath.Join(ws, name), clone)
			gc.Env = withoutGitEnv(os.Environ())
			if out, err := gc.CombinedOutput(); err != nil {
				t.Fatalf("clone: %v\n%s", err, out)
			}
			cmd := exec.Command("linespec", "provenance", "lint", "--format", "json", "--all")
			cmd.Dir = clone
			cmd.Env = withoutGitEnv(os.Environ())
			out, _ := cmd.Output() // lint exits non-zero when it finds errors
			var ls struct {
				Enforcement string `json:"enforcement"`
				Issues      []struct {
					RecordID string `json:"record_id"`
					Field    string `json:"field"`
					Message  string `json:"message"`
					Severity string `json:"severity"`
				} `json:"issues"`
			}
			if err := json.NewDecoder(strings.NewReader(string(out))).Decode(&ls); err != nil {
				t.Fatalf("linespec output: %v\n%.500s", err, out)
			}

			var cfg lsConfig
			data, _ := os.ReadFile(filepath.Join(clone, ".linespec.yml"))
			if err := yaml.Unmarshal(data, &cfg); err != nil {
				t.Fatal(err)
			}
			dir := strings.TrimRight(cfg.Provenance.Dir, "/")
			if dir == "" {
				dir = "provenance"
			}
			opts := Options{Root: clone, Enforcement: ls.Enforcement, ExcludePaths: cfg.Provenance.ExcludePaths, Shared: map[string]*record.Record{}}
			home, _ := os.UserHomeDir()
			for _, sr := range cfg.Provenance.SharedRepos {
				opts.SharedRepos = append(opts.SharedRepos, sr.Name)
				cache := filepath.Join(home, ".linespec", "cache", fmt.Sprintf("%x", sha256.Sum256([]byte(sr.URL))), "provenance")
				for id, r := range loadTree(t, cache) {
					opts.Shared[id] = r
				}
			}
			res, err := Lint(&record.FilesStore{Root: clone, Dir: dir}, opts)
			if err != nil {
				t.Fatal(err)
			}

			want := map[string]int{}
			for _, i := range ls.Issues {
				if !retiredFields[i.Field] {
					want[key(i.RecordID, i.Field, i.Severity, i.Message)]++
				}
			}
			got := map[string]int{}
			for _, i := range res.Issues {
				if !newRules[i.Rule] {
					got[key(i.RecordID, i.Field, string(i.Severity), i.Message)]++
				}
			}
			diff(t, got, want)
			t.Logf("%d carried issues match (%d linespec issues, %d retired)", sum(want), len(ls.Issues), len(ls.Issues)-sum(want))
		})
	}
}

func key(id, field, sev, msg string) string { return id + " | " + field + " | " + sev + " | " + msg }

func sum(m map[string]int) int {
	n := 0
	for _, v := range m {
		n += v
	}
	return n
}

func diff(t *testing.T, got, want map[string]int) {
	t.Helper()
	var extra, missing []string
	for k, n := range got {
		if want[k] < n {
			extra = append(extra, fmt.Sprintf("%dx %s", n-want[k], k))
		}
	}
	for k, n := range want {
		if got[k] < n {
			missing = append(missing, fmt.Sprintf("%dx %s", n-got[k], k))
		}
	}
	sort.Strings(extra)
	sort.Strings(missing)
	for i, s := range extra {
		if i == 15 {
			t.Errorf("... %d more extra", len(extra)-15)
			break
		}
		t.Errorf("extra (ours only): %s", s)
	}
	for i, s := range missing {
		if i == 15 {
			t.Errorf("... %d more missing", len(missing)-15)
			break
		}
		t.Errorf("missing (linespec only): %s", s)
	}
}

// loadTree reads prov-*.yml records under dir recursively, as linespec's
// loader does for shared repo caches.
func loadTree(t *testing.T, dir string) map[string]*record.Record {
	out := map[string]*record.Record{}
	_ = filepath.WalkDir(dir, func(p string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() || !strings.HasPrefix(d.Name(), "prov-") ||
			!(strings.HasSuffix(d.Name(), ".yml") || strings.HasSuffix(d.Name(), ".yaml")) {
			return nil
		}
		data, err := os.ReadFile(p)
		if err != nil {
			t.Fatal(err)
		}
		r, err := record.Decode(data)
		if err != nil {
			t.Fatalf("%s: %v", p, err)
		}
		out[r.ID] = r
		return nil
	})
	return out
}

func withoutGitEnv(env []string) []string {
	out := make([]string, 0, len(env))
	for _, kv := range env {
		if !strings.HasPrefix(kv, "GIT_") {
			out = append(out, kv)
		}
	}
	return out
}
