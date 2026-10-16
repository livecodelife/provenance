package main

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-11876813","gen":"sha256:090b658e8472ed58a1fb75e4a4101e7f16690a1a70b99759f485f9af7dfaeb18","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/next_run.go","path":"encoding/json"}}
import "encoding/json"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-11876813","gen":"sha256:090b658e8472ed58a1fb75e4a4101e7f16690a1a70b99759f485f9af7dfaeb18","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/next_run.go","path":"fmt"}}
import "fmt"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-11876813","gen":"sha256:090b658e8472ed58a1fb75e4a4101e7f16690a1a70b99759f485f9af7dfaeb18","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/next_run.go","path":"io"}}
import "io"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-11876813","gen":"sha256:090b658e8472ed58a1fb75e4a4101e7f16690a1a70b99759f485f9af7dfaeb18","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/next_run.go","path":"path/filepath"}}
import "path/filepath"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-11876813","gen":"sha256:090b658e8472ed58a1fb75e4a4101e7f16690a1a70b99759f485f9af7dfaeb18","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/next_run.go","path":"sort"}}
import "sort"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-11876813","gen":"sha256:090b658e8472ed58a1fb75e4a4101e7f16690a1a70b99759f485f9af7dfaeb18","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/next_run.go","path":"strings"}}
import "strings"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-11876813","gen":"sha256:090b658e8472ed58a1fb75e4a4101e7f16690a1a70b99759f485f9af7dfaeb18","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/next_run.go","path":"github.com/spf13/cobra"}}
import "github.com/spf13/cobra"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-11876813","gen":"sha256:090b658e8472ed58a1fb75e4a4101e7f16690a1a70b99759f485f9af7dfaeb18","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/next_run.go","path":"github.com/livecodelife/provenance/record"}}
import "github.com/livecodelife/provenance/record"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-11876813","gen":"sha256:090b658e8472ed58a1fb75e4a4101e7f16690a1a70b99759f485f9af7dfaeb18","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/next_run.go","path":"github.com/livecodelife/provenance/scope"}}
import "github.com/livecodelife/provenance/scope"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setType {"tier":"owned","record":"prov-2026-11876813","gen":"sha256:090b658e8472ed58a1fb75e4a4101e7f16690a1a70b99759f485f9af7dfaeb18","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type nextAction struct {\n\tKind      string   `json:\"kind\"`\n\tCommand   string   `json:\"command,omitempty\"`\n\tReason    string   `json:\"reason\"`\n\tRecordID  string   `json:\"record_id,omitempty\"`\n\tGoverning []string `json:\"governing,omitempty\"`\n\tDetail    string   `json:\"detail,omitempty\"`\n}","file":"cmd/provenance/next_run.go","name":"nextAction"}}
type nextAction struct {
	Kind      string   `json:"kind"`
	Command   string   `json:"command,omitempty"`
	Reason    string   `json:"reason"`
	RecordID  string   `json:"record_id,omitempty"`
	Governing []string `json:"governing,omitempty"`
	Detail    string   `json:"detail,omitempty"`
}

// /sedum:setType
// sedum:setType {"tier":"owned","record":"prov-2026-11876813","gen":"sha256:090b658e8472ed58a1fb75e4a4101e7f16690a1a70b99759f485f9af7dfaeb18","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type nextState struct {\n\tRecords           []*record.Record\n\tStagedFiles       []string\n\tChangedFiles      []string\n\tIntendedFiles     []string\n\tCommitTagRequired bool\n\tStale             int\n}","file":"cmd/provenance/next_run.go","name":"nextState"}}
type nextState struct {
	Records           []*record.Record
	StagedFiles       []string
	ChangedFiles      []string
	IntendedFiles     []string
	CommitTagRequired bool
	Stale             int
}

// /sedum:setType
// sedum:setFunction {"tier":"owned","record":"prov-2026-11876813","gen":"sha256:090b658e8472ed58a1fb75e4a4101e7f16690a1a70b99759f485f9af7dfaeb18","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func nextTopLevel(rec *record.Record) bool {\n\treturn rec.EffectiveType() != record.TypeImprint\n}","file":"cmd/provenance/next_run.go","name":"nextTopLevel"}}
func nextTopLevel(rec *record.Record) bool {
	return rec.EffectiveType() != record.TypeImprint
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-11876813","gen":"sha256:090b658e8472ed58a1fb75e4a4101e7f16690a1a70b99759f485f9af7dfaeb18","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func nextGoverns(rec *record.Record, file string) bool {\n\tif scope.Mode(rec) != \"allowlist\" {\n\t\treturn false\n\t}\n\tok, err := scope.InScope(rec, file)\n\treturn err == nil && ok\n}","file":"cmd/provenance/next_run.go","name":"nextGoverns"}}
func nextGoverns(rec *record.Record, file string) bool {
	if scope.Mode(rec) != "allowlist" {
		return false
	}
	ok, err := scope.InScope(rec, file)
	return err == nil && ok
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-11876813","gen":"sha256:090b658e8472ed58a1fb75e4a4101e7f16690a1a70b99759f485f9af7dfaeb18","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func nextDedupe(in []string) []string {\n\tseen := map[string]bool{}\n\tvar out []string\n\tfor _, s := range in {\n\t\tif !seen[s] {\n\t\t\tseen[s] = true\n\t\t\tout = append(out, s)\n\t\t}\n\t}\n\treturn out\n}","file":"cmd/provenance/next_run.go","name":"nextDedupe"}}
func nextDedupe(in []string) []string {
	seen := map[string]bool{}
	var out []string
	for _, s := range in {
		if !seen[s] {
			seen[s] = true
			out = append(out, s)
		}
	}
	return out
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-11876813","gen":"sha256:090b658e8472ed58a1fb75e4a4101e7f16690a1a70b99759f485f9af7dfaeb18","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func nextFiles(st nextState) []string {\n\tif len(st.IntendedFiles) > 0 {\n\t\treturn nextDedupe(st.IntendedFiles)\n\t}\n\treturn nextDedupe(append(append([]string{}, st.StagedFiles...), st.ChangedFiles...))\n}","file":"cmd/provenance/next_run.go","name":"nextFiles"}}
func nextFiles(st nextState) []string {
	if len(st.IntendedFiles) > 0 {
		return nextDedupe(st.IntendedFiles)
	}
	return nextDedupe(append(append([]string{}, st.StagedFiles...), st.ChangedFiles...))
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-11876813","gen":"sha256:090b658e8472ed58a1fb75e4a4101e7f16690a1a70b99759f485f9af7dfaeb18","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func nextByStatus(recs []*record.Record, status record.Status) []*record.Record {\n\tvar out []*record.Record\n\tfor _, rec := range recs {\n\t\tif rec.Status == status {\n\t\t\tout = append(out, rec)\n\t\t}\n\t}\n\treturn out\n}","file":"cmd/provenance/next_run.go","name":"nextByStatus"}}
func nextByStatus(recs []*record.Record, status record.Status) []*record.Record {
	var out []*record.Record
	for _, rec := range recs {
		if rec.Status == status {
			out = append(out, rec)
		}
	}
	return out
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-11876813","gen":"sha256:090b658e8472ed58a1fb75e4a4101e7f16690a1a70b99759f485f9af7dfaeb18","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func nextIDs(recs []*record.Record) []string {\n\tids := make([]string, 0, len(recs))\n\tfor _, rec := range recs {\n\t\tids = append(ids, rec.ID)\n\t}\n\tsort.Strings(ids)\n\treturn ids\n}","file":"cmd/provenance/next_run.go","name":"nextIDs"}}
func nextIDs(recs []*record.Record) []string {
	ids := make([]string, 0, len(recs))
	for _, rec := range recs {
		ids = append(ids, rec.ID)
	}
	sort.Strings(ids)
	return ids
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-11876813","gen":"sha256:090b658e8472ed58a1fb75e4a4101e7f16690a1a70b99759f485f9af7dfaeb18","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func nextGoverning(recs []*record.Record, files []string) []*record.Record {\n\tvar hit []*record.Record\n\tfor _, rec := range recs {\n\t\tfor _, f := range files {\n\t\t\tif nextGoverns(rec, f) {\n\t\t\t\thit = append(hit, rec)\n\t\t\t\tbreak\n\t\t\t}\n\t\t}\n\t}\n\tparents := map[string]bool{}\n\tfor _, rec := range hit {\n\t\tif rec.Implements != \"\" {\n\t\t\tparents[rec.Implements] = true\n\t\t}\n\t}\n\tvar out []*record.Record\n\tfor _, rec := range hit {\n\t\tif !parents[rec.ID] {\n\t\t\tout = append(out, rec)\n\t\t}\n\t}\n\tsort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })\n\treturn out\n}","file":"cmd/provenance/next_run.go","name":"nextGoverning"}}
func nextGoverning(recs []*record.Record, files []string) []*record.Record {
	var hit []*record.Record
	for _, rec := range recs {
		for _, f := range files {
			if nextGoverns(rec, f) {
				hit = append(hit, rec)
				break
			}
		}
	}
	parents := map[string]bool{}
	for _, rec := range hit {
		if rec.Implements != "" {
			parents[rec.Implements] = true
		}
	}
	var out []*record.Record
	for _, rec := range hit {
		if !parents[rec.ID] {
			out = append(out, rec)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })
	return out
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-11876813","gen":"sha256:090b658e8472ed58a1fb75e4a4101e7f16690a1a70b99759f485f9af7dfaeb18","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func nextUncovered(active *record.Record, files []string) []string {\n\tvar out []string\n\tfor _, f := range files {\n\t\tif !nextGoverns(active, f) {\n\t\t\tout = append(out, f)\n\t\t}\n\t}\n\treturn out\n}","file":"cmd/provenance/next_run.go","name":"nextUncovered"}}
func nextUncovered(active *record.Record, files []string) []string {
	var out []string
	for _, f := range files {
		if !nextGoverns(active, f) {
			out = append(out, f)
		}
	}
	return out
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-11876813","gen":"sha256:090b658e8472ed58a1fb75e4a4101e7f16690a1a70b99759f485f9af7dfaeb18","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func nextPendingImprints(recs []*record.Record, blueprintID string) []*record.Record {\n\tvar out []*record.Record\n\tfor _, rec := range recs {\n\t\tif rec.EffectiveType() == record.TypeImprint && rec.Implements == blueprintID && rec.Status != record.StatusImplemented {\n\t\t\tout = append(out, rec)\n\t\t}\n\t}\n\tsort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })\n\treturn out\n}","file":"cmd/provenance/next_run.go","name":"nextPendingImprints"}}
func nextPendingImprints(recs []*record.Record, blueprintID string) []*record.Record {
	var out []*record.Record
	for _, rec := range recs {
		if rec.EffectiveType() == record.TypeImprint && rec.Implements == blueprintID && rec.Status != record.StatusImplemented {
			out = append(out, rec)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })
	return out
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-11876813","gen":"sha256:090b658e8472ed58a1fb75e4a4101e7f16690a1a70b99759f485f9af7dfaeb18","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func nextChoose(recs []*record.Record, files []string) nextAction {\n\tids := nextIDs(recs)\n\twhere := \"the files in play\"\n\tif len(files) > 0 {\n\t\twhere = strings.Join(files, \", \")\n\t}\n\treturn nextAction{Kind: \"choose_record\", Governing: ids, Reason: fmt.Sprintf(\"Multiple open records govern %s: %s. Work under ONE of them and tag your commits with it.\", where, strings.Join(ids, \", \"))}\n}","file":"cmd/provenance/next_run.go","name":"nextChoose"}}
func nextChoose(recs []*record.Record, files []string) nextAction {
	ids := nextIDs(recs)
	where := "the files in play"
	if len(files) > 0 {
		where = strings.Join(files, ", ")
	}
	return nextAction{Kind: "choose_record", Governing: ids, Reason: fmt.Sprintf("Multiple open records govern %s: %s. Work under ONE of them and tag your commits with it.", where, strings.Join(ids, ", "))}
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-11876813","gen":"sha256:090b658e8472ed58a1fb75e4a4101e7f16690a1a70b99759f485f9af7dfaeb18","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func nextOnActive(st nextState, active *record.Record) []nextAction {\n\tif missing := nextUncovered(active, nextFiles(st)); len(missing) > 0 {\n\t\treturn []nextAction{{\n\t\t\tKind:     \"add_scope\",\n\t\t\tRecordID: active.ID,\n\t\t\tCommand:  fmt.Sprintf(`provenance edit %s --add-scope %s --reason \"...\"`, active.ID, strings.Join(missing, \" --add-scope \")),\n\t\t\tReason:   fmt.Sprintf(\"%s not in open record %s's affected_scope: widen it (or run provenance add-scope --record %s for files already committed under it).\", strings.Join(missing, \", \"), active.ID, active.ID),\n\t\t}}\n\t}\n\tif len(active.AssociatedSpecs) == 0 {\n\t\treturn []nextAction{{\n\t\t\tKind:     \"add_spec\",\n\t\t\tRecordID: active.ID,\n\t\t\tReason:   \"Open record \" + active.ID + \" has no associated_specs: add a proof artifact before completing.\",\n\t\t\tDetail:   \"# 1. create the proof file first\\n# 2. add it to the record's associated_specs in \" + active.ID + \".yml:\\nassociated_specs:\\n  - path: path/to/proof_test.go\\n    run_command: go test ./... # or any command\\n# 3. provenance edit \" + active.ID + \" --refresh-specs --reason \\\"add proof\\\"\",\n\t\t}}\n\t}\n\tif len(st.StagedFiles) > 0 {\n\t\treason := \"You have staged changes: commit them tagged with \" + active.ID + \".\"\n\t\tif !st.CommitTagRequired {\n\t\t\treason += \" (the tag is optional in this repo, but recommended.)\"\n\t\t}\n\t\treturn []nextAction{{Kind: \"commit\", RecordID: active.ID, Command: fmt.Sprintf(`git commit -m \"<description> [%s@%d]\"`, active.ID, active.Version), Reason: reason}}\n\t}\n\tif pending := nextPendingImprints(st.Records, active.ID); len(pending) > 0 {\n\t\treturn []nextAction{{Kind: \"implement_imprint\", RecordID: pending[0].ID, Command: \"provenance complete --record \" + pending[0].ID, Reason: \"Imprint \" + pending[0].ID + \" of \" + active.ID + \" is not implemented yet: finish and complete it first.\"}}\n\t}\n\treturn []nextAction{{Kind: \"complete\", RecordID: active.ID, Command: \"provenance complete --record \" + active.ID, Reason: \"Specs are present and no imprints are pending: show proof, then complete \" + active.ID + \".\"}}\n}","file":"cmd/provenance/next_run.go","name":"nextOnActive"}}
func nextOnActive(st nextState, active *record.Record) []nextAction {
	if missing := nextUncovered(active, nextFiles(st)); len(missing) > 0 {
		return []nextAction{{
			Kind:     "add_scope",
			RecordID: active.ID,
			Command:  fmt.Sprintf(`provenance edit %s --add-scope %s --reason "..."`, active.ID, strings.Join(missing, " --add-scope ")),
			Reason:   fmt.Sprintf("%s not in open record %s's affected_scope: widen it (or run provenance add-scope --record %s for files already committed under it).", strings.Join(missing, ", "), active.ID, active.ID),
		}}
	}
	if len(active.AssociatedSpecs) == 0 {
		return []nextAction{{
			Kind:     "add_spec",
			RecordID: active.ID,
			Reason:   "Open record " + active.ID + " has no associated_specs: add a proof artifact before completing.",
			Detail:   "# 1. create the proof file first\n# 2. add it to the record's associated_specs in " + active.ID + ".yml:\nassociated_specs:\n  - path: path/to/proof_test.go\n    run_command: go test ./... # or any command\n# 3. provenance edit " + active.ID + " --refresh-specs --reason \"add proof\"",
		}}
	}
	if len(st.StagedFiles) > 0 {
		reason := "You have staged changes: commit them tagged with " + active.ID + "."
		if !st.CommitTagRequired {
			reason += " (the tag is optional in this repo, but recommended.)"
		}
		return []nextAction{{Kind: "commit", RecordID: active.ID, Command: fmt.Sprintf(`git commit -m "<description> [%s@%d]"`, active.ID, active.Version), Reason: reason}}
	}
	if pending := nextPendingImprints(st.Records, active.ID); len(pending) > 0 {
		return []nextAction{{Kind: "implement_imprint", RecordID: pending[0].ID, Command: "provenance complete --record " + pending[0].ID, Reason: "Imprint " + pending[0].ID + " of " + active.ID + " is not implemented yet: finish and complete it first."}}
	}
	return []nextAction{{Kind: "complete", RecordID: active.ID, Command: "provenance complete --record " + active.ID, Reason: "Specs are present and no imprints are pending: show proof, then complete " + active.ID + "."}}
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-11876813","gen":"sha256:090b658e8472ed58a1fb75e4a4101e7f16690a1a70b99759f485f9af7dfaeb18","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func adviseNext(st nextState) []nextAction {\n\tfiles := nextFiles(st)\n\topen := nextByStatus(st.Records, record.StatusOpen)\n\tvar actions []nextAction\n\tvar active *record.Record\n\tgoverning := nextGoverning(open, files)\n\tswitch {\n\tcase len(governing) == 1:\n\t\tactive = governing[0]\n\tcase len(governing) > 1:\n\t\treturn append([]nextAction{nextChoose(governing, files)}, nextFollowUps(st)...)\n\tcase len(files) == 0:\n\t\tvar top []*record.Record\n\t\tfor _, rec := range open {\n\t\t\tif nextTopLevel(rec) {\n\t\t\t\ttop = append(top, rec)\n\t\t\t}\n\t\t}\n\t\tif len(top) == 1 {\n\t\t\tactive = top[0]\n\t\t} else if len(top) > 1 {\n\t\t\treturn append([]nextAction{nextChoose(top, nil)}, nextFollowUps(st)...)\n\t\t}\n\t}\n\tif active != nil {\n\t\treturn append(nextOnActive(st, active), nextFollowUps(st)...)\n\t}\n\tdrafts := nextByStatus(st.Records, record.StatusDraft)\n\tdraft := nextGoverning(drafts, files)\n\tif len(draft) > 0 {\n\t\tactions = []nextAction{{Kind: \"open\", RecordID: draft[0].ID, Command: \"provenance open --record \" + draft[0].ID, Reason: \"Draft \" + draft[0].ID + \" covers this work: open it to enable scope and spec enforcement.\"}}\n\t} else if len(files) == 0 {\n\t\tactions = []nextAction{{Kind: \"none\", Command: `provenance create --title \"...\" --type blueprint --no-edit`, Reason: \"No work in flight and no changed files. Start a blueprint when you begin a change.\"}}\n\t} else if sealed := nextGoverning(nextByStatus(st.Records, record.StatusImplemented), files); len(sealed) > 0 {\n\t\tactions = []nextAction{{Kind: \"revise\", RecordID: sealed[0].ID, Governing: nextIDs(sealed), Command: fmt.Sprintf(`provenance revise %s --reason \"...\"`, sealed[0].ID), Reason: \"These files are governed by implemented record \" + sealed[0].ID + \": revise it to start a new version, then tag the commit with that version.\"}}\n\t} else {\n\t\tactions = []nextAction{{Kind: \"create\", Command: `provenance create --title \"...\" --type blueprint --no-edit`, Reason: \"No record governs these files: create one blueprint covering them.\"}}\n\t}\n\treturn append(actions, nextFollowUps(st)...)\n}","file":"cmd/provenance/next_run.go","name":"adviseNext"}}
func adviseNext(st nextState) []nextAction {
	files := nextFiles(st)
	open := nextByStatus(st.Records, record.StatusOpen)
	var actions []nextAction
	var active *record.Record
	governing := nextGoverning(open, files)
	switch {
	case len(governing) == 1:
		active = governing[0]
	case len(governing) > 1:
		return append([]nextAction{nextChoose(governing, files)}, nextFollowUps(st)...)
	case len(files) == 0:
		var top []*record.Record
		for _, rec := range open {
			if nextTopLevel(rec) {
				top = append(top, rec)
			}
		}
		if len(top) == 1 {
			active = top[0]
		} else if len(top) > 1 {
			return append([]nextAction{nextChoose(top, nil)}, nextFollowUps(st)...)
		}
	}
	if active != nil {
		return append(nextOnActive(st, active), nextFollowUps(st)...)
	}
	drafts := nextByStatus(st.Records, record.StatusDraft)
	draft := nextGoverning(drafts, files)
	if len(draft) > 0 {
		actions = []nextAction{{Kind: "open", RecordID: draft[0].ID, Command: "provenance open --record " + draft[0].ID, Reason: "Draft " + draft[0].ID + " covers this work: open it to enable scope and spec enforcement."}}
	} else if len(files) == 0 {
		actions = []nextAction{{Kind: "none", Command: `provenance create --title "..." --type blueprint --no-edit`, Reason: "No work in flight and no changed files. Start a blueprint when you begin a change."}}
	} else if sealed := nextGoverning(nextByStatus(st.Records, record.StatusImplemented), files); len(sealed) > 0 {
		actions = []nextAction{{Kind: "revise", RecordID: sealed[0].ID, Governing: nextIDs(sealed), Command: fmt.Sprintf(`provenance revise %s --reason "..."`, sealed[0].ID), Reason: "These files are governed by implemented record " + sealed[0].ID + ": revise it to start a new version, then tag the commit with that version."}}
	} else {
		actions = []nextAction{{Kind: "create", Command: `provenance create --title "..." --type blueprint --no-edit`, Reason: "No record governs these files: create one blueprint covering them."}}
	}
	return append(actions, nextFollowUps(st)...)
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-11876813","gen":"sha256:090b658e8472ed58a1fb75e4a4101e7f16690a1a70b99759f485f9af7dfaeb18","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func nextFollowUps(st nextState) []nextAction {\n\tif st.Stale == 0 {\n\t\treturn nil\n\t}\n\treturn []nextAction{{Kind: \"stale\", Command: \"provenance stale\", Reason: fmt.Sprintf(\"%d record(s) are behind their parent: list them with provenance stale, then repin or revise.\", st.Stale)}}\n}","file":"cmd/provenance/next_run.go","name":"nextFollowUps"}}
func nextFollowUps(st nextState) []nextAction {
	if st.Stale == 0 {
		return nil
	}
	return []nextAction{{Kind: "stale", Command: "provenance stale", Reason: fmt.Sprintf("%d record(s) are behind their parent: list them with provenance stale, then repin or revise.", st.Stale)}}
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-11876813","gen":"sha256:090b658e8472ed58a1fb75e4a4101e7f16690a1a70b99759f485f9af7dfaeb18","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func nextRelevant(r *repo, file string) bool {\n\treturn file != \"\" && file != manifestPath && !stateExempt(r, file) && !strings.HasPrefix(file, filepath.ToSlash(r.Config.Dir)+\"/\") && !scope.Excluded(file, r.Config.ExcludePaths)\n}","file":"cmd/provenance/next_run.go","name":"nextRelevant"}}
func nextRelevant(r *repo, file string) bool {
	return file != "" && file != manifestPath && !stateExempt(r, file) && !strings.HasPrefix(file, filepath.ToSlash(r.Config.Dir)+"/") && !scope.Excluded(file, r.Config.ExcludePaths)
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-11876813","gen":"sha256:090b658e8472ed58a1fb75e4a4101e7f16690a1a70b99759f485f9af7dfaeb18","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func nextGitFiles(r *repo, args ...string) []string {\n\tout, err := r.git(args...)\n\tif err != nil {\n\t\treturn nil\n\t}\n\tvar files []string\n\tfor _, f := range strings.Split(out, \"\\n\") {\n\t\tif nextRelevant(r, f) {\n\t\t\tfiles = append(files, f)\n\t\t}\n\t}\n\treturn files\n}","file":"cmd/provenance/next_run.go","name":"nextGitFiles"}}
func nextGitFiles(r *repo, args ...string) []string {
	out, err := r.git(args...)
	if err != nil {
		return nil
	}
	var files []string
	for _, f := range strings.Split(out, "\n") {
		if nextRelevant(r, f) {
			files = append(files, f)
		}
	}
	return files
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-11876813","gen":"sha256:090b658e8472ed58a1fb75e4a4101e7f16690a1a70b99759f485f9af7dfaeb18","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func gatherNext(r *repo, intended []string) (nextState, error) {\n\tst := nextState{IntendedFiles: intended, CommitTagRequired: r.Config.CommitTagRequired}\n\tids, err := r.Store.List()\n\tif err != nil {\n\t\treturn st, err\n\t}\n\tfor _, id := range ids {\n\t\trec, _, err := r.Store.Load(id)\n\t\tif err != nil {\n\t\t\treturn st, err\n\t\t}\n\t\tst.Records = append(st.Records, rec)\n\t}\n\tst.StagedFiles = nextGitFiles(r, \"diff\", \"--cached\", \"--name-only\", \"--no-renames\")\n\tst.ChangedFiles = append(nextGitFiles(r, \"diff\", \"--name-only\", \"HEAD\"), nextGitFiles(r, \"ls-files\", \"--others\", \"--exclude-standard\")...)\n\tst.Stale = len(behindSet(r))\n\treturn st, nil\n}","file":"cmd/provenance/next_run.go","name":"gatherNext"}}
func gatherNext(r *repo, intended []string) (nextState, error) {
	st := nextState{IntendedFiles: intended, CommitTagRequired: r.Config.CommitTagRequired}
	ids, err := r.Store.List()
	if err != nil {
		return st, err
	}
	for _, id := range ids {
		rec, _, err := r.Store.Load(id)
		if err != nil {
			return st, err
		}
		st.Records = append(st.Records, rec)
	}
	st.StagedFiles = nextGitFiles(r, "diff", "--cached", "--name-only", "--no-renames")
	st.ChangedFiles = append(nextGitFiles(r, "diff", "--name-only", "HEAD"), nextGitFiles(r, "ls-files", "--others", "--exclude-standard")...)
	st.Stale = len(behindSet(r))
	return st, nil
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-11876813","gen":"sha256:090b658e8472ed58a1fb75e4a4101e7f16690a1a70b99759f485f9af7dfaeb18","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func printNext(out io.Writer, actions []nextAction) {\n\tfor i, a := range actions {\n\t\tlabel := \"NEXT\"\n\t\tif i > 0 {\n\t\t\tlabel = \"ALSO\"\n\t\t}\n\t\tfmt.Fprintf(out, \"\\n%s: %s\\n  %s\\n\", label, a.Kind, a.Reason)\n\t\tif a.Command != \"\" {\n\t\t\tfmt.Fprintf(out, \"  Run: %s\\n\", a.Command)\n\t\t}\n\t\tif a.Detail != \"\" {\n\t\t\tfmt.Fprint(out, \"\\n\")\n\t\t\tfor _, line := range strings.Split(a.Detail, \"\\n\") {\n\t\t\t\tfmt.Fprintf(out, \"    %s\\n\", line)\n\t\t\t}\n\t\t}\n\t}\n\tfmt.Fprint(out, \"\\n\")\n}","file":"cmd/provenance/next_run.go","name":"printNext"}}
func printNext(out io.Writer, actions []nextAction) {
	for i, a := range actions {
		label := "NEXT"
		if i > 0 {
			label = "ALSO"
		}
		fmt.Fprintf(out, "\n%s: %s\n  %s\n", label, a.Kind, a.Reason)
		if a.Command != "" {
			fmt.Fprintf(out, "  Run: %s\n", a.Command)
		}
		if a.Detail != "" {
			fmt.Fprint(out, "\n")
			for _, line := range strings.Split(a.Detail, "\n") {
				fmt.Fprintf(out, "    %s\n", line)
			}
		}
	}
	fmt.Fprint(out, "\n")
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-11876813","gen":"sha256:090b658e8472ed58a1fb75e4a4101e7f16690a1a70b99759f485f9af7dfaeb18","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func runNext(out io.Writer, cmd *cobra.Command, args []string) error {\n\tr, err := openRepo(cmd)\n\tif err != nil {\n\t\treturn err\n\t}\n\tfiles := splitFileList(mustStringArray(cmd, \"files\"))\n\tplan := splitFileList(mustStringArray(cmd, \"plan\"))\n\tasJSON, _ := cmd.Flags().GetBool(\"json\")\n\tformat, _ := cmd.Flags().GetString(\"format\")\n\tif format != \"\" && format != \"human\" && format != \"json\" {\n\t\treturn fmt.Errorf(\"unknown format %q (human|json)\", format)\n\t}\n\tintended := append(append(append([]string{}, files...), plan...), args...)\n\tst, err := gatherNext(r, intended)\n\tif err != nil {\n\t\treturn err\n\t}\n\tactions := adviseNext(st)\n\tif asJSON || format == \"json\" {\n\t\tenc := json.NewEncoder(out)\n\t\tenc.SetIndent(\"\", \"  \")\n\t\treturn enc.Encode(map[string]any{\"actions\": actions})\n\t}\n\tprintNext(out, actions)\n\treturn nil\n}","file":"cmd/provenance/next_run.go","name":"runNext"}}
func runNext(out io.Writer, cmd *cobra.Command, args []string) error {
	r, err := openRepo(cmd)
	if err != nil {
		return err
	}
	files := splitFileList(mustStringArray(cmd, "files"))
	plan := splitFileList(mustStringArray(cmd, "plan"))
	asJSON, _ := cmd.Flags().GetBool("json")
	format, _ := cmd.Flags().GetString("format")
	if format != "" && format != "human" && format != "json" {
		return fmt.Errorf("unknown format %q (human|json)", format)
	}
	intended := append(append(append([]string{}, files...), plan...), args...)
	st, err := gatherNext(r, intended)
	if err != nil {
		return err
	}
	actions := adviseNext(st)
	if asJSON || format == "json" {
		enc := json.NewEncoder(out)
		enc.SetIndent("", "  ")
		return enc.Encode(map[string]any{"actions": actions})
	}
	printNext(out, actions)
	return nil
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-11876813","gen":"sha256:090b658e8472ed58a1fb75e4a4101e7f16690a1a70b99759f485f9af7dfaeb18","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func mustStringArray(cmd *cobra.Command, name string) []string {\n\tvalues, _ := cmd.Flags().GetStringArray(name)\n\treturn values\n}","file":"cmd/provenance/next_run.go","name":"mustStringArray"}}
func mustStringArray(cmd *cobra.Command, name string) []string {
	values, _ := cmd.Flags().GetStringArray(name)
	return values
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-11876813","gen":"sha256:090b658e8472ed58a1fb75e4a4101e7f16690a1a70b99759f485f9af7dfaeb18","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func splitFileList(items []string) []string {\n\tvar out []string\n\tfor _, item := range items {\n\t\tfor _, f := range strings.Split(item, \",\") {\n\t\t\tif f = strings.TrimSpace(f); f != \"\" {\n\t\t\t\tout = append(out, f)\n\t\t\t}\n\t\t}\n\t}\n\treturn out\n}","file":"cmd/provenance/next_run.go","name":"splitFileList"}}
func splitFileList(items []string) []string {
	var out []string
	for _, item := range items {
		for _, f := range strings.Split(item, ",") {
			if f = strings.TrimSpace(f); f != "" {
				out = append(out, f)
			}
		}
	}
	return out
}

// /sedum:setFunction
