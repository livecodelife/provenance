package main

import (
	"os"

	"github.com/spf13/cobra"
)

func newRootCmd() *cobra.Command {
	root := &cobra.Command{SilenceUsage: true}
	// sedum:anchor:root
	// sedum:describeRoot {"tier":"owned","record":"prov-2026-0000c0b1","gen":"sha256:548e4aad36d237333a5e40aeca4bb3e36cff37295c0814441f79016f399ecb56","pkg":"sha256:5a322fab40b482e662e745110007c5b5c546721328585695e4325ba5626a9685","kwargs":{"bin":"demo","short":"Demo CLI grown by the cobra package","use":"demo"}}
	root.Use = "demo"
	root.Short = "Demo CLI grown by the cobra package"
	// /sedum:describeRoot
	// sedum:anchor:persistent_flags
	// sedum:addPersistentFlag:string {"tier":"owned","record":"prov-2026-0000c0b2","gen":"sha256:f12b8a04384dadc992a27af33c513f1a34362ccf9bc6a3e87192bbb0c9d9f508","pkg":"sha256:67d9fd15ecda53fb7be1b9d541b0b822634ff87778b31aeaeaa7f824971422dc","kwargs":{"bin":"demo","default":"","flag":"config","shorthand":"c","type":"string","usage":"Config file"}}
	root.PersistentFlags().StringP("config", "c", "", "Config file")
	// /sedum:addPersistentFlag:string
	// sedum:addPersistentFlag:bool {"tier":"owned","record":"prov-2026-0000c0b2","gen":"sha256:f12b8a04384dadc992a27af33c513f1a34362ccf9bc6a3e87192bbb0c9d9f508","pkg":"sha256:67d9fd15ecda53fb7be1b9d541b0b822634ff87778b31aeaeaa7f824971422dc","kwargs":{"bin":"demo","default":"","flag":"verbose","shorthand":"","type":"bool","usage":"Verbose output"}}
	root.PersistentFlags().BoolP("verbose", "", false, "Verbose output")
	// /sedum:addPersistentFlag:bool
	// sedum:anchor:commands
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-0000c0b1","gen":"sha256:548e4aad36d237333a5e40aeca4bb3e36cff37295c0814441f79016f399ecb56","pkg":"sha256:3e4cafddc21b78aee3cfb68894433b350206b34709a8701a3932b4908bdc287a","kwargs":{"bin":"demo","command":"greet"}}
	root.AddCommand(newGreetCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-0000c0b1","gen":"sha256:548e4aad36d237333a5e40aeca4bb3e36cff37295c0814441f79016f399ecb56","pkg":"sha256:3e4cafddc21b78aee3cfb68894433b350206b34709a8701a3932b4908bdc287a","kwargs":{"bin":"demo","command":"echo"}}
	root.AddCommand(newEchoCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-0000c0b1","gen":"sha256:548e4aad36d237333a5e40aeca4bb3e36cff37295c0814441f79016f399ecb56","pkg":"sha256:3e4cafddc21b78aee3cfb68894433b350206b34709a8701a3932b4908bdc287a","kwargs":{"bin":"demo","command":"list"}}
	root.AddCommand(newListCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-0000c0b2","gen":"sha256:f12b8a04384dadc992a27af33c513f1a34362ccf9bc6a3e87192bbb0c9d9f508","pkg":"sha256:3e4cafddc21b78aee3cfb68894433b350206b34709a8701a3932b4908bdc287a","kwargs":{"bin":"demo","command":"hooks"}}
	root.AddCommand(newHooksCmd())
	// /sedum:registerCommand
	return root
}

func main() {
	if err := newRootCmd().Execute(); err != nil {
		os.Exit(1)
	}
}
