package remote

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-17d5965a","gen":"sha256:574500cb4ef448e7012a00efbbdafd37be7a068b37baf7219e427bd37401b9b7","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"remote/publish.go","path":"context"}}
import "context"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-17d5965a","gen":"sha256:574500cb4ef448e7012a00efbbdafd37be7a068b37baf7219e427bd37401b9b7","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"remote/publish.go","path":"errors"}}
import "errors"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-17d5965a","gen":"sha256:574500cb4ef448e7012a00efbbdafd37be7a068b37baf7219e427bd37401b9b7","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"remote/publish.go","path":"fmt"}}
import "fmt"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-17d5965a","gen":"sha256:574500cb4ef448e7012a00efbbdafd37be7a068b37baf7219e427bd37401b9b7","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"remote/publish.go","path":"github.com/livecodelife/provenance/record"}}
import "github.com/livecodelife/provenance/record"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-17d5965a","gen":"sha256:574500cb4ef448e7012a00efbbdafd37be7a068b37baf7219e427bd37401b9b7","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"remote/publish.go","path":"regexp"}}
import "regexp"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"tier":"owned","record":"prov-2026-17d5965a","gen":"sha256:574500cb4ef448e7012a00efbbdafd37be7a068b37baf7219e427bd37401b9b7","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"// statusRank orders the lifecycle: a head status may only move forward.\nfunc statusRank(s string) int {\n\tswitch s {\n\tcase \"draft\":\n\t\treturn 0\n\tcase \"open\":\n\t\treturn 1\n\tcase \"implemented\":\n\t\treturn 2\n\tcase \"deprecated\":\n\t\treturn 3\n\t}\n\treturn -1\n}","file":"remote/publish.go","name":"statusRank"}}
// statusRank orders the lifecycle: a head status may only move forward.
func statusRank(s string) int {
	switch s {
	case "draft":
		return 0
	case "open":
		return 1
	case "implemented":
		return 2
	case "deprecated":
		return 3
	}
	return -1
}

// /sedum:setFunction
// sedum:setType {"tier":"owned","record":"prov-2026-17d5965a","gen":"sha256:574500cb4ef448e7012a00efbbdafd37be7a068b37baf7219e427bd37401b9b7","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"// PublishItem is what publishing one record would write.\ntype PublishItem struct {\n\tID    string\n\tTitle string\n\t// Snaps are the versions to write, oldest first.\n\tSnaps []*record.Snapshot\n\t// Status is the local status; the last written version carries it.\n\tStatus string\n\t// HeadCID and HeadVersion describe the remote head, empty and 0 when unpublished.\n\tHeadCID     string\n\tHeadVersion int\n\t// HeadStatus is set when only the head needs updating (no new version).\n\tHeadStatus string\n\t// Source is where the record lives in git; every write carries it when set.\n\tSource *record.PublishedSource\n\t// SetProject and SetSource are set when the record has nothing new to write\n\t// but its head lacks the wanted project or source: a head-only update.\n\tSetProject string\n\tSetSource  *record.PublishedSource\n}","file":"remote/publish.go","name":"PublishItem"}}
// PublishItem is what publishing one record would write.
type PublishItem struct {
	ID    string
	Title string
	// Snaps are the versions to write, oldest first.
	Snaps []*record.Snapshot
	// Status is the local status; the last written version carries it.
	Status string
	// HeadCID and HeadVersion describe the remote head, empty and 0 when unpublished.
	HeadCID     string
	HeadVersion int
	// HeadStatus is set when only the head needs updating (no new version).
	HeadStatus string
	// Source is where the record lives in git; every write carries it when set.
	Source *record.PublishedSource
	// SetProject and SetSource are set when the record has nothing new to write
	// but its head lacks the wanted project or source: a head-only update.
	SetProject string
	SetSource  *record.PublishedSource
}

// /sedum:setType
// sedum:setType {"tier":"owned","record":"prov-2026-17d5965a","gen":"sha256:574500cb4ef448e7012a00efbbdafd37be7a068b37baf7219e427bd37401b9b7","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"// PublishPlan lists what a publish would write.\ntype PublishPlan struct {\n\tItems []PublishItem\n\t// UpToDate are records the server already has in full, with their status.\n\tUpToDate []string\n\t// Problems are records that were skipped; the others are still planned.\n\tProblems []PublishProblem\n\t// Skipped are records with nothing to publish yet, with the reason.\n\tSkipped []string\n}","file":"remote/publish.go","name":"PublishPlan"}}
// PublishPlan lists what a publish would write.
type PublishPlan struct {
	Items []PublishItem
	// UpToDate are records the server already has in full, with their status.
	UpToDate []string
	// Problems are records that were skipped; the others are still planned.
	Problems []PublishProblem
	// Skipped are records with nothing to publish yet, with the reason.
	Skipped []string
}

// /sedum:setType
// sedum:setMethod {"tier":"owned","record":"prov-2026-17d5965a","gen":"sha256:574500cb4ef448e7012a00efbbdafd37be7a068b37baf7219e427bd37401b9b7","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"// Empty reports whether the plan writes nothing.\nfunc (p *PublishPlan) Empty() bool { return len(p.Items) == 0 }","file":"remote/publish.go","name":"Empty","receiver":"PublishPlan"}}
// Empty reports whether the plan writes nothing.
func (p *PublishPlan) Empty() bool { return len(p.Items) == 0 }

// /sedum:setMethod
// sedum:setMethod {"tier":"owned","record":"prov-2026-17d5965a","gen":"sha256:574500cb4ef448e7012a00efbbdafd37be7a068b37baf7219e427bd37401b9b7","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"// Versions counts the versions the plan would write.\nfunc (p *PublishPlan) Versions() int {\n\tn := 0\n\tfor _, it := range p.Items {\n\t\tn += len(it.Snaps)\n\t}\n\treturn n\n}","file":"remote/publish.go","name":"Versions","receiver":"PublishPlan"}}
// Versions counts the versions the plan would write.
func (p *PublishPlan) Versions() int {
	n := 0
	for _, it := range p.Items {
		n += len(it.Snaps)
	}
	return n
}

// /sedum:setMethod
// sedum:setFunction {"tier":"owned","record":"prov-2026-17d5965a","gen":"sha256:574500cb4ef448e7012a00efbbdafd37be7a068b37baf7219e427bd37401b9b7","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"// PlanPublish is PlanPublishWith for callers that want no project or source on the heads.\nfunc PlanPublish(ctx context.Context, local record.Store, ids []string, c *Client) (*PublishPlan, error) {\n\treturn PlanPublishWith(ctx, local, ids, c, HeadWant{})\n}","file":"remote/publish.go","name":"PlanPublish"}}
// PlanPublish is PlanPublishWith for callers that want no project or source on the heads.
func PlanPublish(ctx context.Context, local record.Store, ids []string, c *Client) (*PublishPlan, error) {
	return PlanPublishWith(ctx, local, ids, c, HeadWant{})
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-17d5965a","gen":"sha256:574500cb4ef448e7012a00efbbdafd37be7a068b37baf7219e427bd37401b9b7","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"// ApplyPublish writes the plan: versions oldest first, each compare-and-swap\n// against the head the previous write returned, the last carrying the local\n// status; or a head-only update. commitOf maps (id, version) to the git\n// commit tagged for it, \"\" when unknown. Every write carries the item's\n// Source when it has one. progress, if not nil, hears each write.\nfunc ApplyPublish(ctx context.Context, p *PublishPlan, c *Client, project string, commitOf func(id string, version int) string, progress func(id string, version int)) error {\n\tfor _, it := range p.Items {\n\t\tcid := it.HeadCID\n\t\tfor i, snap := range it.Snaps {\n\t\t\tvar expected *string\n\t\t\tif cid != \"\" {\n\t\t\t\texpected = &cid\n\t\t\t}\n\t\t\tstatus := \"\"\n\t\t\tif i == len(it.Snaps)-1 {\n\t\t\t\tstatus = it.Status\n\t\t\t}\n\t\t\tcommit := \"\"\n\t\t\tif commitOf != nil {\n\t\t\t\tcommit = commitOf(it.ID, snap.Version)\n\t\t\t}\n\t\t\treq := VersionRequest(snap, expected, status, commit, project)\n\t\t\treq.Source = it.Source\n\t\t\thead, _, err := c.PutVersion(ctx, it.ID, req)\n\t\t\tif err != nil {\n\t\t\t\treturn fmt.Errorf(\"%s version %d: %w\", it.ID, snap.Version, err)\n\t\t\t}\n\t\t\tcid = head.CID\n\t\t\tif progress != nil {\n\t\t\t\tprogress(it.ID, snap.Version)\n\t\t\t}\n\t\t}\n\t\tif it.HeadStatus != \"\" || it.SetProject != \"\" || it.SetSource != nil {\n\t\t\tif _, err := c.PutHead(ctx, it.ID, PutHeadRequest{ExpectedHeadCid: cid, Version: it.HeadVersion, Status: it.HeadStatus, Source: it.Source, Project: project}); err != nil {\n\t\t\t\treturn fmt.Errorf(\"%s head: %w\", it.ID, err)\n\t\t\t}\n\t\t\tif progress != nil {\n\t\t\t\tprogress(it.ID, 0)\n\t\t\t}\n\t\t}\n\t}\n\treturn nil\n}","file":"remote/publish.go","name":"ApplyPublish"}}
// ApplyPublish writes the plan: versions oldest first, each compare-and-swap
// against the head the previous write returned, the last carrying the local
// status; or a head-only update. commitOf maps (id, version) to the git
// commit tagged for it, "" when unknown. Every write carries the item's
// Source when it has one. progress, if not nil, hears each write.
func ApplyPublish(ctx context.Context, p *PublishPlan, c *Client, project string, commitOf func(id string, version int) string, progress func(id string, version int)) error {
	for _, it := range p.Items {
		cid := it.HeadCID
		for i, snap := range it.Snaps {
			var expected *string
			if cid != "" {
				expected = &cid
			}
			status := ""
			if i == len(it.Snaps)-1 {
				status = it.Status
			}
			commit := ""
			if commitOf != nil {
				commit = commitOf(it.ID, snap.Version)
			}
			req := VersionRequest(snap, expected, status, commit, project)
			req.Source = it.Source
			head, _, err := c.PutVersion(ctx, it.ID, req)
			if err != nil {
				return fmt.Errorf("%s version %d: %w", it.ID, snap.Version, err)
			}
			cid = head.CID
			if progress != nil {
				progress(it.ID, snap.Version)
			}
		}
		if it.HeadStatus != "" || it.SetProject != "" || it.SetSource != nil {
			if _, err := c.PutHead(ctx, it.ID, PutHeadRequest{ExpectedHeadCid: cid, Version: it.HeadVersion, Status: it.HeadStatus, Source: it.Source, Project: project}); err != nil {
				return fmt.Errorf("%s head: %w", it.ID, err)
			}
			if progress != nil {
				progress(it.ID, 0)
			}
		}
	}
	return nil
}

// /sedum:setFunction
// sedum:setType {"tier":"owned","record":"prov-2026-17d5965a","gen":"sha256:574500cb4ef448e7012a00efbbdafd37be7a068b37baf7219e427bd37401b9b7","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"// PublishProblem is a record that cannot be published until the user acts:\n// Kind is \"server ahead\" (run import), \"diverged\" (import, or revise) or\n// \"invalid id\" (the service rejects the id; it cannot be published).\ntype PublishProblem struct {\n\tID     string\n\tKind   string\n\tReason string\n}","file":"remote/publish.go","name":"PublishProblem"}}
// PublishProblem is a record that cannot be published until the user acts:
// Kind is "server ahead" (run import), "diverged" (import, or revise) or
// "invalid id" (the service rejects the id; it cannot be published).
type PublishProblem struct {
	ID     string
	Kind   string
	Reason string
}

// /sedum:setType
// sedum:setType {"tier":"owned","record":"prov-2026-17d5965a","gen":"sha256:574500cb4ef448e7012a00efbbdafd37be7a068b37baf7219e427bd37401b9b7","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"// HeadWant is what the caller wants heads to carry besides content.\ntype HeadWant struct {\n\t// Project is the project at-uri, \"\" to leave the heads' project alone.\n\tProject string\n\t// Source says where a record lives in git; nil to leave the heads' source alone.\n\tSource func(id string) *record.PublishedSource\n}","file":"remote/publish.go","name":"HeadWant"}}
// HeadWant is what the caller wants heads to carry besides content.
type HeadWant struct {
	// Project is the project at-uri, "" to leave the heads' project alone.
	Project string
	// Source says where a record lives in git; nil to leave the heads' source alone.
	Source func(id string) *record.PublishedSource
}

// /sedum:setType
// sedum:setFunction {"tier":"owned","record":"prov-2026-17d5965a","gen":"sha256:574500cb4ef448e7012a00efbbdafd37be7a068b37baf7219e427bd37401b9b7","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"// PlanPublishWith compares each local record with its remote head. A remote head\n// ahead of the local record, or one that differs from the local version of\n// the same number, is reported in Problems and the record is skipped; only a\n// failed request is an error. An id the service would reject is reported as an\n// \"invalid id\" problem before the record is loaded or any request is sent. A record with nothing new whose head lacks the\n// wanted project or source gets a head-only update.\nfunc PlanPublishWith(ctx context.Context, local record.Store, ids []string, c *Client, want HeadWant) (*PublishPlan, error) {\n\tplan := &PublishPlan{}\n\tfor _, id := range ids {\n\t\tif !ValidServerID(id) {\n\t\t\tplan.Problems = append(plan.Problems, PublishProblem{ID: id, Kind: \"invalid id\", Reason: \"id must look like prov-2026-1a2b3c4d (the record service rejects any other shape)\"})\n\t\t\tcontinue\n\t\t}\n\t\trec, _, err := local.Load(id)\n\t\tif err != nil {\n\t\t\treturn nil, err\n\t\t}\n\t\tversions, err := local.Versions(id)\n\t\tif err != nil {\n\t\t\treturn nil, err\n\t\t}\n\t\tif len(versions) == 0 {\n\t\t\tplan.Skipped = append(plan.Skipped, id+\": a draft with no version yet (open it first)\")\n\t\t\tcontinue\n\t\t}\n\t\titem := PublishItem{ID: id, Title: rec.Title, Status: string(rec.Status)}\n\t\tif want.Source != nil {\n\t\t\titem.Source = want.Source(id)\n\t\t}\n\t\thead, err := c.GetRecord(ctx, id)\n\t\tif err != nil && !errors.Is(err, ErrNotPublished) {\n\t\t\treturn nil, err\n\t\t}\n\t\tif err == nil {\n\t\t\titem.HeadCID, item.HeadVersion = head.CID, head.Version\n\t\t}\n\t\tlast := versions[len(versions)-1]\n\t\tif item.HeadVersion > last {\n\t\t\tplan.Problems = append(plan.Problems, PublishProblem{ID: id, Kind: \"server ahead\", Reason: fmt.Sprintf(\"the server is at version %d, ahead of the local %d; run provenance import\", head.Version, last)})\n\t\t\tcontinue\n\t\t}\n\t\tif item.HeadVersion > 0 {\n\t\t\tsnap, err := local.Snapshot(id, item.HeadVersion)\n\t\t\tif err != nil {\n\t\t\t\treturn nil, err\n\t\t\t}\n\t\t\tif snap.ContentHash != head.ContentHash {\n\t\t\t\tplan.Problems = append(plan.Problems, PublishProblem{ID: id, Kind: \"diverged\", Reason: fmt.Sprintf(\"version %d differs between the server and this repository; run provenance import, or revise\", item.HeadVersion)})\n\t\t\t\tcontinue\n\t\t\t}\n\t\t}\n\t\tfor _, v := range versions {\n\t\t\tif v > item.HeadVersion {\n\t\t\t\tsnap, err := local.Snapshot(id, v)\n\t\t\t\tif err != nil {\n\t\t\t\t\treturn nil, err\n\t\t\t\t}\n\t\t\t\titem.Snaps = append(item.Snaps, snap)\n\t\t\t}\n\t\t}\n\t\tif len(item.Snaps) == 0 {\n\t\t\tif want.Project != \"\" && head.Project != want.Project {\n\t\t\t\titem.SetProject = want.Project\n\t\t\t}\n\t\t\tif item.Source != nil && (head.Source == nil || head.Source.Repo != item.Source.Repo || head.Source.Path != item.Source.Path) {\n\t\t\t\titem.SetSource = item.Source\n\t\t\t}\n\t\t\tif statusRank(item.Status) > statusRank(head.Status) {\n\t\t\t\titem.HeadStatus = item.Status\n\t\t\t} else if item.SetProject == \"\" && item.SetSource == nil {\n\t\t\t\tplan.UpToDate = append(plan.UpToDate, id)\n\t\t\t\tcontinue\n\t\t\t}\n\t\t}\n\t\tplan.Items = append(plan.Items, item)\n\t}\n\treturn plan, nil\n}","file":"remote/publish.go","name":"PlanPublishWith"}}
// PlanPublishWith compares each local record with its remote head. A remote head
// ahead of the local record, or one that differs from the local version of
// the same number, is reported in Problems and the record is skipped; only a
// failed request is an error. An id the service would reject is reported as an
// "invalid id" problem before the record is loaded or any request is sent. A record with nothing new whose head lacks the
// wanted project or source gets a head-only update.
func PlanPublishWith(ctx context.Context, local record.Store, ids []string, c *Client, want HeadWant) (*PublishPlan, error) {
	plan := &PublishPlan{}
	for _, id := range ids {
		if !ValidServerID(id) {
			plan.Problems = append(plan.Problems, PublishProblem{ID: id, Kind: "invalid id", Reason: "id must look like prov-2026-1a2b3c4d (the record service rejects any other shape)"})
			continue
		}
		rec, _, err := local.Load(id)
		if err != nil {
			return nil, err
		}
		versions, err := local.Versions(id)
		if err != nil {
			return nil, err
		}
		if len(versions) == 0 {
			plan.Skipped = append(plan.Skipped, id+": a draft with no version yet (open it first)")
			continue
		}
		item := PublishItem{ID: id, Title: rec.Title, Status: string(rec.Status)}
		if want.Source != nil {
			item.Source = want.Source(id)
		}
		head, err := c.GetRecord(ctx, id)
		if err != nil && !errors.Is(err, ErrNotPublished) {
			return nil, err
		}
		if err == nil {
			item.HeadCID, item.HeadVersion = head.CID, head.Version
		}
		last := versions[len(versions)-1]
		if item.HeadVersion > last {
			plan.Problems = append(plan.Problems, PublishProblem{ID: id, Kind: "server ahead", Reason: fmt.Sprintf("the server is at version %d, ahead of the local %d; run provenance import", head.Version, last)})
			continue
		}
		if item.HeadVersion > 0 {
			snap, err := local.Snapshot(id, item.HeadVersion)
			if err != nil {
				return nil, err
			}
			if snap.ContentHash != head.ContentHash {
				plan.Problems = append(plan.Problems, PublishProblem{ID: id, Kind: "diverged", Reason: fmt.Sprintf("version %d differs between the server and this repository; run provenance import, or revise", item.HeadVersion)})
				continue
			}
		}
		for _, v := range versions {
			if v > item.HeadVersion {
				snap, err := local.Snapshot(id, v)
				if err != nil {
					return nil, err
				}
				item.Snaps = append(item.Snaps, snap)
			}
		}
		if len(item.Snaps) == 0 {
			if want.Project != "" && head.Project != want.Project {
				item.SetProject = want.Project
			}
			if item.Source != nil && (head.Source == nil || head.Source.Repo != item.Source.Repo || head.Source.Path != item.Source.Path) {
				item.SetSource = item.Source
			}
			if statusRank(item.Status) > statusRank(head.Status) {
				item.HeadStatus = item.Status
			} else if item.SetProject == "" && item.SetSource == nil {
				plan.UpToDate = append(plan.UpToDate, id)
				continue
			}
		}
		plan.Items = append(plan.Items, item)
	}
	return plan, nil
}

// /sedum:setFunction
// sedum:setConst {"tier":"owned","record":"prov-2026-17d5965a","gen":"sha256:574500cb4ef448e7012a00efbbdafd37be7a068b37baf7219e427bd37401b9b7","pkg":"sha256:51be1c18998d63cc558c253e211a6bd188d96cb251dad68c6ea88e03c2d5cc4c","kwargs":{"code":"// ServerIDPattern is the record id pattern the record service accepts. If the\n// service widens it, change it here and nowhere else.\nconst ServerIDPattern = `^prov-\\d{4}-([0-9a-f]{8}|\\d{3})$`","file":"remote/publish.go","name":"ServerIDPattern"}}
// ServerIDPattern is the record id pattern the record service accepts. If the
// service widens it, change it here and nowhere else.
const ServerIDPattern = `^prov-\d{4}-([0-9a-f]{8}|\d{3})$`

// /sedum:setConst
// sedum:setVar {"tier":"owned","record":"prov-2026-17d5965a","gen":"sha256:574500cb4ef448e7012a00efbbdafd37be7a068b37baf7219e427bd37401b9b7","pkg":"sha256:c7b9d9ac1655e8211b8b2ca8c49cf34b8d510d8d823d4f5eb60ec7e1018aa2eb","kwargs":{"code":"var serverID = regexp.MustCompile(ServerIDPattern)","file":"remote/publish.go","name":"serverID"}}
var serverID = regexp.MustCompile(ServerIDPattern)

// /sedum:setVar
// sedum:setFunction {"tier":"owned","record":"prov-2026-17d5965a","gen":"sha256:574500cb4ef448e7012a00efbbdafd37be7a068b37baf7219e427bd37401b9b7","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"// ValidServerID reports whether the record service accepts id.\nfunc ValidServerID(id string) bool { return serverID.MatchString(id) }","file":"remote/publish.go","name":"ValidServerID"}}
// ValidServerID reports whether the record service accepts id.
func ValidServerID(id string) bool { return serverID.MatchString(id) }

// /sedum:setFunction
