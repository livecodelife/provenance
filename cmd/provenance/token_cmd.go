package main

import "github.com/spf13/cobra"

func newTokenCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"tier":"owned","record":"prov-2026-7f96263a","gen":"sha256:7d219ded22bdca52e8acfbc153485b2bb82cf7673bdc34a38acd2ef12afcf9dd","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","kwargs":{"bin":"provenance","command":"token","short":"Manage Provisio personal tokens","use":"token"}}
	cmd.Use = "token"
	cmd.Short = "Manage Provisio personal tokens"
	// /sedum:describeCommand
	// sedum:setArgs:none {"tier":"owned","record":"prov-2026-7f96263a","gen":"sha256:7d219ded22bdca52e8acfbc153485b2bb82cf7673bdc34a38acd2ef12afcf9dd","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","kwargs":{"arity":"none","bin":"provenance","command":"token"}}
	cmd.Args = cobra.NoArgs
	// /sedum:setArgs:none
	// sedum:anchor:flags
	// sedum:anchor:subcommands
	// sedum:registerSubcommand {"tier":"owned","record":"prov-2026-7f96263a","gen":"sha256:7d219ded22bdca52e8acfbc153485b2bb82cf7673bdc34a38acd2ef12afcf9dd","pkg":"sha256:8a9e0a8ee811e7c629264e26f09a626602890b11fd82710cf10ae73e3da1cf7a","kwargs":{"bin":"provenance","command":"token-create","parent":"token"}}
	cmd.AddCommand(newTokenCreateCmd())
	// /sedum:registerSubcommand
	// sedum:registerSubcommand {"tier":"owned","record":"prov-2026-7f96263a","gen":"sha256:7d219ded22bdca52e8acfbc153485b2bb82cf7673bdc34a38acd2ef12afcf9dd","pkg":"sha256:8a9e0a8ee811e7c629264e26f09a626602890b11fd82710cf10ae73e3da1cf7a","kwargs":{"bin":"provenance","command":"token-list","parent":"token"}}
	cmd.AddCommand(newTokenListCmd())
	// /sedum:registerSubcommand
	// sedum:registerSubcommand {"tier":"owned","record":"prov-2026-7f96263a","gen":"sha256:7d219ded22bdca52e8acfbc153485b2bb82cf7673bdc34a38acd2ef12afcf9dd","pkg":"sha256:8a9e0a8ee811e7c629264e26f09a626602890b11fd82710cf10ae73e3da1cf7a","kwargs":{"bin":"provenance","command":"token-revoke","parent":"token"}}
	cmd.AddCommand(newTokenRevokeCmd())
	// /sedum:registerSubcommand
	return cmd
}
