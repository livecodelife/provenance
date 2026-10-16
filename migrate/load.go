// Package migrate plans and applies the v1 -> v2 record migration settled in
// linespec-migration-dry-run/docs/MIGRATION_DRY_RUN.md: supersession chains
// and bugs become versions of one record, deprecated bugs and kept-separate
// edges stay separate, and children re-pin to the version they implemented.
//
// Plan is read-only. Apply writes only through a record.Store.
package migrate

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"

	"github.com/livecodelife/provenance/record"
)

// SharedRepo is one entry of the v1 provenance block's shared_repos.
type SharedRepo struct {
	Name string `yaml:"name"`
	URL  string `yaml:"url"`
	Ref  string `yaml:"ref"`
	Dir  string `yaml:"dir"`
}

// Config is the part of a v1 .linespec.yml provenance block migration reads.
type Config struct {
	Dir         string       `yaml:"dir"`
	SharedRepos []SharedRepo `yaml:"shared_repos"`
}

// Source is one record as loaded for migration.
type Source struct {
	Rec  *record.Record
	Doc  *record.Document
	File string // path relative to the repo root
	// Added is when git first added the file; nil when git does not know.
	Added *GitTime

	// Set for schema 2 records (already migrated), read from their snapshots.
	// Latest is the newest snapshot's version, 0 when there are none (a v2
	// draft). LatestHash is that snapshot's content_hash. Drifted reports
	// that the record's content_hash no longer matches it (PROV024).
	Latest     int
	LatestHash string
	Drifted    bool
}

// V2 reports a record that is already migrated.
func (s *Source) V2() bool { return s.Rec.Schema == 2 }

// Type returns the record's type with the v1 default applied.
func (s *Source) Type() record.Type { return s.Rec.EffectiveType() }

// Repo is a repository's records loaded for migration.
type Repo struct {
	Name   string
	Root   string
	Dir    string // records directory relative to Root
	Config Config
	// Records are the local records by id.
	Records map[string]*Source
	// Shared are records from local checkouts of shared_repos, by id, used
	// only to resolve references. SharedRepoOf names the repo each came from.
	Shared       map[string]*Source
	SharedRepoOf map[string]string
}

// LoadOptions controls where shared repos are looked up and how git is read.
type LoadOptions struct {
	// Workspace holds local checkouts of shared repos, named by the last
	// segment of their URL. Defaults to the repo root's parent directory.
	Workspace string
	// Git reads first-add times; nil uses the git CLI.
	Git Git
	// Config, when set, is used instead of reading .provenance.yml or
	// .linespec.yml.
	Config *Config
}

// Load reads a v1 (or already migrated) repository. It never writes.
func Load(root string, opts LoadOptions) (*Repo, error) {
	var cfg Config
	if opts.Config != nil {
		cfg = *opts.Config
	} else {
		var err error
		if cfg, err = readConfig(root); err != nil {
			return nil, err
		}
	}
	if opts.Workspace == "" {
		opts.Workspace = filepath.Dir(root)
	}
	if opts.Git == nil {
		opts.Git = GitCLI{}
	}
	dir := strings.TrimRight(cfg.Dir, "/")
	if dir == "" {
		dir = "provenance"
	}
	r := &Repo{Name: filepath.Base(root), Root: root, Dir: dir, Config: cfg, Shared: map[string]*Source{}, SharedRepoOf: map[string]string{}}
	recs, err := loadDir(filepath.Join(root, dir), root)
	if err != nil {
		return nil, err
	}
	r.Records = recs
	if err := r.readSnapshots(); err != nil {
		return nil, err
	}
	added, err := opts.Git.FirstAdded(root, dir)
	if err != nil {
		return nil, err
	}
	for _, s := range r.Records {
		if t, ok := added[filepath.Base(s.File)]; ok {
			t := t
			s.Added = &t
		} else if t, ok := opts.Git.FollowFirst(root, filepath.ToSlash(filepath.Join(dir, filepath.Base(s.File)))); ok {
			s.Added = &t
		}
	}
	for _, sr := range cfg.SharedRepos {
		local := filepath.Join(opts.Workspace, strings.TrimSuffix(lastSegment(sr.URL), ".git"))
		sdir := sr.Dir
		if sdir == "" {
			sdir = "provenance"
		}
		d := filepath.Join(local, sdir)
		if st, err := os.Stat(d); err != nil || !st.IsDir() {
			continue
		}
		recs, err := loadDir(d, d)
		if err != nil {
			return nil, fmt.Errorf("shared repo %s: %w", sr.Name, err)
		}
		for id, s := range recs {
			r.Shared[id] = s
			r.SharedRepoOf[id] = sr.Name
		}
	}
	return r, nil
}

func lastSegment(url string) string {
	url = strings.TrimRight(url, "/")
	if i := strings.LastIndexByte(url, '/'); i >= 0 {
		return url[i+1:]
	}
	return url
}

// readConfig reads the repository's provenance config: .provenance.yml (the
// config itself, written by `provenance migrate`) first, else the provenance
// block of .linespec.yml. Neither is an error naming both.
func readConfig(root string) (Config, error) {
	if data, err := os.ReadFile(filepath.Join(root, ".provenance.yml")); err == nil {
		var c Config
		if err := yaml.Unmarshal(data, &c); err != nil {
			return Config{}, fmt.Errorf(".provenance.yml: %w", err)
		}
		return c, nil
	} else if !os.IsNotExist(err) {
		return Config{}, err
	}
	data, err := os.ReadFile(filepath.Join(root, ".linespec.yml"))
	if os.IsNotExist(err) {
		return Config{}, fmt.Errorf("no provenance config: neither .provenance.yml nor .linespec.yml exists in %s", root)
	}
	if err != nil {
		return Config{}, err
	}
	var c struct {
		Provenance *Config `yaml:"provenance"`
	}
	if err := yaml.Unmarshal(data, &c); err != nil {
		return Config{}, fmt.Errorf(".linespec.yml: %w", err)
	}
	if c.Provenance == nil {
		return Config{}, fmt.Errorf(".linespec.yml has no provenance block (and there is no .provenance.yml)")
	}
	return *c.Provenance, nil
}

// readSnapshots fills in the snapshot facts of every v2 record.
func (r *Repo) readSnapshots() error {
	store := &record.FilesStore{Root: r.Root, Dir: r.Dir}
	specs := record.FileSpecResolver(r.Root)
	for id, s := range r.Records {
		if !s.V2() {
			continue
		}
		versions, err := store.Versions(id)
		if err != nil {
			return err
		}
		if len(versions) == 0 {
			continue
		}
		s.Latest = versions[len(versions)-1]
		snap, err := store.Snapshot(id, s.Latest)
		if err != nil {
			return err
		}
		s.LatestHash = snap.ContentHash
		h, _ := s.Rec.ContentHash(specs)
		s.Drifted = h != snap.ContentHash || s.Rec.Version != s.Latest
	}
	return nil
}

// loadDir reads every *.yml and *.yaml record in dir; File is relative to rel.
func loadDir(dir, rel string) (map[string]*Source, error) {
	var files []string
	for _, pat := range []string{"*.yml", "*.yaml"} {
		m, err := filepath.Glob(filepath.Join(dir, pat))
		if err != nil {
			return nil, err
		}
		files = append(files, m...)
	}
	sort.Strings(files)
	out := map[string]*Source{}
	for _, f := range files {
		doc, err := record.Load(f)
		if err != nil {
			return nil, err
		}
		rec := doc.Record()
		if rec.ID == "" {
			return nil, fmt.Errorf("%s: no id", f)
		}
		if prev, dup := out[rec.ID]; dup {
			return nil, fmt.Errorf("%s: duplicate id %s (also %s)", f, rec.ID, prev.File)
		}
		relPath, _ := filepath.Rel(rel, f)
		out[rec.ID] = &Source{Rec: rec, Doc: doc, File: filepath.ToSlash(relPath)}
	}
	return out, nil
}

// ref returns where a referenced id resolves: "local", "shared" or "".
func (r *Repo) ref(id string) string {
	if _, ok := r.Records[id]; ok {
		return "local"
	}
	if _, ok := r.Shared[id]; ok {
		return "shared"
	}
	return ""
}
