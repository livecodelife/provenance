package main

// sedum:anchor:imports
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-0b7a4f45","gen":"sha256:dbf78db9ba625aefeed1e964eefa4d37f002ff9c2ff20d6c17bd027babfeb7d1","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/check_sarif_run.go","path":"errors"},"kwargs":"sha256:c6847c9e96e85448dd199f4a64250a75fdcd7c5257c4834335167647fb4b48db","body":"sha256:2b718237ba0a56467c8fd494cc131e942ef410e6228c03b644d70476c71862dc"}
import "errors"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-0b7a4f45","gen":"sha256:dbf78db9ba625aefeed1e964eefa4d37f002ff9c2ff20d6c17bd027babfeb7d1","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/check_sarif_run.go","path":"fmt"},"kwargs":"sha256:c9aaa765d9c21827be8bc2e6a1a257539295f3c33365f2793523e134dbb3084a","body":"sha256:bcbac51883cbf8064e9196b371cb880cee429866652447ad557533d817857609"}
import "fmt"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-0b7a4f45","gen":"sha256:dbf78db9ba625aefeed1e964eefa4d37f002ff9c2ff20d6c17bd027babfeb7d1","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/check_sarif_run.go","path":"io"},"kwargs":"sha256:a3e3bd21579fe43518fdaf6d7b749dea011806ae4d4d718a93876804d33b483f","body":"sha256:ef8e8abf4ef5ee40a19dc85b49bf09c5df1fe32abfe834e0ab4c2175baf80ddf"}
import "io"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-0b7a4f45","gen":"sha256:dbf78db9ba625aefeed1e964eefa4d37f002ff9c2ff20d6c17bd027babfeb7d1","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/check_sarif_run.go","path":"os"},"kwargs":"sha256:b8641f7321b832e0068160e1376e33a4bd9149bf01c82ae4725afe67a217c47f","body":"sha256:7c6e309eb2575d4aa451ec2ab83161d13e9ae03bbca3a4d2ecf30ce12adf0905"}
import "os"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-0b7a4f45","gen":"sha256:dbf78db9ba625aefeed1e964eefa4d37f002ff9c2ff20d6c17bd027babfeb7d1","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/check_sarif_run.go","path":"path/filepath"},"kwargs":"sha256:ca19e06f2fb125b9cf3734f2f6db84c0a1532042cd68a589972129d61c337ba3","body":"sha256:a6d196b558491e646be93f56a321f339068c64c76aa9d49bb18046e7f9c777a1"}
import "path/filepath"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-0b7a4f45","gen":"sha256:dbf78db9ba625aefeed1e964eefa4d37f002ff9c2ff20d6c17bd027babfeb7d1","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/check_sarif_run.go","path":"strings"},"kwargs":"sha256:89b792421ffc30ade6cf0d7beafe2928c302e18c14a983d3914df4c3d655e0d3","body":"sha256:a9af3c0bc4bf9342df9d7e89a324523fdba7941a9ca4190270e68ff3c4a25ea6"}
import "strings"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-0b7a4f45","gen":"sha256:dbf78db9ba625aefeed1e964eefa4d37f002ff9c2ff20d6c17bd027babfeb7d1","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/check_sarif_run.go","path":"github.com/spf13/cobra"},"kwargs":"sha256:a6600c9b822e603dc28ce15b307ce99815c246df36a9544014b5324376a85873","body":"sha256:f361c02d0f8e1999cf3fb715393e16310eeec9debd5aaa13936af7617e2b46aa"}
import "github.com/spf13/cobra"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-0b7a4f45","gen":"sha256:dbf78db9ba625aefeed1e964eefa4d37f002ff9c2ff20d6c17bd027babfeb7d1","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/check_sarif_run.go","name":"classifyViolation"},"kwargs":"sha256:54975216275c1f49342c816f49227d2e2dced736905299d8654668b76948bf0f","body":"sha256:9dc6c39dac03d5a4801aea78fe5d63e5e66d882cbd9e91ab54e5394176eea185"}
func classifyViolation(v violation) string {
	m := v.Message
	switch {
	case strings.Contains(m, "forbids changes to"):
		return "SCOPE001"
	case strings.HasPrefix(m, "Commit tag required"), strings.Contains(m, "is at version"):
		return "TAG001"
	case strings.Contains(m, "is implemented: run 'provenance revise"):
		return "REVISE001"
	default:
		return "STATE001"
	}
}

// /sedum:setFunction
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-0b7a4f45","gen":"sha256:dbf78db9ba625aefeed1e964eefa4d37f002ff9c2ff20d6c17bd027babfeb7d1","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/check_sarif_run.go","name":"violationFinding"},"kwargs":"sha256:b00d66c9928bc16d062f16862e6f64fb48c9f05b23dca7acb583c6540778f6a3","body":"sha256:c9715257250fe34eac7925a90de3f4b876ecb8a629d85c6be83bb70e23cf9864"}
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
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-0b7a4f45","gen":"sha256:dbf78db9ba625aefeed1e964eefa4d37f002ff9c2ff20d6c17bd027babfeb7d1","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/check_sarif_run.go","name":"evidenceFindings"},"kwargs":"sha256:d7e47e77aea0cea2b70cfdec2078a101eea793656dec5c22a384dfabd6b2aaff","body":"sha256:c3c3f6e7544d25e8944255d93fd9b239443115a4f4ca604e0bd8b96f689e7ae4"}
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
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-0b7a4f45","gen":"sha256:dbf78db9ba625aefeed1e964eefa4d37f002ff9c2ff20d6c17bd027babfeb7d1","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/check_sarif_run.go","name":"checkSarif"},"kwargs":"sha256:d0abfc45278248929ce94468ec583f37185523d3255cfc6b3a09b66936be8663","body":"sha256:dd7e88466b807c02aaeeb4de1ce7fa017e66e52ed14b688763bb94d6355b4477"}
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
