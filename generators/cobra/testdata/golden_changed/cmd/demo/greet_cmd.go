package main

import "github.com/spf13/cobra"

func newGreetCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"tier":"owned","record":"prov-2026-0000c0b1","gen":"sha256:5062a337e1576589604ab1efde948f0a19edcbda6e2a5a31ece0ee2edfec42ca","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","kwargs":{"bin":"demo","command":"greet","short":"Greet someone warmly","use":"greet <name>"}}
	cmd.Use = "greet <name>"
	cmd.Short = "Greet someone warmly"
	// /sedum:describeCommand
	// sedum:setRun {"tier":"owned","record":"prov-2026-0000c0b1","gen":"sha256:5062a337e1576589604ab1efde948f0a19edcbda6e2a5a31ece0ee2edfec42ca","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","kwargs":{"bin":"demo","command":"greet"}}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runGreet(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:max_one {"tier":"owned","record":"prov-2026-0000c0b1","gen":"sha256:5062a337e1576589604ab1efde948f0a19edcbda6e2a5a31ece0ee2edfec42ca","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","kwargs":{"arity":"max_one","bin":"demo","command":"greet"}}
	cmd.Args = cobra.MaximumNArgs(1)
	// /sedum:setArgs:max_one
	// sedum:anchor:flags
	// sedum:addDefaultFlag {"tier":"owned","record":"prov-2026-0000c0b1","gen":"sha256:5062a337e1576589604ab1efde948f0a19edcbda6e2a5a31ece0ee2edfec42ca","pkg":"sha256:f3c83dc533de1ffaec737cb62c02055fa2b4fb8d87929314763f7ddb9ff23364","kwargs":{"bin":"demo","command":"greet","default":"hello","flag":"greeting","usage":"Greeting word"}}
	cmd.Flags().String("greeting", "hello", "Greeting word")
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
