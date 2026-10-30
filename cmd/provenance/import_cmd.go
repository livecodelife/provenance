package main

import "github.com/spf13/cobra"

func newImportCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"v":2,"tier":"owned","record":"prov-2026-871605a9","gen":"sha256:c7b209fbc9bbd544f0e8d90fa47cd5bb0b654ddba7423c834cd61299b3619190","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","id":{"bin":"provenance","command":"import"},"kwargs":"sha256:397832591a17da7cd71a018fe2c814ca74bc014e77dfd100ac2fbb489d7ee9eb","body":"sha256:8f6d399f82d45a300820f5a7ed31600b55749a0d5fdcb65538ed9dc86eb22597"}
	cmd.Use = "import [id...]"
	cmd.Short = "Bring versions made on Provisio (for example in the UI) into this project"
	// /sedum:describeCommand
	// sedum:setRun {"v":2,"tier":"owned","record":"prov-2026-871605a9","gen":"sha256:c7b209fbc9bbd544f0e8d90fa47cd5bb0b654ddba7423c834cd61299b3619190","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","id":{"bin":"provenance","command":"import"},"kwargs":"sha256:fdac7a72d2523614b7ac53d3a6c788f888330139f0261b2683cedf30c4b7886c","body":"sha256:c33187c40413943a1c64485fe6118d26f58e922cd642522cd72ef4e571def9bc"}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runImport(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:any {"v":2,"tier":"owned","record":"prov-2026-871605a9","gen":"sha256:c7b209fbc9bbd544f0e8d90fa47cd5bb0b654ddba7423c834cd61299b3619190","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","id":{"bin":"provenance","command":"import"},"kwargs":"sha256:8f38f06d1839a3cf16031766658bc2d5418cf131c387a1ab7f1cf4128df3d674","body":"sha256:b30b0c92d5bd257e52e3c8256ebc76c0673d14e4c2e996925a6699b264276231"}
	cmd.Args = cobra.ArbitraryArgs
	// /sedum:setArgs:any
	// sedum:anchor:flags
	// sedum:addFlag:bool {"v":2,"tier":"owned","record":"prov-2026-871605a9","gen":"sha256:c7b209fbc9bbd544f0e8d90fa47cd5bb0b654ddba7423c834cd61299b3619190","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"import","flag":"all"},"kwargs":"sha256:0e9d1e09f2c37b5eb8a8e92e369ed1821af0cb74269204c82452d26c92b124c9","body":"sha256:86cd8478a1b7eeaa82d774b502938e0513b77c9ce3e98489e04fd6255290183e"}
	cmd.Flags().Bool("all", false, "Import into every record")
	// /sedum:addFlag:bool
	// sedum:anchor:subcommands
	return cmd
}
