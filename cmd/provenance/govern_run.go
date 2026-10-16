package main

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-9ec4a843","gen":"sha256:cc0301a10fcd903006e4f2bed32a32b6a082f192eb2c42bb2c80462870e50daf","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/govern_run.go","path":"encoding/json"}}
import "encoding/json"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-9ec4a843","gen":"sha256:cc0301a10fcd903006e4f2bed32a32b6a082f192eb2c42bb2c80462870e50daf","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/govern_run.go","path":"fmt"}}
import "fmt"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-9ec4a843","gen":"sha256:cc0301a10fcd903006e4f2bed32a32b6a082f192eb2c42bb2c80462870e50daf","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/govern_run.go","path":"io"}}
import "io"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-9ec4a843","gen":"sha256:cc0301a10fcd903006e4f2bed32a32b6a082f192eb2c42bb2c80462870e50daf","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/govern_run.go","path":"sort"}}
import "sort"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-9ec4a843","gen":"sha256:cc0301a10fcd903006e4f2bed32a32b6a082f192eb2c42bb2c80462870e50daf","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/govern_run.go","path":"strings"}}
import "strings"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-9ec4a843","gen":"sha256:cc0301a10fcd903006e4f2bed32a32b6a082f192eb2c42bb2c80462870e50daf","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/govern_run.go","path":"github.com/spf13/cobra"}}
import "github.com/spf13/cobra"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-9ec4a843","gen":"sha256:cc0301a10fcd903006e4f2bed32a32b6a082f192eb2c42bb2c80462870e50daf","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/govern_run.go","path":"github.com/livecodelife/provenance/record"}}
import "github.com/livecodelife/provenance/record"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setType {"tier":"owned","record":"prov-2026-9ec4a843","gen":"sha256:cc0301a10fcd903006e4f2bed32a32b6a082f192eb2c42bb2c80462870e50daf","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type governRecord struct {\n\tID      string   `json:\"id\"`\n\tType    string   `json:\"type\"`\n\tStatus  string   `json:\"status\"`\n\tVersion int      `json:\"version\"`\n\tTitle   string   `json:\"title\"`\n\tFiles   []string `json:\"files\"`\n}","file":"cmd/provenance/govern_run.go","name":"governRecord"}}
type governRecord struct {
	ID      string   `json:"id"`
	Type    string   `json:"type"`
	Status  string   `json:"status"`
	Version int      `json:"version"`
	Title   string   `json:"title"`
	Files   []string `json:"files"`
}

// /sedum:setType
// sedum:setType {"tier":"owned","record":"prov-2026-9ec4a843","gen":"sha256:cc0301a10fcd903006e4f2bed32a32b6a082f192eb2c42bb2c80462870e50daf","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type governDoc struct {\n\tFiles     []string       `json:\"files\"`\n\tExempt    []string       `json:\"exempt\"`\n\tGoverning []governRecord `json:\"governing\"`\n\tNext      *nextAction    `json:\"next,omitempty\"`\n}","file":"cmd/provenance/govern_run.go","name":"governDoc"}}
type governDoc struct {
	Files     []string       `json:"files"`
	Exempt    []string       `json:"exempt"`
	Governing []governRecord `json:"governing"`
	Next      *nextAction    `json:"next,omitempty"`
}

// /sedum:setType
// sedum:setFunction {"tier":"owned","record":"prov-2026-9ec4a843","gen":"sha256:cc0301a10fcd903006e4f2bed32a32b6a082f192eb2c42bb2c80462870e50daf","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func buildGovern(r *repo, files []string) (governDoc, error) {\n\tdoc := governDoc{Files: files, Exempt: []string{}, Governing: []governRecord{}}\n\tids, err := r.Store.List()\n\tif err != nil {\n\t\treturn doc, err\n\t}\n\tvar active []*record.Record\n\tfor _, id := range ids {\n\t\trec, _, err := r.Store.Load(id)\n\t\tif err != nil {\n\t\t\treturn doc, err\n\t\t}\n\t\tif rec.Status == record.StatusOpen || rec.Status == record.StatusImplemented {\n\t\t\tactive = append(active, rec)\n\t\t}\n\t}\n\tsort.Slice(active, func(i, j int) bool { return active[i].ID < active[j].ID })\n\tcovers := map[string][]string{}\n\tfor _, file := range files {\n\t\tpath := strings.TrimPrefix(file, \"./\")\n\t\tif !nextRelevant(r, path) {\n\t\t\tdoc.Exempt = append(doc.Exempt, file)\n\t\t\tcontinue\n\t\t}\n\t\tfor _, rec := range active {\n\t\t\tif nextGoverns(rec, path) {\n\t\t\t\tcovers[rec.ID] = append(covers[rec.ID], file)\n\t\t\t}\n\t\t}\n\t}\n\tfor _, rec := range active {\n\t\tif got := covers[rec.ID]; len(got) > 0 {\n\t\t\tdoc.Governing = append(doc.Governing, governRecord{ID: rec.ID, Type: string(rec.EffectiveType()), Status: string(rec.Status), Version: rec.Version, Title: rec.Title, Files: got})\n\t\t}\n\t}\n\tst, err := gatherNext(r, files)\n\tif err != nil {\n\t\treturn doc, err\n\t}\n\tif actions := adviseNext(st); len(actions) > 0 {\n\t\tdoc.Next = &actions[0]\n\t}\n\treturn doc, nil\n}","file":"cmd/provenance/govern_run.go","name":"buildGovern"}}
func buildGovern(r *repo, files []string) (governDoc, error) {
	doc := governDoc{Files: files, Exempt: []string{}, Governing: []governRecord{}}
	ids, err := r.Store.List()
	if err != nil {
		return doc, err
	}
	var active []*record.Record
	for _, id := range ids {
		rec, _, err := r.Store.Load(id)
		if err != nil {
			return doc, err
		}
		if rec.Status == record.StatusOpen || rec.Status == record.StatusImplemented {
			active = append(active, rec)
		}
	}
	sort.Slice(active, func(i, j int) bool { return active[i].ID < active[j].ID })
	covers := map[string][]string{}
	for _, file := range files {
		path := strings.TrimPrefix(file, "./")
		if !nextRelevant(r, path) {
			doc.Exempt = append(doc.Exempt, file)
			continue
		}
		for _, rec := range active {
			if nextGoverns(rec, path) {
				covers[rec.ID] = append(covers[rec.ID], file)
			}
		}
	}
	for _, rec := range active {
		if got := covers[rec.ID]; len(got) > 0 {
			doc.Governing = append(doc.Governing, governRecord{ID: rec.ID, Type: string(rec.EffectiveType()), Status: string(rec.Status), Version: rec.Version, Title: rec.Title, Files: got})
		}
	}
	st, err := gatherNext(r, files)
	if err != nil {
		return doc, err
	}
	if actions := adviseNext(st); len(actions) > 0 {
		doc.Next = &actions[0]
	}
	return doc, nil
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-9ec4a843","gen":"sha256:cc0301a10fcd903006e4f2bed32a32b6a082f192eb2c42bb2c80462870e50daf","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func printGovern(out io.Writer, doc governDoc) {\n\tfmt.Fprintf(out, \"\\nGOVERN: %s\\n\", ctxClean(strings.Join(doc.Files, \", \"), 300))\n\tif len(doc.Governing) == 0 {\n\t\tfmt.Fprint(out, \"  no active record governs these files\\n\")\n\t}\n\tfor _, g := range doc.Governing {\n\t\tfmt.Fprintf(out, \"  %s@%d  %s  %s  %s  (%s)\\n\", g.ID, g.Version, g.Status, g.Type, ctxClean(g.Title, 100), ctxClean(strings.Join(g.Files, \", \"), 200))\n\t}\n\tif len(doc.Exempt) > 0 {\n\t\tfmt.Fprintf(out, \"  exempt: %s\\n\", ctxClean(strings.Join(doc.Exempt, \", \"), 300))\n\t}\n\tif doc.Next != nil {\n\t\tfmt.Fprintf(out, \"  Next: %s: %s\\n\", doc.Next.Kind, ctxClean(doc.Next.Reason, 300))\n\t\tif doc.Next.Command != \"\" {\n\t\t\tfmt.Fprintf(out, \"    Run: %s\\n\", ctxClean(doc.Next.Command, 300))\n\t\t}\n\t}\n\tfmt.Fprint(out, \"\\n\")\n}","file":"cmd/provenance/govern_run.go","name":"printGovern"}}
func printGovern(out io.Writer, doc governDoc) {
	fmt.Fprintf(out, "\nGOVERN: %s\n", ctxClean(strings.Join(doc.Files, ", "), 300))
	if len(doc.Governing) == 0 {
		fmt.Fprint(out, "  no active record governs these files\n")
	}
	for _, g := range doc.Governing {
		fmt.Fprintf(out, "  %s@%d  %s  %s  %s  (%s)\n", g.ID, g.Version, g.Status, g.Type, ctxClean(g.Title, 100), ctxClean(strings.Join(g.Files, ", "), 200))
	}
	if len(doc.Exempt) > 0 {
		fmt.Fprintf(out, "  exempt: %s\n", ctxClean(strings.Join(doc.Exempt, ", "), 300))
	}
	if doc.Next != nil {
		fmt.Fprintf(out, "  Next: %s: %s\n", doc.Next.Kind, ctxClean(doc.Next.Reason, 300))
		if doc.Next.Command != "" {
			fmt.Fprintf(out, "    Run: %s\n", ctxClean(doc.Next.Command, 300))
		}
	}
	fmt.Fprint(out, "\n")
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-9ec4a843","gen":"sha256:cc0301a10fcd903006e4f2bed32a32b6a082f192eb2c42bb2c80462870e50daf","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func runGovern(out io.Writer, cmd *cobra.Command, args []string) error {\n\tr, err := openRepo(cmd)\n\tif err != nil {\n\t\treturn err\n\t}\n\tfiles := splitFileList(mustStringArray(cmd, \"files\"))\n\tformat, _ := cmd.Flags().GetString(\"format\")\n\tasJSON, _ := cmd.Flags().GetBool(\"json\")\n\tif format != \"\" && format != \"human\" && format != \"json\" {\n\t\treturn fmt.Errorf(\"unknown format %q (human|json)\", format)\n\t}\n\tgiven := nextDedupe(append(append([]string{}, args...), files...))\n\tif len(given) == 0 {\n\t\treturn fmt.Errorf(\"no files: give file paths as arguments or use --files\")\n\t}\n\tdoc, err := buildGovern(r, given)\n\tif err != nil {\n\t\treturn err\n\t}\n\tif asJSON || format == \"json\" {\n\t\tenc := json.NewEncoder(out)\n\t\tenc.SetIndent(\"\", \"  \")\n\t\treturn enc.Encode(doc)\n\t}\n\tprintGovern(out, doc)\n\treturn nil\n}","file":"cmd/provenance/govern_run.go","name":"runGovern"}}
func runGovern(out io.Writer, cmd *cobra.Command, args []string) error {
	r, err := openRepo(cmd)
	if err != nil {
		return err
	}
	files := splitFileList(mustStringArray(cmd, "files"))
	format, _ := cmd.Flags().GetString("format")
	asJSON, _ := cmd.Flags().GetBool("json")
	if format != "" && format != "human" && format != "json" {
		return fmt.Errorf("unknown format %q (human|json)", format)
	}
	given := nextDedupe(append(append([]string{}, args...), files...))
	if len(given) == 0 {
		return fmt.Errorf("no files: give file paths as arguments or use --files")
	}
	doc, err := buildGovern(r, given)
	if err != nil {
		return err
	}
	if asJSON || format == "json" {
		enc := json.NewEncoder(out)
		enc.SetIndent("", "  ")
		return enc.Encode(doc)
	}
	printGovern(out, doc)
	return nil
}

// /sedum:setFunction
