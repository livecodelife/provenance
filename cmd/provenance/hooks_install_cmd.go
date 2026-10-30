package main

import "github.com/spf13/cobra"

func newHooksInstallCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"v":2,"tier":"owned","record":"prov-2026-d1d789de","gen":"sha256:6202ce5f10060b0fa61443fc52df17f80e57f28bdc0e09ce5ef4e680dc50ebff","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","id":{"bin":"provenance","command":"hooks-install"},"kwargs":"sha256:72b76d31b5ff70e0ddba9bfb32e2ecd7702d370f46626d23b4cfb22ea7c6c6ee","body":"sha256:4947331aa1ae2e4dd29b489e61d4446a97855375a324083fc245420ed6dc3b64"}
	cmd.Use = "install"
	cmd.Short = "Install the pre-commit and commit-msg hooks (existing hooks are kept as .bak)"
	// /sedum:describeCommand
	// sedum:setRun {"v":2,"tier":"owned","record":"prov-2026-d1d789de","gen":"sha256:6202ce5f10060b0fa61443fc52df17f80e57f28bdc0e09ce5ef4e680dc50ebff","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","id":{"bin":"provenance","command":"hooks-install"},"kwargs":"sha256:db8e729d9bb217241a019e1da509ed2ea71cdfb800ba67ddda13c3f203ede09f","body":"sha256:1275aac5b9e2533e265f9104a4a7241dfd818a995ab5ca7113e78a10a05bf72c"}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runHooksInstall(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:none {"v":2,"tier":"owned","record":"prov-2026-d1d789de","gen":"sha256:6202ce5f10060b0fa61443fc52df17f80e57f28bdc0e09ce5ef4e680dc50ebff","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","id":{"bin":"provenance","command":"hooks-install"},"kwargs":"sha256:db2858bb08e04e5dc08a5161987cd1fb7ad6765c19f624ee5e7856e858a481a2","body":"sha256:778b1e3eb6004fe613ecb8793efaf401a29f1afdc382707cea55c05601aefd79"}
	cmd.Args = cobra.NoArgs
	// /sedum:setArgs:none
	// sedum:anchor:flags
	// sedum:anchor:subcommands
	return cmd
}
