package main

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-351c5684","gen":"sha256:334771655a864adc9e05113854894ef002e9f5f243ffc85e5c3eba0d7439d6d4","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/regen_run.go","path":"os"}}
import "os"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-351c5684","gen":"sha256:334771655a864adc9e05113854894ef002e9f5f243ffc85e5c3eba0d7439d6d4","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/regen_run.go","path":"os/exec"}}
import "os/exec"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-351c5684","gen":"sha256:334771655a864adc9e05113854894ef002e9f5f243ffc85e5c3eba0d7439d6d4","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/regen_run.go","path":"path/filepath"}}
import "path/filepath"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-351c5684","gen":"sha256:334771655a864adc9e05113854894ef002e9f5f243ffc85e5c3eba0d7439d6d4","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/regen_run.go","path":"strings"}}
import "strings"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-f86ede24","gen":"sha256:8b2b72b5afdc5184b37104fe49da95224d1f2b147532d0e8d071f58a4c0d60d9","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/regen_run.go","path":"github.com/livecodelife/provenance/record"}}
import "github.com/livecodelife/provenance/record"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-f86ede24","gen":"sha256:8b2b72b5afdc5184b37104fe49da95224d1f2b147532d0e8d071f58a4c0d60d9","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/regen_run.go","path":"github.com/livecodelife/provenance/scope"}}
import "github.com/livecodelife/provenance/scope"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"tier":"owned","record":"prov-2026-0b95d9fb","gen":"sha256:ef15cc5263f0776c1be6009c6c55fb596e7fc5aedbe497d464b2a7b1011aa054","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func splitRegions(content string) (string, []regionMarker, bool) {\n\tconst placeholder = \"// sedum-region\"\n\tisWS := func(c byte) bool { return c == ' ' || c == '\\t' || c == '\\r' }\n\ttrimRight := func(s string) string {\n\t\tend := len(s)\n\t\tfor end > 0 && isWS(s[end-1]) {\n\t\t\tend--\n\t\t}\n\t\treturn s[:end]\n\t}\n\tskipWS := func(s string, pos int) int {\n\t\tfor pos < len(s) && isWS(s[pos]) {\n\t\t\tpos++\n\t\t}\n\t\treturn pos\n\t}\n\tparseString := func(s string, pos int) (string, int, bool) {\n\t\tif pos >= len(s) || s[pos] != '\"' {\n\t\t\treturn \"\", pos, false\n\t\t}\n\t\tvar buf []byte\n\t\ti := pos + 1\n\t\tfor i < len(s) {\n\t\t\tc := s[i]\n\t\t\tif c == '\"' {\n\t\t\t\treturn string(buf), i + 1, true\n\t\t\t}\n\t\t\tif c == '\\\\' {\n\t\t\t\tif i+1 >= len(s) {\n\t\t\t\t\treturn \"\", pos, false\n\t\t\t\t}\n\t\t\t\te := s[i+1]\n\t\t\t\tswitch e {\n\t\t\t\tcase 'n':\n\t\t\t\t\tbuf = append(buf, '\\n')\n\t\t\t\tcase 't':\n\t\t\t\t\tbuf = append(buf, '\\t')\n\t\t\t\tcase 'r':\n\t\t\t\t\tbuf = append(buf, '\\r')\n\t\t\t\tcase 'b', 'f':\n\t\t\t\tcase 'u':\n\t\t\t\t\tif i+5 >= len(s) {\n\t\t\t\t\t\treturn \"\", pos, false\n\t\t\t\t\t}\n\t\t\t\t\ti += 4\n\t\t\t\tdefault:\n\t\t\t\t\tbuf = append(buf, e)\n\t\t\t\t}\n\t\t\t\ti += 2\n\t\t\t\tcontinue\n\t\t\t}\n\t\t\tbuf = append(buf, c)\n\t\t\ti++\n\t\t}\n\t\treturn \"\", pos, false\n\t}\n\tskipValue := func(s string, pos int) (int, bool) {\n\t\tdepth := 0\n\t\ti := pos\n\t\tfor i < len(s) {\n\t\t\tc := s[i]\n\t\t\tif c == '\"' {\n\t\t\t\t_, j, ok := parseString(s, i)\n\t\t\t\tif !ok {\n\t\t\t\t\treturn pos, false\n\t\t\t\t}\n\t\t\t\ti = j\n\t\t\t\tcontinue\n\t\t\t}\n\t\t\tif c == '{' || c == '[' {\n\t\t\t\tdepth++\n\t\t\t} else if c == '}' || c == ']' {\n\t\t\t\tif depth == 0 {\n\t\t\t\t\tbreak\n\t\t\t\t}\n\t\t\t\tdepth--\n\t\t\t} else if c == ',' && depth == 0 {\n\t\t\t\tbreak\n\t\t\t}\n\t\t\ti++\n\t\t}\n\t\tif i == pos || depth != 0 {\n\t\t\treturn pos, false\n\t\t}\n\t\treturn i, true\n\t}\n\tparseJSON := func(s string) (map[string]string, bool) {\n\t\tstrs := map[string]string{}\n\t\ti := skipWS(s, 0)\n\t\tif i >= len(s) || s[i] != '{' {\n\t\t\treturn nil, false\n\t\t}\n\t\ti = skipWS(s, i+1)\n\t\tif i < len(s) && s[i] == '}' {\n\t\t\ti++\n\t\t} else {\n\t\t\tfor {\n\t\t\t\tkey, j, ok := parseString(s, i)\n\t\t\t\tif !ok {\n\t\t\t\t\treturn nil, false\n\t\t\t\t}\n\t\t\t\ti = skipWS(s, j)\n\t\t\t\tif i >= len(s) || s[i] != ':' {\n\t\t\t\t\treturn nil, false\n\t\t\t\t}\n\t\t\t\ti = skipWS(s, i+1)\n\t\t\t\tif i < len(s) && s[i] == '\"' {\n\t\t\t\t\tv, j2, ok2 := parseString(s, i)\n\t\t\t\t\tif !ok2 {\n\t\t\t\t\t\treturn nil, false\n\t\t\t\t\t}\n\t\t\t\t\tstrs[key] = v\n\t\t\t\t\ti = j2\n\t\t\t\t} else {\n\t\t\t\t\tj2, ok2 := skipValue(s, i)\n\t\t\t\t\tif !ok2 {\n\t\t\t\t\t\treturn nil, false\n\t\t\t\t\t}\n\t\t\t\t\ti = j2\n\t\t\t\t}\n\t\t\t\ti = skipWS(s, i)\n\t\t\t\tif i < len(s) && s[i] == ',' {\n\t\t\t\t\ti = skipWS(s, i+1)\n\t\t\t\t\tcontinue\n\t\t\t\t}\n\t\t\t\tif i < len(s) && s[i] == '}' {\n\t\t\t\t\ti++\n\t\t\t\t\tbreak\n\t\t\t\t}\n\t\t\t\treturn nil, false\n\t\t\t}\n\t\t}\n\t\tif skipWS(s, i) != len(s) {\n\t\t\treturn nil, false\n\t\t}\n\t\treturn strs, true\n\t}\n\tparseOpen := func(line string) (string, string, regionMarker, bool) {\n\t\ti := skipWS(line, 0)\n\t\trest := line[i:]\n\t\tvar leader string\n\t\tswitch {\n\t\tcase len(rest) >= 2 && rest[0] == '/' && rest[1] == '/':\n\t\t\tleader = \"//\"\n\t\tcase len(rest) >= 1 && rest[0] == '#':\n\t\t\tleader = \"#\"\n\t\tdefault:\n\t\t\treturn \"\", \"\", regionMarker{}, false\n\t\t}\n\t\trest = rest[len(leader):]\n\t\tconst prefix = \" sedum:\"\n\t\tif len(rest) < len(prefix) || rest[:len(prefix)] != prefix {\n\t\t\treturn \"\", \"\", regionMarker{}, false\n\t\t}\n\t\trest = rest[len(prefix):]\n\t\tsp := -1\n\t\tfor k := 0; k < len(rest); k++ {\n\t\t\tif rest[k] == ' ' {\n\t\t\t\tsp = k\n\t\t\t\tbreak\n\t\t\t}\n\t\t}\n\t\tif sp <= 0 {\n\t\t\treturn \"\", \"\", regionMarker{}, false\n\t\t}\n\t\taction := rest[:sp]\n\t\tstrs, ok := parseJSON(trimRight(rest[sp+1:]))\n\t\tif !ok {\n\t\t\treturn \"\", \"\", regionMarker{}, false\n\t\t}\n\t\ttier, ok1 := strs[\"tier\"]\n\t\trecord, ok2 := strs[\"record\"]\n\t\tgen, ok3 := strs[\"gen\"]\n\t\t_, ok4 := strs[\"pkg\"]\n\t\tif !ok1 || !ok2 || !ok3 || !ok4 {\n\t\t\treturn \"\", \"\", regionMarker{}, false\n\t\t}\n\t\treturn leader, action, regionMarker{Tier: tier, Record: record, Gen: gen}, true\n\t}\n\tcloses := func(line, leader, action string) bool {\n\t\ti := skipWS(line, 0)\n\t\trest := line[i:]\n\t\tif len(rest) < len(leader)+1 || rest[:len(leader)] != leader || rest[len(leader)] != ' ' {\n\t\t\treturn false\n\t\t}\n\t\trest = rest[len(leader)+1:]\n\t\twant := \"/sedum:\" + action\n\t\tif len(rest) < len(want) || rest[:len(want)] != want {\n\t\t\treturn false\n\t\t}\n\t\treturn trimRight(rest[len(want):]) == \"\"\n\t}\n\tvar lines []string\n\tstart := 0\n\tfor i := 0; i <= len(content); i++ {\n\t\tif i == len(content) || content[i] == '\\n' {\n\t\t\tlines = append(lines, content[start:i])\n\t\t\tstart = i + 1\n\t\t}\n\t}\n\tvar out []string\n\tmarkers := []regionMarker{}\n\tinRegion := false\n\tvar leader, action string\n\tvar marker regionMarker\n\tfor _, line := range lines {\n\t\tif !inRegion {\n\t\t\tl, a, m, ok := parseOpen(line)\n\t\t\tif ok {\n\t\t\t\tinRegion = true\n\t\t\t\tleader, action, marker = l, a, m\n\t\t\t\tout = append(out, placeholder)\n\t\t\t\tcontinue\n\t\t\t}\n\t\t\tout = append(out, line)\n\t\t\tcontinue\n\t\t}\n\t\tif closes(line, leader, action) {\n\t\t\tinRegion = false\n\t\t\tmarkers = append(markers, marker)\n\t\t\tcontinue\n\t\t}\n\t\tif _, _, _, ok := parseOpen(line); ok {\n\t\t\treturn \"\", nil, false\n\t\t}\n\t}\n\tif inRegion {\n\t\treturn \"\", nil, false\n\t}\n\tvar buf []byte\n\tfor i, l := range out {\n\t\tif i > 0 {\n\t\t\tbuf = append(buf, '\\n')\n\t\t}\n\t\tbuf = append(buf, l...)\n\t}\n\treturn string(buf), markers, true\n}","file":"cmd/provenance/regen_run.go","name":"splitRegions"}}
func splitRegions(content string) (string, []regionMarker, bool) {
	const placeholder = "// sedum-region"
	isWS := func(c byte) bool { return c == ' ' || c == '\t' || c == '\r' }
	trimRight := func(s string) string {
		end := len(s)
		for end > 0 && isWS(s[end-1]) {
			end--
		}
		return s[:end]
	}
	skipWS := func(s string, pos int) int {
		for pos < len(s) && isWS(s[pos]) {
			pos++
		}
		return pos
	}
	parseString := func(s string, pos int) (string, int, bool) {
		if pos >= len(s) || s[pos] != '"' {
			return "", pos, false
		}
		var buf []byte
		i := pos + 1
		for i < len(s) {
			c := s[i]
			if c == '"' {
				return string(buf), i + 1, true
			}
			if c == '\\' {
				if i+1 >= len(s) {
					return "", pos, false
				}
				e := s[i+1]
				switch e {
				case 'n':
					buf = append(buf, '\n')
				case 't':
					buf = append(buf, '\t')
				case 'r':
					buf = append(buf, '\r')
				case 'b', 'f':
				case 'u':
					if i+5 >= len(s) {
						return "", pos, false
					}
					i += 4
				default:
					buf = append(buf, e)
				}
				i += 2
				continue
			}
			buf = append(buf, c)
			i++
		}
		return "", pos, false
	}
	skipValue := func(s string, pos int) (int, bool) {
		depth := 0
		i := pos
		for i < len(s) {
			c := s[i]
			if c == '"' {
				_, j, ok := parseString(s, i)
				if !ok {
					return pos, false
				}
				i = j
				continue
			}
			if c == '{' || c == '[' {
				depth++
			} else if c == '}' || c == ']' {
				if depth == 0 {
					break
				}
				depth--
			} else if c == ',' && depth == 0 {
				break
			}
			i++
		}
		if i == pos || depth != 0 {
			return pos, false
		}
		return i, true
	}
	parseJSON := func(s string) (map[string]string, bool) {
		strs := map[string]string{}
		i := skipWS(s, 0)
		if i >= len(s) || s[i] != '{' {
			return nil, false
		}
		i = skipWS(s, i+1)
		if i < len(s) && s[i] == '}' {
			i++
		} else {
			for {
				key, j, ok := parseString(s, i)
				if !ok {
					return nil, false
				}
				i = skipWS(s, j)
				if i >= len(s) || s[i] != ':' {
					return nil, false
				}
				i = skipWS(s, i+1)
				if i < len(s) && s[i] == '"' {
					v, j2, ok2 := parseString(s, i)
					if !ok2 {
						return nil, false
					}
					strs[key] = v
					i = j2
				} else {
					j2, ok2 := skipValue(s, i)
					if !ok2 {
						return nil, false
					}
					i = j2
				}
				i = skipWS(s, i)
				if i < len(s) && s[i] == ',' {
					i = skipWS(s, i+1)
					continue
				}
				if i < len(s) && s[i] == '}' {
					i++
					break
				}
				return nil, false
			}
		}
		if skipWS(s, i) != len(s) {
			return nil, false
		}
		return strs, true
	}
	parseOpen := func(line string) (string, string, regionMarker, bool) {
		i := skipWS(line, 0)
		rest := line[i:]
		var leader string
		switch {
		case len(rest) >= 2 && rest[0] == '/' && rest[1] == '/':
			leader = "//"
		case len(rest) >= 1 && rest[0] == '#':
			leader = "#"
		default:
			return "", "", regionMarker{}, false
		}
		rest = rest[len(leader):]
		const prefix = " sedum:"
		if len(rest) < len(prefix) || rest[:len(prefix)] != prefix {
			return "", "", regionMarker{}, false
		}
		rest = rest[len(prefix):]
		sp := -1
		for k := 0; k < len(rest); k++ {
			if rest[k] == ' ' {
				sp = k
				break
			}
		}
		if sp <= 0 {
			return "", "", regionMarker{}, false
		}
		action := rest[:sp]
		strs, ok := parseJSON(trimRight(rest[sp+1:]))
		if !ok {
			return "", "", regionMarker{}, false
		}
		tier, ok1 := strs["tier"]
		record, ok2 := strs["record"]
		gen, ok3 := strs["gen"]
		_, ok4 := strs["pkg"]
		if !ok1 || !ok2 || !ok3 || !ok4 {
			return "", "", regionMarker{}, false
		}
		return leader, action, regionMarker{Tier: tier, Record: record, Gen: gen}, true
	}
	closes := func(line, leader, action string) bool {
		i := skipWS(line, 0)
		rest := line[i:]
		if len(rest) < len(leader)+1 || rest[:len(leader)] != leader || rest[len(leader)] != ' ' {
			return false
		}
		rest = rest[len(leader)+1:]
		want := "/sedum:" + action
		if len(rest) < len(want) || rest[:len(want)] != want {
			return false
		}
		return trimRight(rest[len(want):]) == ""
	}
	var lines []string
	start := 0
	for i := 0; i <= len(content); i++ {
		if i == len(content) || content[i] == '\n' {
			lines = append(lines, content[start:i])
			start = i + 1
		}
	}
	var out []string
	markers := []regionMarker{}
	inRegion := false
	var leader, action string
	var marker regionMarker
	for _, line := range lines {
		if !inRegion {
			l, a, m, ok := parseOpen(line)
			if ok {
				inRegion = true
				leader, action, marker = l, a, m
				out = append(out, placeholder)
				continue
			}
			out = append(out, line)
			continue
		}
		if closes(line, leader, action) {
			inRegion = false
			markers = append(markers, marker)
			continue
		}
		if _, _, _, ok := parseOpen(line); ok {
			return "", nil, false
		}
	}
	if inRegion {
		return "", nil, false
	}
	var buf []byte
	for i, l := range out {
		if i > 0 {
			buf = append(buf, '\n')
		}
		buf = append(buf, l...)
	}
	return string(buf), markers, true
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-342fc2b3","gen":"sha256:e92503a1d3d6b29081e84e521ec493532b51606d714f6db62c68849449bd827d","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func regionsAgree(base, staged string) bool {\n\tbaseOutside, baseMarkers, ok := splitRegions(base)\n\tif !ok {\n\t\treturn false\n\t}\n\tstagedOutside, stagedMarkers, ok := splitRegions(staged)\n\tif !ok {\n\t\treturn false\n\t}\n\tif len(baseMarkers) == 0 {\n\t\treturn false\n\t}\n\tif baseOutside != stagedOutside {\n\t\treturn false\n\t}\n\tif len(baseMarkers) != len(stagedMarkers) {\n\t\treturn false\n\t}\n\tfor i := range baseMarkers {\n\t\ta, b := baseMarkers[i], stagedMarkers[i]\n\t\tif a.Tier != b.Tier || a.Record != b.Record || a.Gen != b.Gen {\n\t\t\treturn false\n\t\t}\n\t}\n\treturn true\n}","file":"cmd/provenance/regen_run.go","name":"regionsAgree"}}
func regionsAgree(base, staged string) bool {
	baseOutside, baseMarkers, ok := splitRegions(base)
	if !ok {
		return false
	}
	stagedOutside, stagedMarkers, ok := splitRegions(staged)
	if !ok {
		return false
	}
	if len(baseMarkers) == 0 {
		return false
	}
	if baseOutside != stagedOutside {
		return false
	}
	if len(baseMarkers) != len(stagedMarkers) {
		return false
	}
	for i := range baseMarkers {
		a, b := baseMarkers[i], stagedMarkers[i]
		if a.Tier != b.Tier || a.Record != b.Record || a.Gen != b.Gen {
			return false
		}
	}
	return true
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-351c5684","gen":"sha256:334771655a864adc9e05113854894ef002e9f5f243ffc85e5c3eba0d7439d6d4","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func sedumCurrent(r *repo, treePath string, recordID string) bool {\n\tgens := filepath.Join(treePath, \"generators\")\n\tif info, err := os.Stat(gens); err != nil || !info.IsDir() {\n\t\ttmp, err := os.MkdirTemp(\"\", \"sedum-empty-\")\n\t\tif err != nil {\n\t\t\treturn false\n\t\t}\n\t\tdefer os.RemoveAll(tmp)\n\t\tgens = tmp\n\t}\n\n\tbin := os.Getenv(\"SEDUM_BIN\")\n\tif bin == \"\" {\n\t\tbin = \"sedum\"\n\t}\n\n\tcmd := exec.Command(bin, \"grow\", \"--check\", \"--generators\", gens, \"--records\", filepath.Join(treePath, r.Config.Dir), \"--output\", treePath, \"--only\", recordID)\n\tcmd.Dir = treePath\n\tout, err := cmd.Output()\n\tif err != nil {\n\t\treturn false\n\t}\n\n\tfor _, line := range strings.Split(string(out), \"\\n\") {\n\t\tif !strings.HasPrefix(line, \"current\") {\n\t\t\tcontinue\n\t\t}\n\t\trest := line[len(\"current\"):]\n\t\tif rest == \"\" || (rest[0] != ' ' && rest[0] != '\\t') {\n\t\t\tcontinue\n\t\t}\n\t\trest = strings.TrimLeft(rest, \" \\t\")\n\t\tif !strings.HasPrefix(rest, recordID) {\n\t\t\tcontinue\n\t\t}\n\t\tafter := rest[len(recordID):]\n\t\tif after == \"\" || after[0] == ' ' || after[0] == '\\t' || after[0] == '\\r' {\n\t\t\treturn true\n\t\t}\n\t}\n\treturn false\n}\n","file":"cmd/provenance/regen_run.go","name":"sedumCurrent"}}
func sedumCurrent(r *repo, treePath string, recordID string) bool {
	gens := filepath.Join(treePath, "generators")
	if info, err := os.Stat(gens); err != nil || !info.IsDir() {
		tmp, err := os.MkdirTemp("", "sedum-empty-")
		if err != nil {
			return false
		}
		defer os.RemoveAll(tmp)
		gens = tmp
	}

	bin := os.Getenv("SEDUM_BIN")
	if bin == "" {
		bin = "sedum"
	}

	cmd := exec.Command(bin, "grow", "--check", "--generators", gens, "--records", filepath.Join(treePath, r.Config.Dir), "--output", treePath, "--only", recordID)
	cmd.Dir = treePath
	out, err := cmd.Output()
	if err != nil {
		return false
	}

	for _, line := range strings.Split(string(out), "\n") {
		if !strings.HasPrefix(line, "current") {
			continue
		}
		rest := line[len("current"):]
		if rest == "" || (rest[0] != ' ' && rest[0] != '\t') {
			continue
		}
		rest = strings.TrimLeft(rest, " \t")
		if !strings.HasPrefix(rest, recordID) {
			continue
		}
		after := rest[len(recordID):]
		if after == "" || after[0] == ' ' || after[0] == '\t' || after[0] == '\r' {
			return true
		}
	}
	return false
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-f86ede24","gen":"sha256:8b2b72b5afdc5184b37104fe49da95224d1f2b147532d0e8d071f58a4c0d60d9","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func regenerationOnly(r *repo, base string, treePath string, changed []string, _ *record.Record) bool {\n\tvar governed []string\n\tfor _, p := range changed {\n\t\tif recordFileID(r, p) != \"\" {\n\t\t\treturn false\n\t\t}\n\t\tif stateExempt(r, p) {\n\t\t\treturn false\n\t\t}\n\t\tif scope.Excluded(p, r.Config.ExcludePaths) {\n\t\t\tcontinue\n\t\t}\n\t\tgoverned = append(governed, p)\n\t}\n\tif len(governed) == 0 {\n\t\treturn false\n\t}\n\n\tvar stagedAll []string\n\tfor _, p := range governed {\n\t\tif !headHas(r, base, p) {\n\t\t\treturn false\n\t\t}\n\t\tstaged := filepath.Join(treePath, filepath.FromSlash(p))\n\t\tdata, err := os.ReadFile(staged)\n\t\tif err != nil {\n\t\t\treturn false\n\t\t}\n\t\tcmd := exec.Command(\"git\", \"cat-file\", \"blob\", base+\":\"+p)\n\t\tcmd.Dir = r.Root\n\t\tcmd.Env = r.gitEnv()\n\t\tbaseOut, err := cmd.Output()\n\t\tif err != nil {\n\t\t\treturn false\n\t\t}\n\t\tif !regionsAgree(string(baseOut), string(data)) {\n\t\t\treturn false\n\t\t}\n\t\tstagedAll = append(stagedAll, string(data))\n\t}\n\n\tseen := map[string]bool{}\n\tvar ids []string\n\tfor _, content := range stagedAll {\n\t\t_, markers, _ := splitRegions(content)\n\t\tfor _, m := range markers {\n\t\t\tif !seen[m.Record] {\n\t\t\t\tseen[m.Record] = true\n\t\t\t\tids = append(ids, m.Record)\n\t\t\t}\n\t\t}\n\t}\n\tfor _, id := range ids {\n\t\tif !sedumCurrent(r, treePath, id) {\n\t\t\treturn false\n\t\t}\n\t}\n\treturn true\n}","file":"cmd/provenance/regen_run.go","name":"regenerationOnly"}}
func regenerationOnly(r *repo, base string, treePath string, changed []string, _ *record.Record) bool {
	var governed []string
	for _, p := range changed {
		if recordFileID(r, p) != "" {
			return false
		}
		if stateExempt(r, p) {
			return false
		}
		if scope.Excluded(p, r.Config.ExcludePaths) {
			continue
		}
		governed = append(governed, p)
	}
	if len(governed) == 0 {
		return false
	}

	var stagedAll []string
	for _, p := range governed {
		if !headHas(r, base, p) {
			return false
		}
		staged := filepath.Join(treePath, filepath.FromSlash(p))
		data, err := os.ReadFile(staged)
		if err != nil {
			return false
		}
		cmd := exec.Command("git", "cat-file", "blob", base+":"+p)
		cmd.Dir = r.Root
		cmd.Env = r.gitEnv()
		baseOut, err := cmd.Output()
		if err != nil {
			return false
		}
		if !regionsAgree(string(baseOut), string(data)) {
			return false
		}
		stagedAll = append(stagedAll, string(data))
	}

	seen := map[string]bool{}
	var ids []string
	for _, content := range stagedAll {
		_, markers, _ := splitRegions(content)
		for _, m := range markers {
			if !seen[m.Record] {
				seen[m.Record] = true
				ids = append(ids, m.Record)
			}
		}
	}
	for _, id := range ids {
		if !sedumCurrent(r, treePath, id) {
			return false
		}
	}
	return true
}

// /sedum:setFunction
// sedum:setType {"tier":"owned","record":"prov-2026-fd7c2130","gen":"sha256:968c4936ca8799bdfe2ed91167ca8f072bc08a814dc5b1fe81bc24cb5e82140f","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type regionMarker struct {\n\tTier   string\n\tRecord string\n\tGen    string\n}","file":"cmd/provenance/regen_run.go","name":"regionMarker"}}
type regionMarker struct {
	Tier   string
	Record string
	Gen    string
}

// /sedum:setType
