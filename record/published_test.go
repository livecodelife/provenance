package record

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const (
	sampleDir  = "../lexicons/samples"
	sampleRec  = "at://did:plc:ragnj6bgwbfkhz4luybnthat/dev.provisio.provenance.record/prov-2026-0000a001"
	sampleVer  = "at://did:plc:ragnj6bgwbfkhz4luybnthat/dev.provisio.provenance.version/"
	sampleCIDs = "bafyreib2rxk3rybk3aobmv5cjuql3bm2twh4jo5uxgf5kpqpnxxcrwcyg4"
)

// sampleChain is a record with three versions, built the way the CLI builds
// them: v1, v2 changes the intent, v3 is a fix after a spec file changed.
func sampleChain() (snaps []*Snapshot, head *Record, resolve SpecResolver) {
	files := map[string]string{"specs/users.linespec": "TEST users\nEXPECT 200\n"}
	resolve = func(p string) (string, bool) {
		c, ok := files[p]
		if !ok {
			return "", false
		}
		return SpecHash([]byte(c)), true
	}
	r := &Record{
		Schema: 2, ID: "prov-2026-0000a001", Title: "Users index", Type: TypeImprint, Status: StatusOpen, Version: 1,
		CreatedAt: "2026-10-01", Author: "dev@example.com", Intent: "Add the users index action.",
		Constraints: []string{"Return JSON"}, AffectedScope: []string{"app/users.rb"},
		Implements: "prov-2026-0000b001", ImplementsRepo: "shared", ImplementsVersion: 1,
		Related: []string{"prov-2026-0000c001"}, Tags: []string{"api"}, Aliases: []string{"prov-2026-0000d001"},
		AssociatedSpecs: []AssociatedSpec{
			{Path: "specs/users.linespec", Type: "linespec", RunCommand: "linespec run"},
			{Path: "provenance/prov-2026-0000e001.yml"},
			{Type: "linespec"},
		},
	}
	add := func(reason, kind, at string) {
		parent := ""
		if n := len(snaps); n > 0 {
			parent = snaps[n-1].ContentHash
		}
		snaps = append(snaps, NewSnapshot(r, r.Version, parent, reason, kind, "dev@example.com", at, resolve))
	}
	add("", "", "2026-10-01T09:00:00Z")
	r.Version, r.Intent, r.Status = 2, "Add the users index action with a total count.", StatusImplemented
	add("Clients need the total count", "", "2026-10-02T09:00:00Z")
	files["specs/users.linespec"] = "TEST users\nEXPECT 200\nEXPECT total\n"
	r.Version = 3
	add("The spec now checks the total", ReasonKindFix, "2026-10-03T09:00:00Z")
	r.ImplementsVersion = 2 // repin: moves the head only
	return snaps, r, resolve
}

func sampleJSON(t *testing.T) map[string][]byte {
	t.Helper()
	snaps, head, resolve := sampleChain()
	out := map[string][]byte{}
	enc := func(name string, v any) {
		b, err := json.MarshalIndent(v, "", "  ")
		if err != nil {
			t.Fatal(err)
		}
		out[name] = append(b, '\n')
	}
	tids := []string{"3kqzh3m5vxd2a", "3kqzh3m5vyc2b", "3kqzh3m5vzb2c"}
	var parent *StrongRef
	for i, s := range snaps {
		enc("version-"+string(rune('1'+i))+".json", PublishedVersionOf(s, PublishContext{RecordURI: sampleRec, Parent: parent, Commit: strings.Repeat(string(rune('a'+i)), 40)}))
		parent = &StrongRef{URI: sampleVer + tids[i], CID: sampleCIDs}
	}
	enc("head.json", PublishedHeadOf(head, resolve, *parent,
		&PublishedSource{Repo: "https://github.com/example/app", Path: "provenance/prov-2026-0000a001.yml", Commit: strings.Repeat("c", 40)},
		"at://did:plc:ragnj6bgwbfkhz4luybnthat/dev.provisio.project/3kqzh3m5v2k2d", "2026-10-03T09:05:00Z"))
	return out
}

// The samples under lexicons/samples are what this code produces; set
// UPDATE_SAMPLES=1 to rewrite them.
func TestPublishedSamples(t *testing.T) {
	for name, want := range sampleJSON(t) {
		path := filepath.Join(sampleDir, name)
		if os.Getenv("UPDATE_SAMPLES") != "" {
			if err := os.WriteFile(path, want, 0o644); err != nil {
				t.Fatal(err)
			}
			continue
		}
		got, err := os.ReadFile(path)
		if err != nil || !bytes.Equal(got, want) {
			t.Errorf("%s differs from the mapping's output (%v); run with UPDATE_SAMPLES=1 after checking the change", name, err)
		}
	}
}

// Every sample version parses, verifies, and maps back to the snapshot it came from.
func TestPublishedVersionRoundTrip(t *testing.T) {
	snaps, _, _ := sampleChain()
	for i, s := range snaps {
		data, err := os.ReadFile(filepath.Join(sampleDir, "version-"+string(rune('1'+i))+".json"))
		if err != nil {
			t.Fatal(err)
		}
		_, back, err := ParsePublishedVersion(data)
		if err != nil {
			t.Fatalf("v%d: %v", i+1, err)
		}
		a, _ := MarshalSnapshot(s)
		b, _ := MarshalSnapshot(back)
		if !bytes.Equal(a, b) {
			t.Errorf("v%d does not round trip:\n%s\n%s", i+1, a, b)
		}
	}
}

// An unresolved spec hash is "" in the published form and nil in the snapshot.
func TestPublishedUnresolvedSpecHash(t *testing.T) {
	snaps, _, _ := sampleChain()
	v := PublishedVersionOf(snaps[0], PublishContext{})
	if len(v.SpecHashes) != 3 || v.SpecHashes[1] != RecordFileMarker || v.SpecHashes[2] != "" {
		t.Fatalf("spec hashes %q", v.SpecHashes)
	}
	s, err := v.Snapshot()
	if err != nil || s.SpecHashes[2] != nil || s.SpecHashes[1] == nil || *s.SpecHashes[1] != RecordFileMarker {
		t.Fatalf("back: %v %+v", err, s)
	}
}

// Content that does not hash to the stated contentHash is rejected.
func TestPublishedRejectsTampering(t *testing.T) {
	snaps, _, _ := sampleChain()
	for name, change := range map[string]func(*PublishedVersion){
		"intent":       func(v *PublishedVersion) { v.Content.Intent += " (changed)" },
		"scope":        func(v *PublishedVersion) { v.Content.AffectedScope = append(v.Content.AffectedScope, "x.rb") },
		"spec hash":    func(v *PublishedVersion) { v.SpecHashes[0] = SpecHash([]byte("other")) },
		"spec count":   func(v *PublishedVersion) { v.SpecHashes = v.SpecHashes[:1] },
		"content hash": func(v *PublishedVersion) { v.ContentHash = SpecHash([]byte("x")) },
	} {
		v := PublishedVersionOf(snaps[0], PublishContext{})
		change(v)
		if _, err := v.Snapshot(); err == nil {
			t.Errorf("%s: tampered version verified", name)
		}
	}
	// Fields the hash does not cover may change freely.
	v := PublishedVersionOf(snaps[0], PublishContext{})
	v.Content.Title, v.Content.Tags, v.Content.ImplementsVersion = "Renamed", nil, 9
	if _, err := v.Snapshot(); err != nil {
		t.Errorf("a label change must not break the hash: %v", err)
	}
}

// Every snapshot in this repository survives the published form unchanged.
func TestPublishedAllRepoSnapshots(t *testing.T) {
	store := &FilesStore{Root: ".."}
	ids, err := store.List()
	if err != nil || len(ids) == 0 {
		t.Skipf("no records (%v)", err)
	}
	n := 0
	for _, id := range ids {
		vs, _ := store.Versions(id)
		for _, ver := range vs {
			s, err := store.Snapshot(id, ver)
			if err != nil {
				t.Fatal(err)
			}
			data, err := json.Marshal(PublishedVersionOf(s, PublishContext{RecordURI: sampleRec}))
			if err != nil {
				t.Fatal(err)
			}
			_, back, err := ParsePublishedVersion(data)
			if err != nil {
				t.Fatalf("%s@%d: %v", id, ver, err)
			}
			a, _ := MarshalSnapshot(s)
			b, _ := MarshalSnapshot(back)
			if !bytes.Equal(a, b) {
				t.Fatalf("%s@%d does not round trip", id, ver)
			}
			n++
		}
	}
	t.Logf("%d snapshots round-tripped", n)
}
