package document

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-400f65dc","gen":"sha256:f2ea87ada4c5613576689a6782313b60daf8e401e5fcc802a5de66bb460160bd","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"document/render.go","path":"strings"}}
import "strings"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-f40f3a4a","gen":"sha256:926745d7d9fd808dbc039734a69f175728b8c27d2558a26dc904f0ed503a8f7c","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"document/render.go","path":"sort"}}
import "sort"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-f40f3a4a","gen":"sha256:926745d7d9fd808dbc039734a69f175728b8c27d2558a26dc904f0ed503a8f7c","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"document/render.go","path":"strconv"}}
import "strconv"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-f40f3a4a","gen":"sha256:926745d7d9fd808dbc039734a69f175728b8c27d2558a26dc904f0ed503a8f7c","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"document/render.go","path":"github.com/livecodelife/provenance/scope"}}
import "github.com/livecodelife/provenance/scope"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-f40f3a4a","gen":"sha256:926745d7d9fd808dbc039734a69f175728b8c27d2558a26dc904f0ed503a8f7c","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"document/render.go","path":"github.com/livecodelife/provenance/record"}}
import "github.com/livecodelife/provenance/record"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setType {"tier":"owned","record":"prov-2026-2b7b6b93","gen":"sha256:c87e3eb4991db295a018bc915719ab304c8c0dcd12d66f28c1b116352086ea14","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type Change struct {\n\tID                 string\n\tTitle              string\n\tAnchoredVersion    int\n\tCurrentVersion     int\n\tIntentChanged      bool\n\tConstraintsChanged bool\n\tIntent             string\n\tConstraints        []string\n}","file":"document/render.go","name":"Change"}}
type Change struct {
	ID                 string
	Title              string
	AnchoredVersion    int
	CurrentVersion     int
	IntentChanged      bool
	ConstraintsChanged bool
	Intent             string
	Constraints        []string
}

// /sedum:setType
// sedum:setFunction {"tier":"owned","record":"prov-2026-400f65dc","gen":"sha256:f2ea87ada4c5613576689a6782313b60daf8e401e5fcc802a5de66bb460160bd","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func isMarkerLine(line string) bool {\n\tlead := strings.TrimLeft(line, \" \\t\")\n\tif strings.HasPrefix(lead, \"## \") || strings.HasPrefix(lead, \"### \") || strings.HasPrefix(lead, \"<!-- provenance:\") {\n\t\treturn true\n\t}\n\treturn strings.TrimSpace(line) == \"**Constraints**\"\n}","file":"document/render.go","name":"isMarkerLine"}}
func isMarkerLine(line string) bool {
	lead := strings.TrimLeft(line, " \t")
	if strings.HasPrefix(lead, "## ") || strings.HasPrefix(lead, "### ") || strings.HasPrefix(lead, "<!-- provenance:") {
		return true
	}
	return strings.TrimSpace(line) == "**Constraints**"
}

// /sedum:setFunction
// sedum:setType {"tier":"owned","record":"prov-2026-477b93e0","gen":"sha256:d933a82d00718de7b4eb9a6b8e1e97779d83b0d8c809f3121443bad283290ff4","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type RenderOptions struct {\n\tScopes []string\n\tIDs    bool\n}","file":"document/render.go","name":"RenderOptions"}}
type RenderOptions struct {
	Scopes []string
	IDs    bool
}

// /sedum:setType
// sedum:setFunction {"tier":"owned","record":"prov-2026-5e341891","gen":"sha256:d6fe143223cdc69bba6cafcf15ada60ea3948b01b9f8136044d383eb9e08ee43","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func escapeLine(line string) string {\n\ttrimmed := strings.TrimLeft(line, \"\\\\\")\n\tif isMarkerLine(trimmed) {\n\t\treturn \"\\\\\" + line\n\t}\n\treturn line\n}","file":"document/render.go","name":"escapeLine"}}
func escapeLine(line string) string {
	trimmed := strings.TrimLeft(line, "\\")
	if isMarkerLine(trimmed) {
		return "\\" + line
	}
	return line
}

// /sedum:setFunction
// sedum:setType {"tier":"owned","record":"prov-2026-6eeaa2bc","gen":"sha256:85066f8dcbc8074c00cd5308255455f52b8996a0a6c81769cdbb3bdecc5740fd","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type Plan struct {\n\tChanges   []Change\n\tConflicts []Problem\n\tProblems  []Problem\n\tUnchanged []string\n}","file":"document/render.go","name":"Plan"}}
type Plan struct {
	Changes   []Change
	Conflicts []Problem
	Problems  []Problem
	Unchanged []string
}

// /sedum:setType
// sedum:setType {"tier":"owned","record":"prov-2026-b396ccea","gen":"sha256:f937af628fac653774d7d299772d396db915c0a672cae94ee63fce275d93dc72","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type Problem struct {\n\tKind            string\n\tID              string\n\tLine            int\n\tAnchoredVersion int\n\tCurrentVersion  int\n\tMessage         string\n}","file":"document/render.go","name":"Problem"}}
type Problem struct {
	Kind            string
	ID              string
	Line            int
	AnchoredVersion int
	CurrentVersion  int
	Message         string
}

// /sedum:setType
// sedum:setType {"tier":"owned","record":"prov-2026-b98853b3","gen":"sha256:e84b3fc119d69341d66e566c7536233e244f141dee239ae94a389c2ecfff85c4","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type Section struct {\n\tHeading        string\n\tLevel          int\n\tLine           int\n\tID             string\n\tVersion        int\n\tAnchored       bool\n\tIntent         string\n\tConstraints    []string\n\tHasConstraints bool\n}","file":"document/render.go","name":"Section"}}
type Section struct {
	Heading        string
	Level          int
	Line           int
	ID             string
	Version        int
	Anchored       bool
	Intent         string
	Constraints    []string
	HasConstraints bool
}

// /sedum:setType
// sedum:setFunction {"tier":"owned","record":"prov-2026-f40f3a4a","gen":"sha256:926745d7d9fd808dbc039734a69f175728b8c27d2558a26dc904f0ed503a8f7c","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func Render(store record.Store, opts RenderOptions) (string, error) {\n\tids, err := store.List()\n\tif err != nil {\n\t\treturn \"\", err\n\t}\n\tvar briefs, blueprints []*record.Record\n\tfor _, id := range ids {\n\t\trec, _, err := store.Load(id)\n\t\tif err != nil {\n\t\t\treturn \"\", err\n\t\t}\n\t\ttyp := rec.EffectiveType()\n\t\tif typ != record.TypeBrief && typ != record.TypeBlueprint {\n\t\t\tcontinue\n\t\t}\n\t\tif rec.Status == record.StatusDeprecated || rec.Status == record.StatusSuperseded {\n\t\t\tcontinue\n\t\t}\n\t\tif len(opts.Scopes) > 0 {\n\t\t\tkeep := false\n\t\t\tfor _, pattern := range rec.AffectedScope {\n\t\t\t\tfor _, path := range opts.Scopes {\n\t\t\t\t\tok, err := scope.Match(path, pattern)\n\t\t\t\t\tif err != nil {\n\t\t\t\t\t\treturn \"\", err\n\t\t\t\t\t}\n\t\t\t\t\tif ok {\n\t\t\t\t\t\tkeep = true\n\t\t\t\t\t\tbreak\n\t\t\t\t\t}\n\t\t\t\t}\n\t\t\t\tif keep {\n\t\t\t\t\tbreak\n\t\t\t\t}\n\t\t\t}\n\t\t\tif !keep {\n\t\t\t\tcontinue\n\t\t\t}\n\t\t}\n\t\tif typ == record.TypeBrief {\n\t\t\tbriefs = append(briefs, rec)\n\t\t} else {\n\t\t\tblueprints = append(blueprints, rec)\n\t\t}\n\t}\n\tsort.SliceStable(briefs, func(i, j int) bool { return briefs[i].ID < briefs[j].ID })\n\tsort.SliceStable(blueprints, func(i, j int) bool { return blueprints[i].ID < blueprints[j].ID })\n\tattached := make(map[string]bool)\n\tvar sections []string\n\tfor _, brief := range briefs {\n\t\tsections = append(sections, renderSection(brief, \"## \", opts.IDs))\n\t\tfor _, bp := range blueprints {\n\t\t\tif bp.ImplementsRepo == \"\" && bp.Implements == brief.ID {\n\t\t\t\tattached[bp.ID] = true\n\t\t\t\tsections = append(sections, renderSection(bp, \"### \", opts.IDs))\n\t\t\t}\n\t\t}\n\t}\n\tfor _, bp := range blueprints {\n\t\tif !attached[bp.ID] {\n\t\t\tsections = append(sections, renderSection(bp, \"### \", opts.IDs))\n\t\t}\n\t}\n\tif len(sections) == 0 {\n\t\treturn \"\", nil\n\t}\n\treturn strings.Join(sections, \"\\n\\n\") + \"\\n\", nil\n}","file":"document/render.go","name":"Render"}}
func Render(store record.Store, opts RenderOptions) (string, error) {
	ids, err := store.List()
	if err != nil {
		return "", err
	}
	var briefs, blueprints []*record.Record
	for _, id := range ids {
		rec, _, err := store.Load(id)
		if err != nil {
			return "", err
		}
		typ := rec.EffectiveType()
		if typ != record.TypeBrief && typ != record.TypeBlueprint {
			continue
		}
		if rec.Status == record.StatusDeprecated || rec.Status == record.StatusSuperseded {
			continue
		}
		if len(opts.Scopes) > 0 {
			keep := false
			for _, pattern := range rec.AffectedScope {
				for _, path := range opts.Scopes {
					ok, err := scope.Match(path, pattern)
					if err != nil {
						return "", err
					}
					if ok {
						keep = true
						break
					}
				}
				if keep {
					break
				}
			}
			if !keep {
				continue
			}
		}
		if typ == record.TypeBrief {
			briefs = append(briefs, rec)
		} else {
			blueprints = append(blueprints, rec)
		}
	}
	sort.SliceStable(briefs, func(i, j int) bool { return briefs[i].ID < briefs[j].ID })
	sort.SliceStable(blueprints, func(i, j int) bool { return blueprints[i].ID < blueprints[j].ID })
	attached := make(map[string]bool)
	var sections []string
	for _, brief := range briefs {
		sections = append(sections, renderSection(brief, "## ", opts.IDs))
		for _, bp := range blueprints {
			if bp.ImplementsRepo == "" && bp.Implements == brief.ID {
				attached[bp.ID] = true
				sections = append(sections, renderSection(bp, "### ", opts.IDs))
			}
		}
	}
	for _, bp := range blueprints {
		if !attached[bp.ID] {
			sections = append(sections, renderSection(bp, "### ", opts.IDs))
		}
	}
	if len(sections) == 0 {
		return "", nil
	}
	return strings.Join(sections, "\n\n") + "\n", nil
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-f40f3a4a","gen":"sha256:926745d7d9fd808dbc039734a69f175728b8c27d2558a26dc904f0ed503a8f7c","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func renderSection(rec *record.Record, prefix string, withIDs bool) string {\n\ttitle := strings.Join(strings.Fields(rec.Title), \" \")\n\theading := prefix + title\n\tif withIDs {\n\t\theading += \" (\" + rec.ID + \")\"\n\t}\n\tlines := []string{\n\t\theading,\n\t\t\"<!-- provenance:\" + rec.ID + \"@\" + strconv.Itoa(rec.Version) + \" -->\",\n\t\t\"\",\n\t}\n\tif intent := strings.TrimSpace(rec.Intent); intent != \"\" {\n\t\tfor _, line := range strings.Split(intent, \"\\n\") {\n\t\t\tlines = append(lines, escapeLine(line))\n\t\t}\n\t\tlines = append(lines, \"\")\n\t}\n\tlines = append(lines, \"**Constraints**\")\n\tfor _, constraint := range rec.Constraints {\n\t\tparts := strings.Split(strings.TrimSpace(constraint), \"\\n\")\n\t\tfor i, part := range parts {\n\t\t\tswitch {\n\t\t\tcase i == 0:\n\t\t\t\tlines = append(lines, \"- \"+part)\n\t\t\tcase part == \"\":\n\t\t\t\tlines = append(lines, \"  \")\n\t\t\tdefault:\n\t\t\t\tlines = append(lines, \"  \"+part)\n\t\t\t}\n\t\t}\n\t}\n\treturn strings.Join(lines, \"\\n\")\n}","file":"document/render.go","name":"renderSection"}}
func renderSection(rec *record.Record, prefix string, withIDs bool) string {
	title := strings.Join(strings.Fields(rec.Title), " ")
	heading := prefix + title
	if withIDs {
		heading += " (" + rec.ID + ")"
	}
	lines := []string{
		heading,
		"<!-- provenance:" + rec.ID + "@" + strconv.Itoa(rec.Version) + " -->",
		"",
	}
	if intent := strings.TrimSpace(rec.Intent); intent != "" {
		for _, line := range strings.Split(intent, "\n") {
			lines = append(lines, escapeLine(line))
		}
		lines = append(lines, "")
	}
	lines = append(lines, "**Constraints**")
	for _, constraint := range rec.Constraints {
		parts := strings.Split(strings.TrimSpace(constraint), "\n")
		for i, part := range parts {
			switch {
			case i == 0:
				lines = append(lines, "- "+part)
			case part == "":
				lines = append(lines, "  ")
			default:
				lines = append(lines, "  "+part)
			}
		}
	}
	return strings.Join(lines, "\n")
}

// /sedum:setFunction
