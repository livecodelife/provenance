package main

import "github.com/spf13/cobra"

func newInitCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"v":2,"tier":"owned","record":"prov-2026-1a5d2549","gen":"sha256:95782fcfe653edd478163f4fae54335b865cf426a0e0fccd5f1fc4c3bb633ffe","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","id":{"bin":"provenance","command":"init"},"kwargs":"sha256:616c2521e2e3693b63365c5ac0e5211472b7226fee27ec6287d70c6a1c029de9","body":"sha256:0f859a120600b794dfbb62b14b0a4f9b7617d715fe5180d546525e60ae01620f"}
	cmd.Use = "init"
	cmd.Short = "Scaffold provenance in a new project"
	// /sedum:describeCommand
	// sedum:setRun {"v":2,"tier":"owned","record":"prov-2026-1a5d2549","gen":"sha256:95782fcfe653edd478163f4fae54335b865cf426a0e0fccd5f1fc4c3bb633ffe","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","id":{"bin":"provenance","command":"init"},"kwargs":"sha256:49ba7e13aa36d288fc35e0b7a35dcf581f71d967e1d1fae31cb713f30cdcb3db","body":"sha256:e9215e5f8cdb5164de91ba5cc5b614b095935461ffc46375bdf7300747a8a7e1"}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runInit(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:none {"v":2,"tier":"owned","record":"prov-2026-1a5d2549","gen":"sha256:95782fcfe653edd478163f4fae54335b865cf426a0e0fccd5f1fc4c3bb633ffe","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","id":{"bin":"provenance","command":"init"},"kwargs":"sha256:7f166a5a5ba629238d67aa4895508fd7149565b1c58446387a511eddb6d9972e","body":"sha256:778b1e3eb6004fe613ecb8793efaf401a29f1afdc382707cea55c05601aefd79"}
	cmd.Args = cobra.NoArgs
	// /sedum:setArgs:none
	// sedum:anchor:flags
	// sedum:addFlag:string {"v":2,"tier":"owned","record":"prov-2026-1a5d2549","gen":"sha256:95782fcfe653edd478163f4fae54335b865cf426a0e0fccd5f1fc4c3bb633ffe","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"init","flag":"dir"},"kwargs":"sha256:8a0e30f8111c8a51532565ff1a806958856b2f12ee0a6a9762aea2b42ca0ca18","body":"sha256:298031fc91a7fd37d6bae588e5e7885393f0e8c7d2c618365cd1dafc6c7c7b9c"}
	cmd.Flags().String("dir", "", "Project directory (default: the current directory)")
	// /sedum:addFlag:string
	// sedum:addFlag:bool {"v":2,"tier":"owned","record":"prov-2026-1a5d2549","gen":"sha256:95782fcfe653edd478163f4fae54335b865cf426a0e0fccd5f1fc4c3bb633ffe","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"init","flag":"force"},"kwargs":"sha256:b43e77285c20c9425eeee823b87e6e4bd73e670c349de40c944a64cd405d0ce9","body":"sha256:855f39ba14d382fe3cc84a7f8fa5b32540d20213d733b4a18cc22177c5c3a17a"}
	cmd.Flags().Bool("force", false, "Overwrite an existing .provenance.yml (the old one is saved as .provenance.yml.bak)")
	// /sedum:addFlag:bool
	// sedum:anchor:subcommands
	return cmd
}
