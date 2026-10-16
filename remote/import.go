package remote

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-59f0401e","gen":"sha256:d0d148bfaf34cdd4b7939607576df7e9948cff90ed6056c6bd54951a5f24d9db","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"remote/import.go","path":"context"}}
import "context"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-59f0401e","gen":"sha256:d0d148bfaf34cdd4b7939607576df7e9948cff90ed6056c6bd54951a5f24d9db","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"remote/import.go","path":"errors"}}
import "errors"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-59f0401e","gen":"sha256:d0d148bfaf34cdd4b7939607576df7e9948cff90ed6056c6bd54951a5f24d9db","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"remote/import.go","path":"fmt"}}
import "fmt"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-59f0401e","gen":"sha256:d0d148bfaf34cdd4b7939607576df7e9948cff90ed6056c6bd54951a5f24d9db","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"remote/import.go","path":"github.com/livecodelife/provenance/record"}}
import "github.com/livecodelife/provenance/record"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setType {"tier":"owned","record":"prov-2026-59f0401e","gen":"sha256:d0d148bfaf34cdd4b7939607576df7e9948cff90ed6056c6bd54951a5f24d9db","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"// ImportItem is what importing one record would do.\ntype ImportItem struct {\n\tID    string\n\tTitle string\n\t// Snaps are the remote versions beyond the local latest, oldest first.\n\tSnaps []*record.Snapshot\n\t// Origins and Authors describe each, for the report.\n\tOrigins []string\n}","file":"remote/import.go","name":"ImportItem"}}
// ImportItem is what importing one record would do.
type ImportItem struct {
	ID    string
	Title string
	// Snaps are the remote versions beyond the local latest, oldest first.
	Snaps []*record.Snapshot
	// Origins and Authors describe each, for the report.
	Origins []string
}

// /sedum:setType
// sedum:setType {"tier":"owned","record":"prov-2026-59f0401e","gen":"sha256:d0d148bfaf34cdd4b7939607576df7e9948cff90ed6056c6bd54951a5f24d9db","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"// ImportPlan lists what an import would bring in.\ntype ImportPlan struct {\n\tItems []ImportItem\n\t// Refused are records that cannot be imported, with the reason.\n\tRefused []string\n}","file":"remote/import.go","name":"ImportPlan"}}
// ImportPlan lists what an import would bring in.
type ImportPlan struct {
	Items []ImportItem
	// Refused are records that cannot be imported, with the reason.
	Refused []string
}

// /sedum:setType
// sedum:setFunction {"tier":"owned","record":"prov-2026-59f0401e","gen":"sha256:d0d148bfaf34cdd4b7939607576df7e9948cff90ed6056c6bd54951a5f24d9db","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"// PlanImport finds, for each local record, the remote versions that extend\n// it. A record whose local or remote history differs from the other's at a\n// shared version is refused, never merged.\nfunc PlanImport(ctx context.Context, local record.Store, ids []string, c *Client) (*ImportPlan, error) {\n\tplan := &ImportPlan{}\n\tfor _, id := range ids {\n\t\trec, _, err := local.Load(id)\n\t\tif err != nil {\n\t\t\treturn nil, err\n\t\t}\n\t\tversions, err := local.Versions(id)\n\t\tif err != nil {\n\t\t\treturn nil, err\n\t\t}\n\t\tremote, err := c.ListVersions(ctx, id)\n\t\tif errors.Is(err, ErrNotPublished) {\n\t\t\tcontinue\n\t\t}\n\t\tif err != nil {\n\t\t\treturn nil, err\n\t\t}\n\t\titem := ImportItem{ID: id, Title: rec.Title}\n\t\trefused := \"\"\n\t\tfor i := range remote {\n\t\t\trv := &remote[i]\n\t\t\tif rv.Version <= len(versions) {\n\t\t\t\tsnap, err := local.Snapshot(id, rv.Version)\n\t\t\t\tif err != nil {\n\t\t\t\t\treturn nil, err\n\t\t\t\t}\n\t\t\t\tif snap.ContentHash != rv.ContentHash {\n\t\t\t\t\trefused = fmt.Sprintf(\"%s: version %d differs between the server and this repository (the local record moved past the remote base); revise it locally\", id, rv.Version)\n\t\t\t\t\tbreak\n\t\t\t\t}\n\t\t\t\tcontinue\n\t\t\t}\n\t\t\tsnap, err := rv.PublishedVersion.Snapshot()\n\t\t\tif err != nil {\n\t\t\t\treturn nil, err\n\t\t\t}\n\t\t\titem.Snaps = append(item.Snaps, snap)\n\t\t\titem.Origins = append(item.Origins, rv.Origin)\n\t\t}\n\t\tswitch {\n\t\tcase refused != \"\":\n\t\t\tplan.Refused = append(plan.Refused, refused)\n\t\tcase len(item.Snaps) > 0:\n\t\t\tplan.Items = append(plan.Items, item)\n\t\t}\n\t}\n\treturn plan, nil\n}","file":"remote/import.go","name":"PlanImport"}}
// PlanImport finds, for each local record, the remote versions that extend
// it. A record whose local or remote history differs from the other's at a
// shared version is refused, never merged.
func PlanImport(ctx context.Context, local record.Store, ids []string, c *Client) (*ImportPlan, error) {
	plan := &ImportPlan{}
	for _, id := range ids {
		rec, _, err := local.Load(id)
		if err != nil {
			return nil, err
		}
		versions, err := local.Versions(id)
		if err != nil {
			return nil, err
		}
		remote, err := c.ListVersions(ctx, id)
		if errors.Is(err, ErrNotPublished) {
			continue
		}
		if err != nil {
			return nil, err
		}
		item := ImportItem{ID: id, Title: rec.Title}
		refused := ""
		for i := range remote {
			rv := &remote[i]
			if rv.Version <= len(versions) {
				snap, err := local.Snapshot(id, rv.Version)
				if err != nil {
					return nil, err
				}
				if snap.ContentHash != rv.ContentHash {
					refused = fmt.Sprintf("%s: version %d differs between the server and this repository (the local record moved past the remote base); revise it locally", id, rv.Version)
					break
				}
				continue
			}
			snap, err := rv.PublishedVersion.Snapshot()
			if err != nil {
				return nil, err
			}
			item.Snaps = append(item.Snaps, snap)
			item.Origins = append(item.Origins, rv.Origin)
		}
		switch {
		case refused != "":
			plan.Refused = append(plan.Refused, refused)
		case len(item.Snaps) > 0:
			plan.Items = append(plan.Items, item)
		}
	}
	return plan, nil
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-59f0401e","gen":"sha256:d0d148bfaf34cdd4b7939607576df7e9948cff90ed6056c6bd54951a5f24d9db","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"// ApplyImport appends the planned snapshots to the local store and updates\n// each current record file to the newest imported content. A record that was\n// implemented returns to open, as any new version does.\nfunc ApplyImport(p *ImportPlan, local record.Store) error {\n\tfor _, it := range p.Items {\n\t\tcur, doc, err := local.Load(it.ID)\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n\t\tfor _, snap := range it.Snaps {\n\t\t\tif err := local.Append(snap); err != nil {\n\t\t\t\treturn err\n\t\t\t}\n\t\t}\n\t\tlast := it.Snaps[len(it.Snaps)-1]\n\t\tnext := *last.Content\n\t\tnext.ID, next.Schema = it.ID, 2\n\t\tnext.Version, next.Reason, next.ReasonKind = last.Version, last.Reason, last.ReasonKind\n\t\tnext.Status = cur.Status\n\t\tif next.Status == record.StatusImplemented {\n\t\t\tnext.Status = record.StatusOpen\n\t\t}\n\t\truns := map[string]string{}\n\t\tfor _, s := range cur.AssociatedSpecs {\n\t\t\truns[s.Path] = s.RunCommand\n\t\t}\n\t\tnext.AssociatedSpecs = append([]record.AssociatedSpec(nil), last.Content.AssociatedSpecs...)\n\t\tfor i := range next.AssociatedSpecs {\n\t\t\tnext.AssociatedSpecs[i].RunCommand = runs[next.AssociatedSpecs[i].Path]\n\t\t}\n\t\tif err := local.Save(&next, doc); err != nil {\n\t\t\treturn fmt.Errorf(\"%s: versions written but the record file was not saved: %w\", it.ID, err)\n\t\t}\n\t}\n\treturn nil\n}","file":"remote/import.go","name":"ApplyImport"}}
// ApplyImport appends the planned snapshots to the local store and updates
// each current record file to the newest imported content. A record that was
// implemented returns to open, as any new version does.
func ApplyImport(p *ImportPlan, local record.Store) error {
	for _, it := range p.Items {
		cur, doc, err := local.Load(it.ID)
		if err != nil {
			return err
		}
		for _, snap := range it.Snaps {
			if err := local.Append(snap); err != nil {
				return err
			}
		}
		last := it.Snaps[len(it.Snaps)-1]
		next := *last.Content
		next.ID, next.Schema = it.ID, 2
		next.Version, next.Reason, next.ReasonKind = last.Version, last.Reason, last.ReasonKind
		next.Status = cur.Status
		if next.Status == record.StatusImplemented {
			next.Status = record.StatusOpen
		}
		runs := map[string]string{}
		for _, s := range cur.AssociatedSpecs {
			runs[s.Path] = s.RunCommand
		}
		next.AssociatedSpecs = append([]record.AssociatedSpec(nil), last.Content.AssociatedSpecs...)
		for i := range next.AssociatedSpecs {
			next.AssociatedSpecs[i].RunCommand = runs[next.AssociatedSpecs[i].Path]
		}
		if err := local.Save(&next, doc); err != nil {
			return fmt.Errorf("%s: versions written but the record file was not saved: %w", it.ID, err)
		}
	}
	return nil
}

// /sedum:setFunction
