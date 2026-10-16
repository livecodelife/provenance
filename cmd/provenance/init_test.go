package main

import (
	"bytes"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"gopkg.in/yaml.v3"
)

// configYAMLKeys are the yaml keys of the CLI's own config type: the keys its
// loader knows.
func configYAMLKeys() map[string]bool {
	keys := map[string]bool{}
	typ := reflect.TypeOf(config{})
	for i := 0; i < typ.NumField(); i++ {
		name := strings.Split(typ.Field(i).Tag.Get("yaml"), ",")[0]
		if name != "" {
			keys[name] = true
		}
	}
	return keys
}

func TestInitWritesAConfigTheLoaderKnows(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("SEDUM_BIN", filepath.Join(dir, "no-such-sedum"))
	out := mustRun(t, "init", "--dir", dir)
	for _, want := range []string{"created    .provenance.yml", "created    provenance/.gitkeep", "provenance hooks install", "provenance install-skills", "provenance install-plugin", "provenance create --type blueprint", "provenance next", "not a git repository"} {
		if !strings.Contains(out, want) {
			t.Errorf("output misses %q:\n%s", want, out)
		}
	}
	path := filepath.Join(dir, ".provenance.yml")
	data := readFile(t, path)

	// Strictly into the CLI's own config type: an unknown key is an error.
	dec := yaml.NewDecoder(bytes.NewReader([]byte(data)))
	dec.KnownFields(true)
	var strict config
	if err := dec.Decode(&strict); err != nil {
		t.Fatalf("a key init writes is unknown to the loader: %v", err)
	}
	var raw map[string]any
	if err := yaml.Unmarshal([]byte(data), &raw); err != nil {
		t.Fatal(err)
	}
	known := configYAMLKeys()
	for key := range raw {
		if !known[key] {
			t.Errorf("init writes the key %q, which the loader does not know", key)
		}
	}

	cfg, err := readConfig(path, false)
	if err != nil {
		t.Fatal(err)
	}
	if cfg.Enforcement != "warn" || cfg.Dir != "provenance" || !cfg.RunAssociatedSpecsOnComplete || cfg.CommitTagRequired || cfg.CommitOnStatusChange || cfg.WriteRestriction == nil || *cfg.WriteRestriction || cfg.OverlapSpecsOnComplete != "warn" {
		t.Fatalf("defaults: %+v", cfg)
	}
	if _, err := os.Stat(filepath.Join(dir, ".git", "hooks")); err == nil {
		t.Fatal("init must not install hooks")
	}
	if _, err := os.Stat(filepath.Join(dir, ".sedum")); err == nil {
		t.Fatal("without sedum nothing is written under .sedum")
	}
	if out := mustRun(t, "-c", path, "lint"); !strings.Contains(out, "0 errors") {
		t.Fatalf("a fresh project lints clean:\n%s", out)
	}
}

func TestInitRefusalsForceAndIdempotence(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("SEDUM_BIN", filepath.Join(dir, "no-such-sedum"))
	mustRun(t, "init", "--dir", dir)
	if out := mustRun(t, "init", "--dir", dir); !strings.Contains(out, "unchanged  .provenance.yml") {
		t.Fatalf("an identical rerun:\n%s", out)
	}

	path := filepath.Join(dir, ".provenance.yml")
	writeFile(t, path, "enforcement: strict\n")
	if _, err := run(t, "init", "--dir", dir); err == nil || readFile(t, path) != "enforcement: strict\n" {
		t.Fatalf("an existing different config must not be overwritten (err %v)", err)
	}
	out := mustRun(t, "init", "--dir", dir, "--force")
	if !strings.Contains(out, "overwrote  .provenance.yml") || readFile(t, path+".bak") != "enforcement: strict\n" || !strings.Contains(readFile(t, path), "enforcement: warn") {
		t.Fatalf("--force keeps a .bak and writes the defaults:\n%s", out)
	}

	legacy := t.TempDir()
	writeFile(t, filepath.Join(legacy, ".linespec.yml"), "provenance:\n  dir: provenance\n")
	if _, err := run(t, "init", "--dir", legacy); err == nil || !strings.Contains(err.Error(), "provenance migrate") {
		t.Fatalf("a v1 config must point at migrate (err %v)", err)
	}
	if _, err := os.Stat(filepath.Join(legacy, ".provenance.yml")); err == nil {
		t.Fatal("nothing is written when init refuses")
	}
	if _, err := run(t, "init", "--dir", filepath.Join(dir, "missing")); err == nil {
		t.Fatal("a missing --dir must be an error")
	}
}

func TestInitSedumStubAndGitNote(t *testing.T) {
	dir := t.TempDir()
	fake := filepath.Join(dir, "fake-sedum")
	writeFile(t, fake, "#!/bin/sh\n")
	if err := os.Chmod(fake, 0o755); err != nil {
		t.Fatal(err)
	}
	t.Setenv("SEDUM_BIN", fake)
	gitRepo(t, dir)
	out := mustRun(t, "init", "--dir", dir)
	ctx := filepath.Join(dir, ".sedum", "context.md")
	if !strings.Contains(out, "created    .sedum/context.md") || !strings.Contains(out, ".sedum/run.log") || strings.Contains(out, "not a git repository") {
		t.Fatalf("with sedum, in a git repository:\n%s", out)
	}
	writeFile(t, ctx, "mine\n")
	mustRun(t, "init", "--dir", dir)
	if readFile(t, ctx) != "mine\n" {
		t.Fatal("an existing .sedum/context.md must never be overwritten")
	}
}

func TestInitExcludesGeneratorFilesWhenSedumIsPresent(t *testing.T) {
	excludes := func(dir string) []string {
		cfg, err := readConfig(filepath.Join(dir, ".provenance.yml"), false)
		if err != nil {
			t.Fatal(err)
		}
		return cfg.ExcludePaths
	}
	has := func(list []string, want string) bool {
		for _, s := range list {
			if s == want {
				return true
			}
		}
		return false
	}

	plain := t.TempDir()
	t.Setenv("SEDUM_BIN", filepath.Join(plain, "no-such-sedum"))
	mustRun(t, "init", "--dir", plain)
	if got := excludes(plain); has(got, "generators/**") || has(got, ".sedum/**") {
		t.Fatalf("without sedum: %v", got)
	}

	dir := t.TempDir()
	fake := filepath.Join(dir, "fake-sedum")
	writeFile(t, fake, "#!/bin/sh\n")
	if err := os.Chmod(fake, 0o755); err != nil {
		t.Fatal(err)
	}
	t.Setenv("SEDUM_BIN", fake)
	writeFile(t, filepath.Join(dir, "a.go"), "package a\n")
	gitRepo(t, dir)
	mustRun(t, "init", "--dir", dir)
	got := excludes(dir)
	if !has(got, "generators/**") || !has(got, ".sedum/**") {
		t.Fatalf("with sedum: %v", got)
	}
	dec := yaml.NewDecoder(bytes.NewReader([]byte(readFile(t, filepath.Join(dir, ".provenance.yml")))))
	dec.KnownFields(true)
	var strict config
	if err := dec.Decode(&strict); err != nil {
		t.Fatalf("the sedum variant has an unknown key: %v", err)
	}
	if out := mustRun(t, "init", "--dir", dir); !strings.Contains(out, "unchanged  .provenance.yml") {
		t.Fatalf("an identical rerun with sedum:\n%s", out)
	}

	cfg := filepath.Join(dir, ".provenance.yml")
	mustRun(t, "-c", cfg, "create", "--type", "blueprint", "--no-edit", "--title", "First", "--intent", "Prove the scope check", "--scope", "a.go", "--scope", "proof.sh")
	ids, _ := filepath.Glob(filepath.Join(dir, "provenance", "prov-*.yml"))
	id := strings.TrimSuffix(filepath.Base(ids[0]), ".yml")
	writeFile(t, filepath.Join(dir, "proof.sh"), "exit 0\n")
	recFile := filepath.Join(dir, "provenance", id+".yml")
	writeFile(t, recFile, strings.Replace(readFile(t, recFile), "associated_specs: []", "associated_specs:\n  - path: proof.sh\n    run_command: sh", 1))
	mustRun(t, "-c", cfg, "open", "--record", id)
	writeFile(t, filepath.Join(dir, "generators", "go", "template.go"), "package x\n")
	gitIn(t, dir, "add", "-A")
	msg := filepath.Join(dir, ".git", "TEST_MSG")
	writeFile(t, msg, "Vendor the generators ["+id+"]\n")
	// The generator and sedum files are tooling: they need no scope. (provenance/.gitkeep
	// is init's placeholder for an empty directory; it is not part of this change.)
	gitIn(t, dir, "reset", "-q", "--", "provenance/.gitkeep")
	if out, err := run(t, "-c", cfg, "check", "--staged", "--message-file", msg); err != nil {
		t.Fatalf("generator and sedum files must pass the scope check:\n%s", out)
	}
	writeFile(t, filepath.Join(dir, "other.go"), "package a\n")
	gitIn(t, dir, "add", "other.go")
	if _, err := run(t, "-c", cfg, "check", "--staged", "--message-file", msg); err == nil {
		t.Fatal("a governed file outside the record's scope must still fail")
	}
}
