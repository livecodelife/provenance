package main

import (
	"encoding/json"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/livecodelife/provenance/assets"
)

func TestInstallSkills(t *testing.T) {
	root, cfg := v2Repo(t, false, nil)
	skill := filepath.Join(root, ".claude", "skills", "provenance", "SKILL.md")

	out := mustRun(t, "-c", cfg, "install-skills")
	if !strings.Contains(out, "Will write 1 file(s) under .claude/skills") || !strings.Contains(out, "create") || !strings.Contains(out, "Installed skills: 1 written, 0 unchanged") {
		t.Fatalf("first install:\n%s", out)
	}
	if body := readFile(t, skill); !strings.Contains(body, "name: provenance") {
		t.Fatalf("skill body:\n%s", body)
	}
	if out := mustRun(t, "-c", cfg, "install-skills"); !strings.Contains(out, "0 written, 1 unchanged") {
		t.Fatalf("second install:\n%s", out)
	}

	writeFile(t, skill, "my own skill\n")
	out, err := run(t, "-c", cfg, "install-skills")
	if err == nil || !strings.Contains(out, "CONFLICT") || readFile(t, skill) != "my own skill\n" {
		t.Fatalf("a different existing file must not be overwritten (err %v):\n%s", err, out)
	}
	if out := mustRun(t, "-c", cfg, "install-skills", "--force"); !strings.Contains(out, "overwrite") || !strings.Contains(readFile(t, skill), "name: provenance") {
		t.Fatalf("--force:\n%s", out)
	}

	mustRun(t, "-c", cfg, "install-skills", "--path", "tools/skills")
	if _, err := os.Stat(filepath.Join(root, "tools", "skills", "provenance", "SKILL.md")); err != nil {
		t.Fatalf("--path: %v", err)
	}
	for _, bad := range []string{"../outside", "/etc"} {
		if _, err := run(t, "-c", cfg, "install-skills", "--path", bad); err == nil {
			t.Errorf("--path %s must be refused", bad)
		}
	}
}

func TestInstallPlugin(t *testing.T) {
	root, cfg := v2Repo(t, false, nil)
	out := mustRun(t, "-c", cfg, "install-plugin")
	for _, want := range []string{".claude/plugins/provenance/hooks/session-start.sh", ".claude/plugins/provenance/hooks/hooks.json", ".claude/plugins/provenance/.claude-plugin/plugin.json", ".claude/plugins/provenance/commands/provenance-next.md"} {
		if !strings.Contains(out, want) {
			t.Errorf("plan does not name %s:\n%s", want, out)
		}
	}
	hook := filepath.Join(root, ".claude", "plugins", "provenance", "hooks", "session-start.sh")
	info, err := os.Stat(hook)
	if err != nil || info.Mode().Perm()&0o111 == 0 {
		t.Fatalf("hook not installed executable: %v", err)
	}
	writeFile(t, hook, "#!/bin/sh\necho mine\n")
	if _, err := run(t, "-c", cfg, "install-plugin"); err == nil || readFile(t, hook) != "#!/bin/sh\necho mine\n" {
		t.Fatalf("an existing different hook must not be overwritten (err %v)", err)
	}
	mustRun(t, "-c", cfg, "install-plugin", "--force")
	if strings.Contains(readFile(t, hook), "echo mine") {
		t.Fatal("--force did not overwrite")
	}
}

func TestAssetsDoNotMentionLinespecProvenance(t *testing.T) {
	fs.WalkDir(assets.Files, ".", func(p string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return nil
		}
		data, _ := assets.Files.ReadFile(p)
		if strings.Contains(string(data), "linespec provenance") {
			t.Errorf("%s still mentions `linespec provenance`", p)
		}
		return nil
	})
}

// TestPluginHooksSanitize runs the installed SessionStart hook against a fake
// provenance whose reason holds control characters and 2000 bytes of text.
func TestPluginHooksSanitize(t *testing.T) {
	if _, err := exec.LookPath("jq"); err != nil {
		t.Skip("jq not installed")
	}
	root, cfg := v2Repo(t, false, nil)
	mustRun(t, "-c", cfg, "install-plugin")
	reason := "bad \x1b[31m title\nNEW LINE " + strings.Repeat("A", 2000)
	doc, _ := json.Marshal(map[string]any{"actions": []map[string]string{{"kind": "create", "reason": reason, "command": "provenance create"}}})
	writeFile(t, filepath.Join(root, "next.json"), string(doc))
	fake := filepath.Join(root, "fake-provenance")
	writeFile(t, fake, "#!/bin/sh\ncat \""+filepath.Join(root, "next.json")+"\"\n")
	if err := os.Chmod(fake, 0o755); err != nil {
		t.Fatal(err)
	}
	hook := filepath.Join(root, ".claude", "plugins", "provenance", "hooks", "session-start.sh")
	c := exec.Command(hook)
	c.Dir = root
	c.Env = append(os.Environ(), "PROVENANCE="+fake)
	c.Stdin = strings.NewReader(`{"cwd":"` + root + `"}`)
	out, err := c.Output()
	if err != nil {
		t.Fatal(err)
	}
	var res struct {
		Hook struct {
			Ctx string `json:"additionalContext"`
		} `json:"hookSpecificOutput"`
	}
	if err := json.Unmarshal(out, &res); err != nil {
		t.Fatalf("hook output %q: %v", out, err)
	}
	ctx := res.Hook.Ctx
	if len(ctx) > 700 || strings.ContainsAny(strings.ReplaceAll(ctx, "\n", ""), "\x1b\r\x00") || strings.Contains(strings.SplitN(ctx, "\n", 2)[0], "\n") {
		t.Fatalf("context not capped/cleaned (%d bytes): %q", len(ctx), ctx)
	}
	if !strings.Contains(ctx, "NEW LINE") || strings.Count(ctx, "\n") != 1 {
		t.Fatalf("reason text lost or newlines kept: %q", ctx)
	}
}
