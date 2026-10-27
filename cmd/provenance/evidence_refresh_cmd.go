package main

import "github.com/spf13/cobra"

func newEvidenceRefreshCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"v":2,"tier":"owned","record":"prov-2026-15f42a00","gen":"sha256:9960743cde810036bdd5d9af367a4c54dcc99657e529725ea71d385c210784ac","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","id":{"bin":"provenance","command":"evidence-refresh"},"kwargs":"sha256:d8053264fc4d99eab7109a1292a7e0cc8ec7fe303fc940eaafd425f6f33781df","body":"sha256:dd65af3f39403064f1382c3d5914fd99f1f85782d390222dbd6aab95582f898f"}
	cmd.Use = "refresh [id]"
	cmd.Short = "Append fresh evidence to an implemented record without a new version"
	// /sedum:describeCommand
	// sedum:setRun {"v":2,"tier":"owned","record":"prov-2026-15f42a00","gen":"sha256:9960743cde810036bdd5d9af367a4c54dcc99657e529725ea71d385c210784ac","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","id":{"bin":"provenance","command":"evidence-refresh"},"kwargs":"sha256:767dd818f3d2980328b07aa22167455982c6bbe0c2738695e103b1f14fc268b3","body":"sha256:c1d045267c1c097c973d75eef26affe1b155af75627795fdd6d4dc5dae328346"}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runEvidenceRefresh(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:max_one {"v":2,"tier":"owned","record":"prov-2026-15f42a00","gen":"sha256:9960743cde810036bdd5d9af367a4c54dcc99657e529725ea71d385c210784ac","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","id":{"bin":"provenance","command":"evidence-refresh"},"kwargs":"sha256:0833834244ddd7f492358d994adc16be9eb3b18d026b264ee9a644c6e2fd5585","body":"sha256:186046a2435bbf4d0eae67e61e8a3909c91fd72c18b73dd2ff749798e3517ef3"}
	cmd.Args = cobra.MaximumNArgs(1)
	// /sedum:setArgs:max_one
	// sedum:anchor:flags
	// sedum:addFlag:string {"v":2,"tier":"owned","record":"prov-2026-15f42a00","gen":"sha256:9960743cde810036bdd5d9af367a4c54dcc99657e529725ea71d385c210784ac","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"evidence-refresh","flag":"record"},"kwargs":"sha256:3ae1f884c2e913e147635494b6f9bb17642b225cb31bfb8cd834ea318c15854d","body":"sha256:a4ccec74e3753c0c26186857caf9184fd13ce4e0a844b3e17aa997b4c72f5baf"}
	cmd.Flags().String("record", "", "Record to refresh (or give it as the argument)")
	// /sedum:addFlag:string
	// sedum:addFlag:string {"v":2,"tier":"owned","record":"prov-2026-15f42a00","gen":"sha256:9960743cde810036bdd5d9af367a4c54dcc99657e529725ea71d385c210784ac","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"evidence-refresh","flag":"reason"},"kwargs":"sha256:9e559060c37e40d774ca0794a0e81da78f9888e1fb4c948f381f4f095e67dcd4","body":"sha256:ef47d66260b64b938e622fe3c84868b2cdf3d284775568d0c469227013b00fee"}
	cmd.Flags().String("reason", "", "Why the evidence is refreshed (printed in the output)")
	// /sedum:addFlag:string
	// sedum:anchor:subcommands
	return cmd
}
