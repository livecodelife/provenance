package record

import (
	"strings"
	"testing"
)

const parentURI = "at://did:plc:ragnj6bgwbfkhz4luybnthat/dev.provisio.provenance.record/prov-2026-0000b001"

func uriRecord() *Record {
	return &Record{
		Schema: 2, ID: "prov-2026-0000a001", Title: "Users index", Type: TypeImprint, Status: StatusOpen, Version: 1,
		CreatedAt: "2026-10-01", Intent: "Add the users index action.",
		AffectedScope: []string{"app/users.rb"},
		Implements:    "prov-2026-0000b001", ImplementsVersion: 2,
	}
}

// implements_uri is lineage like implements: setting it must not move the hash.
func TestImplementsURIIsNotContent(t *testing.T) {
	plain := uriRecord()
	linked := uriRecord()
	linked.ImplementsURI = parentURI
	a, _ := plain.ContentHash(nil)
	b, _ := linked.ContentHash(nil)
	if a != b {
		t.Fatalf("implements_uri changed the content hash: %s != %s", a, b)
	}
}

func TestImplementsURIRoundTripsThroughYAML(t *testing.T) {
	r := uriRecord()
	r.ImplementsURI = parentURI
	data, err := Marshal(r)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(data), "implements_uri: "+parentURI) {
		t.Fatalf("implements_uri missing from the file:\n%s", data)
	}
	if i, j := strings.Index(string(data), "implements:"), strings.Index(string(data), "implements_uri:"); !(i < j) {
		t.Fatalf("implements_uri is not written after implements:\n%s", data)
	}
	doc, err := Parse(data)
	if err != nil {
		t.Fatal(err)
	}
	if got := doc.Record().ImplementsURI; got != parentURI {
		t.Fatalf("parsed implements_uri = %q", got)
	}
	if plain, _ := Marshal(uriRecord()); strings.Contains(string(plain), "implements_uri") {
		t.Fatalf("an unset implements_uri is written:\n%s", plain)
	}
}

func TestImplementsURIRoundTripsThroughPublishedContent(t *testing.T) {
	r := uriRecord()
	r.ImplementsURI = parentURI
	c := PublishedContentOf(r)
	if c.ImplementsURI != parentURI {
		t.Fatalf("published content implementsUri = %q", c.ImplementsURI)
	}
	if got := c.Record().ImplementsURI; got != parentURI {
		t.Fatalf("round trip implements_uri = %q", got)
	}
}

func TestImplementsURIWithoutImplementsIsInvalid(t *testing.T) {
	r := uriRecord()
	r.Implements, r.ImplementsVersion, r.ImplementsURI = "", 0, parentURI
	if err := r.Validate(); err == nil {
		t.Fatal("implements_uri without implements validated")
	}
}
