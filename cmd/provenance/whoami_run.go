package main

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-215c273d","gen":"sha256:7d6dead1563884c13e30140873f1cd8a63b81cc841aedfb88922843a11509c33","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/whoami_run.go","path":"encoding/json"}}
import "encoding/json"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-215c273d","gen":"sha256:7d6dead1563884c13e30140873f1cd8a63b81cc841aedfb88922843a11509c33","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/whoami_run.go","path":"errors"}}
import "errors"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-215c273d","gen":"sha256:7d6dead1563884c13e30140873f1cd8a63b81cc841aedfb88922843a11509c33","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/whoami_run.go","path":"fmt"}}
import "fmt"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-215c273d","gen":"sha256:7d6dead1563884c13e30140873f1cd8a63b81cc841aedfb88922843a11509c33","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/whoami_run.go","path":"io"}}
import "io"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-215c273d","gen":"sha256:7d6dead1563884c13e30140873f1cd8a63b81cc841aedfb88922843a11509c33","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/whoami_run.go","path":"github.com/spf13/cobra"}}
import "github.com/spf13/cobra"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-215c273d","gen":"sha256:7d6dead1563884c13e30140873f1cd8a63b81cc841aedfb88922843a11509c33","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/whoami_run.go","path":"github.com/livecodelife/provenance/auth"}}
import "github.com/livecodelife/provenance/auth"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"tier":"owned","record":"prov-2026-215c273d","gen":"sha256:7d6dead1563884c13e30140873f1cd8a63b81cc841aedfb88922843a11509c33","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func runWhoami(out io.Writer, cmd *cobra.Command, args []string) error {\n\tasJSON, _ := cmd.Flags().GetBool(\"json\")\n\tm, err := sessionManager()\n\tif err != nil {\n\t\treturn err\n\t}\n\tctx := cmd.Context()\n\ttok, err := m.Token(ctx)\n\tif errors.Is(err, auth.ErrNoSession) {\n\t\tfmt.Fprintln(out, \"Not signed in.\")\n\t\treturn errors.New(\"not signed in\")\n\t}\n\tif err != nil {\n\t\treturn sessionError(err)\n\t}\n\tvar info *auth.SessionInfo\n\terr = m.Do(ctx, func(bearer string) error {\n\t\tvar err error\n\t\tinfo, err = m.Client.Session(ctx, bearer)\n\t\treturn err\n\t})\n\tif err != nil {\n\t\treturn sessionError(err)\n\t}\n\texpires := \"\"\n\tif t := info.Expiry(); !t.IsZero() {\n\t\texpires = t.UTC().Format(\"2006-01-02T15:04:05Z\")\n\t}\n\tif asJSON {\n\t\tview := map[string]any{\"handle\": info.Handle, \"did\": info.DID, \"source\": string(tok.Source), \"expires_at\": expires}\n\t\tif info.Token != nil {\n\t\t\tview[\"token\"] = map[string]any{\"id\": info.Token.ID, \"name\": info.Token.Name, \"scopes\": info.Token.Scopes}\n\t\t}\n\t\tenc := json.NewEncoder(out)\n\t\tenc.SetIndent(\"\", \"  \")\n\t\treturn enc.Encode(view)\n\t}\n\tfmt.Fprintf(out, \"Handle:  @%s\\nDID:     %s\\nSource:  %s\\n\", info.Handle, info.DID, tok.Source)\n\tif expires != \"\" {\n\t\tfmt.Fprintf(out, \"Expires: %s\\n\", expires)\n\t}\n\treturn nil\n}","file":"cmd/provenance/whoami_run.go","name":"runWhoami"}}
func runWhoami(out io.Writer, cmd *cobra.Command, args []string) error {
	asJSON, _ := cmd.Flags().GetBool("json")
	m, err := sessionManager()
	if err != nil {
		return err
	}
	ctx := cmd.Context()
	tok, err := m.Token(ctx)
	if errors.Is(err, auth.ErrNoSession) {
		fmt.Fprintln(out, "Not signed in.")
		return errors.New("not signed in")
	}
	if err != nil {
		return sessionError(err)
	}
	var info *auth.SessionInfo
	err = m.Do(ctx, func(bearer string) error {
		var err error
		info, err = m.Client.Session(ctx, bearer)
		return err
	})
	if err != nil {
		return sessionError(err)
	}
	expires := ""
	if t := info.Expiry(); !t.IsZero() {
		expires = t.UTC().Format("2006-01-02T15:04:05Z")
	}
	if asJSON {
		view := map[string]any{"handle": info.Handle, "did": info.DID, "source": string(tok.Source), "expires_at": expires}
		if info.Token != nil {
			view["token"] = map[string]any{"id": info.Token.ID, "name": info.Token.Name, "scopes": info.Token.Scopes}
		}
		enc := json.NewEncoder(out)
		enc.SetIndent("", "  ")
		return enc.Encode(view)
	}
	fmt.Fprintf(out, "Handle:  @%s\nDID:     %s\nSource:  %s\n", info.Handle, info.DID, tok.Source)
	if expires != "" {
		fmt.Fprintf(out, "Expires: %s\n", expires)
	}
	return nil
}

// /sedum:setFunction
