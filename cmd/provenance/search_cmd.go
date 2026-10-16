package main

import "github.com/spf13/cobra"

func newSearchCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"tier":"owned","record":"prov-2026-4eb70ef7","gen":"sha256:8dd0cf07c10ca2de12963c41bee8f3709b76949e3c89a1471e2831d9ab3e96fe","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","kwargs":{"bin":"provenance","command":"search","short":"Search records by words (offline, lexical)","use":"search [query]"}}
	cmd.Use = "search [query]"
	cmd.Short = "Search records by words (offline, lexical)"
	// /sedum:describeCommand
	// sedum:setRun {"tier":"owned","record":"prov-2026-4eb70ef7","gen":"sha256:8dd0cf07c10ca2de12963c41bee8f3709b76949e3c89a1471e2831d9ab3e96fe","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","kwargs":{"bin":"provenance","command":"search"}}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runSearch(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:max_one {"tier":"owned","record":"prov-2026-4eb70ef7","gen":"sha256:8dd0cf07c10ca2de12963c41bee8f3709b76949e3c89a1471e2831d9ab3e96fe","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","kwargs":{"arity":"max_one","bin":"provenance","command":"search"}}
	cmd.Args = cobra.MaximumNArgs(1)
	// /sedum:setArgs:max_one
	// sedum:anchor:flags
	// sedum:addFlag:string {"tier":"owned","record":"prov-2026-4eb70ef7","gen":"sha256:8dd0cf07c10ca2de12963c41bee8f3709b76949e3c89a1471e2831d9ab3e96fe","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"search","flag":"query","type":"string","usage":"Search words (alternative to the argument)"}}
	cmd.Flags().String("query", "", "Search words (alternative to the argument)")
	// /sedum:addFlag:string
	// sedum:addFlag:int {"tier":"owned","record":"prov-2026-4eb70ef7","gen":"sha256:8dd0cf07c10ca2de12963c41bee8f3709b76949e3c89a1471e2831d9ab3e96fe","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"search","flag":"limit","type":"int","usage":"Maximum results (default 5, at most 50)"}}
	cmd.Flags().Int("limit", 0, "Maximum results (default 5, at most 50)")
	// /sedum:addFlag:int
	// sedum:addFlag:string {"tier":"owned","record":"prov-2026-4eb70ef7","gen":"sha256:8dd0cf07c10ca2de12963c41bee8f3709b76949e3c89a1471e2831d9ab3e96fe","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"search","flag":"status","type":"string","usage":"Only records with this status (draft|open|implemented|deprecated)"}}
	cmd.Flags().String("status", "", "Only records with this status (draft|open|implemented|deprecated)")
	// /sedum:addFlag:string
	// sedum:addFlag:string {"tier":"owned","record":"prov-2026-4eb70ef7","gen":"sha256:8dd0cf07c10ca2de12963c41bee8f3709b76949e3c89a1471e2831d9ab3e96fe","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"search","flag":"type","type":"string","usage":"Only records of this tier (brief|blueprint|imprint)"}}
	cmd.Flags().String("type", "", "Only records of this tier (brief|blueprint|imprint)")
	// /sedum:addFlag:string
	// sedum:addFlag:string {"tier":"owned","record":"prov-2026-4eb70ef7","gen":"sha256:8dd0cf07c10ca2de12963c41bee8f3709b76949e3c89a1471e2831d9ab3e96fe","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"search","flag":"format","type":"string","usage":"Output format (human|json)"}}
	cmd.Flags().String("format", "", "Output format (human|json)")
	// /sedum:addFlag:string
	// sedum:anchor:subcommands
	return cmd
}
