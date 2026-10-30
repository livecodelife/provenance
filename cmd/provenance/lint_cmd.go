package main

import "github.com/spf13/cobra"

func newLintCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"v":2,"tier":"owned","record":"prov-2026-a03e6bc3","gen":"sha256:08032897102932b2a0dd0a941644e2f7fce15c3bdb19714d078c7ce333926547","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","id":{"bin":"provenance","command":"lint"},"kwargs":"sha256:f7103f855358d59035b84b127deb4d84e457ad840977bc171a77bc0a5f4c03ac","body":"sha256:4253586e743edeb55371f9ee395195bcb05088c2d3e26e0b66ec8d089ea6361a"}
	cmd.Use = "lint"
	cmd.Short = "Validate provenance records"
	// /sedum:describeCommand
	// sedum:setRun {"v":2,"tier":"owned","record":"prov-2026-a03e6bc3","gen":"sha256:08032897102932b2a0dd0a941644e2f7fce15c3bdb19714d078c7ce333926547","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","id":{"bin":"provenance","command":"lint"},"kwargs":"sha256:a651b403bbf345bf15702d90556c49d517e43617fec543c2143ab42c7bee93cb","body":"sha256:0ceafafa103d075c8c790577041dbc299059b3f97873cb7eb69d4b82e9b8b097"}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runLint(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:none {"v":2,"tier":"owned","record":"prov-2026-a03e6bc3","gen":"sha256:08032897102932b2a0dd0a941644e2f7fce15c3bdb19714d078c7ce333926547","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","id":{"bin":"provenance","command":"lint"},"kwargs":"sha256:225eb198e41738d607d8e3f46e4a1656ab12b4798847d204f8c882921ee2358c","body":"sha256:778b1e3eb6004fe613ecb8793efaf401a29f1afdc382707cea55c05601aefd79"}
	cmd.Args = cobra.NoArgs
	// /sedum:setArgs:none
	// sedum:anchor:flags
	// sedum:addFlag:string {"v":2,"tier":"owned","record":"prov-2026-a03e6bc3","gen":"sha256:08032897102932b2a0dd0a941644e2f7fce15c3bdb19714d078c7ce333926547","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"lint","flag":"record"},"kwargs":"sha256:70946d9d71bbd33c1a74ce32881f48725d565c8520a855ff5d58099444814812","body":"sha256:712aea69ded66df19de91a347801a46314ae0f98a29d17d976472a0c0fae92b2"}
	cmd.Flags().String("record", "", "Lint a single record")
	// /sedum:addFlag:string
	// sedum:addFlag:string {"v":2,"tier":"owned","record":"prov-2026-a03e6bc3","gen":"sha256:08032897102932b2a0dd0a941644e2f7fce15c3bdb19714d078c7ce333926547","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"lint","flag":"enforcement"},"kwargs":"sha256:6238ae54de604c4b8e797727f9d6ad77962a9a7738057eb367576b5b46c1f61b","body":"sha256:03a377de3e34533a7ed61f2890295c52002350f34414d6241ce0c902f503d746"}
	cmd.Flags().String("enforcement", "", "Override enforcement (none|warn|strict)")
	// /sedum:addFlag:string
	// sedum:addFlag:string {"v":2,"tier":"owned","record":"prov-2026-a03e6bc3","gen":"sha256:08032897102932b2a0dd0a941644e2f7fce15c3bdb19714d078c7ce333926547","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"lint","flag":"format"},"kwargs":"sha256:55298609010c0fd57d749e1f13bf149c7c51e511d67e2758c2b8781e950c46b2","body":"sha256:f382c95db73029bc4aed2663135d12178a9024cf01bd034f28f31eb0cbe24706"}
	cmd.Flags().String("format", "", "Output format (human|json|sarif)")
	// /sedum:addFlag:string
	// sedum:addFlag:bool {"v":2,"tier":"owned","record":"prov-2026-a03e6bc3","gen":"sha256:08032897102932b2a0dd0a941644e2f7fce15c3bdb19714d078c7ce333926547","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"lint","flag":"warn"},"kwargs":"sha256:db7971b456ee39cb97dbecdc13276b4ea3ec7a1abdccbfc56716a909272746ca","body":"sha256:fad65ae7836c2c28787b156933ac15c2bb5e0f8c8cdb614d087196ac30b5c62d"}
	cmd.Flags().Bool("warn", false, "Show only warnings in output")
	// /sedum:addFlag:bool
	// sedum:addFlag:bool {"v":2,"tier":"owned","record":"prov-2026-a03e6bc3","gen":"sha256:08032897102932b2a0dd0a941644e2f7fce15c3bdb19714d078c7ce333926547","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"lint","flag":"info"},"kwargs":"sha256:36f512e0c7311d5d3b937dc332df2e6a76347944dfdaa8ef46dc361b36d3c1d5","body":"sha256:9ee1464290b2ac92693344aa53cb4b4c43d5942fe667581b510524a77aa55626"}
	cmd.Flags().Bool("info", false, "Show only informational hints in output")
	// /sedum:addFlag:bool
	// sedum:addFlag:bool {"v":2,"tier":"owned","record":"prov-2026-a03e6bc3","gen":"sha256:08032897102932b2a0dd0a941644e2f7fce15c3bdb19714d078c7ce333926547","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"lint","flag":"all"},"kwargs":"sha256:2da70850453169f504374fbb727446ee267210c57a4dcff8a1457b1d86f1e9a4","body":"sha256:2c9cb95f2775467db71c46dc09dd923d62c96bf8f3845c2b17abcbf87144829a"}
	cmd.Flags().Bool("all", false, "Show all output (errors, warnings, and hints)")
	// /sedum:addFlag:bool
	// sedum:anchor:subcommands
	return cmd
}
