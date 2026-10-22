package remote

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-17d5965a","gen":"sha256:682c2334ef2511bc732b96205d5b54d93adb5bd52e9b862bebef0452bb5cfac6","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"remote/publish.go","path":"github.com/livecodelife/provenance/record"}}
import "github.com/livecodelife/provenance/record"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-7131dd15","gen":"sha256:494525f907d90c95339c41d2b6241d07357be090e07fe3bbe6ab860c4f9f1f2d","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"remote/publish.go","path":"regexp"}}
import "regexp"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-8c70a58d","gen":"sha256:301cde4b5379aa7a99ce978fe74af64f95c4889b926af2318d4bd248385c189a","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"remote/publish.go","path":"errors"}}
import "errors"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-8c70a58d","gen":"sha256:301cde4b5379aa7a99ce978fe74af64f95c4889b926af2318d4bd248385c189a","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"remote/publish.go","path":"context"}}
import "context"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-8c70a58d","gen":"sha256:301cde4b5379aa7a99ce978fe74af64f95c4889b926af2318d4bd248385c189a","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"remote/publish.go","path":"fmt"}}
import "fmt"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setType {"tier":"owned","record":"prov-2026-17d5965a","gen":"sha256:682c2334ef2511bc732b96205d5b54d93adb5bd52e9b862bebef0452bb5cfac6","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type PublishItem struct {\n\tID          string\n\tTitle       string\n\tSnaps       []*record.Snapshot\n\tStatus      string\n\tHeadCID     string\n\tHeadVersion int\n\tHeadStatus  string\n\tSource      *record.PublishedSource\n\tSetProject  string\n\tSetSource   *record.PublishedSource\n}","file":"remote/publish.go","name":"PublishItem"}}
type PublishItem struct {
	ID          string
	Title       string
	Snaps       []*record.Snapshot
	Status      string
	HeadCID     string
	HeadVersion int
	HeadStatus  string
	Source      *record.PublishedSource
	SetProject  string
	SetSource   *record.PublishedSource
}

// /sedum:setType
// sedum:setType {"tier":"owned","record":"prov-2026-6d72ebe1","gen":"sha256:46eb456fd5781df6c0f5b86f7a57bfa7406c1942c830b80b3166423980a0950d","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"// PublishProblem describes a record that cannot be published until the user acts.\ntype PublishProblem struct {\n\tID     string\n\tKind   string\n\tReason string\n}","file":"remote/publish.go","name":"PublishProblem"}}
// PublishProblem describes a record that cannot be published until the user acts.
type PublishProblem struct {
	ID     string
	Kind   string
	Reason string
}

// /sedum:setType
// sedum:setConst {"tier":"owned","record":"prov-2026-709cff9f","gen":"sha256:6aca47d51a7fb7936c439e3e343d5ae33bc5242d4e736e85d2b0f0a34ad6815d","pkg":"sha256:51be1c18998d63cc558c253e211a6bd188d96cb251dad68c6ea88e03c2d5cc4c","kwargs":{"code":"const ServerIDPattern = `^prov-[0-9]{4}-([0-9a-f]{8}|[0-9]{3})$`","file":"remote/publish.go","name":"ServerIDPattern"}}
const ServerIDPattern = `^prov-[0-9]{4}-([0-9a-f]{8}|[0-9]{3})$`

// /sedum:setConst
// sedum:setVar {"tier":"owned","record":"prov-2026-7131dd15","gen":"sha256:494525f907d90c95339c41d2b6241d07357be090e07fe3bbe6ab860c4f9f1f2d","pkg":"sha256:c7b9d9ac1655e8211b8b2ca8c49cf34b8d510d8d823d4f5eb60ec7e1018aa2eb","kwargs":{"code":"var serverID = regexp.MustCompile(ServerIDPattern)","file":"remote/publish.go","name":"serverID"}}
var serverID = regexp.MustCompile(ServerIDPattern)

// /sedum:setVar
// sedum:setMethod {"tier":"owned","record":"prov-2026-8276db5d","gen":"sha256:d3edbf6fc0d68e38355464e700099379a426b4a3c8831953c33c32d3bd074530","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (p *PublishPlan) Empty() bool {\n\treturn len(p.Items) == 0\n}","file":"remote/publish.go","name":"Empty","receiver":"PublishPlan"}}
func (p *PublishPlan) Empty() bool {
	return len(p.Items) == 0
}

// /sedum:setMethod
// sedum:setFunction {"tier":"owned","record":"prov-2026-8c70a58d","gen":"sha256:301cde4b5379aa7a99ce978fe74af64f95c4889b926af2318d4bd248385c189a","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func PlanPublishWith(ctx context.Context, local record.Store, ids []string, c *Client, want HeadWant) (*PublishPlan, error) {\n\tplan := &PublishPlan{}\n\tfor _, id := range ids {\n\t\tif !ValidServerID(id) {\n\t\t\tplan.Problems = append(plan.Problems, PublishProblem{\n\t\t\t\tID:     id,\n\t\t\t\tKind:   \"invalid id\",\n\t\t\t\tReason: \"id must look like prov-2026-1a2b3c4d (the record service rejects any other shape)\",\n\t\t\t})\n\t\t\tcontinue\n\t\t}\n\t\trec, _, err := local.Load(id)\n\t\tif err != nil {\n\t\t\treturn nil, err\n\t\t}\n\t\tversions, err := local.Versions(id)\n\t\tif err != nil {\n\t\t\treturn nil, err\n\t\t}\n\t\tif len(versions) == 0 {\n\t\t\tplan.Skipped = append(plan.Skipped, id+\": a draft with no version yet (open it first)\")\n\t\t\tcontinue\n\t\t}\n\t\titem := PublishItem{\n\t\t\tID:     id,\n\t\t\tTitle:  rec.Title,\n\t\t\tStatus: string(rec.Status),\n\t\t}\n\t\tif want.Source != nil {\n\t\t\titem.Source = want.Source(id)\n\t\t}\n\t\thead, err := c.GetRecord(ctx, id)\n\t\tif err != nil {\n\t\t\tif !errors.Is(err, ErrNotPublished) {\n\t\t\t\treturn nil, err\n\t\t\t}\n\t\t} else {\n\t\t\titem.HeadCID = head.CID\n\t\t\titem.HeadVersion = head.Version\n\t\t}\n\t\tlast := versions[len(versions)-1]\n\t\tif item.HeadVersion > last {\n\t\t\tplan.Problems = append(plan.Problems, PublishProblem{\n\t\t\t\tID:     id,\n\t\t\t\tKind:   \"server ahead\",\n\t\t\t\tReason: fmt.Sprintf(\"the server is at version %d, ahead of the local %d; run provenance import\", head.Version, last),\n\t\t\t})\n\t\t\tcontinue\n\t\t}\n\t\tif item.HeadVersion > 0 {\n\t\t\tsnap, err := local.Snapshot(id, item.HeadVersion)\n\t\t\tif err != nil {\n\t\t\t\treturn nil, err\n\t\t\t}\n\t\t\tif snap.ContentHash != head.ContentHash {\n\t\t\t\tplan.Problems = append(plan.Problems, PublishProblem{\n\t\t\t\t\tID:     id,\n\t\t\t\t\tKind:   \"diverged\",\n\t\t\t\t\tReason: fmt.Sprintf(\"version %d differs between the server and this repository; run provenance import, or revise\", item.HeadVersion),\n\t\t\t\t})\n\t\t\t\tcontinue\n\t\t\t}\n\t\t}\n\t\tfor _, v := range versions {\n\t\t\tif v > item.HeadVersion {\n\t\t\t\ts, err := local.Snapshot(id, v)\n\t\t\t\tif err != nil {\n\t\t\t\t\treturn nil, err\n\t\t\t\t}\n\t\t\t\titem.Snaps = append(item.Snaps, s)\n\t\t\t}\n\t\t}\n\t\tif len(item.Snaps) == 0 {\n\t\t\tif want.Project != \"\" && want.Project != head.Project {\n\t\t\t\titem.SetProject = want.Project\n\t\t\t}\n\t\t\tif item.Source != nil && (head.Source == nil || head.Source.Repo != item.Source.Repo || head.Source.Path != item.Source.Path) {\n\t\t\t\titem.SetSource = item.Source\n\t\t\t}\n\t\t\tif statusRank(item.Status) > statusRank(head.Status) {\n\t\t\t\titem.HeadStatus = item.Status\n\t\t\t} else if item.SetProject == \"\" && item.SetSource == nil {\n\t\t\t\tplan.UpToDate = append(plan.UpToDate, id)\n\t\t\t\tcontinue\n\t\t\t}\n\t\t}\n\t\tplan.Items = append(plan.Items, item)\n\t}\n\treturn plan, nil\n}\n","file":"remote/publish.go","name":"PlanPublishWith"}}
func PlanPublishWith(ctx context.Context, local record.Store, ids []string, c *Client, want HeadWant) (*PublishPlan, error) {
	plan := &PublishPlan{}
	for _, id := range ids {
		if !ValidServerID(id) {
			plan.Problems = append(plan.Problems, PublishProblem{
				ID:     id,
				Kind:   "invalid id",
				Reason: "id must look like prov-2026-1a2b3c4d (the record service rejects any other shape)",
			})
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
		item := PublishItem{
			ID:     id,
			Title:  rec.Title,
			Status: string(rec.Status),
		}
		if want.Source != nil {
			item.Source = want.Source(id)
		}
		head, err := c.GetRecord(ctx, id)
		if err != nil {
			if !errors.Is(err, ErrNotPublished) {
				return nil, err
			}
		} else {
			item.HeadCID = head.CID
			item.HeadVersion = head.Version
		}
		last := versions[len(versions)-1]
		if item.HeadVersion > last {
			plan.Problems = append(plan.Problems, PublishProblem{
				ID:     id,
				Kind:   "server ahead",
				Reason: fmt.Sprintf("the server is at version %d, ahead of the local %d; run provenance import", head.Version, last),
			})
			continue
		}
		if item.HeadVersion > 0 {
			snap, err := local.Snapshot(id, item.HeadVersion)
			if err != nil {
				return nil, err
			}
			if snap.ContentHash != head.ContentHash {
				plan.Problems = append(plan.Problems, PublishProblem{
					ID:     id,
					Kind:   "diverged",
					Reason: fmt.Sprintf("version %d differs between the server and this repository; run provenance import, or revise", item.HeadVersion),
				})
				continue
			}
		}
		for _, v := range versions {
			if v > item.HeadVersion {
				s, err := local.Snapshot(id, v)
				if err != nil {
					return nil, err
				}
				item.Snaps = append(item.Snaps, s)
			}
		}
		if len(item.Snaps) == 0 {
			if want.Project != "" && want.Project != head.Project {
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
// sedum:setType {"tier":"owned","record":"prov-2026-a40f827c","gen":"sha256:506bc97774e9926b23886bc1f63c7d10d17e501daf5bf117bd098477c37e15b2","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type PublishPlan struct {\n\tItems    []PublishItem\n\tUpToDate []string\n\tProblems []PublishProblem\n\tSkipped  []string\n}","file":"remote/publish.go","name":"PublishPlan"}}
type PublishPlan struct {
	Items    []PublishItem
	UpToDate []string
	Problems []PublishProblem
	Skipped  []string
}

// /sedum:setType
// sedum:setMethod {"tier":"owned","record":"prov-2026-a48197c4","gen":"sha256:41ca96ff20a35145a4d3d94539bb5554717fc20d32b33499d09838756dd0eb23","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (p *PublishPlan) Versions() int {\n\tn := 0\n\tfor _, item := range p.Items {\n\t\tn += len(item.Snaps)\n\t}\n\treturn n\n}","file":"remote/publish.go","name":"Versions","receiver":"PublishPlan"}}
func (p *PublishPlan) Versions() int {
	n := 0
	for _, item := range p.Items {
		n += len(item.Snaps)
	}
	return n
}

// /sedum:setMethod
// sedum:setFunction {"tier":"owned","record":"prov-2026-ddf5f652","gen":"sha256:0ee12bbbc821a9f851af6f66cc0ce25a4ca1faa7b4bd2651c1f99ee92090f7a7","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func PlanPublish(ctx context.Context, local record.Store, ids []string, c *Client) (*PublishPlan, error) {\n\treturn PlanPublishWith(ctx, local, ids, c, HeadWant{})\n}","file":"remote/publish.go","name":"PlanPublish"}}
func PlanPublish(ctx context.Context, local record.Store, ids []string, c *Client) (*PublishPlan, error) {
	return PlanPublishWith(ctx, local, ids, c, HeadWant{})
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-e1669492","gen":"sha256:eb9813daf72f443ea1d4849fa64bf0f2453bb67e1f3180232ff7a37ea7cf0b1d","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func statusRank(s string) int {\n\tswitch s {\n\tcase \"draft\":\n\t\treturn 0\n\tcase \"open\":\n\t\treturn 1\n\tcase \"implemented\":\n\t\treturn 2\n\tcase \"deprecated\":\n\t\treturn 3\n\t}\n\treturn -1\n}","file":"remote/publish.go","name":"statusRank"}}
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
// sedum:setFunction {"tier":"owned","record":"prov-2026-e3891a7a","gen":"sha256:43abb8fdca9036673209190cb8c66c55a5c90709e8e8405537312aeb5cfc1b81","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func ApplyPublish(ctx context.Context, p *PublishPlan, c *Client, project string, commitOf func(id string, version int) string, progress func(id string, version int)) error {\n\tfor _, item := range p.Items {\n\t\thead := item.HeadCID\n\t\tfor i, snap := range item.Snaps {\n\t\t\tvar expected *string\n\t\t\tif head != \"\" {\n\t\t\t\th := head\n\t\t\t\texpected = &h\n\t\t\t}\n\t\t\tstatus := \"\"\n\t\t\tif i == len(item.Snaps)-1 {\n\t\t\t\tstatus = item.Status\n\t\t\t}\n\t\t\tcommit := \"\"\n\t\t\tif commitOf != nil {\n\t\t\t\tcommit = commitOf(item.ID, snap.Version)\n\t\t\t}\n\t\t\treq := VersionRequest(snap, expected, status, commit, project)\n\t\t\treq.Source = item.Source\n\t\t\th, _, err := c.PutVersion(ctx, item.ID, req)\n\t\t\tif err != nil {\n\t\t\t\treturn fmt.Errorf(\"%s version %d: %w\", item.ID, snap.Version, err)\n\t\t\t}\n\t\t\thead = h.CID\n\t\t\tif progress != nil {\n\t\t\t\tprogress(item.ID, snap.Version)\n\t\t\t}\n\t\t}\n\t\tif item.HeadStatus != \"\" || item.SetProject != \"\" || item.SetSource != nil {\n\t\t\tif _, err := c.PutHead(ctx, item.ID, PutHeadRequest{\n\t\t\t\tExpectedHeadCid: head,\n\t\t\t\tVersion:         item.HeadVersion,\n\t\t\t\tStatus:          item.HeadStatus,\n\t\t\t\tSource:          item.Source,\n\t\t\t\tProject:         project,\n\t\t\t}); err != nil {\n\t\t\t\treturn fmt.Errorf(\"%s head: %w\", item.ID, err)\n\t\t\t}\n\t\t\tif progress != nil {\n\t\t\t\tprogress(item.ID, 0)\n\t\t\t}\n\t\t}\n\t}\n\treturn nil\n}","file":"remote/publish.go","name":"ApplyPublish"}}
func ApplyPublish(ctx context.Context, p *PublishPlan, c *Client, project string, commitOf func(id string, version int) string, progress func(id string, version int)) error {
	for _, item := range p.Items {
		head := item.HeadCID
		for i, snap := range item.Snaps {
			var expected *string
			if head != "" {
				h := head
				expected = &h
			}
			status := ""
			if i == len(item.Snaps)-1 {
				status = item.Status
			}
			commit := ""
			if commitOf != nil {
				commit = commitOf(item.ID, snap.Version)
			}
			req := VersionRequest(snap, expected, status, commit, project)
			req.Source = item.Source
			h, _, err := c.PutVersion(ctx, item.ID, req)
			if err != nil {
				return fmt.Errorf("%s version %d: %w", item.ID, snap.Version, err)
			}
			head = h.CID
			if progress != nil {
				progress(item.ID, snap.Version)
			}
		}
		if item.HeadStatus != "" || item.SetProject != "" || item.SetSource != nil {
			if _, err := c.PutHead(ctx, item.ID, PutHeadRequest{
				ExpectedHeadCid: head,
				Version:         item.HeadVersion,
				Status:          item.HeadStatus,
				Source:          item.Source,
				Project:         project,
			}); err != nil {
				return fmt.Errorf("%s head: %w", item.ID, err)
			}
			if progress != nil {
				progress(item.ID, 0)
			}
		}
	}
	return nil
}

// /sedum:setFunction
// sedum:setType {"tier":"owned","record":"prov-2026-e95b8b0f","gen":"sha256:194a384bbd39bacdd5795cb36d4b073bda0560698bc54d6f4c10a9d0b713dd31","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type HeadWant struct {\n\tProject string\n\tSource  func(id string) *record.PublishedSource\n}","file":"remote/publish.go","name":"HeadWant"}}
type HeadWant struct {
	Project string
	Source  func(id string) *record.PublishedSource
}

// /sedum:setType
// sedum:setFunction {"tier":"owned","record":"prov-2026-f97fd932","gen":"sha256:ca69039feef1337c0599f45fb1b3de9e1571700a2a836b5b5a0fd07be286ac3a","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func ValidServerID(id string) bool {\n\tif id == \"\" {\n\t\treturn false\n\t}\n\treturn serverID.MatchString(id)\n}","file":"remote/publish.go","name":"ValidServerID"}}
func ValidServerID(id string) bool {
	if id == "" {
		return false
	}
	return serverID.MatchString(id)
}

// /sedum:setFunction
