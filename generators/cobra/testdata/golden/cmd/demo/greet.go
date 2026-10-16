package main

import "github.com/spf13/cobra"

func newGreetCmd() *cobra.Command {
	cmd := &cobra.Command{RunE: runGreet}
	// sedum:anchor:command
	// sedum:describeCommand {"tier":"owned","record":"prov-2026-0000c0b1","gen":"sha256:548e4aad36d237333a5e40aeca4bb3e36cff37295c0814441f79016f399ecb56","pkg":"sha256:9958f264ddf8e57ee6107f5f5aacb530fb52d42dbf161cd7f31d050f90486529","kwargs":{"bin":"demo","command":"greet","short":"Greet someone","use":"greet <name>"}}
	cmd.Use = "greet <name>"
	cmd.Short = "Greet someone"
	// /sedum:describeCommand
	// sedum:setArgs:one {"tier":"owned","record":"prov-2026-0000c0b1","gen":"sha256:548e4aad36d237333a5e40aeca4bb3e36cff37295c0814441f79016f399ecb56","pkg":"sha256:84522566c0a6296808b5e0750805ad85951fe40c0b9efaa6bbb8bfff853d4eb7","kwargs":{"arity":"one","bin":"demo","command":"greet"}}
	cmd.Args = cobra.ExactArgs(1)
	// /sedum:setArgs:one
	// sedum:anchor:flags
	// sedum:addFlag:string {"tier":"owned","record":"prov-2026-0000c0b1","gen":"sha256:548e4aad36d237333a5e40aeca4bb3e36cff37295c0814441f79016f399ecb56","pkg":"sha256:8f311eea74643a2f16c68f4e5b8853bc4f2710370d7971247c944d75c33b139c","kwargs":{"bin":"demo","command":"greet","default":"hello","flag":"greeting","shorthand":"g","type":"string","usage":"Word to greet with"}}
	cmd.Flags().StringP("greeting", "g", "hello", "Word to greet with")
	// /sedum:addFlag:string
	// sedum:addFlag:bool {"tier":"owned","record":"prov-2026-0000c0b1","gen":"sha256:548e4aad36d237333a5e40aeca4bb3e36cff37295c0814441f79016f399ecb56","pkg":"sha256:8f311eea74643a2f16c68f4e5b8853bc4f2710370d7971247c944d75c33b139c","kwargs":{"bin":"demo","command":"greet","default":"","flag":"shout","shorthand":"","type":"bool","usage":"Print in upper case"}}
	cmd.Flags().BoolP("shout", "", false, "Print in upper case")
	// /sedum:addFlag:bool
	// sedum:addFlag:int {"tier":"owned","record":"prov-2026-0000c0b1","gen":"sha256:548e4aad36d237333a5e40aeca4bb3e36cff37295c0814441f79016f399ecb56","pkg":"sha256:8f311eea74643a2f16c68f4e5b8853bc4f2710370d7971247c944d75c33b139c","kwargs":{"bin":"demo","command":"greet","default":"","flag":"times","shorthand":"","type":"int","usage":"Repeat count"}}
	cmd.Flags().IntP("times", "", 0, "Repeat count")
	// /sedum:addFlag:int
	// sedum:addFlag:strings {"tier":"owned","record":"prov-2026-0000c0b1","gen":"sha256:548e4aad36d237333a5e40aeca4bb3e36cff37295c0814441f79016f399ecb56","pkg":"sha256:8f311eea74643a2f16c68f4e5b8853bc4f2710370d7971247c944d75c33b139c","kwargs":{"bin":"demo","command":"greet","default":"","flag":"tag","shorthand":"","type":"strings","usage":"Tags, repeatable"}}
	cmd.Flags().StringArrayP("tag", "", nil, "Tags, repeatable")
	// /sedum:addFlag:strings
	// sedum:anchor:subcommands
	return cmd
}
