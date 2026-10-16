package main

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-83d57ec1","gen":"sha256:6b4d5a3926de0813e7bebfa8498ef520429d4b6b530252ea1eb108e16c6d8def","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/publish_run.go","path":"fmt"}}
import "fmt"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-83d57ec1","gen":"sha256:6b4d5a3926de0813e7bebfa8498ef520429d4b6b530252ea1eb108e16c6d8def","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/publish_run.go","path":"io"}}
import "io"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-83d57ec1","gen":"sha256:6b4d5a3926de0813e7bebfa8498ef520429d4b6b530252ea1eb108e16c6d8def","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/publish_run.go","path":"strconv"}}
import "strconv"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-83d57ec1","gen":"sha256:6b4d5a3926de0813e7bebfa8498ef520429d4b6b530252ea1eb108e16c6d8def","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/publish_run.go","path":"strings"}}
import "strings"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-83d57ec1","gen":"sha256:6b4d5a3926de0813e7bebfa8498ef520429d4b6b530252ea1eb108e16c6d8def","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/publish_run.go","path":"github.com/spf13/cobra"}}
import "github.com/spf13/cobra"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-83d57ec1","gen":"sha256:6b4d5a3926de0813e7bebfa8498ef520429d4b6b530252ea1eb108e16c6d8def","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/publish_run.go","path":"github.com/livecodelife/provenance/remote"}}
import "github.com/livecodelife/provenance/remote"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"tier":"owned","record":"prov-2026-83d57ec1","gen":"sha256:6b4d5a3926de0813e7bebfa8498ef520429d4b6b530252ea1eb108e16c6d8def","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func runPublish(out io.Writer, cmd *cobra.Command, args []string) error {\n\tall, _ := cmd.Flags().GetBool(\"all\")\n\tyes, _ := cmd.Flags().GetBool(\"yes\")\n\tproject, _ := cmd.Flags().GetString(\"project\")\n\tif project != \"\" && !strings.HasPrefix(project, \"at://\") {\n\t\treturn fmt.Errorf(\"--project must be an at:// URI\")\n\t}\n\tr, err := openRepo(cmd)\n\tif err != nil {\n\t\treturn err\n\t}\n\tids, err := recordIDs(r, args, all)\n\tif err != nil {\n\t\treturn err\n\t}\n\tclient, err := apiClient()\n\tif err != nil {\n\t\treturn err\n\t}\n\tctx := cmd.Context()\n\tplan, err := remote.PlanPublish(ctx, r.Store, ids, client)\n\tif err != nil {\n\t\treturn err\n\t}\n\tfor _, s := range plan.Skipped {\n\t\tfmt.Fprintf(out, \"  skipped  %s\\n\", s)\n\t}\n\tif plan.Empty() {\n\t\tfmt.Fprintln(out, \"\\nNothing to publish: the server already has every version and status.\")\n\t\tfmt.Fprintln(out)\n\t\treturn nil\n\t}\n\tfmt.Fprintf(out, \"\\nThese become PUBLIC on AT Proto (anyone can read them):\\n\\n\")\n\tfor _, it := range plan.Items {\n\t\tif len(it.Snaps) == 0 {\n\t\t\tfmt.Fprintf(out, \"  %s  %s\\n      status → %s (no new version)\\n\", it.ID, it.Title, it.HeadStatus)\n\t\t\tcontinue\n\t\t}\n\t\tnums := make([]string, len(it.Snaps))\n\t\tfor i, s := range it.Snaps {\n\t\t\tnums[i] = \"v\" + strconv.Itoa(s.Version)\n\t\t}\n\t\tfmt.Fprintf(out, \"  %s  %s\\n      versions %s, status %s\\n\", it.ID, it.Title, strings.Join(nums, \", \"), it.Status)\n\t}\n\tfmt.Fprintln(out)\n\tif !yes && !confirm(cmd, out, \"Publish these? [y/N] \") {\n\t\treturn fmt.Errorf(\"not confirmed, nothing was published (run interactively, or pass --yes)\")\n\t}\n\tcommitOf := func(id string, version int) string {\n\t\tsha, err := r.git(\"log\", \"--all\", \"-E\", \"-n\", \"1\", \"--format=%H\", \"--grep\", `\\[`+id+`@`+strconv.Itoa(version)+`\\]`)\n\t\tif err != nil {\n\t\t\treturn \"\"\n\t\t}\n\t\treturn strings.TrimSpace(sha)\n\t}\n\twritten := 0\n\terr = remote.ApplyPublish(ctx, plan, client, project, commitOf, func(id string, version int) {\n\t\twritten++\n\t\tif version == 0 {\n\t\t\tfmt.Fprintf(out, \"  published  %s status\\n\", id)\n\t\t\treturn\n\t\t}\n\t\tfmt.Fprintf(out, \"  published  %s v%d\\n\", id, version)\n\t})\n\tif err != nil {\n\t\treturn err\n\t}\n\tfmt.Fprintf(out, \"\\n✓ Published %d write(s).\\n\\n\", written)\n\treturn nil\n}","file":"cmd/provenance/publish_run.go","name":"runPublish"}}
func runPublish(out io.Writer, cmd *cobra.Command, args []string) error {
	all, _ := cmd.Flags().GetBool("all")
	yes, _ := cmd.Flags().GetBool("yes")
	project, _ := cmd.Flags().GetString("project")
	if project != "" && !strings.HasPrefix(project, "at://") {
		return fmt.Errorf("--project must be an at:// URI")
	}
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
	ctx := cmd.Context()
	plan, err := remote.PlanPublish(ctx, r.Store, ids, client)
	if err != nil {
		return err
	}
	for _, s := range plan.Skipped {
		fmt.Fprintf(out, "  skipped  %s\n", s)
	}
	if plan.Empty() {
		fmt.Fprintln(out, "\nNothing to publish: the server already has every version and status.")
		fmt.Fprintln(out)
		return nil
	}
	fmt.Fprintf(out, "\nThese become PUBLIC on AT Proto (anyone can read them):\n\n")
	for _, it := range plan.Items {
		if len(it.Snaps) == 0 {
			fmt.Fprintf(out, "  %s  %s\n      status → %s (no new version)\n", it.ID, it.Title, it.HeadStatus)
			continue
		}
		nums := make([]string, len(it.Snaps))
		for i, s := range it.Snaps {
			nums[i] = "v" + strconv.Itoa(s.Version)
		}
		fmt.Fprintf(out, "  %s  %s\n      versions %s, status %s\n", it.ID, it.Title, strings.Join(nums, ", "), it.Status)
	}
	fmt.Fprintln(out)
	if !yes && !confirm(cmd, out, "Publish these? [y/N] ") {
		return fmt.Errorf("not confirmed, nothing was published (run interactively, or pass --yes)")
	}
	commitOf := func(id string, version int) string {
		sha, err := r.git("log", "--all", "-E", "-n", "1", "--format=%H", "--grep", `\[`+id+`@`+strconv.Itoa(version)+`\]`)
		if err != nil {
			return ""
		}
		return strings.TrimSpace(sha)
	}
	written := 0
	err = remote.ApplyPublish(ctx, plan, client, project, commitOf, func(id string, version int) {
		written++
		if version == 0 {
			fmt.Fprintf(out, "  published  %s status\n", id)
			return
		}
		fmt.Fprintf(out, "  published  %s v%d\n", id, version)
	})
	if err != nil {
		return err
	}
	fmt.Fprintf(out, "\n✓ Published %d write(s).\n\n", written)
	return nil
}

// /sedum:setFunction
