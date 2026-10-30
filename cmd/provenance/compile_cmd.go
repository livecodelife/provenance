package main

import "github.com/spf13/cobra"

func newCompileCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"v":2,"tier":"owned","record":"prov-2026-d84dc10e","gen":"sha256:2e97312cea5e86a4f4e11ea65cacfe9c4a63da245a1c446b00a2c6740234e92a","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","id":{"bin":"provenance","command":"compile"},"kwargs":"sha256:15cfcc9a7952a69f01c9cbc7dde17c45c41c1d4d3dc5d1ed73500752b919af88","body":"sha256:751786b014e11f89b28c66b61ee1f545917c842d1e0ca80d1d3ed99f72c65c45"}
	cmd.Use = "compile"
	cmd.Short = "Rebuild the version manifest from the record snapshots"
	// /sedum:describeCommand
	// sedum:setRun {"v":2,"tier":"owned","record":"prov-2026-d84dc10e","gen":"sha256:2e97312cea5e86a4f4e11ea65cacfe9c4a63da245a1c446b00a2c6740234e92a","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","id":{"bin":"provenance","command":"compile"},"kwargs":"sha256:538d96b09cc32982dec744e3804baca080430296aa3922d1f2f41ba2e1db3fc7","body":"sha256:81f9d006aa318d297110d43713e5fe67555d03a60fdda6f80a836e50bbf9dd71"}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runCompile(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:none {"v":2,"tier":"owned","record":"prov-2026-d84dc10e","gen":"sha256:2e97312cea5e86a4f4e11ea65cacfe9c4a63da245a1c446b00a2c6740234e92a","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","id":{"bin":"provenance","command":"compile"},"kwargs":"sha256:2f04431b7c01c8f748181987b13c80a55b0e543761bc49efb2cf375034b13ac0","body":"sha256:778b1e3eb6004fe613ecb8793efaf401a29f1afdc382707cea55c05601aefd79"}
	cmd.Args = cobra.NoArgs
	// /sedum:setArgs:none
	// sedum:anchor:flags
	// sedum:addFlag:bool {"v":2,"tier":"owned","record":"prov-2026-d84dc10e","gen":"sha256:2e97312cea5e86a4f4e11ea65cacfe9c4a63da245a1c446b00a2c6740234e92a","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"compile","flag":"check"},"kwargs":"sha256:7cbe7d9a60b12b43bef03a0d56998cf6dcf159082d6beeed3291742308ff9c8a","body":"sha256:824797041a19303ceb0e65de232dcdf56f8ae0b5e1ecbf78425550e96e78a554"}
	cmd.Flags().Bool("check", false, "Fail, changing nothing, when the manifest is out of date")
	// /sedum:addFlag:bool
	// sedum:anchor:subcommands
	return cmd
}
