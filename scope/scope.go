// Package scope matches file paths against record scope patterns and
// exclude_paths entries. It is a verbatim port of linespec pkg/provenance's
// MatchPattern, GlobToRegex, IsInScope and IsPathExcluded, so the provenance
// CLI, lint and linespec agree on which files a record governs.
package scope

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/livecodelife/provenance/record"
)

// IsRegex reports whether a scope pattern is a regex ("re:" prefix).
func IsRegex(pattern string) bool { return len(pattern) > 3 && pattern[:3] == "re:" }

// IsPattern reports whether a scope entry is a glob or regex rather than an
// exact path.
func IsPattern(s string) bool {
	return strings.Contains(s, "*") || strings.Contains(s, "?") || IsRegex(s)
}

// Match reports whether filePath matches a scope pattern: "re:<regex>", a
// glob (* within a segment, ** across segments, ? one character), or an
// exact path.
func Match(filePath, pattern string) (bool, error) {
	if IsRegex(pattern) {
		re, err := regexp.Compile(pattern[3:])
		if err != nil {
			return false, fmt.Errorf("invalid regex pattern %q: %w", pattern, err)
		}
		return re.MatchString(filePath), nil
	}
	if strings.ContainsAny(pattern, "*?") {
		re, err := regexp.Compile(GlobToRegex(pattern))
		if err != nil {
			return false, fmt.Errorf("invalid glob pattern %q: %w", pattern, err)
		}
		return re.MatchString(filePath), nil
	}
	return filePath == pattern, nil
}

// GlobToRegex converts a glob to an anchored regex: ** matches across
// directories, * and ? stay within one path segment.
func GlobToRegex(glob string) string {
	var result strings.Builder
	for i := 0; i < len(glob); i++ {
		c := glob[i]
		switch c {
		case '*':
			if i+1 < len(glob) && glob[i+1] == '*' {
				result.WriteString(".*")
				i++
			} else {
				result.WriteString("[^/]*")
			}
		case '?':
			result.WriteString("[^/]")
		case '.', '+', '(', ')', '|', '^', '$':
			result.WriteString("\\" + string(c))
		case '\\':
			if i+1 < len(glob) {
				result.WriteString("\\" + string(glob[i+1]))
				i++
			}
		default:
			result.WriteByte(c)
		}
	}
	return "^" + result.String() + "$"
}

// Mode is "allowlist" when a record names an affected_scope and "observed"
// when it does not.
func Mode(r *record.Record) string {
	if len(r.AffectedScope) > 0 {
		return "allowlist"
	}
	return "observed"
}

// InScope reports whether r permits a change to filePath: never when
// forbidden_scope matches; in allowlist mode only when affected_scope
// matches; in observed mode always otherwise.
func InScope(r *record.Record, filePath string) (bool, error) {
	for _, pattern := range r.ForbiddenScope {
		ok, err := Match(filePath, pattern)
		if err != nil {
			return false, err
		}
		if ok {
			return false, nil
		}
	}
	if Mode(r) == "allowlist" {
		for _, pattern := range r.AffectedScope {
			ok, err := Match(filePath, pattern)
			if err != nil {
				return false, err
			}
			if ok {
				return true, nil
			}
		}
		return false, nil
	}
	return true, nil
}

// Excluded reports whether filePath matches an exclude_paths entry: a
// /regex/, a glob, an exact path, or a directory prefix ("docs" or "docs/").
// An invalid regex or glob matches nothing.
func Excluded(filePath string, excludePaths []string) bool {
	for _, pattern := range excludePaths {
		if matchesExclude(filePath, pattern) {
			return true
		}
	}
	return false
}

func matchesExclude(filePath, pattern string) bool {
	if pattern == "" {
		return false
	}
	if len(pattern) >= 2 && pattern[0] == '/' && pattern[len(pattern)-1] == '/' {
		re, err := regexp.Compile(pattern[1 : len(pattern)-1])
		if err != nil {
			return false
		}
		return re.MatchString(filePath)
	}
	if strings.ContainsAny(pattern, "*?") {
		ok, err := Match(filePath, pattern)
		return err == nil && ok
	}
	if filePath == pattern {
		return true
	}
	return strings.HasPrefix(filePath, strings.TrimSuffix(pattern, "/")+"/")
}

// Overlap reports whether two scope entries could match the same file. It
// keeps linespec's sampling heuristic for pattern-vs-pattern overlap.
func Overlap(a, b string) bool {
	if a == b {
		return true
	}
	if !IsPattern(a) {
		if ok, _ := Match(a, b); ok {
			return true
		}
	}
	if !IsPattern(b) {
		if ok, _ := Match(b, a); ok {
			return true
		}
	}
	if IsPattern(a) && IsPattern(b) {
		return patternsOverlap(a, b)
	}
	return false
}

func patternsOverlap(a, b string) bool {
	ar, br := toRegex(a), toRegex(b)
	if ar == nil || br == nil {
		return false
	}
	for _, s := range samplePaths(a) {
		if br.MatchString(s) {
			return true
		}
	}
	for _, s := range samplePaths(b) {
		if ar.MatchString(s) {
			return true
		}
	}
	return false
}

func toRegex(pattern string) *regexp.Regexp {
	expr := GlobToRegex(pattern)
	if IsRegex(pattern) {
		expr = pattern[3:]
	}
	re, err := regexp.Compile(expr)
	if err != nil {
		return nil
	}
	return re
}

func samplePaths(pattern string) []string {
	if IsRegex(pattern) {
		return []string{"test/file.go", "pkg/module/file.go", "src/test.go"}
	}
	p1 := strings.Replace(pattern, "**", "pkg/module", 1)
	p1 = strings.Replace(p1, "*", "file", 1)
	p1 = strings.Replace(p1, "?", "f", 1)
	p2 := strings.Replace(pattern, "**", "src/sub", 1)
	p2 = strings.Replace(p2, "*", "test", 1)
	p2 = strings.Replace(p2, "?", "t", 1)
	return []string{p1, p2}
}
