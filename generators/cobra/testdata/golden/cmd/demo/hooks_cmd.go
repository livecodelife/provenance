package main

import "github.com/spf13/cobra"

func newHooksCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"tier":"owned","record":"prov-2026-0000c0b2","gen":"sha256:8b351f5ad014f88b418f9ba4898111b02258f9432172a203a8d2e033607bb469","pkg":"sha256:f8d152763fd9ce930ecfcc140104d898a335e71fb55263e286a4d13a61752868","kwargs":{"bin":"demo","command":"hooks","short":"Manage git hooks","use":"hooks"}}
	cmd.Use = "hooks"
	cmd.Short = "Manage git hooks"
	// /sedum:describeCommand
	// sedum:setArgs:none {"tier":"owned","record":"prov-2026-0000c0b2","gen":"sha256:8b351f5ad014f88b418f9ba4898111b02258f9432172a203a8d2e033607bb469","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","kwargs":{"arity":"none","bin":"demo","command":"hooks"}}
	cmd.Args = cobra.NoArgs
	// /sedum:setArgs:none
	// sedum:anchor:flags
	// sedum:anchor:subcommands
	// sedum:registerSubcommand {"tier":"owned","record":"prov-2026-0000c0b2","gen":"sha256:8b351f5ad014f88b418f9ba4898111b02258f9432172a203a8d2e033607bb469","pkg":"sha256:8a9e0a8ee811e7c629264e26f09a626602890b11fd82710cf10ae73e3da1cf7a","kwargs":{"bin":"demo","command":"hooks-install","parent":"hooks"}}
	cmd.AddCommand(newHooksInstallCmd())
	// /sedum:registerSubcommand
	return cmd
}
