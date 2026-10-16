package main

import (
	"strings"
	"testing"
)

func TestPublishNeedsSessionAndConfirmation(t *testing.T) {
	_, srv := remoteEnv(t)
	_, cfg := checkRepo(t)
	_, _, err := runCLI(t, "", "-c", cfg, "publish", chkID, "--yes")
	if err == nil || exitCode(err) != 3 {
		t.Fatalf("logged out publish: err=%v code=%d", err, exitCode(err))
	}
	signedIn(t)
	out, _, err := runCLI(t, "", "-c", cfg, "publish", chkID)
	if err == nil || !strings.Contains(err.Error(), "not confirmed") || srv.Writes != 0 {
		t.Fatalf("publish without --yes on a non-terminal: err=%v writes=%d\n%s", err, srv.Writes, out)
	}
	if !strings.Contains(out, "PUBLIC") || !strings.Contains(out, chkID) {
		t.Fatalf("the plan must say what becomes public:\n%s", out)
	}
	if _, _, err := runCLI(t, "", "-c", cfg, "publish"); err == nil {
		t.Fatal("publish must name records or --all")
	}
	if _, _, err := runCLI(t, "", "-c", cfg, "publish", chkID, "--yes", "--project", "not-a-uri"); err == nil {
		t.Fatal("--project must be an at:// uri")
	}
}

func TestPublishTwiceWritesNothingSecondTime(t *testing.T) {
	_, srv := remoteEnv(t)
	root, cfg := checkRepo(t)
	signedIn(t)
	out, _, err := runCLI(t, "", "-c", cfg, "publish", "--all", "--yes")
	if err != nil || srv.Versions(chkID) != 1 || !strings.Contains(out, "published  "+chkID+" v1") {
		t.Fatalf("publish: %v versions=%d\n%s", err, srv.Versions(chkID), out)
	}
	writes := srv.Writes
	out, _, err = runCLI(t, "", "-c", cfg, "publish", "--all", "--yes")
	if err != nil || srv.Writes != writes || !strings.Contains(out, "Nothing to publish") {
		t.Fatalf("second publish: %v writes %d→%d\n%s", err, writes, srv.Writes, out)
	}
	mustRun(t, "-c", cfg, "edit", chkID, "--add-constraint", "Second", "--reason", "Tighten it")
	_ = root
	out, _, err = runCLI(t, "", "-c", cfg, "publish", chkID, "--yes")
	if err != nil || srv.Versions(chkID) != 2 || srv.Writes != writes+1 {
		t.Fatalf("third publish: %v versions=%d\n%s", err, srv.Versions(chkID), out)
	}
}
