package main

import "github.com/spf13/cobra"

func newStatusCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"v":2,"tier":"owned","record":"prov-2026-12b1ca97","gen":"sha256:3d5eb5aff0b580425966bd18cb1e038d9fef9018ce5f4611f46a803ad6ee5bc9","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","id":{"bin":"provenance","command":"status"},"kwargs":"sha256:1071bc9b8d8d283262c33f644b6ab987b1a2675e10e1bfd9d73fb3d5151134c8","body":"sha256:20122eb6d7fcdca87943a637719463f1db63d3b13173800e26e471032fea5a04"}
	cmd.Use = "status"
	cmd.Short = "Show record status"
	// /sedum:describeCommand
	// sedum:setRun {"v":2,"tier":"owned","record":"prov-2026-12b1ca97","gen":"sha256:3d5eb5aff0b580425966bd18cb1e038d9fef9018ce5f4611f46a803ad6ee5bc9","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","id":{"bin":"provenance","command":"status"},"kwargs":"sha256:56de959af0e69bfa72739f3fd0ddb7e7b6dd7b860049ef0ba00a67da6d7e9709","body":"sha256:6c230cce80a1d151e80b51da231f18bbf95e9ea25718b33a72235c186ebc16e2"}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runStatus(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:none {"v":2,"tier":"owned","record":"prov-2026-12b1ca97","gen":"sha256:3d5eb5aff0b580425966bd18cb1e038d9fef9018ce5f4611f46a803ad6ee5bc9","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","id":{"bin":"provenance","command":"status"},"kwargs":"sha256:0b6f3adc531a47abd1583d5547718cc0f4d42ba7a0533ea07dc14ef9d5fca4e0","body":"sha256:778b1e3eb6004fe613ecb8793efaf401a29f1afdc382707cea55c05601aefd79"}
	cmd.Args = cobra.NoArgs
	// /sedum:setArgs:none
	// sedum:anchor:flags
	// sedum:addFlag:string {"v":2,"tier":"owned","record":"prov-2026-12b1ca97","gen":"sha256:3d5eb5aff0b580425966bd18cb1e038d9fef9018ce5f4611f46a803ad6ee5bc9","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"status","flag":"record"},"kwargs":"sha256:85dabdf8405cff2a1977f3670cf0dc72714e7ac2a46982dae1a6ddb1d4bf8a2e","body":"sha256:40b2607c60e5a23618cfebef18d5ffbcce95b2678b7baaaece2377238a0e28c4"}
	cmd.Flags().String("record", "", "Show detailed status for a record")
	// /sedum:addFlag:string
	// sedum:addFlag:string {"v":2,"tier":"owned","record":"prov-2026-12b1ca97","gen":"sha256:3d5eb5aff0b580425966bd18cb1e038d9fef9018ce5f4611f46a803ad6ee5bc9","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"status","flag":"filter"},"kwargs":"sha256:18cf2470a392cc31a9a125ed96c33ec74de38fe347f9ea56e62cb0d4f51873e2","body":"sha256:e1df0ab1dd8e8f64c3fc33ebd08300ddd4a53a82a1eb628bca4f632e242e69ce"}
	cmd.Flags().String("filter", "", "Filter by status (draft|open|implemented|deprecated) or tag:name")
	// /sedum:addFlag:string
	// sedum:addFlag:string {"v":2,"tier":"owned","record":"prov-2026-12b1ca97","gen":"sha256:3d5eb5aff0b580425966bd18cb1e038d9fef9018ce5f4611f46a803ad6ee5bc9","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"status","flag":"format"},"kwargs":"sha256:8a8843044b470ce28874a8653539f4305d2332568be6c2bfc90f8ec9a7c2c827","body":"sha256:5a11ae1e68306dbe22a26ffec6af235a5dc44898684981a9cddacca632a7b196"}
	cmd.Flags().String("format", "", "Output format (human|json)")
	// /sedum:addFlag:string
	// sedum:anchor:subcommands
	return cmd
}
