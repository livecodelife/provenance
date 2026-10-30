package main

import "github.com/spf13/cobra"

func newHooksCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"v":2,"tier":"owned","record":"prov-2026-d1d789de","gen":"sha256:6202ce5f10060b0fa61443fc52df17f80e57f28bdc0e09ce5ef4e680dc50ebff","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","id":{"bin":"provenance","command":"hooks"},"kwargs":"sha256:a0375f9a953e4c1799b510288912bf296234bc8cfb192dc159161106013eb530","body":"sha256:c89f2435eb3748b5bf3c5eb000cf994fbfcae767f56c916d6a2eab0e2be55947"}
	cmd.Use = "hooks"
	cmd.Short = "Manage the provenance git hooks"
	// /sedum:describeCommand
	// sedum:setArgs:none {"v":2,"tier":"owned","record":"prov-2026-d1d789de","gen":"sha256:6202ce5f10060b0fa61443fc52df17f80e57f28bdc0e09ce5ef4e680dc50ebff","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","id":{"bin":"provenance","command":"hooks"},"kwargs":"sha256:72f6c7f2e268660de43bb1c1bf664c9f160ed4f62decc01e4e5dd6cebe63e7ec","body":"sha256:778b1e3eb6004fe613ecb8793efaf401a29f1afdc382707cea55c05601aefd79"}
	cmd.Args = cobra.NoArgs
	// /sedum:setArgs:none
	// sedum:anchor:flags
	// sedum:anchor:subcommands
	// sedum:registerSubcommand {"v":2,"tier":"owned","record":"prov-2026-d1d789de","gen":"sha256:6202ce5f10060b0fa61443fc52df17f80e57f28bdc0e09ce5ef4e680dc50ebff","pkg":"sha256:8a9e0a8ee811e7c629264e26f09a626602890b11fd82710cf10ae73e3da1cf7a","id":{"bin":"provenance","command":"hooks-install","parent":"hooks"},"kwargs":"sha256:5c7df9e56d0eff4bf4d47634a2c5b49f763cc81e1110061c6e6a853698b20361","body":"sha256:a4290ec4500700213ae2a454393b6233022efb0d69661756e6c8e2f6e3e8a5c0"}
	cmd.AddCommand(newHooksInstallCmd())
	// /sedum:registerSubcommand
	// sedum:registerSubcommand {"v":2,"tier":"owned","record":"prov-2026-10bc3e89","gen":"sha256:e93ace3473e7af7dd8c92e4f660130b255ccb6bbefff31577a90dcb3982e21d4","pkg":"sha256:8a9e0a8ee811e7c629264e26f09a626602890b11fd82710cf10ae73e3da1cf7a","id":{"bin":"provenance","command":"hooks-pre-push","parent":"hooks"},"kwargs":"sha256:1d1922e6254096aa25ff02b119291cd0de0fa98eaea00a520b1eb75d2f087d0b","body":"sha256:97e4b8df4afe85056e8e31b5a5dff13a6de1d221f7cb7197c2d7d6c2bdc04254"}
	cmd.AddCommand(newHooksPrePushCmd())
	// /sedum:registerSubcommand
	return cmd
}
