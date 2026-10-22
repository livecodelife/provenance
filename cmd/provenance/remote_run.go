package main

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-06aa4464","gen":"sha256:c5827b62399eabddd5c68f8bad7fcbb69d99111ce89e49c23cf775ff621cdc75","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/remote_run.go","path":"bufio"}}
import "bufio"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-06aa4464","gen":"sha256:c5827b62399eabddd5c68f8bad7fcbb69d99111ce89e49c23cf775ff621cdc75","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/remote_run.go","path":"fmt"}}
import "fmt"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-06aa4464","gen":"sha256:c5827b62399eabddd5c68f8bad7fcbb69d99111ce89e49c23cf775ff621cdc75","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/remote_run.go","path":"io"}}
import "io"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-06aa4464","gen":"sha256:c5827b62399eabddd5c68f8bad7fcbb69d99111ce89e49c23cf775ff621cdc75","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/remote_run.go","path":"os"}}
import "os"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-06aa4464","gen":"sha256:c5827b62399eabddd5c68f8bad7fcbb69d99111ce89e49c23cf775ff621cdc75","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/remote_run.go","path":"strings"}}
import "strings"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-06aa4464","gen":"sha256:c5827b62399eabddd5c68f8bad7fcbb69d99111ce89e49c23cf775ff621cdc75","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/remote_run.go","path":"github.com/spf13/cobra"}}
import "github.com/spf13/cobra"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-a63c4d65","gen":"sha256:79d39d88d3221684bef8857bc2a949a8373130aa092a734485d229dc925718dc","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/remote_run.go","path":"context"}}
import "context"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-a63c4d65","gen":"sha256:79d39d88d3221684bef8857bc2a949a8373130aa092a734485d229dc925718dc","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/remote_run.go","path":"github.com/livecodelife/provenance/auth"}}
import "github.com/livecodelife/provenance/auth"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-a63c4d65","gen":"sha256:79d39d88d3221684bef8857bc2a949a8373130aa092a734485d229dc925718dc","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/remote_run.go","path":"github.com/livecodelife/provenance/remote"}}
import "github.com/livecodelife/provenance/remote"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"tier":"owned","record":"prov-2026-06aa4464","gen":"sha256:c5827b62399eabddd5c68f8bad7fcbb69d99111ce89e49c23cf775ff621cdc75","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func confirm(cmd *cobra.Command, out io.Writer, prompt string) bool {\n\tin, ok := cmd.InOrStdin().(*os.File)\n\tif !ok {\n\t\treturn false\n\t}\n\tinfo, err := in.Stat()\n\tif err != nil || info.Mode()&os.ModeCharDevice == 0 {\n\t\treturn false\n\t}\n\tfmt.Fprint(out, prompt)\n\tline, _ := bufio.NewReader(in).ReadString('\\n')\n\tanswer := strings.ToLower(strings.TrimSpace(line))\n\treturn answer == \"y\" || answer == \"yes\"\n}","file":"cmd/provenance/remote_run.go","name":"confirm"}}
func confirm(cmd *cobra.Command, out io.Writer, prompt string) bool {
	in, ok := cmd.InOrStdin().(*os.File)
	if !ok {
		return false
	}
	info, err := in.Stat()
	if err != nil || info.Mode()&os.ModeCharDevice == 0 {
		return false
	}
	fmt.Fprint(out, prompt)
	line, _ := bufio.NewReader(in).ReadString('\n')
	answer := strings.ToLower(strings.TrimSpace(line))
	return answer == "y" || answer == "yes"
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-2b608452","gen":"sha256:b14179c9398678aa53b8d748c176a7ce22d68610ec7cafd0557c691e94470101","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func recordIDs(r *repo, args []string, all bool) ([]string, error) {\n\tif all {\n\t\treturn r.Store.List()\n\t}\n\tif len(args) == 0 {\n\t\treturn nil, fmt.Errorf(\"name the records, or pass --all\")\n\t}\n\tids := make([]string, 0, len(args))\n\tfor _, arg := range args {\n\t\trec, _, err := r.resolve(arg)\n\t\tif err != nil {\n\t\t\treturn nil, err\n\t\t}\n\t\tids = append(ids, rec.ID)\n\t}\n\treturn ids, nil\n}","file":"cmd/provenance/remote_run.go","name":"recordIDs"}}
func recordIDs(r *repo, args []string, all bool) ([]string, error) {
	if all {
		return r.Store.List()
	}
	if len(args) == 0 {
		return nil, fmt.Errorf("name the records, or pass --all")
	}
	ids := make([]string, 0, len(args))
	for _, arg := range args {
		rec, _, err := r.resolve(arg)
		if err != nil {
			return nil, err
		}
		ids = append(ids, rec.ID)
	}
	return ids, nil
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-a63c4d65","gen":"sha256:79d39d88d3221684bef8857bc2a949a8373130aa092a734485d229dc925718dc","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func apiClient() (*remote.Client, error) {\n\tmgr, err := sessionManager()\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\tbase := os.Getenv(remote.EnvBase)\n\tif base == \"\" {\n\t\tbase = remote.DefaultBase\n\t}\n\tif err := auth.CheckServer(base); err != nil {\n\t\treturn nil, err\n\t}\n\treturn &remote.Client{\n\t\tBase: base,\n\t\tAuth: func(ctx context.Context, fn func(string) error) error {\n\t\t\treturn sessionError(mgr.Do(ctx, fn))\n\t\t},\n\t}, nil\n}","file":"cmd/provenance/remote_run.go","name":"apiClient"}}
func apiClient() (*remote.Client, error) {
	mgr, err := sessionManager()
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
	return &remote.Client{
		Base: base,
		Auth: func(ctx context.Context, fn func(string) error) error {
			return sessionError(mgr.Do(ctx, fn))
		},
	}, nil
}

// /sedum:setFunction
