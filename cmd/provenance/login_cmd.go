package main

import "github.com/spf13/cobra"

func newLoginCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"v":2,"tier":"owned","record":"prov-2026-635b25ac","gen":"sha256:1174732e27a95ca5695b7d4d1dad8b7063f0e7895a238818a555e8756ca8880f","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","id":{"bin":"provenance","command":"login"},"kwargs":"sha256:bcd2891e50e19ae301d02d35de15abe9827737a374904960f01ab009f3af4d8d","body":"sha256:7245d9a2ec513b20e9595bcf1b69a351be71719a113375defe66442509c6acab"}
	cmd.Use = "login"
	cmd.Short = "Sign in to Provisio with a device code"
	// /sedum:describeCommand
	// sedum:setRun {"v":2,"tier":"owned","record":"prov-2026-635b25ac","gen":"sha256:1174732e27a95ca5695b7d4d1dad8b7063f0e7895a238818a555e8756ca8880f","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","id":{"bin":"provenance","command":"login"},"kwargs":"sha256:e1156eb1407a559e28bb114a95d2636f4a261bde5bd2a58041f7ca8ddfd21c48","body":"sha256:b245b1ce931f922ca9fac9bb87259efff26a7dff2c48e2eb3e84d940c300fa07"}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runLogin(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:none {"v":2,"tier":"owned","record":"prov-2026-635b25ac","gen":"sha256:1174732e27a95ca5695b7d4d1dad8b7063f0e7895a238818a555e8756ca8880f","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","id":{"bin":"provenance","command":"login"},"kwargs":"sha256:5b6894acb165928a630285e227a1c154bc28424eb898870208efa50c09da8ba4","body":"sha256:778b1e3eb6004fe613ecb8793efaf401a29f1afdc382707cea55c05601aefd79"}
	cmd.Args = cobra.NoArgs
	// /sedum:setArgs:none
	// sedum:anchor:flags
	// sedum:addFlag:bool {"v":2,"tier":"owned","record":"prov-2026-635b25ac","gen":"sha256:1174732e27a95ca5695b7d4d1dad8b7063f0e7895a238818a555e8756ca8880f","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"login","flag":"no-browser"},"kwargs":"sha256:203a6b2aea8e0701bf1e017555fb99e12684e8ab3008ee70ee09a27570e11751","body":"sha256:33f2369b75866ef520f2f98d7f35223167d6a4b4505d13fa0cca23a20d0b7439"}
	cmd.Flags().Bool("no-browser", false, "Only print the verification URL and code; do not open a browser")
	// /sedum:addFlag:bool
	// sedum:addFlag:bool {"v":2,"tier":"owned","record":"prov-2026-635b25ac","gen":"sha256:1174732e27a95ca5695b7d4d1dad8b7063f0e7895a238818a555e8756ca8880f","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"login","flag":"with-token"},"kwargs":"sha256:854d32e8a2cc322d67e9f1b170b8c0f541ccb533404892925e7339a20ee8af9f","body":"sha256:7172ede72b48b697371f8017f7c341a49c13823b0ccd2d0e1e127b6855fb0a47"}
	cmd.Flags().Bool("with-token", false, "Read a personal token (pvt_...) from stdin and store it instead of signing in")
	// /sedum:addFlag:bool
	// sedum:addFlag:string {"v":2,"tier":"owned","record":"prov-2026-635b25ac","gen":"sha256:1174732e27a95ca5695b7d4d1dad8b7063f0e7895a238818a555e8756ca8880f","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"login","flag":"server"},"kwargs":"sha256:1d01cfd750688761fd46d6bf80c8866312744e1d85ad780a5e2926c7836db2c8","body":"sha256:ca16890f7d9aa172deca35bfcb21ddc06cf7e545ca1de48dc3f5b6b6735e14bd"}
	cmd.Flags().String("server", "", "Auth service URL (default $PROVISIO_AUTH_URL, else https://auth.provisio.dev)")
	// /sedum:addFlag:string
	// sedum:addFlag:string {"v":2,"tier":"owned","record":"prov-2026-635b25ac","gen":"sha256:1174732e27a95ca5695b7d4d1dad8b7063f0e7895a238818a555e8756ca8880f","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"login","flag":"provider"},"kwargs":"sha256:9f9f2de5036e7a32283574b1766269e95b115e080c395cf56b7c99f9d3a148fe","body":"sha256:e3e8d76f38da78d567f041b5525f50971892d39e9e9b51132d181817a396a6c3"}
	cmd.Flags().String("provider", "", "Sign-in hint: provisio or bluesky (default provisio)")
	// /sedum:addFlag:string
	// sedum:anchor:subcommands
	return cmd
}
