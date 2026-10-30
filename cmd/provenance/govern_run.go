package main

// sedum:anchor:imports
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-9ec4a843","gen":"sha256:cc0301a10fcd903006e4f2bed32a32b6a082f192eb2c42bb2c80462870e50daf","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/govern_run.go","path":"encoding/json"},"kwargs":"sha256:23902abbab222d43b3b7b93a19632d0234ff3ad9fc4663757c205ac40a1f6789","body":"sha256:872d7c3e9dee608e79a7f8ef57930d24b5e30f2a57d891a52549179ace45946a"}
import "encoding/json"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-9ec4a843","gen":"sha256:cc0301a10fcd903006e4f2bed32a32b6a082f192eb2c42bb2c80462870e50daf","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/govern_run.go","path":"fmt"},"kwargs":"sha256:e2c6e7b45be9ee3d9984f4af3293fc60bf5928811d1808e5256b571f5b4a2cd9","body":"sha256:bcbac51883cbf8064e9196b371cb880cee429866652447ad557533d817857609"}
import "fmt"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-9ec4a843","gen":"sha256:cc0301a10fcd903006e4f2bed32a32b6a082f192eb2c42bb2c80462870e50daf","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/govern_run.go","path":"io"},"kwargs":"sha256:8d7e016ef6ac3bb20ead47e3c442082f96d26801fd3370940a7505d516bdf05d","body":"sha256:ef8e8abf4ef5ee40a19dc85b49bf09c5df1fe32abfe834e0ab4c2175baf80ddf"}
import "io"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-9ec4a843","gen":"sha256:cc0301a10fcd903006e4f2bed32a32b6a082f192eb2c42bb2c80462870e50daf","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/govern_run.go","path":"sort"},"kwargs":"sha256:c68f51f933b399bb261bda82e2474a4772d190c885b1ba39618b93ffec0bc63d","body":"sha256:f197172617ae21ef23b1855b09aafb14571020ae81bff877d63d3ff81f4bf63c"}
import "sort"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-9ec4a843","gen":"sha256:cc0301a10fcd903006e4f2bed32a32b6a082f192eb2c42bb2c80462870e50daf","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/govern_run.go","path":"strings"},"kwargs":"sha256:58cbf1c69f63c4c64144d1cb94c51ecbd0a3a02cd8b65989e162661761df92dd","body":"sha256:a9af3c0bc4bf9342df9d7e89a324523fdba7941a9ca4190270e68ff3c4a25ea6"}
import "strings"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-9ec4a843","gen":"sha256:cc0301a10fcd903006e4f2bed32a32b6a082f192eb2c42bb2c80462870e50daf","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/govern_run.go","path":"github.com/spf13/cobra"},"kwargs":"sha256:507a8c7bb5caf80af8a88a27b60acedec2b7af8eb512334577b534b4cc99d2fd","body":"sha256:f361c02d0f8e1999cf3fb715393e16310eeec9debd5aaa13936af7617e2b46aa"}
import "github.com/spf13/cobra"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-9ec4a843","gen":"sha256:cc0301a10fcd903006e4f2bed32a32b6a082f192eb2c42bb2c80462870e50daf","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/govern_run.go","path":"github.com/livecodelife/provenance/record"},"kwargs":"sha256:84289dfe866cc8a6063ff014fe0f13722d4b0c143f16f384d4ccc9fbfb5f18b8","body":"sha256:acbf8628061291d11aa39e9ea64ecc91095a117057cfe6f1958ef2c437fac003"}
import "github.com/livecodelife/provenance/record"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setType {"v":2,"tier":"owned","record":"prov-2026-9ec4a843","gen":"sha256:cc0301a10fcd903006e4f2bed32a32b6a082f192eb2c42bb2c80462870e50daf","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","id":{"file":"cmd/provenance/govern_run.go","name":"governRecord"},"kwargs":"sha256:52048997955ed8818487655bb04f11cca44a3de0830da75c07e7d465f8599d6f","body":"sha256:ec7a5bf1f65959c9989e986178e1ac8c48b575318daca7f60371188f75a96c66"}
type governRecord struct {
	ID      string   `json:"id"`
	Type    string   `json:"type"`
	Status  string   `json:"status"`
	Version int      `json:"version"`
	Title   string   `json:"title"`
	Files   []string `json:"files"`
}

// /sedum:setType
// sedum:setType {"v":2,"tier":"owned","record":"prov-2026-9ec4a843","gen":"sha256:cc0301a10fcd903006e4f2bed32a32b6a082f192eb2c42bb2c80462870e50daf","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","id":{"file":"cmd/provenance/govern_run.go","name":"governDoc"},"kwargs":"sha256:9919645db4de56ab9bfa95d1d3716ad18a4309cee22fbbd1c6afb9b51f761a71","body":"sha256:26a02dda6579b2100879e705b11e41eb099088e7a3e79c0807af642d31f5b6a4"}
type governDoc struct {
	Files     []string       `json:"files"`
	Exempt    []string       `json:"exempt"`
	Governing []governRecord `json:"governing"`
	Next      *nextAction    `json:"next,omitempty"`
}

// /sedum:setType
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-9ec4a843","gen":"sha256:cc0301a10fcd903006e4f2bed32a32b6a082f192eb2c42bb2c80462870e50daf","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/govern_run.go","name":"buildGovern"},"kwargs":"sha256:9c2174f81e67549f858c00dc2c00c94d0146f434649cd97908e551a91ecf6950","body":"sha256:5f199f467f1deec230bec009d458cab6513709faa8512ba9e0c5a83cde5ee460"}
func buildGovern(r *repo, files []string) (governDoc, error) {
	doc := governDoc{Files: files, Exempt: []string{}, Governing: []governRecord{}}
	ids, err := r.Store.List()
	if err != nil {
		return doc, err
	}
	var active []*record.Record
	for _, id := range ids {
		rec, _, err := r.Store.Load(id)
		if err != nil {
			return doc, err
		}
		if rec.Status == record.StatusOpen || rec.Status == record.StatusImplemented {
			active = append(active, rec)
		}
	}
	sort.Slice(active, func(i, j int) bool { return active[i].ID < active[j].ID })
	covers := map[string][]string{}
	for _, file := range files {
		path := strings.TrimPrefix(file, "./")
		if !nextRelevant(r, path) {
			doc.Exempt = append(doc.Exempt, file)
			continue
		}
		for _, rec := range active {
			if nextGoverns(rec, path) {
				covers[rec.ID] = append(covers[rec.ID], file)
			}
		}
	}
	for _, rec := range active {
		if got := covers[rec.ID]; len(got) > 0 {
			doc.Governing = append(doc.Governing, governRecord{ID: rec.ID, Type: string(rec.EffectiveType()), Status: string(rec.Status), Version: rec.Version, Title: rec.Title, Files: got})
		}
	}
	st, err := gatherNext(r, files)
	if err != nil {
		return doc, err
	}
	if actions := adviseNext(st); len(actions) > 0 {
		doc.Next = &actions[0]
	}
	return doc, nil
}

// /sedum:setFunction
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-9ec4a843","gen":"sha256:cc0301a10fcd903006e4f2bed32a32b6a082f192eb2c42bb2c80462870e50daf","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/govern_run.go","name":"printGovern"},"kwargs":"sha256:b0e091948cc8512254bf5a0229b9a688407ac82eabe25e3446e79397728246ea","body":"sha256:f09b10b71df5cae52b8747eaee14778de0233b34ac4eca6ca4f92e13c7a1d509"}
func printGovern(out io.Writer, doc governDoc) {
	fmt.Fprintf(out, "\nGOVERN: %s\n", ctxClean(strings.Join(doc.Files, ", "), 300))
	if len(doc.Governing) == 0 {
		fmt.Fprint(out, "  no active record governs these files\n")
	}
	for _, g := range doc.Governing {
		fmt.Fprintf(out, "  %s@%d  %s  %s  %s  (%s)\n", g.ID, g.Version, g.Status, g.Type, ctxClean(g.Title, 100), ctxClean(strings.Join(g.Files, ", "), 200))
	}
	if len(doc.Exempt) > 0 {
		fmt.Fprintf(out, "  exempt: %s\n", ctxClean(strings.Join(doc.Exempt, ", "), 300))
	}
	if doc.Next != nil {
		fmt.Fprintf(out, "  Next: %s: %s\n", doc.Next.Kind, ctxClean(doc.Next.Reason, 300))
		if doc.Next.Command != "" {
			fmt.Fprintf(out, "    Run: %s\n", ctxClean(doc.Next.Command, 300))
		}
	}
	fmt.Fprint(out, "\n")
}

// /sedum:setFunction
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-9ec4a843","gen":"sha256:cc0301a10fcd903006e4f2bed32a32b6a082f192eb2c42bb2c80462870e50daf","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/govern_run.go","name":"runGovern"},"kwargs":"sha256:b6ca606d312db03d077bae04ac38e15b3383f09578b2ca62640410e413c172bc","body":"sha256:957d5406a238968124fb771c0262a72c68238854d52f4fb487848c8fbb8ac955"}
func runGovern(out io.Writer, cmd *cobra.Command, args []string) error {
	r, err := openRepo(cmd)
	if err != nil {
		return err
	}
	files := splitFileList(mustStringArray(cmd, "files"))
	format, _ := cmd.Flags().GetString("format")
	asJSON, _ := cmd.Flags().GetBool("json")
	if format != "" && format != "human" && format != "json" {
		return fmt.Errorf("unknown format %q (human|json)", format)
	}
	given := nextDedupe(append(append([]string{}, args...), files...))
	if len(given) == 0 {
		return fmt.Errorf("no files: give file paths as arguments or use --files")
	}
	doc, err := buildGovern(r, given)
	if err != nil {
		return err
	}
	if asJSON || format == "json" {
		enc := json.NewEncoder(out)
		enc.SetIndent("", "  ")
		return enc.Encode(doc)
	}
	printGovern(out, doc)
	return nil
}

// /sedum:setFunction
