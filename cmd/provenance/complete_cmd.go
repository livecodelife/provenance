package main

import "github.com/spf13/cobra"

func newCompleteCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"v":2,"tier":"owned","record":"prov-2026-41fb548c","gen":"sha256:953e06b4e4abb9f509dd88530863fa800eeb22ece21fb38f9e9dfc6e57e8e5cf","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","id":{"bin":"provenance","command":"complete"},"kwargs":"sha256:92a70c32d28987dcd19d48829d0a419920696d210d867433e3eefefb5b4e17f0","body":"sha256:88b24665aa1ac1de17e0573e7948e3324fef912627731cbc1d0feca9e69ee540"}
	cmd.Use = "complete"
	cmd.Short = "Mark an open record as implemented (seals it at HEAD)"
	// /sedum:describeCommand
	// sedum:setRun {"v":2,"tier":"owned","record":"prov-2026-41fb548c","gen":"sha256:953e06b4e4abb9f509dd88530863fa800eeb22ece21fb38f9e9dfc6e57e8e5cf","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","id":{"bin":"provenance","command":"complete"},"kwargs":"sha256:8448526f028d904a182d3f00323d53140a08444c90b13d1b14b2b97926d4053d","body":"sha256:0201b5cbfbe3dde38b2ca3457eb0c2ac4fc75f4c32feb7907e774de1d82e814f"}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runComplete(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:none {"v":2,"tier":"owned","record":"prov-2026-41fb548c","gen":"sha256:953e06b4e4abb9f509dd88530863fa800eeb22ece21fb38f9e9dfc6e57e8e5cf","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","id":{"bin":"provenance","command":"complete"},"kwargs":"sha256:1a94419a661d6516db002a1b7f10d64acb3d5ae6b1396362b5ab2289fce97ac1","body":"sha256:778b1e3eb6004fe613ecb8793efaf401a29f1afdc382707cea55c05601aefd79"}
	cmd.Args = cobra.NoArgs
	// /sedum:setArgs:none
	// sedum:anchor:flags
	// sedum:addFlag:string {"v":2,"tier":"owned","record":"prov-2026-41fb548c","gen":"sha256:953e06b4e4abb9f509dd88530863fa800eeb22ece21fb38f9e9dfc6e57e8e5cf","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"complete","flag":"record"},"kwargs":"sha256:7b642574a6db23b745c433649c01fc62f16a65a4ed003be8b427581f801ba69a","body":"sha256:0c3d1c300cb14063f1053cf6aea4f72908fa1b3cbfcaf3219e7b51231de49802"}
	cmd.Flags().String("record", "", "Record to mark as implemented")
	// /sedum:addFlag:string
	// sedum:addFlag:bool {"v":2,"tier":"owned","record":"prov-2026-41fb548c","gen":"sha256:953e06b4e4abb9f509dd88530863fa800eeb22ece21fb38f9e9dfc6e57e8e5cf","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"complete","flag":"force"},"kwargs":"sha256:690a0bd687b20372b24b355c09698c9deb013157d692bce90a664c637a255a9f","body":"sha256:9fafcb48002f0d6f0f59cf3c37dd5f1bf97d5681165af2b640c588a97cf13afe"}
	cmd.Flags().Bool("force", false, "Skip the spec existence check and spec runs")
	// /sedum:addFlag:bool
	// sedum:anchor:subcommands
	return cmd
}
