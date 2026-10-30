package main

import "github.com/spf13/cobra"

func newInstallPluginCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"v":2,"tier":"owned","record":"prov-2026-d71d0066","gen":"sha256:cae7a2d3a080d43f1ba12537ce93365db84c39195d99ec904268e9c15daab6a9","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","id":{"bin":"provenance","command":"install-plugin"},"kwargs":"sha256:f52ae55e9272523e5f0208e04a8cd47d3dffe988d041319a4b654259c3b3bb9f","body":"sha256:722f40b3f227e7ab29fdc9e15e97daae3c4f84e061d8c089aaaff31ad369d6d6"}
	cmd.Use = "install-plugin"
	cmd.Short = "Install the Claude Code provenance plugin into the repository"
	// /sedum:describeCommand
	// sedum:setRun {"v":2,"tier":"owned","record":"prov-2026-d71d0066","gen":"sha256:cae7a2d3a080d43f1ba12537ce93365db84c39195d99ec904268e9c15daab6a9","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","id":{"bin":"provenance","command":"install-plugin"},"kwargs":"sha256:a37c12b3ba1f3e61e7be3096ac0f798f2ce02af41bc83034dd38e0d0bb396be8","body":"sha256:af7f76a32d3a4cbe55387ac2c06caf5c960cd82a1691b40b072693265ffd475a"}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runInstallPlugin(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:none {"v":2,"tier":"owned","record":"prov-2026-d71d0066","gen":"sha256:cae7a2d3a080d43f1ba12537ce93365db84c39195d99ec904268e9c15daab6a9","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","id":{"bin":"provenance","command":"install-plugin"},"kwargs":"sha256:9c0a9d9d9b06adf4a9aa3766469f6f130f15cadf1f4a6cc00d7a733e657fffcd","body":"sha256:778b1e3eb6004fe613ecb8793efaf401a29f1afdc382707cea55c05601aefd79"}
	cmd.Args = cobra.NoArgs
	// /sedum:setArgs:none
	// sedum:anchor:flags
	// sedum:addFlag:string {"v":2,"tier":"owned","record":"prov-2026-d71d0066","gen":"sha256:cae7a2d3a080d43f1ba12537ce93365db84c39195d99ec904268e9c15daab6a9","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"install-plugin","flag":"path"},"kwargs":"sha256:6e1735b27f3ff91861c1480caef22842d08846849767057a65396615c05cadf9","body":"sha256:c91f1158db83069175fd91ec2654ca1b1812b132a333bdd964ca92f06b2adb03"}
	cmd.Flags().String("path", "", "Target plugins directory relative to the repository root (default: .claude/plugins)")
	// /sedum:addFlag:string
	// sedum:addFlag:bool {"v":2,"tier":"owned","record":"prov-2026-d71d0066","gen":"sha256:cae7a2d3a080d43f1ba12537ce93365db84c39195d99ec904268e9c15daab6a9","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"install-plugin","flag":"force"},"kwargs":"sha256:a9e51fee0821b77a4bc95641118bb371393e7a9faca0bd6ceeab545974c7547c","body":"sha256:d9034ed699def2ce79a49ab206ac85c2439e521ec8468e054a5d469c9b8f05e1"}
	cmd.Flags().Bool("force", false, "Overwrite existing files that differ")
	// /sedum:addFlag:bool
	// sedum:anchor:subcommands
	return cmd
}
