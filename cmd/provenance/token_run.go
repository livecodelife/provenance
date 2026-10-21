package main

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-7f96263a","gen":"sha256:7d219ded22bdca52e8acfbc153485b2bb82cf7673bdc34a38acd2ef12afcf9dd","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/token_run.go","path":"encoding/json"}}
import "encoding/json"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-7f96263a","gen":"sha256:7d219ded22bdca52e8acfbc153485b2bb82cf7673bdc34a38acd2ef12afcf9dd","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/token_run.go","path":"errors"}}
import "errors"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-7f96263a","gen":"sha256:7d219ded22bdca52e8acfbc153485b2bb82cf7673bdc34a38acd2ef12afcf9dd","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/token_run.go","path":"fmt"}}
import "fmt"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-7f96263a","gen":"sha256:7d219ded22bdca52e8acfbc153485b2bb82cf7673bdc34a38acd2ef12afcf9dd","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/token_run.go","path":"io"}}
import "io"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-7f96263a","gen":"sha256:7d219ded22bdca52e8acfbc153485b2bb82cf7673bdc34a38acd2ef12afcf9dd","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/token_run.go","path":"path"}}
import "path"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-7f96263a","gen":"sha256:7d219ded22bdca52e8acfbc153485b2bb82cf7673bdc34a38acd2ef12afcf9dd","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/token_run.go","path":"path/filepath"}}
import "path/filepath"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-7f96263a","gen":"sha256:7d219ded22bdca52e8acfbc153485b2bb82cf7673bdc34a38acd2ef12afcf9dd","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/token_run.go","path":"strings"}}
import "strings"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-7f96263a","gen":"sha256:7d219ded22bdca52e8acfbc153485b2bb82cf7673bdc34a38acd2ef12afcf9dd","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/token_run.go","path":"text/tabwriter"}}
import "text/tabwriter"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-7f96263a","gen":"sha256:7d219ded22bdca52e8acfbc153485b2bb82cf7673bdc34a38acd2ef12afcf9dd","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/token_run.go","path":"github.com/spf13/cobra"}}
import "github.com/spf13/cobra"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-7f96263a","gen":"sha256:7d219ded22bdca52e8acfbc153485b2bb82cf7673bdc34a38acd2ef12afcf9dd","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/token_run.go","path":"github.com/livecodelife/provenance/auth"}}
import "github.com/livecodelife/provenance/auth"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setConst {"tier":"owned","record":"prov-2026-7f96263a","gen":"sha256:7d219ded22bdca52e8acfbc153485b2bb82cf7673bdc34a38acd2ef12afcf9dd","pkg":"sha256:51be1c18998d63cc558c253e211a6bd188d96cb251dad68c6ea88e03c2d5cc4c","kwargs":{"code":"// msgTokenSession is shown when the service refuses a token command because the\n// caller is a personal token, which cannot manage tokens.\nconst msgTokenSession = \"Managing tokens needs a signed-in session, not a personal token. Run provenance login (and unset PROVISIO_TOKEN).\"","file":"cmd/provenance/token_run.go","name":"msgTokenSession"}}
// msgTokenSession is shown when the service refuses a token command because the
// caller is a personal token, which cannot manage tokens.
const msgTokenSession = "Managing tokens needs a signed-in session, not a personal token. Run provenance login (and unset PROVISIO_TOKEN)."

// /sedum:setConst
// sedum:setVar {"tier":"owned","record":"prov-2026-7f96263a","gen":"sha256:7d219ded22bdca52e8acfbc153485b2bb82cf7673bdc34a38acd2ef12afcf9dd","pkg":"sha256:c7b9d9ac1655e8211b8b2ca8c49cf34b8d510d8d823d4f5eb60ec7e1018aa2eb","kwargs":{"code":"// tokenScopes are the scopes the service accepts.\nvar tokenScopes = map[string]bool{\"records:read\": true, \"records:write\": true}","file":"cmd/provenance/token_run.go","name":"tokenScopes"}}
// tokenScopes are the scopes the service accepts.
var tokenScopes = map[string]bool{"records:read": true, "records:write": true}

// /sedum:setVar
// sedum:setFunction {"tier":"owned","record":"prov-2026-7f96263a","gen":"sha256:7d219ded22bdca52e8acfbc153485b2bb82cf7673bdc34a38acd2ef12afcf9dd","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"// tokenError turns a failed /tokens call into the user-facing error: a 403 means\n// a personal token was used (exit status 3 with the way out), a service message is\n// shown with its code, and session failures are mapped by sessionError.\nfunc tokenError(err error) error {\n\tvar oe *auth.OAuthError\n\tif errors.As(err, &oe) {\n\t\tif oe.Status == 403 {\n\t\t\treturn &exitError{3, errors.New(msgTokenSession)}\n\t\t}\n\t\tif oe.Description != \"\" {\n\t\t\treturn fmt.Errorf(\"%s (HTTP %d): %s\", oe.Code, oe.Status, oe.Description)\n\t\t}\n\t}\n\treturn sessionError(err)\n}","file":"cmd/provenance/token_run.go","name":"tokenError"}}
// tokenError turns a failed /tokens call into the user-facing error: a 403 means
// a personal token was used (exit status 3 with the way out), a service message is
// shown with its code, and session failures are mapped by sessionError.
func tokenError(err error) error {
	var oe *auth.OAuthError
	if errors.As(err, &oe) {
		if oe.Status == 403 {
			return &exitError{3, errors.New(msgTokenSession)}
		}
		if oe.Description != "" {
			return fmt.Errorf("%s (HTTP %d): %s", oe.Code, oe.Status, oe.Description)
		}
	}
	return sessionError(err)
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-7f96263a","gen":"sha256:7d219ded22bdca52e8acfbc153485b2bb82cf7673bdc34a38acd2ef12afcf9dd","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"// defaultTokenName is ci-<repository>: the origin remote's repository name, else\n// the directory of the repository. Outside a repository there is no default.\nfunc defaultTokenName(cmd *cobra.Command) (string, error) {\n\tr, err := openRepo(cmd)\n\tif err != nil {\n\t\treturn \"\", errors.New(\"--name is required outside a repository\")\n\t}\n\ttop, err := r.git(\"rev-parse\", \"--show-toplevel\")\n\tif err != nil {\n\t\treturn \"\", errors.New(\"--name is required outside a repository\")\n\t}\n\tname := filepath.Base(strings.TrimSpace(top))\n\tif origin, err := r.git(\"remote\", \"get-url\", \"origin\"); err == nil {\n\t\tif u := repoURL(origin); u != \"\" {\n\t\t\tname = path.Base(u)\n\t\t}\n\t}\n\treturn \"ci-\" + name, nil\n}","file":"cmd/provenance/token_run.go","name":"defaultTokenName"}}
// defaultTokenName is ci-<repository>: the origin remote's repository name, else
// the directory of the repository. Outside a repository there is no default.
func defaultTokenName(cmd *cobra.Command) (string, error) {
	r, err := openRepo(cmd)
	if err != nil {
		return "", errors.New("--name is required outside a repository")
	}
	top, err := r.git("rev-parse", "--show-toplevel")
	if err != nil {
		return "", errors.New("--name is required outside a repository")
	}
	name := filepath.Base(strings.TrimSpace(top))
	if origin, err := r.git("remote", "get-url", "origin"); err == nil {
		if u := repoURL(origin); u != "" {
			name = path.Base(u)
		}
	}
	return "ci-" + name, nil
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-7f96263a","gen":"sha256:7d219ded22bdca52e8acfbc153485b2bb82cf7673bdc34a38acd2ef12afcf9dd","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"// tokenRequest reads and validates the create flags before any request is made.\nfunc tokenRequest(cmd *cobra.Command) (auth.TokenRequest, error) {\n\tname, _ := cmd.Flags().GetString(\"name\")\n\tscopes, _ := cmd.Flags().GetStringArray(\"scope\")\n\tdays, _ := cmd.Flags().GetInt(\"expires-in-days\")\n\tname = strings.TrimSpace(name)\n\tif name == \"\" && !cmd.Flags().Changed(\"name\") {\n\t\tvar err error\n\t\tif name, err = defaultTokenName(cmd); err != nil {\n\t\t\treturn auth.TokenRequest{}, err\n\t\t}\n\t}\n\tif len(name) < 1 || len(name) > 60 {\n\t\treturn auth.TokenRequest{}, errors.New(\"--name must be 1 to 60 characters\")\n\t}\n\tif len(scopes) == 0 {\n\t\tscopes = []string{\"records:write\"}\n\t}\n\tfor _, s := range scopes {\n\t\tif !tokenScopes[s] {\n\t\t\treturn auth.TokenRequest{}, fmt.Errorf(\"--scope %q: must be records:read or records:write\", s)\n\t\t}\n\t}\n\tif !cmd.Flags().Changed(\"expires-in-days\") {\n\t\tdays = 90\n\t}\n\tif days < 1 || days > 365 {\n\t\treturn auth.TokenRequest{}, errors.New(\"--expires-in-days must be 1 to 365\")\n\t}\n\treturn auth.TokenRequest{Name: name, Scopes: scopes, ExpiresInDays: days}, nil\n}","file":"cmd/provenance/token_run.go","name":"tokenRequest"}}
// tokenRequest reads and validates the create flags before any request is made.
func tokenRequest(cmd *cobra.Command) (auth.TokenRequest, error) {
	name, _ := cmd.Flags().GetString("name")
	scopes, _ := cmd.Flags().GetStringArray("scope")
	days, _ := cmd.Flags().GetInt("expires-in-days")
	name = strings.TrimSpace(name)
	if name == "" && !cmd.Flags().Changed("name") {
		var err error
		if name, err = defaultTokenName(cmd); err != nil {
			return auth.TokenRequest{}, err
		}
	}
	if len(name) < 1 || len(name) > 60 {
		return auth.TokenRequest{}, errors.New("--name must be 1 to 60 characters")
	}
	if len(scopes) == 0 {
		scopes = []string{"records:write"}
	}
	for _, s := range scopes {
		if !tokenScopes[s] {
			return auth.TokenRequest{}, fmt.Errorf("--scope %q: must be records:read or records:write", s)
		}
	}
	if !cmd.Flags().Changed("expires-in-days") {
		days = 90
	}
	if days < 1 || days > 365 {
		return auth.TokenRequest{}, errors.New("--expires-in-days must be 1 to 365")
	}
	return auth.TokenRequest{Name: name, Scopes: scopes, ExpiresInDays: days}, nil
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-7f96263a","gen":"sha256:7d219ded22bdca52e8acfbc153485b2bb82cf7673bdc34a38acd2ef12afcf9dd","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"// dash is \"-\" for an empty cell.\nfunc dash(s string) string {\n\tif s == \"\" {\n\t\treturn \"-\"\n\t}\n\treturn s\n}","file":"cmd/provenance/token_run.go","name":"dash"}}
// dash is "-" for an empty cell.
func dash(s string) string {
	if s == "" {
		return "-"
	}
	return s
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-7f96263a","gen":"sha256:7d219ded22bdca52e8acfbc153485b2bb82cf7673bdc34a38acd2ef12afcf9dd","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func runTokenCreate(out io.Writer, cmd *cobra.Command, args []string) error {\n\treq, err := tokenRequest(cmd)\n\tif err != nil {\n\t\treturn err\n\t}\n\tasJSON, _ := cmd.Flags().GetBool(\"json\")\n\tsecretOnly, _ := cmd.Flags().GetBool(\"secret-only\")\n\tm, err := sessionManager()\n\tif err != nil {\n\t\treturn err\n\t}\n\tctx := cmd.Context()\n\tvar tv *auth.TokenView\n\terr = m.Do(ctx, func(bearer string) error {\n\t\tvar err error\n\t\ttv, err = m.Client.CreateToken(ctx, bearer, req)\n\t\treturn err\n\t})\n\tif err != nil {\n\t\treturn tokenError(err)\n\t}\n\tswitch {\n\tcase secretOnly:\n\t\tfmt.Fprintln(out, tv.Token)\n\tcase asJSON:\n\t\tenc := json.NewEncoder(out)\n\t\tenc.SetIndent(\"\", \"  \")\n\t\treturn enc.Encode(tv)\n\tdefault:\n\t\tfmt.Fprintf(out, \"Created token %s (%s)\\n  scopes:  %s\\n  expires: %s\\n\\n%s\\n\\n\", tv.ID, tv.Name, strings.Join(tv.Scopes, \", \"), dash(tv.ExpiresAt), tv.Token)\n\t\tfmt.Fprintln(out, \"This secret will not be shown again. Store it now, for example:\")\n\t\tfmt.Fprintln(out, \"  provenance token create --name NAME --secret-only | gh secret set PROVISIO_TOKEN --env provisio-publish -R OWNER/REPO\")\n\t}\n\treturn nil\n}","file":"cmd/provenance/token_run.go","name":"runTokenCreate"}}
func runTokenCreate(out io.Writer, cmd *cobra.Command, args []string) error {
	req, err := tokenRequest(cmd)
	if err != nil {
		return err
	}
	asJSON, _ := cmd.Flags().GetBool("json")
	secretOnly, _ := cmd.Flags().GetBool("secret-only")
	m, err := sessionManager()
	if err != nil {
		return err
	}
	ctx := cmd.Context()
	var tv *auth.TokenView
	err = m.Do(ctx, func(bearer string) error {
		var err error
		tv, err = m.Client.CreateToken(ctx, bearer, req)
		return err
	})
	if err != nil {
		return tokenError(err)
	}
	switch {
	case secretOnly:
		fmt.Fprintln(out, tv.Token)
	case asJSON:
		enc := json.NewEncoder(out)
		enc.SetIndent("", "  ")
		return enc.Encode(tv)
	default:
		fmt.Fprintf(out, "Created token %s (%s)\n  scopes:  %s\n  expires: %s\n\n%s\n\n", tv.ID, tv.Name, strings.Join(tv.Scopes, ", "), dash(tv.ExpiresAt), tv.Token)
		fmt.Fprintln(out, "This secret will not be shown again. Store it now, for example:")
		fmt.Fprintln(out, "  provenance token create --name NAME --secret-only | gh secret set PROVISIO_TOKEN --env provisio-publish -R OWNER/REPO")
	}
	return nil
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-7f96263a","gen":"sha256:7d219ded22bdca52e8acfbc153485b2bb82cf7673bdc34a38acd2ef12afcf9dd","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func runTokenList(out io.Writer, cmd *cobra.Command, args []string) error {\n\tasJSON, _ := cmd.Flags().GetBool(\"json\")\n\tm, err := sessionManager()\n\tif err != nil {\n\t\treturn err\n\t}\n\tctx := cmd.Context()\n\tvar list []auth.TokenView\n\terr = m.Do(ctx, func(bearer string) error {\n\t\tvar err error\n\t\tlist, err = m.Client.ListTokens(ctx, bearer)\n\t\treturn err\n\t})\n\tif err != nil {\n\t\treturn tokenError(err)\n\t}\n\tif asJSON {\n\t\tif list == nil {\n\t\t\tlist = []auth.TokenView{}\n\t\t}\n\t\tenc := json.NewEncoder(out)\n\t\tenc.SetIndent(\"\", \"  \")\n\t\treturn enc.Encode(list)\n\t}\n\tif len(list) == 0 {\n\t\tfmt.Fprintln(out, \"No tokens.\")\n\t\treturn nil\n\t}\n\tw := tabwriter.NewWriter(out, 0, 4, 2, ' ', 0)\n\tfmt.Fprintln(w, \"ID\\tNAME\\tSCOPES\\tCREATED\\tLAST USED\\tEXPIRES\")\n\tfor _, t := range list {\n\t\tfmt.Fprintf(w, \"%s\\t%s\\t%s\\t%s\\t%s\\t%s\\n\", t.ID, t.Name, strings.Join(t.Scopes, \",\"), dash(t.CreatedAt), dash(t.LastUsedAt), dash(t.ExpiresAt))\n\t}\n\treturn w.Flush()\n}","file":"cmd/provenance/token_run.go","name":"runTokenList"}}
func runTokenList(out io.Writer, cmd *cobra.Command, args []string) error {
	asJSON, _ := cmd.Flags().GetBool("json")
	m, err := sessionManager()
	if err != nil {
		return err
	}
	ctx := cmd.Context()
	var list []auth.TokenView
	err = m.Do(ctx, func(bearer string) error {
		var err error
		list, err = m.Client.ListTokens(ctx, bearer)
		return err
	})
	if err != nil {
		return tokenError(err)
	}
	if asJSON {
		if list == nil {
			list = []auth.TokenView{}
		}
		enc := json.NewEncoder(out)
		enc.SetIndent("", "  ")
		return enc.Encode(list)
	}
	if len(list) == 0 {
		fmt.Fprintln(out, "No tokens.")
		return nil
	}
	w := tabwriter.NewWriter(out, 0, 4, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\tSCOPES\tCREATED\tLAST USED\tEXPIRES")
	for _, t := range list {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\n", t.ID, t.Name, strings.Join(t.Scopes, ","), dash(t.CreatedAt), dash(t.LastUsedAt), dash(t.ExpiresAt))
	}
	return w.Flush()
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-7f96263a","gen":"sha256:7d219ded22bdca52e8acfbc153485b2bb82cf7673bdc34a38acd2ef12afcf9dd","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func runTokenRevoke(out io.Writer, cmd *cobra.Command, args []string) error {\n\tid := args[0]\n\tm, err := sessionManager()\n\tif err != nil {\n\t\treturn err\n\t}\n\tctx := cmd.Context()\n\terr = m.Do(ctx, func(bearer string) error {\n\t\treturn m.Client.RevokeToken(ctx, bearer, id)\n\t})\n\tvar oe *auth.OAuthError\n\tif errors.As(err, &oe) && oe.Status == 404 {\n\t\treturn fmt.Errorf(\"no token with id %q (provenance token list shows them)\", id)\n\t}\n\tif err != nil {\n\t\treturn tokenError(err)\n\t}\n\tfmt.Fprintf(out, \"Revoked token %s. It stops working at once.\\n\", id)\n\treturn nil\n}","file":"cmd/provenance/token_run.go","name":"runTokenRevoke"}}
func runTokenRevoke(out io.Writer, cmd *cobra.Command, args []string) error {
	id := args[0]
	m, err := sessionManager()
	if err != nil {
		return err
	}
	ctx := cmd.Context()
	err = m.Do(ctx, func(bearer string) error {
		return m.Client.RevokeToken(ctx, bearer, id)
	})
	var oe *auth.OAuthError
	if errors.As(err, &oe) && oe.Status == 404 {
		return fmt.Errorf("no token with id %q (provenance token list shows them)", id)
	}
	if err != nil {
		return tokenError(err)
	}
	fmt.Fprintf(out, "Revoked token %s. It stops working at once.\n", id)
	return nil
}

// /sedum:setFunction
