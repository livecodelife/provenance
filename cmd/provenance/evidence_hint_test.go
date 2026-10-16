package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// noRegionsSedum is a sedum that predates the regions command.
func noRegionsSedum(t *testing.T, root string) {
	t.Helper()
	fake := filepath.Join(root, "old-sedum")
	writeFile(t, fake, "#!/bin/sh\necho 'unknown command \"regions\"' >&2\nexit 1\n")
	if err := os.Chmod(fake, 0o755); err != nil {
		t.Fatal(err)
	}
	t.Setenv("SEDUM_BIN", fake)
}

func TestCompleteWarnsAndCheckHintsOnUnknownCodeHash(t *testing.T) {
	root, cfg := completeRepo(t, "run_associated_specs_on_complete: true\n", "true\n")
	noRegionsSedum(t, root)
	out := mustRun(t, "-c", cfg, "complete", "--record", chkID)
	for _, want := range []string{"⚠ The code hash of " + chkID + " is unknown", "may predate the regions command", "provenance evidence refresh " + chkID} {
		if !strings.Contains(out, want) {
			t.Errorf("complete's output misses %q:\n%s", want, out)
		}
	}
	ev := mustRun(t, "-c", cfg, "check", "--evidence")
	if !strings.Contains(ev, "run: provenance evidence refresh "+chkID) || !strings.Contains(ev, "install a current sedum") {
		t.Fatalf("check --evidence hints on the unknown row and at the end:\n%s", ev)
	}
	if js := mustRun(t, "-c", cfg, "check", "--evidence", "--json"); strings.Contains(js, "refresh") {
		t.Fatalf("the JSON output must not change:\n%s", js)
	}

	// A sedum that answers fixes it: refresh, then neither hint appears.
	good := filepath.Join(root, "good-sedum")
	writeFile(t, good, "#!/bin/sh\necho '{\"format\":\"sedum.regions/1\",\"regions\":[]}'\n")
	if err := os.Chmod(good, 0o755); err != nil {
		t.Fatal(err)
	}
	t.Setenv("SEDUM_BIN", good)
	mustRun(t, "-c", cfg, "evidence", "refresh", chkID)
	if ev := mustRun(t, "-c", cfg, "check", "--evidence"); strings.Contains(ev, "refresh") || !strings.Contains(ev, "1 fresh") {
		t.Fatalf("after the refresh:\n%s", ev)
	}
}

func TestNoRegionsRecordWithWorkingSedumGetsNoWarning(t *testing.T) {
	root, cfg := completeRepo(t, "run_associated_specs_on_complete: true\n", "true\n")
	good := filepath.Join(root, "good-sedum")
	writeFile(t, good, "#!/bin/sh\necho '{\"format\":\"sedum.regions/1\",\"regions\":[]}'\n")
	if err := os.Chmod(good, 0o755); err != nil {
		t.Fatal(err)
	}
	t.Setenv("SEDUM_BIN", good)
	if out := mustRun(t, "-c", cfg, "complete", "--record", chkID); strings.Contains(out, "code hash") {
		t.Fatalf("a record with no regions has a defined (empty) code hash:\n%s", out)
	}
	if ev := mustRun(t, "-c", cfg, "check", "--evidence"); strings.Contains(ev, "refresh") {
		t.Fatalf("no hint expected:\n%s", ev)
	}
}
