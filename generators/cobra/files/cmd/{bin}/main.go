package main

import (
	"errors"
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

// exitCoder is implemented by an error that carries the process exit status.
type exitCoder interface{ ExitCode() int }

func main() {
	if err := newRootCmd().Execute(); err != nil {
		var coded exitCoder
		if errors.As(err, &coded) {
			os.Exit(coded.ExitCode())
		}
		os.Exit(1)
	}
}
