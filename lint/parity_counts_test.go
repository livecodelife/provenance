package lint

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"gopkg.in/yaml.v3"

	"github.com/livecodelife/provenance/record"
)

// TestParityCounts compares the totals, not just the issues: for each repo,
// the record count (local plus shared) and the error, warning and hint counts
// must equal linespec's, once the rules this package retires or adds are
// taken out. A repo linted without its shared records would show up here as
// a record-count mismatch.
func TestParityCounts(t *testing.T) {
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
			gc := exec.Command("git", "clone", "--quiet", "--local", filepath.Join(ws, name), clone)
			gc.Env = withoutGitEnv(os.Environ())
			if out, err := gc.CombinedOutput(); err != nil {
				t.Fatalf("clone: %v\n%s", err, out)
			}
			if migratedToV2(clone) {
				t.Skipf("%s is migrated to v2: linespec no longer reads its records, so there is nothing to compare (parity was proven before it migrated, prov-2026-e135b1f5)", name)
			}
			cmd := exec.Command("linespec", "provenance", "lint", "--format", "json", "--all")
			cmd.Dir = clone
			cmd.Env = withoutGitEnv(os.Environ())
			out, _ := cmd.Output()
			var ls struct {
				Enforcement string `json:"enforcement"`
				Passed      int    `json:"passed"`
				Errors      int    `json:"errors"`
				Warnings    int    `json:"warnings"`
				Issues      []struct {
					Field    string `json:"field"`
					Severity string `json:"severity"`
				} `json:"issues"`
			}
			// linespec prints its report twice; the first is the real one.
			if err := json.NewDecoder(strings.NewReader(string(out))).Decode(&ls); err != nil {
				t.Fatalf("linespec output: %v", err)
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

			// linespec's own counts, minus the issues of retired rules.
			var wantErr, wantWarn, wantHint int
			for _, i := range ls.Issues {
				if retiredFields[i.Field] {
					continue
				}
				switch i.Severity {
				case "error":
					wantErr++
				case "warning":
					wantWarn++
				case "hint":
					wantHint++
				}
			}
			var gotErr, gotWarn, gotHint int
			for _, i := range res.Issues {
				if newRules[i.Rule] {
					continue
				}
				switch i.Severity {
				case SeverityError:
					gotErr++
				case SeverityWarning:
					gotWarn++
				case SeverityHint:
					gotHint++
				}
			}
			// linespec: passed = records - errors, so records = passed + errors.
			if want := ls.Passed + ls.Errors; res.Records != want {
				t.Errorf("records %d, linespec %d (%d local + shared)", res.Records, want, len(opts.Shared))
			}
			if gotErr != wantErr || gotWarn != wantWarn || gotHint != wantHint {
				t.Errorf("errors/warnings/hints %d/%d/%d, linespec %d/%d/%d", gotErr, gotWarn, gotHint, wantErr, wantWarn, wantHint)
			}
			t.Logf("%d records (%d shared); errors %d, warnings %d, hints %d", res.Records, len(opts.Shared), gotErr, gotWarn, gotHint)
		})
	}
}
