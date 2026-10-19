package main

import (
	"path/filepath"
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

const secondID = "prov-2026-0000f002"

// addRecord adds a second open record, a copy of the first with another id.
func addRecord(t *testing.T, root, cfg, id string) {
	t.Helper()
	rec := strings.Replace(v2Draft(id, "draft"), "affected_scope:\n  - a.go\n", "affected_scope:\n  - a.go\n  - spec.sh\n", 1)
	rec = strings.Replace(rec, "associated_specs: []\n", "associated_specs:\n  - path: spec.sh\n    run_command: sh\n", 1)
	writeFile(t, filepath.Join(root, "provenance", id+".yml"), rec)
	mustRun(t, "-c", cfg, "open", "--record", id)
	gitIn(t, root, "add", "-A")
	gitIn(t, root, "commit", "-q", "-m", "second record")
}

func TestPublishPlanWritesNothing(t *testing.T) {
	_, srv := remoteEnv(t)
	_, cfg := checkRepo(t)
	signedIn(t)
	out, _, err := runCLI(t, "", "-c", cfg, "publish", "--all", "--plan")
	if err != nil || srv.Writes != 0 || srv.Versions(chkID) != 0 {
		t.Fatalf("plan: %v writes=%d\n%s", err, srv.Writes, out)
	}
	if !strings.Contains(out, "unpublished") || !strings.Contains(out, chkID) || !strings.Contains(out, "1 unpublished") {
		t.Fatalf("the plan must list the record as unpublished:\n%s", out)
	}
	if _, _, err := runCLI(t, "", "-c", cfg, "publish", "--all", "--yes"); err != nil {
		t.Fatal(err)
	}
	writes := srv.Writes
	out, _, err = runCLI(t, "", "-c", cfg, "publish", "--all", "--plan")
	if err != nil || srv.Writes != writes || !strings.Contains(out, "up to date") || !strings.Contains(out, "1 up to date") {
		t.Fatalf("plan after publishing: %v\n%s", err, out)
	}
}

func TestPublishSkipsWhatDivergedAndPublishesTheRest(t *testing.T) {
	_, srv := remoteEnv(t)
	signedIn(t)
	// Another clone put version 2 of the first record on the server.
	_, cfgB := checkRepo(t)
	mustRun(t, "-c", cfgB, "edit", chkID, "--add-constraint", "Elsewhere", "--reason", "Another clone")
	if _, _, err := runCLI(t, "", "-c", cfgB, "publish", chkID, "--yes"); err != nil {
		t.Fatal(err)
	}
	// This clone has the first record at version 1 only, and a new second one.
	rootA, cfgA := checkRepo(t)
	addRecord(t, rootA, cfgA, secondID)

	out, _, err := runCLI(t, "", "-c", cfgA, "publish", "--all", "--plan")
	if err != nil || !strings.Contains(out, "server ahead") || !strings.Contains(out, "unpublished") || !strings.Contains(out, "1 server ahead") {
		t.Fatalf("plan: %v\n%s", err, out)
	}
	before := srv.Writes
	out, _, err = runCLI(t, "", "-c", cfgA, "publish", "--all", "--yes")
	if err == nil || !strings.Contains(err.Error(), "skipped") {
		t.Fatalf("a skipped record must end in an error: %v\n%s", err, out)
	}
	if srv.Versions(secondID) != 1 || srv.Writes != before+1 || srv.Versions(chkID) != 2 {
		t.Fatalf("the rest must still be published: second=%d writes %d→%d first=%d\n%s", srv.Versions(secondID), before, srv.Writes, srv.Versions(chkID), out)
	}
	if !strings.Contains(out, "skipped  "+chkID) || !strings.Contains(out, "run provenance import") {
		t.Fatalf("the skipped record and the way out must be named:\n%s", out)
	}
}

func TestPublishSendsTheSource(t *testing.T) {
	_, srv := remoteEnv(t)
	root, cfg := checkRepo(t)
	signedIn(t)
	out, _, err := runCLI(t, "", "-c", cfg, "publish", chkID, "--yes")
	if err != nil || srv.Source(chkID) != nil || !strings.Contains(out, "no usable origin remote") {
		t.Fatalf("without an origin: %v source=%+v\n%s", err, srv.Source(chkID), out)
	}
	mustRun(t, "-c", cfg, "edit", chkID, "--add-constraint", "Second", "--reason", "Tighten it")
	gitIn(t, root, "add", "-A")
	gitIn(t, root, "commit", "-q", "-m", "second version")
	gitIn(t, root, "remote", "add", "origin", "git@github.com:acme/widgets.git")
	writeFile(t, filepath.Join(root, "a.go"), "package a // dirty\n")
	out, _, err = runCLI(t, "", "-c", cfg, "publish", chkID, "--yes")
	if err != nil {
		t.Fatal(err, out)
	}
	got := srv.Source(chkID)
	if got == nil || got.Repo != "https://github.com/acme/widgets" || got.Path != "provenance/"+chkID+".yml" || got.Commit != gitIn(t, root, "rev-parse", "HEAD") {
		t.Fatalf("source = %+v", got)
	}
	if !strings.Contains(out, "uncommitted") {
		t.Fatalf("a dirty tree must be mentioned:\n%s", out)
	}
}

func TestRepoURL(t *testing.T) {
	for in, want := range map[string]string{
		"git@github.com:acme/widgets.git":               "https://github.com/acme/widgets",
		"https://github.com/acme/widgets.git":           "https://github.com/acme/widgets",
		"https://token@github.com/acme/widgets":         "https://github.com/acme/widgets",
		"ssh://git@gitlab.com/group/sub/proj.git":       "https://gitlab.com/group/sub/proj",
		"ssh://git@host.example:2222/acme/widgets.git/": "https://host.example/acme/widgets",
		"/home/me/repos/widgets":                        "",
		"../widgets":                                    "",
		"":                                              "",
		"https://github.com":                            "",
	} {
		if got := repoURL(in); got != want {
			t.Errorf("repoURL(%q) = %q, want %q", in, got, want)
		}
	}
}
