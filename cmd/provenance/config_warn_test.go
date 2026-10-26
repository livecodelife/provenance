package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// warnRepo is a git repository with no config at its top level and the
// given files, all committed, so git tracks them.
func warnRepo(t *testing.T, files map[string]string) string {
	t.Helper()
	root := t.TempDir()
	writeFile(t, filepath.Join(root, "README.md"), "hi\n")
	for f, body := range files {
		writeFile(t, filepath.Join(root, f), body)
	}
	gitRepo(t, root)
	return root
}

func nestedWarning(t *testing.T, root string) string {
	t.Helper()
	var buf bytes.Buffer
	warnNestedConfigs(&repo{Root: root}, &buf)
	return buf.String()
}

func TestWarnNestedConfigsNamesTrackedConfigsBelowTheRoot(t *testing.T) {
	root := warnRepo(t, map[string]string{
		"pkg-b/.provenance.yml": "dir: provenance\n",
		"pkg-a/.linespec.yml":   "provenance:\n  dir: provenance\n",
	})
	out := nestedWarning(t, root)
	for _, want := range []string{
		"Warning: no .provenance.yml or .linespec.yml at the repository root (" + root + "); using defaults.",
		"Found config below it: pkg-a/.linespec.yml, pkg-b/.provenance.yml.",
		"Pass -c <path> to use one.",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("warning lacks %q:\n%s", want, out)
		}
	}
	if strings.Count(out, "\n") != 1 || !strings.HasSuffix(out, "\n") {
		t.Errorf("want exactly one line ending in a newline, got %q", out)
	}
}

func TestWarnNestedConfigsIgnoresALegacyFileWithoutAProvenanceBlock(t *testing.T) {
	root := warnRepo(t, map[string]string{
		"pkg-a/.linespec.yml": "linespec:\n  dir: specs\n",
		"pkg-b/.linespec.yml": "provenance: [\n",
	})
	if out := nestedWarning(t, root); out != "" {
		t.Errorf("a legacy file without a provenance block, or an unreadable one, was reported:\n%s", out)
	}
}

func TestWarnNestedConfigsCapsTheListAtFive(t *testing.T) {
	files := map[string]string{}
	for i := 1; i <= 7; i++ {
		files[fmt.Sprintf("pkg-%d/.provenance.yml", i)] = "dir: provenance\n"
	}
	out := nestedWarning(t, warnRepo(t, files))
	want := "Found config below it: pkg-1/.provenance.yml, pkg-2/.provenance.yml, pkg-3/.provenance.yml, pkg-4/.provenance.yml, pkg-5/.provenance.yml, ...."
	if !strings.Contains(out, want) {
		t.Errorf("want %q in:\n%s", want, out)
	}
	if strings.Contains(out, "pkg-6") {
		t.Errorf("more than five configs listed:\n%s", out)
	}
}

func TestWarnNestedConfigsIsSilentWhenThereIsNothingToSay(t *testing.T) {
	t.Run("only an untracked nested config", func(t *testing.T) {
		root := warnRepo(t, nil)
		writeFile(t, filepath.Join(root, "pkg-a", ".provenance.yml"), "dir: provenance\n")
		if out := nestedWarning(t, root); out != "" {
			t.Errorf("an untracked config was reported:\n%s", out)
		}
	})
	t.Run("only a root config", func(t *testing.T) {
		root := warnRepo(t, map[string]string{".provenance.yml": "dir: provenance\n"})
		if out := nestedWarning(t, root); out != "" {
			t.Errorf("a root config was reported:\n%s", out)
		}
	})
	t.Run("a config is already in use", func(t *testing.T) {
		root := warnRepo(t, map[string]string{"pkg-a/.provenance.yml": "dir: provenance\n"})
		var buf bytes.Buffer
		warnNestedConfigs(&repo{Root: root, ConfigPath: filepath.Join(root, "pkg-a", ".provenance.yml")}, &buf)
		if buf.Len() != 0 {
			t.Errorf("a command that used a config was warned:\n%s", buf.String())
		}
	})
	t.Run("not a git repository", func(t *testing.T) {
		dir := t.TempDir()
		writeFile(t, filepath.Join(dir, "pkg-a", ".provenance.yml"), "dir: provenance\n")
		if out := nestedWarning(t, dir); out != "" {
			t.Errorf("a directory outside git was reported:\n%s", out)
		}
	})
}

// stderrOf runs fn with os.Stderr captured.
func stderrOf(t *testing.T, fn func()) string {
	t.Helper()
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	saved := os.Stderr
	os.Stderr = w
	done := make(chan string)
	go func() {
		data, _ := io.ReadAll(r)
		done <- string(data)
	}()
	fn()
	os.Stderr = saved
	w.Close()
	return <-done
}

// TestOpenRepoWarnsOnlyWhenItUsedNoConfig: a command run in a repository with
// a nested config and no config at the root warns on stderr and still runs
// with defaults; with -c it does not warn.
func TestOpenRepoWarnsOnlyWhenItUsedNoConfig(t *testing.T) {
	root := warnRepo(t, map[string]string{"pkg-a/.provenance.yml": "dir: provenance\n"})
	t.Chdir(root)

	var out string
	var err error
	got := stderrOf(t, func() { out, err = run(t, "lint") })
	if err != nil || !strings.Contains(out, "Linting 0 provenance records") {
		t.Errorf("lint without a config did not run with defaults (err %v):\n%s", err, out)
	}
	if !strings.Contains(got, "Found config below it: pkg-a/.provenance.yml") {
		t.Errorf("no warning on stderr:\n%s", got)
	}
	if strings.Contains(out, "Warning:") {
		t.Errorf("the warning went to stdout:\n%s", out)
	}

	got = stderrOf(t, func() { _, err = run(t, "-c", filepath.Join(root, "pkg-a", ".provenance.yml"), "lint") })
	if err != nil || strings.Contains(got, "Warning: no .provenance.yml") {
		t.Errorf("a command given -c was warned (err %v):\n%s", err, got)
	}
}
