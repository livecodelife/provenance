package main

import "github.com/spf13/cobra"

func newGreetCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"tier":"owned","record":"prov-2026-0000c0b1","gen":"sha256:5062a337e1576589604ab1efde948f0a19edcbda6e2a5a31ece0ee2edfec42ca","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","kwargs":{"bin":"demo","command":"greet","short":"Greet someone","use":"greet <name>"}}
	cmd.Use = "greet <name>"
	cmd.Short = "Greet someone"
	// /sedum:describeCommand
	// sedum:setRun {"tier":"owned","record":"prov-2026-0000c0b1","gen":"sha256:5062a337e1576589604ab1efde948f0a19edcbda6e2a5a31ece0ee2edfec42ca","pkg":"sha256:fadae715925d582c263ce084f15c491995be26d77f0489abc836c658e094878f","kwargs":{"bin":"demo","command":"greet"}}
	cmd.RunE = runGreet
	// /sedum:setRun
	// sedum:setArgs:one {"tier":"owned","record":"prov-2026-0000c0b1","gen":"sha256:5062a337e1576589604ab1efde948f0a19edcbda6e2a5a31ece0ee2edfec42ca","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","kwargs":{"arity":"one","bin":"demo","command":"greet"}}
	cmd.Args = cobra.ExactArgs(1)
	// /sedum:setArgs:one
	// sedum:anchor:flags
	// sedum:addDefaultFlag {"tier":"owned","record":"prov-2026-0000c0b1","gen":"sha256:5062a337e1576589604ab1efde948f0a19edcbda6e2a5a31ece0ee2edfec42ca","pkg":"sha256:f3c83dc533de1ffaec737cb62c02055fa2b4fb8d87929314763f7ddb9ff23364","kwargs":{"bin":"demo","command":"greet","default":"hello","flag":"greeting","usage":"Word to greet with"}}
	cmd.Flags().String("greeting", "hello", "Word to greet with")
	// /sedum:addDefaultFlag
	// sedum:addFlag:bool {"tier":"owned","record":"prov-2026-0000c0b1","gen":"sha256:5062a337e1576589604ab1efde948f0a19edcbda6e2a5a31ece0ee2edfec42ca","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"demo","command":"greet","flag":"shout","type":"bool","usage":"Print in upper case"}}
	cmd.Flags().Bool("shout", false, "Print in upper case")
	// /sedum:addFlag:bool
	// sedum:addFlag:int {"tier":"owned","record":"prov-2026-0000c0b1","gen":"sha256:5062a337e1576589604ab1efde948f0a19edcbda6e2a5a31ece0ee2edfec42ca","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"demo","command":"greet","flag":"times","type":"int","usage":"Repeat count"}}
	cmd.Flags().Int("times", 0, "Repeat count")
	// /sedum:addFlag:int
	// sedum:addFlag:strings {"tier":"owned","record":"prov-2026-0000c0b1","gen":"sha256:5062a337e1576589604ab1efde948f0a19edcbda6e2a5a31ece0ee2edfec42ca","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"demo","command":"greet","flag":"tag","type":"strings","usage":"Tags, repeatable"}}
	cmd.Flags().StringArray("tag", nil, "Tags, repeatable")
	// /sedum:addFlag:strings
	// sedum:anchor:subcommands
	return cmd
}
