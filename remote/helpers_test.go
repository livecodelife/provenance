package remote

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/livecodelife/provenance/record"
	"github.com/livecodelife/provenance/remote/remotetest"
)

const testID = "prov-2026-0000a001"

// newClient returns a client for the fake with a fixed bearer token.
func newClient(t *testing.T) (*Client, *remotetest.Server) {
	t.Helper()
	srv := remotetest.New(t)
	return &Client{Base: srv.URL(), Auth: func(ctx context.Context, fn func(string) error) error { return fn("tok") }}, srv
}

// localStore is a files store in a temp dir holding testID at version n.
func localStore(t *testing.T, n int) *record.FilesStore {
	t.Helper()
	root := t.TempDir()
	if err := os.MkdirAll(filepath.Join(root, "provenance"), 0o755); err != nil {
		t.Fatal(err)
	}
	st := &record.FilesStore{Root: root, Dir: "provenance"}
	r := &record.Record{Schema: 2, ID: testID, Title: "Thing", Type: record.TypeImprint, Status: record.StatusOpen, Intent: "Do the thing.", AffectedScope: []string{"a.go"}}
	if err := st.Save(r, nil); err != nil {
		t.Fatal(err)
	}
	for v := 1; v <= n; v++ {
		r.Intent = "Do the thing, take " + string(rune('0'+v)) + "."
		var doc *record.Document
		reason := ""
		if v > 1 {
			reason = "change " + string(rune('0'+v))
		}
		if _, err := record.Commit(st, r, doc, record.Bump{Reason: reason, Author: "t@example.com", Now: time.Date(2026, 10, 16, 0, 0, v, 0, time.UTC)}); err != nil {
			t.Fatalf("commit %d: %v", v, err)
		}
		r, _, _ = st.Load(testID)
	}
	return st
}

// commitMore bumps r (already edited) to the next version with a reason.
func commitMore(st *record.FilesStore, r *record.Record, doc *record.Document) (*record.Snapshot, error) {
	return record.Commit(st, r, doc, record.Bump{Reason: "another change", Author: "t@example.com", Now: time.Date(2026, 10, 16, 0, 5, 0, 0, time.UTC)})
}
