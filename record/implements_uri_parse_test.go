package record

import "testing"

func TestParseImplementsURI(t *testing.T) {
	const did = "at://did:plc:ragnj6bgwbfkhz4luybnthat/"
	cases := []struct {
		uri, id string
		ok      bool
	}{
		{did + "dev.provisio.provenance.record/prov-2026-0000b001", "prov-2026-0000b001", true},
		{"at://did:web:example.com/dev.provisio.provenance.record/prov-2026-0000b001-svc", "prov-2026-0000b001-svc", true},
		{did + "dev.provisio.provenance.version/prov-2026-0000b001", "", false},
		{"https://example.com/dev.provisio.provenance.record/x", "", false},
		{did + "dev.provisio.provenance.record/", "", false},
		{did + "dev.provisio.provenance.record/a/b", "", false},
		{"", "", false},
	}
	for _, c := range cases {
		id, ok := ParseImplementsURI(c.uri)
		if id != c.id || ok != c.ok {
			t.Errorf("ParseImplementsURI(%q) = %q, %v; want %q, %v", c.uri, id, ok, c.id, c.ok)
		}
	}
}
