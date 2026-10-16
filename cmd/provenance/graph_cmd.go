package main

import "github.com/spf13/cobra"

func newGraphCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"tier":"owned","record":"prov-2026-d17b9e44","gen":"sha256:315c66a50f0a199a267746ec3105d78a1b8a619606b138f294efc1ae2c242dba","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","kwargs":{"bin":"provenance","command":"graph","short":"Render the record graph: records, parents, pins and versions","use":"graph"}}
	cmd.Use = "graph"
	cmd.Short = "Render the record graph: records, parents, pins and versions"
	// /sedum:describeCommand
	// sedum:setRun {"tier":"owned","record":"prov-2026-d17b9e44","gen":"sha256:315c66a50f0a199a267746ec3105d78a1b8a619606b138f294efc1ae2c242dba","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","kwargs":{"bin":"provenance","command":"graph"}}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runGraph(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:none {"tier":"owned","record":"prov-2026-d17b9e44","gen":"sha256:315c66a50f0a199a267746ec3105d78a1b8a619606b138f294efc1ae2c242dba","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","kwargs":{"arity":"none","bin":"provenance","command":"graph"}}
	cmd.Args = cobra.NoArgs
	// /sedum:setArgs:none
	// sedum:anchor:flags
	// sedum:addFlag:string {"tier":"owned","record":"prov-2026-d17b9e44","gen":"sha256:315c66a50f0a199a267746ec3105d78a1b8a619606b138f294efc1ae2c242dba","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"graph","flag":"root","type":"string","usage":"Start from one record: its parent (one level up) and its descendants"}}
	cmd.Flags().String("root", "", "Start from one record: its parent (one level up) and its descendants")
	// /sedum:addFlag:string
	// sedum:addFlag:string {"tier":"owned","record":"prov-2026-d17b9e44","gen":"sha256:315c66a50f0a199a267746ec3105d78a1b8a619606b138f294efc1ae2c242dba","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"graph","flag":"filter","type":"string","usage":"Show only records with this status (draft|open|implemented|deprecated)"}}
	cmd.Flags().String("filter", "", "Show only records with this status (draft|open|implemented|deprecated)")
	// /sedum:addFlag:string
	// sedum:addFlag:string {"tier":"owned","record":"prov-2026-d17b9e44","gen":"sha256:315c66a50f0a199a267746ec3105d78a1b8a619606b138f294efc1ae2c242dba","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"graph","flag":"format","type":"string","usage":"Output format (human|json|dot)"}}
	cmd.Flags().String("format", "", "Output format (human|json|dot)")
	// /sedum:addFlag:string
	// sedum:anchor:subcommands
	return cmd
}
