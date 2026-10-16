package main

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-f476b640","gen":"sha256:24167071b71953f56b7a9351b440d09a67530fb3dc4508cd3e405b0ad9bb4806","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/stale_run.go","path":"encoding/json"}}
import "encoding/json"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-f476b640","gen":"sha256:24167071b71953f56b7a9351b440d09a67530fb3dc4508cd3e405b0ad9bb4806","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/stale_run.go","path":"fmt"}}
import "fmt"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-f476b640","gen":"sha256:24167071b71953f56b7a9351b440d09a67530fb3dc4508cd3e405b0ad9bb4806","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/stale_run.go","path":"io"}}
import "io"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-f476b640","gen":"sha256:24167071b71953f56b7a9351b440d09a67530fb3dc4508cd3e405b0ad9bb4806","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/stale_run.go","path":"github.com/spf13/cobra"}}
import "github.com/spf13/cobra"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-f476b640","gen":"sha256:24167071b71953f56b7a9351b440d09a67530fb3dc4508cd3e405b0ad9bb4806","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/stale_run.go","path":"github.com/livecodelife/provenance/stale"}}
import "github.com/livecodelife/provenance/stale"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setType {"tier":"owned","record":"prov-2026-f476b640","gen":"sha256:24167071b71953f56b7a9351b440d09a67530fb3dc4508cd3e405b0ad9bb4806","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type staleFinding struct {\n\tID            string `json:\"id\"`\n\tVersion       int    `json:\"version\"`\n\tParent        string `json:\"parent,omitempty\"`\n\tPinned        int    `json:\"pinned\"`\n\tParentVersion int    `json:\"parent_version\"`\n\tKind          string `json:\"kind\"`\n\tCause         string `json:\"cause,omitempty\"`\n\tWhy           string `json:\"why,omitempty\"`\n}","file":"cmd/provenance/stale_run.go","name":"staleFinding"}}
type staleFinding struct {
	ID            string `json:"id"`
	Version       int    `json:"version"`
	Parent        string `json:"parent,omitempty"`
	Pinned        int    `json:"pinned"`
	ParentVersion int    `json:"parent_version"`
	Kind          string `json:"kind"`
	Cause         string `json:"cause,omitempty"`
	Why           string `json:"why,omitempty"`
}

// /sedum:setType
// sedum:setFunction {"tier":"owned","record":"prov-2026-f476b640","gen":"sha256:24167071b71953f56b7a9351b440d09a67530fb3dc4508cd3e405b0ad9bb4806","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func staleFindings(r *repo) ([]staleFinding, error) {\n\tshared, _ := loadShared(r)\n\tfound, err := stale.FindShared(r.Store, shared)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\tout := []staleFinding{}\n\tfor _, f := range found {\n\t\trec, _, err := r.Store.Load(f.ID)\n\t\tif err != nil {\n\t\t\treturn nil, err\n\t\t}\n\t\tout = append(out, staleFinding{ID: f.ID, Version: rec.Version, Parent: f.Parent, Pinned: f.Pinned, ParentVersion: f.ParentVersion, Kind: string(f.Kind), Cause: f.Cause, Why: f.Why})\n\t}\n\treturn out, nil\n}","file":"cmd/provenance/stale_run.go","name":"staleFindings"}}
func staleFindings(r *repo) ([]staleFinding, error) {
	shared, _ := loadShared(r)
	found, err := stale.FindShared(r.Store, shared)
	if err != nil {
		return nil, err
	}
	out := []staleFinding{}
	for _, f := range found {
		rec, _, err := r.Store.Load(f.ID)
		if err != nil {
			return nil, err
		}
		out = append(out, staleFinding{ID: f.ID, Version: rec.Version, Parent: f.Parent, Pinned: f.Pinned, ParentVersion: f.ParentVersion, Kind: string(f.Kind), Cause: f.Cause, Why: f.Why})
	}
	return out, nil
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-f476b640","gen":"sha256:24167071b71953f56b7a9351b440d09a67530fb3dc4508cd3e405b0ad9bb4806","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func printStale(out io.Writer, findings []staleFinding) {\n\tif len(findings) == 0 {\n\t\tfmt.Fprint(out, \"\\n✓ No stale records\\n\\n\")\n\t\treturn\n\t}\n\tfmt.Fprint(out, \"\\nSTALE RECORDS\\n\\n\")\n\tbehind := 0\n\tfor _, f := range findings {\n\t\tswitch f.Kind {\n\t\tcase string(stale.Direct):\n\t\t\tbehind++\n\t\t\tfmt.Fprintf(out, \"  %s@%d  pinned to %s@%d, parent is at %d  (direct)\\n\", f.ID, f.Version, f.Parent, f.Pinned, f.ParentVersion)\n\t\tcase string(stale.Transitive):\n\t\t\tbehind++\n\t\t\tfmt.Fprintf(out, \"  %s@%d  pinned to %s@%d  (transitive, via %s)\\n\", f.ID, f.Version, f.Parent, f.Pinned, f.Cause)\n\t\tdefault:\n\t\t\tfmt.Fprintf(out, \"  %s@%d  cannot be judged: %s\\n\", f.ID, f.Version, f.Why)\n\t\t}\n\t}\n\tif behind > 0 {\n\t\tfmt.Fprint(out, \"\\nConfirm one with: provenance repin <id> --reason \\\"...\\\"\\n\")\n\t}\n\tfmt.Fprint(out, \"\\n\")\n}","file":"cmd/provenance/stale_run.go","name":"printStale"}}
func printStale(out io.Writer, findings []staleFinding) {
	if len(findings) == 0 {
		fmt.Fprint(out, "\n✓ No stale records\n\n")
		return
	}
	fmt.Fprint(out, "\nSTALE RECORDS\n\n")
	behind := 0
	for _, f := range findings {
		switch f.Kind {
		case string(stale.Direct):
			behind++
			fmt.Fprintf(out, "  %s@%d  pinned to %s@%d, parent is at %d  (direct)\n", f.ID, f.Version, f.Parent, f.Pinned, f.ParentVersion)
		case string(stale.Transitive):
			behind++
			fmt.Fprintf(out, "  %s@%d  pinned to %s@%d  (transitive, via %s)\n", f.ID, f.Version, f.Parent, f.Pinned, f.Cause)
		default:
			fmt.Fprintf(out, "  %s@%d  cannot be judged: %s\n", f.ID, f.Version, f.Why)
		}
	}
	if behind > 0 {
		fmt.Fprint(out, "\nConfirm one with: provenance repin <id> --reason \"...\"\n")
	}
	fmt.Fprint(out, "\n")
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-f476b640","gen":"sha256:24167071b71953f56b7a9351b440d09a67530fb3dc4508cd3e405b0ad9bb4806","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func runStale(out io.Writer, cmd *cobra.Command, args []string) error {\n\tr, err := openRepo(cmd)\n\tif err != nil {\n\t\treturn err\n\t}\n\tasJSON, _ := cmd.Flags().GetBool(\"json\")\n\texitCode, _ := cmd.Flags().GetBool(\"exit-code\")\n\tfindings, err := staleFindings(r)\n\tif err != nil {\n\t\treturn err\n\t}\n\tbehind := 0\n\tfor _, f := range findings {\n\t\tif f.Kind != string(stale.Unknown) {\n\t\t\tbehind++\n\t\t}\n\t}\n\tif asJSON {\n\t\tenc := json.NewEncoder(out)\n\t\tenc.SetIndent(\"\", \"  \")\n\t\tif err := enc.Encode(map[string]any{\"findings\": findings}); err != nil {\n\t\t\treturn err\n\t\t}\n\t} else {\n\t\tprintStale(out, findings)\n\t}\n\tif exitCode && behind > 0 {\n\t\treturn fmt.Errorf(\"%d stale records\", behind)\n\t}\n\treturn nil\n}","file":"cmd/provenance/stale_run.go","name":"runStale"}}
func runStale(out io.Writer, cmd *cobra.Command, args []string) error {
	r, err := openRepo(cmd)
	if err != nil {
		return err
	}
	asJSON, _ := cmd.Flags().GetBool("json")
	exitCode, _ := cmd.Flags().GetBool("exit-code")
	findings, err := staleFindings(r)
	if err != nil {
		return err
	}
	behind := 0
	for _, f := range findings {
		if f.Kind != string(stale.Unknown) {
			behind++
		}
	}
	if asJSON {
		enc := json.NewEncoder(out)
		enc.SetIndent("", "  ")
		if err := enc.Encode(map[string]any{"findings": findings}); err != nil {
			return err
		}
	} else {
		printStale(out, findings)
	}
	if exitCode && behind > 0 {
		return fmt.Errorf("%d stale records", behind)
	}
	return nil
}

// /sedum:setFunction
