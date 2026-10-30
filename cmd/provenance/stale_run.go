package main

// sedum:anchor:imports
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-f476b640","gen":"sha256:24167071b71953f56b7a9351b440d09a67530fb3dc4508cd3e405b0ad9bb4806","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/stale_run.go","path":"encoding/json"},"kwargs":"sha256:a8dbed6e99f6351614e90492b2d8674d7276024f6bbe5b2c69f37e74153d1056","body":"sha256:872d7c3e9dee608e79a7f8ef57930d24b5e30f2a57d891a52549179ace45946a"}
import "encoding/json"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-f476b640","gen":"sha256:24167071b71953f56b7a9351b440d09a67530fb3dc4508cd3e405b0ad9bb4806","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/stale_run.go","path":"fmt"},"kwargs":"sha256:daf2cf46d251a82d2533f04f12d62ce53305311250338c62f1af2fb691c88b88","body":"sha256:bcbac51883cbf8064e9196b371cb880cee429866652447ad557533d817857609"}
import "fmt"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-f476b640","gen":"sha256:24167071b71953f56b7a9351b440d09a67530fb3dc4508cd3e405b0ad9bb4806","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/stale_run.go","path":"io"},"kwargs":"sha256:92672a3059fdcc9c86c993af2ed2ee726fbaca5ec436f33eeae2d1bc63310286","body":"sha256:ef8e8abf4ef5ee40a19dc85b49bf09c5df1fe32abfe834e0ab4c2175baf80ddf"}
import "io"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-f476b640","gen":"sha256:24167071b71953f56b7a9351b440d09a67530fb3dc4508cd3e405b0ad9bb4806","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/stale_run.go","path":"github.com/spf13/cobra"},"kwargs":"sha256:f58d4ed8e40615db99e7abd019fde5eb3ced1e0a11ef56c7f425711ce5af96e9","body":"sha256:f361c02d0f8e1999cf3fb715393e16310eeec9debd5aaa13936af7617e2b46aa"}
import "github.com/spf13/cobra"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-f476b640","gen":"sha256:24167071b71953f56b7a9351b440d09a67530fb3dc4508cd3e405b0ad9bb4806","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/stale_run.go","path":"github.com/livecodelife/provenance/stale"},"kwargs":"sha256:b0ccf629ed18c137aaa8684fdb46af0a1dc57e286e4bb25c9b65d15af4eac171","body":"sha256:9d7308eb26aa328e6ffd9c6c6d9334e95ead3e237675b3953a96d84096c468d0"}
import "github.com/livecodelife/provenance/stale"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setType {"v":2,"tier":"owned","record":"prov-2026-f476b640","gen":"sha256:24167071b71953f56b7a9351b440d09a67530fb3dc4508cd3e405b0ad9bb4806","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","id":{"file":"cmd/provenance/stale_run.go","name":"staleFinding"},"kwargs":"sha256:57efb982a5e17ea8e12edae95dd6bf216dbada19a6ba4a275a501e43b638c11e","body":"sha256:096d246f065f3d47bfddb38e23ee33987e0067a397502baf15f0ad5f8cf5fe06"}
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
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-f476b640","gen":"sha256:24167071b71953f56b7a9351b440d09a67530fb3dc4508cd3e405b0ad9bb4806","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/stale_run.go","name":"staleFindings"},"kwargs":"sha256:d3ff38d21fad2c4d770c2d14dc3e610b72c2aa5f8cf01523cb20e73af47e1730","body":"sha256:ab40bd4aea5e19f2cb1c56d3c0c2715ae17f4a129c687060df9968f7fe710b9e"}
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
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-f476b640","gen":"sha256:24167071b71953f56b7a9351b440d09a67530fb3dc4508cd3e405b0ad9bb4806","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/stale_run.go","name":"printStale"},"kwargs":"sha256:040aece69deb49b55fa228777322e2aa690e5ad2f4f2165f67b8d99ceca4bda9","body":"sha256:77391f684cc4ea4ea3358eebc3bafa0b51dc8b802a9eb829f7bd96b9e673fb1a"}
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
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-f476b640","gen":"sha256:24167071b71953f56b7a9351b440d09a67530fb3dc4508cd3e405b0ad9bb4806","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/stale_run.go","name":"runStale"},"kwargs":"sha256:dda0141cd3e5dc7cae0d8c5b710de465edda6d2fea98371e0a95e0048648d573","body":"sha256:5a3f9ac75333fd256af44f9e257ce1dc3c3eb9a8cab3ab4eb57b96ccbd2b361f"}
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
