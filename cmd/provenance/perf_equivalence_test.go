package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"testing"

	"github.com/livecodelife/provenance/record"
)

// TestPerfEquivalence runs two provenance binaries over the same repository
// and requires the same output and exit code from both: lint and lint --record
// as the working tree stands, then check --staged for a set of staged changes.
// It is how a speed-up is shown to change nothing but the time.
//
//	PROVENANCE_EQUIV_REPO  a repository to run in (required; it is edited and
//	                       restored, and its index is reset: use a throwaway
//	                       worktree, never one with work in it)
//	PROVENANCE_EQUIV_OLD   the reference binary (required)
//	PROVENANCE_EQUIV_NEW   the binary under test (default: this checkout, built)
//	PROVENANCE_EQUIV_SAMPLE how many records of each status get a lint --record
//	                       run (default 6)
func TestPerfEquivalence(t *testing.T) {
	repo, oldBin := os.Getenv("PROVENANCE_EQUIV_REPO"), os.Getenv("PROVENANCE_EQUIV_OLD")
	if repo == "" || oldBin == "" {
		t.Skip("set PROVENANCE_EQUIV_REPO and PROVENANCE_EQUIV_OLD to compare two binaries")
	}
	newBin := os.Getenv("PROVENANCE_EQUIV_NEW")
	if newBin == "" {
		newBin = filepath.Join(t.TempDir(), "provenance-new")
		if out, err := exec.Command("go", "build", "-o", newBin, ".").CombinedOutput(); err != nil {
			t.Fatalf("build: %v\n%s", err, out)
		}
	}
	eq := &equiv{t: t, repo: repo, oldBin: oldBin, newBin: newBin}
	if eq.git("status", "--porcelain", "--untracked-files=no") != "" || eq.git("diff", "--cached", "--name-only") != "" {
		t.Fatal("the repository has changes: run on a clean throwaway worktree")
	}
	eq.pick()

	// Lint as the tree stands.
	t.Run("tree", func(t *testing.T) {
		eq.t = t
		for _, args := range [][]string{
			{"lint"}, {"lint", "--all"}, {"lint", "--warn"}, {"lint", "--info"},
			{"lint", "--format", "json", "--all"}, {"lint", "--format", "sarif"},
			{"lint", "--record", "prov-2026-doesnotexist"},
		} {
			eq.compare("tree: "+strings.Join(args, " "), args...)
		}
		for _, id := range eq.sample {
			eq.compare("tree: lint --record "+id, "lint", "--record", id, "--all")
		}
		eq.compare("tree: lint --record JSON", "lint", "--record", eq.sample[0], "--format", "json")
	})

	// Lint with a record edited in the working tree.
	for _, e := range eq.edits() {
		e := e
		t.Run("edited/"+e.name, func(t *testing.T) {
			eq.t = t
			defer eq.restore()
			e.apply(eq)
			eq.compare("edited "+e.name+": lint --record", "lint", "--record", e.id, "--all")
			eq.compare("edited "+e.name+": lint", "lint", "--all")
		})
	}

	// check --staged.
	for _, sc := range eq.scenarios() {
		sc := sc
		t.Run("staged/"+sc.name, func(t *testing.T) {
			eq.t = t
			defer eq.restore()
			sc.apply(eq)
			eq.git("add", "-A", "--", ".")
			msg := filepath.Join(t.TempDir(), "msg")
			if err := os.WriteFile(msg, []byte(sc.message+"\n"), 0o644); err != nil {
				t.Fatal(err)
			}
			eq.compare("staged "+sc.name, "check", "--staged", "--message-file", msg)
		})
	}
}

type equiv struct {
	t              *testing.T
	repo           string
	oldBin, newBin string

	open, implemented, draft []*record.Record
	sample                   []string
	inScope                  string   // a file the first open record governs
	others                   []string // tracked files outside any record's exact scope
	touched, created         []string
}

func (e *equiv) git(args ...string) string {
	e.t.Helper()
	cmd := exec.Command("git", args...)
	cmd.Dir = e.repo
	var out, errb bytes.Buffer
	cmd.Stdout, cmd.Stderr = &out, &errb
	if err := cmd.Run(); err != nil {
		e.t.Fatalf("git %v: %v\n%s", args, err, errb.String())
	}
	return strings.TrimSpace(out.String())
}

func (e *equiv) run(bin string, args ...string) (string, int) {
	cmd := exec.Command(bin, args...)
	cmd.Dir = e.repo
	var out bytes.Buffer
	cmd.Stdout, cmd.Stderr = &out, &out
	err := cmd.Run()
	code := 0
	var ee *exec.ExitError
	if errors.As(err, &ee) {
		code = ee.ExitCode()
	} else if err != nil {
		e.t.Fatalf("%s: %v", bin, err)
	}
	return out.String(), code
}

func (e *equiv) compare(name string, args ...string) {
	e.t.Helper()
	oldOut, oldCode := e.run(e.oldBin, args...)
	newOut, newCode := e.run(e.newBin, args...)
	// SARIF names the tool version, which is the build's VCS stamp.
	// check --staged names its throwaway checkout, which has a random name.
	for _, re := range []struct {
		re   *regexp.Regexp
		with string
	}{{toolVersion, `"version": "-"`}, {stagedDir, "provenance-staged-N"}} {
		oldOut, newOut = re.re.ReplaceAllString(oldOut, re.with), re.re.ReplaceAllString(newOut, re.with)
	}
	if oldOut != newOut || oldCode != newCode {
		e.t.Errorf("%s: outputs differ (exit %d vs %d)\n--- old\n%s\n--- new\n%s", name, oldCode, newCode, clip(oldOut), clip(newOut))
		return
	}
	e.t.Logf("%s: identical (exit %d, %d bytes)", name, newCode, len(newOut))
}

var (
	toolVersion = regexp.MustCompile(`"version": "v[^"]*"`)
	stagedDir   = regexp.MustCompile(`provenance-staged-[0-9]+`)
)

func clip(s string) string {
	if len(s) > 3000 {
		return s[:3000] + "\n…"
	}
	return s
}

// pick chooses the records and files the scenarios work on, in a fixed order.
func (e *equiv) pick() {
	store := &record.FilesStore{Root: e.repo}
	ids, err := store.List()
	if err != nil || len(ids) == 0 {
		e.t.Fatalf("no records in %s: %v", e.repo, err)
	}
	n := 6
	if v := os.Getenv("PROVENANCE_EQUIV_SAMPLE"); v != "" {
		fmt.Sscan(v, &n)
	}
	exact := map[string]bool{}
	for _, id := range ids {
		r, _, err := store.Load(id)
		if err != nil || r.Schema != 2 {
			continue
		}
		switch r.Status {
		case record.StatusOpen:
			e.open = append(e.open, r)
		case record.StatusImplemented:
			e.implemented = append(e.implemented, r)
		case record.StatusDraft:
			e.draft = append(e.draft, r)
		}
		for _, p := range r.AffectedScope {
			if !strings.ContainsAny(p, "*?") && !strings.HasPrefix(p, "re:") {
				exact[p] = true
			}
		}
	}
	if len(e.open) == 0 || len(e.implemented) == 0 {
		e.t.Fatal("need at least one open and one implemented v2 record")
	}
	spread := func(rs []*record.Record) {
		for i := 0; i < n && i < len(rs); i++ {
			e.sample = append(e.sample, rs[i*len(rs)/min(n, len(rs))].ID)
		}
	}
	spread(e.open)
	spread(e.implemented)
	spread(e.draft)
	for _, r := range e.open {
		for _, p := range r.AffectedScope {
			if exact[p] && e.fileExists(p) && e.inScope == "" {
				e.inScope = p
			}
		}
	}
	if e.inScope == "" {
		e.t.Fatal("no open record governs an existing file by exact path")
	}
	listed := strings.Split(e.git("ls-files", "-z", "--", ":!provenance", ":!.provenance", ":!.sedum"), "\x00")
	sort.Strings(listed)
	for _, f := range listed {
		if f != "" && !exact[f] && e.isText(f) && len(e.others) < 12 {
			e.others = append(e.others, f)
		}
	}
}

func (e *equiv) fileExists(p string) bool {
	st, err := os.Stat(filepath.Join(e.repo, p))
	return err == nil && st.Mode().IsRegular()
}

func (e *equiv) isText(p string) bool {
	data, err := os.ReadFile(filepath.Join(e.repo, p))
	return err == nil && len(data) < 1<<20 && !bytes.Contains(data, []byte{0})
}

func (e *equiv) edit(p string, f func(string) string) {
	full := filepath.Join(e.repo, p)
	data, err := os.ReadFile(full)
	if err != nil {
		e.t.Fatal(err)
	}
	e.touched = append(e.touched, p)
	if err := os.Chmod(full, 0o644); err != nil {
		e.t.Fatal(err)
	}
	if err := os.WriteFile(full, []byte(f(string(data))), 0o644); err != nil {
		e.t.Fatal(err)
	}
}

// addConstraint changes a record's content without a version bump.
func addConstraint(s string) string {
	return strings.Replace(s, "\nconstraints:\n", "\nconstraints:\n  - an extra constraint\n", 1)
}

func (e *equiv) touch(p string) {
	e.edit(p, func(s string) string { return s + "\n// perf-equivalence\n" })
}

func (e *equiv) create(p, body string) {
	full := filepath.Join(e.repo, p)
	_ = os.MkdirAll(filepath.Dir(full), 0o755)
	if err := os.WriteFile(full, []byte(body), 0o644); err != nil {
		e.t.Fatal(err)
	}
	e.created = append(e.created, p)
}

// restore puts the index and the tree back as they were.
func (e *equiv) restore() {
	e.git("reset", "-q")
	if len(e.touched) > 0 {
		e.git(append([]string{"checkout", "-q", "--"}, e.touched...)...)
	}
	for _, p := range e.created {
		_ = os.Remove(filepath.Join(e.repo, p))
	}
	e.touched, e.created = nil, nil
}

type edit struct {
	name, id string
	apply    func(*equiv)
}

func equivRecordFile(r *record.Record) string { return "provenance/" + r.ID + ".yml" }

func (e *equiv) edits() []edit {
	op, im := e.open[0], e.implemented[0]
	var imprint *record.Record
	for _, r := range e.open {
		if r.Implements != "" {
			imprint = r
			break
		}
	}
	out := []edit{
		{"content-change-no-bump", op.ID, func(e *equiv) {
			e.edit(equivRecordFile(op), addConstraint)
		}},
		{"implemented-reopened", im.ID, func(e *equiv) {
			e.edit(equivRecordFile(im), func(s string) string { return strings.Replace(s, "status: implemented", "status: open", 1) })
		}},
		{"scope-path-missing", op.ID, func(e *equiv) {
			e.edit(equivRecordFile(op), func(s string) string {
				return strings.Replace(s, "affected_scope:\n  - ", "affected_scope:\n  - zz/missing/", 1)
			})
		}},
		{"scope-bad-regex", op.ID, func(e *equiv) {
			e.edit(equivRecordFile(op), func(s string) string {
				return strings.Replace(s, "affected_scope:\n  - ", "affected_scope:\n  - re:(\n  - ", 1)
			})
		}},
		{"spec-missing", op.ID, func(e *equiv) {
			e.edit(equivRecordFile(op), func(s string) string {
				return strings.Replace(s, "associated_specs:", "associated_specs:\n  - path: zz/absent_spec.sh", 1)
			})
		}},
		{"unreadable-record", e.implemented[len(e.implemented)-1].ID, func(e *equiv) {
			e.edit(equivRecordFile(e.implemented[len(e.implemented)-1]), func(string) string { return "id: [unclosed\n" })
		}},
		{"overlap-with-another-open", op.ID, func(e *equiv) {
			e.edit(equivRecordFile(op), func(s string) string {
				return strings.Replace(s, "affected_scope:\n  - ", "affected_scope:\n  - src/**\n  - ", 1)
			})
		}},
	}
	if imprint != nil {
		out = append(out,
			edit{"implements-unknown", imprint.ID, func(e *equiv) {
				e.edit(equivRecordFile(imprint), func(s string) string {
					return strings.Replace(s, "implements: "+imprint.Implements, "implements: prov-2026-deadbeef", 1)
				})
			}},
			edit{"implements-pin-ahead", imprint.ID, func(e *equiv) {
				e.edit(equivRecordFile(imprint), func(s string) string {
					return strings.Replace(s, fmt.Sprintf("implements_version: %d", imprint.ImplementsVersion), "implements_version: 99", 1)
				})
			}})
	}
	return out
}

type scenario struct {
	name, message string
	apply         func(*equiv)
}

func (e *equiv) scenarios() []scenario {
	op, im := e.open[0], e.implemented[0]
	tagOp := fmt.Sprintf("work [%s@%d]", op.ID, op.Version)
	tagIm := fmt.Sprintf("work [%s@%d]", im.ID, im.Version)
	ten := func(e *equiv) {
		e.touch(e.inScope)
		for _, f := range e.others[:min(9, len(e.others))] {
			e.touch(f)
		}
	}
	var out []scenario
	add := func(name, msg string, f func(*equiv)) { out = append(out, scenario{name, msg, f}) }

	add("one-file-in-scope", tagOp, func(e *equiv) { e.touch(e.inScope) })
	add("one-file-in-scope-untagged-version", "work ["+op.ID+"]", func(e *equiv) { e.touch(e.inScope) })
	add("one-file-out-of-scope", tagOp, func(e *equiv) { e.touch(e.others[0]) })
	add("one-file-no-tag", "no tag", func(e *equiv) { e.touch(e.inScope) })
	add("one-file-unknown-tag", "work [prov-2026-ffffffff@1]", func(e *equiv) { e.touch(e.inScope) })
	add("one-file-wrong-version-tag", fmt.Sprintf("work [%s@%d]", op.ID, op.Version+5), func(e *equiv) { e.touch(e.inScope) })
	add("ten-files-open-record", tagOp, ten)
	add("ten-files-implemented-record", tagIm, ten)
	add("ten-files-no-tag", "none", ten)
	add("delete-in-scope-file", tagOp, func(e *equiv) {
		e.touched = append(e.touched, e.inScope)
		_ = os.Remove(filepath.Join(e.repo, e.inScope))
	})
	add("new-file-in-no-scope", tagOp, func(e *equiv) { e.create("perf-equivalence-new.txt", "new\n") })
	add("excluded-only", "docs only", func(e *equiv) {
		for _, f := range []string{"README.md", ".gitignore"} {
			if e.fileExists(f) {
				e.touch(f)
				return
			}
		}
		e.t.Skip("no excluded file to touch")
	})
	for _, ed := range e.edits() {
		ed := ed
		rec := ed.id
		var ver int
		for _, r := range append(append([]*record.Record{}, e.open...), e.implemented...) {
			if r.ID == rec {
				ver = r.Version
			}
		}
		add("record-"+ed.name+"-tagged", fmt.Sprintf("edit [%s@%d]", rec, ver), ed.apply)
		add("record-"+ed.name+"-untagged", "edit", ed.apply)
		add("record-"+ed.name+"-plus-code-file", fmt.Sprintf("edit [%s@%d]", rec, ver), func(e *equiv) {
			ed.apply(e)
			e.touch(e.inScope)
		})
	}
	add("two-records-edited", fmt.Sprintf("edit [%s@%d] [%s@%d]", op.ID, op.Version, im.ID, im.Version), func(e *equiv) {
		e.edit(equivRecordFile(op), addConstraint)
		e.edit(equivRecordFile(im), addConstraint)
	})
	add("snapshot-rewritten", tagOp, func(e *equiv) {
		files, _ := filepath.Glob(filepath.Join(e.repo, "provenance", ".versions", op.ID, "*.yml"))
		if len(files) == 0 {
			e.t.Skip("open record has no snapshot")
		}
		rel, _ := filepath.Rel(e.repo, files[0])
		e.touch(rel)
	})
	add("manifest-edited", "state", func(e *equiv) {
		if !e.fileExists(".provenance/versions.jsonl") {
			e.t.Skip("no manifest")
		}
		e.edit(".provenance/versions.jsonl", func(s string) string {
			return s + `{"id":"prov-2026-ffffffff","version":1,"content_hash":"sha256:00"}` + "\n"
		})
	})
	return out
}
