package remote

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-571ab1fb","gen":"sha256:625acefb954b89538195c6b98780772252c9f6d4b7e2e47533d4e8ef3a3d9a34","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"remote/import.go","path":"context"}}
import "context"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-571ab1fb","gen":"sha256:625acefb954b89538195c6b98780772252c9f6d4b7e2e47533d4e8ef3a3d9a34","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"remote/import.go","path":"errors"}}
import "errors"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-571ab1fb","gen":"sha256:625acefb954b89538195c6b98780772252c9f6d4b7e2e47533d4e8ef3a3d9a34","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"remote/import.go","path":"fmt"}}
import "fmt"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-571ab1fb","gen":"sha256:625acefb954b89538195c6b98780772252c9f6d4b7e2e47533d4e8ef3a3d9a34","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"remote/import.go","path":"github.com/livecodelife/provenance/record"}}
import "github.com/livecodelife/provenance/record"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setType {"tier":"owned","record":"prov-2026-1692f078","gen":"sha256:de6e711b2bc59c460bd4d12c533eba3701a4d7190b4c1aaaee734f84a022ff11","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type ImportPlan struct {\n\tItems   []ImportItem\n\tRefused []string\n}","file":"remote/import.go","name":"ImportPlan"}}
type ImportPlan struct {
	Items   []ImportItem
	Refused []string
}

// /sedum:setType
// sedum:setFunction {"tier":"owned","record":"prov-2026-571ab1fb","gen":"sha256:625acefb954b89538195c6b98780772252c9f6d4b7e2e47533d4e8ef3a3d9a34","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func PlanImport(ctx context.Context, local record.Store, ids []string, c *Client) (*ImportPlan, error) {\n\tplan := &ImportPlan{}\n\tfor _, id := range ids {\n\t\trec, _, err := local.Load(id)\n\t\tif err != nil {\n\t\t\treturn nil, err\n\t\t}\n\t\tlocalVersions, err := local.Versions(id)\n\t\tif err != nil {\n\t\t\treturn nil, err\n\t\t}\n\t\tremoteVersions, err := c.ListVersions(ctx, id)\n\t\tif err != nil {\n\t\t\tif errors.Is(err, ErrNotPublished) {\n\t\t\t\tcontinue\n\t\t\t}\n\t\t\treturn nil, err\n\t\t}\n\t\titem := ImportItem{ID: id, Title: rec.Title}\n\t\trefusal := \"\"\n\t\tfor _, v := range remoteVersions {\n\t\t\tif v.Version <= len(localVersions) {\n\t\t\t\tlocalSnap, err := local.Snapshot(id, v.Version)\n\t\t\t\tif err != nil {\n\t\t\t\t\treturn nil, err\n\t\t\t\t}\n\t\t\t\tif localSnap.ContentHash == v.ContentHash {\n\t\t\t\t\tcontinue\n\t\t\t\t}\n\t\t\t\trefusal = fmt.Sprintf(\"%s: version %d differs between the server and this repository (the local record moved past the remote base); revise it locally\", id, v.Version)\n\t\t\t\tbreak\n\t\t\t}\n\t\t\tsnap, err := v.Snapshot()\n\t\t\tif err != nil {\n\t\t\t\treturn nil, err\n\t\t\t}\n\t\t\titem.Snaps = append(item.Snaps, snap)\n\t\t\titem.Origins = append(item.Origins, v.Origin)\n\t\t}\n\t\tif refusal != \"\" {\n\t\t\tplan.Refused = append(plan.Refused, refusal)\n\t\t\tcontinue\n\t\t}\n\t\tif len(item.Snaps) > 0 {\n\t\t\tplan.Items = append(plan.Items, item)\n\t\t}\n\t}\n\treturn plan, nil\n}","file":"remote/import.go","name":"PlanImport"}}
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
// sedum:setFunction {"tier":"owned","record":"prov-2026-5ba86fb3","gen":"sha256:3364ea82757de634569eda416e5be36bb245680bd56a634af1c67e3e2fa5077c","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func ApplyImport(p *ImportPlan, local record.Store) error {\n\tfor _, item := range p.Items {\n\t\tcur, doc, err := local.Load(item.ID)\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n\t\tfor _, snap := range item.Snaps {\n\t\t\tif err := local.Append(snap); err != nil {\n\t\t\t\treturn err\n\t\t\t}\n\t\t}\n\t\tif len(item.Snaps) == 0 {\n\t\t\tcontinue\n\t\t}\n\t\tlast := item.Snaps[len(item.Snaps)-1]\n\t\tcopied := *last.Content\n\t\tcopied.ID = item.ID\n\t\tcopied.Schema = 2\n\t\tcopied.Version = last.Version\n\t\tcopied.Reason = last.Reason\n\t\tcopied.ReasonKind = last.ReasonKind\n\t\tcopied.Status = cur.Status\n\t\tif cur.Status == record.StatusImplemented {\n\t\t\tcopied.Status = record.StatusOpen\n\t\t}\n\t\trunCmds := make(map[string]string, len(cur.AssociatedSpecs))\n\t\tfor _, s := range cur.AssociatedSpecs {\n\t\t\trunCmds[s.Path] = s.RunCommand\n\t\t}\n\t\tspecs := make([]record.AssociatedSpec, 0, len(last.Content.AssociatedSpecs))\n\t\tfor _, s := range last.Content.AssociatedSpecs {\n\t\t\ts.RunCommand = runCmds[s.Path]\n\t\t\tspecs = append(specs, s)\n\t\t}\n\t\tcopied.AssociatedSpecs = specs\n\t\tif err := local.Save(&copied, doc); err != nil {\n\t\t\treturn fmt.Errorf(\"%s: versions written but the record file was not saved: %w\", item.ID, err)\n\t\t}\n\t}\n\treturn nil\n}","file":"remote/import.go","name":"ApplyImport"}}
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
// sedum:setType {"tier":"owned","record":"prov-2026-6b9ee651","gen":"sha256:f4f707525e500a11e8bdbf8eb554e40032540457a926167d3e8e4677da5de1e8","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type ImportItem struct {\n\tID      string\n\tTitle   string\n\tSnaps   []*record.Snapshot\n\tOrigins []string\n}","file":"remote/import.go","name":"ImportItem"}}
type ImportItem struct {
	ID      string
	Title   string
	Snaps   []*record.Snapshot
	Origins []string
}

// /sedum:setType
