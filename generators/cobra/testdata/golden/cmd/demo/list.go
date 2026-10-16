package main

import "github.com/spf13/cobra"

func newListCmd() *cobra.Command {
	cmd := &cobra.Command{RunE: runList}
	// sedum:anchor:command
	// sedum:describeCommand {"tier":"owned","record":"prov-2026-0000c0b1","gen":"sha256:548e4aad36d237333a5e40aeca4bb3e36cff37295c0814441f79016f399ecb56","pkg":"sha256:9958f264ddf8e57ee6107f5f5aacb530fb52d42dbf161cd7f31d050f90486529","kwargs":{"bin":"demo","command":"list","short":"List items","use":"list [items]"}}
	cmd.Use = "list [items]"
	cmd.Short = "List items"
	// /sedum:describeCommand
	// sedum:setArgs:any {"tier":"owned","record":"prov-2026-0000c0b1","gen":"sha256:548e4aad36d237333a5e40aeca4bb3e36cff37295c0814441f79016f399ecb56","pkg":"sha256:84522566c0a6296808b5e0750805ad85951fe40c0b9efaa6bbb8bfff853d4eb7","kwargs":{"arity":"any","bin":"demo","command":"list"}}
	cmd.Args = cobra.ArbitraryArgs
	// /sedum:setArgs:any
	// sedum:anchor:flags
	// sedum:anchor:subcommands
	return cmd
}
