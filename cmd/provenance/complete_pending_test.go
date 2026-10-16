package main

import (
	"path/filepath"
	"strings"
	"testing"
)

const pendingChild = "prov-2026-0000c020"

// pendingRepo is completeRepo plus an imprint of chkID in the given status.
func pendingRepo(t *testing.T, childStatus string) (root, cfg string) {
	t.Helper()
	root, cfg = completeRepo(t, "run_associated_specs_on_complete: true\n", "true\n")
	child := strings.Replace(v2Draft(pendingChild, childStatus), "type: blueprint", "type: imprint\nimplements: "+chkID+"\nimplements_version: 1", 1)
	child = strings.Replace(child, "title: Draft "+pendingChild, "title: Unfinished imprint", 1)
	child = strings.Replace(child, "  - a.go\n", "  - b.go\n", 1)
	writeFile(t, filepath.Join(root, "provenance", pendingChild+".yml"), child)
	return root, cfg
}

func TestCompleteWarnsAboutUnfinishedImprints(t *testing.T) {
	_, cfg := pendingRepo(t, "draft")
	out := mustRun(t, "-c", cfg, "complete", "--record", chkID)
	for _, want := range []string{"1 record(s) that implement " + chkID + " are not implemented yet", pendingChild, "draft", "Unfinished imprint"} {
		if !strings.Contains(out, want) {
			t.Errorf("the warning misses %q:\n%s", want, out)
		}
	}
	if !strings.Contains(out, "✓ "+chkID+" marked as implemented") {
		t.Fatalf("the warning must not block the completion:\n%s", out)
	}
}

func TestCompleteDoesNotWarnWithoutPendingImprints(t *testing.T) {
	// A deprecated imprint is not pending.
	_, cfg := pendingRepo(t, "deprecated")
	if out := mustRun(t, "-c", cfg, "complete", "--record", chkID); strings.Contains(out, "not implemented yet") {
		t.Fatalf("a deprecated imprint is not unfinished:\n%s", out)
	}
	// Nothing implements a record with no imprints.
	_, cfg = completeRepo(t, "run_associated_specs_on_complete: true\n", "true\n")
	if out := mustRun(t, "-c", cfg, "complete", "--record", chkID); strings.Contains(out, "not implemented yet") {
		t.Fatalf("no imprints, no warning:\n%s", out)
	}
}
