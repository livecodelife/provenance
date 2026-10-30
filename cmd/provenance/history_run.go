package main

// sedum:anchor:imports
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-d435c27b","gen":"sha256:4f36ba1a5a6a79f262293754b5a46ef8739347c5fb4b6df037af455f19b551ba","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/history_run.go","path":"encoding/json"},"kwargs":"sha256:bc63593eee4af215cbf3ee352e82b0496169d42adea55251d05ff562889f48a2","body":"sha256:872d7c3e9dee608e79a7f8ef57930d24b5e30f2a57d891a52549179ace45946a"}
import "encoding/json"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-d435c27b","gen":"sha256:4f36ba1a5a6a79f262293754b5a46ef8739347c5fb4b6df037af455f19b551ba","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/history_run.go","path":"fmt"},"kwargs":"sha256:353fcc28a17503281be07d70cc79fb809801ea028894215d1ba20af8f697dcf7","body":"sha256:bcbac51883cbf8064e9196b371cb880cee429866652447ad557533d817857609"}
import "fmt"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-d435c27b","gen":"sha256:4f36ba1a5a6a79f262293754b5a46ef8739347c5fb4b6df037af455f19b551ba","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/history_run.go","path":"io"},"kwargs":"sha256:6ffc3405830ddd7727b8f463ad0529298804c396eb80a0fe440c2b0ffde664fe","body":"sha256:ef8e8abf4ef5ee40a19dc85b49bf09c5df1fe32abfe834e0ab4c2175baf80ddf"}
import "io"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-d435c27b","gen":"sha256:4f36ba1a5a6a79f262293754b5a46ef8739347c5fb4b6df037af455f19b551ba","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/history_run.go","path":"strconv"},"kwargs":"sha256:22dbc38d67659e5dd417eaaae7a563f9dc18773016ef839ecf9f23743bd02d5a","body":"sha256:f33f981c64a42d78a5daa7ff3fca15b10045080cccd282fef544b59eaf7f6dc8"}
import "strconv"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-d435c27b","gen":"sha256:4f36ba1a5a6a79f262293754b5a46ef8739347c5fb4b6df037af455f19b551ba","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/history_run.go","path":"strings"},"kwargs":"sha256:cc6b3e9cd8925ebb037a18bd10be5fd2dbf14abb2a0264472c04724ab157e69a","body":"sha256:a9af3c0bc4bf9342df9d7e89a324523fdba7941a9ca4190270e68ff3c4a25ea6"}
import "strings"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-d435c27b","gen":"sha256:4f36ba1a5a6a79f262293754b5a46ef8739347c5fb4b6df037af455f19b551ba","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/history_run.go","path":"github.com/spf13/cobra"},"kwargs":"sha256:dd95dba12c749d200e0ee44e9e28df0dce0712981a6e16fb3389083ae0d5adbe","body":"sha256:f361c02d0f8e1999cf3fb715393e16310eeec9debd5aaa13936af7617e2b46aa"}
import "github.com/spf13/cobra"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-d435c27b","gen":"sha256:4f36ba1a5a6a79f262293754b5a46ef8739347c5fb4b6df037af455f19b551ba","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/history_run.go","path":"github.com/livecodelife/provenance/record"},"kwargs":"sha256:3be02db87421d7977d9a63dde216a9e27a6c2ca96645653dbff874536106e27d","body":"sha256:acbf8628061291d11aa39e9ea64ecc91095a117057cfe6f1958ef2c437fac003"}
import "github.com/livecodelife/provenance/record"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setType {"v":2,"tier":"owned","record":"prov-2026-d435c27b","gen":"sha256:4f36ba1a5a6a79f262293754b5a46ef8739347c5fb4b6df037af455f19b551ba","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","id":{"file":"cmd/provenance/history_run.go","name":"historyVersion"},"kwargs":"sha256:3f0dcb7cbe35a7b44218e45856c550f06f68f1126dcd87b37ae59c36c804c21a","body":"sha256:d60208287fb17411465bd910e3d68e0b7d53709c97d3509a2807403f299f1bd9"}
type historyVersion struct {
	Version     int      `json:"version"`
	CreatedAt   string   `json:"created_at"`
	Author      string   `json:"author"`
	Reason      string   `json:"reason"`
	ReasonKind  string   `json:"reason_kind,omitempty"`
	ContentHash string   `json:"content_hash"`
	Commits     []string `json:"commits"`
}

// /sedum:setType
// sedum:setType {"v":2,"tier":"owned","record":"prov-2026-d435c27b","gen":"sha256:4f36ba1a5a6a79f262293754b5a46ef8739347c5fb4b6df037af455f19b551ba","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","id":{"file":"cmd/provenance/history_run.go","name":"historyJSON"},"kwargs":"sha256:eef43818f0d55b95586c176fbb18c9794254b4186511a8e51c50864386ad7cb5","body":"sha256:352e09380664c8b20e776ba62444808a15d0eed268c71ae240cb10a2027b72cf"}
type historyJSON struct {
	ID       string           `json:"id"`
	Title    string           `json:"title"`
	Versions []historyVersion `json:"versions"`
	Untagged []string         `json:"untagged_commits"`
}

// /sedum:setType
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-d435c27b","gen":"sha256:4f36ba1a5a6a79f262293754b5a46ef8739347c5fb4b6df037af455f19b551ba","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/history_run.go","name":"grepCommits"},"kwargs":"sha256:9b89394e95ee8a0af16a79785c8487cd094e54d50c31ec9839b4ecc846b8348c","body":"sha256:77998befd630f622c2445c92b6635eef5a8e2f0872d5eb06bf2b2236f22781e8"}
func grepCommits(r *repo, pattern string) []string {
	out, err := r.git("log", "--all", "-E", "--grep", pattern, "--format=%h %s")
	if err != nil {
		return nil
	}
	commits := []string{}
	for _, line := range strings.Split(out, "\n") {
		if line != "" {
			commits = append(commits, line)
		}
	}
	return commits
}

// /sedum:setFunction
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-d435c27b","gen":"sha256:4f36ba1a5a6a79f262293754b5a46ef8739347c5fb4b6df037af455f19b551ba","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/history_run.go","name":"runHistory"},"kwargs":"sha256:87c946a62f6cda0bcb87283bb7f654b08c5e5aafca51b741b3297ce4f593f8bc","body":"sha256:56c58e0cb9a977baaac795e983e2ea17f73c5acee496aa747427fc891b0f5d9c"}
func runHistory(out io.Writer, cmd *cobra.Command, args []string) error {
	r, err := openRepo(cmd)
	if err != nil {
		return err
	}
	id, err := recordArg(cmd, args)
	if err != nil {
		return err
	}
	rec, _, err := r.resolve(id)
	if err != nil {
		return err
	}
	asJSON, _ := cmd.Flags().GetBool("json")
	vs, err := r.Store.Versions(rec.ID)
	if err != nil {
		return err
	}
	h := historyJSON{ID: rec.ID, Title: rec.Title, Versions: []historyVersion{}}
	for _, v := range vs {
		snap, err := r.Store.Snapshot(rec.ID, v)
		if err != nil {
			return err
		}
		commits := grepCommits(r, `\[`+rec.ID+`@`+strconv.Itoa(v)+`\]`)
		h.Versions = append(h.Versions, historyVersion{v, snap.CreatedAt, snap.Author, snap.Reason, snap.ReasonKind, snap.ContentHash, commits})
	}
	ids := append([]string{rec.ID}, rec.Aliases...)
	h.Untagged = grepCommits(r, `\[(`+strings.Join(ids, "|")+`)\]`)
	if asJSON {
		enc := json.NewEncoder(out)
		enc.SetIndent("", "  ")
		return enc.Encode(h)
	}
	fmt.Fprintf(out, "\n%s  %s\n", rec.ID, rec.Title)
	if len(vs) == 0 {
		fmt.Fprintf(out, "\n  %s has no versions yet (a draft: open writes version 1)\n\n", rec.ID)
		return nil
	}
	for _, v := range h.Versions {
		fmt.Fprintf(out, "\n  v%d  %s  %s\n", v.Version, v.CreatedAt, v.Author)
		if v.Reason != "" {
			fmt.Fprintf(out, "      reason: %s", v.Reason)
			if v.ReasonKind == record.ReasonKindFix {
				fmt.Fprint(out, " [fix]")
			}
			fmt.Fprint(out, "\n")
		}
		hash := v.ContentHash
		if len(hash) > 19 {
			hash = hash[:19]
		}
		fmt.Fprintf(out, "      content: %s\n", hash)
		for _, c := range v.Commits {
			fmt.Fprintf(out, "      · %s\n", c)
		}
	}
	if len(h.Untagged) > 0 {
		fmt.Fprint(out, "\n  Commits tagged without a version:\n")
		for _, c := range h.Untagged {
			fmt.Fprintf(out, "      · %s\n", c)
		}
	}
	fmt.Fprint(out, "\n")
	return nil
}

// /sedum:setFunction
