package main

import "github.com/spf13/cobra"

func newDeprecateCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"v":2,"tier":"owned","record":"prov-2026-35439261","gen":"sha256:bb73fb8a27febfb6636afc91c6f46ef087ceb45d41e95db14a66c7469f2256d4","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","id":{"bin":"provenance","command":"deprecate"},"kwargs":"sha256:e11a40f831f40df4547776041da3786cc58f8a74544ddd615bebdca42c279d8e","body":"sha256:7856fdc8165bde6c130aa5ee58eab61015ed687f6113b74a355b8d2aef91bd8d"}
	cmd.Use = "deprecate"
	cmd.Short = "Mark a record as deprecated (terminal: it takes no new versions)"
	// /sedum:describeCommand
	// sedum:setRun {"v":2,"tier":"owned","record":"prov-2026-35439261","gen":"sha256:bb73fb8a27febfb6636afc91c6f46ef087ceb45d41e95db14a66c7469f2256d4","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","id":{"bin":"provenance","command":"deprecate"},"kwargs":"sha256:88ddf979fa7ecf3af2be46dba0b7f74f8715cd7fad98af43b3f47ac2142f3ad1","body":"sha256:22f350155913c373b30a417fb70c2829d543e9eed1ec0b2fd1a5a03be4c2edd2"}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runDeprecate(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:none {"v":2,"tier":"owned","record":"prov-2026-35439261","gen":"sha256:bb73fb8a27febfb6636afc91c6f46ef087ceb45d41e95db14a66c7469f2256d4","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","id":{"bin":"provenance","command":"deprecate"},"kwargs":"sha256:6516d4c3228c43515df78e5b4e5a7dc633d1c89c5597219a6769708d8903124f","body":"sha256:778b1e3eb6004fe613ecb8793efaf401a29f1afdc382707cea55c05601aefd79"}
	cmd.Args = cobra.NoArgs
	// /sedum:setArgs:none
	// sedum:anchor:flags
	// sedum:addFlag:string {"v":2,"tier":"owned","record":"prov-2026-35439261","gen":"sha256:bb73fb8a27febfb6636afc91c6f46ef087ceb45d41e95db14a66c7469f2256d4","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"deprecate","flag":"record"},"kwargs":"sha256:4a18a6c71f91cf73be1ce5d04d4bb2a41c2bb7a2e0a7e4f2fdb3f370b47abf07","body":"sha256:e503b0d560a83a70fe309b3f05b6116005b711a61aa4cbc21d2521609bf84925"}
	cmd.Flags().String("record", "", "Record to deprecate")
	// /sedum:addFlag:string
	// sedum:addFlag:string {"v":2,"tier":"owned","record":"prov-2026-35439261","gen":"sha256:bb73fb8a27febfb6636afc91c6f46ef087ceb45d41e95db14a66c7469f2256d4","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"deprecate","flag":"reason"},"kwargs":"sha256:3ec46bc2a082020c7863b4146daada693b8d23c6c534355943dff4859aecd0cf","body":"sha256:e8978a26ae443690749665c9b19b09a941d463e0bff53631f387ed9be1dcf14a"}
	cmd.Flags().String("reason", "", "Why the record is deprecated (recorded in the commit message)")
	// /sedum:addFlag:string
	// sedum:anchor:subcommands
	return cmd
}
