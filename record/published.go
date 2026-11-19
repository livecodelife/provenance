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

// PublishedShapeParam, PublishedShapeDecl, PublishedShapeStub and
// PublishedShapeFile are dev.provisio.provenance.defs#shapeParam, #shapeDecl,
// #shapeStub and #shapeFile: a record's shape as published. Their JSON keys are
// the skeleton's own (snake_case), not camelCase, because the shape is hashed
// as written: renaming a key would change gen_hash.
type PublishedShapeParam struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type PublishedShapeDecl struct {
	Kind     string `json:"kind,omitempty"`
	Name     string `json:"name,omitempty"`
	Receiver string `json:"receiver,omitempty"`
	Code     string `json:"code,omitempty"`
}

type PublishedShapeStub struct {
	Name         string                `json:"name,omitempty"`
	ReceiverName string                `json:"receiver_name,omitempty"`
	ReceiverType string                `json:"receiver_type,omitempty"`
	Params       []PublishedShapeParam `json:"params,omitempty"`
	Results      []string              `json:"results,omitempty"`
	Doc          string                `json:"doc,omitempty"`
}

type PublishedShapeFile struct {
	Path    string               `json:"path,omitempty"`
	Package string               `json:"package,omitempty"`
	Imports []string             `json:"imports,omitempty"`
	Decls   []PublishedShapeDecl `json:"decls,omitempty"`
	Stubs   []PublishedShapeStub `json:"stubs,omitempty"`
}

// PublishedContent is dev.provisio.provenance.defs#content. Absent and empty
// are the same value, so empty lists and strings are omitted, except the two
// required fields.
type PublishedContent struct {
	Title          string   `json:"title"`
	Type           string   `json:"type,omitempty"`
	CreatedAt      string   `json:"createdAt,omitempty"`
	Author         string   `json:"author,omitempty"`
	Intent         string   `json:"intent"`
	Constraints    []string `json:"constraints,omitempty"`
	AffectedScope  []string `json:"affectedScope,omitempty"`
	ForbiddenScope []string `json:"forbiddenScope,omitempty"`
	// Shape is optional: absent or empty means no shape (and the four-key gen_hash).
	Shape             []PublishedShapeFile `json:"shape,omitempty"`
	Related           []string             `json:"related,omitempty"`
	Tags              []string             `json:"tags,omitempty"`
	Aliases           []string             `json:"aliases,omitempty"`
	Implements        string               `json:"implements,omitempty"`
	ImplementsRepo    string               `json:"implementsRepo,omitempty"`
	ImplementsURI     string               `json:"implementsUri,omitempty"`
	ImplementsVersion int                  `json:"implementsVersion,omitempty"`
	AssociatedSpecs   []PublishedSpec      `json:"associatedSpecs,omitempty"`
	AssociatedTraces  []string             `json:"associatedTraces,omitempty"`
	Monitors          []string             `json:"monitors,omitempty"`
	Locked            bool                 `json:"locked,omitempty"`
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
		Shape:   publishedShape(r.Shape),
		Related: r.Related, Tags: r.Tags, Aliases: r.Aliases, Implements: r.Implements,
		ImplementsRepo: r.ImplementsRepo, ImplementsURI: r.ImplementsURI, ImplementsVersion: r.ImplementsVersion,
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
		Shape:   recordShape(c.Shape),
		Related: c.Related, Tags: c.Tags, Aliases: c.Aliases, Implements: c.Implements,
		ImplementsRepo: c.ImplementsRepo, ImplementsURI: c.ImplementsURI, ImplementsVersion: c.ImplementsVersion,
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

// publishedShape maps a record's shape tree to the typed published form.
// Empty entries are dropped (they hash as absent). Keys the lexicon does not
// name, and empty values, are not carried: CheckPublishable reports them.
func publishedShape(shape []any) []PublishedShapeFile {
	var out []PublishedShapeFile
	for _, e := range PruneShape(shape) {
		b, err := json.Marshal(e)
		if err != nil {
			continue
		}
		var f PublishedShapeFile
		if json.Unmarshal(b, &f) != nil {
			continue
		}
		out = append(out, f)
	}
	return out
}

// recordShape is the inverse: the typed entries as the generic tree a record
// holds (strings, lists and maps, with the same keys).
func recordShape(files []PublishedShapeFile) []any {
	var out []any
	for _, f := range files {
		b, err := json.Marshal(f)
		if err != nil {
			continue
		}
		var v map[string]any
		if json.Unmarshal(b, &v) != nil || len(v) == 0 {
			continue
		}
		out = append(out, jsonTree(v))
	}
	return out
}

// jsonTree converts decoded JSON to the shape tree types ([]any, map[string]any, string).
func jsonTree(v any) any {
	switch v := v.(type) {
	case map[string]any:
		out := make(map[string]any, len(v))
		for k, x := range v {
			out[k] = jsonTree(x)
		}
		return out
	case []any:
		out := make([]any, len(v))
		for i, x := range v {
			out[i] = jsonTree(x)
		}
		return out
	}
	return v
}

// CheckPublishable reports whether r's shape survives publishing unchanged:
// the typed lexicon carries path, package, imports, decls (kind, name,
// receiver, code) and stubs (name, receiver_name, receiver_type, params,
// results, doc), and omits empty values. A shape with other keys, or with an
// empty value that is not absent, would come back with another gen_hash, so
// the server would refuse it; this refuses it first, with the reason, rather
// than dropping anything silently.
func CheckPublishable(r *Record) error {
	if len(PruneShape(r.Shape)) == 0 {
		return nil
	}
	if err := ValidateShape(r.Shape); err != nil {
		return err
	}
	back := PublishedContentOf(r).Record()
	if back.GenHash() != r.GenHash() {
		return fmt.Errorf("%s: the shape has keys or empty values the published lexicon cannot carry (it carries path, package, imports, decls, stubs with their documented keys, and omits empty values), so its gen_hash would change", r.ID)
	}
	return nil
}
