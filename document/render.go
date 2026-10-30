package document

// sedum:anchor:imports
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-400f65dc","gen":"sha256:f2ea87ada4c5613576689a6782313b60daf8e401e5fcc802a5de66bb460160bd","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"document/render.go","path":"strings"},"kwargs":"sha256:6bbc271b90b64d420c3c953b160c5877cef152df8a987bfbd7a2fa8353070778","body":"sha256:a9af3c0bc4bf9342df9d7e89a324523fdba7941a9ca4190270e68ff3c4a25ea6"}
import "strings"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-f40f3a4a","gen":"sha256:926745d7d9fd808dbc039734a69f175728b8c27d2558a26dc904f0ed503a8f7c","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"document/render.go","path":"sort"},"kwargs":"sha256:42eccd6bd74495a35864fe61694c8a21fa14adca60d73d96fc5ad453def90674","body":"sha256:f197172617ae21ef23b1855b09aafb14571020ae81bff877d63d3ff81f4bf63c"}
import "sort"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-f40f3a4a","gen":"sha256:926745d7d9fd808dbc039734a69f175728b8c27d2558a26dc904f0ed503a8f7c","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"document/render.go","path":"strconv"},"kwargs":"sha256:6332da0672dc7b8fd3695a7b7b66c7b1f9c7098643e6f326297e8dc80c58f2f7","body":"sha256:f33f981c64a42d78a5daa7ff3fca15b10045080cccd282fef544b59eaf7f6dc8"}
import "strconv"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-f40f3a4a","gen":"sha256:926745d7d9fd808dbc039734a69f175728b8c27d2558a26dc904f0ed503a8f7c","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"document/render.go","path":"github.com/livecodelife/provenance/scope"},"kwargs":"sha256:fe28c168d7ec7eac9bee7f1f521ac73663e5d8c41efc5aa954e76d5468a1508a","body":"sha256:4ec0116f2b0d829a4ed0e1f2a40f17d465f6f981c35c8fd5427721c94a450e53"}
import "github.com/livecodelife/provenance/scope"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-f40f3a4a","gen":"sha256:926745d7d9fd808dbc039734a69f175728b8c27d2558a26dc904f0ed503a8f7c","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"document/render.go","path":"github.com/livecodelife/provenance/record"},"kwargs":"sha256:917bf7ced41d87ff5762258f62ec09679f6a0a9f1cc45128f515e3cb68819363","body":"sha256:acbf8628061291d11aa39e9ea64ecc91095a117057cfe6f1958ef2c437fac003"}
import "github.com/livecodelife/provenance/record"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setType {"v":2,"tier":"owned","record":"prov-2026-2b7b6b93","gen":"sha256:c87e3eb4991db295a018bc915719ab304c8c0dcd12d66f28c1b116352086ea14","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","id":{"file":"document/render.go","name":"Change"},"kwargs":"sha256:16275c7dab399a3f168cb7b947cae871a15e685a78c0bb0878a7a16bd70b3c4d","body":"sha256:75fd113c028392eb72ff5655768f7deace9e0603829e4b75da4c6f8f974b1151"}
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
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-400f65dc","gen":"sha256:f2ea87ada4c5613576689a6782313b60daf8e401e5fcc802a5de66bb460160bd","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"document/render.go","name":"isMarkerLine"},"kwargs":"sha256:5314b32ab65a3e91ecdf4efa2d6f97d02ba570ca75cc2308cd511a7d7a77a64c","body":"sha256:fe5c665046d03ba8adb3a6ca7860b7501991b6123c6fdc786e372a7dc4d5c07c"}
func isMarkerLine(line string) bool {
	lead := strings.TrimLeft(line, " \t")
	if strings.HasPrefix(lead, "## ") || strings.HasPrefix(lead, "### ") || strings.HasPrefix(lead, "<!-- provenance:") {
		return true
	}
	return strings.TrimSpace(line) == "**Constraints**"
}

// /sedum:setFunction
// sedum:setType {"v":2,"tier":"owned","record":"prov-2026-477b93e0","gen":"sha256:d933a82d00718de7b4eb9a6b8e1e97779d83b0d8c809f3121443bad283290ff4","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","id":{"file":"document/render.go","name":"RenderOptions"},"kwargs":"sha256:ca49a3a212428920acff1bcc590233ba5e512af7a80472e2bca58ec9087356ec","body":"sha256:baeebfde73103cbbdb3fe8c2fb9ed6fc1e1e8d0e29d166676fb8aa94d2c04de1"}
type RenderOptions struct {
	Scopes []string
	IDs    bool
}

// /sedum:setType
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-5e341891","gen":"sha256:d6fe143223cdc69bba6cafcf15ada60ea3948b01b9f8136044d383eb9e08ee43","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"document/render.go","name":"escapeLine"},"kwargs":"sha256:b64e65a1373cba239ef192a8036647c5fd6d0b00e00c10bb4e44a63896c4b228","body":"sha256:46c8ccc652db62e5e76a651c4206f5059de6f67a03d59d9ffe5bb4ff9eb544e4"}
func escapeLine(line string) string {
	trimmed := strings.TrimLeft(line, "\\")
	if isMarkerLine(trimmed) {
		return "\\" + line
	}
	return line
}

// /sedum:setFunction
// sedum:setType {"v":2,"tier":"owned","record":"prov-2026-6eeaa2bc","gen":"sha256:85066f8dcbc8074c00cd5308255455f52b8996a0a6c81769cdbb3bdecc5740fd","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","id":{"file":"document/render.go","name":"Plan"},"kwargs":"sha256:0dedd482dbc486dbc616de8fef5df39c6a9692a85bb1173005ea337e2081cb3b","body":"sha256:889a0007606c616d86cd401e86af4db697bb03398fc0effc9479e14f54da20bb"}
type Plan struct {
	Changes   []Change
	Conflicts []Problem
	Problems  []Problem
	Unchanged []string
}

// /sedum:setType
// sedum:setType {"v":2,"tier":"owned","record":"prov-2026-b396ccea","gen":"sha256:f937af628fac653774d7d299772d396db915c0a672cae94ee63fce275d93dc72","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","id":{"file":"document/render.go","name":"Problem"},"kwargs":"sha256:59f8752aa465162d4c498cf4867f5c7015cef320c6958b0f8e8d683e1253d632","body":"sha256:9f42922bde193c170011d6836169b5fc4224bb41c94e74ef161f7dcb6a47db80"}
type Problem struct {
	Kind            string
	ID              string
	Line            int
	AnchoredVersion int
	CurrentVersion  int
	Message         string
}

// /sedum:setType
// sedum:setType {"v":2,"tier":"owned","record":"prov-2026-b98853b3","gen":"sha256:e84b3fc119d69341d66e566c7536233e244f141dee239ae94a389c2ecfff85c4","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","id":{"file":"document/render.go","name":"Section"},"kwargs":"sha256:ec48a77e7d6e14c221b826ec04a699f63a4bfde8ed0a65394a10d3e78ec2958d","body":"sha256:0233529c636b90593beb660dd7d840b762aada64acd16288916fe54826df0f05"}
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
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-f40f3a4a","gen":"sha256:926745d7d9fd808dbc039734a69f175728b8c27d2558a26dc904f0ed503a8f7c","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"document/render.go","name":"Render"},"kwargs":"sha256:d55d613101545af60fcfd7ceeb46878d05375d707f463fc8092a55bd1038f448","body":"sha256:24cb937f5193fb86b74e0dffdeffd2663c1a53f870983a66dc1b35d61cbd988a"}
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
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-f40f3a4a","gen":"sha256:926745d7d9fd808dbc039734a69f175728b8c27d2558a26dc904f0ed503a8f7c","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"document/render.go","name":"renderSection"},"kwargs":"sha256:d6e95dc69ddc88278cf3cea305bdb573e80f54068461a9e44f0eb656aea1fa74","body":"sha256:7a60dd329a2c79f9f36feaaa3aba1390177e25021f81c38d9c6f521eac9ba2b3"}
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
