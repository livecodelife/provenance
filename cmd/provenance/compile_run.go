package main

// sedum:anchor:imports
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-d84dc10e","gen":"sha256:2e97312cea5e86a4f4e11ea65cacfe9c4a63da245a1c446b00a2c6740234e92a","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/compile_run.go","path":"bytes"},"kwargs":"sha256:d6d2d4dbd4b3a3e3708f887f6152869f987728d8b6ed161a41136bf92bf9414c","body":"sha256:468731f6dd513df0a66690d4649dc9e7244c96de227d8e6ead98cc5e9747c57a"}
import "bytes"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-d84dc10e","gen":"sha256:2e97312cea5e86a4f4e11ea65cacfe9c4a63da245a1c446b00a2c6740234e92a","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/compile_run.go","path":"fmt"},"kwargs":"sha256:dc758c6059ce568b9ab6a8e6b9e5ed2d38a3916fd985b068d88ee2beaf03d639","body":"sha256:bcbac51883cbf8064e9196b371cb880cee429866652447ad557533d817857609"}
import "fmt"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-d84dc10e","gen":"sha256:2e97312cea5e86a4f4e11ea65cacfe9c4a63da245a1c446b00a2c6740234e92a","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/compile_run.go","path":"io"},"kwargs":"sha256:591593ddf0445eb79356be70daa005166da27de550e2ed5cabc3ed3f7b34cc7d","body":"sha256:ef8e8abf4ef5ee40a19dc85b49bf09c5df1fe32abfe834e0ab4c2175baf80ddf"}
import "io"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-d84dc10e","gen":"sha256:2e97312cea5e86a4f4e11ea65cacfe9c4a63da245a1c446b00a2c6740234e92a","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/compile_run.go","path":"os"},"kwargs":"sha256:b069ffc5e77dde2a38b04afa3b5dd254f79574dab0170b76e511552fa993a1cd","body":"sha256:7c6e309eb2575d4aa451ec2ab83161d13e9ae03bbca3a4d2ecf30ce12adf0905"}
import "os"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-d84dc10e","gen":"sha256:2e97312cea5e86a4f4e11ea65cacfe9c4a63da245a1c446b00a2c6740234e92a","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/compile_run.go","path":"path/filepath"},"kwargs":"sha256:9a31d964fe839fb4afbfeea4eef3b6d91bb578277d321426ed36356efc768cc4","body":"sha256:a6d196b558491e646be93f56a321f339068c64c76aa9d49bb18046e7f9c777a1"}
import "path/filepath"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-d84dc10e","gen":"sha256:2e97312cea5e86a4f4e11ea65cacfe9c4a63da245a1c446b00a2c6740234e92a","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/compile_run.go","path":"github.com/spf13/cobra"},"kwargs":"sha256:4818c93585cfd3b136e0274af2a20a4cd8ec2543b65d79c58ad000c48f0f2905","body":"sha256:f361c02d0f8e1999cf3fb715393e16310eeec9debd5aaa13936af7617e2b46aa"}
import "github.com/spf13/cobra"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-d84dc10e","gen":"sha256:2e97312cea5e86a4f4e11ea65cacfe9c4a63da245a1c446b00a2c6740234e92a","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/compile_run.go","path":"github.com/livecodelife/provenance/record"},"kwargs":"sha256:af45d9acc966164c18c19aca4f8bf0a89fd4b58620ce1d1b1822e493be008e1e","body":"sha256:acbf8628061291d11aa39e9ea64ecc91095a117057cfe6f1958ef2c437fac003"}
import "github.com/livecodelife/provenance/record"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-d84dc10e","gen":"sha256:2e97312cea5e86a4f4e11ea65cacfe9c4a63da245a1c446b00a2c6740234e92a","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/compile_run.go","path":"errors"},"kwargs":"sha256:e5caf32f2b64ba481b74a5e7cd5849448016db4521120edf96ed348ad21faf89","body":"sha256:2b718237ba0a56467c8fd494cc131e942ef410e6228c03b644d70476c71862dc"}
import "errors"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-d84dc10e","gen":"sha256:2e97312cea5e86a4f4e11ea65cacfe9c4a63da245a1c446b00a2c6740234e92a","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/compile_run.go","name":"runCompile"},"kwargs":"sha256:449df4265c3deff5987b510230b08b256dfc6abdf9b57b1a79b7300fe698f4ef","body":"sha256:a8542fc9e21e087f2542772025a8e00a1b849116323578672c1176c47f145774"}
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
	if check, _ := cmd.Flags().GetBool("check"); check {
		fmt.Fprint(out, "\n✗ Version manifest is out of date. It is derived data; fix it with:\n\n  provenance compile\n  git add .provenance/versions.jsonl\n  git commit -m \"Recompile version manifest\"\n\n")
		return errors.New("version manifest is out of date")
	}
	if err := r.writeManifest(); err != nil {
		return err
	}
	fmt.Fprintf(out, "\n✓ Version manifest compiled (%d records)\n\n", len(lines))
	return nil
}

// /sedum:setFunction
