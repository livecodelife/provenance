package scope

import (
	"testing"

	"github.com/livecodelife/provenance/record"
)

// Cases ported from linespec pkg/provenance provenance_test.go and
// exclude_paths_test.go, so the port keeps their behaviour.
func TestMatch(t *testing.T) {
	tests := []struct {
		path, pattern string
		want          bool
	}{
		{"src/main.go", "src/main.go", true},
		{"src/main.go", "src/other.go", false},
		{"src/main.go", "src/*.go", true},
		{"src/sub/file.go", "src/**/*.go", true},
		{"src/main.rb", "src/*.go", false},
		{"test/file_test.go", "**/*_test.go", true},
		{"src/test_123.go", `re:src/test_\d+\.go`, true},
		{"src/test_abc.go", `re:src/test_\d+\.go`, false},
		{"src/sub/main.go", "src/*.go", false},
		{"a.b", "a?b", true},
		{"a/b", "a?b", false},
	}
	for _, tt := range tests {
		got, err := Match(tt.path, tt.pattern)
		if err != nil || got != tt.want {
			t.Errorf("Match(%q, %q) = %v, %v; want %v", tt.path, tt.pattern, got, err, tt.want)
		}
	}
	if _, err := Match("x", "re:("); err == nil {
		t.Error("invalid regex not reported")
	}
}

func TestInScope(t *testing.T) {
	r := &record.Record{AffectedScope: []string{"src/**"}, ForbiddenScope: []string{"src/secret/*"}}
	for path, want := range map[string]bool{"src/a.go": true, "src/secret/key.go": false, "docs/x.md": false} {
		if got, _ := InScope(r, path); got != want {
			t.Errorf("InScope(%q) = %v, want %v", path, got, want)
		}
	}
	observed := &record.Record{ForbiddenScope: []string{"db/**"}}
	if got, _ := InScope(observed, "anything.go"); !got || Mode(observed) != "observed" {
		t.Error("observed mode should permit everything not forbidden")
	}
	if got, _ := InScope(observed, "db/schema.sql"); got {
		t.Error("forbidden_scope must win in observed mode")
	}
}

func TestExcluded(t *testing.T) {
	tests := []struct {
		path     string
		patterns []string
		want     bool
	}{
		{"CHANGELOG.md", []string{"CHANGELOG.md"}, true},
		{"README.md", []string{"CHANGELOG.md"}, false},
		{"docs/api.md", []string{"docs/"}, true},
		{"docs/api.md", []string{"docs"}, true},
		{"docs-old/file.md", []string{"docs"}, false},
		{"docs", []string{"docs/"}, false},
		{"migrations/v2/001_add_users.sql", []string{"migrations/"}, true},
		{"gen/x_gen.go", []string{`/.*_gen\.go$/`}, true},
		{"gen/x.go", []string{`/.*_gen\.go$/`}, false},
		{"pkg/a.pb.go", []string{"**/*.pb.go"}, true},
		{"x.go", []string{"/(/"}, false},
		{"x.go", []string{""}, false},
	}
	for _, tt := range tests {
		if got := Excluded(tt.path, tt.patterns); got != tt.want {
			t.Errorf("Excluded(%q, %v) = %v, want %v", tt.path, tt.patterns, got, tt.want)
		}
	}
}

func TestOverlap(t *testing.T) {
	tests := []struct {
		a, b string
		want bool
	}{
		{"src/a.go", "src/a.go", true},
		{"src/a.go", "src/*.go", true},
		{"src/*.go", "src/**", true},
		{"src/*.go", "docs/*.md", false},
		{"src/a.go", "docs/a.go", false},
	}
	for _, tt := range tests {
		if got := Overlap(tt.a, tt.b); got != tt.want {
			t.Errorf("Overlap(%q, %q) = %v, want %v", tt.a, tt.b, got, tt.want)
		}
	}
}
