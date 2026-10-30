package main

import "github.com/spf13/cobra"

func newCreateCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"v":2,"tier":"owned","record":"prov-2026-44a07768","gen":"sha256:38d5d332ecbcb41d5a1af3b535e6cb5d9305d6ac235c5bb64e2cb02ccd551adc","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","id":{"bin":"provenance","command":"create"},"kwargs":"sha256:926b6771110c18f94263c57348ad33886ed5be458a443361f63a84cfbf008e46","body":"sha256:ebb65ece73f54ed9248fa76892d59f887a3f42a8fcb4d9afcd9de44cd51094ee"}
	cmd.Use = "create"
	cmd.Short = "Create a new provenance record (defaults to draft status)"
	// /sedum:describeCommand
	// sedum:setRun {"v":2,"tier":"owned","record":"prov-2026-44a07768","gen":"sha256:38d5d332ecbcb41d5a1af3b535e6cb5d9305d6ac235c5bb64e2cb02ccd551adc","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","id":{"bin":"provenance","command":"create"},"kwargs":"sha256:35cdc1872ce1e8beecbbfd356b12d9183ff5d7ccaa026fae76eb7c4b4e77238f","body":"sha256:9004acafc43c6e4c2037a62ec71656dee660bf67c94a8cd4bb78393da43b3a8d"}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runCreate(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:none {"v":2,"tier":"owned","record":"prov-2026-44a07768","gen":"sha256:38d5d332ecbcb41d5a1af3b535e6cb5d9305d6ac235c5bb64e2cb02ccd551adc","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","id":{"bin":"provenance","command":"create"},"kwargs":"sha256:0cfbd81f00249f6ed45dc26a3803d7f8349daa14c9751fbd5d43b78769d82f0b","body":"sha256:778b1e3eb6004fe613ecb8793efaf401a29f1afdc382707cea55c05601aefd79"}
	cmd.Args = cobra.NoArgs
	// /sedum:setArgs:none
	// sedum:anchor:flags
	// sedum:addFlag:string {"v":2,"tier":"owned","record":"prov-2026-44a07768","gen":"sha256:38d5d332ecbcb41d5a1af3b535e6cb5d9305d6ac235c5bb64e2cb02ccd551adc","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"create","flag":"title"},"kwargs":"sha256:4666cf8abc9ccb02ec96a10032538abed09d55098a38b508f4f2fecb232024e6","body":"sha256:9ac6c7895295c04233d903e381aaf815f09e699e5902e695e3a8211b07058cc8"}
	cmd.Flags().String("title", "", "Pre-populate the title field")
	// /sedum:addFlag:string
	// sedum:addFlag:string {"v":2,"tier":"owned","record":"prov-2026-44a07768","gen":"sha256:38d5d332ecbcb41d5a1af3b535e6cb5d9305d6ac235c5bb64e2cb02ccd551adc","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"create","flag":"intent"},"kwargs":"sha256:75759b8232f99fd9ab454c7055a13b2754b4de1205c8f36c91a44485f1663adf","body":"sha256:bc5ce741cb245f5804e698b03d419108e4011258c9e68e9e1a34706b3e00a383"}
	cmd.Flags().String("intent", "", "Pre-populate the intent field")
	// /sedum:addFlag:string
	// sedum:addFlag:string {"v":2,"tier":"owned","record":"prov-2026-44a07768","gen":"sha256:38d5d332ecbcb41d5a1af3b535e6cb5d9305d6ac235c5bb64e2cb02ccd551adc","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"create","flag":"type"},"kwargs":"sha256:194cf19da851c9fcf11359558551581fde0875928df3901f567c21d147d4fccb","body":"sha256:d7f2476b5ecc1d2f073b01bbfac8c1942f1aac8f87198afdeb541c2e1e12030e"}
	cmd.Flags().String("type", "", "Set the tier type of the record (brief|blueprint|imprint)")
	// /sedum:addFlag:string
	// sedum:addFlag:strings {"v":2,"tier":"owned","record":"prov-2026-44a07768","gen":"sha256:38d5d332ecbcb41d5a1af3b535e6cb5d9305d6ac235c5bb64e2cb02ccd551adc","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"create","flag":"tag"},"kwargs":"sha256:748b5fd25ffb424dbed8ce8ad829bd3419eb64d475761d489eca7b57a1b7af9f","body":"sha256:165d0c3da5d56dfd2efcf22b96c41001b199a80d75df35eab025db0e338f1a10"}
	cmd.Flags().StringArray("tag", nil, "Pre-populate tags (comma-separated; repeatable)")
	// /sedum:addFlag:strings
	// sedum:addFlag:strings {"v":2,"tier":"owned","record":"prov-2026-44a07768","gen":"sha256:38d5d332ecbcb41d5a1af3b535e6cb5d9305d6ac235c5bb64e2cb02ccd551adc","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"create","flag":"scope"},"kwargs":"sha256:ed9cfc37433025318f64a30e058d99ef9d3550e779699d08b81080bc6b7e2994","body":"sha256:ffde5565678cf7b7d21f8ee9048c3f0f2edc4a8654dd40b75f55767c01d33e0a"}
	cmd.Flags().StringArray("scope", nil, "Pre-populate affected_scope (repeatable)")
	// /sedum:addFlag:strings
	// sedum:addFlag:string {"v":2,"tier":"owned","record":"prov-2026-44a07768","gen":"sha256:38d5d332ecbcb41d5a1af3b535e6cb5d9305d6ac235c5bb64e2cb02ccd551adc","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"create","flag":"implements"},"kwargs":"sha256:fade36c10e8c5f74e25f87325dea70357f355a277930ac24222e8a2e514ae49c","body":"sha256:04c1350998ce55a1c6d7c35b80130ca19d7da7fd11ea68c7e7c5ccb8aaf3f805"}
	cmd.Flags().String("implements", "", "Parent record id, or its at-uri (at://<did>/dev.provisio.provenance.record/<id>); pins implements_version to a local parent's current version")
	// /sedum:addFlag:string
	// sedum:addFlag:bool {"v":2,"tier":"owned","record":"prov-2026-44a07768","gen":"sha256:38d5d332ecbcb41d5a1af3b535e6cb5d9305d6ac235c5bb64e2cb02ccd551adc","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"create","flag":"locked"},"kwargs":"sha256:91b44ade96607716b46622f355d585d09aa75803a81a98525b809cdf8e3c1cfc","body":"sha256:fbe126636e87e53b4e09a58d27e67abce7059a2263f787ab732dc9732bf0dd09"}
	cmd.Flags().Bool("locked", false, "Create the record as a locked layer")
	// /sedum:addFlag:bool
	// sedum:addFlag:bool {"v":2,"tier":"owned","record":"prov-2026-44a07768","gen":"sha256:38d5d332ecbcb41d5a1af3b535e6cb5d9305d6ac235c5bb64e2cb02ccd551adc","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"create","flag":"no-edit"},"kwargs":"sha256:9b55f836a15d37a55b23475d228003891c56ec8cc9a7c4d7419a0575e9b843ef","body":"sha256:a928134623e5d85ff6f5fa5fcf7814c300cf1fe42d27489b675a9838e878cb90"}
	cmd.Flags().Bool("no-edit", false, "Write without opening editor")
	// /sedum:addFlag:bool
	// sedum:addShortFlag:string {"v":2,"tier":"owned","record":"prov-2026-44a07768","gen":"sha256:38d5d332ecbcb41d5a1af3b535e6cb5d9305d6ac235c5bb64e2cb02ccd551adc","pkg":"sha256:317b97244b0fe957a6e0b80cdfa819a8afaae50b42b10385c9cd6729cb74fd3a","id":{"bin":"provenance","command":"create","flag":"id-suffix"},"kwargs":"sha256:aebb11c054c0cee7b5f314ed3920d4ff114298851a60163301f00031589867ef","body":"sha256:ab700e5e34c5c21a9631b1f7480cd4a30c2814d06f1a7f44d4f75c18c69372de"}
	cmd.Flags().StringP("id-suffix", "i", "", "Append service suffix to ID (e.g., user-service)")
	// /sedum:addShortFlag:string
	// sedum:anchor:subcommands
	return cmd
}
