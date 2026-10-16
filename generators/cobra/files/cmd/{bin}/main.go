package main

import (
	"os"

	"github.com/spf13/cobra"
)

func newRootCmd() *cobra.Command {
	root := &cobra.Command{SilenceUsage: true}
	// sedum:anchor:root
	// sedum:anchor:persistent_flags
	// sedum:anchor:commands
	return root
}

func main() {
	if err := newRootCmd().Execute(); err != nil {
		os.Exit(exitCode(err))
	}
}
