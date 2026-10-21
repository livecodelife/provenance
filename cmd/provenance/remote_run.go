package main

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-83d57ec1","gen":"sha256:17956fb64b77b361bcbf762257b64c9341f58a44819468b083c8b9e970540a94","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/remote_run.go","path":"bufio"}}
import "bufio"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-83d57ec1","gen":"sha256:17956fb64b77b361bcbf762257b64c9341f58a44819468b083c8b9e970540a94","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/remote_run.go","path":"context"}}
import "context"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-83d57ec1","gen":"sha256:17956fb64b77b361bcbf762257b64c9341f58a44819468b083c8b9e970540a94","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/remote_run.go","path":"fmt"}}
import "fmt"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-83d57ec1","gen":"sha256:17956fb64b77b361bcbf762257b64c9341f58a44819468b083c8b9e970540a94","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/remote_run.go","path":"io"}}
import "io"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-83d57ec1","gen":"sha256:17956fb64b77b361bcbf762257b64c9341f58a44819468b083c8b9e970540a94","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/remote_run.go","path":"os"}}
import "os"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-83d57ec1","gen":"sha256:17956fb64b77b361bcbf762257b64c9341f58a44819468b083c8b9e970540a94","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/remote_run.go","path":"strings"}}
import "strings"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-83d57ec1","gen":"sha256:17956fb64b77b361bcbf762257b64c9341f58a44819468b083c8b9e970540a94","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/remote_run.go","path":"github.com/spf13/cobra"}}
import "github.com/spf13/cobra"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-83d57ec1","gen":"sha256:17956fb64b77b361bcbf762257b64c9341f58a44819468b083c8b9e970540a94","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/remote_run.go","path":"github.com/livecodelife/provenance/auth"}}
import "github.com/livecodelife/provenance/auth"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-83d57ec1","gen":"sha256:17956fb64b77b361bcbf762257b64c9341f58a44819468b083c8b9e970540a94","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/remote_run.go","path":"github.com/livecodelife/provenance/remote"}}
import "github.com/livecodelife/provenance/remote"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"tier":"owned","record":"prov-2026-83d57ec1","gen":"sha256:17956fb64b77b361bcbf762257b64c9341f58a44819468b083c8b9e970540a94","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"// apiClient is a record-server client whose requests carry the session's\n// bearer token and fail with exit status 3 when there is no session.\nfunc apiClient() (*remote.Client, error) {\n\tm, err := sessionManager()\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\tbase := os.Getenv(remote.EnvBase)\n\tif base == \"\" {\n\t\tbase = remote.DefaultBase\n\t}\n\tif err := auth.CheckServer(base); err != nil {\n\t\treturn nil, err\n\t}\n\treturn &remote.Client{Base: base, Auth: func(ctx context.Context, fn func(string) error) error {\n\t\treturn sessionError(m.Do(ctx, fn))\n\t}}, nil\n}","file":"cmd/provenance/remote_run.go","name":"apiClient"}}
// apiClient is a record-server client whose requests carry the session's
// bearer token and fail with exit status 3 when there is no session.
func apiClient() (*remote.Client, error) {
	m, err := sessionManager()
	if err != nil {
		return nil, err
	}
	base := os.Getenv(remote.EnvBase)
	if base == "" {
		base = remote.DefaultBase
	}
	if err := auth.CheckServer(base); err != nil {
		return nil, err
	}
	return &remote.Client{Base: base, Auth: func(ctx context.Context, fn func(string) error) error {
		return sessionError(m.Do(ctx, fn))
	}}, nil
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-83d57ec1","gen":"sha256:17956fb64b77b361bcbf762257b64c9341f58a44819468b083c8b9e970540a94","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"// confirm asks prompt on out and reads the answer from the command's input.\n// It is false whenever the input is not a terminal: nothing outward-facing\n// happens unattended without an explicit --yes.\nfunc confirm(cmd *cobra.Command, out io.Writer, prompt string) bool {\n\tf, ok := cmd.InOrStdin().(*os.File)\n\tif !ok {\n\t\treturn false\n\t}\n\tif info, err := f.Stat(); err != nil || info.Mode()&os.ModeCharDevice == 0 {\n\t\treturn false\n\t}\n\tfmt.Fprint(out, prompt)\n\tline, _ := bufio.NewReader(f).ReadString('\\n')\n\tanswer := strings.ToLower(strings.TrimSpace(line))\n\treturn answer == \"y\" || answer == \"yes\"\n}","file":"cmd/provenance/remote_run.go","name":"confirm"}}
// confirm asks prompt on out and reads the answer from the command's input.
// It is false whenever the input is not a terminal: nothing outward-facing
// happens unattended without an explicit --yes.
func confirm(cmd *cobra.Command, out io.Writer, prompt string) bool {
	f, ok := cmd.InOrStdin().(*os.File)
	if !ok {
		return false
	}
	if info, err := f.Stat(); err != nil || info.Mode()&os.ModeCharDevice == 0 {
		return false
	}
	fmt.Fprint(out, prompt)
	line, _ := bufio.NewReader(f).ReadString('\n')
	answer := strings.ToLower(strings.TrimSpace(line))
	return answer == "y" || answer == "yes"
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-83d57ec1","gen":"sha256:17956fb64b77b361bcbf762257b64c9341f58a44819468b083c8b9e970540a94","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"// recordIDs resolves the record arguments (ids or aliases), or every record\n// when all is set; it fails when neither names any.\nfunc recordIDs(r *repo, args []string, all bool) ([]string, error) {\n\tif all {\n\t\treturn r.Store.List()\n\t}\n\tif len(args) == 0 {\n\t\treturn nil, fmt.Errorf(\"name the records, or pass --all\")\n\t}\n\tids := make([]string, 0, len(args))\n\tfor _, a := range args {\n\t\trec, _, err := r.resolve(a)\n\t\tif err != nil {\n\t\t\treturn nil, err\n\t\t}\n\t\tids = append(ids, rec.ID)\n\t}\n\treturn ids, nil\n}","file":"cmd/provenance/remote_run.go","name":"recordIDs"}}
// recordIDs resolves the record arguments (ids or aliases), or every record
// when all is set; it fails when neither names any.
func recordIDs(r *repo, args []string, all bool) ([]string, error) {
	if all {
		return r.Store.List()
	}
	if len(args) == 0 {
		return nil, fmt.Errorf("name the records, or pass --all")
	}
	ids := make([]string, 0, len(args))
	for _, a := range args {
		rec, _, err := r.resolve(a)
		if err != nil {
			return nil, err
		}
		ids = append(ids, rec.ID)
	}
	return ids, nil
}

// /sedum:setFunction
