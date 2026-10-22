package main

import "github.com/spf13/cobra"

func newTokenRevokeCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"tier":"owned","record":"prov-2026-a8ec3e30","gen":"sha256:32397270f026f7499dfffc6c69c80d69185c2e1390ca78f0ee9ecba7837ebbdd","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","kwargs":{"bin":"provenance","command":"token-revoke","short":"Revoke a personal token","use":"revoke <id>"}}
	cmd.Use = "revoke <id>"
	cmd.Short = "Revoke a personal token"
	// /sedum:describeCommand
	// sedum:setRun {"tier":"owned","record":"prov-2026-a8ec3e30","gen":"sha256:32397270f026f7499dfffc6c69c80d69185c2e1390ca78f0ee9ecba7837ebbdd","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","kwargs":{"bin":"provenance","command":"token-revoke"}}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runTokenRevoke(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:one {"tier":"owned","record":"prov-2026-a8ec3e30","gen":"sha256:32397270f026f7499dfffc6c69c80d69185c2e1390ca78f0ee9ecba7837ebbdd","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","kwargs":{"arity":"one","bin":"provenance","command":"token-revoke"}}
	cmd.Args = cobra.ExactArgs(1)
	// /sedum:setArgs:one
	// sedum:anchor:flags
	// sedum:anchor:subcommands
	return cmd
}
