package record

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

// ManifestPath is where the version manifest lives, relative to the repo root.
const ManifestPath = ".provenance/versions.jsonl"

// ManifestLine is one record's entry in the version manifest: the latest
// snapshot's version and content hash. The manifest replaces v1's hash
// manifest; it keeps bumps visible in git diffs and gives hooks a fast check.
type ManifestLine struct {
	ID          string `json:"id"`
	Version     int    `json:"version"`
	ContentHash string `json:"content_hash"`
}

// BuildManifest lists the latest snapshot of every record in store that has
// one, sorted by id.
func BuildManifest(store Store) ([]ManifestLine, error) {
	ids, err := store.List() // sorted
	if err != nil {
		return nil, err
	}
	var out []ManifestLine
	for _, id := range ids {
		vs, err := store.Versions(id)
		if err != nil {
			return nil, err
		}
		if len(vs) == 0 {
			continue
		}
		snap, err := store.Snapshot(id, vs[len(vs)-1])
		if err != nil {
			return nil, err
		}
		out = append(out, ManifestLine{ID: id, Version: snap.Version, ContentHash: snap.ContentHash})
	}
	return out, nil
}

// MarshalManifest encodes lines as the manifest file: one compact JSON
// object per line ({"id","version","content_hash"} in that order), each
// ending in a newline. No lines is an empty file.
func MarshalManifest(lines []ManifestLine) []byte {
	var b bytes.Buffer
	for _, l := range lines {
		data, _ := json.Marshal(l) // cannot fail for this struct
		b.Write(data)
		b.WriteByte('\n')
	}
	return b.Bytes()
}

// ParseManifest decodes a manifest strictly: every line must be exactly how
// MarshalManifest writes it, ids strictly ascending.
func ParseManifest(data []byte) ([]ManifestLine, error) {
	if len(data) == 0 {
		return nil, nil
	}
	if !bytes.HasSuffix(data, []byte("\n")) {
		return nil, fmt.Errorf("manifest does not end with a newline")
	}
	var out []ManifestLine
	for i, text := range strings.Split(strings.TrimSuffix(string(data), "\n"), "\n") {
		var l ManifestLine
		dec := json.NewDecoder(strings.NewReader(text))
		dec.DisallowUnknownFields()
		if err := dec.Decode(&l); err != nil {
			return nil, fmt.Errorf("manifest line %d: %w", i+1, err)
		}
		if canon, _ := json.Marshal(l); string(canon) != text {
			return nil, fmt.Errorf("manifest line %d is not canonical: want %s", i+1, canon)
		}
		if n := len(out); n > 0 && out[n-1].ID >= l.ID {
			return nil, fmt.Errorf("manifest line %d: %s is not after %s (ids must be sorted and unique)", i+1, l.ID, out[n-1].ID)
		}
		out = append(out, l)
	}
	return out, nil
}
