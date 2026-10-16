package main

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-75646f74","gen":"sha256:bfc4a4b4ff7397d2474dbe459bfda6f85459580329efbe550b3e2379a9c827c6","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/clone_run.go","path":"fmt"}}
import "fmt"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-75646f74","gen":"sha256:bfc4a4b4ff7397d2474dbe459bfda6f85459580329efbe550b3e2379a9c827c6","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/clone_run.go","path":"io"}}
import "io"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-75646f74","gen":"sha256:bfc4a4b4ff7397d2474dbe459bfda6f85459580329efbe550b3e2379a9c827c6","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/clone_run.go","path":"os"}}
import "os"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-75646f74","gen":"sha256:bfc4a4b4ff7397d2474dbe459bfda6f85459580329efbe550b3e2379a9c827c6","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/clone_run.go","path":"path/filepath"}}
import "path/filepath"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-75646f74","gen":"sha256:bfc4a4b4ff7397d2474dbe459bfda6f85459580329efbe550b3e2379a9c827c6","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/clone_run.go","path":"github.com/spf13/cobra"}}
import "github.com/spf13/cobra"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-75646f74","gen":"sha256:bfc4a4b4ff7397d2474dbe459bfda6f85459580329efbe550b3e2379a9c827c6","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/clone_run.go","path":"github.com/livecodelife/provenance/record"}}
import "github.com/livecodelife/provenance/record"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-75646f74","gen":"sha256:bfc4a4b4ff7397d2474dbe459bfda6f85459580329efbe550b3e2379a9c827c6","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/clone_run.go","path":"github.com/livecodelife/provenance/remote"}}
import "github.com/livecodelife/provenance/remote"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"tier":"owned","record":"prov-2026-75646f74","gen":"sha256:bfc4a4b4ff7397d2474dbe459bfda6f85459580329efbe550b3e2379a9c827c6","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func runClone(out io.Writer, cmd *cobra.Command, args []string) error {\n\tdir, ids := args[0], args[1:]\n\tif entries, err := os.ReadDir(dir); err == nil && len(entries) > 0 {\n\t\treturn fmt.Errorf(\"%s is not empty\", dir)\n\t}\n\tclient, err := apiClient()\n\tif err != nil {\n\t\treturn err\n\t}\n\troot, err := filepath.Abs(dir)\n\tif err != nil {\n\t\treturn err\n\t}\n\tif err := os.MkdirAll(root, 0o755); err != nil {\n\t\treturn err\n\t}\n\tstore := &record.FilesStore{Root: root, Dir: \"provenance\"}\n\tn, err := remote.CloneRecords(cmd.Context(), client, store, ids)\n\tif err != nil {\n\t\treturn err\n\t}\n\tif err := os.WriteFile(filepath.Join(root, configFile), []byte(initConfig), 0o644); err != nil {\n\t\treturn err\n\t}\n\tlines, err := record.BuildManifest(store)\n\tif err != nil {\n\t\treturn err\n\t}\n\tif err := os.MkdirAll(filepath.Join(root, filepath.Dir(manifestPath)), 0o755); err != nil {\n\t\treturn err\n\t}\n\tif err := os.WriteFile(filepath.Join(root, manifestPath), record.MarshalManifest(lines), 0o644); err != nil {\n\t\treturn err\n\t}\n\tif _, err := ensureGitattributes(root, \"provenance\"); err != nil {\n\t\treturn err\n\t}\n\tfmt.Fprintf(out, \"\\n✓ Cloned %d record(s) into %s\\n  Every version was verified against its content hash.\\n  Spec files are not published: lint reports them until the repository's files are present.\\n  Run git init and commit when you are ready.\\n\\n\", n, dir)\n\treturn nil\n}","file":"cmd/provenance/clone_run.go","name":"runClone"}}
func runClone(out io.Writer, cmd *cobra.Command, args []string) error {
	dir, ids := args[0], args[1:]
	if entries, err := os.ReadDir(dir); err == nil && len(entries) > 0 {
		return fmt.Errorf("%s is not empty", dir)
	}
	client, err := apiClient()
	if err != nil {
		return err
	}
	root, err := filepath.Abs(dir)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(root, 0o755); err != nil {
		return err
	}
	store := &record.FilesStore{Root: root, Dir: "provenance"}
	n, err := remote.CloneRecords(cmd.Context(), client, store, ids)
	if err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(root, configFile), []byte(initConfig), 0o644); err != nil {
		return err
	}
	lines, err := record.BuildManifest(store)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Join(root, filepath.Dir(manifestPath)), 0o755); err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(root, manifestPath), record.MarshalManifest(lines), 0o644); err != nil {
		return err
	}
	if _, err := ensureGitattributes(root, "provenance"); err != nil {
		return err
	}
	fmt.Fprintf(out, "\n✓ Cloned %d record(s) into %s\n  Every version was verified against its content hash.\n  Spec files are not published: lint reports them until the repository's files are present.\n  Run git init and commit when you are ready.\n\n", n, dir)
	return nil
}

// /sedum:setFunction
