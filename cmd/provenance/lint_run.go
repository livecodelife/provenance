package main

// sedum:anchor:imports
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-a03e6bc3","gen":"sha256:0b9b6ef508b9729ba05441fd94364705dd7752053dab523430bedbe724f30b75","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/lint_run.go","path":"encoding/json"},"kwargs":"sha256:f7937ba90821d241a70b71ef5823d7c940e06d6fd6b266a084f0ec5b5c42e718","body":"sha256:872d7c3e9dee608e79a7f8ef57930d24b5e30f2a57d891a52549179ace45946a"}
import "encoding/json"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-a03e6bc3","gen":"sha256:0b9b6ef508b9729ba05441fd94364705dd7752053dab523430bedbe724f30b75","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/lint_run.go","path":"errors"},"kwargs":"sha256:131e8028eb897b536f38a6235cbc5361c8ff8464ae18040a3ec84ed654cc2490","body":"sha256:2b718237ba0a56467c8fd494cc131e942ef410e6228c03b644d70476c71862dc"}
import "errors"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-a03e6bc3","gen":"sha256:0b9b6ef508b9729ba05441fd94364705dd7752053dab523430bedbe724f30b75","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/lint_run.go","path":"fmt"},"kwargs":"sha256:8c7b23453788ed18062336d380f9859527c3405af0ac5ce9e38c06f76a351069","body":"sha256:bcbac51883cbf8064e9196b371cb880cee429866652447ad557533d817857609"}
import "fmt"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-a03e6bc3","gen":"sha256:0b9b6ef508b9729ba05441fd94364705dd7752053dab523430bedbe724f30b75","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/lint_run.go","path":"io"},"kwargs":"sha256:a715bc7038894193f4361bb2a991022ed0451e55dd756ba6ca68cf4b573add1f","body":"sha256:ef8e8abf4ef5ee40a19dc85b49bf09c5df1fe32abfe834e0ab4c2175baf80ddf"}
import "io"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-a03e6bc3","gen":"sha256:0b9b6ef508b9729ba05441fd94364705dd7752053dab523430bedbe724f30b75","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/lint_run.go","path":"github.com/spf13/cobra"},"kwargs":"sha256:f07a9e25645ec73449b90eb27cf9b820cd9035b85867937a715e1bbaa43e4bec","body":"sha256:f361c02d0f8e1999cf3fb715393e16310eeec9debd5aaa13936af7617e2b46aa"}
import "github.com/spf13/cobra"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-a03e6bc3","gen":"sha256:0b9b6ef508b9729ba05441fd94364705dd7752053dab523430bedbe724f30b75","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/lint_run.go","path":"github.com/livecodelife/provenance/lint"},"kwargs":"sha256:736a811db3f65ed6eb388e367a8bc4e88a4497d091d7f0f2fe155da24455f28d","body":"sha256:9abff5a2d82dd8a4177704699227186fed48d7e0d093d2d2d5c565b09ff787d7"}
import "github.com/livecodelife/provenance/lint"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setType {"v":2,"tier":"owned","record":"prov-2026-a03e6bc3","gen":"sha256:0b9b6ef508b9729ba05441fd94364705dd7752053dab523430bedbe724f30b75","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","id":{"file":"cmd/provenance/lint_run.go","name":"lintIssueJSON"},"kwargs":"sha256:c738a181e5a78d1c58f6c1e220e4d0971a86495a4d3332988e9ad2d73241db6d","body":"sha256:96a972552ba63e8d3c00801c17ff41d867a3d0dd1e52dfa9ac44d8704e17eb3f"}
type lintIssueJSON struct {
	RecordID string `json:"record_id"`
	Field    string `json:"field"`
	Message  string `json:"message"`
	Severity string `json:"severity"`
	Rule     string `json:"rule"`
}

// /sedum:setType
// sedum:setType {"v":2,"tier":"owned","record":"prov-2026-a03e6bc3","gen":"sha256:0b9b6ef508b9729ba05441fd94364705dd7752053dab523430bedbe724f30b75","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","id":{"file":"cmd/provenance/lint_run.go","name":"lintJSON"},"kwargs":"sha256:96a885b4abe60429637acb71267aad858888c62a77c52efcae21b935da6f90d4","body":"sha256:0cbda94cd3eff14cdd861dcc1aabf03ed68d9fdbebc244d7a614e2d8f63207e7"}
type lintJSON struct {
	Enforcement string          `json:"enforcement"`
	Total       int             `json:"total"`
	Passed      int             `json:"passed"`
	Warnings    int             `json:"warnings"`
	Errors      int             `json:"errors"`
	Issues      []lintIssueJSON `json:"issues"`
}

// /sedum:setType
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-a03e6bc3","gen":"sha256:0b9b6ef508b9729ba05441fd94364705dd7752053dab523430bedbe724f30b75","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/lint_run.go","name":"lintRepo"},"kwargs":"sha256:c60db182caad779f67de8581a53b211dae2c2b5dca6de57b28844b48b98a5acc","body":"sha256:a71aec413abb6eef36c0eb4efa6217d97ae59795a1f3c31ce1f1c97f9b32db52"}
func lintRepo(r *repo, enforcement string) (*lint.Result, error) {
	return lintRepoFocus(r, enforcement, nil)
}

// /sedum:setFunction
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-a03e6bc3","gen":"sha256:0b9b6ef508b9729ba05441fd94364705dd7752053dab523430bedbe724f30b75","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/lint_run.go","name":"recordLint"},"kwargs":"sha256:042eb97d01d8ee0a8e54e4d7045f91beac7cf00a767619e931803dafd401a536","body":"sha256:03af1d569d2397a07c2a6db70ddfde0d6a51795576da0dd1a4413f358741b4d0"}
func recordLint(res *lint.Result, id string) *lint.Result {
	out := &lint.Result{Records: 1, Enforcement: res.Enforcement}
	for _, issue := range res.Issues {
		if issue.RecordID != id {
			continue
		}
		out.Issues = append(out.Issues, issue)
		switch issue.Severity {
		case lint.SeverityError:
			out.ErrorCount++
		case lint.SeverityWarning:
			out.WarningCount++
		case lint.SeverityHint:
			out.HintCount++
		}
	}
	if out.ErrorCount == 0 && out.WarningCount == 0 {
		out.PassedCount = 1
	}
	return out
}

// /sedum:setFunction
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-a03e6bc3","gen":"sha256:0b9b6ef508b9729ba05441fd94364705dd7752053dab523430bedbe724f30b75","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/lint_run.go","name":"runLint"},"kwargs":"sha256:5fba3419c6c5f0abe53c7171a4cd1b9cdb8287966396772242da6baa6f7e40ba","body":"sha256:ad264539b18c4042b45a77b0f014b2474f9a110377e7412c59cebaaabd82b379"}
func runLint(out io.Writer, cmd *cobra.Command, args []string) error {
	r, err := openRepo(cmd)
	if err != nil {
		return err
	}
	record, _ := cmd.Flags().GetString("record")
	enforcement, _ := cmd.Flags().GetString("enforcement")
	format, _ := cmd.Flags().GetString("format")
	warnOnly, _ := cmd.Flags().GetBool("warn")
	infoOnly, _ := cmd.Flags().GetBool("info")
	showAll, _ := cmd.Flags().GetBool("all")
	if format != "" && format != "human" && format != "json" && format != "sarif" {
		return fmt.Errorf("unknown format %q (human|json|sarif)", format)
	}
	var focus map[string]bool
	if record != "" {
		if loaded, _, lerr := r.Store.Load(record); lerr == nil {
			focus = map[string]bool{loaded.ID: true}
		}
	}
	res, err := lintRepoFocus(r, enforcement, focus)
	if err != nil {
		return err
	}
	if record != "" {
		rec, _, err := r.resolve(record)
		if err != nil {
			return err
		}
		res = recordLint(res, rec.ID)
	}
	if format == "sarif" {
		if err := printSarif(out, buildSarif(r, lintFindings(r, res))); err != nil {
			return err
		}
		if res.ErrorCount > 0 {
			return errors.New("lint failed")
		}
		return nil
	}
	total := res.PassedCount + res.WarningCount + res.ErrorCount
	if format == "json" {
		issues := []lintIssueJSON{}
		for _, issue := range res.Issues {
			issues = append(issues, lintIssueJSON{
				RecordID: issue.RecordID,
				Field:    issue.Field,
				Message:  issue.Message,
				Severity: string(issue.Severity),
				Rule:     issue.Rule,
			})
		}
		enc := json.NewEncoder(out)
		enc.SetIndent("", "  ")
		if err := enc.Encode(lintJSON{
			Enforcement: res.Enforcement,
			Total:       total,
			Passed:      res.PassedCount,
			Warnings:    res.WarningCount,
			Errors:      res.ErrorCount,
			Issues:      issues,
		}); err != nil {
			return err
		}
	} else {
		show := func(sev lint.Severity) bool {
			switch {
			case showAll:
				return true
			case warnOnly:
				return sev == lint.SeverityWarning
			case infoOnly:
				return sev == lint.SeverityHint
			default:
				return sev == lint.SeverityError
			}
		}
		fmt.Fprintf(out, "\n✓ Linting %d provenance records (enforcement: %s)\n\n", total, res.Enforcement)
		_, warnings := loadShared(r)
		for _, w := range warnings {
			fmt.Fprintf(out, "  ⚠ %s\n", w)
		}
		for _, issue := range res.Issues {
			if !show(issue.Severity) {
				continue
			}
			symbol := "→"
			switch issue.Severity {
			case lint.SeverityError:
				symbol = "✗"
			case lint.SeverityWarning:
				symbol = "⚠"
			}
			field := ""
			if issue.Field != "" {
				field = "(" + issue.Field + ")"
			}
			fmt.Fprintf(out, "  %s %-15s %s %s\n", symbol, issue.RecordID, issue.Message, field)
		}
		fmt.Fprintf(out, "\n✓ %d passed  ⚠ %d warnings  ✗ %d errors\n", res.PassedCount, res.WarningCount, res.ErrorCount)
		if res.ErrorCount > 0 && res.Enforcement == "strict" {
			fmt.Fprint(out, "\n  Hint: Associate LineSpec files with open records or set status to\n       'implemented' if proof already exists outside of LineSpec.\n")
		}
		fmt.Fprint(out, "\n")
	}
	if res.ErrorCount > 0 {
		return errors.New("lint failed")
	}
	return nil
}

// /sedum:setFunction
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-a03e6bc3","gen":"sha256:0b9b6ef508b9729ba05441fd94364705dd7752053dab523430bedbe724f30b75","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/lint_run.go","name":"lintRepoFocus"},"kwargs":"sha256:9ab9e9011f57555d6bb70524adcaf66afa74cb3c0d88704841f3235ae4ca5222","body":"sha256:36513597447fad183bb44d269caf39302a28a0f60d9829393fa14d900e8856e7"}
func lintRepoFocus(r *repo, enforcement string, focus map[string]bool) (*lint.Result, error) {
	if enforcement == "" {
		enforcement = r.Config.Enforcement
	}
	if enforcement == "" {
		enforcement = "none"
	}
	shared, _ := loadShared(r)
	return lint.Lint(r.Store, lint.Options{
		Shared:       shared,
		Root:         r.Root,
		Enforcement:  enforcement,
		ExcludePaths: r.Config.ExcludePaths,
		SharedRepos:  sharedNames(r),
		ManifestPath: manifestPath,
		Focus:        focus,
	})
}

// /sedum:setFunction
