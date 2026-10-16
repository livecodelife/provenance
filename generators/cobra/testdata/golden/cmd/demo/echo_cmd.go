package main

import "github.com/spf13/cobra"

func newEchoCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"tier":"owned","record":"prov-2026-0000c0b1","gen":"sha256:5062a337e1576589604ab1efde948f0a19edcbda6e2a5a31ece0ee2edfec42ca","pkg":"sha256:f8d152763fd9ce930ecfcc140104d898a335e71fb55263e286a4d13a61752868","kwargs":{"bin":"demo","command":"echo","short":"Echo words","use":"echo <word>..."}}
	cmd.Use = "echo <word>..."
	cmd.Short = "Echo words"
	// /sedum:describeCommand
	// sedum:setRun {"tier":"owned","record":"prov-2026-0000c0b1","gen":"sha256:5062a337e1576589604ab1efde948f0a19edcbda6e2a5a31ece0ee2edfec42ca","pkg":"sha256:fadae715925d582c263ce084f15c491995be26d77f0489abc836c658e094878f","kwargs":{"bin":"demo","command":"echo"}}
	cmd.RunE = runEcho
	// /sedum:setRun
	// sedum:setArgs:min_one {"tier":"owned","record":"prov-2026-0000c0b1","gen":"sha256:5062a337e1576589604ab1efde948f0a19edcbda6e2a5a31ece0ee2edfec42ca","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","kwargs":{"arity":"min_one","bin":"demo","command":"echo"}}
	cmd.Args = cobra.MinimumNArgs(1)
	// /sedum:setArgs:min_one
	// sedum:anchor:flags
	// sedum:anchor:subcommands
	return cmd
}
