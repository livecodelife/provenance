package migrate

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/livecodelife/provenance/record"
)

// fakeGit supplies fixed first-add times.
type fakeGit map[string]int64

func (g fakeGit) FirstAdded(root, dir string) (map[string]GitTime, error) {
	out := map[string]GitTime{}
	for name, ts := range g {
		out[name] = GitTime{Unix: ts, SHA: name}
	}
	return out, nil
}

func (fakeGit) FollowFirst(root, relpath string) (GitTime, bool) { return GitTime{}, false }

func writeFile(t *testing.T, path, body string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
}

func v1(id, typ, status, extra string) string {
	return "id: " + id + "\ntitle: T " + id + "\nstatus: " + status + "\ncreated_at: \"2026-01-01\"\nauthor: a@b\n" +
		"intent: Intent of " + id + "\nconstraints:\n  - Shared\naffected_scope:\n  - a.go\nforbidden_scope: []\ntype: " + typ + "\n" +
		"supersedes: \"\"\nsuperseded_by: \"\"\nrelated: []\nsealed_at_sha: \"\"\nassociated_specs: []\nassociated_traces: []\nmonitors: []\ntags: []\n" + extra
}

// fixtureRepo: blueprint bp1 superseded by bp2; bug b1 extends bp2 and adds a
// file and a constraint; imprint im1 implements bp1 (so it re-pins to bp2 at
// version 1 and shows stale); deprecated bug b2 extends bp2 and never folds.
func fixtureRepo(t *testing.T) string {
	root := t.TempDir()
	writeFile(t, filepath.Join(root, ".linespec.yml"), "provenance:\n  dir: provenance\n")
	d := filepath.Join(root, "provenance")
	writeFile(t, filepath.Join(d, "prov-2026-0000b001.yml"), strings.Replace(v1("prov-2026-0000b001", "blueprint", "superseded", ""), `superseded_by: ""`, "superseded_by: prov-2026-0000b002", 1))
	writeFile(t, filepath.Join(d, "prov-2026-0000b002.yml"), strings.Replace(v1("prov-2026-0000b002", "blueprint", "implemented", ""), `supersedes: ""`, "supersedes: prov-2026-0000b001", 1))
	bug := v1("prov-2026-0000f001", "bug", "implemented", "extends: prov-2026-0000b002\n")
	bug = strings.Replace(bug, "  - Shared\n", "  - Shared  \n  - Also validate\n", 1)
	bug = strings.Replace(bug, "  - a.go\n", "  - a.go\n  - b.go\n", 1)
	writeFile(t, filepath.Join(d, "prov-2026-0000f001.yml"), bug)
	writeFile(t, filepath.Join(d, "prov-2026-0000f002.yml"), v1("prov-2026-0000f002", "bug", "deprecated", "extends: prov-2026-0000b002\n"))
	writeFile(t, filepath.Join(d, "prov-2026-0000a001.yml"), v1("prov-2026-0000a001", "imprint", "implemented", "implements: prov-2026-0000b001\n"))
	return root
}

func TestApplyFixture(t *testing.T) {
	root := fixtureRepo(t)
	repo, err := Load(root, LoadOptions{Git: fakeGit{
		"prov-2026-0000b001.yml": 100, "prov-2026-0000b002.yml": 200, "prov-2026-0000f001.yml": 300,
		"prov-2026-0000f002.yml": 400, "prov-2026-0000a001.yml": 150,
	}})
	if err != nil {
		t.Fatal(err)
	}
	p, err := MakePlan(repo)
	if err != nil {
		t.Fatal(err)
	}
	if got := strings.Join(p.Versions["prov-2026-0000b002"], ","); got != "prov-2026-0000b001,prov-2026-0000b002,prov-2026-0000f001" {
		t.Fatalf("version order %s", got)
	}
	store := &record.FilesStore{Root: root}
	res, err := Apply(p, Options{}, store)
	if err != nil {
		t.Fatal(err)
	}
	if res.Records != 3 || res.Snapshots != 5 || len(res.Removed) != 2 {
		t.Fatalf("result %+v", res)
	}

	head, _, err := store.Load("prov-2026-0000b002")
	if err != nil {
		t.Fatal(err)
	}
	if head.Schema != 2 || head.Version != 3 || head.ReasonKind != record.ReasonKindFix || head.Reason != "Intent of prov-2026-0000f001" {
		t.Fatalf("head %+v", head)
	}
	if strings.Join(head.AffectedScope, ",") != "a.go,b.go" || strings.Join(head.Constraints, "|") != "Shared|Also validate" {
		t.Fatalf("bug union: scope %v constraints %v", head.AffectedScope, head.Constraints)
	}
	if strings.Join(head.Aliases, ",") != "prov-2026-0000b001,prov-2026-0000f001" || head.Status != record.StatusImplemented {
		t.Fatalf("aliases %v status %s", head.Aliases, head.Status)
	}
	v2, err := store.Snapshot("prov-2026-0000b002", 2)
	if err != nil {
		t.Fatal(err)
	}
	if v2.Reason != "Supersedes prov-2026-0000b001: T prov-2026-0000b002" || v2.ReasonKind != "" {
		t.Fatalf("rule-2 reason %q kind %q", v2.Reason, v2.ReasonKind)
	}
	if v1, _ := store.Snapshot("prov-2026-0000b002", 1); v1.Reason != "" || v1.Content.Title != "T prov-2026-0000b001" {
		t.Fatalf("v1 %+v", v1)
	}

	im, _, _ := store.Load("prov-2026-0000a001")
	if im.Implements != "prov-2026-0000b002" || im.ImplementsVersion != 1 {
		t.Fatalf("child pin %s@%d, want prov-2026-0000b002@1", im.Implements, im.ImplementsVersion)
	}
	kept, _, _ := store.Load("prov-2026-0000f002")
	if kept.Type != record.TypeImprint || kept.Status != record.StatusDeprecated || kept.Implements != "prov-2026-0000b002" || kept.ImplementsVersion != 2 || kept.ReasonKind != record.ReasonKindFix {
		t.Fatalf("deprecated bug as imprint: %+v", kept)
	}

	checkMigrated(t, root, "provenance", store)
}

// checkMigrated asserts the invariants of a migrated copy: schema 2 records
// that validate, verifying snapshot chains whose head matches the file, and
// an empty re-plan. It returns the number of reason-only versions by record.
func checkMigrated(t *testing.T, root, dir string, store *record.FilesStore) map[string]int {
	t.Helper()
	ids, err := store.List()
	if err != nil {
		t.Fatal(err)
	}
	reasonOnly := map[string]int{}
	for _, id := range ids {
		r, _, err := store.Load(id)
		if err != nil {
			t.Fatal(err)
		}
		if r.Schema != 2 {
			t.Errorf("%s: schema %d", id, r.Schema)
		}
		if err := r.Validate(); err != nil {
			t.Error(err)
		}
		vs, err := store.Versions(id)
		if err != nil || len(vs) != r.Version {
			t.Errorf("%s: versions %v for record version %d (%v)", id, vs, r.Version, err)
			continue
		}
		var last *record.Snapshot
		for _, v := range vs {
			s, err := store.Snapshot(id, v)
			if err != nil {
				t.Error(err)
				break
			}
			if last != nil && s.ContentHash == last.ContentHash {
				reasonOnly[id]++
			}
			last = s
		}
		if last != nil {
			if h, _ := r.ContentHash(store.Specs()); h != last.ContentHash {
				t.Errorf("%s: file content_hash %s, head snapshot %s", id, h, last.ContentHash)
			}
		}
	}
	again, err := Load(root, LoadOptions{Git: fakeGit{}})
	if err != nil {
		t.Fatal(err)
	}
	rp, err := MakePlan(again)
	if err != nil {
		t.Fatal(err)
	}
	if len(rp.Entries) != 0 || len(rp.KeptDeprecated) != 0 {
		t.Errorf("re-plan not empty: %d entries, %d kept bugs", len(rp.Entries), len(rp.KeptDeprecated))
	}
	return reasonOnly
}

// TestApplyRealReposOnCopies migrates temp copies of the four repositories.
// The repositories themselves are only read.
func TestApplyRealReposOnCopies(t *testing.T) {
	ws := os.Getenv("PROVENANCE_REAL_REPOS")
	if ws == "" {
		t.Skip("set PROVENANCE_REAL_REPOS to run")
	}
	want := map[string][2]int{"linespec": {425, 465}, "sedum": {124, 124}, "scion": {218, 221}, "linespec-provenance-ui": {171, 171}}
	// Records whose superseded status is mapped, by mode (provisio-mgr 58929).
	wantMapped := map[string][2]int{"linespec": {0, 5}, "scion": {0, 3}}
	for _, name := range sortedKeys(want) {
		for i, opts := range []Options{{}, {AcceptSuggestions: true}} {
			t.Run(name+map[int]string{0: "/fold-all", 1: "/accept-suggestions"}[i], func(t *testing.T) {
				src := filepath.Join(ws, name)
				repo, err := Load(src, LoadOptions{Workspace: ws})
				if err != nil {
					t.Fatal(err)
				}
				p, err := MakePlan(repo)
				if err != nil {
					t.Fatal(err)
				}
				tmp := t.TempDir()
				copyDir(t, filepath.Join(src, repo.Dir), filepath.Join(tmp, repo.Dir))
				cfg, _ := os.ReadFile(filepath.Join(src, ".linespec.yml"))
				writeFile(t, filepath.Join(tmp, ".linespec.yml"), string(cfg))
				store := &record.FilesStore{Root: tmp, Dir: repo.Dir, SpecRoot: src}
				res, err := Apply(p, opts, store)
				if err != nil {
					t.Fatal(err)
				}
				if n := len(p.StatusMappings(opts)); n != wantMapped[name][i] {
					t.Errorf("status mappings %d, want %d", n, wantMapped[name][i])
				}
				for _, w := range p.StatusMappings(opts) {
					if w.Fields["to"] != string(record.StatusImplemented) {
						t.Errorf("%v: sealed record not mapped to implemented", w.Fields)
					}
				}
				if res.Records != want[name][i] {
					t.Errorf("records after %d, want %d", res.Records, want[name][i])
				}
				reasonOnly := checkMigrated(t, tmp, repo.Dir, store)
				t.Logf("%d records, %d snapshots, %d removed; reason-only versions: %v", res.Records, res.Snapshots, len(res.Removed), reasonOnly)
			})
		}
	}
}

func copyDir(t *testing.T, from, to string) {
	t.Helper()
	entries, err := os.ReadDir(from)
	if err != nil {
		t.Fatal(err)
	}
	for _, e := range entries {
		if e.IsDir() || !(strings.HasSuffix(e.Name(), ".yml") || strings.HasSuffix(e.Name(), ".yaml")) {
			continue
		}
		data, err := os.ReadFile(filepath.Join(from, e.Name()))
		if err != nil {
			t.Fatal(err)
		}
		writeFile(t, filepath.Join(to, e.Name()), string(data))
	}
}
