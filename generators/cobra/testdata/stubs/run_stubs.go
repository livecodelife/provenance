package main

import (
	"io"

	"github.com/spf13/cobra"
)

func runGreet(out io.Writer, cmd *cobra.Command, args []string) error        { return nil }
func runEcho(out io.Writer, cmd *cobra.Command, args []string) error         { return nil }
func runList(out io.Writer, cmd *cobra.Command, args []string) error         { return nil }
func runShowID(out io.Writer, cmd *cobra.Command, args []string) error       { return nil }
func runHooksInstall(out io.Writer, cmd *cobra.Command, args []string) error { return nil }
