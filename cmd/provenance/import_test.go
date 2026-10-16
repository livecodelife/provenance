package main

import (
	"context"
	"strings"
	"testing"

	"github.com/livecodelife/provenance/record"
	"github.com/livecodelife/provenance/remote"
)

// uiEdit writes the next version of id to the server the way the UI does.
func uiEdit(t *testing.T, root, id, intent string) {
	t.Helper()
	c, err := apiClient()
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	head, err := c.GetRecord(ctx, id)
	if err != nil {
		t.Fatal(err)
	}
	st := &record.FilesStore{Root: root, Dir: "provenance"}
	r, _, _ := st.Load(id)
	r.Intent = intent
	snap := record.NewSnapshot(r, head.Version+1, head.ContentHash, "edited in the UI", "", "did:plc:test", "2026-10-16T01:00:00Z", st.Specs())
	req := remote.VersionRequest(snap, &head.CID, "", "", "")
	req.Origin, req.BaseVersion, req.BaseContentHash = "ui", head.Version, head.ContentHash
	if _, _, err := c.PutVersion(ctx, id, req); err != nil {
		t.Fatal(err)
	}
}

func TestImportBringsAUIVersion(t *testing.T) {
	remoteEnv(t)
	root, cfg := checkRepo(t)
	signedIn(t)
	if _, _, err := runCLI(t, "", "-c", cfg, "publish", "--all", "--yes"); err != nil {
		t.Fatal(err)
	}
	uiEdit(t, root, chkID, "Intent changed in the UI.")
	out, _, err := runCLI(t, "", "-c", cfg, "import", "--all")
	if err != nil || !strings.Contains(out, chkID+" v2") || !strings.Contains(out, "edited in the UI") {
		t.Fatalf("import: %v\n%s", err, out)
	}
	hist := mustRun(t, "-c", cfg, "history", chkID)
	if !strings.Contains(hist, "v2") || !strings.Contains(hist, "edited in the UI") || !strings.Contains(hist, "did:plc:test") {
		t.Fatalf("history:\n%s", hist)
	}
	if !strings.Contains(readFile(t, root+"/provenance/"+chkID+".yml"), "Intent changed in the UI.") {
		t.Fatal("the record file was not updated")
	}
	if !strings.Contains(manifestLines(t, root), `"version":2`) {
		t.Fatal("the manifest was not updated")
	}
	out, _, err = runCLI(t, "", "-c", cfg, "import", "--all")
	if err != nil || !strings.Contains(out, "Nothing to import") {
		t.Fatalf("second import: %v\n%s", err, out)
	}
}

func TestImportRefusesAMovedRecord(t *testing.T) {
	remoteEnv(t)
	root, cfg := checkRepo(t)
	signedIn(t)
	if _, _, err := runCLI(t, "", "-c", cfg, "publish", "--all", "--yes"); err != nil {
		t.Fatal(err)
	}
	uiEdit(t, root, chkID, "UI take.")
	mustRun(t, "-c", cfg, "edit", chkID, "--add-constraint", "Local take", "--reason", "Local change")
	_, _, err := runCLI(t, "", "-c", cfg, "import", chkID)
	if err == nil || !strings.Contains(err.Error(), "refused") {
		t.Fatalf("err = %v", err)
	}
	if strings.Contains(readFile(t, root+"/provenance/"+chkID+".yml"), "UI take.") {
		t.Fatal("a refused import must change nothing")
	}
}
