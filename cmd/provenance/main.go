package main

import (
	"os"

	"github.com/spf13/cobra"
)

func newRootCmd() *cobra.Command {
	root := &cobra.Command{SilenceUsage: true}
	// sedum:anchor:root
	// sedum:describeRoot {"tier":"owned","record":"prov-2026-76a125e6","gen":"sha256:5ef3ff16494617aabe7af9742864e5d8974512a0189aae79aff406d6c661f23d","pkg":"sha256:0fc81c42ecd7f40945040b8bc7254475c058082a098a34b5f2c8660225ff618e","kwargs":{"bin":"provenance","short":"Versioned provenance records","use":"provenance"}}
	root.Use = "provenance"
	root.Short = "Versioned provenance records"
	// /sedum:describeRoot
	// sedum:anchor:persistent_flags
	// sedum:addShortPersistentFlag:string {"tier":"owned","record":"prov-2026-76a125e6","gen":"sha256:5ef3ff16494617aabe7af9742864e5d8974512a0189aae79aff406d6c661f23d","pkg":"sha256:9a5acc9c3ccb4219037b31eecd6f40891828ebeb30d9dcc2e18be6a46d59dcae","kwargs":{"bin":"provenance","flag":"config","shorthand":"c","type":"string","usage":"Config file (default .provenance.yml, else the provenance block of .linespec.yml)"}}
	root.PersistentFlags().StringP("config", "c", "", "Config file (default .provenance.yml, else the provenance block of .linespec.yml)")
	// /sedum:addShortPersistentFlag:string
	// sedum:anchor:commands
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-44a07768","gen":"sha256:f38eb00e2aad6097d9b8e6e92ef680d7801181f69cad8b513265bb2f4fad72d6","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"provenance","command":"create"}}
	root.AddCommand(newCreateCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-57b8a95b","gen":"sha256:03858cfa7091f66aef9fd6bec812b41738d06e7d9cb80e64c32135b367f4832b","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"provenance","command":"edit"}}
	root.AddCommand(newEditCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-5f69977b","gen":"sha256:cd83ad7d20c1e101894b6c1b65aa2fd28ee22c28745a2adbb5fab482eef7cc5e","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"provenance","command":"revise"}}
	root.AddCommand(newReviseCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-6db88b75","gen":"sha256:cf831631b6137cd400fe64495057052c30d745b64a30dbfcb2ef52004d5e57e8","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"provenance","command":"open"}}
	root.AddCommand(newOpenCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-76a125e6","gen":"sha256:5ef3ff16494617aabe7af9742864e5d8974512a0189aae79aff406d6c661f23d","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"provenance","command":"migrate"}}
	root.AddCommand(newMigrateCmd())
	// /sedum:registerCommand
	return root
}

func main() {
	if err := newRootCmd().Execute(); err != nil {
		os.Exit(1)
	}
}
