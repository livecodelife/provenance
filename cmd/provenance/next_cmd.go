package main

import "github.com/spf13/cobra"

func newNextCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"v":2,"tier":"owned","record":"prov-2026-11876813","gen":"sha256:090b658e8472ed58a1fb75e4a4101e7f16690a1a70b99759f485f9af7dfaeb18","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","id":{"bin":"provenance","command":"next"},"kwargs":"sha256:ddfd3f594effa3cb7d0bad52f7856653e39db0c9766ff49bd7a02a3c0959f481","body":"sha256:0f3364b4f3b7efa76c16a79829fc0fff44e46e527e94426fd018903d0e294c39"}
	cmd.Use = "next [files...]"
	cmd.Short = "Compute the single correct next provenance action"
	// /sedum:describeCommand
	// sedum:setRun {"v":2,"tier":"owned","record":"prov-2026-11876813","gen":"sha256:090b658e8472ed58a1fb75e4a4101e7f16690a1a70b99759f485f9af7dfaeb18","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","id":{"bin":"provenance","command":"next"},"kwargs":"sha256:395e34d4f4eed74c05856688100d17a90d3ecd7be4c02fc203eb3b0b15e3a8d6","body":"sha256:bb44922898d44b90a2a97d9d05b0aeb2291dff922ab57ce4f87fb0ec125253b7"}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runNext(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:any {"v":2,"tier":"owned","record":"prov-2026-11876813","gen":"sha256:090b658e8472ed58a1fb75e4a4101e7f16690a1a70b99759f485f9af7dfaeb18","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","id":{"bin":"provenance","command":"next"},"kwargs":"sha256:8ae21b7f4266e32b4e7d874bcc6b802249558d04a4ef66835abddc5e05915073","body":"sha256:b30b0c92d5bd257e52e3c8256ebc76c0673d14e4c2e996925a6699b264276231"}
	cmd.Args = cobra.ArbitraryArgs
	// /sedum:setArgs:any
	// sedum:anchor:flags
	// sedum:addFlag:strings {"v":2,"tier":"owned","record":"prov-2026-11876813","gen":"sha256:090b658e8472ed58a1fb75e4a4101e7f16690a1a70b99759f485f9af7dfaeb18","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"next","flag":"files"},"kwargs":"sha256:c6d1a77f167f60d0f44697ca60e016b7e06bf960b5e1180c0c37ececfe327c5c","body":"sha256:f2bbd755d847cf71245100bb2d7a15424a0d0e0c26e4d165a886a65fe84d137c"}
	cmd.Flags().StringArray("files", nil, "Files you intend to change (plan before editing)")
	// /sedum:addFlag:strings
	// sedum:addFlag:strings {"v":2,"tier":"owned","record":"prov-2026-11876813","gen":"sha256:090b658e8472ed58a1fb75e4a4101e7f16690a1a70b99759f485f9af7dfaeb18","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"next","flag":"plan"},"kwargs":"sha256:bbd5bf7163838f29cb78f32f2b44fe6fdbf9b2193d8761298ce1a61a4c703bc7","body":"sha256:127dfeb6efa5575629a9970ad324939a738ae33913a96f328c71959a689142df"}
	cmd.Flags().StringArray("plan", nil, "Alias for --files")
	// /sedum:addFlag:strings
	// sedum:addFlag:bool {"v":2,"tier":"owned","record":"prov-2026-11876813","gen":"sha256:090b658e8472ed58a1fb75e4a4101e7f16690a1a70b99759f485f9af7dfaeb18","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"next","flag":"json"},"kwargs":"sha256:96b49c69d0dfef3a30231eebc8cf578b902b3fdf2e007cafe70b7e0cc432dc5f","body":"sha256:bea1c7577b5a599a972886858936b7d7e9cf68ab7660996d6318dfa45b27826a"}
	cmd.Flags().Bool("json", false, "Machine-readable output (for hooks and agents)")
	// /sedum:addFlag:bool
	// sedum:addFlag:string {"v":2,"tier":"owned","record":"prov-2026-11876813","gen":"sha256:090b658e8472ed58a1fb75e4a4101e7f16690a1a70b99759f485f9af7dfaeb18","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"next","flag":"format"},"kwargs":"sha256:988f0354bcb2ab6542140eb9a35fbe704f39ee07011078bf558ea194ed821fd4","body":"sha256:5a11ae1e68306dbe22a26ffec6af235a5dc44898684981a9cddacca632a7b196"}
	cmd.Flags().String("format", "", "Output format (human|json)")
	// /sedum:addFlag:string
	// sedum:anchor:subcommands
	return cmd
}
