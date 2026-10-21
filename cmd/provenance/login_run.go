package main

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-24bcc587","gen":"sha256:bfad0894f57a43e6157a73b53759ec52ada87137ad1f9a464c1e36e3db2b4bf1","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/login_run.go","path":"errors"}}
import "errors"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-24bcc587","gen":"sha256:bfad0894f57a43e6157a73b53759ec52ada87137ad1f9a464c1e36e3db2b4bf1","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/login_run.go","path":"fmt"}}
import "fmt"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-24bcc587","gen":"sha256:bfad0894f57a43e6157a73b53759ec52ada87137ad1f9a464c1e36e3db2b4bf1","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/login_run.go","path":"io"}}
import "io"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-24bcc587","gen":"sha256:bfad0894f57a43e6157a73b53759ec52ada87137ad1f9a464c1e36e3db2b4bf1","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/login_run.go","path":"os"}}
import "os"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-24bcc587","gen":"sha256:bfad0894f57a43e6157a73b53759ec52ada87137ad1f9a464c1e36e3db2b4bf1","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/login_run.go","path":"strings"}}
import "strings"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-24bcc587","gen":"sha256:bfad0894f57a43e6157a73b53759ec52ada87137ad1f9a464c1e36e3db2b4bf1","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/login_run.go","path":"time"}}
import "time"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-24bcc587","gen":"sha256:bfad0894f57a43e6157a73b53759ec52ada87137ad1f9a464c1e36e3db2b4bf1","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/login_run.go","path":"github.com/livecodelife/provenance/auth"}}
import "github.com/livecodelife/provenance/auth"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-24bcc587","gen":"sha256:bfad0894f57a43e6157a73b53759ec52ada87137ad1f9a464c1e36e3db2b4bf1","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/login_run.go","path":"github.com/spf13/cobra"}}
import "github.com/spf13/cobra"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-976f20d3","gen":"sha256:8a68f83d469030f32a001b925994d0b6f56c32a7616883a7f8ccb325a0eee851","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/login_run.go","path":"os/exec"}}
import "os/exec"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-976f20d3","gen":"sha256:8a68f83d469030f32a001b925994d0b6f56c32a7616883a7f8ccb325a0eee851","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/login_run.go","path":"runtime"}}
import "runtime"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"tier":"owned","record":"prov-2026-24bcc587","gen":"sha256:bfad0894f57a43e6157a73b53759ec52ada87137ad1f9a464c1e36e3db2b4bf1","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func runLogin(out io.Writer, cmd *cobra.Command, args []string) error {\n\tprovider, _ := cmd.Flags().GetString(\"provider\")\n\tif !cmd.Flags().Changed(\"provider\") {\n\t\tprovider = \"provisio\"\n\t}\n\tif provider != \"provisio\" && provider != \"bluesky\" {\n\t\treturn errors.New(\"--provider must be provisio or bluesky\")\n\t}\n\n\tnoBrowser, _ := cmd.Flags().GetBool(\"no-browser\")\n\twithToken, _ := cmd.Flags().GetBool(\"with-token\")\n\tserver, _ := cmd.Flags().GetString(\"server\")\n\tif server == \"\" {\n\t\tserver = serverURL()\n\t}\n\n\tif err := auth.CheckServer(server); err != nil {\n\t\treturn err\n\t}\n\n\tstore, err := auth.DefaultStore()\n\tif err != nil {\n\t\treturn err\n\t}\n\n\tctx := cmd.Context()\n\tclient := &auth.Client{Server: server}\n\tvar cred *auth.Credential\n\n\tif withToken {\n\t\tdata, err := io.ReadAll(io.LimitReader(cmd.InOrStdin(), 4096))\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n\t\ttoken := strings.TrimSpace(string(data))\n\t\tif !strings.HasPrefix(token, \"pvt_\") {\n\t\t\treturn errors.New(\"expected a personal token (pvt_...) on stdin\")\n\t\t}\n\t\tinfo, err := client.Session(ctx, token)\n\t\tif err != nil {\n\t\t\treturn fmt.Errorf(\"the server did not accept that token: %w\", err)\n\t\t}\n\t\tcred = &auth.Credential{\n\t\t\tServer:        server,\n\t\t\tPersonalToken: token,\n\t\t\tDID:           info.DID,\n\t\t\tHandle:        info.Handle,\n\t\t}\n\t} else {\n\t\thostname, _ := os.Hostname()\n\t\tclient.Hostname = hostname\n\t\tclient.Provider = provider\n\n\t\tdc, err := client.RequestDeviceCode(ctx)\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n\n\t\tfmt.Fprintln(out)\n\t\tfmt.Fprintln(out, \"To sign in, open \"+dc.VerificationURI)\n\t\tfmt.Fprintln(out, \"and check that the code is \"+dc.UserCode)\n\t\tfmt.Fprintln(out)\n\n\t\topened := false\n\t\tif !noBrowser {\n\t\t\topened = openBrowser(dc.VerificationURIComplete)\n\t\t}\n\t\tif opened {\n\t\t\tfmt.Fprintln(out, \"Opened your browser. Waiting for approval...\")\n\t\t} else {\n\t\t\tfmt.Fprintln(out, \"Waiting for approval...\")\n\t\t}\n\n\t\tpollCtx, cancel := ctx, func() {}\n\t\tif dc.ExpiresIn > 0 {\n\t\t\tpollCtx, cancel = withTimeout(ctx, time.Duration(dc.ExpiresIn)*time.Second)\n\t\t}\n\t\tdefer cancel()\n\n\t\ttok, err := client.PollToken(pollCtx, dc, nil)\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n\n\t\tnow := time.Now()\n\t\tcred = &auth.Credential{\n\t\t\tServer:           server,\n\t\t\tAccessToken:      tok.AccessToken,\n\t\t\tRefreshToken:     tok.RefreshToken,\n\t\t\tDID:              tok.DID,\n\t\t\tHandle:           tok.Handle,\n\t\t\tAccessExpiresAt:  now.Add(time.Duration(tok.ExpiresIn) * time.Second),\n\t\t\tRefreshExpiresAt: now.Add(time.Duration(tok.RefreshExpiresIn) * time.Second),\n\t\t}\n\t}\n\n\tif err := store.Save(cred); err != nil {\n\t\treturn err\n\t}\n\n\tfmt.Fprintln(out)\n\tfmt.Fprintf(out, \"\\u2713 Signed in as @%s (%s)\\n\", cred.Handle, cred.DID)\n\tif os.Getenv(auth.EnvToken) != \"\" {\n\t\tfmt.Fprintln(out, \"  Note: PROVISIO_TOKEN is set and takes precedence over this sign-in.\")\n\t}\n\tfmt.Fprintln(out)\n\treturn nil\n}","file":"cmd/provenance/login_run.go","name":"runLogin"}}
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

	ctx := cmd.Context()
	client := &auth.Client{Server: server}
	var cred *auth.Credential

	if withToken {
		data, err := io.ReadAll(io.LimitReader(cmd.InOrStdin(), 4096))
		if err != nil {
			return err
		}
		token := strings.TrimSpace(string(data))
		if !strings.HasPrefix(token, "pvt_") {
			return errors.New("expected a personal token (pvt_...) on stdin")
		}
		info, err := client.Session(ctx, token)
		if err != nil {
			return fmt.Errorf("the server did not accept that token: %w", err)
		}
		cred = &auth.Credential{
			Server:        server,
			PersonalToken: token,
			DID:           info.DID,
			Handle:        info.Handle,
		}
	} else {
		hostname, _ := os.Hostname()
		client.Hostname = hostname
		client.Provider = provider

		dc, err := client.RequestDeviceCode(ctx)
		if err != nil {
			return err
		}

		fmt.Fprintln(out)
		fmt.Fprintln(out, "To sign in, open "+dc.VerificationURI)
		fmt.Fprintln(out, "and check that the code is "+dc.UserCode)
		fmt.Fprintln(out)

		opened := false
		if !noBrowser {
			opened = openBrowser(dc.VerificationURIComplete)
		}
		if opened {
			fmt.Fprintln(out, "Opened your browser. Waiting for approval...")
		} else {
			fmt.Fprintln(out, "Waiting for approval...")
		}

		pollCtx, cancel := ctx, func() {}
		if dc.ExpiresIn > 0 {
			pollCtx, cancel = withTimeout(ctx, time.Duration(dc.ExpiresIn)*time.Second)
		}
		defer cancel()

		tok, err := client.PollToken(pollCtx, dc, nil)
		if err != nil {
			return err
		}

		now := time.Now()
		cred = &auth.Credential{
			Server:           server,
			AccessToken:      tok.AccessToken,
			RefreshToken:     tok.RefreshToken,
			DID:              tok.DID,
			Handle:           tok.Handle,
			AccessExpiresAt:  now.Add(time.Duration(tok.ExpiresIn) * time.Second),
			RefreshExpiresAt: now.Add(time.Duration(tok.RefreshExpiresIn) * time.Second),
		}
	}

	if err := store.Save(cred); err != nil {
		return err
	}

	fmt.Fprintln(out)
	fmt.Fprintf(out, "\u2713 Signed in as @%s (%s)\n", cred.Handle, cred.DID)
	if os.Getenv(auth.EnvToken) != "" {
		fmt.Fprintln(out, "  Note: PROVISIO_TOKEN is set and takes precedence over this sign-in.")
	}
	fmt.Fprintln(out)
	return nil
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-976f20d3","gen":"sha256:8a68f83d469030f32a001b925994d0b6f56c32a7616883a7f8ccb325a0eee851","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func openBrowser(url string) bool {\n\tvar cmd *exec.Cmd\n\tswitch runtime.GOOS {\n\tcase \"darwin\":\n\t\tcmd = exec.Command(\"open\", url)\n\tcase \"windows\":\n\t\tcmd = exec.Command(\"rundll32\", \"url.dll,FileProtocolHandler\", url)\n\tdefault:\n\t\tif os.Getenv(\"DISPLAY\") == \"\" && os.Getenv(\"WAYLAND_DISPLAY\") == \"\" {\n\t\t\treturn false\n\t\t}\n\t\tcmd = exec.Command(\"xdg-open\", url)\n\t}\n\tif err := cmd.Start(); err != nil {\n\t\treturn false\n\t}\n\treturn true\n}","file":"cmd/provenance/login_run.go","name":"openBrowser"}}
func openBrowser(url string) bool {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", url)
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	default:
		if os.Getenv("DISPLAY") == "" && os.Getenv("WAYLAND_DISPLAY") == "" {
			return false
		}
		cmd = exec.Command("xdg-open", url)
	}
	if err := cmd.Start(); err != nil {
		return false
	}
	return true
}

// /sedum:setFunction
