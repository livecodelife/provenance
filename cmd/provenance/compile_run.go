package main

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-d84dc10e","gen":"sha256:f20a2537833e4924b62273583dc6d4c9f1678ed42ca83ff10d54da85435d4d91","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/compile_run.go","path":"bytes"}}
import "bytes"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-d84dc10e","gen":"sha256:f20a2537833e4924b62273583dc6d4c9f1678ed42ca83ff10d54da85435d4d91","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/compile_run.go","path":"fmt"}}
import "fmt"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-d84dc10e","gen":"sha256:f20a2537833e4924b62273583dc6d4c9f1678ed42ca83ff10d54da85435d4d91","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/compile_run.go","path":"io"}}
import "io"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-d84dc10e","gen":"sha256:f20a2537833e4924b62273583dc6d4c9f1678ed42ca83ff10d54da85435d4d91","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/compile_run.go","path":"os"}}
import "os"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-d84dc10e","gen":"sha256:f20a2537833e4924b62273583dc6d4c9f1678ed42ca83ff10d54da85435d4d91","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/compile_run.go","path":"path/filepath"}}
import "path/filepath"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-d84dc10e","gen":"sha256:f20a2537833e4924b62273583dc6d4c9f1678ed42ca83ff10d54da85435d4d91","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/compile_run.go","path":"github.com/spf13/cobra"}}
import "github.com/spf13/cobra"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-d84dc10e","gen":"sha256:f20a2537833e4924b62273583dc6d4c9f1678ed42ca83ff10d54da85435d4d91","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/compile_run.go","path":"github.com/livecodelife/provenance/record"}}
import "github.com/livecodelife/provenance/record"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"tier":"owned","record":"prov-2026-d84dc10e","gen":"sha256:f20a2537833e4924b62273583dc6d4c9f1678ed42ca83ff10d54da85435d4d91","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func runCompile(out io.Writer, cmd *cobra.Command, args []string) error {\n\tr, err := openRepo(cmd)\n\tif err != nil {\n\t\treturn err\n\t}\n\tlines, err := record.BuildManifest(r.Store)\n\tif err != nil {\n\t\treturn err\n\t}\n\thave, _ := os.ReadFile(filepath.Join(r.Root, manifestPath))\n\tif bytes.Equal(have, record.MarshalManifest(lines)) {\n\t\tfmt.Fprintf(out, \"\\n✓ Version manifest is up to date (%d records)\\n\\n\", len(lines))\n\t\treturn nil\n\t}\n\tif err := r.writeManifest(); err != nil {\n\t\treturn err\n\t}\n\tfmt.Fprintf(out, \"\\n✓ Version manifest compiled (%d records)\\n\\n\", len(lines))\n\treturn nil\n}","file":"cmd/provenance/compile_run.go","name":"runCompile"}}
func runCompile(out io.Writer, cmd *cobra.Command, args []string) error {
	r, err := openRepo(cmd)
	if err != nil {
		return err
	}
	lines, err := record.BuildManifest(r.Store)
	if err != nil {
		return err
	}
	have, _ := os.ReadFile(filepath.Join(r.Root, manifestPath))
	if bytes.Equal(have, record.MarshalManifest(lines)) {
		fmt.Fprintf(out, "\n✓ Version manifest is up to date (%d records)\n\n", len(lines))
		return nil
	}
	if err := r.writeManifest(); err != nil {
		return err
	}
	fmt.Fprintf(out, "\n✓ Version manifest compiled (%d records)\n\n", len(lines))
	return nil
}

// /sedum:setFunction
