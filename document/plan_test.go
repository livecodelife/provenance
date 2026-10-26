package document

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/livecodelife/provenance/record"
)

const (
	pBrief  = "prov-2026-0000b001"
	pBP     = "prov-2026-0000c001"
	pVP     = "prov-2026-0000c002" // versioned to 3
	pImp    = "prov-2026-0000e001"
	pDep    = "prov-2026-0000f001"
	pSup    = "prov-2026-0000f002"
	pBug    = "prov-2026-0000f003"
	pAbsent = "prov-2026-00000001" // not in the store
)

// planStore holds the records the BuildPlan tests anchor against.
func planStore(t *testing.T) *record.FilesStore {
	t.Helper()
	s := newStore(t)
	putAll(t, s,
		rec{ID: pBrief, Type: "brief", Title: "Brief", Intent: "Brief intent", Constraints: []string{"B1", "B2"}},
		rec{ID: pBP, Type: "blueprint", Title: "Blueprint", Intent: "Blueprint intent", Constraints: []string{"P1"}, Implements: pBrief},
		rec{ID: pImp, Type: "imprint", Title: "Imprint", Intent: "i", Implements: pBP},
		rec{ID: pDep, Type: "blueprint", Title: "Deprecated", Status: "deprecated"},
		rec{ID: pSup, Type: "blueprint", Title: "Superseded", Status: "superseded", Schema1: true},
		rec{ID: pBug, Type: "bug", Title: "Bug", Status: "open", Schema1: true},
	)
	putVersioned(t, s, rec{ID: pVP, Title: "Versioned", Intent: "V intent", Constraints: []string{"V1"}}, 3)
	return s
}

// prob is the observable part of a Problem that every rule fixes.
type prob struct {
	Kind, ID string
	Line     int
	Message  string
}

func probs(ps []Problem) []prob {
	var out []prob
	for _, p := range ps {
		out = append(out, prob{p.Kind, p.ID, p.Line, p.Message})
	}
	return out
}

func mustPlan(t *testing.T, s record.Store, doc string) *Plan {
	t.Helper()
	p, err := BuildPlan(s, doc)
	if err != nil || p == nil {
		t.Fatalf("BuildPlan: %v, %v", p, err)
	}
	return p
}

// TestBuildPlanRules pins imprint prov-2026-c92e5c96: each rule gives a Problem
// whose Kind is the rule name, with the anchored id (empty when unanchored), the
// heading's line and the fixed message, listed in Problems only.
func TestBuildPlanRules(t *testing.T) {
	s := planStore(t)
	for _, c := range []struct {
		name string
		doc  string
		want prob
	}{
		{"unanchored", "## Blueprint\nno anchor here\n", prob{"unanchored", "", 1, "heading has no anchor"}},
		{"unknown", sec(3, "Blueprint", pAbsent, 1, "x", "y"), prob{"unknown", pAbsent, 1, "no record with this id"}},
		{"imprint", sec(3, "Imprint", pImp, 1, "i"), prob{"imprint", pImp, 1, "the record is not a brief or blueprint"}},
		{"bug type is not a brief or blueprint", sec(3, "Bug", pBug, 1, "i"), prob{"imprint", pBug, 1, "the record is not a brief or blueprint"}},
		{"deprecated", sec(3, "Deprecated", pDep, 1, "i"), prob{"inactive", pDep, 1, "the record is deprecated or superseded"}},
		{"superseded", sec(3, "Superseded", pSup, 1, "i"), prob{"inactive", pSup, 1, "the record is deprecated or superseded"}},
		{"brief at level 3", sec(3, "Brief", pBrief, 1, "x", "y"), prob{"level", pBrief, 1, "heading level does not match the record type"}},
		{"blueprint at level 2", sec(2, "Blueprint", pBP, 1, "x", "y"), prob{"level", pBP, 1, "heading level does not match the record type"}},
		{"heading differs", sec(3, "Another title", pBP, 1, "x", "y"), prob{"heading", pBP, 1, "heading text differs from the record title"}},
		{"heading case differs", sec(3, "blueprint", pBP, 1, "x", "y"), prob{"heading", pBP, 1, "heading text differs from the record title"}},
		{"heading with another id", sec(3, "Blueprint ("+pBrief+")", pBP, 1, "x", "y"), prob{"heading", pBP, 1, "heading text differs from the record title"}},
		{"no constraints marker", "### Blueprint\n<!-- provenance:" + pBP + "@1 -->\n\nintent\n", prob{"malformed", pBP, 1, "section has no constraints marker"}},
	} {
		t.Run(c.name, func(t *testing.T) {
			p := mustPlan(t, s, c.doc)
			if got := probs(p.Problems); !reflect.DeepEqual(got, []prob{c.want}) {
				t.Fatalf("problems %+v, want %+v", got, c.want)
			}
			if p.Changes != nil || p.Conflicts != nil || p.Unchanged != nil {
				t.Fatalf("a problem section is listed nowhere else: %+v", *p)
			}
		})
	}
}

// TestBuildPlanProblemLineNumbers pins imprint prov-2026-c92e5c96: Line is the
// heading's line number in the text.
func TestBuildPlanProblemLineNumbers(t *testing.T) {
	s := planStore(t)
	var d docBuilder
	d.add("preamble\n\n")
	l := d.add(sec(3, "Wrong", pBP, 1, "x"))
	p := mustPlan(t, s, d.String())
	if len(p.Problems) != 1 || p.Problems[0].Line != l || l != 3 {
		t.Fatalf("problems %+v, want line %d", p.Problems, l)
	}
}

// TestBuildPlanConflict pins imprint prov-2026-c92e5c96 and blueprint
// prov-2026-9000e51c: an anchored version above or below the record's gives a
// conflict with both versions set, listed only in Conflicts, never as a change
// or an unchanged record, even when the text also changed.
func TestBuildPlanConflict(t *testing.T) {
	s := planStore(t)
	for _, c := range []struct {
		anchored int
		msg      string
	}{
		{1, "the record is at version 3, the section is from version 1"},
		{2, "the record is at version 3, the section is from version 2"},
		{5, "the record is at version 3, the section is from version 5"},
	} {
		p := mustPlan(t, s, sec(3, "Versioned", pVP, c.anchored, "changed text", "changed"))
		if len(p.Conflicts) != 1 || p.Problems != nil || p.Changes != nil || p.Unchanged != nil {
			t.Fatalf("anchored %d: %+v", c.anchored, *p)
		}
		got := p.Conflicts[0]
		if got.Kind != "conflict" || got.ID != pVP || got.Line != 1 || got.Message != c.msg ||
			got.AnchoredVersion != c.anchored || got.CurrentVersion != 3 {
			t.Fatalf("anchored %d: conflict %+v", c.anchored, got)
		}
	}
}

// TestBuildPlanPrecedence pins imprint prov-2026-c92e5c96: the first applicable
// rule decides, in the order unanchored, unknown, imprint, inactive, duplicate,
// level, heading, malformed, conflict, then changed or unchanged.
func TestBuildPlanPrecedence(t *testing.T) {
	s := planStore(t)
	anchor := func(id string, v int) string { return "<!-- provenance:" + id + "@" + strconv.Itoa(v) + " -->\n" }
	for _, c := range []struct {
		name string
		doc  string
		want []prob
		conf int
	}{
		{"unknown before level", sec(3, "X", pAbsent, 1, "x"), []prob{{"unknown", pAbsent, 1, "no record with this id"}}, 0},
		{"unknown before duplicate", sec(3, "X", pAbsent, 1, "x") + sec(3, "X", pAbsent, 1, "x"),
			[]prob{{"unknown", pAbsent, 1, "no record with this id"}, {"unknown", pAbsent, 8, "no record with this id"}}, 0},
		{"imprint before inactive", sec(3, "Imprint", pImp, 1, "i"), []prob{{"imprint", pImp, 1, "the record is not a brief or blueprint"}}, 0},
		{"inactive before duplicate", sec(3, "Deprecated", pDep, 1, "i") + sec(3, "Deprecated", pDep, 1, "i"),
			[]prob{{"inactive", pDep, 1, "the record is deprecated or superseded"}, {"inactive", pDep, 8, "the record is deprecated or superseded"}}, 0},
		{"inactive before level", sec(2, "Deprecated", pDep, 1, "i"), []prob{{"inactive", pDep, 1, "the record is deprecated or superseded"}}, 0},
		{"level before heading", sec(2, "Wrong title", pBP, 1, "x", "y"), []prob{{"level", pBP, 1, "heading level does not match the record type"}}, 0},
		{"heading before malformed", "### Wrong title\n" + anchor(pBP, 1) + "\ntext\n", []prob{{"heading", pBP, 1, "heading text differs from the record title"}}, 0},
		{"malformed before conflict", "### Versioned\n" + anchor(pVP, 1) + "\ntext\n", []prob{{"malformed", pVP, 1, "section has no constraints marker"}}, 0},
		{"heading before conflict", sec(3, "Wrong title", pVP, 1, "x", "y"), []prob{{"heading", pVP, 1, "heading text differs from the record title"}}, 0},
		{"level before conflict", sec(2, "Versioned", pVP, 1, "x", "y"), []prob{{"level", pVP, 1, "heading level does not match the record type"}}, 0},
		{"conflict before change", sec(3, "Versioned", pVP, 2, "changed", "changed"), nil, 1},
	} {
		t.Run(c.name, func(t *testing.T) {
			p := mustPlan(t, s, c.doc)
			if got := probs(p.Problems); !reflect.DeepEqual(got, c.want) {
				t.Fatalf("problems %+v, want %+v", got, c.want)
			}
			if len(p.Conflicts) != c.conf || p.Changes != nil || p.Unchanged != nil {
				t.Fatalf("plan %+v", *p)
			}
		})
	}
}

// TestBuildPlanDuplicate pins imprint prov-2026-c92e5c96: a later section with an
// already anchored id is a duplicate, the earlier one is judged by the normal
// rules (here unchanged, or here a heading problem), and the duplicate is
// reported even when the earlier section had a problem.
func TestBuildPlanDuplicate(t *testing.T) {
	s := planStore(t)
	var d docBuilder
	d.add(sec(3, "Blueprint", pBP, 1, "Blueprint intent", "P1"))
	l2 := d.add(sec(3, "Blueprint", pBP, 1, "Blueprint intent", "P1"))
	p := mustPlan(t, s, d.String())
	want := []prob{{"duplicate", pBP, l2, "this record is anchored more than once"}}
	if got := probs(p.Problems); !reflect.DeepEqual(got, want) {
		t.Fatalf("problems %+v, want %+v", got, want)
	}
	if !reflect.DeepEqual(p.Unchanged, []string{pBP}) || p.Changes != nil || p.Conflicts != nil {
		t.Fatalf("the first section is judged normally: %+v", *p)
	}

	d = docBuilder{}
	l1 := d.add(sec(3, "Wrong title", pBP, 1, "x"))
	l2 = d.add(sec(3, "Blueprint", pBP, 1, "Blueprint intent", "P1"))
	p = mustPlan(t, s, d.String())
	want = []prob{{"heading", pBP, l1, "heading text differs from the record title"}, {"duplicate", pBP, l2, "this record is anchored more than once"}}
	if got := probs(p.Problems); !reflect.DeepEqual(got, want) {
		t.Fatalf("problems %+v, want %+v", got, want)
	}
	if p.Unchanged != nil || p.Changes != nil {
		t.Fatalf("plan %+v", *p)
	}
}

// TestBuildPlanHeadingForms pins imprint prov-2026-c92e5c96: after normalize the
// heading may be the record's title or the title followed by a space and the id
// in parentheses; the record is found by anchor, never by heading.
func TestBuildPlanHeadingForms(t *testing.T) {
	s := newStore(t)
	put(t, s, rec{ID: pBP, Title: "Spaced  Title", Intent: "i", Constraints: []string{"c"}})
	for _, h := range []string{"Spaced Title", "Spaced  Title", "Spaced Title (" + pBP + ")", "Spaced    Title   (" + pBP + ")"} {
		p := mustPlan(t, s, sec(3, h, pBP, 1, "i", "c"))
		if p.Problems != nil || !reflect.DeepEqual(p.Unchanged, []string{pBP}) {
			t.Errorf("heading %q: %+v", h, *p)
		}
	}
	for _, h := range []string{"Spaced", "Spaced Title (" + pBP + ") extra", "(" + pBP + ") Spaced Title", "Spaced Title (" + pBP + "x)"} {
		p := mustPlan(t, s, sec(3, h, pBP, 1, "i", "c"))
		if len(p.Problems) != 1 || p.Problems[0].Kind != "heading" {
			t.Errorf("heading %q: %+v", h, *p)
		}
	}
	// another record's title does not select that record
	put(t, s, rec{ID: pVP, Title: "Other", Intent: "i"})
	p := mustPlan(t, s, sec(3, "Other", pBP, 1, "i", "c"))
	if len(p.Problems) != 1 || p.Problems[0].Kind != "heading" || p.Problems[0].ID != pBP {
		t.Fatalf("plan %+v", *p)
	}
}

// TestBuildPlanChanges pins imprint prov-2026-c92e5c96 and blueprint
// prov-2026-9000e51c: what counts as an intent or constraints change, the
// Change fields, and that unchanged sections are listed as Unchanged.
func TestBuildPlanChanges(t *testing.T) {
	s := planStore(t)
	for _, c := range []struct {
		name        string
		intent      string
		cons        []string
		wantIntent  bool
		wantCons    bool
		wantChanged bool
	}{
		{"identical", "Blueprint intent", []string{"P1"}, false, false, false},
		{"whitespace only in intent", "  Blueprint \t  intent  ", []string{"P1"}, false, false, false},
		{"line wrapping only in intent", "Blueprint\nintent", []string{"P1"}, false, false, false},
		{"whitespace only in constraint", "Blueprint intent", []string{"  P1  "}, false, false, false},
		{"intent changed", "New intent", []string{"P1"}, true, false, true},
		{"intent emptied", "", []string{"P1"}, true, false, true},
		{"constraint text changed", "Blueprint intent", []string{"P1 changed"}, false, true, true},
		{"constraint added", "Blueprint intent", []string{"P1", "P2"}, false, true, true},
		{"constraint removed", "Blueprint intent", nil, false, true, true},
		{"both changed", "New intent", []string{"P9"}, true, true, true},
	} {
		t.Run(c.name, func(t *testing.T) {
			p := mustPlan(t, s, sec(3, "Blueprint", pBP, 1, c.intent, c.cons...))
			if p.Problems != nil || p.Conflicts != nil {
				t.Fatalf("plan %+v", *p)
			}
			if !c.wantChanged {
				if p.Changes != nil || !reflect.DeepEqual(p.Unchanged, []string{pBP}) {
					t.Fatalf("want unchanged: %+v", *p)
				}
				return
			}
			if p.Unchanged != nil || len(p.Changes) != 1 {
				t.Fatalf("want one change: %+v", *p)
			}
			ch := p.Changes[0]
			if ch.ID != pBP || ch.Title != "Blueprint" || ch.AnchoredVersion != 1 || ch.CurrentVersion != 1 ||
				ch.IntentChanged != c.wantIntent || ch.ConstraintsChanged != c.wantCons {
				t.Fatalf("change %+v", ch)
			}
		})
	}
}

// TestBuildPlanConstraintOrder pins imprint prov-2026-c92e5c96: constraints are
// compared by position, so a reorder is a change.
func TestBuildPlanConstraintOrder(t *testing.T) {
	s := planStore(t)
	p := mustPlan(t, s, sec(2, "Brief", pBrief, 1, "Brief intent", "B2", "B1"))
	if len(p.Changes) != 1 || !p.Changes[0].ConstraintsChanged || p.Changes[0].IntentChanged {
		t.Fatalf("plan %+v", *p)
	}
}

// TestBuildPlanChangeCarriesParsedText pins imprint prov-2026-c92e5c96: the
// Change holds the section's Intent and Constraints as parsed (not normalized),
// the record's Title (not the heading) and both versions.
func TestBuildPlanChangeCarriesParsedText(t *testing.T) {
	s := planStore(t)
	doc := "### Versioned (" + pVP + ")\n<!-- provenance:" + pVP + "@3 -->\n\nNew   intent\n\n  with an indented line\n\n**Constraints**\n- first  \n  second\n- third\n"
	p := mustPlan(t, s, doc)
	want := Change{ID: pVP, Title: "Versioned", AnchoredVersion: 3, CurrentVersion: 3, IntentChanged: true, ConstraintsChanged: true,
		Intent: "New   intent\n\n  with an indented line", Constraints: []string{"first  \nsecond", "third"}}
	if len(p.Changes) != 1 || !reflect.DeepEqual(p.Changes[0], want) {
		t.Fatalf("changes %+v\nwant %+v", p.Changes, want)
	}
}

// TestBuildPlanSorting pins imprint prov-2026-c92e5c96 and blueprint
// prov-2026-9000e51c: Changes, Conflicts and Unchanged sort by id; Problems sort
// by id with empty ids first, ties by line; absent records are not listed.
func TestBuildPlanSorting(t *testing.T) {
	s := planStore(t)
	put(t, s, rec{ID: "prov-2026-0000c003", Title: "Third", Intent: "t"})
	put(t, s, rec{ID: "prov-2026-0000c004", Title: "Fourth", Intent: "f"})
	putVersioned(t, s, rec{ID: "prov-2026-0000c005", Title: "Fifth", Intent: "f"}, 2)
	var d docBuilder
	d.add(sec(3, "Fourth", "prov-2026-0000c004", 1, "f changed"))        // change
	d.add(sec(3, "Versioned", pVP, 1, "stale"))                          // conflict (record at 3)
	lu1 := d.add("### Floating\ntext only\n\n")                          // unanchored
	d.add(sec(3, "Third", "prov-2026-0000c003", 1, "t"))                 // unchanged
	d.add(sec(3, "Fifth", "prov-2026-0000c005", 1, "stale"))             // conflict (record at 2)
	d.add(sec(3, "Blueprint", pBP, 1, "Blueprint intent changed", "P1")) // change
	lk := d.add(sec(3, "Ghost", pAbsent, 1, "x"))                        // unknown
	lu2 := d.add("## Another floating\n")                                // unanchored
	ld := d.add(sec(3, "Third", "prov-2026-0000c003", 1, "t"))           // duplicate
	p := mustPlan(t, s, d.String())

	var changeIDs, conflictIDs []string
	for _, c := range p.Changes {
		changeIDs = append(changeIDs, c.ID)
	}
	for _, c := range p.Conflicts {
		conflictIDs = append(conflictIDs, c.ID)
	}
	if want := []string{pBP, "prov-2026-0000c004"}; !reflect.DeepEqual(changeIDs, want) {
		t.Errorf("changes %v, want %v", changeIDs, want)
	}
	if want := []string{pVP, "prov-2026-0000c005"}; !reflect.DeepEqual(conflictIDs, want) {
		t.Errorf("conflicts %v, want %v", conflictIDs, want)
	}
	if want := []string{"prov-2026-0000c003"}; !reflect.DeepEqual(p.Unchanged, want) {
		t.Errorf("unchanged %v, want %v", p.Unchanged, want)
	}
	want := []prob{
		{"unanchored", "", lu1, "heading has no anchor"},
		{"unanchored", "", lu2, "heading has no anchor"},
		{"unknown", pAbsent, lk, "no record with this id"},
		{"duplicate", "prov-2026-0000c003", ld, "this record is anchored more than once"},
	}
	if got := probs(p.Problems); !reflect.DeepEqual(got, want) {
		t.Errorf("problems %+v\nwant %+v", got, want)
	}
}

// TestBuildPlanEmptySlicesAreNil pins imprint prov-2026-c92e5c96: plan slices
// without entries are nil, and an empty document gives an empty plan.
func TestBuildPlanEmptySlicesAreNil(t *testing.T) {
	s := planStore(t)
	for _, doc := range []string{"", "no headings at all\n"} {
		p := mustPlan(t, s, doc)
		if p.Changes != nil || p.Conflicts != nil || p.Problems != nil || p.Unchanged != nil {
			t.Fatalf("%q: %+v", doc, *p)
		}
	}
	p := mustPlan(t, s, sec(3, "Blueprint", pBP, 1, "Blueprint intent", "P1"))
	if p.Changes != nil || p.Conflicts != nil || p.Problems != nil || len(p.Unchanged) != 1 {
		t.Fatalf("%+v", *p)
	}
}

// TestBuildPlanAbsentRecords pins imprint prov-2026-c92e5c96: a record absent
// from the document is neither listed nor changed.
func TestBuildPlanAbsentRecords(t *testing.T) {
	s := planStore(t)
	before := tree(t, s.Root)
	p := mustPlan(t, s, sec(3, "Blueprint", pBP, 1, "Blueprint intent changed", "P1"))
	if len(p.Changes) != 1 || p.Changes[0].ID != pBP || p.Unchanged != nil || p.Conflicts != nil || p.Problems != nil {
		t.Fatalf("plan %+v", *p)
	}
	sameTree(t, before, tree(t, s.Root))
}

// TestBuildPlanLineEndings pins blueprint prov-2026-632550bc through BuildPlan:
// a CRLF document gives the same plan as the LF one.
func TestBuildPlanLineEndings(t *testing.T) {
	s := planStore(t)
	doc := sec(2, "Brief", pBrief, 1, "Changed\n\nintent", "B1", "B2 changed") + sec(3, "Blueprint", pBP, 1, "Blueprint intent", "P1")
	a := mustPlan(t, s, doc)
	b := mustPlan(t, s, strings.ReplaceAll(doc, "\n", "\r\n"))
	if !reflect.DeepEqual(a, b) || len(a.Changes) != 1 {
		t.Fatalf("LF %+v\nCRLF %+v", *a, *b)
	}
}

// TestBuildPlanWritesNothing pins imprint prov-2026-c92e5c96: BuildPlan writes
// nothing to the store.
func TestBuildPlanWritesNothing(t *testing.T) {
	s := planStore(t)
	fs := guard(s)
	before := tree(t, s.Root)
	doc := sec(2, "Brief", pBrief, 1, "changed", "x") + sec(3, "Versioned", pVP, 1, "stale") + "### floating\n"
	if _, err := BuildPlan(fs, doc); err != nil {
		t.Fatal(err)
	}
	if len(fs.writes) != 0 {
		t.Fatalf("writes %v", fs.writes)
	}
	sameTree(t, before, tree(t, s.Root))
}

// TestBuildPlanStoreErrors pins imprint prov-2026-c92e5c96: a nil Plan and the
// error when listing or loading a record fails.
func TestBuildPlanStoreErrors(t *testing.T) {
	boom := errors.New("boom")
	s := planStore(t)
	doc := sec(3, "Blueprint", pBP, 1, "Blueprint intent", "P1")
	if p, err := BuildPlan(&faultStore{Store: s, listErr: boom}, doc); p != nil || !errors.Is(err, boom) {
		t.Fatalf("list failure: %v, %v", p, err)
	}
	if p, err := BuildPlan(&faultStore{Store: s, loadErr: map[string]error{pBP: boom}}, doc); p != nil || !errors.Is(err, boom) {
		t.Fatalf("load failure: %v, %v", p, err)
	}
}

// TestNormalize pins imprint prov-2026-4dc24541 (normalize).
func TestNormalize(t *testing.T) {
	for _, c := range []struct{ in, want string }{
		{"", ""},
		{"   ", ""},
		{"\n\t \n", ""},
		{"a", "a"},
		{"  a  ", "a"},
		{"a b", "a b"},
		{"a   b", "a b"},
		{"a\tb\nc\r\nd", "a b c d"},
		{" \n a \t\t b \n\n c \n", "a b c"},
		{"A  B", "A B"},
	} {
		if got := normalize(c.in); got != c.want {
			t.Errorf("normalize(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}

// TestSharedTypesExist pins the field sets of imprints prov-2026-2b7b6b93
// (Change), prov-2026-b396ccea (Problem), prov-2026-6eeaa2bc (Plan),
// prov-2026-b98853b3 (Section) and prov-2026-477b93e0 (RenderOptions): the code
// below fails to compile when a field is missing or has another type.
func TestSharedTypesExist(t *testing.T) {
	_ = Change{ID: "", Title: "", AnchoredVersion: 0, CurrentVersion: 0, IntentChanged: false, ConstraintsChanged: false, Intent: "", Constraints: []string(nil)}
	_ = Problem{Kind: "", ID: "", Line: 0, AnchoredVersion: 0, CurrentVersion: 0, Message: ""}
	_ = Plan{Changes: []Change(nil), Conflicts: []Problem(nil), Problems: []Problem(nil), Unchanged: []string(nil)}
	_ = Section{Heading: "", Level: 0, Line: 0, ID: "", Version: 0, Anchored: false, Intent: "", Constraints: []string(nil), HasConstraints: false}
	_ = RenderOptions{Scopes: []string(nil), IDs: false}
	if !reflect.DeepEqual(RenderOptions{}, RenderOptions{Scopes: nil, IDs: false}) {
		t.Fatal("the zero RenderOptions means no filter and no ids")
	}
}
