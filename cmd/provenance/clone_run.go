package main

// sedum:anchor:imports
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-6c411bf8","gen":"sha256:51691496e6409a36b3805de7a143010c350733f87e88eddd1891639d22b9cbb0","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/clone_run.go","path":"fmt"},"kwargs":"sha256:1c88f8d8b3c4c056f7c0d72880031b609dc1696b83d8b040b72b5b62232f27c9","body":"sha256:bcbac51883cbf8064e9196b371cb880cee429866652447ad557533d817857609"}
import "fmt"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-6c411bf8","gen":"sha256:51691496e6409a36b3805de7a143010c350733f87e88eddd1891639d22b9cbb0","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/clone_run.go","path":"io"},"kwargs":"sha256:4d004b6654764c20350db06c3a78f807f475702c1f5602d490c34b2d3f5d02b4","body":"sha256:ef8e8abf4ef5ee40a19dc85b49bf09c5df1fe32abfe834e0ab4c2175baf80ddf"}
import "io"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-6c411bf8","gen":"sha256:51691496e6409a36b3805de7a143010c350733f87e88eddd1891639d22b9cbb0","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/clone_run.go","path":"os"},"kwargs":"sha256:d3d5351d3bbd8b192b9a4fd8fd253984b7465558c6594bd20326d1105cba7a9b","body":"sha256:7c6e309eb2575d4aa451ec2ab83161d13e9ae03bbca3a4d2ecf30ce12adf0905"}
import "os"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-6c411bf8","gen":"sha256:51691496e6409a36b3805de7a143010c350733f87e88eddd1891639d22b9cbb0","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/clone_run.go","path":"path/filepath"},"kwargs":"sha256:a83dcd8f07f440401edf4256a1bf6fe5cd20c660f3f895b0e23a73971b1c4be7","body":"sha256:a6d196b558491e646be93f56a321f339068c64c76aa9d49bb18046e7f9c777a1"}
import "path/filepath"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-6c411bf8","gen":"sha256:51691496e6409a36b3805de7a143010c350733f87e88eddd1891639d22b9cbb0","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/clone_run.go","path":"github.com/spf13/cobra"},"kwargs":"sha256:738eef5ddede5a54fd50ac646031df8da7dffda5923ce5a09db0baaedaf44220","body":"sha256:f361c02d0f8e1999cf3fb715393e16310eeec9debd5aaa13936af7617e2b46aa"}
import "github.com/spf13/cobra"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-6c411bf8","gen":"sha256:51691496e6409a36b3805de7a143010c350733f87e88eddd1891639d22b9cbb0","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/clone_run.go","path":"github.com/livecodelife/provenance/record"},"kwargs":"sha256:21a45b30cd4805747740b29a03aabdd5c94837e8029515d23c9e95fc9189c411","body":"sha256:acbf8628061291d11aa39e9ea64ecc91095a117057cfe6f1958ef2c437fac003"}
import "github.com/livecodelife/provenance/record"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-6c411bf8","gen":"sha256:51691496e6409a36b3805de7a143010c350733f87e88eddd1891639d22b9cbb0","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/clone_run.go","path":"github.com/livecodelife/provenance/remote"},"kwargs":"sha256:4c62aac848cb644c17f043dfcf335ba3b638d94c6aa90cdf4272507fa24caf79","body":"sha256:8d05ee0496574b6404cf7e9fbb5899ea912303f93ead4d7ded94d22f09a229d2"}
import "github.com/livecodelife/provenance/remote"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-6c411bf8","gen":"sha256:51691496e6409a36b3805de7a143010c350733f87e88eddd1891639d22b9cbb0","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/clone_run.go","name":"runClone"},"kwargs":"sha256:866413575e632430033392d1fb0301d2de7e34f798ba5469e36d87219168416e","body":"sha256:7aec879e8e1ce883179755f7cab119246c326911a851243396a4f60ddba3d74e"}
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
