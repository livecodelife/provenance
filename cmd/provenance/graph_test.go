package main

import (
	"encoding/json"
	"strings"
	"testing"
)

func graphJSON(t *testing.T, cfg string, args ...string) graphDoc {
	t.Helper()
	out := mustRun(t, append([]string{"-c", cfg, "graph", "--format", "json"}, args...)...)
	var doc graphDoc
	if err := json.Unmarshal([]byte(out), &doc); err != nil {
		t.Fatalf("json: %v\n%s", err, out)
	}
	return doc
}

func TestGraphHumanTree(t *testing.T) {
	_, cfg := staleRepo(t)
	out := mustRun(t, "-c", cfg, "graph")
	parent := "  " + staleParent + "@2  open  blueprint  Draft " + staleParent
	child := "    " + staleChild + "@1  open  imprint  Draft " + staleChild + "  -> " + staleParent + "@1  stale (direct, parent at 2)"
	if !strings.Contains(out, "PROVENANCE GRAPH") || !strings.Contains(out, parent) || !strings.Contains(out, child) {
		t.Fatalf("tree:\n%s", out)
	}
	if strings.Index(out, parent) > strings.Index(out, child) {
		t.Fatalf("a child is listed under its parent:\n%s", out)
	}
	mustRun(t, "-c", cfg, "repin", staleChild, "--reason", "unaffected")
	if out := mustRun(t, "-c", cfg, "graph"); strings.Contains(out, "stale") || !strings.Contains(out, "-> "+staleParent+"@2") {
		t.Fatalf("after repin:\n%s", out)
	}
}

func TestGraphJSONAndSelection(t *testing.T) {
	_, cfg := staleRepo(t)
	doc := graphJSON(t, cfg)
	if len(doc.Nodes) != 2 || doc.Nodes[0].ID != staleParent || doc.Nodes[1].ID != staleChild {
		t.Fatalf("nodes %+v", doc.Nodes)
	}
	if v := doc.Nodes[0].Versions; len(v) != 2 || v[1].Version != 2 || v[1].Reason != "parent moved" {
		t.Fatalf("versions %+v", v)
	}
	if len(doc.Edges) != 1 {
		t.Fatalf("edges %+v", doc.Edges)
	}
	if e := doc.Edges[0]; e.From != staleChild || e.To != staleParent || e.Edge != "implements" || e.Pinned != 1 || e.ParentVersion != 2 || e.Stale != "direct" {
		t.Fatalf("edge %+v", e)
	}
	if again := mustRun(t, "-c", cfg, "graph", "--format", "json"); again != mustRun(t, "-c", cfg, "graph", "--format", "json") {
		t.Fatal("not deterministic")
	}
	if d := graphJSON(t, cfg, "--root", staleChild); len(d.Nodes) != 2 {
		t.Fatalf("--root on the child shows its parent too: %+v", d.Nodes)
	}
	if d := graphJSON(t, cfg, "--root", staleParent); len(d.Nodes) != 2 || len(d.Edges) != 1 {
		t.Fatalf("--root on the parent shows its descendants: %+v", d)
	}
	if d := graphJSON(t, cfg, "--filter", "draft"); len(d.Nodes) != 0 || len(d.Edges) != 0 {
		t.Fatalf("--filter draft: %+v", d)
	}
	if _, err := run(t, "-c", cfg, "graph", "--root", "prov-2026-ffffffff"); err == nil {
		t.Error("an unknown --root must fail")
	}
	if _, err := run(t, "-c", cfg, "graph", "--filter", "bogus"); err == nil {
		t.Error("an unknown --filter must fail")
	}
	if _, err := run(t, "-c", cfg, "graph", "--format", "xml"); err == nil {
		t.Error("an unknown --format must fail")
	}
}

func TestGraphDot(t *testing.T) {
	_, cfg := staleRepo(t)
	out := mustRun(t, "-c", cfg, "graph", "--format", "dot")
	for _, want := range []string{"digraph provenance {", "color=orange", `"` + staleChild + `" -> "` + staleParent + `"`, "style=dashed, color=red"} {
		if !strings.Contains(out, want) {
			t.Errorf("dot misses %q:\n%s", want, out)
		}
	}
}

func TestGraphSharedParent(t *testing.T) {
	remote := sharedRemote(t)
	_, cfg := sharedConsumer(t, remote)
	mustRun(t, "-c", cfg, "sync")
	doc := graphJSON(t, cfg)
	var shared bool
	for _, n := range doc.Nodes {
		shared = shared || (n.ID == sharedParent && n.Shared)
	}
	if !shared || len(doc.Edges) != 1 || doc.Edges[0].To != sharedParent || doc.Edges[0].ParentVersion == 0 {
		t.Fatalf("shared parent: %+v", doc)
	}
}
