package main

import (
	"encoding/json"
	"path/filepath"
	"strings"
	"testing"
)

// historyRepo: chkID at version 2 (edit, fix), with commits tagged @1, @2
// and one bare tag.
func historyRepo(t *testing.T) (root, cfg string) {
	t.Helper()
	root, cfg = checkRepo(t)
	gitIn(t, root, "commit", "-q", "--allow-empty", "-m", "Work on it ["+chkID+"@1]")
	gitIn(t, root, "commit", "-q", "--allow-empty", "-m", "Bare tag ["+chkID+"]")
	mustRun(t, "-c", cfg, "edit", chkID, "--add-constraint", "Second", "--reason", "Tighten it", "--fix")
	gitIn(t, root, "add", "-A")
	gitIn(t, root, "commit", "-q", "-m", "Revise ["+chkID+"@2]")
	return root, cfg
}

func TestHistory(t *testing.T) {
	_, cfg := historyRepo(t)
	out := mustRun(t, "-c", cfg, "history", chkID)
	for _, want := range []string{
		chkID, "v1", "v2", "Tighten it", "[fix]", "test@example.com",
		"Work on it [" + chkID + "@1]", "Revise [" + chkID + "@2]", "Bare tag [" + chkID + "]",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("history lacks %q:\n%s", want, out)
		}
	}
	// Version 2 comes after version 1, and the bare tag is listed apart.
	if strings.Index(out, "v1") > strings.Index(out, "v2") || strings.Index(out, "Revise") < strings.Index(out, "v2") {
		t.Fatalf("order:\n%s", out)
	}
}

func TestHistoryJSON(t *testing.T) {
	_, cfg := historyRepo(t)
	var h struct {
		ID       string `json:"id"`
		Versions []struct {
			Version     int      `json:"version"`
			Reason      string   `json:"reason"`
			ReasonKind  string   `json:"reason_kind"`
			ContentHash string   `json:"content_hash"`
			Commits     []string `json:"commits"`
		} `json:"versions"`
		Untagged []string `json:"untagged_commits"`
	}
	if err := json.Unmarshal([]byte(mustRun(t, "-c", cfg, "history", "--record", chkID, "--json")), &h); err != nil {
		t.Fatal(err)
	}
	if h.ID != chkID || len(h.Versions) != 2 || h.Versions[1].Reason != "Tighten it" || h.Versions[1].ReasonKind != "fix" || !strings.HasPrefix(h.Versions[1].ContentHash, "sha256:") {
		t.Fatalf("history json %+v", h)
	}
	if len(h.Versions[1].Commits) != 1 || len(h.Untagged) != 1 {
		t.Fatalf("commits v2 %v untagged %v", h.Versions[1].Commits, h.Untagged)
	}
}

func TestHistoryDraft(t *testing.T) {
	_, cfg := v2Repo(t, false, map[string]string{draftID: v2Draft(draftID, "draft")})
	if out := mustRun(t, "-c", cfg, "history", draftID); !strings.Contains(out, "no versions yet") {
		t.Fatalf("draft history:\n%s", out)
	}
}

func TestDiff(t *testing.T) {
	root, cfg := historyRepo(t)
	out := mustRun(t, "-c", cfg, "diff", chkID)
	for _, want := range []string{"v1 → v2", "constraints:", "+ Second"} {
		if !strings.Contains(out, want) {
			t.Errorf("default diff lacks %q:\n%s", want, out)
		}
	}
	if out := mustRun(t, "-c", cfg, "diff", chkID, "1", "1"); !strings.Contains(out, "no differences") {
		t.Fatalf("v1 vs v1:\n%s", out)
	}
	if out := mustRun(t, "-c", cfg, "diff", "--record", chkID, "2", "1"); !strings.Contains(out, "- Second") {
		t.Fatalf("v2 → v1:\n%s", out)
	}
	// An unbumped edit is compared against the working copy.
	path := filepath.Join(root, "provenance", chkID+".yml")
	writeFile(t, path, strings.Replace(readFile(t, path), "intent: Do the thing\n", "intent: Do it differently\n", 1))
	out = mustRun(t, "-c", cfg, "diff", chkID)
	for _, want := range []string{"v2 → working copy", "intent:", "- Do the thing", "+ Do it differently"} {
		if !strings.Contains(out, want) {
			t.Errorf("working-copy diff lacks %q:\n%s", want, out)
		}
	}
	for name, args := range map[string][]string{
		"one version":     {"diff", chkID, "1"},
		"missing version": {"diff", chkID, "1", "9"},
		"not a number":    {"diff", chkID, "a", "b"},
	} {
		if out, err := run(t, append([]string{"-c", cfg}, args...)...); err == nil {
			t.Errorf("%s: accepted\n%s", name, out)
		}
	}
}
