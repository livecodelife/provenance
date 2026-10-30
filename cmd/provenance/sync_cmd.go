package main

import "github.com/spf13/cobra"

func newSyncCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"v":2,"tier":"owned","record":"prov-2026-f807707a","gen":"sha256:3f6d1802d83ba9dcc72ed02c8927ae3f79e56242263a8c87803b425321b88f73","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","id":{"bin":"provenance","command":"sync"},"kwargs":"sha256:2cb56d8d5da691c6a7ec8816fe0d95eb2483c64d8308cdce579042afea694379","body":"sha256:58ee2784f1f2ec0c371692855ef0ac58e5d2a535c53c6a83c8018b32e309fe45"}
	cmd.Use = "sync"
	cmd.Short = "Fetch the records of the configured shared repos into the local cache"
	// /sedum:describeCommand
	// sedum:setRun {"v":2,"tier":"owned","record":"prov-2026-f807707a","gen":"sha256:3f6d1802d83ba9dcc72ed02c8927ae3f79e56242263a8c87803b425321b88f73","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","id":{"bin":"provenance","command":"sync"},"kwargs":"sha256:5620389fe1e53752ef29ea0eb88502349079dd89819f18d985bd0f2702a2fe60","body":"sha256:f2fc53ca59a944455aca13166d2794aeee57272b1ec13aeca91e6793ce1787ad"}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runSync(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:none {"v":2,"tier":"owned","record":"prov-2026-f807707a","gen":"sha256:3f6d1802d83ba9dcc72ed02c8927ae3f79e56242263a8c87803b425321b88f73","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","id":{"bin":"provenance","command":"sync"},"kwargs":"sha256:f0580426a3a6a8e7ef1254c53280be296e797fc24fb55c4830c03d1e9154e278","body":"sha256:778b1e3eb6004fe613ecb8793efaf401a29f1afdc382707cea55c05601aefd79"}
	cmd.Args = cobra.NoArgs
	// /sedum:setArgs:none
	// sedum:anchor:flags
	// sedum:addFlag:bool {"v":2,"tier":"owned","record":"prov-2026-f807707a","gen":"sha256:3f6d1802d83ba9dcc72ed02c8927ae3f79e56242263a8c87803b425321b88f73","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"sync","flag":"force"},"kwargs":"sha256:0fb9280b7324a06056447a5b3c7e7ae072c2724b80167f78f766935df2af8bde","body":"sha256:8c9c7a644b91a4d79d6f20e7035bee89f24167664b2a2ee6a46e648d12530ec7"}
	cmd.Flags().Bool("force", false, "Ignore the cache TTL and fetch again")
	// /sedum:addFlag:bool
	// sedum:anchor:subcommands
	return cmd
}
