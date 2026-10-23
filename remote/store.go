package remote

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-ba01059c","gen":"sha256:cfb1e5f04fe72ca7a9ab21435377aeebb7282ecbb9616bbef3fd963ab9594528","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"remote/store.go","path":"context"}}
import "context"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-3c6f9961","gen":"sha256:2868a704194a74ada9af2c9cb6a49495db773909d01d7ebb9535ec42c937003f","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"remote/store.go","path":"github.com/livecodelife/provenance/record"}}
import "github.com/livecodelife/provenance/record"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-6dd26112","gen":"sha256:0cb0771a2afb7228e1f8b956134fecff4e853aa5894dae21b185d5cc1f052e9b","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"remote/store.go","path":"sort"}}
import "sort"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setType {"tier":"owned","record":"prov-2026-ba01059c","gen":"sha256:cfb1e5f04fe72ca7a9ab21435377aeebb7282ecbb9616bbef3fd963ab9594528","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type Store struct {\n\tClient *Client\n\tCtx    context.Context\n}","file":"remote/store.go","name":"Store"}}
type Store struct {
	Client *Client
	Ctx    context.Context
}

// /sedum:setType
// sedum:setMethod {"tier":"owned","record":"prov-2026-3c6f9961","gen":"sha256:2868a704194a74ada9af2c9cb6a49495db773909d01d7ebb9535ec42c937003f","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (s *Store) Append(snap *record.Snapshot) error {\n\thead, err := s.Client.GetRecord(s.ctx(), snap.ID)\n\tvar expected *string\n\tif err != nil {\n\t\tif err != ErrNotPublished {\n\t\t\treturn err\n\t\t}\n\t\texpected = nil\n\t} else {\n\t\texpected = &head.CID\n\t}\n\treq := VersionRequest(snap, expected, \"\", \"\", \"\")\n\t_, _, err = s.Client.PutVersion(s.ctx(), snap.ID, req)\n\treturn err\n}","file":"remote/store.go","name":"Append","receiver":"Store"}}
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
// sedum:setMethod {"tier":"owned","record":"prov-2026-5586fce0","gen":"sha256:6ce7d8a73c60555ec44ba0d6013e73d7615ceae5e8a0d812ab0685a3896fc8b5","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (s *Store) Specs() record.SpecResolver {\n\treturn func(string) (string, bool) {\n\t\treturn \"\", false\n\t}\n}","file":"remote/store.go","name":"Specs","receiver":"Store"}}
func (s *Store) Specs() record.SpecResolver {
	return func(string) (string, bool) {
		return "", false
	}
}

// /sedum:setMethod
// sedum:setMethod {"tier":"owned","record":"prov-2026-6dd26112","gen":"sha256:0cb0771a2afb7228e1f8b956134fecff4e853aa5894dae21b185d5cc1f052e9b","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (s *Store) List() ([]string, error) {\n\theads, err := s.Client.ListRecords(s.ctx())\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\tids := make([]string, 0, len(heads))\n\tfor _, h := range heads {\n\t\tids = append(ids, h.ID)\n\t}\n\tsort.Strings(ids)\n\treturn ids, nil\n}","file":"remote/store.go","name":"List","receiver":"Store"}}
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
// sedum:setMethod {"tier":"owned","record":"prov-2026-769879ec","gen":"sha256:2942c269f387dc23dd83dda9d8cfc85ba5db1062d0312a7546b287b191b8110a","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (s *Store) Versions(id string) ([]int, error) {\n\tvs, err := s.Client.ListVersions(s.ctx(), id)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\tout := make([]int, 0, len(vs))\n\tfor _, v := range vs {\n\t\tout = append(out, v.Version)\n\t}\n\treturn out, nil\n}","file":"remote/store.go","name":"Versions","receiver":"Store"}}
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
// sedum:setFunction {"tier":"owned","record":"prov-2026-ae2d7d30","gen":"sha256:a30d12be018d95b843c2b1a618de0a6d11f6889586b6982ba8d5472209cade27","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func VersionRequest(snap *record.Snapshot, expected *string, status, commit, project string) PutVersionRequest {\n\tpv := record.PublishedVersionOf(snap, record.PublishContext{})\n\treturn PutVersionRequest{\n\t\tExpectedHeadCid: expected,\n\t\tVersion:         snap.Version,\n\t\tContentHash:     snap.ContentHash,\n\t\tGenHash:         snap.Content.GenHash(),\n\t\tSpecHashes:      pv.SpecHashes,\n\t\tContent:         pv.Content,\n\t\tStatus:          status,\n\t\tReason:          snap.Reason,\n\t\tReasonKind:      snap.ReasonKind,\n\t\tAuthor:          snap.Author,\n\t\tCommit:          commit,\n\t\tProject:         project,\n\t\tOrigin:          \"git\",\n\t}\n}","file":"remote/store.go","name":"VersionRequest"}}
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
// sedum:setMethod {"tier":"owned","record":"prov-2026-bfe1318d","gen":"sha256:65505461e3bc765d4cdf9b85abfc1bbc4adc0f5260fb7796b625777999250625","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (*Store) Save(*record.Record, *record.Document) error {\n\treturn ErrUnsupported\n}","file":"remote/store.go","name":"Save","receiver":"Store"}}
func (*Store) Save(*record.Record, *record.Document) error {
	return ErrUnsupported
}

// /sedum:setMethod
// sedum:setMethod {"tier":"owned","record":"prov-2026-c5ed6f2d","gen":"sha256:32b01e153c8635746c3e485a2629c0d363ba533fa386dc5aaea315693bbc47e2","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (s *Store) Load(id string) (*record.Record, *record.Document, error) {\n\thead, err := s.Client.GetRecord(s.ctx(), id)\n\tif err != nil {\n\t\treturn nil, nil, err\n\t}\n\trec := head.Content.Record()\n\trec.ID = id\n\trec.Status = record.Status(head.Status)\n\trec.Version = head.Version\n\tversions, err := s.Client.ListVersions(s.ctx(), id)\n\tif err == nil && len(versions) > 0 {\n\t\tlast := versions[len(versions)-1]\n\t\trec.Reason = last.Reason\n\t\trec.ReasonKind = last.ReasonKind\n\t}\n\treturn rec, nil, nil\n}","file":"remote/store.go","name":"Load","receiver":"Store"}}
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
// sedum:setMethod {"tier":"owned","record":"prov-2026-d4afa3bb","gen":"sha256:7b0d35b2f11640f3e09ab8950da7c41f95a6e0b8b9d34be330498089a1e7f900","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (s *Store) Snapshot(id string, version int) (*record.Snapshot, error) {\n\tversions, err := s.Client.ListVersions(s.ctx(), id)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\tfor _, v := range versions {\n\t\tif v.Version == version {\n\t\t\treturn v.PublishedVersion.Snapshot()\n\t\t}\n\t}\n\treturn nil, ErrNotPublished\n}","file":"remote/store.go","name":"Snapshot","receiver":"Store"}}
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
// sedum:setMethod {"tier":"owned","record":"prov-2026-dd5cb132","gen":"sha256:3d314b882b6739243f6c9a5464e296a4ad8aeb387f013566600d5e85623987f7","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (*Store) Remove(string) error {\n\treturn ErrUnsupported\n}","file":"remote/store.go","name":"Remove","receiver":"Store"}}
func (*Store) Remove(string) error {
	return ErrUnsupported
}

// /sedum:setMethod
// sedum:setMethod {"tier":"owned","record":"prov-2026-fc4134cb","gen":"sha256:82ad8da0bda278fde2ae04e5b798b42cfd6091ccac09ab00e4739e0c06c52a03","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (s *Store) ctx() context.Context {\n\tif s.Ctx != nil {\n\t\treturn s.Ctx\n\t}\n\treturn context.Background()\n}","file":"remote/store.go","name":"ctx","receiver":"Store"}}
func (s *Store) ctx() context.Context {
	if s.Ctx != nil {
		return s.Ctx
	}
	return context.Background()
}

// /sedum:setMethod
