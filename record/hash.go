package record

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"unicode/utf8"

	"golang.org/x/text/unicode/norm"
)

// hashPrefix starts every hash this package produces.
const hashPrefix = "sha256:"

// GenInput is exactly what sedum reads from a record, and so exactly what
// gen_hash covers. A nil slice and an empty slice are the same value.
type GenInput struct {
	Intent                                     string
	Constraints, AffectedScope, ForbiddenScope []string
}

// GenHash returns the gen_hash of in as "sha256:<64 lowercase hex>"
// (docs/HASHING.md §3).
func GenHash(in GenInput) string {
	return hashObject(map[string]any{
		"intent":          normalize(in.Intent),
		"constraints":     normalizeList(in.Constraints),
		"affected_scope":  normalizeList(in.AffectedScope),
		"forbidden_scope": normalizeList(in.ForbiddenScope),
	})
}

// GenInput returns the fields of r that gen_hash covers.
func (r *Record) GenInput() GenInput {
	return GenInput{
		Intent:         r.Intent,
		Constraints:    r.Constraints,
		AffectedScope:  r.AffectedScope,
		ForbiddenScope: r.ForbiddenScope,
	}
}

// GenHash returns r's gen_hash.
func (r *Record) GenHash() string { return GenHash(r.GenInput()) }

// SpecResolver returns the spec hash of the file at path (relative to the
// repository root), or ok=false when the file cannot be read.
type SpecResolver func(path string) (hash string, ok bool)

// FileSpecResolver resolves spec paths against root on disk.
func FileSpecResolver(root string) SpecResolver {
	return func(path string) (string, bool) {
		data, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(path)))
		if err != nil {
			return "", false
		}
		return SpecHash(data), true
	}
}

// SpecHash returns the hash of a spec file's content, with CRLF pairs folded
// to LF (docs/HASHING.md §4).
func SpecHash(content []byte) string {
	return hashBytes(bytes.ReplaceAll(content, []byte("\r\n"), []byte("\n")))
}

// RecordFileMarker is what a spec that is a record file contributes to
// content_hash instead of its bytes (docs/HASHING.md section 4). It has no
// "sha256:" prefix, so it never equals a real spec hash.
const RecordFileMarker = "record-file/v1"

var recordFileName = regexp.MustCompile(`^prov-[0-9]{4}-[0-9a-f]{8}(-[A-Za-z0-9._-]+)?\.yml$`)

// IsRecordFilePath reports whether path names a record file: its base name
// (after the last "/") has the form prov-YYYY-<hex8>[-suffix].yml.
func IsRecordFilePath(path string) bool {
	if i := strings.LastIndex(path, "/"); i >= 0 {
		path = path[i+1:]
	}
	return recordFileName.MatchString(path)
}

// SpecHashes resolves each associated spec in order. A spec with no path or
// an unreadable file yields nil, and its path (or "" when there is none) is
// listed in missing.
func (r *Record) SpecHashes(specs SpecResolver) (hashes []*string, missing []string) {
	hashes = make([]*string, len(r.AssociatedSpecs))
	for i, s := range r.AssociatedSpecs {
		if IsRecordFilePath(s.Path) {
			m := RecordFileMarker
			hashes[i] = &m
			continue
		}
		if s.Path == "" || specs == nil {
			missing = append(missing, s.Path)
			continue
		}
		h, ok := specs(s.Path)
		if !ok {
			missing = append(missing, s.Path)
			continue
		}
		hashes[i] = &h
	}
	return hashes, missing
}

// ContentHash returns r's content_hash. Specs that cannot be resolved count
// as null and are returned in missing, so the caller can warn.
func (r *Record) ContentHash(specs SpecResolver) (hash string, missing []string) {
	hashes, missing := r.SpecHashes(specs)
	return ContentHashWithSpecs(r, hashes), missing
}

// ContentHashWithSpecs computes content_hash from already-resolved spec
// hashes, one per associated spec (nil for unresolved). Snapshots use it to
// recompute a hash without the spec files.
func ContentHashWithSpecs(r *Record, specHashes []*string) string {
	return hashObject(contentFields(r, specHashes))
}

// contentFields is the content_hash projection. It must match the table in
// docs/HASHING.md §4: behavioral fields only.
func contentFields(r *Record, specHashes []*string) map[string]any {
	specs := make([]any, len(specHashes))
	for i, h := range specHashes {
		if h != nil {
			specs[i] = *h
		}
	}
	return map[string]any{
		"intent":           normalize(r.Intent),
		"constraints":      normalizeList(r.Constraints),
		"affected_scope":   normalizeList(r.AffectedScope),
		"forbidden_scope":  normalizeList(r.ForbiddenScope),
		"associated_specs": specs,
	}
}

// Normalize applies docs/HASHING.md §2 to s: NFC, collapse each run of
// White_Space to one space, trim. Exported for corpus tests and for callers
// that compare normalized text.
func Normalize(s string) string { return normalize(s) }

func normalize(s string) string {
	if !utf8.ValidString(s) {
		// Defensive only: the decoder rejects invalid UTF-8. Ranging over the
		// string yields U+FFFD for each invalid byte.
		var b strings.Builder
		for _, r := range s {
			b.WriteRune(r)
		}
		s = b.String()
	}
	s = norm.NFC.String(s)
	var b strings.Builder
	b.Grow(len(s))
	inSpace := false
	for _, r := range s {
		if isWhiteSpace(r) {
			if !inSpace {
				b.WriteByte(' ')
			}
			inSpace = true
			continue
		}
		inSpace = false
		b.WriteRune(r)
	}
	return strings.TrimPrefix(strings.TrimSuffix(b.String(), " "), " ")
}

// isWhiteSpace is the Unicode White_Space property, listed explicitly so the
// set cannot drift with a library (docs/HASHING.md §2).
func isWhiteSpace(r rune) bool {
	switch {
	case r >= 0x09 && r <= 0x0D, r == 0x20, r == 0x85, r == 0xA0, r == 0x1680,
		r >= 0x2000 && r <= 0x200A, r == 0x2028, r == 0x2029, r == 0x202F,
		r == 0x205F, r == 0x3000:
		return true
	}
	return false
}

func normalizeList(in []string) []any {
	out := make([]any, len(in))
	for i, s := range in {
		out[i] = normalize(s)
	}
	return out
}

func hashObject(obj map[string]any) string {
	var b bytes.Buffer
	writeCanonical(&b, obj)
	return hashBytes(b.Bytes())
}

// Canonical returns the canonical JSON encoding of a projected object
// (docs/HASHING.md §5). Values must be string, bool, int, int64, uint64, nil, []any or
// map[string]any; floats and every other type panic.
func Canonical(v any) string {
	var b bytes.Buffer
	writeCanonical(&b, v)
	return b.String()
}

// GenCanonical and ContentCanonical expose the exact bytes each hash covers,
// for debugging hash mismatches against another implementation.
func GenCanonical(in GenInput) string {
	return Canonical(map[string]any{
		"intent":          normalize(in.Intent),
		"constraints":     normalizeList(in.Constraints),
		"affected_scope":  normalizeList(in.AffectedScope),
		"forbidden_scope": normalizeList(in.ForbiddenScope),
	})
}

func ContentCanonical(r *Record, specHashes []*string) string {
	return Canonical(contentFields(r, specHashes))
}

func writeCanonical(b *bytes.Buffer, v any) {
	switch v := v.(type) {
	case nil:
		b.WriteString("null")
	case string:
		writeString(b, v)
	case bool:
		writeBool(b, v)
	case int:
		writeInt(b, int64(v))
	case int64:
		writeInt(b, v)
	case uint64:
		writeUint(b, v)
	case []any:
		b.WriteByte('[')
		for i, x := range v {
			if i > 0 {
				b.WriteByte(',')
			}
			writeCanonical(b, x)
		}
		b.WriteByte(']')
	case map[string]any:
		keys := make([]string, 0, len(v))
		for k := range v {
			keys = append(keys, k)
		}
		sort.Strings(keys) // byte order
		b.WriteByte('{')
		for i, k := range keys {
			if i > 0 {
				b.WriteByte(',')
			}
			writeString(b, k)
			b.WriteByte(':')
			writeCanonical(b, v[k])
		}
		b.WriteByte('}')
	default:
		panic("record: non-canonical value type")
	}
}

func writeString(b *bytes.Buffer, s string) {
	const hexDigits = "0123456789abcdef"
	b.WriteByte('"')
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch {
		case c == '"':
			b.WriteString(`\"`)
		case c == '\\':
			b.WriteString(`\\`)
		case c < 0x20:
			b.WriteString(`\u00`)
			b.WriteByte(hexDigits[c>>4])
			b.WriteByte(hexDigits[c&0xF])
		default:
			b.WriteByte(c)
		}
	}
	b.WriteByte('"')
}

func hashBytes(data []byte) string {
	sum := sha256.Sum256(data)
	return hashPrefix + hex.EncodeToString(sum[:])
}
