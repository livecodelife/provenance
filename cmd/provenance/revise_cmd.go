package main

import "github.com/spf13/cobra"

func newReviseCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"tier":"owned","record":"prov-2026-5f69977b","gen":"sha256:cd83ad7d20c1e101894b6c1b65aa2fd28ee22c28745a2adbb5fab482eef7cc5e","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","kwargs":{"bin":"provenance","command":"revise","short":"Edit a record's content in $EDITOR and save it as a new version","use":"revise [id]"}}
	cmd.Use = "revise [id]"
	cmd.Short = "Edit a record's content in $EDITOR and save it as a new version"
	// /sedum:describeCommand
	// sedum:setRun {"tier":"owned","record":"prov-2026-5f69977b","gen":"sha256:cd83ad7d20c1e101894b6c1b65aa2fd28ee22c28745a2adbb5fab482eef7cc5e","pkg":"sha256:fadae715925d582c263ce084f15c491995be26d77f0489abc836c658e094878f","kwargs":{"bin":"provenance","command":"revise"}}
	cmd.RunE = runRevise
	// /sedum:setRun
	// sedum:setArgs:max_one {"tier":"owned","record":"prov-2026-5f69977b","gen":"sha256:cd83ad7d20c1e101894b6c1b65aa2fd28ee22c28745a2adbb5fab482eef7cc5e","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","kwargs":{"arity":"max_one","bin":"provenance","command":"revise"}}
	cmd.Args = cobra.MaximumNArgs(1)
	// /sedum:setArgs:max_one
	// sedum:anchor:flags
	// sedum:addFlag:string {"tier":"owned","record":"prov-2026-5f69977b","gen":"sha256:cd83ad7d20c1e101894b6c1b65aa2fd28ee22c28745a2adbb5fab482eef7cc5e","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"revise","flag":"record","type":"string","usage":"Record to revise (or give it as the argument)"}}
	cmd.Flags().String("record", "", "Record to revise (or give it as the argument)")
	// /sedum:addFlag:string
	// sedum:addFlag:string {"tier":"owned","record":"prov-2026-5f69977b","gen":"sha256:cd83ad7d20c1e101894b6c1b65aa2fd28ee22c28745a2adbb5fab482eef7cc5e","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"revise","flag":"reason","type":"string","usage":"Why this version exists (required for a new version)"}}
	cmd.Flags().String("reason", "", "Why this version exists (required for a new version)")
	// /sedum:addFlag:string
	// sedum:addFlag:bool {"tier":"owned","record":"prov-2026-5f69977b","gen":"sha256:cd83ad7d20c1e101894b6c1b65aa2fd28ee22c28745a2adbb5fab482eef7cc5e","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"revise","flag":"fix","type":"bool","usage":"Mark this version as a fix (reason_kind: fix)"}}
	cmd.Flags().Bool("fix", false, "Mark this version as a fix (reason_kind: fix)")
	// /sedum:addFlag:bool
	// sedum:anchor:subcommands
	return cmd
}
