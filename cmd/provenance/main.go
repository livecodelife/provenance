package main

import (
	"errors"
	"os"

	"github.com/spf13/cobra"
)

func newRootCmd() *cobra.Command {
	root := &cobra.Command{SilenceUsage: true}
	// sedum:anchor:root
	// sedum:describeRoot {"tier":"owned","record":"prov-2026-76a125e6","gen":"sha256:68d57b1b46d0b0deed02c7130fec8249449d9587617ea80f4e9ce204f34362a1","pkg":"sha256:0fc81c42ecd7f40945040b8bc7254475c058082a098a34b5f2c8660225ff618e","kwargs":{"bin":"provenance","short":"Versioned provenance records","use":"provenance"}}
	root.Use = "provenance"
	root.Short = "Versioned provenance records"
	// /sedum:describeRoot
	// sedum:anchor:persistent_flags
	// sedum:addShortPersistentFlag:string {"tier":"owned","record":"prov-2026-76a125e6","gen":"sha256:68d57b1b46d0b0deed02c7130fec8249449d9587617ea80f4e9ce204f34362a1","pkg":"sha256:9a5acc9c3ccb4219037b31eecd6f40891828ebeb30d9dcc2e18be6a46d59dcae","kwargs":{"bin":"provenance","flag":"config","shorthand":"c","type":"string","usage":"Config file (default .provenance.yml, else the provenance block of .linespec.yml)"}}
	root.PersistentFlags().StringP("config", "c", "", "Config file (default .provenance.yml, else the provenance block of .linespec.yml)")
	// /sedum:addShortPersistentFlag:string
	// sedum:anchor:commands
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-063f1395","gen":"sha256:be8cdfa71ba526e7a2b3721fabfc6aa6d5cd8952b1ae6a78999dd073f1d63c4d","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"provenance","command":"check"}}
	root.AddCommand(newCheckCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-12b1ca97","gen":"sha256:3d5eb5aff0b580425966bd18cb1e038d9fef9018ce5f4611f46a803ad6ee5bc9","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"provenance","command":"status"}}
	root.AddCommand(newStatusCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-35439261","gen":"sha256:bb73fb8a27febfb6636afc91c6f46ef087ceb45d41e95db14a66c7469f2256d4","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"provenance","command":"deprecate"}}
	root.AddCommand(newDeprecateCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-41fb548c","gen":"sha256:03af2c7fbe865130ed6609ebc92a9c3677f9826266df64175537a4fedf05032d","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"provenance","command":"complete"}}
	root.AddCommand(newCompleteCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-44a07768","gen":"sha256:38d5d332ecbcb41d5a1af3b535e6cb5d9305d6ac235c5bb64e2cb02ccd551adc","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"provenance","command":"create"}}
	root.AddCommand(newCreateCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-57b8a95b","gen":"sha256:ecce14df038164301d4c4728b274460b7a56a8972e302aacf33c69ff352be339","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"provenance","command":"edit"}}
	root.AddCommand(newEditCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-5f69977b","gen":"sha256:8b7b52e1252d8b9aa1024924c8dc2aac14d21148579ea2b89a1653897fffab8a","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"provenance","command":"revise"}}
	root.AddCommand(newReviseCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-6db88b75","gen":"sha256:18e81b445bc0135242f33754116c5073aa843ca855024412b23f848a00f1084b","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"provenance","command":"open"}}
	root.AddCommand(newOpenCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-76a125e6","gen":"sha256:68d57b1b46d0b0deed02c7130fec8249449d9587617ea80f4e9ce204f34362a1","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"provenance","command":"migrate"}}
	root.AddCommand(newMigrateCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-9a3aad45","gen":"sha256:d19012536182a12da7b036a73f8eea6a7a589bd985a0ee0352375454ff233111","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"provenance","command":"reconcile"}}
	root.AddCommand(newReconcileCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-9bfa6107","gen":"sha256:e937012cd7f04b4dea168e2f24be21b5aadd0783afbe0dda58babc69031e3dbc","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"provenance","command":"repin"}}
	root.AddCommand(newRepinCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-a03e6bc3","gen":"sha256:08032897102932b2a0dd0a941644e2f7fce15c3bdb19714d078c7ce333926547","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"provenance","command":"lint"}}
	root.AddCommand(newLintCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-bfd55535","gen":"sha256:98fb8e2f8fd67f02e2cabc1153d5dfb84f1f27f5c19d0582f9236cd5fa8e4f10","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"provenance","command":"diff"}}
	root.AddCommand(newDiffCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-c31c9c59","gen":"sha256:9a18a4eb62cedc8b94c8ee92552042184ff66c2dcd6e7dd094c735af31f5de23","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"provenance","command":"run-specs"}}
	root.AddCommand(newRunSpecsCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-d1d789de","gen":"sha256:6202ce5f10060b0fa61443fc52df17f80e57f28bdc0e09ce5ef4e680dc50ebff","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"provenance","command":"hooks"}}
	root.AddCommand(newHooksCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-d435c27b","gen":"sha256:4f36ba1a5a6a79f262293754b5a46ef8739347c5fb4b6df037af455f19b551ba","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"provenance","command":"history"}}
	root.AddCommand(newHistoryCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-f476b640","gen":"sha256:24167071b71953f56b7a9351b440d09a67530fb3dc4508cd3e405b0ad9bb4806","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"provenance","command":"stale"}}
	root.AddCommand(newStaleCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-f807707a","gen":"sha256:3f6d1802d83ba9dcc72ed02c8927ae3f79e56242263a8c87803b425321b88f73","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"provenance","command":"sync"}}
	root.AddCommand(newSyncCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-11876813","gen":"sha256:090b658e8472ed58a1fb75e4a4101e7f16690a1a70b99759f485f9af7dfaeb18","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"provenance","command":"next"}}
	root.AddCommand(newNextCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-15f42a00","gen":"sha256:0c0506c22fde6bea7a03238d33ecb1685877efd1f0d133e1b1730b1e64cd7996","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"provenance","command":"evidence"}}
	root.AddCommand(newEvidenceCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-1a5d2549","gen":"sha256:e344efddd9441a7734ecfef807e1e2c0d0a91b33f8cffcc5e6d16b03b1229edb","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"provenance","command":"init"}}
	root.AddCommand(newInitCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-215c273d","gen":"sha256:7d6dead1563884c13e30140873f1cd8a63b81cc841aedfb88922843a11509c33","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"provenance","command":"whoami"}}
	root.AddCommand(newWhoamiCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-32201524","gen":"sha256:da838940ae1e0b662456e90348132411fdd4321c7c2aa1a859c590f248a94d7c","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"provenance","command":"logout"}}
	root.AddCommand(newLogoutCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-4eb70ef7","gen":"sha256:8dd0cf07c10ca2de12963c41bee8f3709b76949e3c89a1471e2831d9ab3e96fe","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"provenance","command":"search"}}
	root.AddCommand(newSearchCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:57a095967ea4cbd86875cea0f96e504d2e1d138c4f31fba23d7b01ba011caa69","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"provenance","command":"context"}}
	root.AddCommand(newContextCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-635b25ac","gen":"sha256:47717a6e305353fab65b1cd72cdd7e2343a55090484ccb9c090403fd17b3a020","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"provenance","command":"login"}}
	root.AddCommand(newLoginCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-75646f74","gen":"sha256:bfc4a4b4ff7397d2474dbe459bfda6f85459580329efbe550b3e2379a9c827c6","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"provenance","command":"clone"}}
	root.AddCommand(newCloneCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-83d57ec1","gen":"sha256:6b4d5a3926de0813e7bebfa8498ef520429d4b6b530252ea1eb108e16c6d8def","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"provenance","command":"publish"}}
	root.AddCommand(newPublishCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-871605a9","gen":"sha256:a837b61e0b1282ffa2184bd15712bd692804fd3785f63fd676f9036654ebea2d","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"provenance","command":"import"}}
	root.AddCommand(newImportCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-9c8a0b2a","gen":"sha256:9c65c23d7893e39a1725f1f00b3344e10ee8e1823f7702b3722d3bef7130ba0c","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"provenance","command":"add-scope"}}
	root.AddCommand(newAddScopeCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-9ec4a843","gen":"sha256:cc0301a10fcd903006e4f2bed32a32b6a082f192eb2c42bb2c80462870e50daf","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"provenance","command":"govern"}}
	root.AddCommand(newGovernCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-d17b9e44","gen":"sha256:ac383f099a8cbb5db3d785a4c63bd601a4b3df79dd4e69ecde937bf6b98aeca0","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"provenance","command":"graph"}}
	root.AddCommand(newGraphCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-d71d0066","gen":"sha256:cae7a2d3a080d43f1ba12537ce93365db84c39195d99ec904268e9c15daab6a9","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"provenance","command":"install-skills"}}
	root.AddCommand(newInstallSkillsCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-d71d0066","gen":"sha256:cae7a2d3a080d43f1ba12537ce93365db84c39195d99ec904268e9c15daab6a9","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"provenance","command":"install-plugin"}}
	root.AddCommand(newInstallPluginCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-d84dc10e","gen":"sha256:2e97312cea5e86a4f4e11ea65cacfe9c4a63da245a1c446b00a2c6740234e92a","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"provenance","command":"compile"}}
	root.AddCommand(newCompileCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-adabd65f","gen":"sha256:4ae46f3089d333bd12c54eda92a2d20ae90a3ad1ebf7c65d3eda625c0356ee02","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"provenance","command":"version"}}
	root.AddCommand(newVersionCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-d84dc10e","gen":"sha256:2e97312cea5e86a4f4e11ea65cacfe9c4a63da245a1c446b00a2c6740234e92a","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"provenance","command":"merge-manifest"}}
	root.AddCommand(newMergeManifestCmd())
	// /sedum:registerCommand
	return root
}

// exitCoder is implemented by an error that carries the process exit status.
type exitCoder interface{ ExitCode() int }

func main() {
	if err := newRootCmd().Execute(); err != nil {
		var coded exitCoder
		if errors.As(err, &coded) {
			os.Exit(coded.ExitCode())
		}
		os.Exit(1)
	}
}
