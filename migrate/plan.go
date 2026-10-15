package migrate

import (
	"container/heap"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/livecodelife/provenance/record"
)

// Rule names which migration rule folds an absorbed id into its target.
const (
	Rule1  = "1"  // imprint superseded by an imprint with the same implements
	Rule2  = "2"  // blueprint superseded by a blueprint
	Rule3  = "3"  // bug supersedes a blueprint: the bug folds into the blueprint
	Rule4  = "4"  // bug extends a record
	RuleC3 = "C3" // bug supersedes a bug: follows rule 4 through the older bug
)

// Keep-separate thresholds (MIGRATION_DRY_RUN.md, settled C1 and A5).
const (
	KeepSeparateJaccard   = 1.0 / 3 // rule-2 edges sharing at most this much scope
	KeepSeparateExtenders = 10      // targets with this many bug extenders or more
)

// Entry folds an absorbed id into Target. Suggest is non-empty when --plan
// should propose keeping the two separate, and says why.
type Entry struct {
	Target  string
	Rule    string
	Suggest string
}

// Suggestion is a keep-separate proposal for one plan entry.
type Suggestion struct {
	ID, Target, Rule, Why string
	// AsImprint is set for rule-4 suggestions: the shape the bug takes if kept.
	AsImprint *KeptImprint
}

// KeptImprint is a bug kept separate, as the imprint it becomes.
type KeptImprint struct {
	Implements string // empty when unresolvable
	RelatedAdd []string
	Note       string
	HasNote    bool // the oracle distinguishes a null note from no note
}

// KeptBug is a deprecated (won't-fix) bug, which never folds.
type KeptBug struct {
	ID, Extends string
	AsImprint   KeptImprint
}

// Warning is one thing --plan reports. Fields mirror the dry run's issue
// objects so they can be compared directly.
type Warning struct {
	Kind   string
	Fields map[string]any
}

// Plan is the read-only result of planning a repository's migration.
type Plan struct {
	Repo *Repo
	// Entries maps each absorbed id to where it folds.
	Entries map[string][]*Entry
	// Survivor maps every local id to the id it ends up in, ignoring
	// keep-separate suggestions; "" when unresolvable.
	Survivor map[string]string
	// Versions lists each survivor's ids in version order (survivors with
	// two or more versions only).
	Versions map[string][]string
	// Suggestions are the keep-separate proposals.
	Suggestions []Suggestion
	// KeptDeprecated are deprecated bugs, which become deprecated imprints.
	KeptDeprecated []KeptBug
	Warnings       []Warning
}

// Options are the --plan choices that change what Apply does.
type Options struct {
	// AcceptSuggestions keeps every suggested entry separate.
	AcceptSuggestions bool
	// KeepSeparate overrides per absorbed id: true keeps it separate, false
	// folds it even when suggested.
	KeepSeparate map[string]bool
}

func (o Options) kept(id string, e *Entry) bool {
	if v, ok := o.KeepSeparate[id]; ok {
		return v
	}
	return o.AcceptSuggestions && e.Suggest != ""
}

func ref(s string) string { return strings.TrimSpace(s) }

// MakePlan plans the migration of a loaded repository. It never writes.
func MakePlan(repo *Repo) (*Plan, error) {
	p := &Plan{Repo: repo, Entries: map[string][]*Entry{}, Survivor: map[string]string{}, Versions: map[string][]string{}}
	R := repo.Records
	ids := sortedIDs(R)

	// Supersession edges old -> new from both fields, local ends only.
	type edge struct{ old, new string }
	edgeSet := map[edge]bool{}
	for _, id := range ids {
		r := R[id].Rec
		if o := ref(r.Supersedes); o != "" {
			edgeSet[edge{o, id}] = true
		}
		if n := ref(r.SupersededBy); n != "" {
			edgeSet[edge{id, n}] = true
		}
	}
	succ, pred := map[string][]string{}, map[string][]string{}
	var edges []edge
	for e := range edgeSet {
		for _, end := range []string{e.old, e.new} {
			switch repo.ref(end) {
			case "":
				p.warn("dangling_ref", "field", "supersession", "target", end)
			case "shared":
				p.warn("cross_repo_ref", "field", "supersession", "target", end)
			}
		}
		if _, ok := R[e.old]; !ok {
			continue
		}
		if _, ok := R[e.new]; !ok {
			continue
		}
		edges = append(edges, e)
		succ[e.old] = append(succ[e.old], e.new)
		pred[e.new] = append(pred[e.new], e.old)
	}
	sort.Slice(edges, func(i, j int) bool {
		if edges[i].old != edges[j].old {
			return edges[i].old < edges[j].old
		}
		return edges[i].new < edges[j].new
	})
	for u, vs := range succ {
		if len(vs) > 1 {
			p.warn("branch_split", "old", u, "successors", sortedCopy(vs))
		}
	}
	for v, us := range pred {
		if len(us) > 1 {
			p.warn("branch_merge", "new", v, "predecessors", sortedCopy(us))
		}
	}
	for _, e := range edges {
		p.classifyEdge(R[e.old], R[e.new])
	}

	for _, id := range ids {
		r := R[id].Rec
		for _, f := range []struct{ name, target string }{{"extends", ref(r.Extends)}, {"implements", ref(r.Implements)}} {
			if f.target == "" {
				continue
			}
			switch repo.ref(f.target) {
			case "":
				p.warn("dangling_ref", "field", f.name, "src", id, "target", f.target)
			case "shared":
				p.warn("cross_repo_ref", "field", f.name, "src", id, "target", f.target)
			}
		}
		if ref(r.Extends) != "" {
			p.classifyExtends(R[id])
		}
	}

	p.suggestKeepSeparate()

	// Survivors and version order.
	members := map[string][]string{}
	for _, id := range ids {
		s := p.survivor(id, Options{})
		p.Survivor[id] = s
		if s != "" {
			members[s] = append(members[s], id)
		}
	}
	for _, s := range sortedKeys(members) {
		m := members[s]
		if len(m) < 2 {
			continue
		}
		order := orderVersions(R, m)
		p.Versions[s] = order
		p.statusWarnings(s, order)
	}
	for _, old := range sortedKeys(p.Entries) {
		targets := map[string]bool{}
		for _, e := range p.Entries[old] {
			targets[e.Target] = true
		}
		if len(targets) > 1 {
			p.warn("alias_collision_multi_target", "alias", old, "targets", sortedKeys(targets))
		}
		if p.Survivor[old] == "" {
			p.warn("survivor_unresolvable", "id", old)
		}
	}
	return p, nil
}

func (p *Plan) warn(kind string, kv ...any) {
	f := map[string]any{}
	for i := 0; i+1 < len(kv); i += 2 {
		f[kv[i].(string)] = kv[i+1]
	}
	p.Warnings = append(p.Warnings, Warning{Kind: kind, Fields: f})
}

func (p *Plan) add(absorbed, target, rule string) {
	p.Entries[absorbed] = append(p.Entries[absorbed], &Entry{Target: target, Rule: rule})
}

func (p *Plan) classifyEdge(old, new *Source) {
	ot, nt := old.Type(), new.Type()
	switch {
	case ot == record.TypeImprint && nt == record.TypeImprint:
		oi, ni := ref(old.Rec.Implements), ref(new.Rec.Implements)
		if oi != "" && oi == ni {
			p.add(old.Rec.ID, new.Rec.ID, Rule1)
		} else {
			p.warn("imprint_superseded_different_parent", "old", old.Rec.ID, "new", new.Rec.ID)
		}
	case ot == record.TypeBlueprint && nt == record.TypeBlueprint:
		p.add(old.Rec.ID, new.Rec.ID, Rule2)
	case ot == record.TypeBlueprint && nt == record.TypeBug:
		p.add(new.Rec.ID, old.Rec.ID, Rule3)
	case ot == record.TypeBug && nt == record.TypeBug && new.Rec.Status != record.StatusDeprecated:
		p.add(new.Rec.ID, old.Rec.ID, RuleC3)
	default:
		p.warn("cross_type_supersession", "old", old.Rec.ID, "new", new.Rec.ID, "old_type", string(ot), "new_type", string(nt))
	}
}

func (p *Plan) classifyExtends(s *Source) {
	r := s.Rec
	if s.Type() != record.TypeBug {
		p.warn("non_bug_extends", "id", r.ID, "extends", ref(r.Extends))
		return
	}
	t := ref(r.Extends)
	if p.Repo.ref(t) == "" {
		return
	}
	if r.Status == record.StatusDeprecated {
		p.KeptDeprecated = append(p.KeptDeprecated, KeptBug{ID: r.ID, Extends: t, AsImprint: keptBugAsImprint(p.Repo.Records, s)})
		return
	}
	if _, local := p.Repo.Records[t]; local {
		p.add(r.ID, t, Rule4)
	}
}

// keptBugAsImprint is the settled shape of a bug kept separate: implements
// the extended record if it is a brief or blueprint, else that record's own
// implements target, with the extended record added to related.
func keptBugAsImprint(R map[string]*Source, bug *Source) KeptImprint {
	t, ok := R[ref(bug.Rec.Extends)]
	if !ok {
		return KeptImprint{RelatedAdd: []string{}, Note: "extended record not local", HasNote: true}
	}
	switch t.Type() {
	case record.TypeBrief, record.TypeBlueprint:
		return KeptImprint{Implements: t.Rec.ID, RelatedAdd: []string{}}
	}
	k := KeptImprint{Implements: ref(t.Rec.Implements), RelatedAdd: []string{t.Rec.ID}, HasNote: true}
	if t.Type() != record.TypeImprint {
		k.Note = "extended record is a " + string(t.Type())
	}
	return k
}

// jaccard mirrors the dry run: the ratio rounded to two decimals (a shared
// set of 0 elements counts as full overlap).
func jaccard(a, b []string) float64 {
	sa, sb := map[string]bool{}, map[string]bool{}
	for _, x := range a {
		sa[x] = true
	}
	for _, x := range b {
		sb[x] = true
	}
	union, inter := len(sa), 0
	for x := range sb {
		if sa[x] {
			inter++
		} else {
			union++
		}
	}
	if union == 0 {
		return 1
	}
	// Correctly rounded to 2 decimals, like Python's round(x, 2).
	v, _ := strconv.ParseFloat(strconv.FormatFloat(float64(inter)/float64(union), 'f', 2, 64), 64)
	return v
}

func (p *Plan) suggestKeepSeparate() {
	R := p.Repo.Records
	nExt := map[string]int{}
	for _, s := range R {
		if t := ref(s.Rec.Extends); t != "" {
			nExt[t]++
		}
	}
	for _, id := range sortedIDs(R) {
		for _, e := range p.Entries[id] {
			why := ""
			switch {
			case e.Rule == Rule2 && jaccard(R[id].Rec.AffectedScope, R[e.Target].Rec.AffectedScope) <= KeepSeparateJaccard:
				why = "low scope overlap"
			case e.Rule == Rule4 && nExt[e.Target] >= KeepSeparateExtenders:
				why = fmt.Sprintf("target has %d bug extenders", nExt[e.Target])
			}
			if why == "" {
				continue
			}
			e.Suggest = why
			sg := Suggestion{ID: id, Target: e.Target, Rule: e.Rule, Why: why}
			if e.Rule == Rule4 {
				k := keptBugAsImprint(R, R[id])
				sg.AsImprint = &k
			}
			p.Suggestions = append(p.Suggestions, sg)
		}
	}
}

// survivor follows absorbed -> target links; "" on a cycle or ambiguity.
func (p *Plan) survivor(id string, opts Options) string {
	seen := map[string]bool{}
	x := id
	for {
		entries, ok := p.Entries[x]
		if !ok {
			return x
		}
		targets := map[string]bool{}
		for _, e := range entries {
			if !opts.kept(x, e) {
				targets[e.Target] = true
			}
		}
		if len(targets) == 0 {
			return x
		}
		if len(targets) != 1 || seen[x] {
			return ""
		}
		seen[x] = true
		for t := range targets {
			x = t
		}
	}
}

// SurvivorWith returns where id ends up once opts' keep-separate choices are
// applied.
func (p *Plan) SurvivorWith(id string, opts Options) string { return p.survivor(id, opts) }

// RecordsAfter counts the records left once opts are applied.
func (p *Plan) RecordsAfter(opts Options) int {
	n := 0
	for id := range p.Repo.Records {
		if p.survivor(id, opts) == id {
			n++
		}
	}
	return n
}

var statusRank = map[record.Status]int{record.StatusDraft: 0, record.StatusOpen: 1, record.StatusImplemented: 2}

func (p *Plan) statusWarnings(s string, order []string) {
	R := p.Repo.Records
	var live []record.Status
	statuses := make([]string, len(order))
	for i, x := range order {
		st := R[x].Rec.Status
		statuses[i] = string(st)
		if _, ok := statusRank[st]; ok {
			live = append(live, st)
		}
	}
	if len(live) > 0 {
		max := 0
		for _, st := range live {
			if statusRank[st] > max {
				max = statusRank[st]
			}
		}
		if statusRank[live[len(live)-1]] < max {
			p.warn("status_goes_backwards", "survivor", s, "final", string(live[len(live)-1]), "newest", order[len(order)-1], "statuses", statuses)
		}
	}
	for i, x := range order[:len(order)-1] {
		st := R[x].Rec.Status
		if st != record.StatusDraft && st != record.StatusOpen {
			continue
		}
		later := false
		for _, y := range order[i+1:] {
			if R[y].Rec.Status == record.StatusImplemented {
				later = true
			}
		}
		if !later {
			continue
		}
		suggested := false
		for _, e := range p.Entries[x] {
			if e.Suggest != "" {
				suggested = true
			}
		}
		p.warn("unfinished_version_buried", "survivor", s, "id", x, "status", string(st), "suggested_keep_separate", suggested)
	}
	if own, ok := R[s]; ok && (own.Rec.Status == record.StatusDraft || own.Rec.Status == record.StatusOpen) &&
		len(live) > 0 && live[len(live)-1] == record.StatusImplemented {
		p.warn("implemented_only_through_fixes", "survivor", s, "own_status", string(own.Rec.Status))
	}
}

// timeKey is the settled tiebreak: git first-add time, then v1 created_at,
// then id.
type timeKey struct {
	unix    int64
	created string
	id      string
}

func keyOf(R map[string]*Source, id string) timeKey {
	k := timeKey{unix: 1 << 40, created: R[id].Rec.CreatedAt, id: id}
	if a := R[id].Added; a != nil {
		k.unix = a.Unix
	}
	return k
}

func (a timeKey) less(b timeKey) bool {
	if a.unix != b.unix {
		return a.unix < b.unix
	}
	if a.created != b.created {
		return a.created < b.created
	}
	return a.id < b.id
}

type keyHeap struct {
	R   map[string]*Source
	ids []string
}

func (h *keyHeap) Len() int           { return len(h.ids) }
func (h *keyHeap) Less(i, j int) bool { return keyOf(h.R, h.ids[i]).less(keyOf(h.R, h.ids[j])) }
func (h *keyHeap) Swap(i, j int)      { h.ids[i], h.ids[j] = h.ids[j], h.ids[i] }
func (h *keyHeap) Push(x any)         { h.ids = append(h.ids, x.(string)) }
func (h *keyHeap) Pop() any {
	x := h.ids[len(h.ids)-1]
	h.ids = h.ids[:len(h.ids)-1]
	return x
}

// orderVersions orders a survivor's members links first (a version follows
// what it supersedes or extends, and what names it in superseded_by), then by
// timeKey. A cycle falls back to time order for what remains.
func orderVersions(R map[string]*Source, ids []string) []string {
	in := map[string]bool{}
	for _, x := range ids {
		in[x] = true
	}
	deps := map[string]map[string]bool{}
	for _, x := range ids {
		deps[x] = map[string]bool{}
		for _, d := range []string{ref(R[x].Rec.Supersedes), ref(R[x].Rec.Extends)} {
			if in[d] && d != x {
				deps[x][d] = true
			}
		}
	}
	for _, x := range ids {
		if n := ref(R[x].Rec.SupersededBy); in[n] && n != x {
			deps[n][x] = true
		}
	}
	users := map[string][]string{}
	indeg := map[string]int{}
	for x, ds := range deps {
		indeg[x] = len(ds)
		for d := range ds {
			users[d] = append(users[d], x)
		}
	}
	h := &keyHeap{R: R}
	for _, x := range ids {
		if indeg[x] == 0 {
			h.ids = append(h.ids, x)
		}
	}
	heap.Init(h)
	var out []string
	for h.Len() > 0 {
		x := heap.Pop(h).(string)
		out = append(out, x)
		for _, u := range users[x] {
			indeg[u]--
			if indeg[u] == 0 {
				heap.Push(h, u)
			}
		}
	}
	if len(out) != len(ids) {
		done := map[string]bool{}
		for _, x := range out {
			done[x] = true
		}
		var rest []string
		for _, x := range ids {
			if !done[x] {
				rest = append(rest, x)
			}
		}
		sort.Slice(rest, func(i, j int) bool { return keyOf(R, rest[i]).less(keyOf(R, rest[j])) })
		out = append(out, rest...)
	}
	return out
}

func sortedIDs(m map[string]*Source) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

func sortedKeys[V any](m map[string]V) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

func sortedCopy(s []string) []string {
	out := append([]string(nil), s...)
	sort.Strings(out)
	return out
}

// StatusMappings lists the records whose v1 status superseded has to map to
// a v2 status under opts: a record left standing because its successor is
// kept separate (or not local). --plan shows them so a person can review.
func (p *Plan) StatusMappings(opts Options) []Warning {
	R := p.Repo.Records
	members := map[string][]string{}
	for _, id := range sortedIDs(R) {
		if s := p.survivor(id, opts); s != "" {
			members[s] = append(members[s], id)
		}
	}
	var out []Warning
	for _, s := range sortedKeys(members) {
		order := orderVersions(R, members[s])
		last := R[order[len(order)-1]].Rec
		if last.Status == record.StatusSuperseded {
			out = append(out, Warning{Kind: "superseded_status_mapped", Fields: map[string]any{
				"id": s, "from": string(last.Status), "to": string(v2Status(last)),
			}})
		}
	}
	return out
}
