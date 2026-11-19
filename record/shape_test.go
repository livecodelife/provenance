package record

import (
	"bytes"
	"strings"
	"testing"
)

// Hashes below were produced by the code before shape existed (origin/main at
// c6c11e2d); a record with no shape must keep them forever.
func TestGenHashWithoutShapeIsUnchanged(t *testing.T) {
	// The sedum marker pins this record's gen_hash (check_sarif_run.go).
	d, err := Load("../provenance/prov-2026-0b7a4f45.yml")
	if err != nil {
		t.Fatal(err)
	}
	r := d.Record()
	if r.Shape != nil {
		t.Fatal("record unexpectedly has a shape")
	}
	if got, want := r.GenHash(), "sha256:dbf78db9ba625aefeed1e964eefa4d37f002ff9c2ff20d6c17bd027babfeb7d1"; got != want {
		t.Errorf("gen_hash %s, want %s", got, want)
	}
}

func TestEmptyShapeIsAbsent(t *testing.T) {
	base := GenInput{Intent: "x", Constraints: []string{"c"}}
	want := GenHash(base)
	for _, sh := range [][]any{nil, {}, {map[string]any{}}} {
		in := base
		in.Shape = sh
		if got := GenHash(in); got != want {
			t.Errorf("shape %v: %s, want %s", sh, got, want)
		}
		if strings.Contains(GenCanonical(in), "shape") {
			t.Errorf("shape %v appears in the canonical input", sh)
		}
	}
	in := base
	in.Shape = []any{map[string]any{"path": "a.go"}}
	if GenHash(in) == want || !strings.Contains(GenCanonical(in), `"shape":[{"path":"a.go"}]`) {
		t.Errorf("non-empty shape not hashed: %s", GenCanonical(in))
	}
}

func TestShapeIsNotContent(t *testing.T) {
	r := &Record{Schema: 2, Intent: "x"}
	before, _ := r.ContentHash(nil)
	r.Shape = []any{map[string]any{"path": "a.go"}}
	if after, _ := r.ContentHash(nil); after != before {
		t.Error("shape moved content_hash")
	}
}

const shapeRecord = `id: prov-2026-0000a1b1
title: T
status: open
version: 1
created_at: "2026-11-19"
author: a
intent: do it
constraints: []
affected_scope:
  - a.go
forbidden_scope: []
shape:
  - decls:
      - code: |-
          type T struct {
          	N int
          }
        kind: type
        name: T
    imports:
      - fmt
    path: a.go
    stubs:
      - doc: F does x.
        name: F
        params:
          - name: x
            type: int
        returns: error
tags: []
`

func TestShapeRoundTrip(t *testing.T) {
	d, err := Parse([]byte(shapeRecord))
	if err != nil {
		t.Fatal(err)
	}
	r := d.Record()
	if len(r.Shape) != 1 {
		t.Fatalf("shape %v", r.Shape)
	}
	out, err := d.Encode(r)
	if err != nil || !bytes.Equal(out, []byte(shapeRecord)) {
		t.Fatalf("unchanged record changed (err %v):\n%s", err, out)
	}
	// An edit rewrites only the shape key; key order and other keys stay.
	r.Shape[0].(map[string]any)["imports"] = []any{"fmt", "os"}
	out, err = d.Encode(r)
	if err != nil {
		t.Fatal(err)
	}
	if want := strings.Replace(shapeRecord, "      - fmt\n", "      - fmt\n      - os\n", 1); string(out) != want {
		t.Errorf("edit:\n%s\nwant:\n%s", out, want)
	}
	// A fresh marshal puts shape after forbidden_scope and decodes back equal.
	fresh, err := Marshal(r)
	if err != nil {
		t.Fatal(err)
	}
	back, err := Decode(fresh)
	if err != nil {
		t.Fatal(err)
	}
	if back.GenHash() != r.GenHash() {
		t.Error("round trip moved gen_hash")
	}
	if i, j := bytes.Index(fresh, []byte("\nforbidden_scope:")), bytes.Index(fresh, []byte("\nshape:")); i < 0 || j < i {
		t.Errorf("shape not after forbidden_scope:\n%s", fresh)
	}
	// A record with no shape never writes the key.
	if b, _ := Marshal(&Record{Schema: 2, ID: "prov-2026-0000a1b2", Version: 1}); bytes.Contains(b, []byte("shape")) {
		t.Errorf("shape written for a record without one:\n%s", b)
	}
}

// A shape-only change keeps content_hash, so it is a reason-only version.
func TestShapeOnlyRevise(t *testing.T) {
	s := &FilesStore{Root: t.TempDir()}
	r := &Record{Schema: 2, ID: "prov-2026-0000a1b3", Status: StatusImplemented, Version: 1, Intent: "x", AffectedScope: []string{"a.go"}}
	if _, err := Commit(s, r, nil, Bump{}); err != nil {
		t.Fatal(err)
	}
	r.Status = StatusImplemented
	r.Shape = []any{map[string]any{"path": "a.go"}}
	if _, err := Commit(s, r, nil, Bump{}); err != ErrNoChange {
		t.Fatalf("no reason: %v", err)
	}
	snap, err := Commit(s, r, nil, Bump{Reason: "import skeleton", Author: "graft"})
	if err != nil {
		t.Fatal(err)
	}
	if snap.Version != 2 || r.Status != StatusOpen {
		t.Errorf("version %d status %s", snap.Version, r.Status)
	}
	got, err := s.Snapshot(r.ID, 2)
	if err != nil || len(got.Content.Shape) != 1 {
		t.Fatalf("snapshot lost the shape: %v %v", err, got)
	}
}
