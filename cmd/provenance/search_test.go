package main

import (
	"encoding/json"
	"strings"
	"testing"
)

const (
	srchA = "prov-2026-0000a001"
	srchB = "prov-2026-0000a002"
	srchC = "prov-2026-0000a003"
)

func srchRecord(id, status, title, intent string) string {
	rec := strings.Replace(v2Draft(id, status), "title: Draft "+id, "title: "+title, 1)
	return strings.Replace(rec, "intent: Do the thing", "intent: "+intent, 1)
}

func searchRepo(t *testing.T) string {
	t.Helper()
	_, cfg := v2Repo(t, false, map[string]string{
		srchA: srchRecord(srchA, "draft", "Reconcile file write permissions", "Lock tracked files that no open record covers"),
		srchB: srchRecord(srchB, "draft", "Graph output", "Render records and parents; permissions are not involved here"),
		srchC: srchRecord(srchC, "deprecated", "Old reconcile design", "Superseded reconcile notes"),
	})
	return cfg
}

func searchJSON(t *testing.T, cfg string, args ...string) searchResult {
	t.Helper()
	out := mustRun(t, append([]string{"-c", cfg, "search", "--format", "json"}, args...)...)
	var res searchResult
	if err := json.Unmarshal([]byte(out), &res); err != nil {
		t.Fatalf("json: %v\n%s", err, out)
	}
	return res
}

func TestSearchRanksAndFilters(t *testing.T) {
	cfg := searchRepo(t)
	res := searchJSON(t, cfg, "reconcile permissions")
	if res.Backend != "lexical" || res.Query != "reconcile permissions" || len(res.Results) != 2 {
		t.Fatalf("result %+v", res)
	}
	if res.Results[0].ID != srchA || res.Results[1].ID != srchB || res.Results[0].Score <= res.Results[1].Score {
		t.Fatalf("ranking %+v", res.Results)
	}
	if h := res.Results[0]; len(h.Matched) < 1 || h.Matched[0] != "title" || !strings.Contains(h.Excerpt, "Reconcile") {
		t.Fatalf("hit %+v", h)
	}
	for _, h := range res.Results {
		if h.ID == srchC {
			t.Fatal("a deprecated record is listed without --status deprecated")
		}
	}
	if d := searchJSON(t, cfg, "reconcile", "--status", "deprecated"); len(d.Results) != 1 || d.Results[0].ID != srchC {
		t.Fatalf("--status deprecated: %+v", d.Results)
	}
	if d := searchJSON(t, cfg, "reconcile", "--type", "imprint"); len(d.Results) != 0 {
		t.Fatalf("--type imprint: %+v", d.Results)
	}
	if one := searchJSON(t, cfg, "permissions", "--limit", "1"); len(one.Results) != 1 {
		t.Fatalf("--limit 1: %+v", one.Results)
	}
	again := mustRun(t, "-c", cfg, "search", "permissions", "--format", "json")
	if again != mustRun(t, "-c", cfg, "search", "permissions", "--format", "json") {
		t.Fatal("not deterministic")
	}
}

func TestSearchIDQueryAndErrors(t *testing.T) {
	cfg := searchRepo(t)
	if res := searchJSON(t, cfg, "0000a002"); len(res.Results) == 0 || res.Results[0].ID != srchB {
		t.Fatalf("an id fragment finds its record first: %+v", res.Results)
	}
	if res := searchJSON(t, cfg, "--query", "graph"); len(res.Results) != 1 || res.Results[0].ID != srchB {
		t.Fatalf("--query: %+v", res.Results)
	}
	for name, args := range map[string][]string{
		"empty query":     {"search"},
		"only stop words": {"search", "the of and"},
		"limit too big":   {"search", "graph", "--limit", "99"},
		"bad status":      {"search", "graph", "--status", "bogus"},
		"bad type":        {"search", "graph", "--type", "bogus"},
		"bad format":      {"search", "graph", "--format", "xml"},
	} {
		if _, err := run(t, append([]string{"-c", cfg}, args...)...); err == nil {
			t.Errorf("%s must be an error", name)
		}
	}
	if out := mustRun(t, "-c", cfg, "search", "zzzzzz"); !strings.Contains(out, "no matching records") || !strings.Contains(out, "(lexical)") {
		t.Fatalf("no match:\n%s", out)
	}
}

func TestSearchTextIsDataAndSharedRecords(t *testing.T) {
	remote := sharedRemote(t)
	root, cfg := sharedConsumer(t, remote)
	mustRun(t, "-c", cfg, "sync")
	res := searchJSON(t, cfg, "thing")
	var shared bool
	for _, h := range res.Results {
		shared = shared || (h.ID == sharedParent && h.Shared)
	}
	if !shared {
		t.Fatalf("a cached shared-repo record is searched and marked shared: %+v", res.Results)
	}

	path := root + "/provenance/" + sharedChild + ".yml"
	writeFile(t, path, strings.Replace(readFile(t, path), "title: Draft "+sharedChild, "title: \"escape \\e[31m"+strings.Repeat("x", 300)+" zebra\"", 1))
	out := mustRun(t, "-c", cfg, "search", "zebra")
	if strings.Contains(out, "\x1b") || strings.Contains(out, strings.Repeat("x", 120)) || !strings.Contains(out, "escape [31m") {
		t.Fatalf("control characters and long text must not reach the human output:\n%s", out)
	}
}
