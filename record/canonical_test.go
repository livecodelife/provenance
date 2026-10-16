package record

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"
)

// Proof for prov-2026-32ee2643: the canonical-kind cases of the corpus
// (docs/HASHING.md section 5, booleans) pass, and a number still panics.

type canonicalCase struct {
	ID        string          `json:"id"`
	Kind      string          `json:"kind"`
	Value     json.RawMessage `json:"value"`
	Canonical string          `json:"canonical"`
	Hash      string          `json:"hash"`
	Error     bool            `json:"error"`
}

func canonicalCases(t *testing.T) []canonicalCase {
	t.Helper()
	data, err := os.ReadFile(corpusPath)
	if err != nil {
		t.Fatal(err)
	}
	var c struct {
		Version int             `json:"corpus_version"`
		Cases   []canonicalCase `json:"cases"`
	}
	if err := json.Unmarshal(data, &c); err != nil {
		t.Fatal(err)
	}
	if c.Version < 4 {
		t.Fatalf("corpus_version %d, want at least 4", c.Version)
	}
	var out []canonicalCase
	for _, k := range c.Cases {
		if k.Kind == "canonical" {
			out = append(out, k)
		}
	}
	if len(out) < 7 {
		t.Fatalf("%d canonical cases, want at least 7", len(out))
	}
	return out
}

// decodeValue yields the types Canonical accepts, with numbers kept as
// json.Number so they reach the encoder instead of being rounded away.
func decodeValue(t *testing.T, raw json.RawMessage) any {
	t.Helper()
	dec := json.NewDecoder(bytes.NewReader(raw))
	dec.UseNumber()
	var v any
	if err := dec.Decode(&v); err != nil {
		t.Fatal(err)
	}
	return v
}

func TestCanonicalCorpus(t *testing.T) {
	for _, c := range canonicalCases(t) {
		t.Run(c.ID, func(t *testing.T) {
			v := decodeValue(t, c.Value)
			if c.Error {
				defer func() {
					if recover() == nil {
						t.Errorf("Canonical(%s) did not panic", c.Value)
					}
				}()
				Canonical(v)
				return
			}
			if got := Canonical(v); got != c.Canonical {
				t.Errorf("canonical\n got %s\nwant %s", got, c.Canonical)
			}
			if got := hashBytes([]byte(c.Canonical)); got != c.Hash {
				t.Errorf("hash %s, want %s", got, c.Hash)
			}
		})
	}
}

func TestCanonicalBoolLiterals(t *testing.T) {
	for in, want := range map[bool]string{true: "true", false: "false"} {
		if got := Canonical(in); got != want {
			t.Errorf("Canonical(%v) = %s, want %s", in, got, want)
		}
	}
	if got := Canonical(map[string]any{"b": true, "a": []any{false, "true", nil}}); got != `{"a":[false,"true",null],"b":true}` {
		t.Errorf("mixed: %s", got)
	}
}
