package record

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Store is where records and their version snapshots live. The files backend
// keeps them in the repository; a Provisio API backend can implement the same
// interface later, so nothing above it changes when the canonical store moves.
type Store interface {
	// List returns every record id, sorted.
	List() ([]string, error)
	// Load returns the current record and the parsed document to save it with.
	Load(id string) (*Record, *Document, error)
	// Save writes the current record. doc may be nil for a new record.
	Save(r *Record, doc *Document) error
	// Versions returns the record's snapshot versions, ascending.
	Versions(id string) ([]int, error)
	// Snapshot reads and verifies one version.
	Snapshot(id string, version int) (*Snapshot, error)
	// Append writes a new snapshot. It never overwrites one.
	Append(s *Snapshot) error
	// Remove deletes a record's current file (migration folding an absorbed
	// id). Its snapshots, if any, are kept.
	Remove(id string) error
	// Specs resolves associated spec paths for content hashing.
	Specs() SpecResolver
}

var (
	// ErrDeprecated: a deprecated record is terminal and takes no new versions.
	ErrDeprecated = errors.New("record is deprecated and takes no new versions")
	// ErrNoChange: neither the content nor the reason gives a reason to bump.
	ErrNoChange = errors.New("content unchanged and no reason given")
	// ErrReasonRequired: every version after the first needs a reason.
	ErrReasonRequired = errors.New("a new version needs a reason")
	// ErrNotV2: records must be migrated before they can be versioned.
	ErrNotV2 = errors.New("record is schema 1; migrate it before versioning")
)

// Bump describes a new version.
type Bump struct {
	Reason     string
	ReasonKind string
	Author     string
	Now        time.Time
}

// Commit is the one path that creates a version. It snapshots r as the next
// version when its content_hash changed since the latest snapshot, or when
// b carries a reason (a reason-only version). An implemented record returns
// to open. On success r's Version and Reason are updated and r is saved.
func Commit(s Store, r *Record, doc *Document, b Bump) (*Snapshot, error) {
	if r.Schema != 2 {
		return nil, ErrNotV2
	}
	if r.Status == StatusDeprecated {
		return nil, ErrDeprecated
	}
	versions, err := s.Versions(r.ID)
	if err != nil {
		return nil, err
	}
	hash, _ := r.ContentHash(s.Specs())
	next, parent := 1, ""
	if len(versions) > 0 {
		latest, err := s.Snapshot(r.ID, versions[len(versions)-1])
		if err != nil {
			return nil, err
		}
		next, parent = latest.Version+1, latest.ContentHash
		if hash == parent && b.Reason == "" {
			return nil, ErrNoChange
		}
		if b.Reason == "" {
			return nil, ErrReasonRequired
		}
	}
	if b.ReasonKind != "" && b.ReasonKind != ReasonKindFix {
		return nil, fmt.Errorf("invalid reason kind %q", b.ReasonKind)
	}
	now := b.Now
	if now.IsZero() {
		now = time.Now()
	}
	snap := NewSnapshot(r, next, parent, b.Reason, b.ReasonKind, b.Author, now.UTC().Format(time.RFC3339), s.Specs())
	if err := s.Append(snap); err != nil {
		return nil, err
	}
	r.Version, r.Reason, r.ReasonKind = next, b.Reason, b.ReasonKind
	if r.Status == StatusImplemented {
		r.Status = StatusOpen
		r.SealedAtSHA = ""
	}
	if err := s.Save(r, doc); err != nil {
		return nil, fmt.Errorf("snapshot %s written but record not saved: %w", snap.FileName(), err)
	}
	return snap, nil
}

// FilesStore keeps records as provenance/<id>.yml under Root and snapshots
// under provenance/.versions/<id>/.
type FilesStore struct {
	// Root is the repository root; spec paths resolve against it.
	Root string
	// Dir is the records directory relative to Root; "provenance" if empty.
	Dir string
	// SpecRoot resolves associated spec paths; Root if empty. Migration sets
	// it to the source repository while writing to a copy.
	SpecRoot string
}

func (f *FilesStore) dir() string {
	d := f.Dir
	if d == "" {
		d = "provenance"
	}
	return filepath.Join(f.Root, d)
}

func (f *FilesStore) recordPath(id string) string { return filepath.Join(f.dir(), id+".yml") }

func (f *FilesStore) versionsDir(id string) string {
	return filepath.Join(f.dir(), ".versions", id)
}

func (f *FilesStore) Specs() SpecResolver {
	if f.SpecRoot != "" {
		return FileSpecResolver(f.SpecRoot)
	}
	return FileSpecResolver(f.Root)
}

func (f *FilesStore) Remove(id string) error { return os.Remove(f.recordPath(id)) }

func (f *FilesStore) List() ([]string, error) {
	matches, err := filepath.Glob(filepath.Join(f.dir(), "*.yml"))
	if err != nil {
		return nil, err
	}
	ids := make([]string, 0, len(matches))
	for _, m := range matches {
		ids = append(ids, strings.TrimSuffix(filepath.Base(m), ".yml"))
	}
	sort.Strings(ids)
	return ids, nil
}

func (f *FilesStore) Load(id string) (*Record, *Document, error) {
	doc, err := Load(f.recordPath(id))
	if err != nil {
		return nil, nil, err
	}
	return doc.Record(), doc, nil
}

func (f *FilesStore) Save(r *Record, doc *Document) error {
	return Save(f.recordPath(r.ID), doc, r)
}

// snapshotFiles maps version number to file name for one record.
func (f *FilesStore) snapshotFiles(id string) (map[int]string, error) {
	entries, err := os.ReadDir(f.versionsDir(id))
	if os.IsNotExist(err) {
		return map[int]string{}, nil
	}
	if err != nil {
		return nil, err
	}
	out := map[int]string{}
	for _, e := range entries {
		name := e.Name()
		dash := strings.IndexByte(name, '-')
		if e.IsDir() || dash < 1 || !strings.HasSuffix(name, ".yml") {
			continue
		}
		v, err := strconv.Atoi(name[:dash])
		if err != nil {
			continue
		}
		if prev, dup := out[v]; dup {
			return nil, fmt.Errorf("%s: two snapshots for version %d: %s and %s", id, v, prev, name)
		}
		out[v] = name
	}
	return out, nil
}

func (f *FilesStore) Versions(id string) ([]int, error) {
	files, err := f.snapshotFiles(id)
	if err != nil {
		return nil, err
	}
	vs := make([]int, 0, len(files))
	for v := range files {
		vs = append(vs, v)
	}
	sort.Ints(vs)
	return vs, nil
}

func (f *FilesStore) Snapshot(id string, version int) (*Snapshot, error) {
	files, err := f.snapshotFiles(id)
	if err != nil {
		return nil, err
	}
	name, ok := files[version]
	if !ok {
		return nil, fmt.Errorf("%s has no version %d", id, version)
	}
	data, err := os.ReadFile(filepath.Join(f.versionsDir(id), name))
	if err != nil {
		return nil, err
	}
	s, err := UnmarshalSnapshot(data)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", name, err)
	}
	if s.ID != id || s.Version != version || s.FileName() != name {
		return nil, fmt.Errorf("%s: envelope (%s v%d) does not match its file name", name, s.ID, s.Version)
	}
	return s, nil
}

// Append writes s. It checks the chain (version = latest+1, parent_hash =
// latest content_hash), refuses to overwrite a different snapshot, and is a
// no-op for an identical one.
func (f *FilesStore) Append(s *Snapshot) error {
	if err := s.Verify(); err != nil {
		return err
	}
	data, err := MarshalSnapshot(s)
	if err != nil {
		return err
	}
	path := filepath.Join(f.versionsDir(s.ID), s.FileName())
	if existing, err := os.ReadFile(path); err == nil {
		if sameBytes(existing, data) {
			return nil
		}
		return fmt.Errorf("snapshot %s already exists with different content", path)
	}
	versions, err := f.Versions(s.ID)
	if err != nil {
		return err
	}
	wantParent := ""
	if n := len(versions); n > 0 {
		latest, err := f.Snapshot(s.ID, versions[n-1])
		if err != nil {
			return err
		}
		if s.Version != latest.Version+1 {
			return fmt.Errorf("%s: version %d does not follow %d", s.ID, s.Version, latest.Version)
		}
		wantParent = latest.ContentHash
	} else if s.Version != 1 {
		return fmt.Errorf("%s: first snapshot must be version 1, got %d", s.ID, s.Version)
	}
	if s.ParentHash != wantParent {
		return fmt.Errorf("%s v%d: parent_hash %q, want %q", s.ID, s.Version, s.ParentHash, wantParent)
	}
	if err := os.MkdirAll(f.versionsDir(s.ID), 0o755); err != nil {
		return err
	}
	// O_EXCL makes a concurrent writer of the same version fail rather than
	// overwrite.
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o444)
	if err != nil {
		return err
	}
	if _, err := file.Write(data); err != nil {
		file.Close()
		return err
	}
	return file.Close()
}

var _ Store = (*FilesStore)(nil)
