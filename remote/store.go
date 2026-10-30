package remote

// sedum:anchor:imports
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-ba01059c","gen":"sha256:cfb1e5f04fe72ca7a9ab21435377aeebb7282ecbb9616bbef3fd963ab9594528","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"remote/store.go","path":"context"},"kwargs":"sha256:8b2a00954da88a94ed260919200dda2bc1d2930060ef8ca6eca31548205aa7bd","body":"sha256:0f7443fb25d7ad6a970617d4d0df6123283e24801b94c1d84fbe431230f31e1e"}
import "context"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-3c6f9961","gen":"sha256:2868a704194a74ada9af2c9cb6a49495db773909d01d7ebb9535ec42c937003f","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"remote/store.go","path":"github.com/livecodelife/provenance/record"},"kwargs":"sha256:5965184d49cf7b94993f731d9018e15632804b354778950b7d3bef6ab4b2dd6e","body":"sha256:acbf8628061291d11aa39e9ea64ecc91095a117057cfe6f1958ef2c437fac003"}
import "github.com/livecodelife/provenance/record"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-6dd26112","gen":"sha256:0cb0771a2afb7228e1f8b956134fecff4e853aa5894dae21b185d5cc1f052e9b","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"remote/store.go","path":"sort"},"kwargs":"sha256:fb36c17a1670900e0dd307e19dcf7077f47a64435cc998b115e83e6d070d0073","body":"sha256:f197172617ae21ef23b1855b09aafb14571020ae81bff877d63d3ff81f4bf63c"}
import "sort"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setType {"v":2,"tier":"owned","record":"prov-2026-ba01059c","gen":"sha256:cfb1e5f04fe72ca7a9ab21435377aeebb7282ecbb9616bbef3fd963ab9594528","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","id":{"file":"remote/store.go","name":"Store"},"kwargs":"sha256:31113b3a061da9640bdff95ecaddd65e6e5d0cdcf0b5470f9c9869cf651b805e","body":"sha256:002518df4c8b4185ab833aab1adb31d989c007a5a2f719e9a689184de7696884"}
type Store struct {
	Client *Client
	Ctx    context.Context
}

// /sedum:setType
// sedum:setMethod {"v":2,"tier":"owned","record":"prov-2026-3c6f9961","gen":"sha256:2868a704194a74ada9af2c9cb6a49495db773909d01d7ebb9535ec42c937003f","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","id":{"file":"remote/store.go","name":"Append","receiver":"Store"},"kwargs":"sha256:54b2eafc05dfe1504d96ae24d2948a03aa54a33eeb6af556d19e1e8d5f440be2","body":"sha256:ec801d34429ce66e32064e16a597fa164f8b0008f5a61c2f690413abad4fa3ae"}
func (s *Store) Append(snap *record.Snapshot) error {
	head, err := s.Client.GetRecord(s.ctx(), snap.ID)
	var expected *string
	if err != nil {
		if err != ErrNotPublished {
			return err
		}
		expected = nil
	} else {
		expected = &head.CID
	}
	req := VersionRequest(snap, expected, "", "", "")
	_, _, err = s.Client.PutVersion(s.ctx(), snap.ID, req)
	return err
}

// /sedum:setMethod
// sedum:setMethod {"v":2,"tier":"owned","record":"prov-2026-5586fce0","gen":"sha256:6ce7d8a73c60555ec44ba0d6013e73d7615ceae5e8a0d812ab0685a3896fc8b5","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","id":{"file":"remote/store.go","name":"Specs","receiver":"Store"},"kwargs":"sha256:190a49d6a3de8a451560485bc3f1b78003482deb730355e52bafedbe55147f89","body":"sha256:856004834ee4231905cce248d8027d19b15f556c214bc1fc728e129bf1d3b059"}
func (s *Store) Specs() record.SpecResolver {
	return func(string) (string, bool) {
		return "", false
	}
}

// /sedum:setMethod
// sedum:setMethod {"v":2,"tier":"owned","record":"prov-2026-6dd26112","gen":"sha256:0cb0771a2afb7228e1f8b956134fecff4e853aa5894dae21b185d5cc1f052e9b","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","id":{"file":"remote/store.go","name":"List","receiver":"Store"},"kwargs":"sha256:fbf0dc2f79927b2f09acd9d8dd5af579ae41648e0d216172d489e7ac72ce5ae8","body":"sha256:9c968032c3155f5954033c11da92bfc03cad32e8cf907fd5fc328bf1b8eb4459"}
func (s *Store) List() ([]string, error) {
	heads, err := s.Client.ListRecords(s.ctx())
	if err != nil {
		return nil, err
	}
	ids := make([]string, 0, len(heads))
	for _, h := range heads {
		ids = append(ids, h.ID)
	}
	sort.Strings(ids)
	return ids, nil
}

// /sedum:setMethod
// sedum:setMethod {"v":2,"tier":"owned","record":"prov-2026-769879ec","gen":"sha256:2942c269f387dc23dd83dda9d8cfc85ba5db1062d0312a7546b287b191b8110a","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","id":{"file":"remote/store.go","name":"Versions","receiver":"Store"},"kwargs":"sha256:06574c2a1f1fb71dc957ebff19952af72caa619fa5ed91227f8698e59a0d21df","body":"sha256:5f82d4de43253c8012db162e167bd6b66576c2315b4be08e818b37d56181bb59"}
func (s *Store) Versions(id string) ([]int, error) {
	vs, err := s.Client.ListVersions(s.ctx(), id)
	if err != nil {
		return nil, err
	}
	out := make([]int, 0, len(vs))
	for _, v := range vs {
		out = append(out, v.Version)
	}
	return out, nil
}

// /sedum:setMethod
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-ae2d7d30","gen":"sha256:a30d12be018d95b843c2b1a618de0a6d11f6889586b6982ba8d5472209cade27","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"remote/store.go","name":"VersionRequest"},"kwargs":"sha256:14f970831d6eb4058e9c5f38a7113dda8a300895ce4247fc8a86875c1a4cbb27","body":"sha256:0af851f79cf6a5bda4667d1a37dcd15090bf92ecf3252d47b9f65a1ecec8df88"}
func VersionRequest(snap *record.Snapshot, expected *string, status, commit, project string) PutVersionRequest {
	pv := record.PublishedVersionOf(snap, record.PublishContext{})
	return PutVersionRequest{
		ExpectedHeadCid: expected,
		Version:         snap.Version,
		ContentHash:     snap.ContentHash,
		GenHash:         snap.Content.GenHash(),
		SpecHashes:      pv.SpecHashes,
		Content:         pv.Content,
		Status:          status,
		Reason:          snap.Reason,
		ReasonKind:      snap.ReasonKind,
		Author:          snap.Author,
		Commit:          commit,
		Project:         project,
		Origin:          "git",
	}
}

// /sedum:setFunction
// sedum:setMethod {"v":2,"tier":"owned","record":"prov-2026-bfe1318d","gen":"sha256:65505461e3bc765d4cdf9b85abfc1bbc4adc0f5260fb7796b625777999250625","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","id":{"file":"remote/store.go","name":"Save","receiver":"Store"},"kwargs":"sha256:b05a51176ecb13d5ed26bdecdb317efe3db98d39ddf3238ef018c7c4c60dec23","body":"sha256:19c03d6f25279fb888f8d29c57913b18cb808aa07609dccb7996a5ecd80def8b"}
func (*Store) Save(*record.Record, *record.Document) error {
	return ErrUnsupported
}

// /sedum:setMethod
// sedum:setMethod {"v":2,"tier":"owned","record":"prov-2026-c5ed6f2d","gen":"sha256:32b01e153c8635746c3e485a2629c0d363ba533fa386dc5aaea315693bbc47e2","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","id":{"file":"remote/store.go","name":"Load","receiver":"Store"},"kwargs":"sha256:2be4dc0374c089c07c7aa65bad352a26e2f80e219f7bf05596b07423df860829","body":"sha256:71435313d93e3787af5eaca017baefbb91f688ed6c98bc4bc3878e8f65b35eab"}
func (s *Store) Load(id string) (*record.Record, *record.Document, error) {
	head, err := s.Client.GetRecord(s.ctx(), id)
	if err != nil {
		return nil, nil, err
	}
	rec := head.Content.Record()
	rec.ID = id
	rec.Status = record.Status(head.Status)
	rec.Version = head.Version
	versions, err := s.Client.ListVersions(s.ctx(), id)
	if err == nil && len(versions) > 0 {
		last := versions[len(versions)-1]
		rec.Reason = last.Reason
		rec.ReasonKind = last.ReasonKind
	}
	return rec, nil, nil
}

// /sedum:setMethod
// sedum:setMethod {"v":2,"tier":"owned","record":"prov-2026-d4afa3bb","gen":"sha256:7b0d35b2f11640f3e09ab8950da7c41f95a6e0b8b9d34be330498089a1e7f900","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","id":{"file":"remote/store.go","name":"Snapshot","receiver":"Store"},"kwargs":"sha256:43c1459901d31f5c616eac694b140c2ddc33ae1b3c69c48c211dc7f68aeee35d","body":"sha256:ba8eca1340866c82f9b3a39d851aff337c87a85b49b26152638a5fe801d33af3"}
func (s *Store) Snapshot(id string, version int) (*record.Snapshot, error) {
	versions, err := s.Client.ListVersions(s.ctx(), id)
	if err != nil {
		return nil, err
	}
	for _, v := range versions {
		if v.Version == version {
			return v.PublishedVersion.Snapshot()
		}
	}
	return nil, ErrNotPublished
}

// /sedum:setMethod
// sedum:setMethod {"v":2,"tier":"owned","record":"prov-2026-dd5cb132","gen":"sha256:3d314b882b6739243f6c9a5464e296a4ad8aeb387f013566600d5e85623987f7","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","id":{"file":"remote/store.go","name":"Remove","receiver":"Store"},"kwargs":"sha256:c1ccb0bb410bd4d2a218c982cb57c54732b7f0554f008421bafe76d40fa85a74","body":"sha256:ad7629ad0a4a508a139acf67e7920e46831644362cccc600dd30cf5c46ef568b"}
func (*Store) Remove(string) error {
	return ErrUnsupported
}

// /sedum:setMethod
// sedum:setMethod {"v":2,"tier":"owned","record":"prov-2026-fc4134cb","gen":"sha256:82ad8da0bda278fde2ae04e5b798b42cfd6091ccac09ab00e4739e0c06c52a03","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","id":{"file":"remote/store.go","name":"ctx","receiver":"Store"},"kwargs":"sha256:9f3114d58cf563eb1d2927d6b7190d67c2d7a97f837410fd9b19b678eb133e95","body":"sha256:25705fca71a49a03eafc410ae0f8f17092159140b0b9c17bc507cb2a8da7fbc4"}
func (s *Store) ctx() context.Context {
	if s.Ctx != nil {
		return s.Ctx
	}
	return context.Background()
}

// /sedum:setMethod
