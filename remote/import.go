package remote

// sedum:anchor:imports
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-571ab1fb","gen":"sha256:625acefb954b89538195c6b98780772252c9f6d4b7e2e47533d4e8ef3a3d9a34","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"remote/import.go","path":"context"},"kwargs":"sha256:c57458050b68687cd1e419c594038b60b7198622b32c601dbadbad67e9b40fa8","body":"sha256:0f7443fb25d7ad6a970617d4d0df6123283e24801b94c1d84fbe431230f31e1e"}
import "context"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-571ab1fb","gen":"sha256:625acefb954b89538195c6b98780772252c9f6d4b7e2e47533d4e8ef3a3d9a34","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"remote/import.go","path":"errors"},"kwargs":"sha256:028cac4c57bf976b8f6ce95a1f85dbe887d8108554a53077550a7082fff226b6","body":"sha256:2b718237ba0a56467c8fd494cc131e942ef410e6228c03b644d70476c71862dc"}
import "errors"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-571ab1fb","gen":"sha256:625acefb954b89538195c6b98780772252c9f6d4b7e2e47533d4e8ef3a3d9a34","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"remote/import.go","path":"fmt"},"kwargs":"sha256:394568f1e2ff9b9a0ccfef2bcf81bbc82b47a4c6e0e89354ec6a519a80ff65dd","body":"sha256:bcbac51883cbf8064e9196b371cb880cee429866652447ad557533d817857609"}
import "fmt"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-571ab1fb","gen":"sha256:625acefb954b89538195c6b98780772252c9f6d4b7e2e47533d4e8ef3a3d9a34","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"remote/import.go","path":"github.com/livecodelife/provenance/record"},"kwargs":"sha256:2172942ee9a3a509c150bf6a759267a5f14a10571e1ea587c26e154a4e0e201a","body":"sha256:acbf8628061291d11aa39e9ea64ecc91095a117057cfe6f1958ef2c437fac003"}
import "github.com/livecodelife/provenance/record"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setType {"v":2,"tier":"owned","record":"prov-2026-1692f078","gen":"sha256:de6e711b2bc59c460bd4d12c533eba3701a4d7190b4c1aaaee734f84a022ff11","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","id":{"file":"remote/import.go","name":"ImportPlan"},"kwargs":"sha256:8c213b308c4dcf93adb51d0484f9e46e95394e1cc3382b9f787914470c5e9112","body":"sha256:b207dd4910833dc073ff2639f68177c64dad2b7c5050932df3886a3c6abaf5c3"}
type ImportPlan struct {
	Items   []ImportItem
	Refused []string
}

// /sedum:setType
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-571ab1fb","gen":"sha256:625acefb954b89538195c6b98780772252c9f6d4b7e2e47533d4e8ef3a3d9a34","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"remote/import.go","name":"PlanImport"},"kwargs":"sha256:77922316b10cd64dec544cca34241d326c14019f0fcd4740c47496e6a661a881","body":"sha256:cba93bcc0ab3430c2d9cf67b3dccf5987ba82e4f25f3554c51c83ec7fd012c3a"}
func PlanImport(ctx context.Context, local record.Store, ids []string, c *Client) (*ImportPlan, error) {
	plan := &ImportPlan{}
	for _, id := range ids {
		rec, _, err := local.Load(id)
		if err != nil {
			return nil, err
		}
		localVersions, err := local.Versions(id)
		if err != nil {
			return nil, err
		}
		remoteVersions, err := c.ListVersions(ctx, id)
		if err != nil {
			if errors.Is(err, ErrNotPublished) {
				continue
			}
			return nil, err
		}
		item := ImportItem{ID: id, Title: rec.Title}
		refusal := ""
		for _, v := range remoteVersions {
			if v.Version <= len(localVersions) {
				localSnap, err := local.Snapshot(id, v.Version)
				if err != nil {
					return nil, err
				}
				if localSnap.ContentHash == v.ContentHash {
					continue
				}
				refusal = fmt.Sprintf("%s: version %d differs between the server and this repository (the local record moved past the remote base); revise it locally", id, v.Version)
				break
			}
			snap, err := v.Snapshot()
			if err != nil {
				return nil, err
			}
			item.Snaps = append(item.Snaps, snap)
			item.Origins = append(item.Origins, v.Origin)
		}
		if refusal != "" {
			plan.Refused = append(plan.Refused, refusal)
			continue
		}
		if len(item.Snaps) > 0 {
			plan.Items = append(plan.Items, item)
		}
	}
	return plan, nil
}

// /sedum:setFunction
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-5ba86fb3","gen":"sha256:3364ea82757de634569eda416e5be36bb245680bd56a634af1c67e3e2fa5077c","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"remote/import.go","name":"ApplyImport"},"kwargs":"sha256:06c2114eb6c15a9e0344d6a8589790d666eced44c0efc1982138f33c9d62be39","body":"sha256:8ba4b7f7e2cf7bbb99896f68125d851869a3f798c15e174370f1876272957afb"}
func ApplyImport(p *ImportPlan, local record.Store) error {
	for _, item := range p.Items {
		cur, doc, err := local.Load(item.ID)
		if err != nil {
			return err
		}
		for _, snap := range item.Snaps {
			if err := local.Append(snap); err != nil {
				return err
			}
		}
		if len(item.Snaps) == 0 {
			continue
		}
		last := item.Snaps[len(item.Snaps)-1]
		copied := *last.Content
		copied.ID = item.ID
		copied.Schema = 2
		copied.Version = last.Version
		copied.Reason = last.Reason
		copied.ReasonKind = last.ReasonKind
		copied.Status = cur.Status
		if cur.Status == record.StatusImplemented {
			copied.Status = record.StatusOpen
		}
		runCmds := make(map[string]string, len(cur.AssociatedSpecs))
		for _, s := range cur.AssociatedSpecs {
			runCmds[s.Path] = s.RunCommand
		}
		specs := make([]record.AssociatedSpec, 0, len(last.Content.AssociatedSpecs))
		for _, s := range last.Content.AssociatedSpecs {
			s.RunCommand = runCmds[s.Path]
			specs = append(specs, s)
		}
		copied.AssociatedSpecs = specs
		if err := local.Save(&copied, doc); err != nil {
			return fmt.Errorf("%s: versions written but the record file was not saved: %w", item.ID, err)
		}
	}
	return nil
}

// /sedum:setFunction
// sedum:setType {"v":2,"tier":"owned","record":"prov-2026-6b9ee651","gen":"sha256:f4f707525e500a11e8bdbf8eb554e40032540457a926167d3e8e4677da5de1e8","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","id":{"file":"remote/import.go","name":"ImportItem"},"kwargs":"sha256:32fe728c52ecaf381e5289cfb423e2c2e5acaa008eb68b0693c42d3cc16ad7a4","body":"sha256:f8e338877bd5b23e693fec4dc87f427b2a093bbb2a06bbc1a2b0a904017447b8"}
type ImportItem struct {
	ID      string
	Title   string
	Snaps   []*record.Snapshot
	Origins []string
}

// /sedum:setType
