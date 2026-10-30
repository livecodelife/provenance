package main

import "github.com/spf13/cobra"

func newContextCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"v":2,"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:57a095967ea4cbd86875cea0f96e504d2e1d138c4f31fba23d7b01ba011caa69","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","id":{"bin":"provenance","command":"context"},"kwargs":"sha256:1e0fa9950bab0acfd13f5d0876a67f65cee130250654b0c9ee4f867fc079c7db","body":"sha256:5c93ec9d6da17ec29b4c75f5a2512156d6d49a94a7ca8f1d25f7df4eb801457b"}
	cmd.Use = "context [targets...]"
	cmd.Short = "Show the provenance context for paths or records"
	// /sedum:describeCommand
	// sedum:setRun {"v":2,"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:57a095967ea4cbd86875cea0f96e504d2e1d138c4f31fba23d7b01ba011caa69","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","id":{"bin":"provenance","command":"context"},"kwargs":"sha256:02691fe13fca3226c8f871e0260117973260716b0757d776e7ed4d7c05848465","body":"sha256:b5176234d9791394efaca4592f5c1d5ec5e97b4da8b8625c8919c007e4173f5f"}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runContext(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:any {"v":2,"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:57a095967ea4cbd86875cea0f96e504d2e1d138c4f31fba23d7b01ba011caa69","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","id":{"bin":"provenance","command":"context"},"kwargs":"sha256:7165f1299e989f3aa55385b5ceb6cbe0d60e1b97da35ac5a05d3420fbd1795d5","body":"sha256:b30b0c92d5bd257e52e3c8256ebc76c0673d14e4c2e996925a6699b264276231"}
	cmd.Args = cobra.ArbitraryArgs
	// /sedum:setArgs:any
	// sedum:anchor:flags
	// sedum:addFlag:strings {"v":2,"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:57a095967ea4cbd86875cea0f96e504d2e1d138c4f31fba23d7b01ba011caa69","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"context","flag":"files"},"kwargs":"sha256:966e4196c206bcaf76009c1e7aab3c74f5f6a171517f3d30bf50239d840766bc","body":"sha256:a735f2acc5b854691378d94be2aec862885596d513cba267a9fbbab8c3e7c8fb"}
	cmd.Flags().StringArray("files", nil, "Explicit target list (alternative to arguments)")
	// /sedum:addFlag:strings
	// sedum:addFlag:string {"v":2,"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:57a095967ea4cbd86875cea0f96e504d2e1d138c4f31fba23d7b01ba011caa69","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"context","flag":"format"},"kwargs":"sha256:3c526c3b9e851100ea9d983b4d78c08c139fbdbc115e7220a909f6a5a98e0e37","body":"sha256:3709144a8aec1ce2560358b9297df5cf2fdb7620d10df5c4e0b0a60deeec1327"}
	cmd.Flags().String("format", "", "Output format (human|compact|json)")
	// /sedum:addFlag:string
	// sedum:anchor:subcommands
	return cmd
}
