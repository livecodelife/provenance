package main

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-871605a9","gen":"sha256:a837b61e0b1282ffa2184bd15712bd692804fd3785f63fd676f9036654ebea2d","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/import_run.go","path":"fmt"}}
import "fmt"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-871605a9","gen":"sha256:a837b61e0b1282ffa2184bd15712bd692804fd3785f63fd676f9036654ebea2d","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/import_run.go","path":"io"}}
import "io"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-871605a9","gen":"sha256:a837b61e0b1282ffa2184bd15712bd692804fd3785f63fd676f9036654ebea2d","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/import_run.go","path":"strconv"}}
import "strconv"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-871605a9","gen":"sha256:a837b61e0b1282ffa2184bd15712bd692804fd3785f63fd676f9036654ebea2d","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/import_run.go","path":"github.com/spf13/cobra"}}
import "github.com/spf13/cobra"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-871605a9","gen":"sha256:a837b61e0b1282ffa2184bd15712bd692804fd3785f63fd676f9036654ebea2d","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/import_run.go","path":"github.com/livecodelife/provenance/remote"}}
import "github.com/livecodelife/provenance/remote"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"tier":"owned","record":"prov-2026-871605a9","gen":"sha256:a837b61e0b1282ffa2184bd15712bd692804fd3785f63fd676f9036654ebea2d","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func runImport(out io.Writer, cmd *cobra.Command, args []string) error {\n\tall, _ := cmd.Flags().GetBool(\"all\")\n\tr, err := openRepo(cmd)\n\tif err != nil {\n\t\treturn err\n\t}\n\tids, err := recordIDs(r, args, all)\n\tif err != nil {\n\t\treturn err\n\t}\n\tclient, err := apiClient()\n\tif err != nil {\n\t\treturn err\n\t}\n\tplan, err := remote.PlanImport(cmd.Context(), r.Store, ids, client)\n\tif err != nil {\n\t\treturn err\n\t}\n\tfor _, msg := range plan.Refused {\n\t\tr.fail(\"Not imported: %s\", msg)\n\t}\n\tif len(plan.Items) == 0 {\n\t\tif len(plan.Refused) > 0 {\n\t\t\treturn fmt.Errorf(\"%d record(s) refused\", len(plan.Refused))\n\t\t}\n\t\tfmt.Fprintln(out, \"\\nNothing to import: every published version is already here.\")\n\t\tfmt.Fprintln(out)\n\t\treturn nil\n\t}\n\tif err := remote.ApplyImport(plan, r.Store); err != nil {\n\t\treturn err\n\t}\n\tif err := r.writeManifest(); err != nil {\n\t\treturn err\n\t}\n\tfmt.Fprintf(out, \"\\n✓ Imported:\\n\\n\")\n\tfor _, it := range plan.Items {\n\t\tfor i, s := range it.Snaps {\n\t\t\torigin := it.Origins[i]\n\t\t\tif origin == \"\" {\n\t\t\t\torigin = \"git\"\n\t\t\t}\n\t\t\tfmt.Fprintf(out, \"  %s v%s  %s  (%s, %s)\\n\", it.ID, strconv.Itoa(s.Version), s.Reason, s.Author, origin)\n\t\t}\n\t}\n\tfmt.Fprintf(out, \"\\n  Commit the changed record files, %s and %s.\\n\\n\", r.Config.Dir+\"/.versions\", manifestPath)\n\tif len(plan.Refused) > 0 {\n\t\treturn fmt.Errorf(\"%d record(s) refused\", len(plan.Refused))\n\t}\n\treturn nil\n}","file":"cmd/provenance/import_run.go","name":"runImport"}}
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
	for _, msg := range plan.Refused {
		r.fail("Not imported: %s", msg)
	}
	if len(plan.Items) == 0 {
		if len(plan.Refused) > 0 {
			return fmt.Errorf("%d record(s) refused", len(plan.Refused))
		}
		fmt.Fprintln(out, "\nNothing to import: every published version is already here.")
		fmt.Fprintln(out)
		return nil
	}
	if err := remote.ApplyImport(plan, r.Store); err != nil {
		return err
	}
	if err := r.writeManifest(); err != nil {
		return err
	}
	fmt.Fprintf(out, "\n✓ Imported:\n\n")
	for _, it := range plan.Items {
		for i, s := range it.Snaps {
			origin := it.Origins[i]
			if origin == "" {
				origin = "git"
			}
			fmt.Fprintf(out, "  %s v%s  %s  (%s, %s)\n", it.ID, strconv.Itoa(s.Version), s.Reason, s.Author, origin)
		}
	}
	fmt.Fprintf(out, "\n  Commit the changed record files, %s and %s.\n\n", r.Config.Dir+"/.versions", manifestPath)
	if len(plan.Refused) > 0 {
		return fmt.Errorf("%d record(s) refused", len(plan.Refused))
	}
	return nil
}

// /sedum:setFunction
