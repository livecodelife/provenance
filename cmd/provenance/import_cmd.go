package main

import "github.com/spf13/cobra"

func newImportCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"tier":"owned","record":"prov-2026-871605a9","gen":"sha256:a837b61e0b1282ffa2184bd15712bd692804fd3785f63fd676f9036654ebea2d","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","kwargs":{"bin":"provenance","command":"import","short":"Bring versions made on Provisio (for example in the UI) into this project","use":"import [id...]"}}
	cmd.Use = "import [id...]"
	cmd.Short = "Bring versions made on Provisio (for example in the UI) into this project"
	// /sedum:describeCommand
	// sedum:setRun {"tier":"owned","record":"prov-2026-871605a9","gen":"sha256:a837b61e0b1282ffa2184bd15712bd692804fd3785f63fd676f9036654ebea2d","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","kwargs":{"bin":"provenance","command":"import"}}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runImport(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:any {"tier":"owned","record":"prov-2026-871605a9","gen":"sha256:a837b61e0b1282ffa2184bd15712bd692804fd3785f63fd676f9036654ebea2d","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","kwargs":{"arity":"any","bin":"provenance","command":"import"}}
	cmd.Args = cobra.ArbitraryArgs
	// /sedum:setArgs:any
	// sedum:anchor:flags
	// sedum:addFlag:bool {"tier":"owned","record":"prov-2026-871605a9","gen":"sha256:a837b61e0b1282ffa2184bd15712bd692804fd3785f63fd676f9036654ebea2d","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"import","flag":"all","type":"bool","usage":"Import into every record"}}
	cmd.Flags().Bool("all", false, "Import into every record")
	// /sedum:addFlag:bool
	// sedum:anchor:subcommands
	return cmd
}
