package main

import "github.com/spf13/cobra"

func newHooksInstallCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"tier":"owned","record":"prov-2026-0000c0b2","gen":"sha256:8b351f5ad014f88b418f9ba4898111b02258f9432172a203a8d2e033607bb469","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","kwargs":{"bin":"demo","command":"hooks-install","short":"Install git hooks","use":"install [dir]"}}
	cmd.Use = "install [dir]"
	cmd.Short = "Install git hooks"
	// /sedum:describeCommand
	// sedum:setRun {"tier":"owned","record":"prov-2026-0000c0b2","gen":"sha256:8b351f5ad014f88b418f9ba4898111b02258f9432172a203a8d2e033607bb469","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","kwargs":{"bin":"demo","command":"hooks-install"}}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runHooksInstall(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:max_one {"tier":"owned","record":"prov-2026-0000c0b2","gen":"sha256:8b351f5ad014f88b418f9ba4898111b02258f9432172a203a8d2e033607bb469","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","kwargs":{"arity":"max_one","bin":"demo","command":"hooks-install"}}
	cmd.Args = cobra.MaximumNArgs(1)
	// /sedum:setArgs:max_one
	// sedum:anchor:flags
	// sedum:addShortFlag:bool {"tier":"owned","record":"prov-2026-0000c0b2","gen":"sha256:8b351f5ad014f88b418f9ba4898111b02258f9432172a203a8d2e033607bb469","pkg":"sha256:317b97244b0fe957a6e0b80cdfa819a8afaae50b42b10385c9cd6729cb74fd3a","kwargs":{"bin":"demo","command":"hooks-install","flag":"force","shorthand":"f","type":"bool","usage":"Overwrite existing hooks"}}
	cmd.Flags().BoolP("force", "f", false, "Overwrite existing hooks")
	// /sedum:addShortFlag:bool
	// sedum:anchor:subcommands
	return cmd
}
