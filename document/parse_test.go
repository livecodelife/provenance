package document

import (
	"math"
	"reflect"
	"slices"
	"strconv"
	"strings"
	"testing"
)

// TestUnescapeLine pins imprint prov-2026-c3d52774 (unescapeLine).
func TestUnescapeLine(t *testing.T) {
	for _, c := range []struct{ in, want string }{
		{`\## x`, "## x"},
		{`\### x`, "### x"},
		{`\  ## x`, "  ## x"},
		{`\<!-- provenance:abc@1 -->`, "<!-- provenance:abc@1 -->"},
		{`\**Constraints**`, "**Constraints**"},
		{`\\## x`, `\## x`},
		{`\\\## x`, `\\## x`},
		{`\\**Constraints**`, `\**Constraints**`},
		{`\ ## x`, " ## x"},
		// unchanged: no leading backslash, or the rest is not a marker line
		{"## x", "## x"},
		{"**Constraints**", "**Constraints**"},
		{"", ""},
		{"plain", "plain"},
		{`\plain`, `\plain`},
		{`\\plain`, `\\plain`},
		{`\`, `\`},
		{`\\`, `\\`},
		{`\#### x`, `\#### x`},
		{`\# x`, `\# x`},
		{`\## `, "## "},
		{`x\## y`, `x\## y`},
	} {
		if got := unescapeLine(c.in); got != c.want {
			t.Errorf("unescapeLine(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}

// TestParseAnchor pins imprint prov-2026-8ccbee91 (parseAnchor).
func TestParseAnchor(t *testing.T) {
	maxInt := strconv.Itoa(math.MaxInt)
	for _, c := range []struct {
		name    string
		line    string
		id      string
		version int
		ok      bool
	}{
		{"plain", "<!-- provenance:prov-2026-4c78e82d@1 -->", "prov-2026-4c78e82d", 1, true},
		{"multi digit", "<!-- provenance:prov-2026-4c78e82d@12 -->", "prov-2026-4c78e82d", 12, true},
		{"trimmed", "  \t<!-- provenance:abc@7 -->  \t", "abc", 7, true},
		{"split at the last at-sign", "<!-- provenance:a@b@3 -->", "a@b", 3, true},
		{"id with dashes and dots", "<!-- provenance:x-y.z_1@2 -->", "x-y.z_1", 2, true},
		{"largest int", "<!-- provenance:abc@" + maxInt + " -->", "abc", math.MaxInt, true},
		{"too big for an int", "<!-- provenance:abc@" + maxInt + "0 -->", "", 0, false},
		{"far too big", "<!-- provenance:abc@99999999999999999999999 -->", "", 0, false},
		{"zero", "<!-- provenance:abc@0 -->", "", 0, false},
		{"leading zero", "<!-- provenance:abc@01 -->", "", 0, false},
		{"leading zeros on zero", "<!-- provenance:abc@00 -->", "", 0, false},
		{"empty version", "<!-- provenance:abc@ -->", "", 0, false},
		{"negative version", "<!-- provenance:abc@-1 -->", "", 0, false},
		{"plus sign", "<!-- provenance:abc@+1 -->", "", 0, false},
		{"letters in version", "<!-- provenance:abc@1a -->", "", 0, false},
		{"decimal point", "<!-- provenance:abc@1.5 -->", "", 0, false},
		{"non-ascii digit", "<!-- provenance:abc@\u0663 -->", "", 0, false},
		{"no at-sign", "<!-- provenance:abc -->", "", 0, false},
		{"empty id", "<!-- provenance:@1 -->", "", 0, false},
		{"whitespace in id", "<!-- provenance:a b@1 -->", "", 0, false},
		{"tab in id", "<!-- provenance:a\tb@1 -->", "", 0, false},
		{"no space after opening", "<!--provenance:abc@1 -->", "", 0, false},
		{"two spaces after opening", "<!--  provenance:abc@1 -->", "", 0, false},
		{"no space before closing", "<!-- provenance:abc@1-->", "", 0, false},
		{"two spaces before closing", "<!-- provenance:abc@1  -->", "", 0, false},
		{"not closed", "<!-- provenance:abc@1", "", 0, false},
		{"trailing text", "<!-- provenance:abc@1 --> text", "", 0, false},
		{"leading text", "text <!-- provenance:abc@1 -->", "", 0, false},
		{"wrong keyword", "<!-- Provenance:abc@1 -->", "", 0, false},
		{"other comment", "<!-- note -->", "", 0, false},
		{"empty line", "", "", 0, false},
		{"plain text", "Some text", "", 0, false},
	} {
		t.Run(c.name, func(t *testing.T) {
			id, v, ok := parseAnchor(c.line)
			if ok != c.ok {
				t.Fatalf("parseAnchor(%q) ok = %v, want %v", c.line, ok, c.ok)
			}
			if ok && (id != c.id || v != c.version) {
				t.Fatalf("parseAnchor(%q) = %q, %d; want %q, %d", c.line, id, v, c.id, c.version)
			}
		})
	}
}

// wantSec is the expected shape of a Section. Constraints is compared by
// content (nil and empty are the same) unless NilConstraints is set.
type wantSec struct {
	Heading        string
	Level          int
	Line           int
	ID             string
	Version        int
	Anchored       bool
	Intent         string
	Constraints    []string
	HasConstraints bool
	NilConstraints bool
}

func checkSections(t *testing.T, got []Section, want []wantSec) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("%d sections, want %d: %+v", len(got), len(want), got)
	}
	for i, w := range want {
		g := got[i]
		if g.Heading != w.Heading || g.Level != w.Level || g.Line != w.Line || g.ID != w.ID ||
			g.Version != w.Version || g.Anchored != w.Anchored || g.Intent != w.Intent ||
			g.HasConstraints != w.HasConstraints || !slices.Equal(g.Constraints, w.Constraints) {
			t.Errorf("section %d:\n got  %#v\n want %#v", i, g, w)
		}
		if w.NilConstraints && g.Constraints != nil {
			t.Errorf("section %d: Constraints must be nil, got %#v", i, g.Constraints)
		}
	}
}

// unanchored is the expected section for a heading that is not anchored: only
// Heading, Level and Line are set.
func unanchored(heading string, level, line int) wantSec {
	return wantSec{Heading: heading, Level: level, Line: line, NilConstraints: true}
}

const parseDoc = "# Document title\n" + // 1
	"\n" + // 2
	"## Alpha\n" + // 3
	"<!-- provenance:prov-2026-00000001@2 -->\n" + // 4
	"\n" + // 5
	"Alpha intent line one\n" + // 6
	"line two\n" + // 7
	"\n" + // 8
	"**Constraints**\n" + // 9
	"- first\n" + // 10
	"- second\n" + // 11
	"  continued\n" + // 12
	"\n" + // 13
	"### Beta\n" + // 14
	"<!-- provenance:prov-2026-00000002@10 -->\n" + // 15
	"\n" + // 16
	"**Constraints**\n" // 17

func parseDocWant() []wantSec {
	return []wantSec{
		{Heading: "Alpha", Level: 2, Line: 3, ID: "prov-2026-00000001", Version: 2, Anchored: true,
			Intent: "Alpha intent line one\nline two", Constraints: []string{"first", "second\ncontinued"}, HasConstraints: true},
		{Heading: "Beta", Level: 3, Line: 14, ID: "prov-2026-00000002", Version: 10, Anchored: true,
			Intent: "", HasConstraints: true},
	}
}

// TestParseSections pins imprint prov-2026-6f0b5e46 (Parse): sections in document
// order, level, one-based line, heading, anchor, intent, constraints, text
// before the first heading ignored.
func TestParseSections(t *testing.T) {
	got, err := Parse(parseDoc)
	if err != nil {
		t.Fatal(err)
	}
	checkSections(t, got, parseDocWant())
}

// TestParseLineEndings pins blueprint prov-2026-632550bc and imprint
// prov-2026-6f0b5e46: newline and carriage-return-plus-newline endings give the
// same sections, including line numbers, with and without a final newline.
func TestParseLineEndings(t *testing.T) {
	for name, doc := range map[string]string{
		"crlf":                 strings.ReplaceAll(parseDoc, "\n", "\r\n"),
		"lf without final eol": strings.TrimSuffix(parseDoc, "\n"),
		"crlf without final":   strings.TrimSuffix(strings.ReplaceAll(parseDoc, "\n", "\r\n"), "\r\n"),
	} {
		t.Run(name, func(t *testing.T) {
			got, err := Parse(doc)
			if err != nil {
				t.Fatal(err)
			}
			checkSections(t, got, parseDocWant())
		})
	}
	// A blank continuation line (two spaces) survives CRLF.
	doc := "## T\r\n<!-- provenance:abc@1 -->\r\n\r\n**Constraints**\r\n- a\r\n  \r\n  b\r\n"
	got, _ := Parse(doc)
	checkSections(t, got, []wantSec{{Heading: "T", Level: 2, Line: 1, ID: "abc", Version: 1, Anchored: true,
		Constraints: []string{"a\n\nb"}, HasConstraints: true}})
}

// TestParseEmpty pins imprint prov-2026-6f0b5e46: text with no heading has no
// section, and no error.
func TestParseEmpty(t *testing.T) {
	for _, doc := range []string{"", "\n", "just text\nmore text\n", "# one hash\n#### four hashes\n##no space\n  ## indented\n"} {
		got, err := Parse(doc)
		if err != nil || len(got) != 0 {
			t.Errorf("Parse(%q) = %+v, %v", doc, got, err)
		}
	}
}

// TestParseHeadings pins imprint prov-2026-6f0b5e46: a section starts only at a
// line with two or three number signs and a space at column zero; the heading
// is the text after them, trimmed.
func TestParseHeadings(t *testing.T) {
	doc := "## Two\n" + // 1
		"<!-- provenance:a@1 -->\n" + // 2
		"\n" + // 3
		"#### four hashes stay in the intent\n" + // 4
		"##no space stays\n" + // 5
		" ## indented stays\n" + // 6
		"# one hash stays\n" + // 7
		"**Constraints**\n" + // 8
		"###   Three, with spaces   \n" + // 9
		"<!-- provenance:b@1 -->\n" + // 10
		"##  \n" // 11
	got, err := Parse(doc)
	if err != nil {
		t.Fatal(err)
	}
	checkSections(t, got, []wantSec{
		{Heading: "Two", Level: 2, Line: 1, ID: "a", Version: 1, Anchored: true,
			Intent: "#### four hashes stay in the intent\n##no space stays\n ## indented stays\n# one hash stays", HasConstraints: true},
		{Heading: "Three, with spaces", Level: 3, Line: 9, ID: "b", Version: 1, Anchored: true},
		unanchored("", 2, 11),
	})
}

// TestParseUnanchored pins imprint prov-2026-6f0b5e46: a section is Anchored
// only when the first non-blank line after the heading is accepted by
// parseAnchor; otherwise ID, Version, Intent, Constraints and HasConstraints
// are all zero even if a constraints marker follows.
func TestParseUnanchored(t *testing.T) {
	doc := "## No lines after\n" + // 1
		"## Text first\n" + // 2
		"Some intent\n" + // 3
		"<!-- provenance:late@1 -->\n" + // 4
		"**Constraints**\n" + // 5
		"- c\n" + // 6
		"### Bad anchor spacing\n" + // 7
		"<!--provenance:x@1 -->\n" + // 8
		"**Constraints**\n" + // 9
		"- c\n" + // 10
		"## Version zero\n" + // 11
		"<!-- provenance:x@0 -->\n" + // 12
		"**Constraints**\n" + // 13
		"## Blank lines only\n" + // 14
		"\n" + // 15
		"   \n" + // 16
		"### Marker first\n" + // 17
		"**Constraints**\n" + // 18
		"- c\n" // 19
	got, err := Parse(doc)
	if err != nil {
		t.Fatal(err)
	}
	checkSections(t, got, []wantSec{
		unanchored("No lines after", 2, 1),
		unanchored("Text first", 2, 2),
		unanchored("Bad anchor spacing", 3, 7),
		unanchored("Version zero", 2, 11),
		unanchored("Blank lines only", 2, 14),
		unanchored("Marker first", 3, 17),
	})
}

// TestParseAnchorAfterBlankLines pins imprint prov-2026-6f0b5e46: blank lines
// between heading and anchor are skipped; the anchor line is trimmed by
// parseAnchor.
func TestParseAnchorAfterBlankLines(t *testing.T) {
	doc := "## T\n\n   \n\t\n  <!-- provenance:abc@3 -->  \n\nIntent\n\n**Constraints**\n- c\n"
	got, _ := Parse(doc)
	checkSections(t, got, []wantSec{{Heading: "T", Level: 2, Line: 1, ID: "abc", Version: 3, Anchored: true,
		Intent: "Intent", Constraints: []string{"c"}, HasConstraints: true}})
}

// TestParseWithoutConstraintsMarker pins imprint prov-2026-6f0b5e46: without a
// Constraints line the intent is all lines after the anchor, HasConstraints is
// false and Constraints is nil.
func TestParseWithoutConstraintsMarker(t *testing.T) {
	doc := "## T\n<!-- provenance:abc@1 -->\n\nLine one\n\n- looks like an item\nLine three\n\n## Next\n<!-- provenance:def@1 -->\n**Constraints** extra\n"
	got, _ := Parse(doc)
	checkSections(t, got, []wantSec{
		{Heading: "T", Level: 2, Line: 1, ID: "abc", Version: 1, Anchored: true,
			Intent: "Line one\n\n- looks like an item\nLine three", NilConstraints: true},
		{Heading: "Next", Level: 2, Line: 9, ID: "def", Version: 1, Anchored: true,
			Intent: "**Constraints** extra", NilConstraints: true},
	})
}

// TestParseIntentAndMarker pins imprint prov-2026-6f0b5e46: the Constraints
// line is the first line equal to **Constraints** after trimming; the intent is
// the lines before it, unescaped, joined with newlines and trimmed.
func TestParseIntentAndMarker(t *testing.T) {
	doc := "## T\n<!-- provenance:abc@1 -->\n\n" +
		"  first line indented\n" +
		"\n" +
		`\## escaped heading` + "\n" +
		`\\## double` + "\n" +
		`\<!-- provenance:x@1 -->` + "\n" +
		`\**Constraints**` + "\n" +
		`\plain backslash` + "\n" +
		"last  \n" +
		"\n" +
		"  **Constraints**  \n" +
		"- c\n"
	got, _ := Parse(doc)
	wantIntent := "first line indented\n\n## escaped heading\n" + `\## double` + "\n<!-- provenance:x@1 -->\n**Constraints**\n" + `\plain backslash` + "\nlast"
	checkSections(t, got, []wantSec{{Heading: "T", Level: 2, Line: 1, ID: "abc", Version: 1, Anchored: true,
		Intent: wantIntent, Constraints: []string{"c"}, HasConstraints: true}})
}

// TestParseConstraints pins imprint prov-2026-6f0b5e46: a constraint starts at a
// line with a hyphen and a space; lines beginning with two spaces continue it
// (two spaces removed, a line of only two spaces adds an empty line); the whole
// constraint is trimmed after joining; other lines are ignored; a completely
// empty line between constraints is skipped.
func TestParseConstraints(t *testing.T) {
	for _, c := range []struct {
		name string
		body string // lines after the Constraints line
		want []string
	}{
		{"two items", "- a\n- b\n", []string{"a", "b"}},
		{"empty line between items", "- a\n\n- b\n", []string{"a", "b"}},
		{"continuation", "- a\n  b\n  c\n", []string{"a\nb\nc"}},
		{"blank line inside", "- a\n  \n  b\n", []string{"a\n\nb"}},
		{"two blank lines inside", "- a\n  \n  \n  b\n", []string{"a\n\n\nb"}},
		{"trailing blank continuation is trimmed", "- a\n  \n  \n", []string{"a"}},
		{"trimmed only after joining", "- a  \n  b  \n", []string{"a  \nb"}},
		{"leading spaces after hyphen", "-   padded\n", []string{"padded"}},
		{"indent beyond two spaces is kept", "- a\n      deeper\n", []string{"a\n    deeper"}},
		{"continuation that looks like an item", "- a\n  - nested\n", []string{"a\n- nested"}},
		{"continuation that looks like structure", "- a\n  ## h\n  <!-- provenance:x@1 -->\n  **Constraints**\n", []string{"a\n## h\n<!-- provenance:x@1 -->\n**Constraints**"}},
		{"item that looks like structure", "- ## h\n- **Constraints**\n", []string{"## h", "**Constraints**"}},
		{"constraint text is not unescaped", "- " + `\## x` + "\n  " + `\## y` + "\n", []string{`\## x` + "\n" + `\## y`}},
		{"hyphen without space is ignored", "-a\n- b\n", []string{"b"}},
		{"star bullet is ignored", "* a\n- b\n", []string{"b"}},
		{"other line is ignored", "text\n- a\ntext\n- b\n", []string{"a", "b"}},
		{"second marker line is ignored", "- a\n**Constraints**\n- b\n", []string{"a", "b"}},
		{"no items", "\n", nil},
	} {
		t.Run(c.name, func(t *testing.T) {
			doc := "## T\n<!-- provenance:abc@1 -->\n\nIntent\n\n**Constraints**\n" + c.body
			got, err := Parse(doc)
			if err != nil {
				t.Fatal(err)
			}
			checkSections(t, got, []wantSec{{Heading: "T", Level: 2, Line: 1, ID: "abc", Version: 1, Anchored: true,
				Intent: "Intent", Constraints: c.want, HasConstraints: true}})
		})
	}
}

// TestParseSectionEnds pins imprint prov-2026-6f0b5e46: a section runs until the
// next section-starting line, so constraints after a marker end at the next
// heading and the next section is read on its own.
func TestParseSectionEnds(t *testing.T) {
	doc := "## A\n<!-- provenance:a@1 -->\n\n**Constraints**\n- a1\n- a2\n### B\n<!-- provenance:b@2 -->\n\nB intent\n\n**Constraints**\n- b1\n## C\n"
	got, _ := Parse(doc)
	checkSections(t, got, []wantSec{
		{Heading: "A", Level: 2, Line: 1, ID: "a", Version: 1, Anchored: true, Constraints: []string{"a1", "a2"}, HasConstraints: true},
		{Heading: "B", Level: 3, Line: 7, ID: "b", Version: 2, Anchored: true, Intent: "B intent", Constraints: []string{"b1"}, HasConstraints: true},
		unanchored("C", 2, 14),
	})
}

// TestParseReadsTextOnly pins blueprint prov-2026-632550bc: Parse is a pure
// function of its text; the same text gives the same sections every time.
func TestParseReadsTextOnly(t *testing.T) {
	a, err1 := Parse(parseDoc)
	b, err2 := Parse(parseDoc)
	if err1 != nil || err2 != nil || !reflect.DeepEqual(a, b) {
		t.Fatalf("Parse is not deterministic: %v %v", err1, err2)
	}
}
