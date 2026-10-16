package remote

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/livecodelife/provenance/record"
)

var _ record.Store = (*Store)(nil)

func publishAll(t *testing.T, c *Client, st record.Store) {
	t.Helper()
	plan, err := PlanPublish(context.Background(), st, []string{testID}, c)
	if err != nil {
		t.Fatal(err)
	}
	if err := ApplyPublish(context.Background(), plan, c, "", nil, nil); err != nil {
		t.Fatal(err)
	}
}

func TestCloneReproducesHashes(t *testing.T) {
	c, _ := newClient(t)
	src := localStore(t, 3)
	publishAll(t, c, src)
	dst := &record.FilesStore{Root: t.TempDir(), Dir: "provenance"}
	n, err := CloneRecords(context.Background(), c, dst, nil)
	if err != nil || n != 1 {
		t.Fatalf("clone: %d %v", n, err)
	}
	vs, _ := dst.Versions(testID)
	if len(vs) != 3 {
		t.Fatalf("versions = %v", vs)
	}
	for _, v := range vs {
		a, _ := src.Snapshot(testID, v)
		b, err := dst.Snapshot(testID, v)
		if err != nil || a.ContentHash != b.ContentHash || a.ParentHash != b.ParentHash || a.Reason != b.Reason || a.Author != b.Author {
			t.Fatalf("version %d differs: %+v vs %+v (%v)", v, a, b, err)
		}
	}
	r, _, err := dst.Load(testID)
	if err != nil || r.Version != 3 || r.Status != record.StatusOpen || r.Title != "Thing" {
		t.Fatalf("record = %+v %v", r, err)
	}
}

func TestStoreReadsAndAppends(t *testing.T) {
	c, srv := newClient(t)
	publishAll(t, c, localStore(t, 2))
	rs := &Store{Client: c}
	ids, err := rs.List()
	if err != nil || len(ids) != 1 || ids[0] != testID {
		t.Fatalf("list = %v %v", ids, err)
	}
	vs, _ := rs.Versions(testID)
	if len(vs) != 2 {
		t.Fatalf("versions = %v", vs)
	}
	snap, err := rs.Snapshot(testID, 2)
	if err != nil || snap.Version != 2 {
		t.Fatalf("snapshot: %+v %v", snap, err)
	}
	r, _, _ := rs.Load(testID)
	r.Intent = "A third, written through the remote store."
	next := record.NewSnapshot(r, 3, snap.ContentHash, "via remote store", "", "t@example.com", "2026-10-16T00:00:03Z", rs.Specs())
	if err := rs.Append(next); err != nil || srv.Versions(testID) != 3 {
		t.Fatalf("append: %v versions=%d", err, srv.Versions(testID))
	}
	if rs.Save(r, nil) != ErrUnsupported || rs.Remove(testID) != ErrUnsupported {
		t.Fatal("Save and Remove must be unsupported")
	}
}

// uiVersion writes version n+1 through the API as a UI edit would.
func uiVersion(t *testing.T, c *Client, st record.Store, intent string) {
	t.Helper()
	ctx := context.Background()
	head, err := c.GetRecord(ctx, testID)
	if err != nil {
		t.Fatal(err)
	}
	r, _, _ := st.Load(testID)
	r.Intent = intent
	snap := record.NewSnapshot(r, head.Version+1, head.ContentHash, "edited in the UI", "", "did:plc:test", "2026-10-16T01:00:00Z", st.Specs())
	req := VersionRequest(snap, &head.CID, "", "", "")
	req.Origin = "ui"
	req.BaseVersion = head.Version
	req.BaseContentHash = head.ContentHash
	if _, _, err := c.PutVersion(ctx, testID, req); err != nil {
		t.Fatal(err)
	}
}

func TestImportBringsUIVersion(t *testing.T) {
	c, _ := newClient(t)
	local := localStore(t, 2)
	publishAll(t, c, local)
	uiVersion(t, c, local, "Intent changed in the UI.")
	plan, err := PlanImport(context.Background(), local, []string{testID}, c)
	if err != nil || len(plan.Items) != 1 || len(plan.Items[0].Snaps) != 1 || plan.Items[0].Origins[0] != "ui" {
		t.Fatalf("plan = %+v %v", plan, err)
	}
	if err := ApplyImport(plan, local); err != nil {
		t.Fatal(err)
	}
	vs, _ := local.Versions(testID)
	if len(vs) != 3 {
		t.Fatalf("versions = %v", vs)
	}
	snap, _ := local.Snapshot(testID, 3)
	if snap.Reason != "edited in the UI" || snap.Author != "did:plc:test" {
		t.Fatalf("snapshot = %+v", snap)
	}
	r, _, _ := local.Load(testID)
	if r.Version != 3 || r.Intent != "Intent changed in the UI." || r.Reason != "edited in the UI" {
		t.Fatalf("record = %+v", r)
	}
	if again, _ := PlanImport(context.Background(), local, []string{testID}, c); len(again.Items) != 0 {
		t.Fatal("nothing left to import")
	}
}

func TestImportRefusesWhenLocalMovedPast(t *testing.T) {
	c, _ := newClient(t)
	local := localStore(t, 2)
	publishAll(t, c, local)
	uiVersion(t, c, local, "UI take.")
	// The local record also made its own version 3.
	r, doc, _ := local.Load(testID)
	r.Intent = "Local take."
	if _, err := commitMore(local, r, doc); err != nil {
		t.Fatal(err)
	}
	plan, err := PlanImport(context.Background(), local, []string{testID}, c)
	if err != nil || len(plan.Items) != 0 || len(plan.Refused) != 1 || !strings.Contains(plan.Refused[0], "revise") {
		t.Fatalf("plan = %+v %v", plan, err)
	}
}

func TestImportKeepsLocalRunCommandsAndDemotesImplemented(t *testing.T) {
	c, _ := newClient(t)
	local := localStore(t, 1)
	r, doc, _ := local.Load(testID)
	r.AssociatedSpecs = []record.AssociatedSpec{{Path: "spec.sh", RunCommand: "sh spec.sh"}}
	if err := os.WriteFile(filepath.Join(local.Root, "spec.sh"), []byte("true\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := commitMore(local, r, doc); err != nil {
		t.Fatal(err)
	}
	r, doc, _ = local.Load(testID)
	r.Status = record.StatusImplemented
	_ = local.Save(r, doc)
	publishAll(t, c, local)
	uiVersion(t, c, local, "UI take on an implemented record.")
	plan, err := PlanImport(context.Background(), local, []string{testID}, c)
	if err != nil {
		t.Fatal(err)
	}
	if err := ApplyImport(plan, local); err != nil {
		t.Fatal(err)
	}
	r, _, _ = local.Load(testID)
	if r.Status != record.StatusOpen || len(r.AssociatedSpecs) != 1 || r.AssociatedSpecs[0].RunCommand != "sh spec.sh" {
		t.Fatalf("record = %+v", r)
	}
}
