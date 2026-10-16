package evidence

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/livecodelife/provenance/record"
)

func TestCodeHashCorpus(t *testing.T) {
	data, err := os.ReadFile("../conformance/hashing/cases.json")
	if err != nil {
		t.Fatal(err)
	}
	var c struct {
		Cases []struct {
			ID        string   `json:"id"`
			Kind      string   `json:"kind"`
			Regions   []Region `json:"regions"`
			Canonical string   `json:"canonical"`
			Hash      string   `json:"hash"`
			Error     bool     `json:"error"`
		} `json:"cases"`
	}
	if err := json.Unmarshal(data, &c); err != nil {
		t.Fatal(err)
	}
	n := 0
	for _, tc := range c.Cases {
		if tc.Kind != "code_hash" {
			continue
		}
		n++
		h, err := CodeHash(tc.Regions)
		if tc.Error {
			if !errors.Is(err, ErrUnstamped) {
				t.Errorf("%s: err = %v, want ErrUnstamped", tc.ID, err)
			}
			continue
		}
		if err != nil || h != tc.Hash || CodeCanonical(tc.Regions) != tc.Canonical {
			t.Errorf("%s: got %s %s (%v), want %s %s", tc.ID, h, CodeCanonical(tc.Regions), err, tc.Hash, tc.Canonical)
		}
	}
	if n < 6 {
		t.Fatalf("only %d code_hash cases", n)
	}
}

func testRecord() *record.Record {
	return &record.Record{Schema: 2, ID: "prov-2026-0000a001", Version: 1, Intent: "x", Constraints: []string{"c"}}
}

var regions = []Region{{Gen: "sha256:aa", Pkg: "sha256:01"}}

func TestCheck(t *testing.T) {
	r := testRecord()
	cur := CurrentFor(r, nil, regions)
	e := New(r, cur, true, "2026-10-15T00:00:00Z", "test")
	if v := Check(e, cur); !v.Verified() || len(v.Reasons) != 0 {
		t.Fatalf("fresh run: %+v", v)
	}
	if v := Check(nil, cur); v.State != Missing {
		t.Fatalf("no evidence: %+v", v)
	}

	changed := testRecord()
	changed.Constraints = append(changed.Constraints, "more")
	if v := Check(e, CurrentFor(changed, nil, regions)); v.State != Stale || strings.Join(v.Reasons, ",") != "content" {
		t.Fatalf("content change: %+v", v)
	}
	r2 := testRecord()
	r2.AssociatedSpecs = []record.AssociatedSpec{{Path: "a.linespec"}}
	withSpec := CurrentFor(r2, func(string) (string, bool) { return "sha256:s1", true }, regions)
	e2 := New(r2, withSpec, true, "2026-10-15T00:00:00Z", "")
	editedSpec := CurrentFor(r2, func(string) (string, bool) { return "sha256:s2", true }, regions)
	if v := Check(e2, editedSpec); v.State != Stale || !contains(v.Reasons, "specs") {
		t.Fatalf("spec change: %+v", v)
	}
	if v := Check(e, CurrentFor(r, nil, []Region{{Gen: "sha256:aa", Pkg: "sha256:09"}})); v.State != Stale || strings.Join(v.Reasons, ",") != "code" {
		t.Fatalf("code change: %+v", v)
	}
	unstamped := CurrentFor(r, nil, []Region{{Gen: "sha256:aa"}})
	if v := Check(e, unstamped); v.State != Unknown || v.Verified() || strings.Join(v.Reasons, ",") != "code_unknown" {
		t.Fatalf("unstamped code: %+v", v)
	}
	behind := cur
	behind.Behind = true
	if v := Check(e, behind); v.State != Stale || strings.Join(v.Reasons, ",") != "parent" {
		t.Fatalf("parent behind: %+v", v)
	}
	failed := New(r, cur, false, "2026-10-15T00:00:00Z", "")
	if v := Check(failed, cur); v.State != Fresh || v.Verified() {
		t.Fatalf("failed run counts as verified: %+v", v)
	}
}

func contains(l []string, s string) bool {
	for _, x := range l {
		if x == s {
			return true
		}
	}
	return false
}

func TestStore(t *testing.T) {
	s := &Store{Root: t.TempDir()}
	r := testRecord()
	cur := CurrentFor(r, nil, regions)
	t0 := time.Date(2026, 10, 15, 1, 0, 0, 0, time.UTC)
	for i, passed := range []bool{false, true} {
		if err := s.Append(New(r, cur, passed, t0.Add(time.Duration(i)*time.Minute).Format(time.RFC3339Nano), "")); err != nil {
			t.Fatal(err)
		}
	}
	r.Version = 2
	if err := s.Append(New(r, cur, false, t0.Add(time.Hour).Format(time.RFC3339Nano), "")); err != nil {
		t.Fatal(err)
	}
	if runs, _ := s.List(r.ID); len(runs) != 3 {
		t.Fatalf("List = %d runs", len(runs))
	}
	if e, _ := s.Latest(r.ID, 1); e == nil || !e.Passed {
		t.Fatalf("latest v1 = %+v, want the passing run", e)
	}
	if v, _ := s.Status(r, cur); v.State != Fresh || v.Passed {
		t.Fatalf("status of v2 = %+v, want its failed run", v)
	}
	if e, _ := s.Latest(r.ID, 3); e != nil {
		t.Fatal("evidence for a version that was never run")
	}
	// Append-only: the same run twice is refused.
	r.Version = 1
	if err := s.Append(New(r, cur, true, t0.Format(time.RFC3339Nano), "")); err == nil {
		t.Fatal("overwrote an existing run")
	}
}
