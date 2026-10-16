package lint

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/livecodelife/provenance/record"
)

// PROV028 (provisio-mgr): an unmigrated repo gets ONE repo-level warning
// saying how many records use the v1 schema; each record's own entry is only
// a hint, so lint output is not flooded.
func TestNotMigratedIsOneWarning(t *testing.T) {
	root := t.TempDir()
	s := &record.FilesStore{Root: root}
	os.MkdirAll(filepath.Join(root, "provenance"), 0o755)
	for _, id := range []string{"prov-2026-0000e001", "prov-2026-0000e002"} {
		body := "id: " + id + "\ntitle: Old\nstatus: draft\ncreated_at: \"2026-01-01\"\nauthor: a\nintent: x\nconstraints: [c]\nsupersedes: \"\"\n"
		if err := os.WriteFile(filepath.Join(root, "provenance", id+".yml"), []byte(body), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	migrated := newRecord("prov-2026-0000a001", record.StatusDraft)
	if err := s.Save(migrated, nil); err != nil {
		t.Fatal(err)
	}

	res := lintStore(t, s)
	var warnings, hints []Issue
	for _, i := range res.Issues {
		if i.Rule != RuleNotMigrated {
			continue
		}
		switch i.Severity {
		case SeverityWarning:
			warnings = append(warnings, i)
		case SeverityHint:
			hints = append(hints, i)
		default:
			t.Errorf("PROV028 as %s: %+v", i.Severity, i)
		}
	}
	if len(warnings) != 1 || warnings[0].RecordID != "" || !strings.Contains(warnings[0].Message, "2 records use the v1 schema") ||
		!strings.Contains(warnings[0].Message, "provenance migrate") {
		t.Fatalf("repo-level warning: %+v", warnings)
	}
	if len(hints) != 2 {
		t.Fatalf("per-record hints: %+v", hints)
	}
	for _, h := range hints {
		if h.RecordID == migrated.ID || h.RecordID == "" {
			t.Errorf("hint on the wrong record: %+v", h)
		}
	}

	// A fully migrated repo has none.
	clean := &record.FilesStore{Root: t.TempDir()}
	if err := clean.Save(newRecord("prov-2026-0000a002", record.StatusDraft), nil); err != nil {
		t.Fatal(err)
	}
	if n := rules(lintStore(t, clean), "")[RuleNotMigrated]; n != 0 {
		t.Fatalf("migrated repo has %d PROV028 issues", n)
	}
}
