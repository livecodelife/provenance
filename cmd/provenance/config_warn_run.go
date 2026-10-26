package main

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-78ed3e94","gen":"sha256:e715c38ac37ef8a6e260b14fadbdb77be6baf9fa59a7a5e29915ab6a286d2fca","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/config_warn_run.go","path":"fmt"}}
import "fmt"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-78ed3e94","gen":"sha256:e715c38ac37ef8a6e260b14fadbdb77be6baf9fa59a7a5e29915ab6a286d2fca","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/config_warn_run.go","path":"io"}}
import "io"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-78ed3e94","gen":"sha256:e715c38ac37ef8a6e260b14fadbdb77be6baf9fa59a7a5e29915ab6a286d2fca","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/config_warn_run.go","path":"os"}}
import "os"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-78ed3e94","gen":"sha256:e715c38ac37ef8a6e260b14fadbdb77be6baf9fa59a7a5e29915ab6a286d2fca","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/config_warn_run.go","path":"path/filepath"}}
import "path/filepath"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-78ed3e94","gen":"sha256:e715c38ac37ef8a6e260b14fadbdb77be6baf9fa59a7a5e29915ab6a286d2fca","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/config_warn_run.go","path":"sort"}}
import "sort"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-78ed3e94","gen":"sha256:e715c38ac37ef8a6e260b14fadbdb77be6baf9fa59a7a5e29915ab6a286d2fca","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/config_warn_run.go","path":"strings"}}
import "strings"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-78ed3e94","gen":"sha256:e715c38ac37ef8a6e260b14fadbdb77be6baf9fa59a7a5e29915ab6a286d2fca","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/config_warn_run.go","path":"gopkg.in/yaml.v3"}}
import "gopkg.in/yaml.v3"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"tier":"owned","record":"prov-2026-78ed3e94","gen":"sha256:e715c38ac37ef8a6e260b14fadbdb77be6baf9fa59a7a5e29915ab6a286d2fca","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func warnNestedConfigs(r *repo, w io.Writer) {\n\tif r.ConfigPath != \"\" {\n\t\treturn\n\t}\n\tout, err := r.git(\"ls-files\", \"-z\", \"--\", \":(glob)**/.provenance.yml\", \":(glob)**/.linespec.yml\")\n\tif err != nil {\n\t\treturn\n\t}\n\tvar candidates []string\n\tfor _, p := range strings.Split(out, \"\\x00\") {\n\t\tif p == \"\" || !strings.Contains(p, \"/\") {\n\t\t\tcontinue\n\t\t}\n\t\tif filepath.Base(p) == \".linespec.yml\" {\n\t\t\tdata, err := os.ReadFile(filepath.Join(r.Root, filepath.FromSlash(p)))\n\t\t\tif err != nil {\n\t\t\t\tcontinue\n\t\t\t}\n\t\t\tvar doc map[string]any\n\t\t\tif err := yaml.Unmarshal(data, &doc); err != nil {\n\t\t\t\tcontinue\n\t\t\t}\n\t\t\tif _, ok := doc[\"provenance\"]; !ok {\n\t\t\t\tcontinue\n\t\t\t}\n\t\t}\n\t\tcandidates = append(candidates, p)\n\t}\n\tif len(candidates) == 0 {\n\t\treturn\n\t}\n\tsort.Strings(candidates)\n\tshown := candidates\n\tmore := false\n\tif len(shown) > 5 {\n\t\tshown = shown[:5]\n\t\tmore = true\n\t}\n\tlist := strings.Join(shown, \", \")\n\tif more {\n\t\tlist += \", ...\"\n\t}\n\tfmt.Fprintf(w, \"Warning: no .provenance.yml or .linespec.yml at the repository root (%s); using defaults. Found config below it: %s. Pass -c <path> to use one.\\n\", r.Root, list)\n}","file":"cmd/provenance/config_warn_run.go","name":"warnNestedConfigs"}}
func warnNestedConfigs(r *repo, w io.Writer) {
	if r.ConfigPath != "" {
		return
	}
	out, err := r.git("ls-files", "-z", "--", ":(glob)**/.provenance.yml", ":(glob)**/.linespec.yml")
	if err != nil {
		return
	}
	var candidates []string
	for _, p := range strings.Split(out, "\x00") {
		if p == "" || !strings.Contains(p, "/") {
			continue
		}
		if filepath.Base(p) == ".linespec.yml" {
			data, err := os.ReadFile(filepath.Join(r.Root, filepath.FromSlash(p)))
			if err != nil {
				continue
			}
			var doc map[string]any
			if err := yaml.Unmarshal(data, &doc); err != nil {
				continue
			}
			if _, ok := doc["provenance"]; !ok {
				continue
			}
		}
		candidates = append(candidates, p)
	}
	if len(candidates) == 0 {
		return
	}
	sort.Strings(candidates)
	shown := candidates
	more := false
	if len(shown) > 5 {
		shown = shown[:5]
		more = true
	}
	list := strings.Join(shown, ", ")
	if more {
		list += ", ..."
	}
	fmt.Fprintf(w, "Warning: no .provenance.yml or .linespec.yml at the repository root (%s); using defaults. Found config below it: %s. Pass -c <path> to use one.\n", r.Root, list)
}

// /sedum:setFunction
