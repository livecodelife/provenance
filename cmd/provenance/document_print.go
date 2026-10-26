package main

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-29e9b1b1","gen":"sha256:2e8980f0d01f958ea0e637515d0e79bca72234b95c5792c8e21bf54452d8760f","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/document_print.go","path":"fmt"}}
import "fmt"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-29e9b1b1","gen":"sha256:2e8980f0d01f958ea0e637515d0e79bca72234b95c5792c8e21bf54452d8760f","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/document_print.go","path":"io"}}
import "io"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-29e9b1b1","gen":"sha256:2e8980f0d01f958ea0e637515d0e79bca72234b95c5792c8e21bf54452d8760f","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/document_print.go","path":"sort"}}
import "sort"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-29e9b1b1","gen":"sha256:2e8980f0d01f958ea0e637515d0e79bca72234b95c5792c8e21bf54452d8760f","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/document_print.go","path":"github.com/livecodelife/provenance/document"}}
import "github.com/livecodelife/provenance/document"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"tier":"owned","record":"prov-2026-29e9b1b1","gen":"sha256:2e8980f0d01f958ea0e637515d0e79bca72234b95c5792c8e21bf54452d8760f","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func printDocumentPlan(w io.Writer, plan *document.Plan) {\n\tif plan == nil {\n\t\treturn\n\t}\n\ttype entry struct {\n\t\tkey  string\n\t\tcat  int\n\t\tline int\n\t\ttext string\n\t}\n\tvar entries []entry\n\tadd := func(id string, cat, line int, rest string) {\n\t\tkey, display := id, id\n\t\tif id == \"\" {\n\t\t\tdisplay = \"-\"\n\t\t}\n\t\tentries = append(entries, entry{key: key, cat: cat, line: line, text: display + \" \" + rest})\n\t}\n\tfor _, c := range plan.Changes {\n\t\tdetail := \"intent\"\n\t\tswitch {\n\t\tcase c.IntentChanged && c.ConstraintsChanged:\n\t\t\tdetail = \"intent and constraints\"\n\t\tcase c.ConstraintsChanged:\n\t\t\tdetail = \"constraints\"\n\t\t}\n\t\tadd(c.ID, 0, 0, \"change \"+detail)\n\t}\n\tfor _, c := range plan.Conflicts {\n\t\tadd(c.ID, 1, c.Line, fmt.Sprintf(\"conflict anchored v%d current v%d\", c.AnchoredVersion, c.CurrentVersion))\n\t}\n\tfor _, p := range plan.Problems {\n\t\tadd(p.ID, 2, p.Line, fmt.Sprintf(\"%s line %d: %s\", p.Kind, p.Line, p.Message))\n\t}\n\tsort.SliceStable(entries, func(i, j int) bool {\n\t\tif entries[i].key != entries[j].key {\n\t\t\treturn entries[i].key < entries[j].key\n\t\t}\n\t\tif entries[i].cat != entries[j].cat {\n\t\t\treturn entries[i].cat < entries[j].cat\n\t\t}\n\t\treturn entries[i].line < entries[j].line\n\t})\n\tfor _, e := range entries {\n\t\tfmt.Fprintln(w, e.text)\n\t}\n\tfmt.Fprintf(w, \"%d changed, %d conflicts, %d problems, %d unchanged\\n\", len(plan.Changes), len(plan.Conflicts), len(plan.Problems), len(plan.Unchanged))\n}","file":"cmd/provenance/document_print.go","name":"printDocumentPlan"}}
func printDocumentPlan(w io.Writer, plan *document.Plan) {
	if plan == nil {
		return
	}
	type entry struct {
		key  string
		cat  int
		line int
		text string
	}
	var entries []entry
	add := func(id string, cat, line int, rest string) {
		key, display := id, id
		if id == "" {
			display = "-"
		}
		entries = append(entries, entry{key: key, cat: cat, line: line, text: display + " " + rest})
	}
	for _, c := range plan.Changes {
		detail := "intent"
		switch {
		case c.IntentChanged && c.ConstraintsChanged:
			detail = "intent and constraints"
		case c.ConstraintsChanged:
			detail = "constraints"
		}
		add(c.ID, 0, 0, "change "+detail)
	}
	for _, c := range plan.Conflicts {
		add(c.ID, 1, c.Line, fmt.Sprintf("conflict anchored v%d current v%d", c.AnchoredVersion, c.CurrentVersion))
	}
	for _, p := range plan.Problems {
		add(p.ID, 2, p.Line, fmt.Sprintf("%s line %d: %s", p.Kind, p.Line, p.Message))
	}
	sort.SliceStable(entries, func(i, j int) bool {
		if entries[i].key != entries[j].key {
			return entries[i].key < entries[j].key
		}
		if entries[i].cat != entries[j].cat {
			return entries[i].cat < entries[j].cat
		}
		return entries[i].line < entries[j].line
	})
	for _, e := range entries {
		fmt.Fprintln(w, e.text)
	}
	fmt.Fprintf(w, "%d changed, %d conflicts, %d problems, %d unchanged\n", len(plan.Changes), len(plan.Conflicts), len(plan.Problems), len(plan.Unchanged))
}

// /sedum:setFunction
