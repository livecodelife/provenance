package main

import "github.com/spf13/cobra"

func newEvidenceCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"tier":"owned","record":"prov-2026-15f42a00","gen":"sha256:c80bc449c18db599e38113fc2c49132fd1fd245096f19d2387e80dc6dcd91188","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","kwargs":{"bin":"provenance","command":"evidence","short":"Work with a record's verification evidence","use":"evidence"}}
	cmd.Use = "evidence"
	cmd.Short = "Work with a record's verification evidence"
	// /sedum:describeCommand
	// sedum:setArgs:none {"tier":"owned","record":"prov-2026-15f42a00","gen":"sha256:c80bc449c18db599e38113fc2c49132fd1fd245096f19d2387e80dc6dcd91188","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","kwargs":{"arity":"none","bin":"provenance","command":"evidence"}}
	cmd.Args = cobra.NoArgs
	// /sedum:setArgs:none
	// sedum:anchor:flags
	// sedum:anchor:subcommands
	// sedum:registerSubcommand {"tier":"owned","record":"prov-2026-15f42a00","gen":"sha256:c80bc449c18db599e38113fc2c49132fd1fd245096f19d2387e80dc6dcd91188","pkg":"sha256:8a9e0a8ee811e7c629264e26f09a626602890b11fd82710cf10ae73e3da1cf7a","kwargs":{"bin":"provenance","command":"evidence-refresh","parent":"evidence"}}
	cmd.AddCommand(newEvidenceRefreshCmd())
	// /sedum:registerSubcommand
	return cmd
}
