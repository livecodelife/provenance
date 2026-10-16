package main

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-15f42a00","gen":"sha256:c80bc449c18db599e38113fc2c49132fd1fd245096f19d2387e80dc6dcd91188","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/evidence_refresh_run.go","path":"fmt"}}
import "fmt"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-15f42a00","gen":"sha256:c80bc449c18db599e38113fc2c49132fd1fd245096f19d2387e80dc6dcd91188","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/evidence_refresh_run.go","path":"io"}}
import "io"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-15f42a00","gen":"sha256:c80bc449c18db599e38113fc2c49132fd1fd245096f19d2387e80dc6dcd91188","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/evidence_refresh_run.go","path":"os"}}
import "os"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-15f42a00","gen":"sha256:c80bc449c18db599e38113fc2c49132fd1fd245096f19d2387e80dc6dcd91188","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/evidence_refresh_run.go","path":"path/filepath"}}
import "path/filepath"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-15f42a00","gen":"sha256:c80bc449c18db599e38113fc2c49132fd1fd245096f19d2387e80dc6dcd91188","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/evidence_refresh_run.go","path":"strings"}}
import "strings"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-15f42a00","gen":"sha256:c80bc449c18db599e38113fc2c49132fd1fd245096f19d2387e80dc6dcd91188","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/evidence_refresh_run.go","path":"time"}}
import "time"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-15f42a00","gen":"sha256:c80bc449c18db599e38113fc2c49132fd1fd245096f19d2387e80dc6dcd91188","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/evidence_refresh_run.go","path":"github.com/spf13/cobra"}}
import "github.com/spf13/cobra"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-15f42a00","gen":"sha256:c80bc449c18db599e38113fc2c49132fd1fd245096f19d2387e80dc6dcd91188","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/evidence_refresh_run.go","path":"github.com/livecodelife/provenance/evidence"}}
import "github.com/livecodelife/provenance/evidence"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-15f42a00","gen":"sha256:c80bc449c18db599e38113fc2c49132fd1fd245096f19d2387e80dc6dcd91188","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/evidence_refresh_run.go","path":"github.com/livecodelife/provenance/record"}}
import "github.com/livecodelife/provenance/record"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"tier":"owned","record":"prov-2026-15f42a00","gen":"sha256:c80bc449c18db599e38113fc2c49132fd1fd245096f19d2387e80dc6dcd91188","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func codeOnlyStale(row evidenceRow) bool {\n\tif row.State != string(evidence.Stale) {\n\t\treturn false\n\t}\n\thasCode := false\n\tfor _, reason := range row.Reasons {\n\t\tswitch reason {\n\t\tcase \"content\", \"specs\", \"parent\":\n\t\t\treturn false\n\t\tcase \"code\":\n\t\t\thasCode = true\n\t\t}\n\t}\n\treturn hasCode\n}","file":"cmd/provenance/evidence_refresh_run.go","name":"codeOnlyStale"}}
func codeOnlyStale(row evidenceRow) bool {
	if row.State != string(evidence.Stale) {
		return false
	}
	hasCode := false
	for _, reason := range row.Reasons {
		switch reason {
		case "content", "specs", "parent":
			return false
		case "code":
			hasCode = true
		}
	}
	return hasCode
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-15f42a00","gen":"sha256:c80bc449c18db599e38113fc2c49132fd1fd245096f19d2387e80dc6dcd91188","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func refreshEvidence(out io.Writer, r *repo, id, reason string) error {\n\trec, _, err := r.Store.Load(id)\n\tif err != nil {\n\t\treturn err\n\t}\n\tif rec.Status != record.StatusImplemented {\n\t\treturn fmt.Errorf(\"%s is %s: only an implemented record's evidence is refreshed (use provenance complete)\", id, rec.Status)\n\t}\n\tstore := &evidence.Store{Root: r.Root, Dir: r.Config.Dir}\n\tlatest, err := store.Latest(rec.ID, rec.Version)\n\tif err != nil {\n\t\treturn err\n\t}\n\tif latest == nil {\n\t\treturn fmt.Errorf(\"%s has no evidence for version %d to refresh: run provenance complete or revise it first\", id, rec.Version)\n\t}\n\tbehind := behindSet(r)\n\tverdict := evidence.Check(latest, currentEvidence(r, rec, behind))\n\tvar moved []string\n\tfor _, why := range verdict.Reasons {\n\t\tif why == \"content\" || why == \"specs\" {\n\t\t\tmoved = append(moved, why)\n\t\t}\n\t}\n\tif len(moved) > 0 {\n\t\treturn fmt.Errorf(\"%s changed since its latest evidence (%s): use provenance revise or edit --refresh-specs, not a refresh\", id, strings.Join(moved, \", \"))\n\t}\n\tran, failed, _ := runRecordSpecs(out, r, rec, nil)\n\tif !ran {\n\t\treturn fmt.Errorf(\"%s has no runnable associated specs to refresh its evidence with\", id)\n\t}\n\tif failed {\n\t\treturn fmt.Errorf(\"%s: associated specs failed; no evidence was written\", id)\n\t}\n\tbefore := evidenceFiles(r, rec.ID)\n\tcur := currentEvidence(r, rec, behind)\n\tat := time.Now().UTC().Truncate(time.Second)\n\tif prev, err := time.Parse(time.RFC3339, latest.At); err == nil && !at.After(prev) {\n\t\tat = prev.Add(time.Second)\n\t}\n\tfresh := evidence.New(rec, cur, true, at.Format(time.RFC3339), \"provenance evidence refresh\")\n\tif err := store.Append(fresh); err != nil {\n\t\treturn err\n\t}\n\tvar added []string\n\tfor f := range evidenceFiles(r, rec.ID) {\n\t\tif !before[f] {\n\t\t\tadded = append(added, f)\n\t\t}\n\t}\n\tif r.Config.CommitOnStatusChange {\n\t\tvar paths []string\n\t\tfor _, f := range added {\n\t\t\tif rel, err := filepath.Rel(r.Root, f); err == nil {\n\t\t\t\tpaths = append(paths, rel)\n\t\t\t}\n\t\t}\n\t\tif err := r.commit(\"Refresh evidence of \"+rec.ID+\" \"+tag(rec.ID, rec.Version), paths...); err != nil {\n\t\t\tfor _, f := range added {\n\t\t\t\tos.Remove(f)\n\t\t\t}\n\t\t\tr.unstage(paths...)\n\t\t\treturn fmt.Errorf(\"%s: the auto-commit was rejected: %w\", id, err)\n\t\t}\n\t}\n\tnow := evidence.Check(fresh, cur)\n\tfmt.Fprintf(out, \"\\n✓ Refreshed evidence of %s@%d (%s)\\n\", rec.ID, rec.Version, now.State)\n\tif reason != \"\" {\n\t\tfmt.Fprintf(out, \"  Reason: %s\\n\", reason)\n\t}\n\tfmt.Fprint(out, \"  No new version was written; the record is unchanged.\\n\\n\")\n\treturn nil\n}","file":"cmd/provenance/evidence_refresh_run.go","name":"refreshEvidence"}}
func refreshEvidence(out io.Writer, r *repo, id, reason string) error {
	rec, _, err := r.Store.Load(id)
	if err != nil {
		return err
	}
	if rec.Status != record.StatusImplemented {
		return fmt.Errorf("%s is %s: only an implemented record's evidence is refreshed (use provenance complete)", id, rec.Status)
	}
	store := &evidence.Store{Root: r.Root, Dir: r.Config.Dir}
	latest, err := store.Latest(rec.ID, rec.Version)
	if err != nil {
		return err
	}
	if latest == nil {
		return fmt.Errorf("%s has no evidence for version %d to refresh: run provenance complete or revise it first", id, rec.Version)
	}
	behind := behindSet(r)
	verdict := evidence.Check(latest, currentEvidence(r, rec, behind))
	var moved []string
	for _, why := range verdict.Reasons {
		if why == "content" || why == "specs" {
			moved = append(moved, why)
		}
	}
	if len(moved) > 0 {
		return fmt.Errorf("%s changed since its latest evidence (%s): use provenance revise or edit --refresh-specs, not a refresh", id, strings.Join(moved, ", "))
	}
	ran, failed, _ := runRecordSpecs(out, r, rec, nil)
	if !ran {
		return fmt.Errorf("%s has no runnable associated specs to refresh its evidence with", id)
	}
	if failed {
		return fmt.Errorf("%s: associated specs failed; no evidence was written", id)
	}
	before := evidenceFiles(r, rec.ID)
	cur := currentEvidence(r, rec, behind)
	at := time.Now().UTC().Truncate(time.Second)
	if prev, err := time.Parse(time.RFC3339, latest.At); err == nil && !at.After(prev) {
		at = prev.Add(time.Second)
	}
	fresh := evidence.New(rec, cur, true, at.Format(time.RFC3339), "provenance evidence refresh")
	if err := store.Append(fresh); err != nil {
		return err
	}
	var added []string
	for f := range evidenceFiles(r, rec.ID) {
		if !before[f] {
			added = append(added, f)
		}
	}
	if r.Config.CommitOnStatusChange {
		var paths []string
		for _, f := range added {
			if rel, err := filepath.Rel(r.Root, f); err == nil {
				paths = append(paths, rel)
			}
		}
		if err := r.commit("Refresh evidence of "+rec.ID+" "+tag(rec.ID, rec.Version), paths...); err != nil {
			for _, f := range added {
				os.Remove(f)
			}
			r.unstage(paths...)
			return fmt.Errorf("%s: the auto-commit was rejected: %w", id, err)
		}
	}
	now := evidence.Check(fresh, cur)
	fmt.Fprintf(out, "\n✓ Refreshed evidence of %s@%d (%s)\n", rec.ID, rec.Version, now.State)
	if reason != "" {
		fmt.Fprintf(out, "  Reason: %s\n", reason)
	}
	fmt.Fprint(out, "  No new version was written; the record is unchanged.\n\n")
	return nil
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-15f42a00","gen":"sha256:c80bc449c18db599e38113fc2c49132fd1fd245096f19d2387e80dc6dcd91188","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func runEvidenceRefresh(out io.Writer, cmd *cobra.Command, args []string) error {\n\tr, err := openRepo(cmd)\n\tif err != nil {\n\t\treturn err\n\t}\n\tall, _ := cmd.Flags().GetBool(\"all\")\n\treason, _ := cmd.Flags().GetString(\"reason\")\n\tvar ids []string\n\tif all {\n\t\trows, err := checkEvidence(r)\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n\t\tfor _, row := range rows {\n\t\t\tif codeOnlyStale(row) {\n\t\t\t\tids = append(ids, row.ID)\n\t\t\t}\n\t\t}\n\t\tif len(ids) == 0 {\n\t\t\tfmt.Fprint(out, \"\\n✓ Nothing to refresh: no implemented record is stale only because its code moved\\n\\n\")\n\t\t\treturn nil\n\t\t}\n\t} else {\n\t\tid, err := recordArg(cmd, args)\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n\t\trec, _, err := r.resolve(id)\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n\t\tids = []string{rec.ID}\n\t}\n\tfailed := 0\n\tfor _, id := range ids {\n\t\tif err := refreshEvidence(out, r, id, reason); err != nil {\n\t\t\tfailed++\n\t\t\tfmt.Fprintf(out, \"\\n✗ %v\\n\\n\", err)\n\t\t}\n\t}\n\tif failed > 0 {\n\t\treturn fmt.Errorf(\"%d of %d records could not be refreshed\", failed, len(ids))\n\t}\n\treturn nil\n}","file":"cmd/provenance/evidence_refresh_run.go","name":"runEvidenceRefresh"}}
func runEvidenceRefresh(out io.Writer, cmd *cobra.Command, args []string) error {
	r, err := openRepo(cmd)
	if err != nil {
		return err
	}
	all, _ := cmd.Flags().GetBool("all")
	reason, _ := cmd.Flags().GetString("reason")
	var ids []string
	if all {
		rows, err := checkEvidence(r)
		if err != nil {
			return err
		}
		for _, row := range rows {
			if codeOnlyStale(row) {
				ids = append(ids, row.ID)
			}
		}
		if len(ids) == 0 {
			fmt.Fprint(out, "\n✓ Nothing to refresh: no implemented record is stale only because its code moved\n\n")
			return nil
		}
	} else {
		id, err := recordArg(cmd, args)
		if err != nil {
			return err
		}
		rec, _, err := r.resolve(id)
		if err != nil {
			return err
		}
		ids = []string{rec.ID}
	}
	failed := 0
	for _, id := range ids {
		if err := refreshEvidence(out, r, id, reason); err != nil {
			failed++
			fmt.Fprintf(out, "\n✗ %v\n\n", err)
		}
	}
	if failed > 0 {
		return fmt.Errorf("%d of %d records could not be refreshed", failed, len(ids))
	}
	return nil
}

// /sedum:setFunction
