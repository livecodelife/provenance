package main

import (
	"strings"
	"testing"
)

// The M4 check in one run: sign in, publish a record, publish again (nothing
// is written), clone it elsewhere (same content hashes), and bring a UI edit
// back as a new local version. No token reaches any output on the way.
func TestM4Flow(t *testing.T) {
	f, srv := remoteEnv(t)
	root, cfg := checkRepo(t)
	var seen []string
	run := func(stdin string, args ...string) string {
		out, errOut, err := runCLI(t, stdin, args...)
		seen = append(seen, out, errOut)
		if err != nil {
			t.Fatalf("provenance %s: %v\n%s", strings.Join(args, " "), err, out)
		}
		return out
	}
	run("", "login", "--no-browser")
	mustRun(t, "-c", cfg, "edit", chkID, "--add-constraint", "Second", "--reason", "Tighten it")
	run("", "-c", cfg, "publish", chkID, "--yes")
	writes := srv.Writes
	run("", "-c", cfg, "publish", chkID, "--yes")
	if srv.Writes != writes {
		t.Fatalf("a second publish wrote %d time(s)", srv.Writes-writes)
	}
	dir := t.TempDir() + "/clone"
	run("", "clone", dir)
	if !strings.Contains(manifestLines(t, root), strings.TrimSpace(manifestLines(t, dir))) {
		t.Fatalf("clone manifest differs:\n%s\n%s", manifestLines(t, dir), manifestLines(t, root))
	}
	uiEdit(t, root, chkID, "Changed in the UI.")
	run("", "-c", cfg, "import", "--all")
	if !strings.Contains(mustRun(t, "-c", cfg, "history", chkID), "edited in the UI") {
		t.Fatal("the UI version is not in the local history")
	}
	all := strings.Join(seen, "\n")
	for _, s := range f.Secrets {
		if strings.Contains(all, s) {
			t.Errorf("secret %q reached the output", s)
		}
	}
}
