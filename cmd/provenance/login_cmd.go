package main

import "github.com/spf13/cobra"

func newLoginCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"tier":"owned","record":"prov-2026-635b25ac","gen":"sha256:3dd07349442894e371a607dad8f67454397b701f2bd4e8d8d0edfb17be96c5a0","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","kwargs":{"bin":"provenance","command":"login","short":"Sign in to Provisio with a device code","use":"login"}}
	cmd.Use = "login"
	cmd.Short = "Sign in to Provisio with a device code"
	// /sedum:describeCommand
	// sedum:setRun {"tier":"owned","record":"prov-2026-635b25ac","gen":"sha256:3dd07349442894e371a607dad8f67454397b701f2bd4e8d8d0edfb17be96c5a0","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","kwargs":{"bin":"provenance","command":"login"}}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runLogin(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:none {"tier":"owned","record":"prov-2026-635b25ac","gen":"sha256:3dd07349442894e371a607dad8f67454397b701f2bd4e8d8d0edfb17be96c5a0","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","kwargs":{"arity":"none","bin":"provenance","command":"login"}}
	cmd.Args = cobra.NoArgs
	// /sedum:setArgs:none
	// sedum:anchor:flags
	// sedum:addFlag:bool {"tier":"owned","record":"prov-2026-635b25ac","gen":"sha256:3dd07349442894e371a607dad8f67454397b701f2bd4e8d8d0edfb17be96c5a0","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"login","flag":"no-browser","type":"bool","usage":"Only print the verification URL and code; do not open a browser"}}
	cmd.Flags().Bool("no-browser", false, "Only print the verification URL and code; do not open a browser")
	// /sedum:addFlag:bool
	// sedum:addFlag:bool {"tier":"owned","record":"prov-2026-635b25ac","gen":"sha256:3dd07349442894e371a607dad8f67454397b701f2bd4e8d8d0edfb17be96c5a0","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"login","flag":"with-token","type":"bool","usage":"Read a personal token (pvt_...) from stdin and store it instead of signing in"}}
	cmd.Flags().Bool("with-token", false, "Read a personal token (pvt_...) from stdin and store it instead of signing in")
	// /sedum:addFlag:bool
	// sedum:addFlag:string {"tier":"owned","record":"prov-2026-635b25ac","gen":"sha256:3dd07349442894e371a607dad8f67454397b701f2bd4e8d8d0edfb17be96c5a0","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"login","flag":"server","type":"string","usage":"Auth service URL (default $PROVISIO_AUTH_URL, else https://auth.provisio.dev)"}}
	cmd.Flags().String("server", "", "Auth service URL (default $PROVISIO_AUTH_URL, else https://auth.provisio.dev)")
	// /sedum:addFlag:string
	// sedum:addFlag:string {"tier":"owned","record":"prov-2026-635b25ac","gen":"sha256:3dd07349442894e371a607dad8f67454397b701f2bd4e8d8d0edfb17be96c5a0","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","kwargs":{"bin":"provenance","command":"login","flag":"provider","type":"string","usage":"Sign-in hint: provisio or bluesky (default provisio)"}}
	cmd.Flags().String("provider", "", "Sign-in hint: provisio or bluesky (default provisio)")
	// /sedum:addFlag:string
	// sedum:anchor:subcommands
	return cmd
}
