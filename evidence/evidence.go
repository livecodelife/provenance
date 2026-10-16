// Package evidence keeps verification evidence per record version: which
// content, specs and code a spec run checked, whether it passed, and when.
// Evidence goes stale as soon as any of the three hashes moves, so Provisio
// can show "v4 verified, v5 unverified" (DESIGN.md, Associated specs).
package evidence

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/livecodelife/provenance/record"
)

// Region is one sedum region of a record, as `sedum regions --json` lists
// it. Only gen and pkg feed the code hash; other keys are ignored.
type Region struct {
	Gen string `json:"gen"`
	Pkg string `json:"pkg"`
	// Tier is the region's tier; only owned regions count. Empty counts as
	// owned, so a caller may filter beforehand.
	Tier string `json:"tier,omitempty"`
}

func owned(regions []Region) []Region {
	var out []Region
	for _, r := range regions {
		if r.Tier == "" || r.Tier == "owned" {
			out = append(out, r)
		}
	}
	return out
}

// ErrUnstamped: a region lacks gen or pkg, so its code cannot be hashed.
// Absent means stale on markers; `sedum grow --stamp` fills them in.
var ErrUnstamped = errors.New("a region has no gen or pkg; run sedum grow --stamp")

// CodeHash hashes a record's regions per docs/HASHING.md §7: each region's
// canonical {gen, pkg} object, sorted by bytes, as a canonical JSON array.
// Only owned regions count. It refuses an owned region without gen or pkg.
func CodeHash(regions []Region) (string, error) {
	regions = owned(regions)
	for _, r := range regions {
		if r.Gen == "" || r.Pkg == "" {
			return "", ErrUnstamped
		}
	}
	return hashOf(CodeCanonical(regions)), nil
}

// CodeCanonical returns the exact bytes CodeHash covers.
func CodeCanonical(regions []Region) string {
	regions = owned(regions)
	enc := make([]string, len(regions))
	for i, r := range regions {
		enc[i] = record.Canonical(map[string]any{"gen": r.Gen, "pkg": r.Pkg})
	}
	sort.Strings(enc)
	return "[" + strings.Join(enc, ",") + "]"
}

func hashOf(s string) string {
	sum := sha256.Sum256([]byte(s))
	return "sha256:" + hex.EncodeToString(sum[:])
}

// Evidence is one verification run of one record version.
type Evidence struct {
	RecordID    string    `json:"record_id"`
	Version     int       `json:"version"`
	ContentHash string    `json:"content_hash"`
	SpecHashes  []*string `json:"spec_hashes"`
	CodeHash    string    `json:"code_hash"`
	Passed      bool      `json:"passed"`
	// At is an RFC 3339 UTC timestamp.
	At string `json:"at"`
	// Runner says what ran the specs, for people (e.g. "provenance complete").
	Runner string `json:"runner,omitempty"`
}

// Current is what a record looks like now, to compare evidence against.
// CodeHash is empty when the code is unknown (an unstamped region). Behind
// is set by the caller when stale.Find lists the record (its pin, or an
// ancestor's, is behind), since a parent's change does not move the child's
// own hashes but does invalidate what was verified.
type Current struct {
	ContentHash string
	SpecHashes  []*string
	CodeHash    string
	Behind      bool
}

// CurrentFor computes r's present hashes from its spec files and regions.
func CurrentFor(r *record.Record, specs record.SpecResolver, regions []Region) Current {
	hashes, _ := r.SpecHashes(specs)
	code, _ := CodeHash(regions)
	return Current{ContentHash: record.ContentHashWithSpecs(r, hashes), SpecHashes: hashes, CodeHash: code}
}

// New records a run of r's current version against cur. A run against
// unknown code (empty CodeHash) is still recorded; it can never be fresh.
func New(r *record.Record, cur Current, passed bool, at, runner string) *Evidence {
	return &Evidence{RecordID: r.ID, Version: r.Version, ContentHash: cur.ContentHash,
		SpecHashes: cur.SpecHashes, CodeHash: cur.CodeHash, Passed: passed, At: at, Runner: runner}
}

// State is how evidence relates to the record now.
type State string

const (
	Fresh   State = "fresh"   // all three hashes still match
	Stale   State = "stale"   // at least one moved; Reasons says which
	Unknown State = "unknown" // nothing moved that can be checked, but the code is unknown
	Missing State = "missing" // no evidence at all
)

// Verdict is the result of Check.
type Verdict struct {
	State State
	// Reasons lists what moved ("content", "specs", "code", "parent") and
	// "code_unknown" when an unstamped region hides the code.
	Reasons []string
	// Passed is the run's outcome, meaningful when there is evidence.
	Passed   bool
	Evidence *Evidence
}

// Verified reports fresh evidence of a passing run.
func (v Verdict) Verified() bool { return v.State == Fresh && v.Passed }

// Check compares evidence with the record's current hashes. Evidence is
// judged by hashes, not version numbers, so it stays fresh across a
// reason-only version and goes stale on any change to what was checked.
func Check(e *Evidence, cur Current) Verdict {
	if e == nil {
		return Verdict{State: Missing}
	}
	v := Verdict{State: Fresh, Passed: e.Passed, Evidence: e}
	if e.ContentHash != cur.ContentHash {
		v.Reasons = append(v.Reasons, "content")
	}
	if !sameSpecs(e.SpecHashes, cur.SpecHashes) {
		v.Reasons = append(v.Reasons, "specs")
	}
	if cur.Behind {
		v.Reasons = append(v.Reasons, "parent")
	}
	codeUnknown := cur.CodeHash == "" || e.CodeHash == ""
	if !codeUnknown && e.CodeHash != cur.CodeHash {
		v.Reasons = append(v.Reasons, "code")
	}
	switch {
	case len(v.Reasons) > 0:
		v.State = Stale
	case codeUnknown:
		v.State = Unknown
	}
	if codeUnknown {
		v.Reasons = append(v.Reasons, "code_unknown")
	}
	return v
}

func sameSpecs(a, b []*string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if (a[i] == nil) != (b[i] == nil) || (a[i] != nil && *a[i] != *b[i]) {
			return false
		}
	}
	return true
}

// Store keeps evidence as provenance/.evidence/<id>/<version>-<unix-nanos>.json
// under Root. It is append-only: a run never replaces another.
type Store struct {
	Root string
	// Dir is the records directory relative to Root; "provenance" if empty.
	Dir string
}

func (s *Store) dir(id string) string {
	d := s.Dir
	if d == "" {
		d = "provenance"
	}
	return filepath.Join(s.Root, d, ".evidence", id)
}

// Append writes a run. Its file name comes from the version and At; when two
// runs share the name, the later one takes the next free nanosecond (see
// createExclusive), and an existing file is never overwritten.
func (s *Store) Append(e *Evidence) error {
	t, err := parseTime(e.At)
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return err
	}
	if err := os.MkdirAll(s.dir(e.RecordID), 0o755); err != nil {
		return err
	}
	f, err := createExclusive(s.dir(e.RecordID), e.Version, t)
	if err != nil {
		return err
	}
	if _, err := f.Write(append(data, '\n')); err != nil {
		f.Close()
		return err
	}
	return f.Close()
}

// List returns every run for id, oldest first.
func (s *Store) List(id string) ([]*Evidence, error) {
	entries, err := os.ReadDir(s.dir(id))
	if os.IsNotExist(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	type named struct {
		version int
		nanos   int64
		e       *Evidence
	}
	var runs []named
	for _, ent := range entries {
		name := strings.TrimSuffix(ent.Name(), ".json")
		v, n, ok := strings.Cut(name, "-")
		if ent.IsDir() || !ok || name == ent.Name() {
			continue
		}
		version, err1 := strconv.Atoi(v)
		nanos, err2 := strconv.ParseInt(n, 10, 64)
		if err1 != nil || err2 != nil {
			continue
		}
		data, err := os.ReadFile(filepath.Join(s.dir(id), ent.Name()))
		if err != nil {
			return nil, err
		}
		var e Evidence
		dec := json.NewDecoder(bytes.NewReader(data))
		dec.DisallowUnknownFields()
		if err := dec.Decode(&e); err != nil {
			return nil, fmt.Errorf("%s: %w", ent.Name(), err)
		}
		if e.RecordID != id || e.Version != version {
			return nil, fmt.Errorf("%s: content (%s v%d) does not match its name", ent.Name(), e.RecordID, e.Version)
		}
		runs = append(runs, named{version, nanos, &e})
	}
	sort.Slice(runs, func(i, j int) bool {
		if runs[i].nanos != runs[j].nanos {
			return runs[i].nanos < runs[j].nanos
		}
		return runs[i].version < runs[j].version
	})
	out := make([]*Evidence, len(runs))
	for i, r := range runs {
		out[i] = r.e
	}
	return out, nil
}

// Latest returns the newest run for id at version, or nil.
func (s *Store) Latest(id string, version int) (*Evidence, error) {
	runs, err := s.List(id)
	if err != nil {
		return nil, err
	}
	for i := len(runs) - 1; i >= 0; i-- {
		if runs[i].Version == version {
			return runs[i], nil
		}
	}
	return nil, nil
}

// Status checks r against the newest evidence for its current version.
func (s *Store) Status(r *record.Record, cur Current) (Verdict, error) {
	e, err := s.Latest(r.ID, r.Version)
	if err != nil {
		return Verdict{}, err
	}
	return Check(e, cur), nil
}

func parseTime(at string) (int64, error) {
	t, err := time.Parse(time.RFC3339Nano, at)
	if err != nil {
		return 0, fmt.Errorf("evidence time %q: %w", at, err)
	}
	return t.UnixNano(), nil
}
