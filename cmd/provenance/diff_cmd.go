package main

import "github.com/spf13/cobra"

func newDiffCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"v":2,"tier":"owned","record":"prov-2026-bfd55535","gen":"sha256:98fb8e2f8fd67f02e2cabc1153d5dfb84f1f27f5c19d0582f9236cd5fa8e4f10","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","id":{"bin":"provenance","command":"diff"},"kwargs":"sha256:82cd4c5b1b869a138eb947c7b16eab1b62445bc1e13b1df864d1a07d26a43793","body":"sha256:e026e90c2f7059be96cfee67a902f4029086cf6ea1447d5d5b09af806c8944df"}
	cmd.Use = "diff [id] [vA vB]"
	cmd.Short = "Show what changed between two versions of a record"
	// /sedum:describeCommand
	// sedum:setRun {"v":2,"tier":"owned","record":"prov-2026-bfd55535","gen":"sha256:98fb8e2f8fd67f02e2cabc1153d5dfb84f1f27f5c19d0582f9236cd5fa8e4f10","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","id":{"bin":"provenance","command":"diff"},"kwargs":"sha256:e65b2c6b081d304a2ef912bd19d07fd3f207b6678322ba0067ad0a18a03fae31","body":"sha256:ae764439f66321386d97f6eb77c725487d9d1bf8675729333a8e241e85f33e07"}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runDiff(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:any {"v":2,"tier":"owned","record":"prov-2026-bfd55535","gen":"sha256:98fb8e2f8fd67f02e2cabc1153d5dfb84f1f27f5c19d0582f9236cd5fa8e4f10","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","id":{"bin":"provenance","command":"diff"},"kwargs":"sha256:29a7d4d546c2baac56139d46bc3d98993189e44c960be37ee678f316440b624c","body":"sha256:b30b0c92d5bd257e52e3c8256ebc76c0673d14e4c2e996925a6699b264276231"}
	cmd.Args = cobra.ArbitraryArgs
	// /sedum:setArgs:any
	// sedum:anchor:flags
	// sedum:addFlag:string {"v":2,"tier":"owned","record":"prov-2026-bfd55535","gen":"sha256:98fb8e2f8fd67f02e2cabc1153d5dfb84f1f27f5c19d0582f9236cd5fa8e4f10","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"diff","flag":"record"},"kwargs":"sha256:c69d8af6e7cafd71cdd91f8bdc59aa1ec74bc1c6d76b5260c1cba762d8864410","body":"sha256:fa055f65c644e5cbf6153d106247f1f1ba0ab5547c6221db296e2befa4097112"}
	cmd.Flags().String("record", "", "Record to diff (or give it as the first argument)")
	// /sedum:addFlag:string
	// sedum:anchor:subcommands
	return cmd
}
