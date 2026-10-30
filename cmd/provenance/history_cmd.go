package main

import "github.com/spf13/cobra"

func newHistoryCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"v":2,"tier":"owned","record":"prov-2026-d435c27b","gen":"sha256:4f36ba1a5a6a79f262293754b5a46ef8739347c5fb4b6df037af455f19b551ba","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","id":{"bin":"provenance","command":"history"},"kwargs":"sha256:f2526527f75355465c2692968f4ea933d7f01c316db93d2aaa5cfef4e7da028f","body":"sha256:e167611d2532e7d47236d5c8822fa8fae8d4c82eb6b8d9227bdf710e8c0416bf"}
	cmd.Use = "history [id]"
	cmd.Short = "List a record's versions with reasons and commits"
	// /sedum:describeCommand
	// sedum:setRun {"v":2,"tier":"owned","record":"prov-2026-d435c27b","gen":"sha256:4f36ba1a5a6a79f262293754b5a46ef8739347c5fb4b6df037af455f19b551ba","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","id":{"bin":"provenance","command":"history"},"kwargs":"sha256:f16fb5be7b018724f49ff95d556f8723a1c38582dcddce2243e123e7426fb28b","body":"sha256:32240ddf8391bea8af0effcd00aa9c46ce50a35919ae901cdebcb9a4c86ec45d"}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runHistory(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:max_one {"v":2,"tier":"owned","record":"prov-2026-d435c27b","gen":"sha256:4f36ba1a5a6a79f262293754b5a46ef8739347c5fb4b6df037af455f19b551ba","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","id":{"bin":"provenance","command":"history"},"kwargs":"sha256:0d4ae52b0bec1c8d523e7a08c6d080e0e3adf872b8c0f1fd46e7d9f91d44dd70","body":"sha256:186046a2435bbf4d0eae67e61e8a3909c91fd72c18b73dd2ff749798e3517ef3"}
	cmd.Args = cobra.MaximumNArgs(1)
	// /sedum:setArgs:max_one
	// sedum:anchor:flags
	// sedum:addFlag:string {"v":2,"tier":"owned","record":"prov-2026-d435c27b","gen":"sha256:4f36ba1a5a6a79f262293754b5a46ef8739347c5fb4b6df037af455f19b551ba","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"history","flag":"record"},"kwargs":"sha256:3e6e5fc4e4ef7e3b0726a79e12d9d32e40bb440730b09c7910b2b296b833caf4","body":"sha256:99c4a5d5026aad1799f22932d686ec7292865495f60ab0c63ead1baecd134255"}
	cmd.Flags().String("record", "", "Record to show (or give it as the argument)")
	// /sedum:addFlag:string
	// sedum:addFlag:bool {"v":2,"tier":"owned","record":"prov-2026-d435c27b","gen":"sha256:4f36ba1a5a6a79f262293754b5a46ef8739347c5fb4b6df037af455f19b551ba","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"history","flag":"json"},"kwargs":"sha256:686bc37f58b242fddd152184784422bf0e4bf217c3c235b73329db828d0261c6","body":"sha256:d039c1f2caf44bc7dd84431a8617dd2453d265a0337b351f1488a67d6da3b9ce"}
	cmd.Flags().Bool("json", false, "Print as JSON")
	// /sedum:addFlag:bool
	// sedum:anchor:subcommands
	return cmd
}
