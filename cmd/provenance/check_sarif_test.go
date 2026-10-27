package main

import (
	"encoding/json"
	"path/filepath"
	"strings"
	"testing"
)

func sarifResults(t *testing.T, out string) []sarifResult {
	t.Helper()
	if err := validateSarif(t, out); err != nil {
		t.Fatalf("not valid SARIF 2.1.0: %v\n%s", err, out)
	}
	var doc sarifDoc
	if err := json.Unmarshal([]byte(out), &doc); err != nil {
		t.Fatal(err)
	}
	return doc.Runs[0].Results
}

func TestCheckStagedSarif(t *testing.T) {
	root, cfg := checkRepo(t)
	msg := filepath.Join(root, ".git", "TEST_MSG")

	stage(t, root, map[string]string{"a.go": "package a // ok\n"})
	writeFile(t, msg, "Change a ["+chkID+"]\n")
	out, err := run(t, "-c", cfg, "check", "--staged", "--message-file", msg, "--format", "sarif")
	if err != nil || len(sarifResults(t, out)) != 0 {
		t.Fatalf("a clean change gives an empty, valid document and exit 0 (err %v):\n%s", err, out)
	}

	stage(t, root, map[string]string{"b.go": "package b // out of scope\n"})
	writeFile(t, msg, "Change b ["+chkID+"]\n")
	out, err = run(t, "-c", cfg, "check", "--staged", "--message-file", msg, "--format", "sarif")
	if err == nil {
		t.Fatal("violations must still fail the command")
	}
	var found bool
	for _, res := range sarifResults(t, out) {
		if res.RuleID == "SCOPE001" && res.Locations[0].PhysicalLocation.ArtifactLocation.URI == "b.go" && res.Level == "error" && res.Properties["recordId"] == chkID {
			found = true
		}
	}
	if !found {
		t.Fatalf("no SCOPE001 result for b.go:\n%s", out)
	}

	writeFile(t, msg, "No tag at all\n")
	out, _ = run(t, "-c", cfg, "check", "--staged", "--message-file", msg, "--format", "sarif")
	var tag bool
	for _, res := range sarifResults(t, out) {
		tag = tag || res.RuleID == "TAG001"
	}
	if !tag {
		t.Fatalf("a missing tag is TAG001:\n%s", out)
	}
	if _, err := run(t, "-c", cfg, "check", "--staged", "--format", "xml"); err == nil {
		t.Error("an unknown --format was accepted")
	}
}

func TestCheckCommitSarif(t *testing.T) {
	root, cfg := checkRepo(t)
	base := gitIn(t, root, "rev-parse", "HEAD")
	stage(t, root, map[string]string{"b.go": "package b // two\n"})
	gitIn(t, root, "commit", "-q", "-m", "Change b ["+chkID+"]")
	bad := gitIn(t, root, "rev-parse", "--short=7", "HEAD")
	out, err := run(t, "-c", cfg, "check", "--range", base+"..HEAD", "--format", "sarif")
	if err == nil {
		t.Fatal("a violating commit must fail")
	}
	res := sarifResults(t, out)
	if len(res) == 0 || res[0].RuleID != "SCOPE001" || !strings.HasPrefix(res[0].Message.Text, bad+": ") {
		t.Fatalf("results %+v", res)
	}
}

func TestCheckEvidenceSarif(t *testing.T) {
	root, cfg := refreshRepo(t)
	if out, err := run(t, "-c", cfg, "check", "--evidence", "--format", "sarif"); err != nil || len(sarifResults(t, out)) != 0 {
		t.Fatalf("fresh evidence: no findings (err %v):\n%s", err, out)
	}
	writeFile(t, filepath.Join(root, "spec.sh"), "[ -z \"$REFRESH_FAIL\" ] # changed\n")
	out, err := run(t, "-c", cfg, "check", "--evidence", "--format", "sarif")
	if err == nil {
		t.Fatal("stale evidence must fail")
	}
	res := sarifResults(t, out)
	if len(res) != 1 || res[0].RuleID != "EVID001" || !strings.Contains(res[0].Message.Text, "specs") {
		t.Fatalf("results %+v", res)
	}
}
