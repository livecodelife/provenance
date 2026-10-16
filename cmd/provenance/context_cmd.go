package main

import "github.com/spf13/cobra"

func newContextCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:0a18ffebf7eba674f7ee6e1817baa18d76bbad63cd34d4668bea4d9b2c20df40","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","kwargs":{"bin":"provenance","command":"context","short":"Show the provenance context for paths or records","use":"context [targets...]"}}
	cmd.Use = "context [targets...]"
	cmd.Short = "Show the provenance context for paths or records"
	// /sedum:describeCommand
	// sedum:setRun {"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:0a18ffebf7eba674f7ee6e1817baa18d76bbad63cd34d4668bea4d9b2c20df40","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","kwargs":{"bin":"provenance","command":"context"}}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runContext(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:any {"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:0a18ffebf7eba674f7ee6e1817baa18d76bbad63cd34d4668bea4d9b2c20df40","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","kwargs":{"arity":"any","bin":"provenance","command":"context"}}
	cmd.Args = cobra.ArbitraryArgs
	// /sedum:setArgs:any
	// sedum:anchor:flags
	// sedum:addFlag:strings {"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:0a18ffebf7eba674f7ee6e1817baa18d76bbad63cd34d4668bea4d9b2c20df40","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"context","flag":"files","type":"strings","usage":"Explicit target list (alternative to arguments)"}}
	cmd.Flags().StringArray("files", nil, "Explicit target list (alternative to arguments)")
	// /sedum:addFlag:strings
	// sedum:addFlag:string {"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:0a18ffebf7eba674f7ee6e1817baa18d76bbad63cd34d4668bea4d9b2c20df40","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"context","flag":"format","type":"string","usage":"Output format (human|compact|json)"}}
	cmd.Flags().String("format", "", "Output format (human|compact|json)")
	// /sedum:addFlag:string
	// sedum:anchor:subcommands
	return cmd
}
