package main

import "github.com/spf13/cobra"

func newMigrateCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"v":2,"tier":"owned","record":"prov-2026-76a125e6","gen":"sha256:68d57b1b46d0b0deed02c7130fec8249449d9587617ea80f4e9ce204f34362a1","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","id":{"bin":"provenance","command":"migrate"},"kwargs":"sha256:17b08eeda82a8fbcbe0bb6158f2eff3152d8e1299095a56955b421f71da49c29","body":"sha256:372aab2c64839a4c97e058e1e67a52e9f88a61c4e2c2563d491965cc74af83e8"}
	cmd.Use = "migrate"
	cmd.Short = "Convert v1 records to versioned v2 records"
	// /sedum:describeCommand
	// sedum:setRun {"v":2,"tier":"owned","record":"prov-2026-76a125e6","gen":"sha256:68d57b1b46d0b0deed02c7130fec8249449d9587617ea80f4e9ce204f34362a1","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","id":{"bin":"provenance","command":"migrate"},"kwargs":"sha256:c4c99cb10c851126b8f1d4aa2af1ce7a664b1224a96a21975ba3f3649f2aa818","body":"sha256:8ec23118097ddb6b1b1a00e00327517aa06c453cb5a59189e3139fd0edabece6"}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runMigrate(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:none {"v":2,"tier":"owned","record":"prov-2026-76a125e6","gen":"sha256:68d57b1b46d0b0deed02c7130fec8249449d9587617ea80f4e9ce204f34362a1","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","id":{"bin":"provenance","command":"migrate"},"kwargs":"sha256:b4630c9642cbd947c84b6b8d11adfef451bd3cda84ac2ab8a873ca1a74dc6bfd","body":"sha256:778b1e3eb6004fe613ecb8793efaf401a29f1afdc382707cea55c05601aefd79"}
	cmd.Args = cobra.NoArgs
	// /sedum:setArgs:none
	// sedum:anchor:flags
	// sedum:addFlag:bool {"v":2,"tier":"owned","record":"prov-2026-76a125e6","gen":"sha256:68d57b1b46d0b0deed02c7130fec8249449d9587617ea80f4e9ce204f34362a1","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"migrate","flag":"plan"},"kwargs":"sha256:77fa477d37248e0a95dce4af901b723e7419e5417eb69950a7f36efd25261d7a","body":"sha256:3efba23a392a8f2fe8cbc6b84e5a6f21ad6c3486e08679ad794581efbba6f302"}
	cmd.Flags().Bool("plan", false, "Print the migration plan and write nothing")
	// /sedum:addFlag:bool
	// sedum:addFlag:bool {"v":2,"tier":"owned","record":"prov-2026-76a125e6","gen":"sha256:68d57b1b46d0b0deed02c7130fec8249449d9587617ea80f4e9ce204f34362a1","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"migrate","flag":"json"},"kwargs":"sha256:56f4d08572cdd9329ccaf7aba35061ae7eaf186cf990e0f439ad3cc6b5107c9d","body":"sha256:6d86730a855587624d40cd64bf4f9d688062e00b5d4c39f20106ca224974a92d"}
	cmd.Flags().Bool("json", false, "Print the plan as JSON (with --plan)")
	// /sedum:addFlag:bool
	// sedum:addFlag:bool {"v":2,"tier":"owned","record":"prov-2026-76a125e6","gen":"sha256:68d57b1b46d0b0deed02c7130fec8249449d9587617ea80f4e9ce204f34362a1","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"migrate","flag":"accept-suggestions"},"kwargs":"sha256:35259fae1172f8e38817b87e4228fd9893704a6b590142d0aa198115add7a67a","body":"sha256:23fd371f195af82596fd58b84b891cf653c8a1b4c505ec1f00dde7fe5f8c4dce"}
	cmd.Flags().Bool("accept-suggestions", false, "Keep every suggested record separate")
	// /sedum:addFlag:bool
	// sedum:addFlag:strings {"v":2,"tier":"owned","record":"prov-2026-76a125e6","gen":"sha256:68d57b1b46d0b0deed02c7130fec8249449d9587617ea80f4e9ce204f34362a1","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"migrate","flag":"keep-separate"},"kwargs":"sha256:145490c3ba30d68eab1000668200697d7daeb558039ecda52102c333270db72c","body":"sha256:e628441c7cd4c94532c167194c2a30ac3c949ade67a18f705816251f994ab5a1"}
	cmd.Flags().StringArray("keep-separate", nil, "Keep this absorbed record separate (repeatable)")
	// /sedum:addFlag:strings
	// sedum:addFlag:strings {"v":2,"tier":"owned","record":"prov-2026-76a125e6","gen":"sha256:68d57b1b46d0b0deed02c7130fec8249449d9587617ea80f4e9ce204f34362a1","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"migrate","flag":"fold"},"kwargs":"sha256:34cefe980016fef6050e18ecdb62be8776c13357f3401fa7c369354e64d0b4f6","body":"sha256:84489e81275b2564e77b8f181b9e7050e60e7ff6b09b69502a5f9eeacea18255"}
	cmd.Flags().StringArray("fold", nil, "Fold this record even when suggested to keep it separate (repeatable)")
	// /sedum:addFlag:strings
	// sedum:addFlag:string {"v":2,"tier":"owned","record":"prov-2026-76a125e6","gen":"sha256:68d57b1b46d0b0deed02c7130fec8249449d9587617ea80f4e9ce204f34362a1","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"migrate","flag":"workspace"},"kwargs":"sha256:8ef1cb8cd8d6e9657ad35e8bfc1f2ebe6c94e5abb2cf16ce8c38946247ae0510","body":"sha256:b33ecc238e665cb3589ca4f94a3f3ab0e7a7e1723efed8f8068855e4b9f6bc40"}
	cmd.Flags().String("workspace", "", "Directory holding local checkouts of shared repos (default: the repo's parent)")
	// /sedum:addFlag:string
	// sedum:anchor:subcommands
	return cmd
}
