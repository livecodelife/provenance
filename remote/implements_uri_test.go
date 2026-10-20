package remote

import (
	"context"
	"testing"

	"github.com/livecodelife/provenance/record"
)

const testParentURI = "at://did:plc:ragnj6bgwbfkhz4luybnthat/dev.provisio.provenance.record/prov-2026-0000b001"

// A cross-account parent link travels with the record: publish sends it in
// every version's content, and clone and import bring it back.
func TestImplementsURITravelsThroughPublishAndClone(t *testing.T) {
	c, _ := newClient(t)
	src := localStore(t, 1)
	r, doc, _ := src.Load(testID)
	r.Implements, r.ImplementsURI, r.ImplementsVersion = "prov-2026-0000b001", testParentURI, 2
	r.Intent = "Do the thing with a parent elsewhere."
	if _, err := record.Commit(src, r, doc, record.Bump{Reason: "link the parent", Author: "t@example.com"}); err != nil {
		t.Fatal(err)
	}
	publishAll(t, c, src)

	ctx := context.Background()
	vs, err := c.ListVersions(ctx, testID)
	if err != nil || len(vs) != 2 {
		t.Fatalf("versions = %d %v", len(vs), err)
	}
	if vs[0].Content.ImplementsURI != "" || vs[1].Content.ImplementsURI != testParentURI {
		t.Fatalf("published implementsUri: v1 %q, v2 %q", vs[0].Content.ImplementsURI, vs[1].Content.ImplementsURI)
	}

	dst := &record.FilesStore{Root: t.TempDir(), Dir: "provenance"}
	if _, err := CloneRecords(ctx, c, dst, nil); err != nil {
		t.Fatal(err)
	}
	got, _, err := dst.Load(testID)
	if err != nil || got.ImplementsURI != testParentURI {
		t.Fatalf("cloned implements_uri = %q (%v)", got.ImplementsURI, err)
	}
	snap, _ := dst.Snapshot(testID, 2)
	want, _ := src.Snapshot(testID, 2)
	if snap.ContentHash != want.ContentHash || snap.Content.ImplementsURI != testParentURI {
		t.Fatalf("cloned snapshot differs: %+v", snap)
	}
}
