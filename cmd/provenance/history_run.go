package main

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-d435c27b","gen":"sha256:4f36ba1a5a6a79f262293754b5a46ef8739347c5fb4b6df037af455f19b551ba","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/history_run.go","path":"encoding/json"}}
import "encoding/json"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-d435c27b","gen":"sha256:4f36ba1a5a6a79f262293754b5a46ef8739347c5fb4b6df037af455f19b551ba","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/history_run.go","path":"fmt"}}
import "fmt"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-d435c27b","gen":"sha256:4f36ba1a5a6a79f262293754b5a46ef8739347c5fb4b6df037af455f19b551ba","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/history_run.go","path":"io"}}
import "io"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-d435c27b","gen":"sha256:4f36ba1a5a6a79f262293754b5a46ef8739347c5fb4b6df037af455f19b551ba","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/history_run.go","path":"strconv"}}
import "strconv"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-d435c27b","gen":"sha256:4f36ba1a5a6a79f262293754b5a46ef8739347c5fb4b6df037af455f19b551ba","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/history_run.go","path":"strings"}}
import "strings"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-d435c27b","gen":"sha256:4f36ba1a5a6a79f262293754b5a46ef8739347c5fb4b6df037af455f19b551ba","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/history_run.go","path":"github.com/spf13/cobra"}}
import "github.com/spf13/cobra"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-d435c27b","gen":"sha256:4f36ba1a5a6a79f262293754b5a46ef8739347c5fb4b6df037af455f19b551ba","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/history_run.go","path":"github.com/livecodelife/provenance/record"}}
import "github.com/livecodelife/provenance/record"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setType {"tier":"owned","record":"prov-2026-d435c27b","gen":"sha256:4f36ba1a5a6a79f262293754b5a46ef8739347c5fb4b6df037af455f19b551ba","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type historyVersion struct {\n\tVersion     int      `json:\"version\"`\n\tCreatedAt   string   `json:\"created_at\"`\n\tAuthor      string   `json:\"author\"`\n\tReason      string   `json:\"reason\"`\n\tReasonKind  string   `json:\"reason_kind,omitempty\"`\n\tContentHash string   `json:\"content_hash\"`\n\tCommits     []string `json:\"commits\"`\n}","file":"cmd/provenance/history_run.go","name":"historyVersion"}}
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
// sedum:setType {"tier":"owned","record":"prov-2026-d435c27b","gen":"sha256:4f36ba1a5a6a79f262293754b5a46ef8739347c5fb4b6df037af455f19b551ba","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type historyJSON struct {\n\tID       string           `json:\"id\"`\n\tTitle    string           `json:\"title\"`\n\tVersions []historyVersion `json:\"versions\"`\n\tUntagged []string         `json:\"untagged_commits\"`\n}","file":"cmd/provenance/history_run.go","name":"historyJSON"}}
type historyJSON struct {
	ID       string           `json:"id"`
	Title    string           `json:"title"`
	Versions []historyVersion `json:"versions"`
	Untagged []string         `json:"untagged_commits"`
}

// /sedum:setType
// sedum:setFunction {"tier":"owned","record":"prov-2026-d435c27b","gen":"sha256:4f36ba1a5a6a79f262293754b5a46ef8739347c5fb4b6df037af455f19b551ba","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func grepCommits(r *repo, pattern string) []string {\n\tout, err := r.git(\"log\", \"--all\", \"-E\", \"--grep\", pattern, \"--format=%h %s\")\n\tif err != nil {\n\t\treturn nil\n\t}\n\tcommits := []string{}\n\tfor _, line := range strings.Split(out, \"\\n\") {\n\t\tif line != \"\" {\n\t\t\tcommits = append(commits, line)\n\t\t}\n\t}\n\treturn commits\n}","file":"cmd/provenance/history_run.go","name":"grepCommits"}}
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
// sedum:setFunction {"tier":"owned","record":"prov-2026-d435c27b","gen":"sha256:4f36ba1a5a6a79f262293754b5a46ef8739347c5fb4b6df037af455f19b551ba","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func runHistory(out io.Writer, cmd *cobra.Command, args []string) error {\n\tr, err := openRepo(cmd)\n\tif err != nil {\n\t\treturn err\n\t}\n\tid, err := recordArg(cmd, args)\n\tif err != nil {\n\t\treturn err\n\t}\n\trec, _, err := r.resolve(id)\n\tif err != nil {\n\t\treturn err\n\t}\n\tasJSON, _ := cmd.Flags().GetBool(\"json\")\n\tvs, err := r.Store.Versions(rec.ID)\n\tif err != nil {\n\t\treturn err\n\t}\n\th := historyJSON{ID: rec.ID, Title: rec.Title, Versions: []historyVersion{}}\n\tfor _, v := range vs {\n\t\tsnap, err := r.Store.Snapshot(rec.ID, v)\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n\t\tcommits := grepCommits(r, `\\[`+rec.ID+`@`+strconv.Itoa(v)+`\\]`)\n\t\th.Versions = append(h.Versions, historyVersion{v, snap.CreatedAt, snap.Author, snap.Reason, snap.ReasonKind, snap.ContentHash, commits})\n\t}\n\tids := append([]string{rec.ID}, rec.Aliases...)\n\th.Untagged = grepCommits(r, `\\[(`+strings.Join(ids, \"|\")+`)\\]`)\n\tif asJSON {\n\t\tenc := json.NewEncoder(out)\n\t\tenc.SetIndent(\"\", \"  \")\n\t\treturn enc.Encode(h)\n\t}\n\tfmt.Fprintf(out, \"\\n%s  %s\\n\", rec.ID, rec.Title)\n\tif len(vs) == 0 {\n\t\tfmt.Fprintf(out, \"\\n  %s has no versions yet (a draft: open writes version 1)\\n\\n\", rec.ID)\n\t\treturn nil\n\t}\n\tfor _, v := range h.Versions {\n\t\tfmt.Fprintf(out, \"\\n  v%d  %s  %s\\n\", v.Version, v.CreatedAt, v.Author)\n\t\tif v.Reason != \"\" {\n\t\t\tfmt.Fprintf(out, \"      reason: %s\", v.Reason)\n\t\t\tif v.ReasonKind == record.ReasonKindFix {\n\t\t\t\tfmt.Fprint(out, \" [fix]\")\n\t\t\t}\n\t\t\tfmt.Fprint(out, \"\\n\")\n\t\t}\n\t\thash := v.ContentHash\n\t\tif len(hash) > 19 {\n\t\t\thash = hash[:19]\n\t\t}\n\t\tfmt.Fprintf(out, \"      content: %s\\n\", hash)\n\t\tfor _, c := range v.Commits {\n\t\t\tfmt.Fprintf(out, \"      · %s\\n\", c)\n\t\t}\n\t}\n\tif len(h.Untagged) > 0 {\n\t\tfmt.Fprint(out, \"\\n  Commits tagged without a version:\\n\")\n\t\tfor _, c := range h.Untagged {\n\t\t\tfmt.Fprintf(out, \"      · %s\\n\", c)\n\t\t}\n\t}\n\tfmt.Fprint(out, \"\\n\")\n\treturn nil\n}","file":"cmd/provenance/history_run.go","name":"runHistory"}}
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
