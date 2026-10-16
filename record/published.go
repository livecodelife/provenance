package record

import (
	"encoding/json"
	"fmt"
)

// The published form of a record (docs/LEXICONS.md): the JSON shapes of the
// lexicons dev.provisio.provenance.version and .record. Field names are the
// record's snake_case names in camelCase. This file is the one mapping between
// a Snapshot and those shapes, shared by publish, import and the record
// server, so none of them re-derives it.

// StrongRef is AT Proto's {uri, cid} pair.
type StrongRef struct {
	URI string `json:"uri"`
	CID string `json:"cid"`
}

// PublishedSpec is an associated spec as published; the run command is local.
type PublishedSpec struct {
	Path string `json:"path"`
	Type string `json:"type,omitempty"`
}

// PublishedContent is dev.provisio.provenance.defs#content. Absent and empty
// are the same value, so empty lists and strings are omitted, except the two
// required fields.
type PublishedContent struct {
	Title             string          `json:"title"`
	Type              string          `json:"type,omitempty"`
	CreatedAt         string          `json:"createdAt,omitempty"`
	Author            string          `json:"author,omitempty"`
	Intent            string          `json:"intent"`
	Constraints       []string        `json:"constraints,omitempty"`
	AffectedScope     []string        `json:"affectedScope,omitempty"`
	ForbiddenScope    []string        `json:"forbiddenScope,omitempty"`
	Related           []string        `json:"related,omitempty"`
	Tags              []string        `json:"tags,omitempty"`
	Aliases           []string        `json:"aliases,omitempty"`
	Implements        string          `json:"implements,omitempty"`
	ImplementsRepo    string          `json:"implementsRepo,omitempty"`
	ImplementsVersion int             `json:"implementsVersion,omitempty"`
	AssociatedSpecs   []PublishedSpec `json:"associatedSpecs,omitempty"`
	AssociatedTraces  []string        `json:"associatedTraces,omitempty"`
	Monitors          []string        `json:"monitors,omitempty"`
	Locked            bool            `json:"locked,omitempty"`
}

// PublishedVersion is dev.provisio.provenance.version. SpecHashes uses ""
// where the snapshot has none (unresolved), because a lexicon has no null.
type PublishedVersion struct {
	Record          string           `json:"record"`
	RecordID        string           `json:"recordId"`
	Version         int              `json:"version"`
	ContentHash     string           `json:"contentHash"`
	ParentHash      string           `json:"parentHash,omitempty"`
	Parent          *StrongRef       `json:"parent,omitempty"`
	Reason          string           `json:"reason,omitempty"`
	ReasonKind      string           `json:"reasonKind,omitempty"`
	Author          string           `json:"author"`
	CreatedAt       string           `json:"createdAt"`
	Commit          string           `json:"commit,omitempty"`
	SpecHashes      []string         `json:"specHashes"`
	Content         PublishedContent `json:"content"`
	Origin          string           `json:"origin,omitempty"`
	BaseVersion     int              `json:"baseVersion,omitempty"`
	BaseContentHash string           `json:"baseContentHash,omitempty"`
}

// PublishedSource is where the record lives in git.
type PublishedSource struct {
	Repo   string `json:"repo,omitempty"`
	Path   string `json:"path,omitempty"`
	Commit string `json:"commit,omitempty"`
}

// PublishedHead is dev.provisio.provenance.record.
type PublishedHead struct {
	Status      string           `json:"status"`
	Version     int              `json:"version"`
	ContentHash string           `json:"contentHash"`
	GenHash     string           `json:"genHash"`
	Latest      StrongRef        `json:"latest"`
	Content     PublishedContent `json:"content"`
	Source      *PublishedSource `json:"source,omitempty"`
	Project     string           `json:"project,omitempty"`
	UpdatedAt   string           `json:"updatedAt"`
}

// PublishedContentOf maps a record's content fields; lifecycle and envelope
// fields (status, version, reason, sealed commit, v1 graph fields) are left out.
func PublishedContentOf(r *Record) PublishedContent {
	c := PublishedContent{
		Title: r.Title, Type: string(r.Type), CreatedAt: r.CreatedAt, Author: r.Author, Intent: r.Intent,
		Constraints: r.Constraints, AffectedScope: r.AffectedScope, ForbiddenScope: r.ForbiddenScope,
		Related: r.Related, Tags: r.Tags, Aliases: r.Aliases, Implements: r.Implements,
		ImplementsRepo: r.ImplementsRepo, ImplementsVersion: r.ImplementsVersion,
		AssociatedTraces: r.AssociatedTraces, Monitors: r.Monitors, Locked: r.Locked,
	}
	for _, s := range r.AssociatedSpecs {
		c.AssociatedSpecs = append(c.AssociatedSpecs, PublishedSpec{Path: s.Path, Type: s.Type})
	}
	return c
}

// Record is the inverse of PublishedContentOf: a schema 2 record carrying
// only content fields.
func (c PublishedContent) Record() *Record {
	r := &Record{
		Schema: 2, Title: c.Title, Type: Type(c.Type), CreatedAt: c.CreatedAt, Author: c.Author, Intent: c.Intent,
		Constraints: c.Constraints, AffectedScope: c.AffectedScope, ForbiddenScope: c.ForbiddenScope,
		Related: c.Related, Tags: c.Tags, Aliases: c.Aliases, Implements: c.Implements,
		ImplementsRepo: c.ImplementsRepo, ImplementsVersion: c.ImplementsVersion,
		AssociatedTraces: c.AssociatedTraces, Monitors: c.Monitors, Locked: c.Locked,
	}
	for _, s := range c.AssociatedSpecs {
		r.AssociatedSpecs = append(r.AssociatedSpecs, AssociatedSpec{Path: s.Path, Type: s.Type})
	}
	return r
}

// PublishContext is what a snapshot does not know: where it is published.
type PublishContext struct {
	// RecordURI is the at-uri of the head record.
	RecordURI string
	// Parent is the previous version record; nil for version 1.
	Parent *StrongRef
	// Commit is the git commit the version was published from.
	Commit string
}

// PublishedVersionOf maps a snapshot to its published form.
func PublishedVersionOf(s *Snapshot, ctx PublishContext) *PublishedVersion {
	v := &PublishedVersion{
		Record: ctx.RecordURI, RecordID: s.ID, Version: s.Version, ContentHash: s.ContentHash,
		ParentHash: s.ParentHash, Parent: ctx.Parent, Reason: s.Reason, ReasonKind: s.ReasonKind,
		Author: s.Author, CreatedAt: s.CreatedAt, Commit: ctx.Commit, SpecHashes: []string{},
	}
	for _, h := range s.SpecHashes {
		if h == nil {
			v.SpecHashes = append(v.SpecHashes, "")
		} else {
			v.SpecHashes = append(v.SpecHashes, *h)
		}
	}
	if s.Content != nil {
		v.Content = PublishedContentOf(s.Content)
	}
	return v
}

// Snapshot is the inverse of PublishedVersionOf and checks the result: the
// content hash must be the one recomputed from content and spec hashes.
func (v *PublishedVersion) Snapshot() (*Snapshot, error) {
	if len(v.SpecHashes) != len(v.Content.AssociatedSpecs) {
		return nil, fmt.Errorf("version %s@%d: %d spec hashes for %d specs", v.RecordID, v.Version, len(v.SpecHashes), len(v.Content.AssociatedSpecs))
	}
	s := &Snapshot{
		ID: v.RecordID, Version: v.Version, ContentHash: v.ContentHash, ParentHash: v.ParentHash,
		Reason: v.Reason, ReasonKind: v.ReasonKind, Author: v.Author, CreatedAt: v.CreatedAt,
		SpecHashes: make([]*string, len(v.SpecHashes)), Content: v.Content.Record(),
	}
	s.Content.ID = v.RecordID
	for i := range v.SpecHashes {
		if v.SpecHashes[i] != "" {
			h := v.SpecHashes[i]
			s.SpecHashes[i] = &h
		}
	}
	if err := s.Verify(); err != nil {
		return nil, err
	}
	return s, nil
}

// ParsePublishedVersion decodes a version record's JSON and verifies it.
func ParsePublishedVersion(data []byte) (*PublishedVersion, *Snapshot, error) {
	var v PublishedVersion
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, nil, err
	}
	s, err := v.Snapshot()
	if err != nil {
		return nil, nil, err
	}
	return &v, s, nil
}

// PublishedHeadOf maps a record, as it is now, to its head. latest is the
// newest version record; updatedAt an RFC 3339 UTC time.
func PublishedHeadOf(r *Record, specs SpecResolver, latest StrongRef, source *PublishedSource, project, updatedAt string) PublishedHead {
	hash, _ := r.ContentHash(specs)
	return PublishedHead{
		Status: string(r.Status), Version: r.Version, ContentHash: hash, GenHash: r.GenHash(),
		Latest: latest, Content: PublishedContentOf(r), Source: source, Project: project, UpdatedAt: updatedAt,
	}
}
