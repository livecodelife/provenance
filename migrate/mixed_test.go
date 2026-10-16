package migrate

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/livecodelife/provenance/record"
)

// ids used by the mixed fixture.
const (
	mB1  = "prov-2026-0000b001" // v2 blueprint, a bug folds in
	mB2  = "prov-2026-0000b002" // v2 blueprint, a v1 successor folds in
	mD1  = "prov-2026-0000d001" // v2 draft: no snapshot to chain from
	mU1  = "prov-2026-0000e001" // v2 blueprint nothing touches
	mC1  = "prov-2026-0000c001" // v2 imprint implementing a v1 id that is absorbed
	mF1  = "prov-2026-0000f001" // v1 bug extending B1
	mS1  = "prov-2026-0000f002" // v1 blueprint superseding B2
	mX1  = "prov-2026-0000f003" // v1 bug extending the draft D1
	mVP  = "prov-2026-0000a010" // v1 blueprint superseded by VQ
	mVQ  = "prov-2026-0000a011" // v1 blueprint superseding VP
	mP1  = "prov-2026-0000a020" // v1 imprint implementing B1
	mW1  = "prov-2026-0000a030" // plain v1 blueprint
	cfgY = "dir: provenance\n"
)

var now = time.Date(2026, 10, 15, 20, 0, 0, 0, time.UTC)

func v2rec(id string, typ record.Type, status record.Status) *record.Record {
	return &record.Record{Schema: 2, ID: id, Title: "T " + id, Type: typ, Status: status, Version: 1,
		CreatedAt: "2026-01-01", Author: "a@b", Intent: "Intent of " + id,
		Constraints: []string{"Shared"}, AffectedScope: []string{"a.go"}}
}

// mixedRepo builds a repository the way a rebased branch looks after main's
// migration: .provenance.yml, migrated records with snapshots, and v1 records.
func mixedRepo(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	writeFile(t, filepath.Join(root, ".provenance.yml"), cfgY)
	store := &record.FilesStore{Root: root}
	for _, r := range []*record.Record{
		v2rec(mB1, record.TypeBlueprint, record.StatusOpen),
		v2rec(mB2, record.TypeBlueprint, record.StatusOpen),
		v2rec(mU1, record.TypeBlueprint, record.StatusOpen),
	} {
		if _, err := record.Commit(store, r, nil, record.Bump{Author: "a@b", Now: now}); err != nil {
			t.Fatal(err)
		}
	}
	c1 := v2rec(mC1, record.TypeImprint, record.StatusOpen)
	c1.Implements, c1.ImplementsVersion = mVP, 1
	if _, err := record.Commit(store, c1, nil, record.Bump{Author: "a@b", Now: now}); err != nil {
		t.Fatal(err)
	}
	if err := store.Save(v2rec(mD1, record.TypeBlueprint, record.StatusDraft), nil); err != nil {
		t.Fatal(err)
	}

	dir := filepath.Join(root, "provenance")
	rep := func(s, from, to string) string { return strings.Replace(s, from, to, 1) }
	bug := rep(v1(mF1, "bug", "implemented", "extends: "+mB1+"\n"), "  - Shared\n", "  - Shared  \n  - Also validate\n")
	bug = rep(bug, "  - a.go\n", "  - a.go\n  - b1.go\n")
	writeFile(t, filepath.Join(dir, mF1+".yml"), bug)
	writeFile(t, filepath.Join(dir, mS1+".yml"), rep(v1(mS1, "blueprint", "implemented", ""), `supersedes: ""`, "supersedes: "+mB2))
	writeFile(t, filepath.Join(dir, mX1+".yml"), v1(mX1, "bug", "implemented", "extends: "+mD1+"\n"))
	writeFile(t, filepath.Join(dir, mVP+".yml"), rep(v1(mVP, "blueprint", "superseded", ""), `superseded_by: ""`, "superseded_by: "+mVQ))
	writeFile(t, filepath.Join(dir, mVQ+".yml"), rep(v1(mVQ, "blueprint", "implemented", ""), `supersedes: ""`, "supersedes: "+mVP))
	writeFile(t, filepath.Join(dir, mP1+".yml"), v1(mP1, "imprint", "implemented", "implements: "+mB1+"\n"))
	writeFile(t, filepath.Join(dir, mW1+".yml"), v1(mW1, "blueprint", "implemented", ""))
	return root
}

// tree reads every file under the records directory, path to content.
func tree(t *testing.T, root string) map[string]string {
	t.Helper()
	out := map[string]string{}
	err := filepath.WalkDir(filepath.Join(root, "provenance"), func(p string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		b, rerr := os.ReadFile(p)
		if rerr != nil {
			return rerr
		}
		rel, _ := filepath.Rel(root, p)
		out[rel] = string(b)
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	return out
}

func loadPlan(t *testing.T, root string) (*Repo, *Plan) {
	t.Helper()
	repo, err := Load(root, LoadOptions{Git: fakeGit{}})
	if err != nil {
		t.Fatal(err)
	}
	p, err := MakePlan(repo)
	if err != nil {
		t.Fatal(err)
	}
	return repo, p
}

func warningsOf(p *Plan, kind string) []Warning {
	var out []Warning
	for _, w := range p.Warnings {
		if w.Kind == kind {
			out = append(out, w)
		}
	}
	return out
}

func TestMixedPlan(t *testing.T) {
	_, p := loadPlan(t, mixedRepo(t))

	want := map[string]string{mF1: mB1 + "/4", mS1: mB2 + "/2", mVP: mVQ + "/2", mX1: mD1 + "/4"}
	for id, w := range want {
		es := p.Entries[id]
		if len(es) != 1 || es[0].Target+"/"+es[0].Rule != w {
			t.Errorf("entry for %s = %+v, want %s", id, es, w)
		}
	}
	if len(p.Entries) != len(want) {
		t.Errorf("entries %v, want only %v", p.Entries, want)
	}
	if e := p.Entries[mX1][0]; e.Blocked == "" {
		t.Error("fold into the snapshot-less v2 draft is not blocked")
	}
	for id, s := range map[string]string{mF1: mB1, mS1: mB2, mVP: mVQ, mX1: mX1, mB1: mB1, mD1: mD1, mC1: mC1} {
		if p.Survivor[id] != s {
			t.Errorf("survivor of %s = %s, want %s", id, p.Survivor[id], s)
		}
	}
	if got := strings.Join(p.Versions[mB1], ","); got != mB1+","+mF1 {
		t.Errorf("B1 versions %s", got)
	}
	if v := p.versionOf(mF1); v != 2 {
		t.Errorf("F1 becomes version %d of B1, want 2", v)
	}

	if w := warningsOf(p, "migrated_target_without_snapshot"); len(w) != 1 || w[0].Fields["id"] != mX1 {
		t.Errorf("draft-target warning: %+v", w)
	}
	w := warningsOf(p, "migrated_child_points_at_absorbed")
	if len(w) != 1 || w[0].Fields["child"] != mC1 || w[0].Fields["survivor"] != mVQ {
		t.Fatalf("absorbed-parent warning: %+v", w)
	}
	if hint, _ := w[0].Fields["hint"].(string); !strings.Contains(hint, "provenance revise "+mC1) {
		t.Errorf("hint lacks the revise command: %q", hint)
	}
	// Migrated records are never plan sources.
	for _, v2 := range []string{mB1, mB2, mU1, mC1, mD1} {
		if _, ok := p.Entries[v2]; ok {
			t.Errorf("migrated record %s has an entry", v2)
		}
	}
}

func TestMixedApply(t *testing.T) {
	root := mixedRepo(t)
	before := tree(t, root)
	repo, p := loadPlan(t, root)
	store := &record.FilesStore{Root: root, Dir: repo.Dir}
	res, err := Apply(p, Options{}, store)
	if err != nil {
		t.Fatal(err)
	}
	after := tree(t, root)

	// Untouched migrated records: files and snapshots byte-identical.
	for path, content := range before {
		touched := strings.Contains(path, mB1) || strings.Contains(path, mB2)
		gone := strings.Contains(path, mF1) || strings.Contains(path, mS1) || strings.Contains(path, mVP)
		switch {
		case gone:
			if _, ok := after[path]; ok && strings.HasSuffix(path, ".yml") && !strings.Contains(path, ".versions") {
				t.Errorf("%s should have been removed", path)
			}
		case !touched && strings.Contains(path, ".versions") || strings.Contains(path, mU1) || strings.Contains(path, mC1) || strings.Contains(path, mD1):
			if after[path] != content {
				t.Errorf("%s changed", path)
			}
		}
	}
	// B1/B2 keep their existing snapshot byte for byte.
	for path, content := range before {
		if strings.Contains(path, ".versions") && (strings.Contains(path, mB1) || strings.Contains(path, mB2)) && after[path] != content {
			t.Errorf("existing snapshot %s changed", path)
		}
	}

	b1, _, _ := store.Load(mB1)
	if b1.Version != 2 || b1.ReasonKind != record.ReasonKindFix || b1.Reason != "Intent of "+mF1 {
		t.Fatalf("B1 = v%d %q %q", b1.Version, b1.Reason, b1.ReasonKind)
	}
	if strings.Join(b1.Aliases, ",") != mF1 || strings.Join(b1.AffectedScope, ",") != "a.go,b1.go" || strings.Join(b1.Constraints, "|") != "Shared|Also validate" {
		t.Errorf("B1 aliases %v scope %v constraints %v", b1.Aliases, b1.AffectedScope, b1.Constraints)
	}
	s1, err := store.Snapshot(mB1, 1)
	if err != nil {
		t.Fatal(err)
	}
	s2, err := store.Snapshot(mB1, 2)
	if err != nil {
		t.Fatal(err)
	}
	if s2.ParentHash != s1.ContentHash || s2.ContentHash == s1.ContentHash {
		t.Errorf("B1 snapshot chain: v1 %s, v2 parent %s hash %s", s1.ContentHash, s2.ParentHash, s2.ContentHash)
	}

	b2, _, _ := store.Load(mB2)
	if b2.Version != 2 || b2.Reason != "Supersedes "+mB2+": T "+mS1 || b2.ReasonKind != "" || b2.Title != "T "+mS1 {
		t.Errorf("B2 = v%d %q %q title %q", b2.Version, b2.Reason, b2.ReasonKind, b2.Title)
	}
	if strings.Join(b2.Aliases, ",") != mS1 {
		t.Errorf("B2 aliases %v", b2.Aliases)
	}

	x1, _, err := store.Load(mX1)
	if err != nil {
		t.Fatal(err)
	}
	if x1.Schema != 2 || x1.Type != record.TypeImprint || x1.Implements != mD1 || x1.ImplementsVersion != 0 {
		t.Errorf("blocked bug X1 = %+v", x1)
	}
	p1, _, _ := store.Load(mP1)
	if p1.Implements != mB1 || p1.ImplementsVersion != 1 {
		t.Errorf("P1 pin %s@%d, want %s@1 (stale against B1 v2)", p1.Implements, p1.ImplementsVersion, mB1)
	}
	vq, _, _ := store.Load(mVQ)
	if vq.Version != 2 || strings.Join(vq.Aliases, ",") != mVP {
		t.Errorf("VQ v%d aliases %v", vq.Version, vq.Aliases)
	}
	c1, _, _ := store.Load(mC1)
	if c1.Implements != mVP {
		t.Errorf("C1 was rewritten: implements %s", c1.Implements)
	}

	if strings.Join(res.Removed, ",") != mVP+","+mF1+","+mS1 {
		t.Errorf("removed %v", res.Removed)
	}
	for _, id := range []string{mF1, mS1, mVP} {
		if _, err := os.Stat(filepath.Join(root, "provenance", id+".yml")); !os.IsNotExist(err) {
			t.Errorf("%s still exists", id)
		}
	}

	checkMixed(t, store)

	// Running it again is a no-op: empty plan, no writes.
	repo2, p2 := loadPlan(t, root)
	if len(p2.Entries) != 0 || len(p2.KeptDeprecated) != 0 || len(warningsOf(p2, "migrated_target_without_snapshot")) != 0 {
		t.Fatalf("second plan not empty: %+v", p2.Entries)
	}
	snapshot := tree(t, root)
	res2, err := Apply(p2, Options{}, &record.FilesStore{Root: root, Dir: repo2.Dir})
	if err != nil {
		t.Fatal(err)
	}
	if res2.Records != 0 || res2.Snapshots != 0 || len(res2.Removed) != 0 {
		t.Errorf("second Apply wrote: %+v", res2)
	}
	again := tree(t, root)
	if len(again) != len(snapshot) {
		t.Errorf("file count %d -> %d", len(snapshot), len(again))
	}
	for path, c := range snapshot {
		if again[path] != c {
			t.Errorf("second Apply changed %s", path)
		}
	}
}

// checkMixed: every record is schema 2 and valid, and every record with
// snapshots has a verifying chain whose head matches the file.
func checkMixed(t *testing.T, store *record.FilesStore) {
	t.Helper()
	ids, err := store.List()
	if err != nil {
		t.Fatal(err)
	}
	for _, id := range ids {
		r, _, err := store.Load(id)
		if err != nil {
			t.Fatal(err)
		}
		if r.Schema != 2 {
			t.Errorf("%s is still schema %d", id, r.Schema)
		}
		if err := r.Validate(); err != nil {
			t.Error(err)
		}
		vs, err := store.Versions(id)
		if err != nil {
			t.Fatal(err)
		}
		if len(vs) == 0 {
			continue // a v2 draft
		}
		var last *record.Snapshot
		for i, v := range vs {
			s, err := store.Snapshot(id, v)
			if err != nil || v != i+1 {
				t.Errorf("%s snapshot %d: %v", id, v, err)
				break
			}
			last = s
		}
		if last == nil {
			continue
		}
		if h, _ := r.ContentHash(store.Specs()); r.Version != last.Version || h != last.ContentHash {
			t.Errorf("%s: file v%d %s, head snapshot v%d %s", id, r.Version, h, last.Version, last.ContentHash)
		}
	}
}

func TestMixedDriftIsAnError(t *testing.T) {
	root := mixedRepo(t)
	path := filepath.Join(root, "provenance", mB1+".yml")
	data, _ := os.ReadFile(path)
	if err := os.WriteFile(path, []byte(strings.Replace(string(data), "Intent of "+mB1, "Edited without a bump", 1)), 0o644); err != nil {
		t.Fatal(err)
	}
	repo, err := Load(root, LoadOptions{Git: fakeGit{}})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := MakePlan(repo); err == nil || !strings.Contains(err.Error(), "changed since its latest snapshot") || !strings.Contains(err.Error(), "provenance revise "+mB1) {
		t.Fatalf("err = %v", err)
	}
	// Drift in a record nothing folds into is not this migration's business.
	path = filepath.Join(root, "provenance", mU1+".yml")
	data2, _ := os.ReadFile(path)
	_ = os.WriteFile(filepath.Join(root, "provenance", mB1+".yml"), data, 0o644)
	if err := os.WriteFile(path, []byte(strings.Replace(string(data2), "Intent of "+mU1, "Edited", 1)), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := MakePlan(mustLoad(t, root)); err != nil {
		t.Fatalf("drift in an unrelated record failed the plan: %v", err)
	}
}

func mustLoad(t *testing.T, root string) *Repo {
	t.Helper()
	repo, err := Load(root, LoadOptions{Git: fakeGit{}})
	if err != nil {
		t.Fatal(err)
	}
	return repo
}

func TestMixedLoadConfig(t *testing.T) {
	mk := func(files map[string]string) string {
		root := t.TempDir()
		for name, body := range files {
			writeFile(t, filepath.Join(root, name), body)
		}
		return root
	}
	rec := func(dir string) string {
		return v1("prov-2026-0000a001", "blueprint", "implemented", "") + "# " + dir + "\n"
	}

	// Only .provenance.yml.
	r1 := mk(map[string]string{".provenance.yml": "dir: records\n", "records/prov-2026-0000a001.yml": rec("records")})
	if repo, err := Load(r1, LoadOptions{Git: fakeGit{}}); err != nil || repo.Dir != "records" || len(repo.Records) != 1 {
		t.Fatalf("only .provenance.yml: %v %+v", err, repo)
	}
	// Only .linespec.yml (a v1 repository).
	r2 := mk(map[string]string{".linespec.yml": "provenance:\n  dir: provenance\n", "provenance/prov-2026-0000a001.yml": rec("v1")})
	if repo, err := Load(r2, LoadOptions{Git: fakeGit{}}); err != nil || repo.Dir != "provenance" {
		t.Fatalf("only .linespec.yml: %v", err)
	}
	// Both: .provenance.yml wins.
	r3 := mk(map[string]string{
		".provenance.yml": "dir: records\n", ".linespec.yml": "provenance:\n  dir: provenance\n",
		"records/prov-2026-0000a001.yml": rec("records"), "provenance/prov-2026-0000a002.yml": v1("prov-2026-0000a002", "blueprint", "implemented", ""),
	})
	if repo, err := Load(r3, LoadOptions{Git: fakeGit{}}); err != nil || repo.Dir != "records" || len(repo.Records) != 1 {
		t.Fatalf("both: %v %+v", err, repo)
	}
	// Neither: the error names both files.
	if _, err := Load(t.TempDir(), LoadOptions{Git: fakeGit{}}); err == nil || !strings.Contains(err.Error(), ".provenance.yml") || !strings.Contains(err.Error(), ".linespec.yml") {
		t.Fatalf("neither: %v", err)
	}
	// An explicit Config overrides both.
	cfg := Config{Dir: "records"}
	if repo, err := Load(r2, LoadOptions{Git: fakeGit{}, Config: &cfg}); err != nil || repo.Dir != "records" {
		t.Fatalf("explicit config: %v", err)
	}
}

// migrateInPlace does what `provenance migrate` does around the library:
// plan, apply, then move the config to .provenance.yml.
func migrateInPlace(t *testing.T, root string, g Git) *Result {
	t.Helper()
	repo, err := Load(root, LoadOptions{Git: g})
	if err != nil {
		t.Fatal(err)
	}
	p, err := MakePlan(repo)
	if err != nil {
		t.Fatal(err)
	}
	res, err := Apply(p, Options{}, &record.FilesStore{Root: root, Dir: repo.Dir})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(root, ".provenance.yml")); os.IsNotExist(err) {
		writeFile(t, filepath.Join(root, ".provenance.yml"), "dir: "+repo.Dir+"\n")
		os.Remove(filepath.Join(root, ".linespec.yml"))
	}
	return res
}

// TestMixedBranchRebase: a branch cut before the migration carries a v1
// record; main migrates; the branch rebases onto it and migrates again. The
// repositories are throwaway git repos under t.TempDir() (COMMON.md git rule).
func TestMixedBranchRebase(t *testing.T) {
	tmp := t.TempDir()
	home := filepath.Join(tmp, "home")
	repo := filepath.Join(tmp, "repo")
	for _, d := range []string{home, repo} {
		os.MkdirAll(d, 0o755)
	}
	git := func(args ...string) { isolatedGit(t, home, repo, args...) }
	const P, Q, F = "prov-2026-0000a001", "prov-2026-0000a002", "prov-2026-0000f001"
	dir := filepath.Join(repo, "provenance")

	git("init", "-q", "-b", "main")
	writeFile(t, filepath.Join(repo, ".linespec.yml"), "provenance:\n  dir: provenance\n")
	writeFile(t, filepath.Join(dir, P+".yml"), v1(P, "blueprint", "implemented", ""))
	writeFile(t, filepath.Join(dir, Q+".yml"), v1(Q, "imprint", "implemented", "implements: "+P+"\n"))
	git("add", "-A")
	git("commit", "-q", "-m", "v1 records")

	git("checkout", "-q", "-b", "feature")
	writeFile(t, filepath.Join(dir, F+".yml"), v1(F, "bug", "implemented", "extends: "+P+"\n"))
	git("add", "-A")
	git("commit", "-q", "-m", "a bug cut before the migration")

	git("checkout", "-q", "main")
	if res := migrateInPlace(t, repo, GitCLI{}); res.Records != 2 {
		t.Fatalf("main migration: %+v", res)
	}
	git("add", "-A")
	git("commit", "-q", "-m", "migrate main to v2")

	git("checkout", "-q", "feature")
	git("rebase", "-q", "main")
	if _, err := os.Stat(filepath.Join(repo, ".linespec.yml")); !os.IsNotExist(err) {
		t.Fatal("rebase should have brought main's config move")
	}
	qBefore, _ := os.ReadFile(filepath.Join(dir, Q+".yml"))
	snapsBefore := tree(t, repo)

	res := migrateInPlace(t, repo, GitCLI{})
	if res.Records != 1 || res.Snapshots != 1 || strings.Join(res.Removed, ",") != F {
		t.Fatalf("branch migration: %+v", res)
	}
	store := &record.FilesStore{Root: repo}
	pr, _, err := store.Load(P)
	if err != nil {
		t.Fatal(err)
	}
	if pr.Version != 2 || strings.Join(pr.Aliases, ",") != F {
		t.Errorf("P v%d aliases %v", pr.Version, pr.Aliases)
	}
	if qAfter, _ := os.ReadFile(filepath.Join(dir, Q+".yml")); string(qAfter) != string(qBefore) {
		t.Error("Q (migrated on main, untouched on the branch) changed")
	}
	for path, c := range snapsBefore {
		if strings.Contains(path, ".versions") && tree(t, repo)[path] != c {
			t.Errorf("snapshot %s changed", path)
		}
	}
	checkMixed(t, store)

	// And again: nothing left to do.
	if res := migrateInPlace(t, repo, GitCLI{}); res.Records != 0 || res.Snapshots != 0 {
		t.Errorf("third migration wrote: %+v", res)
	}
}
