package record

import (
	"encoding/json"
	"os"
	"reflect"
	"strings"
	"testing"
)

// The corpus is read in place, never copied: a disagreement is a bug in this
// package or in the corpus (docs/HASHING.md §6).
const corpusPath = "../conformance/hashing/cases.json"

type corpusCase struct {
	ID         string                     `json:"id"`
	Kind       string                     `json:"kind"`
	Input      string                     `json:"input"`
	Output     string                     `json:"output"`
	Record     map[string]json.RawMessage `json:"record"`
	SpecFiles  map[string]string          `json:"spec_files"`
	Canonical  string                     `json:"canonical"`
	Hash       string                     `json:"hash"`
	Content    string                     `json:"content"`
	YAML       string                     `json:"yaml"`
	Fields     map[string]json.RawMessage `json:"fields"`
	Error      bool                       `json:"error"`
	ErrorNames string                     `json:"error_names"`
}

func loadCorpus(t *testing.T) []corpusCase {
	t.Helper()
	data, err := os.ReadFile(corpusPath)
	if err != nil {
		t.Fatal(err)
	}
	var c struct {
		Format string       `json:"format"`
		Cases  []corpusCase `json:"cases"`
	}
	if err := json.Unmarshal(data, &c); err != nil {
		t.Fatal(err)
	}
	if c.Format != "provenance.hashing" || len(c.Cases) == 0 {
		t.Fatalf("unexpected corpus format %q with %d cases", c.Format, len(c.Cases))
	}
	return c.Cases
}

// recordFromJSON builds a Record from a case's record object, the hash input
// model: hashed keys are strings, lists (null items allowed) or null.
func recordFromJSON(t *testing.T, obj map[string]json.RawMessage) *Record {
	t.Helper()
	r := &Record{Schema: 2}
	str := func(key string) string {
		var s *string
		if raw, ok := obj[key]; ok {
			if err := json.Unmarshal(raw, &s); err != nil {
				t.Fatalf("%s: %v", key, err)
			}
		}
		if s == nil {
			return ""
		}
		return *s
	}
	list := func(key string) []string {
		var items []*string
		if raw, ok := obj[key]; ok {
			if err := json.Unmarshal(raw, &items); err != nil {
				t.Fatalf("%s: %v", key, err)
			}
		}
		out := make([]string, len(items))
		for i, s := range items {
			if s != nil {
				out[i] = *s
			}
		}
		return out
	}
	r.Intent = str("intent")
	r.Constraints = list("constraints")
	r.AffectedScope = list("affected_scope")
	r.ForbiddenScope = list("forbidden_scope")
	r.Type = Type(str("type"))
	r.Implements = str("implements")
	r.Title = str("title")
	r.Tags = list("tags")
	r.Related = list("related")
	if raw, ok := obj["associated_specs"]; ok {
		var specs []map[string]any
		if err := json.Unmarshal(raw, &specs); err != nil {
			t.Fatalf("associated_specs: %v", err)
		}
		for _, s := range specs {
			p, _ := s["path"].(string)
			r.AssociatedSpecs = append(r.AssociatedSpecs, AssociatedSpec{Path: p})
		}
	}
	return r
}

func TestConformance(t *testing.T) {
	for _, c := range loadCorpus(t) {
		t.Run(c.ID, func(t *testing.T) {
			switch c.Kind {
			case "normalize":
				if got := Normalize(c.Input); got != c.Output {
					t.Errorf("Normalize(%q) = %q, want %q", c.Input, got, c.Output)
				}
			case "gen_hash":
				r := recordFromJSON(t, c.Record)
				if got := GenCanonical(r.GenInput()); got != c.Canonical {
					t.Errorf("canonical\n got %s\nwant %s", got, c.Canonical)
				}
				if got := GenHash(r.GenInput()); got != c.Hash {
					t.Errorf("hash %s, want %s", got, c.Hash)
				}
			case "content_hash":
				r := recordFromJSON(t, c.Record)
				resolve := func(path string) (string, bool) {
					content, ok := c.SpecFiles[path]
					if !ok {
						return "", false
					}
					return SpecHash([]byte(content)), true
				}
				specs, _ := r.SpecHashes(resolve)
				if got := ContentCanonical(r, specs); got != c.Canonical {
					t.Errorf("canonical\n got %s\nwant %s", got, c.Canonical)
				}
				if got, _ := r.ContentHash(resolve); got != c.Hash {
					t.Errorf("hash %s, want %s", got, c.Hash)
				}
			case "spec_hash":
				if got := SpecHash([]byte(c.Content)); got != c.Hash {
					t.Errorf("hash %s, want %s", got, c.Hash)
				}
			case "canonical":
				// Run by TestCanonicalCorpus in canonical_test.go.
			case "decode":
				fields, err := DecodeFields([]byte(c.YAML))
				if c.Error {
					if err == nil || !strings.Contains(err.Error(), c.ErrorNames) {
						t.Fatalf("err = %v, want one containing %q", err, c.ErrorNames)
					}
					return
				}
				if err != nil {
					t.Fatal(err)
				}
				for key, raw := range c.Fields {
					var want any
					_ = json.Unmarshal(raw, &want)
					if got := valueJSON(fields[key]); !reflect.DeepEqual(got, want) {
						t.Errorf("%s = %#v, want %#v", key, got, want)
					}
				}
			case "code_hash":
				t.Skip("checked by the evidence package")
			default:
				t.Fatalf("unknown kind %q", c.Kind)
			}
		})
	}
}

// valueJSON renders a decoded Value the way the corpus writes it.
func valueJSON(v Value) any {
	if v.Null {
		return nil
	}
	if !v.IsList {
		return v.Str
	}
	out := make([]any, len(v.List))
	for i, s := range v.List {
		if s != nil {
			out[i] = *s
		}
	}
	return out
}
