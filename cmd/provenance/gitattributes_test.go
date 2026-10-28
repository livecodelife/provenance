package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var wantAttrLines = []string{
	"provenance/.versions/** linguist-generated=true -diff",
	"provenance/.evidence/** linguist-generated=true -diff",
	".provenance/versions.jsonl linguist-generated=true -diff",
}

func attrFile(root string) string { return filepath.Join(root, ".gitattributes") }

func TestEnsureGitattributesCreatesAppendsAndIsIdempotent(t *testing.T) {
	root := t.TempDir()
	status, err := ensureGitattributes(root, "provenance")
	if err != nil || status != "created" {
		t.Fatalf("first run: %q, %v", status, err)
	}
	if got := readFile(t, attrFile(root)); got != strings.Join(wantAttrLines, "\n")+"\n" {
		t.Fatalf("created file:\n%s", got)
	}
	if status, err := ensureGitattributes(root, "provenance"); err != nil || status != "unchanged" {
		t.Fatalf("second run: %q, %v", status, err)
	}

	// Existing content (no final newline, one line already present) is kept; only
	// the missing lines are appended.
	other := t.TempDir()
	writeFile(t, attrFile(other), "*.png binary\n  "+wantAttrLines[1]+"  \nmerge=union")
	if status, err := ensureGitattributes(other, "provenance"); err != nil || status != "updated" {
		t.Fatalf("append run: %q, %v", status, err)
	}
	want := "*.png binary\n  " + wantAttrLines[1] + "  \nmerge=union\n" + wantAttrLines[0] + "\n" + wantAttrLines[2] + "\n"
	if got := readFile(t, attrFile(other)); got != want {
		t.Fatalf("appended file:\n%q\nwant\n%q", got, want)
	}
}

func TestGitattributesFollowsTheRecordsDir(t *testing.T) {
	root := t.TempDir()
	if _, err := ensureGitattributes(root, "specs/"); err != nil {
		t.Fatal(err)
	}
	got := readFile(t, attrFile(root))
	for _, want := range []string{"specs/.versions/** linguist-generated=true -diff\n", "specs/.evidence/** linguist-generated=true -diff\n", ".provenance/versions.jsonl linguist-generated=true -diff\n"} {
		if !strings.Contains(got, want) {
			t.Errorf("missing %q in:\n%s", want, got)
		}
	}
}

func TestInitWritesGitattributes(t *testing.T) {
	dir := t.TempDir()
	out := mustRun(t, "init", "--dir", dir)
	if !strings.Contains(out, "created    .gitattributes") {
		t.Errorf("init output misses the created line:\n%s", out)
	}
	if got := readFile(t, attrFile(dir)); got != strings.Join(wantAttrLines, "\n")+"\n" {
		t.Errorf(".gitattributes:\n%s", got)
	}
	again := mustRun(t, "init", "--dir", dir)
	if !strings.Contains(again, "unchanged  .gitattributes") {
		t.Errorf("second init output:\n%s", again)
	}
}

func TestMigrateWritesGitattributesButPlanDoesNot(t *testing.T) {
	root := migrateFixture(t)
	cfg := filepath.Join(root, ".linespec.yml")
	mustRun(t, "-c", cfg, "migrate", "--plan")
	if _, err := os.Stat(attrFile(root)); !os.IsNotExist(err) {
		t.Fatalf("--plan wrote .gitattributes (%v)", err)
	}
	out := mustRun(t, "-c", cfg, "migrate")
	if !strings.Contains(out, ".gitattributes created") {
		t.Errorf("migrate output misses the .gitattributes line:\n%s", out)
	}
	if got := readFile(t, attrFile(root)); got != strings.Join(wantAttrLines, "\n")+"\n" {
		t.Errorf(".gitattributes:\n%s", got)
	}
}

func TestMigrateOnV2HintsAtMissingLinesAndWritesNothing(t *testing.T) {
	root := migrateFixture(t)
	mustRun(t, "-c", filepath.Join(root, ".linespec.yml"), "migrate")
	if err := os.Remove(attrFile(root)); err != nil {
		t.Fatal(err)
	}
	gitIn(t, root, "add", "-A")
	gitIn(t, root, "commit", "-q", "-m", "migrate")
	cfg := filepath.Join(root, ".provenance.yml")
	out := mustRun(t, "-c", cfg, "migrate")
	if !strings.Contains(out, "Nothing to migrate") || !strings.Contains(out, "Hint: add these lines to .gitattributes") {
		t.Errorf("output:\n%s", out)
	}
	if _, err := os.Stat(attrFile(root)); !os.IsNotExist(err) {
		t.Fatalf("the already-v2 path wrote .gitattributes (%v)", err)
	}
	writeFile(t, attrFile(root), strings.Join(wantAttrLines, "\n")+"\n")
	if out := mustRun(t, "-c", cfg, "migrate"); strings.Contains(out, "Hint") {
		t.Errorf("hint printed although the lines are present:\n%s", out)
	}
}
