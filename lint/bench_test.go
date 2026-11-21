package lint

import (
	"fmt"
	"os"
	"testing"

	"github.com/livecodelife/provenance/record"
)

// benchRepo is the repository the benchmarks run on: PROVENANCE_BENCH_REPO if
// set (a real checkout, read only), else a generated one.
func benchRepo(b *testing.B) (*record.FilesStore, Options, string) {
	b.Helper()
	if root := os.Getenv("PROVENANCE_BENCH_REPO"); root != "" {
		s := &record.FilesStore{Root: root}
		ids, err := s.List()
		if err != nil || len(ids) == 0 {
			b.Fatalf("no records in %s: %v", root, err)
		}
		return s, Options{Root: root, Enforcement: "strict", ManifestPath: record.ManifestPath}, ids[len(ids)/2]
	}
	root := b.TempDir()
	s := &record.FilesStore{Root: root}
	var last string
	for i := 0; i < 300; i++ {
		r := newRecord(fmt.Sprintf("prov-2026-%08x", 0x1000+i), record.StatusOpen)
		r.AffectedScope = []string{fmt.Sprintf("src/m%d/**", i%40), fmt.Sprintf("src/m%d/f%d.go", i%40, i), "re:^lib/x" + fmt.Sprint(i%7)}
		if _, err := record.Commit(s, r, nil, record.Bump{Author: "a@b", Reason: "r"}); err != nil {
			b.Fatal(err)
		}
		last = r.ID
	}
	lines, _ := record.BuildManifest(s)
	if err := os.MkdirAll(root+"/.provenance", 0o755); err != nil {
		b.Fatal(err)
	}
	if err := os.WriteFile(root+"/"+record.ManifestPath, record.MarshalManifest(lines), 0o644); err != nil {
		b.Fatal(err)
	}
	return s, Options{Root: root, Enforcement: "strict", ManifestPath: record.ManifestPath}, last
}

// BenchmarkLintFull is a whole-repository lint.
func BenchmarkLintFull(b *testing.B) {
	s, opts, _ := benchRepo(b)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := Lint(s, opts); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkLintFocused is a lint limited to one record, as `lint --record` and
// `check --staged` run it.
func BenchmarkLintFocused(b *testing.B) {
	s, opts, id := benchRepo(b)
	opts.Focus = map[string]bool{id: true}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := Lint(s, opts); err != nil {
			b.Fatal(err)
		}
	}
}
