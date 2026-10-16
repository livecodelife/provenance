package provenancelex

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

const samples = "../../samples"

func loadSample(t *testing.T, name string) []byte {
	t.Helper()
	data, err := os.ReadFile(filepath.Join(samples, name))
	if err != nil {
		t.Fatal(err)
	}
	return data
}

// sameJSON compares two JSON documents as values, ignoring $type, which the
// generated types add.
func sameJSON(t *testing.T, got, want []byte) {
	t.Helper()
	var g, w map[string]any
	if err := json.Unmarshal(got, &g); err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(want, &w); err != nil {
		t.Fatal(err)
	}
	delete(g, "$type")
	delete(w, "$type")
	if !reflect.DeepEqual(g, w) {
		t.Fatalf("round trip changed the document:\n got %s\nwant %s", got, want)
	}
}

// The sample head and three versions (lexicons/samples, written by the Go
// mapping in the provenance module) survive the generated types unchanged.
func TestVersionsRoundTrip(t *testing.T) {
	for _, name := range []string{"version-1.json", "version-2.json", "version-3.json"} {
		t.Run(name, func(t *testing.T) {
			data := loadSample(t, name)
			var v Version
			if err := json.Unmarshal(data, &v); err != nil {
				t.Fatal(err)
			}
			if v.RecordId != "prov-2026-0000a001" || v.Content == nil || v.ContentHash == nil || len(v.SpecHashes) != 3 {
				t.Fatalf("decoded %+v", v)
			}
			// The unresolved spec hash is the empty string, not a missing entry.
			if v.SpecHashes[2] == nil || *v.SpecHashes[2] != "" {
				t.Fatalf("unresolved spec hash decoded as %v", v.SpecHashes[2])
			}
			v.LexiconTypeID = "dev.provisio.provenance.version"
			back, err := json.Marshal(v)
			if err != nil {
				t.Fatal(err)
			}
			sameJSON(t, back, data)
		})
	}
}

func TestChainLinks(t *testing.T) {
	var v1, v2 Version
	if err := json.Unmarshal(loadSample(t, "version-1.json"), &v1); err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(loadSample(t, "version-2.json"), &v2); err != nil {
		t.Fatal(err)
	}
	if v1.Parent != nil || v1.ParentHash != nil {
		t.Fatal("version 1 has a parent")
	}
	if v2.ParentHash == nil || *v2.ParentHash != *v1.ContentHash {
		t.Fatalf("version 2 does not chain to version 1's content hash")
	}
}

func TestHeadRoundTrip(t *testing.T) {
	data := loadSample(t, "head.json")
	var h Record
	if err := json.Unmarshal(data, &h); err != nil {
		t.Fatal(err)
	}
	if h.Version != 3 || h.Latest == nil || h.Content == nil || h.Project == nil {
		t.Fatalf("decoded %+v", h)
	}
	h.LexiconTypeID = "dev.provisio.provenance.record"
	back, err := json.Marshal(h)
	if err != nil {
		t.Fatal(err)
	}
	sameJSON(t, back, data)
}

// A record is stored on the PDS as DAG-CBOR: the same documents survive that too.
func TestCBORRoundTrip(t *testing.T) {
	for _, name := range []string{"version-1.json", "version-2.json", "version-3.json"} {
		data := loadSample(t, name)
		var v Version
		if err := json.Unmarshal(data, &v); err != nil {
			t.Fatal(err)
		}
		var buf bytes.Buffer
		if err := v.MarshalCBOR(&buf); err != nil {
			t.Fatalf("%s: %v", name, err)
		}
		var back Version
		if err := back.UnmarshalCBOR(&buf); err != nil {
			t.Fatalf("%s: %v", name, err)
		}
		out, err := json.Marshal(back)
		if err != nil {
			t.Fatal(err)
		}
		sameJSON(t, out, data)
	}
	data := loadSample(t, "head.json")
	var h Record
	if err := json.Unmarshal(data, &h); err != nil {
		t.Fatal(err)
	}
	var buf bytes.Buffer
	if err := h.MarshalCBOR(&buf); err != nil {
		t.Fatal(err)
	}
	var back Record
	if err := back.UnmarshalCBOR(&buf); err != nil {
		t.Fatal(err)
	}
	out, _ := json.Marshal(back)
	sameJSON(t, out, data)
}
