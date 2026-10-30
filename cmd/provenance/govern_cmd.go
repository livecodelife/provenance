package main

import "github.com/spf13/cobra"

func newGovernCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"v":2,"tier":"owned","record":"prov-2026-9ec4a843","gen":"sha256:cc0301a10fcd903006e4f2bed32a32b6a082f192eb2c42bb2c80462870e50daf","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","id":{"bin":"provenance","command":"govern"},"kwargs":"sha256:26210074b7b9a5ad1c115d4d0e5cc95de98bad3a9f821ce28f13be2944d62190","body":"sha256:a60aade87829f2642235faf57214dcbd69852e0b2dc42e2cb5ae119e1fdf917e"}
	cmd.Use = "govern [files...]"
	cmd.Short = "List the active records that govern the given files"
	// /sedum:describeCommand
	// sedum:setRun {"v":2,"tier":"owned","record":"prov-2026-9ec4a843","gen":"sha256:cc0301a10fcd903006e4f2bed32a32b6a082f192eb2c42bb2c80462870e50daf","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","id":{"bin":"provenance","command":"govern"},"kwargs":"sha256:fe1df7947e25554f2ffebb2c2e3b799532ec6d224cf3847e3ad4751ad6b85c0b","body":"sha256:a2e79d54e24733fd13cb11defe47c6c44323ef4383b4b0cc23fd5248283f93e1"}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runGovern(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:any {"v":2,"tier":"owned","record":"prov-2026-9ec4a843","gen":"sha256:cc0301a10fcd903006e4f2bed32a32b6a082f192eb2c42bb2c80462870e50daf","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","id":{"bin":"provenance","command":"govern"},"kwargs":"sha256:d2b96f64e49bc3285b1ee2313fdacadaa994ab55763d510bca6b77d72c5bfe20","body":"sha256:b30b0c92d5bd257e52e3c8256ebc76c0673d14e4c2e996925a6699b264276231"}
	cmd.Args = cobra.ArbitraryArgs
	// /sedum:setArgs:any
	// sedum:anchor:flags
	// sedum:addFlag:strings {"v":2,"tier":"owned","record":"prov-2026-9ec4a843","gen":"sha256:cc0301a10fcd903006e4f2bed32a32b6a082f192eb2c42bb2c80462870e50daf","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"govern","flag":"files"},"kwargs":"sha256:eb83df61c7c0744e4ecfcd7db978b4edfa9f1cb231776fcb95815191ba1682a0","body":"sha256:ed6bc2de76c960c82e30b4a516bde8cb4a8b49a1dd0bbb310e11aae797cb3280"}
	cmd.Flags().StringArray("files", nil, "Explicit file list (alternative to arguments)")
	// /sedum:addFlag:strings
	// sedum:addFlag:string {"v":2,"tier":"owned","record":"prov-2026-9ec4a843","gen":"sha256:cc0301a10fcd903006e4f2bed32a32b6a082f192eb2c42bb2c80462870e50daf","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"govern","flag":"format"},"kwargs":"sha256:dfffc727349c8b6db24e62db996cee5a2a45da904705ea022984fe3afc24e39b","body":"sha256:5a11ae1e68306dbe22a26ffec6af235a5dc44898684981a9cddacca632a7b196"}
	cmd.Flags().String("format", "", "Output format (human|json)")
	// /sedum:addFlag:string
	// sedum:addFlag:bool {"v":2,"tier":"owned","record":"prov-2026-9ec4a843","gen":"sha256:cc0301a10fcd903006e4f2bed32a32b6a082f192eb2c42bb2c80462870e50daf","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"govern","flag":"json"},"kwargs":"sha256:7fe58a5a5b0e68b766043715a3ca5fa37be68e7084baf804bbf6abf140672de2","body":"sha256:13068068bc38fbfc704cdead532232eb306160aeaaa873e4b6576fd93cb18f81"}
	cmd.Flags().Bool("json", false, "Alias for --format json")
	// /sedum:addFlag:bool
	// sedum:anchor:subcommands
	return cmd
}
