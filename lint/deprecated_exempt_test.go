package lint

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/livecodelife/provenance/record"
)

// A deprecated record takes no new versions, so a change to its content or to
// the files its specs name can never be answered with a bump: PROV024 must not
// report it.

func committedOpen(t *testing.T, s *record.FilesStore, id string) *record.Record {
	t.Helper()
	r := newRecord(id, record.StatusOpen)
	r.AssociatedSpecs = []record.AssociatedSpec{{Path: "spec.sh", RunCommand: "sh"}}
	if err := os.WriteFile(filepath.Join(s.Root, "spec.sh"), []byte("true\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := s.Save(r, nil); err != nil {
		t.Fatal(err)
	}
	if _, err := record.Commit(s, r, nil, record.Bump{Author: "a@b"}); err != nil {
		t.Fatal(err)
	}
	return r
}

func deprecate(t *testing.T, s *record.FilesStore, id string) {
	t.Helper()
	r, doc, err := s.Load(id)
	if err != nil {
		t.Fatal(err)
	}
	r.Status = record.StatusDeprecated
	if err := s.Save(r, doc); err != nil {
		t.Fatal(err)
	}
}

func TestDeprecatedRecordIsNotFlaggedForEditedContent(t *testing.T) {
	s := &record.FilesStore{Root: t.TempDir()}
	r := committedOpen(t, s, "prov-2026-0000d101")
	r, doc, _ := s.Load(r.ID)
	r.Constraints = append(r.Constraints, "Edited after deprecation")
	r.Status = record.StatusDeprecated
	if err := s.Save(r, doc); err != nil {
		t.Fatal(err)
	}
	if n := rules(lintStore(t, s), r.ID)[RuleUnbumpedChange]; n != 0 {
		t.Fatalf("deprecated record with edited content: PROV024 x%d, want 0", n)
	}
}

func TestDeprecatedRecordIsNotFlaggedForEditedSpecFile(t *testing.T) {
	s := &record.FilesStore{Root: t.TempDir()}
	r := committedOpen(t, s, "prov-2026-0000d102")
	deprecate(t, s, r.ID)
	if n := rules(lintStore(t, s), r.ID)[RuleUnbumpedChange]; n != 0 {
		t.Fatalf("deprecated record, nothing changed: PROV024 x%d, want 0", n)
	}
	if err := os.WriteFile(filepath.Join(s.Root, "spec.sh"), []byte("echo changed\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if n := rules(lintStore(t, s), r.ID)[RuleUnbumpedChange]; n != 0 {
		t.Fatalf("deprecated record with an edited spec file: PROV024 x%d, want 0", n)
	}
}

func TestOpenRecordIsStillFlaggedForEditedSpecFile(t *testing.T) {
	s := &record.FilesStore{Root: t.TempDir()}
	r := committedOpen(t, s, "prov-2026-0000d103")
	if err := os.WriteFile(filepath.Join(s.Root, "spec.sh"), []byte("echo changed\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if n := rules(lintStore(t, s), r.ID)[RuleUnbumpedChange]; n != 1 {
		t.Fatalf("open record with an edited spec file: PROV024 x%d, want 1", n)
	}
}

func TestImplementedRecordIsStillFlaggedForEditedContent(t *testing.T) {
	s := &record.FilesStore{Root: t.TempDir()}
	r := committedOpen(t, s, "prov-2026-0000d104")
	r, doc, _ := s.Load(r.ID)
	r.Status = record.StatusImplemented
	r.Constraints = append(r.Constraints, "Edited")
	if err := s.Save(r, doc); err != nil {
		t.Fatal(err)
	}
	if n := rules(lintStore(t, s), r.ID)[RuleUnbumpedChange]; n != 1 {
		t.Fatalf("implemented record with edited content: PROV024 x%d, want 1", n)
	}
}
