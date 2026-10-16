package main

import "github.com/spf13/cobra"

func newEditCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"tier":"owned","record":"prov-2026-57b8a95b","gen":"sha256:03858cfa7091f66aef9fd6bec812b41738d06e7d9cb80e64c32135b367f4832b","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","kwargs":{"bin":"provenance","command":"edit","short":"Change a record's content fields without an editor and save a new version","use":"edit [id]"}}
	cmd.Use = "edit [id]"
	cmd.Short = "Change a record's content fields without an editor and save a new version"
	// /sedum:describeCommand
	// sedum:setRun {"tier":"owned","record":"prov-2026-57b8a95b","gen":"sha256:03858cfa7091f66aef9fd6bec812b41738d06e7d9cb80e64c32135b367f4832b","pkg":"sha256:fadae715925d582c263ce084f15c491995be26d77f0489abc836c658e094878f","kwargs":{"bin":"provenance","command":"edit"}}
	cmd.RunE = runEdit
	// /sedum:setRun
	// sedum:setArgs:max_one {"tier":"owned","record":"prov-2026-57b8a95b","gen":"sha256:03858cfa7091f66aef9fd6bec812b41738d06e7d9cb80e64c32135b367f4832b","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","kwargs":{"arity":"max_one","bin":"provenance","command":"edit"}}
	cmd.Args = cobra.MaximumNArgs(1)
	// /sedum:setArgs:max_one
	// sedum:anchor:flags
	// sedum:addFlag:string {"tier":"owned","record":"prov-2026-57b8a95b","gen":"sha256:03858cfa7091f66aef9fd6bec812b41738d06e7d9cb80e64c32135b367f4832b","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"edit","flag":"record","type":"string","usage":"Record to edit (or give it as the argument)"}}
	cmd.Flags().String("record", "", "Record to edit (or give it as the argument)")
	// /sedum:addFlag:string
	// sedum:addFlag:string {"tier":"owned","record":"prov-2026-57b8a95b","gen":"sha256:03858cfa7091f66aef9fd6bec812b41738d06e7d9cb80e64c32135b367f4832b","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"edit","flag":"title","type":"string","usage":"New title"}}
	cmd.Flags().String("title", "", "New title")
	// /sedum:addFlag:string
	// sedum:addFlag:string {"tier":"owned","record":"prov-2026-57b8a95b","gen":"sha256:03858cfa7091f66aef9fd6bec812b41738d06e7d9cb80e64c32135b367f4832b","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"edit","flag":"intent","type":"string","usage":"New intent"}}
	cmd.Flags().String("intent", "", "New intent")
	// /sedum:addFlag:string
	// sedum:addFlag:strings {"tier":"owned","record":"prov-2026-57b8a95b","gen":"sha256:03858cfa7091f66aef9fd6bec812b41738d06e7d9cb80e64c32135b367f4832b","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"edit","flag":"add-constraint","type":"strings","usage":"Append a constraint (repeatable)"}}
	cmd.Flags().StringArray("add-constraint", nil, "Append a constraint (repeatable)")
	// /sedum:addFlag:strings
	// sedum:addFlag:strings {"tier":"owned","record":"prov-2026-57b8a95b","gen":"sha256:03858cfa7091f66aef9fd6bec812b41738d06e7d9cb80e64c32135b367f4832b","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"edit","flag":"remove-constraint","type":"strings","usage":"Remove constraint number n, counting from 1 (repeatable)"}}
	cmd.Flags().StringArray("remove-constraint", nil, "Remove constraint number n, counting from 1 (repeatable)")
	// /sedum:addFlag:strings
	// sedum:addFlag:strings {"tier":"owned","record":"prov-2026-57b8a95b","gen":"sha256:03858cfa7091f66aef9fd6bec812b41738d06e7d9cb80e64c32135b367f4832b","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"edit","flag":"add-scope","type":"strings","usage":"Append an affected_scope entry (repeatable)"}}
	cmd.Flags().StringArray("add-scope", nil, "Append an affected_scope entry (repeatable)")
	// /sedum:addFlag:strings
	// sedum:addFlag:strings {"tier":"owned","record":"prov-2026-57b8a95b","gen":"sha256:03858cfa7091f66aef9fd6bec812b41738d06e7d9cb80e64c32135b367f4832b","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"edit","flag":"remove-scope","type":"strings","usage":"Remove an affected_scope entry by value (repeatable)"}}
	cmd.Flags().StringArray("remove-scope", nil, "Remove an affected_scope entry by value (repeatable)")
	// /sedum:addFlag:strings
	// sedum:addFlag:strings {"tier":"owned","record":"prov-2026-57b8a95b","gen":"sha256:03858cfa7091f66aef9fd6bec812b41738d06e7d9cb80e64c32135b367f4832b","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"edit","flag":"add-forbidden","type":"strings","usage":"Append a forbidden_scope entry (repeatable)"}}
	cmd.Flags().StringArray("add-forbidden", nil, "Append a forbidden_scope entry (repeatable)")
	// /sedum:addFlag:strings
	// sedum:addFlag:strings {"tier":"owned","record":"prov-2026-57b8a95b","gen":"sha256:03858cfa7091f66aef9fd6bec812b41738d06e7d9cb80e64c32135b367f4832b","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"edit","flag":"remove-forbidden","type":"strings","usage":"Remove a forbidden_scope entry by value (repeatable)"}}
	cmd.Flags().StringArray("remove-forbidden", nil, "Remove a forbidden_scope entry by value (repeatable)")
	// /sedum:addFlag:strings
	// sedum:addFlag:string {"tier":"owned","record":"prov-2026-57b8a95b","gen":"sha256:03858cfa7091f66aef9fd6bec812b41738d06e7d9cb80e64c32135b367f4832b","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"edit","flag":"reason","type":"string","usage":"Why this version exists (required for a new version)"}}
	cmd.Flags().String("reason", "", "Why this version exists (required for a new version)")
	// /sedum:addFlag:string
	// sedum:addFlag:bool {"tier":"owned","record":"prov-2026-57b8a95b","gen":"sha256:03858cfa7091f66aef9fd6bec812b41738d06e7d9cb80e64c32135b367f4832b","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"edit","flag":"fix","type":"bool","usage":"Mark this version as a fix (reason_kind: fix)"}}
	cmd.Flags().Bool("fix", false, "Mark this version as a fix (reason_kind: fix)")
	// /sedum:addFlag:bool
	// sedum:anchor:subcommands
	return cmd
}
