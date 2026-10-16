package main

import "github.com/spf13/cobra"

func newEvidenceRefreshCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"tier":"owned","record":"prov-2026-15f42a00","gen":"sha256:c80bc449c18db599e38113fc2c49132fd1fd245096f19d2387e80dc6dcd91188","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","kwargs":{"bin":"provenance","command":"evidence-refresh","short":"Append fresh evidence to an implemented record without a new version","use":"refresh [id]"}}
	cmd.Use = "refresh [id]"
	cmd.Short = "Append fresh evidence to an implemented record without a new version"
	// /sedum:describeCommand
	// sedum:setRun {"tier":"owned","record":"prov-2026-15f42a00","gen":"sha256:c80bc449c18db599e38113fc2c49132fd1fd245096f19d2387e80dc6dcd91188","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","kwargs":{"bin":"provenance","command":"evidence-refresh"}}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runEvidenceRefresh(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:max_one {"tier":"owned","record":"prov-2026-15f42a00","gen":"sha256:c80bc449c18db599e38113fc2c49132fd1fd245096f19d2387e80dc6dcd91188","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","kwargs":{"arity":"max_one","bin":"provenance","command":"evidence-refresh"}}
	cmd.Args = cobra.MaximumNArgs(1)
	// /sedum:setArgs:max_one
	// sedum:anchor:flags
	// sedum:addFlag:string {"tier":"owned","record":"prov-2026-15f42a00","gen":"sha256:c80bc449c18db599e38113fc2c49132fd1fd245096f19d2387e80dc6dcd91188","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"evidence-refresh","flag":"record","type":"string","usage":"Record to refresh (or give it as the argument)"}}
	cmd.Flags().String("record", "", "Record to refresh (or give it as the argument)")
	// /sedum:addFlag:string
	// sedum:addFlag:string {"tier":"owned","record":"prov-2026-15f42a00","gen":"sha256:c80bc449c18db599e38113fc2c49132fd1fd245096f19d2387e80dc6dcd91188","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"evidence-refresh","flag":"reason","type":"string","usage":"Why the evidence is refreshed (printed in the output)"}}
	cmd.Flags().String("reason", "", "Why the evidence is refreshed (printed in the output)")
	// /sedum:addFlag:string
	// sedum:addFlag:bool {"tier":"owned","record":"prov-2026-15f42a00","gen":"sha256:c80bc449c18db599e38113fc2c49132fd1fd245096f19d2387e80dc6dcd91188","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"evidence-refresh","flag":"all","type":"bool","usage":"Refresh every implemented record that is stale only because its code moved"}}
	cmd.Flags().Bool("all", false, "Refresh every implemented record that is stale only because its code moved")
	// /sedum:addFlag:bool
	// sedum:anchor:subcommands
	return cmd
}
