package main

import "github.com/spf13/cobra"

func newReviseCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"v":2,"tier":"owned","record":"prov-2026-5f69977b","gen":"sha256:8b7b52e1252d8b9aa1024924c8dc2aac14d21148579ea2b89a1653897fffab8a","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","id":{"bin":"provenance","command":"revise"},"kwargs":"sha256:8f43a3e4335cc2622ee3bcd26b645223b9805a29e9523a6a1dad4f978c75df9a","body":"sha256:cd712fb242c4fe871ecb82943f2acf82419fe56a55d78cd407f92f8e945a2735"}
	cmd.Use = "revise [id]"
	cmd.Short = "Edit a record's content in $EDITOR and save it as a new version"
	// /sedum:describeCommand
	// sedum:setRun {"v":2,"tier":"owned","record":"prov-2026-5f69977b","gen":"sha256:8b7b52e1252d8b9aa1024924c8dc2aac14d21148579ea2b89a1653897fffab8a","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","id":{"bin":"provenance","command":"revise"},"kwargs":"sha256:cfec1d8d1c6288021f821c41bee3664956aed03a3cf2fded8708a654b312ef27","body":"sha256:eb2f03eda5b54e595928705deb1b3677c11c749f1db60024d29f302ec942822d"}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runRevise(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:max_one {"v":2,"tier":"owned","record":"prov-2026-5f69977b","gen":"sha256:8b7b52e1252d8b9aa1024924c8dc2aac14d21148579ea2b89a1653897fffab8a","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","id":{"bin":"provenance","command":"revise"},"kwargs":"sha256:f1ee85c2568df3e265ecf81f095716a0b69643a685b98cf3961732c9e0a7f2ff","body":"sha256:186046a2435bbf4d0eae67e61e8a3909c91fd72c18b73dd2ff749798e3517ef3"}
	cmd.Args = cobra.MaximumNArgs(1)
	// /sedum:setArgs:max_one
	// sedum:anchor:flags
	// sedum:addFlag:string {"v":2,"tier":"owned","record":"prov-2026-5f69977b","gen":"sha256:8b7b52e1252d8b9aa1024924c8dc2aac14d21148579ea2b89a1653897fffab8a","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"revise","flag":"record"},"kwargs":"sha256:9d665d9730cfd85d1187d0bbece9b66931f10e27bc482bf04dfdcae86f6abf1a","body":"sha256:052950afb969aa0ebbee6b8db0650300a6c2fd2a9bef8bec15ff2c0d6dc81518"}
	cmd.Flags().String("record", "", "Record to revise (or give it as the argument)")
	// /sedum:addFlag:string
	// sedum:addFlag:string {"v":2,"tier":"owned","record":"prov-2026-5f69977b","gen":"sha256:8b7b52e1252d8b9aa1024924c8dc2aac14d21148579ea2b89a1653897fffab8a","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"revise","flag":"reason"},"kwargs":"sha256:c2f35441d7f7ce57b7313b381f8b0d2163d4258e4cb48f60f2421ed4ba5dcd36","body":"sha256:1782ad1d7e46e74b0540daee6d858c6e7774ab94014119b47df3727d579a2bbd"}
	cmd.Flags().String("reason", "", "Why this version exists (required for a new version)")
	// /sedum:addFlag:string
	// sedum:addFlag:bool {"v":2,"tier":"owned","record":"prov-2026-5f69977b","gen":"sha256:8b7b52e1252d8b9aa1024924c8dc2aac14d21148579ea2b89a1653897fffab8a","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"revise","flag":"fix"},"kwargs":"sha256:64f7971f99d8ca76987b2f62dccd8d927878f19a4e96dcec0bbb7a6b7cfb3378","body":"sha256:d2d89233f213edf28f7f946dbe76e7043dbde41d346537552e2bfd7bd449b4e8"}
	cmd.Flags().Bool("fix", false, "Mark this version as a fix (reason_kind: fix)")
	// /sedum:addFlag:bool
	// sedum:anchor:subcommands
	return cmd
}
