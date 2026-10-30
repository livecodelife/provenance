package main

import "github.com/spf13/cobra"

func newTokenCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"v":2,"tier":"owned","record":"prov-2026-7f96263a","gen":"sha256:d80ee2d48ee6cd63e3793ad5ad19e755f8443d038715bdaab730d623ed577073","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","id":{"bin":"provenance","command":"token"},"kwargs":"sha256:e339b6acaa95ecb51cea0db3ff61dca478ed5baeddca45bada21252775d89982","body":"sha256:f3dcc0888a9709635ce1d7c75a2e2f3fd9702d529210241f7efbdc24af9fe762"}
	cmd.Use = "token"
	cmd.Short = "Manage Provisio personal tokens"
	// /sedum:describeCommand
	// sedum:setArgs:none {"v":2,"tier":"owned","record":"prov-2026-7f96263a","gen":"sha256:d80ee2d48ee6cd63e3793ad5ad19e755f8443d038715bdaab730d623ed577073","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","id":{"bin":"provenance","command":"token"},"kwargs":"sha256:13526d477b310d0560dd87dd46e7ad2637745bf85a8a04cf6b8b6a9e0e3855ae","body":"sha256:778b1e3eb6004fe613ecb8793efaf401a29f1afdc382707cea55c05601aefd79"}
	cmd.Args = cobra.NoArgs
	// /sedum:setArgs:none
	// sedum:anchor:flags
	// sedum:anchor:subcommands
	// sedum:registerSubcommand {"v":2,"tier":"owned","record":"prov-2026-a8ec3e30","gen":"sha256:32397270f026f7499dfffc6c69c80d69185c2e1390ca78f0ee9ecba7837ebbdd","pkg":"sha256:8a9e0a8ee811e7c629264e26f09a626602890b11fd82710cf10ae73e3da1cf7a","id":{"bin":"provenance","command":"token-revoke","parent":"token"},"kwargs":"sha256:ea1d559a2e009d757de0763fa97ffb3c9011e0eb9d4bac68ed1ea158396009b4","body":"sha256:b090b356e61ed84c55192180108ee83b5cc05d0a588f12ead831c33972c8aae8"}
	cmd.AddCommand(newTokenRevokeCmd())
	// /sedum:registerSubcommand
	// sedum:registerSubcommand {"v":2,"tier":"owned","record":"prov-2026-b071f42a","gen":"sha256:07b2170a3890a057d5f8c16a5193181369591558ff57477d9219a65601fd0761","pkg":"sha256:8a9e0a8ee811e7c629264e26f09a626602890b11fd82710cf10ae73e3da1cf7a","id":{"bin":"provenance","command":"token-list","parent":"token"},"kwargs":"sha256:bef10a6dc9d4aa717923697f6014960ec4c6ea3a458f589fd7e1ce24014ada09","body":"sha256:006a51788cf8c0820b4ccb21b9371d94c598e1722c745872a76addd7f978df9b"}
	cmd.AddCommand(newTokenListCmd())
	// /sedum:registerSubcommand
	// sedum:registerSubcommand {"v":2,"tier":"owned","record":"prov-2026-f1bc0a28","gen":"sha256:ec8ea3f85631b4d9593ab92af954f0e8b7a92678909682ad65a2d932ebd94f0a","pkg":"sha256:8a9e0a8ee811e7c629264e26f09a626602890b11fd82710cf10ae73e3da1cf7a","id":{"bin":"provenance","command":"token-create","parent":"token"},"kwargs":"sha256:a98e1e446e232cf0bb4f7b3e6ce4f127da90419058bd7057998e65cf77e00cdc","body":"sha256:a4d89aa702a8d1b2af36a80d66ebd0e9f2f1905abfeefaf3ac4ced3947499f67"}
	cmd.AddCommand(newTokenCreateCmd())
	// /sedum:registerSubcommand
	return cmd
}
