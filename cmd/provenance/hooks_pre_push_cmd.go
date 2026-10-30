package main

import "github.com/spf13/cobra"

func newHooksPrePushCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"v":2,"tier":"owned","record":"prov-2026-10bc3e89","gen":"sha256:e93ace3473e7af7dd8c92e4f660130b255ccb6bbefff31577a90dcb3982e21d4","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","id":{"bin":"provenance","command":"hooks-pre-push"},"kwargs":"sha256:02440b6e14ef69645642fe76c96513f3668e009b1d871b94035dbf53a4516ca3","body":"sha256:ba4e697b6c70ff091c2f0bf8074163c0d3c0b9d093c9d1c71346b871a23ffc8d"}
	cmd.Use = "pre-push"
	cmd.Short = "Run pre_push_command in a clean checkout of every pushed commit (the pre-push hook)"
	// /sedum:describeCommand
	// sedum:setRun {"v":2,"tier":"owned","record":"prov-2026-10bc3e89","gen":"sha256:e93ace3473e7af7dd8c92e4f660130b255ccb6bbefff31577a90dcb3982e21d4","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","id":{"bin":"provenance","command":"hooks-pre-push"},"kwargs":"sha256:4c184471d774416296230349316a110a125fe792b041426bcb8eebf6a16fd807","body":"sha256:0df709386f575b036e69b4278dac16edc128b95f5bd88051e6774b2d5395001d"}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runHooksPrePush(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:any {"v":2,"tier":"owned","record":"prov-2026-10bc3e89","gen":"sha256:e93ace3473e7af7dd8c92e4f660130b255ccb6bbefff31577a90dcb3982e21d4","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","id":{"bin":"provenance","command":"hooks-pre-push"},"kwargs":"sha256:f19ab1e27dfe8492498645f193678844022a051a6fd855b20142ff3d4639bfbe","body":"sha256:b30b0c92d5bd257e52e3c8256ebc76c0673d14e4c2e996925a6699b264276231"}
	cmd.Args = cobra.ArbitraryArgs
	// /sedum:setArgs:any
	// sedum:anchor:flags
	// sedum:anchor:subcommands
	return cmd
}
