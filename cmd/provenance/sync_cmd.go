package main

import "github.com/spf13/cobra"

func newSyncCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"tier":"owned","record":"prov-2026-f807707a","gen":"sha256:3f6d1802d83ba9dcc72ed02c8927ae3f79e56242263a8c87803b425321b88f73","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","kwargs":{"bin":"provenance","command":"sync","short":"Fetch the records of the configured shared repos into the local cache","use":"sync"}}
	cmd.Use = "sync"
	cmd.Short = "Fetch the records of the configured shared repos into the local cache"
	// /sedum:describeCommand
	// sedum:setRun {"tier":"owned","record":"prov-2026-f807707a","gen":"sha256:3f6d1802d83ba9dcc72ed02c8927ae3f79e56242263a8c87803b425321b88f73","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","kwargs":{"bin":"provenance","command":"sync"}}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runSync(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:none {"tier":"owned","record":"prov-2026-f807707a","gen":"sha256:3f6d1802d83ba9dcc72ed02c8927ae3f79e56242263a8c87803b425321b88f73","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","kwargs":{"arity":"none","bin":"provenance","command":"sync"}}
	cmd.Args = cobra.NoArgs
	// /sedum:setArgs:none
	// sedum:anchor:flags
	// sedum:addFlag:bool {"tier":"owned","record":"prov-2026-f807707a","gen":"sha256:3f6d1802d83ba9dcc72ed02c8927ae3f79e56242263a8c87803b425321b88f73","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"sync","flag":"force","type":"bool","usage":"Ignore the cache TTL and fetch again"}}
	cmd.Flags().Bool("force", false, "Ignore the cache TTL and fetch again")
	// /sedum:addFlag:bool
	// sedum:anchor:subcommands
	return cmd
}
