package main

import (
	"errors"
	"os"

	"github.com/spf13/cobra"
)

func newRootCmd() *cobra.Command {
	root := &cobra.Command{SilenceUsage: true}
	// sedum:anchor:root
	// sedum:describeRoot {"v":2,"tier":"owned","record":"prov-2026-76a125e6","gen":"sha256:68d57b1b46d0b0deed02c7130fec8249449d9587617ea80f4e9ce204f34362a1","pkg":"sha256:0fc81c42ecd7f40945040b8bc7254475c058082a098a34b5f2c8660225ff618e","id":{"bin":"provenance"},"kwargs":"sha256:cd590f8c44a47ce6c5328e1fc5202640aab16948bbc42e619168744840ca6464","body":"sha256:afce34e714b3952a50f3a6b8207e986c86f0253bbb74a0e7eff3ed41d0d8db84"}
	root.Use = "provenance"
	root.Short = "Versioned provenance records"
	// /sedum:describeRoot
	// sedum:anchor:persistent_flags
	// sedum:addShortPersistentFlag:string {"v":2,"tier":"owned","record":"prov-2026-76a125e6","gen":"sha256:68d57b1b46d0b0deed02c7130fec8249449d9587617ea80f4e9ce204f34362a1","pkg":"sha256:9a5acc9c3ccb4219037b31eecd6f40891828ebeb30d9dcc2e18be6a46d59dcae","id":{"bin":"provenance","flag":"config"},"kwargs":"sha256:64d3f843dd410bee37f964b7925189bb9be651bb80ce7145de3e06c795936ac0","body":"sha256:ef25ef027d967af36ed7f7aa31ccf4385f5c7cbb88a36ddcf2a4a06f57532446"}
	root.PersistentFlags().StringP("config", "c", "", "Config file (default .provenance.yml, else the provenance block of .linespec.yml)")
	// /sedum:addShortPersistentFlag:string
	// sedum:anchor:commands
	// sedum:registerCommand {"v":2,"tier":"owned","record":"prov-2026-063f1395","gen":"sha256:4239f75c5db7efc2453be01f153d0012da926ff1bd1870547582214fbf406657","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","id":{"bin":"provenance","command":"check"},"kwargs":"sha256:170ac74438ebb2fec298dfb7ee87090d28eb803209cc56fe7e1c62a4d971f2b1","body":"sha256:ec8348f67339c5b211f135ea40ecb26739c7dee7db6d4224a9d09a5deb524383"}
	root.AddCommand(newCheckCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"v":2,"tier":"owned","record":"prov-2026-12b1ca97","gen":"sha256:3d5eb5aff0b580425966bd18cb1e038d9fef9018ce5f4611f46a803ad6ee5bc9","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","id":{"bin":"provenance","command":"status"},"kwargs":"sha256:56de959af0e69bfa72739f3fd0ddb7e7b6dd7b860049ef0ba00a67da6d7e9709","body":"sha256:1c18d07d90d7f73d88e01be32fddda390e5d10616790f15fd61f5a4ab8ae8ae8"}
	root.AddCommand(newStatusCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"v":2,"tier":"owned","record":"prov-2026-35439261","gen":"sha256:bb73fb8a27febfb6636afc91c6f46ef087ceb45d41e95db14a66c7469f2256d4","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","id":{"bin":"provenance","command":"deprecate"},"kwargs":"sha256:88ddf979fa7ecf3af2be46dba0b7f74f8715cd7fad98af43b3f47ac2142f3ad1","body":"sha256:4677a2f8e33f906680cb54f9f7c5b2332778216f1cecfd5b7e54bcbdcec08031"}
	root.AddCommand(newDeprecateCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"v":2,"tier":"owned","record":"prov-2026-41fb548c","gen":"sha256:953e06b4e4abb9f509dd88530863fa800eeb22ece21fb38f9e9dfc6e57e8e5cf","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","id":{"bin":"provenance","command":"complete"},"kwargs":"sha256:8448526f028d904a182d3f00323d53140a08444c90b13d1b14b2b97926d4053d","body":"sha256:d1ab52b21850e5d3abeba557ad94b992d8f53ca5efb24fae31935b8412d43e7d"}
	root.AddCommand(newCompleteCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"v":2,"tier":"owned","record":"prov-2026-44a07768","gen":"sha256:38d5d332ecbcb41d5a1af3b535e6cb5d9305d6ac235c5bb64e2cb02ccd551adc","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","id":{"bin":"provenance","command":"create"},"kwargs":"sha256:35cdc1872ce1e8beecbbfd356b12d9183ff5d7ccaa026fae76eb7c4b4e77238f","body":"sha256:98d62ff305d14f93403f8c4a74acdc07b42b8db52c4bd3c74b2c5657b68f1c1c"}
	root.AddCommand(newCreateCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"v":2,"tier":"owned","record":"prov-2026-57b8a95b","gen":"sha256:ecce14df038164301d4c4728b274460b7a56a8972e302aacf33c69ff352be339","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","id":{"bin":"provenance","command":"edit"},"kwargs":"sha256:e7a4b6073ea9164a358b8fa6132903dec3f2b70ef55af18ce7bc6a72083c3820","body":"sha256:321869383d89042f0345b5c997ea905e039d7dc1ad340b2889eadfb37b35462d"}
	root.AddCommand(newEditCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"v":2,"tier":"owned","record":"prov-2026-5f69977b","gen":"sha256:8b7b52e1252d8b9aa1024924c8dc2aac14d21148579ea2b89a1653897fffab8a","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","id":{"bin":"provenance","command":"revise"},"kwargs":"sha256:cfec1d8d1c6288021f821c41bee3664956aed03a3cf2fded8708a654b312ef27","body":"sha256:f2fcd25e90b66bcaeb4c0831667675765d234d692227b33e234053a3e090b2ce"}
	root.AddCommand(newReviseCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"v":2,"tier":"owned","record":"prov-2026-6db88b75","gen":"sha256:18e81b445bc0135242f33754116c5073aa843ca855024412b23f848a00f1084b","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","id":{"bin":"provenance","command":"open"},"kwargs":"sha256:e1bcceda3bda344b4c39739d5d9da6f5fdedca5f6b793a889959657bf7672faf","body":"sha256:4aad7e5b705a43a33d96d66453eea9b6beb9d57af81b5ca890b7e4977f189d46"}
	root.AddCommand(newOpenCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"v":2,"tier":"owned","record":"prov-2026-76a125e6","gen":"sha256:68d57b1b46d0b0deed02c7130fec8249449d9587617ea80f4e9ce204f34362a1","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","id":{"bin":"provenance","command":"migrate"},"kwargs":"sha256:c4c99cb10c851126b8f1d4aa2af1ce7a664b1224a96a21975ba3f3649f2aa818","body":"sha256:38cd69ce3498bd21fb2d0b5801f7ee13643cd7340369ff0ace184f43a6769279"}
	root.AddCommand(newMigrateCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"v":2,"tier":"owned","record":"prov-2026-9a3aad45","gen":"sha256:d19012536182a12da7b036a73f8eea6a7a589bd985a0ee0352375454ff233111","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","id":{"bin":"provenance","command":"reconcile"},"kwargs":"sha256:781b147f4c3bc78c9c38c1fba0900748ac441ae3e3e2f9aad20bd6fb838a6d08","body":"sha256:05fa48d9f61b9e443785718d9be2f6976b2d21e4757721cef0a6e82f5464b0a9"}
	root.AddCommand(newReconcileCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"v":2,"tier":"owned","record":"prov-2026-9bfa6107","gen":"sha256:e937012cd7f04b4dea168e2f24be21b5aadd0783afbe0dda58babc69031e3dbc","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","id":{"bin":"provenance","command":"repin"},"kwargs":"sha256:a86827076b82167baa7d8f2b0862bd26bf0a132a3fd4f0c0fcb0c0541d2230be","body":"sha256:26eaa6a6ce7d77183b8d1822f367f9e0779bc0cff1bd4f5f80e96126bd685256"}
	root.AddCommand(newRepinCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"v":2,"tier":"owned","record":"prov-2026-a03e6bc3","gen":"sha256:08032897102932b2a0dd0a941644e2f7fce15c3bdb19714d078c7ce333926547","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","id":{"bin":"provenance","command":"lint"},"kwargs":"sha256:a651b403bbf345bf15702d90556c49d517e43617fec543c2143ab42c7bee93cb","body":"sha256:6a203ba5501d27235570394aeabdc9433bb7e2c08609f0b3984b4b0861e9c316"}
	root.AddCommand(newLintCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"v":2,"tier":"owned","record":"prov-2026-bfd55535","gen":"sha256:98fb8e2f8fd67f02e2cabc1153d5dfb84f1f27f5c19d0582f9236cd5fa8e4f10","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","id":{"bin":"provenance","command":"diff"},"kwargs":"sha256:e65b2c6b081d304a2ef912bd19d07fd3f207b6678322ba0067ad0a18a03fae31","body":"sha256:cd2ea7a989c1888c5c8ea1f43814d9b055d8055a008293c1a372d68a85f0a149"}
	root.AddCommand(newDiffCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"v":2,"tier":"owned","record":"prov-2026-c31c9c59","gen":"sha256:9a18a4eb62cedc8b94c8ee92552042184ff66c2dcd6e7dd094c735af31f5de23","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","id":{"bin":"provenance","command":"run-specs"},"kwargs":"sha256:97c669f9c235b598b8daac53c90a35a727a81ea2c1a96d2e907443fc656d1812","body":"sha256:f8e499ab6c717f926f71cdffaba47be2ef0058afe89fd4338130934712ab4e6d"}
	root.AddCommand(newRunSpecsCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"v":2,"tier":"owned","record":"prov-2026-d1d789de","gen":"sha256:6202ce5f10060b0fa61443fc52df17f80e57f28bdc0e09ce5ef4e680dc50ebff","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","id":{"bin":"provenance","command":"hooks"},"kwargs":"sha256:989114f30898b1941d09eb09f8a5e404218d69f30fe924f10edf0c024a82ae15","body":"sha256:ae3f5f225c89a66318fc7e3b5543e5ed55e937e4a705eb06dfb71f7868e9ff53"}
	root.AddCommand(newHooksCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"v":2,"tier":"owned","record":"prov-2026-d435c27b","gen":"sha256:4f36ba1a5a6a79f262293754b5a46ef8739347c5fb4b6df037af455f19b551ba","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","id":{"bin":"provenance","command":"history"},"kwargs":"sha256:f16fb5be7b018724f49ff95d556f8723a1c38582dcddce2243e123e7426fb28b","body":"sha256:38f696ff50c2d463ed17a5dfb1941700fc92bfe7274fd27914a3e8f0ecbf0365"}
	root.AddCommand(newHistoryCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"v":2,"tier":"owned","record":"prov-2026-f476b640","gen":"sha256:24167071b71953f56b7a9351b440d09a67530fb3dc4508cd3e405b0ad9bb4806","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","id":{"bin":"provenance","command":"stale"},"kwargs":"sha256:be070c465b03cea9756295b67c88eea09bdd1300b0fa0c027b5f86ebdce83d72","body":"sha256:fc330b9ab907019b90aa24cb184bdbd775edf9f4e88cc4e7cb9f36eb54ce07a8"}
	root.AddCommand(newStaleCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"v":2,"tier":"owned","record":"prov-2026-f807707a","gen":"sha256:3f6d1802d83ba9dcc72ed02c8927ae3f79e56242263a8c87803b425321b88f73","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","id":{"bin":"provenance","command":"sync"},"kwargs":"sha256:5620389fe1e53752ef29ea0eb88502349079dd89819f18d985bd0f2702a2fe60","body":"sha256:8d264c5e522b6d83d4bc1f406249110a829d3e1078a411aa4048515670544a9b"}
	root.AddCommand(newSyncCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"v":2,"tier":"owned","record":"prov-2026-11876813","gen":"sha256:090b658e8472ed58a1fb75e4a4101e7f16690a1a70b99759f485f9af7dfaeb18","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","id":{"bin":"provenance","command":"next"},"kwargs":"sha256:395e34d4f4eed74c05856688100d17a90d3ecd7be4c02fc203eb3b0b15e3a8d6","body":"sha256:a4dff6199c7d243d38e39d70509b5ab4168feaf2ece985b204d2332698e63266"}
	root.AddCommand(newNextCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"v":2,"tier":"owned","record":"prov-2026-15f42a00","gen":"sha256:9960743cde810036bdd5d9af367a4c54dcc99657e529725ea71d385c210784ac","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","id":{"bin":"provenance","command":"evidence"},"kwargs":"sha256:ea423a1b16c9cd960b471e9795e835596769c36c6825242874714caae67a8be2","body":"sha256:6c066b87cc2b807f50c67fd238c4adc0be30e2853351d7287562e1bfd6d2dc47"}
	root.AddCommand(newEvidenceCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"v":2,"tier":"owned","record":"prov-2026-1a5d2549","gen":"sha256:95782fcfe653edd478163f4fae54335b865cf426a0e0fccd5f1fc4c3bb633ffe","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","id":{"bin":"provenance","command":"init"},"kwargs":"sha256:49ba7e13aa36d288fc35e0b7a35dcf581f71d967e1d1fae31cb713f30cdcb3db","body":"sha256:2cd83f6ef1adfc205809548e08106da00f0e06e0393d177864fe678df693bd96"}
	root.AddCommand(newInitCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"v":2,"tier":"owned","record":"prov-2026-215c273d","gen":"sha256:35eaf90eaf676b31114c953aa2ccef582eceebd10849617fd9691d67499118f6","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","id":{"bin":"provenance","command":"whoami"},"kwargs":"sha256:353cc45eeeff17da10c47383fe7f132d84e480899052b58137903c8bb55032d9","body":"sha256:b7b9d03509911f9703ff7ec9d1c0515e94a6694cc050411ba0087a98a19ea5de"}
	root.AddCommand(newWhoamiCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"v":2,"tier":"owned","record":"prov-2026-32201524","gen":"sha256:ee05402713f727aa111554dd0aacd7f0e84ce70ae0e346c620bc80a4c4bad7c7","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","id":{"bin":"provenance","command":"logout"},"kwargs":"sha256:a826e9b22284c312384c899f3543700e7932e58cb10ca43b9fc05531b31a0e5b","body":"sha256:01f1901a127853de9b401c7ba20f4608cc34043186502b39fddb3f04c05d8668"}
	root.AddCommand(newLogoutCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"v":2,"tier":"owned","record":"prov-2026-4eb70ef7","gen":"sha256:8dd0cf07c10ca2de12963c41bee8f3709b76949e3c89a1471e2831d9ab3e96fe","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","id":{"bin":"provenance","command":"search"},"kwargs":"sha256:f459f79a47743e4473f02649cad86eb130430d3b00fcf2a0c7949acfc6128ac0","body":"sha256:a5d49cf05e1074bdf7662aee26d2797a2065502cf6a849d8adb453a916d5de65"}
	root.AddCommand(newSearchCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"v":2,"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:57a095967ea4cbd86875cea0f96e504d2e1d138c4f31fba23d7b01ba011caa69","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","id":{"bin":"provenance","command":"context"},"kwargs":"sha256:02691fe13fca3226c8f871e0260117973260716b0757d776e7ed4d7c05848465","body":"sha256:6a0cb0e626873946e9e3bd60d6652b8f0b1377c1029429ffe31fff6464ab109b"}
	root.AddCommand(newContextCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"v":2,"tier":"owned","record":"prov-2026-635b25ac","gen":"sha256:3dd07349442894e371a607dad8f67454397b701f2bd4e8d8d0edfb17be96c5a0","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","id":{"bin":"provenance","command":"login"},"kwargs":"sha256:e1156eb1407a559e28bb114a95d2636f4a261bde5bd2a58041f7ca8ddfd21c48","body":"sha256:aad1eb3f9e8af2fc775f4fd9ba163b7dbeb39de4dcd74b60d6262e4183025241"}
	root.AddCommand(newLoginCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"v":2,"tier":"owned","record":"prov-2026-75646f74","gen":"sha256:335539008ab4bfd313eddaf0df95cf3ef61bd753a11aa929a86ff10195eee983","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","id":{"bin":"provenance","command":"clone"},"kwargs":"sha256:50442a1d4ef2709ff51ce9d849fc939851bb4354df1e4af8f4c156a8d2eaec3d","body":"sha256:568e53ac4b5f19b147a4727fb270222188fa5665a03216a67dfafde75453811c"}
	root.AddCommand(newCloneCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"v":2,"tier":"owned","record":"prov-2026-83d57ec1","gen":"sha256:671154c039eaa7dc75b6ff179083a4dcf289b36501a3ddf2b0cbcd42ba3d794b","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","id":{"bin":"provenance","command":"publish"},"kwargs":"sha256:dd6ccb786f4fd36221547b4839e6b254fed767f515f1b02999611022bf9e5527","body":"sha256:987f2f58f38aec653255935eea7836af90edcff38ce43a81906310d49a54e588"}
	root.AddCommand(newPublishCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"v":2,"tier":"owned","record":"prov-2026-871605a9","gen":"sha256:c7b209fbc9bbd544f0e8d90fa47cd5bb0b654ddba7423c834cd61299b3619190","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","id":{"bin":"provenance","command":"import"},"kwargs":"sha256:fdac7a72d2523614b7ac53d3a6c788f888330139f0261b2683cedf30c4b7886c","body":"sha256:d1ef19e1bed8e903e65ae4565674ca236bbd348e9080d128b646e932f0f9d94c"}
	root.AddCommand(newImportCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"v":2,"tier":"owned","record":"prov-2026-9c8a0b2a","gen":"sha256:9c65c23d7893e39a1725f1f00b3344e10ee8e1823f7702b3722d3bef7130ba0c","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","id":{"bin":"provenance","command":"add-scope"},"kwargs":"sha256:6177d19c9717aaf86087f74a48007d50c9005f8526a4b2f77c1b2d6e95c89318","body":"sha256:84a9190113da17dbf19d8765f9a8c2f5b6b538307e6f446cbaa5c793fed1eb0e"}
	root.AddCommand(newAddScopeCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"v":2,"tier":"owned","record":"prov-2026-9ec4a843","gen":"sha256:cc0301a10fcd903006e4f2bed32a32b6a082f192eb2c42bb2c80462870e50daf","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","id":{"bin":"provenance","command":"govern"},"kwargs":"sha256:fe1df7947e25554f2ffebb2c2e3b799532ec6d224cf3847e3ad4751ad6b85c0b","body":"sha256:be57724db563aee272ba828403bcb1f9da7c87612f97e68354b6d9f84aa29b2a"}
	root.AddCommand(newGovernCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"v":2,"tier":"owned","record":"prov-2026-d17b9e44","gen":"sha256:ac383f099a8cbb5db3d785a4c63bd601a4b3df79dd4e69ecde937bf6b98aeca0","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","id":{"bin":"provenance","command":"graph"},"kwargs":"sha256:f4edc5b1ba3b04fef802ac55265fabfa223f2ecb82c119302a25af7ba7ae773c","body":"sha256:267f66fc59e20e43831af82983600c4c2e4583ed51f85dc0ae0ed3c0650a4cdd"}
	root.AddCommand(newGraphCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"v":2,"tier":"owned","record":"prov-2026-d71d0066","gen":"sha256:cae7a2d3a080d43f1ba12537ce93365db84c39195d99ec904268e9c15daab6a9","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","id":{"bin":"provenance","command":"install-skills"},"kwargs":"sha256:19eaf2c9b5ac0bdfab63f51a29db850be48c22487cb56ea8221a3172d13de22a","body":"sha256:35f07c1741ad0f8ea9f149fa8c43aa9fb785b006a48c6dbc2c02da244f89205e"}
	root.AddCommand(newInstallSkillsCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"v":2,"tier":"owned","record":"prov-2026-d71d0066","gen":"sha256:cae7a2d3a080d43f1ba12537ce93365db84c39195d99ec904268e9c15daab6a9","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","id":{"bin":"provenance","command":"install-plugin"},"kwargs":"sha256:a37c12b3ba1f3e61e7be3096ac0f798f2ce02af41bc83034dd38e0d0bb396be8","body":"sha256:008f699f0f8784e8f211d634e444f944a9a8f06621f3c2c6e30cfa9b5ca845fd"}
	root.AddCommand(newInstallPluginCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"v":2,"tier":"owned","record":"prov-2026-d84dc10e","gen":"sha256:2e97312cea5e86a4f4e11ea65cacfe9c4a63da245a1c446b00a2c6740234e92a","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","id":{"bin":"provenance","command":"compile"},"kwargs":"sha256:538d96b09cc32982dec744e3804baca080430296aa3922d1f2f41ba2e1db3fc7","body":"sha256:7368c92c5e771d7940d295619e9c7e479c4ce1429efc459968ee433f42202b0e"}
	root.AddCommand(newCompileCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"v":2,"tier":"owned","record":"prov-2026-adabd65f","gen":"sha256:4ae46f3089d333bd12c54eda92a2d20ae90a3ad1ebf7c65d3eda625c0356ee02","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","id":{"bin":"provenance","command":"version"},"kwargs":"sha256:005f3d08f0a921fb396d39e7d20c07e4b9db37f8fcea0735027b669bb1f413b1","body":"sha256:d1dd14dc04f492395835d41ece1523b584df00825247297fc37eb9aee3e7aa06"}
	root.AddCommand(newVersionCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"v":2,"tier":"owned","record":"prov-2026-d84dc10e","gen":"sha256:2e97312cea5e86a4f4e11ea65cacfe9c4a63da245a1c446b00a2c6740234e92a","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","id":{"bin":"provenance","command":"merge-manifest"},"kwargs":"sha256:897c107fa37d5c660df2ccc483d2118ad5f1bcf5a7858ac6747201b3b148c8b1","body":"sha256:a1db4c2c4018556c7782c29733de816f820778c4fde8e6f79c6d05321935ab48"}
	root.AddCommand(newMergeManifestCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"v":2,"tier":"owned","record":"prov-2026-7f96263a","gen":"sha256:d80ee2d48ee6cd63e3793ad5ad19e755f8443d038715bdaab730d623ed577073","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","id":{"bin":"provenance","command":"token"},"kwargs":"sha256:57545d27641decef3f27d3454cd635438c05c6ac5c799de06452c69c1584424b","body":"sha256:5fec0ee3ed627876155a60248d4d62766224853f15da308cdd7cedbc5cfddc99"}
	root.AddCommand(newTokenCmd())
	// /sedum:registerCommand
	// sedum:registerCommand {"v":2,"tier":"owned","record":"prov-2026-36b81f10","gen":"sha256:da0a6a84bf97ac25549e05b29429583669ff3d83fc5d84d6e6f3383a7d3b06f8","pkg":"sha256:f769c9a5d9d50a8fe780086d0e8c77985476b765245e10101df65f52f7c23f70","id":{"bin":"provenance","command":"document"},"kwargs":"sha256:b487580e966a25087263ad732d40f1dd0911fbef89042c7787e3a4b6d80762a1","body":"sha256:b7057985489bf05172c96b50e3c6d6e0c1807d9347bbb42ec8dbdb72ad0423a9"}
	root.AddCommand(newDocumentCmd())
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
