package main

import "github.com/spf13/cobra"

func newCloneCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"tier":"owned","record":"prov-2026-75646f74","gen":"sha256:335539008ab4bfd313eddaf0df95cf3ef61bd753a11aa929a86ff10195eee983","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","kwargs":{"bin":"provenance","command":"clone","short":"Create a local project from your published records","use":"clone DIR [id...]"}}
	cmd.Use = "clone DIR [id...]"
	cmd.Short = "Create a local project from your published records"
	// /sedum:describeCommand
	// sedum:setRun {"tier":"owned","record":"prov-2026-75646f74","gen":"sha256:335539008ab4bfd313eddaf0df95cf3ef61bd753a11aa929a86ff10195eee983","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","kwargs":{"bin":"provenance","command":"clone"}}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runClone(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:min_one {"tier":"owned","record":"prov-2026-75646f74","gen":"sha256:335539008ab4bfd313eddaf0df95cf3ef61bd753a11aa929a86ff10195eee983","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","kwargs":{"arity":"min_one","bin":"provenance","command":"clone"}}
	cmd.Args = cobra.MinimumNArgs(1)
	// /sedum:setArgs:min_one
	// sedum:anchor:flags
	// sedum:anchor:subcommands
	return cmd
}
