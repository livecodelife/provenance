package document

// sedum:anchor:imports
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-4dc24541","gen":"sha256:be56d89178fda13194f238070158f35db3fbc7565b43a68147fa30e2c1db87c8","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"document/plan.go","path":"strings"},"kwargs":"sha256:198a4768f17bf0255ba1e251d5d419dd140b0a45f85925fcfce70cf7ad5a57a8","body":"sha256:a9af3c0bc4bf9342df9d7e89a324523fdba7941a9ca4190270e68ff3c4a25ea6"}
import "strings"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-c92e5c96","gen":"sha256:4fa5d953628f759c4267c2a785c776489d62ce8fb11779c36962a12ae57c81a4","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"document/plan.go","path":"sort"},"kwargs":"sha256:53e97781c464fa19ab6b84caebe82ca711425f4449a58d01b8a0baa3c762b820","body":"sha256:f197172617ae21ef23b1855b09aafb14571020ae81bff877d63d3ff81f4bf63c"}
import "sort"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-c92e5c96","gen":"sha256:4fa5d953628f759c4267c2a785c776489d62ce8fb11779c36962a12ae57c81a4","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"document/plan.go","path":"github.com/livecodelife/provenance/record"},"kwargs":"sha256:0eac8f0e66441874ba9e00fd3862db4b3b2ded39a4383c3076a8e69c8ee399c4","body":"sha256:acbf8628061291d11aa39e9ea64ecc91095a117057cfe6f1958ef2c437fac003"}
import "github.com/livecodelife/provenance/record"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-4dc24541","gen":"sha256:be56d89178fda13194f238070158f35db3fbc7565b43a68147fa30e2c1db87c8","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"document/plan.go","name":"normalize"},"kwargs":"sha256:5ae948dd8ff52d399bc537c9a92174d9578a1ae0ae4cc3d7058228d98b2f6efc","body":"sha256:70afae29de8508eea6d54ce188f5e650fd8ce2d7e43df40ae41d16f1c9eebd9a"}
func normalize(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

// /sedum:setFunction
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-c92e5c96","gen":"sha256:4fa5d953628f759c4267c2a785c776489d62ce8fb11779c36962a12ae57c81a4","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"document/plan.go","name":"BuildPlan"},"kwargs":"sha256:983da30479d27060f5079c4e5f6f3fbd5a1babe8b9dda860ba7bd9060c1fcb80","body":"sha256:5db37cb3d7431e1445defd89730594e26f780f0e4211b10edbbede308eec49be"}
func BuildPlan(store record.Store, text string) (*Plan, error) {
	ids, err := store.List()
	if err != nil {
		return nil, err
	}
	known := make(map[string]bool, len(ids))
	for _, id := range ids {
		known[id] = true
	}
	itoa := func(n int) string {
		if n == 0 {
			return "0"
		}
		var b []byte
		for n > 0 {
			b = append([]byte{byte('0' + n%10)}, b...)
			n /= 10
		}
		return string(b)
	}
	plan := &Plan{}
	seen := make(map[string]bool)
	add := func(kind, id string, line int, msg string) {
		plan.Problems = append(plan.Problems, Problem{Kind: kind, ID: id, Line: line, Message: msg})
	}
	for _, s := range Parse(text) {
		if !s.Anchored {
			add("unanchored", "", s.Line, "heading has no anchor")
			continue
		}
		if !known[s.ID] {
			add("unknown", s.ID, s.Line, "no record with this id")
			continue
		}
		rec, _, err := store.Load(s.ID)
		if err != nil {
			return nil, err
		}
		rtype := rec.EffectiveType()
		if rtype != "brief" && rtype != "blueprint" {
			add("imprint", s.ID, s.Line, "the record is not a brief or blueprint")
			continue
		}
		if rec.Status == "deprecated" || rec.Status == "superseded" {
			add("inactive", s.ID, s.Line, "the record is deprecated or superseded")
			continue
		}
		if seen[s.ID] {
			add("duplicate", s.ID, s.Line, "this record is anchored more than once")
			continue
		}
		seen[s.ID] = true
		if (rtype == "brief" && s.Level == 3) || (rtype == "blueprint" && s.Level == 2) {
			add("level", s.ID, s.Line, "heading level does not match the record type")
			continue
		}
		heading := normalize(s.Heading)
		if heading != normalize(rec.Title) && heading != normalize(rec.Title+" ("+rec.ID+")") {
			add("heading", s.ID, s.Line, "heading text differs from the record title")
			continue
		}
		if !s.HasConstraints {
			add("malformed", s.ID, s.Line, "section has no constraints marker")
			continue
		}
		if s.Version != rec.Version {
			plan.Conflicts = append(plan.Conflicts, Problem{
				Kind:            "conflict",
				ID:              s.ID,
				Line:            s.Line,
				Message:         "the record is at version " + itoa(rec.Version) + ", the section is from version " + itoa(s.Version),
				AnchoredVersion: s.Version,
				CurrentVersion:  rec.Version,
			})
			continue
		}
		intentChanged := normalize(s.Intent) != normalize(rec.Intent)
		constraintsChanged := len(s.Constraints) != len(rec.Constraints)
		if !constraintsChanged {
			for i := range s.Constraints {
				if normalize(s.Constraints[i]) != normalize(rec.Constraints[i]) {
					constraintsChanged = true
					break
				}
			}
		}
		if !intentChanged && !constraintsChanged {
			plan.Unchanged = append(plan.Unchanged, s.ID)
			continue
		}
		plan.Changes = append(plan.Changes, Change{
			ID:                 rec.ID,
			Title:              rec.Title,
			AnchoredVersion:    s.Version,
			CurrentVersion:     rec.Version,
			IntentChanged:      intentChanged,
			ConstraintsChanged: constraintsChanged,
			Intent:             s.Intent,
			Constraints:        s.Constraints,
		})
	}
	sort.SliceStable(plan.Changes, func(i, j int) bool { return plan.Changes[i].ID < plan.Changes[j].ID })
	sort.SliceStable(plan.Conflicts, func(i, j int) bool { return plan.Conflicts[i].ID < plan.Conflicts[j].ID })
	sort.Strings(plan.Unchanged)
	sort.SliceStable(plan.Problems, func(i, j int) bool {
		a, b := plan.Problems[i], plan.Problems[j]
		if a.ID != b.ID {
			return a.ID < b.ID
		}
		return a.Line < b.Line
	})
	return plan, nil
}

// /sedum:setFunction
