package record

import (
	"bytes"
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

// Snapshot is one immutable version of a record. The files backend keeps it
// at provenance/.versions/<id>/<version>-<hex>.yml, where hex is ContentHash
// without its "sha256:" prefix.
type Snapshot struct {
	ID          string
	Version     int
	ContentHash string
	// ParentHash is the previous version's ContentHash; empty for version 1.
	ParentHash string
	Reason     string
	ReasonKind string
	Author     string
	// CreatedAt is an RFC 3339 UTC timestamp.
	CreatedAt string
	// SpecHashes holds each associated spec's hash at this version, nil when
	// unresolved, so ContentHash can be recomputed without the spec files.
	SpecHashes []*string
	// Content is the record at this version without its lifecycle fields
	// (status, sealed_at_sha), without version and reason, which the envelope
	// carries, and without any spec run_command, which stays local.
	Content *Record
}

// NewSnapshot builds the snapshot of r as the given version. createdAt is
// an RFC 3339 UTC timestamp. Specs are resolved now, so the snapshot can
// verify its content_hash later without the files.
func NewSnapshot(r *Record, version int, parentHash, reason, reasonKind, author, createdAt string, specs SpecResolver) *Snapshot {
	specHashes, _ := r.SpecHashes(specs)
	return &Snapshot{
		ID: r.ID, Version: version, ContentHash: ContentHashWithSpecs(r, specHashes),
		ParentHash: parentHash, Reason: reason, ReasonKind: reasonKind,
		Author: author, CreatedAt: createdAt,
		SpecHashes: specHashes, Content: snapshotContent(r),
	}
}

// snapshotContent strips what the snapshot envelope or the lifecycle owns.
func snapshotContent(r *Record) *Record {
	c := cloneRecord(r)
	c.Status, c.SealedAtSHA = "", ""
	c.Version, c.Reason, c.ReasonKind = 0, "", ""
	c.Schema = 2
	c.Supersedes, c.SupersededBy, c.Extends = "", "", ""
	// run_command is local-only and never published (DESIGN.md, Associated
	// specs); snapshots are what publishing copies.
	for i := range c.AssociatedSpecs {
		c.AssociatedSpecs[i].RunCommand = ""
	}
	return c
}

// Verify recomputes the content hash from Content and SpecHashes.
func (s *Snapshot) Verify() error {
	if s.Content == nil {
		return fmt.Errorf("snapshot %s@%d has no content", s.ID, s.Version)
	}
	if len(s.SpecHashes) != len(s.Content.AssociatedSpecs) {
		return fmt.Errorf("snapshot %s@%d: %d spec hashes for %d specs", s.ID, s.Version, len(s.SpecHashes), len(s.Content.AssociatedSpecs))
	}
	if got := ContentHashWithSpecs(s.Content, s.SpecHashes); got != s.ContentHash {
		return fmt.Errorf("snapshot %s@%d: content hashes to %s, envelope says %s", s.ID, s.Version, got, s.ContentHash)
	}
	return nil
}

// FileName is the snapshot's name inside provenance/.versions/<id>/.
func (s *Snapshot) FileName() string {
	return fmt.Sprintf("%d-%s.yml", s.Version, strings.TrimPrefix(s.ContentHash, hashPrefix))
}

// MarshalSnapshot encodes a snapshot. The encoding is deterministic, so the
// same snapshot always produces the same bytes.
func MarshalSnapshot(s *Snapshot) ([]byte, error) {
	node := &yaml.Node{Kind: yaml.MappingNode}
	add := func(key string, v any) {
		k, val := keyValueNodes(key, v)
		node.Content = append(node.Content, k, val)
	}
	add("id", s.ID)
	add("version", s.Version)
	add("content_hash", s.ContentHash)
	add("parent_hash", s.ParentHash)
	add("reason", s.Reason)
	if s.ReasonKind != "" {
		add("reason_kind", s.ReasonKind)
	}
	add("author", s.Author)
	add("created_at", s.CreatedAt)
	specs := &yaml.Node{Kind: yaml.SequenceNode}
	if len(s.SpecHashes) == 0 {
		specs.Style = yaml.FlowStyle
	}
	for _, h := range s.SpecHashes {
		if h == nil {
			specs.Content = append(specs.Content, &yaml.Node{Kind: yaml.ScalarNode, Tag: "!!null", Value: "null"})
		} else {
			specs.Content = append(specs.Content, strNode(*h))
		}
	}
	node.Content = append(node.Content, &yaml.Node{Kind: yaml.ScalarNode, Value: "spec_hashes"}, specs)
	content := recordNode(s.Content)
	// The content mapping drops keys that are empty, so a snapshot shows what
	// the version said rather than the schema's every field.
	trimmed := &yaml.Node{Kind: yaml.MappingNode}
	for i := 0; i+1 < len(content.Content); i += 2 {
		if isZero(fieldValue(s.Content, content.Content[i].Value)) {
			continue
		}
		trimmed.Content = append(trimmed.Content, content.Content[i], content.Content[i+1])
	}
	node.Content = append(node.Content, &yaml.Node{Kind: yaml.ScalarNode, Value: "content"}, trimmed)
	return encodeNode(node)
}

// UnmarshalSnapshot decodes a snapshot and verifies its content hash.
func UnmarshalSnapshot(data []byte) (*Snapshot, error) {
	root, err := parseRoot(data)
	if err != nil {
		return nil, err
	}
	s := &Snapshot{}
	for i := 0; i+1 < len(root.Content); i += 2 {
		key, val := root.Content[i].Value, root.Content[i+1]
		var dst *string
		switch key {
		case "id":
			dst = &s.ID
		case "content_hash":
			dst = &s.ContentHash
		case "parent_hash":
			dst = &s.ParentHash
		case "reason":
			dst = &s.Reason
		case "reason_kind":
			dst = &s.ReasonKind
		case "author":
			dst = &s.Author
		case "created_at":
			dst = &s.CreatedAt
		case "version":
			if s.Version, err = decodeInt(key, val); err != nil {
				return nil, err
			}
		case "spec_hashes":
			v, err := decodeList(key, val)
			if err != nil {
				return nil, err
			}
			s.SpecHashes = v.List
		case "content":
			val = deref(val)
			if val.Kind != yaml.MappingNode {
				return nil, fmt.Errorf("snapshot content must be a mapping")
			}
			if s.Content, err = recordFromRoot(val); err != nil {
				return nil, err
			}
			s.Content.Schema, s.Content.Version = 2, 0
		}
		if dst != nil {
			v, err := decodeString(key, val)
			if err != nil {
				return nil, err
			}
			*dst = v.Str
		}
	}
	if s.Content == nil {
		s.Content = &Record{Schema: 2}
	}
	if s.SpecHashes == nil {
		s.SpecHashes = []*string{}
	}
	if err := s.Verify(); err != nil {
		return nil, err
	}
	return s, nil
}

// sameBytes is used to make a repeated Append of an identical snapshot a
// no-op.
func sameBytes(a, b []byte) bool { return bytes.Equal(a, b) }
