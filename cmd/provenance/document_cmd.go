package main

import "github.com/spf13/cobra"

func newDocumentCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"v":2,"tier":"owned","record":"prov-2026-36b81f10","gen":"sha256:da0a6a84bf97ac25549e05b29429583669ff3d83fc5d84d6e6f3383a7d3b06f8","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","id":{"bin":"provenance","command":"document"},"kwargs":"sha256:ef4a674c3f6fab124e5eb3fed3d0286ce339bcfda01f824c6f7e9ca9754f2ab9","body":"sha256:90a7ad99cbe5b901c0d2a85cce300c43b72ba36074d7ac05f4c388c0548c88f1"}
	cmd.Use = "document"
	cmd.Short = "Print the behaviour document of the briefs and blueprints, or apply an edited one"
	// /sedum:describeCommand
	// sedum:setRun {"v":2,"tier":"owned","record":"prov-2026-36b81f10","gen":"sha256:da0a6a84bf97ac25549e05b29429583669ff3d83fc5d84d6e6f3383a7d3b06f8","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","id":{"bin":"provenance","command":"document"},"kwargs":"sha256:b487580e966a25087263ad732d40f1dd0911fbef89042c7787e3a4b6d80762a1","body":"sha256:6f7068521ac807517be9f37595dae7e56c59bb475f0ff7b8e74131236e3dcb94"}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runDocument(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:none {"v":2,"tier":"owned","record":"prov-2026-36b81f10","gen":"sha256:da0a6a84bf97ac25549e05b29429583669ff3d83fc5d84d6e6f3383a7d3b06f8","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","id":{"bin":"provenance","command":"document"},"kwargs":"sha256:15a897c156fc125a536f4004ba5fced03622c538cc5f639c9fa7315c27a4b4ea","body":"sha256:778b1e3eb6004fe613ecb8793efaf401a29f1afdc382707cea55c05601aefd79"}
	cmd.Args = cobra.NoArgs
	// /sedum:setArgs:none
	// sedum:anchor:flags
	// sedum:addFlag:strings {"v":2,"tier":"owned","record":"prov-2026-36b81f10","gen":"sha256:da0a6a84bf97ac25549e05b29429583669ff3d83fc5d84d6e6f3383a7d3b06f8","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"document","flag":"scope"},"kwargs":"sha256:49487d9e327c965a240ada886c85fd52107788c89413566104571c087b6040e8","body":"sha256:af83587aa3e4ff49404f55f68b71ac24d84191927067dcb06b70ab43314e6d00"}
	cmd.Flags().StringArray("scope", nil, "Only records whose affected_scope covers this path (repeatable)")
	// /sedum:addFlag:strings
	// sedum:addDefaultFlag {"v":2,"tier":"owned","record":"prov-2026-36b81f10","gen":"sha256:da0a6a84bf97ac25549e05b29429583669ff3d83fc5d84d6e6f3383a7d3b06f8","pkg":"sha256:f3c83dc533de1ffaec737cb62c02055fa2b4fb8d87929314763f7ddb9ff23364","id":{"bin":"provenance","command":"document","flag":"format"},"kwargs":"sha256:543b73137dad2686d15406e346e47fdd0ada620493a79ca1b6890e631f769354","body":"sha256:c7f36cffe2f665f9b7f58658d7af16a71e72a66dfc4b8a5677cebedc0ad14898"}
	cmd.Flags().String("format", "md", "Output format (only md)")
	// /sedum:addDefaultFlag
	// sedum:addShortFlag:string {"v":2,"tier":"owned","record":"prov-2026-36b81f10","gen":"sha256:da0a6a84bf97ac25549e05b29429583669ff3d83fc5d84d6e6f3383a7d3b06f8","pkg":"sha256:317b97244b0fe957a6e0b80cdfa819a8afaae50b42b10385c9cd6729cb74fd3a","id":{"bin":"provenance","command":"document","flag":"out"},"kwargs":"sha256:c62fdd4d61fb28e24cf019b1b71b018daa5a356f10177cd397a221326fbc1c29","body":"sha256:fa10478399082a3e49bc79dec3ddd701092cc50a0331336ae892309c5d36872d"}
	cmd.Flags().StringP("out", "o", "", "Write the document to this file instead of standard output")
	// /sedum:addShortFlag:string
	// sedum:addFlag:bool {"v":2,"tier":"owned","record":"prov-2026-36b81f10","gen":"sha256:da0a6a84bf97ac25549e05b29429583669ff3d83fc5d84d6e6f3383a7d3b06f8","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"document","flag":"ids"},"kwargs":"sha256:4d7611fc03f126c94bfca25f8cf4cb134ec073a58254bc9295d345bdefeb5742","body":"sha256:fbdf708ad30e72cde01582a1eafe238c03aa7d8ebca77aa32889cc8f6c7108d2"}
	cmd.Flags().Bool("ids", false, "Show record ids in the headings")
	// /sedum:addFlag:bool
	// sedum:addFlag:string {"v":2,"tier":"owned","record":"prov-2026-36b81f10","gen":"sha256:da0a6a84bf97ac25549e05b29429583669ff3d83fc5d84d6e6f3383a7d3b06f8","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"document","flag":"apply"},"kwargs":"sha256:26faf6ba3fcdb1430f28fb560d36fed83767f85ac6b2a6fb1d15e68e946fcf9f","body":"sha256:f58259bb3aaa4663e279d656b528c34fce5d44e32d6848d1d10518a5e22aada8"}
	cmd.Flags().String("apply", "", "Plan the changes made in this edited document file")
	// /sedum:addFlag:string
	// sedum:addFlag:bool {"v":2,"tier":"owned","record":"prov-2026-36b81f10","gen":"sha256:da0a6a84bf97ac25549e05b29429583669ff3d83fc5d84d6e6f3383a7d3b06f8","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"document","flag":"write"},"kwargs":"sha256:24a932c1af9363d2457c14ad1a49743cfeccfbbedac67342e87171088d1f86fb","body":"sha256:a49e58e991d22eb91bea1267929544e28c02a985c64176f37c0f8357e29adc4f"}
	cmd.Flags().Bool("write", false, "Apply the plan to the records (needs --apply and --reason)")
	// /sedum:addFlag:bool
	// sedum:addFlag:string {"v":2,"tier":"owned","record":"prov-2026-36b81f10","gen":"sha256:da0a6a84bf97ac25549e05b29429583669ff3d83fc5d84d6e6f3383a7d3b06f8","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"document","flag":"reason"},"kwargs":"sha256:cb30f1421a03f0015dcbac7f79bd46f6a2ece135dcc8b4f0b51c9d787cfdd8f3","body":"sha256:9372fee8d86e416100940e35c1866246ca144c26a0f0ada4a589615066b90fd7"}
	cmd.Flags().String("reason", "", "Why these revisions exist (required with --write)")
	// /sedum:addFlag:string
	// sedum:addFlag:bool {"v":2,"tier":"owned","record":"prov-2026-36b81f10","gen":"sha256:da0a6a84bf97ac25549e05b29429583669ff3d83fc5d84d6e6f3383a7d3b06f8","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"document","flag":"skip-conflicts"},"kwargs":"sha256:9b181248c9ac3540ccc62b17f187e0a9566c1c507feb5a14922d1548d102a975","body":"sha256:da441ac619787b7d80a2fb68fd9eef1c4ff74de419585086bc6c8dbbd9d65bf5"}
	cmd.Flags().Bool("skip-conflicts", false, "With --write, apply the other changes and skip conflicting sections")
	// /sedum:addFlag:bool
	// sedum:anchor:subcommands
	return cmd
}
