package main

import "github.com/spf13/cobra"

func newShowIDCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"tier":"owned","record":"prov-2026-0000c0b1","gen":"sha256:5062a337e1576589604ab1efde948f0a19edcbda6e2a5a31ece0ee2edfec42ca","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","kwargs":{"bin":"demo","command":"show-id","short":"Show the demo id","use":"show-id"}}
	cmd.Use = "show-id"
	cmd.Short = "Show the demo id"
	// /sedum:describeCommand
	// sedum:setRun {"tier":"owned","record":"prov-2026-0000c0b1","gen":"sha256:5062a337e1576589604ab1efde948f0a19edcbda6e2a5a31ece0ee2edfec42ca","pkg":"sha256:fadae715925d582c263ce084f15c491995be26d77f0489abc836c658e094878f","kwargs":{"bin":"demo","command":"show-id"}}
	cmd.RunE = runShowID
	// /sedum:setRun
	// sedum:setArgs:none {"tier":"owned","record":"prov-2026-0000c0b1","gen":"sha256:5062a337e1576589604ab1efde948f0a19edcbda6e2a5a31ece0ee2edfec42ca","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","kwargs":{"arity":"none","bin":"demo","command":"show-id"}}
	cmd.Args = cobra.NoArgs
	// /sedum:setArgs:none
	// sedum:anchor:flags
	// sedum:anchor:subcommands
	return cmd
}
