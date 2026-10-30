package main

import "github.com/spf13/cobra"

func newInstallSkillsCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"v":2,"tier":"owned","record":"prov-2026-d71d0066","gen":"sha256:cae7a2d3a080d43f1ba12537ce93365db84c39195d99ec904268e9c15daab6a9","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","id":{"bin":"provenance","command":"install-skills"},"kwargs":"sha256:912f7b3e1744bff66c7488a880db1b6cf433a717c5c48c36a51d63850645329a","body":"sha256:f2f2fac6416e5a1fb8ef509c6d2ba67bb1b82c20c993782cb526cca14ff5d173"}
	cmd.Use = "install-skills"
	cmd.Short = "Install the Claude Code provenance skill into the repository"
	// /sedum:describeCommand
	// sedum:setRun {"v":2,"tier":"owned","record":"prov-2026-d71d0066","gen":"sha256:cae7a2d3a080d43f1ba12537ce93365db84c39195d99ec904268e9c15daab6a9","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","id":{"bin":"provenance","command":"install-skills"},"kwargs":"sha256:19eaf2c9b5ac0bdfab63f51a29db850be48c22487cb56ea8221a3172d13de22a","body":"sha256:628fa2aa3247fd5f3644489e97aa2058b72ab2001a19f4685b181dec8e6685b0"}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runInstallSkills(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:none {"v":2,"tier":"owned","record":"prov-2026-d71d0066","gen":"sha256:cae7a2d3a080d43f1ba12537ce93365db84c39195d99ec904268e9c15daab6a9","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","id":{"bin":"provenance","command":"install-skills"},"kwargs":"sha256:189f47cf45dd3d9155feb9c05e6bb974aca6c7016801a01cf3f2018f346e141f","body":"sha256:778b1e3eb6004fe613ecb8793efaf401a29f1afdc382707cea55c05601aefd79"}
	cmd.Args = cobra.NoArgs
	// /sedum:setArgs:none
	// sedum:anchor:flags
	// sedum:addFlag:string {"v":2,"tier":"owned","record":"prov-2026-d71d0066","gen":"sha256:cae7a2d3a080d43f1ba12537ce93365db84c39195d99ec904268e9c15daab6a9","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"install-skills","flag":"path"},"kwargs":"sha256:037358be6c5f670fa31f0fcba549203ae9b7d9f4c819ed22487fb1df168fa919","body":"sha256:6d3075776bdbef5e756d7f07741094adc66af56a8eb0e82cf208ee39cbc2862f"}
	cmd.Flags().String("path", "", "Target directory relative to the repository root (default: .claude/skills)")
	// /sedum:addFlag:string
	// sedum:addFlag:bool {"v":2,"tier":"owned","record":"prov-2026-d71d0066","gen":"sha256:cae7a2d3a080d43f1ba12537ce93365db84c39195d99ec904268e9c15daab6a9","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"install-skills","flag":"force"},"kwargs":"sha256:be58ca897d636cc5e4eb18a689e7a401c33a4247b4d50354cc9b236d20f04a21","body":"sha256:d9034ed699def2ce79a49ab206ac85c2439e521ec8468e054a5d469c9b8f05e1"}
	cmd.Flags().Bool("force", false, "Overwrite existing files that differ")
	// /sedum:addFlag:bool
	// sedum:anchor:subcommands
	return cmd
}
