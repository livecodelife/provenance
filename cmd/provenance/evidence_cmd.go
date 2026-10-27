package main

import "github.com/spf13/cobra"

func newEvidenceCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"v":2,"tier":"owned","record":"prov-2026-15f42a00","gen":"sha256:9960743cde810036bdd5d9af367a4c54dcc99657e529725ea71d385c210784ac","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","id":{"bin":"provenance","command":"evidence"},"kwargs":"sha256:9b78f4945b86e120d5bd830001ad0b3d81cc6e14c63a760b15078faf0aaf6a90","body":"sha256:238366f59e9f90e4bdb4ddb834b47278b2189926fa97557df4f2de089ff47672"}
	cmd.Use = "evidence"
	cmd.Short = "Work with a record's verification evidence"
	// /sedum:describeCommand
	// sedum:setArgs:none {"v":2,"tier":"owned","record":"prov-2026-15f42a00","gen":"sha256:9960743cde810036bdd5d9af367a4c54dcc99657e529725ea71d385c210784ac","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","id":{"bin":"provenance","command":"evidence"},"kwargs":"sha256:5caafe366b3ef519f2aae16f663caaf7afed4f5bba9eb4ac0b2579fe9ef50757","body":"sha256:778b1e3eb6004fe613ecb8793efaf401a29f1afdc382707cea55c05601aefd79"}
	cmd.Args = cobra.NoArgs
	// /sedum:setArgs:none
	// sedum:anchor:flags
	// sedum:anchor:subcommands
	// sedum:registerSubcommand {"v":2,"tier":"owned","record":"prov-2026-15f42a00","gen":"sha256:9960743cde810036bdd5d9af367a4c54dcc99657e529725ea71d385c210784ac","pkg":"sha256:8a9e0a8ee811e7c629264e26f09a626602890b11fd82710cf10ae73e3da1cf7a","id":{"bin":"provenance","command":"evidence-refresh","parent":"evidence"},"kwargs":"sha256:37f80a0604d2ed102359037149d7a6c05fb85aeba5265d253912da56e60a94aa","body":"sha256:7b72556f38adb157643077bdacd2568976de5260b08dcc5f9478429baadf4edd"}
	cmd.AddCommand(newEvidenceRefreshCmd())
	// /sedum:registerSubcommand
	return cmd
}
