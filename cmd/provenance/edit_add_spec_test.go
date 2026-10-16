package main

import (
	"path/filepath"
	"strings"
	"testing"
)

func specSummary(t *testing.T, root, id string) []string {
	t.Helper()
	var out []string
	for _, s := range load(t, root, id).AssociatedSpecs {
		out = append(out, s.Path+"|"+s.RunCommand)
	}
	return out
}

func TestEditAddAndRemoveSpec(t *testing.T) {
	root, cfg := v2Repo(t, false, map[string]string{draftID: v2Draft(draftID, "draft")})
	writeFile(t, filepath.Join(root, "proof.sh"), "exit 0\n")
	writeFile(t, filepath.Join(root, "other.sh"), "exit 0 # other\n")
	writeFile(t, filepath.Join(root, "third.sh"), "exit 0 # third\n")

	// A draft saves in place.
	out := mustRun(t, "-c", cfg, "edit", draftID, "--add-spec", "proof.sh:sh proof.sh")
	if !strings.Contains(out, "Updated "+draftID) {
		t.Fatalf("output:\n%s", out)
	}
	if got := strings.Join(specSummary(t, root, draftID), ","); got != "proof.sh|sh proof.sh" {
		t.Fatalf("specs %q", got)
	}
	for name, args := range map[string][]string{
		"duplicate":    {"edit", draftID, "--add-spec", "proof.sh"},
		"missing file": {"edit", draftID, "--add-spec", "nope.sh"},
		"empty path":   {"edit", draftID, "--add-spec", ":sh"},
		"unknown":      {"edit", draftID, "--remove-spec", "nope.sh"},
	} {
		if _, err := run(t, append([]string{"-c", cfg}, args...)...); err == nil {
			t.Errorf("%s must be an error", name)
		}
	}
	if got := strings.Join(specSummary(t, root, draftID), ","); got != "proof.sh|sh proof.sh" {
		t.Fatalf("a refused edit changed the record: %q", got)
	}

	// Repeatable, and add plus remove in one call.
	mustRun(t, "-c", cfg, "edit", draftID, "--add-spec", "other.sh", "--add-spec", "third.sh:sh", "--remove-spec", "proof.sh")
	if got := strings.Join(specSummary(t, root, draftID), ","); got != "other.sh|,third.sh|sh" {
		t.Fatalf("specs %q", got)
	}

	// On an open record a change is a new version and needs a reason.
	mustRun(t, "-c", cfg, "open", "--record", draftID)
	if _, err := run(t, "-c", cfg, "edit", draftID, "--remove-spec", "third.sh"); err == nil {
		t.Fatal("removing a spec of an open record without --reason must fail")
	}
	out = mustRun(t, "-c", cfg, "edit", draftID, "--remove-spec", "third.sh", "--add-spec", "proof.sh:sh", "--reason", "swap the proof")
	if !strings.Contains(out, draftID+" is now version 2") || strings.Join(specSummary(t, root, draftID), ",") != "other.sh|,proof.sh|sh" {
		t.Fatalf("open record:\n%s\n%v", out, specSummary(t, root, draftID))
	}
}
