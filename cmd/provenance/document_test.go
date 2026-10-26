package main

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/spf13/pflag"

	"github.com/livecodelife/provenance/document"
	"github.com/livecodelife/provenance/record"
)

// TestDocumentCommandShape pins imprints prov-2026-36b81f10 (newDocumentCmd) and
// prov-2026-7931302e: use, short, no arguments, subcommands, flags with type,
// shorthand, default and usage text, and the registration order.
func TestDocumentCommandShape(t *testing.T) {
	cmd := newDocumentCmd()
	if cmd.Use != "document" {
		t.Errorf("Use %q", cmd.Use)
	}
	if want := "Print the behaviour document of the briefs and blueprints, or apply an edited one"; cmd.Short != want {
		t.Errorf("Short %q", cmd.Short)
	}
	if cmd.Args == nil || cmd.Args(cmd, nil) != nil || cmd.Args(cmd, []string{"x"}) == nil {
		t.Error("the command takes no positional arguments")
	}
	if len(cmd.Commands()) != 0 {
		t.Error("the command has no subcommands")
	}
	if cmd.RunE == nil && cmd.Run == nil {
		t.Error("no run function")
	}
	type fl struct{ name, short, typ, def, usage string }
	want := []fl{
		{"scope", "", "stringSlice", "[]", "Only records whose affected_scope covers this path (repeatable)"},
		{"format", "", "string", "md", "Output format (only md)"},
		{"out", "o", "string", "", "Write the document to this file instead of standard output"},
		{"ids", "", "bool", "false", "Show record ids in the headings"},
		{"apply", "", "string", "", "Plan the changes made in this edited document file"},
		{"write", "", "bool", "false", "Apply the plan to the records (needs --apply and --reason)"},
		{"reason", "", "string", "", "Why these revisions exist (required with --write)"},
		{"skip-conflicts", "", "bool", "false", "With --write, apply the other changes and skip conflicting sections"},
	}
	cmd.Flags().SortFlags = false
	var got []fl
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		got = append(got, fl{f.Name, f.Shorthand, f.Value.Type(), f.DefValue, f.Usage})
	})
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("flags (in registration order):\n%+v\nwant:\n%+v", got, want)
	}
}

// TestDocumentIsRegisteredOnRoot pins imprint prov-2026-36b81f10: the command is
// a top-level command of the root, and a positional argument is refused.
func TestDocumentIsRegisteredOnRoot(t *testing.T) {
	c, _, err := newRootCmd().Find([]string{"document"})
	if err != nil || c == nil || c.Use != "document" || c.Parent() == nil {
		t.Fatalf("Find: %v, %v", c, err)
	}
	_, cfg := docRepo(t, nil)
	if out, err := run(t, "-c", cfg, "document", "extra"); err == nil {
		t.Fatalf("a positional argument was accepted:\n%s", out)
	}
}

// TestDocumentFlagErrors pins imprint prov-2026-cabce335: every flag-check error
// message, the check order (format, apply combined, write without apply, reason
// without write, skip-conflicts without write, empty reason), and that all of
// them are raised before the repository is opened (the config here cannot be
// read, yet the flag error is what comes back).
func TestDocumentFlagErrors(t *testing.T) {
	cfg := missingConfig(t)
	for _, c := range []struct {
		name string
		args []string
		want string
	}{
		{"format", []string{"--format", "txt"}, `unsupported format "txt": only md is accepted`},
		{"format is case sensitive", []string{"--format", "MD"}, `unsupported format "MD": only md is accepted`},
		{"format empty", []string{"--format", ""}, `unsupported format "": only md is accepted`},
		{"format before apply combined", []string{"--format", "html", "--apply", "f", "--scope", "x"}, `unsupported format "html": only md is accepted`},
		{"apply with scope", []string{"--apply", "f", "--scope", "x"}, "--apply cannot be combined with --scope, --ids or -o"},
		{"apply with ids", []string{"--apply", "f", "--ids"}, "--apply cannot be combined with --scope, --ids or -o"},
		{"apply with out", []string{"--apply", "f", "-o", "x.md"}, "--apply cannot be combined with --scope, --ids or -o"},
		{"apply combined before write checks", []string{"--apply", "f", "--ids", "--write"}, "--apply cannot be combined with --scope, --ids or -o"},
		{"write without apply", []string{"--write"}, "--write needs --apply"},
		{"write without apply before reason", []string{"--write", "--reason", "r", "--skip-conflicts"}, "--write needs --apply"},
		{"reason without write", []string{"--reason", "because"}, "--reason needs --write"},
		{"blank reason without write", []string{"--reason", "  "}, "--reason needs --write"},
		{"reason without write with apply", []string{"--apply", "f", "--reason", "r"}, "--reason needs --write"},
		{"reason before skip-conflicts", []string{"--reason", "r", "--skip-conflicts"}, "--reason needs --write"},
		{"skip-conflicts without write", []string{"--skip-conflicts"}, "--skip-conflicts needs --write"},
		{"skip-conflicts without write with apply", []string{"--apply", "f", "--skip-conflicts"}, "--skip-conflicts needs --write"},
		{"empty reason", []string{"--apply", "f", "--write"}, "--write needs a non-empty --reason"},
		{"blank reason", []string{"--apply", "f", "--write", "--reason", " \t "}, "--write needs a non-empty --reason"},
		{"empty reason with skip-conflicts", []string{"--apply", "f", "--write", "--skip-conflicts"}, "--write needs a non-empty --reason"},
	} {
		t.Run(c.name, func(t *testing.T) {
			out, err := run(t, append([]string{"-c", cfg, "document"}, c.args...)...)
			if err == nil || err.Error() != c.want {
				t.Fatalf("error %v, want %q", err, c.want)
			}
			if out != "" {
				t.Fatalf("stdout %q", out)
			}
		})
	}
	// An explicit --format md and a bad config: only then does openRepo fail.
	if _, err := run(t, "-c", cfg, "document", "--format", "md"); err == nil || strings.Contains(err.Error(), "needs") || strings.Contains(err.Error(), "unsupported") {
		t.Fatalf("a valid combination must reach openRepo: %v", err)
	}
}

// libRender renders the repository through the library, the oracle for stdout.
func libRender(t *testing.T, root string, o document.RenderOptions) string {
	t.Helper()
	out, err := document.Render(&record.FilesStore{Root: root}, o)
	if err != nil {
		t.Fatal(err)
	}
	return out
}

// TestDocumentPrintsToStdout pins blueprint prov-2026-65888774 and imprint
// prov-2026-cabce335: without --apply the document built by the library is
// written to standard output, and nothing else is.
func TestDocumentPrintsToStdout(t *testing.T) {
	root, cfg := docRepo(t, nil)
	before := treeOf(t, root)
	out := mustRun(t, "-c", cfg, "document")
	want := "## Brief\n<!-- provenance:" + docBrief + "@1 -->\n\nBrief intent\n\n**Constraints**\n- B constraint\n\n" +
		"### Blueprint one\n<!-- provenance:" + docBP1 + "@1 -->\n\nBlueprint one intent\n\n**Constraints**\n- P1 constraint\n\n" +
		"### Blueprint two\n<!-- provenance:" + docBP2 + "@1 -->\n\nBlueprint two intent\n\n**Constraints**\n- P2 constraint\n"
	if out != want {
		t.Fatalf("stdout:\n%q\nwant:\n%q", out, want)
	}
	if out != libRender(t, root, document.RenderOptions{}) {
		t.Fatal("stdout differs from the library's Render")
	}
	if got := mustRun(t, "-c", cfg, "document", "--format", "md"); got != want {
		t.Fatal("--format md changes the output")
	}
	sameTreeOf(t, before, treeOf(t, root))
}

// TestDocumentIDsAndScope pins blueprint prov-2026-65888774 and imprint
// prov-2026-cabce335: --ids and the repeatable --scope are passed to Render as
// IDs and Scopes.
func TestDocumentIDsAndScope(t *testing.T) {
	other := docRecord("prov-2026-0000d0e1", "brief", "Elsewhere", "Other intent", "", "O")
	other = strings.Replace(other, "  - a.go\n", "  - pkg/*.go\n", 1)
	root, cfg := docRepo(t, map[string]string{"prov-2026-0000d0e1": other})

	out := mustRun(t, "-c", cfg, "document", "--ids")
	if !strings.Contains(out, "## Brief ("+docBrief+")\n") || !strings.Contains(out, "### Blueprint one ("+docBP1+")\n") {
		t.Fatalf("--ids:\n%s", out)
	}
	if out != libRender(t, root, document.RenderOptions{IDs: true}) {
		t.Fatal("--ids output differs from the library's")
	}
	for _, c := range []struct {
		args []string
		opts document.RenderOptions
		want []string
	}{
		{[]string{"--scope", "a.go"}, document.RenderOptions{Scopes: []string{"a.go"}}, []string{"## Brief", "### Blueprint one", "### Blueprint two"}},
		{[]string{"--scope", "pkg/x.go"}, document.RenderOptions{Scopes: []string{"pkg/x.go"}}, []string{"## Elsewhere"}},
		{[]string{"--scope", "a.go", "--scope", "pkg/x.go", "--ids"}, document.RenderOptions{Scopes: []string{"a.go", "pkg/x.go"}, IDs: true},
			[]string{"## Brief (" + docBrief + ")", "### Blueprint one (" + docBP1 + ")", "### Blueprint two (" + docBP2 + ")", "## Elsewhere (prov-2026-0000d0e1)"}},
	} {
		got := mustRun(t, append([]string{"-c", cfg, "document"}, c.args...)...)
		if got != libRender(t, root, c.opts) {
			t.Errorf("%v: output differs from the library's", c.args)
		}
		var heads []string
		for _, l := range strings.Split(got, "\n") {
			if strings.HasPrefix(l, "## ") || strings.HasPrefix(l, "### ") {
				heads = append(heads, l)
			}
		}
		if !reflect.DeepEqual(heads, c.want) {
			t.Errorf("%v: headings %q, want %q", c.args, heads, c.want)
		}
	}
	// nothing matches: empty output, success
	if got, err := run(t, "-c", cfg, "document", "--scope", "zzz/none.go"); err != nil || got != "" {
		t.Fatalf("empty document: %q, %v", got, err)
	}
}

// TestDocumentOutFile pins blueprint prov-2026-65888774 and imprint
// prov-2026-cabce335: -o writes the document to the file with permission 0644,
// overwrites an existing file, prints nothing, and relative paths are relative
// to the working directory.
func TestDocumentOutFile(t *testing.T) {
	withUmask(t, 0o022)
	root, cfg := docRepo(t, nil)
	want := libRender(t, root, document.RenderOptions{})
	dir := t.TempDir()
	p := filepath.Join(dir, "doc.md")
	writeFile(t, p, strings.Repeat("old content that is longer than the document\n", 400))
	out, err := run(t, "-c", cfg, "document", "-o", p)
	if err != nil || out != "" {
		t.Fatalf("stdout %q, %v", out, err)
	}
	if got := readFile(t, p); got != want {
		t.Fatalf("file:\n%q\nwant:\n%q", got, want)
	}
	if info, _ := os.Stat(p); info.Mode().Perm() != 0o644 {
		t.Fatalf("mode %v", info.Mode().Perm())
	}
	// long form, new file, with --ids, relative to the working directory
	wd, _ := os.Getwd()
	if err := os.Chdir(dir); err != nil {
		t.Fatal(err)
	}
	defer os.Chdir(wd)
	if out, err := run(t, "-c", cfg, "document", "--out", "rel.md", "--ids"); err != nil || out != "" {
		t.Fatalf("%q, %v", out, err)
	}
	if got := readFile(t, filepath.Join(dir, "rel.md")); got != libRender(t, root, document.RenderOptions{IDs: true}) {
		t.Fatalf("relative file:\n%s", got)
	}
}

// TestDocumentOutFileError pins imprint prov-2026-cabce335: an unwritable -o
// gives "cannot write PATH: ERR" with the system error text, and prints nothing.
func TestDocumentOutFileError(t *testing.T) {
	_, cfg := docRepo(t, nil)
	p := filepath.Join(t.TempDir(), "no-such-dir", "doc.md")
	sys := os.WriteFile(p, []byte("x"), 0o644)
	if sys == nil {
		t.Fatal("fixture path is writable")
	}
	out, err := run(t, "-c", cfg, "document", "-o", p)
	if err == nil || err.Error() != "cannot write "+p+": "+sys.Error() || out != "" {
		t.Fatalf("%q, %v (system error %q)", out, err, sys)
	}
}

// TestDocumentRenderFailure pins imprint prov-2026-cabce335: when Render fails the
// error is returned and nothing is written, neither to stdout nor to -o.
func TestDocumentRenderFailure(t *testing.T) {
	bad := strings.Replace(docRecord("prov-2026-0000d0e1", "brief", "Bad", "i", ""), "  - a.go\n", "  - \"re:(\"\n", 1)
	_, cfg := docRepo(t, map[string]string{"prov-2026-0000d0e1": bad})
	p := filepath.Join(t.TempDir(), "doc.md")
	out, err := run(t, "-c", cfg, "document", "--scope", "a.go", "-o", p)
	if err == nil || !strings.Contains(err.Error(), "invalid regex") || out != "" {
		t.Fatalf("%q, %v", out, err)
	}
	if _, err := os.Stat(p); err == nil {
		t.Fatal("-o file was written")
	}
	out, err = run(t, "-c", cfg, "document", "--scope", "a.go")
	if err == nil || out != "" {
		t.Fatalf("%q, %v", out, err)
	}
	// a corrupt record file fails the same way
	_, cfg = docRepo(t, map[string]string{"prov-2026-0000d0e2": "{{{ not a record\n"})
	if out, err = run(t, "-c", cfg, "document"); err == nil || out != "" {
		t.Fatalf("corrupt record: %q, %v", out, err)
	}
}

// planLines is the printed plan for a one-change repository.
func planSummary(changed, conflicts, problems, unchanged int) string {
	return strconv.Itoa(changed) + " changed, " + strconv.Itoa(conflicts) + " conflicts, " + strconv.Itoa(problems) + " problems, " + strconv.Itoa(unchanged) + " unchanged\n"
}

// TestDocumentApplyPlansOnly pins blueprint prov-2026-65888774 and imprint
// prov-2026-cabce335: --apply reads the file, prints the plan, writes nothing
// without --write, and returns nil when the plan is clean.
func TestDocumentApplyPlansOnly(t *testing.T) {
	root, cfg := docRepo(t, nil, docBrief, docBP1, docBP2)
	doc := mustRun(t, "-c", cfg, "document")
	edited := editDoc(t, editDoc(t, doc, "Brief intent", "Brief intent, edited"), "- P1 constraint", "- P1 constraint\n- P1 added")
	file := docFile(t, "edited.md", edited)
	before := treeOf(t, root)
	out, err := run(t, "-c", cfg, "document", "--apply", file)
	want := docBrief + " change intent\n" + docBP1 + " change constraints\n" + planSummary(2, 0, 0, 1)
	if err != nil || out != want {
		t.Fatalf("%q, %v\nwant %q", out, err, want)
	}
	sameTreeOf(t, before, treeOf(t, root))

	both := editDoc(t, editDoc(t, doc, "Blueprint two intent", "changed"), "- P2 constraint", "- P2 changed")
	out, err = run(t, "-c", cfg, "document", "--apply", docFile(t, "e2.md", both), "--format", "md")
	if err != nil || out != docBP2+" change intent and constraints\n"+planSummary(1, 0, 0, 2) {
		t.Fatalf("%q, %v", out, err)
	}
	// an unedited document: nothing to do
	out, err = run(t, "-c", cfg, "document", "--apply", docFile(t, "same.md", doc))
	if err != nil || out != planSummary(0, 0, 0, 3) {
		t.Fatalf("%q, %v", out, err)
	}
	sameTreeOf(t, before, treeOf(t, root))
}

// TestDocumentApplyRelativePath pins imprint prov-2026-cabce335: the apply path
// is used as given, relative to the working directory.
func TestDocumentApplyRelativePath(t *testing.T) {
	_, cfg := docRepo(t, nil)
	doc := mustRun(t, "-c", cfg, "document")
	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, "d.md"), doc)
	wd, _ := os.Getwd()
	if err := os.Chdir(dir); err != nil {
		t.Fatal(err)
	}
	defer os.Chdir(wd)
	out, err := run(t, "-c", cfg, "document", "--apply", "d.md")
	if err != nil || out != planSummary(0, 0, 0, 3) {
		t.Fatalf("%q, %v", out, err)
	}
}

// TestDocumentApplyFileUnreadable pins imprint prov-2026-cabce335: the message
// "cannot read PATH: ERR" with the system error text, and nothing printed.
func TestDocumentApplyFileUnreadable(t *testing.T) {
	_, cfg := docRepo(t, nil)
	p := filepath.Join(t.TempDir(), "missing.md")
	_, sys := os.ReadFile(p)
	out, err := run(t, "-c", cfg, "document", "--apply", p)
	if err == nil || err.Error() != "cannot read "+p+": "+sys.Error() || out != "" {
		t.Fatalf("%q, %v", out, err)
	}
}

// TestDocumentApplyPlanFailure pins imprint prov-2026-cabce335: when BuildPlan
// fails its error is returned and no plan is printed.
func TestDocumentApplyPlanFailure(t *testing.T) {
	_, cfg := docRepo(t, map[string]string{"prov-2026-0000d0e2": "{{{ not a record\n"})
	out, err := run(t, "-c", cfg, "document", "--apply", docFile(t, "d.md", "### Corrupt\n<!-- provenance:prov-2026-0000d0e2@1 -->\n\n**Constraints**\n"))
	if err == nil || out != "" {
		t.Fatalf("%q, %v", out, err)
	}
}

// TestDocumentApplyExitsNonZeroOnProblems pins blueprint prov-2026-65888774 and
// imprints prov-2026-cabce335, prov-2026-57558206: even when only planning, a
// plan with conflicts or problems prints the plan and returns the planError.
func TestDocumentApplyExitsNonZeroOnProblems(t *testing.T) {
	root, cfg := docRepo(t, nil, docBrief, docBP1, docBP2)
	doc := mustRun(t, "-c", cfg, "document")
	before := treeOf(t, root)

	bad := editDoc(t, doc, "Brief intent", "Edited") + "## Floating\ntext\n\n### Ghost\n<!-- provenance:prov-2026-0000ffff@1 -->\n\n**Constraints**\n"
	line := strings.Count(editDoc(t, doc, "Brief intent", "Edited"), "\n") + 1 // "## Floating" follows the document directly
	out, err := run(t, "-c", cfg, "document", "--apply", docFile(t, "bad.md", bad))
	want := "- unanchored line " + strconv.Itoa(line) + ": heading has no anchor\n" +
		docBrief + " change intent\n" +
		"prov-2026-0000ffff unknown line " + strconv.Itoa(line+3) + ": no record with this id\n" +
		planSummary(1, 0, 2, 2)
	if out != want {
		t.Fatalf("stdout:\n%q\nwant:\n%q", out, want)
	}
	if err == nil || err.Error() != "document plan has 0 conflicts and 2 problems" {
		t.Fatalf("error %v", err)
	}

	bumpRecord(t, root, docBP1)
	before = treeOf(t, root)
	out, err = run(t, "-c", cfg, "document", "--apply", docFile(t, "same.md", doc))
	want = docBP1 + " conflict anchored v1 current v2\n" + planSummary(0, 1, 0, 2)
	if out != want || err == nil || err.Error() != "document plan has 1 conflicts and 0 problems" {
		t.Fatalf("%q, %v", out, err)
	}
	sameTreeOf(t, before, treeOf(t, root))
}

// TestDocumentWriteDrafts pins blueprint prov-2026-65888774 and imprint
// prov-2026-cabce335: --write applies the plan, drafts are saved in place and
// printed as "saved ID (draft)" after the plan, the manifest is rewritten, and
// only the changed record is touched.
func TestDocumentWriteDrafts(t *testing.T) {
	root, cfg := docRepo(t, nil)
	doc := mustRun(t, "-c", cfg, "document")
	file := docFile(t, "e.md", editDoc(t, doc, "Blueprint one intent", "Rewritten"))
	before := treeOf(t, root)
	out, err := run(t, "-c", cfg, "document", "--apply", file, "--write", "--reason", "Why")
	want := docBP1 + " change intent\n" + planSummary(1, 0, 0, 2) + "saved " + docBP1 + " (draft)\n"
	if err != nil || out != want {
		t.Fatalf("%q, %v\nwant %q", out, err, want)
	}
	rec := load(t, root, docBP1)
	if rec.Intent != "Rewritten" || rec.Version != 1 || rec.Reason != "" || len(versions(t, root, docBP1)) != 0 {
		t.Fatalf("record %+v", rec)
	}
	after := treeOf(t, root)
	for _, id := range []string{docBrief, docBP2} {
		p := filepath.Join("provenance", id+".yml")
		if after[p] != before[p] {
			t.Errorf("%s was touched", id)
		}
	}
	if _, ok := after[manifestPath]; !ok {
		t.Error("the manifest was not rewritten after a revision")
	}
}

// TestDocumentWriteRevisesVersions pins blueprint prov-2026-65888774 and imprints
// prov-2026-cabce335, prov-2026-6b2ceecb: non-draft records get a new version with
// the --reason and the repository's author, the line is "revised ID vN TAG", the
// manifest is rewritten, lines follow the order of the changes, and records not
// changed are left alone.
func TestDocumentWriteRevisesVersions(t *testing.T) {
	root, cfg := docRepo(t, nil, docBrief, docBP1)
	doc := mustRun(t, "-c", cfg, "document")
	edited := editDoc(t, editDoc(t, doc, "Brief intent", "Brief intent v2"), "- P1 constraint", "- P1 constraint v2")
	// docBP2 stays a draft and is edited as well: drafts and versions mix
	edited = editDoc(t, edited, "Blueprint two intent", "Two v2")
	out, err := run(t, "-c", cfg, "document", "--apply", docFile(t, "e.md", edited), "--write", "--reason", "Spec moved")
	want := docBrief + " change intent\n" + docBP1 + " change constraints\n" + docBP2 + " change intent\n" + planSummary(3, 0, 0, 0) +
		"revised " + docBrief + " v2 [" + docBrief + "@2]\n" +
		"revised " + docBP1 + " v2 [" + docBP1 + "@2]\n" +
		"saved " + docBP2 + " (draft)\n"
	if err != nil || out != want {
		t.Fatalf("%q, %v\nwant %q", out, err, want)
	}
	r := load(t, root, docBrief)
	if r.Version != 2 || r.Reason != "Spec moved" || r.Intent != "Brief intent v2" {
		t.Fatalf("brief %+v", r)
	}
	snap, err := (&record.FilesStore{Root: root}).Snapshot(docBrief, 2)
	if err != nil || snap.Reason != "Spec moved" || snap.Author != "test@example.com" {
		t.Fatalf("snapshot %+v, %v", snap, err)
	}
	if !reflect.DeepEqual(load(t, root, docBP1).Constraints, []string{"P1 constraint v2"}) {
		t.Fatalf("blueprint constraints %q", load(t, root, docBP1).Constraints)
	}
	manifest := readFile(t, filepath.Join(root, manifestPath))
	if !strings.Contains(manifest, `"id":"`+docBrief+`","version":2`) || !strings.Contains(manifest, `"id":"`+docBP1+`","version":2`) {
		t.Fatalf("manifest not rewritten:\n%s", manifest)
	}

	// only the changed record is revised: edit one, the others keep version and bytes
	root2, cfg2 := docRepo(t, nil, docBrief, docBP1, docBP2)
	doc2 := mustRun(t, "-c", cfg2, "document")
	before2 := treeOf(t, root2)
	out, err = run(t, "-c", cfg2, "document", "--apply", docFile(t, "e.md", editDoc(t, doc2, "Blueprint two intent", "Only this")), "--write", "--reason", "One")
	if err != nil || !strings.HasSuffix(out, "revised "+docBP2+" v2 ["+docBP2+"@2]\n") || strings.Count(out, "revised ") != 1 {
		t.Fatalf("%q, %v", out, err)
	}
	after2 := treeOf(t, root2)
	for _, id := range []string{docBrief, docBP1} {
		p := filepath.Join("provenance", id+".yml")
		if after2[p] != before2[p] || len(versions(t, root2, id)) != 1 {
			t.Errorf("%s was touched", id)
		}
	}
}

// TestDocumentWriteProblemsDoNotBlock pins blueprint prov-2026-65888774 and
// imprint prov-2026-cabce335: with --write, problems other than conflicts do not
// block the valid changes; the plan, then the saved line are printed and the
// planError is returned.
func TestDocumentWriteProblemsDoNotBlock(t *testing.T) {
	root, cfg := docRepo(t, nil)
	doc := mustRun(t, "-c", cfg, "document")
	edited := editDoc(t, doc, "Brief intent", "Brief edited")
	line := strings.Count(edited, "\n") + 1
	edited += "## Extra\ntext\n"
	out, err := run(t, "-c", cfg, "document", "--apply", docFile(t, "e.md", edited), "--write", "--reason", "Why")
	want := "- unanchored line " + strconv.Itoa(line) + ": heading has no anchor\n" +
		docBrief + " change intent\n" + planSummary(1, 0, 1, 2) + "saved " + docBrief + " (draft)\n"
	if out != want || err == nil || err.Error() != "document plan has 0 conflicts and 1 problems" {
		t.Fatalf("%q, %v\nwant %q", out, err, want)
	}
	if load(t, root, docBrief).Intent != "Brief edited" {
		t.Fatal("the valid change was not applied")
	}
}

// TestDocumentWriteConflictsBlock pins blueprint prov-2026-65888774 and imprint
// prov-2026-cabce335: with --write and no --skip-conflicts a conflict blocks every
// change; Apply's error is returned as is, the plan is printed and nothing is
// written.
func TestDocumentWriteConflictsBlock(t *testing.T) {
	root, cfg := docRepo(t, nil, docBrief, docBP1, docBP2)
	doc := mustRun(t, "-c", cfg, "document")
	bumpRecord(t, root, docBP1)
	edited := editDoc(t, editDoc(t, doc, "Brief intent", "Edited"), "Blueprint one intent", "Edited too")
	before := treeOf(t, root)
	out, err := run(t, "-c", cfg, "document", "--apply", docFile(t, "e.md", edited), "--write", "--reason", "Why")
	want := docBrief + " change intent\n" + docBP1 + " conflict anchored v1 current v2\n" + planSummary(1, 1, 0, 1)
	if out != want || err == nil || err.Error() != "the plan has conflicting sections: 1 conflicts" {
		t.Fatalf("%q, %v\nwant %q", out, err, want)
	}
	sameTreeOf(t, before, treeOf(t, root))
}

// TestDocumentWriteSkipConflicts pins blueprint prov-2026-65888774 and imprints
// prov-2026-cabce335, prov-2026-57558206: --skip-conflicts applies the other changes,
// leaves the conflicting record alone, and the planError is still returned.
func TestDocumentWriteSkipConflicts(t *testing.T) {
	root, cfg := docRepo(t, nil, docBrief, docBP1, docBP2)
	doc := mustRun(t, "-c", cfg, "document")
	bumpRecord(t, root, docBP1)
	edited := editDoc(t, editDoc(t, doc, "Brief intent", "Edited"), "Blueprint one intent", "Edited too")
	bp1 := readFile(t, filepath.Join(root, "provenance", docBP1+".yml"))
	out, err := run(t, "-c", cfg, "document", "--apply", docFile(t, "e.md", edited), "--write", "--reason", "Why", "--skip-conflicts")
	want := docBrief + " change intent\n" + docBP1 + " conflict anchored v1 current v2\n" + planSummary(1, 1, 0, 1) +
		"revised " + docBrief + " v2 [" + docBrief + "@2]\n"
	if out != want || err == nil || err.Error() != "document plan has 1 conflicts and 0 problems" {
		t.Fatalf("%q, %v\nwant %q", out, err, want)
	}
	if got := readFile(t, filepath.Join(root, "provenance", docBP1+".yml")); got != bp1 {
		t.Fatal("the conflicting record was touched")
	}
	if load(t, root, docBrief).Version != 2 {
		t.Fatal("the other change was not applied")
	}
}

// TestDocumentWriteNothingRevised pins imprint prov-2026-cabce335: when Apply
// returns no Revised, the manifest is not written, and the result is the planError.
func TestDocumentWriteNothingRevised(t *testing.T) {
	root, cfg := docRepo(t, nil)
	doc := mustRun(t, "-c", cfg, "document")
	out, err := run(t, "-c", cfg, "document", "--apply", docFile(t, "same.md", doc), "--write", "--reason", "Why")
	if err != nil || out != planSummary(0, 0, 0, 3) {
		t.Fatalf("%q, %v", out, err)
	}
	if _, err := os.Stat(filepath.Join(root, manifestPath)); err == nil {
		t.Fatal("manifest written although nothing was revised")
	}
	// a conflict skipped, nothing else: still nothing revised, the manifest is untouched
	root, cfg = docRepo(t, nil, docBrief)
	doc = mustRun(t, "-c", cfg, "document")
	bumpRecord(t, root, docBrief)
	writeFile(t, filepath.Join(root, manifestPath), "garbage\n")
	out, err = run(t, "-c", cfg, "document", "--apply", docFile(t, "e.md", editDoc(t, doc, "Brief intent", "X")), "--write", "--reason", "Why", "--skip-conflicts")
	if err == nil || err.Error() != "document plan has 1 conflicts and 0 problems" || strings.Contains(out, "revised") || strings.Contains(out, "saved") {
		t.Fatalf("%q, %v", out, err)
	}
	if readFile(t, filepath.Join(root, manifestPath)) != "garbage\n" {
		t.Fatal("manifest rewritten although nothing was revised")
	}
}

// TestDocumentWritePartialFailure pins imprint prov-2026-cabce335 and blueprint
// prov-2026-65888774: when Apply fails after revising some records, the manifest
// is still rewritten for them, the revised lines are printed, then Apply's error
// is returned; later changes are untouched.
func TestDocumentWritePartialFailure(t *testing.T) {
	v1rec := v1(docV1, "blueprint", "draft", "implements: "+docBrief+"\n")
	root, cfg := docRepo(t, map[string]string{
		docV1:   v1rec,
		docLast: docRecord(docLast, "blueprint", "Last", "Last intent", docBrief, "L"),
	}, docBrief)
	doc := mustRun(t, "-c", cfg, "document")
	edited := editDoc(t, editDoc(t, editDoc(t, doc, "Brief intent", "Brief v2"), "Intent of "+docV1, "Changed v1"), "Last intent", "Never applied")
	writeFile(t, filepath.Join(root, manifestPath), "stale manifest\n")
	last := readFile(t, filepath.Join(root, "provenance", docLast+".yml"))
	v1before := readFile(t, filepath.Join(root, "provenance", docV1+".yml"))

	out, err := run(t, "-c", cfg, "document", "--apply", docFile(t, "e.md", edited), "--write", "--reason", "Why")
	wantErr := docV1 + ": " + record.ErrNotV2.Error()
	if err == nil || err.Error() != wantErr {
		t.Fatalf("error %v, want %q", err, wantErr)
	}
	wantOut := docBrief + " change intent\n" + docV1 + " change intent\n" + docLast + " change intent\n" + planSummary(3, 0, 0, 2) +
		"revised " + docBrief + " v2 [" + docBrief + "@2]\n"
	if out != wantOut {
		t.Fatalf("stdout:\n%q\nwant:\n%q", out, wantOut)
	}
	if m := readFile(t, filepath.Join(root, manifestPath)); !strings.Contains(m, `"id":"`+docBrief+`","version":2`) || strings.Contains(m, "stale") {
		t.Fatalf("manifest not rewritten after the partial failure:\n%s", m)
	}
	if readFile(t, filepath.Join(root, "provenance", docLast+".yml")) != last || readFile(t, filepath.Join(root, "provenance", docV1+".yml")) != v1before {
		t.Fatal("records after the failure were touched")
	}
}

// TestDocumentWriteManifestFailure pins imprint prov-2026-cabce335: when the
// manifest cannot be written the manifest state is restored, the revised records
// stay saved, the writeManifest error is returned and no revised line is printed.
func TestDocumentWriteManifestFailure(t *testing.T) {
	root, cfg := docRepo(t, nil)
	doc := mustRun(t, "-c", cfg, "document")
	// .provenance is a regular file, so the manifest directory cannot be made
	writeFile(t, filepath.Join(root, ".provenance"), "not a directory\n")
	out, err := run(t, "-c", cfg, "document", "--apply", docFile(t, "e.md", editDoc(t, doc, "Brief intent", "Saved anyway")), "--write", "--reason", "Why")
	if err == nil || strings.HasPrefix(err.Error(), "document plan has") {
		t.Fatalf("error %v", err)
	}
	if strings.Contains(out, "saved ") || strings.Contains(out, "revised ") {
		t.Fatalf("revised lines are printed only after the manifest is written:\n%s", out)
	}
	if load(t, root, docBrief).Intent != "Saved anyway" {
		t.Fatal("the saved record was rolled back")
	}
	if readFile(t, filepath.Join(root, ".provenance")) != "not a directory\n" {
		t.Fatal("manifest state not restored")
	}
}

// TestPrintDocumentPlan pins imprint prov-2026-29e9b1b1 (and blueprint
// prov-2026-f5fb0c02): line templates, a hyphen for an empty id, sort by id with
// hyphen lines first, change before conflict before problem, then line number,
// unchanged records counted but not printed, the summary line.
func TestPrintDocumentPlan(t *testing.T) {
	const a, b, c = "prov-2026-0000a001", "prov-2026-0000b001", "prov-2026-0000c001"
	plan := &document.Plan{
		Changes: []document.Change{
			{ID: c, IntentChanged: true},
			{ID: a, ConstraintsChanged: true},
			{ID: b, IntentChanged: true, ConstraintsChanged: true},
		},
		Conflicts: []document.Problem{
			{Kind: "conflict", ID: b, Line: 8, AnchoredVersion: 2, CurrentVersion: 3, Message: "ignored"},
			{Kind: "conflict", ID: b, Line: 4, AnchoredVersion: 1, CurrentVersion: 3, Message: "ignored"},
		},
		Problems: []document.Problem{
			{Kind: "heading", ID: b, Line: 9, Message: "heading text differs from the record title"},
			{Kind: "unanchored", ID: "", Line: 12, Message: "heading has no anchor"},
			{Kind: "unanchored", ID: "", Line: 3, Message: "heading has no anchor"},
			{Kind: "unknown", ID: a, Line: 7, Message: "no record with this id"},
			{Kind: "level", ID: b, Line: 2, Message: "heading level does not match the record type"},
		},
		Unchanged: []string{"prov-2026-0000d001", "prov-2026-0000d002"},
	}
	var buf bytes.Buffer
	printDocumentPlan(&buf, plan)
	want := "- unanchored line 3: heading has no anchor\n" +
		"- unanchored line 12: heading has no anchor\n" +
		a + " change constraints\n" +
		a + " unknown line 7: no record with this id\n" +
		b + " change intent and constraints\n" +
		b + " conflict anchored v1 current v3\n" +
		b + " conflict anchored v2 current v3\n" +
		b + " level line 2: heading level does not match the record type\n" +
		b + " heading line 9: heading text differs from the record title\n" +
		c + " change intent\n" +
		"3 changed, 2 conflicts, 5 problems, 2 unchanged\n"
	if buf.String() != want {
		t.Fatalf("output:\n%s\nwant:\n%s", buf.String(), want)
	}
}

// TestPrintDocumentPlanSmall pins imprint prov-2026-29e9b1b1: an empty plan prints
// only the summary, a nil plan prints nothing, and a lone change prints its line.
func TestPrintDocumentPlanSmall(t *testing.T) {
	var buf bytes.Buffer
	printDocumentPlan(&buf, nil)
	if buf.Len() != 0 {
		t.Fatalf("nil plan printed %q", buf.String())
	}
	printDocumentPlan(&buf, &document.Plan{})
	if buf.String() != "0 changed, 0 conflicts, 0 problems, 0 unchanged\n" {
		t.Fatalf("%q", buf.String())
	}
	buf.Reset()
	printDocumentPlan(&buf, &document.Plan{Changes: []document.Change{{ID: "prov-2026-0000a001", IntentChanged: true}}, Unchanged: []string{"x"}})
	if buf.String() != "prov-2026-0000a001 change intent\n1 changed, 0 conflicts, 0 problems, 1 unchanged\n" {
		t.Fatalf("%q", buf.String())
	}
	buf.Reset()
	printDocumentPlan(&buf, &document.Plan{Problems: []document.Problem{{Kind: "unanchored", Line: 1, Message: "heading has no anchor"}}})
	if buf.String() != "- unanchored line 1: heading has no anchor\n0 changed, 0 conflicts, 1 problems, 0 unchanged\n" {
		t.Fatalf("%q", buf.String())
	}
}

// TestPlanError pins imprint prov-2026-57558206: nil without conflicts and
// problems, otherwise "document plan has N conflicts and M problems".
func TestPlanError(t *testing.T) {
	p := func(n int) []document.Problem { return make([]document.Problem, n) }
	if err := planError(&document.Plan{Changes: []document.Change{{ID: "x"}}, Unchanged: []string{"y"}}); err != nil {
		t.Fatalf("clean plan: %v", err)
	}
	for _, c := range []struct {
		conflicts, problems int
		want                string
	}{
		{1, 0, "document plan has 1 conflicts and 0 problems"},
		{0, 1, "document plan has 0 conflicts and 1 problems"},
		{2, 3, "document plan has 2 conflicts and 3 problems"},
		{10, 12, "document plan has 10 conflicts and 12 problems"},
	} {
		err := planError(&document.Plan{Conflicts: p(c.conflicts), Problems: p(c.problems)})
		if err == nil || err.Error() != c.want {
			t.Errorf("%d/%d: %v, want %q", c.conflicts, c.problems, err, c.want)
		}
	}
	if errors.Is(planError(&document.Plan{Conflicts: p(1)}), document.ErrConflicts) {
		t.Error("planError is its own error, not the library's")
	}
}
