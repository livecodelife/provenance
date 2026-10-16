package lint

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/livecodelife/provenance/record"
)

// Tag is one provenance reference in a commit message. Version is 0 for a
// bare [prov-…] tag, which resolves to the record's version at that commit.
type Tag struct {
	ID      string
	Version int
}

// tagPattern extends linespec's commit tag pattern with an optional @version.
var tagPattern = regexp.MustCompile(`\[(prov-\d{4}-(?:\d{3}|[a-f0-9]{8})(?:-[a-z0-9-]+)?)(?:@([1-9]\d*))?\]`)

// ParseCommitTags returns every [prov-…] and [prov-…@N] tag in msg, in
// order. Text that only resembles a tag (bad id, @0, @x) is not a tag.
func ParseCommitTags(msg string) []Tag {
	var tags []Tag
	for _, m := range tagPattern.FindAllStringSubmatch(msg, -1) {
		t := Tag{ID: m[1]}
		if m[2] != "" {
			t.Version, _ = strconv.Atoi(m[2])
		}
		tags = append(tags, t)
	}
	return tags
}

// String renders the tag as it appears in a commit message.
func (t Tag) String() string {
	if t.Version == 0 {
		return "[" + t.ID + "]"
	}
	return fmt.Sprintf("[%s@%d]", t.ID, t.Version)
}

// Resolved is the record version a commit tag names.
type Resolved struct {
	ID      string
	Version int
	// Snapshot is that version's snapshot; nil for a bare tag on a record
	// that has none yet (a draft).
	Snapshot *record.Snapshot
	// ViaAlias reports that the tag named an old id folded into ID.
	ViaAlias bool
}

// ResolveTag finds the record version a tag names. [id@N] is snapshot N of
// id (an error if it does not exist). A bare [id] is id's current version.
// An id that only appears in some record's aliases resolves to that record
// (an error if two records list it).
//
// A bare tag means "the record's version at that commit". ResolveTag cannot
// know commit time, so it answers with the current version, which is right
// for a commit being made (check --staged). To resolve a historical commit,
// the caller works out the version as of that commit (e.g. from git) and
// passes it as Tag.Version.
func ResolveTag(store record.Store, t Tag) (*Resolved, error) {
	ids, err := store.List()
	if err != nil {
		return nil, err
	}
	res := &Resolved{ID: t.ID}
	found := false
	for _, id := range ids {
		if id == t.ID {
			found = true
			break
		}
	}
	if !found {
		var owners []string
		for _, id := range ids {
			r, _, err := store.Load(id)
			if err != nil {
				return nil, err
			}
			for _, a := range r.Aliases {
				if a == t.ID {
					owners = append(owners, r.ID)
				}
			}
		}
		switch len(owners) {
		case 0:
			return nil, fmt.Errorf("%s is not a record or an alias of one", t.ID)
		case 1:
			res.ID, res.ViaAlias = owners[0], true
		default:
			return nil, fmt.Errorf("%s is an alias of more than one record: %v", t.ID, owners)
		}
	}
	r, _, err := store.Load(res.ID)
	if err != nil {
		return nil, err
	}
	res.Version = r.Version
	if t.Version != 0 {
		res.Version = t.Version
	}
	versions, err := store.Versions(res.ID)
	if err != nil {
		return nil, err
	}
	for _, v := range versions {
		if v == res.Version {
			snap, err := store.Snapshot(res.ID, v)
			if err != nil {
				return nil, err
			}
			res.Snapshot = snap
			return res, nil
		}
	}
	if t.Version != 0 {
		return nil, fmt.Errorf("%s has no version %d", res.ID, t.Version)
	}
	return res, nil
}

// Transition is one allowed status change. Bump means the change happens only
// together with a new version (revise or edit reopening a record).
type Transition struct {
	From, To record.Status
	Bump     bool
}

// Transitions are the v2 lifecycle, shared by the CLI and the UI. Unlike v1,
// an implemented record is not frozen: a new version reopens it. Deprecated is
// terminal.
var Transitions = []Transition{
	{record.StatusDraft, record.StatusOpen, false},
	{record.StatusDraft, record.StatusDeprecated, false},
	{record.StatusOpen, record.StatusImplemented, false},
	{record.StatusOpen, record.StatusDeprecated, false},
	{record.StatusImplemented, record.StatusOpen, true},
	{record.StatusImplemented, record.StatusDeprecated, false},
}

// CheckTransition reports whether a record may move from one status to
// another; bump says whether a new version is being written in the same step.
func CheckTransition(from, to record.Status, bump bool) error {
	if from == to {
		return nil
	}
	for _, t := range Transitions {
		if t.From == from && t.To == to {
			if t.Bump && !bump {
				return fmt.Errorf("%s -> %s needs a new version: use revise or edit", from, to)
			}
			return nil
		}
	}
	if from == record.StatusDeprecated {
		return fmt.Errorf("deprecated is terminal: a deprecated record takes no new versions or statuses")
	}
	return fmt.Errorf("%s -> %s is not an allowed transition", from, to)
}
