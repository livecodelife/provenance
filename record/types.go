// Package record is the one definition of a provenance record: the v2 schema,
// a decoder that also reads v1 records, a formatting-preserving YAML codec,
// the canonical gen_hash and content_hash (docs/HASHING.md), version
// snapshots, and the record-store interface with its files backend.
package record

import (
	"fmt"
	"regexp"

	"gopkg.in/yaml.v3"
)

// Status is the lifecycle state of a record's current version.
type Status string

const (
	StatusDraft       Status = "draft"
	StatusOpen        Status = "open"
	StatusImplemented Status = "implemented"
	// StatusDeprecated is terminal: a deprecated record takes no new versions.
	StatusDeprecated Status = "deprecated"
	// StatusSuperseded exists only in v1 records. Versioning replaces it.
	StatusSuperseded Status = "superseded"
)

// Valid reports whether s is a legal status for a record of the given schema.
func (s Status) Valid(schema int) bool {
	switch s {
	case StatusDraft, StatusOpen, StatusImplemented, StatusDeprecated:
		return true
	case StatusSuperseded:
		return schema == 1
	}
	return false
}

// Type is a record's tier.
type Type string

const (
	TypeBrief     Type = "brief"
	TypeBlueprint Type = "blueprint"
	TypeImprint   Type = "imprint"
	// TypeBug exists only in v1 records. Every bug becomes a version.
	TypeBug Type = "bug"
)

// Valid reports whether t is a legal type for a record of the given schema.
// The empty type means blueprint, as in v1.
func (t Type) Valid(schema int) bool {
	switch t {
	case "", TypeBrief, TypeBlueprint, TypeImprint:
		return true
	case TypeBug:
		return schema == 1
	}
	return false
}

// ReasonKindFix marks a version that fixes a defect, which keeps fixes
// visible once the bug type retires.
const ReasonKindFix = "fix"

// AssociatedSpec is a proof artifact attached to a record.
type AssociatedSpec struct {
	Path       string
	Type       string
	RunCommand string
	// Extra holds keys the schema does not model, in file order. They are
	// written back unchanged.
	Extra []KeyNode
}

// KeyNode is one unmodelled mapping entry, kept as parsed.
type KeyNode struct {
	Key   string
	Value *yaml.Node
}

// Record is a provenance record. Schema is 2 for a versioned record and 1 for
// a record read through the v1 compatibility path.
type Record struct {
	Schema int

	ID        string
	Title     string
	Type      Type
	Status    Status
	CreatedAt string
	Author    string

	// Version starts at 1 and is bumped whenever content_hash changes. A v1
	// record reads as version 1.
	Version int
	// Reason is why the current version exists. Required on every bump.
	Reason     string
	ReasonKind string

	// Aliases are old ids folded into this record during migration, so old
	// commit tags still resolve.
	Aliases []string

	Intent      string
	Constraints []string

	AffectedScope  []string
	ForbiddenScope []string

	// Implements is the parent record's id.
	Implements string
	// ImplementsRepo names the repository holding the parent when it is not
	// local (a v1 shared_repos target). Empty means the parent is local.
	ImplementsRepo string
	// ImplementsURI is the at-uri of the parent's head record
	// (at://<did>/dev.provisio.provenance.record/<id>) when the parent is
	// published in another account. Its record id is Implements. Lineage only:
	// it is not part of the content hash.
	ImplementsURI string
	// ImplementsVersion is the parent version this record was last confirmed
	// against. Zero means unpinned.
	ImplementsVersion int
	Related           []string

	SealedAtSHA      string
	AssociatedSpecs  []AssociatedSpec
	AssociatedTraces []string
	Monitors         []string
	Tags             []string
	Locked           bool

	// V1-only graph fields, kept so migration can read them. Never hashed and
	// never written for a schema 2 record.
	Supersedes   string
	SupersededBy string
	Extends      string
}

// EffectiveType returns the record's type, with the v1 default applied.
func (r *Record) EffectiveType() Type {
	if r.Type == "" {
		return TypeBlueprint
	}
	return r.Type
}

// IDPattern matches valid record ids: prov-YYYY-NNN or prov-YYYY-XXXXXXXX,
// with an optional service suffix.
var IDPattern = regexp.MustCompile(`^prov-(\d{4})-(\d{3}|[a-f0-9]{8})(?:-[a-z0-9-]+)?$`)

// Validate checks schema-level invariants that do not need other records.
func (r *Record) Validate() error {
	if !IDPattern.MatchString(r.ID) {
		return fmt.Errorf("invalid id %q", r.ID)
	}
	if !r.Status.Valid(r.Schema) {
		return fmt.Errorf("%s: invalid status %q for schema %d", r.ID, r.Status, r.Schema)
	}
	if !r.Type.Valid(r.Schema) {
		return fmt.Errorf("%s: invalid type %q for schema %d", r.ID, r.Type, r.Schema)
	}
	if r.Schema == 2 {
		if r.Version < 1 {
			return fmt.Errorf("%s: version must be at least 1", r.ID)
		}
		if r.Version > 1 && r.Reason == "" {
			return fmt.Errorf("%s: version %d has no reason", r.ID, r.Version)
		}
		if r.ReasonKind != "" && r.ReasonKind != ReasonKindFix {
			return fmt.Errorf("%s: invalid reason_kind %q", r.ID, r.ReasonKind)
		}
		if (r.ImplementsVersion != 0 || r.ImplementsRepo != "" || r.ImplementsURI != "") && r.Implements == "" {
			return fmt.Errorf("%s: implements_version, implements_repo or implements_uri without implements", r.ID)
		}
		if r.Supersedes != "" || r.SupersededBy != "" || r.Extends != "" {
			return fmt.Errorf("%s: supersedes, superseded_by and extends are v1-only", r.ID)
		}
	}
	return nil
}
