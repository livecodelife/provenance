package main

// sedum:anchor:imports
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-12b1ca97","gen":"sha256:3d5eb5aff0b580425966bd18cb1e038d9fef9018ce5f4611f46a803ad6ee5bc9","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/status_run.go","path":"encoding/json"},"kwargs":"sha256:74a0a342507b7d2133976aa161af2fffadcd3d5ff78c2bdda1d2f2be51eabf06","body":"sha256:872d7c3e9dee608e79a7f8ef57930d24b5e30f2a57d891a52549179ace45946a"}
import "encoding/json"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-12b1ca97","gen":"sha256:3d5eb5aff0b580425966bd18cb1e038d9fef9018ce5f4611f46a803ad6ee5bc9","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/status_run.go","path":"fmt"},"kwargs":"sha256:937182fefd86fa3f95ce93ac86d03dae0e229c2bf177fb630ca077402df3f334","body":"sha256:bcbac51883cbf8064e9196b371cb880cee429866652447ad557533d817857609"}
import "fmt"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-12b1ca97","gen":"sha256:3d5eb5aff0b580425966bd18cb1e038d9fef9018ce5f4611f46a803ad6ee5bc9","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/status_run.go","path":"io"},"kwargs":"sha256:4581fe515f5272ca7db8d639e5a67b5ce438ef7f2ce21804c5aa091f15d33315","body":"sha256:ef8e8abf4ef5ee40a19dc85b49bf09c5df1fe32abfe834e0ab4c2175baf80ddf"}
import "io"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-12b1ca97","gen":"sha256:3d5eb5aff0b580425966bd18cb1e038d9fef9018ce5f4611f46a803ad6ee5bc9","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/status_run.go","path":"sort"},"kwargs":"sha256:76089311753472bfcb5f3e0a50d902a8141d2a5171a640d9f9b897410604948a","body":"sha256:f197172617ae21ef23b1855b09aafb14571020ae81bff877d63d3ff81f4bf63c"}
import "sort"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-12b1ca97","gen":"sha256:3d5eb5aff0b580425966bd18cb1e038d9fef9018ce5f4611f46a803ad6ee5bc9","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/status_run.go","path":"strings"},"kwargs":"sha256:28fd8ac86b11531ff021b87f13ef40a81c8572034d13921f57b0113904776350","body":"sha256:a9af3c0bc4bf9342df9d7e89a324523fdba7941a9ca4190270e68ff3c4a25ea6"}
import "strings"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-12b1ca97","gen":"sha256:3d5eb5aff0b580425966bd18cb1e038d9fef9018ce5f4611f46a803ad6ee5bc9","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/status_run.go","path":"github.com/spf13/cobra"},"kwargs":"sha256:3d4dbedefc9bae00efac244125cec258757599b893d58d0aae93560811ac7f44","body":"sha256:f361c02d0f8e1999cf3fb715393e16310eeec9debd5aaa13936af7617e2b46aa"}
import "github.com/spf13/cobra"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-12b1ca97","gen":"sha256:3d5eb5aff0b580425966bd18cb1e038d9fef9018ce5f4611f46a803ad6ee5bc9","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/status_run.go","path":"github.com/livecodelife/provenance/record"},"kwargs":"sha256:ddd7f2659fd59242091696795f830b0fe07215cd35f882cf6bb02519408e9b42","body":"sha256:acbf8628061291d11aa39e9ea64ecc91095a117057cfe6f1958ef2c437fac003"}
import "github.com/livecodelife/provenance/record"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setType {"v":2,"tier":"owned","record":"prov-2026-12b1ca97","gen":"sha256:3d5eb5aff0b580425966bd18cb1e038d9fef9018ce5f4611f46a803ad6ee5bc9","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","id":{"file":"cmd/provenance/status_run.go","name":"statusSpec"},"kwargs":"sha256:908cf6a9470073e6473b9075332610852465d48a433d42e99bd9abbccca0dc19","body":"sha256:488e18bf625c6f16808e13eb8313ee74e2045a89563086e22d78fc7828f29dd6"}
type statusSpec struct {
	Path       string `json:"path"`
	Type       string `json:"type,omitempty"`
	RunCommand string `json:"run_command,omitempty"`
}

// /sedum:setType
// sedum:setType {"v":2,"tier":"owned","record":"prov-2026-12b1ca97","gen":"sha256:3d5eb5aff0b580425966bd18cb1e038d9fef9018ce5f4611f46a803ad6ee5bc9","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","id":{"file":"cmd/provenance/status_run.go","name":"statusRecord"},"kwargs":"sha256:c78570674461e3c4e72b655e029f6f27a6d17d8be4e062ec95509fbf3a21f3ca","body":"sha256:8632fea58fe99744d958f5537eeb18894f1569aed0d962b602528d851e2b6192"}
type statusRecord struct {
	ID                string       `json:"id"`
	Title             string       `json:"title"`
	Type              string       `json:"type"`
	Status            string       `json:"status"`
	Version           int          `json:"version"`
	Reason            string       `json:"reason"`
	ReasonKind        string       `json:"reason_kind,omitempty"`
	CreatedAt         string       `json:"created_at"`
	Author            string       `json:"author"`
	Aliases           []string     `json:"aliases,omitempty"`
	Intent            string       `json:"intent"`
	Constraints       []string     `json:"constraints"`
	AffectedScope     []string     `json:"affected_scope"`
	ForbiddenScope    []string     `json:"forbidden_scope"`
	Implements        string       `json:"implements,omitempty"`
	ImplementsVersion int          `json:"implements_version,omitempty"`
	Related           []string     `json:"related"`
	Tags              []string     `json:"tags"`
	AssociatedSpecs   []statusSpec `json:"associated_specs"`
	SealedAtSHA       string       `json:"sealed_at_sha,omitempty"`
}

// /sedum:setType
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-12b1ca97","gen":"sha256:3d5eb5aff0b580425966bd18cb1e038d9fef9018ce5f4611f46a803ad6ee5bc9","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/status_run.go","name":"toStatusRecord"},"kwargs":"sha256:b4c57a667ffe25d76e68aec8e4fe2f431ebd13d5c01a306c5005410913cec050","body":"sha256:834ab21a70897733b8ba23eb43779ebb5628a4dd8ef5a3932a3ac8acd4b01695"}
func toStatusRecord(r *record.Record) statusRecord {
	specs := []statusSpec{}
	for _, s := range r.AssociatedSpecs {
		specs = append(specs, statusSpec{Path: s.Path, Type: s.Type, RunCommand: s.RunCommand})
	}
	return statusRecord{
		ID:                r.ID,
		Title:             r.Title,
		Type:              string(r.EffectiveType()),
		Status:            string(r.Status),
		Version:           r.Version,
		Reason:            r.Reason,
		ReasonKind:        r.ReasonKind,
		CreatedAt:         r.CreatedAt,
		Author:            r.Author,
		Aliases:           append([]string{}, r.Aliases...),
		Intent:            r.Intent,
		Constraints:       append([]string{}, r.Constraints...),
		AffectedScope:     append([]string{}, r.AffectedScope...),
		ForbiddenScope:    append([]string{}, r.ForbiddenScope...),
		Implements:        r.Implements,
		ImplementsVersion: r.ImplementsVersion,
		Related:           append([]string{}, r.Related...),
		Tags:              append([]string{}, r.Tags...),
		AssociatedSpecs:   specs,
		SealedAtSHA:       r.SealedAtSHA,
	}
}

// /sedum:setFunction
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-12b1ca97","gen":"sha256:3d5eb5aff0b580425966bd18cb1e038d9fef9018ce5f4611f46a803ad6ee5bc9","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/status_run.go","name":"printTable"},"kwargs":"sha256:e919ebae2bca2f90ffc5a8d7999df5d2467c7fac642d68052527d92693b8cf2c","body":"sha256:8a7d10dfa5341d2f87543c18b281341126c78a35625e2ea2694a4a460113f27b"}
func printTable(out io.Writer, enforcement string, recs []*record.Record) {
	fmt.Fprint(out, "\nPROVENANCE RECORDS\n\n")
	if enforcement == "" {
		enforcement = "none"
	}
	fmt.Fprintf(out, "  Enforcement: %s\n\n", enforcement)
	fmt.Fprintf(out, "  %-26s %-14s %-45s %s\n", "ID", "STATUS", "TITLE", "LINESPECS")
	fmt.Fprint(out, "  "+strings.Repeat("-", 96)+"\n")
	openNoSpecs := 0
	for _, rec := range recs {
		id := fmt.Sprintf("%s@%d", rec.ID, rec.Version)
		status := string(rec.Status)
		if rec.Status == record.StatusOpen && len(rec.AssociatedSpecs) == 0 {
			status = "open ⚠"
			openNoSpecs++
		}
		title := rec.Title
		if len(title) > 42 {
			title = title[:39] + "..."
		}
		fmt.Fprintf(out, "  %-26s %-14s %-45s %d\n", id, status, title, len(rec.AssociatedSpecs))
	}
	if openNoSpecs > 0 {
		fmt.Fprintf(out, "\n⚠ %d open records have no associated LineSpecs\n", openNoSpecs)
	}
	fmt.Fprint(out, "\n")
}

// /sedum:setFunction
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-12b1ca97","gen":"sha256:3d5eb5aff0b580425966bd18cb1e038d9fef9018ce5f4611f46a803ad6ee5bc9","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/status_run.go","name":"printDetail"},"kwargs":"sha256:52fed9ebb6b90ee68036cb6252cb5bf59f478a7a57d08278dd91cb71b72f6f82","body":"sha256:bebc2015c8950915955c2c3a6b7546940814680c42f973751a6f0ed3bf85e96a"}
func printDetail(out io.Writer, rec *record.Record, all []*record.Record) {
	status := string(rec.Status)
	if rec.Status == record.StatusOpen && len(rec.AssociatedSpecs) == 0 {
		status = "open ⚠"
	}
	fmt.Fprintf(out, "\n%s@%d  ·  %s\n", rec.ID, rec.Version, status)
	fmt.Fprintf(out, "%s\n\n", strings.Repeat("─", 60))
	fmt.Fprintf(out, "Title:        %s\n", rec.Title)
	fmt.Fprintf(out, "Author:       %s\n", rec.Author)
	fmt.Fprintf(out, "Created:      %s\n", rec.CreatedAt)
	fmt.Fprintf(out, "Version:      %d", rec.Version)
	switch {
	case rec.ReasonKind == record.ReasonKindFix:
		fmt.Fprintf(out, " (fix: %s)", rec.Reason)
	case rec.Reason != "":
		fmt.Fprintf(out, " (%s)", rec.Reason)
	}
	fmt.Fprintln(out)
	if len(rec.Aliases) > 0 {
		fmt.Fprintf(out, "Aliases:      %s\n", strings.Join(rec.Aliases, ", "))
	}
	if rec.Status == record.StatusImplemented && rec.SealedAtSHA != "" {
		sha := rec.SealedAtSHA
		if len(sha) > 7 {
			sha = sha[:7]
		}
		fmt.Fprintf(out, "Sealed at:    %s\n", sha)
	}
	if rec.Implements != "" {
		line := rec.Implements
		if rec.ImplementsVersion != 0 {
			line = fmt.Sprintf("%s@%d", rec.Implements, rec.ImplementsVersion)
		}
		for _, o := range all {
			if o.ID == rec.Implements {
				line += "  " + o.Title
				break
			}
		}
		fmt.Fprintf(out, "Implements:   %s\n", line)
	} else {
		fmt.Fprintln(out, "Implements:   —")
	}
	var children []*record.Record
	for _, o := range all {
		if o.Implements == rec.ID {
			children = append(children, o)
		}
	}
	if len(children) == 0 {
		fmt.Fprintln(out, "Implemented by: —")
	} else {
		fmt.Fprintln(out, "Implemented by:")
		for _, o := range children {
			fmt.Fprintf(out, "  · %s  [%s] %s\n", o.ID, o.EffectiveType(), o.Title)
		}
	}
	if len(rec.Tags) > 0 {
		fmt.Fprintf(out, "Tags:         %s\n", strings.Join(rec.Tags, ", "))
	}
	fmt.Fprintln(out)
	if rec.Intent != "" {
		fmt.Fprintln(out, "Intent:")
		for _, line := range strings.Split(strings.TrimRight(rec.Intent, "\n"), "\n") {
			fmt.Fprintf(out, "  %s\n", line)
		}
		fmt.Fprintln(out)
	}
	section := func(title string, items []string) {
		if len(items) == 0 {
			return
		}
		fmt.Fprintln(out, title)
		for _, it := range items {
			fmt.Fprintf(out, "  · %s\n", it)
		}
		fmt.Fprintln(out)
	}
	section("Constraints:", rec.Constraints)
	section("Allowed Scope:", rec.AffectedScope)
	section("Forbidden Scope (explicit):", rec.ForbiddenScope)
	var paths []string
	for _, s := range rec.AssociatedSpecs {
		paths = append(paths, s.Path)
	}
	section("Associated Specs:", paths)
}

// /sedum:setFunction
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-12b1ca97","gen":"sha256:3d5eb5aff0b580425966bd18cb1e038d9fef9018ce5f4611f46a803ad6ee5bc9","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/status_run.go","name":"runStatus"},"kwargs":"sha256:f6ec433de453422b834a74965a430854ac41245c250053b87aa5ec76110645b4","body":"sha256:4516431fcd6ddfc4e3a296322f43c8e930ee99a615aa1079dc3bd27d7b86cf77"}
func runStatus(out io.Writer, cmd *cobra.Command, args []string) error {
	r, err := openRepo(cmd)
	if err != nil {
		return err
	}
	recordID, err := cmd.Flags().GetString("record")
	if err != nil {
		return err
	}
	filter, err := cmd.Flags().GetString("filter")
	if err != nil {
		return err
	}
	format, err := cmd.Flags().GetString("format")
	if err != nil {
		return err
	}
	if format != "" && format != "human" && format != "json" {
		return fmt.Errorf("unknown format %q (human|json)", format)
	}
	ids, err := r.Store.List()
	if err != nil {
		return err
	}
	sort.Strings(ids)
	var all []*record.Record
	for _, id := range ids {
		rec, _, err := r.Store.Load(id)
		if err != nil {
			return err
		}
		all = append(all, rec)
	}
	if recordID != "" {
		rec, _, err := r.resolve(recordID)
		if err != nil {
			return err
		}
		if format == "json" {
			enc := json.NewEncoder(out)
			enc.SetIndent("", "  ")
			return enc.Encode(map[string]statusRecord{"record": toStatusRecord(rec)})
		}
		printDetail(out, rec, all)
		return nil
	}
	var filtered []*record.Record
	switch {
	case filter == "":
		filtered = all
	case strings.HasPrefix(filter, "tag:"):
		name := strings.TrimPrefix(filter, "tag:")
		for _, rec := range all {
			for _, t := range rec.Tags {
				if t == name {
					filtered = append(filtered, rec)
					break
				}
			}
		}
	default:
		switch filter {
		case "draft", "open", "implemented", "deprecated":
			for _, rec := range all {
				if string(rec.Status) == filter {
					filtered = append(filtered, rec)
				}
			}
		default:
			return fmt.Errorf("unknown filter %q (draft|open|implemented|deprecated or tag:name)", filter)
		}
	}
	if format == "json" {
		items := make([]statusRecord, 0, len(filtered))
		for _, rec := range filtered {
			items = append(items, toStatusRecord(rec))
		}
		enc := json.NewEncoder(out)
		enc.SetIndent("", "  ")
		return enc.Encode(map[string][]statusRecord{"records": items})
	}
	printTable(out, r.Config.Enforcement, filtered)
	return nil
}

// /sedum:setFunction
