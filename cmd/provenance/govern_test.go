package main

import (
	"encoding/json"
	"strings"
	"testing"
)

func governDocOf(t *testing.T, cfg string, args ...string) governDoc {
	t.Helper()
	out := mustRun(t, append([]string{"-c", cfg, "govern", "--json"}, args...)...)
	var doc governDoc
	if err := json.Unmarshal([]byte(out), &doc); err != nil {
		t.Fatalf("json: %v\n%s", err, out)
	}
	return doc
}

func TestGovernOpenRecordAndExempt(t *testing.T) {
	_, cfg := checkRepo(t)
	doc := governDocOf(t, cfg, "a.go", "b.go", "README.md")
	if len(doc.Governing) != 1 || doc.Governing[0].ID != chkID || doc.Governing[0].Status != "open" || strings.Join(doc.Governing[0].Files, ",") != "a.go" {
		t.Fatalf("only a.go is covered: %+v", doc.Governing)
	}
	if len(doc.Exempt) != 1 || doc.Exempt[0] != "README.md" || len(doc.Files) != 3 {
		t.Fatalf("exempt %v files %v", doc.Exempt, doc.Files)
	}
	var nx struct {
		Actions []nextAction `json:"actions"`
	}
	if err := json.Unmarshal([]byte(mustRun(t, "-c", cfg, "next", "--json", "a.go", "b.go", "README.md")), &nx); err != nil {
		t.Fatal(err)
	}
	if doc.Next == nil || doc.Next.Kind != nx.Actions[0].Kind || doc.Next.Command != nx.Actions[0].Command || doc.Next.Reason != nx.Actions[0].Reason {
		t.Fatalf("govern next %+v, next --json %+v", doc.Next, nx.Actions[0])
	}
	if d := governDocOf(t, cfg, "b.go"); len(d.Governing) != 0 || d.Governing == nil || d.Next == nil || d.Next.Kind != "create" {
		t.Fatalf("an ungoverned file: %+v", d)
	}
	if again := mustRun(t, "-c", cfg, "govern", "a.go", "--format", "json"); again != mustRun(t, "-c", cfg, "govern", "--files", "a.go", "--format", "json") {
		t.Fatal("arguments and --files must give the same answer")
	}
}

func TestGovernActiveOnlyAndForbidden(t *testing.T) {
	// A draft does not govern yet; a deprecated record no longer does.
	_, cfg := v2Repo(t, false, map[string]string{
		draftID:              v2Draft(draftID, "draft"),
		"prov-2026-0000d00a": v2Draft("prov-2026-0000d00a", "deprecated"),
	})
	if d := governDocOf(t, cfg, "a.go"); len(d.Governing) != 0 {
		t.Fatalf("draft and deprecated records are not active: %+v", d.Governing)
	}

	// A record whose forbidden_scope matches the file does not govern it.
	forbidding := strings.Replace(v2Draft("prov-2026-0000d00b", "open"), "forbidden_scope: []", "forbidden_scope:\n  - a.go", 1)
	_, cfg = v2Repo(t, false, map[string]string{"prov-2026-0000d00b": forbidding})
	if d := governDocOf(t, cfg, "a.go"); len(d.Governing) != 0 {
		t.Fatalf("forbidden_scope: %+v", d.Governing)
	}

	// An implemented record governs.
	_, cfg = completeRepo(t, "run_associated_specs_on_complete: true\n", "true\n")
	mustRun(t, "-c", cfg, "complete", "--record", chkID)
	if d := governDocOf(t, cfg, "a.go"); len(d.Governing) != 1 || d.Governing[0].Status != "implemented" || d.Governing[0].Version != 1 {
		t.Fatalf("implemented: %+v", d.Governing)
	}
	out := mustRun(t, "-c", cfg, "govern", "a.go")
	if !strings.Contains(out, "GOVERN: a.go") || !strings.Contains(out, chkID+"@1  implemented") {
		t.Fatalf("human:\n%s", out)
	}
	if _, err := run(t, "-c", cfg, "govern"); err == nil {
		t.Error("no files must be an error")
	}
	if _, err := run(t, "-c", cfg, "govern", "a.go", "--format", "xml"); err == nil {
		t.Error("an unknown format must be an error")
	}
}
