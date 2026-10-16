package main

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-0b7a4f45","gen":"sha256:b66c51d55d04feb335cbf3ac17e504427cbd9f7c62759735d74359869231d324","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/check_sarif_run.go","path":"errors"}}
import "errors"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-0b7a4f45","gen":"sha256:b66c51d55d04feb335cbf3ac17e504427cbd9f7c62759735d74359869231d324","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/check_sarif_run.go","path":"fmt"}}
import "fmt"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-0b7a4f45","gen":"sha256:b66c51d55d04feb335cbf3ac17e504427cbd9f7c62759735d74359869231d324","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/check_sarif_run.go","path":"io"}}
import "io"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-0b7a4f45","gen":"sha256:b66c51d55d04feb335cbf3ac17e504427cbd9f7c62759735d74359869231d324","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/check_sarif_run.go","path":"os"}}
import "os"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-0b7a4f45","gen":"sha256:b66c51d55d04feb335cbf3ac17e504427cbd9f7c62759735d74359869231d324","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/check_sarif_run.go","path":"path/filepath"}}
import "path/filepath"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-0b7a4f45","gen":"sha256:b66c51d55d04feb335cbf3ac17e504427cbd9f7c62759735d74359869231d324","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/check_sarif_run.go","path":"strings"}}
import "strings"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-0b7a4f45","gen":"sha256:b66c51d55d04feb335cbf3ac17e504427cbd9f7c62759735d74359869231d324","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/check_sarif_run.go","path":"github.com/spf13/cobra"}}
import "github.com/spf13/cobra"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"tier":"owned","record":"prov-2026-0b7a4f45","gen":"sha256:b66c51d55d04feb335cbf3ac17e504427cbd9f7c62759735d74359869231d324","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func classifyViolation(v violation) string {\n\tm := v.Message\n\tswitch {\n\tcase strings.Contains(m, \"forbids changes to\"):\n\t\treturn \"SCOPE001\"\n\tcase strings.HasPrefix(m, \"Commit tag required\"), strings.Contains(m, \"is at version\"):\n\t\treturn \"TAG001\"\n\tdefault:\n\t\treturn \"STATE001\"\n\t}\n}","file":"cmd/provenance/check_sarif_run.go","name":"classifyViolation"}}
func classifyViolation(v violation) string {
	m := v.Message
	switch {
	case strings.Contains(m, "forbids changes to"):
		return "SCOPE001"
	case strings.HasPrefix(m, "Commit tag required"), strings.Contains(m, "is at version"):
		return "TAG001"
	default:
		return "STATE001"
	}
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-0b7a4f45","gen":"sha256:b66c51d55d04feb335cbf3ac17e504427cbd9f7c62759735d74359869231d324","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func violationFinding(r *repo, v violation) sarifFinding {\n\tf := sarifFinding{RuleID: classifyViolation(v), Level: \"error\", Message: v.Message, RecordID: v.RecordID, File: v.File}\n\tif strings.HasPrefix(v.Message, \"lint PROV\") {\n\t\tif id, msg, ok := strings.Cut(strings.TrimPrefix(v.Message, \"lint \"), \": \"); ok {\n\t\t\tf.RuleID = id\n\t\t\tf.Message = msg\n\t\t}\n\t}\n\tif f.File == \"\" && v.RecordID != \"\" {\n\t\tf.File = filepath.ToSlash(r.recordPath(v.RecordID))\n\t}\n\treturn f\n}","file":"cmd/provenance/check_sarif_run.go","name":"violationFinding"}}
func violationFinding(r *repo, v violation) sarifFinding {
	f := sarifFinding{RuleID: classifyViolation(v), Level: "error", Message: v.Message, RecordID: v.RecordID, File: v.File}
	if strings.HasPrefix(v.Message, "lint PROV") {
		if id, msg, ok := strings.Cut(strings.TrimPrefix(v.Message, "lint "), ": "); ok {
			f.RuleID = id
			f.Message = msg
		}
	}
	if f.File == "" && v.RecordID != "" {
		f.File = filepath.ToSlash(r.recordPath(v.RecordID))
	}
	return f
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-0b7a4f45","gen":"sha256:b66c51d55d04feb335cbf3ac17e504427cbd9f7c62759735d74359869231d324","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func evidenceFindings(r *repo, rows []evidenceRow) []sarifFinding {\n\tvar out []sarifFinding\n\tfor _, row := range rows {\n\t\tfile := filepath.ToSlash(r.recordPath(row.ID))\n\t\tswitch {\n\t\tcase row.State == \"stale\":\n\t\t\tout = append(out, sarifFinding{RuleID: \"EVID001\", Level: \"error\", Message: fmt.Sprintf(\"%s@%d evidence is stale: %s\", row.ID, row.Version, strings.Join(row.Reasons, \", \")), File: file, RecordID: row.ID})\n\t\tcase row.State != \"missing\" && !row.Passed:\n\t\t\tout = append(out, sarifFinding{RuleID: \"EVID002\", Level: \"error\", Message: fmt.Sprintf(\"%s@%d latest evidence is a failing run\", row.ID, row.Version), File: file, RecordID: row.ID})\n\t\t}\n\t}\n\treturn out\n}","file":"cmd/provenance/check_sarif_run.go","name":"evidenceFindings"}}
func evidenceFindings(r *repo, rows []evidenceRow) []sarifFinding {
	var out []sarifFinding
	for _, row := range rows {
		file := filepath.ToSlash(r.recordPath(row.ID))
		switch {
		case row.State == "stale":
			out = append(out, sarifFinding{RuleID: "EVID001", Level: "error", Message: fmt.Sprintf("%s@%d evidence is stale: %s", row.ID, row.Version, strings.Join(row.Reasons, ", ")), File: file, RecordID: row.ID})
		case row.State != "missing" && !row.Passed:
			out = append(out, sarifFinding{RuleID: "EVID002", Level: "error", Message: fmt.Sprintf("%s@%d latest evidence is a failing run", row.ID, row.Version), File: file, RecordID: row.ID})
		}
	}
	return out
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-0b7a4f45","gen":"sha256:b66c51d55d04feb335cbf3ac17e504427cbd9f7c62759735d74359869231d324","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func checkSarif(out io.Writer, cmd *cobra.Command, r *repo) error {\n\tstaged, _ := cmd.Flags().GetBool(\"staged\")\n\tcommit, _ := cmd.Flags().GetString(\"commit\")\n\trng, _ := cmd.Flags().GetString(\"range\")\n\tev, _ := cmd.Flags().GetBool(\"evidence\")\n\tif !staged && commit == \"\" && rng == \"\" && !ev {\n\t\treturn errors.New(\"pass --staged, --commit <sha>, --range <a..b> or --evidence\")\n\t}\n\tvar findings []sarifFinding\n\tif ev {\n\t\trows, err := checkEvidence(r)\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n\t\tfindings = append(findings, evidenceFindings(r, rows)...)\n\t}\n\tvar shas []string\n\tif commit != \"\" {\n\t\tsha, err := r.git(\"rev-parse\", \"--verify\", commit+\"^{commit}\")\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n\t\tshas = append(shas, sha)\n\t}\n\tif rng != \"\" {\n\t\tlist, err := r.git(\"rev-list\", \"--reverse\", rng)\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n\t\tshas = append(shas, strings.Fields(list)...)\n\t}\n\tfor _, sha := range shas {\n\t\tvs, err := checkCommit(r, sha)\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n\t\tfor _, v := range vs {\n\t\t\tf := violationFinding(r, v)\n\t\t\tf.Message = sha[:7] + \": \" + f.Message\n\t\t\tfindings = append(findings, f)\n\t\t}\n\t}\n\tif staged {\n\t\tmsgFile, _ := cmd.Flags().GetString(\"message-file\")\n\t\tvar message string\n\t\tif msgFile != \"\" {\n\t\t\tdata, err := os.ReadFile(msgFile)\n\t\t\tif err != nil {\n\t\t\t\treturn err\n\t\t\t}\n\t\t\tmessage = strings.TrimSpace(string(data))\n\t\t} else {\n\t\t\tvar err error\n\t\t\tmessage, err = r.git(\"log\", \"-1\", \"--format=%B\")\n\t\t\tif err != nil {\n\t\t\t\treturn err\n\t\t\t}\n\t\t}\n\t\tvs, err := checkStaged(r, message)\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n\t\tfor _, v := range vs {\n\t\t\tfindings = append(findings, violationFinding(r, v))\n\t\t}\n\t}\n\tif err := printSarif(out, buildSarif(r, findings)); err != nil {\n\t\treturn err\n\t}\n\tif len(findings) > 0 {\n\t\treturn fmt.Errorf(\"%d findings\", len(findings))\n\t}\n\treturn nil\n}","file":"cmd/provenance/check_sarif_run.go","name":"checkSarif"}}
func checkSarif(out io.Writer, cmd *cobra.Command, r *repo) error {
	staged, _ := cmd.Flags().GetBool("staged")
	commit, _ := cmd.Flags().GetString("commit")
	rng, _ := cmd.Flags().GetString("range")
	ev, _ := cmd.Flags().GetBool("evidence")
	if !staged && commit == "" && rng == "" && !ev {
		return errors.New("pass --staged, --commit <sha>, --range <a..b> or --evidence")
	}
	var findings []sarifFinding
	if ev {
		rows, err := checkEvidence(r)
		if err != nil {
			return err
		}
		findings = append(findings, evidenceFindings(r, rows)...)
	}
	var shas []string
	if commit != "" {
		sha, err := r.git("rev-parse", "--verify", commit+"^{commit}")
		if err != nil {
			return err
		}
		shas = append(shas, sha)
	}
	if rng != "" {
		list, err := r.git("rev-list", "--reverse", rng)
		if err != nil {
			return err
		}
		shas = append(shas, strings.Fields(list)...)
	}
	for _, sha := range shas {
		vs, err := checkCommit(r, sha)
		if err != nil {
			return err
		}
		for _, v := range vs {
			f := violationFinding(r, v)
			f.Message = sha[:7] + ": " + f.Message
			findings = append(findings, f)
		}
	}
	if staged {
		msgFile, _ := cmd.Flags().GetString("message-file")
		var message string
		if msgFile != "" {
			data, err := os.ReadFile(msgFile)
			if err != nil {
				return err
			}
			message = strings.TrimSpace(string(data))
		} else {
			var err error
			message, err = r.git("log", "-1", "--format=%B")
			if err != nil {
				return err
			}
		}
		vs, err := checkStaged(r, message)
		if err != nil {
			return err
		}
		for _, v := range vs {
			findings = append(findings, violationFinding(r, v))
		}
	}
	if err := printSarif(out, buildSarif(r, findings)); err != nil {
		return err
	}
	if len(findings) > 0 {
		return fmt.Errorf("%d findings", len(findings))
	}
	return nil
}

// /sedum:setFunction
