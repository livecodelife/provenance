package main

import "github.com/spf13/cobra"

func newEchoCmd() *cobra.Command {
	cmd := &cobra.Command{RunE: runEcho}
	// sedum:anchor:command
	// sedum:describeCommand {"tier":"owned","record":"prov-2026-0000c0b1","gen":"sha256:548e4aad36d237333a5e40aeca4bb3e36cff37295c0814441f79016f399ecb56","pkg":"sha256:9958f264ddf8e57ee6107f5f5aacb530fb52d42dbf161cd7f31d050f90486529","kwargs":{"bin":"demo","command":"echo","short":"Echo words","use":"echo <word>..."}}
	cmd.Use = "echo <word>..."
	cmd.Short = "Echo words"
	// /sedum:describeCommand
	// sedum:setArgs:min_one {"tier":"owned","record":"prov-2026-0000c0b1","gen":"sha256:548e4aad36d237333a5e40aeca4bb3e36cff37295c0814441f79016f399ecb56","pkg":"sha256:84522566c0a6296808b5e0750805ad85951fe40c0b9efaa6bbb8bfff853d4eb7","kwargs":{"arity":"min_one","bin":"demo","command":"echo"}}
	cmd.Args = cobra.MinimumNArgs(1)
	// /sedum:setArgs:min_one
	// sedum:anchor:flags
	// sedum:anchor:subcommands
	return cmd
}
