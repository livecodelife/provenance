package main

import "github.com/spf13/cobra"

func newRepinCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"tier":"owned","record":"prov-2026-9bfa6107","gen":"sha256:e937012cd7f04b4dea168e2f24be21b5aadd0783afbe0dda58babc69031e3dbc","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","kwargs":{"bin":"provenance","command":"repin","short":"Confirm a record against a newer parent version without a new version","use":"repin [id]"}}
	cmd.Use = "repin [id]"
	cmd.Short = "Confirm a record against a newer parent version without a new version"
	// /sedum:describeCommand
	// sedum:setRun {"tier":"owned","record":"prov-2026-9bfa6107","gen":"sha256:e937012cd7f04b4dea168e2f24be21b5aadd0783afbe0dda58babc69031e3dbc","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","kwargs":{"bin":"provenance","command":"repin"}}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runRepin(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:max_one {"tier":"owned","record":"prov-2026-9bfa6107","gen":"sha256:e937012cd7f04b4dea168e2f24be21b5aadd0783afbe0dda58babc69031e3dbc","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","kwargs":{"arity":"max_one","bin":"provenance","command":"repin"}}
	cmd.Args = cobra.MaximumNArgs(1)
	// /sedum:setArgs:max_one
	// sedum:anchor:flags
	// sedum:addFlag:string {"tier":"owned","record":"prov-2026-9bfa6107","gen":"sha256:e937012cd7f04b4dea168e2f24be21b5aadd0783afbe0dda58babc69031e3dbc","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"repin","flag":"record","type":"string","usage":"Record to repin (or give it as the argument)"}}
	cmd.Flags().String("record", "", "Record to repin (or give it as the argument)")
	// /sedum:addFlag:string
	// sedum:addFlag:string {"tier":"owned","record":"prov-2026-9bfa6107","gen":"sha256:e937012cd7f04b4dea168e2f24be21b5aadd0783afbe0dda58babc69031e3dbc","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"repin","flag":"reason","type":"string","usage":"Why the record is still correct against the parent (required)"}}
	cmd.Flags().String("reason", "", "Why the record is still correct against the parent (required)")
	// /sedum:addFlag:string
	// sedum:addFlag:int {"tier":"owned","record":"prov-2026-9bfa6107","gen":"sha256:e937012cd7f04b4dea168e2f24be21b5aadd0783afbe0dda58babc69031e3dbc","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"repin","flag":"to","type":"int","usage":"Parent version to pin to (default: the parent's current version)"}}
	cmd.Flags().Int("to", 0, "Parent version to pin to (default: the parent's current version)")
	// /sedum:addFlag:int
	// sedum:anchor:subcommands
	return cmd
}
