package record

import (
	"strings"
	"testing"
)

func TestManifestBuildAndMarshal(t *testing.T) {
	s, r := newStore(t)
	draft := &Record{Schema: 2, ID: "prov-2026-00000001", Title: "D", Status: StatusDraft, Version: 1, Intent: "x"}
	if err := s.Save(draft, nil); err != nil {
		t.Fatal(err)
	}
	v1, err := Commit(s, r, nil, Bump{Now: t0})
	if err != nil {
		t.Fatal(err)
	}
	lines, err := BuildManifest(s)
	if err != nil {
		t.Fatal(err)
	}
	// The draft has no snapshot, so it has no line.
	if len(lines) != 1 || lines[0] != (ManifestLine{ID: r.ID, Version: 1, ContentHash: v1.ContentHash}) {
		t.Fatalf("lines %+v", lines)
	}
	// Byte format of prov-cli's writeManifest: compact json.Marshal of
	// {id, version, content_hash}, one per line, newline-terminated.
	want := `{"id":"prov-2026-0badc0de","version":1,"content_hash":"` + v1.ContentHash + `"}` + "\n"
	if got := string(MarshalManifest(lines)); got != want {
		t.Fatalf("got %q\nwant %q", got, want)
	}
	if len(MarshalManifest(nil)) != 0 {
		t.Fatal("empty manifest should be an empty file")
	}
	back, err := ParseManifest(MarshalManifest(lines))
	if err != nil || len(back) != 1 || back[0] != lines[0] {
		t.Fatalf("round trip: %+v %v", back, err)
	}
}

func TestParseManifestStrict(t *testing.T) {
	ok := `{"id":"prov-2026-00000001","version":1,"content_hash":"sha256:aa"}` + "\n" +
		`{"id":"prov-2026-00000002","version":3,"content_hash":"sha256:bb"}` + "\n"
	if _, err := ParseManifest([]byte(ok)); err != nil {
		t.Fatal(err)
	}
	for name, bad := range map[string]string{
		"no trailing newline": strings.TrimSuffix(ok, "\n"),
		"unsorted":            `{"id":"prov-2026-00000002","version":1,"content_hash":"sha256:bb"}` + "\n" + `{"id":"prov-2026-00000001","version":1,"content_hash":"sha256:aa"}` + "\n",
		"duplicate":           `{"id":"prov-2026-00000001","version":1,"content_hash":"sha256:aa"}` + "\n" + `{"id":"prov-2026-00000001","version":2,"content_hash":"sha256:aa"}` + "\n",
		"unknown key":         `{"id":"prov-2026-00000001","version":1,"content_hash":"sha256:aa","x":1}` + "\n",
		"not canonical":       `{"version":1,"id":"prov-2026-00000001","content_hash":"sha256:aa"}` + "\n",
		"blank line":          ok + "\n",
		"not json":            "nope\n",
	} {
		if _, err := ParseManifest([]byte(bad)); err == nil {
			t.Errorf("%s: accepted", name)
		}
	}
}
