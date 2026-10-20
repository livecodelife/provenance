package main

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/livecodelife/provenance/record"
)

const createParentURI = "at://did:plc:ragnj6bgwbfkhz4luybnthat/dev.provisio.provenance.record/prov-2026-0000b001"

// A parent in another account is named by its at-uri: the record id goes to
// implements, the uri to implements_uri, the pin stays 0 and the user is told.
func TestCreateImplementsAtURI(t *testing.T) {
	root, cfg := v2Repo(t, false, nil)
	out := mustRun(t, "-c", cfg, "create", "--no-edit", "--title", "Child", "--type", "imprint", "--implements", createParentURI)
	id := created(t, out)
	rec, _, err := (&record.FilesStore{Root: root}).Load(id)
	if err != nil {
		t.Fatal(err)
	}
	if rec.Implements != "prov-2026-0000b001" || rec.ImplementsURI != createParentURI || rec.ImplementsVersion != 0 {
		t.Fatalf("implements = %q, uri = %q, version = %d", rec.Implements, rec.ImplementsURI, rec.ImplementsVersion)
	}
	if !strings.Contains(out, "unpinned") || !strings.Contains(out, "provenance repin") {
		t.Fatalf("no unpinned hint:\n%s", out)
	}
}

// A malformed at-uri fails at create time and writes nothing.
func TestCreateImplementsBadAtURIIsRefused(t *testing.T) {
	root, cfg := v2Repo(t, false, nil)
	for _, bad := range []string{
		"at://did:plc:ragnj6bgwbfkhz4luybnthat/dev.provisio.provenance.version/prov-2026-0000b001",
		"at://not-a-did/dev.provisio.provenance.record/prov-2026-0000b001",
		"at://did:plc:ragnj6bgwbfkhz4luybnthat/dev.provisio.provenance.record/",
	} {
		out, err := run(t, "-c", cfg, "create", "--no-edit", "--implements", bad)
		if err == nil || !strings.Contains(out+err.Error(), "invalid --implements") {
			t.Errorf("%s: accepted or unclear (%v)\n%s", bad, err, out)
		}
	}
	if files, _ := filepath.Glob(filepath.Join(root, "provenance", "*.yml")); len(files) != 0 {
		t.Fatalf("refused creates wrote %v", files)
	}
}
