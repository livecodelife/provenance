package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/livecodelife/provenance/record"
)

// openedRepo is a migrated repo whose record id has been opened (version 1
// snapshot written), not under git.
func openedRepo(t *testing.T, id string) (root, cfg string) {
	t.Helper()
	root, cfg = v2Repo(t, false, map[string]string{id: v2Draft(id, "draft")})
	mustRun(t, "-c", cfg, "open", "--record", id)
	return root, cfg
}

// editorScript sets $EDITOR to a script that runs sed with expr on the file.
func editorScript(t *testing.T, dir, expr string) {
	t.Helper()
	p := filepath.Join(dir, "editor-"+strings.NewReplacer("/", "_", " ", "_").Replace(expr)+".sh")
	writeFile(t, p, "#!/bin/sh\nsed -i.bak '"+expr+"' \"$1\" && rm -f \"$1.bak\"\n")
	if err := os.Chmod(p, 0o755); err != nil {
		t.Fatal(err)
	}
	t.Setenv("EDITOR", p)
}

func load(t *testing.T, root, id string) *record.Record {
	t.Helper()
	rec, _, err := (&record.FilesStore{Root: root}).Load(id)
	if err != nil {
		t.Fatal(err)
	}
	return rec
}

func versions(t *testing.T, root, id string) []int {
	t.Helper()
	vs, err := (&record.FilesStore{Root: root}).Versions(id)
	if err != nil {
		t.Fatal(err)
	}
	return vs
}

func TestReviseBumpsVersion(t *testing.T) {
	root, cfg := openedRepo(t, draftID)
	editorScript(t, root, "s/^intent: Do the thing$/intent: Do the other thing/")
	out := mustRun(t, "-c", cfg, "revise", draftID, "--reason", "Scope changed", "--fix")
	if !strings.Contains(out, draftID+" is now version 2") || !strings.Contains(out, "["+draftID+"@2]") {
		t.Fatalf("output:\n%s", out)
	}
	rec := load(t, root, draftID)
	if rec.Version != 2 || rec.Reason != "Scope changed" || rec.ReasonKind != record.ReasonKindFix || rec.Intent != "Do the other thing" {
		t.Fatalf("record %+v", rec)
	}
	snap, err := (&record.FilesStore{Root: root}).Snapshot(draftID, 2)
	if err != nil || snap.Reason != "Scope changed" || snap.Content.Intent != "Do the other thing" {
		t.Fatalf("snapshot %+v (%v)", snap, err)
	}
	if !strings.Contains(readFile(t, filepath.Join(root, manifestPath)), `"version":2`) {
		t.Fatalf("manifest not rewritten:\n%s", readFile(t, filepath.Join(root, manifestPath)))
	}
}

func TestReviseUnchangedWritesNothing(t *testing.T) {
	root, cfg := openedRepo(t, draftID)
	t.Setenv("EDITOR", "true")
	before := readFile(t, filepath.Join(root, "provenance", draftID+".yml"))
	out := mustRun(t, "-c", cfg, "revise", "--record", draftID, "--reason", "nothing")
	if !strings.Contains(out, "No change") {
		t.Fatalf("output:\n%s", out)
	}
	if got := readFile(t, filepath.Join(root, "provenance", draftID+".yml")); got != before || len(versions(t, root, draftID)) != 1 {
		t.Fatalf("unchanged revise wrote something")
	}
}

func TestReviseNeedsReason(t *testing.T) {
	root, cfg := openedRepo(t, draftID)
	editorScript(t, root, "s/Keep it small/Keep it tiny/")
	before := readFile(t, filepath.Join(root, "provenance", draftID+".yml"))
	if out, err := run(t, "-c", cfg, "revise", draftID); err == nil {
		t.Fatalf("content change without --reason accepted\n%s", out)
	}
	if got := readFile(t, filepath.Join(root, "provenance", draftID+".yml")); got != before || len(versions(t, root, draftID)) != 1 {
		t.Fatalf("refused revise wrote something")
	}
}

func TestReviseReopensImplemented(t *testing.T) {
	root, cfg := openedRepo(t, draftID)
	path := filepath.Join(root, "provenance", draftID+".yml")
	writeFile(t, path, strings.Replace(strings.Replace(readFile(t, path), "status: open", "status: implemented", 1), `sealed_at_sha: ""`, "sealed_at_sha: abc1234", 1))
	editorScript(t, root, "s/Do the thing/Do it better/")
	out := mustRun(t, "-c", cfg, "revise", draftID, "--reason", "Better")
	rec := load(t, root, draftID)
	if rec.Status != record.StatusOpen || rec.SealedAtSHA != "" || rec.Version != 2 || !strings.Contains(out, "reopened") {
		t.Fatalf("status %s sealed %q version %d\n%s", rec.Status, rec.SealedAtSHA, rec.Version, out)
	}
}

func TestReviseDraftInPlace(t *testing.T) {
	root, cfg := v2Repo(t, false, map[string]string{draftID: v2Draft(draftID, "draft")})
	editorScript(t, root, "s/Do the thing/Do a different thing/")
	mustRun(t, "-c", cfg, "revise", draftID)
	rec := load(t, root, draftID)
	if rec.Intent != "Do a different thing" || rec.Version != 1 || rec.Status != record.StatusDraft || len(versions(t, root, draftID)) != 0 {
		t.Fatalf("draft revise: %+v versions %v", rec, versions(t, root, draftID))
	}
}

func TestReviseTitleOnlyNoBump(t *testing.T) {
	root, cfg := openedRepo(t, draftID)
	editorScript(t, root, "s/^title: .*/title: Renamed/")
	mustRun(t, "-c", cfg, "revise", draftID)
	rec := load(t, root, draftID)
	if rec.Title != "Renamed" || rec.Version != 1 || len(versions(t, root, draftID)) != 1 {
		t.Fatalf("title-only revise: title %q version %d", rec.Title, rec.Version)
	}
}

func TestReviseRefusals(t *testing.T) {
	dep := "prov-2026-0000d004"
	root, cfg := openedRepo(t, draftID)
	writeFile(t, filepath.Join(root, "provenance", dep+".yml"), v2Draft(dep, "deprecated"))
	editorScript(t, root, "s/^status: open$/status: implemented/")
	before := readFile(t, filepath.Join(root, "provenance", draftID+".yml"))
	for name, args := range map[string][]string{
		"no id":            {"revise", "--reason", "x"},
		"mismatched ids":   {"revise", draftID, "--record", dep, "--reason", "x"},
		"deprecated":       {"revise", dep, "--reason", "x"},
		"lifecycle edited": {"revise", draftID, "--reason", "x"},
		"two args":         {"revise", draftID, dep},
	} {
		if out, err := run(t, append([]string{"-c", cfg}, args...)...); err == nil {
			t.Errorf("%s: accepted\n%s", name, out)
		}
	}
	if got := readFile(t, filepath.Join(root, "provenance", draftID+".yml")); got != before || len(versions(t, root, draftID)) != 1 {
		t.Fatalf("refused revise wrote something")
	}
}

func TestReviseResolvesAlias(t *testing.T) {
	root, cfg := openedRepo(t, draftID)
	path := filepath.Join(root, "provenance", draftID+".yml")
	writeFile(t, path, strings.Replace(readFile(t, path), "author: a@b\n", "author: a@b\naliases:\n  - prov-2026-0000a11a\n", 1))
	editorScript(t, root, "s/Do the thing/Via alias/")
	mustRun(t, "-c", cfg, "revise", "prov-2026-0000a11a", "--reason", "alias")
	if rec := load(t, root, draftID); rec.Version != 2 || rec.Intent != "Via alias" {
		t.Fatalf("alias revise: %+v", rec)
	}
}

// specRecord is v2Draft with one associated spec, spec.sh.
func specRecord(id string) string {
	return strings.Replace(v2Draft(id, "draft"), "associated_specs: []",
		"associated_specs:\n  - path: spec.sh\n    run_command: sh spec.sh", 1)
}

func TestReviseSpecOnlyChange(t *testing.T) {
	root, cfg := v2Repo(t, false, map[string]string{draftID: specRecord(draftID)})
	writeFile(t, filepath.Join(root, "spec.sh"), "exit 0\n")
	mustRun(t, "-c", cfg, "open", "--record", draftID)
	t.Setenv("EDITOR", "true")

	out := mustRun(t, "-c", cfg, "revise", draftID, "--reason", "nothing")
	if !strings.Contains(out, "No change") || len(versions(t, root, draftID)) != 1 {
		t.Fatalf("untouched record:\n%s", out)
	}

	// The record file is byte-identical, only the spec content moved.
	writeFile(t, filepath.Join(root, "spec.sh"), "exit 0 # changed\n")
	if _, err := run(t, "-c", cfg, "revise", draftID); err == nil {
		t.Fatal("a spec-only change without --reason must fail")
	}
	out = mustRun(t, "-c", cfg, "revise", draftID, "--reason", "spec changed on disk")
	if !strings.Contains(out, draftID+" is now version 2") || len(versions(t, root, draftID)) != 2 {
		t.Fatalf("spec-only change:\n%s", out)
	}
}

func TestEditRefreshSpecs(t *testing.T) {
	root, cfg := v2Repo(t, false, map[string]string{draftID: specRecord(draftID)})
	writeFile(t, filepath.Join(root, "spec.sh"), "exit 0\n")
	mustRun(t, "-c", cfg, "open", "--record", draftID)

	out := mustRun(t, "-c", cfg, "edit", draftID, "--refresh-specs", "--reason", "none")
	if !strings.Contains(out, "No change") || len(versions(t, root, draftID)) != 1 {
		t.Fatalf("no spec change:\n%s", out)
	}
	writeFile(t, filepath.Join(root, "spec.sh"), "exit 0 # changed\n")
	if _, err := run(t, "-c", cfg, "edit", draftID, "--refresh-specs"); err == nil {
		t.Fatal("--refresh-specs without --reason must fail")
	}
	out = mustRun(t, "-c", cfg, "edit", draftID, "--refresh-specs", "--reason", "spec changed on disk")
	if !strings.Contains(out, draftID+" is now version 2") || len(versions(t, root, draftID)) != 2 {
		t.Fatalf("refresh:\n%s", out)
	}
}

func TestOpenRefusesSelfRewrittenSpec(t *testing.T) {
	for _, spec := range []string{manifestPath, "provenance/" + draftID + ".yml", "provenance/.versions/" + draftID + "/1-x.yml"} {
		rec := strings.Replace(specRecord(draftID), "path: spec.sh", "path: "+spec, 1)
		root, cfg := v2Repo(t, false, map[string]string{draftID: rec})
		if _, err := run(t, "-c", cfg, "open", "--record", draftID); err == nil {
			t.Fatalf("open must refuse the spec %s", spec)
		}
		if got := versions(t, root, draftID); len(got) != 0 {
			t.Fatalf("%s: versions written: %v", spec, got)
		}
	}
}
