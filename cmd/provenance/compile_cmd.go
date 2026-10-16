package main

import "github.com/spf13/cobra"

func newCompileCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"tier":"owned","record":"prov-2026-d84dc10e","gen":"sha256:f20a2537833e4924b62273583dc6d4c9f1678ed42ca83ff10d54da85435d4d91","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","kwargs":{"bin":"provenance","command":"compile","short":"Rebuild the version manifest from the record snapshots","use":"compile"}}
	cmd.Use = "compile"
	cmd.Short = "Rebuild the version manifest from the record snapshots"
	// /sedum:describeCommand
	// sedum:setRun {"tier":"owned","record":"prov-2026-d84dc10e","gen":"sha256:f20a2537833e4924b62273583dc6d4c9f1678ed42ca83ff10d54da85435d4d91","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","kwargs":{"bin":"provenance","command":"compile"}}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runCompile(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:none {"tier":"owned","record":"prov-2026-d84dc10e","gen":"sha256:f20a2537833e4924b62273583dc6d4c9f1678ed42ca83ff10d54da85435d4d91","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","kwargs":{"arity":"none","bin":"provenance","command":"compile"}}
	cmd.Args = cobra.NoArgs
	// /sedum:setArgs:none
	// sedum:anchor:flags
	// sedum:anchor:subcommands
	return cmd
}
