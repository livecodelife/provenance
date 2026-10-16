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
