package remote

import (
	"context"
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	"github.com/livecodelife/provenance/record"
)

// localStoreWithShape is localStore(t, 1) plus a second version that adds a
// shape (as graft's skeleton import does).
func localStoreWithShape(t *testing.T) (*record.FilesStore, []any) {
	t.Helper()
	st := localStore(t, 1)
	r, doc, err := st.Load(testID)
	if err != nil {
		t.Fatal(err)
	}
	shape := []any{
		map[string]any{"path": "running.go", "imports": []any{"sync"},
			"decls": []any{map[string]any{"kind": "type", "name": "T", "code": "type T struct{ N int }"}},
			"stubs": []any{map[string]any{"name": "Post", "receiver_name": "l", "receiver_type": "*L",
				"params": []any{map[string]any{"name": "d", "type": "int64"}}, "results": []any{"error"}, "doc": "Post changes a balance."}}},
		map[string]any{"path": "a.go", "stubs": []any{map[string]any{"name": "init"}}},
	}
	r.Shape = shape
	if _, err := commitMore(st, r, doc); err != nil {
		t.Fatal(err)
	}
	return st, shape
}

// A record with a shape publishes without hash_mismatch (the server
// recomputes genHash from the published content, shape included), and a clone
// has the identical shape and gen_hash.
func TestPublishAndCloneRoundTripShape(t *testing.T) {
	c, srv := newClient(t)
	src, _ := localStoreWithShape(t)
	publishAll(t, c, src)
	if srv.Versions(testID) != 2 {
		t.Fatalf("server has %d versions", srv.Versions(testID))
	}
	want, _, err := src.Load(testID)
	if err != nil {
		t.Fatal(err)
	}
	dst := &record.FilesStore{Root: t.TempDir(), Dir: "provenance"}
	if n, err := CloneRecords(context.Background(), c, dst, nil); err != nil || n != 1 {
		t.Fatalf("clone: %d %v", n, err)
	}
	got, _, err := dst.Load(testID)
	if err != nil {
		t.Fatal(err)
	}
	if got.GenHash() != want.GenHash() {
		t.Errorf("clone gen_hash %s, want %s", got.GenHash(), want.GenHash())
	}
	if record.GenCanonical(got.GenInput()) != record.GenCanonical(want.GenInput()) {
		t.Error("the cloned shape differs")
	}
	if !reflect.DeepEqual(got.Shape, want.Shape) && len(got.Shape) != 2 {
		t.Errorf("shape = %v", got.Shape)
	}
	// Every version's snapshot kept its shape, and verifies.
	for _, v := range []int{1, 2} {
		a, _ := src.Snapshot(testID, v)
		b, err := dst.Snapshot(testID, v)
		if err != nil || a.ContentHash != b.ContentHash || a.Content.GenHash() != b.Content.GenHash() {
			t.Fatalf("version %d: %v", v, err)
		}
	}
	// The head the server holds says the same.
	head, err := c.GetRecord(context.Background(), testID)
	if err != nil || head.GenHash != want.GenHash() || len(head.Content.Shape) != 2 {
		t.Fatalf("head: %+v %v", head, err)
	}
}

// A record without a shape is untouched by all of this: its request body has
// no shape key and the server still agrees on the hash.
func TestPublishWithoutShapeSendsNoShape(t *testing.T) {
	c, srv := newClient(t)
	src := localStore(t, 2)
	publishAll(t, c, src)
	snap, _ := src.Snapshot(testID, 2)
	b, err := jsonOf(VersionRequest(snap, nil, "open", "", ""))
	if err != nil || contains(b, "shape") {
		t.Fatalf("request mentions shape: %s %v", b, err)
	}
	if srv.Versions(testID) != 2 {
		t.Fatalf("server has %d versions", srv.Versions(testID))
	}
}

func jsonOf(v any) (string, error) {
	b, err := json.Marshal(v)
	return string(b), err
}

func contains(s, sub string) bool { return strings.Contains(s, sub) }

// A shape the lexicon cannot carry unchanged is a plan problem naming the
// record, and nothing is sent for it (the server would answer hash_mismatch).
func TestPublishPlanRefusesAShapeTheLexiconCannotCarry(t *testing.T) {
	c, srv := newClient(t)
	st := localStore(t, 1)
	r, doc, err := st.Load(testID)
	if err != nil {
		t.Fatal(err)
	}
	r.Shape = []any{map[string]any{"path": "a.go", "package": ""}} // valid keys, but an empty value the lexicon omits
	if _, err := commitMore(st, r, doc); err != nil {
		t.Fatal(err)
	}
	plan, err := PlanPublish(context.Background(), st, []string{testID}, c)
	if err != nil || len(plan.Items) != 0 || len(plan.Problems) != 1 || plan.Problems[0].Kind != "shape not publishable" || !strings.Contains(plan.Problems[0].Reason, testID) {
		t.Fatalf("plan = %+v %v", plan, err)
	}
	if srv.Versions(testID) != 0 {
		t.Error("something was published")
	}
}
