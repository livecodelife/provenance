package document

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-8ccbee91","gen":"sha256:c0f66011da8a1028a690ce2879d028f1712647a56603fa6e7649d85dc9dafa52","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"document/parse.go","path":"strconv"}}
import "strconv"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-8ccbee91","gen":"sha256:c0f66011da8a1028a690ce2879d028f1712647a56603fa6e7649d85dc9dafa52","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"document/parse.go","path":"strings"}}
import "strings"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"tier":"owned","record":"prov-2026-6f0b5e46","gen":"sha256:d7020c4ccb2e8b9e9d9836f4576e40ed3121fe530e41d744acc76fb3d0aee79f","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func Parse(text string) []Section {\n\tlines := strings.Split(text, \"\\n\")\n\tfor i, line := range lines {\n\t\tlines[i] = strings.TrimSuffix(line, \"\\r\")\n\t}\n\n\theadingAt := func(line string) (int, string, bool) {\n\t\tswitch {\n\t\tcase strings.HasPrefix(line, \"### \"):\n\t\t\treturn 3, strings.TrimSpace(line[4:]), true\n\t\tcase strings.HasPrefix(line, \"## \"):\n\t\t\treturn 2, strings.TrimSpace(line[3:]), true\n\t\t}\n\t\treturn 0, \"\", false\n\t}\n\n\tvar sections []Section\n\tfor i := 0; i < len(lines); i++ {\n\t\tlevel, heading, ok := headingAt(lines[i])\n\t\tif !ok {\n\t\t\tcontinue\n\t\t}\n\n\t\tend := i + 1\n\t\tfor end < len(lines) {\n\t\t\tif _, _, next := headingAt(lines[end]); next {\n\t\t\t\tbreak\n\t\t\t}\n\t\t\tend++\n\t\t}\n\n\t\tbody := lines[i+1 : end]\n\t\tsec := Section{Heading: heading, Level: level, Line: i + 1}\n\n\t\tanchorIdx := -1\n\t\tfor j, line := range body {\n\t\t\tif strings.TrimSpace(line) == \"\" {\n\t\t\t\tcontinue\n\t\t\t}\n\t\t\tif id, version, anchored := parseAnchor(line); anchored {\n\t\t\t\tsec.Anchored = true\n\t\t\t\tsec.ID = id\n\t\t\t\tsec.Version = version\n\t\t\t\tanchorIdx = j\n\t\t\t}\n\t\t\tbreak\n\t\t}\n\n\t\tif sec.Anchored {\n\t\t\trest := body[anchorIdx+1:]\n\n\t\t\tconstraintsAt := -1\n\t\t\tfor j, line := range rest {\n\t\t\t\tif strings.TrimSpace(line) == \"**Constraints**\" {\n\t\t\t\t\tconstraintsAt = j\n\t\t\t\t\tbreak\n\t\t\t\t}\n\t\t\t}\n\n\t\t\tintentLines := rest\n\t\t\tif constraintsAt >= 0 {\n\t\t\t\tintentLines = rest[:constraintsAt]\n\t\t\t\tsec.HasConstraints = true\n\t\t\t}\n\n\t\t\tvar intent []string\n\t\t\tfor _, line := range intentLines {\n\t\t\t\tintent = append(intent, unescapeLine(line))\n\t\t\t}\n\t\t\tsec.Intent = strings.TrimSpace(strings.Join(intent, \"\\n\"))\n\n\t\t\tif constraintsAt >= 0 {\n\t\t\t\tvar constraints []string\n\t\t\t\tfor _, line := range rest[constraintsAt+1:] {\n\t\t\t\t\tswitch {\n\t\t\t\t\tcase strings.HasPrefix(line, \"- \"):\n\t\t\t\t\t\tconstraints = append(constraints, line[2:])\n\t\t\t\t\tcase strings.HasPrefix(line, \"  \") && len(constraints) > 0:\n\t\t\t\t\t\tconstraints[len(constraints)-1] += \"\\n\" + line[2:]\n\t\t\t\t\t}\n\t\t\t\t}\n\t\t\t\tfor k := range constraints {\n\t\t\t\t\tconstraints[k] = strings.TrimSpace(constraints[k])\n\t\t\t\t}\n\t\t\t\tsec.Constraints = constraints\n\t\t\t}\n\t\t}\n\n\t\tsections = append(sections, sec)\n\t\ti = end - 1\n\t}\n\treturn sections\n}","file":"document/parse.go","name":"Parse"}}
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
// sedum:setFunction {"tier":"owned","record":"prov-2026-8ccbee91","gen":"sha256:c0f66011da8a1028a690ce2879d028f1712647a56603fa6e7649d85dc9dafa52","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func parseAnchor(line string) (string, int, bool) {\n\ts := strings.TrimSpace(line)\n\tconst prefix = \"<!-- provenance:\"\n\tconst suffix = \" -->\"\n\tif !strings.HasPrefix(s, prefix) || !strings.HasSuffix(s, suffix) || len(s) < len(prefix)+len(suffix) {\n\t\treturn \"\", 0, false\n\t}\n\tbody := s[len(prefix) : len(s)-len(suffix)]\n\tat := strings.LastIndex(body, \"@\")\n\tif at < 0 {\n\t\treturn \"\", 0, false\n\t}\n\tid := body[:at]\n\tver := body[at+1:]\n\tif len(strings.Fields(id)) != 1 || strings.Fields(id)[0] != id {\n\t\treturn \"\", 0, false\n\t}\n\tif ver == \"\" || ver[0] == '0' {\n\t\treturn \"\", 0, false\n\t}\n\tfor i := 0; i < len(ver); i++ {\n\t\tif ver[i] < '0' || ver[i] > '9' {\n\t\t\treturn \"\", 0, false\n\t\t}\n\t}\n\tn, err := strconv.Atoi(ver)\n\tif err != nil {\n\t\treturn \"\", 0, false\n\t}\n\treturn id, n, true\n}\n","file":"document/parse.go","name":"parseAnchor"}}
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
// sedum:setFunction {"tier":"owned","record":"prov-2026-c3d52774","gen":"sha256:eeb125dce4a509b78a0e2e90dbde4d0c1576cbf080d144b7d24cdec123f546e5","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func unescapeLine(line string) string {\n\ttrimmed := strings.TrimLeft(line, \"\\\\\")\n\tif len(trimmed) == len(line) || !isMarkerLine(trimmed) {\n\t\treturn line\n\t}\n\treturn line[1:]\n}","file":"document/parse.go","name":"unescapeLine"}}
func unescapeLine(line string) string {
	trimmed := strings.TrimLeft(line, "\\")
	if len(trimmed) == len(line) || !isMarkerLine(trimmed) {
		return line
	}
	return line[1:]
}

// /sedum:setFunction
