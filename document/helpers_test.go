package document

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/livecodelife/provenance/record"
)

// rec describes one record file written into a test store. The zero value is
// a schema 2 draft blueprint at version 1 with no intent and no constraints.
type rec struct {
	ID             string
	Title          string
	Type           string // "" leaves the type key out (a blueprint)
	Status         string // "" means draft
	Version        int    // 0 means 1; written as is, with Reason when above 1
	Reason         string
	Intent         string
	Constraints    []string
	Scope          []string
	Implements     string
	ImplementsRepo string
	Schema1        bool // write a v1 file (no version key)
}

func quote(s string) string { return strconv.Quote(s) }

func yamlList(items []string) string {
	if len(items) == 0 {
		return "[]"
	}
	var b strings.Builder
	for _, it := range items {
		b.WriteString("\n  - " + quote(it))
	}
	return b.String()
}

// yaml renders the record file. Strings are double-quoted so any fixture text
// (newlines, markers, backslashes) survives the round trip through the codec.
func (r rec) yaml() string {
	status := r.Status
	if status == "" {
		status = "draft"
	}
	var b strings.Builder
	w := func(k, v string) { b.WriteString(k + ": " + v + "\n") }
	w("id", r.ID)
	w("title", quote(r.Title))
	if r.Schema1 {
		w("status", status)
		w("created_at", `"2026-01-01"`)
		w("author", "a@b")
		w("intent", quote(r.Intent))
		w("constraints", yamlList(r.Constraints))
		w("affected_scope", yamlList(r.Scope))
		w("forbidden_scope", "[]")
		w("type", orBlueprint(r.Type))
		w("supersedes", `""`)
		w("superseded_by", `""`)
		if r.Implements != "" {
			w("implements", r.Implements)
		}
		w("related", "[]")
		w("sealed_at_sha", `""`)
		w("associated_specs", "[]")
		w("associated_traces", "[]")
		w("monitors", "[]")
		w("tags", "[]")
		return b.String()
	}
	if r.Type != "" {
		w("type", r.Type)
	}
	w("status", status)
	v := r.Version
	if v == 0 {
		v = 1
	}
	w("version", strconv.Itoa(v))
	w("reason", quote(r.Reason))
	w("created_at", `"2026-10-15"`)
	w("author", "a@b")
	w("intent", quote(r.Intent))
	w("constraints", yamlList(r.Constraints))
	w("affected_scope", yamlList(r.Scope))
	w("forbidden_scope", "[]")
	if r.Implements != "" {
		w("implements", r.Implements)
		if r.ImplementsRepo != "" {
			w("implements_repo", r.ImplementsRepo)
		}
	}
	w("related", "[]")
	w("sealed_at_sha", `""`)
	w("associated_specs", "[]")
	w("associated_traces", "[]")
	w("monitors", "[]")
	w("tags", "[]")
	return b.String()
}

func orBlueprint(t string) string {
	if t == "" {
		return "blueprint"
	}
	return t
}

// newStore is an empty files store under a fresh temp dir. The library under
// test must not need git, so the directory is not a repository.
func newStore(t *testing.T) *record.FilesStore {
	t.Helper()
	root := t.TempDir()
	if err := os.MkdirAll(filepath.Join(root, "provenance"), 0o755); err != nil {
		t.Fatal(err)
	}
	return &record.FilesStore{Root: root}
}

// put writes r as provenance/<id>.yml.
func put(t *testing.T, s *record.FilesStore, r rec) {
	t.Helper()
	p := filepath.Join(s.Root, "provenance", r.ID+".yml")
	if err := os.WriteFile(p, []byte(r.yaml()), 0o644); err != nil {
		t.Fatal(err)
	}
}

// putAll writes every record.
func putAll(t *testing.T, s *record.FilesStore, rs ...rec) {
	t.Helper()
	for _, r := range rs {
		put(t, s, r)
	}
}

// putVersioned writes r as an open record and commits versions 1..upTo through
// record.Commit (versions above 1 are reason-only), so the store holds real
// snapshots and the record file ends at version upTo.
func putVersioned(t *testing.T, s *record.FilesStore, r rec, upTo int) {
	t.Helper()
	r.Status, r.Version, r.Reason = "open", 1, ""
	put(t, s, r)
	for v := 1; v <= upTo; v++ {
		rc, doc, err := s.Load(r.ID)
		if err != nil {
			t.Fatal(err)
		}
		b := record.Bump{Author: "fixture@example.com"}
		if v > 1 {
			b.Reason = fmt.Sprintf("fixture reason %d", v)
		}
		if _, err := record.Commit(s, rc, doc, b); err != nil {
			t.Fatalf("commit %s v%d: %v", r.ID, v, err)
		}
	}
}

// loadRec loads one record.
func loadRec(t *testing.T, s *record.FilesStore, id string) *record.Record {
	t.Helper()
	r, _, err := s.Load(id)
	if err != nil {
		t.Fatal(err)
	}
	return r
}

// versionsOf lists a record's snapshot versions.
func versionsOf(t *testing.T, s *record.FilesStore, id string) []int {
	t.Helper()
	vs, err := s.Versions(id)
	if err != nil {
		t.Fatal(err)
	}
	return vs
}

// tree reads every file under root into a map keyed by relative path, with
// directories keyed as "dir/" so created directories count as changes.
func tree(t *testing.T, root string) map[string]string {
	t.Helper()
	out := map[string]string{}
	err := filepath.WalkDir(root, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, _ := filepath.Rel(root, p)
		if rel == "." {
			return nil
		}
		if d.IsDir() {
			out[rel+"/"] = ""
			return nil
		}
		data, err := os.ReadFile(p)
		if err != nil {
			return err
		}
		out[rel] = string(data)
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	return out
}

// sameTree fails the test when two tree snapshots differ.
func sameTree(t *testing.T, before, after map[string]string) {
	t.Helper()
	for k, v := range before {
		if got, ok := after[k]; !ok {
			t.Errorf("%s was removed", k)
		} else if got != v {
			t.Errorf("%s was modified", k)
		}
	}
	for k := range after {
		if _, ok := before[k]; !ok {
			t.Errorf("%s was created", k)
		}
	}
}

// faultStore wraps a store, injects errors and records every write call.
type faultStore struct {
	record.Store
	listErr     error
	loadErr     map[string]error
	versionsErr map[string]error
	saveErr     map[string]error
	writes      []string
}

func (f *faultStore) List() ([]string, error) {
	if f.listErr != nil {
		return nil, f.listErr
	}
	return f.Store.List()
}

func (f *faultStore) Load(id string) (*record.Record, *record.Document, error) {
	if err := f.loadErr[id]; err != nil {
		return nil, nil, err
	}
	return f.Store.Load(id)
}

func (f *faultStore) Versions(id string) ([]int, error) {
	if err := f.versionsErr[id]; err != nil {
		return nil, err
	}
	return f.Store.Versions(id)
}

func (f *faultStore) Save(r *record.Record, d *record.Document) error {
	f.writes = append(f.writes, "save "+r.ID)
	if err := f.saveErr[r.ID]; err != nil {
		return err
	}
	return f.Store.Save(r, d)
}

func (f *faultStore) Append(s *record.Snapshot) error {
	f.writes = append(f.writes, "append "+s.ID)
	return f.Store.Append(s)
}

func (f *faultStore) Remove(id string) error {
	f.writes = append(f.writes, "remove "+id)
	return f.Store.Remove(id)
}

// guard wraps s so any write is recorded.
func guard(s record.Store) *faultStore { return &faultStore{Store: s} }

// headings lists the heading lines of a rendered document, "## " and "### "
// prefixes kept, in order.
func headings(doc string) []string {
	var out []string
	for _, l := range strings.Split(doc, "\n") {
		if strings.HasPrefix(l, "## ") || strings.HasPrefix(l, "### ") {
			out = append(out, l)
		}
	}
	return out
}

// docBuilder assembles a document text and remembers the one-based line each
// piece starts on.
type docBuilder struct{ b strings.Builder }

func (d *docBuilder) line() int { return strings.Count(d.b.String(), "\n") + 1 }

// add appends text and returns the line it starts on.
func (d *docBuilder) add(text string) int {
	n := d.line()
	d.b.WriteString(text)
	return n
}

func (d *docBuilder) String() string { return d.b.String() }

// sec is one well-formed section: heading, anchor, intent and constraints.
func sec(level int, heading, id string, version int, intent string, cons ...string) string {
	var b strings.Builder
	b.WriteString(strings.Repeat("#", level) + " " + heading + "\n")
	b.WriteString(fmt.Sprintf("<!-- provenance:%s@%d -->\n\n", id, version))
	if intent != "" {
		b.WriteString(intent + "\n\n")
	}
	b.WriteString("**Constraints**\n")
	for _, c := range cons {
		b.WriteString("- " + c + "\n")
	}
	b.WriteString("\n")
	return b.String()
}
