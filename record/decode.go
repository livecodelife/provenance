package record

import (
	"fmt"
	"strconv"
	"unicode/utf8"

	"gopkg.in/yaml.v3"
)

// Value is one decoded field in the hash input model (docs/HASHING.md §1): a
// string, a list of strings with null items allowed, or null.
type Value struct {
	Null bool
	Str  string
	// List is set (possibly empty) when the field is a sequence. A nil item
	// is a null list item.
	List []*string
	// IsList distinguishes an empty sequence from a string.
	IsList bool
}

// fieldKind is how a known top-level key decodes.
type fieldKind int

const (
	kindString fieldKind = iota
	kindList
	kindInt
	kindBool
	kindSpecs
)

// knownFields maps every key the schema models, v1 and v2, to its kind.
// Any other top-level key is unknown: preserved, never hashed.
var knownFields = map[string]fieldKind{
	"id": kindString, "title": kindString, "type": kindString, "status": kindString,
	"created_at": kindString, "author": kindString,
	"version": kindInt, "reason": kindString, "reason_kind": kindString,
	"aliases": kindList,
	"intent":  kindString, "constraints": kindList,
	"affected_scope": kindList, "forbidden_scope": kindList,
	"implements": kindString, "implements_repo": kindString, "implements_version": kindInt,
	"related": kindList, "sealed_at_sha": kindString,
	"associated_specs": kindSpecs, "associated_traces": kindList,
	"monitors": kindList, "tags": kindList, "locked": kindBool,
	"supersedes": kindString, "superseded_by": kindString, "extends": kindString,
}

// Decode parses a record file. A file with a top-level version key is schema
// 2; any other file is read through the v1 compatibility path as version 1.
func Decode(data []byte) (*Record, error) {
	doc, err := Parse(data)
	if err != nil {
		return nil, err
	}
	return doc.Record(), nil
}

// DecodeFields decodes the known top-level keys of a record file into the
// hash input model, keeping null as null. The conformance corpus's decode
// cases check this layer.
func DecodeFields(data []byte) (map[string]Value, error) {
	root, err := parseRoot(data)
	if err != nil {
		return nil, err
	}
	out := map[string]Value{}
	for i := 0; i+1 < len(root.Content); i += 2 {
		key, val := root.Content[i].Value, root.Content[i+1]
		kind, known := knownFields[key]
		if !known {
			continue
		}
		switch kind {
		case kindString:
			v, err := decodeString(key, val)
			if err != nil {
				return nil, err
			}
			out[key] = v
		case kindList:
			v, err := decodeList(key, val)
			if err != nil {
				return nil, err
			}
			out[key] = v
		}
	}
	return out, nil
}

func parseRoot(data []byte) (*yaml.Node, error) {
	if !utf8.Valid(data) {
		return nil, fmt.Errorf("record is not valid UTF-8")
	}
	var doc yaml.Node
	if err := yaml.Unmarshal(data, &doc); err != nil {
		return nil, fmt.Errorf("invalid YAML: %w", err)
	}
	if doc.Kind != yaml.DocumentNode || len(doc.Content) != 1 || doc.Content[0].Kind != yaml.MappingNode {
		return nil, fmt.Errorf("record must be a YAML mapping")
	}
	root := doc.Content[0]
	seen := map[string]bool{}
	for i := 0; i+1 < len(root.Content); i += 2 {
		k := root.Content[i]
		if k.Kind != yaml.ScalarNode {
			return nil, fmt.Errorf("line %d: top-level key must be a scalar", k.Line)
		}
		if seen[k.Value] {
			return nil, fmt.Errorf("line %d: duplicate key %q", k.Line, k.Value)
		}
		seen[k.Value] = true
	}
	return root, nil
}

func deref(n *yaml.Node) *yaml.Node {
	for n != nil && n.Kind == yaml.AliasNode {
		n = n.Alias
	}
	return n
}

func isNull(n *yaml.Node) bool {
	return n.Kind == yaml.ScalarNode && n.ShortTag() == "!!null"
}

// decodeString takes a scalar's source text (after quoting and block-style
// resolution), never a typed value: plain 1.50 stays "1.50".
func decodeString(key string, n *yaml.Node) (Value, error) {
	n = deref(n)
	if n.Kind != yaml.ScalarNode {
		return Value{}, fmt.Errorf("line %d: %s must be a string", n.Line, key)
	}
	if isNull(n) {
		return Value{Null: true}, nil
	}
	return Value{Str: n.Value}, nil
}

func decodeList(key string, n *yaml.Node) (Value, error) {
	n = deref(n)
	if isNull(n) {
		return Value{Null: true}, nil
	}
	if n.Kind != yaml.SequenceNode {
		return Value{}, fmt.Errorf("line %d: %s must be a list", n.Line, key)
	}
	v := Value{IsList: true, List: make([]*string, 0, len(n.Content))}
	for _, item := range n.Content {
		s, err := decodeString(key+" item", item)
		if err != nil {
			return Value{}, err
		}
		if s.Null {
			v.List = append(v.List, nil)
		} else {
			str := s.Str
			v.List = append(v.List, &str)
		}
	}
	return v, nil
}

func (v Value) strings() []string {
	if !v.IsList {
		return nil
	}
	out := make([]string, len(v.List))
	for i, s := range v.List {
		if s != nil {
			out[i] = *s
		}
	}
	return out
}

func decodeInt(key string, n *yaml.Node) (int, error) {
	n = deref(n)
	if isNull(n) {
		return 0, nil
	}
	if n.Kind != yaml.ScalarNode {
		return 0, fmt.Errorf("line %d: %s must be an integer", n.Line, key)
	}
	i, err := strconv.Atoi(n.Value)
	if err != nil {
		return 0, fmt.Errorf("line %d: %s must be an integer: %q", n.Line, key, n.Value)
	}
	return i, nil
}

func decodeBool(key string, n *yaml.Node) (bool, error) {
	n = deref(n)
	if isNull(n) {
		return false, nil
	}
	var b bool
	if n.Kind != yaml.ScalarNode || n.Decode(&b) != nil {
		return false, fmt.Errorf("line %d: %s must be a boolean", n.Line, key)
	}
	return b, nil
}

func decodeSpecs(n *yaml.Node) ([]AssociatedSpec, error) {
	n = deref(n)
	if isNull(n) {
		return nil, nil
	}
	if n.Kind != yaml.SequenceNode {
		return nil, fmt.Errorf("line %d: associated_specs must be a list", n.Line)
	}
	specs := make([]AssociatedSpec, 0, len(n.Content))
	for _, item := range n.Content {
		item = deref(item)
		if item.Kind == yaml.ScalarNode && !isNull(item) {
			// A bare path is accepted for leniency; it saves back as a mapping.
			specs = append(specs, AssociatedSpec{Path: item.Value})
			continue
		}
		if item.Kind != yaml.MappingNode {
			return nil, fmt.Errorf("line %d: associated_specs item must be a mapping", item.Line)
		}
		var s AssociatedSpec
		for i := 0; i+1 < len(item.Content); i += 2 {
			k, v := item.Content[i].Value, item.Content[i+1]
			var dst *string
			switch k {
			case "path":
				dst = &s.Path
			case "type":
				dst = &s.Type
			case "run_command":
				dst = &s.RunCommand
			default:
				s.Extra = append(s.Extra, KeyNode{Key: k, Value: v})
				continue
			}
			str, err := decodeString("associated_specs."+k, v)
			if err != nil {
				return nil, err
			}
			*dst = str.Str
		}
		specs = append(specs, s)
	}
	return specs, nil
}

// recordFromRoot builds a Record from a validated root mapping.
func recordFromRoot(root *yaml.Node) (*Record, error) {
	r := &Record{}
	hasVersion := false
	for i := 0; i+1 < len(root.Content); i += 2 {
		key, val := root.Content[i].Value, root.Content[i+1]
		kind, known := knownFields[key]
		if !known {
			continue
		}
		var err error
		switch kind {
		case kindString:
			var v Value
			if v, err = decodeString(key, val); err == nil {
				*stringField(r, key) = v.Str
			}
		case kindList:
			var v Value
			if v, err = decodeList(key, val); err == nil {
				*listField(r, key) = v.strings()
			}
		case kindInt:
			var n int
			if n, err = decodeInt(key, val); err == nil {
				if key == "version" {
					hasVersion = true
					r.Version = n
				} else {
					r.ImplementsVersion = n
				}
			}
		case kindBool:
			r.Locked, err = decodeBool(key, val)
		case kindSpecs:
			r.AssociatedSpecs, err = decodeSpecs(val)
		}
		if err != nil {
			return nil, err
		}
	}
	if hasVersion {
		r.Schema = 2
	} else {
		r.Schema = 1
		r.Version = 1
	}
	return r, nil
}

func stringField(r *Record, key string) *string {
	switch key {
	case "id":
		return &r.ID
	case "title":
		return &r.Title
	case "type":
		return (*string)(&r.Type)
	case "status":
		return (*string)(&r.Status)
	case "created_at":
		return &r.CreatedAt
	case "author":
		return &r.Author
	case "reason":
		return &r.Reason
	case "reason_kind":
		return &r.ReasonKind
	case "intent":
		return &r.Intent
	case "implements":
		return &r.Implements
	case "implements_repo":
		return &r.ImplementsRepo
	case "sealed_at_sha":
		return &r.SealedAtSHA
	case "supersedes":
		return &r.Supersedes
	case "superseded_by":
		return &r.SupersededBy
	case "extends":
		return &r.Extends
	}
	panic("record: no string field " + key)
}

func listField(r *Record, key string) *[]string {
	switch key {
	case "aliases":
		return &r.Aliases
	case "constraints":
		return &r.Constraints
	case "affected_scope":
		return &r.AffectedScope
	case "forbidden_scope":
		return &r.ForbiddenScope
	case "related":
		return &r.Related
	case "associated_traces":
		return &r.AssociatedTraces
	case "monitors":
		return &r.Monitors
	case "tags":
		return &r.Tags
	}
	panic("record: no list field " + key)
}
