package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"testing"

	"github.com/livecodelife/provenance/record"
)

// Ids of the records in the document command fixtures. They sort in this
// order, which is also the order of the rendered sections of the brief and
// its blueprints.
const (
	docBrief = "prov-2026-0000d0b1"
	docBP1   = "prov-2026-0000d0c1"
	docBP2   = "prov-2026-0000d0c2"
	docV1    = "prov-2026-0000d0c3" // a schema 1 draft blueprint
	docLast  = "prov-2026-0000d0c4"
)

// docRecord is a schema 2 draft record with scope a.go, version 1 and the
// given content, as v2Draft writes one.
func docRecord(id, typ, title, intent, implements string, cons ...string) string {
	var b strings.Builder
	b.WriteString("id: " + id + "\ntitle: " + strconv.Quote(title) + "\nstatus: draft\nversion: 1\nreason: \"\"\ncreated_at: \"2026-10-15\"\nauthor: a@b\n")
	b.WriteString("intent: " + strconv.Quote(intent) + "\nconstraints:")
	if len(cons) == 0 {
		b.WriteString(" []")
	}
	for _, c := range cons {
		b.WriteString("\n  - " + strconv.Quote(c))
	}
	b.WriteString("\naffected_scope:\n  - a.go\nforbidden_scope: []\ntype: " + typ + "\n")
	if implements != "" {
		b.WriteString("implements: " + implements + "\n")
	}
	b.WriteString("related: []\nsealed_at_sha: \"\"\nassociated_specs: []\nassociated_traces: []\nmonitors: []\ntags: []\n")
	return b.String()
}

// docRepo is a migrated, ungit'd repository holding a brief and two blueprints
// (all drafts) plus the extra records; the ids in open are then opened, which
// writes their version 1 snapshot and the manifest.
func docRepo(t *testing.T, extra map[string]string, open ...string) (root, cfg string) {
	t.Helper()
	recs := map[string]string{
		docBrief: docRecord(docBrief, "brief", "Brief", "Brief intent", "", "B constraint"),
		docBP1:   docRecord(docBP1, "blueprint", "Blueprint one", "Blueprint one intent", docBrief, "P1 constraint"),
		docBP2:   docRecord(docBP2, "blueprint", "Blueprint two", "Blueprint two intent", docBrief, "P2 constraint"),
	}
	for k, v := range extra {
		recs[k] = v
	}
	root, cfg = v2Repo(t, false, recs)
	for _, id := range open {
		mustRun(t, "-c", cfg, "open", "--record", id)
	}
	return root, cfg
}

// bumpRecord gives an opened record a reason-only next version through the
// library, so documents rendered before it carry an out-of-date anchor.
func bumpRecord(t *testing.T, root, id string) {
	t.Helper()
	s := &record.FilesStore{Root: root}
	rec, doc, err := s.Load(id)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := record.Commit(s, rec, doc, record.Bump{Reason: "moved on", Author: "x@y"}); err != nil {
		t.Fatal(err)
	}
}

// editDoc replaces the first occurrence of old, which must exist.
func editDoc(t *testing.T, doc, old, repl string) string {
	t.Helper()
	if !strings.Contains(doc, old) {
		t.Fatalf("%q not in document:\n%s", old, doc)
	}
	return strings.Replace(doc, old, repl, 1)
}

// docFile writes text into a fresh temp dir outside any repository and returns
// its path.
func docFile(t *testing.T, name, text string) string {
	t.Helper()
	p := filepath.Join(t.TempDir(), name)
	writeFile(t, p, text)
	return p
}

// treeOf reads every file under root (ignoring .git) into a map.
func treeOf(t *testing.T, root string) map[string]string {
	t.Helper()
	out := map[string]string{}
	err := filepath.WalkDir(root, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, _ := filepath.Rel(root, p)
		if d.IsDir() {
			if rel == ".git" {
				return filepath.SkipDir
			}
			if rel != "." {
				out[rel+"/"] = ""
			}
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

func sameTreeOf(t *testing.T, before, after map[string]string) {
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

// withUmask sets the process umask for one test.
func withUmask(t *testing.T, mask int) {
	t.Helper()
	old := syscall.Umask(mask)
	t.Cleanup(func() { syscall.Umask(old) })
}

// missingConfig is a -c value that cannot be opened: a command that reaches
// openRepo with it fails, so an error that is not about the config proves it
// was raised before the repository was opened.
func missingConfig(t *testing.T) string {
	return filepath.Join(t.TempDir(), "no", "such", ".provenance.yml")
}
