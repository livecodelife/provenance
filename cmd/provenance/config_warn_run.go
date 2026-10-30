package main

// sedum:anchor:imports
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-78ed3e94","gen":"sha256:e715c38ac37ef8a6e260b14fadbdb77be6baf9fa59a7a5e29915ab6a286d2fca","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/config_warn_run.go","path":"fmt"},"kwargs":"sha256:3af267300132df785cd24a2e492cff8de17010ec4b0e09b98cc6c84913792ead","body":"sha256:bcbac51883cbf8064e9196b371cb880cee429866652447ad557533d817857609"}
import "fmt"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-78ed3e94","gen":"sha256:e715c38ac37ef8a6e260b14fadbdb77be6baf9fa59a7a5e29915ab6a286d2fca","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/config_warn_run.go","path":"io"},"kwargs":"sha256:dfdfba33f21d2907e8782cf7d80279832d79a74382d977f4f2dee7fae8e3e7c7","body":"sha256:ef8e8abf4ef5ee40a19dc85b49bf09c5df1fe32abfe834e0ab4c2175baf80ddf"}
import "io"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-78ed3e94","gen":"sha256:e715c38ac37ef8a6e260b14fadbdb77be6baf9fa59a7a5e29915ab6a286d2fca","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/config_warn_run.go","path":"os"},"kwargs":"sha256:2c46465fa8a419cc2f95c214c64f019962ae3c953d0a542d432ea03f7fd5ffb6","body":"sha256:7c6e309eb2575d4aa451ec2ab83161d13e9ae03bbca3a4d2ecf30ce12adf0905"}
import "os"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-78ed3e94","gen":"sha256:e715c38ac37ef8a6e260b14fadbdb77be6baf9fa59a7a5e29915ab6a286d2fca","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/config_warn_run.go","path":"path/filepath"},"kwargs":"sha256:a39977dd010b7e01e4212caad925eff848f2940fb6c1fb95c2878a4d8b6445f4","body":"sha256:a6d196b558491e646be93f56a321f339068c64c76aa9d49bb18046e7f9c777a1"}
import "path/filepath"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-78ed3e94","gen":"sha256:e715c38ac37ef8a6e260b14fadbdb77be6baf9fa59a7a5e29915ab6a286d2fca","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/config_warn_run.go","path":"sort"},"kwargs":"sha256:a86ee4d744a51d82f66573c1db05cb5b72cdfffe3914ff4ac875d8bc1cf2869c","body":"sha256:f197172617ae21ef23b1855b09aafb14571020ae81bff877d63d3ff81f4bf63c"}
import "sort"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-78ed3e94","gen":"sha256:e715c38ac37ef8a6e260b14fadbdb77be6baf9fa59a7a5e29915ab6a286d2fca","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/config_warn_run.go","path":"strings"},"kwargs":"sha256:704e5303240bd2e7ffbd23aa40085db2566b98477510df1fa663bb8616b0d20c","body":"sha256:a9af3c0bc4bf9342df9d7e89a324523fdba7941a9ca4190270e68ff3c4a25ea6"}
import "strings"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-78ed3e94","gen":"sha256:e715c38ac37ef8a6e260b14fadbdb77be6baf9fa59a7a5e29915ab6a286d2fca","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/config_warn_run.go","path":"gopkg.in/yaml.v3"},"kwargs":"sha256:a417d3125f32e88d1abebf66a8cbd51c7e82a80dc67b5b77ef9b60a88484325c","body":"sha256:a4214cbff710690b396b73d2b1ccc9aa8b55a807d07f73aa80f10af0c41e13b9"}
import "gopkg.in/yaml.v3"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-78ed3e94","gen":"sha256:e715c38ac37ef8a6e260b14fadbdb77be6baf9fa59a7a5e29915ab6a286d2fca","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/config_warn_run.go","name":"warnNestedConfigs"},"kwargs":"sha256:251541575f30fdcf0590b953dbf9a4976078d21d1cc7af8679c3050745df4525","body":"sha256:3b295f6b31b3723d05f8bac14cdb1577d76a4ba60e86456db2b8c411b487affb"}
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
