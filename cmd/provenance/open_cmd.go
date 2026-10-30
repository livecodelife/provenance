package main

import "github.com/spf13/cobra"

func newOpenCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"v":2,"tier":"owned","record":"prov-2026-6db88b75","gen":"sha256:18e81b445bc0135242f33754116c5073aa843ca855024412b23f848a00f1084b","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","id":{"bin":"provenance","command":"open"},"kwargs":"sha256:430b0173535321390663c7c4e16f3e32e6f9be22170182f41438456535b93274","body":"sha256:a1a60f92c4e96ebcebf5d5ce2b67b069bbc56196403026f692bb67f2bab003b7"}
	cmd.Use = "open"
	cmd.Short = "Transition a record from draft to open (enables enforcement)"
	// /sedum:describeCommand
	// sedum:setRun {"v":2,"tier":"owned","record":"prov-2026-6db88b75","gen":"sha256:18e81b445bc0135242f33754116c5073aa843ca855024412b23f848a00f1084b","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","id":{"bin":"provenance","command":"open"},"kwargs":"sha256:e1bcceda3bda344b4c39739d5d9da6f5fdedca5f6b793a889959657bf7672faf","body":"sha256:b05f00b66ea175c64b97e3303ca7e70dbf224d05dae8eee9b2e61ba4de2c7a6f"}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runOpen(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:none {"v":2,"tier":"owned","record":"prov-2026-6db88b75","gen":"sha256:18e81b445bc0135242f33754116c5073aa843ca855024412b23f848a00f1084b","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","id":{"bin":"provenance","command":"open"},"kwargs":"sha256:4f57cfaa85cdf507fba992c9ca1ae05f00634460c540db0461b41df4cf65a95e","body":"sha256:778b1e3eb6004fe613ecb8793efaf401a29f1afdc382707cea55c05601aefd79"}
	cmd.Args = cobra.NoArgs
	// /sedum:setArgs:none
	// sedum:anchor:flags
	// sedum:addFlag:string {"v":2,"tier":"owned","record":"prov-2026-6db88b75","gen":"sha256:18e81b445bc0135242f33754116c5073aa843ca855024412b23f848a00f1084b","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"open","flag":"record"},"kwargs":"sha256:00946a59e9024ed4aabdcf66fc32ac3d9cd0de1c37bfdcb8af219adaa93f7f43","body":"sha256:f672cf1fc189a6519d916eea2d745271ae54ec3b84ce1f99159d6d7806cb0e53"}
	cmd.Flags().String("record", "", "Record to transition from draft to open")
	// /sedum:addFlag:string
	// sedum:anchor:subcommands
	return cmd
}
