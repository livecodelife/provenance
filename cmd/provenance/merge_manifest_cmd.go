package main

import "github.com/spf13/cobra"

func newMergeManifestCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"v":2,"tier":"owned","record":"prov-2026-d84dc10e","gen":"sha256:2e97312cea5e86a4f4e11ea65cacfe9c4a63da245a1c446b00a2c6740234e92a","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","id":{"bin":"provenance","command":"merge-manifest"},"kwargs":"sha256:81942f975968594dcb5225a9d6477fad69ca19b6e7d075019f4d7c0f8c60fecd","body":"sha256:9792410a50e54b5cc1cf1b3c5fee17e1f5cdc01160680854c040bd9543f1e009"}
	cmd.Use = "merge-manifest"
	cmd.Short = "Merge three versions of the version manifest (git merge driver: BASE OURS THEIRS)"
	// /sedum:describeCommand
	// sedum:setRun {"v":2,"tier":"owned","record":"prov-2026-d84dc10e","gen":"sha256:2e97312cea5e86a4f4e11ea65cacfe9c4a63da245a1c446b00a2c6740234e92a","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","id":{"bin":"provenance","command":"merge-manifest"},"kwargs":"sha256:897c107fa37d5c660df2ccc483d2118ad5f1bcf5a7858ac6747201b3b148c8b1","body":"sha256:9b83dd21e1a6d4d0fa06b1c8b510a2b911dbdf0d6f40b644e96ad5c3b04ac912"}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runMergeManifest(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:any {"v":2,"tier":"owned","record":"prov-2026-d84dc10e","gen":"sha256:2e97312cea5e86a4f4e11ea65cacfe9c4a63da245a1c446b00a2c6740234e92a","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","id":{"bin":"provenance","command":"merge-manifest"},"kwargs":"sha256:664d7ec16496a75324b349c0b79313c4d402227a2dad9ccd0049dd6ba34faf6f","body":"sha256:b30b0c92d5bd257e52e3c8256ebc76c0673d14e4c2e996925a6699b264276231"}
	cmd.Args = cobra.ArbitraryArgs
	// /sedum:setArgs:any
	// sedum:anchor:flags
	// sedum:anchor:subcommands
	return cmd
}
