package document

// sedum:anchor:imports
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-2a640ada","gen":"sha256:52f742d9bf131f6a93dc56d698a431e545e822fbbf32baaaf87b6111f45edaa6","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"document/apply.go","path":"errors"},"kwargs":"sha256:249a95f2a795621d72358e02c445c0d3d9c8a1e280891eba22540750033f8b34","body":"sha256:2b718237ba0a56467c8fd494cc131e942ef410e6228c03b644d70476c71862dc"}
import "errors"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-6b2ceecb","gen":"sha256:a1af9e73e0651d18088c29216aa26fbfe9fa6d1c43db78b15a6c2b589b11ec10","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"document/apply.go","path":"fmt"},"kwargs":"sha256:b960631aa246fad89bda36db1f16179e6f869b5056045681d9a717e4d667d129","body":"sha256:bcbac51883cbf8064e9196b371cb880cee429866652447ad557533d817857609"}
import "fmt"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-6b2ceecb","gen":"sha256:a1af9e73e0651d18088c29216aa26fbfe9fa6d1c43db78b15a6c2b589b11ec10","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"document/apply.go","path":"strings"},"kwargs":"sha256:d636acd8090daabdc6b060503f84eb6bfdee920b851dbc99c47d4322820bf5ef","body":"sha256:a9af3c0bc4bf9342df9d7e89a324523fdba7941a9ca4190270e68ff3c4a25ea6"}
import "strings"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-6b2ceecb","gen":"sha256:a1af9e73e0651d18088c29216aa26fbfe9fa6d1c43db78b15a6c2b589b11ec10","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"document/apply.go","path":"github.com/livecodelife/provenance/record"},"kwargs":"sha256:343e7a0232ac4b5ecbee9a4adc7248bbec83da7c8f5a0219e1d45fe33073ea1d","body":"sha256:acbf8628061291d11aa39e9ea64ecc91095a117057cfe6f1958ef2c437fac003"}
import "github.com/livecodelife/provenance/record"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setVar {"v":2,"tier":"owned","record":"prov-2026-2a640ada","gen":"sha256:52f742d9bf131f6a93dc56d698a431e545e822fbbf32baaaf87b6111f45edaa6","pkg":"sha256:c7b9d9ac1655e8211b8b2ca8c49cf34b8d510d8d823d4f5eb60ec7e1018aa2eb","id":{"file":"document/apply.go","name":"ErrNoReason"},"kwargs":"sha256:77605943d9c650cf5d6d70142dd5142917abb1a5e36602a082efce89ccb821bf","body":"sha256:7c91f9754a4a0159d50e2b6cffb77df2ab096fd4726dbb101bee5463bd034618"}
var ErrNoReason = errors.New("a reason is required")

// /sedum:setVar
// sedum:setType {"v":2,"tier":"owned","record":"prov-2026-44fce5f4","gen":"sha256:990c33c8ef5f32da36a8b5d627e6a67e380c8ec17d2154fa5f512ddec06ba907","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","id":{"file":"document/apply.go","name":"ApplyOptions"},"kwargs":"sha256:493884a9be8c49e4e7c03bbc53a5aebf45eda500cf591611778f7f6eea21e1b7","body":"sha256:a232f645f6d8d58b53c3776fdfaff9abccdc5728d3e611c48eb948b91a7e3a62"}
type ApplyOptions struct {
	Reason        string
	Author        string
	SkipConflicts bool
}

// /sedum:setType
// sedum:setType {"v":2,"tier":"owned","record":"prov-2026-658e3026","gen":"sha256:4a81ce9a48abddcdca4bd927a238fa8ee41ba5542db58f8942ab59f0a3ff88d2","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","id":{"file":"document/apply.go","name":"Revised"},"kwargs":"sha256:067726ba89339b9db0ca822cfe77233da90d78229cee3146322e8346160a7f3a","body":"sha256:e71f6d440e58ed2ed26ce5950cdfd54c82505d1f1f24b699a980a0331d573d14"}
type Revised struct {
	ID      string
	Version int
	Draft   bool
}

// /sedum:setType
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-6b2ceecb","gen":"sha256:a1af9e73e0651d18088c29216aa26fbfe9fa6d1c43db78b15a6c2b589b11ec10","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"document/apply.go","name":"Apply"},"kwargs":"sha256:0e7f6f94df2ff4e1ae8410f7bf07921b6962a6fb24ea70db0979979e6c1b7641","body":"sha256:adffe0c072b267231bd0edc8c1b40f74959140a5f60a815c2664f5006b2e6a8e"}
func Apply(store record.Store, plan *Plan, opts ApplyOptions) ([]Revised, error) {
	if plan == nil {
		return nil, fmt.Errorf("nil plan")
	}
	reason := strings.TrimSpace(opts.Reason)
	if reason == "" {
		return nil, ErrNoReason
	}
	if len(plan.Conflicts) > 0 && !opts.SkipConflicts {
		return nil, fmt.Errorf("%w: %d conflicts", ErrConflicts, len(plan.Conflicts))
	}
	var revised []Revised
	for _, ch := range plan.Changes {
		rec, doc, err := store.Load(ch.ID)
		if err != nil {
			return revised, fmt.Errorf("%s: %w", ch.ID, err)
		}
		if rec.Version != ch.CurrentVersion {
			return revised, fmt.Errorf("%s: %w", ch.ID, ErrStale)
		}
		if ch.IntentChanged {
			rec.Intent = ch.Intent
		}
		if ch.ConstraintsChanged {
			rec.Constraints = ch.Constraints
		}
		versions, err := store.Versions(ch.ID)
		if err != nil {
			return revised, fmt.Errorf("%s: %w", ch.ID, err)
		}
		if len(versions) == 0 {
			if rec.Schema != 2 {
				return revised, fmt.Errorf("%s: %w", ch.ID, record.ErrNotV2)
			}
			if err := store.Save(rec, doc); err != nil {
				return revised, fmt.Errorf("%s: %w", ch.ID, err)
			}
			revised = append(revised, Revised{ID: ch.ID, Version: rec.Version, Draft: true})
			continue
		}
		snap, err := record.Commit(store, rec, doc, record.Bump{Reason: reason, Author: opts.Author})
		if err != nil {
			return revised, fmt.Errorf("%s: %w", ch.ID, err)
		}
		revised = append(revised, Revised{ID: ch.ID, Version: snap.Version, Draft: false})
	}
	return revised, nil
}

// /sedum:setFunction
// sedum:setVar {"v":2,"tier":"owned","record":"prov-2026-9b7dc23c","gen":"sha256:29de6917dd3707e23d4774cdb633e7359789fddaaa0907bf49c1ce937c537d4c","pkg":"sha256:c7b9d9ac1655e8211b8b2ca8c49cf34b8d510d8d823d4f5eb60ec7e1018aa2eb","id":{"file":"document/apply.go","name":"ErrConflicts"},"kwargs":"sha256:4c03a26bba5af50d4e45b350cf4d5a75c16bf236079b568117a67ce753b0a747","body":"sha256:4a6dcc531a942ee56e40ef30aa9464cf1c2dedee68847d3808f301e0b74116e9"}
var ErrConflicts = errors.New("the plan has conflicting sections")

// /sedum:setVar
// sedum:setVar {"v":2,"tier":"owned","record":"prov-2026-9ca50ce0","gen":"sha256:965a4de386318c18b6f7ae783e5a7094d4f7c5327b976ae802766665617e7c1d","pkg":"sha256:c7b9d9ac1655e8211b8b2ca8c49cf34b8d510d8d823d4f5eb60ec7e1018aa2eb","id":{"file":"document/apply.go","name":"ErrStale"},"kwargs":"sha256:ecf1c4b129482198c416e219a4d7a8d82f50274cd7fe530aa456af77f8e5c7f8","body":"sha256:e38fc37785b416c773293219931e9b2348b1407d4e23c78fc3867b7cea10c15b"}
var ErrStale = errors.New("the record changed since the plan was made")

// /sedum:setVar
