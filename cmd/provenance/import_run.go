package main

// sedum:anchor:imports
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-98753860","gen":"sha256:e2ab13cad210f5eb275ac9de2f1efc5b1ad105c8886fe10c86c00c7e4a57de4a","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/import_run.go","path":"fmt"},"kwargs":"sha256:4fab6fd64efd4e95a0d882d0cd04b537f886f73375a62b6adf36e85415834563","body":"sha256:bcbac51883cbf8064e9196b371cb880cee429866652447ad557533d817857609"}
import "fmt"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-98753860","gen":"sha256:e2ab13cad210f5eb275ac9de2f1efc5b1ad105c8886fe10c86c00c7e4a57de4a","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/import_run.go","path":"io"},"kwargs":"sha256:deaf7591c1a00ea08192bad3bfdf5865183d943e6b79c0ba95ffd8f2992a1311","body":"sha256:ef8e8abf4ef5ee40a19dc85b49bf09c5df1fe32abfe834e0ab4c2175baf80ddf"}
import "io"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-98753860","gen":"sha256:e2ab13cad210f5eb275ac9de2f1efc5b1ad105c8886fe10c86c00c7e4a57de4a","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/import_run.go","path":"strconv"},"kwargs":"sha256:c2eefb96f2a2bc4bb01cd88697f23d88c3f1fff3107cb685cc40d6543f44edba","body":"sha256:f33f981c64a42d78a5daa7ff3fca15b10045080cccd282fef544b59eaf7f6dc8"}
import "strconv"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-98753860","gen":"sha256:e2ab13cad210f5eb275ac9de2f1efc5b1ad105c8886fe10c86c00c7e4a57de4a","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/import_run.go","path":"github.com/spf13/cobra"},"kwargs":"sha256:519f4df7b7a757c9a099862e086a9c479ca90c70333be913fd8885d7dce7fd35","body":"sha256:f361c02d0f8e1999cf3fb715393e16310eeec9debd5aaa13936af7617e2b46aa"}
import "github.com/spf13/cobra"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-98753860","gen":"sha256:e2ab13cad210f5eb275ac9de2f1efc5b1ad105c8886fe10c86c00c7e4a57de4a","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/import_run.go","path":"github.com/livecodelife/provenance/remote"},"kwargs":"sha256:f6e51b7afb60f8347ca6ef8e35a229dac0677fdc520974c57a680c15a8136de1","body":"sha256:8d05ee0496574b6404cf7e9fbb5899ea912303f93ead4d7ded94d22f09a229d2"}
import "github.com/livecodelife/provenance/remote"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-98753860","gen":"sha256:e2ab13cad210f5eb275ac9de2f1efc5b1ad105c8886fe10c86c00c7e4a57de4a","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/import_run.go","name":"runImport"},"kwargs":"sha256:68ff31bdb270a3eab6f55a33ba03cf55a18ca323362c6e04e19308fc7472c347","body":"sha256:e9fa05b0cbb6a2ab5661d382a6ceea10f9ab640096ac63376a8186ddc5a0c59b"}
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
