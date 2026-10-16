package main

import "github.com/spf13/cobra"

func newOpenCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"tier":"owned","record":"prov-2026-6db88b75","gen":"sha256:cf831631b6137cd400fe64495057052c30d745b64a30dbfcb2ef52004d5e57e8","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","kwargs":{"bin":"provenance","command":"open","short":"Transition a record from draft to open (enables enforcement)","use":"open"}}
	cmd.Use = "open"
	cmd.Short = "Transition a record from draft to open (enables enforcement)"
	// /sedum:describeCommand
	// sedum:setRun {"tier":"owned","record":"prov-2026-6db88b75","gen":"sha256:cf831631b6137cd400fe64495057052c30d745b64a30dbfcb2ef52004d5e57e8","pkg":"sha256:fadae715925d582c263ce084f15c491995be26d77f0489abc836c658e094878f","kwargs":{"bin":"provenance","command":"open"}}
	cmd.RunE = runOpen
	// /sedum:setRun
	// sedum:setArgs:none {"tier":"owned","record":"prov-2026-6db88b75","gen":"sha256:cf831631b6137cd400fe64495057052c30d745b64a30dbfcb2ef52004d5e57e8","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","kwargs":{"arity":"none","bin":"provenance","command":"open"}}
	cmd.Args = cobra.NoArgs
	// /sedum:setArgs:none
	// sedum:anchor:flags
	// sedum:addFlag:string {"tier":"owned","record":"prov-2026-6db88b75","gen":"sha256:cf831631b6137cd400fe64495057052c30d745b64a30dbfcb2ef52004d5e57e8","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"open","flag":"record","type":"string","usage":"Record to transition from draft to open"}}
	cmd.Flags().String("record", "", "Record to transition from draft to open")
	// /sedum:addFlag:string
	// sedum:anchor:subcommands
	return cmd
}
