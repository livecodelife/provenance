package lint

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/livecodelife/provenance/record"
)

func gitInit(t *testing.T, root string) {
	t.Helper()
	if out, err := exec.Command("git", "-C", root, "init", "-q").CombinedOutput(); err != nil {
		t.Skipf("git unavailable: %v %s", err, out)
	}
}

func ignoredIssues(res *Result) []Issue {
	var out []Issue
	for _, i := range res.Issues {
		if i.Rule == RuleSpecIgnored {
			out = append(out, i)
		}
	}
	return out
}

func specRepo(t *testing.T, specs ...string) (*record.FilesStore, *record.Record) {
	t.Helper()
	root := t.TempDir()
	gitInit(t, root)
	for name, body := range map[string]string{".gitignore": "built/\n", "built/page.html": "x", "tracked.go": "x"} {
		p := filepath.Join(root, name)
		os.MkdirAll(filepath.Dir(p), 0o755)
		if err := os.WriteFile(p, []byte(body), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	r := newRecord("prov-2026-0000c001", record.StatusOpen)
	for _, s := range specs {
		r.AssociatedSpecs = append(r.AssociatedSpecs, record.AssociatedSpec{Path: s, Type: "unit"})
	}
	s := &record.FilesStore{Root: root}
	if err := s.Save(r, nil); err != nil {
		t.Fatal(err)
	}
	return s, r
}

// A spec on a gitignored file is a warning that names the spec and says lint
// depends on the checkout; a tracked spec is not.
func TestIgnoredSpecWarns(t *testing.T) {
	s, r := specRepo(t, "built/page.html", "tracked.go")
	got := ignoredIssues(lintStore(t, s))
	if len(got) != 1 {
		t.Fatalf("issues %+v", got)
	}
	i := got[0]
	if i.Severity != SeverityWarning || i.RecordID != r.ID || !strings.Contains(i.Message, "built/page.html") || !strings.Contains(i.Message, "depends on the checkout") {
		t.Fatalf("issue %+v", i)
	}
}

// A tracked file that matches an ignore pattern is not ignored.
func TestTrackedIgnoredMatchIsQuiet(t *testing.T) {
	s, _ := specRepo(t, "built/page.html")
	if out, err := exec.Command("git", "-C", s.Root, "add", "-f", "built/page.html").CombinedOutput(); err != nil {
		t.Fatalf("%v %s", err, out)
	}
	if got := ignoredIssues(lintStore(t, s)); len(got) != 0 {
		t.Fatalf("tracked file flagged: %+v", got)
	}
}

// Outside a git work tree lint reports nothing about ignores.
func TestIgnoredSpecOutsideGit(t *testing.T) {
	root := t.TempDir()
	os.WriteFile(filepath.Join(root, "a.go"), []byte("x"), 0o644)
	r := newRecord("prov-2026-0000c002", record.StatusOpen)
	r.AssociatedSpecs = []record.AssociatedSpec{{Path: "a.go", Type: "unit"}}
	s := &record.FilesStore{Root: root}
	if err := s.Save(r, nil); err != nil {
		t.Fatal(err)
	}
	if got := ignoredIssues(lintStore(t, s)); len(got) != 0 {
		t.Fatalf("flagged outside git: %+v", got)
	}
}
