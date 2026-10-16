package main

import "github.com/spf13/cobra"

func newInitCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"tier":"owned","record":"prov-2026-1a5d2549","gen":"sha256:e344efddd9441a7734ecfef807e1e2c0d0a91b33f8cffcc5e6d16b03b1229edb","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","kwargs":{"bin":"provenance","command":"init","short":"Scaffold provenance in a new project","use":"init"}}
	cmd.Use = "init"
	cmd.Short = "Scaffold provenance in a new project"
	// /sedum:describeCommand
	// sedum:setRun {"tier":"owned","record":"prov-2026-1a5d2549","gen":"sha256:e344efddd9441a7734ecfef807e1e2c0d0a91b33f8cffcc5e6d16b03b1229edb","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","kwargs":{"bin":"provenance","command":"init"}}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runInit(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:none {"tier":"owned","record":"prov-2026-1a5d2549","gen":"sha256:e344efddd9441a7734ecfef807e1e2c0d0a91b33f8cffcc5e6d16b03b1229edb","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","kwargs":{"arity":"none","bin":"provenance","command":"init"}}
	cmd.Args = cobra.NoArgs
	// /sedum:setArgs:none
	// sedum:anchor:flags
	// sedum:addFlag:string {"tier":"owned","record":"prov-2026-1a5d2549","gen":"sha256:e344efddd9441a7734ecfef807e1e2c0d0a91b33f8cffcc5e6d16b03b1229edb","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"init","flag":"dir","type":"string","usage":"Project directory (default: the current directory)"}}
	cmd.Flags().String("dir", "", "Project directory (default: the current directory)")
	// /sedum:addFlag:string
	// sedum:addFlag:bool {"tier":"owned","record":"prov-2026-1a5d2549","gen":"sha256:e344efddd9441a7734ecfef807e1e2c0d0a91b33f8cffcc5e6d16b03b1229edb","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"init","flag":"force","type":"bool","usage":"Overwrite an existing .provenance.yml (the old one is saved as .provenance.yml.bak)"}}
	cmd.Flags().Bool("force", false, "Overwrite an existing .provenance.yml (the old one is saved as .provenance.yml.bak)")
	// /sedum:addFlag:bool
	// sedum:anchor:subcommands
	return cmd
}
