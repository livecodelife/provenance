package record

import (
	"strings"
	"testing"
)

const v1Fixture = `# Leading comment survives.
id: prov-2026-1a2b3c4d
title: 'Users index'
status: open
created_at: "2026-10-15"
author: a@example.com
intent: |
  Add the users index.

  status: this line is inside a block scalar and must not be touched.
constraints:
  - Return JSON   # trailing comment
affected_scope:
  - app/users.rb
forbidden_scope: []
type: imprint
supersedes: ""
superseded_by: ""
implements: prov-2026-00000001
related: []
sealed_at_sha: ""
associated_specs:
  - path: specs/users.linespec
    type: linespec
    description: kept although unmodelled
associated_traces: []
monitors: []
tags: []

# A comment before an unknown key.
x_custom:
  nested: true
`

func mustParse(t *testing.T, src string) *Document {
	t.Helper()
	d, err := Parse([]byte(src))
	if err != nil {
		t.Fatal(err)
	}
	return d
}

func TestCodecUnchangedSaveIsByteIdentical(t *testing.T) {
	d := mustParse(t, v1Fixture)
	out, err := d.Encode(d.Record())
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != v1Fixture {
		t.Fatalf("unchanged save rewrote the file:\n%s", out)
	}
}

func TestCodecDecodesV1(t *testing.T) {
	r := mustParse(t, v1Fixture).Record()
	if r.Schema != 1 || r.Version != 1 {
		t.Fatalf("schema %d version %d, want 1/1", r.Schema, r.Version)
	}
	if r.Implements != "prov-2026-00000001" || r.Type != TypeImprint {
		t.Fatalf("implements %q type %q", r.Implements, r.Type)
	}
	if len(r.AssociatedSpecs) != 1 || len(r.AssociatedSpecs[0].Extra) != 1 {
		t.Fatalf("spec extra keys not kept: %+v", r.AssociatedSpecs)
	}
	if err := r.Validate(); err != nil {
		t.Fatal(err)
	}
}

func TestCodecScalarChangeTouchesOnlyThatLine(t *testing.T) {
	d := mustParse(t, v1Fixture)
	r := d.Record()
	r.Status = StatusImplemented
	r.SealedAtSHA = "abc123"
	out, err := d.Encode(r)
	if err != nil {
		t.Fatal(err)
	}
	want := strings.Replace(v1Fixture, "status: open\n", "status: implemented\n", 1)
	want = strings.Replace(want, `sealed_at_sha: ""`, "sealed_at_sha: abc123", 1)
	if string(out) != want {
		t.Fatalf("got:\n%s\nwant:\n%s", out, want)
	}
}

func TestCodecContentChangeSplicesBlock(t *testing.T) {
	d := mustParse(t, v1Fixture)
	r := d.Record()
	r.Intent = "Add the users index, paginated."
	r.Constraints = []string{"Return JSON", "Paginate at 50"}
	out, err := d.Encode(r)
	if err != nil {
		t.Fatal(err)
	}
	s := string(out)
	for _, keep := range []string{"# Leading comment survives.", "title: 'Users index'", "# A comment before an unknown key.", "x_custom:\n  nested: true", "description: kept although unmodelled"} {
		if !strings.Contains(s, keep) {
			t.Errorf("lost %q", keep)
		}
	}
	if !strings.Contains(s, "intent: Add the users index, paginated.\nconstraints:\n  - Return JSON\n  - Paginate at 50\naffected_scope:") {
		t.Errorf("content keys not spliced in place:\n%s", s)
	}
	back, err := Decode(out)
	if err != nil {
		t.Fatal(err)
	}
	if back.GenHash() != r.GenHash() {
		t.Error("re-decoded record hashes differently")
	}
}

func TestCodecMigrateToV2(t *testing.T) {
	d := mustParse(t, v1Fixture)
	r := d.Record()
	r.Schema = 2
	r.Version = 2
	r.Reason = "Fold bug prov-2026-deadbeef"
	r.ReasonKind = ReasonKindFix
	r.ImplementsVersion = 3
	r.ImplementsRepo = "github.com/livecodelife/linespec"
	r.Aliases = []string{"prov-2026-deadbeef"}
	r.Supersedes, r.SupersededBy = "", ""
	out, err := d.Encode(r)
	if err != nil {
		t.Fatal(err)
	}
	s := string(out)
	if strings.Contains(s, "supersedes:") || strings.Contains(s, "superseded_by:") {
		t.Errorf("v1-only keys survived migration:\n%s", s)
	}
	if !strings.Contains(s, "status: open\nversion: 2\nreason: Fold bug prov-2026-deadbeef\nreason_kind: fix\n") {
		t.Errorf("version/reason not placed after status:\n%s", s)
	}
	if !strings.Contains(s, "implements: prov-2026-00000001\nimplements_repo: github.com/livecodelife/linespec\nimplements_version: 3\n") {
		t.Errorf("implements pin not placed after implements:\n%s", s)
	}
	back, err := Decode(out)
	if err != nil {
		t.Fatal(err)
	}
	if back.Schema != 2 || back.Version != 2 || back.ImplementsVersion != 3 || len(back.Aliases) != 1 {
		t.Fatalf("round trip lost v2 fields: %+v", back)
	}
	if err := back.Validate(); err != nil {
		t.Fatal(err)
	}
	if again, _ := mustParse(t, s).Encode(back); string(again) != s {
		t.Error("second save of migrated record is not stable")
	}
}

func TestCodecMarshalNewRecordRoundTrips(t *testing.T) {
	r := &Record{
		Schema: 2, ID: "prov-2026-0badc0de", Title: "New", Type: TypeBlueprint,
		Status: StatusDraft, Version: 1, CreatedAt: "2026-10-15", Author: "a@b",
		Intent: "Multi\nline\n", Constraints: []string{"1.50", "true", ""},
		AssociatedSpecs: []AssociatedSpec{{Path: "x.linespec", Type: "linespec"}},
	}
	out, err := Marshal(r)
	if err != nil {
		t.Fatal(err)
	}
	back, err := Decode(out)
	if err != nil {
		t.Fatal(err)
	}
	if back.Intent != r.Intent || strings.Join(back.Constraints, "|") != "1.50|true|" || back.CreatedAt != "2026-10-15" {
		t.Fatalf("round trip changed values:\n%s", out)
	}
	if back.GenHash() != r.GenHash() {
		t.Fatal("round trip changed gen_hash")
	}
	d := mustParse(t, string(out))
	if again, _ := d.Encode(back); string(again) != string(out) {
		t.Fatal("unchanged save of marshalled record rewrote it")
	}
}

func TestCodecNoTrailingNewline(t *testing.T) {
	src := "id: prov-2026-1a2b3c4d\nstatus: draft\nintent: x"
	d := mustParse(t, src)
	r := d.Record()
	r.Intent = "z"
	r.Tags = []string{"t"}
	out, err := d.Encode(r)
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != "id: prov-2026-1a2b3c4d\nstatus: draft\nintent: z\ntags:\n  - t\n" {
		t.Fatalf("got %q", out)
	}
}

func TestCodecInsertBeforeRemovedKey(t *testing.T) {
	// version is inserted after status, exactly where the removed v1 key
	// supersedes begins.
	src := "id: prov-2026-1a2b3c4d\nstatus: open\nsupersedes: prov-2026-00000001\nintent: x\n"
	d := mustParse(t, src)
	r := d.Record()
	r.Schema, r.Version, r.Supersedes = 2, 1, ""
	out, err := d.Encode(r)
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != "id: prov-2026-1a2b3c4d\nstatus: open\nversion: 1\nintent: x\n" {
		t.Fatalf("got %q", out)
	}
}
