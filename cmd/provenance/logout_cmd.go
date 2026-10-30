package main

import "github.com/spf13/cobra"

func newLogoutCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"v":2,"tier":"owned","record":"prov-2026-32201524","gen":"sha256:ee05402713f727aa111554dd0aacd7f0e84ce70ae0e346c620bc80a4c4bad7c7","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","id":{"bin":"provenance","command":"logout"},"kwargs":"sha256:f57817f3671cd7ac0170ef90be6b29dd95ab13a1f688ca9f6ed270e8d50fadca","body":"sha256:ff8b3f452c138faf1f3a702167922f24f6afd1e146ffd24fd9aebec5aa4f9b68"}
	cmd.Use = "logout"
	cmd.Short = "Sign out of Provisio and delete the local credential"
	// /sedum:describeCommand
	// sedum:setRun {"v":2,"tier":"owned","record":"prov-2026-32201524","gen":"sha256:ee05402713f727aa111554dd0aacd7f0e84ce70ae0e346c620bc80a4c4bad7c7","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","id":{"bin":"provenance","command":"logout"},"kwargs":"sha256:a826e9b22284c312384c899f3543700e7932e58cb10ca43b9fc05531b31a0e5b","body":"sha256:f6a836d341b79c1ed3fa76899e3c9bb97f4561fef5abfebdf8f740406d933064"}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runLogout(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:none {"v":2,"tier":"owned","record":"prov-2026-32201524","gen":"sha256:ee05402713f727aa111554dd0aacd7f0e84ce70ae0e346c620bc80a4c4bad7c7","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","id":{"bin":"provenance","command":"logout"},"kwargs":"sha256:3c5edbf654c1297d00c8b648f0e3c6202da5067eaa8cee505379ae10f7abe1df","body":"sha256:778b1e3eb6004fe613ecb8793efaf401a29f1afdc382707cea55c05601aefd79"}
	cmd.Args = cobra.NoArgs
	// /sedum:setArgs:none
	// sedum:anchor:flags
	// sedum:anchor:subcommands
	return cmd
}
