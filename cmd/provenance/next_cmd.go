package main

import "github.com/spf13/cobra"

func newNextCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"tier":"owned","record":"prov-2026-11876813","gen":"sha256:c325742da33a6c9e275df6fdfe823a0b9eccf202a038d29287ebcb298253db17","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","kwargs":{"bin":"provenance","command":"next","short":"Compute the single correct next provenance action","use":"next [files...]"}}
	cmd.Use = "next [files...]"
	cmd.Short = "Compute the single correct next provenance action"
	// /sedum:describeCommand
	// sedum:setRun {"tier":"owned","record":"prov-2026-11876813","gen":"sha256:c325742da33a6c9e275df6fdfe823a0b9eccf202a038d29287ebcb298253db17","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","kwargs":{"bin":"provenance","command":"next"}}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runNext(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:any {"tier":"owned","record":"prov-2026-11876813","gen":"sha256:c325742da33a6c9e275df6fdfe823a0b9eccf202a038d29287ebcb298253db17","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","kwargs":{"arity":"any","bin":"provenance","command":"next"}}
	cmd.Args = cobra.ArbitraryArgs
	// /sedum:setArgs:any
	// sedum:anchor:flags
	// sedum:addFlag:strings {"tier":"owned","record":"prov-2026-11876813","gen":"sha256:c325742da33a6c9e275df6fdfe823a0b9eccf202a038d29287ebcb298253db17","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"next","flag":"files","type":"strings","usage":"Files you intend to change (plan before editing)"}}
	cmd.Flags().StringArray("files", nil, "Files you intend to change (plan before editing)")
	// /sedum:addFlag:strings
	// sedum:addFlag:strings {"tier":"owned","record":"prov-2026-11876813","gen":"sha256:c325742da33a6c9e275df6fdfe823a0b9eccf202a038d29287ebcb298253db17","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"next","flag":"plan","type":"strings","usage":"Alias for --files"}}
	cmd.Flags().StringArray("plan", nil, "Alias for --files")
	// /sedum:addFlag:strings
	// sedum:addFlag:bool {"tier":"owned","record":"prov-2026-11876813","gen":"sha256:c325742da33a6c9e275df6fdfe823a0b9eccf202a038d29287ebcb298253db17","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"next","flag":"json","type":"bool","usage":"Machine-readable output (for hooks and agents)"}}
	cmd.Flags().Bool("json", false, "Machine-readable output (for hooks and agents)")
	// /sedum:addFlag:bool
	// sedum:addFlag:string {"tier":"owned","record":"prov-2026-11876813","gen":"sha256:c325742da33a6c9e275df6fdfe823a0b9eccf202a038d29287ebcb298253db17","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"next","flag":"format","type":"string","usage":"Output format (human|json)"}}
	cmd.Flags().String("format", "", "Output format (human|json)")
	// /sedum:addFlag:string
	// sedum:anchor:subcommands
	return cmd
}
