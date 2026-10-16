package main

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-9cf58fc4","gen":"sha256:b8a28407718aa5c94529f2febaecce67a0f06ee4c6f3bc1891ee5e818472e196","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/evidence_hint_run.go","path":"fmt"}}
import "fmt"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-9cf58fc4","gen":"sha256:b8a28407718aa5c94529f2febaecce67a0f06ee4c6f3bc1891ee5e818472e196","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/evidence_hint_run.go","path":"io"}}
import "io"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-9cf58fc4","gen":"sha256:b8a28407718aa5c94529f2febaecce67a0f06ee4c6f3bc1891ee5e818472e196","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/evidence_hint_run.go","path":"os"}}
import "os"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-9cf58fc4","gen":"sha256:b8a28407718aa5c94529f2febaecce67a0f06ee4c6f3bc1891ee5e818472e196","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/evidence_hint_run.go","path":"os/exec"}}
import "os/exec"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"tier":"owned","record":"prov-2026-9cf58fc4","gen":"sha256:b8a28407718aa5c94529f2febaecce67a0f06ee4c6f3bc1891ee5e818472e196","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func codeUnknownWhy() string {\n\tbin := os.Getenv(\"SEDUM_BIN\")\n\tif bin == \"\" {\n\t\tbin = \"sedum\"\n\t}\n\tif _, err := exec.LookPath(bin); err != nil {\n\t\treturn fmt.Sprintf(\"sedum was not found (tried %q)\", bin)\n\t}\n\treturn fmt.Sprintf(\"%q could not list this record's regions (it may predate the regions command, or a region is unstamped: run sedum grow --stamp)\", bin)\n}","file":"cmd/provenance/evidence_hint_run.go","name":"codeUnknownWhy"}}
func codeUnknownWhy() string {
	bin := os.Getenv("SEDUM_BIN")
	if bin == "" {
		bin = "sedum"
	}
	if _, err := exec.LookPath(bin); err != nil {
		return fmt.Sprintf("sedum was not found (tried %q)", bin)
	}
	return fmt.Sprintf("%q could not list this record's regions (it may predate the regions command, or a region is unstamped: run sedum grow --stamp)", bin)
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-9cf58fc4","gen":"sha256:b8a28407718aa5c94529f2febaecce67a0f06ee4c6f3bc1891ee5e818472e196","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func printCodeUnknownWarning(out io.Writer, id string) {\n\tfmt.Fprintf(out, \"\\n⚠ The code hash of %s is unknown: %s.\\n  provenance check --evidence will report it as unknown until this is fixed.\\n  Fix: install a current sedum (or point SEDUM_BIN at one), then run: provenance evidence refresh %s\\n\\n\", id, codeUnknownWhy(), id)\n}","file":"cmd/provenance/evidence_hint_run.go","name":"printCodeUnknownWarning"}}
func printCodeUnknownWarning(out io.Writer, id string) {
	fmt.Fprintf(out, "\n⚠ The code hash of %s is unknown: %s.\n  provenance check --evidence will report it as unknown until this is fixed.\n  Fix: install a current sedum (or point SEDUM_BIN at one), then run: provenance evidence refresh %s\n\n", id, codeUnknownWhy(), id)
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-9cf58fc4","gen":"sha256:b8a28407718aa5c94529f2febaecce67a0f06ee4c6f3bc1891ee5e818472e196","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func evidenceCodeUnknown(row evidenceRow) bool {\n\tif row.State != \"unknown\" {\n\t\treturn false\n\t}\n\tfor _, reason := range row.Reasons {\n\t\tif reason == \"code_unknown\" {\n\t\t\treturn true\n\t\t}\n\t}\n\treturn false\n}","file":"cmd/provenance/evidence_hint_run.go","name":"evidenceCodeUnknown"}}
func evidenceCodeUnknown(row evidenceRow) bool {
	if row.State != "unknown" {
		return false
	}
	for _, reason := range row.Reasons {
		if reason == "code_unknown" {
			return true
		}
	}
	return false
}

// /sedum:setFunction
