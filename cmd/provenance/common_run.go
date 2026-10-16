package main

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-be1d9ebc","gen":"sha256:90da47fc1c4c1d3f9654ed8965388d1e53056cb1b0b73857755dc47a7343f72d","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/common_run.go","path":"bytes"}}
import "bytes"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-be1d9ebc","gen":"sha256:90da47fc1c4c1d3f9654ed8965388d1e53056cb1b0b73857755dc47a7343f72d","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/common_run.go","path":"errors"}}
import "errors"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-be1d9ebc","gen":"sha256:90da47fc1c4c1d3f9654ed8965388d1e53056cb1b0b73857755dc47a7343f72d","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/common_run.go","path":"fmt"}}
import "fmt"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-be1d9ebc","gen":"sha256:90da47fc1c4c1d3f9654ed8965388d1e53056cb1b0b73857755dc47a7343f72d","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/common_run.go","path":"os"}}
import "os"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-be1d9ebc","gen":"sha256:90da47fc1c4c1d3f9654ed8965388d1e53056cb1b0b73857755dc47a7343f72d","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/common_run.go","path":"os/exec"}}
import "os/exec"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-be1d9ebc","gen":"sha256:90da47fc1c4c1d3f9654ed8965388d1e53056cb1b0b73857755dc47a7343f72d","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/common_run.go","path":"path/filepath"}}
import "path/filepath"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-be1d9ebc","gen":"sha256:90da47fc1c4c1d3f9654ed8965388d1e53056cb1b0b73857755dc47a7343f72d","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/common_run.go","path":"strings"}}
import "strings"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-be1d9ebc","gen":"sha256:90da47fc1c4c1d3f9654ed8965388d1e53056cb1b0b73857755dc47a7343f72d","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/common_run.go","path":"github.com/spf13/cobra"}}
import "github.com/spf13/cobra"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-be1d9ebc","gen":"sha256:90da47fc1c4c1d3f9654ed8965388d1e53056cb1b0b73857755dc47a7343f72d","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/common_run.go","path":"gopkg.in/yaml.v3"}}
import "gopkg.in/yaml.v3"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-be1d9ebc","gen":"sha256:90da47fc1c4c1d3f9654ed8965388d1e53056cb1b0b73857755dc47a7343f72d","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/common_run.go","path":"github.com/livecodelife/provenance/record"}}
import "github.com/livecodelife/provenance/record"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setConst {"tier":"owned","record":"prov-2026-be1d9ebc","gen":"sha256:90da47fc1c4c1d3f9654ed8965388d1e53056cb1b0b73857755dc47a7343f72d","pkg":"sha256:51be1c18998d63cc558c253e211a6bd188d96cb251dad68c6ea88e03c2d5cc4c","kwargs":{"code":"const (\n\tconfigFile       = \".provenance.yml\"\n\tlegacyConfigFile = \".linespec.yml\"\n\tmanifestPath     = \".provenance/versions.jsonl\"\n\tlegacyManifest   = \".linespec/hash_manifest.json\"\n)","file":"cmd/provenance/common_run.go","name":"configFile"}}
const (
	configFile       = ".provenance.yml"
	legacyConfigFile = ".linespec.yml"
	manifestPath     = ".provenance/versions.jsonl"
	legacyManifest   = ".linespec/hash_manifest.json"
)

// /sedum:setConst
// sedum:setType {"tier":"owned","record":"prov-2026-be1d9ebc","gen":"sha256:90da47fc1c4c1d3f9654ed8965388d1e53056cb1b0b73857755dc47a7343f72d","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"// sharedRepo is one entry of shared_repos.\ntype sharedRepo struct {\n\tName string `yaml:\"name\"`\n\tURL  string `yaml:\"url\"`\n\tRef  string `yaml:\"ref\"`\n\tDir  string `yaml:\"dir\"`\n}","file":"cmd/provenance/common_run.go","name":"sharedRepo"}}
// sharedRepo is one entry of shared_repos.
type sharedRepo struct {
	Name string `yaml:"name"`
	URL  string `yaml:"url"`
	Ref  string `yaml:"ref"`
	Dir  string `yaml:"dir"`
}

// /sedum:setType
// sedum:setType {"tier":"owned","record":"prov-2026-be1d9ebc","gen":"sha256:90da47fc1c4c1d3f9654ed8965388d1e53056cb1b0b73857755dc47a7343f72d","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"// config is the provenance configuration: the top level of .provenance.yml,\n// or the provenance block of .linespec.yml before a repo is migrated.\ntype config struct {\n\tEnforcement                  string       `yaml:\"enforcement\"`\n\tDir                          string       `yaml:\"dir\"`\n\tExcludePaths                 []string     `yaml:\"exclude_paths\"`\n\tSharedRepos                  []sharedRepo `yaml:\"shared_repos\"`\n\tCacheTTLMinutes              int          `yaml:\"cache_ttl_minutes\"`\n\tCommitTagRequired            bool         `yaml:\"commit_tag_required\"`\n\tAutoAffectedScope            bool         `yaml:\"auto_affected_scope\"`\n\tRunAssociatedSpecsOnComplete bool         `yaml:\"run_associated_specs_on_complete\"`\n\tOverlapSpecsOnComplete       string       `yaml:\"overlap_specs_on_complete\"`\n\tCommitOnStatusChange         bool         `yaml:\"commit_on_status_change\"`\n\tWriteRestriction             *bool        `yaml:\"write_restriction,omitempty\"`\n}","file":"cmd/provenance/common_run.go","name":"config"}}
// config is the provenance configuration: the top level of .provenance.yml,
// or the provenance block of .linespec.yml before a repo is migrated.
type config struct {
	Enforcement                  string       `yaml:"enforcement"`
	Dir                          string       `yaml:"dir"`
	ExcludePaths                 []string     `yaml:"exclude_paths"`
	SharedRepos                  []sharedRepo `yaml:"shared_repos"`
	CacheTTLMinutes              int          `yaml:"cache_ttl_minutes"`
	CommitTagRequired            bool         `yaml:"commit_tag_required"`
	AutoAffectedScope            bool         `yaml:"auto_affected_scope"`
	RunAssociatedSpecsOnComplete bool         `yaml:"run_associated_specs_on_complete"`
	OverlapSpecsOnComplete       string       `yaml:"overlap_specs_on_complete"`
	CommitOnStatusChange         bool         `yaml:"commit_on_status_change"`
	WriteRestriction             *bool        `yaml:"write_restriction,omitempty"`
}

// /sedum:setType
// sedum:setType {"tier":"owned","record":"prov-2026-be1d9ebc","gen":"sha256:90da47fc1c4c1d3f9654ed8965388d1e53056cb1b0b73857755dc47a7343f72d","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"// repo is the repository a command acts on.\ntype repo struct {\n\tRoot   string\n\tConfig config\n\t// ConfigPath is the file the config came from; empty when defaults apply.\n\tConfigPath string\n\t// Legacy is set when the config came from .linespec.yml.\n\tLegacy bool\n\tStore  *record.FilesStore\n\tout    *os.File\n\t// env caches gitEnv.\n\tenv []string\n}","file":"cmd/provenance/common_run.go","name":"repo"}}
// repo is the repository a command acts on.
type repo struct {
	Root   string
	Config config
	// ConfigPath is the file the config came from; empty when defaults apply.
	ConfigPath string
	// Legacy is set when the config came from .linespec.yml.
	Legacy bool
	Store  *record.FilesStore
	out    *os.File
	// env caches gitEnv.
	env []string
}

// /sedum:setType
// sedum:setFunction {"tier":"owned","record":"prov-2026-be1d9ebc","gen":"sha256:90da47fc1c4c1d3f9654ed8965388d1e53056cb1b0b73857755dc47a7343f72d","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"// openRepo finds the repository and its config. -c/--config names the file;\n// otherwise the repo root is git's top level (or the working directory) and\n// .provenance.yml wins over .linespec.yml's provenance block.\nfunc openRepo(cmd *cobra.Command) (*repo, error) {\n\tr := &repo{out: os.Stdout}\n\texplicit, _ := cmd.Flags().GetString(\"config\")\n\tif explicit != \"\" {\n\t\tabs, err := filepath.Abs(explicit)\n\t\tif err != nil {\n\t\t\treturn nil, err\n\t\t}\n\t\tr.Root = filepath.Dir(abs)\n\t\tr.ConfigPath = abs\n\t\tr.Legacy = filepath.Base(abs) == legacyConfigFile\n\t} else {\n\t\tr.Root = repoRoot()\n\t\tfor _, name := range []string{configFile, legacyConfigFile} {\n\t\t\tp := filepath.Join(r.Root, name)\n\t\t\tif _, err := os.Stat(p); err == nil {\n\t\t\t\tr.ConfigPath, r.Legacy = p, name == legacyConfigFile\n\t\t\t\tbreak\n\t\t\t}\n\t\t}\n\t}\n\tif r.ConfigPath != \"\" {\n\t\tcfg, err := readConfig(r.ConfigPath, r.Legacy)\n\t\tif err != nil {\n\t\t\treturn nil, err\n\t\t}\n\t\tr.Config = cfg\n\t}\n\tif r.Config.Dir == \"\" {\n\t\tr.Config.Dir = \"provenance\"\n\t}\n\tr.Config.Dir = strings.TrimRight(r.Config.Dir, \"/\")\n\tr.Store = &record.FilesStore{Root: r.Root, Dir: r.Config.Dir}\n\treturn r, nil\n}","file":"cmd/provenance/common_run.go","name":"openRepo"}}
// openRepo finds the repository and its config. -c/--config names the file;
// otherwise the repo root is git's top level (or the working directory) and
// .provenance.yml wins over .linespec.yml's provenance block.
func openRepo(cmd *cobra.Command) (*repo, error) {
	r := &repo{out: os.Stdout}
	explicit, _ := cmd.Flags().GetString("config")
	if explicit != "" {
		abs, err := filepath.Abs(explicit)
		if err != nil {
			return nil, err
		}
		r.Root = filepath.Dir(abs)
		r.ConfigPath = abs
		r.Legacy = filepath.Base(abs) == legacyConfigFile
	} else {
		r.Root = repoRoot()
		for _, name := range []string{configFile, legacyConfigFile} {
			p := filepath.Join(r.Root, name)
			if _, err := os.Stat(p); err == nil {
				r.ConfigPath, r.Legacy = p, name == legacyConfigFile
				break
			}
		}
	}
	if r.ConfigPath != "" {
		cfg, err := readConfig(r.ConfigPath, r.Legacy)
		if err != nil {
			return nil, err
		}
		r.Config = cfg
	}
	if r.Config.Dir == "" {
		r.Config.Dir = "provenance"
	}
	r.Config.Dir = strings.TrimRight(r.Config.Dir, "/")
	r.Store = &record.FilesStore{Root: r.Root, Dir: r.Config.Dir}
	return r, nil
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-be1d9ebc","gen":"sha256:90da47fc1c4c1d3f9654ed8965388d1e53056cb1b0b73857755dc47a7343f72d","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func repoRoot() string {\n\tif out, err := exec.Command(\"git\", \"rev-parse\", \"--show-toplevel\").Output(); err == nil {\n\t\treturn strings.TrimSpace(string(out))\n\t}\n\twd, _ := os.Getwd()\n\treturn wd\n}","file":"cmd/provenance/common_run.go","name":"repoRoot"}}
func repoRoot() string {
	if out, err := exec.Command("git", "rev-parse", "--show-toplevel").Output(); err == nil {
		return strings.TrimSpace(string(out))
	}
	wd, _ := os.Getwd()
	return wd
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-be1d9ebc","gen":"sha256:90da47fc1c4c1d3f9654ed8965388d1e53056cb1b0b73857755dc47a7343f72d","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func readConfig(path string, legacy bool) (config, error) {\n\tvar cfg config\n\tdata, err := os.ReadFile(path)\n\tif err != nil {\n\t\treturn cfg, err\n\t}\n\tif legacy {\n\t\tvar wrap struct {\n\t\t\tProvenance config `yaml:\"provenance\"`\n\t\t}\n\t\tif err := yaml.Unmarshal(data, &wrap); err != nil {\n\t\t\treturn cfg, fmt.Errorf(\"%s: %w\", path, err)\n\t\t}\n\t\treturn wrap.Provenance, nil\n\t}\n\tif err := yaml.Unmarshal(data, &cfg); err != nil {\n\t\treturn cfg, fmt.Errorf(\"%s: %w\", path, err)\n\t}\n\treturn cfg, nil\n}","file":"cmd/provenance/common_run.go","name":"readConfig"}}
func readConfig(path string, legacy bool) (config, error) {
	var cfg config
	data, err := os.ReadFile(path)
	if err != nil {
		return cfg, err
	}
	if legacy {
		var wrap struct {
			Provenance config `yaml:"provenance"`
		}
		if err := yaml.Unmarshal(data, &wrap); err != nil {
			return cfg, fmt.Errorf("%s: %w", path, err)
		}
		return wrap.Provenance, nil
	}
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return cfg, fmt.Errorf("%s: %w", path, err)
	}
	return cfg, nil
}

// /sedum:setFunction
// sedum:setVar {"tier":"owned","record":"prov-2026-be1d9ebc","gen":"sha256:90da47fc1c4c1d3f9654ed8965388d1e53056cb1b0b73857755dc47a7343f72d","pkg":"sha256:c7b9d9ac1655e8211b8b2ca8c49cf34b8d510d8d823d4f5eb60ec7e1018aa2eb","kwargs":{"code":"// errNotFound is returned, after printing linespec's message, when an id\n// matches neither a record nor any record's aliases.\nvar errNotFound = errors.New(\"record not found\")","file":"cmd/provenance/common_run.go","name":"errNotFound"}}
// errNotFound is returned, after printing linespec's message, when an id
// matches neither a record nor any record's aliases.
var errNotFound = errors.New("record not found")

// /sedum:setVar
// sedum:setMethod {"tier":"owned","record":"prov-2026-be1d9ebc","gen":"sha256:90da47fc1c4c1d3f9654ed8965388d1e53056cb1b0b73857755dc47a7343f72d","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"// resolve loads the record named by id, or the record that lists id among\n// its aliases.\nfunc (r *repo) resolve(id string) (*record.Record, *record.Document, error) {\n\trec, doc, err := r.Store.Load(id)\n\tif err == nil {\n\t\treturn rec, doc, nil\n\t}\n\tif !os.IsNotExist(errors.Unwrap(err)) && !os.IsNotExist(err) {\n\t\treturn nil, nil, err\n\t}\n\tids, err := r.Store.List()\n\tif err != nil {\n\t\treturn nil, nil, err\n\t}\n\tfor _, other := range ids {\n\t\trec, doc, err := r.Store.Load(other)\n\t\tif err != nil {\n\t\t\tcontinue\n\t\t}\n\t\tfor _, a := range rec.Aliases {\n\t\t\tif a == id {\n\t\t\t\treturn rec, doc, nil\n\t\t\t}\n\t\t}\n\t}\n\tr.fail(\"Record not found: %s\", id)\n\treturn nil, nil, errNotFound\n}","file":"cmd/provenance/common_run.go","name":"resolve","receiver":"repo"}}
// resolve loads the record named by id, or the record that lists id among
// its aliases.
func (r *repo) resolve(id string) (*record.Record, *record.Document, error) {
	rec, doc, err := r.Store.Load(id)
	if err == nil {
		return rec, doc, nil
	}
	if !os.IsNotExist(errors.Unwrap(err)) && !os.IsNotExist(err) {
		return nil, nil, err
	}
	ids, err := r.Store.List()
	if err != nil {
		return nil, nil, err
	}
	for _, other := range ids {
		rec, doc, err := r.Store.Load(other)
		if err != nil {
			continue
		}
		for _, a := range rec.Aliases {
			if a == id {
				return rec, doc, nil
			}
		}
	}
	r.fail("Record not found: %s", id)
	return nil, nil, errNotFound
}

// /sedum:setMethod
// sedum:setMethod {"tier":"owned","record":"prov-2026-be1d9ebc","gen":"sha256:90da47fc1c4c1d3f9654ed8965388d1e53056cb1b0b73857755dc47a7343f72d","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"// recordPath is the record file's path relative to the repo root.\nfunc (r *repo) recordPath(id string) string {\n\treturn filepath.Join(r.Config.Dir, id+\".yml\")\n}","file":"cmd/provenance/common_run.go","name":"recordPath","receiver":"repo"}}
// recordPath is the record file's path relative to the repo root.
func (r *repo) recordPath(id string) string {
	return filepath.Join(r.Config.Dir, id+".yml")
}

// /sedum:setMethod
// sedum:setMethod {"tier":"owned","record":"prov-2026-be1d9ebc","gen":"sha256:90da47fc1c4c1d3f9654ed8965388d1e53056cb1b0b73857755dc47a7343f72d","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"// writeManifest rewrites .provenance/versions.jsonl from the latest snapshot\n// of every record that has one. The format is the library's, so the writer\n// and lint's PROV026 check share one definition.\nfunc (r *repo) writeManifest() error {\n\tlines, err := record.BuildManifest(r.Store)\n\tif err != nil {\n\t\treturn err\n\t}\n\tpath := filepath.Join(r.Root, manifestPath)\n\tif err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {\n\t\treturn err\n\t}\n\treturn os.WriteFile(path, record.MarshalManifest(lines), 0o644)\n}","file":"cmd/provenance/common_run.go","name":"writeManifest","receiver":"repo"}}
// writeManifest rewrites .provenance/versions.jsonl from the latest snapshot
// of every record that has one. The format is the library's, so the writer
// and lint's PROV026 check share one definition.
func (r *repo) writeManifest() error {
	lines, err := record.BuildManifest(r.Store)
	if err != nil {
		return err
	}
	path := filepath.Join(r.Root, manifestPath)
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	return os.WriteFile(path, record.MarshalManifest(lines), 0o644)
}

// /sedum:setMethod
// sedum:setMethod {"tier":"owned","record":"prov-2026-be1d9ebc","gen":"sha256:90da47fc1c4c1d3f9654ed8965388d1e53056cb1b0b73857755dc47a7343f72d","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"// versionsDir is a record's snapshot directory relative to the repo root.\nfunc (r *repo) versionsDir(id string) string {\n\treturn filepath.Join(r.Config.Dir, \".versions\", id)\n}","file":"cmd/provenance/common_run.go","name":"versionsDir","receiver":"repo"}}
// versionsDir is a record's snapshot directory relative to the repo root.
func (r *repo) versionsDir(id string) string {
	return filepath.Join(r.Config.Dir, ".versions", id)
}

// /sedum:setMethod
// sedum:setMethod {"tier":"owned","record":"prov-2026-be1d9ebc","gen":"sha256:90da47fc1c4c1d3f9654ed8965388d1e53056cb1b0b73857755dc47a7343f72d","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"// gitEnv is the environment for git commands on r.Root. A hook exports\n// GIT_DIR, GIT_INDEX_FILE and friends for the repository being committed.\n// They are kept when that repository is r.Root's own, because check --staged\n// must read the index being committed (git commit -a and git commit <paths>\n// use a temporary one), and dropped otherwise, so a command acting on another\n// repository never reads or writes the hooked one (the 2026-10-15 incident:\n// fixture git commands rewrote the real .git/config).\nfunc (r *repo) gitEnv() []string {\n\tif r.env != nil {\n\t\treturn r.env\n\t}\n\tenv := os.Environ()\n\tvar clean []string\n\tinherited := false\n\tfor _, kv := range env {\n\t\tif strings.HasPrefix(kv, \"GIT_\") {\n\t\t\tinherited = true\n\t\t\tcontinue\n\t\t}\n\t\tclean = append(clean, kv)\n\t}\n\tr.env = clean\n\tif !inherited {\n\t\tr.env = env\n\t\treturn r.env\n\t}\n\thooked := os.Getenv(\"GIT_DIR\")\n\tif hooked == \"\" {\n\t\twd, _ := os.Getwd()\n\t\thooked = gitDirOf(wd, clean)\n\t} else if !filepath.IsAbs(hooked) {\n\t\twd, _ := os.Getwd()\n\t\thooked = filepath.Join(wd, hooked)\n\t}\n\tif own := gitDirOf(r.Root, clean); own != \"\" && samePath(own, hooked) {\n\t\tr.env = env\n\t}\n\treturn r.env\n}","file":"cmd/provenance/common_run.go","name":"gitEnv","receiver":"repo"}}
// gitEnv is the environment for git commands on r.Root. A hook exports
// GIT_DIR, GIT_INDEX_FILE and friends for the repository being committed.
// They are kept when that repository is r.Root's own, because check --staged
// must read the index being committed (git commit -a and git commit <paths>
// use a temporary one), and dropped otherwise, so a command acting on another
// repository never reads or writes the hooked one (the 2026-10-15 incident:
// fixture git commands rewrote the real .git/config).
func (r *repo) gitEnv() []string {
	if r.env != nil {
		return r.env
	}
	env := os.Environ()
	var clean []string
	inherited := false
	for _, kv := range env {
		if strings.HasPrefix(kv, "GIT_") {
			inherited = true
			continue
		}
		clean = append(clean, kv)
	}
	r.env = clean
	if !inherited {
		r.env = env
		return r.env
	}
	hooked := os.Getenv("GIT_DIR")
	if hooked == "" {
		wd, _ := os.Getwd()
		hooked = gitDirOf(wd, clean)
	} else if !filepath.IsAbs(hooked) {
		wd, _ := os.Getwd()
		hooked = filepath.Join(wd, hooked)
	}
	if own := gitDirOf(r.Root, clean); own != "" && samePath(own, hooked) {
		r.env = env
	}
	return r.env
}

// /sedum:setMethod
// sedum:setFunction {"tier":"owned","record":"prov-2026-be1d9ebc","gen":"sha256:90da47fc1c4c1d3f9654ed8965388d1e53056cb1b0b73857755dc47a7343f72d","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"// gitDirOf is the absolute git dir of the repository holding dir, found with\n// env (which carries no GIT_* variables), or \"\" when there is none.\nfunc gitDirOf(dir string, env []string) string {\n\tc := exec.Command(\"git\", \"-C\", dir, \"rev-parse\", \"--absolute-git-dir\")\n\tc.Env = env\n\tout, err := c.Output()\n\tif err != nil {\n\t\treturn \"\"\n\t}\n\treturn strings.TrimSpace(string(out))\n}","file":"cmd/provenance/common_run.go","name":"gitDirOf"}}
// gitDirOf is the absolute git dir of the repository holding dir, found with
// env (which carries no GIT_* variables), or "" when there is none.
func gitDirOf(dir string, env []string) string {
	c := exec.Command("git", "-C", dir, "rev-parse", "--absolute-git-dir")
	c.Env = env
	out, err := c.Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(out))
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-be1d9ebc","gen":"sha256:90da47fc1c4c1d3f9654ed8965388d1e53056cb1b0b73857755dc47a7343f72d","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func samePath(a, b string) bool {\n\tra, errA := filepath.EvalSymlinks(a)\n\trb, errB := filepath.EvalSymlinks(b)\n\treturn errA == nil && errB == nil && ra == rb\n}","file":"cmd/provenance/common_run.go","name":"samePath"}}
func samePath(a, b string) bool {
	ra, errA := filepath.EvalSymlinks(a)
	rb, errB := filepath.EvalSymlinks(b)
	return errA == nil && errB == nil && ra == rb
}

// /sedum:setFunction
// sedum:setMethod {"tier":"owned","record":"prov-2026-be1d9ebc","gen":"sha256:90da47fc1c4c1d3f9654ed8965388d1e53056cb1b0b73857755dc47a7343f72d","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"// git runs git in the repo root and returns trimmed stdout.\nfunc (r *repo) git(args ...string) (string, error) {\n\tc := exec.Command(\"git\", args...)\n\tc.Dir = r.Root\n\tc.Env = r.gitEnv()\n\tvar stderr bytes.Buffer\n\tc.Stderr = &stderr\n\tout, err := c.Output()\n\tif err != nil {\n\t\treturn \"\", fmt.Errorf(\"git %s: %w\\n%s\", strings.Join(args, \" \"), err, stderr.String())\n\t}\n\treturn strings.TrimSpace(string(out)), nil\n}","file":"cmd/provenance/common_run.go","name":"git","receiver":"repo"}}
// git runs git in the repo root and returns trimmed stdout.
func (r *repo) git(args ...string) (string, error) {
	c := exec.Command("git", args...)
	c.Dir = r.Root
	c.Env = r.gitEnv()
	var stderr bytes.Buffer
	c.Stderr = &stderr
	out, err := c.Output()
	if err != nil {
		return "", fmt.Errorf("git %s: %w\n%s", strings.Join(args, " "), err, stderr.String())
	}
	return strings.TrimSpace(string(out)), nil
}

// /sedum:setMethod
// sedum:setMethod {"tier":"owned","record":"prov-2026-be1d9ebc","gen":"sha256:90da47fc1c4c1d3f9654ed8965388d1e53056cb1b0b73857755dc47a7343f72d","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"// author is git's user.email, as linespec records it.\nfunc (r *repo) author() string {\n\tif email, err := r.git(\"config\", \"user.email\"); err == nil && email != \"\" {\n\t\treturn email\n\t}\n\treturn \"unknown@example.com\"\n}","file":"cmd/provenance/common_run.go","name":"author","receiver":"repo"}}
// author is git's user.email, as linespec records it.
func (r *repo) author() string {
	if email, err := r.git("config", "user.email"); err == nil && email != "" {
		return email
	}
	return "unknown@example.com"
}

// /sedum:setMethod
// sedum:setMethod {"tier":"owned","record":"prov-2026-be1d9ebc","gen":"sha256:90da47fc1c4c1d3f9654ed8965388d1e53056cb1b0b73857755dc47a7343f72d","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"// commit stages paths and commits them, letting hooks print to the terminal.\nfunc (r *repo) commit(message string, paths ...string) error {\n\tif _, err := r.git(append([]string{\"add\", \"--\"}, paths...)...); err != nil {\n\t\treturn err\n\t}\n\tc := exec.Command(\"git\", \"commit\", \"-m\", message)\n\tc.Dir = r.Root\n\tc.Env = r.gitEnv()\n\tc.Stdout, c.Stderr = os.Stdout, os.Stderr\n\tif err := c.Run(); err != nil {\n\t\treturn fmt.Errorf(\"git commit failed: %w\", err)\n\t}\n\treturn nil\n}","file":"cmd/provenance/common_run.go","name":"commit","receiver":"repo"}}
// commit stages paths and commits them, letting hooks print to the terminal.
func (r *repo) commit(message string, paths ...string) error {
	if _, err := r.git(append([]string{"add", "--"}, paths...)...); err != nil {
		return err
	}
	c := exec.Command("git", "commit", "-m", message)
	c.Dir = r.Root
	c.Env = r.gitEnv()
	c.Stdout, c.Stderr = os.Stdout, os.Stderr
	if err := c.Run(); err != nil {
		return fmt.Errorf("git commit failed: %w", err)
	}
	return nil
}

// /sedum:setMethod
// sedum:setMethod {"tier":"owned","record":"prov-2026-be1d9ebc","gen":"sha256:90da47fc1c4c1d3f9654ed8965388d1e53056cb1b0b73857755dc47a7343f72d","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"// unstage undoes commit's `git add` after a rejected commit.\nfunc (r *repo) unstage(paths ...string) {\n\tif _, err := r.git(append([]string{\"reset\", \"--quiet\", \"HEAD\", \"--\"}, paths...)...); err != nil {\n\t\tfmt.Fprintf(os.Stderr, \"Warning: failed to unstage during rollback: %v\\n\", err)\n\t}\n}","file":"cmd/provenance/common_run.go","name":"unstage","receiver":"repo"}}
// unstage undoes commit's `git add` after a rejected commit.
func (r *repo) unstage(paths ...string) {
	if _, err := r.git(append([]string{"reset", "--quiet", "HEAD", "--"}, paths...)...); err != nil {
		fmt.Fprintf(os.Stderr, "Warning: failed to unstage during rollback: %v\n", err)
	}
}

// /sedum:setMethod
// sedum:setType {"tier":"owned","record":"prov-2026-be1d9ebc","gen":"sha256:90da47fc1c4c1d3f9654ed8965388d1e53056cb1b0b73857755dc47a7343f72d","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"// fileState is a file's bytes before a transition, for rollback. A missing\n// file is restored by removing it.\ntype fileState struct {\n\tpath   string\n\tdata   []byte\n\texists bool\n}","file":"cmd/provenance/common_run.go","name":"fileState"}}
// fileState is a file's bytes before a transition, for rollback. A missing
// file is restored by removing it.
type fileState struct {
	path   string
	data   []byte
	exists bool
}

// /sedum:setType
// sedum:setFunction {"tier":"owned","record":"prov-2026-be1d9ebc","gen":"sha256:90da47fc1c4c1d3f9654ed8965388d1e53056cb1b0b73857755dc47a7343f72d","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func snapshotFiles(paths ...string) []fileState {\n\tout := make([]fileState, 0, len(paths))\n\tfor _, p := range paths {\n\t\tdata, err := os.ReadFile(p)\n\t\tout = append(out, fileState{path: p, data: data, exists: err == nil})\n\t}\n\treturn out\n}","file":"cmd/provenance/common_run.go","name":"snapshotFiles"}}
func snapshotFiles(paths ...string) []fileState {
	out := make([]fileState, 0, len(paths))
	for _, p := range paths {
		data, err := os.ReadFile(p)
		out = append(out, fileState{path: p, data: data, exists: err == nil})
	}
	return out
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-be1d9ebc","gen":"sha256:90da47fc1c4c1d3f9654ed8965388d1e53056cb1b0b73857755dc47a7343f72d","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func restoreFiles(states []fileState) {\n\tfor _, s := range states {\n\t\tif s.exists {\n\t\t\t_ = os.WriteFile(s.path, s.data, 0o644)\n\t\t} else {\n\t\t\t_ = os.RemoveAll(s.path)\n\t\t}\n\t}\n}","file":"cmd/provenance/common_run.go","name":"restoreFiles"}}
func restoreFiles(states []fileState) {
	for _, s := range states {
		if s.exists {
			_ = os.WriteFile(s.path, s.data, 0o644)
		} else {
			_ = os.RemoveAll(s.path)
		}
	}
}

// /sedum:setFunction
// sedum:setMethod {"tier":"owned","record":"prov-2026-be1d9ebc","gen":"sha256:90da47fc1c4c1d3f9654ed8965388d1e53056cb1b0b73857755dc47a7343f72d","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"// fail prints an error the way linespec's formatter does.\nfunc (r *repo) fail(format string, args ...any) {\n\tfmt.Fprintf(os.Stderr, \"\\n✗ \"+format+\"\\n\\n\", args...)\n}","file":"cmd/provenance/common_run.go","name":"fail","receiver":"repo"}}
// fail prints an error the way linespec's formatter does.
func (r *repo) fail(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "\n✗ "+format+"\n\n", args...)
}

// /sedum:setMethod
// sedum:setFunction {"tier":"owned","record":"prov-2026-be1d9ebc","gen":"sha256:90da47fc1c4c1d3f9654ed8965388d1e53056cb1b0b73857755dc47a7343f72d","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"// tag is the commit tag for a record version.\nfunc tag(id string, version int) string {\n\treturn fmt.Sprintf(\"[%s@%d]\", id, version)\n}","file":"cmd/provenance/common_run.go","name":"tag"}}
// tag is the commit tag for a record version.
func tag(id string, version int) string {
	return fmt.Sprintf("[%s@%d]", id, version)
}

// /sedum:setFunction
