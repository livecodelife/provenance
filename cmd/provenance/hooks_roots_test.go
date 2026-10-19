package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

// A .linespec.yml without a provenance block is LineSpec's own runner config,
// not a provenance root: committing next to one must still be judged by the
// repository's .provenance.yml.
func TestHooksIgnoreLinespecRunnerConfig(t *testing.T) {
	bin := filepath.Join(t.TempDir(), "provenance")
	build := exec.Command("go", "build", "-o", bin, ".")
	build.Env = append(os.Environ(), "GOFLAGS=-modcacherw")
	if out, err := build.CombinedOutput(); err != nil {
		t.Fatalf("build: %v\n%s", err, out)
	}
	t.Setenv("PROVENANCE", bin)
	root, cfg, _ := hooksRepo(t)
	writeFile(t, cfg, "dir: provenance\ncommit_tag_required: true\nexclude_paths: ['README.md', 'linespec/**']\n")
	mustRun(t, "-c", cfg, "hooks", "install")

	writeFile(t, filepath.Join(root, "linespec", ".linespec.yml"), "specs:\n  - name: health\n    run: curl -f localhost:8000/healthz\n")
	stage(t, root, map[string]string{"a.go": "package a // 1\n", "linespec/health.yml": "name: health\n"})
	c := exec.Command("git", "-C", root, "commit", "-q", "-m", "Change a ["+chkID+"]")
	c.Env = append(os.Environ(), "GIT_DIR="+filepath.Join(root, ".git"), "GIT_WORK_TREE="+root)
	if out, err := c.CombinedOutput(); err != nil {
		t.Fatalf("a tagged commit next to a LineSpec runner config was refused:\n%s", out)
	}
}
