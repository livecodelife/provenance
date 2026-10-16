package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// refreshRepo is a completed record whose evidence was taken against
// sedum regions {gen: aa}: a fake sedum prints regions.json from its own
// directory, so a test moves the code hash by editing that file.
func refreshRepo(t *testing.T) (root, cfg string) {
	t.Helper()
	root, cfg = completeRepo(t, "run_associated_specs_on_complete: true\n", "[ -z \"$REFRESH_FAIL\" ]\n")
	setRegions(t, root, "aa")
	fake := filepath.Join(root, "fake-sedum")
	writeFile(t, fake, "#!/bin/sh\ncat \"$(dirname \"$0\")/regions.json\"\n")
	if err := os.Chmod(fake, 0o755); err != nil {
		t.Fatal(err)
	}
	t.Setenv("SEDUM_BIN", fake)
	mustRun(t, "-c", cfg, "complete", "--record", chkID)
	return root, cfg
}

func setRegions(t *testing.T, root, gen string) {
	t.Helper()
	writeFile(t, filepath.Join(root, "regions.json"), `{"format":"sedum.regions/1","regions":[{"record":"x","tier":"owned","gen":"sha256:`+gen+`","pkg":"sha256:bb"}]}`)
}

func evidenceRuns(t *testing.T, root string) int {
	t.Helper()
	runs, _ := filepath.Glob(filepath.Join(root, "provenance", ".evidence", chkID, "*"))
	return len(runs)
}

func TestEvidenceRefreshAfterCodeMoved(t *testing.T) {
	root, cfg := refreshRepo(t)
	if out := mustRun(t, "-c", cfg, "check", "--evidence"); !strings.Contains(out, "fresh") {
		t.Fatalf("fresh after complete:\n%s", out)
	}
	before := recordFile(t, root, chkID)
	setRegions(t, root, "cc")
	if out, err := run(t, "-c", cfg, "check", "--evidence"); err == nil || !strings.Contains(out, "stale") || !strings.Contains(out, "code") {
		t.Fatalf("a moved code hash must be stale (err %v):\n%s", err, out)
	}
	out := mustRun(t, "-c", cfg, "evidence", "refresh", chkID, "--reason", "template change plus replay")
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
	setRegions(t, root, "cc")
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

func TestEvidenceRefreshNotImplementedAndAll(t *testing.T) {
	root, cfg := checkRepo(t)
	if _, err := run(t, "-c", cfg, "evidence", "refresh", chkID); err == nil {
		t.Fatal("an open record must be refused")
	}
	_ = root

	root, cfg = refreshRepo(t)
	if out := mustRun(t, "-c", cfg, "evidence", "refresh", "--all"); !strings.Contains(out, "Nothing to refresh") {
		t.Fatalf("--all with all fresh:\n%s", out)
	}
	setRegions(t, root, "dd")
	if out := mustRun(t, "-c", cfg, "evidence", "refresh", "--all"); !strings.Contains(out, "Refreshed evidence of "+chkID) {
		t.Fatalf("--all refreshes code-only stale records:\n%s", out)
	}
	if evidenceRuns(t, root) != 2 {
		t.Fatalf("evidence files: %d", evidenceRuns(t, root))
	}
}
