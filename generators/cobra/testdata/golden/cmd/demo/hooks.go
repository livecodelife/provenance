package main

import "github.com/spf13/cobra"

func newHooksCmd() *cobra.Command {
	cmd := &cobra.Command{RunE: runHooks}
	// sedum:anchor:command
	// sedum:describeCommand {"tier":"owned","record":"prov-2026-0000c0b2","gen":"sha256:f12b8a04384dadc992a27af33c513f1a34362ccf9bc6a3e87192bbb0c9d9f508","pkg":"sha256:9958f264ddf8e57ee6107f5f5aacb530fb52d42dbf161cd7f31d050f90486529","kwargs":{"bin":"demo","command":"hooks","short":"Manage git hooks","use":"hooks"}}
	cmd.Use = "hooks"
	cmd.Short = "Manage git hooks"
	// /sedum:describeCommand
	// sedum:setArgs:none {"tier":"owned","record":"prov-2026-0000c0b2","gen":"sha256:f12b8a04384dadc992a27af33c513f1a34362ccf9bc6a3e87192bbb0c9d9f508","pkg":"sha256:84522566c0a6296808b5e0750805ad85951fe40c0b9efaa6bbb8bfff853d4eb7","kwargs":{"arity":"none","bin":"demo","command":"hooks"}}
	cmd.Args = cobra.NoArgs
	// /sedum:setArgs:none
	// sedum:anchor:flags
	// sedum:anchor:subcommands
	// sedum:registerSubcommand {"tier":"owned","record":"prov-2026-0000c0b2","gen":"sha256:f12b8a04384dadc992a27af33c513f1a34362ccf9bc6a3e87192bbb0c9d9f508","pkg":"sha256:51768e25849e99146291ab75a000f04608bbef880dc56bcaa595bbc516a2ca3c","kwargs":{"bin":"demo","command":"install","parent":"hooks"}}
	cmd.AddCommand(newInstallCmd())
	// /sedum:registerSubcommand
	return cmd
}
