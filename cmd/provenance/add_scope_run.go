package main

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-9c8a0b2a","gen":"sha256:9c65c23d7893e39a1725f1f00b3344e10ee8e1823f7702b3722d3bef7130ba0c","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/add_scope_run.go","path":"errors"}}
import "errors"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-9c8a0b2a","gen":"sha256:9c65c23d7893e39a1725f1f00b3344e10ee8e1823f7702b3722d3bef7130ba0c","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/add_scope_run.go","path":"fmt"}}
import "fmt"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-9c8a0b2a","gen":"sha256:9c65c23d7893e39a1725f1f00b3344e10ee8e1823f7702b3722d3bef7130ba0c","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/add_scope_run.go","path":"io"}}
import "io"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-9c8a0b2a","gen":"sha256:9c65c23d7893e39a1725f1f00b3344e10ee8e1823f7702b3722d3bef7130ba0c","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/add_scope_run.go","path":"path/filepath"}}
import "path/filepath"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-9c8a0b2a","gen":"sha256:9c65c23d7893e39a1725f1f00b3344e10ee8e1823f7702b3722d3bef7130ba0c","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/add_scope_run.go","path":"sort"}}
import "sort"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-9c8a0b2a","gen":"sha256:9c65c23d7893e39a1725f1f00b3344e10ee8e1823f7702b3722d3bef7130ba0c","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/add_scope_run.go","path":"strings"}}
import "strings"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-9c8a0b2a","gen":"sha256:9c65c23d7893e39a1725f1f00b3344e10ee8e1823f7702b3722d3bef7130ba0c","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/add_scope_run.go","path":"github.com/spf13/cobra"}}
import "github.com/spf13/cobra"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-9c8a0b2a","gen":"sha256:9c65c23d7893e39a1725f1f00b3344e10ee8e1823f7702b3722d3bef7130ba0c","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/add_scope_run.go","path":"github.com/livecodelife/provenance/record"}}
import "github.com/livecodelife/provenance/record"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-9c8a0b2a","gen":"sha256:9c65c23d7893e39a1725f1f00b3344e10ee8e1823f7702b3722d3bef7130ba0c","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/add_scope_run.go","path":"github.com/livecodelife/provenance/scope"}}
import "github.com/livecodelife/provenance/scope"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"tier":"owned","record":"prov-2026-9c8a0b2a","gen":"sha256:9c65c23d7893e39a1725f1f00b3344e10ee8e1823f7702b3722d3bef7130ba0c","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func scopeCandidates(r *repo, rec *record.Record) []string {\n\town := filepath.ToSlash(r.recordPath(rec.ID))\n\tdir := filepath.ToSlash(r.Config.Dir) + \"/\"\n\tvar added []string\n\tseen := map[string]bool{}\n\tfor _, f := range filesInCommits(r, commitsForRecord(r, rec)) {\n\t\tif seen[f] || f == own || strings.HasPrefix(f, dir) || f == manifestPath || stateExempt(r, f) || scope.Excluded(f, r.Config.ExcludePaths) {\n\t\t\tcontinue\n\t\t}\n\t\tseen[f] = true\n\t\tif ok, err := scope.InScope(rec, f); err == nil && ok {\n\t\t\tcontinue\n\t\t}\n\t\tadded = append(added, f)\n\t}\n\tsort.Strings(added)\n\treturn added\n}","file":"cmd/provenance/add_scope_run.go","name":"scopeCandidates"}}
func scopeCandidates(r *repo, rec *record.Record) []string {
	own := filepath.ToSlash(r.recordPath(rec.ID))
	dir := filepath.ToSlash(r.Config.Dir) + "/"
	var added []string
	seen := map[string]bool{}
	for _, f := range filesInCommits(r, commitsForRecord(r, rec)) {
		if seen[f] || f == own || strings.HasPrefix(f, dir) || f == manifestPath || stateExempt(r, f) || scope.Excluded(f, r.Config.ExcludePaths) {
			continue
		}
		seen[f] = true
		if ok, err := scope.InScope(rec, f); err == nil && ok {
			continue
		}
		added = append(added, f)
	}
	sort.Strings(added)
	return added
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-9c8a0b2a","gen":"sha256:9c65c23d7893e39a1725f1f00b3344e10ee8e1823f7702b3722d3bef7130ba0c","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func runAddScope(out io.Writer, cmd *cobra.Command, args []string) error {\n\tr, err := openRepo(cmd)\n\tif err != nil {\n\t\treturn err\n\t}\n\tid, _ := cmd.Flags().GetString(\"record\")\n\tif id == \"\" {\n\t\treturn errors.New(`required flag(s) \"record\" not set`)\n\t}\n\tdry, _ := cmd.Flags().GetBool(\"dry-run\")\n\treason, _ := cmd.Flags().GetString(\"reason\")\n\trec, doc, err := r.resolve(id)\n\tif err != nil {\n\t\treturn err\n\t}\n\tif err := versionable(r, rec); err != nil {\n\t\treturn err\n\t}\n\tif rec.Status != record.StatusOpen {\n\t\tr.fail(\"Cannot add scope to %s: it is %s (add-scope applies to an open record).\", rec.ID, rec.Status)\n\t\treturn errors.New(\"record is not open\")\n\t}\n\tif scope.Mode(rec) != \"allowlist\" {\n\t\tr.fail(\"%s has no affected_scope yet: give it one with provenance edit --add-scope.\", rec.ID)\n\t\treturn errors.New(\"record is not in allowlist mode\")\n\t}\n\tadded := scopeCandidates(r, rec)\n\tif len(added) == 0 {\n\t\tr.fail(\"%s has no newly committed files to add to affected_scope\", rec.ID)\n\t\treturn errors.New(\"nothing new to add\")\n\t}\n\tif dry {\n\t\tfmt.Fprintf(out, \"\\nWould add to the affected_scope of %s:\\n\", rec.ID)\n\t\tfor _, f := range added {\n\t\t\tfmt.Fprintf(out, \"    · %s\\n\", f)\n\t\t}\n\t\tfmt.Fprint(out, \"\\n\")\n\t\treturn nil\n\t}\n\tif reason == \"\" {\n\t\treason = \"add-scope: files committed under \" + rec.ID\n\t}\n\tfmt.Fprintf(out, \"\\nAdding %d file(s) to the affected_scope of %s:\\n\", len(added), rec.ID)\n\tfor _, f := range added {\n\t\tfmt.Fprintf(out, \"    · %s\\n\", f)\n\t}\n\trec.AffectedScope = append(rec.AffectedScope, added...)\n\treturn saveVersion(out, r, rec, doc, reason, false)\n}","file":"cmd/provenance/add_scope_run.go","name":"runAddScope"}}
func runAddScope(out io.Writer, cmd *cobra.Command, args []string) error {
	r, err := openRepo(cmd)
	if err != nil {
		return err
	}
	id, _ := cmd.Flags().GetString("record")
	if id == "" {
		return errors.New(`required flag(s) "record" not set`)
	}
	dry, _ := cmd.Flags().GetBool("dry-run")
	reason, _ := cmd.Flags().GetString("reason")
	rec, doc, err := r.resolve(id)
	if err != nil {
		return err
	}
	if err := versionable(r, rec); err != nil {
		return err
	}
	if rec.Status != record.StatusOpen {
		r.fail("Cannot add scope to %s: it is %s (add-scope applies to an open record).", rec.ID, rec.Status)
		return errors.New("record is not open")
	}
	if scope.Mode(rec) != "allowlist" {
		r.fail("%s has no affected_scope yet: give it one with provenance edit --add-scope.", rec.ID)
		return errors.New("record is not in allowlist mode")
	}
	added := scopeCandidates(r, rec)
	if len(added) == 0 {
		r.fail("%s has no newly committed files to add to affected_scope", rec.ID)
		return errors.New("nothing new to add")
	}
	if dry {
		fmt.Fprintf(out, "\nWould add to the affected_scope of %s:\n", rec.ID)
		for _, f := range added {
			fmt.Fprintf(out, "    · %s\n", f)
		}
		fmt.Fprint(out, "\n")
		return nil
	}
	if reason == "" {
		reason = "add-scope: files committed under " + rec.ID
	}
	fmt.Fprintf(out, "\nAdding %d file(s) to the affected_scope of %s:\n", len(added), rec.ID)
	for _, f := range added {
		fmt.Fprintf(out, "    · %s\n", f)
	}
	rec.AffectedScope = append(rec.AffectedScope, added...)
	return saveVersion(out, r, rec, doc, reason, false)
}

// /sedum:setFunction
