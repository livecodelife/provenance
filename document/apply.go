package document

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-2a640ada","gen":"sha256:52f742d9bf131f6a93dc56d698a431e545e822fbbf32baaaf87b6111f45edaa6","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"document/apply.go","path":"errors"}}
import "errors"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-6b2ceecb","gen":"sha256:a1af9e73e0651d18088c29216aa26fbfe9fa6d1c43db78b15a6c2b589b11ec10","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"document/apply.go","path":"fmt"}}
import "fmt"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-6b2ceecb","gen":"sha256:a1af9e73e0651d18088c29216aa26fbfe9fa6d1c43db78b15a6c2b589b11ec10","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"document/apply.go","path":"strings"}}
import "strings"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-6b2ceecb","gen":"sha256:a1af9e73e0651d18088c29216aa26fbfe9fa6d1c43db78b15a6c2b589b11ec10","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"document/apply.go","path":"github.com/livecodelife/provenance/record"}}
import "github.com/livecodelife/provenance/record"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setVar {"tier":"owned","record":"prov-2026-2a640ada","gen":"sha256:52f742d9bf131f6a93dc56d698a431e545e822fbbf32baaaf87b6111f45edaa6","pkg":"sha256:c7b9d9ac1655e8211b8b2ca8c49cf34b8d510d8d823d4f5eb60ec7e1018aa2eb","kwargs":{"code":"var ErrNoReason = errors.New(\"a reason is required\")","file":"document/apply.go","name":"ErrNoReason"}}
var ErrNoReason = errors.New("a reason is required")

// /sedum:setVar
// sedum:setType {"tier":"owned","record":"prov-2026-44fce5f4","gen":"sha256:990c33c8ef5f32da36a8b5d627e6a67e380c8ec17d2154fa5f512ddec06ba907","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type ApplyOptions struct {\n\tReason        string\n\tAuthor        string\n\tSkipConflicts bool\n}","file":"document/apply.go","name":"ApplyOptions"}}
type ApplyOptions struct {
	Reason        string
	Author        string
	SkipConflicts bool
}

// /sedum:setType
// sedum:setType {"tier":"owned","record":"prov-2026-658e3026","gen":"sha256:4a81ce9a48abddcdca4bd927a238fa8ee41ba5542db58f8942ab59f0a3ff88d2","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type Revised struct {\n\tID      string\n\tVersion int\n\tDraft   bool\n}","file":"document/apply.go","name":"Revised"}}
type Revised struct {
	ID      string
	Version int
	Draft   bool
}

// /sedum:setType
// sedum:setFunction {"tier":"owned","record":"prov-2026-6b2ceecb","gen":"sha256:a1af9e73e0651d18088c29216aa26fbfe9fa6d1c43db78b15a6c2b589b11ec10","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func Apply(store record.Store, plan *Plan, opts ApplyOptions) ([]Revised, error) {\n\tif plan == nil {\n\t\treturn nil, fmt.Errorf(\"nil plan\")\n\t}\n\treason := strings.TrimSpace(opts.Reason)\n\tif reason == \"\" {\n\t\treturn nil, ErrNoReason\n\t}\n\tif len(plan.Conflicts) > 0 && !opts.SkipConflicts {\n\t\treturn nil, fmt.Errorf(\"%w: %d conflicts\", ErrConflicts, len(plan.Conflicts))\n\t}\n\tvar revised []Revised\n\tfor _, ch := range plan.Changes {\n\t\trec, doc, err := store.Load(ch.ID)\n\t\tif err != nil {\n\t\t\treturn revised, fmt.Errorf(\"%s: %w\", ch.ID, err)\n\t\t}\n\t\tif rec.Version != ch.CurrentVersion {\n\t\t\treturn revised, fmt.Errorf(\"%s: %w\", ch.ID, ErrStale)\n\t\t}\n\t\tif ch.IntentChanged {\n\t\t\trec.Intent = ch.Intent\n\t\t}\n\t\tif ch.ConstraintsChanged {\n\t\t\trec.Constraints = ch.Constraints\n\t\t}\n\t\tversions, err := store.Versions(ch.ID)\n\t\tif err != nil {\n\t\t\treturn revised, fmt.Errorf(\"%s: %w\", ch.ID, err)\n\t\t}\n\t\tif len(versions) == 0 {\n\t\t\tif rec.Schema != 2 {\n\t\t\t\treturn revised, fmt.Errorf(\"%s: %w\", ch.ID, record.ErrNotV2)\n\t\t\t}\n\t\t\tif err := store.Save(rec, doc); err != nil {\n\t\t\t\treturn revised, fmt.Errorf(\"%s: %w\", ch.ID, err)\n\t\t\t}\n\t\t\trevised = append(revised, Revised{ID: ch.ID, Version: rec.Version, Draft: true})\n\t\t\tcontinue\n\t\t}\n\t\tsnap, err := record.Commit(store, rec, doc, record.Bump{Reason: reason, Author: opts.Author})\n\t\tif err != nil {\n\t\t\treturn revised, fmt.Errorf(\"%s: %w\", ch.ID, err)\n\t\t}\n\t\trevised = append(revised, Revised{ID: ch.ID, Version: snap.Version, Draft: false})\n\t}\n\treturn revised, nil\n}","file":"document/apply.go","name":"Apply"}}
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
// sedum:setVar {"tier":"owned","record":"prov-2026-9b7dc23c","gen":"sha256:29de6917dd3707e23d4774cdb633e7359789fddaaa0907bf49c1ce937c537d4c","pkg":"sha256:c7b9d9ac1655e8211b8b2ca8c49cf34b8d510d8d823d4f5eb60ec7e1018aa2eb","kwargs":{"code":"var ErrConflicts = errors.New(\"the plan has conflicting sections\")","file":"document/apply.go","name":"ErrConflicts"}}
var ErrConflicts = errors.New("the plan has conflicting sections")

// /sedum:setVar
// sedum:setVar {"tier":"owned","record":"prov-2026-9ca50ce0","gen":"sha256:965a4de386318c18b6f7ae783e5a7094d4f7c5327b976ae802766665617e7c1d","pkg":"sha256:c7b9d9ac1655e8211b8b2ca8c49cf34b8d510d8d823d4f5eb60ec7e1018aa2eb","kwargs":{"code":"var ErrStale = errors.New(\"the record changed since the plan was made\")","file":"document/apply.go","name":"ErrStale"}}
var ErrStale = errors.New("the record changed since the plan was made")

// /sedum:setVar
