package main

import "github.com/spf13/cobra"

func newRepinCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"v":2,"tier":"owned","record":"prov-2026-9bfa6107","gen":"sha256:e937012cd7f04b4dea168e2f24be21b5aadd0783afbe0dda58babc69031e3dbc","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","id":{"bin":"provenance","command":"repin"},"kwargs":"sha256:5407fce34520c36c34fd17fa7bcd4c37240553252b6b7214e2d778b1828e1eb6","body":"sha256:5d2cd7233aff502d36fbec3d502d0c899a56962081100abfb5352f73b951aa90"}
	cmd.Use = "repin [id]"
	cmd.Short = "Confirm a record against a newer parent version without a new version"
	// /sedum:describeCommand
	// sedum:setRun {"v":2,"tier":"owned","record":"prov-2026-9bfa6107","gen":"sha256:e937012cd7f04b4dea168e2f24be21b5aadd0783afbe0dda58babc69031e3dbc","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","id":{"bin":"provenance","command":"repin"},"kwargs":"sha256:a86827076b82167baa7d8f2b0862bd26bf0a132a3fd4f0c0fcb0c0541d2230be","body":"sha256:0f6e1849f62d577481ae1d32adfd8d3afcdee938ce8e4bd532a18e69af7d0647"}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runRepin(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:max_one {"v":2,"tier":"owned","record":"prov-2026-9bfa6107","gen":"sha256:e937012cd7f04b4dea168e2f24be21b5aadd0783afbe0dda58babc69031e3dbc","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","id":{"bin":"provenance","command":"repin"},"kwargs":"sha256:e6694e914aa61f7ab41b59763085e2363090c5c0941ceddef5d206824290f0e5","body":"sha256:186046a2435bbf4d0eae67e61e8a3909c91fd72c18b73dd2ff749798e3517ef3"}
	cmd.Args = cobra.MaximumNArgs(1)
	// /sedum:setArgs:max_one
	// sedum:anchor:flags
	// sedum:addFlag:string {"v":2,"tier":"owned","record":"prov-2026-9bfa6107","gen":"sha256:e937012cd7f04b4dea168e2f24be21b5aadd0783afbe0dda58babc69031e3dbc","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"repin","flag":"record"},"kwargs":"sha256:599bc35b5f4ebe545bdbb2fd15b7022a683845b5f716eb15a1c1b84f57f61fd7","body":"sha256:358cdb90224e37ea23dff05a3546011fc9c7a08f348b4b5ccb2207816105e281"}
	cmd.Flags().String("record", "", "Record to repin (or give it as the argument)")
	// /sedum:addFlag:string
	// sedum:addFlag:string {"v":2,"tier":"owned","record":"prov-2026-9bfa6107","gen":"sha256:e937012cd7f04b4dea168e2f24be21b5aadd0783afbe0dda58babc69031e3dbc","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"repin","flag":"reason"},"kwargs":"sha256:77c94e3d8dfaaa867094564860028766dceb1ed5e024d9f8355a7e3e38fef11b","body":"sha256:933b2964ceea12a58a5d5202986077c9dc36cf6fc4ae6a92498704d1c48c2d22"}
	cmd.Flags().String("reason", "", "Why the record is still correct against the parent (required)")
	// /sedum:addFlag:string
	// sedum:addFlag:int {"v":2,"tier":"owned","record":"prov-2026-9bfa6107","gen":"sha256:e937012cd7f04b4dea168e2f24be21b5aadd0783afbe0dda58babc69031e3dbc","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"repin","flag":"to"},"kwargs":"sha256:d661ae4e9882a3a963e765d3d4360fb430bdc58b5388525fda1bd4228c2b1a97","body":"sha256:443d1a5da5152f043519fafb6ea7ee1e9dc6951a3742330eb8619229cd038552"}
	cmd.Flags().Int("to", 0, "Parent version to pin to (default: the parent's current version)")
	// /sedum:addFlag:int
	// sedum:anchor:subcommands
	return cmd
}
