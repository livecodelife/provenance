package main

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestEditSetSpec(t *testing.T) {
	root, cfg := v2Repo(t, false, map[string]string{draftID: specRecord(draftID)})
	writeFile(t, filepath.Join(root, "spec.sh"), "exit 0\n")
	writeFile(t, filepath.Join(root, "other.sh"), "exit 0 # other\n")
	mustRun(t, "-c", cfg, "open", "--record", draftID)

	out := mustRun(t, "-c", cfg, "edit", draftID, "--set-spec", "spec.sh=other.sh", "--reason", "move the proof")
	if !strings.Contains(out, draftID+" is now version 2") {
		t.Fatalf("output:\n%s", out)
	}
	specs := load(t, root, draftID).AssociatedSpecs
	if len(specs) != 1 || specs[0].Path != "other.sh" || specs[0].RunCommand != "sh spec.sh" {
		t.Fatalf("specs after path change: %+v", specs)
	}

	mustRun(t, "-c", cfg, "edit", draftID, "--set-spec", "other.sh=spec.sh:sh spec.sh -x", "--reason", "move back, new command")
	specs = load(t, root, draftID).AssociatedSpecs
	if specs[0].Path != "spec.sh" || specs[0].RunCommand != "sh spec.sh -x" {
		t.Fatalf("specs after path and command change: %+v", specs)
	}

	out, err := run(t, "-c", cfg, "edit", draftID, "--set-spec", "nope.sh=x.sh", "--reason", "x")
	if err == nil || len(versions(t, root, draftID)) != 3 {
		t.Fatalf("an unknown old path must fail and write nothing (err %v)\n%s", err, out)
	}
	if _, err := run(t, "-c", cfg, "edit", draftID, "--set-spec", "spec.sh", "--reason", "x"); err == nil {
		t.Fatal("a malformed --set-spec must fail")
	}
}
