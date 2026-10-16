package main

import (
	"os"

	"github.com/spf13/cobra"
)

func newRootCmd() *cobra.Command {
	root := &cobra.Command{SilenceUsage: true}
	// sedum:anchor:root
	// sedum:describeRoot {"tier":"owned","record":"prov-2026-0000c0b1","gen":"sha256:5062a337e1576589604ab1efde948f0a19edcbda6e2a5a31ece0ee2edfec42ca","pkg":"sha256:0fc81c42ecd7f40945040b8bc7254475c058082a098a34b5f2c8660225ff618e","kwargs":{"bin":"demo","short":"Demo CLI grown by the cobra package","use":"demo"}}
	root.Use = "demo"
	root.Short = "Demo CLI grown by the cobra package"
	// /sedum:describeRoot
	// sedum:anchor:persistent_flags
	// sedum:addShortPersistentFlag:string {"tier":"owned","record":"prov-2026-0000c0b2","gen":"sha256:8b351f5ad014f88b418f9ba4898111b02258f9432172a203a8d2e033607bb469","pkg":"sha256:9a5acc9c3ccb4219037b31eecd6f40891828ebeb30d9dcc2e18be6a46d59dcae","kwargs":{"bin":"demo","flag":"config","shorthand":"c","type":"string","usage":"Config file"}}
	root.PersistentFlags().StringP("config", "c", "", "Config file")
	// /sedum:addShortPersistentFlag:string
	// sedum:addPersistentFlag:bool {"tier":"owned","record":"prov-2026-0000c0b2","gen":"sha256:8b351f5ad014f88b418f9ba4898111b02258f9432172a203a8d2e033607bb469","pkg":"sha256:91bf4789adc094a18f09294be9dd754f48adf139730ce2485393ef023022172f","kwargs":{"bin":"demo","flag":"verbose","type":"bool","usage":"Verbose output"}}
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
