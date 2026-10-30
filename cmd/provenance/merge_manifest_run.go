package main

// sedum:anchor:imports
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-d84dc10e","gen":"sha256:2e97312cea5e86a4f4e11ea65cacfe9c4a63da245a1c446b00a2c6740234e92a","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/merge_manifest_run.go","path":"errors"},"kwargs":"sha256:b09e72ab39bd2966400a77b37f4a9a1fa0e5d098f95fee4ec65d20ca1f309866","body":"sha256:2b718237ba0a56467c8fd494cc131e942ef410e6228c03b644d70476c71862dc"}
import "errors"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-d84dc10e","gen":"sha256:2e97312cea5e86a4f4e11ea65cacfe9c4a63da245a1c446b00a2c6740234e92a","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/merge_manifest_run.go","path":"fmt"},"kwargs":"sha256:bf7c199ea0699d8f859ab160f317d69454227340309f2312edc1b588b44200c5","body":"sha256:bcbac51883cbf8064e9196b371cb880cee429866652447ad557533d817857609"}
import "fmt"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-d84dc10e","gen":"sha256:2e97312cea5e86a4f4e11ea65cacfe9c4a63da245a1c446b00a2c6740234e92a","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/merge_manifest_run.go","path":"io"},"kwargs":"sha256:2d9dd953c94449b3687393ee961f0e51a3d1d9ff75cda7821c8940f1cdb0475f","body":"sha256:ef8e8abf4ef5ee40a19dc85b49bf09c5df1fe32abfe834e0ab4c2175baf80ddf"}
import "io"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-d84dc10e","gen":"sha256:2e97312cea5e86a4f4e11ea65cacfe9c4a63da245a1c446b00a2c6740234e92a","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/merge_manifest_run.go","path":"os"},"kwargs":"sha256:572671cc4827d231645987384b19aa92a5d55e4e60ccdefb0fd6431d94ec55cc","body":"sha256:7c6e309eb2575d4aa451ec2ab83161d13e9ae03bbca3a4d2ecf30ce12adf0905"}
import "os"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-d84dc10e","gen":"sha256:2e97312cea5e86a4f4e11ea65cacfe9c4a63da245a1c446b00a2c6740234e92a","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/merge_manifest_run.go","path":"sort"},"kwargs":"sha256:4662cf146eef627bbac810913e1a99439f9ea2e4d1ebf6eaef4e875311befee8","body":"sha256:f197172617ae21ef23b1855b09aafb14571020ae81bff877d63d3ff81f4bf63c"}
import "sort"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-d84dc10e","gen":"sha256:2e97312cea5e86a4f4e11ea65cacfe9c4a63da245a1c446b00a2c6740234e92a","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/merge_manifest_run.go","path":"github.com/spf13/cobra"},"kwargs":"sha256:6a9c7bf02b62bbdf273a8f1222c152b189f4b7a5fd534cb5be73eb9726d6aebd","body":"sha256:f361c02d0f8e1999cf3fb715393e16310eeec9debd5aaa13936af7617e2b46aa"}
import "github.com/spf13/cobra"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-d84dc10e","gen":"sha256:2e97312cea5e86a4f4e11ea65cacfe9c4a63da245a1c446b00a2c6740234e92a","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/merge_manifest_run.go","path":"github.com/livecodelife/provenance/record"},"kwargs":"sha256:3082c47ff627011d36b4a6369ba1c4f2224865432f6a7f005b8f2b8fce7c2330","body":"sha256:acbf8628061291d11aa39e9ea64ecc91095a117057cfe6f1958ef2c437fac003"}
import "github.com/livecodelife/provenance/record"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-d84dc10e","gen":"sha256:2e97312cea5e86a4f4e11ea65cacfe9c4a63da245a1c446b00a2c6740234e92a","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/merge_manifest_run.go","name":"describeManifestLine"},"kwargs":"sha256:af834c51b6b0c72a53f0c4a123a4d3b62a593fda230f2525250f233a4f0db712","body":"sha256:b74898961fffec13d275f04775726c432c7208e5791026960834dde7f9df3ac1"}
// describeManifestLine words one side of a conflicting manifest entry.
func describeManifestLine(l record.ManifestLine, present bool) string {
	if !present {
		return "removed"
	}
	return fmt.Sprintf("version %d, %s", l.Version, l.ContentHash)
}

// /sedum:setFunction
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-d84dc10e","gen":"sha256:2e97312cea5e86a4f4e11ea65cacfe9c4a63da245a1c446b00a2c6740234e92a","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/merge_manifest_run.go","name":"mergeManifests"},"kwargs":"sha256:a2bb379f50f05b99f9140029170e60f1013c6cc746ad947f55e55a4061b14aff","body":"sha256:a1894a0f866482f6b1d27e1e43fb516fc5999a48cfb5246e8464c5f88b1d4233"}
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
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-d84dc10e","gen":"sha256:2e97312cea5e86a4f4e11ea65cacfe9c4a63da245a1c446b00a2c6740234e92a","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/merge_manifest_run.go","name":"runMergeManifest"},"kwargs":"sha256:cc92e592dde1ea882e41abab3467f00b0d6e689db2bce6936fe2212b4550d08d","body":"sha256:eb421e719f90839ebcdb76dbc0fa5e7d37f3e28030aa9fb780c79820e9c8b9c1"}
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
