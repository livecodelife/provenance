package main

import "github.com/spf13/cobra"

func newVersionCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"v":2,"tier":"owned","record":"prov-2026-adabd65f","gen":"sha256:4ae46f3089d333bd12c54eda92a2d20ae90a3ad1ebf7c65d3eda625c0356ee02","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","id":{"bin":"provenance","command":"version"},"kwargs":"sha256:75ce99902705dcfea3160f36d149f5434753f5b6f595cf6410daf1a734768de3","body":"sha256:a961aadc506874fce71b8a6f21016fd928242037c55e1b2560b6daa86228977c"}
	cmd.Use = "version"
	cmd.Short = "Print the provenance version"
	// /sedum:describeCommand
	// sedum:setRun {"v":2,"tier":"owned","record":"prov-2026-adabd65f","gen":"sha256:4ae46f3089d333bd12c54eda92a2d20ae90a3ad1ebf7c65d3eda625c0356ee02","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","id":{"bin":"provenance","command":"version"},"kwargs":"sha256:005f3d08f0a921fb396d39e7d20c07e4b9db37f8fcea0735027b669bb1f413b1","body":"sha256:252cf2ffdaacac15031d14ecfecdd5b17fe70f3f27937dd5af5eb7e95e26a711"}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runVersion(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:none {"v":2,"tier":"owned","record":"prov-2026-adabd65f","gen":"sha256:4ae46f3089d333bd12c54eda92a2d20ae90a3ad1ebf7c65d3eda625c0356ee02","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","id":{"bin":"provenance","command":"version"},"kwargs":"sha256:d4ee72d0e671be0a31b81c04678e1ef4cfa891a336ac89b06c9c707642af07d1","body":"sha256:778b1e3eb6004fe613ecb8793efaf401a29f1afdc382707cea55c05601aefd79"}
	cmd.Args = cobra.NoArgs
	// /sedum:setArgs:none
	// sedum:anchor:flags
	// sedum:anchor:subcommands
	return cmd
}
