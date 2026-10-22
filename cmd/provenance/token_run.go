package main

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-21fa7b7f","gen":"sha256:ed5483bb6e3c61f0fc809cd66d1bec320b1886d3b1ec685acc5a8b4815175b9e","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/token_run.go","path":"encoding/json"}}
import "encoding/json"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-21fa7b7f","gen":"sha256:ed5483bb6e3c61f0fc809cd66d1bec320b1886d3b1ec685acc5a8b4815175b9e","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/token_run.go","path":"io"}}
import "io"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-8ef3d267","gen":"sha256:d9de083abd0f66b5668ed69bf5ac4f03c48d6216edab3c2fd0242de8eb528fcd","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/token_run.go","path":"errors"}}
import "errors"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-8ef3d267","gen":"sha256:d9de083abd0f66b5668ed69bf5ac4f03c48d6216edab3c2fd0242de8eb528fcd","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/token_run.go","path":"fmt"}}
import "fmt"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-8ef3d267","gen":"sha256:d9de083abd0f66b5668ed69bf5ac4f03c48d6216edab3c2fd0242de8eb528fcd","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/token_run.go","path":"github.com/livecodelife/provenance/auth"}}
import "github.com/livecodelife/provenance/auth"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-d90240d6","gen":"sha256:c9129118fd6768a40a7558e681d443ccf2da0b756ecb96cc86f84c70c66b8595","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/token_run.go","path":"path"}}
import "path"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-d90240d6","gen":"sha256:c9129118fd6768a40a7558e681d443ccf2da0b756ecb96cc86f84c70c66b8595","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/token_run.go","path":"path/filepath"}}
import "path/filepath"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-d90240d6","gen":"sha256:c9129118fd6768a40a7558e681d443ccf2da0b756ecb96cc86f84c70c66b8595","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/token_run.go","path":"strings"}}
import "strings"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-d90240d6","gen":"sha256:c9129118fd6768a40a7558e681d443ccf2da0b756ecb96cc86f84c70c66b8595","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/token_run.go","path":"github.com/spf13/cobra"}}
import "github.com/spf13/cobra"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-d906587d","gen":"sha256:9ffe591a871dc124f129444b234fe8ce32b1b42f924659d22ded66953b004b77","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/token_run.go","path":"text/tabwriter"}}
import "text/tabwriter"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"tier":"owned","record":"prov-2026-21fa7b7f","gen":"sha256:ed5483bb6e3c61f0fc809cd66d1bec320b1886d3b1ec685acc5a8b4815175b9e","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func runTokenCreate(out io.Writer, cmd *cobra.Command, args []string) error {\n\treq, err := tokenRequest(cmd)\n\tif err != nil {\n\t\treturn err\n\t}\n\tasJSON, _ := cmd.Flags().GetBool(\"json\")\n\tsecretOnly, _ := cmd.Flags().GetBool(\"secret-only\")\n\tm, err := sessionManager()\n\tif err != nil {\n\t\treturn err\n\t}\n\tctx := cmd.Context()\n\tvar view *auth.TokenView\n\terr = m.Do(ctx, func(bearer string) error {\n\t\tv, err := m.Client.CreateToken(ctx, bearer, req)\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n\t\tview = v\n\t\treturn nil\n\t})\n\tif err != nil {\n\t\treturn tokenError(err)\n\t}\n\tif secretOnly {\n\t\tfmt.Fprintln(out, view.Token)\n\t\treturn nil\n\t}\n\tif asJSON {\n\t\tenc := json.NewEncoder(out)\n\t\tenc.SetIndent(\"\", \"  \")\n\t\treturn enc.Encode(view)\n\t}\n\tfmt.Fprintf(out, \"Created token %s (%s)\\n\", view.ID, view.Name)\n\tfmt.Fprintf(out, \"  scopes: %s\\n\", strings.Join(view.Scopes, \", \"))\n\tfmt.Fprintf(out, \"  expires: %s\\n\", dash(view.ExpiresAt))\n\tfmt.Fprintf(out, \"\\n%s\\n\\n\", view.Token)\n\tfmt.Fprintln(out, \"This secret will not be shown again. Store it now, for example:\")\n\tfmt.Fprintf(out, \"  provenance token create --name %s --secret-only | gh secret set PROVISIO_TOKEN --env provisio-publish -R OWNER/REPO\\n\", view.Name)\n\treturn nil\n}","file":"cmd/provenance/token_run.go","name":"runTokenCreate"}}
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
	var view *auth.TokenView
	err = m.Do(ctx, func(bearer string) error {
		v, err := m.Client.CreateToken(ctx, bearer, req)
		if err != nil {
			return err
		}
		view = v
		return nil
	})
	if err != nil {
		return tokenError(err)
	}
	if secretOnly {
		fmt.Fprintln(out, view.Token)
		return nil
	}
	if asJSON {
		enc := json.NewEncoder(out)
		enc.SetIndent("", "  ")
		return enc.Encode(view)
	}
	fmt.Fprintf(out, "Created token %s (%s)\n", view.ID, view.Name)
	fmt.Fprintf(out, "  scopes: %s\n", strings.Join(view.Scopes, ", "))
	fmt.Fprintf(out, "  expires: %s\n", dash(view.ExpiresAt))
	fmt.Fprintf(out, "\n%s\n\n", view.Token)
	fmt.Fprintln(out, "This secret will not be shown again. Store it now, for example:")
	fmt.Fprintf(out, "  provenance token create --name %s --secret-only | gh secret set PROVISIO_TOKEN --env provisio-publish -R OWNER/REPO\n", view.Name)
	return nil
}

// /sedum:setFunction
// sedum:setVar {"tier":"owned","record":"prov-2026-4ccf3cb6","gen":"sha256:40c740fd19b52b3dc642b6addc230dd38bbe082a6d605e53fd1a8f610491f7ba","pkg":"sha256:c7b9d9ac1655e8211b8b2ca8c49cf34b8d510d8d823d4f5eb60ec7e1018aa2eb","kwargs":{"code":"var tokenScopes = map[string]bool{\n\t\"records:read\":  true,\n\t\"records:write\": true,\n}","file":"cmd/provenance/token_run.go","name":"tokenScopes"}}
var tokenScopes = map[string]bool{
	"records:read":  true,
	"records:write": true,
}

// /sedum:setVar
// sedum:setFunction {"tier":"owned","record":"prov-2026-77a96139","gen":"sha256:7a6aa6893246f4975a87d9b6b57a326e9b19e20bb42fbdf8c085a20b5627e2d4","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func dash(s string) string {\n\tif s == \"\" {\n\t\treturn \"-\"\n\t}\n\treturn s\n}","file":"cmd/provenance/token_run.go","name":"dash"}}
func dash(s string) string {
	if s == "" {
		return "-"
	}
	return s
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-8ef3d267","gen":"sha256:d9de083abd0f66b5668ed69bf5ac4f03c48d6216edab3c2fd0242de8eb528fcd","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func tokenError(err error) error {\n\tvar oauthErr *auth.OAuthError\n\tif errors.As(err, &oauthErr) {\n\t\tif oauthErr.Status == 403 {\n\t\t\treturn &exitError{code: 3, err: errors.New(msgTokenSession)}\n\t\t}\n\t\tif oauthErr.Description != \"\" {\n\t\t\treturn fmt.Errorf(\"%s (HTTP %d): %s\", oauthErr.Code, oauthErr.Status, oauthErr.Description)\n\t\t}\n\t}\n\treturn sessionError(err)\n}","file":"cmd/provenance/token_run.go","name":"tokenError"}}
func tokenError(err error) error {
	var oauthErr *auth.OAuthError
	if errors.As(err, &oauthErr) {
		if oauthErr.Status == 403 {
			return &exitError{code: 3, err: errors.New(msgTokenSession)}
		}
		if oauthErr.Description != "" {
			return fmt.Errorf("%s (HTTP %d): %s", oauthErr.Code, oauthErr.Status, oauthErr.Description)
		}
	}
	return sessionError(err)
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-a0f84eec","gen":"sha256:f9864342d8a1f22f49671f96a826f82f9fb3114fe5bb4e2076306eed79fbec84","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func tokenRequest(cmd *cobra.Command) (auth.TokenRequest, error) {\n\tname, _ := cmd.Flags().GetString(\"name\")\n\tscopes, _ := cmd.Flags().GetStringArray(\"scope\")\n\tdays, _ := cmd.Flags().GetInt(\"expires-in-days\")\n\n\tname = strings.TrimSpace(name)\n\tif name == \"\" && !cmd.Flags().Changed(\"name\") {\n\t\tdef, err := defaultTokenName(cmd)\n\t\tif err != nil {\n\t\t\treturn auth.TokenRequest{}, err\n\t\t}\n\t\tname = def\n\t}\n\tif len(name) < 1 || len(name) > 60 {\n\t\treturn auth.TokenRequest{}, errors.New(\"--name must be 1 to 60 characters\")\n\t}\n\n\tif len(scopes) == 0 {\n\t\tscopes = []string{\"records:write\"}\n\t}\n\tfor _, s := range scopes {\n\t\tif !tokenScopes[s] {\n\t\t\treturn auth.TokenRequest{}, fmt.Errorf(\"--scope %q: must be records:read or records:write\", s)\n\t\t}\n\t}\n\n\tif !cmd.Flags().Changed(\"expires-in-days\") {\n\t\tdays = 90\n\t}\n\tif days < 1 || days > 365 {\n\t\treturn auth.TokenRequest{}, errors.New(\"--expires-in-days must be 1 to 365\")\n\t}\n\n\treturn auth.TokenRequest{Name: name, Scopes: scopes, ExpiresInDays: days}, nil\n}","file":"cmd/provenance/token_run.go","name":"tokenRequest"}}
func tokenRequest(cmd *cobra.Command) (auth.TokenRequest, error) {
	name, _ := cmd.Flags().GetString("name")
	scopes, _ := cmd.Flags().GetStringArray("scope")
	days, _ := cmd.Flags().GetInt("expires-in-days")

	name = strings.TrimSpace(name)
	if name == "" && !cmd.Flags().Changed("name") {
		def, err := defaultTokenName(cmd)
		if err != nil {
			return auth.TokenRequest{}, err
		}
		name = def
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
// sedum:setFunction {"tier":"owned","record":"prov-2026-cb2eb740","gen":"sha256:9f6c71d5da246bef57684ef8dc6865bbc2f161d1232ebea5f527221e8cb37320","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func runTokenRevoke(out io.Writer, cmd *cobra.Command, args []string) error {\n\tid := args[0]\n\tm, err := sessionManager()\n\tif err != nil {\n\t\treturn err\n\t}\n\tctx := cmd.Context()\n\terr = m.Do(ctx, func(bearer string) error {\n\t\treturn m.Client.RevokeToken(ctx, bearer, id)\n\t})\n\tvar oauthErr *auth.OAuthError\n\tif errors.As(err, &oauthErr) && oauthErr.Status == 404 {\n\t\treturn fmt.Errorf(\"no token with id %q (provenance token list shows them)\", id)\n\t}\n\tif err != nil {\n\t\treturn tokenError(err)\n\t}\n\tfmt.Fprintf(out, \"Revoked token %s. It stops working at once.\\n\", id)\n\treturn nil\n}","file":"cmd/provenance/token_run.go","name":"runTokenRevoke"}}
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
	var oauthErr *auth.OAuthError
	if errors.As(err, &oauthErr) && oauthErr.Status == 404 {
		return fmt.Errorf("no token with id %q (provenance token list shows them)", id)
	}
	if err != nil {
		return tokenError(err)
	}
	fmt.Fprintf(out, "Revoked token %s. It stops working at once.\n", id)
	return nil
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-d90240d6","gen":"sha256:c9129118fd6768a40a7558e681d443ccf2da0b756ecb96cc86f84c70c66b8595","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func defaultTokenName(cmd *cobra.Command) (string, error) {\n\tr, err := openRepo(cmd)\n\tif err != nil {\n\t\treturn \"\", errors.New(\"--name is required outside a repository\")\n\t}\n\ttop, err := r.git(\"rev-parse\", \"--show-toplevel\")\n\tif err != nil {\n\t\treturn \"\", errors.New(\"--name is required outside a repository\")\n\t}\n\tname := filepath.Base(strings.TrimSpace(top))\n\tif out, err := r.git(\"remote\", \"get-url\", \"origin\"); err == nil {\n\t\tif u := repoURL(out); u != \"\" {\n\t\t\tname = path.Base(u)\n\t\t}\n\t}\n\treturn \"ci-\" + name, nil\n}","file":"cmd/provenance/token_run.go","name":"defaultTokenName"}}
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
	if out, err := r.git("remote", "get-url", "origin"); err == nil {
		if u := repoURL(out); u != "" {
			name = path.Base(u)
		}
	}
	return "ci-" + name, nil
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-d906587d","gen":"sha256:9ffe591a871dc124f129444b234fe8ce32b1b42f924659d22ded66953b004b77","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func runTokenList(out io.Writer, cmd *cobra.Command, args []string) error {\n\tasJSON, _ := cmd.Flags().GetBool(\"json\")\n\tmgr, err := sessionManager()\n\tif err != nil {\n\t\treturn err\n\t}\n\tctx := cmd.Context()\n\tvar list []auth.TokenView\n\terr = mgr.Do(ctx, func(bearer string) error {\n\t\tl, err := mgr.Client.ListTokens(ctx, bearer)\n\t\tlist = l\n\t\treturn err\n\t})\n\tif err != nil {\n\t\treturn tokenError(err)\n\t}\n\tif asJSON {\n\t\tif list == nil {\n\t\t\tlist = []auth.TokenView{}\n\t\t}\n\t\tenc := json.NewEncoder(out)\n\t\tenc.SetIndent(\"\", \"  \")\n\t\treturn enc.Encode(list)\n\t}\n\tif len(list) == 0 {\n\t\tfmt.Fprintln(out, \"No tokens.\")\n\t\treturn nil\n\t}\n\ttw := tabwriter.NewWriter(out, 0, 4, 2, ' ', 0)\n\tfmt.Fprintln(tw, \"ID\\tNAME\\tSCOPES\\tCREATED\\tLAST USED\\tEXPIRES\")\n\tfor _, t := range list {\n\t\tfmt.Fprintf(tw, \"%s\\t%s\\t%s\\t%s\\t%s\\t%s\\n\", t.ID, t.Name, strings.Join(t.Scopes, \",\"), dash(t.CreatedAt), dash(t.LastUsedAt), dash(t.ExpiresAt))\n\t}\n\treturn tw.Flush()\n}\n","file":"cmd/provenance/token_run.go","name":"runTokenList"}}
func runTokenList(out io.Writer, cmd *cobra.Command, args []string) error {
	asJSON, _ := cmd.Flags().GetBool("json")
	mgr, err := sessionManager()
	if err != nil {
		return err
	}
	ctx := cmd.Context()
	var list []auth.TokenView
	err = mgr.Do(ctx, func(bearer string) error {
		l, err := mgr.Client.ListTokens(ctx, bearer)
		list = l
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
	tw := tabwriter.NewWriter(out, 0, 4, 2, ' ', 0)
	fmt.Fprintln(tw, "ID\tNAME\tSCOPES\tCREATED\tLAST USED\tEXPIRES")
	for _, t := range list {
		fmt.Fprintf(tw, "%s\t%s\t%s\t%s\t%s\t%s\n", t.ID, t.Name, strings.Join(t.Scopes, ","), dash(t.CreatedAt), dash(t.LastUsedAt), dash(t.ExpiresAt))
	}
	return tw.Flush()
}

// /sedum:setFunction
// sedum:setConst {"tier":"owned","record":"prov-2026-e7a14eb3","gen":"sha256:5bdd41ea5b536a9d96fa44bec3b33217d78d436641a1fb52c0c82085722075e1","pkg":"sha256:51be1c18998d63cc558c253e211a6bd188d96cb251dad68c6ea88e03c2d5cc4c","kwargs":{"code":"const msgTokenSession = \"Managing tokens needs a signed-in session, not a personal token. Run provenance login (and unset PROVISIO_TOKEN).\"","file":"cmd/provenance/token_run.go","name":"msgTokenSession"}}
const msgTokenSession = "Managing tokens needs a signed-in session, not a personal token. Run provenance login (and unset PROVISIO_TOKEN)."

// /sedum:setConst
