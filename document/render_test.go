package document

import (
	"errors"
	"reflect"
	"strings"
	"testing"

	"github.com/livecodelife/provenance/record"
)

// ids used by the Render fixtures. Their order is chosen so that sorting by id
// differs from insertion order and from the document order.
const (
	rBrief1   = "prov-2026-00000b01"
	rBrief2   = "prov-2026-00000b02"
	rDepBrief = "prov-2026-00000b03"
)

// TestIsMarkerLine pins imprint prov-2026-400f65dc (isMarkerLine).
func TestIsMarkerLine(t *testing.T) {
	for _, c := range []struct {
		line string
		want bool
	}{
		{"## x", true},
		{"### x", true},
		{"## ", true},
		{"### ", true},
		{"  ## x", true},
		{"\t### x", true},
		{" \t ## x", true},
		{"<!-- provenance:abc@1 -->", true},
		{"   <!-- provenance:abc@1 -->", true},
		{"<!-- provenance:", true},
		{"**Constraints**", true},
		{"  **Constraints**  ", true},
		{"\t**Constraints**\t", true},
		// everything else is false
		{"", false},
		{"plain text", false},
		{"# x", false},
		{"#### x", false},
		{"##", false},
		{"###", false},
		{"##x", false},
		{"###x", false},
		{"<!--provenance:abc@1 -->", false},
		{"<!-- Provenance:abc@1 -->", false},
		{"text <!-- provenance:abc@1 -->", false},
		{"**Constraints** extra", false},
		{"extra **Constraints**", false},
		{"**constraints**", false},
		{"*Constraints*", false},
		{"- **Constraints**", false},
		{`\## x`, false},
		{`\**Constraints**`, false},
	} {
		if got := isMarkerLine(c.line); got != c.want {
			t.Errorf("isMarkerLine(%q) = %v, want %v", c.line, got, c.want)
		}
	}
}

// TestEscapeLine pins imprint prov-2026-5e341891 (escapeLine).
func TestEscapeLine(t *testing.T) {
	for _, c := range []struct{ in, want string }{
		{"## x", `\## x`},
		{"### x", `\### x`},
		{"  ## x", `\  ## x`},
		{"<!-- provenance:abc@1 -->", `\<!-- provenance:abc@1 -->`},
		{"**Constraints**", `\**Constraints**`},
		{" **Constraints** ", `\ **Constraints** `},
		// a line that already starts with backslashes before a marker gets one more
		{`\## x`, `\\## x`},
		{`\\## x`, `\\\## x`},
		{`\\\**Constraints**`, `\\\\**Constraints**`},
		{`\<!-- provenance:abc@1 -->`, `\\<!-- provenance:abc@1 -->`},
		// the backslashes are removed before the marker test, so a space after them counts
		{`\ ## x`, `\\ ## x`},
		// unchanged lines
		{"", ""},
		{"plain", "plain"},
		{"#### x", "#### x"},
		{"# x", "# x"},
		{`\plain`, `\plain`},
		{`\`, `\`},
		{`\\`, `\\`},
		{`\#### x`, `\#### x`},
		{"- item", "- item"},
	} {
		if got := escapeLine(c.in); got != c.want {
			t.Errorf("escapeLine(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}

// TestEscapeThenUnescapeRestoresEveryLine pins imprints prov-2026-5e341891 and
// prov-2026-c3d52774: the two functions are inverse on every line.
func TestEscapeThenUnescapeRestoresEveryLine(t *testing.T) {
	for _, line := range []string{
		"", "plain", "## x", "### x", "#### x", "  ## x", "<!-- provenance:a@1 -->", "**Constraints**",
		`\`, `\\`, `\## x`, `\\## x`, `\\\## x`, `\<!-- provenance:a@1 -->`, `\**Constraints**`,
		`\ ## x`, `\plain`, `\#### x`, "\t### x", " **Constraints** ", "- item", "text ## x",
	} {
		if got := unescapeLine(escapeLine(line)); got != line {
			t.Errorf("unescapeLine(escapeLine(%q)) = %q", line, got)
		}
	}
}

// renderFixture is the store of the golden test: a brief, a blueprint that
// implements it (at version 2), an orphan blueprint with no constraints, and
// records Render must leave out.
func renderFixture(t *testing.T) *record.FilesStore {
	t.Helper()
	s := newStore(t)
	put(t, s, rec{ID: rBrief1, Type: "brief", Title: "Brief One", Intent: "Brief intent.",
		Constraints: []string{"First constraint.", "Second constraint\ncontinues here."}})
	putVersioned(t, s, rec{ID: "prov-2026-0000c001", Type: "blueprint", Title: "Blueprint One",
		Implements: rBrief1, Intent: "Blueprint intent\nspans two lines.", Constraints: []string{"Only one."}}, 2)
	put(t, s, rec{ID: "prov-2026-0000a001", Type: "blueprint", Title: "Orphan", Intent: "Orphan intent."})
	put(t, s, rec{ID: "prov-2026-0000e001", Type: "imprint", Title: "An imprint", Implements: "prov-2026-0000c001",
		Intent: "imprint intent", Constraints: []string{"imprint constraint"}})
	return s
}

// TestRenderGolden pins imprint prov-2026-f40f3a4a (Render): the exact bytes of
// the document, headings by type, anchor lines, blank lines, constraint list
// items, continuation indent, section separation and the final newline.
func TestRenderGolden(t *testing.T) {
	st := renderFixture(t)
	fs := guard(st)
	want := "## Brief One\n" +
		"<!-- provenance:prov-2026-00000b01@1 -->\n" +
		"\n" +
		"Brief intent.\n" +
		"\n" +
		"**Constraints**\n" +
		"- First constraint.\n" +
		"- Second constraint\n" +
		"  continues here.\n" +
		"\n" +
		"### Blueprint One\n" +
		"<!-- provenance:prov-2026-0000c001@2 -->\n" +
		"\n" +
		"Blueprint intent\n" +
		"spans two lines.\n" +
		"\n" +
		"**Constraints**\n" +
		"- Only one.\n" +
		"\n" +
		"### Orphan\n" +
		"<!-- provenance:prov-2026-0000a001@1 -->\n" +
		"\n" +
		"Orphan intent.\n" +
		"\n" +
		"**Constraints**\n"
	got, err := Render(fs, RenderOptions{})
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Fatalf("document:\n%q\nwant:\n%q", got, want)
	}
	// Render never creates, edits or deletes a record (imprint f40f3a4a, blueprint 3e9a2764).
	if len(fs.writes) != 0 {
		t.Fatalf("Render wrote: %v", fs.writes)
	}
}

// TestRenderStoreRootUntouched is the on-disk half of the previous contract.
func TestRenderStoreRootUntouched(t *testing.T) {
	s := newStore(t)
	putAll(t, s, rec{ID: rBrief1, Type: "brief", Title: "B", Intent: "i", Constraints: []string{"c"}})
	putVersioned(t, s, rec{ID: "prov-2026-0000c001", Title: "P", Implements: rBrief1, Intent: "i"}, 2)
	before := tree(t, s.Root)
	for _, o := range []RenderOptions{{}, {IDs: true}, {Scopes: []string{"x"}}} {
		if _, err := Render(s, o); err != nil {
			t.Fatal(err)
		}
	}
	sameTree(t, before, tree(t, s.Root))
}

// TestRenderOrdering pins imprint prov-2026-f40f3a4a: briefs by id, each
// followed by its local blueprints by id, the remaining blueprints last by id;
// blueprints with an ImplementsRepo or whose parent is not a brief in the
// document are remaining; an empty type is a blueprint; deprecated, superseded,
// bug and imprint records are left out.
func TestRenderOrdering(t *testing.T) {
	s := newStore(t)
	putAll(t, s,
		rec{ID: rBrief2, Type: "brief", Title: "Brief B"},
		rec{ID: rBrief1, Type: "brief", Title: "Brief A"},
		rec{ID: rDepBrief, Type: "brief", Title: "Deprecated brief", Status: "deprecated"},
		rec{ID: "prov-2026-00000a01", Title: "Under B2 first", Implements: rBrief2},
		rec{ID: "prov-2026-00000f01", Title: "Under B2 second", Implements: rBrief2},
		rec{ID: "prov-2026-00000c01", Title: "Under B1", Type: "blueprint", Implements: rBrief1},
		rec{ID: "prov-2026-00000d03", Title: "No type", Implements: rBrief1}, // type key absent
		rec{ID: "prov-2026-00000e01", Title: "Orphan e"},
		rec{ID: "prov-2026-00000d01", Title: "Orphan d", Implements: "prov-2026-00000999"},
		rec{ID: "prov-2026-00000b99", Title: "Remote parent", Implements: rBrief1, ImplementsRepo: "other"},
		rec{ID: "prov-2026-00000d02", Title: "Child of blueprint", Implements: "prov-2026-00000c01"},
		rec{ID: "prov-2026-00000d00", Title: "Child of deprecated", Implements: rDepBrief},
		rec{ID: "prov-2026-00000aa1", Title: "A deprecated blueprint", Status: "deprecated"},
		rec{ID: "prov-2026-00000aa2", Title: "A superseded blueprint", Status: "superseded", Schema1: true},
		rec{ID: "prov-2026-00000aa3", Title: "A bug", Type: "bug", Status: "open", Schema1: true},
		rec{ID: "prov-2026-00000aa4", Title: "An imprint", Type: "imprint", Implements: rBrief1},
	)
	out, err := Render(s, RenderOptions{})
	if err != nil {
		t.Fatal(err)
	}
	want := []string{
		"## Brief A",
		"### Under B1",
		"### No type",
		"## Brief B",
		"### Under B2 first",
		"### Under B2 second",
		"### Remote parent",
		"### Child of deprecated",
		"### Orphan d",
		"### Child of blueprint",
		"### Orphan e",
	}
	if got := headings(out); !reflect.DeepEqual(got, want) {
		t.Fatalf("headings:\n%s\nwant:\n%s", strings.Join(got, "\n"), strings.Join(want, "\n"))
	}
}

// TestRenderScopes pins imprint prov-2026-f40f3a4a and blueprint
// prov-2026-3e9a2764: with Scopes a record is kept only when one of its
// affected_scope patterns matches one of the paths (scope.Match(path,
// pattern)); an empty affected_scope matches nothing; a blueprint whose brief is
// filtered out moves to the remaining blueprints.
func TestRenderScopes(t *testing.T) {
	s := newStore(t)
	putAll(t, s,
		rec{ID: "prov-2026-00005b01", Type: "brief", Title: "In Go", Scope: []string{"pkg/*.go"}},
		rec{ID: "prov-2026-00005b02", Type: "brief", Title: "In docs", Scope: []string{"docs/**"}},
		rec{ID: "prov-2026-00005c03", Title: "Both", Implements: "prov-2026-00005b01", Scope: []string{"pkg/a.go", "docs/x.md"}},
		rec{ID: "prov-2026-00005c04", Title: "None", Implements: "prov-2026-00005b01"},
		rec{ID: "prov-2026-00005c05", Title: "Pat", Implements: "prov-2026-00005b02", Scope: []string{"re:^cmd/"}},
	)
	for _, c := range []struct {
		name   string
		scopes []string
		want   []string
	}{
		{"no filter", nil, []string{"## In Go", "### Both", "### None", "## In docs", "### Pat"}},
		{"empty list is no filter", []string{}, []string{"## In Go", "### Both", "### None", "## In docs", "### Pat"}},
		{"glob and exact", []string{"pkg/a.go"}, []string{"## In Go", "### Both"}},
		{"double star glob, brief filtered out", []string{"docs/x.md"}, []string{"## In docs", "### Both"}},
		{"regex pattern", []string{"cmd/main.go"}, []string{"### Pat"}},
		{"several paths are a union", []string{"pkg/a.go", "docs/x.md"}, []string{"## In Go", "### Both", "## In docs"}},
		{"single star stays in one segment", []string{"pkg/sub/a.go"}, nil},
		{"nothing matches", []string{"other/file.txt"}, nil},
	} {
		t.Run(c.name, func(t *testing.T) {
			out, err := Render(s, RenderOptions{Scopes: c.scopes})
			if err != nil {
				t.Fatal(err)
			}
			if got := headings(out); !reflect.DeepEqual(got, c.want) {
				t.Fatalf("headings %q, want %q", got, c.want)
			}
			if len(c.want) == 0 && out != "" {
				t.Fatalf("a document with no section must be the empty string, got %q", out)
			}
		})
	}
}

// TestRenderMatchError pins imprint prov-2026-f40f3a4a: when Match returns an
// error Render returns an empty string and that error; without Scopes no match
// is attempted so the bad pattern is harmless.
func TestRenderMatchError(t *testing.T) {
	s := newStore(t)
	put(t, s, rec{ID: rBrief1, Type: "brief", Title: "Bad pattern", Scope: []string{"re:("}})
	out, err := Render(s, RenderOptions{Scopes: []string{"a.go"}})
	if err == nil || out != "" {
		t.Fatalf("got %q, %v; want empty string and an error", out, err)
	}
	if !strings.Contains(err.Error(), "invalid regex pattern") {
		t.Fatalf("error is not scope.Match's: %v", err)
	}
	if out, err = Render(s, RenderOptions{}); err != nil || !strings.HasPrefix(out, "## Bad pattern\n") {
		t.Fatalf("without Scopes: %q, %v", out, err)
	}
}

// TestRenderStoreErrors pins imprint prov-2026-f40f3a4a: an empty string and the
// error when listing or loading fails.
func TestRenderStoreErrors(t *testing.T) {
	boom := errors.New("boom")
	s := newStore(t)
	putAll(t, s,
		rec{ID: rBrief1, Type: "brief", Title: "One"},
		rec{ID: rBrief2, Type: "brief", Title: "Two"},
	)
	f := &faultStore{Store: s, listErr: boom}
	if out, err := Render(f, RenderOptions{}); out != "" || !errors.Is(err, boom) {
		t.Fatalf("list failure: %q, %v", out, err)
	}
	f = &faultStore{Store: s, loadErr: map[string]error{rBrief2: boom}}
	if out, err := Render(f, RenderOptions{}); out != "" || !errors.Is(err, boom) {
		t.Fatalf("load failure: %q, %v", out, err)
	}
}

// TestRenderEmpty pins imprint prov-2026-f40f3a4a: the document is the empty
// string when it holds no section, and ends in exactly one newline otherwise.
func TestRenderEmpty(t *testing.T) {
	s := newStore(t)
	if out, err := Render(s, RenderOptions{}); err != nil || out != "" {
		t.Fatalf("empty store: %q, %v", out, err)
	}
	put(t, s, rec{ID: "prov-2026-00000aa4", Title: "Only an imprint", Type: "imprint"})
	put(t, s, rec{ID: "prov-2026-00000aa5", Title: "Only deprecated", Status: "deprecated"})
	if out, err := Render(s, RenderOptions{}); err != nil || out != "" {
		t.Fatalf("no brief or blueprint: %q, %v", out, err)
	}
	put(t, s, rec{ID: rBrief1, Type: "brief", Title: "One", Intent: "i", Constraints: []string{"c"}})
	out, err := Render(s, RenderOptions{})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasSuffix(out, "- c\n") || strings.HasSuffix(out, "\n\n") {
		t.Fatalf("must end with exactly one newline: %q", out)
	}
}

// TestRenderIDs pins imprint prov-2026-f40f3a4a and blueprint
// prov-2026-3e9a2764: ids appear only in the anchor line unless IDs is set,
// then also in the heading as " (ID)".
func TestRenderIDs(t *testing.T) {
	s := newStore(t)
	putAll(t, s,
		rec{ID: rBrief1, Type: "brief", Title: "Brief One", Intent: "i"},
		rec{ID: "prov-2026-0000c001", Title: "Blueprint One", Implements: rBrief1, Intent: "j"},
	)
	plain, err := Render(s, RenderOptions{})
	if err != nil {
		t.Fatal(err)
	}
	withIDs, err := Render(s, RenderOptions{IDs: true})
	if err != nil {
		t.Fatal(err)
	}
	for _, id := range []string{rBrief1, "prov-2026-0000c001"} {
		if n := strings.Count(plain, id); n != 1 {
			t.Errorf("without IDs, %s appears %d times:\n%s", id, n, plain)
		}
		if n := strings.Count(withIDs, id); n != 2 {
			t.Errorf("with IDs, %s appears %d times:\n%s", id, n, withIDs)
		}
	}
	want := []string{"## Brief One (" + rBrief1 + ")", "### Blueprint One (prov-2026-0000c001)"}
	if got := headings(withIDs); !reflect.DeepEqual(got, want) {
		t.Fatalf("headings %q, want %q", got, want)
	}
	if !strings.Contains(withIDs, "## Brief One ("+rBrief1+")\n<!-- provenance:"+rBrief1+"@1 -->\n") {
		t.Fatalf("anchor must follow the heading directly:\n%s", withIDs)
	}
}

// TestRenderTitleWhitespace pins imprint prov-2026-f40f3a4a: every run of
// whitespace in a title collapses to one space.
func TestRenderTitleWhitespace(t *testing.T) {
	s := newStore(t)
	put(t, s, rec{ID: rBrief1, Type: "brief", Title: "  Spaced \t  Out   Title "})
	out, err := Render(s, RenderOptions{})
	if err != nil {
		t.Fatal(err)
	}
	if got, want := headings(out), []string{"## Spaced Out Title"}; !reflect.DeepEqual(got, want) {
		t.Fatalf("headings %q, want %q", got, want)
	}
	if out, _ = Render(s, RenderOptions{IDs: true}); !strings.HasPrefix(out, "## Spaced Out Title ("+rBrief1+")\n") {
		t.Fatalf("with IDs: %q", out)
	}
}

// TestRenderTrimsAndFormatsText pins imprint prov-2026-f40f3a4a: the intent
// and each constraint are written with surrounding whitespace removed; a
// multi-line constraint continues with two spaces, a blank line inside it is
// two spaces; the Constraints line is written even without constraints.
func TestRenderTrimsAndFormatsText(t *testing.T) {
	s := newStore(t)
	put(t, s, rec{ID: rBrief1, Type: "brief", Title: "T", Intent: "\n  Padded intent  \n\n",
		Constraints: []string{"  padded  ", "para one\n\npara two\n  indented more", "\n\nlast\n"}})
	put(t, s, rec{ID: "prov-2026-0000c001", Title: "No constraints", Intent: "x"})
	out, err := Render(s, RenderOptions{})
	if err != nil {
		t.Fatal(err)
	}
	want := "## T\n<!-- provenance:" + rBrief1 + "@1 -->\n\nPadded intent\n\n**Constraints**\n" +
		"- padded\n" +
		"- para one\n  \n  para two\n    indented more\n" +
		"- last\n" +
		"\n" +
		"### No constraints\n<!-- provenance:prov-2026-0000c001@1 -->\n\nx\n\n**Constraints**\n"
	if out != want {
		t.Fatalf("document:\n%q\nwant:\n%q", out, want)
	}
}

// TestRenderEmptyIntent pins imprint prov-2026-f40f3a4a for the part that is
// unambiguous: an empty intent writes no intent line and the section still
// parses back with an empty intent and its constraints. (The exact blank-line
// count between anchor and marker is not pinned; see the ambiguity report.)
func TestRenderEmptyIntent(t *testing.T) {
	s := newStore(t)
	put(t, s, rec{ID: rBrief1, Type: "brief", Title: "T", Constraints: []string{"c"}})
	out, err := Render(s, RenderOptions{})
	if err != nil {
		t.Fatal(err)
	}
	head := "## T\n<!-- provenance:" + rBrief1 + "@1 -->\n"
	if !strings.HasPrefix(out, head) || !strings.HasSuffix(out, "**Constraints**\n- c\n") {
		t.Fatalf("document: %q", out)
	}
	mid := strings.TrimSuffix(strings.TrimPrefix(out, head), "**Constraints**\n- c\n")
	if strings.Trim(mid, "\n") != "" {
		t.Fatalf("only blank lines may separate anchor and marker: %q", mid)
	}
	secs, err := Parse(out)
	if err != nil || len(secs) != 1 || secs[0].Intent != "" || !secs[0].HasConstraints || !reflect.DeepEqual(secs[0].Constraints, []string{"c"}) {
		t.Fatalf("parsed back: %+v, %v", secs, err)
	}
}

// TestRenderEscapesStructure pins blueprint prov-2026-3e9a2764 and imprint
// prov-2026-f40f3a4a: intent lines that look like document structure are
// written through escapeLine.
func TestRenderEscapesStructure(t *testing.T) {
	s := newStore(t)
	intent := "Intro\n## Not a heading\n### Nor this\n  ## indented\n<!-- provenance:prov-2026-deadbeef@9 -->\n**Constraints**\n" + `\## already escaped` + "\nplain"
	put(t, s, rec{ID: rBrief1, Type: "brief", Title: "T", Intent: intent, Constraints: []string{"c"}})
	out, err := Render(s, RenderOptions{})
	if err != nil {
		t.Fatal(err)
	}
	want := "## T\n<!-- provenance:" + rBrief1 + "@1 -->\n\n" +
		"Intro\n" +
		`\## Not a heading` + "\n" +
		`\### Nor this` + "\n" +
		`\  ## indented` + "\n" +
		`\<!-- provenance:prov-2026-deadbeef@9 -->` + "\n" +
		`\**Constraints**` + "\n" +
		`\\## already escaped` + "\n" +
		"plain\n\n**Constraints**\n- c\n"
	if out != want {
		t.Fatalf("document:\n%q\nwant:\n%q", out, want)
	}
	secs, err := Parse(out)
	if err != nil || len(secs) != 1 || secs[0].Intent != intent {
		t.Fatalf("parse did not restore the intent: %+v, %v", secs, err)
	}
}

// TestRenderIsDeterministic pins blueprint prov-2026-3e9a2764 and imprint
// prov-2026-f40f3a4a: the same store and options give byte-identical output.
func TestRenderIsDeterministic(t *testing.T) {
	fs := renderFixture(t)
	for _, o := range []RenderOptions{{}, {IDs: true}, {Scopes: []string{"a.go"}}} {
		a, err1 := Render(fs, o)
		b, err2 := Render(fs, o)
		if err1 != nil || err2 != nil || a != b {
			t.Fatalf("%+v: renders differ or fail: %v %v", o, err1, err2)
		}
	}
}

// roundTripStore holds records with every awkward text shape Render must
// survive: structure-looking intent, backslashes, blank lines, multi-paragraph
// constraints, odd titles, an empty intent and no constraints.
func roundTripStore(t *testing.T) *record.FilesStore {
	t.Helper()
	s := newStore(t)
	putAll(t, s,
		rec{ID: rBrief1, Type: "brief", Title: "  A   brief \t title ",
			Intent: "Line one\n\nLine three after a blank\n## fake heading\n  ### indented fake\n**Constraints**\n  **Constraints**  \n<!-- provenance:x@1 -->\n" + `\## kept` + "\n" + `\\### two` + "\n   indented line\ntrailing spaces  \nlast",
			Constraints: []string{
				"simple",
				"multi\nline",
				"para one\n\npara two",
				"## not a heading\n<!-- provenance:y@2 -->\n**Constraints**\n" + `\## back` + "\n- looks like an item",
				"- starts with a dash",
			}},
		rec{ID: rBrief2, Type: "brief", Title: "Empty intent", Constraints: []string{"only"}},
		rec{ID: "prov-2026-0000c001", Title: "No constraints", Implements: rBrief1, Intent: "Has intent"},
		rec{ID: "prov-2026-0000c002", Title: "Nothing at all", Implements: rBrief1},
		rec{ID: "prov-2026-0000c003", Title: "Title (with parens)", Implements: rBrief2, Intent: "x", Constraints: []string{"y"}},
	)
	putVersioned(t, s, rec{ID: "prov-2026-0000c004", Title: "Versioned", Implements: rBrief2, Intent: "v", Constraints: []string{"w"}}, 3)
	return s
}

// TestRenderParseBuildPlanRoundTrip pins blueprint prov-2026-3e9a2764: rendering
// a store, parsing the output and planning it against the same store gives a
// plan with no changes, no conflicts and no problems (every rendered record
// unchanged), with and without ids in the headings.
func TestRenderParseBuildPlanRoundTrip(t *testing.T) {
	st := roundTripStore(t)
	fs := guard(st)
	for _, o := range []RenderOptions{{}, {IDs: true}} {
		out, err := Render(fs, o)
		if err != nil {
			t.Fatal(err)
		}
		secs, err := Parse(out)
		if err != nil {
			t.Fatal(err)
		}
		if len(secs) != 6 {
			t.Fatalf("%+v: %d sections, want 6:\n%s", o, len(secs), out)
		}
		plan, err := BuildPlan(fs, out)
		if err != nil {
			t.Fatal(err)
		}
		if plan.Changes != nil || plan.Conflicts != nil || plan.Problems != nil {
			t.Fatalf("%+v: round trip plan is not clean: %+v", o, *plan)
		}
		wantUnchanged := []string{rBrief1, rBrief2, "prov-2026-0000c001", "prov-2026-0000c002", "prov-2026-0000c003", "prov-2026-0000c004"}
		if len(wantUnchanged) != len(plan.Unchanged) || !reflect.DeepEqual(plan.Unchanged, wantUnchanged) {
			t.Fatalf("%+v: unchanged %v, want %v", o, plan.Unchanged, wantUnchanged)
		}
	}
	if len(fs.writes) != 0 {
		t.Fatalf("writes: %v", fs.writes)
	}
}

// TestRenderParseRestoresTextExactly pins blueprint prov-2026-632550bc and
// imprint prov-2026-6f0b5e46: after Render, Parse gives back each intent and
// each constraint as written (trimmed), including blank lines and
// structure-looking lines.
func TestRenderParseRestoresTextExactly(t *testing.T) {
	st := roundTripStore(t)
	out, err := Render(st, RenderOptions{})
	if err != nil {
		t.Fatal(err)
	}
	secs, err := Parse(out)
	if err != nil {
		t.Fatal(err)
	}
	if len(secs) != 6 {
		t.Fatalf("%d sections, want 6", len(secs))
	}
	for _, sc := range secs {
		r := loadRec(t, st, sc.ID)
		if !sc.Anchored || sc.Version != r.Version {
			t.Errorf("%s: anchored %v version %d, record version %d", sc.ID, sc.Anchored, sc.Version, r.Version)
		}
		if want := strings.TrimSpace(r.Intent); sc.Intent != want {
			t.Errorf("%s: intent %q, want %q", sc.ID, sc.Intent, want)
		}
		var want []string
		for _, c := range r.Constraints {
			want = append(want, strings.TrimSpace(c))
		}
		if len(sc.Constraints) != len(want) || (len(want) > 0 && !reflect.DeepEqual(sc.Constraints, want)) {
			t.Errorf("%s: constraints %q, want %q", sc.ID, sc.Constraints, want)
		}
		if !sc.HasConstraints {
			t.Errorf("%s: Render always writes the Constraints line", sc.ID)
		}
	}
}
