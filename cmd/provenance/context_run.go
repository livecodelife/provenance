package main

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:00fcaf0ee06fe8975724189121d1fe777e13a1cb2bdd0442ca7e97169a62f987","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/context_run.go","path":"encoding/json"}}
import "encoding/json"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:00fcaf0ee06fe8975724189121d1fe777e13a1cb2bdd0442ca7e97169a62f987","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/context_run.go","path":"fmt"}}
import "fmt"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:00fcaf0ee06fe8975724189121d1fe777e13a1cb2bdd0442ca7e97169a62f987","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/context_run.go","path":"io"}}
import "io"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:00fcaf0ee06fe8975724189121d1fe777e13a1cb2bdd0442ca7e97169a62f987","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/context_run.go","path":"sort"}}
import "sort"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:00fcaf0ee06fe8975724189121d1fe777e13a1cb2bdd0442ca7e97169a62f987","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/context_run.go","path":"strings"}}
import "strings"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:00fcaf0ee06fe8975724189121d1fe777e13a1cb2bdd0442ca7e97169a62f987","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/context_run.go","path":"github.com/spf13/cobra"}}
import "github.com/spf13/cobra"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:00fcaf0ee06fe8975724189121d1fe777e13a1cb2bdd0442ca7e97169a62f987","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/context_run.go","path":"github.com/livecodelife/provenance/record"}}
import "github.com/livecodelife/provenance/record"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:00fcaf0ee06fe8975724189121d1fe777e13a1cb2bdd0442ca7e97169a62f987","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/context_run.go","path":"github.com/livecodelife/provenance/stale"}}
import "github.com/livecodelife/provenance/stale"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setType {"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:00fcaf0ee06fe8975724189121d1fe777e13a1cb2bdd0442ca7e97169a62f987","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type ctxRecord struct {\n\tID                string   `json:\"id\"`\n\tType              string   `json:\"type\"`\n\tStatus            string   `json:\"status\"`\n\tVersion           int      `json:\"version\"`\n\tTitle             string   `json:\"title\"`\n\tLatestReason      string   `json:\"latest_reason,omitempty\"`\n\tImplements        string   `json:\"implements,omitempty\"`\n\tImplementsVersion int      `json:\"implements_version,omitempty\"`\n\tStale             string   `json:\"stale,omitempty\"`\n\tEvidence          string   `json:\"evidence,omitempty\"`\n\tSpecs             []string `json:\"specs\"`\n}","file":"cmd/provenance/context_run.go","name":"ctxRecord"}}
type ctxRecord struct {
	ID                string   `json:"id"`
	Type              string   `json:"type"`
	Status            string   `json:"status"`
	Version           int      `json:"version"`
	Title             string   `json:"title"`
	LatestReason      string   `json:"latest_reason,omitempty"`
	Implements        string   `json:"implements,omitempty"`
	ImplementsVersion int      `json:"implements_version,omitempty"`
	Stale             string   `json:"stale,omitempty"`
	Evidence          string   `json:"evidence,omitempty"`
	Specs             []string `json:"specs"`
}

// /sedum:setType
// sedum:setType {"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:00fcaf0ee06fe8975724189121d1fe777e13a1cb2bdd0442ca7e97169a62f987","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type ctxAncestor struct {\n\tID      string `json:\"id\"`\n\tVersion int    `json:\"version\"`\n\tStatus  string `json:\"status\"`\n\tTitle   string `json:\"title\"`\n\tShared  bool   `json:\"shared,omitempty\"`\n}","file":"cmd/provenance/context_run.go","name":"ctxAncestor"}}
type ctxAncestor struct {
	ID      string `json:"id"`
	Version int    `json:"version"`
	Status  string `json:"status"`
	Title   string `json:"title"`
	Shared  bool   `json:"shared,omitempty"`
}

// /sedum:setType
// sedum:setType {"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:00fcaf0ee06fe8975724189121d1fe777e13a1cb2bdd0442ca7e97169a62f987","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type ctxVersion struct {\n\tVersion int    `json:\"version\"`\n\tReason  string `json:\"reason,omitempty\"`\n}","file":"cmd/provenance/context_run.go","name":"ctxVersion"}}
type ctxVersion struct {
	Version int    `json:"version"`
	Reason  string `json:"reason,omitempty"`
}

// /sedum:setType
// sedum:setType {"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:00fcaf0ee06fe8975724189121d1fe777e13a1cb2bdd0442ca7e97169a62f987","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type ctxChild struct {\n\tID      string `json:\"id\"`\n\tVersion int    `json:\"version\"`\n\tPinned  int    `json:\"pinned\"`\n\tStale   string `json:\"stale,omitempty\"`\n}","file":"cmd/provenance/context_run.go","name":"ctxChild"}}
type ctxChild struct {
	ID      string `json:"id"`
	Version int    `json:"version"`
	Pinned  int    `json:"pinned"`
	Stale   string `json:"stale,omitempty"`
}

// /sedum:setType
// sedum:setType {"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:00fcaf0ee06fe8975724189121d1fe777e13a1cb2bdd0442ca7e97169a62f987","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type ctxRecordDetail struct {\n\tctxRecord\n\tIntent         string        `json:\"intent\"`\n\tConstraints    []string      `json:\"constraints\"`\n\tAffectedScope  []string      `json:\"affected_scope\"`\n\tForbiddenScope []string      `json:\"forbidden_scope\"`\n\tVersions       []ctxVersion  `json:\"versions\"`\n\tParents        []ctxAncestor `json:\"parents\"`\n\tChildren       []ctxChild    `json:\"children\"`\n\tCommits        []string      `json:\"commits\"`\n}","file":"cmd/provenance/context_run.go","name":"ctxRecordDetail"}}
type ctxRecordDetail struct {
	ctxRecord
	Intent         string        `json:"intent"`
	Constraints    []string      `json:"constraints"`
	AffectedScope  []string      `json:"affected_scope"`
	ForbiddenScope []string      `json:"forbidden_scope"`
	Versions       []ctxVersion  `json:"versions"`
	Parents        []ctxAncestor `json:"parents"`
	Children       []ctxChild    `json:"children"`
	Commits        []string      `json:"commits"`
}

// /sedum:setType
// sedum:setType {"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:00fcaf0ee06fe8975724189121d1fe777e13a1cb2bdd0442ca7e97169a62f987","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type ctxTarget struct {\n\tTarget    string                   `json:\"target\"`\n\tKind      string                   `json:\"kind\"`\n\tExempt    bool                     `json:\"exempt,omitempty\"`\n\tGoverning []ctxRecord              `json:\"governing,omitempty\"`\n\tAncestry  map[string][]ctxAncestor `json:\"ancestry,omitempty\"`\n\tConflicts []string                 `json:\"conflicts,omitempty\"`\n\tNext      *nextAction              `json:\"next,omitempty\"`\n\tRecord    *ctxRecordDetail         `json:\"record,omitempty\"`\n}","file":"cmd/provenance/context_run.go","name":"ctxTarget"}}
type ctxTarget struct {
	Target    string                   `json:"target"`
	Kind      string                   `json:"kind"`
	Exempt    bool                     `json:"exempt,omitempty"`
	Governing []ctxRecord              `json:"governing,omitempty"`
	Ancestry  map[string][]ctxAncestor `json:"ancestry,omitempty"`
	Conflicts []string                 `json:"conflicts,omitempty"`
	Next      *nextAction              `json:"next,omitempty"`
	Record    *ctxRecordDetail         `json:"record,omitempty"`
}

// /sedum:setType
// sedum:setType {"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:00fcaf0ee06fe8975724189121d1fe777e13a1cb2bdd0442ca7e97169a62f987","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type ctxDoc struct {\n\tTargets []ctxTarget `json:\"targets\"`\n}","file":"cmd/provenance/context_run.go","name":"ctxDoc"}}
type ctxDoc struct {
	Targets []ctxTarget `json:"targets"`
}

// /sedum:setType
// sedum:setType {"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:00fcaf0ee06fe8975724189121d1fe777e13a1cb2bdd0442ca7e97169a62f987","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type ctxWorld struct {\n\tr        *repo\n\tlocal    map[string]*record.Record\n\tshared   map[string]*record.Record\n\tstale    map[string]string\n\tevidence map[string]string\n}","file":"cmd/provenance/context_run.go","name":"ctxWorld"}}
type ctxWorld struct {
	r        *repo
	local    map[string]*record.Record
	shared   map[string]*record.Record
	stale    map[string]string
	evidence map[string]string
}

// /sedum:setType
// sedum:setFunction {"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:00fcaf0ee06fe8975724189121d1fe777e13a1cb2bdd0442ca7e97169a62f987","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func ctxClean(s string, max int) string {\n\tvar b strings.Builder\n\tfor _, c := range s {\n\t\tswitch {\n\t\tcase c == '\\n' || c == '\\t':\n\t\t\tb.WriteRune(' ')\n\t\tcase c < 32 || (c >= 127 && c <= 159):\n\t\tdefault:\n\t\t\tb.WriteRune(c)\n\t\t}\n\t}\n\tout := b.String()\n\tif max > 0 && len([]rune(out)) > max {\n\t\treturn string([]rune(out)[:max-3]) + \"...\"\n\t}\n\treturn out\n}","file":"cmd/provenance/context_run.go","name":"ctxClean"}}
func ctxClean(s string, max int) string {
	var b strings.Builder
	for _, c := range s {
		switch {
		case c == '\n' || c == '\t':
			b.WriteRune(' ')
		case c < 32 || (c >= 127 && c <= 159):
		default:
			b.WriteRune(c)
		}
	}
	out := b.String()
	if max > 0 && len([]rune(out)) > max {
		return string([]rune(out)[:max-3]) + "..."
	}
	return out
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:00fcaf0ee06fe8975724189121d1fe777e13a1cb2bdd0442ca7e97169a62f987","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func loadCtxWorld(r *repo) (*ctxWorld, error) {\n\tw := &ctxWorld{r: r, local: map[string]*record.Record{}, stale: map[string]string{}, evidence: map[string]string{}}\n\tids, err := r.Store.List()\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\tfor _, id := range ids {\n\t\trec, _, err := r.Store.Load(id)\n\t\tif err != nil {\n\t\t\treturn nil, err\n\t\t}\n\t\tw.local[id] = rec\n\t}\n\tw.shared, _ = loadShared(r)\n\tfound, err := stale.Find(r.Store)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\tfor _, f := range found {\n\t\tw.stale[f.ID] = string(f.Kind)\n\t}\n\trows, err := checkEvidence(r)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\tfor _, row := range rows {\n\t\tw.evidence[row.ID] = row.State\n\t}\n\treturn w, nil\n}","file":"cmd/provenance/context_run.go","name":"loadCtxWorld"}}
func loadCtxWorld(r *repo) (*ctxWorld, error) {
	w := &ctxWorld{r: r, local: map[string]*record.Record{}, stale: map[string]string{}, evidence: map[string]string{}}
	ids, err := r.Store.List()
	if err != nil {
		return nil, err
	}
	for _, id := range ids {
		rec, _, err := r.Store.Load(id)
		if err != nil {
			return nil, err
		}
		w.local[id] = rec
	}
	w.shared, _ = loadShared(r)
	found, err := stale.Find(r.Store)
	if err != nil {
		return nil, err
	}
	for _, f := range found {
		w.stale[f.ID] = string(f.Kind)
	}
	rows, err := checkEvidence(r)
	if err != nil {
		return nil, err
	}
	for _, row := range rows {
		w.evidence[row.ID] = row.State
	}
	return w, nil
}

// /sedum:setFunction
// sedum:setMethod {"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:00fcaf0ee06fe8975724189121d1fe777e13a1cb2bdd0442ca7e97169a62f987","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (w *ctxWorld) summary(rec *record.Record) ctxRecord {\n\tc := ctxRecord{ID: rec.ID, Type: string(rec.EffectiveType()), Status: string(rec.Status), Version: rec.Version, Title: rec.Title, Implements: rec.Implements, ImplementsVersion: rec.ImplementsVersion, Stale: w.stale[rec.ID], Evidence: w.evidence[rec.ID], Specs: []string{}}\n\tif versions, err := w.r.Store.Versions(rec.ID); err == nil && len(versions) > 0 {\n\t\tif snap, err := w.r.Store.Snapshot(rec.ID, versions[len(versions)-1]); err == nil {\n\t\t\tc.LatestReason = snap.Reason\n\t\t}\n\t}\n\tfor _, s := range rec.AssociatedSpecs {\n\t\tc.Specs = append(c.Specs, s.Path)\n\t}\n\treturn c\n}","file":"cmd/provenance/context_run.go","name":"summary","receiver":"ctxWorld"}}
func (w *ctxWorld) summary(rec *record.Record) ctxRecord {
	c := ctxRecord{ID: rec.ID, Type: string(rec.EffectiveType()), Status: string(rec.Status), Version: rec.Version, Title: rec.Title, Implements: rec.Implements, ImplementsVersion: rec.ImplementsVersion, Stale: w.stale[rec.ID], Evidence: w.evidence[rec.ID], Specs: []string{}}
	if versions, err := w.r.Store.Versions(rec.ID); err == nil && len(versions) > 0 {
		if snap, err := w.r.Store.Snapshot(rec.ID, versions[len(versions)-1]); err == nil {
			c.LatestReason = snap.Reason
		}
	}
	for _, s := range rec.AssociatedSpecs {
		c.Specs = append(c.Specs, s.Path)
	}
	return c
}

// /sedum:setMethod
// sedum:setMethod {"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:00fcaf0ee06fe8975724189121d1fe777e13a1cb2bdd0442ca7e97169a62f987","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (w *ctxWorld) ancestors(rec *record.Record) []ctxAncestor {\n\tout := []ctxAncestor{}\n\tseen := map[string]bool{rec.ID: true}\n\tfor next := rec.Implements; next != \"\" && !seen[next] && len(out) < 10; {\n\t\tseen[next] = true\n\t\tif p, ok := w.local[next]; ok {\n\t\t\tout = append(out, ctxAncestor{ID: p.ID, Version: p.Version, Status: string(p.Status), Title: p.Title})\n\t\t\tnext = p.Implements\n\t\t} else if p, ok := w.shared[next]; ok {\n\t\t\tout = append(out, ctxAncestor{ID: p.ID, Version: p.Version, Status: string(p.Status), Title: p.Title, Shared: true})\n\t\t\tnext = p.Implements\n\t\t} else {\n\t\t\tout = append(out, ctxAncestor{ID: next, Status: \"unresolved\"})\n\t\t\tbreak\n\t\t}\n\t}\n\treturn out\n}","file":"cmd/provenance/context_run.go","name":"ancestors","receiver":"ctxWorld"}}
func (w *ctxWorld) ancestors(rec *record.Record) []ctxAncestor {
	out := []ctxAncestor{}
	seen := map[string]bool{rec.ID: true}
	for next := rec.Implements; next != "" && !seen[next] && len(out) < 10; {
		seen[next] = true
		if p, ok := w.local[next]; ok {
			out = append(out, ctxAncestor{ID: p.ID, Version: p.Version, Status: string(p.Status), Title: p.Title})
			next = p.Implements
		} else if p, ok := w.shared[next]; ok {
			out = append(out, ctxAncestor{ID: p.ID, Version: p.Version, Status: string(p.Status), Title: p.Title, Shared: true})
			next = p.Implements
		} else {
			out = append(out, ctxAncestor{ID: next, Status: "unresolved"})
			break
		}
	}
	return out
}

// /sedum:setMethod
// sedum:setMethod {"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:00fcaf0ee06fe8975724189121d1fe777e13a1cb2bdd0442ca7e97169a62f987","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (w *ctxWorld) pathTarget(path string) (ctxTarget, error) {\n\tt := ctxTarget{Target: path, Kind: \"path\"}\n\tif !nextRelevant(w.r, path) {\n\t\tt.Exempt = true\n\t\treturn t, nil\n\t}\n\tvar ids []string\n\tfor id := range w.local {\n\t\tids = append(ids, id)\n\t}\n\tsort.Strings(ids)\n\tvar open []string\n\tt.Ancestry = map[string][]ctxAncestor{}\n\tfor _, id := range ids {\n\t\trec := w.local[id]\n\t\tif rec.Status == record.StatusDeprecated || !nextGoverns(rec, path) {\n\t\t\tcontinue\n\t\t}\n\t\tt.Governing = append(t.Governing, w.summary(rec))\n\t\tt.Ancestry[id] = w.ancestors(rec)\n\t\tif rec.Status == record.StatusOpen {\n\t\t\topen = append(open, id)\n\t\t}\n\t}\n\tif len(open) > 1 {\n\t\tt.Conflicts = open\n\t}\n\tst, err := gatherNext(w.r, []string{path})\n\tif err != nil {\n\t\treturn t, err\n\t}\n\tif actions := adviseNext(st); len(actions) > 0 {\n\t\tt.Next = &actions[0]\n\t}\n\treturn t, nil\n}","file":"cmd/provenance/context_run.go","name":"pathTarget","receiver":"ctxWorld"}}
func (w *ctxWorld) pathTarget(path string) (ctxTarget, error) {
	t := ctxTarget{Target: path, Kind: "path"}
	if !nextRelevant(w.r, path) {
		t.Exempt = true
		return t, nil
	}
	var ids []string
	for id := range w.local {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	var open []string
	t.Ancestry = map[string][]ctxAncestor{}
	for _, id := range ids {
		rec := w.local[id]
		if rec.Status == record.StatusDeprecated || !nextGoverns(rec, path) {
			continue
		}
		t.Governing = append(t.Governing, w.summary(rec))
		t.Ancestry[id] = w.ancestors(rec)
		if rec.Status == record.StatusOpen {
			open = append(open, id)
		}
	}
	if len(open) > 1 {
		t.Conflicts = open
	}
	st, err := gatherNext(w.r, []string{path})
	if err != nil {
		return t, err
	}
	if actions := adviseNext(st); len(actions) > 0 {
		t.Next = &actions[0]
	}
	return t, nil
}

// /sedum:setMethod
// sedum:setMethod {"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:00fcaf0ee06fe8975724189121d1fe777e13a1cb2bdd0442ca7e97169a62f987","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (w *ctxWorld) recordTarget(rec *record.Record) ctxTarget {\n\td := &ctxRecordDetail{ctxRecord: w.summary(rec), Intent: rec.Intent, Constraints: append([]string{}, rec.Constraints...), AffectedScope: append([]string{}, rec.AffectedScope...), ForbiddenScope: append([]string{}, rec.ForbiddenScope...), Versions: []ctxVersion{}, Parents: w.ancestors(rec), Children: []ctxChild{}, Commits: []string{}}\n\tif versions, err := w.r.Store.Versions(rec.ID); err == nil {\n\t\tfor _, v := range versions {\n\t\t\tif snap, err := w.r.Store.Snapshot(rec.ID, v); err == nil {\n\t\t\t\td.Versions = append(d.Versions, ctxVersion{Version: v, Reason: snap.Reason})\n\t\t\t}\n\t\t}\n\t}\n\tvar kids []string\n\tfor id, c := range w.local {\n\t\tif c.Implements == rec.ID {\n\t\t\tkids = append(kids, id)\n\t\t}\n\t}\n\tsort.Strings(kids)\n\tfor _, id := range kids {\n\t\tc := w.local[id]\n\t\td.Children = append(d.Children, ctxChild{ID: id, Version: c.Version, Pinned: c.ImplementsVersion, Stale: w.stale[id]})\n\t}\n\tfor _, sha := range commitsForRecord(w.r, rec) {\n\t\tif len(d.Commits) >= 20 {\n\t\t\tbreak\n\t\t}\n\t\tif line, err := w.r.git(\"log\", \"-1\", \"--format=%h %s\", sha); err == nil {\n\t\t\td.Commits = append(d.Commits, line)\n\t\t}\n\t}\n\treturn ctxTarget{Target: rec.ID, Kind: \"record\", Record: d}\n}","file":"cmd/provenance/context_run.go","name":"recordTarget","receiver":"ctxWorld"}}
func (w *ctxWorld) recordTarget(rec *record.Record) ctxTarget {
	d := &ctxRecordDetail{ctxRecord: w.summary(rec), Intent: rec.Intent, Constraints: append([]string{}, rec.Constraints...), AffectedScope: append([]string{}, rec.AffectedScope...), ForbiddenScope: append([]string{}, rec.ForbiddenScope...), Versions: []ctxVersion{}, Parents: w.ancestors(rec), Children: []ctxChild{}, Commits: []string{}}
	if versions, err := w.r.Store.Versions(rec.ID); err == nil {
		for _, v := range versions {
			if snap, err := w.r.Store.Snapshot(rec.ID, v); err == nil {
				d.Versions = append(d.Versions, ctxVersion{Version: v, Reason: snap.Reason})
			}
		}
	}
	var kids []string
	for id, c := range w.local {
		if c.Implements == rec.ID {
			kids = append(kids, id)
		}
	}
	sort.Strings(kids)
	for _, id := range kids {
		c := w.local[id]
		d.Children = append(d.Children, ctxChild{ID: id, Version: c.Version, Pinned: c.ImplementsVersion, Stale: w.stale[id]})
	}
	for _, sha := range commitsForRecord(w.r, rec) {
		if len(d.Commits) >= 20 {
			break
		}
		if line, err := w.r.git("log", "-1", "--format=%h %s", sha); err == nil {
			d.Commits = append(d.Commits, line)
		}
	}
	return ctxTarget{Target: rec.ID, Kind: "record", Record: d}
}

// /sedum:setMethod
// sedum:setFunction {"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:00fcaf0ee06fe8975724189121d1fe777e13a1cb2bdd0442ca7e97169a62f987","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func ctxLine(c ctxRecord, max int) string {\n\tline := fmt.Sprintf(\"%s@%d  %s  %s  %s\", c.ID, c.Version, c.Status, c.Type, ctxClean(c.Title, max))\n\tif c.Stale != \"\" {\n\t\tline += \"  stale (\" + c.Stale + \")\"\n\t}\n\tif c.Evidence != \"\" {\n\t\tline += \"  evidence: \" + c.Evidence\n\t}\n\treturn line\n}","file":"cmd/provenance/context_run.go","name":"ctxLine"}}
func ctxLine(c ctxRecord, max int) string {
	line := fmt.Sprintf("%s@%d  %s  %s  %s", c.ID, c.Version, c.Status, c.Type, ctxClean(c.Title, max))
	if c.Stale != "" {
		line += "  stale (" + c.Stale + ")"
	}
	if c.Evidence != "" {
		line += "  evidence: " + c.Evidence
	}
	return line
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:00fcaf0ee06fe8975724189121d1fe777e13a1cb2bdd0442ca7e97169a62f987","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func printCtxHuman(out io.Writer, t ctxTarget) {\n\tfmt.Fprintf(out, \"\\nCONTEXT: %s\\n\", t.Target)\n\tif t.Kind == \"path\" {\n\t\tif t.Exempt {\n\t\t\tfmt.Fprint(out, \"  Exempt: no provenance rules apply to this path.\\n\")\n\t\t\treturn\n\t\t}\n\t\tif len(t.Governing) == 0 {\n\t\t\tfmt.Fprint(out, \"  No record governs this path.\\n\")\n\t\t}\n\t\tfor _, g := range t.Governing {\n\t\t\tfmt.Fprintf(out, \"  %s\\n\", ctxLine(g, 120))\n\t\t\tif g.LatestReason != \"\" {\n\t\t\t\tfmt.Fprintf(out, \"      latest: %s\\n\", ctxClean(g.LatestReason, 160))\n\t\t\t}\n\t\t\tif len(g.Specs) > 0 {\n\t\t\t\tfmt.Fprintf(out, \"      specs: %s\\n\", ctxClean(strings.Join(g.Specs, \", \"), 200))\n\t\t\t}\n\t\t\tfor _, a := range t.Ancestry[g.ID] {\n\t\t\t\tmark := \"\"\n\t\t\t\tif a.Shared {\n\t\t\t\t\tmark = \" (shared)\"\n\t\t\t\t}\n\t\t\t\tfmt.Fprintf(out, \"      from: %s@%d  %s  %s%s\\n\", a.ID, a.Version, a.Status, ctxClean(a.Title, 120), mark)\n\t\t\t}\n\t\t}\n\t\tif len(t.Conflicts) > 0 {\n\t\t\tfmt.Fprintf(out, \"  Conflict: %d open records cover this path: %s\\n\", len(t.Conflicts), strings.Join(t.Conflicts, \", \"))\n\t\t}\n\t\tif t.Next != nil {\n\t\t\tfmt.Fprintf(out, \"  Next: %s: %s\\n\", t.Next.Kind, ctxClean(t.Next.Reason, 300))\n\t\t\tif t.Next.Command != \"\" {\n\t\t\t\tfmt.Fprintf(out, \"    Run: %s\\n\", ctxClean(t.Next.Command, 300))\n\t\t\t}\n\t\t}\n\t\treturn\n\t}\n\td := t.Record\n\tfmt.Fprintf(out, \"  %s\\n\", ctxLine(d.ctxRecord, 120))\n\tfmt.Fprint(out, \"  Intent:\\n\")\n\tfor _, l := range strings.Split(d.Intent, \"\\n\") {\n\t\tfmt.Fprintf(out, \"    %s\\n\", ctxClean(l, 400))\n\t}\n\tsection := func(name string, items []string) {\n\t\tif len(items) == 0 {\n\t\t\treturn\n\t\t}\n\t\tfmt.Fprintf(out, \"  %s:\\n\", name)\n\t\tfor _, it := range items {\n\t\t\tfmt.Fprintf(out, \"    · %s\\n\", ctxClean(it, 400))\n\t\t}\n\t}\n\tsection(\"Constraints\", d.Constraints)\n\tsection(\"Affected scope\", d.AffectedScope)\n\tsection(\"Forbidden scope\", d.ForbiddenScope)\n\tsection(\"Specs\", d.Specs)\n\tfmt.Fprint(out, \"  Versions:\\n\")\n\tfor _, v := range d.Versions {\n\t\tfmt.Fprintf(out, \"    · %d  %s\\n\", v.Version, ctxClean(v.Reason, 200))\n\t}\n\tfor _, p := range d.Parents {\n\t\tfmt.Fprintf(out, \"  Implements: %s@%d  %s  %s\\n\", p.ID, p.Version, p.Status, ctxClean(p.Title, 120))\n\t}\n\tfor _, c := range d.Children {\n\t\tfmt.Fprintf(out, \"  Implemented by: %s@%d -> pinned @%d %s\\n\", c.ID, c.Version, c.Pinned, c.Stale)\n\t}\n\tsection(\"Commits\", d.Commits)\n}","file":"cmd/provenance/context_run.go","name":"printCtxHuman"}}
func printCtxHuman(out io.Writer, t ctxTarget) {
	fmt.Fprintf(out, "\nCONTEXT: %s\n", t.Target)
	if t.Kind == "path" {
		if t.Exempt {
			fmt.Fprint(out, "  Exempt: no provenance rules apply to this path.\n")
			return
		}
		if len(t.Governing) == 0 {
			fmt.Fprint(out, "  No record governs this path.\n")
		}
		for _, g := range t.Governing {
			fmt.Fprintf(out, "  %s\n", ctxLine(g, 120))
			if g.LatestReason != "" {
				fmt.Fprintf(out, "      latest: %s\n", ctxClean(g.LatestReason, 160))
			}
			if len(g.Specs) > 0 {
				fmt.Fprintf(out, "      specs: %s\n", ctxClean(strings.Join(g.Specs, ", "), 200))
			}
			for _, a := range t.Ancestry[g.ID] {
				mark := ""
				if a.Shared {
					mark = " (shared)"
				}
				fmt.Fprintf(out, "      from: %s@%d  %s  %s%s\n", a.ID, a.Version, a.Status, ctxClean(a.Title, 120), mark)
			}
		}
		if len(t.Conflicts) > 0 {
			fmt.Fprintf(out, "  Conflict: %d open records cover this path: %s\n", len(t.Conflicts), strings.Join(t.Conflicts, ", "))
		}
		if t.Next != nil {
			fmt.Fprintf(out, "  Next: %s: %s\n", t.Next.Kind, ctxClean(t.Next.Reason, 300))
			if t.Next.Command != "" {
				fmt.Fprintf(out, "    Run: %s\n", ctxClean(t.Next.Command, 300))
			}
		}
		return
	}
	d := t.Record
	fmt.Fprintf(out, "  %s\n", ctxLine(d.ctxRecord, 120))
	fmt.Fprint(out, "  Intent:\n")
	for _, l := range strings.Split(d.Intent, "\n") {
		fmt.Fprintf(out, "    %s\n", ctxClean(l, 400))
	}
	section := func(name string, items []string) {
		if len(items) == 0 {
			return
		}
		fmt.Fprintf(out, "  %s:\n", name)
		for _, it := range items {
			fmt.Fprintf(out, "    · %s\n", ctxClean(it, 400))
		}
	}
	section("Constraints", d.Constraints)
	section("Affected scope", d.AffectedScope)
	section("Forbidden scope", d.ForbiddenScope)
	section("Specs", d.Specs)
	fmt.Fprint(out, "  Versions:\n")
	for _, v := range d.Versions {
		fmt.Fprintf(out, "    · %d  %s\n", v.Version, ctxClean(v.Reason, 200))
	}
	for _, p := range d.Parents {
		fmt.Fprintf(out, "  Implements: %s@%d  %s  %s\n", p.ID, p.Version, p.Status, ctxClean(p.Title, 120))
	}
	for _, c := range d.Children {
		fmt.Fprintf(out, "  Implemented by: %s@%d -> pinned @%d %s\n", c.ID, c.Version, c.Pinned, c.Stale)
	}
	section("Commits", d.Commits)
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:00fcaf0ee06fe8975724189121d1fe777e13a1cb2bdd0442ca7e97169a62f987","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func printCtxCompact(out io.Writer, t ctxTarget) {\n\tswitch {\n\tcase t.Kind == \"record\":\n\t\tfmt.Fprintln(out, ctxLine(t.Record.ctxRecord, 80))\n\tcase t.Exempt:\n\t\tfmt.Fprintf(out, \"%s: exempt\\n\", t.Target)\n\tcase len(t.Governing) == 0:\n\t\tfmt.Fprintf(out, \"%s: no governing record\\n\", t.Target)\n\tdefault:\n\t\tfor _, g := range t.Governing {\n\t\t\tfmt.Fprintf(out, \"%s: %s\\n\", t.Target, ctxLine(g, 80))\n\t\t}\n\t}\n}","file":"cmd/provenance/context_run.go","name":"printCtxCompact"}}
func printCtxCompact(out io.Writer, t ctxTarget) {
	switch {
	case t.Kind == "record":
		fmt.Fprintln(out, ctxLine(t.Record.ctxRecord, 80))
	case t.Exempt:
		fmt.Fprintf(out, "%s: exempt\n", t.Target)
	case len(t.Governing) == 0:
		fmt.Fprintf(out, "%s: no governing record\n", t.Target)
	default:
		for _, g := range t.Governing {
			fmt.Fprintf(out, "%s: %s\n", t.Target, ctxLine(g, 80))
		}
	}
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:00fcaf0ee06fe8975724189121d1fe777e13a1cb2bdd0442ca7e97169a62f987","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func runContext(out io.Writer, cmd *cobra.Command, args []string) error {\n\tr, err := openRepo(cmd)\n\tif err != nil {\n\t\treturn err\n\t}\n\tfiles, _ := cmd.Flags().GetStringSlice(\"files\")\n\tformat, _ := cmd.Flags().GetString(\"format\")\n\tswitch format {\n\tcase \"\", \"human\", \"compact\", \"json\":\n\tdefault:\n\t\treturn fmt.Errorf(\"unknown format %q (human|compact|json)\", format)\n\t}\n\ttargets := append(append([]string{}, args...), files...)\n\tif len(targets) == 0 {\n\t\treturn fmt.Errorf(\"no target: give file paths or record ids as arguments, or use --files\")\n\t}\n\tw, err := loadCtxWorld(r)\n\tif err != nil {\n\t\treturn err\n\t}\n\tdoc := ctxDoc{Targets: []ctxTarget{}}\n\tfor _, target := range targets {\n\t\tif rec, ok := w.local[target]; ok {\n\t\t\tdoc.Targets = append(doc.Targets, w.recordTarget(rec))\n\t\t\tcontinue\n\t\t}\n\t\tt, err := w.pathTarget(strings.TrimPrefix(target, \"./\"))\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n\t\tdoc.Targets = append(doc.Targets, t)\n\t}\n\tswitch format {\n\tcase \"json\":\n\t\tenc := json.NewEncoder(out)\n\t\tenc.SetIndent(\"\", \"  \")\n\t\treturn enc.Encode(doc)\n\tcase \"compact\":\n\t\tfor _, t := range doc.Targets {\n\t\t\tprintCtxCompact(out, t)\n\t\t}\n\tdefault:\n\t\tfor _, t := range doc.Targets {\n\t\t\tprintCtxHuman(out, t)\n\t\t}\n\t\tfmt.Fprint(out, \"\\n\")\n\t}\n\treturn nil\n}","file":"cmd/provenance/context_run.go","name":"runContext"}}
func runContext(out io.Writer, cmd *cobra.Command, args []string) error {
	r, err := openRepo(cmd)
	if err != nil {
		return err
	}
	files, _ := cmd.Flags().GetStringSlice("files")
	format, _ := cmd.Flags().GetString("format")
	switch format {
	case "", "human", "compact", "json":
	default:
		return fmt.Errorf("unknown format %q (human|compact|json)", format)
	}
	targets := append(append([]string{}, args...), files...)
	if len(targets) == 0 {
		return fmt.Errorf("no target: give file paths or record ids as arguments, or use --files")
	}
	w, err := loadCtxWorld(r)
	if err != nil {
		return err
	}
	doc := ctxDoc{Targets: []ctxTarget{}}
	for _, target := range targets {
		if rec, ok := w.local[target]; ok {
			doc.Targets = append(doc.Targets, w.recordTarget(rec))
			continue
		}
		t, err := w.pathTarget(strings.TrimPrefix(target, "./"))
		if err != nil {
			return err
		}
		doc.Targets = append(doc.Targets, t)
	}
	switch format {
	case "json":
		enc := json.NewEncoder(out)
		enc.SetIndent("", "  ")
		return enc.Encode(doc)
	case "compact":
		for _, t := range doc.Targets {
			printCtxCompact(out, t)
		}
	default:
		for _, t := range doc.Targets {
			printCtxHuman(out, t)
		}
		fmt.Fprint(out, "\n")
	}
	return nil
}

// /sedum:setFunction
