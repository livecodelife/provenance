package main

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-d84dc10e","gen":"sha256:2e97312cea5e86a4f4e11ea65cacfe9c4a63da245a1c446b00a2c6740234e92a","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/merge_manifest_run.go","path":"errors"}}
import "errors"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-d84dc10e","gen":"sha256:2e97312cea5e86a4f4e11ea65cacfe9c4a63da245a1c446b00a2c6740234e92a","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/merge_manifest_run.go","path":"fmt"}}
import "fmt"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-d84dc10e","gen":"sha256:2e97312cea5e86a4f4e11ea65cacfe9c4a63da245a1c446b00a2c6740234e92a","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/merge_manifest_run.go","path":"io"}}
import "io"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-d84dc10e","gen":"sha256:2e97312cea5e86a4f4e11ea65cacfe9c4a63da245a1c446b00a2c6740234e92a","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/merge_manifest_run.go","path":"os"}}
import "os"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-d84dc10e","gen":"sha256:2e97312cea5e86a4f4e11ea65cacfe9c4a63da245a1c446b00a2c6740234e92a","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/merge_manifest_run.go","path":"sort"}}
import "sort"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-d84dc10e","gen":"sha256:2e97312cea5e86a4f4e11ea65cacfe9c4a63da245a1c446b00a2c6740234e92a","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/merge_manifest_run.go","path":"github.com/spf13/cobra"}}
import "github.com/spf13/cobra"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-d84dc10e","gen":"sha256:2e97312cea5e86a4f4e11ea65cacfe9c4a63da245a1c446b00a2c6740234e92a","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/merge_manifest_run.go","path":"github.com/livecodelife/provenance/record"}}
import "github.com/livecodelife/provenance/record"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"tier":"owned","record":"prov-2026-d84dc10e","gen":"sha256:2e97312cea5e86a4f4e11ea65cacfe9c4a63da245a1c446b00a2c6740234e92a","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"// describeManifestLine words one side of a conflicting manifest entry.\nfunc describeManifestLine(l record.ManifestLine, present bool) string {\n\tif !present {\n\t\treturn \"removed\"\n\t}\n\treturn fmt.Sprintf(\"version %d, %s\", l.Version, l.ContentHash)\n}","file":"cmd/provenance/merge_manifest_run.go","name":"describeManifestLine"}}
// describeManifestLine words one side of a conflicting manifest entry.
func describeManifestLine(l record.ManifestLine, present bool) string {
	if !present {
		return "removed"
	}
	return fmt.Sprintf("version %d, %s", l.Version, l.ContentHash)
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-d84dc10e","gen":"sha256:2e97312cea5e86a4f4e11ea65cacfe9c4a63da245a1c446b00a2c6740234e92a","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"// mergeManifests merges two manifests that both descend from base, one line\n// per record id. A record changed on one side takes that side's line; changed\n// on both, the higher version wins, and the same version with different content\n// (or a removal against a change) is a conflict.\nfunc mergeManifests(base, ours, theirs []record.ManifestLine) ([]record.ManifestLine, error) {\n\tindex := func(ls []record.ManifestLine) map[string]record.ManifestLine {\n\t\tm := make(map[string]record.ManifestLine, len(ls))\n\t\tfor _, l := range ls {\n\t\t\tm[l.ID] = l\n\t\t}\n\t\treturn m\n\t}\n\tb, o, t := index(base), index(ours), index(theirs)\n\tids := map[string]bool{}\n\tfor _, m := range []map[string]record.ManifestLine{b, o, t} {\n\t\tfor id := range m {\n\t\t\tids[id] = true\n\t\t}\n\t}\n\tvar out []record.ManifestLine\n\tfor id := range ids {\n\t\tbl, bok := b[id]\n\t\tol, ook := o[id]\n\t\ttl, tok := t[id]\n\t\tvar pick record.ManifestLine\n\t\tvar ok bool\n\t\tswitch {\n\t\tcase ook == tok && ol == tl:\n\t\t\tpick, ok = ol, ook\n\t\tcase ook == bok && ol == bl:\n\t\t\tpick, ok = tl, tok\n\t\tcase tok == bok && tl == bl:\n\t\t\tpick, ok = ol, ook\n\t\tcase ook && tok && ol.Version != tl.Version:\n\t\t\tpick, ok = ol, true\n\t\t\tif tl.Version > ol.Version {\n\t\t\t\tpick = tl\n\t\t\t}\n\t\tdefault:\n\t\t\treturn nil, fmt.Errorf(\"%s: both branches changed it (ours: %s; theirs: %s); resolve the record files, then run provenance compile\", id, describeManifestLine(ol, ook), describeManifestLine(tl, tok))\n\t\t}\n\t\tif ok {\n\t\t\tout = append(out, pick)\n\t\t}\n\t}\n\tsort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })\n\treturn out, nil\n}","file":"cmd/provenance/merge_manifest_run.go","name":"mergeManifests"}}
// mergeManifests merges two manifests that both descend from base, one line
// per record id. A record changed on one side takes that side's line; changed
// on both, the higher version wins, and the same version with different content
// (or a removal against a change) is a conflict.
func mergeManifests(base, ours, theirs []record.ManifestLine) ([]record.ManifestLine, error) {
	index := func(ls []record.ManifestLine) map[string]record.ManifestLine {
		m := make(map[string]record.ManifestLine, len(ls))
		for _, l := range ls {
			m[l.ID] = l
		}
		return m
	}
	b, o, t := index(base), index(ours), index(theirs)
	ids := map[string]bool{}
	for _, m := range []map[string]record.ManifestLine{b, o, t} {
		for id := range m {
			ids[id] = true
		}
	}
	var out []record.ManifestLine
	for id := range ids {
		bl, bok := b[id]
		ol, ook := o[id]
		tl, tok := t[id]
		var pick record.ManifestLine
		var ok bool
		switch {
		case ook == tok && ol == tl:
			pick, ok = ol, ook
		case ook == bok && ol == bl:
			pick, ok = tl, tok
		case tok == bok && tl == bl:
			pick, ok = ol, ook
		case ook && tok && ol.Version != tl.Version:
			pick, ok = ol, true
			if tl.Version > ol.Version {
				pick = tl
			}
		default:
			return nil, fmt.Errorf("%s: both branches changed it (ours: %s; theirs: %s); resolve the record files, then run provenance compile", id, describeManifestLine(ol, ook), describeManifestLine(tl, tok))
		}
		if ok {
			out = append(out, pick)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })
	return out, nil
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-d84dc10e","gen":"sha256:2e97312cea5e86a4f4e11ea65cacfe9c4a63da245a1c446b00a2c6740234e92a","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"// runMergeManifest is git's merge driver for the version manifest: args are\n// the base, ours and theirs files, and the merged manifest replaces ours.\nfunc runMergeManifest(out io.Writer, cmd *cobra.Command, args []string) error {\n\tif len(args) != 3 {\n\t\treturn errors.New(\"usage: provenance merge-manifest BASE OURS THEIRS\")\n\t}\n\tvar sets [3][]record.ManifestLine\n\tfor i, p := range args {\n\t\tdata, err := os.ReadFile(p)\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n\t\tlines, err := record.ParseManifest(data)\n\t\tif err != nil {\n\t\t\treturn fmt.Errorf(\"%s: %w\", p, err)\n\t\t}\n\t\tsets[i] = lines\n\t}\n\tmerged, err := mergeManifests(sets[0], sets[1], sets[2])\n\tif err != nil {\n\t\treturn err\n\t}\n\treturn os.WriteFile(args[1], record.MarshalManifest(merged), 0o644)\n}","file":"cmd/provenance/merge_manifest_run.go","name":"runMergeManifest"}}
// runMergeManifest is git's merge driver for the version manifest: args are
// the base, ours and theirs files, and the merged manifest replaces ours.
func runMergeManifest(out io.Writer, cmd *cobra.Command, args []string) error {
	if len(args) != 3 {
		return errors.New("usage: provenance merge-manifest BASE OURS THEIRS")
	}
	var sets [3][]record.ManifestLine
	for i, p := range args {
		data, err := os.ReadFile(p)
		if err != nil {
			return err
		}
		lines, err := record.ParseManifest(data)
		if err != nil {
			return fmt.Errorf("%s: %w", p, err)
		}
		sets[i] = lines
	}
	merged, err := mergeManifests(sets[0], sets[1], sets[2])
	if err != nil {
		return err
	}
	return os.WriteFile(args[1], record.MarshalManifest(merged), 0o644)
}

// /sedum:setFunction
