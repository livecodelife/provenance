package main

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-635b25ac","gen":"sha256:3dd07349442894e371a607dad8f67454397b701f2bd4e8d8d0edfb17be96c5a0","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/login_run.go","path":"fmt"}}
import "fmt"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-635b25ac","gen":"sha256:3dd07349442894e371a607dad8f67454397b701f2bd4e8d8d0edfb17be96c5a0","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/login_run.go","path":"io"}}
import "io"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-635b25ac","gen":"sha256:3dd07349442894e371a607dad8f67454397b701f2bd4e8d8d0edfb17be96c5a0","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/login_run.go","path":"os"}}
import "os"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-635b25ac","gen":"sha256:3dd07349442894e371a607dad8f67454397b701f2bd4e8d8d0edfb17be96c5a0","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/login_run.go","path":"os/exec"}}
import "os/exec"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-635b25ac","gen":"sha256:3dd07349442894e371a607dad8f67454397b701f2bd4e8d8d0edfb17be96c5a0","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/login_run.go","path":"runtime"}}
import "runtime"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-635b25ac","gen":"sha256:3dd07349442894e371a607dad8f67454397b701f2bd4e8d8d0edfb17be96c5a0","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/login_run.go","path":"strings"}}
import "strings"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-635b25ac","gen":"sha256:3dd07349442894e371a607dad8f67454397b701f2bd4e8d8d0edfb17be96c5a0","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/login_run.go","path":"time"}}
import "time"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-635b25ac","gen":"sha256:3dd07349442894e371a607dad8f67454397b701f2bd4e8d8d0edfb17be96c5a0","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/login_run.go","path":"github.com/spf13/cobra"}}
import "github.com/spf13/cobra"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-635b25ac","gen":"sha256:3dd07349442894e371a607dad8f67454397b701f2bd4e8d8d0edfb17be96c5a0","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/login_run.go","path":"github.com/livecodelife/provenance/auth"}}
import "github.com/livecodelife/provenance/auth"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-635b25ac","gen":"sha256:3dd07349442894e371a607dad8f67454397b701f2bd4e8d8d0edfb17be96c5a0","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/login_run.go","path":"errors"}}
import "errors"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"tier":"owned","record":"prov-2026-635b25ac","gen":"sha256:3dd07349442894e371a607dad8f67454397b701f2bd4e8d8d0edfb17be96c5a0","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func runLogin(out io.Writer, cmd *cobra.Command, args []string) error {\n\tprovider, _ := cmd.Flags().GetString(\"provider\")\n\tif !cmd.Flags().Changed(\"provider\") {\n\t\tprovider = \"provisio\"\n\t}\n\tif provider != \"provisio\" && provider != \"bluesky\" {\n\t\treturn errors.New(\"--provider must be provisio or bluesky\")\n\t}\n\tnoBrowser, _ := cmd.Flags().GetBool(\"no-browser\")\n\twithToken, _ := cmd.Flags().GetBool(\"with-token\")\n\tserver, _ := cmd.Flags().GetString(\"server\")\n\tif server == \"\" {\n\t\tserver = serverURL()\n\t}\n\tif err := auth.CheckServer(server); err != nil {\n\t\treturn err\n\t}\n\tstore, err := auth.DefaultStore()\n\tif err != nil {\n\t\treturn err\n\t}\n\tclient := &auth.Client{Server: server}\n\tctx := cmd.Context()\n\n\tvar cred *auth.Credential\n\tif withToken {\n\t\traw, err := io.ReadAll(io.LimitReader(cmd.InOrStdin(), 4096))\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n\t\ttoken := strings.TrimSpace(string(raw))\n\t\tif !strings.HasPrefix(token, \"pvt_\") {\n\t\t\treturn fmt.Errorf(\"expected a personal token (pvt_...) on stdin\")\n\t\t}\n\t\tinfo, err := client.Session(ctx, token)\n\t\tif err != nil {\n\t\t\treturn fmt.Errorf(\"the server did not accept that token: %w\", err)\n\t\t}\n\t\tcred = &auth.Credential{Server: server, PersonalToken: token, DID: info.DID, Handle: info.Handle}\n\t} else {\n\t\thost, _ := os.Hostname()\n\t\tclient.Hostname = host\n\t\tclient.Provider = provider\n\t\tdc, err := client.RequestDeviceCode(ctx)\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n\t\tfmt.Fprintf(out, \"\\nTo sign in, open %s\\nand check that the code is %s\\n\\n\", dc.VerificationURI, dc.UserCode)\n\t\tif !noBrowser && openBrowser(dc.VerificationURIComplete) {\n\t\t\tfmt.Fprintln(out, \"Opened your browser. Waiting for approval...\")\n\t\t} else {\n\t\t\tfmt.Fprintln(out, \"Waiting for approval...\")\n\t\t}\n\t\tpollCtx := ctx\n\t\tif dc.ExpiresIn > 0 {\n\t\t\tvar cancel func()\n\t\t\tpollCtx, cancel = withTimeout(ctx, time.Duration(dc.ExpiresIn)*time.Second)\n\t\t\tdefer cancel()\n\t\t}\n\t\ttr, err := client.PollToken(pollCtx, dc, nil)\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n\t\tnow := time.Now()\n\t\tcred = &auth.Credential{\n\t\t\tServer:           server,\n\t\t\tAccessToken:      tr.AccessToken,\n\t\t\tAccessExpiresAt:  now.Add(time.Duration(tr.ExpiresIn) * time.Second),\n\t\t\tRefreshToken:     tr.RefreshToken,\n\t\t\tRefreshExpiresAt: now.Add(time.Duration(tr.RefreshExpiresIn) * time.Second),\n\t\t\tDID:              tr.DID,\n\t\t\tHandle:           tr.Handle,\n\t\t}\n\t}\n\tif err := store.Save(cred); err != nil {\n\t\treturn err\n\t}\n\tfmt.Fprintf(out, \"\\n✓ Signed in as @%s (%s)\\n\", cred.Handle, cred.DID)\n\tif os.Getenv(auth.EnvToken) != \"\" {\n\t\tfmt.Fprintf(out, \"  Note: %s is set and takes precedence over this sign-in.\\n\", auth.EnvToken)\n\t}\n\tfmt.Fprintln(out)\n\treturn nil\n}","file":"cmd/provenance/login_run.go","name":"runLogin"}}
func runLogin(out io.Writer, cmd *cobra.Command, args []string) error {
	provider, _ := cmd.Flags().GetString("provider")
	if !cmd.Flags().Changed("provider") {
		provider = "provisio"
	}
	if provider != "provisio" && provider != "bluesky" {
		return errors.New("--provider must be provisio or bluesky")
	}
	noBrowser, _ := cmd.Flags().GetBool("no-browser")
	withToken, _ := cmd.Flags().GetBool("with-token")
	server, _ := cmd.Flags().GetString("server")
	if server == "" {
		server = serverURL()
	}
	if err := auth.CheckServer(server); err != nil {
		return err
	}
	store, err := auth.DefaultStore()
	if err != nil {
		return err
	}
	client := &auth.Client{Server: server}
	ctx := cmd.Context()

	var cred *auth.Credential
	if withToken {
		raw, err := io.ReadAll(io.LimitReader(cmd.InOrStdin(), 4096))
		if err != nil {
			return err
		}
		token := strings.TrimSpace(string(raw))
		if !strings.HasPrefix(token, "pvt_") {
			return fmt.Errorf("expected a personal token (pvt_...) on stdin")
		}
		info, err := client.Session(ctx, token)
		if err != nil {
			return fmt.Errorf("the server did not accept that token: %w", err)
		}
		cred = &auth.Credential{Server: server, PersonalToken: token, DID: info.DID, Handle: info.Handle}
	} else {
		host, _ := os.Hostname()
		client.Hostname = host
		client.Provider = provider
		dc, err := client.RequestDeviceCode(ctx)
		if err != nil {
			return err
		}
		fmt.Fprintf(out, "\nTo sign in, open %s\nand check that the code is %s\n\n", dc.VerificationURI, dc.UserCode)
		if !noBrowser && openBrowser(dc.VerificationURIComplete) {
			fmt.Fprintln(out, "Opened your browser. Waiting for approval...")
		} else {
			fmt.Fprintln(out, "Waiting for approval...")
		}
		pollCtx := ctx
		if dc.ExpiresIn > 0 {
			var cancel func()
			pollCtx, cancel = withTimeout(ctx, time.Duration(dc.ExpiresIn)*time.Second)
			defer cancel()
		}
		tr, err := client.PollToken(pollCtx, dc, nil)
		if err != nil {
			return err
		}
		now := time.Now()
		cred = &auth.Credential{
			Server:           server,
			AccessToken:      tr.AccessToken,
			AccessExpiresAt:  now.Add(time.Duration(tr.ExpiresIn) * time.Second),
			RefreshToken:     tr.RefreshToken,
			RefreshExpiresAt: now.Add(time.Duration(tr.RefreshExpiresIn) * time.Second),
			DID:              tr.DID,
			Handle:           tr.Handle,
		}
	}
	if err := store.Save(cred); err != nil {
		return err
	}
	fmt.Fprintf(out, "\n✓ Signed in as @%s (%s)\n", cred.Handle, cred.DID)
	if os.Getenv(auth.EnvToken) != "" {
		fmt.Fprintf(out, "  Note: %s is set and takes precedence over this sign-in.\n", auth.EnvToken)
	}
	fmt.Fprintln(out)
	return nil
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-635b25ac","gen":"sha256:3dd07349442894e371a607dad8f67454397b701f2bd4e8d8d0edfb17be96c5a0","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func openBrowser(url string) bool {\n\tvar c *exec.Cmd\n\tswitch runtime.GOOS {\n\tcase \"darwin\":\n\t\tc = exec.Command(\"open\", url)\n\tcase \"windows\":\n\t\tc = exec.Command(\"rundll32\", \"url.dll,FileProtocolHandler\", url)\n\tdefault:\n\t\tif os.Getenv(\"DISPLAY\") == \"\" && os.Getenv(\"WAYLAND_DISPLAY\") == \"\" {\n\t\t\treturn false\n\t\t}\n\t\tc = exec.Command(\"xdg-open\", url)\n\t}\n\treturn c.Start() == nil\n}","file":"cmd/provenance/login_run.go","name":"openBrowser"}}
func openBrowser(url string) bool {
	var c *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		c = exec.Command("open", url)
	case "windows":
		c = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	default:
		if os.Getenv("DISPLAY") == "" && os.Getenv("WAYLAND_DISPLAY") == "" {
			return false
		}
		c = exec.Command("xdg-open", url)
	}
	return c.Start() == nil
}

// /sedum:setFunction
