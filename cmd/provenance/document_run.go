package main

// sedum:anchor:imports
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-57558206","gen":"sha256:5cbafa9abb7457bc939ce881ae8f9409398722bf2fd07340b8b1ea02d16eec57","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/document_run.go","path":"fmt"},"kwargs":"sha256:526eede958049d9cf7bfe2f9b7e6b492b109ab7e8d46cb62c24edc8114bfaec0","body":"sha256:bcbac51883cbf8064e9196b371cb880cee429866652447ad557533d817857609"}
import "fmt"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-57558206","gen":"sha256:5cbafa9abb7457bc939ce881ae8f9409398722bf2fd07340b8b1ea02d16eec57","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/document_run.go","path":"github.com/livecodelife/provenance/document"},"kwargs":"sha256:623190c6a9a115ff64ea73d3e6ff7d7c99ab82e2b4e330ed437533ca22ae13f5","body":"sha256:a04e07b5db3032f088d5a86d2cb05b47d2e9c82d3f4a629bd1e9bd0ee5b5462f"}
import "github.com/livecodelife/provenance/document"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-cabce335","gen":"sha256:3284d5dff3d2c8851d461b1cfea5a711d815b224d97ad99c9b6c0ba605f4fbd9","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/document_run.go","path":"errors"},"kwargs":"sha256:0ded26aafcf81796386227db06cabf321ab5ddbc9cfd6f5d7ad30974c5d3ca05","body":"sha256:2b718237ba0a56467c8fd494cc131e942ef410e6228c03b644d70476c71862dc"}
import "errors"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-cabce335","gen":"sha256:3284d5dff3d2c8851d461b1cfea5a711d815b224d97ad99c9b6c0ba605f4fbd9","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/document_run.go","path":"io"},"kwargs":"sha256:b6aec718d8d3466dffa797bbffc4df32a638c490996b202618f13a399345613c","body":"sha256:ef8e8abf4ef5ee40a19dc85b49bf09c5df1fe32abfe834e0ab4c2175baf80ddf"}
import "io"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-cabce335","gen":"sha256:3284d5dff3d2c8851d461b1cfea5a711d815b224d97ad99c9b6c0ba605f4fbd9","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/document_run.go","path":"os"},"kwargs":"sha256:c1b75071a83753e890584a73151b22fe86a2ebb01241a15d4845e102301030d3","body":"sha256:7c6e309eb2575d4aa451ec2ab83161d13e9ae03bbca3a4d2ecf30ce12adf0905"}
import "os"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-cabce335","gen":"sha256:3284d5dff3d2c8851d461b1cfea5a711d815b224d97ad99c9b6c0ba605f4fbd9","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/document_run.go","path":"path/filepath"},"kwargs":"sha256:e46fad4e63a174ea2283e21f7ad2997579268db9ba7b7422774703ef7cec48a6","body":"sha256:a6d196b558491e646be93f56a321f339068c64c76aa9d49bb18046e7f9c777a1"}
import "path/filepath"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-cabce335","gen":"sha256:3284d5dff3d2c8851d461b1cfea5a711d815b224d97ad99c9b6c0ba605f4fbd9","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/document_run.go","path":"strings"},"kwargs":"sha256:59073da30d9a1a28ce7df057cc65ee56e1fac699306eabd694fcee643912e595","body":"sha256:a9af3c0bc4bf9342df9d7e89a324523fdba7941a9ca4190270e68ff3c4a25ea6"}
import "strings"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-cabce335","gen":"sha256:3284d5dff3d2c8851d461b1cfea5a711d815b224d97ad99c9b6c0ba605f4fbd9","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/document_run.go","path":"github.com/spf13/cobra"},"kwargs":"sha256:4f99f5bfe3f4e63aacd9adc0b988dce2b4ed32f8d5d0882e96a5138a3af51632","body":"sha256:f361c02d0f8e1999cf3fb715393e16310eeec9debd5aaa13936af7617e2b46aa"}
import "github.com/spf13/cobra"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-194d88b9","gen":"sha256:a5020d83de1579e79b5aeeada237b0fd303668f1c24d9c5c9117dc77eea87d97","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/document_run.go","name":"validateDocumentFlags"},"kwargs":"sha256:116846e7a12548e2cf73631a2c2ffc9e9e18b69156fca1dba64bd7748b7e067c","body":"sha256:1bc504ede651bbda18726effce43e2beb12eef76a0ff8b17f47e47bca2cc079a"}
func validateDocumentFlags(format string, scope []string, out string, ids bool, apply string, write bool, reason string, skipConflicts bool) error {
	if format != "md" {
		return fmt.Errorf("unsupported format %q: only md is accepted", format)
	}
	if apply != "" && (len(scope) > 0 || ids || out != "") {
		return errors.New("--apply cannot be combined with --scope, --ids or -o")
	}
	if write && apply == "" {
		return errors.New("--write needs --apply")
	}
	if reason != "" && !write {
		return errors.New("--reason needs --write")
	}
	if skipConflicts && !write {
		return errors.New("--skip-conflicts needs --write")
	}
	if write && strings.TrimSpace(reason) == "" {
		return errors.New("--write needs a non-empty --reason")
	}
	return nil
}

// /sedum:setFunction
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-213e080d","gen":"sha256:95112ceaa3f8f81c59b6d19bfba1cdcb38cd5739c29ce0a076a76f3dbe5d2afd","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/document_run.go","name":"afterApply"},"kwargs":"sha256:5a99d88d61d4d6a44bb64f43aa4c1d853b9ceec6651d4fff62a6fb112d1be5a5","body":"sha256:fd9746dc9f5d2aabfcc677dca0d144956280aa55bb6f8f95c62c0c292cb4b7b1"}
func afterApply(r *repo, states []fileState, revised []document.Revised) error {
	if len(revised) == 0 {
		return nil
	}
	if err := r.writeManifest(); err != nil {
		restoreFiles(states)
		return err
	}
	reconcileQuietly(r)
	return nil
}

// /sedum:setFunction
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-57558206","gen":"sha256:5cbafa9abb7457bc939ce881ae8f9409398722bf2fd07340b8b1ea02d16eec57","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/document_run.go","name":"planError"},"kwargs":"sha256:e018155545a918c3f293f57dd374f7e5c09529fb2f5906df8f6cab7f814785a6","body":"sha256:1e2d87f91172c7759df565893cdebb42cf991dbe65ade861f0158dd85130cf97"}
func planError(p *document.Plan) error {
	if len(p.Conflicts) == 0 && len(p.Problems) == 0 {
		return nil
	}
	return fmt.Errorf("document plan has %d conflicts and %d problems", len(p.Conflicts), len(p.Problems))
}

// /sedum:setFunction
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-a8861e31","gen":"sha256:dbc9c4ff0141e987a9e8889d567f406a848c9b1b696819d7eff8b3067e9c215f","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/document_run.go","name":"printRevised"},"kwargs":"sha256:7a8c8488e143e05f9900d2fdc95d1fc51b5b2fdea0c1feb8f5eef796c6ea43b4","body":"sha256:a00524281806ef306f3b90c463cd27d732bde8037c5ad5e06d3bde3e5bab6145"}
func printRevised(w io.Writer, revs []document.Revised) {
	for _, r := range revs {
		if r.Draft {
			fmt.Fprintf(w, "saved %s (draft)\n", r.ID)
			continue
		}
		fmt.Fprintf(w, "revised %s v%d %s\n", r.ID, r.Version, tag(r.ID, r.Version))
	}
}

// /sedum:setFunction
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-b0549209","gen":"sha256:f750c643d9a2f283194a28ad154d124b43b3bb606724dcb9e1bca4175cfae6c2","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/document_run.go","name":"doRender"},"kwargs":"sha256:57a4e7f949958dda5ee27746c7f505b2af6a803b18a92809dbda93db5978106e","body":"sha256:906a30ff4cccff204784c6edc0371151bf5f4922781c0d8fe3f97310d4eeee94"}
func doRender(w io.Writer, r *repo, scopes []string, ids bool, out string) error {
	text, err := document.Render(r.Store, document.RenderOptions{Scopes: scopes, IDs: ids})
	if err != nil {
		return err
	}
	if out != "" {
		if err := os.WriteFile(out, []byte(text), 0644); err != nil {
			return fmt.Errorf("cannot write %s: %w", out, err)
		}
		return nil
	}
	fmt.Fprint(w, text)
	return nil
}

// /sedum:setFunction
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-b141d7f7","gen":"sha256:ce3da0906b1f5fcf44881b2a2c89bb40db16f645f3372cbffe7905a269fcf77d","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/document_run.go","name":"doApply"},"kwargs":"sha256:576d34f64c63651991f7f4fb9e15b52e53e71bc29ab3c052f53f6ab4d017f34f","body":"sha256:fa226bc52aafcc1158401707bc225b2f6d51e85e7a667a949c255f25de274429"}
func doApply(out io.Writer, r *repo, path string, write bool, reason string, skip bool) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("cannot read %s: %v", path, err)
	}
	plan, err := document.BuildPlan(r.Store, string(data))
	if err != nil {
		return err
	}
	printDocumentPlan(out, plan)
	if !write {
		return planError(plan)
	}
	return applyPlan(out, r, plan, reason, skip)
}

// /sedum:setFunction
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-cabce335","gen":"sha256:3284d5dff3d2c8851d461b1cfea5a711d815b224d97ad99c9b6c0ba605f4fbd9","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/document_run.go","name":"runDocument"},"kwargs":"sha256:ccb209d54f17f0503883b7015895bda16d7629606a421c177bf0f52a715f3875","body":"sha256:7f1b27d8f931a786d819a901aa743e6b4dcd838be728ac3b3f330d5c0e13e383"}
func runDocument(out io.Writer, cmd *cobra.Command, args []string) error {
	_ = errors.New
	_ = os.Stdout
	_ = filepath.Separator
	_ = strings.TrimSpace
	_ = args
	flags := cmd.Flags()
	format, _ := flags.GetString("format")
	scope, _ := flags.GetStringArray("scope")
	outPath, _ := flags.GetString("out")
	ids, _ := flags.GetBool("ids")
	apply, _ := flags.GetString("apply")
	write, _ := flags.GetBool("write")
	reason, _ := flags.GetString("reason")
	skipConflicts, _ := flags.GetBool("skip-conflicts")
	if err := validateDocumentFlags(format, scope, outPath, ids, apply, write, reason, skipConflicts); err != nil {
		return err
	}
	r, err := openRepo(cmd)
	if err != nil {
		return err
	}
	if apply == "" {
		return doRender(out, r, scope, ids, outPath)
	}
	return doApply(out, r, apply, write, reason, skipConflicts)
}

// /sedum:setFunction
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-e22c88da","gen":"sha256:6c39f344684f499c24737c37c2b19045a56f47bba7df1545573698ac0ad7e1ca","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/document_run.go","name":"applyPlan"},"kwargs":"sha256:716600501aca63edfbad115226fc34c975c0fe6b5d5aa6412c33b5548f530fee","body":"sha256:2e85070ddd679d3f3faec1b6aab5e64e8f17997af59f58f6cf4dafe0ca8426e6"}
func applyPlan(out io.Writer, r *repo, plan *document.Plan, reason string, skipConflicts bool) error {
	states := snapshotFiles(filepath.Join(r.Root, manifestPath))
	revised, applyErr := document.Apply(r.Store, plan, document.ApplyOptions{
		Reason:        reason,
		Author:        r.author(),
		SkipConflicts: skipConflicts,
	})
	if err := afterApply(r, states, revised); err != nil {
		return err
	}
	printRevised(out, revised)
	if applyErr != nil {
		return applyErr
	}
	return planError(plan)
}

// /sedum:setFunction
