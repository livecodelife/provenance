package record

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"testing"
)

// Proof for the integer extension of Canonical (docs/HASHING.md section 5):
// int, int64 and uint64 are written as bare base-10 digits; floats stay
// rejected. Corpus kind "canonical_int" carries an integer as {"$int": text}
// so no language has to round-trip a 64-bit value through a float.

type intCase struct {
	ID        string          `json:"id"`
	Kind      string          `json:"kind"`
	Value     json.RawMessage `json:"value"`
	Canonical string          `json:"canonical"`
	Hash      string          `json:"hash"`
	Error     bool            `json:"error"`
}

// fromCorpus rebuilds the value, turning {"$int": "<text>"} into int64 or
// uint64. Text that is not an integer in range is an error.
func fromCorpus(v any) (any, error) {
	switch v := v.(type) {
	case map[string]any:
		if t, ok := v["$int"]; ok && len(v) == 1 {
			text, _ := t.(string)
			if n, err := strconv.ParseInt(text, 10, 64); err == nil {
				return n, nil
			}
			if u, err := strconv.ParseUint(text, 10, 64); err == nil {
				return u, nil
			}
			return nil, fmt.Errorf("%q is not an integer in range", text)
		}
		for k, x := range v {
			y, err := fromCorpus(x)
			if err != nil {
				return nil, err
			}
			v[k] = y
		}
	case []any:
		for i, x := range v {
			y, err := fromCorpus(x)
			if err != nil {
				return nil, err
			}
			v[i] = y
		}
	}
	return v, nil
}

func TestCanonicalIntCorpus(t *testing.T) {
	data, err := os.ReadFile(corpusPath)
	if err != nil {
		t.Fatal(err)
	}
	var c struct {
		Version int       `json:"corpus_version"`
		Cases   []intCase `json:"cases"`
	}
	if err := json.Unmarshal(data, &c); err != nil {
		t.Fatal(err)
	}
	if c.Version < 5 {
		t.Fatalf("corpus_version %d, want at least 5", c.Version)
	}
	n := 0
	for _, k := range c.Cases {
		if k.Kind != "canonical_int" {
			continue
		}
		n++
		t.Run(k.ID, func(t *testing.T) {
			var raw any
			if err := json.Unmarshal(k.Value, &raw); err != nil {
				t.Fatal(err)
			}
			v, err := fromCorpus(raw)
			if k.Error {
				if err != nil {
					return
				}
				defer func() {
					if recover() == nil {
						t.Errorf("Canonical(%s) did not fail", k.Value)
					}
				}()
				Canonical(v)
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			if got := Canonical(v); got != k.Canonical {
				t.Errorf("canonical\n got %s\nwant %s", got, k.Canonical)
			}
			if got := hashBytes([]byte(k.Canonical)); got != k.Hash {
				t.Errorf("hash %s, want %s", got, k.Hash)
			}
		})
	}
	if n < 8 {
		t.Fatalf("%d canonical_int cases, want at least 8", n)
	}
}

func TestCanonicalIntTypes(t *testing.T) {
	for _, tc := range []struct {
		in   any
		want string
	}{
		{0, "0"},
		{-7, "-7"},
		{int64(math.MinInt64), "-9223372036854775808"},
		{int64(math.MaxInt64), "9223372036854775807"},
		{uint64(math.MaxUint64), "18446744073709551615"},
		{map[string]any{"n": 3, "ok": true}, `{"n":3,"ok":true}`},
	} {
		if got := Canonical(tc.in); got != tc.want {
			t.Errorf("Canonical(%v) = %s, want %s", tc.in, got, tc.want)
		}
	}
}

func TestCanonicalRejectsFloats(t *testing.T) {
	for _, v := range []any{1.5, float32(2), float64(3), json.Number("1")} {
		func() {
			defer func() {
				if r := recover(); r == nil || !strings.Contains(fmt.Sprint(r), "non-canonical") {
					t.Errorf("Canonical(%T) recovered %v, want a non-canonical panic", v, r)
				}
			}()
			Canonical(v)
		}()
	}
}
