package record

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

// v2Order is the key order for a new schema 2 file.
var v2Order = []string{
	"id", "title", "type", "status", "version", "reason", "reason_kind",
	"created_at", "author", "aliases", "intent", "constraints",
	"affected_scope", "forbidden_scope", "implements", "implements_repo",
	"implements_version", "related", "sealed_at_sha", "associated_specs",
	"associated_traces", "monitors", "tags", "locked",
}

// v1Order is linespec's marshalled key order, for a new schema 1 file.
var v1Order = []string{
	"id", "title", "status", "created_at", "author", "intent", "constraints",
	"affected_scope", "forbidden_scope", "type", "supersedes", "superseded_by",
	"extends", "implements", "related", "sealed_at_sha", "associated_specs",
	"associated_traces", "monitors", "tags", "locked",
}

// omitWhenZero lists keys a new file leaves out while they hold the zero
// value. Every other key in the order is always written.
var omitWhenZero = map[string]bool{
	"reason_kind": true, "implements": true, "implements_repo": true,
	"implements_version": true, "locked": true, "extends": true,
	"aliases": true, "type": true,
}

// insertAfter places a key that an existing file lacks next to a related one.
// Keys without an entry, or whose anchor is absent, go at the end.
var insertAfter = map[string]string{
	"version": "status", "reason": "version", "reason_kind": "reason",
	"aliases": "author", "implements_repo": "implements",
	"implements_version": "implements_repo",
}

func keysFor(schema int) []string {
	if schema == 1 {
		return v1Order
	}
	return v2Order
}

func validKey(schema int, key string) bool {
	for _, k := range keysFor(schema) {
		if k == key {
			return true
		}
	}
	return false
}

// span is the byte range a top-level key occupies: from the start of its key
// line to the start of the comment and blank lines that precede the next key
// (or the end of the file). Those in-between lines belong to neither key and
// are never rewritten.
type span struct{ start, end int }

// Document is a parsed record file. It remembers the original bytes so that
// saving rewrites only the keys whose values changed.
type Document struct {
	src   []byte
	root  *yaml.Node
	rec   *Record
	spans map[string]span
	order []string // top-level keys in file order
	// block is false when the mapping cannot be spliced (flow style or keys
	// not at column 1); Encode then rewrites the whole file.
	block bool
}

// Parse decodes a record file and indexes its top-level keys.
func Parse(data []byte) (*Document, error) {
	root, err := parseRoot(data)
	if err != nil {
		return nil, err
	}
	rec, err := recordFromRoot(root)
	if err != nil {
		return nil, err
	}
	d := &Document{src: data, root: root, rec: rec, spans: map[string]span{}, block: root.Style&yaml.FlowStyle == 0}
	lineStarts := []int{0}
	for i, c := range data {
		if c == '\n' {
			lineStarts = append(lineStarts, i+1)
		}
	}
	lineOf := func(off int) int { return sort.SearchInts(lineStarts, off+1) - 1 }
	var keyStarts []int
	for i := 0; i < len(root.Content); i += 2 {
		k := root.Content[i]
		if k.Column != 1 {
			d.block = false
		}
		d.order = append(d.order, k.Value)
		if k.Line-1 < len(lineStarts) {
			keyStarts = append(keyStarts, lineStarts[k.Line-1])
		}
	}
	if !d.block || len(keyStarts) != len(d.order) {
		d.block = false
		return d, nil
	}
	for i, key := range d.order {
		end := len(data)
		if i+1 < len(keyStarts) {
			end = keyStarts[i+1]
		}
		// Walk back over blank and column-1 comment lines.
		for line := lineOf(end) - 1; line > lineOf(keyStarts[i]); line-- {
			text := strings.TrimRight(string(data[lineStarts[line]:min(end, lineEnd(data, lineStarts[line]))]), "\r\n")
			if strings.TrimSpace(text) == "" || strings.HasPrefix(text, "#") {
				end = lineStarts[line]
				continue
			}
			break
		}
		d.spans[key] = span{start: keyStarts[i], end: end}
	}
	return d, nil
}

func lineEnd(data []byte, start int) int {
	if i := bytes.IndexByte(data[start:], '\n'); i >= 0 {
		return start + i + 1
	}
	return len(data)
}

// Record returns a copy of the record as decoded.
func (d *Document) Record() *Record {
	return cloneRecord(d.rec)
}

// Encode returns the file bytes for r, rewriting only the top-level keys whose
// values differ from what was parsed. Comments, key order, quoting and
// unchanged or unknown keys stay byte-identical. Saving an unchanged record
// returns the original bytes.
func (d *Document) Encode(r *Record) ([]byte, error) {
	if !d.block {
		return d.encodeWhole(r)
	}
	type edit struct {
		start, end int
		seq        int
		text       string
	}
	var edits []edit
	keys := keysFor(r.Schema)
	// Keys that exist in the file but are invalid for r's schema (v1 fields
	// after migration) are removed.
	for _, key := range d.order {
		if _, known := knownFields[key]; known && !validKey(r.Schema, key) {
			sp := d.spans[key]
			edits = append(edits, edit{start: sp.start, end: sp.end})
		}
	}
	for seq, key := range keys {
		newV := fieldValue(r, key)
		sp, present := d.spans[key]
		if present {
			if equalValue(fieldValue(d.rec, key), newV) {
				continue
			}
			text, err := encodeKey(key, newV)
			if err != nil {
				return nil, err
			}
			if !bytes.HasSuffix(d.src[sp.start:sp.end], []byte("\n")) {
				text = strings.TrimSuffix(text, "\n")
			}
			edits = append(edits, edit{start: sp.start, end: sp.end, seq: seq, text: text})
			continue
		}
		if isZero(newV) {
			continue
		}
		text, err := encodeKey(key, newV)
		if err != nil {
			return nil, err
		}
		at := d.insertionPoint(key)
		if at == len(d.src) && len(d.src) > 0 && d.src[len(d.src)-1] != '\n' {
			text = "\n" + text
		}
		edits = append(edits, edit{start: at, end: at, seq: seq, text: text})
	}
	if len(edits) == 0 {
		return append([]byte(nil), d.src...), nil
	}
	// At one offset, insertions go first: a key inserted after its anchor
	// lands before whatever starts there, including a span being replaced or
	// removed.
	sort.SliceStable(edits, func(i, j int) bool {
		if edits[i].start != edits[j].start {
			return edits[i].start < edits[j].start
		}
		ii, ji := edits[i].start == edits[i].end, edits[j].start == edits[j].end
		if ii != ji {
			return ii
		}
		return edits[i].seq < edits[j].seq
	})
	var out bytes.Buffer
	pos := 0
	for _, e := range edits {
		if e.start < pos {
			return nil, fmt.Errorf("record: overlapping edits at byte %d", e.start)
		}
		out.Write(d.src[pos:e.start])
		out.WriteString(e.text)
		pos = e.end
	}
	out.Write(d.src[pos:])
	return out.Bytes(), nil
}

// insertionPoint returns where a key the file lacks is inserted: after its
// anchor key's span when the file has one, else after the last key's span.
func (d *Document) insertionPoint(key string) int {
	for anchor := insertAfter[key]; anchor != ""; anchor = insertAfter[anchor] {
		if sp, ok := d.spans[anchor]; ok {
			return sp.end
		}
	}
	if len(d.order) == 0 {
		return len(d.src)
	}
	return d.spans[d.order[len(d.order)-1]].end
}

// encodeWhole re-marshals the record for files the splicer cannot handle,
// keeping unknown keys at the end.
func (d *Document) encodeWhole(r *Record) ([]byte, error) {
	node := recordNode(r)
	for i := 0; i+1 < len(d.root.Content); i += 2 {
		if _, known := knownFields[d.root.Content[i].Value]; !known {
			node.Content = append(node.Content, d.root.Content[i], d.root.Content[i+1])
		}
	}
	return encodeNode(node)
}

// Marshal encodes a record as a new file, in canonical key order.
func Marshal(r *Record) ([]byte, error) { return encodeNode(recordNode(r)) }

func recordNode(r *Record) *yaml.Node {
	node := &yaml.Node{Kind: yaml.MappingNode}
	for _, key := range keysFor(r.Schema) {
		v := fieldValue(r, key)
		if omitWhenZero[key] && isZero(v) {
			continue
		}
		k, val := keyValueNodes(key, v)
		node.Content = append(node.Content, k, val)
	}
	return node
}

func encodeKey(key string, v any) (string, error) {
	k, val := keyValueNodes(key, v)
	out, err := encodeNode(&yaml.Node{Kind: yaml.MappingNode, Content: []*yaml.Node{k, val}})
	return string(out), err
}

func encodeNode(n *yaml.Node) ([]byte, error) {
	var b bytes.Buffer
	enc := yaml.NewEncoder(&b)
	enc.SetIndent(2)
	if err := enc.Encode(n); err != nil {
		return nil, err
	}
	if err := enc.Close(); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func keyValueNodes(key string, v any) (*yaml.Node, *yaml.Node) {
	k := &yaml.Node{Kind: yaml.ScalarNode, Value: key}
	var val yaml.Node
	switch v := v.(type) {
	case []string:
		val = yaml.Node{Kind: yaml.SequenceNode}
		if len(v) == 0 {
			val.Style = yaml.FlowStyle
		}
		for _, s := range v {
			val.Content = append(val.Content, strNode(s))
		}
	case []AssociatedSpec:
		val = yaml.Node{Kind: yaml.SequenceNode}
		if len(v) == 0 {
			val.Style = yaml.FlowStyle
		}
		for _, s := range v {
			m := &yaml.Node{Kind: yaml.MappingNode}
			add := func(k, s string) {
				if s != "" {
					m.Content = append(m.Content, &yaml.Node{Kind: yaml.ScalarNode, Value: k}, strNode(s))
				}
			}
			add("path", s.Path)
			add("type", s.Type)
			add("run_command", s.RunCommand)
			for _, e := range s.Extra {
				m.Content = append(m.Content, &yaml.Node{Kind: yaml.ScalarNode, Value: e.Key}, e.Value)
			}
			val.Content = append(val.Content, m)
		}
	case string:
		val = *strNode(v)
	default:
		if err := val.Encode(v); err != nil {
			panic(err) // ints and bools always encode
		}
	}
	return k, &val
}

// strNode encodes a string so that it decodes back to the same text: the
// encoder quotes anything that would otherwise resolve to another type.
func strNode(s string) *yaml.Node {
	var n yaml.Node
	_ = n.Encode(s)
	if s == "" {
		n.Style = yaml.DoubleQuotedStyle
	}
	return &n
}

// fieldValue returns the Go value behind a schema key.
func fieldValue(r *Record, key string) any {
	switch knownFields[key] {
	case kindString:
		return *stringField(r, key)
	case kindList:
		return *listField(r, key)
	case kindSpecs:
		return r.AssociatedSpecs
	case kindBool:
		return r.Locked
	}
	if key == "version" {
		return r.Version
	}
	return r.ImplementsVersion
}

func isZero(v any) bool {
	switch v := v.(type) {
	case []string:
		return len(v) == 0
	case []AssociatedSpec:
		return len(v) == 0
	}
	return reflect.ValueOf(v).IsZero()
}

// equalValue compares two field values, treating nil and empty lists alike
// and spec Extra nodes by their encoded form.
func equalValue(a, b any) bool {
	if isZero(a) && isZero(b) {
		return true
	}
	sa, aok := a.([]AssociatedSpec)
	sb, bok := b.([]AssociatedSpec)
	if aok && bok {
		if len(sa) != len(sb) {
			return false
		}
		for i := range sa {
			if sa[i].Path != sb[i].Path || sa[i].Type != sb[i].Type || sa[i].RunCommand != sb[i].RunCommand || len(sa[i].Extra) != len(sb[i].Extra) {
				return false
			}
			for j := range sa[i].Extra {
				if sa[i].Extra[j].Key != sb[i].Extra[j].Key || sa[i].Extra[j].Value != sb[i].Extra[j].Value {
					return false
				}
			}
		}
		return true
	}
	return reflect.DeepEqual(a, b)
}

func cloneRecord(r *Record) *Record {
	c := *r
	cp := func(s []string) []string {
		if s == nil {
			return nil
		}
		return append([]string(nil), s...)
	}
	c.Aliases, c.Constraints = cp(r.Aliases), cp(r.Constraints)
	c.AffectedScope, c.ForbiddenScope = cp(r.AffectedScope), cp(r.ForbiddenScope)
	c.Related, c.AssociatedTraces = cp(r.Related), cp(r.AssociatedTraces)
	c.Monitors, c.Tags = cp(r.Monitors), cp(r.Tags)
	if r.AssociatedSpecs != nil {
		c.AssociatedSpecs = make([]AssociatedSpec, len(r.AssociatedSpecs))
		for i, s := range r.AssociatedSpecs {
			s.Extra = append([]KeyNode(nil), s.Extra...)
			c.AssociatedSpecs[i] = s
		}
	}
	return &c
}

// Load reads and parses the record file at path.
func Load(path string) (*Document, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	d, err := Parse(data)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", path, err)
	}
	return d, nil
}

// Save writes r to path. When doc is non-nil (the file as last loaded), only
// changed keys are rewritten; otherwise the record is marshalled fresh.
func Save(path string, doc *Document, r *Record) error {
	var out []byte
	var err error
	if doc != nil {
		out, err = doc.Encode(r)
	} else {
		out, err = Marshal(r)
	}
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	return os.WriteFile(path, out, 0o644)
}
