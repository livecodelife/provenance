package record

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

// realRepos are the repositories whose records must all decode, hash and
// round-trip (prov-2026-20c3360e's done criterion).
var realRepos = []string{"linespec", "sedum", "scion", "linespec-provenance-ui"}

// TestRealRepos runs only when PROVENANCE_REAL_REPOS names the workspace
// directory holding those repositories.
func TestRealRepos(t *testing.T) {
	ws := os.Getenv("PROVENANCE_REAL_REPOS")
	if ws == "" {
		t.Skip("set PROVENANCE_REAL_REPOS to the workspace directory to run")
	}
	for _, repo := range realRepos {
		t.Run(repo, func(t *testing.T) {
			root := filepath.Join(ws, repo)
			files, err := filepath.Glob(filepath.Join(root, "provenance", "*.yml"))
			if err != nil || len(files) == 0 {
				t.Fatalf("no records under %s (%v)", root, err)
			}
			var missingSpecs, records int
			resolve := FileSpecResolver(root)
			for _, f := range files {
				data, err := os.ReadFile(f)
				if err != nil {
					t.Fatal(err)
				}
				d, err := Parse(data)
				if err != nil {
					t.Errorf("%s: %v", f, err)
					continue
				}
				r := d.Record()
				if err := r.Validate(); err != nil {
					t.Errorf("%s: %v", f, err)
				}
				if h := r.GenHash(); len(h) != len(hashPrefix)+64 {
					t.Errorf("%s: bad gen_hash %q", f, h)
				}
				h, missing := r.ContentHash(resolve)
				if len(h) != len(hashPrefix)+64 {
					t.Errorf("%s: bad content_hash %q", f, h)
				}
				missingSpecs += len(missing)
				out, err := d.Encode(r)
				if err != nil {
					t.Errorf("%s: encode: %v", f, err)
				} else if !bytes.Equal(out, data) {
					t.Errorf("%s: unchanged save is not byte-identical", f)
				}
				// A status flip must leave both hashes alone and touch one line.
				flipped := d.Record()
				flipped.Status = StatusDeprecated
				out, err = d.Encode(flipped)
				if err != nil {
					t.Errorf("%s: encode status flip: %v", f, err)
					continue
				}
				back, err := Decode(out)
				if err != nil {
					t.Errorf("%s: re-decode after status flip: %v", f, err)
					continue
				}
				if back.Status != StatusDeprecated || back.GenHash() != r.GenHash() {
					t.Errorf("%s: status flip changed more than status", f)
				}
				if h2, _ := back.ContentHash(resolve); h2 != h {
					t.Errorf("%s: status flip changed content_hash", f)
				}
				records++
			}
			t.Logf("%s: %d records decoded, hashed and round-tripped; %d unresolvable spec paths (null)", repo, records, missingSpecs)
		})
	}
}
