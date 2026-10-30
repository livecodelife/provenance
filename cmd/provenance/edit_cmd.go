package main

import "github.com/spf13/cobra"

func newEditCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"v":2,"tier":"owned","record":"prov-2026-57b8a95b","gen":"sha256:ecce14df038164301d4c4728b274460b7a56a8972e302aacf33c69ff352be339","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","id":{"bin":"provenance","command":"edit"},"kwargs":"sha256:2d2809f9cbdcb4880098ec9d5b268c9e93172c4443b5c3e9340d187b01fd0f44","body":"sha256:83ab15ad14e18656c2389898f0571c8d25c0bce86b54c0c8279b4c1cab574235"}
	cmd.Use = "edit [id]"
	cmd.Short = "Change a record's content fields without an editor and save a new version"
	// /sedum:describeCommand
	// sedum:setRun {"v":2,"tier":"owned","record":"prov-2026-57b8a95b","gen":"sha256:ecce14df038164301d4c4728b274460b7a56a8972e302aacf33c69ff352be339","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","id":{"bin":"provenance","command":"edit"},"kwargs":"sha256:e7a4b6073ea9164a358b8fa6132903dec3f2b70ef55af18ce7bc6a72083c3820","body":"sha256:4db603288b2159a003262d32c941071220b01c7a684aad585ac72e2b6642b3d7"}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runEdit(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:max_one {"v":2,"tier":"owned","record":"prov-2026-57b8a95b","gen":"sha256:ecce14df038164301d4c4728b274460b7a56a8972e302aacf33c69ff352be339","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","id":{"bin":"provenance","command":"edit"},"kwargs":"sha256:f41a3aa4395af496218c8f6b6c495fd2b542beea41d9ffe152327330efb13464","body":"sha256:186046a2435bbf4d0eae67e61e8a3909c91fd72c18b73dd2ff749798e3517ef3"}
	cmd.Args = cobra.MaximumNArgs(1)
	// /sedum:setArgs:max_one
	// sedum:anchor:flags
	// sedum:addFlag:string {"v":2,"tier":"owned","record":"prov-2026-57b8a95b","gen":"sha256:ecce14df038164301d4c4728b274460b7a56a8972e302aacf33c69ff352be339","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"edit","flag":"record"},"kwargs":"sha256:a1de3a402f73d323a67163273a30de65ca701cb3a5396868852af604e72bc63c","body":"sha256:5ec7e9882a1013462147983624156c68d20f02ab341d4fcb320e0469d0abb23d"}
	cmd.Flags().String("record", "", "Record to edit (or give it as the argument)")
	// /sedum:addFlag:string
	// sedum:addFlag:string {"v":2,"tier":"owned","record":"prov-2026-57b8a95b","gen":"sha256:ecce14df038164301d4c4728b274460b7a56a8972e302aacf33c69ff352be339","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"edit","flag":"title"},"kwargs":"sha256:ce1375e0af64b8d75ad07f3339e37ef63d19369f84648a614ee6baa67a203ee6","body":"sha256:264b11a63b5f37afa614ed7f456793456f661d840e731e0833882e272e036f98"}
	cmd.Flags().String("title", "", "New title")
	// /sedum:addFlag:string
	// sedum:addFlag:string {"v":2,"tier":"owned","record":"prov-2026-57b8a95b","gen":"sha256:ecce14df038164301d4c4728b274460b7a56a8972e302aacf33c69ff352be339","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"edit","flag":"intent"},"kwargs":"sha256:5084186c2e521e3cfa8f959bb1b1e9f2db486146000d4603e680ee69a3b8e1d7","body":"sha256:f38b49e598ceb40fead5ae00e3c52f93ac903a967bbf1ee0758534ea97527070"}
	cmd.Flags().String("intent", "", "New intent")
	// /sedum:addFlag:string
	// sedum:addFlag:strings {"v":2,"tier":"owned","record":"prov-2026-57b8a95b","gen":"sha256:ecce14df038164301d4c4728b274460b7a56a8972e302aacf33c69ff352be339","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"edit","flag":"add-constraint"},"kwargs":"sha256:0999f895b5055fbe0535a0c160e7f8761fdd5f40beb7ea34d32138b8074232e8","body":"sha256:ae0aca961790c1496278b72d1b87578db78d4d79b43419db23f222a52a5e6e66"}
	cmd.Flags().StringArray("add-constraint", nil, "Append a constraint (repeatable)")
	// /sedum:addFlag:strings
	// sedum:addFlag:strings {"v":2,"tier":"owned","record":"prov-2026-57b8a95b","gen":"sha256:ecce14df038164301d4c4728b274460b7a56a8972e302aacf33c69ff352be339","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"edit","flag":"remove-constraint"},"kwargs":"sha256:331c75fff019e8ab5f633f353e05bb0b294e10199cba5f4009816da0d2c9c77d","body":"sha256:3e882adf663d55d23506389706e7a2dbc93670112f4702162db73e43d1917328"}
	cmd.Flags().StringArray("remove-constraint", nil, "Remove constraint number n, counting from 1 (repeatable)")
	// /sedum:addFlag:strings
	// sedum:addFlag:strings {"v":2,"tier":"owned","record":"prov-2026-57b8a95b","gen":"sha256:ecce14df038164301d4c4728b274460b7a56a8972e302aacf33c69ff352be339","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"edit","flag":"add-scope"},"kwargs":"sha256:621d345cb171e45bc789c7dad83f4fd41e6ad3b115761731a275394163aa87a4","body":"sha256:32925a1382196ae0e5ca22a4e73b6128959256384a4705b5f4633aa48722acfa"}
	cmd.Flags().StringArray("add-scope", nil, "Append an affected_scope entry (repeatable)")
	// /sedum:addFlag:strings
	// sedum:addFlag:strings {"v":2,"tier":"owned","record":"prov-2026-57b8a95b","gen":"sha256:ecce14df038164301d4c4728b274460b7a56a8972e302aacf33c69ff352be339","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"edit","flag":"remove-scope"},"kwargs":"sha256:6524fb9aabe75f5c0bf1d16ec0f76d22c68aa793a4057978ff8a2097ec25a5b8","body":"sha256:75b61b46735c8435160666b85b4bd4beaa98d8a0a474a799666504f37211f083"}
	cmd.Flags().StringArray("remove-scope", nil, "Remove an affected_scope entry by value (repeatable)")
	// /sedum:addFlag:strings
	// sedum:addFlag:strings {"v":2,"tier":"owned","record":"prov-2026-57b8a95b","gen":"sha256:ecce14df038164301d4c4728b274460b7a56a8972e302aacf33c69ff352be339","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"edit","flag":"add-forbidden"},"kwargs":"sha256:4647c5ea46cbd715b2ef073918b5dfb98667c167e8d475bb9bb3a193a167e1dc","body":"sha256:8e66614b5bf346218d893765705c84fcb75207ff729131056258c9f34588f509"}
	cmd.Flags().StringArray("add-forbidden", nil, "Append a forbidden_scope entry (repeatable)")
	// /sedum:addFlag:strings
	// sedum:addFlag:strings {"v":2,"tier":"owned","record":"prov-2026-57b8a95b","gen":"sha256:ecce14df038164301d4c4728b274460b7a56a8972e302aacf33c69ff352be339","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"edit","flag":"remove-forbidden"},"kwargs":"sha256:12c8a4b2a0708588db5f0a43d55203946d4689bd040a7d564fa16ce4ef5aa6db","body":"sha256:7955c5c30fc445ef7eb66763f7af79f13d37cb1dc3709444e743eb4e28749a44"}
	cmd.Flags().StringArray("remove-forbidden", nil, "Remove a forbidden_scope entry by value (repeatable)")
	// /sedum:addFlag:strings
	// sedum:addFlag:string {"v":2,"tier":"owned","record":"prov-2026-57b8a95b","gen":"sha256:ecce14df038164301d4c4728b274460b7a56a8972e302aacf33c69ff352be339","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"edit","flag":"reason"},"kwargs":"sha256:95be052d00e17a1f49929ed7d899511041f46f0491ea821f9d5a5eb61c1e42a1","body":"sha256:1782ad1d7e46e74b0540daee6d858c6e7774ab94014119b47df3727d579a2bbd"}
	cmd.Flags().String("reason", "", "Why this version exists (required for a new version)")
	// /sedum:addFlag:string
	// sedum:addFlag:bool {"v":2,"tier":"owned","record":"prov-2026-57b8a95b","gen":"sha256:ecce14df038164301d4c4728b274460b7a56a8972e302aacf33c69ff352be339","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"edit","flag":"fix"},"kwargs":"sha256:c9da288c0fbf7f00db4b944c897011855e8d217617087e8b0a17d765b1349000","body":"sha256:d2d89233f213edf28f7f946dbe76e7043dbde41d346537552e2bfd7bd449b4e8"}
	cmd.Flags().Bool("fix", false, "Mark this version as a fix (reason_kind: fix)")
	// /sedum:addFlag:bool
	// sedum:addFlag:bool {"v":2,"tier":"owned","record":"prov-2026-57b8a95b","gen":"sha256:ecce14df038164301d4c4728b274460b7a56a8972e302aacf33c69ff352be339","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"edit","flag":"refresh-specs"},"kwargs":"sha256:4c0e4c8f5f66c6da50c96b728a573c5f4a9d645e9f85cadf56924b390b09783a","body":"sha256:a47729fdd98113e50c88489aaf61d009a90edda4fb418f409667ccd9235d5dc4"}
	cmd.Flags().Bool("refresh-specs", false, "Record that the associated specs' content changed on disk (a new version; needs --reason)")
	// /sedum:addFlag:bool
	// sedum:addFlag:strings {"v":2,"tier":"owned","record":"prov-2026-57b8a95b","gen":"sha256:ecce14df038164301d4c4728b274460b7a56a8972e302aacf33c69ff352be339","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"edit","flag":"set-spec"},"kwargs":"sha256:b45734c11c73064334a162007e028cadedf58e7d4aa886cbc99b391d2f835634","body":"sha256:55cfa61656f88506218f606a670e499ef51e9f17ce10b07576995996c11176ac"}
	cmd.Flags().StringArray("set-spec", nil, "Change an associated spec: <old>=<new> or <old>=<new>:<run_command> (repeatable)")
	// /sedum:addFlag:strings
	// sedum:addFlag:strings {"v":2,"tier":"owned","record":"prov-2026-57b8a95b","gen":"sha256:ecce14df038164301d4c4728b274460b7a56a8972e302aacf33c69ff352be339","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"edit","flag":"add-spec"},"kwargs":"sha256:c957337a0b3e7f14cca656eb6f48f0def99015c861dc5c88ba8b7609aa0c77f5","body":"sha256:0e581dd40c2477914dded542f80d49699466e99be9708352b8b25e2f7a2b4d9a"}
	cmd.Flags().StringArray("add-spec", nil, "Add an associated spec: <path> or <path>:<run_command> (repeatable; the file must exist)")
	// /sedum:addFlag:strings
	// sedum:addFlag:strings {"v":2,"tier":"owned","record":"prov-2026-57b8a95b","gen":"sha256:ecce14df038164301d4c4728b274460b7a56a8972e302aacf33c69ff352be339","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"edit","flag":"remove-spec"},"kwargs":"sha256:da9267a057e0fbf983de00c356890ea368da2fd21a5648aa418588f8a165fa36","body":"sha256:37af42f5d86117c8d1990fa8b6b265b96e5f079026f0bc945da67192d5ab787f"}
	cmd.Flags().StringArray("remove-spec", nil, "Remove an associated spec by path (repeatable)")
	// /sedum:addFlag:strings
	// sedum:anchor:subcommands
	return cmd
}
