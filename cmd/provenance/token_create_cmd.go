package main

import "github.com/spf13/cobra"

func newTokenCreateCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"tier":"owned","record":"prov-2026-f1bc0a28","gen":"sha256:ec8ea3f85631b4d9593ab92af954f0e8b7a92678909682ad65a2d932ebd94f0a","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","kwargs":{"bin":"provenance","command":"token-create","short":"Create a personal token (the secret is shown once)","use":"create"}}
	cmd.Use = "create"
	cmd.Short = "Create a personal token (the secret is shown once)"
	// /sedum:describeCommand
	// sedum:setRun {"tier":"owned","record":"prov-2026-f1bc0a28","gen":"sha256:ec8ea3f85631b4d9593ab92af954f0e8b7a92678909682ad65a2d932ebd94f0a","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","kwargs":{"bin":"provenance","command":"token-create"}}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runTokenCreate(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:none {"tier":"owned","record":"prov-2026-f1bc0a28","gen":"sha256:ec8ea3f85631b4d9593ab92af954f0e8b7a92678909682ad65a2d932ebd94f0a","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","kwargs":{"arity":"none","bin":"provenance","command":"token-create"}}
	cmd.Args = cobra.NoArgs
	// /sedum:setArgs:none
	// sedum:anchor:flags
	// sedum:addFlag:string {"tier":"owned","record":"prov-2026-f1bc0a28","gen":"sha256:ec8ea3f85631b4d9593ab92af954f0e8b7a92678909682ad65a2d932ebd94f0a","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"token-create","flag":"name","type":"string","usage":"Token name (default ci-<repository>)"}}
	cmd.Flags().String("name", "", "Token name (default ci-<repository>)")
	// /sedum:addFlag:string
	// sedum:addFlag:strings {"tier":"owned","record":"prov-2026-f1bc0a28","gen":"sha256:ec8ea3f85631b4d9593ab92af954f0e8b7a92678909682ad65a2d932ebd94f0a","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"token-create","flag":"scope","type":"strings","usage":"Scope: records:read or records:write (repeatable; default records:write)"}}
	cmd.Flags().StringArray("scope", nil, "Scope: records:read or records:write (repeatable; default records:write)")
	// /sedum:addFlag:strings
	// sedum:addFlag:int {"tier":"owned","record":"prov-2026-f1bc0a28","gen":"sha256:ec8ea3f85631b4d9593ab92af954f0e8b7a92678909682ad65a2d932ebd94f0a","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"token-create","flag":"expires-in-days","type":"int","usage":"Days until the token expires, 1 to 365 (default 90)"}}
	cmd.Flags().Int("expires-in-days", 0, "Days until the token expires, 1 to 365 (default 90)")
	// /sedum:addFlag:int
	// sedum:addFlag:bool {"tier":"owned","record":"prov-2026-f1bc0a28","gen":"sha256:ec8ea3f85631b4d9593ab92af954f0e8b7a92678909682ad65a2d932ebd94f0a","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"token-create","flag":"json","type":"bool","usage":"Print as JSON"}}
	cmd.Flags().Bool("json", false, "Print as JSON")
	// /sedum:addFlag:bool
	// sedum:addFlag:bool {"tier":"owned","record":"prov-2026-f1bc0a28","gen":"sha256:ec8ea3f85631b4d9593ab92af954f0e8b7a92678909682ad65a2d932ebd94f0a","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"token-create","flag":"secret-only","type":"bool","usage":"Print only the secret (for piping to gh secret set)"}}
	cmd.Flags().Bool("secret-only", false, "Print only the secret (for piping to gh secret set)")
	// /sedum:addFlag:bool
	// sedum:anchor:subcommands
	return cmd
}
