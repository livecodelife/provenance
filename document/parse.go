package document

// sedum:anchor:imports
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-8ccbee91","gen":"sha256:c0f66011da8a1028a690ce2879d028f1712647a56603fa6e7649d85dc9dafa52","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"document/parse.go","path":"strconv"},"kwargs":"sha256:91895a035b7bb1663870b582ae0fe2c419d2fafd93ae385f8db4aaf072d44311","body":"sha256:f33f981c64a42d78a5daa7ff3fca15b10045080cccd282fef544b59eaf7f6dc8"}
import "strconv"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-8ccbee91","gen":"sha256:c0f66011da8a1028a690ce2879d028f1712647a56603fa6e7649d85dc9dafa52","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"document/parse.go","path":"strings"},"kwargs":"sha256:f3bcc729ea67bb4594497dc030700afeae5215181493c5fa42ec1b088eff1caf","body":"sha256:a9af3c0bc4bf9342df9d7e89a324523fdba7941a9ca4190270e68ff3c4a25ea6"}
import "strings"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-6f0b5e46","gen":"sha256:d7020c4ccb2e8b9e9d9836f4576e40ed3121fe530e41d744acc76fb3d0aee79f","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"document/parse.go","name":"Parse"},"kwargs":"sha256:11800d98ed17232e08a6d62ed4e47777d539dc422cb7341f98915540fbc87b47","body":"sha256:a759e24fe8f9ea390833f149ca7b82d6abf460335c81ae37267777a1aff9aed7"}
func Parse(text string) []Section {
	lines := strings.Split(text, "\n")
	for i, line := range lines {
		lines[i] = strings.TrimSuffix(line, "\r")
	}

	headingAt := func(line string) (int, string, bool) {
		switch {
		case strings.HasPrefix(line, "### "):
			return 3, strings.TrimSpace(line[4:]), true
		case strings.HasPrefix(line, "## "):
			return 2, strings.TrimSpace(line[3:]), true
		}
		return 0, "", false
	}

	var sections []Section
	for i := 0; i < len(lines); i++ {
		level, heading, ok := headingAt(lines[i])
		if !ok {
			continue
		}

		end := i + 1
		for end < len(lines) {
			if _, _, next := headingAt(lines[end]); next {
				break
			}
			end++
		}

		body := lines[i+1 : end]
		sec := Section{Heading: heading, Level: level, Line: i + 1}

		anchorIdx := -1
		for j, line := range body {
			if strings.TrimSpace(line) == "" {
				continue
			}
			if id, version, anchored := parseAnchor(line); anchored {
				sec.Anchored = true
				sec.ID = id
				sec.Version = version
				anchorIdx = j
			}
			break
		}

		if sec.Anchored {
			rest := body[anchorIdx+1:]

			constraintsAt := -1
			for j, line := range rest {
				if strings.TrimSpace(line) == "**Constraints**" {
					constraintsAt = j
					break
				}
			}

			intentLines := rest
			if constraintsAt >= 0 {
				intentLines = rest[:constraintsAt]
				sec.HasConstraints = true
			}

			var intent []string
			for _, line := range intentLines {
				intent = append(intent, unescapeLine(line))
			}
			sec.Intent = strings.TrimSpace(strings.Join(intent, "\n"))

			if constraintsAt >= 0 {
				var constraints []string
				for _, line := range rest[constraintsAt+1:] {
					switch {
					case strings.HasPrefix(line, "- "):
						constraints = append(constraints, line[2:])
					case strings.HasPrefix(line, "  ") && len(constraints) > 0:
						constraints[len(constraints)-1] += "\n" + line[2:]
					}
				}
				for k := range constraints {
					constraints[k] = strings.TrimSpace(constraints[k])
				}
				sec.Constraints = constraints
			}
		}

		sections = append(sections, sec)
		i = end - 1
	}
	return sections
}

// /sedum:setFunction
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-8ccbee91","gen":"sha256:c0f66011da8a1028a690ce2879d028f1712647a56603fa6e7649d85dc9dafa52","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"document/parse.go","name":"parseAnchor"},"kwargs":"sha256:a2b5c6ad13f4ba13261814f1eb101418264369f7c3d8ff6518433af66d851099","body":"sha256:a3e3d6778bcba17a6549f176b137f5be29c2ffc9f8898329bd88b2e242daf71a"}
func parseAnchor(line string) (string, int, bool) {
	s := strings.TrimSpace(line)
	const prefix = "<!-- provenance:"
	const suffix = " -->"
	if !strings.HasPrefix(s, prefix) || !strings.HasSuffix(s, suffix) || len(s) < len(prefix)+len(suffix) {
		return "", 0, false
	}
	body := s[len(prefix) : len(s)-len(suffix)]
	at := strings.LastIndex(body, "@")
	if at < 0 {
		return "", 0, false
	}
	id := body[:at]
	ver := body[at+1:]
	if len(strings.Fields(id)) != 1 || strings.Fields(id)[0] != id {
		return "", 0, false
	}
	if ver == "" || ver[0] == '0' {
		return "", 0, false
	}
	for i := 0; i < len(ver); i++ {
		if ver[i] < '0' || ver[i] > '9' {
			return "", 0, false
		}
	}
	n, err := strconv.Atoi(ver)
	if err != nil {
		return "", 0, false
	}
	return id, n, true
}

// /sedum:setFunction
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-c3d52774","gen":"sha256:eeb125dce4a509b78a0e2e90dbde4d0c1576cbf080d144b7d24cdec123f546e5","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"document/parse.go","name":"unescapeLine"},"kwargs":"sha256:8a256744ddebc2fba68dd806189cfb987482bb1f8f2963e532379d2bfd70a012","body":"sha256:2793b244364fe62708b553e04a2e84ce332ca04b77286c76e526c75d51c98a4a"}
func unescapeLine(line string) string {
	trimmed := strings.TrimLeft(line, "\\")
	if len(trimmed) == len(line) || !isMarkerLine(trimmed) {
		return line
	}
	return line[1:]
}

// /sedum:setFunction
