package main

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-e81ba1b2","gen":"sha256:014d047a0dc61aaff6ee2513ee870940a9b66340a759448be05682ede7193127","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/whoami_run.go","path":"errors"}}
import "errors"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-e81ba1b2","gen":"sha256:014d047a0dc61aaff6ee2513ee870940a9b66340a759448be05682ede7193127","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/whoami_run.go","path":"fmt"}}
import "fmt"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-e81ba1b2","gen":"sha256:014d047a0dc61aaff6ee2513ee870940a9b66340a759448be05682ede7193127","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/whoami_run.go","path":"io"}}
import "io"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-e81ba1b2","gen":"sha256:014d047a0dc61aaff6ee2513ee870940a9b66340a759448be05682ede7193127","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/whoami_run.go","path":"encoding/json"}}
import "encoding/json"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-e81ba1b2","gen":"sha256:014d047a0dc61aaff6ee2513ee870940a9b66340a759448be05682ede7193127","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/whoami_run.go","path":"github.com/spf13/cobra"}}
import "github.com/spf13/cobra"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-e81ba1b2","gen":"sha256:014d047a0dc61aaff6ee2513ee870940a9b66340a759448be05682ede7193127","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/whoami_run.go","path":"github.com/livecodelife/provenance/auth"}}
import "github.com/livecodelife/provenance/auth"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"tier":"owned","record":"prov-2026-e81ba1b2","gen":"sha256:014d047a0dc61aaff6ee2513ee870940a9b66340a759448be05682ede7193127","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func runWhoami(out io.Writer, cmd *cobra.Command, args []string) error {\n\tasJSON, _ := cmd.Flags().GetBool(\"json\")\n\n\tmgr, err := sessionManager()\n\tif err != nil {\n\t\treturn err\n\t}\n\n\tctx := cmd.Context()\n\ttok, err := mgr.Token(ctx)\n\tif err != nil {\n\t\tif errors.Is(err, auth.ErrNoSession) {\n\t\t\tfmt.Fprintln(out, \"Not signed in.\")\n\t\t\treturn errors.New(\"not signed in\")\n\t\t}\n\t\treturn sessionError(err)\n\t}\n\n\tvar info *auth.SessionInfo\n\tif err := mgr.Do(ctx, func(bearer string) error {\n\t\tvar err error\n\t\tinfo, err = mgr.Client.Session(ctx, bearer)\n\t\treturn err\n\t}); err != nil {\n\t\treturn sessionError(err)\n\t}\n\n\texpiresAt := \"\"\n\tif exp := info.Expiry(); !exp.IsZero() {\n\t\texpiresAt = exp.UTC().Format(\"2006-01-02T15:04:05Z\")\n\t}\n\n\tif asJSON {\n\t\tm := map[string]any{\n\t\t\t\"handle\":     info.Handle,\n\t\t\t\"did\":        info.DID,\n\t\t\t\"source\":     string(tok.Source),\n\t\t\t\"expires_at\": expiresAt,\n\t\t}\n\t\tif info.Token != nil {\n\t\t\tm[\"token\"] = map[string]any{\n\t\t\t\t\"id\":     info.Token.ID,\n\t\t\t\t\"name\":   info.Token.Name,\n\t\t\t\t\"scopes\": info.Token.Scopes,\n\t\t\t}\n\t\t}\n\t\tenc := json.NewEncoder(out)\n\t\tenc.SetIndent(\"\", \"  \")\n\t\treturn enc.Encode(m)\n\t}\n\n\tfmt.Fprintf(out, \"Handle:  @%s\\n\", info.Handle)\n\tfmt.Fprintf(out, \"DID:     %s\\n\", info.DID)\n\tfmt.Fprintf(out, \"Source:  %s\\n\", string(tok.Source))\n\tif expiresAt != \"\" {\n\t\tfmt.Fprintf(out, \"Expires: %s\\n\", expiresAt)\n\t}\n\treturn nil\n}","file":"cmd/provenance/whoami_run.go","name":"runWhoami"}}
func runWhoami(out io.Writer, cmd *cobra.Command, args []string) error {
	asJSON, _ := cmd.Flags().GetBool("json")

	mgr, err := sessionManager()
	if err != nil {
		return err
	}

	ctx := cmd.Context()
	tok, err := mgr.Token(ctx)
	if err != nil {
		if errors.Is(err, auth.ErrNoSession) {
			fmt.Fprintln(out, "Not signed in.")
			return errors.New("not signed in")
		}
		return sessionError(err)
	}

	var info *auth.SessionInfo
	if err := mgr.Do(ctx, func(bearer string) error {
		var err error
		info, err = mgr.Client.Session(ctx, bearer)
		return err
	}); err != nil {
		return sessionError(err)
	}

	expiresAt := ""
	if exp := info.Expiry(); !exp.IsZero() {
		expiresAt = exp.UTC().Format("2006-01-02T15:04:05Z")
	}

	if asJSON {
		m := map[string]any{
			"handle":     info.Handle,
			"did":        info.DID,
			"source":     string(tok.Source),
			"expires_at": expiresAt,
		}
		if info.Token != nil {
			m["token"] = map[string]any{
				"id":     info.Token.ID,
				"name":   info.Token.Name,
				"scopes": info.Token.Scopes,
			}
		}
		enc := json.NewEncoder(out)
		enc.SetIndent("", "  ")
		return enc.Encode(m)
	}

	fmt.Fprintf(out, "Handle:  @%s\n", info.Handle)
	fmt.Fprintf(out, "DID:     %s\n", info.DID)
	fmt.Fprintf(out, "Source:  %s\n", string(tok.Source))
	if expiresAt != "" {
		fmt.Fprintf(out, "Expires: %s\n", expiresAt)
	}
	return nil
}

// /sedum:setFunction
