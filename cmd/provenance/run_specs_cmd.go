package main

import "github.com/spf13/cobra"

func newRunSpecsCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"v":2,"tier":"owned","record":"prov-2026-c31c9c59","gen":"sha256:9a18a4eb62cedc8b94c8ee92552042184ff66c2dcd6e7dd094c735af31f5de23","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","id":{"bin":"provenance","command":"run-specs"},"kwargs":"sha256:6bb988768c7eaa90e1b4aef2bd375f71ef15bfe7e97e67917e4ea3724cd9407a","body":"sha256:f9fa7685842bb3b2026e019f9f46577ad30cfd0d7d48355d4ecf962f5e85e384"}
	cmd.Use = "run-specs"
	cmd.Short = "Run a record's associated specs (used by the pre-commit hook)"
	// /sedum:describeCommand
	// sedum:setRun {"v":2,"tier":"owned","record":"prov-2026-c31c9c59","gen":"sha256:9a18a4eb62cedc8b94c8ee92552042184ff66c2dcd6e7dd094c735af31f5de23","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","id":{"bin":"provenance","command":"run-specs"},"kwargs":"sha256:97c669f9c235b598b8daac53c90a35a727a81ea2c1a96d2e907443fc656d1812","body":"sha256:89dc69bb2470082354a5500d69b4895c0a351a23f9c2da7d90358127f6c0f51e"}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runRunSpecs(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:none {"v":2,"tier":"owned","record":"prov-2026-c31c9c59","gen":"sha256:9a18a4eb62cedc8b94c8ee92552042184ff66c2dcd6e7dd094c735af31f5de23","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","id":{"bin":"provenance","command":"run-specs"},"kwargs":"sha256:51552a7bbe0caa31b2657209587c6664b70c944016f7e51f6ce862aa3f8681f9","body":"sha256:778b1e3eb6004fe613ecb8793efaf401a29f1afdc382707cea55c05601aefd79"}
	cmd.Args = cobra.NoArgs
	// /sedum:setArgs:none
	// sedum:anchor:flags
	// sedum:addFlag:string {"v":2,"tier":"owned","record":"prov-2026-c31c9c59","gen":"sha256:9a18a4eb62cedc8b94c8ee92552042184ff66c2dcd6e7dd094c735af31f5de23","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"run-specs","flag":"record"},"kwargs":"sha256:9f0205b49a86abddfa8958fa2a9d9578a21febf449909a343362263b5fddea95","body":"sha256:c4cd2965036bd46308614515d0f96371b7e4fd75574ba02714415f8cb24e09e1"}
	cmd.Flags().String("record", "", "Record whose associated specs to run")
	// /sedum:addFlag:string
	// sedum:anchor:subcommands
	return cmd
}
