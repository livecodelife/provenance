package main

import "github.com/spf13/cobra"

func newTokenCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"tier":"owned","record":"prov-2026-7f96263a","gen":"sha256:d80ee2d48ee6cd63e3793ad5ad19e755f8443d038715bdaab730d623ed577073","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","kwargs":{"bin":"provenance","command":"token","short":"Manage Provisio personal tokens","use":"token"}}
	cmd.Use = "token"
	cmd.Short = "Manage Provisio personal tokens"
	// /sedum:describeCommand
	// sedum:setArgs:none {"tier":"owned","record":"prov-2026-7f96263a","gen":"sha256:d80ee2d48ee6cd63e3793ad5ad19e755f8443d038715bdaab730d623ed577073","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","kwargs":{"arity":"none","bin":"provenance","command":"token"}}
	cmd.Args = cobra.NoArgs
	// /sedum:setArgs:none
	// sedum:anchor:flags
	// sedum:anchor:subcommands
	// sedum:registerSubcommand {"tier":"owned","record":"prov-2026-a8ec3e30","gen":"sha256:32397270f026f7499dfffc6c69c80d69185c2e1390ca78f0ee9ecba7837ebbdd","pkg":"sha256:8a9e0a8ee811e7c629264e26f09a626602890b11fd82710cf10ae73e3da1cf7a","kwargs":{"bin":"provenance","command":"token-revoke","parent":"token"}}
	cmd.AddCommand(newTokenRevokeCmd())
	// /sedum:registerSubcommand
	// sedum:registerSubcommand {"tier":"owned","record":"prov-2026-b071f42a","gen":"sha256:07b2170a3890a057d5f8c16a5193181369591558ff57477d9219a65601fd0761","pkg":"sha256:8a9e0a8ee811e7c629264e26f09a626602890b11fd82710cf10ae73e3da1cf7a","kwargs":{"bin":"provenance","command":"token-list","parent":"token"}}
	cmd.AddCommand(newTokenListCmd())
	// /sedum:registerSubcommand
	// sedum:registerSubcommand {"tier":"owned","record":"prov-2026-f1bc0a28","gen":"sha256:ec8ea3f85631b4d9593ab92af954f0e8b7a92678909682ad65a2d932ebd94f0a","pkg":"sha256:8a9e0a8ee811e7c629264e26f09a626602890b11fd82710cf10ae73e3da1cf7a","kwargs":{"bin":"provenance","command":"token-create","parent":"token"}}
	cmd.AddCommand(newTokenCreateCmd())
	// /sedum:registerSubcommand
	return cmd
}
