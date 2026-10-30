package main

import "github.com/spf13/cobra"

func newPublishCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"v":2,"tier":"owned","record":"prov-2026-83d57ec1","gen":"sha256:671154c039eaa7dc75b6ff179083a4dcf289b36501a3ddf2b0cbcd42ba3d794b","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","id":{"bin":"provenance","command":"publish"},"kwargs":"sha256:a6d35515e892360de19799502930bba710697ce6dd24f50e3a0c4a077cfa8dc8","body":"sha256:4fd77f0682047753bcbda15488027265ed5316b50c1f72e5a10f3478b4c4acee"}
	cmd.Use = "publish [id...]"
	cmd.Short = "Publish records to Provisio (they become public on AT Proto)"
	// /sedum:describeCommand
	// sedum:setRun {"v":2,"tier":"owned","record":"prov-2026-83d57ec1","gen":"sha256:671154c039eaa7dc75b6ff179083a4dcf289b36501a3ddf2b0cbcd42ba3d794b","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","id":{"bin":"provenance","command":"publish"},"kwargs":"sha256:dd6ccb786f4fd36221547b4839e6b254fed767f515f1b02999611022bf9e5527","body":"sha256:676d32958a48b21fa038ecb407e492f88eb0dc5809edd09631b09a26bfed3224"}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runPublish(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:any {"v":2,"tier":"owned","record":"prov-2026-83d57ec1","gen":"sha256:671154c039eaa7dc75b6ff179083a4dcf289b36501a3ddf2b0cbcd42ba3d794b","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","id":{"bin":"provenance","command":"publish"},"kwargs":"sha256:e1fcd39639a1448bfe4e0fd7344fce1820dc08e068be0bee1586193994faacfc","body":"sha256:b30b0c92d5bd257e52e3c8256ebc76c0673d14e4c2e996925a6699b264276231"}
	cmd.Args = cobra.ArbitraryArgs
	// /sedum:setArgs:any
	// sedum:anchor:flags
	// sedum:addFlag:bool {"v":2,"tier":"owned","record":"prov-2026-83d57ec1","gen":"sha256:671154c039eaa7dc75b6ff179083a4dcf289b36501a3ddf2b0cbcd42ba3d794b","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"publish","flag":"all"},"kwargs":"sha256:7cc0575685d92b1b2fd94b00535437b8881c863f38ed9519438cfa41ff3e64c0","body":"sha256:ad766f7841aad7acb79e6092605bcf87add636b33a6d92006d9aa18fb09a8512"}
	cmd.Flags().Bool("all", false, "Publish every record")
	// /sedum:addFlag:bool
	// sedum:addFlag:bool {"v":2,"tier":"owned","record":"prov-2026-83d57ec1","gen":"sha256:671154c039eaa7dc75b6ff179083a4dcf289b36501a3ddf2b0cbcd42ba3d794b","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"publish","flag":"yes"},"kwargs":"sha256:9cb055d762ad0839ff3eef30de11569a984c76cbaaf5c8a9a5a47f2af295f873","body":"sha256:00d4c89632167e2077977be476726a058dd9a27b82243fcf10bcabdc321053c7"}
	cmd.Flags().Bool("yes", false, "Confirm without asking (for scripts)")
	// /sedum:addFlag:bool
	// sedum:addFlag:string {"v":2,"tier":"owned","record":"prov-2026-83d57ec1","gen":"sha256:671154c039eaa7dc75b6ff179083a4dcf289b36501a3ddf2b0cbcd42ba3d794b","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"publish","flag":"project"},"kwargs":"sha256:55b6ed95ed149132f6dfb8ae4cca9b41887e495f9413c40c884fac4498f1825b","body":"sha256:d9893f998f810a989682d780cc86ac1913ac5f26970e2271ec71e6b156c55552"}
	cmd.Flags().String("project", "", "AT URI of the Provisio project the records belong to")
	// /sedum:addFlag:string
	// sedum:addFlag:bool {"v":2,"tier":"owned","record":"prov-2026-83d57ec1","gen":"sha256:671154c039eaa7dc75b6ff179083a4dcf289b36501a3ddf2b0cbcd42ba3d794b","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"publish","flag":"plan"},"kwargs":"sha256:298f5f1e1e69addba9828c825c57bcbd48987aaa6341f1db8264536083b3fb49","body":"sha256:0fa3ded458d5f1eb430fee24ce607edf70f94cea3b2386de4246c75cc188e1de"}
	cmd.Flags().Bool("plan", false, "List what publishing would write, per record; writes nothing and asks nothing")
	// /sedum:addFlag:bool
	// sedum:anchor:subcommands
	return cmd
}
