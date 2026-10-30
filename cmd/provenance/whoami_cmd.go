package main

import "github.com/spf13/cobra"

func newWhoamiCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"v":2,"tier":"owned","record":"prov-2026-215c273d","gen":"sha256:35eaf90eaf676b31114c953aa2ccef582eceebd10849617fd9691d67499118f6","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","id":{"bin":"provenance","command":"whoami"},"kwargs":"sha256:352ac173f7dce2e70c40efe6d9f6141ea8a63443006824ffa329870e1ecc5a2b","body":"sha256:d1d33b8ca302ee839fcb26eebac19a28982708ec71a9bbd7a7ac7cae635f4d8d"}
	cmd.Use = "whoami"
	cmd.Short = "Show the signed-in Provisio account"
	// /sedum:describeCommand
	// sedum:setRun {"v":2,"tier":"owned","record":"prov-2026-215c273d","gen":"sha256:35eaf90eaf676b31114c953aa2ccef582eceebd10849617fd9691d67499118f6","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","id":{"bin":"provenance","command":"whoami"},"kwargs":"sha256:353cc45eeeff17da10c47383fe7f132d84e480899052b58137903c8bb55032d9","body":"sha256:408cf2a05785a3ff6dbcec117abd8fc68bcd7f22a9905564520732efc8e5108c"}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runWhoami(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:none {"v":2,"tier":"owned","record":"prov-2026-215c273d","gen":"sha256:35eaf90eaf676b31114c953aa2ccef582eceebd10849617fd9691d67499118f6","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","id":{"bin":"provenance","command":"whoami"},"kwargs":"sha256:2ecf8b265d7bd38ab2ba5fc0d439d2c4ef7180793762cd2e5c44fa6fe990c5b1","body":"sha256:778b1e3eb6004fe613ecb8793efaf401a29f1afdc382707cea55c05601aefd79"}
	cmd.Args = cobra.NoArgs
	// /sedum:setArgs:none
	// sedum:anchor:flags
	// sedum:addFlag:bool {"v":2,"tier":"owned","record":"prov-2026-215c273d","gen":"sha256:35eaf90eaf676b31114c953aa2ccef582eceebd10849617fd9691d67499118f6","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"whoami","flag":"json"},"kwargs":"sha256:b2940ebf60af5d48c9ef1ad4fafb7884ad71760e1f88352f8b4c1cafdd6f04b8","body":"sha256:d039c1f2caf44bc7dd84431a8617dd2453d265a0337b351f1488a67d6da3b9ce"}
	cmd.Flags().Bool("json", false, "Print as JSON")
	// /sedum:addFlag:bool
	// sedum:anchor:subcommands
	return cmd
}
