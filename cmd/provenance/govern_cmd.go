package main

import "github.com/spf13/cobra"

func newGovernCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"tier":"owned","record":"prov-2026-9ec4a843","gen":"sha256:cc0301a10fcd903006e4f2bed32a32b6a082f192eb2c42bb2c80462870e50daf","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","kwargs":{"bin":"provenance","command":"govern","short":"List the active records that govern the given files","use":"govern [files...]"}}
	cmd.Use = "govern [files...]"
	cmd.Short = "List the active records that govern the given files"
	// /sedum:describeCommand
	// sedum:setRun {"tier":"owned","record":"prov-2026-9ec4a843","gen":"sha256:cc0301a10fcd903006e4f2bed32a32b6a082f192eb2c42bb2c80462870e50daf","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","kwargs":{"bin":"provenance","command":"govern"}}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runGovern(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:any {"tier":"owned","record":"prov-2026-9ec4a843","gen":"sha256:cc0301a10fcd903006e4f2bed32a32b6a082f192eb2c42bb2c80462870e50daf","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","kwargs":{"arity":"any","bin":"provenance","command":"govern"}}
	cmd.Args = cobra.ArbitraryArgs
	// /sedum:setArgs:any
	// sedum:anchor:flags
	// sedum:addFlag:strings {"tier":"owned","record":"prov-2026-9ec4a843","gen":"sha256:cc0301a10fcd903006e4f2bed32a32b6a082f192eb2c42bb2c80462870e50daf","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"govern","flag":"files","type":"strings","usage":"Explicit file list (alternative to arguments)"}}
	cmd.Flags().StringArray("files", nil, "Explicit file list (alternative to arguments)")
	// /sedum:addFlag:strings
	// sedum:addFlag:string {"tier":"owned","record":"prov-2026-9ec4a843","gen":"sha256:cc0301a10fcd903006e4f2bed32a32b6a082f192eb2c42bb2c80462870e50daf","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"govern","flag":"format","type":"string","usage":"Output format (human|json)"}}
	cmd.Flags().String("format", "", "Output format (human|json)")
	// /sedum:addFlag:string
	// sedum:addFlag:bool {"tier":"owned","record":"prov-2026-9ec4a843","gen":"sha256:cc0301a10fcd903006e4f2bed32a32b6a082f192eb2c42bb2c80462870e50daf","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"govern","flag":"json","type":"bool","usage":"Alias for --format json"}}
	cmd.Flags().Bool("json", false, "Alias for --format json")
	// /sedum:addFlag:bool
	// sedum:anchor:subcommands
	return cmd
}
