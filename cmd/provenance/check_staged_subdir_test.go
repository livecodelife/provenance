package main

import (
	"path/filepath"
	"strings"
	"testing"
)

// subdirRepo is a git repository whose config sits two levels below the top:
// packages/service holds the config, the record and a.go. The record has two
// specs: spec.sh inside the package and ../../spec/outside.sh outside it.
func subdirRepo(t *testing.T) (root string, p nestedPkg) {
	t.Helper()
	root = t.TempDir()
	p = nestedPkg{dir: "packages/service", id: "prov-2026-0000c201"}
	base := filepath.Join(root, p.dir)
	p.cfg = filepath.Join(base, ".provenance.yml")
	writeFile(t, p.cfg, "dir: provenance\ncommit_tag_required: true\nexclude_paths: ['README.md']\n")
	rec := strings.Replace(v2Draft(p.id, "draft"), "affected_scope:\n  - a.go\n", "affected_scope:\n  - a.go\n  - spec.sh\n", 1)
	rec = strings.Replace(rec, "associated_specs: []\n", "associated_specs:\n  - path: spec.sh\n    run_command: sh\n  - path: ../../spec/outside.sh\n    run_command: sh\n", 1)
	writeFile(t, filepath.Join(base, "provenance", p.id+".yml"), rec)
	for f, body := range map[string]string{"a.go": "package a\n", "spec.sh": "true\n"} {
		writeFile(t, filepath.Join(base, f), body)
	}
	writeFile(t, filepath.Join(root, "spec", "outside.sh"), "true\n")
	mustRun(t, "-c", p.cfg, "open", "--record", p.id)
	gitRepo(t, root)
	return root, p
}

// TestCheckStagedSpecOutsideTheConfigDirectory: with the config below the git
// top level, a spec that lives outside the config's directory is found in the
// staged tree, so check --staged raises no changed-spec finding for an
// unchanged spec and agrees with lint.
func TestCheckStagedSpecOutsideTheConfigDirectory(t *testing.T) {
	root, p := subdirRepo(t)
	if out := mustRun(t, "-c", p.cfg, "lint"); !strings.Contains(out, "0 errors") {
		t.Fatalf("lint must pass before any change:\n%s", out)
	}
	stage(t, root, map[string]string{p.dir + "/a.go": "package a // changed\n"})
	out, ok := nestedCheck(t, p, "Change a ["+p.id+"]")
	if !ok || strings.Contains(out, "PROV024") {
		t.Fatalf("an unchanged spec outside the config directory must not be reported (ok %v):\n%s", ok, out)
	}
}

// TestCheckStagedSpecInsideStillFine: a spec inside the package keeps working,
// and a staged change to the outside spec is still reported as a changed spec.
func TestCheckStagedSpecInsideStillFine(t *testing.T) {
	root, p := subdirRepo(t)
	stage(t, root, map[string]string{p.dir + "/spec.sh": "true # changed\n"})
	out, ok := nestedCheck(t, p, "Change spec ["+p.id+"]")
	if ok || !strings.Contains(out, "PROV024") {
		t.Fatalf("a changed spec inside the package must be reported (ok %v):\n%s", ok, out)
	}
	gitIn(t, root, "reset", "-q")

	stage(t, root, map[string]string{"spec/outside.sh": "true # changed\n"})
	out, ok = nestedCheck(t, p, "Change outside spec ["+p.id+"]")
	if ok || !strings.Contains(out, "PROV024") {
		t.Fatalf("a changed spec outside the package must be reported (ok %v):\n%s", ok, out)
	}
}

// TestCheckCommitAndRangeSpecOutsideTheConfigDirectory: --commit and --range
// read the same outside spec from the committed tree and agree with --staged.
func TestCheckCommitAndRangeSpecOutsideTheConfigDirectory(t *testing.T) {
	root, p := subdirRepo(t)
	base := gitIn(t, root, "rev-parse", "HEAD")
	stage(t, root, map[string]string{p.dir + "/a.go": "package a // changed\n"})
	gitIn(t, root, "commit", "-q", "-m", "Change a ["+p.id+"]")
	if out, err := run(t, "-c", p.cfg, "check", "--commit", "HEAD"); err != nil || strings.Contains(out, "PROV024") {
		t.Fatalf("--commit (err %v):\n%s", err, out)
	}
	if out, err := run(t, "-c", p.cfg, "check", "--range", base+"..HEAD"); err != nil || strings.Contains(out, "PROV024") {
		t.Fatalf("--range (err %v):\n%s", err, out)
	}
}

// TestCheckSpecsWithTheConfigAtTheTopLevel: nothing regresses when the config
// sits at the git top level, whether the spec is in a subdirectory or beside
// the config, for --staged, --commit and --range.
func TestCheckSpecsWithTheConfigAtTheTopLevel(t *testing.T) {
	root := t.TempDir()
	id := "prov-2026-0000c202"
	cfg := filepath.Join(root, ".provenance.yml")
	writeFile(t, cfg, "dir: provenance\ncommit_tag_required: true\nexclude_paths: ['README.md']\n")
	rec := strings.Replace(v2Draft(id, "draft"), "affected_scope:\n  - a.go\n", "affected_scope:\n  - a.go\n", 1)
	rec = strings.Replace(rec, "associated_specs: []\n", "associated_specs:\n  - path: spec.sh\n    run_command: sh\n  - path: spec/inner.sh\n    run_command: sh\n", 1)
	writeFile(t, filepath.Join(root, "provenance", id+".yml"), rec)
	for f, body := range map[string]string{"a.go": "package a\n", "spec.sh": "true\n", "spec/inner.sh": "true\n"} {
		writeFile(t, filepath.Join(root, f), body)
	}
	mustRun(t, "-c", cfg, "open", "--record", id)
	gitRepo(t, root)
	p := nestedPkg{dir: ".", cfg: cfg, id: id}

	base := gitIn(t, root, "rev-parse", "HEAD")
	stage(t, root, map[string]string{"a.go": "package a // changed\n"})
	if out, ok := nestedCheck(t, p, "Change a ["+id+"]"); !ok || strings.Contains(out, "PROV024") {
		t.Fatalf("--staged (ok %v):\n%s", ok, out)
	}
	gitIn(t, root, "commit", "-q", "-m", "Change a ["+id+"]")
	if out, err := run(t, "-c", cfg, "check", "--commit", "HEAD"); err != nil || strings.Contains(out, "PROV024") {
		t.Fatalf("--commit (err %v):\n%s", err, out)
	}
	if out, err := run(t, "-c", cfg, "check", "--range", base+"..HEAD"); err != nil || strings.Contains(out, "PROV024") {
		t.Fatalf("--range (err %v):\n%s", err, out)
	}
}
