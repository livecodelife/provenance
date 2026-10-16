package main

import "github.com/spf13/cobra"

func newHooksPrePushCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"tier":"owned","record":"prov-2026-10bc3e89","gen":"sha256:e93ace3473e7af7dd8c92e4f660130b255ccb6bbefff31577a90dcb3982e21d4","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","kwargs":{"bin":"provenance","command":"hooks-pre-push","short":"Run pre_push_command in a clean checkout of every pushed commit (the pre-push hook)","use":"pre-push"}}
	cmd.Use = "pre-push"
	cmd.Short = "Run pre_push_command in a clean checkout of every pushed commit (the pre-push hook)"
	// /sedum:describeCommand
	// sedum:setRun {"tier":"owned","record":"prov-2026-10bc3e89","gen":"sha256:e93ace3473e7af7dd8c92e4f660130b255ccb6bbefff31577a90dcb3982e21d4","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","kwargs":{"bin":"provenance","command":"hooks-pre-push"}}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runHooksPrePush(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:any {"tier":"owned","record":"prov-2026-10bc3e89","gen":"sha256:e93ace3473e7af7dd8c92e4f660130b255ccb6bbefff31577a90dcb3982e21d4","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","kwargs":{"arity":"any","bin":"provenance","command":"hooks-pre-push"}}
	cmd.Args = cobra.ArbitraryArgs
	// /sedum:setArgs:any
	// sedum:anchor:flags
	// sedum:anchor:subcommands
	return cmd
}
