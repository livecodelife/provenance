package main

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-32201524","gen":"sha256:da838940ae1e0b662456e90348132411fdd4321c7c2aa1a859c590f248a94d7c","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/logout_run.go","path":"errors"}}
import "errors"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-32201524","gen":"sha256:da838940ae1e0b662456e90348132411fdd4321c7c2aa1a859c590f248a94d7c","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/logout_run.go","path":"fmt"}}
import "fmt"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-32201524","gen":"sha256:da838940ae1e0b662456e90348132411fdd4321c7c2aa1a859c590f248a94d7c","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/logout_run.go","path":"io"}}
import "io"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-32201524","gen":"sha256:da838940ae1e0b662456e90348132411fdd4321c7c2aa1a859c590f248a94d7c","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/logout_run.go","path":"os"}}
import "os"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-32201524","gen":"sha256:da838940ae1e0b662456e90348132411fdd4321c7c2aa1a859c590f248a94d7c","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/logout_run.go","path":"time"}}
import "time"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-32201524","gen":"sha256:da838940ae1e0b662456e90348132411fdd4321c7c2aa1a859c590f248a94d7c","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/logout_run.go","path":"github.com/spf13/cobra"}}
import "github.com/spf13/cobra"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-32201524","gen":"sha256:da838940ae1e0b662456e90348132411fdd4321c7c2aa1a859c590f248a94d7c","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/logout_run.go","path":"github.com/livecodelife/provenance/auth"}}
import "github.com/livecodelife/provenance/auth"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"tier":"owned","record":"prov-2026-32201524","gen":"sha256:da838940ae1e0b662456e90348132411fdd4321c7c2aa1a859c590f248a94d7c","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func runLogout(out io.Writer, cmd *cobra.Command, args []string) error {\n\tstore, err := auth.DefaultStore()\n\tif err != nil {\n\t\treturn err\n\t}\n\tcred, _, err := store.Load()\n\tif errors.Is(err, auth.ErrNoSession) {\n\t\tfmt.Fprintln(out, \"Not signed in.\")\n\t\treturn nil\n\t}\n\tif err != nil {\n\t\treturn err\n\t}\n\trevokeFailed := \"\"\n\tif cred.RefreshToken != \"\" {\n\t\tctx, cancel := withTimeout(cmd.Context(), 15*time.Second)\n\t\tclient := &auth.Client{Server: cred.Server}\n\t\tif err := client.Logout(ctx, cred.RefreshToken); err != nil {\n\t\t\trevokeFailed = err.Error()\n\t\t}\n\t\tcancel()\n\t}\n\tif err := store.Delete(); err != nil {\n\t\treturn err\n\t}\n\tfmt.Fprintln(out, \"✓ Signed out; the local credential is deleted.\")\n\tif revokeFailed != \"\" {\n\t\tfmt.Fprintf(out, \"  The server could not be reached to revoke the session (%s); it expires on its own.\\n\", revokeFailed)\n\t}\n\tif cred.PersonalToken != \"\" {\n\t\tfmt.Fprintln(out, \"  A personal token stays valid until you delete it in the Provisio UI.\")\n\t}\n\tif os.Getenv(auth.EnvToken) != \"\" {\n\t\tfmt.Fprintf(out, \"  %s is still set in this environment.\\n\", auth.EnvToken)\n\t}\n\treturn nil\n}","file":"cmd/provenance/logout_run.go","name":"runLogout"}}
func runLogout(out io.Writer, cmd *cobra.Command, args []string) error {
	store, err := auth.DefaultStore()
	if err != nil {
		return err
	}
	cred, _, err := store.Load()
	if errors.Is(err, auth.ErrNoSession) {
		fmt.Fprintln(out, "Not signed in.")
		return nil
	}
	if err != nil {
		return err
	}
	revokeFailed := ""
	if cred.RefreshToken != "" {
		ctx, cancel := withTimeout(cmd.Context(), 15*time.Second)
		client := &auth.Client{Server: cred.Server}
		if err := client.Logout(ctx, cred.RefreshToken); err != nil {
			revokeFailed = err.Error()
		}
		cancel()
	}
	if err := store.Delete(); err != nil {
		return err
	}
	fmt.Fprintln(out, "✓ Signed out; the local credential is deleted.")
	if revokeFailed != "" {
		fmt.Fprintf(out, "  The server could not be reached to revoke the session (%s); it expires on its own.\n", revokeFailed)
	}
	if cred.PersonalToken != "" {
		fmt.Fprintln(out, "  A personal token stays valid until you delete it in the Provisio UI.")
	}
	if os.Getenv(auth.EnvToken) != "" {
		fmt.Fprintf(out, "  %s is still set in this environment.\n", auth.EnvToken)
	}
	return nil
}

// /sedum:setFunction
