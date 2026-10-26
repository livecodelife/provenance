package main

import "github.com/spf13/cobra"

func newDocumentCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"tier":"owned","record":"prov-2026-36b81f10","gen":"sha256:da0a6a84bf97ac25549e05b29429583669ff3d83fc5d84d6e6f3383a7d3b06f8","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","kwargs":{"bin":"provenance","command":"document","short":"Print the behaviour document of the briefs and blueprints, or apply an edited one","use":"document"}}
	cmd.Use = "document"
	cmd.Short = "Print the behaviour document of the briefs and blueprints, or apply an edited one"
	// /sedum:describeCommand
	// sedum:setRun {"tier":"owned","record":"prov-2026-36b81f10","gen":"sha256:da0a6a84bf97ac25549e05b29429583669ff3d83fc5d84d6e6f3383a7d3b06f8","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","kwargs":{"bin":"provenance","command":"document"}}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runDocument(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:none {"tier":"owned","record":"prov-2026-36b81f10","gen":"sha256:da0a6a84bf97ac25549e05b29429583669ff3d83fc5d84d6e6f3383a7d3b06f8","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","kwargs":{"arity":"none","bin":"provenance","command":"document"}}
	cmd.Args = cobra.NoArgs
	// /sedum:setArgs:none
	// sedum:anchor:flags
	// sedum:addFlag:strings {"tier":"owned","record":"prov-2026-36b81f10","gen":"sha256:da0a6a84bf97ac25549e05b29429583669ff3d83fc5d84d6e6f3383a7d3b06f8","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"document","flag":"scope","type":"strings","usage":"Only records whose affected_scope covers this path (repeatable)"}}
	cmd.Flags().StringArray("scope", nil, "Only records whose affected_scope covers this path (repeatable)")
	// /sedum:addFlag:strings
	// sedum:addDefaultFlag {"tier":"owned","record":"prov-2026-36b81f10","gen":"sha256:da0a6a84bf97ac25549e05b29429583669ff3d83fc5d84d6e6f3383a7d3b06f8","pkg":"sha256:f3c83dc533de1ffaec737cb62c02055fa2b4fb8d87929314763f7ddb9ff23364","kwargs":{"bin":"provenance","command":"document","default":"md","flag":"format","usage":"Output format (only md)"}}
	cmd.Flags().String("format", "md", "Output format (only md)")
	// /sedum:addDefaultFlag
	// sedum:addShortFlag:string {"tier":"owned","record":"prov-2026-36b81f10","gen":"sha256:da0a6a84bf97ac25549e05b29429583669ff3d83fc5d84d6e6f3383a7d3b06f8","pkg":"sha256:317b97244b0fe957a6e0b80cdfa819a8afaae50b42b10385c9cd6729cb74fd3a","kwargs":{"bin":"provenance","command":"document","flag":"out","shorthand":"o","type":"string","usage":"Write the document to this file instead of standard output"}}
	cmd.Flags().StringP("out", "o", "", "Write the document to this file instead of standard output")
	// /sedum:addShortFlag:string
	// sedum:addFlag:bool {"tier":"owned","record":"prov-2026-36b81f10","gen":"sha256:da0a6a84bf97ac25549e05b29429583669ff3d83fc5d84d6e6f3383a7d3b06f8","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"document","flag":"ids","type":"bool","usage":"Show record ids in the headings"}}
	cmd.Flags().Bool("ids", false, "Show record ids in the headings")
	// /sedum:addFlag:bool
	// sedum:addFlag:string {"tier":"owned","record":"prov-2026-36b81f10","gen":"sha256:da0a6a84bf97ac25549e05b29429583669ff3d83fc5d84d6e6f3383a7d3b06f8","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"document","flag":"apply","type":"string","usage":"Plan the changes made in this edited document file"}}
	cmd.Flags().String("apply", "", "Plan the changes made in this edited document file")
	// /sedum:addFlag:string
	// sedum:addFlag:bool {"tier":"owned","record":"prov-2026-36b81f10","gen":"sha256:da0a6a84bf97ac25549e05b29429583669ff3d83fc5d84d6e6f3383a7d3b06f8","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"document","flag":"write","type":"bool","usage":"Apply the plan to the records (needs --apply and --reason)"}}
	cmd.Flags().Bool("write", false, "Apply the plan to the records (needs --apply and --reason)")
	// /sedum:addFlag:bool
	// sedum:addFlag:string {"tier":"owned","record":"prov-2026-36b81f10","gen":"sha256:da0a6a84bf97ac25549e05b29429583669ff3d83fc5d84d6e6f3383a7d3b06f8","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"document","flag":"reason","type":"string","usage":"Why these revisions exist (required with --write)"}}
	cmd.Flags().String("reason", "", "Why these revisions exist (required with --write)")
	// /sedum:addFlag:string
	// sedum:addFlag:bool {"tier":"owned","record":"prov-2026-36b81f10","gen":"sha256:da0a6a84bf97ac25549e05b29429583669ff3d83fc5d84d6e6f3383a7d3b06f8","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"document","flag":"skip-conflicts","type":"bool","usage":"With --write, apply the other changes and skip conflicting sections"}}
	cmd.Flags().Bool("skip-conflicts", false, "With --write, apply the other changes and skip conflicting sections")
	// /sedum:addFlag:bool
	// sedum:anchor:subcommands
	return cmd
}
