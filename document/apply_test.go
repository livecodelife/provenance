package document

import (
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/livecodelife/provenance/record"
)

const (
	aDraft1 = "prov-2026-0000a001"
	aDraft2 = "prov-2026-0000a002"
	aOpen1  = "prov-2026-0000a003"
	aOpen2  = "prov-2026-0000a004"
	aV1     = "prov-2026-0000a005" // schema 1 draft
)

var okOpts = ApplyOptions{Reason: "Why", Author: "me@example.com"}

// applyStore holds two drafts, two open records (version 1 snapshots) and a
// schema 1 draft. Intent "old <id>", constraints "old1", "old2".
func applyStore(t *testing.T) *record.FilesStore {
	t.Helper()
	s := newStore(t)
	mk := func(id, title string) rec {
		return rec{ID: id, Title: title, Type: "blueprint", Intent: "old " + id, Constraints: []string{"old1", "old2"},
			Scope: []string{"a.go"}, Implements: "prov-2026-0000b001"}
	}
	put(t, s, mk(aDraft1, "Draft one"))
	put(t, s, mk(aDraft2, "Draft two"))
	putVersioned(t, s, mk(aOpen1, "Open one"), 1)
	putVersioned(t, s, mk(aOpen2, "Open two"), 1)
	v1 := mk(aV1, "Schema one")
	v1.Schema1 = true
	put(t, s, v1)
	return s
}

func change(id string, ver int, intent string, cons []string) Change {
	return Change{ID: id, Title: "t", AnchoredVersion: ver, CurrentVersion: ver,
		IntentChanged: true, ConstraintsChanged: true, Intent: intent, Constraints: cons}
}

// TestErrorValues pins imprints prov-2026-2a640ada, prov-2026-9b7dc23c and
// prov-2026-9ca50ce0: the exact messages, and that the three can be told apart.
func TestErrorValues(t *testing.T) {
	for _, c := range []struct {
		err  error
		want string
	}{
		{ErrNoReason, "a reason is required"},
		{ErrConflicts, "the plan has conflicting sections"},
		{ErrStale, "the record changed since the plan was made"},
	} {
		if c.err == nil || c.err.Error() != c.want {
			t.Errorf("message %v, want %q", c.err, c.want)
		}
	}
	if errors.Is(ErrNoReason, ErrConflicts) || errors.Is(ErrConflicts, ErrStale) || errors.Is(ErrStale, ErrNoReason) {
		t.Fatal("the error values must be distinct")
	}
}

// TestApplyTypes pins imprints prov-2026-44fce5f4 (ApplyOptions) and
// prov-2026-658e3026 (Revised): the fields and their types.
func TestApplyTypes(t *testing.T) {
	_ = ApplyOptions{Reason: "", Author: "", SkipConflicts: false}
	_ = Revised{ID: "", Version: 0, Draft: false}
}

// TestApplyNilPlan pins imprint prov-2026-6b2ceecb: no revisions and an error
// with the message "nil plan".
func TestApplyNilPlan(t *testing.T) {
	got, err := Apply(applyStore(t), nil, okOpts)
	if got != nil || err == nil || err.Error() != "nil plan" {
		t.Fatalf("got %v, %v", got, err)
	}
}

// TestApplyRequiresReason pins imprint prov-2026-6b2ceecb and blueprint
// prov-2026-06e5c9b4: an empty or blank reason returns ErrNoReason itself and
// writes nothing.
func TestApplyRequiresReason(t *testing.T) {
	s := applyStore(t)
	before := tree(t, s.Root)
	plan := &Plan{Changes: []Change{change(aDraft1, 1, "new", []string{"n"}), change(aOpen1, 1, "new", []string{"n"})}}
	for _, reason := range []string{"", " ", "  \t\n "} {
		got, err := Apply(s, plan, ApplyOptions{Reason: reason, Author: "me@example.com"})
		if err != ErrNoReason || got != nil {
			t.Fatalf("reason %q: %v, %v", reason, got, err)
		}
	}
	sameTree(t, before, tree(t, s.Root))
}

// TestApplyConflicts pins imprint prov-2026-6b2ceecb and blueprint
// prov-2026-06e5c9b4: without SkipConflicts a plan with conflicts writes nothing
// and returns an error wrapping ErrConflicts whose message is the ErrConflicts
// text, a colon, a space, the count and the word conflicts.
func TestApplyConflicts(t *testing.T) {
	for _, c := range []struct {
		conflicts int
		msg       string
	}{
		{1, "the plan has conflicting sections: 1 conflicts"},
		{2, "the plan has conflicting sections: 2 conflicts"},
		{12, "the plan has conflicting sections: 12 conflicts"},
	} {
		s := applyStore(t)
		before := tree(t, s.Root)
		plan := &Plan{Changes: []Change{change(aDraft1, 1, "new", []string{"n"})}}
		for i := 0; i < c.conflicts; i++ {
			plan.Conflicts = append(plan.Conflicts, Problem{Kind: "conflict", ID: aOpen1, Line: 1, AnchoredVersion: 1, CurrentVersion: 2})
		}
		got, err := Apply(s, plan, okOpts)
		if got != nil || !errors.Is(err, ErrConflicts) || err.Error() != c.msg {
			t.Fatalf("%d conflicts: %v, %v", c.conflicts, got, err)
		}
		if errors.Is(err, ErrStale) || errors.Is(err, ErrNoReason) {
			t.Fatalf("error is also another sentinel: %v", err)
		}
		sameTree(t, before, tree(t, s.Root))
	}
}

// TestApplySkipConflicts pins imprint prov-2026-6b2ceecb and blueprint
// prov-2026-06e5c9b4: with SkipConflicts the conflicts are ignored and the
// other changes are applied; only the changes are touched.
func TestApplySkipConflicts(t *testing.T) {
	s := applyStore(t)
	conflicted := tree(t, s.Root)
	plan := &Plan{
		Changes:   []Change{change(aDraft1, 1, "new intent", []string{"n1"})},
		Conflicts: []Problem{{Kind: "conflict", ID: aOpen1, Line: 3, AnchoredVersion: 1, CurrentVersion: 2}},
	}
	opts := okOpts
	opts.SkipConflicts = true
	got, err := Apply(s, plan, opts)
	if err != nil || !reflect.DeepEqual(got, []Revised{{ID: aDraft1, Version: 1, Draft: true}}) {
		t.Fatalf("%v, %v", got, err)
	}
	if r := loadRec(t, s, aDraft1); r.Intent != "new intent" {
		t.Fatalf("intent %q", r.Intent)
	}
	for _, id := range []string{aOpen1, aOpen2, aDraft2} {
		p := filepath.Join("provenance", id+".yml")
		if tree(t, s.Root)[p] != conflicted[p] {
			t.Errorf("%s was touched", id)
		}
	}
	if !reflect.DeepEqual(versionsOf(t, s, aOpen1), []int{1}) {
		t.Fatal("the conflicting record got a version")
	}
}

// TestApplySkipConflictsOnlyConflicts pins imprint prov-2026-6b2ceecb: when no
// record is revised the returned slice is nil.
func TestApplySkipConflictsOnlyConflicts(t *testing.T) {
	s := applyStore(t)
	before := tree(t, s.Root)
	plan := &Plan{Conflicts: []Problem{{Kind: "conflict", ID: aOpen1}}}
	got, err := Apply(s, plan, ApplyOptions{Reason: "r", SkipConflicts: true})
	if err != nil || got != nil {
		t.Fatalf("%v, %v", got, err)
	}
	got, err = Apply(s, &Plan{}, okOpts)
	if err != nil || got != nil {
		t.Fatalf("empty plan: %v, %v", got, err)
	}
	sameTree(t, before, tree(t, s.Root))
}

// TestApplyIgnoresProblems pins imprint prov-2026-6b2ceecb: Apply does not look
// at the plan's Problems.
func TestApplyIgnoresProblems(t *testing.T) {
	s := applyStore(t)
	plan := &Plan{
		Changes:  []Change{change(aDraft1, 1, "new", []string{"n"})},
		Problems: []Problem{{Kind: "unknown", ID: "prov-2026-00000001", Line: 9, Message: "no record with this id"}, {Kind: "unanchored", Line: 2}},
	}
	got, err := Apply(s, plan, okOpts)
	if err != nil || len(got) != 1 {
		t.Fatalf("%v, %v", got, err)
	}
}

// TestApplyDraft pins imprint prov-2026-6b2ceecb: a record without versions is
// saved in place, no version or reason is created, Revised carries the current
// version and Draft true, and only intent and constraints change.
func TestApplyDraft(t *testing.T) {
	s := newStore(t)
	put(t, s, rec{ID: aDraft1, Title: "Keep title", Type: "blueprint", Version: 2, Reason: "old reason", Intent: "old",
		Constraints: []string{"c1"}, Scope: []string{"a.go", "b/**"}, Implements: "prov-2026-0000b001"})
	put(t, s, rec{ID: aDraft2, Title: "Untouched", Intent: "x"})
	untouched := tree(t, s.Root)[filepath.Join("provenance", aDraft2+".yml")]

	got, err := Apply(s, &Plan{Changes: []Change{change(aDraft1, 2, "new intent", []string{"n1", "n2\nmore"})}}, okOpts)
	if err != nil || !reflect.DeepEqual(got, []Revised{{ID: aDraft1, Version: 2, Draft: true}}) {
		t.Fatalf("%v, %v", got, err)
	}
	r := loadRec(t, s, aDraft1)
	if r.Intent != "new intent" || !reflect.DeepEqual(r.Constraints, []string{"n1", "n2\nmore"}) {
		t.Fatalf("content %q %q", r.Intent, r.Constraints)
	}
	if r.Version != 2 || r.Reason != "old reason" || r.Title != "Keep title" || r.Status != record.StatusDraft ||
		r.Type != record.TypeBlueprint || r.Implements != "prov-2026-0000b001" || !reflect.DeepEqual(r.AffectedScope, []string{"a.go", "b/**"}) {
		t.Fatalf("record %+v", r)
	}
	if vs := versionsOf(t, s, aDraft1); len(vs) != 0 {
		t.Fatalf("a draft gets no version: %v", vs)
	}
	if _, err := os.Stat(filepath.Join(s.Root, "provenance", ".versions", aDraft1)); err == nil {
		t.Fatal("snapshot directory created for a draft")
	}
	if tree(t, s.Root)[filepath.Join("provenance", aDraft2+".yml")] != untouched {
		t.Fatal("another record was touched")
	}
}

// TestApplyVersioned pins imprint prov-2026-6b2ceecb and blueprint
// prov-2026-06e5c9b4: a record with versions goes through record.Commit, which
// gives the next version with the caller's reason and author and new hashes.
func TestApplyVersioned(t *testing.T) {
	s := applyStore(t)
	got, err := Apply(s, &Plan{Changes: []Change{change(aOpen1, 1, "new intent", []string{"n1"})}}, okOpts)
	if err != nil || !reflect.DeepEqual(got, []Revised{{ID: aOpen1, Version: 2, Draft: false}}) {
		t.Fatalf("%v, %v", got, err)
	}
	r := loadRec(t, s, aOpen1)
	if r.Version != 2 || r.Reason != "Why" || r.Intent != "new intent" || !reflect.DeepEqual(r.Constraints, []string{"n1"}) {
		t.Fatalf("record %+v", r)
	}
	if r.Title != "Open one" || r.Implements != "prov-2026-0000b001" || !reflect.DeepEqual(r.AffectedScope, []string{"a.go"}) {
		t.Fatalf("other fields changed: %+v", r)
	}
	if !reflect.DeepEqual(versionsOf(t, s, aOpen1), []int{1, 2}) {
		t.Fatalf("versions %v", versionsOf(t, s, aOpen1))
	}
	snap, err := s.Snapshot(aOpen1, 2)
	if err != nil {
		t.Fatal(err)
	}
	hash, _ := r.ContentHash(s.Specs())
	if snap.Reason != "Why" || snap.Author != "me@example.com" || snap.Content.Intent != "new intent" || snap.ContentHash == "" || snap.ContentHash != hash {
		t.Fatalf("snapshot %+v", snap)
	}
	old, _ := s.Snapshot(aOpen1, 1)
	if old.ContentHash == snap.ContentHash || snap.ParentHash != old.ContentHash {
		t.Fatalf("hash chain: v1 %s v2 %s parent %s", old.ContentHash, snap.ContentHash, snap.ParentHash)
	}
}

// TestApplyFlagsDecideWhatIsReplaced pins blueprint prov-2026-06e5c9b4 and
// imprint prov-2026-6b2ceecb: intent only when IntentChanged, constraints only
// when ConstraintsChanged; the text is set as parsed, not normalized.
func TestApplyFlagsDecideWhatIsReplaced(t *testing.T) {
	s := applyStore(t)
	odd := "  keep   spacing\n\n  and blank lines  "
	oddCons := []string{"a  b\n\n  c", " lead"}
	plan := &Plan{Changes: []Change{
		{ID: aDraft1, CurrentVersion: 1, AnchoredVersion: 1, IntentChanged: true, Intent: odd, Constraints: []string{"must not be used"}},
		{ID: aDraft2, CurrentVersion: 1, AnchoredVersion: 1, ConstraintsChanged: true, Constraints: oddCons, Intent: "must not be used"},
	}}
	if _, err := Apply(s, plan, okOpts); err != nil {
		t.Fatal(err)
	}
	r1, r2 := loadRec(t, s, aDraft1), loadRec(t, s, aDraft2)
	if r1.Intent != odd || !reflect.DeepEqual(r1.Constraints, []string{"old1", "old2"}) {
		t.Fatalf("intent-only: %q %q", r1.Intent, r1.Constraints)
	}
	if r2.Intent != "old "+aDraft2 || !reflect.DeepEqual(r2.Constraints, oddCons) {
		t.Fatalf("constraints-only: %q %q", r2.Intent, r2.Constraints)
	}
}

// TestApplyOrderAndOnlyChanges pins imprint prov-2026-6b2ceecb: changes are
// processed and reported in plan order (drafts and versioned mixed), and no
// record outside the changes is touched.
func TestApplyOrderAndOnlyChanges(t *testing.T) {
	s := applyStore(t)
	before := tree(t, s.Root)
	plan := &Plan{Changes: []Change{
		change(aOpen2, 1, "n", []string{"n"}),
		change(aDraft2, 1, "n", []string{"n"}),
		change(aOpen1, 1, "n", []string{"n"}),
	}}
	got, err := Apply(s, plan, okOpts)
	want := []Revised{{aOpen2, 2, false}, {aDraft2, 1, true}, {aOpen1, 2, false}}
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Fatalf("%v, %v", got, err)
	}
	after := tree(t, s.Root)
	for _, id := range []string{aDraft1, aV1} {
		p := filepath.Join("provenance", id+".yml")
		if after[p] != before[p] {
			t.Errorf("%s was touched", id)
		}
	}
	if _, ok := after[filepath.Join(".provenance", "versions.jsonl")]; ok {
		t.Error("Apply must never write the versions manifest")
	}
	for k := range after {
		if _, ok := before[k]; !ok && !isUnder(k, "provenance/.versions/"+aOpen1) && !isUnder(k, "provenance/.versions/"+aOpen2) {
			t.Errorf("unexpected new file %s", k)
		}
	}
	ids, _ := s.List()
	if len(ids) != 5 {
		t.Fatalf("record count changed: %v", ids)
	}
}

func isUnder(path, dir string) bool {
	return len(path) >= len(dir) && path[:len(dir)] == dir
}

// TestApplyStale pins imprint prov-2026-6b2ceecb: a record whose version
// differs from the Change's CurrentVersion, lower or higher, stops Apply with
// an error wrapping ErrStale whose message starts with the id, before anything
// is written for that record.
func TestApplyStale(t *testing.T) {
	for _, cur := range []int{0, 2, 7} {
		s := applyStore(t)
		before := tree(t, s.Root)
		ch := change(aOpen1, 1, "new", []string{"n"})
		ch.CurrentVersion = cur
		got, err := Apply(s, &Plan{Changes: []Change{ch}}, okOpts)
		if got != nil || !errors.Is(err, ErrStale) || err.Error() != aOpen1+": the record changed since the plan was made" {
			t.Fatalf("current %d: %v, %v", cur, got, err)
		}
		sameTree(t, before, tree(t, s.Root))
	}
	// drafts too
	s := applyStore(t)
	ch := change(aDraft1, 1, "new", []string{"n"})
	ch.CurrentVersion = 2
	if _, err := Apply(s, &Plan{Changes: []Change{ch}}, okOpts); !errors.Is(err, ErrStale) {
		t.Fatalf("draft: %v", err)
	}
}

// TestApplyStopsAtFirstFailure pins imprint prov-2026-6b2ceecb and blueprint
// prov-2026-06e5c9b4: on the first failure Apply stops, returns the records
// revised so far together with the error, and leaves later changes untouched.
func TestApplyStopsAtFirstFailure(t *testing.T) {
	s := applyStore(t)
	before := tree(t, s.Root)
	stale := change(aOpen1, 1, "n", []string{"n"})
	stale.CurrentVersion = 9
	plan := &Plan{Changes: []Change{
		change(aDraft1, 1, "first", []string{"n"}),
		change(aOpen2, 1, "second", []string{"n"}),
		stale,
		change(aDraft2, 1, "never", []string{"n"}),
	}}
	got, err := Apply(s, plan, okOpts)
	want := []Revised{{aDraft1, 1, true}, {aOpen2, 2, false}}
	if !reflect.DeepEqual(got, want) || !errors.Is(err, ErrStale) || err.Error() != aOpen1+": "+ErrStale.Error() {
		t.Fatalf("%v, %v", got, err)
	}
	after := tree(t, s.Root)
	for _, id := range []string{aOpen1, aDraft2} {
		p := filepath.Join("provenance", id+".yml")
		if after[p] != before[p] {
			t.Errorf("%s was touched", id)
		}
	}
	if loadRec(t, s, aDraft1).Intent != "first" || loadRec(t, s, aOpen2).Version != 2 {
		t.Fatal("earlier revisions must stay saved")
	}
}

// TestApplySchema1Draft pins imprint prov-2026-6b2ceecb: a draft that is not a
// schema 2 record is refused with an error wrapping record.ErrNotV2 and saved
// nowhere.
func TestApplySchema1Draft(t *testing.T) {
	s := applyStore(t)
	before := tree(t, s.Root)
	got, err := Apply(s, &Plan{Changes: []Change{change(aV1, 1, "new", []string{"n"})}}, okOpts)
	if got != nil || !errors.Is(err, record.ErrNotV2) || err.Error() != aV1+": "+record.ErrNotV2.Error() {
		t.Fatalf("%v, %v", got, err)
	}
	sameTree(t, before, tree(t, s.Root))
}

// TestApplyWrapsStoreErrors pins imprint prov-2026-6b2ceecb: errors from Load,
// Versions, Save and Commit are wrapped, with the message "ID: " followed by the
// wrapped text; the revisions made so far are returned.
func TestApplyWrapsStoreErrors(t *testing.T) {
	boom := errors.New("boom")
	t.Run("load", func(t *testing.T) {
		s := applyStore(t)
		f := &faultStore{Store: s, loadErr: map[string]error{aDraft2: boom}}
		plan := &Plan{Changes: []Change{change(aDraft1, 1, "a", []string{"a"}), change(aDraft2, 1, "b", []string{"b"})}}
		got, err := Apply(f, plan, okOpts)
		if !errors.Is(err, boom) || err.Error() != aDraft2+": boom" || !reflect.DeepEqual(got, []Revised{{aDraft1, 1, true}}) {
			t.Fatalf("%v, %v", got, err)
		}
	})
	t.Run("missing record", func(t *testing.T) {
		s := applyStore(t)
		got, err := Apply(s, &Plan{Changes: []Change{change("prov-2026-00000001", 1, "a", []string{"a"})}}, okOpts)
		if got != nil || err == nil || len(err.Error()) < len("prov-2026-00000001: ") || err.Error()[:20] != "prov-2026-00000001: " {
			t.Fatalf("%v, %v", got, err)
		}
	})
	t.Run("versions", func(t *testing.T) {
		s := applyStore(t)
		f := &faultStore{Store: s, versionsErr: map[string]error{aDraft1: boom}}
		got, err := Apply(f, &Plan{Changes: []Change{change(aDraft1, 1, "a", []string{"a"})}}, okOpts)
		if !errors.Is(err, boom) || err.Error() != aDraft1+": boom" || got != nil {
			t.Fatalf("%v, %v", got, err)
		}
	})
	t.Run("save", func(t *testing.T) {
		s := applyStore(t)
		f := &faultStore{Store: s, saveErr: map[string]error{aDraft1: boom}}
		got, err := Apply(f, &Plan{Changes: []Change{change(aDraft1, 1, "a", []string{"a"})}}, okOpts)
		if !errors.Is(err, boom) || err.Error() != aDraft1+": boom" || got != nil {
			t.Fatalf("%v, %v", got, err)
		}
	})
	t.Run("commit", func(t *testing.T) {
		s := applyStore(t)
		// a record with versions whose file now says deprecated: Commit refuses it
		p := filepath.Join(s.Root, "provenance", aOpen1+".yml")
		data, _ := os.ReadFile(p)
		if err := os.WriteFile(p, []byte(replaceOnce(string(data), "status: open", "status: deprecated")), 0o644); err != nil {
			t.Fatal(err)
		}
		got, err := Apply(s, &Plan{Changes: []Change{change(aOpen1, 1, "a", []string{"a"})}}, okOpts)
		if got != nil || !errors.Is(err, record.ErrDeprecated) || err.Error() != aOpen1+": "+record.ErrDeprecated.Error() {
			t.Fatalf("%v, %v", got, err)
		}
	})
}

func replaceOnce(s, old, new string) string {
	for i := 0; i+len(old) <= len(s); i++ {
		if s[i:i+len(old)] == old {
			return s[:i] + new + s[i+len(old):]
		}
	}
	return s
}

// TestApplyNeverCreatesOrManifests pins blueprint prov-2026-06e5c9b4: Apply
// never creates a record and never writes the version manifest, even for a
// change that names a record the store does not have.
func TestApplyNeverCreatesOrManifests(t *testing.T) {
	s := applyStore(t)
	_, _ = Apply(s, &Plan{Changes: []Change{change("prov-2026-00000002", 1, "a", []string{"a"})}}, okOpts)
	if _, err := os.Stat(filepath.Join(s.Root, "provenance", "prov-2026-00000002.yml")); err == nil {
		t.Fatal("a record was created")
	}
	if _, err := os.Stat(filepath.Join(s.Root, ".provenance")); err == nil {
		t.Fatal("manifest directory created")
	}
}

// TestEndToEndRenderEditPlanApply pins the blueprints together: render, edit the
// document, plan, apply; only the edited record is revised, with the reason;
// re-rendering shows the edit at the new version; the old plan is then stale
// and the old document conflicts.
func TestEndToEndRenderEditPlanApply(t *testing.T) {
	s := newStore(t)
	putVersioned(t, s, rec{ID: pBrief, Type: "brief", Title: "Brief", Intent: "Old brief intent", Constraints: []string{"B1"}}, 1)
	putVersioned(t, s, rec{ID: pBP, Type: "blueprint", Title: "Blueprint", Implements: pBrief, Intent: "BP intent", Constraints: []string{"P1", "P2"}}, 1)
	doc, err := Render(s, RenderOptions{})
	if err != nil {
		t.Fatal(err)
	}
	edited := replaceOnce(replaceOnce(doc, "Old brief intent", "New brief intent\n\nsecond paragraph"), "- P2\n", "- P2 revised\n")
	plan := mustPlan(t, s, edited)
	if len(plan.Changes) != 2 || plan.Conflicts != nil || plan.Problems != nil {
		t.Fatalf("plan %+v", *plan)
	}
	got, err := Apply(s, plan, ApplyOptions{Reason: "Doc edit", Author: "me@example.com"})
	if err != nil || !reflect.DeepEqual(got, []Revised{{pBrief, 2, false}, {pBP, 2, false}}) {
		t.Fatalf("%v, %v", got, err)
	}
	if r := loadRec(t, s, pBrief); r.Intent != "New brief intent\n\nsecond paragraph" || r.Reason != "Doc edit" {
		t.Fatalf("brief %+v", r)
	}
	if r := loadRec(t, s, pBP); r.Intent != "BP intent" || !reflect.DeepEqual(r.Constraints, []string{"P1", "P2 revised"}) {
		t.Fatalf("blueprint %+v", r)
	}
	again, _ := Render(s, RenderOptions{})
	if want := replaceOnce(replaceOnce(edited, "@1 -->", "@2 -->"), "@1 -->", "@2 -->"); again != want {
		t.Fatalf("re-rendered:\n%q\nwant:\n%q", again, want)
	}
	if _, err := Apply(s, plan, okOpts); !errors.Is(err, ErrStale) {
		t.Fatalf("second apply of the same plan: %v", err)
	}
	stale := mustPlan(t, s, edited)
	if len(stale.Conflicts) != 2 || stale.Changes != nil {
		t.Fatalf("old document against new records: %+v", *stale)
	}
	if clean := mustPlan(t, s, again); clean.Changes != nil || clean.Conflicts != nil || clean.Problems != nil || len(clean.Unchanged) != 2 {
		t.Fatalf("new document: %+v", *clean)
	}
}

// TestEndToEndDraftsStayAtVersion pins imprint prov-2026-6b2ceecb: applying to
// drafts keeps the version, so the edited document plans clean afterwards.
func TestEndToEndDraftsStayAtVersion(t *testing.T) {
	s := newStore(t)
	put(t, s, rec{ID: pBrief, Type: "brief", Title: "Brief", Intent: "Old", Constraints: []string{"B1"}})
	doc, _ := Render(s, RenderOptions{})
	edited := replaceOnce(doc, "Old", "New")
	plan := mustPlan(t, s, edited)
	if _, err := Apply(s, plan, okOpts); err != nil {
		t.Fatal(err)
	}
	after := mustPlan(t, s, edited)
	if after.Changes != nil || after.Conflicts != nil || !reflect.DeepEqual(after.Unchanged, []string{pBrief}) {
		t.Fatalf("%+v", *after)
	}
}
