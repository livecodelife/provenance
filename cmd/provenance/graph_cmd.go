package main

import "github.com/spf13/cobra"

func newGraphCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"v":2,"tier":"owned","record":"prov-2026-d17b9e44","gen":"sha256:ac383f099a8cbb5db3d785a4c63bd601a4b3df79dd4e69ecde937bf6b98aeca0","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","id":{"bin":"provenance","command":"graph"},"kwargs":"sha256:36bd86930bab11fcfa31bdc89889df27731d25ed98ca4f0644299ee2872770c9","body":"sha256:2951bc66dcc9eb69e9132f6a6db3e23e7d82179e351754b7f7239dabe1b7fb63"}
	cmd.Use = "graph"
	cmd.Short = "Render the record graph: records, parents, pins and versions"
	// /sedum:describeCommand
	// sedum:setRun {"v":2,"tier":"owned","record":"prov-2026-d17b9e44","gen":"sha256:ac383f099a8cbb5db3d785a4c63bd601a4b3df79dd4e69ecde937bf6b98aeca0","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","id":{"bin":"provenance","command":"graph"},"kwargs":"sha256:f4edc5b1ba3b04fef802ac55265fabfa223f2ecb82c119302a25af7ba7ae773c","body":"sha256:99e43ce3640dc1da1f860bcc25c570ddd53264256ca69f88d62019cc38560550"}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runGraph(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:none {"v":2,"tier":"owned","record":"prov-2026-d17b9e44","gen":"sha256:ac383f099a8cbb5db3d785a4c63bd601a4b3df79dd4e69ecde937bf6b98aeca0","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","id":{"bin":"provenance","command":"graph"},"kwargs":"sha256:e316555e66bd3372b7baeb2a864a6e1bc4c570644c4aae4774ad94d3b944e1e5","body":"sha256:778b1e3eb6004fe613ecb8793efaf401a29f1afdc382707cea55c05601aefd79"}
	cmd.Args = cobra.NoArgs
	// /sedum:setArgs:none
	// sedum:anchor:flags
	// sedum:addFlag:string {"v":2,"tier":"owned","record":"prov-2026-d17b9e44","gen":"sha256:ac383f099a8cbb5db3d785a4c63bd601a4b3df79dd4e69ecde937bf6b98aeca0","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"graph","flag":"root"},"kwargs":"sha256:22b77904f6afef389b1119aac64f4a75f4bd43bb130fd7028dd84381efdbb162","body":"sha256:cb7c4409d1cc8a6c1d7ed7f50c935e0426c6237b6cd444b7833e65169c10ffc3"}
	cmd.Flags().String("root", "", "Start from one record: its parent (one level up) and its descendants")
	// /sedum:addFlag:string
	// sedum:addFlag:string {"v":2,"tier":"owned","record":"prov-2026-d17b9e44","gen":"sha256:ac383f099a8cbb5db3d785a4c63bd601a4b3df79dd4e69ecde937bf6b98aeca0","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"graph","flag":"filter"},"kwargs":"sha256:03bae03b2fb48dd461d2a4a3c96c1af7cc35fdd2b5f385ede98bdac18af88dd2","body":"sha256:f496c1d51ecc24b4b0e429f6e6fd2c7105b0d321db282a501bc57499115a3d32"}
	cmd.Flags().String("filter", "", "Show only records with this status (draft|open|implemented|deprecated)")
	// /sedum:addFlag:string
	// sedum:addFlag:string {"v":2,"tier":"owned","record":"prov-2026-d17b9e44","gen":"sha256:ac383f099a8cbb5db3d785a4c63bd601a4b3df79dd4e69ecde937bf6b98aeca0","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"graph","flag":"format"},"kwargs":"sha256:235f83ba2ea52cd7038a11400c8b9c11cacf0f9da0f57d18c3a97395c14a76df","body":"sha256:494e1667e3c85ee9914607f991903ed4e9216efaa33f56a5cb051de60fe20295"}
	cmd.Flags().String("format", "", "Output format (human|json|dot)")
	// /sedum:addFlag:string
	// sedum:anchor:subcommands
	return cmd
}
