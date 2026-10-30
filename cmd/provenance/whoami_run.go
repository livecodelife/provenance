package main

// sedum:anchor:imports
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-e81ba1b2","gen":"sha256:014d047a0dc61aaff6ee2513ee870940a9b66340a759448be05682ede7193127","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/whoami_run.go","path":"errors"},"kwargs":"sha256:ed66034bf4d93fc5b7b8a0b8d67790b7654a24cdeefe1f63b5154df42d18ebf0","body":"sha256:2b718237ba0a56467c8fd494cc131e942ef410e6228c03b644d70476c71862dc"}
import "errors"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-e81ba1b2","gen":"sha256:014d047a0dc61aaff6ee2513ee870940a9b66340a759448be05682ede7193127","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/whoami_run.go","path":"fmt"},"kwargs":"sha256:16d2098f5c7274c3f1066a7d0e1455c5fa03f21c0a77ac0e01c7203e3aec531c","body":"sha256:bcbac51883cbf8064e9196b371cb880cee429866652447ad557533d817857609"}
import "fmt"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-e81ba1b2","gen":"sha256:014d047a0dc61aaff6ee2513ee870940a9b66340a759448be05682ede7193127","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/whoami_run.go","path":"io"},"kwargs":"sha256:592fe94df7a22be6285e26f0a3ca4765fe2476b09f920d542107a6728535bf3c","body":"sha256:ef8e8abf4ef5ee40a19dc85b49bf09c5df1fe32abfe834e0ab4c2175baf80ddf"}
import "io"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-e81ba1b2","gen":"sha256:014d047a0dc61aaff6ee2513ee870940a9b66340a759448be05682ede7193127","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/whoami_run.go","path":"encoding/json"},"kwargs":"sha256:6e9e36376d4434e20490c78b05b150b7c44eafb9d5675b1661cacd3f089ab920","body":"sha256:872d7c3e9dee608e79a7f8ef57930d24b5e30f2a57d891a52549179ace45946a"}
import "encoding/json"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-e81ba1b2","gen":"sha256:014d047a0dc61aaff6ee2513ee870940a9b66340a759448be05682ede7193127","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/whoami_run.go","path":"github.com/spf13/cobra"},"kwargs":"sha256:9ae6d3492418731b64910a9cbcdb4838059cd8d0191e033b6175131bf4c1205a","body":"sha256:f361c02d0f8e1999cf3fb715393e16310eeec9debd5aaa13936af7617e2b46aa"}
import "github.com/spf13/cobra"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-e81ba1b2","gen":"sha256:014d047a0dc61aaff6ee2513ee870940a9b66340a759448be05682ede7193127","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/whoami_run.go","path":"github.com/livecodelife/provenance/auth"},"kwargs":"sha256:41da7bc16c4125104d5cec1d199c46695ae3f13f0b9326634b2ae2fa90956644","body":"sha256:57cbe62cd356e3c713dcc35195098163d64975ccdb05ef9a2317cf089e95c547"}
import "github.com/livecodelife/provenance/auth"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-e81ba1b2","gen":"sha256:014d047a0dc61aaff6ee2513ee870940a9b66340a759448be05682ede7193127","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/whoami_run.go","name":"runWhoami"},"kwargs":"sha256:e01a02c5c52c15b747940ee2825196c5905c9a4ec99f91b818e44058f449ed42","body":"sha256:c124d94bbf39bde2e03bc9e87e03262c1769059ffe21e54b02a87bedce53a12c"}
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
