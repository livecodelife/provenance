package main

import "github.com/spf13/cobra"

func newLogoutCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"tier":"owned","record":"prov-2026-32201524","gen":"sha256:ee05402713f727aa111554dd0aacd7f0e84ce70ae0e346c620bc80a4c4bad7c7","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","kwargs":{"bin":"provenance","command":"logout","short":"Sign out of Provisio and delete the local credential","use":"logout"}}
	cmd.Use = "logout"
	cmd.Short = "Sign out of Provisio and delete the local credential"
	// /sedum:describeCommand
	// sedum:setRun {"tier":"owned","record":"prov-2026-32201524","gen":"sha256:ee05402713f727aa111554dd0aacd7f0e84ce70ae0e346c620bc80a4c4bad7c7","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","kwargs":{"bin":"provenance","command":"logout"}}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runLogout(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:none {"tier":"owned","record":"prov-2026-32201524","gen":"sha256:ee05402713f727aa111554dd0aacd7f0e84ce70ae0e346c620bc80a4c4bad7c7","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","kwargs":{"arity":"none","bin":"provenance","command":"logout"}}
	cmd.Args = cobra.NoArgs
	// /sedum:setArgs:none
	// sedum:anchor:flags
	// sedum:anchor:subcommands
	return cmd
}
