package main

import "github.com/spf13/cobra"

func newCompileCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"tier":"owned","record":"prov-2026-d84dc10e","gen":"sha256:2e97312cea5e86a4f4e11ea65cacfe9c4a63da245a1c446b00a2c6740234e92a","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","kwargs":{"bin":"provenance","command":"compile","short":"Rebuild the version manifest from the record snapshots","use":"compile"}}
	cmd.Use = "compile"
	cmd.Short = "Rebuild the version manifest from the record snapshots"
	// /sedum:describeCommand
	// sedum:setRun {"tier":"owned","record":"prov-2026-d84dc10e","gen":"sha256:2e97312cea5e86a4f4e11ea65cacfe9c4a63da245a1c446b00a2c6740234e92a","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","kwargs":{"bin":"provenance","command":"compile"}}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runCompile(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:none {"tier":"owned","record":"prov-2026-d84dc10e","gen":"sha256:2e97312cea5e86a4f4e11ea65cacfe9c4a63da245a1c446b00a2c6740234e92a","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","kwargs":{"arity":"none","bin":"provenance","command":"compile"}}
	cmd.Args = cobra.NoArgs
	// /sedum:setArgs:none
	// sedum:anchor:flags
	// sedum:addFlag:bool {"tier":"owned","record":"prov-2026-d84dc10e","gen":"sha256:2e97312cea5e86a4f4e11ea65cacfe9c4a63da245a1c446b00a2c6740234e92a","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"compile","flag":"check","type":"bool","usage":"Fail, changing nothing, when the manifest is out of date"}}
	cmd.Flags().Bool("check", false, "Fail, changing nothing, when the manifest is out of date")
	// /sedum:addFlag:bool
	// sedum:anchor:subcommands
	return cmd
}
