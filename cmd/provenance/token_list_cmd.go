package main

import "github.com/spf13/cobra"

func newTokenListCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"v":2,"tier":"owned","record":"prov-2026-b071f42a","gen":"sha256:07b2170a3890a057d5f8c16a5193181369591558ff57477d9219a65601fd0761","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","id":{"bin":"provenance","command":"token-list"},"kwargs":"sha256:e90d557a9baa2f54b5cae98079a869eaa0f0d9984e30c3bf8a22918413a3a998","body":"sha256:3b11b14fae8405214059bacfc0edc68ba33444cc257675e5cead4697a14d58ff"}
	cmd.Use = "list"
	cmd.Short = "List personal tokens (never their secrets)"
	// /sedum:describeCommand
	// sedum:setRun {"v":2,"tier":"owned","record":"prov-2026-b071f42a","gen":"sha256:07b2170a3890a057d5f8c16a5193181369591558ff57477d9219a65601fd0761","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","id":{"bin":"provenance","command":"token-list"},"kwargs":"sha256:c490e2e3d710ce45a67e0ef1a58b1ccc9e275f26263a0be21c102542b57192b0","body":"sha256:59c2e256c8f4801a9c64ba073cbbcc6738e710b564ac6fca913eac733dad8b92"}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runTokenList(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:none {"v":2,"tier":"owned","record":"prov-2026-b071f42a","gen":"sha256:07b2170a3890a057d5f8c16a5193181369591558ff57477d9219a65601fd0761","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","id":{"bin":"provenance","command":"token-list"},"kwargs":"sha256:c13243da746572436a693da9889e9fba82a135a256033fbf9bac99cea6f03fd6","body":"sha256:778b1e3eb6004fe613ecb8793efaf401a29f1afdc382707cea55c05601aefd79"}
	cmd.Args = cobra.NoArgs
	// /sedum:setArgs:none
	// sedum:anchor:flags
	// sedum:addFlag:bool {"v":2,"tier":"owned","record":"prov-2026-b071f42a","gen":"sha256:07b2170a3890a057d5f8c16a5193181369591558ff57477d9219a65601fd0761","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"token-list","flag":"json"},"kwargs":"sha256:2f342faea31f88f7d5bafdf0d927ab27aac7d1c680124b5af9f3d2267115e6f1","body":"sha256:d039c1f2caf44bc7dd84431a8617dd2453d265a0337b351f1488a67d6da3b9ce"}
	cmd.Flags().Bool("json", false, "Print as JSON")
	// /sedum:addFlag:bool
	// sedum:anchor:subcommands
	return cmd
}
