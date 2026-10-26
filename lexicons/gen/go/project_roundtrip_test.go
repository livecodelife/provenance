package provenancelex

import (
	"bytes"
	"encoding/json"
	"testing"
)

// The sample project (lexicons/samples/project.json) survives the generated
// type unchanged, and the generated type has CBOR encoding.
func TestProjectRoundTrip(t *testing.T) {
	data := loadSample(t, "project.json")
	var p Project
	if err := json.Unmarshal(data, &p); err != nil {
		t.Fatal(err)
	}
	if p.Name != "Provisio" || p.Description == nil || *p.Description == "" || p.CreatedAt == "" {
		t.Fatalf("decoded %+v", p)
	}
	p.LexiconTypeID = "dev.provisio.project"
	back, err := json.Marshal(p)
	if err != nil {
		t.Fatal(err)
	}
	sameJSON(t, back, data)

	var buf bytes.Buffer
	if err := p.MarshalCBOR(&buf); err != nil {
		t.Fatal(err)
	}
	var again Project
	if err := again.UnmarshalCBOR(&buf); err != nil {
		t.Fatal(err)
	}
	if again.Name != p.Name || again.CreatedAt != p.CreatedAt || again.Description == nil || *again.Description != *p.Description {
		t.Fatalf("CBOR round trip changed the project: %+v", again)
	}
}

// A project without a description still decodes, and the description stays absent.
func TestProjectWithoutDescription(t *testing.T) {
	var p Project
	if err := json.Unmarshal([]byte(`{"name":"Provisio","createdAt":"2026-10-26T12:00:00.000Z"}`), &p); err != nil {
		t.Fatal(err)
	}
	if p.Description != nil {
		t.Fatalf("description decoded as %v", *p.Description)
	}
}
