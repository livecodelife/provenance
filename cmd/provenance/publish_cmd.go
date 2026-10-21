package main

import "github.com/spf13/cobra"

func newPublishCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"tier":"owned","record":"prov-2026-83d57ec1","gen":"sha256:17956fb64b77b361bcbf762257b64c9341f58a44819468b083c8b9e970540a94","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","kwargs":{"bin":"provenance","command":"publish","short":"Publish records to Provisio (they become public on AT Proto)","use":"publish [id...]"}}
	cmd.Use = "publish [id...]"
	cmd.Short = "Publish records to Provisio (they become public on AT Proto)"
	// /sedum:describeCommand
	// sedum:setRun {"tier":"owned","record":"prov-2026-83d57ec1","gen":"sha256:17956fb64b77b361bcbf762257b64c9341f58a44819468b083c8b9e970540a94","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","kwargs":{"bin":"provenance","command":"publish"}}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runPublish(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:any {"tier":"owned","record":"prov-2026-83d57ec1","gen":"sha256:17956fb64b77b361bcbf762257b64c9341f58a44819468b083c8b9e970540a94","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","kwargs":{"arity":"any","bin":"provenance","command":"publish"}}
	cmd.Args = cobra.ArbitraryArgs
	// /sedum:setArgs:any
	// sedum:anchor:flags
	// sedum:addFlag:bool {"tier":"owned","record":"prov-2026-83d57ec1","gen":"sha256:17956fb64b77b361bcbf762257b64c9341f58a44819468b083c8b9e970540a94","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"publish","flag":"all","type":"bool","usage":"Publish every record"}}
	cmd.Flags().Bool("all", false, "Publish every record")
	// /sedum:addFlag:bool
	// sedum:addFlag:bool {"tier":"owned","record":"prov-2026-83d57ec1","gen":"sha256:17956fb64b77b361bcbf762257b64c9341f58a44819468b083c8b9e970540a94","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"publish","flag":"yes","type":"bool","usage":"Confirm without asking (for scripts)"}}
	cmd.Flags().Bool("yes", false, "Confirm without asking (for scripts)")
	// /sedum:addFlag:bool
	// sedum:addFlag:string {"tier":"owned","record":"prov-2026-83d57ec1","gen":"sha256:17956fb64b77b361bcbf762257b64c9341f58a44819468b083c8b9e970540a94","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"publish","flag":"project","type":"string","usage":"AT URI of the Provisio project the records belong to"}}
	cmd.Flags().String("project", "", "AT URI of the Provisio project the records belong to")
	// /sedum:addFlag:string
	// sedum:addFlag:bool {"tier":"owned","record":"prov-2026-83d57ec1","gen":"sha256:17956fb64b77b361bcbf762257b64c9341f58a44819468b083c8b9e970540a94","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"publish","flag":"plan","type":"bool","usage":"List what publishing would write, per record; writes nothing and asks nothing"}}
	cmd.Flags().Bool("plan", false, "List what publishing would write, per record; writes nothing and asks nothing")
	// /sedum:addFlag:bool
	// sedum:anchor:subcommands
	return cmd
}
