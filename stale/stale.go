// Package stale finds children that are behind their parents and lets a
// person confirm a child against a newer parent without a new version
// (DESIGN.md, "When a parent changes").
package stale

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/livecodelife/provenance/record"
)

// Kind classifies a finding.
type Kind string

const (
	// Direct: the child's implements_version is below its parent's version.
	Direct Kind = "direct"
	// Transitive: the child's own pin is current but an ancestor is directly
	// stale. Re-pinning that ancestor clears it; a new version of the
	// ancestor makes the child directly stale instead.
	Transitive Kind = "transitive"
	// Unknown: the pin cannot be judged (unpinned, cross-repo parent, or a
	// parent that does not resolve). Reported, never stale.
	Unknown Kind = "unknown"
)

// Finding is one child that is behind, or whose pin cannot be judged.
type Finding struct {
	ID            string
	Parent        string
	Pinned        int
	ParentVersion int
	Kind          Kind
	// Cause is the nearest directly stale ancestor, for a transitive finding.
	Cause string
	// Why explains an unknown finding.
	Why string
}

// Find lists every non-deprecated record in store that is directly or
// transitively stale, plus those whose pins cannot be judged, sorted by id.
func Find(store record.Store) ([]Finding, error) {
	ids, err := store.List()
	if err != nil {
		return nil, err
	}
	recs := map[string]*record.Record{}
	for _, id := range ids {
		r, _, err := store.Load(id)
		if err != nil {
			return nil, err
		}
		recs[r.ID] = r
	}
	direct := map[string]Finding{}
	unknown := map[string]Finding{}
	for _, id := range ids {
		r := recs[id]
		if r.Implements == "" {
			continue
		}
		f := Finding{ID: r.ID, Parent: r.Implements, Pinned: r.ImplementsVersion}
		parent, local := recs[r.Implements]
		switch {
		case r.ImplementsRepo != "":
			f.Kind, f.Why = Unknown, "parent is in shared repo "+r.ImplementsRepo
		case !local:
			f.Kind, f.Why = Unknown, "parent does not resolve"
		case r.ImplementsVersion == 0:
			f.Kind, f.Why, f.ParentVersion = Unknown, "unpinned", parent.Version
		case parent.Version > r.ImplementsVersion:
			f.Kind, f.ParentVersion = Direct, parent.Version
		default:
			continue
		}
		if f.Kind == Direct {
			direct[r.ID] = f
		} else {
			unknown[r.ID] = f
		}
	}

	var out []Finding
	for _, id := range ids {
		r := recs[id]
		if r.Status == record.StatusDeprecated {
			continue
		}
		if f, ok := direct[id]; ok {
			out = append(out, f)
			continue
		}
		if f, ok := unknown[id]; ok {
			out = append(out, f)
			continue
		}
		if r.Implements == "" {
			continue
		}
		// Walk up: the nearest directly stale ancestor makes this transitive.
		seen := map[string]bool{id: true}
		for p := r.Implements; p != ""; {
			if seen[p] {
				break
			}
			seen[p] = true
			if _, ok := direct[p]; ok {
				parent := recs[r.Implements]
				out = append(out, Finding{ID: id, Parent: r.Implements, Pinned: r.ImplementsVersion,
					ParentVersion: parent.Version, Kind: Transitive, Cause: p})
				break
			}
			if _, ok := unknown[p]; ok {
				break
			}
			anc, ok := recs[p]
			if !ok {
				break
			}
			p = anc.Implements
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })
	return out, nil
}

// Pin is one confirmation that a child is still correct against a parent
// version, written by Repin because a repin creates no version.
type Pin struct {
	At     string `json:"at"`
	Parent string `json:"parent"`
	From   int    `json:"from"`
	To     int    `json:"to"`
	Reason string `json:"reason"`
	Author string `json:"author"`
}

// PinLog is the append-only record of repins.
type PinLog interface {
	Append(id string, p Pin) error
	Read(id string) ([]Pin, error)
}

// FilePinLog keeps pins at <Root>/<Dir>/.versions/<id>/pins.jsonl.
type FilePinLog struct {
	Root string
	Dir  string // records directory; "provenance" if empty
}

func (l FilePinLog) path(id string) string {
	d := l.Dir
	if d == "" {
		d = "provenance"
	}
	return filepath.Join(l.Root, d, ".versions", id, "pins.jsonl")
}

// Append adds one line; existing lines are never rewritten.
func (l FilePinLog) Append(id string, p Pin) error {
	path := l.path(id)
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	line, err := json.Marshal(p)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0o644)
	if err != nil {
		return err
	}
	if _, err := f.Write(append(line, '\n')); err != nil {
		f.Close()
		return err
	}
	return f.Close()
}

// Read returns every pin for id, oldest first.
func (l FilePinLog) Read(id string) ([]Pin, error) {
	f, err := os.Open(l.path(id))
	if os.IsNotExist(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var out []Pin
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		if strings.TrimSpace(sc.Text()) == "" {
			continue
		}
		var p Pin
		if err := json.Unmarshal(sc.Bytes(), &p); err != nil {
			return nil, fmt.Errorf("%s: %w", l.path(id), err)
		}
		out = append(out, p)
	}
	return out, sc.Err()
}

var (
	ErrDeprecated  = errors.New("record is deprecated")
	ErrNoParent    = errors.New("record has no local parent to pin against")
	ErrNoReason    = errors.New("a repin needs a reason")
	ErrAlreadyThat = errors.New("record is already pinned to that version")
)

// Repin confirms id against its parent's version toVersion (the parent's
// current version when 0). It moves only implements_version: no version
// bump, so content_hash and every snapshot stay as they are and nothing
// regenerates. The confirmation and its reason go to the pin log.
func Repin(store record.Store, log PinLog, id string, toVersion int, reason, author string, now time.Time) (*Pin, error) {
	if strings.TrimSpace(reason) == "" {
		return nil, ErrNoReason
	}
	r, doc, err := store.Load(id)
	if err != nil {
		return nil, err
	}
	if r.Status == record.StatusDeprecated {
		return nil, ErrDeprecated
	}
	if r.Implements == "" || r.ImplementsRepo != "" {
		return nil, ErrNoParent
	}
	parent, _, err := store.Load(r.Implements)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrNoParent, err)
	}
	to := toVersion
	if to == 0 {
		to = parent.Version
	}
	if to < 1 || to > parent.Version {
		return nil, fmt.Errorf("version %d is outside %s's versions 1..%d", to, parent.ID, parent.Version)
	}
	if to == r.ImplementsVersion {
		return nil, ErrAlreadyThat
	}
	p := Pin{At: now.UTC().Format(time.RFC3339), Parent: parent.ID, From: r.ImplementsVersion, To: to, Reason: reason, Author: author}
	r.ImplementsVersion = to
	if err := store.Save(r, doc); err != nil {
		return nil, err
	}
	if err := log.Append(id, p); err != nil {
		return nil, fmt.Errorf("pinned %s to %s@%d but could not log it: %w", id, parent.ID, to, err)
	}
	return &p, nil
}
