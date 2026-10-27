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

func TestInitTemplateExplainsThePathPlaceholder(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("SEDUM_BIN", filepath.Join(dir, "no-such-sedum"))
	mustRun(t, "init", "--dir", dir)
	data := readFile(t, filepath.Join(dir, ".provenance.yml"))
	if !strings.Contains(data, "run_command is run as `<run_command> <path>`, or with {{path}} substituted") {
		t.Fatalf("the template must explain {{path}}:\n%s", data)
	}
	dec := yaml.NewDecoder(bytes.NewReader([]byte(data)))
	dec.KnownFields(true)
	var strict config
	if err := dec.Decode(&strict); err != nil {
		t.Fatalf("the comment must add no key: %v", err)
	}
}
