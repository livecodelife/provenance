package main

import "github.com/spf13/cobra"

func newCloneCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"v":2,"tier":"owned","record":"prov-2026-75646f74","gen":"sha256:335539008ab4bfd313eddaf0df95cf3ef61bd753a11aa929a86ff10195eee983","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","id":{"bin":"provenance","command":"clone"},"kwargs":"sha256:9d99210af1eadfde02a6fdb27fd7b19b6facc16eb7e80ae2a8d700d4ff8d6e8b","body":"sha256:8f212f3bf158a4d1c2fee875421090ad687b7691cd696c9dbcc4e4110c011281"}
	cmd.Use = "clone DIR [id...]"
	cmd.Short = "Create a local project from your published records"
	// /sedum:describeCommand
	// sedum:setRun {"v":2,"tier":"owned","record":"prov-2026-75646f74","gen":"sha256:335539008ab4bfd313eddaf0df95cf3ef61bd753a11aa929a86ff10195eee983","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","id":{"bin":"provenance","command":"clone"},"kwargs":"sha256:50442a1d4ef2709ff51ce9d849fc939851bb4354df1e4af8f4c156a8d2eaec3d","body":"sha256:ade336720a35121379cafce6d900605e48728bf313e8f4e60f3fb59fd5439c21"}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runClone(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:min_one {"v":2,"tier":"owned","record":"prov-2026-75646f74","gen":"sha256:335539008ab4bfd313eddaf0df95cf3ef61bd753a11aa929a86ff10195eee983","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","id":{"bin":"provenance","command":"clone"},"kwargs":"sha256:ebf04f21af83cdfdb1b1eb73b1900cb7dd3057ba32c18e049639c3506a2cb0aa","body":"sha256:c6da1a55784cf376fe47cadbebdd6c6bef498da9a05f6d41a5c642a4c8e8f325"}
	cmd.Args = cobra.MinimumNArgs(1)
	// /sedum:setArgs:min_one
	// sedum:anchor:flags
	// sedum:anchor:subcommands
	return cmd
}
