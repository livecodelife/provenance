package main

import "github.com/spf13/cobra"

func newStaleCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"tier":"owned","record":"prov-2026-f476b640","gen":"sha256:24167071b71953f56b7a9351b440d09a67530fb3dc4508cd3e405b0ad9bb4806","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","kwargs":{"bin":"provenance","command":"stale","short":"List records whose implements_version is behind their parent","use":"stale"}}
	cmd.Use = "stale"
	cmd.Short = "List records whose implements_version is behind their parent"
	// /sedum:describeCommand
	// sedum:setRun {"tier":"owned","record":"prov-2026-f476b640","gen":"sha256:24167071b71953f56b7a9351b440d09a67530fb3dc4508cd3e405b0ad9bb4806","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","kwargs":{"bin":"provenance","command":"stale"}}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runStale(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:none {"tier":"owned","record":"prov-2026-f476b640","gen":"sha256:24167071b71953f56b7a9351b440d09a67530fb3dc4508cd3e405b0ad9bb4806","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","kwargs":{"arity":"none","bin":"provenance","command":"stale"}}
	cmd.Args = cobra.NoArgs
	// /sedum:setArgs:none
	// sedum:anchor:flags
	// sedum:addFlag:bool {"tier":"owned","record":"prov-2026-f476b640","gen":"sha256:24167071b71953f56b7a9351b440d09a67530fb3dc4508cd3e405b0ad9bb4806","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"stale","flag":"json","type":"bool","usage":"Print as JSON"}}
	cmd.Flags().Bool("json", false, "Print as JSON")
	// /sedum:addFlag:bool
	// sedum:addFlag:bool {"tier":"owned","record":"prov-2026-f476b640","gen":"sha256:24167071b71953f56b7a9351b440d09a67530fb3dc4508cd3e405b0ad9bb4806","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"stale","flag":"exit-code","type":"bool","usage":"Exit non-zero when any record is directly or transitively stale"}}
	cmd.Flags().Bool("exit-code", false, "Exit non-zero when any record is directly or transitively stale")
	// /sedum:addFlag:bool
	// sedum:anchor:subcommands
	return cmd
}
