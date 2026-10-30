package main

import "github.com/spf13/cobra"

func newAddScopeCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"v":2,"tier":"owned","record":"prov-2026-9c8a0b2a","gen":"sha256:9c65c23d7893e39a1725f1f00b3344e10ee8e1823f7702b3722d3bef7130ba0c","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","id":{"bin":"provenance","command":"add-scope"},"kwargs":"sha256:e9a6ce894c7be244f67c4c75ae328f30e9c22532ea2d820d4241d49f3c72af9a","body":"sha256:953b25b9e99461acb2efc6da391d26b3e008d76f9d5005834fb217f15b3951fa"}
	cmd.Use = "add-scope"
	cmd.Short = "Widen an open record's scope with the files committed under it"
	// /sedum:describeCommand
	// sedum:setRun {"v":2,"tier":"owned","record":"prov-2026-9c8a0b2a","gen":"sha256:9c65c23d7893e39a1725f1f00b3344e10ee8e1823f7702b3722d3bef7130ba0c","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","id":{"bin":"provenance","command":"add-scope"},"kwargs":"sha256:6177d19c9717aaf86087f74a48007d50c9005f8526a4b2f77c1b2d6e95c89318","body":"sha256:a3a3dbc686e22b1f901b38f666bcc5dd5439a26bc72b605fbed943cbf362abf3"}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runAddScope(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:none {"v":2,"tier":"owned","record":"prov-2026-9c8a0b2a","gen":"sha256:9c65c23d7893e39a1725f1f00b3344e10ee8e1823f7702b3722d3bef7130ba0c","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","id":{"bin":"provenance","command":"add-scope"},"kwargs":"sha256:fbf99668915f7ed23ea22c06acd728f331bebe05e45970f6bcac950112d62e8d","body":"sha256:778b1e3eb6004fe613ecb8793efaf401a29f1afdc382707cea55c05601aefd79"}
	cmd.Args = cobra.NoArgs
	// /sedum:setArgs:none
	// sedum:anchor:flags
	// sedum:addFlag:string {"v":2,"tier":"owned","record":"prov-2026-9c8a0b2a","gen":"sha256:9c65c23d7893e39a1725f1f00b3344e10ee8e1823f7702b3722d3bef7130ba0c","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"add-scope","flag":"record"},"kwargs":"sha256:26340f390a9c0dedb6116c4c45be4bd89db06f73fdefd9cf7a50572ae558fa6d","body":"sha256:b40cfc50c55673b60cffccb47fac162a27c45c803c5e3ff555062044cd66cd20"}
	cmd.Flags().String("record", "", "Required. The record to widen")
	// /sedum:addFlag:string
	// sedum:addFlag:bool {"v":2,"tier":"owned","record":"prov-2026-9c8a0b2a","gen":"sha256:9c65c23d7893e39a1725f1f00b3344e10ee8e1823f7702b3722d3bef7130ba0c","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"add-scope","flag":"dry-run"},"kwargs":"sha256:e2def4ec9d6a9bc5713a57d2c73f9a702df55b8456f7ec499d64005e646f564a","body":"sha256:5922d44107215ebc51314f9d2047d803ce98eba1a12991b9e615a20a9c6ca425"}
	cmd.Flags().Bool("dry-run", false, "Print the files that would be added without writing")
	// /sedum:addFlag:bool
	// sedum:addFlag:string {"v":2,"tier":"owned","record":"prov-2026-9c8a0b2a","gen":"sha256:9c65c23d7893e39a1725f1f00b3344e10ee8e1823f7702b3722d3bef7130ba0c","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"add-scope","flag":"reason"},"kwargs":"sha256:ee1812078ee42e2c22105e343aefe64d00085c17aff63c631f7ad27f457c306f","body":"sha256:8c83d2e1ea6f1b09a6b5ac65472af50856ec7323e6bc96dabb6ae6a6d1f0de80"}
	cmd.Flags().String("reason", "", "Why this version exists (default: add-scope: files committed under <id>)")
	// /sedum:addFlag:string
	// sedum:anchor:subcommands
	return cmd
}
