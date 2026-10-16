package remote

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-ba01059c","gen":"sha256:d5686b28170503679400dd6f3d84593b72ebee59ddd07fe1565ddfd8c513a6f3","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"remote/store.go","path":"context"}}
import "context"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-ba01059c","gen":"sha256:d5686b28170503679400dd6f3d84593b72ebee59ddd07fe1565ddfd8c513a6f3","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"remote/store.go","path":"sort"}}
import "sort"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-ba01059c","gen":"sha256:d5686b28170503679400dd6f3d84593b72ebee59ddd07fe1565ddfd8c513a6f3","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"remote/store.go","path":"github.com/livecodelife/provenance/record"}}
import "github.com/livecodelife/provenance/record"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setType {"tier":"owned","record":"prov-2026-ba01059c","gen":"sha256:d5686b28170503679400dd6f3d84593b72ebee59ddd07fe1565ddfd8c513a6f3","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"// Store is the Provisio backend of record.Store: reads come from the record\n// server, Append writes one version, and a remote record cannot be saved or\n// removed any other way.\ntype Store struct {\n\tClient *Client\n\tCtx    context.Context\n}","file":"remote/store.go","name":"Store"}}
// Store is the Provisio backend of record.Store: reads come from the record
// server, Append writes one version, and a remote record cannot be saved or
// removed any other way.
type Store struct {
	Client *Client
	Ctx    context.Context
}

// /sedum:setType
// sedum:setMethod {"tier":"owned","record":"prov-2026-ba01059c","gen":"sha256:d5686b28170503679400dd6f3d84593b72ebee59ddd07fe1565ddfd8c513a6f3","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (s *Store) ctx() context.Context {\n\tif s.Ctx != nil {\n\t\treturn s.Ctx\n\t}\n\treturn context.Background()\n}","file":"remote/store.go","name":"ctx","receiver":"Store"}}
func (s *Store) ctx() context.Context {
	if s.Ctx != nil {
		return s.Ctx
	}
	return context.Background()
}

// /sedum:setMethod
// sedum:setMethod {"tier":"owned","record":"prov-2026-ba01059c","gen":"sha256:d5686b28170503679400dd6f3d84593b72ebee59ddd07fe1565ddfd8c513a6f3","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (s *Store) List() ([]string, error) {\n\theads, err := s.Client.ListRecords(s.ctx())\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\tids := make([]string, 0, len(heads))\n\tfor _, h := range heads {\n\t\tids = append(ids, h.ID)\n\t}\n\tsort.Strings(ids)\n\treturn ids, nil\n}","file":"remote/store.go","name":"List","receiver":"Store"}}
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
// sedum:setMethod {"tier":"owned","record":"prov-2026-ba01059c","gen":"sha256:d5686b28170503679400dd6f3d84593b72ebee59ddd07fe1565ddfd8c513a6f3","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"// Load returns the record as its head and newest version describe it.\nfunc (s *Store) Load(id string) (*record.Record, *record.Document, error) {\n\thead, err := s.Client.GetRecord(s.ctx(), id)\n\tif err != nil {\n\t\treturn nil, nil, err\n\t}\n\tr := head.Content.Record()\n\tr.ID, r.Status, r.Version = id, record.Status(head.Status), head.Version\n\tif vs, err := s.Client.ListVersions(s.ctx(), id); err == nil && len(vs) > 0 {\n\t\tlast := vs[len(vs)-1]\n\t\tr.Reason, r.ReasonKind = last.Reason, last.ReasonKind\n\t}\n\treturn r, nil, nil\n}","file":"remote/store.go","name":"Load","receiver":"Store"}}
// Load returns the record as its head and newest version describe it.
func (s *Store) Load(id string) (*record.Record, *record.Document, error) {
	head, err := s.Client.GetRecord(s.ctx(), id)
	if err != nil {
		return nil, nil, err
	}
	r := head.Content.Record()
	r.ID, r.Status, r.Version = id, record.Status(head.Status), head.Version
	if vs, err := s.Client.ListVersions(s.ctx(), id); err == nil && len(vs) > 0 {
		last := vs[len(vs)-1]
		r.Reason, r.ReasonKind = last.Reason, last.ReasonKind
	}
	return r, nil, nil
}

// /sedum:setMethod
// sedum:setMethod {"tier":"owned","record":"prov-2026-ba01059c","gen":"sha256:d5686b28170503679400dd6f3d84593b72ebee59ddd07fe1565ddfd8c513a6f3","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (s *Store) Versions(id string) ([]int, error) {\n\tvs, err := s.Client.ListVersions(s.ctx(), id)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\tout := make([]int, 0, len(vs))\n\tfor _, v := range vs {\n\t\tout = append(out, v.Version)\n\t}\n\treturn out, nil\n}","file":"remote/store.go","name":"Versions","receiver":"Store"}}
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
// sedum:setMethod {"tier":"owned","record":"prov-2026-ba01059c","gen":"sha256:d5686b28170503679400dd6f3d84593b72ebee59ddd07fe1565ddfd8c513a6f3","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"// Snapshot reads one version and verifies its content hash.\nfunc (s *Store) Snapshot(id string, version int) (*record.Snapshot, error) {\n\tvs, err := s.Client.ListVersions(s.ctx(), id)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\tfor i := range vs {\n\t\tif vs[i].Version == version {\n\t\t\treturn vs[i].PublishedVersion.Snapshot()\n\t\t}\n\t}\n\treturn nil, ErrNotPublished\n}","file":"remote/store.go","name":"Snapshot","receiver":"Store"}}
// Snapshot reads one version and verifies its content hash.
func (s *Store) Snapshot(id string, version int) (*record.Snapshot, error) {
	vs, err := s.Client.ListVersions(s.ctx(), id)
	if err != nil {
		return nil, err
	}
	for i := range vs {
		if vs[i].Version == version {
			return vs[i].PublishedVersion.Snapshot()
		}
	}
	return nil, ErrNotPublished
}

// /sedum:setMethod
// sedum:setMethod {"tier":"owned","record":"prov-2026-ba01059c","gen":"sha256:d5686b28170503679400dd6f3d84593b72ebee59ddd07fe1565ddfd8c513a6f3","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"// Append writes snap as the next version of its record.\nfunc (s *Store) Append(snap *record.Snapshot) error {\n\tvar expected *string\n\thead, err := s.Client.GetRecord(s.ctx(), snap.ID)\n\tif err == nil {\n\t\texpected = &head.CID\n\t} else if err != ErrNotPublished {\n\t\treturn err\n\t}\n\t_, _, err = s.Client.PutVersion(s.ctx(), snap.ID, VersionRequest(snap, expected, \"\", \"\", \"\"))\n\treturn err\n}","file":"remote/store.go","name":"Append","receiver":"Store"}}
// Append writes snap as the next version of its record.
func (s *Store) Append(snap *record.Snapshot) error {
	var expected *string
	head, err := s.Client.GetRecord(s.ctx(), snap.ID)
	if err == nil {
		expected = &head.CID
	} else if err != ErrNotPublished {
		return err
	}
	_, _, err = s.Client.PutVersion(s.ctx(), snap.ID, VersionRequest(snap, expected, "", "", ""))
	return err
}

// /sedum:setMethod
// sedum:setMethod {"tier":"owned","record":"prov-2026-ba01059c","gen":"sha256:d5686b28170503679400dd6f3d84593b72ebee59ddd07fe1565ddfd8c513a6f3","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (s *Store) Save(*record.Record, *record.Document) error { return ErrUnsupported }","file":"remote/store.go","name":"Save","receiver":"Store"}}
func (s *Store) Save(*record.Record, *record.Document) error { return ErrUnsupported }

// /sedum:setMethod
// sedum:setMethod {"tier":"owned","record":"prov-2026-ba01059c","gen":"sha256:d5686b28170503679400dd6f3d84593b72ebee59ddd07fe1565ddfd8c513a6f3","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (s *Store) Remove(string) error { return ErrUnsupported }","file":"remote/store.go","name":"Remove","receiver":"Store"}}
func (s *Store) Remove(string) error { return ErrUnsupported }

// /sedum:setMethod
// sedum:setMethod {"tier":"owned","record":"prov-2026-ba01059c","gen":"sha256:d5686b28170503679400dd6f3d84593b72ebee59ddd07fe1565ddfd8c513a6f3","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"// Specs resolves nothing: a remote snapshot carries its spec hashes.\nfunc (s *Store) Specs() record.SpecResolver {\n\treturn func(string) (string, bool) { return \"\", false }\n}","file":"remote/store.go","name":"Specs","receiver":"Store"}}
// Specs resolves nothing: a remote snapshot carries its spec hashes.
func (s *Store) Specs() record.SpecResolver {
	return func(string) (string, bool) { return "", false }
}

// /sedum:setMethod
// sedum:setFunction {"tier":"owned","record":"prov-2026-ba01059c","gen":"sha256:d5686b28170503679400dd6f3d84593b72ebee59ddd07fe1565ddfd8c513a6f3","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"// VersionRequest builds the PUT body for snap. status, commit and project are\n// sent only when set.\nfunc VersionRequest(snap *record.Snapshot, expected *string, status, commit, project string) PutVersionRequest {\n\tpv := record.PublishedVersionOf(snap, record.PublishContext{})\n\treturn PutVersionRequest{\n\t\tExpectedHeadCid: expected, Version: snap.Version, ContentHash: snap.ContentHash,\n\t\tGenHash: snap.Content.GenHash(), SpecHashes: pv.SpecHashes, Content: pv.Content,\n\t\tStatus: status, Reason: snap.Reason, ReasonKind: snap.ReasonKind, Author: snap.Author,\n\t\tCommit: commit, Project: project, Origin: \"git\",\n\t}\n}","file":"remote/store.go","name":"VersionRequest"}}
// VersionRequest builds the PUT body for snap. status, commit and project are
// sent only when set.
func VersionRequest(snap *record.Snapshot, expected *string, status, commit, project string) PutVersionRequest {
	pv := record.PublishedVersionOf(snap, record.PublishContext{})
	return PutVersionRequest{
		ExpectedHeadCid: expected, Version: snap.Version, ContentHash: snap.ContentHash,
		GenHash: snap.Content.GenHash(), SpecHashes: pv.SpecHashes, Content: pv.Content,
		Status: status, Reason: snap.Reason, ReasonKind: snap.ReasonKind, Author: snap.Author,
		Commit: commit, Project: project, Origin: "git",
	}
}

// /sedum:setFunction
