package main

import "github.com/spf13/cobra"

func runGreet(cmd *cobra.Command, args []string) error   { return nil }
func runEcho(cmd *cobra.Command, args []string) error    { return nil }
func runList(cmd *cobra.Command, args []string) error    { return nil }
func runHooks(cmd *cobra.Command, args []string) error   { return cmd.Help() }
func runInstall(cmd *cobra.Command, args []string) error { return nil }
