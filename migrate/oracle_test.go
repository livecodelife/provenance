package migrate

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strconv"
	"testing"
)

// TestOracle compares Plan with migrate-dry's dry_run.py --json output, the
// done-check of prov-2026-e106ee62. It runs only when PROVENANCE_REAL_REPOS
// names the workspace and PROVENANCE_MIGRATE_ORACLE names that JSON file.
// It only reads the repositories.
func TestOracle(t *testing.T) {
	ws, oraclePath := os.Getenv("PROVENANCE_REAL_REPOS"), os.Getenv("PROVENANCE_MIGRATE_ORACLE")
	if ws == "" || oraclePath == "" {
		t.Skip("set PROVENANCE_REAL_REPOS and PROVENANCE_MIGRATE_ORACLE to run")
	}
	data, err := os.ReadFile(oraclePath)
	if err != nil {
		t.Fatal(err)
	}
	var oracle struct {
		Repos []map[string]json.RawMessage `json:"repos"`
	}
	if err := json.Unmarshal(data, &oracle); err != nil {
		t.Fatal(err)
	}
	totalAfter, totalWith := 0, 0
	for _, o := range oracle.Repos {
		var name string
		_ = json.Unmarshal(o["name"], &name)
		t.Run(name, func(t *testing.T) {
			repo, err := Load(filepath.Join(ws, name), LoadOptions{Workspace: ws})
			if err != nil {
				t.Fatal(err)
			}
			p, err := MakePlan(repo)
			if err != nil {
				t.Fatal(err)
			}
			got := oracleView(p)
			var rules, issues map[string]json.RawMessage
			_ = json.Unmarshal(o["rules"], &rules)
			_ = json.Unmarshal(o["issues"], &issues)
			want := map[string]json.RawMessage{
				"plan":                           o["plan"],
				"survivor_changes":               o["survivor_changes"],
				"records_after":                  o["records_after"],
				"records_after_with_suggestions": o["records_after_with_suggestions"],
				"suggested_keep_separate":        o["suggested_keep_separate"],
				"versioned_records.order":        versionOrders(t, rules["versioned_records"]),
				"bug_kept_separate_deprecated":   orEmpty(rules["bug_kept_separate_deprecated"]),
				"status_goes_backwards":          orEmpty(issues["status_goes_backwards"]),
				"unfinished_version_buried":      orEmpty(issues["unfinished_version_buried"]),
				"implemented_only_through_fixes": orEmpty(issues["implemented_only_through_fixes"]),
			}
			for _, key := range sortedKeys(want) {
				if !jsonEqual(t, got[key], want[key], key != "versioned_records.order") {
					g, _ := json.Marshal(got[key])
					t.Errorf("%s differs\n got %s\nwant %s", key, truncate(g), truncate(want[key]))
				}
			}
			n, _ := strconv.Atoi(string(o["records_after"]))
			m, _ := strconv.Atoi(string(o["records_after_with_suggestions"]))
			totalAfter += n
			totalWith += m
		})
	}
	t.Logf("oracle totals: %d records after, %d with suggestions", totalAfter, totalWith)
}

func orEmpty(raw json.RawMessage) json.RawMessage {
	if raw == nil {
		return json.RawMessage("[]")
	}
	return raw
}

func versionOrders(t *testing.T, raw json.RawMessage) json.RawMessage {
	var vr []struct {
		Survivor string   `json:"survivor"`
		Order    []string `json:"order"`
	}
	_ = json.Unmarshal(raw, &vr)
	m := map[string][]string{}
	for _, v := range vr {
		m[v.Survivor] = v.Order
	}
	out, _ := json.Marshal(m)
	return out
}

// jsonEqual compares as JSON values, with top-level lists compared as
// multisets (the oracle sorts them by its own key).
func jsonEqual(t *testing.T, got any, want json.RawMessage, unordered bool) bool {
	g, err := json.Marshal(got)
	if err != nil {
		t.Fatal(err)
	}
	var a, b any
	_ = json.Unmarshal(g, &a)
	_ = json.Unmarshal(want, &b)
	if !unordered {
		return reflect.DeepEqual(a, b)
	}
	return reflect.DeepEqual(canonList(a), canonList(b))
}

func canonList(v any) any {
	switch v := v.(type) {
	case []any:
		keys := make([]string, len(v))
		for i, x := range v {
			b, _ := json.Marshal(x)
			keys[i] = string(b)
		}
		sort.Strings(keys)
		return keys
	case map[string]any:
		out := map[string]any{}
		for k, x := range v {
			out[k] = canonList(x)
		}
		return out
	}
	return v
}

func truncate(b []byte) string {
	if len(b) > 1500 {
		return string(b[:1500]) + "…"
	}
	return string(b)
}

// oracleView renders a Plan in the shape of the dry run's JSON fields.
func oracleView(p *Plan) map[string]any {
	ruleJSON := func(r string) any {
		if n, err := strconv.Atoi(r); err == nil {
			return n
		}
		return r
	}
	imprint := func(k KeptImprint) map[string]any {
		m := map[string]any{"implements": nilIfEmpty(k.Implements), "related_add": k.RelatedAdd}
		if k.HasNote {
			m["note"] = nilIfEmpty(k.Note)
		}
		return m
	}
	plan := map[string]any{}
	for id, entries := range p.Entries {
		var list []any
		for _, e := range entries {
			m := map[string]any{"target": e.Target, "rule": ruleJSON(e.Rule)}
			if e.Suggest != "" {
				m["suggest_keep_separate"] = e.Suggest
			}
			list = append(list, m)
		}
		plan[id] = list
	}
	changes := map[string]any{}
	for id, s := range p.Survivor {
		if s != id {
			changes[id] = nilIfEmpty(s)
		}
	}
	var sugg []any
	for _, s := range p.Suggestions {
		m := map[string]any{"id": s.ID, "target": s.Target, "rule": ruleJSON(s.Rule), "why": s.Why}
		if s.AsImprint != nil {
			m["as_imprint"] = imprint(*s.AsImprint)
		}
		sugg = append(sugg, m)
	}
	var kept []any
	for _, k := range p.KeptDeprecated {
		kept = append(kept, map[string]any{"id": k.ID, "extends": k.Extends, "as_imprint": imprint(k.AsImprint)})
	}
	warnings := func(kind string) []any {
		out := []any{}
		for _, w := range p.Warnings {
			if w.Kind == kind {
				out = append(out, w.Fields)
			}
		}
		return out
	}
	return map[string]any{
		"plan":                           plan,
		"survivor_changes":               changes,
		"records_after":                  p.RecordsAfter(Options{}),
		"records_after_with_suggestions": p.RecordsAfter(Options{AcceptSuggestions: true}),
		"suggested_keep_separate":        orNil(sugg),
		"versioned_records.order":        p.Versions,
		"bug_kept_separate_deprecated":   orNil(kept),
		"status_goes_backwards":          warnings("status_goes_backwards"),
		"unfinished_version_buried":      warnings("unfinished_version_buried"),
		"implemented_only_through_fixes": warnings("implemented_only_through_fixes"),
	}
}

func nilIfEmpty(s string) any {
	if s == "" {
		return nil
	}
	return s
}

func orNil(l []any) []any {
	if l == nil {
		return []any{}
	}
	return l
}
