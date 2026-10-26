package main

import (
	"path/filepath"
	"strings"
	"testing"
)

// nestedPkg is one package of a repository that has no config at its git top
// level: dir holds the config, which is a .provenance.yml or, with legacy set,
// a .linespec.yml provenance block, and one open record covering a.go.
type nestedPkg struct {
	dir    string
	cfg    string
	id     string
	legacy bool
}

// nestedRepo is a git repository with two packages, pkg-a (.provenance.yml)
// and pkg-b (.linespec.yml), each with its own record and files a.go, b.go,
// spec.sh and README.md (excluded).
func nestedRepo(t *testing.T) (root string, pkgs []nestedPkg) {
	t.Helper()
	root = t.TempDir()
	pkgs = []nestedPkg{
		{dir: "pkg-a", id: "prov-2026-0000c101"},
		{dir: "pkg-b", id: "prov-2026-0000c102", legacy: true},
	}
	for i := range pkgs {
		p := &pkgs[i]
		base := filepath.Join(root, p.dir)
		settings := "dir: provenance\ncommit_tag_required: true\nexclude_paths: ['README.md']\n"
		if p.legacy {
			p.cfg = filepath.Join(base, ".linespec.yml")
			writeFile(t, p.cfg, "provenance:\n  "+strings.ReplaceAll(strings.TrimSuffix(settings, "\n"), "\n", "\n  ")+"\n")
		} else {
			p.cfg = filepath.Join(base, ".provenance.yml")
			writeFile(t, p.cfg, settings)
		}
		rec := strings.Replace(v2Draft(p.id, "draft"), "affected_scope:\n  - a.go\n", "affected_scope:\n  - a.go\n  - spec.sh\n", 1)
		rec = strings.Replace(rec, "associated_specs: []\n", "associated_specs:\n  - path: spec.sh\n    run_command: sh\n", 1)
		writeFile(t, filepath.Join(base, "provenance", p.id+".yml"), rec)
		for f, body := range map[string]string{"a.go": "package a\n", "b.go": "package b\n", "spec.sh": "true\n", "README.md": "hi\n"} {
			writeFile(t, filepath.Join(base, f), body)
		}
		mustRun(t, "-c", p.cfg, "open", "--record", p.id)
	}
	gitRepo(t, root)
	return root, pkgs
}

// nestedCheck runs check --staged for the package's config with message.
func nestedCheck(t *testing.T, p nestedPkg, message string) (string, bool) {
	t.Helper()
	msg := filepath.Join(t.TempDir(), "MSG")
	writeFile(t, msg, message+"\n")
	out, err := run(t, "-c", p.cfg, "check", "--staged", "--message-file", msg)
	return out, err == nil
}

// TestCheckNestedStagedMatchesRelativeToTheConfig: with the config below the
// git top level, a staged file is matched against scope relative to the
// config's directory, and files outside that directory are not governed.
func TestCheckNestedStagedMatchesRelativeToTheConfig(t *testing.T) {
	root, pkgs := nestedRepo(t)
	for _, p := range pkgs {
		stage(t, root, map[string]string{p.dir + "/a.go": "package a // changed\n"})
		if out, ok := nestedCheck(t, p, "Change a ["+p.id+"]"); !ok || !strings.Contains(out, "✓ No forbidden scope violations in staged") {
			t.Errorf("%s: in-scope change refused:\n%s", p.dir, out)
		}
		gitIn(t, root, "reset", "-q")

		stage(t, root, map[string]string{p.dir + "/b.go": "package b // changed\n"})
		out, ok := nestedCheck(t, p, "Change b ["+p.id+"]")
		if ok || !strings.Contains(out, "· b.go") || strings.Contains(out, p.dir+"/b.go") {
			t.Errorf("%s: out-of-scope change passed or named the wrong path (ok %v):\n%s", p.dir, ok, out)
		}
		gitIn(t, root, "reset", "-q")
	}

	other := pkgs[1]
	stage(t, root, map[string]string{other.dir + "/b.go": "package b // elsewhere\n"})
	if out, ok := nestedCheck(t, pkgs[0], "Change elsewhere"); !ok {
		t.Errorf("a file outside pkg-a was governed by pkg-a's config:\n%s", out)
	}
}

// TestCheckNestedCommitAndRange: --commit and --range apply the same rule as
// --staged, including to the record's own files and the version manifest.
func TestCheckNestedCommitAndRange(t *testing.T) {
	root, pkgs := nestedRepo(t)
	for _, p := range pkgs {
		base := gitIn(t, root, "rev-parse", "HEAD")
		stage(t, root, map[string]string{p.dir + "/a.go": "package a // one\n"})
		gitIn(t, root, "commit", "-q", "-m", "Change a ["+p.id+"]")
		stage(t, root, map[string]string{p.dir + "/b.go": "package b // two\n"})
		gitIn(t, root, "commit", "-q", "-m", "Change b ["+p.id+"]")

		if out := mustRun(t, "-c", p.cfg, "check", "--commit", "HEAD~1"); !strings.Contains(out, "✓ No forbidden scope violations in 1 commit(s)") {
			t.Errorf("%s: in-scope commit:\n%s", p.dir, out)
		}
		out, err := run(t, "-c", p.cfg, "check", "--range", base+"..HEAD")
		if err == nil || !strings.Contains(out, "forbids changes to b.go") || strings.Contains(out, p.dir+"/b.go") {
			t.Errorf("%s: range with an out-of-scope commit (err %v):\n%s", p.dir, err, out)
		}
		if out := mustRun(t, "-c", p.cfg, "check", "--range", base+"..HEAD~1"); !strings.Contains(out, "✓") {
			t.Errorf("%s: clean range:\n%s", p.dir, out)
		}
	}
}
