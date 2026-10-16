package migrate

import (
	"fmt"
	"strings"
	"time"

	"github.com/livecodelife/provenance/record"
)

// ManifestLine is one entry of the version manifest (record.ManifestLine).
type ManifestLine = record.ManifestLine

// Result summarizes what Apply wrote.
type Result struct {
	Records   int
	Snapshots int
	Removed   []string
	Manifest  []ManifestLine
}

// version is one snapshot being built for a surviving record.
type version struct {
	content    *record.Record
	reason     string
	reasonKind string
	member     *Source
}

// Apply migrates the planned repository into store, applying opts'
// keep-separate choices. Everything is computed and validated before the
// first write. It writes only through store, so callers point it at a copy.
func Apply(p *Plan, opts Options, store record.Store) (*Result, error) {
	R := p.Repo.Records
	ids := sortedIDs(R)

	surv := map[string]string{}
	members := map[string][]string{}
	for _, id := range ids {
		s := p.survivor(id, opts)
		if s == "" {
			return nil, fmt.Errorf("%s has no single survivor; resolve the plan first", id)
		}
		surv[id] = s
		members[s] = append(members[s], id)
	}
	order := map[string][]string{}
	pos := map[string]int{} // version number each id becomes
	for s, m := range members {
		order[s] = orderMembers(R, s, m)
		base := 1
		if R[s].V2() {
			base = R[s].Latest // its existing versions come first
		}
		for i, id := range order[s] {
			pos[id] = base + i
		}
	}

	// Related links that replace kept-separate non-bug edges.
	relatedAdd := map[string][]string{}
	for _, id := range ids {
		for _, e := range p.Entries[id] {
			if opts.kept(id, e) && R[id].Type() != record.TypeBug {
				// The link lives on the survivor of the pair, except that a
				// migrated record is never edited: then the v1 record holds it.
				holder, link := e.Target, id
				if R[e.Target].V2() {
					holder, link = id, e.Target
				}
				relatedAdd[holder] = append(relatedAdd[holder], link)
			}
		}
	}

	specs := store.Specs()
	type planned struct {
		head  *record.Record
		doc   *record.Document
		snaps []*record.Snapshot
	}
	var out []planned
	for _, s := range sortedKeys(order) {
		// A migrated record that nothing folds into is not touched at all:
		// its file and snapshots stay byte-identical.
		var base *record.Record
		first, parent := 1, ""
		ids := order[s]
		if R[s].V2() {
			if len(ids) == 1 {
				continue
			}
			base = R[s].Doc.Record()
			first, parent = R[s].Latest+1, R[s].LatestHash
			ids = ids[1:]
		}
		vs := p.buildVersions(ids, s, base, surv, pos, relatedAdd)
		var snaps []*record.Snapshot
		for i, v := range vs {
			snap := record.NewSnapshot(v.content, first+i, parent, v.reason, v.reasonKind, v.member.Rec.Author, createdAt(v.member), specs)
			if err := snap.Verify(); err != nil {
				return nil, err
			}
			snaps = append(snaps, snap)
			parent = snap.ContentHash
		}
		last := vs[len(vs)-1]
		head := cloneOf(last.content)
		own := R[s].Rec
		head.ID, head.Author, head.CreatedAt = s, own.Author, own.CreatedAt
		head.Version, head.Reason, head.ReasonKind = first+len(vs)-1, last.reason, last.reasonKind
		head.Status, head.SealedAtSHA = v2Status(last.member.Rec), last.member.Rec.SealedAtSHA
		if head.Status != record.StatusImplemented {
			head.SealedAtSHA = ""
		}
		head.Aliases = appendNew(own.Aliases, order[s]...)
		head.Aliases = without(head.Aliases, s)
		if err := head.Validate(); err != nil {
			return nil, fmt.Errorf("migrated %s: %w", s, err)
		}
		out = append(out, planned{head: head, doc: R[s].Doc, snaps: snaps})
	}

	res := &Result{}
	for _, pl := range out {
		for _, snap := range pl.snaps {
			if err := store.Append(snap); err != nil {
				return res, err
			}
			res.Snapshots++
		}
		if err := store.Save(pl.head, pl.doc); err != nil {
			return res, err
		}
		res.Records++
	}
	for _, id := range ids {
		if surv[id] != id {
			if err := store.Remove(id); err != nil {
				return res, err
			}
			res.Removed = append(res.Removed, id)
		}
	}
	manifest, err := record.BuildManifest(store)
	if err != nil {
		return res, err
	}
	res.Manifest = manifest
	return res, nil
}

// buildVersions turns the members that become versions into versions. For a
// migrated survivor, base is its current content and ids are only the new v1
// members; for a v1 survivor base is nil and ids start with its first version.
func (p *Plan) buildVersions(ids []string, s string, base *record.Record, surv map[string]string, pos map[string]int, relatedAdd map[string][]string) []version {
	R := p.Repo.Records
	var vs []version
	prev, prevID := base, ""
	if base != nil {
		prevID = s
	}
	for _, id := range ids {
		m := R[id]
		switch {
		case prev == nil:
			c := p.normalized(m, surv, pos)
			reason, kind := "", ""
			if m.Type() == record.TypeBug {
				// A bug that survives on its own (kept separate, deprecated,
				// or with no local target) becomes an imprint.
				p.bugAsImprint(c, m, surv, pos)
				reason, kind = strings.TrimSpace(m.Rec.Intent), record.ReasonKindFix
			}
			vs = append(vs, version{content: c, reason: reason, reasonKind: kind, member: m})
		case m.Type() == record.TypeBug:
			c := cloneOf(prev)
			c.AffectedScope = appendNew(c.AffectedScope, m.Rec.AffectedScope...)
			c.Constraints = appendNewNormalized(c.Constraints, m.Rec.Constraints)
			c.AssociatedSpecs = appendSpecs(c.AssociatedSpecs, m.Rec.AssociatedSpecs)
			vs = append(vs, version{content: c, reason: strings.TrimSpace(m.Rec.Intent), reasonKind: record.ReasonKindFix, member: m})
		default:
			c := p.normalized(m, surv, pos)
			vs = append(vs, version{content: c, reason: fmt.Sprintf("Supersedes %s: %s", prevID, record.Normalize(m.Rec.Title)), member: m})
		}
		prev, prevID = vs[len(vs)-1].content, id
	}
	for _, v := range vs {
		v.content.ID = s
		v.content.Related = appendNew(v.content.Related, relatedAdd[s]...)
	}
	return vs
}

// normalized returns a schema-2 copy of a member's record: explicit type,
// no v1-only fields, and implements re-pinned.
func (p *Plan) normalized(m *Source, surv map[string]string, pos map[string]int) *record.Record {
	c := m.Doc.Record()
	c.Schema = 2
	c.Type = m.Type()
	c.Supersedes, c.SupersededBy, c.Extends = "", "", ""
	c.Implements, c.ImplementsRepo, c.ImplementsVersion = p.pin(ref(c.Implements), surv, pos)
	return c
}

// pin resolves an implements target to its survivor and the version the
// target became; a shared target gets its repo name and an unknown version.
func (p *Plan) pin(target string, surv map[string]string, pos map[string]int) (id, repo string, version int) {
	if target == "" {
		return "", "", 0
	}
	if s, ok := surv[target]; ok {
		return s, "", pos[target]
	}
	if name, ok := p.Repo.SharedRepoOf[target]; ok {
		return target, name, 0
	}
	return target, "", 0
}

// bugAsImprint reshapes a surviving bug per the settled kept-bug rule.
func (p *Plan) bugAsImprint(c *record.Record, bug *Source, surv map[string]string, pos map[string]int) {
	c.Type = record.TypeImprint
	ext := ref(bug.Rec.Extends)
	if ext == "" {
		ext = ref(bug.Rec.Supersedes)
	}
	k := keptBugAsImprint(p.Repo.Records, &Source{Rec: &record.Record{Extends: ext}})
	if k.Implements == "" && ext != "" {
		if t, ok := p.Repo.Shared[ext]; ok {
			k.Implements = ext
			if t.Type() == record.TypeImprint {
				k.Implements = ref(t.Rec.Implements)
			}
		}
	}
	c.Implements, c.ImplementsRepo, c.ImplementsVersion = p.pin(k.Implements, surv, pos)
	for _, r := range k.RelatedAdd {
		if s, ok := surv[r]; ok {
			r = s
		}
		c.Related = appendNew(c.Related, r)
	}
}

// v2Status maps a v1 status. superseded does not exist in v2; a record left
// superseded after planning (its successor kept separate or not local) keeps
// the lifecycle it reached: implemented when sealed, else deprecated.
func v2Status(r *record.Record) record.Status {
	if r.Status != record.StatusSuperseded {
		return r.Status
	}
	if r.SealedAtSHA != "" {
		return record.StatusImplemented
	}
	return record.StatusDeprecated
}

func createdAt(m *Source) string {
	if m.Added != nil {
		return time.Unix(m.Added.Unix, 0).UTC().Format(time.RFC3339)
	}
	if t, err := time.Parse("2006-01-02", m.Rec.CreatedAt); err == nil {
		return t.UTC().Format(time.RFC3339)
	}
	return ""
}

func cloneOf(r *record.Record) *record.Record {
	data, err := record.Marshal(r)
	if err != nil {
		panic(err)
	}
	c, err := record.Decode(data)
	if err != nil {
		panic(err)
	}
	c.Schema = r.Schema
	c.Version = r.Version
	return c
}

func appendNew(list []string, items ...string) []string {
	seen := map[string]bool{}
	for _, x := range list {
		seen[x] = true
	}
	out := append([]string(nil), list...)
	for _, x := range items {
		if x != "" && !seen[x] {
			seen[x] = true
			out = append(out, x)
		}
	}
	return out
}

// appendNewNormalized appends constraints whose normalized text is new.
func appendNewNormalized(list, items []string) []string {
	seen := map[string]bool{}
	for _, x := range list {
		seen[record.Normalize(x)] = true
	}
	out := append([]string(nil), list...)
	for _, x := range items {
		if n := record.Normalize(x); !seen[n] {
			seen[n] = true
			out = append(out, x)
		}
	}
	return out
}

func appendSpecs(list, items []record.AssociatedSpec) []record.AssociatedSpec {
	seen := map[string]bool{}
	for _, s := range list {
		seen[s.Path] = true
	}
	out := append([]record.AssociatedSpec(nil), list...)
	for _, s := range items {
		if !seen[s.Path] {
			seen[s.Path] = true
			out = append(out, s)
		}
	}
	return out
}

func without(list []string, x string) []string {
	var out []string
	for _, y := range list {
		if y != x {
			out = append(out, y)
		}
	}
	return out
}
