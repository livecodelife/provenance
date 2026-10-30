package main

import "github.com/spf13/cobra"

func newReconcileCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"v":2,"tier":"owned","record":"prov-2026-9a3aad45","gen":"sha256:d19012536182a12da7b036a73f8eea6a7a589bd985a0ee0352375454ff233111","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","id":{"bin":"provenance","command":"reconcile"},"kwargs":"sha256:bbbcb9b5a0151aed1c440bf0ad2f93b2409f0d39750aa725b1d15bec7403ac4f","body":"sha256:27b057e983fa75f5f9ebb4f4e60601063bd5636616cb7d62520f5770a0d7bb7f"}
	cmd.Use = "reconcile"
	cmd.Short = "Re-derive file write permissions from the open records' scopes"
	// /sedum:describeCommand
	// sedum:setRun {"v":2,"tier":"owned","record":"prov-2026-9a3aad45","gen":"sha256:d19012536182a12da7b036a73f8eea6a7a589bd985a0ee0352375454ff233111","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","id":{"bin":"provenance","command":"reconcile"},"kwargs":"sha256:781b147f4c3bc78c9c38c1fba0900748ac441ae3e3e2f9aad20bd6fb838a6d08","body":"sha256:354e17df8dd78422cb0e91be5744d474519626b8447d65b554774e0547f2798c"}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runReconcile(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:none {"v":2,"tier":"owned","record":"prov-2026-9a3aad45","gen":"sha256:d19012536182a12da7b036a73f8eea6a7a589bd985a0ee0352375454ff233111","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","id":{"bin":"provenance","command":"reconcile"},"kwargs":"sha256:880b86a7b73baa70f7f572e65e35aae8d8502718decc6649e21ec9c403960869","body":"sha256:778b1e3eb6004fe613ecb8793efaf401a29f1afdc382707cea55c05601aefd79"}
	cmd.Args = cobra.NoArgs
	// /sedum:setArgs:none
	// sedum:anchor:flags
	// sedum:addFlag:bool {"v":2,"tier":"owned","record":"prov-2026-9a3aad45","gen":"sha256:d19012536182a12da7b036a73f8eea6a7a589bd985a0ee0352375454ff233111","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"reconcile","flag":"dry-run"},"kwargs":"sha256:9e910962e7a5b0d4ed526bdd71fe0a11ecf002d5844a46116957b80424050efe","body":"sha256:d6777ef130a0db69d95846cf343bc0f17c83f8cb162b920d3572f2638c6cd555"}
	cmd.Flags().Bool("dry-run", false, "Report what would change without touching permissions")
	// /sedum:addFlag:bool
	// sedum:anchor:subcommands
	return cmd
}
