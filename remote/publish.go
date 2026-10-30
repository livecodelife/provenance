package remote

// sedum:anchor:imports
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-17d5965a","gen":"sha256:682c2334ef2511bc732b96205d5b54d93adb5bd52e9b862bebef0452bb5cfac6","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"remote/publish.go","path":"github.com/livecodelife/provenance/record"},"kwargs":"sha256:12d3e1c9c9e6a222eee2cd3062515099a7fa9e566dc2069aabd20352f0f7c983","body":"sha256:acbf8628061291d11aa39e9ea64ecc91095a117057cfe6f1958ef2c437fac003"}
import "github.com/livecodelife/provenance/record"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-7131dd15","gen":"sha256:494525f907d90c95339c41d2b6241d07357be090e07fe3bbe6ab860c4f9f1f2d","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"remote/publish.go","path":"regexp"},"kwargs":"sha256:8fc7f7e3cf5c384e3931878a5454eb83892159d490938eccda960184533cc16d","body":"sha256:b437bd748876cfcdfa00102314b52d7ce19d023c66492336bdfd27288b1329d2"}
import "regexp"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-8c70a58d","gen":"sha256:301cde4b5379aa7a99ce978fe74af64f95c4889b926af2318d4bd248385c189a","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"remote/publish.go","path":"errors"},"kwargs":"sha256:61a560e1efbfa59c76b06baaa92efed6dbe189c0e50fed7f0cdaa7547a067eb6","body":"sha256:2b718237ba0a56467c8fd494cc131e942ef410e6228c03b644d70476c71862dc"}
import "errors"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-8c70a58d","gen":"sha256:301cde4b5379aa7a99ce978fe74af64f95c4889b926af2318d4bd248385c189a","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"remote/publish.go","path":"context"},"kwargs":"sha256:d43aa832fdf5e9af3386e2cabeaf6d3428c7398f79257938f523564ac40c2104","body":"sha256:0f7443fb25d7ad6a970617d4d0df6123283e24801b94c1d84fbe431230f31e1e"}
import "context"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-8c70a58d","gen":"sha256:301cde4b5379aa7a99ce978fe74af64f95c4889b926af2318d4bd248385c189a","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"remote/publish.go","path":"fmt"},"kwargs":"sha256:2e981e87b762a9b6b2dcfbb8bd2466e85efae1e02396af0d3220a338d747703b","body":"sha256:bcbac51883cbf8064e9196b371cb880cee429866652447ad557533d817857609"}
import "fmt"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setType {"v":2,"tier":"owned","record":"prov-2026-17d5965a","gen":"sha256:682c2334ef2511bc732b96205d5b54d93adb5bd52e9b862bebef0452bb5cfac6","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","id":{"file":"remote/publish.go","name":"PublishItem"},"kwargs":"sha256:45e3b6558b1b24ce3411be64c732788efe203c5a6b8f0ea26f2a1bf93b0840d5","body":"sha256:d5268de3c17537d5e7e4e7b8bfda767d7096291fa62563bfc43ab1cfcbac1a75"}
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
// sedum:setType {"v":2,"tier":"owned","record":"prov-2026-6d72ebe1","gen":"sha256:46eb456fd5781df6c0f5b86f7a57bfa7406c1942c830b80b3166423980a0950d","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","id":{"file":"remote/publish.go","name":"PublishProblem"},"kwargs":"sha256:f13984c529444ef26c0bdfb7db72bd4c5147e50d96178bb6c02201e0981df614","body":"sha256:a98524504dc96fe81d713c941bedcf1c2db48d206160a57515274202bb7c7d24"}
// PublishProblem describes a record that cannot be published until the user acts.
type PublishProblem struct {
	ID     string
	Kind   string
	Reason string
}

// /sedum:setType
// sedum:setConst {"v":2,"tier":"owned","record":"prov-2026-709cff9f","gen":"sha256:6aca47d51a7fb7936c439e3e343d5ae33bc5242d4e736e85d2b0f0a34ad6815d","pkg":"sha256:51be1c18998d63cc558c253e211a6bd188d96cb251dad68c6ea88e03c2d5cc4c","id":{"file":"remote/publish.go","name":"ServerIDPattern"},"kwargs":"sha256:f4f4690a39511790ee9e6971412b68c6a23db49255bc7013d2ddc9e379313865","body":"sha256:9ecaea85f4d22e1021771ca78f91a1e056113d1c1a9b38ed5d56a45d979c2198"}
const ServerIDPattern = `^prov-[0-9]{4}-([0-9a-f]{8}|[0-9]{3})$`

// /sedum:setConst
// sedum:setVar {"v":2,"tier":"owned","record":"prov-2026-7131dd15","gen":"sha256:494525f907d90c95339c41d2b6241d07357be090e07fe3bbe6ab860c4f9f1f2d","pkg":"sha256:c7b9d9ac1655e8211b8b2ca8c49cf34b8d510d8d823d4f5eb60ec7e1018aa2eb","id":{"file":"remote/publish.go","name":"serverID"},"kwargs":"sha256:66a82a43477e229ca90eb8b53aa1cdb62a7fcb8124b346fcea8bd377ecfee8c6","body":"sha256:05b17e850e9a3b6c0edf59fc8f3c963aa0241fe33916a6a220f8ffa976b7d61c"}
var serverID = regexp.MustCompile(ServerIDPattern)

// /sedum:setVar
// sedum:setMethod {"v":2,"tier":"owned","record":"prov-2026-8276db5d","gen":"sha256:d3edbf6fc0d68e38355464e700099379a426b4a3c8831953c33c32d3bd074530","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","id":{"file":"remote/publish.go","name":"Empty","receiver":"PublishPlan"},"kwargs":"sha256:c4f2cf0dcbd42761d72ec40cdaf150748d9c0ace70a0791ef4d95cda72b645a2","body":"sha256:0fba64ffce5446b9a81d26404e14bf12405054b398badf231705f061b9ec1855"}
func (p *PublishPlan) Empty() bool {
	return len(p.Items) == 0
}

// /sedum:setMethod
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-8c70a58d","gen":"sha256:301cde4b5379aa7a99ce978fe74af64f95c4889b926af2318d4bd248385c189a","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"remote/publish.go","name":"PlanPublishWith"},"kwargs":"sha256:1cfa484e62acc997965d7b740e57b69fbf00cbbf347dbc5e3639664f0556486e","body":"sha256:e4a853c494b9a0892351b6def270b350ed053b84ae5d7577c0e8c23ce6b06b24"}
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
// sedum:setType {"v":2,"tier":"owned","record":"prov-2026-a40f827c","gen":"sha256:506bc97774e9926b23886bc1f63c7d10d17e501daf5bf117bd098477c37e15b2","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","id":{"file":"remote/publish.go","name":"PublishPlan"},"kwargs":"sha256:b04a8fa93721bbaf0bf1784c3e3c610d13b75f53d38bef36bd061b2808f4d352","body":"sha256:7a3224e8bd5feae7389154257b1b23d8639af24fd94b2d7e228ace80e490940e"}
type PublishPlan struct {
	Items    []PublishItem
	UpToDate []string
	Problems []PublishProblem
	Skipped  []string
}

// /sedum:setType
// sedum:setMethod {"v":2,"tier":"owned","record":"prov-2026-a48197c4","gen":"sha256:41ca96ff20a35145a4d3d94539bb5554717fc20d32b33499d09838756dd0eb23","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","id":{"file":"remote/publish.go","name":"Versions","receiver":"PublishPlan"},"kwargs":"sha256:9ce5fcb872f89a841761d6ce3351017a0cd4e8b873af8ba7c10630ca4d1b6c14","body":"sha256:3c8401b3099ab1ca10b1d205de2866079ee0f94dc14f5e066352448fd7fdda49"}
func (p *PublishPlan) Versions() int {
	n := 0
	for _, item := range p.Items {
		n += len(item.Snaps)
	}
	return n
}

// /sedum:setMethod
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-ddf5f652","gen":"sha256:0ee12bbbc821a9f851af6f66cc0ce25a4ca1faa7b4bd2651c1f99ee92090f7a7","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"remote/publish.go","name":"PlanPublish"},"kwargs":"sha256:684a0abcda997dce71bd4aadfb283a4b048696077f2e0cc8f95efc23745f14d1","body":"sha256:f180f6affe12a88d8db8214d4921fd5ebf4c48bb50b22cc9f8ad6684a64b0bbb"}
func PlanPublish(ctx context.Context, local record.Store, ids []string, c *Client) (*PublishPlan, error) {
	return PlanPublishWith(ctx, local, ids, c, HeadWant{})
}

// /sedum:setFunction
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-e1669492","gen":"sha256:eb9813daf72f443ea1d4849fa64bf0f2453bb67e1f3180232ff7a37ea7cf0b1d","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"remote/publish.go","name":"statusRank"},"kwargs":"sha256:3a588109476a6b86b822c0816d9036e9ba86d9c80feaf6c59a09817ff75cc2a6","body":"sha256:6a70d60fb6f7cee96e9258ce9384fa634574a2e6f6efd9613afd9ffb72c1d112"}
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
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-e3891a7a","gen":"sha256:43abb8fdca9036673209190cb8c66c55a5c90709e8e8405537312aeb5cfc1b81","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"remote/publish.go","name":"ApplyPublish"},"kwargs":"sha256:fe8057369a5390d98370487615774264457c76d4a826b740bfa6d6f7fff4b27d","body":"sha256:47312b53d7c5ca7c3395cd3c82539ccdf694a8323d90f28e614f99d3058bae2e"}
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
// sedum:setType {"v":2,"tier":"owned","record":"prov-2026-e95b8b0f","gen":"sha256:194a384bbd39bacdd5795cb36d4b073bda0560698bc54d6f4c10a9d0b713dd31","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","id":{"file":"remote/publish.go","name":"HeadWant"},"kwargs":"sha256:6edf86ee8c3e7a4732ba6e8afaaa45cf3e102a0414c097ea58ced4cb1bdc822c","body":"sha256:3c900b87e421277c4a54d30927b5380cfa53c5fb15d0054546b345d66a9e6be1"}
type HeadWant struct {
	Project string
	Source  func(id string) *record.PublishedSource
}

// /sedum:setType
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-f97fd932","gen":"sha256:ca69039feef1337c0599f45fb1b3de9e1571700a2a836b5b5a0fd07be286ac3a","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"remote/publish.go","name":"ValidServerID"},"kwargs":"sha256:2c9a4377b5cc625ef35b9d5451b16ef6908d9339f70eea741eba382d96d92d0f","body":"sha256:07198e61be63c299c36af813192a09fc856193d218f41db6d04ae150c44ab97c"}
func ValidServerID(id string) bool {
	if id == "" {
		return false
	}
	return serverID.MatchString(id)
}

// /sedum:setFunction
