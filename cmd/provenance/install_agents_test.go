package main

import (
	"io/fs"
	"path/filepath"
	"strings"
	"testing"

	"github.com/livecodelife/provenance/assets"
)

var pluginAgents = []string{"provenance-records", "provenance-record-reviewer"}

// TestPluginShipsRecordAgents: the plugin embeds the record writer and the
// record reviewer as agents, each named after its file.
func TestPluginShipsRecordAgents(t *testing.T) {
	for _, name := range pluginAgents {
		data, err := assets.Files.ReadFile("plugins/provenance/agents/" + name + ".md")
		if err != nil {
			t.Fatalf("agent %s is not embedded: %v", name, err)
		}
		text := string(data)
		if !strings.HasPrefix(text, "---\n") || !strings.Contains(text, "\nname: "+name+"\n") || !strings.Contains(text, "\ndescription: ") {
			t.Errorf("agent %s needs front matter with its name and a description:\n%.200s", name, text)
		}
	}
}

// TestRecordAgentsAreNeutral: the agent texts say nothing about any generator
// and nothing about regenerating code; they apply to any language.
func TestRecordAgentsAreNeutral(t *testing.T) {
	err := fs.WalkDir(assets.Files, "plugins/provenance/agents", func(p string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		data, _ := assets.Files.ReadFile(p)
		lower := strings.ToLower(string(data))
		for _, word := range []string{"sedum", "regenerat", "haiku", "sonnet", "opus", "grown by", "linespec"} {
			if strings.Contains(lower, word) {
				t.Errorf("%s mentions %q", p, word)
			}
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}

// TestReviewerAgentIsReadOnly: the reviewer cannot write files.
func TestReviewerAgentIsReadOnly(t *testing.T) {
	data, err := assets.Files.ReadFile("plugins/provenance/agents/provenance-record-reviewer.md")
	if err != nil {
		t.Fatal(err)
	}
	for _, line := range strings.Split(string(data), "\n") {
		if strings.HasPrefix(line, "tools:") {
			for _, bad := range []string{"Write", "Edit", "MultiEdit", "NotebookEdit"} {
				if strings.Contains(line, bad) {
					t.Errorf("the reviewer must be read-only, tools line: %s", line)
				}
			}
			return
		}
	}
	t.Error("the reviewer names no tools line")
}

// TestInstallPluginWritesTheAgents: install-plugin plans and writes both agents
// under the plugin's agents directory, and rewrites nothing on a second run.
func TestInstallPluginWritesTheAgents(t *testing.T) {
	root, cfg := v2Repo(t, false, nil)
	out := mustRun(t, "-c", cfg, "install-plugin")
	for _, name := range pluginAgents {
		rel := ".claude/plugins/provenance/agents/" + name + ".md"
		if !strings.Contains(out, rel) {
			t.Errorf("plan does not name %s:\n%s", rel, out)
		}
		want, _ := assets.Files.ReadFile("plugins/provenance/agents/" + name + ".md")
		if got := readFile(t, filepath.Join(root, filepath.FromSlash(rel))); got != string(want) {
			t.Errorf("%s was not installed as embedded", rel)
		}
	}
	if again := mustRun(t, "-c", cfg, "install-plugin"); !strings.Contains(again, "0 written") {
		t.Errorf("a second install must write nothing:\n%s", again)
	}
}

// TestDocsMentionTheRecordAgents: the skill and the README tell the reader
// about both agents.
func TestDocsMentionTheRecordAgents(t *testing.T) {
	skill, err := assets.Files.ReadFile("skills/provenance/SKILL.md")
	if err != nil {
		t.Fatal(err)
	}
	readme := readFile(t, filepath.Join("..", "..", "README.md"))
	for _, name := range pluginAgents {
		if !strings.Contains(string(skill), name) {
			t.Errorf("the skill does not mention %s", name)
		}
		if !strings.Contains(readme, name) {
			t.Errorf("the README does not mention %s", name)
		}
	}
}
