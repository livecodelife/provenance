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
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-f86ede24","gen":"sha256:f5c0aa8b1ff8e113e8b95624f260f999acd8ff00d7882e63fec1278863c8b3ae","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/regen_run.go","path":"github.com/livecodelife/provenance/record"},"kwargs":"sha256:b004d8ceed0d01a217965d6da92086c1bd495157b39ad597d56ba52bbd1dd357","body":"sha256:acbf8628061291d11aa39e9ea64ecc91095a117057cfe6f1958ef2c437fac003"}
import "github.com/livecodelife/provenance/record"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-f86ede24","gen":"sha256:f5c0aa8b1ff8e113e8b95624f260f999acd8ff00d7882e63fec1278863c8b3ae","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/regen_run.go","path":"github.com/livecodelife/provenance/scope"},"kwargs":"sha256:be9d6a3ca0dae11b484e99bbb318f271135ce69249ee965aa7c764c154d129a8","body":"sha256:4ec0116f2b0d829a4ed0e1f2a40f17d465f6f981c35c8fd5427721c94a450e53"}
import "github.com/livecodelife/provenance/scope"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-0b95d9fb","gen":"sha256:b6bfe06e3d79cf3468405ce4758458ea1539c6b4ef1bb41546a090b1c9c78e9d","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/regen_run.go","name":"splitRegions"},"kwargs":"sha256:59d50199895b46e7a8028287516cfa5b0ed7079a69d9ed4bae225bf1aedd18fd","body":"sha256:21ac77abf6f67106cc678feafd7dc4598e3205abb6a352085cfdeba1f473da48"}
func splitRegions(content string) (string, []regionMarker, bool) {
	const placeholder = "// sedum-region"
	isWS := func(c byte) bool { return c == ' ' || c == '\t' || c == '\r' }
	isIndent := func(c byte) bool { return c == ' ' || c == '\t' }
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
		i := 0
		for i < len(line) && isIndent(line[i]) {
			i++
		}
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
		pos := i + len(leader)
		const prefix = " sedum:"
		rest = line[pos:]
		if len(rest) < len(prefix) || rest[:len(prefix)] != prefix {
			return "", "", regionMarker{}, false
		}
		pos += len(prefix)
		rest = line[pos:]
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
		jsonStart := pos + sp + 1
		if jsonStart >= len(line) || line[jsonStart] != '{' {
			return "", "", regionMarker{}, false
		}
		strs, ok := parseJSON(trimRight(line[jsonStart:]))
		if !ok {
			return "", "", regionMarker{}, false
		}
		tier, ok1 := strs["tier"]
		record, ok2 := strs["record"]
		gen, ok3 := strs["gen"]
		pkg, ok4 := strs["pkg"]
		if !ok1 || !ok2 || !ok3 || !ok4 {
			return "", "", regionMarker{}, false
		}
		return leader, action, regionMarker{Tier: tier, Record: record, Gen: gen, Pkg: pkg, Prefix: line[:jsonStart]}, true
	}
	closes := func(line, leader, action string) bool {
		i := 0
		for i < len(line) && isIndent(line[i]) {
			i++
		}
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
	var body string
	for _, line := range lines {
		if !inRegion {
			l, a, m, ok := parseOpen(line)
			if ok {
				inRegion = true
				leader, action, marker = l, a, m
				body = ""
				out = append(out, placeholder)
				continue
			}
			out = append(out, line)
			continue
		}
		if closes(line, leader, action) {
			inRegion = false
			marker.Body = body
			marker.Close = line
			markers = append(markers, marker)
			continue
		}
		if _, _, _, ok := parseOpen(line); ok {
			return "", nil, false
		}
		body += line + "\n"
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
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-f86ede24","gen":"sha256:f5c0aa8b1ff8e113e8b95624f260f999acd8ff00d7882e63fec1278863c8b3ae","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/regen_run.go","name":"regenerationOnly"},"kwargs":"sha256:9b1712ecb17006efc417d5cbbacc47fb4dc13cc0e978e6001a1c2b46551e53f5","body":"sha256:b6d5714bb48fb803df703345f3fa34f4171283b45f37b04b2a4e1a7fc230b222"}
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
		cmd := exec.Command("git", "cat-file", "blob", base+":./"+p)
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
// sedum:setType {"v":2,"tier":"owned","record":"prov-2026-fd7c2130","gen":"sha256:50963fc8d4febe2bd1c6a0155e367ccba71e32cde64fde46f55061a92748333d","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","id":{"file":"cmd/provenance/regen_run.go","name":"regionMarker"},"kwargs":"sha256:dc34bd27761192daaea53290e1d7887fe7ef78bd842dd9822c030af18c4142d4","body":"sha256:a24fc94016203403198ee9699d1265fccc04584c233701aa5ae3459fa0a5d1d6"}
type regionMarker struct {
	Tier   string
	Record string
	Gen    string
	Pkg    string
	Prefix string
	Body   string
	Close  string
}

// /sedum:setType
// sedum:setFunction {"tier":"owned","record":"prov-2026-b171a7da","gen":"sha256:4539bc13c3a0213659c5787d4058bcde1a870f7b05345b0242c24bf12ae5fa40","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func sedumMissing() bool {\n\tprog := os.Getenv(\"SEDUM_BIN\")\n\tif prog == \"\" {\n\t\tprog = \"sedum\"\n\t}\n\t_, err := exec.LookPath(prog)\n\treturn err != nil\n}","file":"cmd/provenance/regen_run.go","name":"sedumMissing"}}
func sedumMissing() bool {
	prog := os.Getenv("SEDUM_BIN")
	if prog == "" {
		prog = "sedum"
	}
	_, err := exec.LookPath(prog)
	return err != nil
}

// /sedum:setFunction
