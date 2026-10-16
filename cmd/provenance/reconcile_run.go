package main

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-9a3aad45","gen":"sha256:d19012536182a12da7b036a73f8eea6a7a589bd985a0ee0352375454ff233111","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/reconcile_run.go","path":"fmt"}}
import "fmt"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-9a3aad45","gen":"sha256:d19012536182a12da7b036a73f8eea6a7a589bd985a0ee0352375454ff233111","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/reconcile_run.go","path":"io"}}
import "io"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-9a3aad45","gen":"sha256:d19012536182a12da7b036a73f8eea6a7a589bd985a0ee0352375454ff233111","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/reconcile_run.go","path":"os"}}
import "os"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-9a3aad45","gen":"sha256:d19012536182a12da7b036a73f8eea6a7a589bd985a0ee0352375454ff233111","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/reconcile_run.go","path":"path/filepath"}}
import "path/filepath"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-9a3aad45","gen":"sha256:d19012536182a12da7b036a73f8eea6a7a589bd985a0ee0352375454ff233111","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/reconcile_run.go","path":"sort"}}
import "sort"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-9a3aad45","gen":"sha256:d19012536182a12da7b036a73f8eea6a7a589bd985a0ee0352375454ff233111","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/reconcile_run.go","path":"strings"}}
import "strings"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-9a3aad45","gen":"sha256:d19012536182a12da7b036a73f8eea6a7a589bd985a0ee0352375454ff233111","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/reconcile_run.go","path":"github.com/spf13/cobra"}}
import "github.com/spf13/cobra"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-9a3aad45","gen":"sha256:d19012536182a12da7b036a73f8eea6a7a589bd985a0ee0352375454ff233111","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/reconcile_run.go","path":"github.com/livecodelife/provenance/record"}}
import "github.com/livecodelife/provenance/record"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-9a3aad45","gen":"sha256:d19012536182a12da7b036a73f8eea6a7a589bd985a0ee0352375454ff233111","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/reconcile_run.go","path":"github.com/livecodelife/provenance/scope"}}
import "github.com/livecodelife/provenance/scope"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setType {"tier":"owned","record":"prov-2026-9a3aad45","gen":"sha256:d19012536182a12da7b036a73f8eea6a7a589bd985a0ee0352375454ff233111","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type reconcileResult struct {\n\tUnlocked []string\n\tLocked   []string\n}","file":"cmd/provenance/reconcile_run.go","name":"reconcileResult"}}
type reconcileResult struct {
	Unlocked []string
	Locked   []string
}

// /sedum:setType
// sedum:setFunction {"tier":"owned","record":"prov-2026-9a3aad45","gen":"sha256:d19012536182a12da7b036a73f8eea6a7a589bd985a0ee0352375454ff233111","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func isSnapshotFile(r *repo, rel string) bool {\n\treturn strings.HasSuffix(rel, \".yml\") && snapshotOwner(r, rel) != \"\"\n}","file":"cmd/provenance/reconcile_run.go","name":"isSnapshotFile"}}
func isSnapshotFile(r *repo, rel string) bool {
	return strings.HasSuffix(rel, ".yml") && snapshotOwner(r, rel) != ""
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-9a3aad45","gen":"sha256:d19012536182a12da7b036a73f8eea6a7a589bd985a0ee0352375454ff233111","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func alwaysWritable(r *repo, recs []*record.Record) []string {\n\tout := append([]string{}, r.Config.ExcludePaths...)\n\tout = append(out, filepath.ToSlash(r.Config.Dir), \".provenance\", configFile, legacyConfigFile)\n\tfor _, rec := range recs {\n\t\tfor _, spec := range rec.AssociatedSpecs {\n\t\t\tif spec.Path != \"\" {\n\t\t\t\tout = append(out, spec.Path)\n\t\t\t}\n\t\t}\n\t}\n\treturn out\n}","file":"cmd/provenance/reconcile_run.go","name":"alwaysWritable"}}
func alwaysWritable(r *repo, recs []*record.Record) []string {
	out := append([]string{}, r.Config.ExcludePaths...)
	out = append(out, filepath.ToSlash(r.Config.Dir), ".provenance", configFile, legacyConfigFile)
	for _, rec := range recs {
		for _, spec := range rec.AssociatedSpecs {
			if spec.Path != "" {
				out = append(out, spec.Path)
			}
		}
	}
	return out
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-9a3aad45","gen":"sha256:d19012536182a12da7b036a73f8eea6a7a589bd985a0ee0352375454ff233111","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func openScope(recs []*record.Record) []string {\n\tseen := map[string]bool{}\n\tvar out []string\n\tfor _, rec := range recs {\n\t\tif rec.Status != record.StatusOpen || scope.Mode(rec) != \"allowlist\" {\n\t\t\tcontinue\n\t\t}\n\t\tfor _, p := range rec.AffectedScope {\n\t\t\tif p != \"\" && !seen[p] {\n\t\t\t\tseen[p] = true\n\t\t\t\tout = append(out, p)\n\t\t\t}\n\t\t}\n\t}\n\tsort.Strings(out)\n\treturn out\n}","file":"cmd/provenance/reconcile_run.go","name":"openScope"}}
func openScope(recs []*record.Record) []string {
	seen := map[string]bool{}
	var out []string
	for _, rec := range recs {
		if rec.Status != record.StatusOpen || scope.Mode(rec) != "allowlist" {
			continue
		}
		for _, p := range rec.AffectedScope {
			if p != "" && !seen[p] {
				seen[p] = true
				out = append(out, p)
			}
		}
	}
	sort.Strings(out)
	return out
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-9a3aad45","gen":"sha256:d19012536182a12da7b036a73f8eea6a7a589bd985a0ee0352375454ff233111","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func writablePath(r *repo, rel string, always, open []string, restricted bool) bool {\n\tif isSnapshotFile(r, rel) {\n\t\treturn false\n\t}\n\tif !restricted || scope.Excluded(rel, always) {\n\t\treturn true\n\t}\n\tfor _, p := range open {\n\t\tif ok, err := scope.Match(rel, p); err == nil && ok {\n\t\t\treturn true\n\t\t}\n\t}\n\treturn false\n}","file":"cmd/provenance/reconcile_run.go","name":"writablePath"}}
func writablePath(r *repo, rel string, always, open []string, restricted bool) bool {
	if isSnapshotFile(r, rel) {
		return false
	}
	if !restricted || scope.Excluded(rel, always) {
		return true
	}
	for _, p := range open {
		if ok, err := scope.Match(rel, p); err == nil && ok {
			return true
		}
	}
	return false
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-9a3aad45","gen":"sha256:d19012536182a12da7b036a73f8eea6a7a589bd985a0ee0352375454ff233111","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func patternBaseDir(pattern string) string {\n\tbody := pattern\n\tmeta := \"*?\"\n\tif scope.IsRegex(pattern) {\n\t\tbody = strings.TrimPrefix(pattern[3:], \"^\")\n\t\tmeta = `.+*?()[]{}^$|\\`\n\t}\n\tsegments := strings.Split(body, \"/\")\n\tvar literal []string\n\tfor _, seg := range segments[:len(segments)-1] {\n\t\tif strings.ContainsAny(seg, meta) {\n\t\t\tbreak\n\t\t}\n\t\tliteral = append(literal, seg)\n\t}\n\treturn strings.Join(literal, \"/\")\n}","file":"cmd/provenance/reconcile_run.go","name":"patternBaseDir"}}
func patternBaseDir(pattern string) string {
	body := pattern
	meta := "*?"
	if scope.IsRegex(pattern) {
		body = strings.TrimPrefix(pattern[3:], "^")
		meta = `.+*?()[]{}^$|\`
	}
	segments := strings.Split(body, "/")
	var literal []string
	for _, seg := range segments[:len(segments)-1] {
		if strings.ContainsAny(seg, meta) {
			break
		}
		literal = append(literal, seg)
	}
	return strings.Join(literal, "/")
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-9a3aad45","gen":"sha256:d19012536182a12da7b036a73f8eea6a7a589bd985a0ee0352375454ff233111","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func reconcileFiles(r *repo, dry bool) (*reconcileResult, error) {\n\tres := &reconcileResult{}\n\trestricted := r.Config.WriteRestriction == nil || *r.Config.WriteRestriction\n\tids, err := r.Store.List()\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\tvar recs []*record.Record\n\tfor _, id := range ids {\n\t\trec, _, err := r.Store.Load(id)\n\t\tif err != nil {\n\t\t\treturn nil, err\n\t\t}\n\t\trecs = append(recs, rec)\n\t}\n\tif restricted && len(recs) == 0 {\n\t\treturn res, nil\n\t}\n\tlisting, err := r.git(\"ls-files\")\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\talways, open := alwaysWritable(r, recs), openScope(recs)\n\tdirs := map[string]bool{}\n\tfor _, rel := range strings.Split(listing, \"\\n\") {\n\t\tif rel == \"\" {\n\t\t\tcontinue\n\t\t}\n\t\twritable := writablePath(r, rel, always, open, restricted)\n\t\tif writable {\n\t\t\tdirs[filepath.Dir(rel)] = true\n\t\t}\n\t\tabs := filepath.Join(r.Root, rel)\n\t\tinfo, err := os.Lstat(abs)\n\t\tif err != nil || info.IsDir() {\n\t\t\tcontinue\n\t\t}\n\t\tcur := info.Mode().Perm()&0o200 != 0\n\t\tswitch {\n\t\tcase writable && !cur:\n\t\t\tif !dry {\n\t\t\t\tif err := os.Chmod(abs, info.Mode().Perm()|0o200); err != nil {\n\t\t\t\t\treturn nil, err\n\t\t\t\t}\n\t\t\t}\n\t\t\tres.Unlocked = append(res.Unlocked, rel)\n\t\tcase !writable && cur:\n\t\t\tif !dry {\n\t\t\t\tif err := os.Chmod(abs, info.Mode().Perm()&^0o222); err != nil {\n\t\t\t\t\treturn nil, err\n\t\t\t\t}\n\t\t\t}\n\t\t\tres.Locked = append(res.Locked, rel)\n\t\t}\n\t}\n\tif restricted {\n\t\tfor _, p := range open {\n\t\t\tif scope.IsPattern(p) {\n\t\t\t\tif d := patternBaseDir(p); d != \"\" {\n\t\t\t\t\tdirs[d] = true\n\t\t\t\t}\n\t\t\t\tcontinue\n\t\t\t}\n\t\t\tif _, err := os.Lstat(filepath.Join(r.Root, p)); os.IsNotExist(err) {\n\t\t\t\tdirs[filepath.Dir(p)] = true\n\t\t\t}\n\t\t}\n\t}\n\tfor d := range dirs {\n\t\tif d == \".\" || d == \"\" || dry {\n\t\t\tcontinue\n\t\t}\n\t\tabs := filepath.Join(r.Root, d)\n\t\tif info, err := os.Lstat(abs); err == nil && info.IsDir() {\n\t\t\tif err := os.Chmod(abs, info.Mode().Perm()|0o300); err != nil {\n\t\t\t\treturn nil, err\n\t\t\t}\n\t\t}\n\t}\n\tsort.Strings(res.Unlocked)\n\tsort.Strings(res.Locked)\n\treturn res, nil\n}","file":"cmd/provenance/reconcile_run.go","name":"reconcileFiles"}}
func reconcileFiles(r *repo, dry bool) (*reconcileResult, error) {
	res := &reconcileResult{}
	restricted := r.Config.WriteRestriction == nil || *r.Config.WriteRestriction
	ids, err := r.Store.List()
	if err != nil {
		return nil, err
	}
	var recs []*record.Record
	for _, id := range ids {
		rec, _, err := r.Store.Load(id)
		if err != nil {
			return nil, err
		}
		recs = append(recs, rec)
	}
	if restricted && len(recs) == 0 {
		return res, nil
	}
	listing, err := r.git("ls-files")
	if err != nil {
		return nil, err
	}
	always, open := alwaysWritable(r, recs), openScope(recs)
	dirs := map[string]bool{}
	for _, rel := range strings.Split(listing, "\n") {
		if rel == "" {
			continue
		}
		writable := writablePath(r, rel, always, open, restricted)
		if writable {
			dirs[filepath.Dir(rel)] = true
		}
		abs := filepath.Join(r.Root, rel)
		info, err := os.Lstat(abs)
		if err != nil || info.IsDir() {
			continue
		}
		cur := info.Mode().Perm()&0o200 != 0
		switch {
		case writable && !cur:
			if !dry {
				if err := os.Chmod(abs, info.Mode().Perm()|0o200); err != nil {
					return nil, err
				}
			}
			res.Unlocked = append(res.Unlocked, rel)
		case !writable && cur:
			if !dry {
				if err := os.Chmod(abs, info.Mode().Perm()&^0o222); err != nil {
					return nil, err
				}
			}
			res.Locked = append(res.Locked, rel)
		}
	}
	if restricted {
		for _, p := range open {
			if scope.IsPattern(p) {
				if d := patternBaseDir(p); d != "" {
					dirs[d] = true
				}
				continue
			}
			if _, err := os.Lstat(filepath.Join(r.Root, p)); os.IsNotExist(err) {
				dirs[filepath.Dir(p)] = true
			}
		}
	}
	for d := range dirs {
		if d == "." || d == "" || dry {
			continue
		}
		abs := filepath.Join(r.Root, d)
		if info, err := os.Lstat(abs); err == nil && info.IsDir() {
			if err := os.Chmod(abs, info.Mode().Perm()|0o300); err != nil {
				return nil, err
			}
		}
	}
	sort.Strings(res.Unlocked)
	sort.Strings(res.Locked)
	return res, nil
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-9a3aad45","gen":"sha256:d19012536182a12da7b036a73f8eea6a7a589bd985a0ee0352375454ff233111","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func reconcileQuietly(r *repo) {\n\tif _, err := reconcileFiles(r, false); err != nil {\n\t\treturn\n\t}\n}","file":"cmd/provenance/reconcile_run.go","name":"reconcileQuietly"}}
func reconcileQuietly(r *repo) {
	if _, err := reconcileFiles(r, false); err != nil {
		return
	}
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-9a3aad45","gen":"sha256:d19012536182a12da7b036a73f8eea6a7a589bd985a0ee0352375454ff233111","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func runReconcile(out io.Writer, cmd *cobra.Command, args []string) error {\n\tr, err := openRepo(cmd)\n\tif err != nil {\n\t\treturn err\n\t}\n\tdry, _ := cmd.Flags().GetBool(\"dry-run\")\n\tres, err := reconcileFiles(r, dry)\n\tif err != nil {\n\t\treturn err\n\t}\n\tverb := \"\"\n\tif dry {\n\t\tverb = \"Would \"\n\t}\n\tfmt.Fprintf(out, \"\\n✓ %sReconciled write permissions: %d unlocked, %d locked\\n\", verb, len(res.Unlocked), len(res.Locked))\n\tfor _, f := range res.Unlocked {\n\t\tfmt.Fprintf(out, \"    unlocked  %s\\n\", f)\n\t}\n\tfor _, f := range res.Locked {\n\t\tfmt.Fprintf(out, \"    locked    %s\\n\", f)\n\t}\n\tfmt.Fprint(out, \"\\n\")\n\treturn nil\n}","file":"cmd/provenance/reconcile_run.go","name":"runReconcile"}}
func runReconcile(out io.Writer, cmd *cobra.Command, args []string) error {
	r, err := openRepo(cmd)
	if err != nil {
		return err
	}
	dry, _ := cmd.Flags().GetBool("dry-run")
	res, err := reconcileFiles(r, dry)
	if err != nil {
		return err
	}
	verb := ""
	if dry {
		verb = "Would "
	}
	fmt.Fprintf(out, "\n✓ %sReconciled write permissions: %d unlocked, %d locked\n", verb, len(res.Unlocked), len(res.Locked))
	for _, f := range res.Unlocked {
		fmt.Fprintf(out, "    unlocked  %s\n", f)
	}
	for _, f := range res.Locked {
		fmt.Fprintf(out, "    locked    %s\n", f)
	}
	fmt.Fprint(out, "\n")
	return nil
}

// /sedum:setFunction
