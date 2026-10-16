package main

import "github.com/spf13/cobra"

func newReconcileCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"tier":"owned","record":"prov-2026-9a3aad45","gen":"sha256:d19012536182a12da7b036a73f8eea6a7a589bd985a0ee0352375454ff233111","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","kwargs":{"bin":"provenance","command":"reconcile","short":"Re-derive file write permissions from the open records' scopes","use":"reconcile"}}
	cmd.Use = "reconcile"
	cmd.Short = "Re-derive file write permissions from the open records' scopes"
	// /sedum:describeCommand
	// sedum:setRun {"tier":"owned","record":"prov-2026-9a3aad45","gen":"sha256:d19012536182a12da7b036a73f8eea6a7a589bd985a0ee0352375454ff233111","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","kwargs":{"bin":"provenance","command":"reconcile"}}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runReconcile(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:none {"tier":"owned","record":"prov-2026-9a3aad45","gen":"sha256:d19012536182a12da7b036a73f8eea6a7a589bd985a0ee0352375454ff233111","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","kwargs":{"arity":"none","bin":"provenance","command":"reconcile"}}
	cmd.Args = cobra.NoArgs
	// /sedum:setArgs:none
	// sedum:anchor:flags
	// sedum:addFlag:bool {"tier":"owned","record":"prov-2026-9a3aad45","gen":"sha256:d19012536182a12da7b036a73f8eea6a7a589bd985a0ee0352375454ff233111","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"reconcile","flag":"dry-run","type":"bool","usage":"Report what would change without touching permissions"}}
	cmd.Flags().Bool("dry-run", false, "Report what would change without touching permissions")
	// /sedum:addFlag:bool
	// sedum:anchor:subcommands
	return cmd
}
