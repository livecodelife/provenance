package remote

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-59f0401e","gen":"sha256:ca4c401c192d7e9987e539d842f6a9b00dcf854221fdfc13c58a58030802de73","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"remote/clone.go","path":"context"}}
import "context"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-59f0401e","gen":"sha256:ca4c401c192d7e9987e539d842f6a9b00dcf854221fdfc13c58a58030802de73","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"remote/clone.go","path":"fmt"}}
import "fmt"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-59f0401e","gen":"sha256:ca4c401c192d7e9987e539d842f6a9b00dcf854221fdfc13c58a58030802de73","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"remote/clone.go","path":"github.com/livecodelife/provenance/record"}}
import "github.com/livecodelife/provenance/record"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"tier":"owned","record":"prov-2026-59f0401e","gen":"sha256:ca4c401c192d7e9987e539d842f6a9b00dcf854221fdfc13c58a58030802de73","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func CloneRecords(ctx context.Context, c *Client, dst *record.FilesStore, ids []string) (int, error) {\n\tif len(ids) == 0 {\n\t\theads, err := c.ListRecords(ctx)\n\t\tif err != nil {\n\t\t\treturn 0, err\n\t\t}\n\t\tfor _, h := range heads {\n\t\t\tids = append(ids, h.ID)\n\t\t}\n\t}\n\n\tfor _, id := range ids {\n\t\thead, err := c.GetRecord(ctx, id)\n\t\tif err != nil {\n\t\t\treturn 0, fmt.Errorf(\"%s: %w\", id, err)\n\t\t}\n\n\t\tversions, err := c.ListVersions(ctx, id)\n\t\tif err != nil {\n\t\t\treturn 0, fmt.Errorf(\"%s: %w\", id, err)\n\t\t}\n\n\t\tfor _, v := range versions {\n\t\t\tsnap, err := v.Snapshot()\n\t\t\tif err != nil {\n\t\t\t\treturn 0, err\n\t\t\t}\n\t\t\tif err := dst.Append(snap); err != nil {\n\t\t\t\treturn 0, err\n\t\t\t}\n\t\t}\n\n\t\tif len(versions) == 0 {\n\t\t\treturn 0, fmt.Errorf(\"%s: the server has no versions of it\", id)\n\t\t}\n\n\t\tnewest := versions[len(versions)-1]\n\t\trec := head.Content.Record()\n\t\trec.ID = id\n\t\trec.Status = record.Status(head.Status)\n\t\trec.Version = head.Version\n\t\trec.Reason = newest.Reason\n\t\trec.ReasonKind = newest.ReasonKind\n\n\t\tif err := dst.Save(rec, nil); err != nil {\n\t\t\treturn 0, err\n\t\t}\n\t}\n\n\treturn len(ids), nil\n}","file":"remote/clone.go","name":"CloneRecords"}}
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

		for _, v := range versions {
			snap, err := v.Snapshot()
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

		newest := versions[len(versions)-1]
		rec := head.Content.Record()
		rec.ID = id
		rec.Status = record.Status(head.Status)
		rec.Version = head.Version
		rec.Reason = newest.Reason
		rec.ReasonKind = newest.ReasonKind

		if err := dst.Save(rec, nil); err != nil {
			return 0, err
		}
	}

	return len(ids), nil
}

// /sedum:setFunction
