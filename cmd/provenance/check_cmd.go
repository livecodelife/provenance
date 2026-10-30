package main

import "github.com/spf13/cobra"

func newCheckCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"v":2,"tier":"owned","record":"prov-2026-063f1395","gen":"sha256:4239f75c5db7efc2453be01f153d0012da926ff1bd1870547582214fbf406657","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","id":{"bin":"provenance","command":"check"},"kwargs":"sha256:f5a31f367ddc14290720fa3998c3cfbeb1469050911aaba837c8953a4403d6bb","body":"sha256:09fa6e7c59e15ef5eac2c26eb0396b36f75b1a96b1001c53689ead8d13f307e2"}
	cmd.Use = "check"
	cmd.Short = "Check staged changes against provenance records (the commit-msg hook)"
	// /sedum:describeCommand
	// sedum:setRun {"v":2,"tier":"owned","record":"prov-2026-063f1395","gen":"sha256:4239f75c5db7efc2453be01f153d0012da926ff1bd1870547582214fbf406657","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","id":{"bin":"provenance","command":"check"},"kwargs":"sha256:170ac74438ebb2fec298dfb7ee87090d28eb803209cc56fe7e1c62a4d971f2b1","body":"sha256:020ab3581e48cbc2fa08a8b2eab2ac9bc81dd19e5ff0583274187757e68c0158"}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runCheck(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:none {"v":2,"tier":"owned","record":"prov-2026-063f1395","gen":"sha256:4239f75c5db7efc2453be01f153d0012da926ff1bd1870547582214fbf406657","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","id":{"bin":"provenance","command":"check"},"kwargs":"sha256:67af3ad9c54e972bd9425ea299a5bffb770b3e38a349da9744b6191e24bb18ec","body":"sha256:778b1e3eb6004fe613ecb8793efaf401a29f1afdc382707cea55c05601aefd79"}
	cmd.Args = cobra.NoArgs
	// /sedum:setArgs:none
	// sedum:anchor:flags
	// sedum:addFlag:bool {"v":2,"tier":"owned","record":"prov-2026-063f1395","gen":"sha256:4239f75c5db7efc2453be01f153d0012da926ff1bd1870547582214fbf406657","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"check","flag":"staged"},"kwargs":"sha256:360f9b484e0239fff0dc5baaf6cd500c6b483247a1b5314ebe8a0201b9c5633f","body":"sha256:468a25eb65a009df445ec4df16326c51b0368545dc8a8c0f21a775f312e54a11"}
	cmd.Flags().Bool("staged", false, "Check the staged changes")
	// /sedum:addFlag:bool
	// sedum:addFlag:string {"v":2,"tier":"owned","record":"prov-2026-063f1395","gen":"sha256:4239f75c5db7efc2453be01f153d0012da926ff1bd1870547582214fbf406657","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"check","flag":"message-file"},"kwargs":"sha256:806fe37de90e5de4264d6faf3e38b6c45218c66ea1b0139602445b13ed479664","body":"sha256:fe30de93f35ce118a57259b04916b952e399706a1c619c44172f7b393bb4ee33"}
	cmd.Flags().String("message-file", "", "Commit message file (as the commit-msg hook passes it)")
	// /sedum:addFlag:string
	// sedum:addFlag:string {"v":2,"tier":"owned","record":"prov-2026-063f1395","gen":"sha256:4239f75c5db7efc2453be01f153d0012da926ff1bd1870547582214fbf406657","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"check","flag":"commit"},"kwargs":"sha256:4457d1cad16e094ea8ce1f265a02f8f771a727e461a3cbe9955993a7b02f0f23","body":"sha256:7adc3692e8143a8e99345a189e9f92317803e193fb96850b0ce63511cb73dbd6"}
	cmd.Flags().String("commit", "", "Check one commit")
	// /sedum:addFlag:string
	// sedum:addFlag:string {"v":2,"tier":"owned","record":"prov-2026-063f1395","gen":"sha256:4239f75c5db7efc2453be01f153d0012da926ff1bd1870547582214fbf406657","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"check","flag":"range"},"kwargs":"sha256:21af246a8c565f71a6f7180f4440017144e8b0b7df2ff48e2a1f16601c6ff01b","body":"sha256:30f2753197d3b6d05b3c1a6d30c81ebdbb848bfbe74cd2afac3b22814038357d"}
	cmd.Flags().String("range", "", "Check every commit in a..b")
	// /sedum:addFlag:string
	// sedum:addFlag:bool {"v":2,"tier":"owned","record":"prov-2026-063f1395","gen":"sha256:4239f75c5db7efc2453be01f153d0012da926ff1bd1870547582214fbf406657","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"check","flag":"evidence"},"kwargs":"sha256:03846841351baf43b06d5d93f92866578738c4a428a17b8f83105339b02ab525","body":"sha256:a354c52c1c24594d5191aa484dc7b5aac536f7480815148ef0425a9e695f7b36"}
	cmd.Flags().Bool("evidence", false, "Report whether each implemented record's evidence is fresh")
	// /sedum:addFlag:bool
	// sedum:addFlag:bool {"v":2,"tier":"owned","record":"prov-2026-063f1395","gen":"sha256:4239f75c5db7efc2453be01f153d0012da926ff1bd1870547582214fbf406657","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"check","flag":"json"},"kwargs":"sha256:4455f333e958569472974f324641497ccc7f556d1d0202547fa7d62c7153d452","body":"sha256:bc2850b9493f25cad4f80d80f4760926ee03693386b99d5bfeb4350dafac0a1d"}
	cmd.Flags().Bool("json", false, "Print --evidence as JSON")
	// /sedum:addFlag:bool
	// sedum:addFlag:string {"v":2,"tier":"owned","record":"prov-2026-063f1395","gen":"sha256:4239f75c5db7efc2453be01f153d0012da926ff1bd1870547582214fbf406657","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"check","flag":"format"},"kwargs":"sha256:da01b608cd1dd412fc364c4924287e47e57059b81cc012aa4f2134798550c698","body":"sha256:63d879c2d8f9371f6d3c99a19a1bdadd015f93503dd836bbff4d2df26eeb3d29"}
	cmd.Flags().String("format", "", "Output format for violations and evidence (human|sarif)")
	// /sedum:addFlag:string
	// sedum:anchor:subcommands
	return cmd
}
