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
