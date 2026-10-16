package main

import (
	"os"

	"github.com/spf13/cobra"
)

func newRootCmd() *cobra.Command {
	root := &cobra.Command{SilenceUsage: true}
	// sedum:anchor:root
	// sedum:describeRoot {"tier":"owned","record":"prov-2026-0000c0b1","gen":"sha256:5062a337e1576589604ab1efde948f0a19edcbda6e2a5a31ece0ee2edfec42ca","pkg":"sha256:dd72f7aa5ae9482d6600af196b08e20f5279b7c0d560da5c20d8c829f2a20e7c","kwargs":{"bin":"demo","short":"Demo CLI grown by the cobra package","use":"demo"}}
	root.Use = "demo"
	root.Short = "Demo CLI grown by the cobra package"
	// /sedum:describeRoot
	// sedum:anchor:persistent_flags
	// sedum:addShortPersistentFlag:string {"tier":"owned","record":"prov-2026-0000c0b2","gen":"sha256:8b351f5ad014f88b418f9ba4898111b02258f9432172a203a8d2e033607bb469","pkg":"sha256:fbd768d277c5ae0145be975e733052e1688c43917ee7e5328de3e03eab6d410f","kwargs":{"bin":"demo","flag":"config","shorthand":"c","type":"string","usage":"Config file"}}
	root.PersistentFlags().StringP("config", "c", "", "Config file")
	// /sedum:addShortPersistentFlag:string
	// sedum:addPersistentFlag:bool {"tier":"owned","record":"prov-2026-0000c0b2","gen":"sha256:8b351f5ad014f88b418f9ba4898111b02258f9432172a203a8d2e033607bb469","pkg":"sha256:c5dec76f3dc6548e7fa86556be84b9872029d3217cca211de191fee0e2077f54","kwargs":{"bin":"demo","flag":"verbose","type":"bool","usage":"Verbose output"}}
	root.PersistentFlags().Bool("verbose", false, "Verbose output")
	// /sedum:addPersistentFlag:bool
	// sedum:anchor:commands
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-0000c0b1","gen":"sha256:5062a337e1576589604ab1efde948f0a19edcbda6e2a5a31ece0ee2edfec42ca","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"demo","command":"greet"}}
	root.AddCommand(newGreetCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-0000c0b1","gen":"sha256:5062a337e1576589604ab1efde948f0a19edcbda6e2a5a31ece0ee2edfec42ca","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"demo","command":"echo"}}
	root.AddCommand(newEchoCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-0000c0b1","gen":"sha256:5062a337e1576589604ab1efde948f0a19edcbda6e2a5a31ece0ee2edfec42ca","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"demo","command":"list"}}
	root.AddCommand(newListCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-0000c0b1","gen":"sha256:5062a337e1576589604ab1efde948f0a19edcbda6e2a5a31ece0ee2edfec42ca","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"demo","command":"show-id"}}
	root.AddCommand(newShowIDCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-0000c0b2","gen":"sha256:8b351f5ad014f88b418f9ba4898111b02258f9432172a203a8d2e033607bb469","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"demo","command":"hooks"}}
	root.AddCommand(newHooksCmd())
	// /sedum:registerCommand
	return root
}

func main() {
	if err := newRootCmd().Execute(); err != nil {
		os.Exit(1)
	}
}
