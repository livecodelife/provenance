package main

import "github.com/spf13/cobra"

func newSearchCmd() *cobra.Command {
	cmd := &cobra.Command{}
	// sedum:anchor:command
	// sedum:describeCommand {"v":2,"tier":"owned","record":"prov-2026-4eb70ef7","gen":"sha256:8dd0cf07c10ca2de12963c41bee8f3709b76949e3c89a1471e2831d9ab3e96fe","pkg":"sha256:cb672494161bfe1ff7809e5af7923edd680b02925293c87ba9726637c12e86af","id":{"bin":"provenance","command":"search"},"kwargs":"sha256:c853ee56252824b74c9c5ca58e71087e1e67c6eb3351188a9a3a5fa015e97173","body":"sha256:baecaa77f6dc91890628fb33757248bd88ba26bde28c61dc5e776fbbde62bf4f"}
	cmd.Use = "search [query]"
	cmd.Short = "Search records by words (offline, lexical)"
	// /sedum:describeCommand
	// sedum:setRun {"v":2,"tier":"owned","record":"prov-2026-4eb70ef7","gen":"sha256:8dd0cf07c10ca2de12963c41bee8f3709b76949e3c89a1471e2831d9ab3e96fe","pkg":"sha256:5e369d6d2a4af49acd0fc42910733893fbbd57e90acd09046714692a30b3819e","id":{"bin":"provenance","command":"search"},"kwargs":"sha256:f459f79a47743e4473f02649cad86eb130430d3b00fcf2a0c7949acfc6128ac0","body":"sha256:ceb5da5728238e3128e00d228fa2c444dddc5a5842d59255b0516f3bb0a5f1ce"}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runSearch(cmd.OutOrStdout(), cmd, args)
	}
	// /sedum:setRun
	// sedum:setArgs:max_one {"v":2,"tier":"owned","record":"prov-2026-4eb70ef7","gen":"sha256:8dd0cf07c10ca2de12963c41bee8f3709b76949e3c89a1471e2831d9ab3e96fe","pkg":"sha256:ca644378490be5811d57a9072cb3caf3c807f5d2c6deee0170e1bc58e0c2d3f9","id":{"bin":"provenance","command":"search"},"kwargs":"sha256:02f2c17aed93b9ec6158d908ad0b48b297629bd25a3488a37d3e972892865e6b","body":"sha256:186046a2435bbf4d0eae67e61e8a3909c91fd72c18b73dd2ff749798e3517ef3"}
	cmd.Args = cobra.MaximumNArgs(1)
	// /sedum:setArgs:max_one
	// sedum:anchor:flags
	// sedum:addFlag:string {"v":2,"tier":"owned","record":"prov-2026-4eb70ef7","gen":"sha256:8dd0cf07c10ca2de12963c41bee8f3709b76949e3c89a1471e2831d9ab3e96fe","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"search","flag":"query"},"kwargs":"sha256:a28e660b9cf27f07642b53be223c911ada0c9dd77f8468f7a6b6c247f7fcef14","body":"sha256:577e44502a543e1858cc4b776afde7d8c0aa590950a16bbcdd31fb1b5ccfb939"}
	cmd.Flags().String("query", "", "Search words (alternative to the argument)")
	// /sedum:addFlag:string
	// sedum:addFlag:int {"v":2,"tier":"owned","record":"prov-2026-4eb70ef7","gen":"sha256:8dd0cf07c10ca2de12963c41bee8f3709b76949e3c89a1471e2831d9ab3e96fe","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"search","flag":"limit"},"kwargs":"sha256:48ae02d1210045f126e12ceae2e5081af53674e55de01d7a7a9643c2ebc5c60a","body":"sha256:0ca50f0cb8bc5ab1d506e9bd26ebdb86c9aa545f78b1e9c9e8964fc6822418cf"}
	cmd.Flags().Int("limit", 0, "Maximum results (default 5, at most 50)")
	// /sedum:addFlag:int
	// sedum:addFlag:string {"v":2,"tier":"owned","record":"prov-2026-4eb70ef7","gen":"sha256:8dd0cf07c10ca2de12963c41bee8f3709b76949e3c89a1471e2831d9ab3e96fe","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"search","flag":"status"},"kwargs":"sha256:cecc4e83e26905cc56833a6d94e86bc79c342b8c487289cc855ffbf2c023e2e4","body":"sha256:220e84b71d76d96e5e696c63825690fecba55f8f11f57a5d387e0950ceef1efa"}
	cmd.Flags().String("status", "", "Only records with this status (draft|open|implemented|deprecated)")
	// /sedum:addFlag:string
	// sedum:addFlag:string {"v":2,"tier":"owned","record":"prov-2026-4eb70ef7","gen":"sha256:8dd0cf07c10ca2de12963c41bee8f3709b76949e3c89a1471e2831d9ab3e96fe","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"search","flag":"type"},"kwargs":"sha256:2fdbdcd537a28b1ba37e0b97568991e197ae0339283fbf5ab432538b87ab2d3b","body":"sha256:678c45b80e66cd4b713aa3fe834c3c0d270fe0d79fc3e9cdc4ebdd8cc585367e"}
	cmd.Flags().String("type", "", "Only records of this tier (brief|blueprint|imprint)")
	// /sedum:addFlag:string
	// sedum:addFlag:string {"v":2,"tier":"owned","record":"prov-2026-4eb70ef7","gen":"sha256:8dd0cf07c10ca2de12963c41bee8f3709b76949e3c89a1471e2831d9ab3e96fe","pkg":"sha256:d873fc5f71035cf617eb56ca6ca05ce36f39842f4f774d2fb6b6816cf95ccc9b","id":{"bin":"provenance","command":"search","flag":"format"},"kwargs":"sha256:6c0e18aa1ad11548148aa1d4f9b092cb2cb81c8be2bbf77ca260e484db1a5416","body":"sha256:5a11ae1e68306dbe22a26ffec6af235a5dc44898684981a9cddacca632a7b196"}
	cmd.Flags().String("format", "", "Output format (human|json)")
	// /sedum:addFlag:string
	// sedum:anchor:subcommands
	return cmd
}
