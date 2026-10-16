package evidence

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

// Two runs of one version with the same At second must both be stored, in
// append order, and the first keeps the plain <version>-<nanos>.json name.
func TestAppendSameSecondDoesNotCollide(t *testing.T) {
	s := &Store{Root: t.TempDir()}
	r := testRecord()
	at := "2026-10-16T09:00:00Z"
	for i, runner := range []string{"first", "second", "third"} {
		if err := s.Append(New(r, Current{ContentHash: "sha256:aa"}, true, at, runner)); err != nil {
			t.Fatalf("append %d: %v", i, err)
		}
	}
	runs, err := s.List(r.ID)
	if err != nil {
		t.Fatal(err)
	}
	if len(runs) != 3 || runs[0].Runner != "first" || runs[1].Runner != "second" || runs[2].Runner != "third" {
		t.Fatalf("runs = %+v, want first, second, third in order", runs)
	}
	latest, err := s.Latest(r.ID, r.Version)
	if err != nil || latest == nil || latest.Runner != "third" {
		t.Fatalf("latest = %+v, %v; want third", latest, err)
	}
	if _, err := os.Stat(filepath.Join(s.dir(r.ID), "1-1792141200000000000.json")); err != nil {
		t.Fatalf("first run lost its plain name: %v", err)
	}
}

// A different version with the same At never needs a suffix, and an existing
// file is never overwritten.
func TestAppendKeepsExistingFiles(t *testing.T) {
	s := &Store{Root: t.TempDir()}
	r := testRecord()
	at := "2026-10-16T09:00:00Z"
	if err := s.Append(New(r, Current{ContentHash: "sha256:aa"}, true, at, "one")); err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(s.dir(r.ID), "1-1792141200000000000.json")
	before, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	r2 := *r
	r2.Version = 2
	if err := s.Append(New(&r2, Current{ContentHash: "sha256:bb"}, true, at, "v2")); err != nil {
		t.Fatal(err)
	}
	if err := s.Append(New(r, Current{ContentHash: "sha256:aa"}, false, at, "again")); err != nil {
		t.Fatal(err)
	}
	after, err := os.ReadFile(path)
	if err != nil || string(after) != string(before) {
		t.Fatalf("first file changed or vanished: %v", err)
	}
	if _, err := os.Stat(filepath.Join(s.dir(r.ID), "2-1792141200000000000.json")); err != nil {
		t.Fatalf("version 2 should keep the plain name: %v", err)
	}
}

// The retry is bounded: with every candidate name taken, Append gives up with
// an error instead of looping, and writes nothing more.
func TestAppendGivesUpAtRetryBound(t *testing.T) {
	s := &Store{Root: t.TempDir()}
	r := testRecord()
	if err := os.MkdirAll(s.dir(r.ID), 0o755); err != nil {
		t.Fatal(err)
	}
	for i := int64(0); i < maxAppendRetries; i++ {
		name := fmt.Sprintf("1-%d.json", 1792141200000000000+i)
		if err := os.WriteFile(filepath.Join(s.dir(r.ID), name), []byte("{}"), 0o444); err != nil {
			t.Fatal(err)
		}
	}
	if err := s.Append(New(r, Current{ContentHash: "sha256:aa"}, true, "2026-10-16T09:00:00Z", "x")); err == nil {
		t.Fatal("Append succeeded with every name taken")
	}
	entries, _ := os.ReadDir(s.dir(r.ID))
	if int64(len(entries)) != maxAppendRetries {
		t.Fatalf("%d files, want %d", len(entries), maxAppendRetries)
	}
}
