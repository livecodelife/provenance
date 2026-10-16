package main

import (
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
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-063f1395","gen":"sha256:1757f92e4668f0e73d83db1120b817ccdd8e464d7f5e59cbfe98102926ec0b12","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"provenance","command":"check"}}
	root.AddCommand(newCheckCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-12b1ca97","gen":"sha256:3d5eb5aff0b580425966bd18cb1e038d9fef9018ce5f4611f46a803ad6ee5bc9","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"provenance","command":"status"}}
	root.AddCommand(newStatusCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-35439261","gen":"sha256:7a7f9652ab699a1d6305ff28ecf9f26e5a23c00aa71b77596c96dedb6c805fe1","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"provenance","command":"deprecate"}}
	root.AddCommand(newDeprecateCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-41fb548c","gen":"sha256:cccc5b62420ee91b0c3cda795fa3d41efaee0b261b6586bfaa6592cb2e629e4a","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"provenance","command":"complete"}}
	root.AddCommand(newCompleteCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-44a07768","gen":"sha256:ba3e6b98ee7f4297af267684cf4d5385650dd9703772a650529ecba804c273f6","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"provenance","command":"create"}}
	root.AddCommand(newCreateCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-57b8a95b","gen":"sha256:8834ae395b2cf22cd905f6349e82c36326110c430c4a58c1e9ca2c6a25561d7d","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"provenance","command":"edit"}}
	root.AddCommand(newEditCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-5f69977b","gen":"sha256:291944a4f33f90d52c7f30e039d94974dab160a76b3b4e05001226d563ef6d90","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"provenance","command":"revise"}}
	root.AddCommand(newReviseCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-6db88b75","gen":"sha256:a536c55853a03849a57674103e15b0ff1cd358706f90eebef3ae8ae03fe09658","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"provenance","command":"open"}}
	root.AddCommand(newOpenCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-76a125e6","gen":"sha256:68d57b1b46d0b0deed02c7130fec8249449d9587617ea80f4e9ce204f34362a1","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"provenance","command":"migrate"}}
	root.AddCommand(newMigrateCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-9bfa6107","gen":"sha256:ec4f7dba3f7e1d4d2bdffa3ede03727f7b0c533cc8c3831f66522a7a81a8ad4e","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"provenance","command":"repin"}}
	root.AddCommand(newRepinCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-a03e6bc3","gen":"sha256:aec5a8b3297a63ed73463101f5712b37b803293254ec536f0c5009748de88f09","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"provenance","command":"lint"}}
	root.AddCommand(newLintCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-bfd55535","gen":"sha256:98fb8e2f8fd67f02e2cabc1153d5dfb84f1f27f5c19d0582f9236cd5fa8e4f10","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"provenance","command":"diff"}}
	root.AddCommand(newDiffCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-c31c9c59","gen":"sha256:9a18a4eb62cedc8b94c8ee92552042184ff66c2dcd6e7dd094c735af31f5de23","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"provenance","command":"run-specs"}}
	root.AddCommand(newRunSpecsCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-d1d789de","gen":"sha256:bb33e281db3fef59d1d98f44c667bf0ad0797c86378e45488297adfe0b0afb14","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"provenance","command":"hooks"}}
	root.AddCommand(newHooksCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-d435c27b","gen":"sha256:4f36ba1a5a6a79f262293754b5a46ef8739347c5fb4b6df037af455f19b551ba","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"provenance","command":"history"}}
	root.AddCommand(newHistoryCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"tier":"owned","record":"prov-2026-f476b640","gen":"sha256:9dea06195024f8f78498344d38c9f26ddee5cdda7c9672b172c420149810bc3c","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","kwargs":{"bin":"provenance","command":"stale"}}
	root.AddCommand(newStaleCmd())
	// /sedum:registerCommand
	return root
}

func main() {
	if err := newRootCmd().Execute(); err != nil {
		os.Exit(1)
	}
}
