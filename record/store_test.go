package record

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func newStore(t *testing.T) (*FilesStore, *Record) {
	t.Helper()
	root := t.TempDir()
	if err := os.MkdirAll(filepath.Join(root, "specs"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "specs", "a.linespec"), []byte("TEST a\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	s := &FilesStore{Root: root}
	r := &Record{
		Schema: 2, ID: "prov-2026-0badc0de", Title: "Users", Type: TypeImprint,
		Status: StatusDraft, Version: 1, CreatedAt: "2026-10-15", Author: "a@b",
		Intent: "Add users.", Constraints: []string{"JSON"},
		Implements: "prov-2026-00000001", ImplementsVersion: 1,
		AssociatedSpecs: []AssociatedSpec{{Path: "specs/a.linespec", Type: "linespec", RunCommand: "rm -rf /"}, {Path: "specs/missing.linespec"}},
	}
	if err := s.Save(r, nil); err != nil {
		t.Fatal(err)
	}
	return s, r
}

var t0 = time.Date(2026, 10, 15, 18, 0, 0, 0, time.FixedZone("x", -7*3600))

func TestStoreCommitChain(t *testing.T) {
	s, r := newStore(t)
	v1, err := Commit(s, r, nil, Bump{Author: "a@b", Now: t0})
	if err != nil {
		t.Fatal(err)
	}
	if v1.Version != 1 || v1.ParentHash != "" || v1.CreatedAt != "2026-10-16T01:00:00Z" {
		t.Fatalf("v1 envelope: %+v", v1)
	}
	if v1.Content.AssociatedSpecs[0].RunCommand != "" {
		t.Fatal("run_command leaked into a snapshot")
	}
	if v1.SpecHashes[0] == nil || v1.SpecHashes[1] != nil {
		t.Fatalf("spec hashes: want [hash, null], got %v", v1.SpecHashes)
	}

	// No content change and no reason: refused.
	r, doc, err := s.Load(r.ID)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := Commit(s, r, doc, Bump{Now: t0}); !errors.Is(err, ErrNoChange) {
		t.Fatalf("err = %v, want ErrNoChange", err)
	}
	// A content change needs a reason.
	r.Constraints = append(r.Constraints, "Paginate")
	if _, err := Commit(s, r, doc, Bump{Now: t0}); !errors.Is(err, ErrReasonRequired) {
		t.Fatalf("err = %v, want ErrReasonRequired", err)
	}
	// Implemented record: a bump reopens it.
	r.Status, r.SealedAtSHA = StatusImplemented, "abc"
	v2, err := Commit(s, r, doc, Bump{Reason: "Paginate", ReasonKind: ReasonKindFix, Now: t0.Add(time.Hour)})
	if err != nil {
		t.Fatal(err)
	}
	if v2.Version != 2 || v2.ParentHash != v1.ContentHash || v2.ContentHash == v1.ContentHash {
		t.Fatalf("v2 chain: %+v", v2)
	}
	saved, doc, err := s.Load(r.ID)
	if err != nil {
		t.Fatal(err)
	}
	if saved.Version != 2 || saved.Reason != "Paginate" || saved.ReasonKind != ReasonKindFix || saved.Status != StatusOpen || saved.SealedAtSHA != "" {
		t.Fatalf("saved record: %+v", saved)
	}

	// Reason-only version: same content hash, legal with a reason.
	v3, err := Commit(s, saved, doc, Bump{Reason: "Folded bug with no new content", Now: t0})
	if err != nil {
		t.Fatal(err)
	}
	if v3.ContentHash != v2.ContentHash || v3.ParentHash != v2.ContentHash {
		t.Fatalf("reason-only version: %+v", v3)
	}

	// A repin (implements_version) or title edit is not a content change.
	saved, doc, _ = s.Load(r.ID)
	saved.ImplementsVersion, saved.Title = 5, "Renamed"
	if _, err := Commit(s, saved, doc, Bump{Now: t0}); !errors.Is(err, ErrNoChange) {
		t.Fatalf("repin bumped: err = %v", err)
	}

	vs, err := s.Versions(r.ID)
	if err != nil || len(vs) != 3 {
		t.Fatalf("versions %v, %v", vs, err)
	}
	for _, v := range vs {
		snap, err := s.Snapshot(r.ID, v)
		if err != nil {
			t.Fatal(err)
		}
		if snap.Content.Title != "Users" || snap.Content.Status != "" || snap.Content.Implements != "prov-2026-00000001" {
			t.Fatalf("v%d content: %+v", v, snap.Content)
		}
	}
	name := filepath.Join(s.versionsDir(r.ID), v2.FileName())
	if !strings.HasPrefix(filepath.Base(name), "2-") || len(filepath.Base(name)) != len("2-.yml")+64 {
		t.Fatalf("snapshot name %s", name)
	}
}

func TestStoreDeprecatedIsTerminal(t *testing.T) {
	s, r := newStore(t)
	r.Status = StatusDeprecated
	if _, err := Commit(s, r, nil, Bump{Reason: "x"}); !errors.Is(err, ErrDeprecated) {
		t.Fatalf("err = %v, want ErrDeprecated", err)
	}
}

func TestStoreRefusesV1(t *testing.T) {
	s, r := newStore(t)
	r.Schema = 1
	if _, err := Commit(s, r, nil, Bump{Reason: "x"}); !errors.Is(err, ErrNotV2) {
		t.Fatalf("err = %v, want ErrNotV2", err)
	}
}

func TestStoreAppendIsImmutable(t *testing.T) {
	s, r := newStore(t)
	v1, err := Commit(s, r, nil, Bump{Now: t0})
	if err != nil {
		t.Fatal(err)
	}
	// Identical bytes: no-op.
	if err := s.Append(v1); err != nil {
		t.Fatalf("re-append identical: %v", err)
	}
	// Same path, different bytes: refused.
	changed := *v1
	changed.Reason = "rewritten history"
	if err := s.Append(&changed); err == nil || !strings.Contains(err.Error(), "already exists") {
		t.Fatalf("overwrite allowed: %v", err)
	}
	// Broken chain: refused.
	bad := *v1
	bad.Version, bad.ParentHash = 3, v1.ContentHash
	if err := s.Append(&bad); err == nil || !strings.Contains(err.Error(), "does not follow") {
		t.Fatalf("gap allowed: %v", err)
	}
	bad.Version, bad.ParentHash = 2, "sha256:wrong"
	if err := s.Append(&bad); err == nil || !strings.Contains(err.Error(), "parent_hash") {
		t.Fatalf("wrong parent allowed: %v", err)
	}
	// Tampered content: hash check fails on read.
	path := filepath.Join(s.versionsDir(r.ID), v1.FileName())
	data, _ := os.ReadFile(path)
	os.Chmod(path, 0o644)
	os.WriteFile(path, []byte(strings.Replace(string(data), "Add users.", "Add admins.", 1)), 0o644)
	if _, err := s.Snapshot(r.ID, 1); err == nil || !strings.Contains(err.Error(), "content hashes to") {
		t.Fatalf("tampering not detected: %v", err)
	}
}

func TestStoreSnapshotRoundTrip(t *testing.T) {
	s, r := newStore(t)
	snap, err := Commit(s, r, nil, Bump{Now: t0})
	if err != nil {
		t.Fatal(err)
	}
	data, err := MarshalSnapshot(snap)
	if err != nil {
		t.Fatal(err)
	}
	back, err := UnmarshalSnapshot(data)
	if err != nil {
		t.Fatal(err)
	}
	again, _ := MarshalSnapshot(back)
	if string(again) != string(data) {
		t.Fatalf("snapshot encoding not stable:\n%s\n---\n%s", data, again)
	}
	ids, err := s.List()
	if err != nil || len(ids) != 1 || ids[0] != r.ID {
		t.Fatalf("List = %v, %v", ids, err)
	}
	if err := s.Remove(r.ID); err != nil {
		t.Fatal(err)
	}
	if ids, _ := s.List(); len(ids) != 0 {
		t.Fatalf("Remove left %v", ids)
	}
	if vs, _ := s.Versions(r.ID); len(vs) != 1 {
		t.Fatal("Remove deleted snapshots")
	}
}
