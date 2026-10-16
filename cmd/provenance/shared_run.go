package main

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-f807707a","gen":"sha256:3f6d1802d83ba9dcc72ed02c8927ae3f79e56242263a8c87803b425321b88f73","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/shared_run.go","path":"crypto/sha256"}}
import "crypto/sha256"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-f807707a","gen":"sha256:3f6d1802d83ba9dcc72ed02c8927ae3f79e56242263a8c87803b425321b88f73","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/shared_run.go","path":"encoding/json"}}
import "encoding/json"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-f807707a","gen":"sha256:3f6d1802d83ba9dcc72ed02c8927ae3f79e56242263a8c87803b425321b88f73","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/shared_run.go","path":"fmt"}}
import "fmt"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-f807707a","gen":"sha256:3f6d1802d83ba9dcc72ed02c8927ae3f79e56242263a8c87803b425321b88f73","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/shared_run.go","path":"io"}}
import "io"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-f807707a","gen":"sha256:3f6d1802d83ba9dcc72ed02c8927ae3f79e56242263a8c87803b425321b88f73","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/shared_run.go","path":"io/fs"}}
import "io/fs"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-f807707a","gen":"sha256:3f6d1802d83ba9dcc72ed02c8927ae3f79e56242263a8c87803b425321b88f73","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/shared_run.go","path":"os"}}
import "os"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-f807707a","gen":"sha256:3f6d1802d83ba9dcc72ed02c8927ae3f79e56242263a8c87803b425321b88f73","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/shared_run.go","path":"os/exec"}}
import "os/exec"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-f807707a","gen":"sha256:3f6d1802d83ba9dcc72ed02c8927ae3f79e56242263a8c87803b425321b88f73","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/shared_run.go","path":"path/filepath"}}
import "path/filepath"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-f807707a","gen":"sha256:3f6d1802d83ba9dcc72ed02c8927ae3f79e56242263a8c87803b425321b88f73","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/shared_run.go","path":"strings"}}
import "strings"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-f807707a","gen":"sha256:3f6d1802d83ba9dcc72ed02c8927ae3f79e56242263a8c87803b425321b88f73","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/shared_run.go","path":"time"}}
import "time"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-f807707a","gen":"sha256:3f6d1802d83ba9dcc72ed02c8927ae3f79e56242263a8c87803b425321b88f73","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/shared_run.go","path":"github.com/spf13/cobra"}}
import "github.com/spf13/cobra"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-f807707a","gen":"sha256:3f6d1802d83ba9dcc72ed02c8927ae3f79e56242263a8c87803b425321b88f73","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/shared_run.go","path":"github.com/livecodelife/provenance/record"}}
import "github.com/livecodelife/provenance/record"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setType {"tier":"owned","record":"prov-2026-f807707a","gen":"sha256:3f6d1802d83ba9dcc72ed02c8927ae3f79e56242263a8c87803b425321b88f73","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type cacheMeta struct {\n\tFetchedAt   time.Time `json:\"fetched_at\"`\n\tRecordCount int       `json:\"record_count\"`\n\tRepoName    string    `json:\"repo_name\"`\n\tRepoURL     string    `json:\"repo_url\"`\n\tRepoRef     string    `json:\"repo_ref\"`\n}","file":"cmd/provenance/shared_run.go","name":"cacheMeta"}}
type cacheMeta struct {
	FetchedAt   time.Time `json:"fetched_at"`
	RecordCount int       `json:"record_count"`
	RepoName    string    `json:"repo_name"`
	RepoURL     string    `json:"repo_url"`
	RepoRef     string    `json:"repo_ref"`
}

// /sedum:setType
// sedum:setFunction {"tier":"owned","record":"prov-2026-f807707a","gen":"sha256:3f6d1802d83ba9dcc72ed02c8927ae3f79e56242263a8c87803b425321b88f73","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func sharedCacheRoot(sr sharedRepo) string {\n\thome, _ := os.UserHomeDir()\n\treturn filepath.Join(home, \".linespec\", \"cache\", fmt.Sprintf(\"%x\", sha256.Sum256([]byte(sr.URL))))\n}","file":"cmd/provenance/shared_run.go","name":"sharedCacheRoot"}}
func sharedCacheRoot(sr sharedRepo) string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".linespec", "cache", fmt.Sprintf("%x", sha256.Sum256([]byte(sr.URL))))
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-f807707a","gen":"sha256:3f6d1802d83ba9dcc72ed02c8927ae3f79e56242263a8c87803b425321b88f73","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func sharedCacheDir(sr sharedRepo) string {\n\treturn filepath.Join(sharedCacheRoot(sr), \"provenance\")\n}","file":"cmd/provenance/shared_run.go","name":"sharedCacheDir"}}
func sharedCacheDir(sr sharedRepo) string {
	return filepath.Join(sharedCacheRoot(sr), "provenance")
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-f807707a","gen":"sha256:3f6d1802d83ba9dcc72ed02c8927ae3f79e56242263a8c87803b425321b88f73","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func readCacheMeta(sr sharedRepo) (cacheMeta, error) {\n\tvar meta cacheMeta\n\tdata, err := os.ReadFile(filepath.Join(sharedCacheRoot(sr), \".meta\"))\n\tif err != nil {\n\t\treturn meta, err\n\t}\n\treturn meta, json.Unmarshal(data, &meta)\n}","file":"cmd/provenance/shared_run.go","name":"readCacheMeta"}}
func readCacheMeta(sr sharedRepo) (cacheMeta, error) {
	var meta cacheMeta
	data, err := os.ReadFile(filepath.Join(sharedCacheRoot(sr), ".meta"))
	if err != nil {
		return meta, err
	}
	return meta, json.Unmarshal(data, &meta)
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-f807707a","gen":"sha256:3f6d1802d83ba9dcc72ed02c8927ae3f79e56242263a8c87803b425321b88f73","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func cacheTTL(r *repo) time.Duration {\n\tminutes := r.Config.CacheTTLMinutes\n\tif minutes <= 0 {\n\t\tminutes = 60\n\t}\n\treturn time.Duration(minutes) * time.Minute\n}","file":"cmd/provenance/shared_run.go","name":"cacheTTL"}}
func cacheTTL(r *repo) time.Duration {
	minutes := r.Config.CacheTTLMinutes
	if minutes <= 0 {
		minutes = 60
	}
	return time.Duration(minutes) * time.Minute
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-f807707a","gen":"sha256:3f6d1802d83ba9dcc72ed02c8927ae3f79e56242263a8c87803b425321b88f73","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func cacheFresh(r *repo, sr sharedRepo) bool {\n\tmeta, err := readCacheMeta(sr)\n\treturn err == nil && time.Since(meta.FetchedAt) < cacheTTL(r)\n}","file":"cmd/provenance/shared_run.go","name":"cacheFresh"}}
func cacheFresh(r *repo, sr sharedRepo) bool {
	meta, err := readCacheMeta(sr)
	return err == nil && time.Since(meta.FetchedAt) < cacheTTL(r)
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-f807707a","gen":"sha256:3f6d1802d83ba9dcc72ed02c8927ae3f79e56242263a8c87803b425321b88f73","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func loadShared(r *repo) (map[string]*record.Record, []string) {\n\tshared := map[string]*record.Record{}\n\tvar warnings []string\n\tfor _, sr := range r.Config.SharedRepos {\n\t\tdir := sharedCacheDir(sr)\n\t\tif _, err := os.Stat(dir); err != nil {\n\t\t\twarnings = append(warnings, fmt.Sprintf(\"shared repo %s has no cache: run 'provenance sync'\", sr.Name))\n\t\t\tcontinue\n\t\t}\n\t\tif !cacheFresh(r, sr) {\n\t\t\twarnings = append(warnings, fmt.Sprintf(\"the cache of shared repo %s is stale: run 'provenance sync'\", sr.Name))\n\t\t}\n\t\tfilepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {\n\t\t\tif err != nil {\n\t\t\t\treturn nil\n\t\t\t}\n\t\t\tname := d.Name()\n\t\t\tif d.IsDir() {\n\t\t\t\tif name == \".versions\" {\n\t\t\t\t\treturn filepath.SkipDir\n\t\t\t\t}\n\t\t\t\treturn nil\n\t\t\t}\n\t\t\tif !strings.HasPrefix(name, \"prov-\") || !(strings.HasSuffix(name, \".yml\") || strings.HasSuffix(name, \".yaml\")) {\n\t\t\t\treturn nil\n\t\t\t}\n\t\t\tdata, err := os.ReadFile(path)\n\t\t\tif err != nil {\n\t\t\t\treturn nil\n\t\t\t}\n\t\t\tif rec, err := record.Decode(data); err == nil && rec.ID != \"\" {\n\t\t\t\tshared[rec.ID] = rec\n\t\t\t}\n\t\t\treturn nil\n\t\t})\n\t}\n\treturn shared, warnings\n}","file":"cmd/provenance/shared_run.go","name":"loadShared"}}
func loadShared(r *repo) (map[string]*record.Record, []string) {
	shared := map[string]*record.Record{}
	var warnings []string
	for _, sr := range r.Config.SharedRepos {
		dir := sharedCacheDir(sr)
		if _, err := os.Stat(dir); err != nil {
			warnings = append(warnings, fmt.Sprintf("shared repo %s has no cache: run 'provenance sync'", sr.Name))
			continue
		}
		if !cacheFresh(r, sr) {
			warnings = append(warnings, fmt.Sprintf("the cache of shared repo %s is stale: run 'provenance sync'", sr.Name))
		}
		filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return nil
			}
			name := d.Name()
			if d.IsDir() {
				if name == ".versions" {
					return filepath.SkipDir
				}
				return nil
			}
			if !strings.HasPrefix(name, "prov-") || !(strings.HasSuffix(name, ".yml") || strings.HasSuffix(name, ".yaml")) {
				return nil
			}
			data, err := os.ReadFile(path)
			if err != nil {
				return nil
			}
			if rec, err := record.Decode(data); err == nil && rec.ID != "" {
				shared[rec.ID] = rec
			}
			return nil
		})
	}
	return shared, warnings
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-f807707a","gen":"sha256:3f6d1802d83ba9dcc72ed02c8927ae3f79e56242263a8c87803b425321b88f73","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func sharedNames(r *repo) []string {\n\tvar names []string\n\tfor _, sr := range r.Config.SharedRepos {\n\t\tnames = append(names, sr.Name)\n\t}\n\treturn names\n}","file":"cmd/provenance/shared_run.go","name":"sharedNames"}}
func sharedNames(r *repo) []string {
	var names []string
	for _, sr := range r.Config.SharedRepos {
		names = append(names, sr.Name)
	}
	return names
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-f807707a","gen":"sha256:3f6d1802d83ba9dcc72ed02c8927ae3f79e56242263a8c87803b425321b88f73","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func cleanGitEnv() []string {\n\tvar env []string\n\tfor _, kv := range os.Environ() {\n\t\tif !strings.HasPrefix(kv, \"GIT_\") {\n\t\t\tenv = append(env, kv)\n\t\t}\n\t}\n\treturn env\n}","file":"cmd/provenance/shared_run.go","name":"cleanGitEnv"}}
func cleanGitEnv() []string {
	var env []string
	for _, kv := range os.Environ() {
		if !strings.HasPrefix(kv, "GIT_") {
			env = append(env, kv)
		}
	}
	return env
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-f807707a","gen":"sha256:3f6d1802d83ba9dcc72ed02c8927ae3f79e56242263a8c87803b425321b88f73","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func fetchShared(sr sharedRepo, ref, dir string) (string, func(), error) {\n\ttmp, err := os.MkdirTemp(\"\", \"provenance-sync-\")\n\tif err != nil {\n\t\treturn \"\", nil, err\n\t}\n\tcleanup := func() { os.RemoveAll(tmp) }\n\tc := exec.Command(\"sh\", \"-c\", `git archive --remote=\"$0\" \"$1\" \"$2/\" | tar -x -C \"$3\"`, sr.URL, ref, dir, tmp)\n\tc.Env = cleanGitEnv()\n\tif err := c.Run(); err == nil {\n\t\tif _, err := os.Stat(filepath.Join(tmp, filepath.FromSlash(dir))); err == nil {\n\t\t\treturn tmp, cleanup, nil\n\t\t}\n\t}\n\tclone, err := os.MkdirTemp(\"\", \"provenance-sync-clone-\")\n\tif err != nil {\n\t\tcleanup()\n\t\treturn \"\", nil, err\n\t}\n\tboth := func() { cleanup(); os.RemoveAll(clone) }\n\tfor _, args := range [][]string{\n\t\t{\"clone\", \"--depth=1\", \"--filter=blob:none\", \"--sparse\", \"--branch=\" + ref, sr.URL, \".\"},\n\t\t{\"sparse-checkout\", \"set\", \"--no-cone\", dir},\n\t} {\n\t\tc := exec.Command(\"git\", args...)\n\t\tc.Dir = clone\n\t\tc.Env = cleanGitEnv()\n\t\tif out, err := c.CombinedOutput(); err != nil {\n\t\t\tboth()\n\t\t\treturn \"\", nil, fmt.Errorf(\"git %s failed for %s: %s\", args[0], sr.Name, strings.TrimSpace(string(out)))\n\t\t}\n\t}\n\treturn clone, both, nil\n}","file":"cmd/provenance/shared_run.go","name":"fetchShared"}}
func fetchShared(sr sharedRepo, ref, dir string) (string, func(), error) {
	tmp, err := os.MkdirTemp("", "provenance-sync-")
	if err != nil {
		return "", nil, err
	}
	cleanup := func() { os.RemoveAll(tmp) }
	c := exec.Command("sh", "-c", `git archive --remote="$0" "$1" "$2/" | tar -x -C "$3"`, sr.URL, ref, dir, tmp)
	c.Env = cleanGitEnv()
	if err := c.Run(); err == nil {
		if _, err := os.Stat(filepath.Join(tmp, filepath.FromSlash(dir))); err == nil {
			return tmp, cleanup, nil
		}
	}
	clone, err := os.MkdirTemp("", "provenance-sync-clone-")
	if err != nil {
		cleanup()
		return "", nil, err
	}
	both := func() { cleanup(); os.RemoveAll(clone) }
	for _, args := range [][]string{
		{"clone", "--depth=1", "--filter=blob:none", "--sparse", "--branch=" + ref, sr.URL, "."},
		{"sparse-checkout", "set", "--no-cone", dir},
	} {
		c := exec.Command("git", args...)
		c.Dir = clone
		c.Env = cleanGitEnv()
		if out, err := c.CombinedOutput(); err != nil {
			both()
			return "", nil, fmt.Errorf("git %s failed for %s: %s", args[0], sr.Name, strings.TrimSpace(string(out)))
		}
	}
	return clone, both, nil
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-f807707a","gen":"sha256:3f6d1802d83ba9dcc72ed02c8927ae3f79e56242263a8c87803b425321b88f73","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func syncShared(r *repo, sr sharedRepo) (int, error) {\n\tref := sr.Ref\n\tif ref == \"\" {\n\t\tref = \"main\"\n\t}\n\tdir := sr.Dir\n\tif dir == \"\" {\n\t\tdir = \"provenance\"\n\t}\n\troot, cleanup, err := fetchShared(sr, ref, dir)\n\tif err != nil {\n\t\treturn 0, err\n\t}\n\tdefer cleanup()\n\tentries, err := os.ReadDir(filepath.Join(root, filepath.FromSlash(dir)))\n\tif err != nil && !os.IsNotExist(err) {\n\t\treturn 0, err\n\t}\n\tprovDir := sharedCacheDir(sr)\n\tif err := os.MkdirAll(provDir, 0o755); err != nil {\n\t\treturn 0, err\n\t}\n\told, _ := os.ReadDir(provDir)\n\tfor _, e := range old {\n\t\tif !e.IsDir() {\n\t\t\tos.Remove(filepath.Join(provDir, e.Name()))\n\t\t}\n\t}\n\tcount := 0\n\tfor _, e := range entries {\n\t\tname := e.Name()\n\t\tif e.IsDir() || !(strings.HasSuffix(name, \".yml\") || strings.HasSuffix(name, \".yaml\")) {\n\t\t\tcontinue\n\t\t}\n\t\tdata, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(dir), name))\n\t\tif err != nil {\n\t\t\treturn count, err\n\t\t}\n\t\tif err := os.WriteFile(filepath.Join(provDir, name), data, 0o644); err != nil {\n\t\t\treturn count, err\n\t\t}\n\t\tcount++\n\t}\n\tmeta, err := json.Marshal(cacheMeta{FetchedAt: time.Now().UTC(), RecordCount: count, RepoName: sr.Name, RepoURL: sr.URL, RepoRef: ref})\n\tif err != nil {\n\t\treturn count, err\n\t}\n\treturn count, os.WriteFile(filepath.Join(sharedCacheRoot(sr), \".meta\"), meta, 0o644)\n}","file":"cmd/provenance/shared_run.go","name":"syncShared"}}
func syncShared(r *repo, sr sharedRepo) (int, error) {
	ref := sr.Ref
	if ref == "" {
		ref = "main"
	}
	dir := sr.Dir
	if dir == "" {
		dir = "provenance"
	}
	root, cleanup, err := fetchShared(sr, ref, dir)
	if err != nil {
		return 0, err
	}
	defer cleanup()
	entries, err := os.ReadDir(filepath.Join(root, filepath.FromSlash(dir)))
	if err != nil && !os.IsNotExist(err) {
		return 0, err
	}
	provDir := sharedCacheDir(sr)
	if err := os.MkdirAll(provDir, 0o755); err != nil {
		return 0, err
	}
	old, _ := os.ReadDir(provDir)
	for _, e := range old {
		if !e.IsDir() {
			os.Remove(filepath.Join(provDir, e.Name()))
		}
	}
	count := 0
	for _, e := range entries {
		name := e.Name()
		if e.IsDir() || !(strings.HasSuffix(name, ".yml") || strings.HasSuffix(name, ".yaml")) {
			continue
		}
		data, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(dir), name))
		if err != nil {
			return count, err
		}
		if err := os.WriteFile(filepath.Join(provDir, name), data, 0o644); err != nil {
			return count, err
		}
		count++
	}
	meta, err := json.Marshal(cacheMeta{FetchedAt: time.Now().UTC(), RecordCount: count, RepoName: sr.Name, RepoURL: sr.URL, RepoRef: ref})
	if err != nil {
		return count, err
	}
	return count, os.WriteFile(filepath.Join(sharedCacheRoot(sr), ".meta"), meta, 0o644)
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-f807707a","gen":"sha256:3f6d1802d83ba9dcc72ed02c8927ae3f79e56242263a8c87803b425321b88f73","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func runSync(out io.Writer, cmd *cobra.Command, args []string) error {\n\tr, err := openRepo(cmd)\n\tif err != nil {\n\t\treturn err\n\t}\n\tforce, _ := cmd.Flags().GetBool(\"force\")\n\tif len(r.Config.SharedRepos) == 0 {\n\t\tfmt.Fprintln(out, \"No shared_repos configured\")\n\t\treturn nil\n\t}\n\tfailed := false\n\tfor _, sr := range r.Config.SharedRepos {\n\t\tif !force && cacheFresh(r, sr) {\n\t\t\tfmt.Fprintf(out, \"✓ %s (%s) — cache is fresh, skipped\\n\", sr.Name, sr.URL)\n\t\t\tcontinue\n\t\t}\n\t\tcount, err := syncShared(r, sr)\n\t\tif err != nil {\n\t\t\tfmt.Fprintf(out, \"✗ %s (%s): %v\\n\", sr.Name, sr.URL, err)\n\t\t\tfailed = true\n\t\t\tcontinue\n\t\t}\n\t\tfmt.Fprintf(out, \"✓ Synced %s (%s) — %d records\\n\", sr.Name, sr.URL, count)\n\t}\n\tif failed {\n\t\treturn fmt.Errorf(\"one or more repos failed to sync\")\n\t}\n\treturn nil\n}","file":"cmd/provenance/shared_run.go","name":"runSync"}}
func runSync(out io.Writer, cmd *cobra.Command, args []string) error {
	r, err := openRepo(cmd)
	if err != nil {
		return err
	}
	force, _ := cmd.Flags().GetBool("force")
	if len(r.Config.SharedRepos) == 0 {
		fmt.Fprintln(out, "No shared_repos configured")
		return nil
	}
	failed := false
	for _, sr := range r.Config.SharedRepos {
		if !force && cacheFresh(r, sr) {
			fmt.Fprintf(out, "✓ %s (%s) — cache is fresh, skipped\n", sr.Name, sr.URL)
			continue
		}
		count, err := syncShared(r, sr)
		if err != nil {
			fmt.Fprintf(out, "✗ %s (%s): %v\n", sr.Name, sr.URL, err)
			failed = true
			continue
		}
		fmt.Fprintf(out, "✓ Synced %s (%s) — %d records\n", sr.Name, sr.URL, count)
	}
	if failed {
		return fmt.Errorf("one or more repos failed to sync")
	}
	return nil
}

// /sedum:setFunction
