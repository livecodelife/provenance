package main

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-dce50c89","gen":"sha256:86f4fd0e90b200b90b4a7bc131e2d6070e9dc6fff73199149ab0ca54592d190c","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/evidence_run.go","path":"encoding/json"}}
import "encoding/json"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-dce50c89","gen":"sha256:86f4fd0e90b200b90b4a7bc131e2d6070e9dc6fff73199149ab0ca54592d190c","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/evidence_run.go","path":"fmt"}}
import "fmt"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-dce50c89","gen":"sha256:86f4fd0e90b200b90b4a7bc131e2d6070e9dc6fff73199149ab0ca54592d190c","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/evidence_run.go","path":"io"}}
import "io"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-dce50c89","gen":"sha256:86f4fd0e90b200b90b4a7bc131e2d6070e9dc6fff73199149ab0ca54592d190c","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/evidence_run.go","path":"os"}}
import "os"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-dce50c89","gen":"sha256:86f4fd0e90b200b90b4a7bc131e2d6070e9dc6fff73199149ab0ca54592d190c","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/evidence_run.go","path":"os/exec"}}
import "os/exec"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-dce50c89","gen":"sha256:86f4fd0e90b200b90b4a7bc131e2d6070e9dc6fff73199149ab0ca54592d190c","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/evidence_run.go","path":"path/filepath"}}
import "path/filepath"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-dce50c89","gen":"sha256:86f4fd0e90b200b90b4a7bc131e2d6070e9dc6fff73199149ab0ca54592d190c","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/evidence_run.go","path":"time"}}
import "time"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-dce50c89","gen":"sha256:86f4fd0e90b200b90b4a7bc131e2d6070e9dc6fff73199149ab0ca54592d190c","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/evidence_run.go","path":"github.com/spf13/cobra"}}
import "github.com/spf13/cobra"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-dce50c89","gen":"sha256:86f4fd0e90b200b90b4a7bc131e2d6070e9dc6fff73199149ab0ca54592d190c","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/evidence_run.go","path":"github.com/livecodelife/provenance/evidence"}}
import "github.com/livecodelife/provenance/evidence"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-dce50c89","gen":"sha256:86f4fd0e90b200b90b4a7bc131e2d6070e9dc6fff73199149ab0ca54592d190c","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/evidence_run.go","path":"github.com/livecodelife/provenance/record"}}
import "github.com/livecodelife/provenance/record"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-dce50c89","gen":"sha256:86f4fd0e90b200b90b4a7bc131e2d6070e9dc6fff73199149ab0ca54592d190c","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/evidence_run.go","path":"github.com/livecodelife/provenance/stale"}}
import "github.com/livecodelife/provenance/stale"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setType {"tier":"owned","record":"prov-2026-dce50c89","gen":"sha256:86f4fd0e90b200b90b4a7bc131e2d6070e9dc6fff73199149ab0ca54592d190c","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type evidenceRow struct {\n\tID      string   `json:\"id\"`\n\tVersion int      `json:\"version\"`\n\tState   string   `json:\"state\"`\n\tPassed  bool     `json:\"passed\"`\n\tReasons []string `json:\"reasons,omitempty\"`\n}","file":"cmd/provenance/evidence_run.go","name":"evidenceRow"}}
type evidenceRow struct {
	ID      string   `json:"id"`
	Version int      `json:"version"`
	State   string   `json:"state"`
	Passed  bool     `json:"passed"`
	Reasons []string `json:"reasons,omitempty"`
}

// /sedum:setType
// sedum:setFunction {"tier":"owned","record":"prov-2026-dce50c89","gen":"sha256:86f4fd0e90b200b90b4a7bc131e2d6070e9dc6fff73199149ab0ca54592d190c","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func sedumRegions(r *repo, id string) ([]evidence.Region, bool) {\n\tgen := filepath.Join(r.Root, \"generators\")\n\tif _, err := os.Stat(gen); err != nil {\n\t\tempty, err := os.MkdirTemp(\"\", \"provenance-nogenerators-\")\n\t\tif err != nil {\n\t\t\treturn nil, false\n\t\t}\n\t\tdefer os.RemoveAll(empty)\n\t\tgen = empty\n\t}\n\tbin := os.Getenv(\"SEDUM_BIN\")\n\tif bin == \"\" {\n\t\tbin = \"sedum\"\n\t}\n\tc := exec.Command(bin, \"regions\", \"--json\", \"--generators\", gen, \"--records\", filepath.Join(r.Root, r.Config.Dir), \"--record\", id)\n\tc.Dir = r.Root\n\tout, err := c.Output()\n\tif err != nil {\n\t\treturn nil, false\n\t}\n\tvar doc struct {\n\t\tRegions []evidence.Region `json:\"regions\"`\n\t}\n\tif json.Unmarshal(out, &doc) != nil {\n\t\treturn nil, false\n\t}\n\treturn doc.Regions, true\n}","file":"cmd/provenance/evidence_run.go","name":"sedumRegions"}}
func sedumRegions(r *repo, id string) ([]evidence.Region, bool) {
	gen := filepath.Join(r.Root, "generators")
	if _, err := os.Stat(gen); err != nil {
		empty, err := os.MkdirTemp("", "provenance-nogenerators-")
		if err != nil {
			return nil, false
		}
		defer os.RemoveAll(empty)
		gen = empty
	}
	bin := os.Getenv("SEDUM_BIN")
	if bin == "" {
		bin = "sedum"
	}
	c := exec.Command(bin, "regions", "--json", "--generators", gen, "--records", filepath.Join(r.Root, r.Config.Dir), "--record", id)
	c.Dir = r.Root
	out, err := c.Output()
	if err != nil {
		return nil, false
	}
	var doc struct {
		Regions []evidence.Region `json:"regions"`
	}
	if json.Unmarshal(out, &doc) != nil {
		return nil, false
	}
	return doc.Regions, true
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-dce50c89","gen":"sha256:86f4fd0e90b200b90b4a7bc131e2d6070e9dc6fff73199149ab0ca54592d190c","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func behindSet(r *repo) map[string]bool {\n\tset := map[string]bool{}\n\tfound, err := stale.Find(r.Store)\n\tif err != nil {\n\t\treturn set\n\t}\n\tfor _, f := range found {\n\t\tif f.Kind != stale.Unknown {\n\t\t\tset[f.ID] = true\n\t\t}\n\t}\n\treturn set\n}","file":"cmd/provenance/evidence_run.go","name":"behindSet"}}
func behindSet(r *repo) map[string]bool {
	set := map[string]bool{}
	found, err := stale.Find(r.Store)
	if err != nil {
		return set
	}
	for _, f := range found {
		if f.Kind != stale.Unknown {
			set[f.ID] = true
		}
	}
	return set
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-dce50c89","gen":"sha256:86f4fd0e90b200b90b4a7bc131e2d6070e9dc6fff73199149ab0ca54592d190c","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func currentEvidence(r *repo, rec *record.Record, behind map[string]bool) evidence.Current {\n\tregions, ok := sedumRegions(r, rec.ID)\n\tcur := evidence.CurrentFor(rec, r.Store.Specs(), regions)\n\tif !ok {\n\t\tcur.CodeHash = \"\"\n\t}\n\tcur.Behind = behind[rec.ID]\n\treturn cur\n}","file":"cmd/provenance/evidence_run.go","name":"currentEvidence"}}
func currentEvidence(r *repo, rec *record.Record, behind map[string]bool) evidence.Current {
	regions, ok := sedumRegions(r, rec.ID)
	cur := evidence.CurrentFor(rec, r.Store.Specs(), regions)
	if !ok {
		cur.CodeHash = ""
	}
	cur.Behind = behind[rec.ID]
	return cur
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-dce50c89","gen":"sha256:86f4fd0e90b200b90b4a7bc131e2d6070e9dc6fff73199149ab0ca54592d190c","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func evidenceFiles(r *repo, id string) map[string]bool {\n\tm := map[string]bool{}\n\tfiles, _ := filepath.Glob(filepath.Join(r.Root, r.Config.Dir, \".evidence\", id, \"*\"))\n\tfor _, f := range files {\n\t\tm[f] = true\n\t}\n\treturn m\n}","file":"cmd/provenance/evidence_run.go","name":"evidenceFiles"}}
func evidenceFiles(r *repo, id string) map[string]bool {
	m := map[string]bool{}
	files, _ := filepath.Glob(filepath.Join(r.Root, r.Config.Dir, ".evidence", id, "*"))
	for _, f := range files {
		m[f] = true
	}
	return m
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-dce50c89","gen":"sha256:86f4fd0e90b200b90b4a7bc131e2d6070e9dc6fff73199149ab0ca54592d190c","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func recordEvidence(r *repo, rec *record.Record, passed bool) ([]string, error) {\n\tbefore := evidenceFiles(r, rec.ID)\n\tcur := currentEvidence(r, rec, behindSet(r))\n\te := evidence.New(rec, cur, passed, time.Now().UTC().Format(time.RFC3339), \"provenance complete\")\n\tif err := (&evidence.Store{Root: r.Root, Dir: r.Config.Dir}).Append(e); err != nil {\n\t\treturn nil, err\n\t}\n\tvar added []string\n\tfor f := range evidenceFiles(r, rec.ID) {\n\t\tif !before[f] {\n\t\t\tadded = append(added, f)\n\t\t}\n\t}\n\treturn added, nil\n}","file":"cmd/provenance/evidence_run.go","name":"recordEvidence"}}
func recordEvidence(r *repo, rec *record.Record, passed bool) ([]string, error) {
	before := evidenceFiles(r, rec.ID)
	cur := currentEvidence(r, rec, behindSet(r))
	e := evidence.New(rec, cur, passed, time.Now().UTC().Format(time.RFC3339), "provenance complete")
	if err := (&evidence.Store{Root: r.Root, Dir: r.Config.Dir}).Append(e); err != nil {
		return nil, err
	}
	var added []string
	for f := range evidenceFiles(r, rec.ID) {
		if !before[f] {
			added = append(added, f)
		}
	}
	return added, nil
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-dce50c89","gen":"sha256:86f4fd0e90b200b90b4a7bc131e2d6070e9dc6fff73199149ab0ca54592d190c","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func checkEvidence(r *repo) ([]evidenceRow, error) {\n\tids, err := r.Store.List()\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\tbehind := behindSet(r)\n\tstore := &evidence.Store{Root: r.Root, Dir: r.Config.Dir}\n\trows := []evidenceRow{}\n\tfor _, id := range ids {\n\t\trec, _, err := r.Store.Load(id)\n\t\tif err != nil {\n\t\t\treturn nil, err\n\t\t}\n\t\tif rec.Status != record.StatusImplemented {\n\t\t\tcontinue\n\t\t}\n\t\tv, err := store.Status(rec, currentEvidence(r, rec, behind))\n\t\tif err != nil {\n\t\t\treturn nil, err\n\t\t}\n\t\trows = append(rows, evidenceRow{ID: rec.ID, Version: rec.Version, State: string(v.State), Passed: v.Passed, Reasons: v.Reasons})\n\t}\n\treturn rows, nil\n}","file":"cmd/provenance/evidence_run.go","name":"checkEvidence"}}
func checkEvidence(r *repo) ([]evidenceRow, error) {
	ids, err := r.Store.List()
	if err != nil {
		return nil, err
	}
	behind := behindSet(r)
	store := &evidence.Store{Root: r.Root, Dir: r.Config.Dir}
	rows := []evidenceRow{}
	for _, id := range ids {
		rec, _, err := r.Store.Load(id)
		if err != nil {
			return nil, err
		}
		if rec.Status != record.StatusImplemented {
			continue
		}
		v, err := store.Status(rec, currentEvidence(r, rec, behind))
		if err != nil {
			return nil, err
		}
		rows = append(rows, evidenceRow{ID: rec.ID, Version: rec.Version, State: string(v.State), Passed: v.Passed, Reasons: v.Reasons})
	}
	return rows, nil
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-dce50c89","gen":"sha256:86f4fd0e90b200b90b4a7bc131e2d6070e9dc6fff73199149ab0ca54592d190c","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func evidenceFailing(row evidenceRow) bool {\n\treturn row.State == string(evidence.Stale) || (row.State != string(evidence.Missing) && !row.Passed)\n}","file":"cmd/provenance/evidence_run.go","name":"evidenceFailing"}}
func evidenceFailing(row evidenceRow) bool {
	return row.State == string(evidence.Stale) || (row.State != string(evidence.Missing) && !row.Passed)
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-dce50c89","gen":"sha256:86f4fd0e90b200b90b4a7bc131e2d6070e9dc6fff73199149ab0ca54592d190c","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func printEvidence(out io.Writer, rows []evidenceRow) {\n\tfmt.Fprint(out, \"\\nEVIDENCE\\n\\n\")\n\tcounts := map[string]int{}\n\tfor _, row := range rows {\n\t\tcounts[row.State]++\n\t\twhy := \"\"\n\t\tif len(row.Reasons) > 0 {\n\t\t\twhy = fmt.Sprintf(\"%v\", row.Reasons)\n\t\t}\n\t\tif row.State != string(evidence.Missing) && !row.Passed {\n\t\t\twhy = \"specs failed \" + why\n\t\t}\n\t\tfmt.Fprintf(out, \"  %-22s %-8s %s\\n\", fmt.Sprintf(\"%s@%d\", row.ID, row.Version), row.State, why)\n\t}\n\tfmt.Fprintf(out, \"\\n  %d fresh, %d stale, %d unknown, %d missing\\n\\n\", counts[\"fresh\"], counts[\"stale\"], counts[\"unknown\"], counts[\"missing\"])\n}","file":"cmd/provenance/evidence_run.go","name":"printEvidence"}}
func printEvidence(out io.Writer, rows []evidenceRow) {
	fmt.Fprint(out, "\nEVIDENCE\n\n")
	counts := map[string]int{}
	for _, row := range rows {
		counts[row.State]++
		why := ""
		if len(row.Reasons) > 0 {
			why = fmt.Sprintf("%v", row.Reasons)
		}
		if row.State != string(evidence.Missing) && !row.Passed {
			why = "specs failed " + why
		}
		fmt.Fprintf(out, "  %-22s %-8s %s\n", fmt.Sprintf("%s@%d", row.ID, row.Version), row.State, why)
	}
	fmt.Fprintf(out, "\n  %d fresh, %d stale, %d unknown, %d missing\n\n", counts["fresh"], counts["stale"], counts["unknown"], counts["missing"])
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-dce50c89","gen":"sha256:86f4fd0e90b200b90b4a7bc131e2d6070e9dc6fff73199149ab0ca54592d190c","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func runCheckEvidence(out io.Writer, cmd *cobra.Command, r *repo) error {\n\trows, err := checkEvidence(r)\n\tif err != nil {\n\t\treturn err\n\t}\n\tasJSON, _ := cmd.Flags().GetBool(\"json\")\n\tif asJSON {\n\t\tenc := json.NewEncoder(out)\n\t\tenc.SetIndent(\"\", \"  \")\n\t\tif err := enc.Encode(map[string]any{\"evidence\": rows}); err != nil {\n\t\t\treturn err\n\t\t}\n\t} else {\n\t\tprintEvidence(out, rows)\n\t}\n\tbad := 0\n\tfor _, row := range rows {\n\t\tif evidenceFailing(row) {\n\t\t\tbad++\n\t\t}\n\t}\n\tif bad > 0 {\n\t\treturn fmt.Errorf(\"%d records have stale or failing evidence\", bad)\n\t}\n\treturn nil\n}","file":"cmd/provenance/evidence_run.go","name":"runCheckEvidence"}}
func runCheckEvidence(out io.Writer, cmd *cobra.Command, r *repo) error {
	rows, err := checkEvidence(r)
	if err != nil {
		return err
	}
	asJSON, _ := cmd.Flags().GetBool("json")
	if asJSON {
		enc := json.NewEncoder(out)
		enc.SetIndent("", "  ")
		if err := enc.Encode(map[string]any{"evidence": rows}); err != nil {
			return err
		}
	} else {
		printEvidence(out, rows)
	}
	bad := 0
	for _, row := range rows {
		if evidenceFailing(row) {
			bad++
		}
	}
	if bad > 0 {
		return fmt.Errorf("%d records have stale or failing evidence", bad)
	}
	return nil
}

// /sedum:setFunction
