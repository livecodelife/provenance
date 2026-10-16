package record

import (
	"encoding/json"
	"os"
	"testing"
)

// The version_hash cases of the shared corpus (docs/HASHING.md §6): a version
// record in its published form hashes from content and specHashes alone.
func TestVersionHashCorpus(t *testing.T) {
	data, err := os.ReadFile(corpusPath)
	if err != nil {
		t.Fatal(err)
	}
	var corpus struct {
		Cases []struct {
			ID        string           `json:"id"`
			Kind      string           `json:"kind"`
			Version   PublishedVersion `json:"version"`
			Canonical string           `json:"canonical"`
			Hash      string           `json:"hash"`
			Valid     bool             `json:"valid"`
			Error     bool             `json:"error"`
		} `json:"cases"`
	}
	if err := json.Unmarshal(data, &corpus); err != nil {
		t.Fatal(err)
	}
	n := 0
	for _, c := range corpus.Cases {
		if c.Kind != "version_hash" {
			continue
		}
		n++
		t.Run(c.ID, func(t *testing.T) {
			v := c.Version
			if c.Error {
				if _, err := v.Snapshot(); err == nil {
					t.Fatal("expected an error")
				}
				return
			}
			hashes := make([]*string, len(v.SpecHashes))
			for i := range v.SpecHashes {
				if v.SpecHashes[i] != "" {
					h := v.SpecHashes[i]
					hashes[i] = &h
				}
			}
			r := v.Content.Record()
			if got := ContentCanonical(r, hashes); got != c.Canonical {
				t.Fatalf("canonical\n got %s\nwant %s", got, c.Canonical)
			}
			if got := ContentHashWithSpecs(r, hashes); got != c.Hash {
				t.Fatalf("hash %s, want %s", got, c.Hash)
			}
			_, err := v.Snapshot()
			if c.Valid != (err == nil) {
				t.Fatalf("valid = %v, want %v (%v)", err == nil, c.Valid, err)
			}
		})
	}
	if n == 0 {
		t.Fatal("no version_hash cases in the corpus")
	}
}
