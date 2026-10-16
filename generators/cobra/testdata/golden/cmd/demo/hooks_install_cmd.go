package main

import "github.com/spf13/cobra"

func newHooksInstallCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"tier":"owned","record":"prov-2026-0000c0b2","gen":"sha256:8b351f5ad014f88b418f9ba4898111b02258f9432172a203a8d2e033607bb469","pkg":"sha256:f8d152763fd9ce930ecfcc140104d898a335e71fb55263e286a4d13a61752868","kwargs":{"bin":"demo","command":"hooks-install","short":"Install git hooks","use":"install [dir]"}}
	cmd.Use = "install [dir]"
	cmd.Short = "Install git hooks"
	// /sedum:describeCommand
	// sedum:setRun {"tier":"owned","record":"prov-2026-0000c0b2","gen":"sha256:8b351f5ad014f88b418f9ba4898111b02258f9432172a203a8d2e033607bb469","pkg":"sha256:fadae715925d582c263ce084f15c491995be26d77f0489abc836c658e094878f","kwargs":{"bin":"demo","command":"hooks-install"}}
	cmd.RunE = runHooksInstall
	// /sedum:setRun
	// sedum:setArgs:max_one {"tier":"owned","record":"prov-2026-0000c0b2","gen":"sha256:8b351f5ad014f88b418f9ba4898111b02258f9432172a203a8d2e033607bb469","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","kwargs":{"arity":"max_one","bin":"demo","command":"hooks-install"}}
	cmd.Args = cobra.MaximumNArgs(1)
	// /sedum:setArgs:max_one
	// sedum:anchor:flags
	// sedum:addShortFlag:bool {"tier":"owned","record":"prov-2026-0000c0b2","gen":"sha256:8b351f5ad014f88b418f9ba4898111b02258f9432172a203a8d2e033607bb469","pkg":"sha256:5f5c7969eebf532c5ae0b4d8c4ed650029df54cc12ae123e5b0f04e45761961b","kwargs":{"bin":"demo","command":"hooks-install","flag":"force","shorthand":"f","type":"bool","usage":"Overwrite existing hooks"}}
	cmd.Flags().BoolP("force", "f", false, "Overwrite existing hooks")
	// /sedum:addShortFlag:bool
	// sedum:anchor:subcommands
	return cmd
}
