package document

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-4dc24541","gen":"sha256:be56d89178fda13194f238070158f35db3fbc7565b43a68147fa30e2c1db87c8","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"document/plan.go","path":"strings"}}
import "strings"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-c92e5c96","gen":"sha256:4fa5d953628f759c4267c2a785c776489d62ce8fb11779c36962a12ae57c81a4","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"document/plan.go","path":"sort"}}
import "sort"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-c92e5c96","gen":"sha256:4fa5d953628f759c4267c2a785c776489d62ce8fb11779c36962a12ae57c81a4","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"document/plan.go","path":"github.com/livecodelife/provenance/record"}}
import "github.com/livecodelife/provenance/record"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"tier":"owned","record":"prov-2026-4dc24541","gen":"sha256:be56d89178fda13194f238070158f35db3fbc7565b43a68147fa30e2c1db87c8","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func normalize(s string) string {\n\treturn strings.Join(strings.Fields(s), \" \")\n}","file":"document/plan.go","name":"normalize"}}
func normalize(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-c92e5c96","gen":"sha256:4fa5d953628f759c4267c2a785c776489d62ce8fb11779c36962a12ae57c81a4","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func BuildPlan(store record.Store, text string) (*Plan, error) {\n\tids, err := store.List()\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\tknown := make(map[string]bool, len(ids))\n\tfor _, id := range ids {\n\t\tknown[id] = true\n\t}\n\titoa := func(n int) string {\n\t\tif n == 0 {\n\t\t\treturn \"0\"\n\t\t}\n\t\tvar b []byte\n\t\tfor n > 0 {\n\t\t\tb = append([]byte{byte('0' + n%10)}, b...)\n\t\t\tn /= 10\n\t\t}\n\t\treturn string(b)\n\t}\n\tplan := &Plan{}\n\tseen := make(map[string]bool)\n\tadd := func(kind, id string, line int, msg string) {\n\t\tplan.Problems = append(plan.Problems, Problem{Kind: kind, ID: id, Line: line, Message: msg})\n\t}\n\tfor _, s := range Parse(text) {\n\t\tif !s.Anchored {\n\t\t\tadd(\"unanchored\", \"\", s.Line, \"heading has no anchor\")\n\t\t\tcontinue\n\t\t}\n\t\tif !known[s.ID] {\n\t\t\tadd(\"unknown\", s.ID, s.Line, \"no record with this id\")\n\t\t\tcontinue\n\t\t}\n\t\trec, _, err := store.Load(s.ID)\n\t\tif err != nil {\n\t\t\treturn nil, err\n\t\t}\n\t\trtype := rec.EffectiveType()\n\t\tif rtype != \"brief\" && rtype != \"blueprint\" {\n\t\t\tadd(\"imprint\", s.ID, s.Line, \"the record is not a brief or blueprint\")\n\t\t\tcontinue\n\t\t}\n\t\tif rec.Status == \"deprecated\" || rec.Status == \"superseded\" {\n\t\t\tadd(\"inactive\", s.ID, s.Line, \"the record is deprecated or superseded\")\n\t\t\tcontinue\n\t\t}\n\t\tif seen[s.ID] {\n\t\t\tadd(\"duplicate\", s.ID, s.Line, \"this record is anchored more than once\")\n\t\t\tcontinue\n\t\t}\n\t\tseen[s.ID] = true\n\t\tif (rtype == \"brief\" && s.Level == 3) || (rtype == \"blueprint\" && s.Level == 2) {\n\t\t\tadd(\"level\", s.ID, s.Line, \"heading level does not match the record type\")\n\t\t\tcontinue\n\t\t}\n\t\theading := normalize(s.Heading)\n\t\tif heading != normalize(rec.Title) && heading != normalize(rec.Title+\" (\"+rec.ID+\")\") {\n\t\t\tadd(\"heading\", s.ID, s.Line, \"heading text differs from the record title\")\n\t\t\tcontinue\n\t\t}\n\t\tif !s.HasConstraints {\n\t\t\tadd(\"malformed\", s.ID, s.Line, \"section has no constraints marker\")\n\t\t\tcontinue\n\t\t}\n\t\tif s.Version != rec.Version {\n\t\t\tplan.Conflicts = append(plan.Conflicts, Problem{\n\t\t\t\tKind:            \"conflict\",\n\t\t\t\tID:              s.ID,\n\t\t\t\tLine:            s.Line,\n\t\t\t\tMessage:         \"the record is at version \" + itoa(rec.Version) + \", the section is from version \" + itoa(s.Version),\n\t\t\t\tAnchoredVersion: s.Version,\n\t\t\t\tCurrentVersion:  rec.Version,\n\t\t\t})\n\t\t\tcontinue\n\t\t}\n\t\tintentChanged := normalize(s.Intent) != normalize(rec.Intent)\n\t\tconstraintsChanged := len(s.Constraints) != len(rec.Constraints)\n\t\tif !constraintsChanged {\n\t\t\tfor i := range s.Constraints {\n\t\t\t\tif normalize(s.Constraints[i]) != normalize(rec.Constraints[i]) {\n\t\t\t\t\tconstraintsChanged = true\n\t\t\t\t\tbreak\n\t\t\t\t}\n\t\t\t}\n\t\t}\n\t\tif !intentChanged && !constraintsChanged {\n\t\t\tplan.Unchanged = append(plan.Unchanged, s.ID)\n\t\t\tcontinue\n\t\t}\n\t\tplan.Changes = append(plan.Changes, Change{\n\t\t\tID:                 rec.ID,\n\t\t\tTitle:              rec.Title,\n\t\t\tAnchoredVersion:    s.Version,\n\t\t\tCurrentVersion:     rec.Version,\n\t\t\tIntentChanged:      intentChanged,\n\t\t\tConstraintsChanged: constraintsChanged,\n\t\t\tIntent:             s.Intent,\n\t\t\tConstraints:        s.Constraints,\n\t\t})\n\t}\n\tsort.SliceStable(plan.Changes, func(i, j int) bool { return plan.Changes[i].ID < plan.Changes[j].ID })\n\tsort.SliceStable(plan.Conflicts, func(i, j int) bool { return plan.Conflicts[i].ID < plan.Conflicts[j].ID })\n\tsort.Strings(plan.Unchanged)\n\tsort.SliceStable(plan.Problems, func(i, j int) bool {\n\t\ta, b := plan.Problems[i], plan.Problems[j]\n\t\tif a.ID != b.ID {\n\t\t\treturn a.ID < b.ID\n\t\t}\n\t\treturn a.Line < b.Line\n\t})\n\treturn plan, nil\n}","file":"document/plan.go","name":"BuildPlan"}}
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
