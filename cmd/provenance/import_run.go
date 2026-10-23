package main

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-98753860","gen":"sha256:e2ab13cad210f5eb275ac9de2f1efc5b1ad105c8886fe10c86c00c7e4a57de4a","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/import_run.go","path":"fmt"}}
import "fmt"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-98753860","gen":"sha256:e2ab13cad210f5eb275ac9de2f1efc5b1ad105c8886fe10c86c00c7e4a57de4a","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/import_run.go","path":"io"}}
import "io"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-98753860","gen":"sha256:e2ab13cad210f5eb275ac9de2f1efc5b1ad105c8886fe10c86c00c7e4a57de4a","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/import_run.go","path":"strconv"}}
import "strconv"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-98753860","gen":"sha256:e2ab13cad210f5eb275ac9de2f1efc5b1ad105c8886fe10c86c00c7e4a57de4a","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/import_run.go","path":"github.com/spf13/cobra"}}
import "github.com/spf13/cobra"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-98753860","gen":"sha256:e2ab13cad210f5eb275ac9de2f1efc5b1ad105c8886fe10c86c00c7e4a57de4a","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/import_run.go","path":"github.com/livecodelife/provenance/remote"}}
import "github.com/livecodelife/provenance/remote"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"tier":"owned","record":"prov-2026-98753860","gen":"sha256:e2ab13cad210f5eb275ac9de2f1efc5b1ad105c8886fe10c86c00c7e4a57de4a","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func runImport(out io.Writer, cmd *cobra.Command, args []string) error {\n\tall, _ := cmd.Flags().GetBool(\"all\")\n\tr, err := openRepo(cmd)\n\tif err != nil {\n\t\treturn err\n\t}\n\tids, err := recordIDs(r, args, all)\n\tif err != nil {\n\t\treturn err\n\t}\n\tclient, err := apiClient()\n\tif err != nil {\n\t\treturn err\n\t}\n\tplan, err := remote.PlanImport(cmd.Context(), r.Store, ids, client)\n\tif err != nil {\n\t\treturn err\n\t}\n\tfor _, id := range plan.Refused {\n\t\tr.fail(\"Not imported: %s\", id)\n\t}\n\tif len(plan.Items) == 0 {\n\t\tif len(plan.Refused) > 0 {\n\t\t\treturn fmt.Errorf(\"%d record(s) refused\", len(plan.Refused))\n\t\t}\n\t\tfmt.Fprintln(out)\n\t\tfmt.Fprintln(out, \"Nothing to import: every published version is already here.\")\n\t\tfmt.Fprintln(out)\n\t\treturn nil\n\t}\n\tif err := remote.ApplyImport(plan, r.Store); err != nil {\n\t\treturn err\n\t}\n\tif err := r.writeManifest(); err != nil {\n\t\treturn err\n\t}\n\tfmt.Fprintln(out)\n\tfmt.Fprintln(out, \"✓ Imported:\")\n\tfmt.Fprintln(out)\n\tfor _, item := range plan.Items {\n\t\tfor i, s := range item.Snaps {\n\t\t\torigin := item.Origins[i]\n\t\t\tif origin == \"\" {\n\t\t\t\torigin = \"git\"\n\t\t\t}\n\t\t\tfmt.Fprintf(out, \"  %s v%s  %s  (%s), %s\\n\", item.ID, strconv.Itoa(s.Version), s.Reason, s.Author, origin)\n\t\t}\n\t}\n\tfmt.Fprintln(out)\n\tfmt.Fprintf(out, \"  Commit the changed record files, %s/.versions and %s.\\n\", r.Config.Dir, manifestPath)\n\tfmt.Fprintln(out)\n\tif len(plan.Refused) > 0 {\n\t\treturn fmt.Errorf(\"%d record(s) refused\", len(plan.Refused))\n\t}\n\treturn nil\n}","file":"cmd/provenance/import_run.go","name":"runImport"}}
func runImport(out io.Writer, cmd *cobra.Command, args []string) error {
	all, _ := cmd.Flags().GetBool("all")
	r, err := openRepo(cmd)
	if err != nil {
		return err
	}
	ids, err := recordIDs(r, args, all)
	if err != nil {
		return err
	}
	client, err := apiClient()
	if err != nil {
		return err
	}
	plan, err := remote.PlanImport(cmd.Context(), r.Store, ids, client)
	if err != nil {
		return err
	}
	for _, id := range plan.Refused {
		r.fail("Not imported: %s", id)
	}
	if len(plan.Items) == 0 {
		if len(plan.Refused) > 0 {
			return fmt.Errorf("%d record(s) refused", len(plan.Refused))
		}
		fmt.Fprintln(out)
		fmt.Fprintln(out, "Nothing to import: every published version is already here.")
		fmt.Fprintln(out)
		return nil
	}
	if err := remote.ApplyImport(plan, r.Store); err != nil {
		return err
	}
	if err := r.writeManifest(); err != nil {
		return err
	}
	fmt.Fprintln(out)
	fmt.Fprintln(out, "✓ Imported:")
	fmt.Fprintln(out)
	for _, item := range plan.Items {
		for i, s := range item.Snaps {
			origin := item.Origins[i]
			if origin == "" {
				origin = "git"
			}
			fmt.Fprintf(out, "  %s v%s  %s  (%s), %s\n", item.ID, strconv.Itoa(s.Version), s.Reason, s.Author, origin)
		}
	}
	fmt.Fprintln(out)
	fmt.Fprintf(out, "  Commit the changed record files, %s/.versions and %s.\n", r.Config.Dir, manifestPath)
	fmt.Fprintln(out)
	if len(plan.Refused) > 0 {
		return fmt.Errorf("%d record(s) refused", len(plan.Refused))
	}
	return nil
}

// /sedum:setFunction
