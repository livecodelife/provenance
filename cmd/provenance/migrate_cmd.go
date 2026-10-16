package main

import "github.com/spf13/cobra"

func newMigrateCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"tier":"owned","record":"prov-2026-76a125e6","gen":"sha256:5ef3ff16494617aabe7af9742864e5d8974512a0189aae79aff406d6c661f23d","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","kwargs":{"bin":"provenance","command":"migrate","short":"Convert v1 records to versioned v2 records","use":"migrate"}}
	cmd.Use = "migrate"
	cmd.Short = "Convert v1 records to versioned v2 records"
	// /sedum:describeCommand
	// sedum:setRun {"tier":"owned","record":"prov-2026-76a125e6","gen":"sha256:5ef3ff16494617aabe7af9742864e5d8974512a0189aae79aff406d6c661f23d","pkg":"sha256:fadae715925d582c263ce084f15c491995be26d77f0489abc836c658e094878f","kwargs":{"bin":"provenance","command":"migrate"}}
	cmd.RunE = runMigrate
	// /sedum:setRun
	// sedum:setArgs:none {"tier":"owned","record":"prov-2026-76a125e6","gen":"sha256:5ef3ff16494617aabe7af9742864e5d8974512a0189aae79aff406d6c661f23d","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","kwargs":{"arity":"none","bin":"provenance","command":"migrate"}}
	cmd.Args = cobra.NoArgs
	// /sedum:setArgs:none
	// sedum:anchor:flags
	// sedum:addFlag:bool {"tier":"owned","record":"prov-2026-76a125e6","gen":"sha256:5ef3ff16494617aabe7af9742864e5d8974512a0189aae79aff406d6c661f23d","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"migrate","flag":"plan","type":"bool","usage":"Print the migration plan and write nothing"}}
	cmd.Flags().Bool("plan", false, "Print the migration plan and write nothing")
	// /sedum:addFlag:bool
	// sedum:addFlag:bool {"tier":"owned","record":"prov-2026-76a125e6","gen":"sha256:5ef3ff16494617aabe7af9742864e5d8974512a0189aae79aff406d6c661f23d","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"migrate","flag":"json","type":"bool","usage":"Print the plan as JSON (with --plan)"}}
	cmd.Flags().Bool("json", false, "Print the plan as JSON (with --plan)")
	// /sedum:addFlag:bool
	// sedum:addFlag:bool {"tier":"owned","record":"prov-2026-76a125e6","gen":"sha256:5ef3ff16494617aabe7af9742864e5d8974512a0189aae79aff406d6c661f23d","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"migrate","flag":"accept-suggestions","type":"bool","usage":"Keep every suggested record separate"}}
	cmd.Flags().Bool("accept-suggestions", false, "Keep every suggested record separate")
	// /sedum:addFlag:bool
	// sedum:addFlag:strings {"tier":"owned","record":"prov-2026-76a125e6","gen":"sha256:5ef3ff16494617aabe7af9742864e5d8974512a0189aae79aff406d6c661f23d","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"migrate","flag":"keep-separate","type":"strings","usage":"Keep this absorbed record separate (repeatable)"}}
	cmd.Flags().StringArray("keep-separate", nil, "Keep this absorbed record separate (repeatable)")
	// /sedum:addFlag:strings
	// sedum:addFlag:strings {"tier":"owned","record":"prov-2026-76a125e6","gen":"sha256:5ef3ff16494617aabe7af9742864e5d8974512a0189aae79aff406d6c661f23d","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"migrate","flag":"fold","type":"strings","usage":"Fold this record even when suggested to keep it separate (repeatable)"}}
	cmd.Flags().StringArray("fold", nil, "Fold this record even when suggested to keep it separate (repeatable)")
	// /sedum:addFlag:strings
	// sedum:addFlag:string {"tier":"owned","record":"prov-2026-76a125e6","gen":"sha256:5ef3ff16494617aabe7af9742864e5d8974512a0189aae79aff406d6c661f23d","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"migrate","flag":"workspace","type":"string","usage":"Directory holding local checkouts of shared repos (default: the repo's parent)"}}
	cmd.Flags().String("workspace", "", "Directory holding local checkouts of shared repos (default: the repo's parent)")
	// /sedum:addFlag:string
	// sedum:anchor:subcommands
	return cmd
}
