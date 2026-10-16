package main

import (
	"encoding/json"
	"path/filepath"
	"strings"
	"testing"
)

func contextDoc(t *testing.T, cfg string, args ...string) ctxDoc {
	t.Helper()
	out := mustRun(t, append([]string{"-c", cfg, "context", "--format", "json"}, args...)...)
	var doc ctxDoc
	if err := json.Unmarshal([]byte(out), &doc); err != nil {
		t.Fatalf("json: %v\n%s", err, out)
	}
	return doc
}

func TestContextPath(t *testing.T) {
	_, cfg := checkRepo(t)
	doc := contextDoc(t, cfg, "a.go")
	if len(doc.Targets) != 1 || doc.Targets[0].Kind != "path" || len(doc.Targets[0].Governing) != 1 {
		t.Fatalf("doc %+v", doc)
	}
	g := doc.Targets[0].Governing[0]
	if g.ID != chkID || g.Status != "open" || g.Version != 1 || len(g.Specs) != 1 || g.Specs[0] != "spec.sh" {
		t.Fatalf("governing %+v", g)
	}
	// `next` is exactly the first action of provenance next for the path.
	var nx struct {
		Actions []nextAction `json:"actions"`
	}
	if err := json.Unmarshal([]byte(mustRun(t, "-c", cfg, "next", "a.go", "--json")), &nx); err != nil {
		t.Fatal(err)
	}
	if got := doc.Targets[0].Next; got == nil || got.Kind != nx.Actions[0].Kind || got.Command != nx.Actions[0].Command || got.Reason != nx.Actions[0].Reason {
		t.Fatalf("context next %+v, next --json %+v", got, nx.Actions[0])
	}
	for _, exempt := range []string{"README.md", ".provenance/versions.jsonl", "provenance/.versions/" + chkID + "/pins.jsonl"} {
		if d := contextDoc(t, cfg, exempt); !d.Targets[0].Exempt || len(d.Targets[0].Governing) != 0 {
			t.Errorf("%s must be exempt: %+v", exempt, d.Targets[0])
		}
	}
	if d := contextDoc(t, cfg, "other.txt"); len(d.Targets[0].Governing) != 0 || d.Targets[0].Next == nil || d.Targets[0].Next.Kind != "create" {
		t.Errorf("an ungoverned path: %+v", d.Targets[0])
	}
	if _, err := run(t, "-c", cfg, "context"); err == nil {
		t.Error("no target must be an error")
	}
	if _, err := run(t, "-c", cfg, "context", "a.go", "--format", "xml"); err == nil {
		t.Error("an unknown format must be an error")
	}
}

func TestContextAncestryConflictsAndRecords(t *testing.T) {
	_, cfg := staleRepo(t)
	d := contextDoc(t, cfg, "a.go").Targets[0]
	if len(d.Conflicts) != 2 || d.Conflicts[0] != staleParent {
		t.Fatalf("two open records cover a.go: %+v", d.Conflicts)
	}
	if anc := d.Ancestry[staleChild]; len(anc) != 1 || anc[0].ID != staleParent || anc[0].Version != 2 {
		t.Fatalf("ancestry %+v", d.Ancestry)
	}
	var child ctxRecord
	for _, g := range d.Governing {
		if g.ID == staleChild {
			child = g
		}
	}
	if child.Stale != "direct" || child.Implements != staleParent || child.ImplementsVersion != 1 {
		t.Fatalf("child %+v", child)
	}

	rec := contextDoc(t, cfg, staleParent, staleChild).Targets
	p, c := rec[0].Record, rec[1].Record
	if rec[0].Kind != "record" || len(p.Versions) != 2 || p.Versions[1].Reason != "parent moved" || len(p.Children) != 1 || p.Children[0].ID != staleChild || p.Children[0].Pinned != 1 || p.Children[0].Stale != "direct" {
		t.Fatalf("parent %+v", p)
	}
	if len(c.Parents) != 1 || c.Parents[0].ID != staleParent || c.Stale != "direct" || !strings.Contains(c.Intent, "Do the thing") {
		t.Fatalf("child record %+v", c)
	}
}

func TestContextCommitsCompactAndHuman(t *testing.T) {
	root, cfg := checkRepo(t)
	stage(t, root, map[string]string{"a.go": "package a // one\n"})
	gitIn(t, root, "commit", "-q", "-m", "Work on a ["+chkID+"]")
	rec := contextDoc(t, cfg, chkID).Targets[0].Record
	if len(rec.Commits) == 0 || !strings.Contains(rec.Commits[0], "Work on a") {
		t.Fatalf("commits %+v", rec.Commits)
	}

	long := strings.Repeat("T", 300)
	path := filepath.Join(root, "provenance", chkID+".yml")
	writeFile(t, path, strings.Replace(readFile(t, path), "title: Draft "+chkID, "title: "+long, 1))
	out := mustRun(t, "-c", cfg, "context", "a.go", "--format", "compact")
	if !strings.Contains(out, "a.go: "+chkID+"@1  open") || strings.Contains(out, strings.Repeat("T", 90)) || !strings.Contains(out, "...") {
		t.Fatalf("compact output is one capped line:\n%s", out)
	}
	writeFile(t, path, strings.Replace(readFile(t, path), "title: "+long, "title: \"hi \\e[31mred\"", 1))
	human := mustRun(t, "-c", cfg, "context", "a.go")
	if !strings.Contains(human, "CONTEXT: a.go") || !strings.Contains(human, "Next:") {
		t.Fatalf("human:\n%s", human)
	}
	if !strings.Contains(human, "hi [31mred") || strings.Contains(human, "\x1b") {
		t.Fatal("control characters in record text must not reach the output")
	}
}

func TestContextFilesFlag(t *testing.T) {
	_, cfg := checkRepo(t)
	doc := contextDoc(t, cfg, "--files", "a.go,README.md")
	if len(doc.Targets) != 2 || doc.Targets[0].Target != "a.go" || len(doc.Targets[0].Governing) != 1 || !doc.Targets[1].Exempt {
		t.Fatalf("--files with a comma list: %+v", doc.Targets)
	}
}
