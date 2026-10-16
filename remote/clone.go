package remote

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-59f0401e","gen":"sha256:d0d148bfaf34cdd4b7939607576df7e9948cff90ed6056c6bd54951a5f24d9db","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"remote/clone.go","path":"context"}}
import "context"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-59f0401e","gen":"sha256:d0d148bfaf34cdd4b7939607576df7e9948cff90ed6056c6bd54951a5f24d9db","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"remote/clone.go","path":"fmt"}}
import "fmt"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-59f0401e","gen":"sha256:d0d148bfaf34cdd4b7939607576df7e9948cff90ed6056c6bd54951a5f24d9db","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"remote/clone.go","path":"github.com/livecodelife/provenance/record"}}
import "github.com/livecodelife/provenance/record"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"tier":"owned","record":"prov-2026-59f0401e","gen":"sha256:d0d148bfaf34cdd4b7939607576df7e9948cff90ed6056c6bd54951a5f24d9db","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"// CloneRecords writes the caller's published records into dst: every version\n// snapshot (each verified against its content hash) and the current record\n// file. ids selects records; empty means all of them. It returns the number\n// of records written.\nfunc CloneRecords(ctx context.Context, c *Client, dst *record.FilesStore, ids []string) (int, error) {\n\tif len(ids) == 0 {\n\t\theads, err := c.ListRecords(ctx)\n\t\tif err != nil {\n\t\t\treturn 0, err\n\t\t}\n\t\tfor _, h := range heads {\n\t\t\tids = append(ids, h.ID)\n\t\t}\n\t}\n\tfor _, id := range ids {\n\t\thead, err := c.GetRecord(ctx, id)\n\t\tif err != nil {\n\t\t\treturn 0, fmt.Errorf(\"%s: %w\", id, err)\n\t\t}\n\t\tversions, err := c.ListVersions(ctx, id)\n\t\tif err != nil {\n\t\t\treturn 0, fmt.Errorf(\"%s: %w\", id, err)\n\t\t}\n\t\tfor i := range versions {\n\t\t\tsnap, err := versions[i].PublishedVersion.Snapshot()\n\t\t\tif err != nil {\n\t\t\t\treturn 0, err\n\t\t\t}\n\t\t\tif err := dst.Append(snap); err != nil {\n\t\t\t\treturn 0, err\n\t\t\t}\n\t\t}\n\t\tif len(versions) == 0 {\n\t\t\treturn 0, fmt.Errorf(\"%s: the server has no versions of it\", id)\n\t\t}\n\t\tlast := versions[len(versions)-1]\n\t\tr := head.Content.Record()\n\t\tr.ID, r.Status, r.Version = id, record.Status(head.Status), head.Version\n\t\tr.Reason, r.ReasonKind = last.Reason, last.ReasonKind\n\t\tif err := dst.Save(r, nil); err != nil {\n\t\t\treturn 0, err\n\t\t}\n\t}\n\treturn len(ids), nil\n}","file":"remote/clone.go","name":"CloneRecords"}}
// CloneRecords writes the caller's published records into dst: every version
// snapshot (each verified against its content hash) and the current record
// file. ids selects records; empty means all of them. It returns the number
// of records written.
func CloneRecords(ctx context.Context, c *Client, dst *record.FilesStore, ids []string) (int, error) {
	if len(ids) == 0 {
		heads, err := c.ListRecords(ctx)
		if err != nil {
			return 0, err
		}
		for _, h := range heads {
			ids = append(ids, h.ID)
		}
	}
	for _, id := range ids {
		head, err := c.GetRecord(ctx, id)
		if err != nil {
			return 0, fmt.Errorf("%s: %w", id, err)
		}
		versions, err := c.ListVersions(ctx, id)
		if err != nil {
			return 0, fmt.Errorf("%s: %w", id, err)
		}
		for i := range versions {
			snap, err := versions[i].PublishedVersion.Snapshot()
			if err != nil {
				return 0, err
			}
			if err := dst.Append(snap); err != nil {
				return 0, err
			}
		}
		if len(versions) == 0 {
			return 0, fmt.Errorf("%s: the server has no versions of it", id)
		}
		last := versions[len(versions)-1]
		r := head.Content.Record()
		r.ID, r.Status, r.Version = id, record.Status(head.Status), head.Version
		r.Reason, r.ReasonKind = last.Reason, last.ReasonKind
		if err := dst.Save(r, nil); err != nil {
			return 0, err
		}
	}
	return len(ids), nil
}

// /sedum:setFunction
