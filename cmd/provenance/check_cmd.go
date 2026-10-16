package main

import "github.com/spf13/cobra"

func newCheckCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"tier":"owned","record":"prov-2026-063f1395","gen":"sha256:c998c2f19898100f6d7e0c2b705122f5b5e0ad542e32c7720ab172b5b901ce4e","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","kwargs":{"bin":"provenance","command":"check","short":"Check staged changes against provenance records (the commit-msg hook)","use":"check"}}
	cmd.Use = "check"
	cmd.Short = "Check staged changes against provenance records (the commit-msg hook)"
	// /sedum:describeCommand
	// sedum:setRun {"tier":"owned","record":"prov-2026-063f1395","gen":"sha256:c998c2f19898100f6d7e0c2b705122f5b5e0ad542e32c7720ab172b5b901ce4e","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","kwargs":{"bin":"provenance","command":"check"}}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runCheck(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:none {"tier":"owned","record":"prov-2026-063f1395","gen":"sha256:c998c2f19898100f6d7e0c2b705122f5b5e0ad542e32c7720ab172b5b901ce4e","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","kwargs":{"arity":"none","bin":"provenance","command":"check"}}
	cmd.Args = cobra.NoArgs
	// /sedum:setArgs:none
	// sedum:anchor:flags
	// sedum:addFlag:bool {"tier":"owned","record":"prov-2026-063f1395","gen":"sha256:c998c2f19898100f6d7e0c2b705122f5b5e0ad542e32c7720ab172b5b901ce4e","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"check","flag":"staged","type":"bool","usage":"Check the staged changes"}}
	cmd.Flags().Bool("staged", false, "Check the staged changes")
	// /sedum:addFlag:bool
	// sedum:addFlag:string {"tier":"owned","record":"prov-2026-063f1395","gen":"sha256:c998c2f19898100f6d7e0c2b705122f5b5e0ad542e32c7720ab172b5b901ce4e","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"check","flag":"message-file","type":"string","usage":"Commit message file (as the commit-msg hook passes it)"}}
	cmd.Flags().String("message-file", "", "Commit message file (as the commit-msg hook passes it)")
	// /sedum:addFlag:string
	// sedum:addFlag:string {"tier":"owned","record":"prov-2026-063f1395","gen":"sha256:c998c2f19898100f6d7e0c2b705122f5b5e0ad542e32c7720ab172b5b901ce4e","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"check","flag":"commit","type":"string","usage":"Check one commit"}}
	cmd.Flags().String("commit", "", "Check one commit")
	// /sedum:addFlag:string
	// sedum:addFlag:string {"tier":"owned","record":"prov-2026-063f1395","gen":"sha256:c998c2f19898100f6d7e0c2b705122f5b5e0ad542e32c7720ab172b5b901ce4e","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"check","flag":"range","type":"string","usage":"Check every commit in a..b"}}
	cmd.Flags().String("range", "", "Check every commit in a..b")
	// /sedum:addFlag:string
	// sedum:addFlag:bool {"tier":"owned","record":"prov-2026-063f1395","gen":"sha256:c998c2f19898100f6d7e0c2b705122f5b5e0ad542e32c7720ab172b5b901ce4e","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"check","flag":"evidence","type":"bool","usage":"Report whether each implemented record's evidence is fresh"}}
	cmd.Flags().Bool("evidence", false, "Report whether each implemented record's evidence is fresh")
	// /sedum:addFlag:bool
	// sedum:addFlag:bool {"tier":"owned","record":"prov-2026-063f1395","gen":"sha256:c998c2f19898100f6d7e0c2b705122f5b5e0ad542e32c7720ab172b5b901ce4e","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"check","flag":"json","type":"bool","usage":"Print --evidence as JSON"}}
	cmd.Flags().Bool("json", false, "Print --evidence as JSON")
	// /sedum:addFlag:bool
	// sedum:addFlag:string {"tier":"owned","record":"prov-2026-063f1395","gen":"sha256:c998c2f19898100f6d7e0c2b705122f5b5e0ad542e32c7720ab172b5b901ce4e","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"check","flag":"format","type":"string","usage":"Output format for violations and evidence (human|sarif)"}}
	cmd.Flags().String("format", "", "Output format for violations and evidence (human|sarif)")
	// /sedum:addFlag:string
	// sedum:anchor:subcommands
	return cmd
}
