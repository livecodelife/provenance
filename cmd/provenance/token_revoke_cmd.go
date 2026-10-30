package main

import "github.com/spf13/cobra"

func newTokenRevokeCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"v":2,"tier":"owned","record":"prov-2026-a8ec3e30","gen":"sha256:32397270f026f7499dfffc6c69c80d69185c2e1390ca78f0ee9ecba7837ebbdd","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","id":{"bin":"provenance","command":"token-revoke"},"kwargs":"sha256:91f1c309ec2788f234d8b0b335253670837ef8b96ded64e6152d23305a79c191","body":"sha256:6c333bdf631b97bb2188bbbcc3215b7713806555316febb22f787a5122df280d"}
	cmd.Use = "revoke <id>"
	cmd.Short = "Revoke a personal token"
	// /sedum:describeCommand
	// sedum:setRun {"v":2,"tier":"owned","record":"prov-2026-a8ec3e30","gen":"sha256:32397270f026f7499dfffc6c69c80d69185c2e1390ca78f0ee9ecba7837ebbdd","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","id":{"bin":"provenance","command":"token-revoke"},"kwargs":"sha256:573833b684d633484f0610ad187ddac8c8a1ec44a079b03c4abe03fcc8a84dbc","body":"sha256:e5f024826d622ba78ca1545d6ae57cc365daa7d05fe46608d63a31ca125bcc66"}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runTokenRevoke(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:one {"v":2,"tier":"owned","record":"prov-2026-a8ec3e30","gen":"sha256:32397270f026f7499dfffc6c69c80d69185c2e1390ca78f0ee9ecba7837ebbdd","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","id":{"bin":"provenance","command":"token-revoke"},"kwargs":"sha256:d341529baf092c2063a4a5fbf007d746a2f2193c6c3fe6ffd47bc5dbc8131362","body":"sha256:ba5f42dd9b6e7c8d7804e37fa0293cf76d4cb2f250b8778e4a36b12c7c17dd95"}
	cmd.Args = cobra.ExactArgs(1)
	// /sedum:setArgs:one
	// sedum:anchor:flags
	// sedum:anchor:subcommands
	return cmd
}
