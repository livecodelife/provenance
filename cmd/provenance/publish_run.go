package main

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-2a8e5d26","gen":"sha256:9a2a0375786ac0e3d5d087a0697dc5388e847d273eefd309880f43efa84de458","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/publish_run.go","path":"regexp"}}
import "regexp"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-2d51bcd8","gen":"sha256:e54095c8862121df6f5379c48e015e47b5beb7fc5c5ac36dc07deae2b9541771","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/publish_run.go","path":"strings"}}
import "strings"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-2d51bcd8","gen":"sha256:e54095c8862121df6f5379c48e015e47b5beb7fc5c5ac36dc07deae2b9541771","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/publish_run.go","path":"github.com/livecodelife/provenance/remote"}}
import "github.com/livecodelife/provenance/remote"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-50c9b7be","gen":"sha256:8091f07ebbb51f363b40f23beed20cec9a67d51de4f3275e0687ef8efee7ada1","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/publish_run.go","path":"fmt"}}
import "fmt"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-50c9b7be","gen":"sha256:8091f07ebbb51f363b40f23beed20cec9a67d51de4f3275e0687ef8efee7ada1","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/publish_run.go","path":"io"}}
import "io"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-50c9b7be","gen":"sha256:8091f07ebbb51f363b40f23beed20cec9a67d51de4f3275e0687ef8efee7ada1","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/publish_run.go","path":"strconv"}}
import "strconv"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-50c9b7be","gen":"sha256:8091f07ebbb51f363b40f23beed20cec9a67d51de4f3275e0687ef8efee7ada1","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/publish_run.go","path":"github.com/spf13/cobra"}}
import "github.com/spf13/cobra"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-647a7038","gen":"sha256:3e57f8c5dd2931ae920108e119bd04522a3beaa46442493d78910a3b420ee6a6","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/publish_run.go","path":"errors"}}
import "errors"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-7163f9f2","gen":"sha256:4c6bc602935bc9046f542c490eed5787bcb5d3211e9cbc9bb29c42009c668f9a","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/publish_run.go","path":"github.com/livecodelife/provenance/record"}}
import "github.com/livecodelife/provenance/record"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-b8cb30d5","gen":"sha256:0604e95152d8f50c4c7576551c6513df7c9b67351e2b87af6e518c82a1684cd0","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/publish_run.go","path":"net/url"}}
import "net/url"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setVar {"tier":"owned","record":"prov-2026-2a8e5d26","gen":"sha256:9a2a0375786ac0e3d5d087a0697dc5388e847d273eefd309880f43efa84de458","pkg":"sha256:c7b9d9ac1655e8211b8b2ca8c49cf34b8d510d8d823d4f5eb60ec7e1018aa2eb","kwargs":{"code":"var didSyntax = regexp.MustCompile(`^did:[a-z]+:[A-Za-z0-9._:%-]*[A-Za-z0-9._%-]$`)","file":"cmd/provenance/publish_run.go","name":"didSyntax"}}
var didSyntax = regexp.MustCompile(`^did:[a-z]+:[A-Za-z0-9._:%-]*[A-Za-z0-9._%-]$`)

// /sedum:setVar
// sedum:setFunction {"tier":"owned","record":"prov-2026-2d51bcd8","gen":"sha256:e54095c8862121df6f5379c48e015e47b5beb7fc5c5ac36dc07deae2b9541771","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func headUpdate(it remote.PublishItem) string {\n\tvar parts []string\n\tif it.HeadStatus != \"\" {\n\t\tparts = append(parts, \"status → \"+it.HeadStatus)\n\t}\n\tif it.SetProject != \"\" {\n\t\tparts = append(parts, \"project → \"+it.SetProject)\n\t}\n\tif it.SetSource != nil {\n\t\tparts = append(parts, \"source → \"+it.SetSource.Repo+\" \"+it.SetSource.Path)\n\t}\n\treturn strings.Join(parts, \", \")\n}","file":"cmd/provenance/publish_run.go","name":"headUpdate"}}
func headUpdate(it remote.PublishItem) string {
	var parts []string
	if it.HeadStatus != "" {
		parts = append(parts, "status → "+it.HeadStatus)
	}
	if it.SetProject != "" {
		parts = append(parts, "project → "+it.SetProject)
	}
	if it.SetSource != nil {
		parts = append(parts, "source → "+it.SetSource.Repo+" "+it.SetSource.Path)
	}
	return strings.Join(parts, ", ")
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-50c9b7be","gen":"sha256:8091f07ebbb51f363b40f23beed20cec9a67d51de4f3275e0687ef8efee7ada1","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func runPublish(out io.Writer, cmd *cobra.Command, args []string) error {\n\tall, _ := cmd.Flags().GetBool(\"all\")\n\tyes, _ := cmd.Flags().GetBool(\"yes\")\n\tplanOnly, _ := cmd.Flags().GetBool(\"plan\")\n\tproject, _ := cmd.Flags().GetString(\"project\")\n\tif project != \"\" {\n\t\tif err := checkProject(project); err != nil {\n\t\t\treturn fmt.Errorf(\"--project %q: %w\", project, err)\n\t\t}\n\t}\n\tr, err := openRepo(cmd)\n\tif err != nil {\n\t\treturn err\n\t}\n\tids, err := recordIDs(r, args, all)\n\tif err != nil {\n\t\treturn err\n\t}\n\tclient, err := apiClient()\n\tif err != nil {\n\t\treturn err\n\t}\n\tsrc, notes := publishSource(r)\n\tpl, err := remote.PlanPublishWith(cmd.Context(), r.Store, ids, client, remote.HeadWant{Project: project, Source: src})\n\tif err != nil {\n\t\treturn err\n\t}\n\tfor _, note := range notes {\n\t\tfmt.Fprintf(out, \"  note: %s\\n\", note)\n\t}\n\tif planOnly {\n\t\tfmt.Fprintln(out)\n\t\twritePlan(out, pl)\n\t\treturn nil\n\t}\n\tfor _, s := range pl.Skipped {\n\t\tfmt.Fprintf(out, \"  skipped  %s\\n\", s)\n\t}\n\tfor _, p := range pl.Problems {\n\t\tfmt.Fprintf(out, \"  skipped  %s: %s\\n\", p.ID, p.Reason)\n\t}\n\tskipErr := fmt.Errorf(\"%d record(s) were skipped (see above); the others were handled.\", len(pl.Problems))\n\tif pl.Empty() {\n\t\tfmt.Fprintln(out)\n\t\tfmt.Fprintln(out, \"Nothing to publish: the server already has every version and status.\")\n\t\tfmt.Fprintln(out)\n\t\tif len(pl.Problems) > 0 {\n\t\t\treturn skipErr\n\t\t}\n\t\treturn nil\n\t}\n\tfmt.Fprintln(out)\n\tfmt.Fprintln(out, \"These become PUBLIC on AT Proto (anyone can read them):\")\n\tfmt.Fprintln(out)\n\tfor _, it := range pl.Items {\n\t\tfmt.Fprintf(out, \"  %s  %s\\n\", it.ID, it.Title)\n\t\tif len(it.Snaps) == 0 {\n\t\t\tfmt.Fprintf(out, \"      %s (no new version)\\n\", headUpdate(it))\n\t\t\tcontinue\n\t\t}\n\t\tlabels := make([]string, 0, len(it.Snaps))\n\t\tfor _, s := range it.Snaps {\n\t\t\tlabels = append(labels, \"v\"+strconv.Itoa(s.Version))\n\t\t}\n\t\tfmt.Fprintf(out, \"      versions %s, status %s\\n\", strings.Join(labels, \", \"), it.Status)\n\t}\n\tfmt.Fprintln(out)\n\tif !yes {\n\t\tif !confirm(cmd, out, \"Publish these? [y/N] \") {\n\t\t\treturn fmt.Errorf(\"not confirmed, nothing was published (run interactively, or pass --yes), and writes nothing to the server.\")\n\t\t}\n\t}\n\tcommitOf := func(id string, version int) string {\n\t\ttagPattern := fmt.Sprintf(`\\[%s@%d\\]`, id, version)\n\t\tsha, err := r.git(\"log\", \"--all\", \"-E\", \"-n\", \"1\", \"--format=%H\", \"--grep\", tagPattern)\n\t\tif err != nil {\n\t\t\treturn \"\"\n\t\t}\n\t\treturn strings.TrimSpace(sha)\n\t}\n\twrites := 0\n\tprogress := func(id string, version int) {\n\t\twrites++\n\t\tif version == 0 {\n\t\t\tfmt.Fprintf(out, \"  published  %s head\\n\", id)\n\t\t\treturn\n\t\t}\n\t\tfmt.Fprintf(out, \"  published  %s v%d\\n\", id, version)\n\t}\n\tif err := remote.ApplyPublish(cmd.Context(), pl, client, project, commitOf, progress); err != nil {\n\t\treturn err\n\t}\n\tfmt.Fprintln(out)\n\tfmt.Fprintf(out, \"\\u2713 Published %d write(s).\\n\", writes)\n\tfmt.Fprintln(out)\n\tif len(pl.Problems) > 0 {\n\t\treturn skipErr\n\t}\n\treturn nil\n}","file":"cmd/provenance/publish_run.go","name":"runPublish"}}
func runPublish(out io.Writer, cmd *cobra.Command, args []string) error {
	all, _ := cmd.Flags().GetBool("all")
	yes, _ := cmd.Flags().GetBool("yes")
	planOnly, _ := cmd.Flags().GetBool("plan")
	project, _ := cmd.Flags().GetString("project")
	if project != "" {
		if err := checkProject(project); err != nil {
			return fmt.Errorf("--project %q: %w", project, err)
		}
	}
	r, err := openRepo(cmd)
	if err != nil {
		return err
	}
	ids, err := recordIDs(r, args, all)
	if err != nil {
		return err
	}
	client, err := apiClient()
	if err != nil {
		return err
	}
	src, notes := publishSource(r)
	pl, err := remote.PlanPublishWith(cmd.Context(), r.Store, ids, client, remote.HeadWant{Project: project, Source: src})
	if err != nil {
		return err
	}
	for _, note := range notes {
		fmt.Fprintf(out, "  note: %s\n", note)
	}
	if planOnly {
		fmt.Fprintln(out)
		writePlan(out, pl)
		return nil
	}
	for _, s := range pl.Skipped {
		fmt.Fprintf(out, "  skipped  %s\n", s)
	}
	for _, p := range pl.Problems {
		fmt.Fprintf(out, "  skipped  %s: %s\n", p.ID, p.Reason)
	}
	skipErr := fmt.Errorf("%d record(s) were skipped (see above); the others were handled.", len(pl.Problems))
	if pl.Empty() {
		fmt.Fprintln(out)
		fmt.Fprintln(out, "Nothing to publish: the server already has every version and status.")
		fmt.Fprintln(out)
		if len(pl.Problems) > 0 {
			return skipErr
		}
		return nil
	}
	fmt.Fprintln(out)
	fmt.Fprintln(out, "These become PUBLIC on AT Proto (anyone can read them):")
	fmt.Fprintln(out)
	for _, it := range pl.Items {
		fmt.Fprintf(out, "  %s  %s\n", it.ID, it.Title)
		if len(it.Snaps) == 0 {
			fmt.Fprintf(out, "      %s (no new version)\n", headUpdate(it))
			continue
		}
		labels := make([]string, 0, len(it.Snaps))
		for _, s := range it.Snaps {
			labels = append(labels, "v"+strconv.Itoa(s.Version))
		}
		fmt.Fprintf(out, "      versions %s, status %s\n", strings.Join(labels, ", "), it.Status)
	}
	fmt.Fprintln(out)
	if !yes {
		if !confirm(cmd, out, "Publish these? [y/N] ") {
			return fmt.Errorf("not confirmed, nothing was published (run interactively, or pass --yes), and writes nothing to the server.")
		}
	}
	commitOf := func(id string, version int) string {
		tagPattern := fmt.Sprintf(`\[%s@%d\]`, id, version)
		sha, err := r.git("log", "--all", "-E", "-n", "1", "--format=%H", "--grep", tagPattern)
		if err != nil {
			return ""
		}
		return strings.TrimSpace(sha)
	}
	writes := 0
	progress := func(id string, version int) {
		writes++
		if version == 0 {
			fmt.Fprintf(out, "  published  %s head\n", id)
			return
		}
		fmt.Fprintf(out, "  published  %s v%d\n", id, version)
	}
	if err := remote.ApplyPublish(cmd.Context(), pl, client, project, commitOf, progress); err != nil {
		return err
	}
	fmt.Fprintln(out)
	fmt.Fprintf(out, "\u2713 Published %d write(s).\n", writes)
	fmt.Fprintln(out)
	if len(pl.Problems) > 0 {
		return skipErr
	}
	return nil
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-647a7038","gen":"sha256:3e57f8c5dd2931ae920108e119bd04522a3beaa46442493d78910a3b420ee6a6","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func checkProject(s string) error {\n\tif !strings.HasPrefix(s, \"at://\") {\n\t\treturn errors.New(\"must start with at://, for example at://did:plc:abc123/dev.provisio.project/name\")\n\t}\n\tif strings.ContainsAny(s, \"<>{}\") {\n\t\treturn errors.New(\"looks like a placeholder; use your real DID, for example at://did:plc:abc123/dev.provisio.project/name\")\n\t}\n\tif strings.Contains(s, \" \") {\n\t\treturn errors.New(\"must not contain a space\")\n\t}\n\tif strings.Contains(s, \"?\") {\n\t\treturn errors.New(\"must not have a query (?)\")\n\t}\n\tif strings.Contains(s, \"#\") {\n\t\treturn errors.New(\"must not have a fragment (#)\")\n\t}\n\tfor _, c := range s {\n\t\tif c < 33 || c > 126 {\n\t\t\treturn errors.New(\"must be ASCII\")\n\t\t}\n\t}\n\trest, _ := strings.CutPrefix(s, \"at://\")\n\tparts := strings.Split(rest, \"/\")\n\tauthority := parts[0]\n\tif !didSyntax.MatchString(authority) && !handleSyntax.MatchString(authority) {\n\t\treturn fmt.Errorf(\"the authority %q is not a DID (did:plc:...) or a handle\", authority)\n\t}\n\tif strings.HasSuffix(s, \"/\") {\n\t\treturn errors.New(\"must not end with a slash\")\n\t}\n\tif len(parts) > 3 {\n\t\treturn errors.New(\"has too many segments: at://<authority>[/<collection>[/<record key>]]\")\n\t}\n\tif len(parts) >= 2 && !nsidSyntax.MatchString(parts[1]) {\n\t\treturn fmt.Errorf(\"the collection %q is not an NSID such as dev.provisio.project\", parts[1])\n\t}\n\tif len(parts) >= 3 && (!rkeySyntax.MatchString(parts[2]) || parts[2] == \".\" || parts[2] == \"..\") {\n\t\treturn fmt.Errorf(\"the record key %q is not valid (ASCII letters, digits and . _ ~ : -)\", parts[2])\n\t}\n\treturn nil\n}","file":"cmd/provenance/publish_run.go","name":"checkProject"}}
func checkProject(s string) error {
	if !strings.HasPrefix(s, "at://") {
		return errors.New("must start with at://, for example at://did:plc:abc123/dev.provisio.project/name")
	}
	if strings.ContainsAny(s, "<>{}") {
		return errors.New("looks like a placeholder; use your real DID, for example at://did:plc:abc123/dev.provisio.project/name")
	}
	if strings.Contains(s, " ") {
		return errors.New("must not contain a space")
	}
	if strings.Contains(s, "?") {
		return errors.New("must not have a query (?)")
	}
	if strings.Contains(s, "#") {
		return errors.New("must not have a fragment (#)")
	}
	for _, c := range s {
		if c < 33 || c > 126 {
			return errors.New("must be ASCII")
		}
	}
	rest, _ := strings.CutPrefix(s, "at://")
	parts := strings.Split(rest, "/")
	authority := parts[0]
	if !didSyntax.MatchString(authority) && !handleSyntax.MatchString(authority) {
		return fmt.Errorf("the authority %q is not a DID (did:plc:...) or a handle", authority)
	}
	if strings.HasSuffix(s, "/") {
		return errors.New("must not end with a slash")
	}
	if len(parts) > 3 {
		return errors.New("has too many segments: at://<authority>[/<collection>[/<record key>]]")
	}
	if len(parts) >= 2 && !nsidSyntax.MatchString(parts[1]) {
		return fmt.Errorf("the collection %q is not an NSID such as dev.provisio.project", parts[1])
	}
	if len(parts) >= 3 && (!rkeySyntax.MatchString(parts[2]) || parts[2] == "." || parts[2] == "..") {
		return fmt.Errorf("the record key %q is not valid (ASCII letters, digits and . _ ~ : -)", parts[2])
	}
	return nil
}

// /sedum:setFunction
// sedum:setVar {"tier":"owned","record":"prov-2026-7155fcd2","gen":"sha256:e81dd0d47582e1ffbae9e3a018150d59d1afc6d67c83c470e85c6ba0b782f0a6","pkg":"sha256:c7b9d9ac1655e8211b8b2ca8c49cf34b8d510d8d823d4f5eb60ec7e1018aa2eb","kwargs":{"code":"var nsidSyntax = regexp.MustCompile(`^[a-zA-Z]([a-zA-Z0-9-]*[a-zA-Z0-9])?(\\.[a-zA-Z0-9]([a-zA-Z0-9-]*[a-zA-Z0-9])?){2,}$`)","file":"cmd/provenance/publish_run.go","name":"nsidSyntax"}}
var nsidSyntax = regexp.MustCompile(`^[a-zA-Z]([a-zA-Z0-9-]*[a-zA-Z0-9])?(\.[a-zA-Z0-9]([a-zA-Z0-9-]*[a-zA-Z0-9])?){2,}$`)

// /sedum:setVar
// sedum:setFunction {"tier":"owned","record":"prov-2026-7163f9f2","gen":"sha256:4c6bc602935bc9046f542c490eed5787bcb5d3211e9cbc9bb29c42009c668f9a","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func publishSource(r *repo) (func(id string) *record.PublishedSource, []string) {\n\tvar notes []string\n\n\tout, err := r.git(\"remote\", \"get-url\", \"origin\")\n\tvar origin string\n\tif err == nil {\n\t\torigin = repoURL(out)\n\t}\n\tif origin == \"\" {\n\t\treturn nil, []string{\"no usable origin remote: the records are published without a source (git remote add origin <url>)\"}\n\t}\n\n\thead, err := r.git(\"rev-parse\", \"HEAD\")\n\tif err != nil {\n\t\thead = \"\"\n\t}\n\n\tstatus, err := r.git(\"status\", \"--porcelain\", \"--untracked-files=no\")\n\tif err == nil && status != \"\" {\n\t\tnotes = append(notes, \"working tree has uncommitted changes; the source commit is still HEAD\")\n\t}\n\n\treturn func(id string) *record.PublishedSource {\n\t\treturn &record.PublishedSource{\n\t\t\tRepo:   origin,\n\t\t\tPath:   r.Config.Dir + \"/\" + id + \".yml\",\n\t\t\tCommit: head,\n\t\t}\n\t}, notes\n}","file":"cmd/provenance/publish_run.go","name":"publishSource"}}
func publishSource(r *repo) (func(id string) *record.PublishedSource, []string) {
	var notes []string

	out, err := r.git("remote", "get-url", "origin")
	var origin string
	if err == nil {
		origin = repoURL(out)
	}
	if origin == "" {
		return nil, []string{"no usable origin remote: the records are published without a source (git remote add origin <url>)"}
	}

	head, err := r.git("rev-parse", "HEAD")
	if err != nil {
		head = ""
	}

	status, err := r.git("status", "--porcelain", "--untracked-files=no")
	if err == nil && status != "" {
		notes = append(notes, "working tree has uncommitted changes; the source commit is still HEAD")
	}

	return func(id string) *record.PublishedSource {
		return &record.PublishedSource{
			Repo:   origin,
			Path:   r.Config.Dir + "/" + id + ".yml",
			Commit: head,
		}
	}, notes
}

// /sedum:setFunction
// sedum:setVar {"tier":"owned","record":"prov-2026-912be1ee","gen":"sha256:95eaf8c3c73641e880a3faf096fe3628552463a8b15fb28351ee19f70c8d98cd","pkg":"sha256:c7b9d9ac1655e8211b8b2ca8c49cf34b8d510d8d823d4f5eb60ec7e1018aa2eb","kwargs":{"code":"var rkeySyntax = regexp.MustCompile(`^[A-Za-z0-9._~:-]{1,512}$`)","file":"cmd/provenance/publish_run.go","name":"rkeySyntax"}}
var rkeySyntax = regexp.MustCompile(`^[A-Za-z0-9._~:-]{1,512}$`)

// /sedum:setVar
// sedum:setFunction {"tier":"owned","record":"prov-2026-b8cb30d5","gen":"sha256:0604e95152d8f50c4c7576551c6513df7c9b67351e2b87af6e518c82a1684cd0","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func repoURL(remote string) string {\n\tremote = strings.TrimSpace(remote)\n\tvar host, path string\n\tif strings.Contains(remote, \"://\") {\n\t\tu, err := url.Parse(remote)\n\t\tif err != nil {\n\t\t\treturn \"\"\n\t\t}\n\t\thost = u.Hostname()\n\t\tpath = u.Path\n\t} else if before, after, ok := strings.Cut(remote, \":\"); ok && strings.Contains(before, \"@\") {\n\t\thost = before[strings.LastIndex(before, \"@\")+1:]\n\t\tpath = after\n\t}\n\tpath = strings.Trim(path, \"/\")\n\tpath = strings.TrimSuffix(path, \".git\")\n\tif host == \"\" || path == \"\" {\n\t\treturn \"\"\n\t}\n\treturn \"https://\" + host + \"/\" + path\n}","file":"cmd/provenance/publish_run.go","name":"repoURL"}}
func repoURL(remote string) string {
	remote = strings.TrimSpace(remote)
	var host, path string
	if strings.Contains(remote, "://") {
		u, err := url.Parse(remote)
		if err != nil {
			return ""
		}
		host = u.Hostname()
		path = u.Path
	} else if before, after, ok := strings.Cut(remote, ":"); ok && strings.Contains(before, "@") {
		host = before[strings.LastIndex(before, "@")+1:]
		path = after
	}
	path = strings.Trim(path, "/")
	path = strings.TrimSuffix(path, ".git")
	if host == "" || path == "" {
		return ""
	}
	return "https://" + host + "/" + path
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-c9353bb5","gen":"sha256:40661903dd4b8b80184335c4c491fc384e5712936e242c2a2803cd8da7e3a9d3","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func writePlan(out io.Writer, plan *remote.PublishPlan) {\n\tunpublished, localAhead, serverAhead, diverged, notPublishable := 0, 0, 0, 0, 0\n\tfor _, p := range plan.Problems {\n\t\tif p.Kind == \"invalid id\" {\n\t\t\tfmt.Fprintf(out, \"  %-12s %s: %s\\n\", \"not publishable\", p.ID, p.Reason)\n\t\t\tnotPublishable++\n\t\t}\n\t}\n\tfor _, item := range plan.Items {\n\t\tdesc := headUpdate(item)\n\t\tlabel := \"local ahead\"\n\t\tif item.HeadVersion == 0 {\n\t\t\tlabel = \"unpublished\"\n\t\t\tunpublished++\n\t\t} else {\n\t\t\tlocalAhead++\n\t\t}\n\t\tfmt.Fprintf(out, \"  %-12s %s  %s (%s)\\n\", label, item.ID, item.Title, desc)\n\t}\n\tfor _, id := range plan.UpToDate {\n\t\tfmt.Fprintf(out, \"  %-12s %s\\n\", \"up to date\", id)\n\t}\n\tfor _, p := range plan.Problems {\n\t\tif p.Kind == \"invalid id\" {\n\t\t\tcontinue\n\t\t}\n\t\tif p.Kind == \"server ahead\" {\n\t\t\tserverAhead++\n\t\t} else {\n\t\t\tdiverged++\n\t\t}\n\t\tfmt.Fprintf(out, \"  %-12s %s: %s\\n\", p.Kind, p.ID, p.Reason)\n\t}\n\tfor _, s := range plan.Skipped {\n\t\tfmt.Fprintf(out, \"  %-12s %s\\n\", \"skipped\", s)\n\t}\n\tfmt.Fprintf(out, \"\\n%d unpublished, %d local ahead, %d up to date, %d server ahead, %d diverged, %d not publishable, %d skipped\\n\\n\", unpublished, localAhead, len(plan.UpToDate), serverAhead, diverged, notPublishable, len(plan.Skipped))\n}","file":"cmd/provenance/publish_run.go","name":"writePlan"}}
func writePlan(out io.Writer, plan *remote.PublishPlan) {
	unpublished, localAhead, serverAhead, diverged, notPublishable := 0, 0, 0, 0, 0
	for _, p := range plan.Problems {
		if p.Kind == "invalid id" {
			fmt.Fprintf(out, "  %-12s %s: %s\n", "not publishable", p.ID, p.Reason)
			notPublishable++
		}
	}
	for _, item := range plan.Items {
		desc := headUpdate(item)
		label := "local ahead"
		if item.HeadVersion == 0 {
			label = "unpublished"
			unpublished++
		} else {
			localAhead++
		}
		fmt.Fprintf(out, "  %-12s %s  %s (%s)\n", label, item.ID, item.Title, desc)
	}
	for _, id := range plan.UpToDate {
		fmt.Fprintf(out, "  %-12s %s\n", "up to date", id)
	}
	for _, p := range plan.Problems {
		if p.Kind == "invalid id" {
			continue
		}
		if p.Kind == "server ahead" {
			serverAhead++
		} else {
			diverged++
		}
		fmt.Fprintf(out, "  %-12s %s: %s\n", p.Kind, p.ID, p.Reason)
	}
	for _, s := range plan.Skipped {
		fmt.Fprintf(out, "  %-12s %s\n", "skipped", s)
	}
	fmt.Fprintf(out, "\n%d unpublished, %d local ahead, %d up to date, %d server ahead, %d diverged, %d not publishable, %d skipped\n\n", unpublished, localAhead, len(plan.UpToDate), serverAhead, diverged, notPublishable, len(plan.Skipped))
}

// /sedum:setFunction
// sedum:setVar {"tier":"owned","record":"prov-2026-ff747ab8","gen":"sha256:d6039e89bf08409719aebdaab3e71f5a4b4d1d73dcbd675b2fc5071b5e3607a2","pkg":"sha256:c7b9d9ac1655e8211b8b2ca8c49cf34b8d510d8d823d4f5eb60ec7e1018aa2eb","kwargs":{"code":"var handleSyntax = regexp.MustCompile(`^(?:[A-Za-z0-9](?:[A-Za-z0-9-]{0,61}[A-Za-z0-9])?\\.)+[A-Za-z](?:[A-Za-z0-9-]{0,61}[A-Za-z0-9])?$`)","file":"cmd/provenance/publish_run.go","name":"handleSyntax"}}
var handleSyntax = regexp.MustCompile(`^(?:[A-Za-z0-9](?:[A-Za-z0-9-]{0,61}[A-Za-z0-9])?\.)+[A-Za-z](?:[A-Za-z0-9-]{0,61}[A-Za-z0-9])?$`)

// /sedum:setVar
