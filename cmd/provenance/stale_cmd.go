package main

import "github.com/spf13/cobra"

func newStaleCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"v":2,"tier":"owned","record":"prov-2026-f476b640","gen":"sha256:24167071b71953f56b7a9351b440d09a67530fb3dc4508cd3e405b0ad9bb4806","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","id":{"bin":"provenance","command":"stale"},"kwargs":"sha256:02a113976a8958d53417495a2efbff8da471d5a2169c7475f734b297e035fad1","body":"sha256:d8d0af2993839c0924ea2c7a98887df4314477b9eb967a6d95a33366e7cd6d61"}
	cmd.Use = "stale"
	cmd.Short = "List records whose implements_version is behind their parent"
	// /sedum:describeCommand
	// sedum:setRun {"v":2,"tier":"owned","record":"prov-2026-f476b640","gen":"sha256:24167071b71953f56b7a9351b440d09a67530fb3dc4508cd3e405b0ad9bb4806","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","id":{"bin":"provenance","command":"stale"},"kwargs":"sha256:be070c465b03cea9756295b67c88eea09bdd1300b0fa0c027b5f86ebdce83d72","body":"sha256:1172b8aa4b62e97c37884513c6723df84d93dbf700684e8ebce1ba9d8c19e503"}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runStale(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:none {"v":2,"tier":"owned","record":"prov-2026-f476b640","gen":"sha256:24167071b71953f56b7a9351b440d09a67530fb3dc4508cd3e405b0ad9bb4806","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","id":{"bin":"provenance","command":"stale"},"kwargs":"sha256:4a20ef2a590dbd97132bb84719e2820db2b179ea52966de8f91a66d8027bd3ee","body":"sha256:778b1e3eb6004fe613ecb8793efaf401a29f1afdc382707cea55c05601aefd79"}
	cmd.Args = cobra.NoArgs
	// /sedum:setArgs:none
	// sedum:anchor:flags
	// sedum:addFlag:bool {"v":2,"tier":"owned","record":"prov-2026-f476b640","gen":"sha256:24167071b71953f56b7a9351b440d09a67530fb3dc4508cd3e405b0ad9bb4806","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"stale","flag":"json"},"kwargs":"sha256:b11a0debf1dfcdad479d94d5d496266271a76735e1c014a31455e3dfa7363959","body":"sha256:d039c1f2caf44bc7dd84431a8617dd2453d265a0337b351f1488a67d6da3b9ce"}
	cmd.Flags().Bool("json", false, "Print as JSON")
	// /sedum:addFlag:bool
	// sedum:addFlag:bool {"v":2,"tier":"owned","record":"prov-2026-f476b640","gen":"sha256:24167071b71953f56b7a9351b440d09a67530fb3dc4508cd3e405b0ad9bb4806","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"stale","flag":"exit-code"},"kwargs":"sha256:5e0580b5f2fdd83707308c67908b4a77ed38a18471d8980cfbac31b9b9aa8f20","body":"sha256:bb984f1d1a903f2b372f2511ea1a1ab7f2743463e94c2a38460e3f7d7d27e5c1"}
	cmd.Flags().Bool("exit-code", false, "Exit non-zero when any record is directly or transitively stale")
	// /sedum:addFlag:bool
	// sedum:anchor:subcommands
	return cmd
}
