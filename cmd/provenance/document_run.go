package main

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-57558206","gen":"sha256:5cbafa9abb7457bc939ce881ae8f9409398722bf2fd07340b8b1ea02d16eec57","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/document_run.go","path":"fmt"}}
import "fmt"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-57558206","gen":"sha256:5cbafa9abb7457bc939ce881ae8f9409398722bf2fd07340b8b1ea02d16eec57","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/document_run.go","path":"github.com/livecodelife/provenance/document"}}
import "github.com/livecodelife/provenance/document"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-cabce335","gen":"sha256:3284d5dff3d2c8851d461b1cfea5a711d815b224d97ad99c9b6c0ba605f4fbd9","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/document_run.go","path":"errors"}}
import "errors"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-cabce335","gen":"sha256:3284d5dff3d2c8851d461b1cfea5a711d815b224d97ad99c9b6c0ba605f4fbd9","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/document_run.go","path":"io"}}
import "io"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-cabce335","gen":"sha256:3284d5dff3d2c8851d461b1cfea5a711d815b224d97ad99c9b6c0ba605f4fbd9","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/document_run.go","path":"os"}}
import "os"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-cabce335","gen":"sha256:3284d5dff3d2c8851d461b1cfea5a711d815b224d97ad99c9b6c0ba605f4fbd9","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/document_run.go","path":"path/filepath"}}
import "path/filepath"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-cabce335","gen":"sha256:3284d5dff3d2c8851d461b1cfea5a711d815b224d97ad99c9b6c0ba605f4fbd9","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/document_run.go","path":"strings"}}
import "strings"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-cabce335","gen":"sha256:3284d5dff3d2c8851d461b1cfea5a711d815b224d97ad99c9b6c0ba605f4fbd9","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/document_run.go","path":"github.com/spf13/cobra"}}
import "github.com/spf13/cobra"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"tier":"owned","record":"prov-2026-194d88b9","gen":"sha256:a5020d83de1579e79b5aeeada237b0fd303668f1c24d9c5c9117dc77eea87d97","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func validateDocumentFlags(format string, scope []string, out string, ids bool, apply string, write bool, reason string, skipConflicts bool) error {\n\tif format != \"md\" {\n\t\treturn fmt.Errorf(\"unsupported format %q: only md is accepted\", format)\n\t}\n\tif apply != \"\" && (len(scope) > 0 || ids || out != \"\") {\n\t\treturn errors.New(\"--apply cannot be combined with --scope, --ids or -o\")\n\t}\n\tif write && apply == \"\" {\n\t\treturn errors.New(\"--write needs --apply\")\n\t}\n\tif reason != \"\" && !write {\n\t\treturn errors.New(\"--reason needs --write\")\n\t}\n\tif skipConflicts && !write {\n\t\treturn errors.New(\"--skip-conflicts needs --write\")\n\t}\n\tif write && strings.TrimSpace(reason) == \"\" {\n\t\treturn errors.New(\"--write needs a non-empty --reason\")\n\t}\n\treturn nil\n}","file":"cmd/provenance/document_run.go","name":"validateDocumentFlags"}}
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
// sedum:setFunction {"tier":"owned","record":"prov-2026-213e080d","gen":"sha256:95112ceaa3f8f81c59b6d19bfba1cdcb38cd5739c29ce0a076a76f3dbe5d2afd","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func afterApply(r *repo, states []fileState, revised []document.Revised) error {\n\tif len(revised) == 0 {\n\t\treturn nil\n\t}\n\tif err := r.writeManifest(); err != nil {\n\t\trestoreFiles(states)\n\t\treturn err\n\t}\n\treconcileQuietly(r)\n\treturn nil\n}","file":"cmd/provenance/document_run.go","name":"afterApply"}}
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
// sedum:setFunction {"tier":"owned","record":"prov-2026-57558206","gen":"sha256:5cbafa9abb7457bc939ce881ae8f9409398722bf2fd07340b8b1ea02d16eec57","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func planError(p *document.Plan) error {\n\tif len(p.Conflicts) == 0 && len(p.Problems) == 0 {\n\t\treturn nil\n\t}\n\treturn fmt.Errorf(\"document plan has %d conflicts and %d problems\", len(p.Conflicts), len(p.Problems))\n}","file":"cmd/provenance/document_run.go","name":"planError"}}
func planError(p *document.Plan) error {
	if len(p.Conflicts) == 0 && len(p.Problems) == 0 {
		return nil
	}
	return fmt.Errorf("document plan has %d conflicts and %d problems", len(p.Conflicts), len(p.Problems))
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-a8861e31","gen":"sha256:dbc9c4ff0141e987a9e8889d567f406a848c9b1b696819d7eff8b3067e9c215f","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func printRevised(w io.Writer, revs []document.Revised) {\n\tfor _, r := range revs {\n\t\tif r.Draft {\n\t\t\tfmt.Fprintf(w, \"saved %s (draft)\\n\", r.ID)\n\t\t\tcontinue\n\t\t}\n\t\tfmt.Fprintf(w, \"revised %s v%d %s\\n\", r.ID, r.Version, tag(r.ID, r.Version))\n\t}\n}","file":"cmd/provenance/document_run.go","name":"printRevised"}}
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
// sedum:setFunction {"tier":"owned","record":"prov-2026-b0549209","gen":"sha256:f750c643d9a2f283194a28ad154d124b43b3bb606724dcb9e1bca4175cfae6c2","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func doRender(w io.Writer, r *repo, scopes []string, ids bool, out string) error {\n\ttext, err := document.Render(r.Store, document.RenderOptions{Scopes: scopes, IDs: ids})\n\tif err != nil {\n\t\treturn err\n\t}\n\tif out != \"\" {\n\t\tif err := os.WriteFile(out, []byte(text), 0644); err != nil {\n\t\t\treturn fmt.Errorf(\"cannot write %s: %w\", out, err)\n\t\t}\n\t\treturn nil\n\t}\n\tfmt.Fprint(w, text)\n\treturn nil\n}","file":"cmd/provenance/document_run.go","name":"doRender"}}
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
// sedum:setFunction {"tier":"owned","record":"prov-2026-b141d7f7","gen":"sha256:ce3da0906b1f5fcf44881b2a2c89bb40db16f645f3372cbffe7905a269fcf77d","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func doApply(out io.Writer, r *repo, path string, write bool, reason string, skip bool) error {\n\tdata, err := os.ReadFile(path)\n\tif err != nil {\n\t\treturn fmt.Errorf(\"cannot read %s: %v\", path, err)\n\t}\n\tplan, err := document.BuildPlan(r.Store, string(data))\n\tif err != nil {\n\t\treturn err\n\t}\n\tprintDocumentPlan(out, plan)\n\tif !write {\n\t\treturn planError(plan)\n\t}\n\treturn applyPlan(out, r, plan, reason, skip)\n}","file":"cmd/provenance/document_run.go","name":"doApply"}}
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
// sedum:setFunction {"tier":"owned","record":"prov-2026-cabce335","gen":"sha256:3284d5dff3d2c8851d461b1cfea5a711d815b224d97ad99c9b6c0ba605f4fbd9","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func runDocument(out io.Writer, cmd *cobra.Command, args []string) error {\n\t_ = errors.New\n\t_ = os.Stdout\n\t_ = filepath.Separator\n\t_ = strings.TrimSpace\n\t_ = args\n\tflags := cmd.Flags()\n\tformat, _ := flags.GetString(\"format\")\n\tscope, _ := flags.GetStringArray(\"scope\")\n\toutPath, _ := flags.GetString(\"out\")\n\tids, _ := flags.GetBool(\"ids\")\n\tapply, _ := flags.GetString(\"apply\")\n\twrite, _ := flags.GetBool(\"write\")\n\treason, _ := flags.GetString(\"reason\")\n\tskipConflicts, _ := flags.GetBool(\"skip-conflicts\")\n\tif err := validateDocumentFlags(format, scope, outPath, ids, apply, write, reason, skipConflicts); err != nil {\n\t\treturn err\n\t}\n\tr, err := openRepo(cmd)\n\tif err != nil {\n\t\treturn err\n\t}\n\tif apply == \"\" {\n\t\treturn doRender(out, r, scope, ids, outPath)\n\t}\n\treturn doApply(out, r, apply, write, reason, skipConflicts)\n}","file":"cmd/provenance/document_run.go","name":"runDocument"}}
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
// sedum:setFunction {"tier":"owned","record":"prov-2026-e22c88da","gen":"sha256:6c39f344684f499c24737c37c2b19045a56f47bba7df1545573698ac0ad7e1ca","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func applyPlan(out io.Writer, r *repo, plan *document.Plan, reason string, skipConflicts bool) error {\n\tstates := snapshotFiles(filepath.Join(r.Root, manifestPath))\n\trevised, applyErr := document.Apply(r.Store, plan, document.ApplyOptions{\n\t\tReason:        reason,\n\t\tAuthor:        r.author(),\n\t\tSkipConflicts: skipConflicts,\n\t})\n\tif err := afterApply(r, states, revised); err != nil {\n\t\treturn err\n\t}\n\tprintRevised(out, revised)\n\tif applyErr != nil {\n\t\treturn applyErr\n\t}\n\treturn planError(plan)\n}","file":"cmd/provenance/document_run.go","name":"applyPlan"}}
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
