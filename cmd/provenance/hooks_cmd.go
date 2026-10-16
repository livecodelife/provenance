package main

import "github.com/spf13/cobra"

func newHooksCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"tier":"owned","record":"prov-2026-d1d789de","gen":"sha256:bb33e281db3fef59d1d98f44c667bf0ad0797c86378e45488297adfe0b0afb14","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","kwargs":{"bin":"provenance","command":"hooks","short":"Manage the provenance git hooks","use":"hooks"}}
	cmd.Use = "hooks"
	cmd.Short = "Manage the provenance git hooks"
	// /sedum:describeCommand
	// sedum:setArgs:none {"tier":"owned","record":"prov-2026-d1d789de","gen":"sha256:bb33e281db3fef59d1d98f44c667bf0ad0797c86378e45488297adfe0b0afb14","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","kwargs":{"arity":"none","bin":"provenance","command":"hooks"}}
	cmd.Args = cobra.NoArgs
	// /sedum:setArgs:none
	// sedum:anchor:flags
	// sedum:anchor:subcommands
	// sedum:registerSubcommand {"tier":"owned","record":"prov-2026-d1d789de","gen":"sha256:bb33e281db3fef59d1d98f44c667bf0ad0797c86378e45488297adfe0b0afb14","pkg":"sha256:8a9e0a8ee811e7c629264e26f09a626602890b11fd82710cf10ae73e3da1cf7a","kwargs":{"bin":"provenance","command":"hooks-install","parent":"hooks"}}
	cmd.AddCommand(newHooksInstallCmd())
	// /sedum:registerSubcommand
	return cmd
}
