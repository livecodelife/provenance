package main

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-adabd65f","gen":"sha256:4ae46f3089d333bd12c54eda92a2d20ae90a3ad1ebf7c65d3eda625c0356ee02","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/version_run.go","path":"fmt"}}
import "fmt"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-adabd65f","gen":"sha256:4ae46f3089d333bd12c54eda92a2d20ae90a3ad1ebf7c65d3eda625c0356ee02","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/version_run.go","path":"io"}}
import "io"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-adabd65f","gen":"sha256:4ae46f3089d333bd12c54eda92a2d20ae90a3ad1ebf7c65d3eda625c0356ee02","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/version_run.go","path":"github.com/spf13/cobra"}}
import "github.com/spf13/cobra"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-adabd65f","gen":"sha256:4ae46f3089d333bd12c54eda92a2d20ae90a3ad1ebf7c65d3eda625c0356ee02","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/version_run.go","path":"github.com/livecodelife/provenance/version"}}
import "github.com/livecodelife/provenance/version"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"tier":"owned","record":"prov-2026-adabd65f","gen":"sha256:4ae46f3089d333bd12c54eda92a2d20ae90a3ad1ebf7c65d3eda625c0356ee02","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func runVersion(out io.Writer, cmd *cobra.Command, args []string) error {\n\t_, err := fmt.Fprintln(out, version.Version)\n\treturn err\n}","file":"cmd/provenance/version_run.go","name":"runVersion"}}
func runVersion(out io.Writer, cmd *cobra.Command, args []string) error {
	_, err := fmt.Fprintln(out, version.Version)
	return err
}

// /sedum:setFunction
