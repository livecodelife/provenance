package main

import "github.com/spf13/cobra"

func new{{name|pascal}}Cmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:anchor:flags
	// sedum:anchor:subcommands
	return cmd
}
