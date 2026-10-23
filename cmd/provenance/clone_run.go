package main

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-6c411bf8","gen":"sha256:51691496e6409a36b3805de7a143010c350733f87e88eddd1891639d22b9cbb0","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/clone_run.go","path":"fmt"}}
import "fmt"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-6c411bf8","gen":"sha256:51691496e6409a36b3805de7a143010c350733f87e88eddd1891639d22b9cbb0","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/clone_run.go","path":"io"}}
import "io"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-6c411bf8","gen":"sha256:51691496e6409a36b3805de7a143010c350733f87e88eddd1891639d22b9cbb0","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/clone_run.go","path":"os"}}
import "os"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-6c411bf8","gen":"sha256:51691496e6409a36b3805de7a143010c350733f87e88eddd1891639d22b9cbb0","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/clone_run.go","path":"path/filepath"}}
import "path/filepath"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-6c411bf8","gen":"sha256:51691496e6409a36b3805de7a143010c350733f87e88eddd1891639d22b9cbb0","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/clone_run.go","path":"github.com/spf13/cobra"}}
import "github.com/spf13/cobra"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-6c411bf8","gen":"sha256:51691496e6409a36b3805de7a143010c350733f87e88eddd1891639d22b9cbb0","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/clone_run.go","path":"github.com/livecodelife/provenance/record"}}
import "github.com/livecodelife/provenance/record"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-6c411bf8","gen":"sha256:51691496e6409a36b3805de7a143010c350733f87e88eddd1891639d22b9cbb0","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/clone_run.go","path":"github.com/livecodelife/provenance/remote"}}
import "github.com/livecodelife/provenance/remote"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"tier":"owned","record":"prov-2026-6c411bf8","gen":"sha256:51691496e6409a36b3805de7a143010c350733f87e88eddd1891639d22b9cbb0","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func runClone(out io.Writer, cmd *cobra.Command, args []string) error {\n\ttarget := args[0]\n\tids := args[1:]\n\n\tentries, err := os.ReadDir(target)\n\tif err == nil && len(entries) > 0 {\n\t\treturn fmt.Errorf(\"%s is not empty\", target)\n\t}\n\n\tclient, err := apiClient()\n\tif err != nil {\n\t\treturn err\n\t}\n\n\tabsDir, err := filepath.Abs(target)\n\tif err != nil {\n\t\treturn err\n\t}\n\tif err := os.MkdirAll(absDir, 0755); err != nil {\n\t\treturn err\n\t}\n\n\tstore := &record.FilesStore{Root: absDir, Dir: \"provenance\"}\n\tcount, err := remote.CloneRecords(cmd.Context(), client, store, ids)\n\tif err != nil {\n\t\treturn err\n\t}\n\n\tif err := os.WriteFile(filepath.Join(absDir, configFile), []byte(initConfig), 0644); err != nil {\n\t\treturn err\n\t}\n\n\tmanifest, err := record.BuildManifest(store)\n\tif err != nil {\n\t\treturn err\n\t}\n\tif err := os.MkdirAll(filepath.Join(absDir, filepath.Dir(manifestPath)), 0755); err != nil {\n\t\treturn err\n\t}\n\tif err := os.WriteFile(filepath.Join(absDir, manifestPath), record.MarshalManifest(manifest), 0644); err != nil {\n\t\treturn err\n\t}\n\n\tif _, err := ensureGitattributes(absDir, \"provenance\"); err != nil {\n\t\treturn err\n\t}\n\n\tfmt.Fprintf(out, \"\\n✓ Cloned %d record(s) into %s\\n\", count, target)\n\tfmt.Fprintf(out, \"  Every version was verified against its content hash.\\n\")\n\tfmt.Fprintf(out, \"  Spec files are not published: lint reports them until the repository's files are present.\\n\")\n\tfmt.Fprintf(out, \"  Run git init and commit when you are ready.\\n\")\n\tfmt.Fprintln(out)\n\treturn nil\n}","file":"cmd/provenance/clone_run.go","name":"runClone"}}
func runClone(out io.Writer, cmd *cobra.Command, args []string) error {
	target := args[0]
	ids := args[1:]

	entries, err := os.ReadDir(target)
	if err == nil && len(entries) > 0 {
		return fmt.Errorf("%s is not empty", target)
	}

	client, err := apiClient()
	if err != nil {
		return err
	}

	absDir, err := filepath.Abs(target)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(absDir, 0755); err != nil {
		return err
	}

	store := &record.FilesStore{Root: absDir, Dir: "provenance"}
	count, err := remote.CloneRecords(cmd.Context(), client, store, ids)
	if err != nil {
		return err
	}

	if err := os.WriteFile(filepath.Join(absDir, configFile), []byte(initConfig), 0644); err != nil {
		return err
	}

	manifest, err := record.BuildManifest(store)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Join(absDir, filepath.Dir(manifestPath)), 0755); err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(absDir, manifestPath), record.MarshalManifest(manifest), 0644); err != nil {
		return err
	}

	if _, err := ensureGitattributes(absDir, "provenance"); err != nil {
		return err
	}

	fmt.Fprintf(out, "\n✓ Cloned %d record(s) into %s\n", count, target)
	fmt.Fprintf(out, "  Every version was verified against its content hash.\n")
	fmt.Fprintf(out, "  Spec files are not published: lint reports them until the repository's files are present.\n")
	fmt.Fprintf(out, "  Run git init and commit when you are ready.\n")
	fmt.Fprintln(out)
	return nil
}

// /sedum:setFunction
