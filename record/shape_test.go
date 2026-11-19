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
        results:
          - error
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

// Literals produced by the code before shape existed (origin/main c6c11e2d).
func TestGenHashLiteralsFromBeforeShape(t *testing.T) {
	for _, c := range []struct {
		in   GenInput
		want string
	}{
		{GenInput{}, "sha256:7923a2ea36ef30a60c1a9cee38850f45e9dc2c661e1a7f052d4e03f784e69b7d"},
		{GenInput{Intent: "x", Constraints: []string{"c"}, AffectedScope: []string{"a.go"}}, "sha256:b1b56ca55ac730bd5bc1f5965236253ca18d6b03cc54730b79d8155d8ea05295"},
		{GenInput{Intent: "Users can\nsign in.\n", Constraints: []string{"No  passwords "}}, "sha256:68d37480ba807fe51b98239aa22a2334f50254d0cab63d313f278f8733624e1f"},
	} {
		for _, sh := range [][]any{nil, {}, {map[string]any{}}} {
			in := c.in
			in.Shape = sh
			if got := GenHash(in); got != c.want {
				t.Errorf("%+v shape %v: %s, want %s", c.in, sh, got, c.want)
			}
		}
	}
}

func TestShapeOrderedByPath(t *testing.T) {
	a, b := map[string]any{"path": "a.go"}, map[string]any{"path": "b.go"}
	if GenHash(GenInput{Shape: []any{b, a}}) != GenHash(GenInput{Shape: []any{a, b}}) {
		t.Error("entry order moved the hash")
	}
	x, y := map[string]any{"k": "2"}, map[string]any{"k": "1"}
	if GenHash(GenInput{Shape: []any{x, y}}) == GenHash(GenInput{Shape: []any{y, x}}) {
		t.Error("entries without a path must keep their order")
	}
}

func TestShapeTypedValuesDoNotPanic(t *testing.T) {
	typed := GenInput{Shape: []any{map[string]any{"path": "a.go", "imports": []string{"fmt"}, "package": nil}, map[string]string{"path": "b.go", "package": "b"}}}
	plain := GenInput{Shape: []any{map[string]any{"path": "a.go", "imports": []any{"fmt"}, "package": ""}, map[string]any{"path": "b.go", "package": "b"}}}
	if GenHash(typed) != GenHash(plain) {
		t.Error("typed and plain values hash differently")
	}
	if err := ValidateShape(typed.Shape); err != nil {
		t.Error(err)
	}
	// The hash accepts a top-level map[string]string, so validation does too;
	// an empty one is dropped like an empty map[string]any.
	if GenHash(GenInput{Shape: []any{map[string]string{}}}) != GenHash(GenInput{}) {
		t.Error("an empty map[string]string entry is not absent")
	}
	for name, v := range map[string]any{
		"float": 1.5, "int": 3, "bool": true, "nil entry": nil,
	} {
		shape := []any{map[string]any{"path": "a.go", "package": v}}
		if name == "nil entry" {
			shape = []any{nil}
		}
		if err := ValidateShape(shape); err == nil {
			t.Errorf("%s was accepted", name)
		}
		_ = GenHash(GenInput{Shape: shape}) // reported by ValidateShape, never a panic
	}
	if err := ValidateShape([]any{map[string]any{"path": "a.go", "extra": "x"}}); err == nil {
		t.Error("an unknown entry key was accepted")
	}
	r := &Record{Schema: 2, ID: "prov-2026-0000a1b1", Version: 1, Shape: []any{map[string]any{"package": 1.5}}}
	if r.Validate() == nil {
		t.Error("Validate accepted a float")
	}
}

func TestDecodeRejectsInvalidShapeAndSchema1Shape(t *testing.T) {
	y := "id: prov-2026-0000a1b1\nversion: 1\nintent: x\n"
	if _, err := Decode([]byte(y + "shape:\n  - path: a.go\n    extra: x\n")); err == nil || !strings.Contains(err.Error(), "extra") {
		t.Errorf("unknown key: %v", err)
	}
	v1 := "id: prov-2026-0000a1b1\nintent: x\n"
	for _, sh := range []string{"shape: []\n", "shape: ~\n", "shape:\n", "shape:\n  - {}\n", "shape:\n  - path: a.go\n"} {
		if _, err := Decode([]byte(v1 + sh)); err == nil || !strings.Contains(err.Error(), "schema 1") {
			t.Errorf("schema 1 %q: %v", sh, err)
		}
	}
}

func TestShapeRejectedInSchema1AndComplexKeys(t *testing.T) {
	if _, err := Decode([]byte("id: prov-2026-0000a1b1\nintent: x\nshape:\n  - path: a.go\n")); err == nil || !strings.Contains(err.Error(), "schema 1") {
		t.Errorf("v1 record with a shape: %v", err)
	}
	if _, err := Decode([]byte("id: prov-2026-0000a1b1\nversion: 1\nshape:\n  - ? [a, b]\n    : x\n")); err == nil || !strings.Contains(err.Error(), "key must be a scalar") {
		t.Errorf("complex key: %v", err)
	}
}

func TestMarshalDropsEmptyShapeEntries(t *testing.T) {
	r := &Record{Schema: 2, ID: "prov-2026-0000a1b1", Version: 1, Shape: []any{map[string]any{}}}
	b, err := Marshal(r)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(b), "shape") {
		t.Errorf("an empty shape entry was written:\n%s", b)
	}
	r.Shape = []any{map[string]any{}, map[string]any{"path": "a.go"}}
	b, _ = Marshal(r)
	back, _ := Decode(b)
	if len(back.Shape) != 1 {
		t.Errorf("shape after round trip: %v", back.Shape)
	}
}
