package main

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/livecodelife/provenance/assets"
)

const discoverPath = "plugins/provenance/commands/discover.md"

func discoverText(t *testing.T) string {
	t.Helper()
	data, err := assets.Files.ReadFile(discoverPath)
	if err != nil {
		t.Fatalf("the /discover command is not embedded: %v", err)
	}
	return string(data)
}

// frontMatter returns the property lines between the leading and the closing
// three-dash lines of text, or fails when it has none.
func frontMatter(t *testing.T, text string) map[string]string {
	t.Helper()
	if !strings.HasPrefix(text, "---\n") {
		t.Fatalf("no front matter:\n%.200s", text)
	}
	end := strings.Index(text[4:], "\n---\n")
	if end < 0 {
		t.Fatalf("front matter is not closed:\n%.200s", text)
	}
	props := map[string]string{}
	for _, line := range strings.Split(text[4:4+end], "\n") {
		if k, v, ok := strings.Cut(line, ":"); ok {
			props[strings.TrimSpace(k)] = strings.TrimSpace(v)
		}
	}
	return props
}

func TestDiscoverCommandIsEmbedded(t *testing.T) {
	text := discoverText(t)
	for _, want := range []string{"$ARGUMENTS", "provenance-records", "provenance-record-reviewer", "draft", "git history", "top-level folder", "already governed", "ask"} {
		if !strings.Contains(strings.ToLower(text), strings.ToLower(want)) {
			t.Errorf("the /discover text does not say %q", want)
		}
	}
}

func TestDiscoverCommandFrontMatter(t *testing.T) {
	props := frontMatter(t, discoverText(t))
	if props["description"] == "" {
		t.Error("front matter has no non-empty description")
	}
	if props["argument-hint"] == "" {
		t.Error("front matter has no argument-hint")
	}
}

func TestDiscoverCommandNamesRealCommands(t *testing.T) {
	found := regexp.MustCompile("`provenance ([a-z][a-z-]*)").FindAllStringSubmatch(discoverText(t), -1)
	if len(found) == 0 {
		t.Fatal("the /discover text names no provenance command")
	}
	for _, m := range found {
		if c, _, err := newRootCmd().Find([]string{m[1]}); err != nil || c.Name() != m[1] {
			t.Errorf("the /discover text names `provenance %s`, which is not a command", m[1])
		}
	}
}

func TestDiscoverCommandIsNeutral(t *testing.T) {
	lower := strings.ToLower(discoverText(t))
	for _, word := range []string{"sedum", "regenerat", "tree-sitter", "linespec"} {
		if strings.Contains(lower, word) {
			t.Errorf("the /discover text mentions %q", word)
		}
	}
}

func TestInstallPluginWritesDiscover(t *testing.T) {
	root, cfg := v2Repo(t, false, nil)
	rel := ".claude/plugins/provenance/commands/discover.md"
	out := mustRun(t, "-c", cfg, "install-plugin")
	if !strings.Contains(out, rel) {
		t.Errorf("plan does not name %s:\n%s", rel, out)
	}
	if got := readFile(t, filepath.Join(root, filepath.FromSlash(rel))); got != discoverText(t) {
		t.Errorf("%s was not installed as embedded", rel)
	}
	dest := filepath.Join(root, filepath.FromSlash(rel))
	info, err := os.Stat(dest)
	if err != nil || info.Mode().Perm() != 0o644 {
		t.Errorf("%s must be written with mode 0644 (err %v)", rel, err)
	}
	if again := mustRun(t, "-c", cfg, "install-plugin"); !strings.Contains(again, "0 written") {
		t.Errorf("an identical file must be left unchanged:\n%s", again)
	}
	writeFile(t, dest, "mine\n")
	if _, err := run(t, "-c", cfg, "install-plugin"); err == nil || readFile(t, dest) != "mine\n" {
		t.Errorf("an existing different file must not be overwritten (err %v)", err)
	}
	mustRun(t, "-c", cfg, "install-plugin", "--force")
	if readFile(t, dest) != discoverText(t) {
		t.Errorf("--force must overwrite the different file")
	}
}

func TestDocsMentionDiscover(t *testing.T) {
	skill, err := assets.Files.ReadFile("skills/provenance/SKILL.md")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(skill), "/discover") {
		t.Error("the skill does not mention /discover")
	}
}
