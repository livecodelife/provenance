package main

import "github.com/spf13/cobra"

func newInstallCmd() *cobra.Command {
	cmd := &cobra.Command{RunE: runInstall}
	// sedum:anchor:command
	// sedum:describeCommand {"tier":"owned","record":"prov-2026-0000c0b2","gen":"sha256:f12b8a04384dadc992a27af33c513f1a34362ccf9bc6a3e87192bbb0c9d9f508","pkg":"sha256:9958f264ddf8e57ee6107f5f5aacb530fb52d42dbf161cd7f31d050f90486529","kwargs":{"bin":"demo","command":"install","short":"Install git hooks","use":"install [dir]"}}
	cmd.Use = "install [dir]"
	cmd.Short = "Install git hooks"
	// /sedum:describeCommand
	// sedum:setArgs:max_one {"tier":"owned","record":"prov-2026-0000c0b2","gen":"sha256:f12b8a04384dadc992a27af33c513f1a34362ccf9bc6a3e87192bbb0c9d9f508","pkg":"sha256:84522566c0a6296808b5e0750805ad85951fe40c0b9efaa6bbb8bfff853d4eb7","kwargs":{"arity":"max_one","bin":"demo","command":"install"}}
	cmd.Args = cobra.MaximumNArgs(1)
	// /sedum:setArgs:max_one
	// sedum:anchor:flags
	// sedum:addFlag:bool {"tier":"owned","record":"prov-2026-0000c0b2","gen":"sha256:f12b8a04384dadc992a27af33c513f1a34362ccf9bc6a3e87192bbb0c9d9f508","pkg":"sha256:8f311eea74643a2f16c68f4e5b8853bc4f2710370d7971247c944d75c33b139c","kwargs":{"bin":"demo","command":"install","default":"","flag":"force","shorthand":"f","type":"bool","usage":"Overwrite existing hooks"}}
	cmd.Flags().BoolP("force", "f", false, "Overwrite existing hooks")
	// /sedum:addFlag:bool
	// sedum:anchor:subcommands
	return cmd
}
