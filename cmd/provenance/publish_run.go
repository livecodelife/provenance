package main

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-83d57ec1","gen":"sha256:17956fb64b77b361bcbf762257b64c9341f58a44819468b083c8b9e970540a94","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/publish_run.go","path":"fmt"}}
import "fmt"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-83d57ec1","gen":"sha256:17956fb64b77b361bcbf762257b64c9341f58a44819468b083c8b9e970540a94","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/publish_run.go","path":"io"}}
import "io"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-83d57ec1","gen":"sha256:17956fb64b77b361bcbf762257b64c9341f58a44819468b083c8b9e970540a94","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/publish_run.go","path":"strconv"}}
import "strconv"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-83d57ec1","gen":"sha256:17956fb64b77b361bcbf762257b64c9341f58a44819468b083c8b9e970540a94","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/publish_run.go","path":"strings"}}
import "strings"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-83d57ec1","gen":"sha256:17956fb64b77b361bcbf762257b64c9341f58a44819468b083c8b9e970540a94","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/publish_run.go","path":"github.com/spf13/cobra"}}
import "github.com/spf13/cobra"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-83d57ec1","gen":"sha256:17956fb64b77b361bcbf762257b64c9341f58a44819468b083c8b9e970540a94","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/publish_run.go","path":"github.com/livecodelife/provenance/remote"}}
import "github.com/livecodelife/provenance/remote"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-83d57ec1","gen":"sha256:17956fb64b77b361bcbf762257b64c9341f58a44819468b083c8b9e970540a94","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/publish_run.go","path":"net/url"}}
import "net/url"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-83d57ec1","gen":"sha256:17956fb64b77b361bcbf762257b64c9341f58a44819468b083c8b9e970540a94","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/publish_run.go","path":"github.com/livecodelife/provenance/record"}}
import "github.com/livecodelife/provenance/record"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-83d57ec1","gen":"sha256:17956fb64b77b361bcbf762257b64c9341f58a44819468b083c8b9e970540a94","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/publish_run.go","path":"errors"}}
import "errors"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-83d57ec1","gen":"sha256:17956fb64b77b361bcbf762257b64c9341f58a44819468b083c8b9e970540a94","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/publish_run.go","path":"regexp"}}
import "regexp"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"tier":"owned","record":"prov-2026-83d57ec1","gen":"sha256:17956fb64b77b361bcbf762257b64c9341f58a44819468b083c8b9e970540a94","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func runPublish(out io.Writer, cmd *cobra.Command, args []string) error {\n\tall, _ := cmd.Flags().GetBool(\"all\")\n\tyes, _ := cmd.Flags().GetBool(\"yes\")\n\tplanOnly, _ := cmd.Flags().GetBool(\"plan\")\n\tproject, _ := cmd.Flags().GetString(\"project\")\n\tif project != \"\" {\n\t\tif err := checkProject(project); err != nil {\n\t\t\treturn fmt.Errorf(\"--project %q: %w\", project, err)\n\t\t}\n\t}\n\tr, err := openRepo(cmd)\n\tif err != nil {\n\t\treturn err\n\t}\n\tids, err := recordIDs(r, args, all)\n\tif err != nil {\n\t\treturn err\n\t}\n\tclient, err := apiClient()\n\tif err != nil {\n\t\treturn err\n\t}\n\tctx := cmd.Context()\n\tsource, notes := publishSource(r)\n\tplan, err := remote.PlanPublishWith(ctx, r.Store, ids, client, remote.HeadWant{Project: project, Source: source})\n\tif err != nil {\n\t\treturn err\n\t}\n\tfor _, n := range notes {\n\t\tfmt.Fprintf(out, \"  note: %s\\n\", n)\n\t}\n\tif planOnly {\n\t\tfmt.Fprintln(out)\n\t\twritePlan(out, plan)\n\t\treturn nil\n\t}\n\tfor _, s := range plan.Skipped {\n\t\tfmt.Fprintf(out, \"  skipped  %s\\n\", s)\n\t}\n\tfor _, p := range plan.Problems {\n\t\tfmt.Fprintf(out, \"  skipped  %s: %s\\n\", p.ID, p.Reason)\n\t}\n\tskipped := fmt.Errorf(\"%d record(s) were skipped (see above); the others were handled\", len(plan.Problems))\n\tif plan.Empty() {\n\t\tfmt.Fprintln(out, \"\\nNothing to publish: the server already has every version and status.\")\n\t\tfmt.Fprintln(out)\n\t\tif len(plan.Problems) > 0 {\n\t\t\treturn skipped\n\t\t}\n\t\treturn nil\n\t}\n\tfmt.Fprintf(out, \"\\nThese become PUBLIC on AT Proto (anyone can read them):\\n\\n\")\n\tfor _, it := range plan.Items {\n\t\tif len(it.Snaps) == 0 {\n\t\t\tfmt.Fprintf(out, \"  %s  %s\\n      %s (no new version)\\n\", it.ID, it.Title, headUpdate(it))\n\t\t\tcontinue\n\t\t}\n\t\tnums := make([]string, len(it.Snaps))\n\t\tfor i, s := range it.Snaps {\n\t\t\tnums[i] = \"v\" + strconv.Itoa(s.Version)\n\t\t}\n\t\tfmt.Fprintf(out, \"  %s  %s\\n      versions %s, status %s\\n\", it.ID, it.Title, strings.Join(nums, \", \"), it.Status)\n\t}\n\tfmt.Fprintln(out)\n\tif !yes && !confirm(cmd, out, \"Publish these? [y/N] \") {\n\t\treturn fmt.Errorf(\"not confirmed, nothing was published (run interactively, or pass --yes)\")\n\t}\n\tcommitOf := func(id string, version int) string {\n\t\tsha, err := r.git(\"log\", \"--all\", \"-E\", \"-n\", \"1\", \"--format=%H\", \"--grep\", `\\[`+id+`@`+strconv.Itoa(version)+`\\]`)\n\t\tif err != nil {\n\t\t\treturn \"\"\n\t\t}\n\t\treturn strings.TrimSpace(sha)\n\t}\n\twritten := 0\n\terr = remote.ApplyPublish(ctx, plan, client, project, commitOf, func(id string, version int) {\n\t\twritten++\n\t\tif version == 0 {\n\t\t\tfmt.Fprintf(out, \"  published  %s head\\n\", id)\n\t\t\treturn\n\t\t}\n\t\tfmt.Fprintf(out, \"  published  %s v%d\\n\", id, version)\n\t})\n\tif err != nil {\n\t\treturn err\n\t}\n\tfmt.Fprintf(out, \"\\n✓ Published %d write(s).\\n\\n\", written)\n\tif len(plan.Problems) > 0 {\n\t\treturn skipped\n\t}\n\treturn nil\n}","file":"cmd/provenance/publish_run.go","name":"runPublish"}}
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
	ctx := cmd.Context()
	source, notes := publishSource(r)
	plan, err := remote.PlanPublishWith(ctx, r.Store, ids, client, remote.HeadWant{Project: project, Source: source})
	if err != nil {
		return err
	}
	for _, n := range notes {
		fmt.Fprintf(out, "  note: %s\n", n)
	}
	if planOnly {
		fmt.Fprintln(out)
		writePlan(out, plan)
		return nil
	}
	for _, s := range plan.Skipped {
		fmt.Fprintf(out, "  skipped  %s\n", s)
	}
	for _, p := range plan.Problems {
		fmt.Fprintf(out, "  skipped  %s: %s\n", p.ID, p.Reason)
	}
	skipped := fmt.Errorf("%d record(s) were skipped (see above); the others were handled", len(plan.Problems))
	if plan.Empty() {
		fmt.Fprintln(out, "\nNothing to publish: the server already has every version and status.")
		fmt.Fprintln(out)
		if len(plan.Problems) > 0 {
			return skipped
		}
		return nil
	}
	fmt.Fprintf(out, "\nThese become PUBLIC on AT Proto (anyone can read them):\n\n")
	for _, it := range plan.Items {
		if len(it.Snaps) == 0 {
			fmt.Fprintf(out, "  %s  %s\n      %s (no new version)\n", it.ID, it.Title, headUpdate(it))
			continue
		}
		nums := make([]string, len(it.Snaps))
		for i, s := range it.Snaps {
			nums[i] = "v" + strconv.Itoa(s.Version)
		}
		fmt.Fprintf(out, "  %s  %s\n      versions %s, status %s\n", it.ID, it.Title, strings.Join(nums, ", "), it.Status)
	}
	fmt.Fprintln(out)
	if !yes && !confirm(cmd, out, "Publish these? [y/N] ") {
		return fmt.Errorf("not confirmed, nothing was published (run interactively, or pass --yes)")
	}
	commitOf := func(id string, version int) string {
		sha, err := r.git("log", "--all", "-E", "-n", "1", "--format=%H", "--grep", `\[`+id+`@`+strconv.Itoa(version)+`\]`)
		if err != nil {
			return ""
		}
		return strings.TrimSpace(sha)
	}
	written := 0
	err = remote.ApplyPublish(ctx, plan, client, project, commitOf, func(id string, version int) {
		written++
		if version == 0 {
			fmt.Fprintf(out, "  published  %s head\n", id)
			return
		}
		fmt.Fprintf(out, "  published  %s v%d\n", id, version)
	})
	if err != nil {
		return err
	}
	fmt.Fprintf(out, "\n✓ Published %d write(s).\n\n", written)
	if len(plan.Problems) > 0 {
		return skipped
	}
	return nil
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-83d57ec1","gen":"sha256:17956fb64b77b361bcbf762257b64c9341f58a44819468b083c8b9e970540a94","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"// repoURL turns a git remote URL into the https URL of its repository\n// (https://host/owner/name, credentials and a trailing .git dropped), or \"\"\n// when the remote is not a host/owner/name remote, as a local path is not.\nfunc repoURL(remote string) string {\n\tremote = strings.TrimSpace(remote)\n\tvar host, path string\n\tif strings.Contains(remote, \"://\") {\n\t\tu, err := url.Parse(remote)\n\t\tif err != nil {\n\t\t\treturn \"\"\n\t\t}\n\t\thost, path = u.Hostname(), u.Path\n\t} else if before, after, ok := strings.Cut(remote, \":\"); ok && strings.Contains(before, \"@\") {\n\t\thost, path = before[strings.LastIndex(before, \"@\")+1:], after\n\t}\n\tpath = strings.TrimSuffix(strings.Trim(path, \"/\"), \".git\")\n\tif host == \"\" || path == \"\" {\n\t\treturn \"\"\n\t}\n\treturn \"https://\" + host + \"/\" + path\n}","file":"cmd/provenance/publish_run.go","name":"repoURL"}}
// repoURL turns a git remote URL into the https URL of its repository
// (https://host/owner/name, credentials and a trailing .git dropped), or ""
// when the remote is not a host/owner/name remote, as a local path is not.
func repoURL(remote string) string {
	remote = strings.TrimSpace(remote)
	var host, path string
	if strings.Contains(remote, "://") {
		u, err := url.Parse(remote)
		if err != nil {
			return ""
		}
		host, path = u.Hostname(), u.Path
	} else if before, after, ok := strings.Cut(remote, ":"); ok && strings.Contains(before, "@") {
		host, path = before[strings.LastIndex(before, "@")+1:], after
	}
	path = strings.TrimSuffix(strings.Trim(path, "/"), ".git")
	if host == "" || path == "" {
		return ""
	}
	return "https://" + host + "/" + path
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-83d57ec1","gen":"sha256:17956fb64b77b361bcbf762257b64c9341f58a44819468b083c8b9e970540a94","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"// publishSource says where r's records live in git: the origin remote, the\n// record's file in the repo, and HEAD. The function is nil when there is no\n// usable origin remote. notes are things to tell the user.\nfunc publishSource(r *repo) (source func(id string) *record.PublishedSource, notes []string) {\n\torigin, err := r.git(\"remote\", \"get-url\", \"origin\")\n\trepoAt := \"\"\n\tif err == nil {\n\t\trepoAt = repoURL(origin)\n\t}\n\tif repoAt == \"\" {\n\t\treturn nil, []string{\"no usable origin remote: the records are published without a source (git remote add origin <url>)\"}\n\t}\n\tcommit, _ := r.git(\"rev-parse\", \"HEAD\")\n\tif dirty, err := r.git(\"status\", \"--porcelain\", \"--untracked-files=no\"); err == nil && dirty != \"\" {\n\t\tnotes = append(notes, \"the working tree has uncommitted changes; the source commit is still HEAD\")\n\t}\n\treturn func(id string) *record.PublishedSource {\n\t\treturn &record.PublishedSource{Repo: repoAt, Path: r.Config.Dir + \"/\" + id + \".yml\", Commit: commit}\n\t}, notes\n}","file":"cmd/provenance/publish_run.go","name":"publishSource"}}
// publishSource says where r's records live in git: the origin remote, the
// record's file in the repo, and HEAD. The function is nil when there is no
// usable origin remote. notes are things to tell the user.
func publishSource(r *repo) (source func(id string) *record.PublishedSource, notes []string) {
	origin, err := r.git("remote", "get-url", "origin")
	repoAt := ""
	if err == nil {
		repoAt = repoURL(origin)
	}
	if repoAt == "" {
		return nil, []string{"no usable origin remote: the records are published without a source (git remote add origin <url>)"}
	}
	commit, _ := r.git("rev-parse", "HEAD")
	if dirty, err := r.git("status", "--porcelain", "--untracked-files=no"); err == nil && dirty != "" {
		notes = append(notes, "the working tree has uncommitted changes; the source commit is still HEAD")
	}
	return func(id string) *record.PublishedSource {
		return &record.PublishedSource{Repo: repoAt, Path: r.Config.Dir + "/" + id + ".yml", Commit: commit}
	}, notes
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-83d57ec1","gen":"sha256:17956fb64b77b361bcbf762257b64c9341f58a44819468b083c8b9e970540a94","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"// writePlan prints the records whose id the service rejects, then one line per\n// record, then the totals. It writes nothing and asks nothing: it is what\n// publish --plan shows.\nfunc writePlan(out io.Writer, plan *remote.PublishPlan) {\n\tnotPublishable := 0\n\tfor _, p := range plan.Problems {\n\t\tif p.Kind == \"invalid id\" {\n\t\t\tnotPublishable++\n\t\t\tfmt.Fprintf(out, \"  %-12s %s: %s\\n\", \"not publishable\", p.ID, p.Reason)\n\t\t}\n\t}\n\tunpublished, ahead := 0, 0\n\tfor _, it := range plan.Items {\n\t\twhat := headUpdate(it)\n\t\tif n := len(it.Snaps); n == 1 {\n\t\t\twhat = \"1 version\"\n\t\t} else if n > 1 {\n\t\t\twhat = strconv.Itoa(n) + \" versions\"\n\t\t}\n\t\tkind := \"local ahead\"\n\t\tif it.HeadVersion == 0 {\n\t\t\tkind = \"unpublished\"\n\t\t\tunpublished++\n\t\t} else {\n\t\t\tahead++\n\t\t}\n\t\tfmt.Fprintf(out, \"  %-12s %s  %s (%s)\\n\", kind, it.ID, it.Title, what)\n\t}\n\tfor _, id := range plan.UpToDate {\n\t\tfmt.Fprintf(out, \"  %-12s %s\\n\", \"up to date\", id)\n\t}\n\tserverAhead, diverged := 0, 0\n\tfor _, p := range plan.Problems {\n\t\tif p.Kind == \"invalid id\" {\n\t\t\tcontinue\n\t\t}\n\t\tif p.Kind == \"server ahead\" {\n\t\t\tserverAhead++\n\t\t} else {\n\t\t\tdiverged++\n\t\t}\n\t\tfmt.Fprintf(out, \"  %-12s %s: %s\\n\", p.Kind, p.ID, p.Reason)\n\t}\n\tfor _, s := range plan.Skipped {\n\t\tfmt.Fprintf(out, \"  %-12s %s\\n\", \"skipped\", s)\n\t}\n\tfmt.Fprintf(out, \"\\n%d unpublished, %d local ahead, %d up to date, %d server ahead, %d diverged, %d not publishable, %d skipped\\n\\n\",\n\t\tunpublished, ahead, len(plan.UpToDate), serverAhead, diverged, notPublishable, len(plan.Skipped))\n}","file":"cmd/provenance/publish_run.go","name":"writePlan"}}
// writePlan prints the records whose id the service rejects, then one line per
// record, then the totals. It writes nothing and asks nothing: it is what
// publish --plan shows.
func writePlan(out io.Writer, plan *remote.PublishPlan) {
	notPublishable := 0
	for _, p := range plan.Problems {
		if p.Kind == "invalid id" {
			notPublishable++
			fmt.Fprintf(out, "  %-12s %s: %s\n", "not publishable", p.ID, p.Reason)
		}
	}
	unpublished, ahead := 0, 0
	for _, it := range plan.Items {
		what := headUpdate(it)
		if n := len(it.Snaps); n == 1 {
			what = "1 version"
		} else if n > 1 {
			what = strconv.Itoa(n) + " versions"
		}
		kind := "local ahead"
		if it.HeadVersion == 0 {
			kind = "unpublished"
			unpublished++
		} else {
			ahead++
		}
		fmt.Fprintf(out, "  %-12s %s  %s (%s)\n", kind, it.ID, it.Title, what)
	}
	for _, id := range plan.UpToDate {
		fmt.Fprintf(out, "  %-12s %s\n", "up to date", id)
	}
	serverAhead, diverged := 0, 0
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
	fmt.Fprintf(out, "\n%d unpublished, %d local ahead, %d up to date, %d server ahead, %d diverged, %d not publishable, %d skipped\n\n",
		unpublished, ahead, len(plan.UpToDate), serverAhead, diverged, notPublishable, len(plan.Skipped))
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-83d57ec1","gen":"sha256:17956fb64b77b361bcbf762257b64c9341f58a44819468b083c8b9e970540a94","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"// headUpdate says what a head-only update of it changes, for example\n// \"status → implemented, project → at://did:plc:x/dev.provisio.project/y\".\nfunc headUpdate(it remote.PublishItem) string {\n\tvar parts []string\n\tif it.HeadStatus != \"\" {\n\t\tparts = append(parts, \"status → \"+it.HeadStatus)\n\t}\n\tif it.SetProject != \"\" {\n\t\tparts = append(parts, \"project → \"+it.SetProject)\n\t}\n\tif it.SetSource != nil {\n\t\tparts = append(parts, \"source → \"+it.SetSource.Repo+\" \"+it.SetSource.Path)\n\t}\n\treturn strings.Join(parts, \", \")\n}","file":"cmd/provenance/publish_run.go","name":"headUpdate"}}
// headUpdate says what a head-only update of it changes, for example
// "status → implemented, project → at://did:plc:x/dev.provisio.project/y".
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
// sedum:setVar {"tier":"owned","record":"prov-2026-83d57ec1","gen":"sha256:17956fb64b77b361bcbf762257b64c9341f58a44819468b083c8b9e970540a94","pkg":"sha256:c7b9d9ac1655e8211b8b2ca8c49cf34b8d510d8d823d4f5eb60ec7e1018aa2eb","kwargs":{"code":"var (\n\tdidSyntax    = regexp.MustCompile(`^did:[a-z]+:[a-zA-Z0-9._:%-]*[a-zA-Z0-9._%-]$`)\n\thandleSyntax = regexp.MustCompile(`^([a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?\\.)+[a-zA-Z]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?$`)\n\tnsidSyntax   = regexp.MustCompile(`^[a-zA-Z]([a-zA-Z0-9-]*[a-zA-Z0-9])?(\\.[a-zA-Z0-9]([a-zA-Z0-9-]*[a-zA-Z0-9])?){2,}$`)\n\trkeySyntax   = regexp.MustCompile(`^[a-zA-Z0-9._~:-]{1,512}$`)\n)","file":"cmd/provenance/publish_run.go","name":"didSyntax"}}
var (
	didSyntax    = regexp.MustCompile(`^did:[a-z]+:[a-zA-Z0-9._:%-]*[a-zA-Z0-9._%-]$`)
	handleSyntax = regexp.MustCompile(`^([a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?$`)
	nsidSyntax   = regexp.MustCompile(`^[a-zA-Z]([a-zA-Z0-9-]*[a-zA-Z0-9])?(\.[a-zA-Z0-9]([a-zA-Z0-9-]*[a-zA-Z0-9])?){2,}$`)
	rkeySyntax   = regexp.MustCompile(`^[a-zA-Z0-9._~:-]{1,512}$`)
)

// /sedum:setVar
// sedum:setFunction {"tier":"owned","record":"prov-2026-83d57ec1","gen":"sha256:17956fb64b77b361bcbf762257b64c9341f58a44819468b083c8b9e970540a94","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"// checkProject says why s is not an AT URI the service accepts as a project\n// (the head lexicon's at-uri format): at://, a DID or handle, then at most a\n// collection NSID and a record key. nil means it is acceptable.\nfunc checkProject(s string) error {\n\trest, ok := strings.CutPrefix(s, \"at://\")\n\tswitch {\n\tcase !ok:\n\t\treturn errors.New(\"must start with at://, for example at://did:plc:abc123/dev.provisio.project/name\")\n\tcase strings.ContainsAny(s, \"<>{}\"):\n\t\treturn errors.New(\"looks like a placeholder; use your real DID, for example at://did:plc:abc123/dev.provisio.project/name\")\n\tcase strings.Contains(s, \" \"):\n\t\treturn errors.New(\"must not contain a space\")\n\tcase strings.Contains(s, \"?\"):\n\t\treturn errors.New(\"must not have a query (?)\")\n\tcase strings.Contains(s, \"#\"):\n\t\treturn errors.New(\"must not have a fragment (#)\")\n\t}\n\tfor _, r := range s {\n\t\tif r < 33 || r > 126 {\n\t\t\treturn errors.New(\"must be ASCII\")\n\t\t}\n\t}\n\tparts := strings.Split(rest, \"/\")\n\tif !didSyntax.MatchString(parts[0]) && !handleSyntax.MatchString(parts[0]) {\n\t\treturn fmt.Errorf(\"the authority %q is not a DID (did:plc:...) or a handle\", parts[0])\n\t}\n\tif strings.HasSuffix(s, \"/\") {\n\t\treturn errors.New(\"must not end with a slash\")\n\t}\n\tif len(parts) > 3 {\n\t\treturn errors.New(\"has too many segments: at://<authority>[/<collection>[/<record key>]]\")\n\t}\n\tif len(parts) > 1 && !nsidSyntax.MatchString(parts[1]) {\n\t\treturn fmt.Errorf(\"the collection %q is not an NSID such as dev.provisio.project\", parts[1])\n\t}\n\tif len(parts) > 2 && (!rkeySyntax.MatchString(parts[2]) || parts[2] == \".\" || parts[2] == \"..\") {\n\t\treturn fmt.Errorf(\"the record key %q is not valid (ASCII letters, digits and . _ ~ : -)\", parts[2])\n\t}\n\treturn nil\n}","file":"cmd/provenance/publish_run.go","name":"checkProject"}}
// checkProject says why s is not an AT URI the service accepts as a project
// (the head lexicon's at-uri format): at://, a DID or handle, then at most a
// collection NSID and a record key. nil means it is acceptable.
func checkProject(s string) error {
	rest, ok := strings.CutPrefix(s, "at://")
	switch {
	case !ok:
		return errors.New("must start with at://, for example at://did:plc:abc123/dev.provisio.project/name")
	case strings.ContainsAny(s, "<>{}"):
		return errors.New("looks like a placeholder; use your real DID, for example at://did:plc:abc123/dev.provisio.project/name")
	case strings.Contains(s, " "):
		return errors.New("must not contain a space")
	case strings.Contains(s, "?"):
		return errors.New("must not have a query (?)")
	case strings.Contains(s, "#"):
		return errors.New("must not have a fragment (#)")
	}
	for _, r := range s {
		if r < 33 || r > 126 {
			return errors.New("must be ASCII")
		}
	}
	parts := strings.Split(rest, "/")
	if !didSyntax.MatchString(parts[0]) && !handleSyntax.MatchString(parts[0]) {
		return fmt.Errorf("the authority %q is not a DID (did:plc:...) or a handle", parts[0])
	}
	if strings.HasSuffix(s, "/") {
		return errors.New("must not end with a slash")
	}
	if len(parts) > 3 {
		return errors.New("has too many segments: at://<authority>[/<collection>[/<record key>]]")
	}
	if len(parts) > 1 && !nsidSyntax.MatchString(parts[1]) {
		return fmt.Errorf("the collection %q is not an NSID such as dev.provisio.project", parts[1])
	}
	if len(parts) > 2 && (!rkeySyntax.MatchString(parts[2]) || parts[2] == "." || parts[2] == "..") {
		return fmt.Errorf("the record key %q is not valid (ASCII letters, digits and . _ ~ : -)", parts[2])
	}
	return nil
}

// /sedum:setFunction
