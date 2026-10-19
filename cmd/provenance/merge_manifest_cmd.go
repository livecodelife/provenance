package main

import "github.com/spf13/cobra"

func newMergeManifestCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"tier":"owned","record":"prov-2026-d84dc10e","gen":"sha256:2e97312cea5e86a4f4e11ea65cacfe9c4a63da245a1c446b00a2c6740234e92a","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","kwargs":{"bin":"provenance","command":"merge-manifest","short":"Merge three versions of the version manifest (git merge driver: BASE OURS THEIRS)","use":"merge-manifest"}}
	cmd.Use = "merge-manifest"
	cmd.Short = "Merge three versions of the version manifest (git merge driver: BASE OURS THEIRS)"
	// /sedum:describeCommand
	// sedum:setRun {"tier":"owned","record":"prov-2026-d84dc10e","gen":"sha256:2e97312cea5e86a4f4e11ea65cacfe9c4a63da245a1c446b00a2c6740234e92a","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","kwargs":{"bin":"provenance","command":"merge-manifest"}}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runMergeManifest(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:any {"tier":"owned","record":"prov-2026-d84dc10e","gen":"sha256:2e97312cea5e86a4f4e11ea65cacfe9c4a63da245a1c446b00a2c6740234e92a","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","kwargs":{"arity":"any","bin":"provenance","command":"merge-manifest"}}
	cmd.Args = cobra.ArbitraryArgs
	// /sedum:setArgs:any
	// sedum:anchor:flags
	// sedum:anchor:subcommands
	return cmd
}
