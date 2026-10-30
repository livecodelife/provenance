package remote

// sedum:anchor:imports
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-59f0401e","gen":"sha256:ca4c401c192d7e9987e539d842f6a9b00dcf854221fdfc13c58a58030802de73","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"remote/clone.go","path":"context"},"kwargs":"sha256:a3bc7136e3b331d0168dc1f2c8b1468f6096917418952b615af1a2d1b1bd0c56","body":"sha256:0f7443fb25d7ad6a970617d4d0df6123283e24801b94c1d84fbe431230f31e1e"}
import "context"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-59f0401e","gen":"sha256:ca4c401c192d7e9987e539d842f6a9b00dcf854221fdfc13c58a58030802de73","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"remote/clone.go","path":"fmt"},"kwargs":"sha256:2004ebae880f2337b5dfda4c275418c3275ed398e7e282f5405b014ebe124d3e","body":"sha256:bcbac51883cbf8064e9196b371cb880cee429866652447ad557533d817857609"}
import "fmt"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-59f0401e","gen":"sha256:ca4c401c192d7e9987e539d842f6a9b00dcf854221fdfc13c58a58030802de73","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"remote/clone.go","path":"github.com/livecodelife/provenance/record"},"kwargs":"sha256:2aef6d9d70f199f5d9cccd6577365ee4febe9973fec19b6fa0058709af11ac23","body":"sha256:acbf8628061291d11aa39e9ea64ecc91095a117057cfe6f1958ef2c437fac003"}
import "github.com/livecodelife/provenance/record"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-59f0401e","gen":"sha256:ca4c401c192d7e9987e539d842f6a9b00dcf854221fdfc13c58a58030802de73","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"remote/clone.go","name":"CloneRecords"},"kwargs":"sha256:8ad7907c456d76b319d23fc9b063511d2e25baf71e125ed79e0fd0f1854fc700","body":"sha256:de3bc0d3bbd81eaedb33c9e27282700b8a76699023f341c3180bcbe8e138c613"}
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
