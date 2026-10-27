package main

import (
	"path/filepath"
	"strings"
	"testing"
)

// refreshRepo is a completed record with its evidence on file.
func refreshRepo(t *testing.T) (root, cfg string) {
	t.Helper()
	root, cfg = completeRepo(t, "run_associated_specs_on_complete: true\n", "[ -z \"$REFRESH_FAIL\" ]\n")
	mustRun(t, "-c", cfg, "complete", "--record", chkID)
	return root, cfg
}

func evidenceRuns(t *testing.T, root string) int {
	t.Helper()
	runs, _ := filepath.Glob(filepath.Join(root, "provenance", ".evidence", chkID, "*"))
	return len(runs)
}

func TestEvidenceRefreshKeepsTheVersion(t *testing.T) {
	root, cfg := refreshRepo(t)
	if out := mustRun(t, "-c", cfg, "check", "--evidence"); !strings.Contains(out, "fresh") {
		t.Fatalf("fresh after complete:\n%s", out)
	}
	before := recordFile(t, root, chkID)
	out := mustRun(t, "-c", cfg, "evidence", "refresh", chkID, "--reason", "specs re-run")
	if !strings.Contains(out, "Refreshed evidence of "+chkID+"@1") || !strings.Contains(out, "No new version") {
		t.Fatalf("output:\n%s", out)
	}
	if recordFile(t, root, chkID) != before || len(versions(t, root, chkID)) != 1 {
		t.Fatal("refresh changed the record or wrote a version")
	}
	if evidenceRuns(t, root) != 2 {
		t.Fatalf("evidence files: %d", evidenceRuns(t, root))
	}
	if out := mustRun(t, "-c", cfg, "check", "--evidence"); !strings.Contains(out, "fresh") {
		t.Fatalf("fresh again after refresh:\n%s", out)
	}
}

func TestEvidenceRefreshRefusals(t *testing.T) {
	root, cfg := refreshRepo(t)

	// A failing spec writes nothing.
	t.Setenv("REFRESH_FAIL", "1")
	if _, err := run(t, "-c", cfg, "evidence", "refresh", chkID); err == nil {
		t.Fatal("a failing spec must fail the refresh")
	}
	if evidenceRuns(t, root) != 1 {
		t.Fatalf("a failed refresh wrote evidence: %d", evidenceRuns(t, root))
	}
	t.Setenv("REFRESH_FAIL", "")

	// A moved spec hash is a real change: refused, naming specs.
	writeFile(t, filepath.Join(root, "spec.sh"), "[ -z \"$REFRESH_FAIL\" ] # changed\n")
	out, err := run(t, "-c", cfg, "evidence", "refresh", chkID)
	if err == nil || evidenceRuns(t, root) != 1 {
		t.Fatalf("a spec change must be refused (err %v, runs %d):\n%s", err, evidenceRuns(t, root), out)
	}
	if _, err = run(t, "-c", cfg, "evidence", "refresh", chkID); err == nil {
		t.Fatal("still refused")
	}
}

func TestEvidenceRefreshNotImplemented(t *testing.T) {
	_, cfg := checkRepo(t)
	if _, err := run(t, "-c", cfg, "evidence", "refresh", chkID); err == nil {
		t.Fatal("an open record must be refused")
	}
}
