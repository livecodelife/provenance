package main

import (
	"fmt"
	"path/filepath"
	"strings"
	"testing"
)

// A merge commit whose tree is exactly what git's own merge of its two parents
// produces needs no tag and no scope: both sides were checked when they were
// committed. These tests build merges of two tagged branches over a.go, the
// file the fixture record governs, with the two sides editing lines far apart
// so git merges them without help.

// mergeFile is a.go with thirty numbered lines, so edits at the top and at the
// bottom merge cleanly.
func mergeFile(top, bottom string) string {
	var b strings.Builder
	b.WriteString("package a\n")
	for i := 1; i <= 30; i++ {
		line := fmt.Sprintf("// line %d", i)
		if i == 2 && top != "" {
			line = top
		}
		if i == 28 && bottom != "" {
			line = bottom
		}
		b.WriteString(line + "\n")
	}
	return b.String()
}

// mergeScenario is checkRepo with a.go holding mergeFile, one tagged commit on
// a feature branch editing the top and one on the main branch editing the
// bottom. The returned merge is started with --no-commit.
func mergeScenario(t *testing.T) (root, cfg string) {
	t.Helper()
	root, cfg = checkRepo(t)
	stage(t, root, map[string]string{"a.go": mergeFile("", "")})
	gitIn(t, root, "commit", "-q", "-m", "Grow a ["+chkID+"]")
	gitIn(t, root, "checkout", "-q", "-b", "feature")
	stage(t, root, map[string]string{"a.go": mergeFile("// top, from the feature", "")})
	gitIn(t, root, "commit", "-q", "-m", "Edit the top ["+chkID+"]")
	gitIn(t, root, "checkout", "-q", "-")
	stage(t, root, map[string]string{"a.go": mergeFile("", "// bottom, from main")})
	gitIn(t, root, "commit", "-q", "-m", "Edit the bottom ["+chkID+"]")
	return root, cfg
}

func startMerge(t *testing.T, root string) {
	t.Helper()
	if out, err := gitTry(t, root, "merge", "-q", "--no-commit", "--no-ff", "feature"); err != nil {
		t.Fatalf("the merge was meant to be clean: %v\n%s", err, out)
	}
}

func checkCommitted(t *testing.T, cfg string) (string, bool) {
	t.Helper()
	out, err := run(t, "-c", cfg, "check", "--commit", "HEAD")
	return out, err == nil
}

func TestCheckMergeEqualToGitsOwnMergePassesUntagged(t *testing.T) {
	root, cfg := mergeScenario(t)
	startMerge(t, root)
	if !strings.Contains(readFile(t, filepath.Join(root, "a.go")), "from the feature") || !strings.Contains(readFile(t, filepath.Join(root, "a.go")), "from main") {
		t.Fatal("the fixture merge should have combined both edits in a.go")
	}
	if out, ok := check(t, cfg, "Merge feature"); !ok || !strings.Contains(out, "✓ No forbidden scope violations in staged") {
		t.Fatalf("a merge equal to git's own must pass untagged:\n%s", out)
	}
	gitIn(t, root, "commit", "-q", "-m", "Merge feature")
	if out, ok := checkCommitted(t, cfg); !ok {
		t.Fatalf("--commit must agree:\n%s", out)
	}
	if out, err := run(t, "-c", cfg, "check", "--range", "HEAD~1..HEAD"); err != nil {
		t.Fatalf("--range must agree: %v\n%s", err, out)
	}
}

func TestCheckMergeWithAnExtraHandEditIsStillChecked(t *testing.T) {
	root, cfg := mergeScenario(t)
	startMerge(t, root)
	stage(t, root, map[string]string{"a.go": readFile(t, filepath.Join(root, "a.go")) + "// hand edit\n"})
	if out, ok := check(t, cfg, "Merge feature"); ok {
		t.Fatalf("a hand edit in the merged file passed untagged:\n%s", out)
	}
	gitIn(t, root, "commit", "-q", "-m", "Merge feature")
	if out, ok := checkCommitted(t, cfg); ok {
		t.Fatalf("--commit let a hand-edited merge through:\n%s", out)
	}
}

func TestCheckMergeWithAConflictResolvedByHandIsStillChecked(t *testing.T) {
	root, cfg := checkRepo(t)
	stage(t, root, map[string]string{"a.go": mergeFile("", "")})
	gitIn(t, root, "commit", "-q", "-m", "Grow a ["+chkID+"]")
	gitIn(t, root, "checkout", "-q", "-b", "feature")
	stage(t, root, map[string]string{"a.go": mergeFile("// one way", "")})
	gitIn(t, root, "commit", "-q", "-m", "Edit the top ["+chkID+"]")
	gitIn(t, root, "checkout", "-q", "-")
	stage(t, root, map[string]string{"a.go": mergeFile("// the other way", "")})
	gitIn(t, root, "commit", "-q", "-m", "Edit the top too ["+chkID+"]")
	if _, err := gitTry(t, root, "merge", "-q", "--no-commit", "--no-ff", "feature"); err == nil {
		t.Fatal("the fixture merge was meant to conflict")
	}
	stage(t, root, map[string]string{"a.go": mergeFile("// resolved by hand", "")})
	if out, ok := check(t, cfg, "Merge feature"); ok {
		t.Fatalf("a conflict resolved in a.go passed untagged:\n%s", out)
	}
	if out, ok := check(t, cfg, "Merge feature ["+chkID+"]"); !ok {
		t.Fatalf("the same resolution tagged with its record must pass:\n%s", out)
	}
}

// manifestScenario adds one line to the manifest on each side, at the same
// place, so the merge conflicts on that file alone.
func manifestScenario(t *testing.T) (root, cfg string) {
	t.Helper()
	root, cfg = checkRepo(t)
	stage(t, root, map[string]string{"a.go": mergeFile("", "")})
	gitIn(t, root, "commit", "-q", "-m", "Grow a ["+chkID+"]")
	manifest := filepath.Join(root, manifestPath)
	base := readFile(t, manifest)
	gitIn(t, root, "checkout", "-q", "-b", "feature")
	stage(t, root, map[string]string{
		"a.go":       mergeFile("// top, from the feature", ""),
		manifestPath: base + `{"id":"prov-2026-0000f001","version":1,"content_hash":"sha256:f1"}` + "\n",
	})
	gitIn(t, root, "commit", "-q", "-m", "Edit the top ["+chkID+"]")
	gitIn(t, root, "checkout", "-q", "-")
	stage(t, root, map[string]string{
		"a.go":       mergeFile("", "// bottom, from main"),
		manifestPath: base + `{"id":"prov-2026-0000f002","version":1,"content_hash":"sha256:f2"}` + "\n",
	})
	gitIn(t, root, "commit", "-q", "-m", "Edit the bottom ["+chkID+"]")
	out, err := gitTry(t, root, "merge", "-q", "--no-commit", "--no-ff", "feature")
	if err == nil || !strings.Contains(out, manifestPath) {
		t.Fatalf("the fixture merge was meant to conflict on the manifest: %v\n%s", err, out)
	}
	return root, cfg
}

func TestCheckMergeResolvingOnlyTheManifestConflictPasses(t *testing.T) {
	root, cfg := manifestScenario(t)
	merged := `{"id":"prov-2026-0000f001","version":1,"content_hash":"sha256:f1"}` + "\n" + `{"id":"prov-2026-0000f002","version":1,"content_hash":"sha256:f2"}` + "\n"
	manifest := filepath.Join(root, manifestPath)
	head := readFile(t, manifest)
	if i := strings.Index(head, "<<<<<<<"); i < 0 {
		t.Fatal("the manifest should hold conflict markers")
	} else {
		head = head[:i]
	}
	stage(t, root, map[string]string{manifestPath: head + merged})
	if out, ok := check(t, cfg, "Merge feature"); !ok {
		t.Fatalf("keeping both manifest lines is the one allowed resolution:\n%s", out)
	}
	gitIn(t, root, "commit", "-q", "-m", "Merge feature")
	if out, ok := checkCommitted(t, cfg); !ok {
		t.Fatalf("--commit must agree:\n%s", out)
	}
}

func TestCheckMergeResolvingTheManifestAndEditingAFileIsStillChecked(t *testing.T) {
	root, cfg := manifestScenario(t)
	manifest := filepath.Join(root, manifestPath)
	head := readFile(t, manifest)
	head = head[:strings.Index(head, "<<<<<<<")]
	stage(t, root, map[string]string{
		manifestPath: head + `{"id":"prov-2026-0000f001","version":1,"content_hash":"sha256:f1"}` + "\n",
		"a.go":       readFile(t, filepath.Join(root, "a.go")) + "// hand edit\n",
	})
	if out, ok := check(t, cfg, "Merge feature"); ok {
		t.Fatalf("a hand edit beside the manifest resolution passed untagged:\n%s", out)
	}
}

func TestCheckMergeWithAHandEditedManifestButNoConflictIsStillChecked(t *testing.T) {
	root, cfg := mergeScenario(t)
	startMerge(t, root)
	stage(t, root, map[string]string{manifestPath: readFile(t, filepath.Join(root, manifestPath)) + `{"id":"prov-2026-0000f003","version":1,"content_hash":"sha256:f3"}` + "\n"})
	if out, ok := check(t, cfg, "Merge feature"); ok {
		t.Fatalf("only a conflicted manifest may differ from git's merge; a hand-edited one passed:\n%s", out)
	}
}

func TestCheckMergeOfThreeParentsIsStillChecked(t *testing.T) {
	root, cfg := checkRepo(t)
	stage(t, root, map[string]string{"a.go": mergeFile("", "")})
	gitIn(t, root, "commit", "-q", "-m", "Grow a ["+chkID+"]")
	base := gitIn(t, root, "rev-parse", "HEAD")
	for i, edit := range []string{"// first", "// second"} {
		name := fmt.Sprintf("side%d", i)
		gitIn(t, root, "checkout", "-q", "-b", name, base)
		stage(t, root, map[string]string{"a.go": strings.Replace(mergeFile("", ""), fmt.Sprintf("// line %d\n", 2+i*10), edit+"\n", 1)})
		gitIn(t, root, "commit", "-q", "-m", "Edit "+name+" ["+chkID+"]")
	}
	gitIn(t, root, "checkout", "-q", "-b", "main2", base)
	stage(t, root, map[string]string{"a.go": strings.Replace(mergeFile("", ""), "// line 28\n", "// third\n", 1)})
	gitIn(t, root, "commit", "-q", "-m", "Edit main2 ["+chkID+"]")
	if out, err := gitTry(t, root, "merge", "-q", "--no-commit", "--no-ff", "side0", "side1"); err != nil {
		t.Fatalf("the octopus merge was meant to be clean: %v\n%s", err, out)
	}
	if out, ok := check(t, cfg, "Merge both sides"); ok {
		t.Fatalf("a merge of three parents has no exemption:\n%s", out)
	}
}
