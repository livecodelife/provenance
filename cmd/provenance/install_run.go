package main

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-d71d0066","gen":"sha256:cae7a2d3a080d43f1ba12537ce93365db84c39195d99ec904268e9c15daab6a9","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/install_run.go","path":"errors"}}
import "errors"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-d71d0066","gen":"sha256:cae7a2d3a080d43f1ba12537ce93365db84c39195d99ec904268e9c15daab6a9","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/install_run.go","path":"fmt"}}
import "fmt"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-d71d0066","gen":"sha256:cae7a2d3a080d43f1ba12537ce93365db84c39195d99ec904268e9c15daab6a9","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/install_run.go","path":"io"}}
import "io"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-d71d0066","gen":"sha256:cae7a2d3a080d43f1ba12537ce93365db84c39195d99ec904268e9c15daab6a9","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/install_run.go","path":"io/fs"}}
import "io/fs"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-d71d0066","gen":"sha256:cae7a2d3a080d43f1ba12537ce93365db84c39195d99ec904268e9c15daab6a9","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/install_run.go","path":"os"}}
import "os"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-d71d0066","gen":"sha256:cae7a2d3a080d43f1ba12537ce93365db84c39195d99ec904268e9c15daab6a9","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/install_run.go","path":"path/filepath"}}
import "path/filepath"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-d71d0066","gen":"sha256:cae7a2d3a080d43f1ba12537ce93365db84c39195d99ec904268e9c15daab6a9","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/install_run.go","path":"sort"}}
import "sort"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-d71d0066","gen":"sha256:cae7a2d3a080d43f1ba12537ce93365db84c39195d99ec904268e9c15daab6a9","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/install_run.go","path":"strings"}}
import "strings"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-d71d0066","gen":"sha256:cae7a2d3a080d43f1ba12537ce93365db84c39195d99ec904268e9c15daab6a9","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/install_run.go","path":"github.com/spf13/cobra"}}
import "github.com/spf13/cobra"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-d71d0066","gen":"sha256:cae7a2d3a080d43f1ba12537ce93365db84c39195d99ec904268e9c15daab6a9","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/install_run.go","path":"github.com/livecodelife/provenance/assets"}}
import "github.com/livecodelife/provenance/assets"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setType {"tier":"owned","record":"prov-2026-d71d0066","gen":"sha256:cae7a2d3a080d43f1ba12537ce93365db84c39195d99ec904268e9c15daab6a9","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type installItem struct {\n\tDest  string\n\tData  []byte\n\tMode  fs.FileMode\n\tState string\n}","file":"cmd/provenance/install_run.go","name":"installItem"}}
type installItem struct {
	Dest  string
	Data  []byte
	Mode  fs.FileMode
	State string
}

// /sedum:setType
// sedum:setFunction {"tier":"owned","record":"prov-2026-d71d0066","gen":"sha256:cae7a2d3a080d43f1ba12537ce93365db84c39195d99ec904268e9c15daab6a9","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func installBase(flag, def string) (string, error) {\n\tbase := flag\n\tif base == \"\" {\n\t\tbase = def\n\t}\n\tif filepath.IsAbs(base) {\n\t\treturn \"\", fmt.Errorf(\"--path %q must be relative to the repository root\", base)\n\t}\n\tclean := filepath.Clean(base)\n\tif clean == \"..\" || strings.HasPrefix(clean, \"..\"+string(filepath.Separator)) {\n\t\treturn \"\", fmt.Errorf(\"--path %q must stay inside the repository\", base)\n\t}\n\treturn clean, nil\n}","file":"cmd/provenance/install_run.go","name":"installBase"}}
func installBase(flag, def string) (string, error) {
	base := flag
	if base == "" {
		base = def
	}
	if filepath.IsAbs(base) {
		return "", fmt.Errorf("--path %q must be relative to the repository root", base)
	}
	clean := filepath.Clean(base)
	if clean == ".." || strings.HasPrefix(clean, ".."+string(filepath.Separator)) {
		return "", fmt.Errorf("--path %q must stay inside the repository", base)
	}
	return clean, nil
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-d71d0066","gen":"sha256:cae7a2d3a080d43f1ba12537ce93365db84c39195d99ec904268e9c15daab6a9","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func installPlan(r *repo, root, base string) ([]installItem, error) {\n\tvar items []installItem\n\terr := fs.WalkDir(assets.Files, root, func(p string, d fs.DirEntry, err error) error {\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n\t\tif d.IsDir() {\n\t\t\treturn nil\n\t\t}\n\t\tdata, err := assets.Files.ReadFile(p)\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n\t\trel := strings.TrimPrefix(p, root+\"/\")\n\t\tmode := fs.FileMode(0o644)\n\t\tif strings.HasSuffix(rel, \".sh\") {\n\t\t\tmode = 0o755\n\t\t}\n\t\tdest := filepath.Join(r.Root, base, filepath.FromSlash(rel))\n\t\tstate := \"create\"\n\t\tif have, err := os.ReadFile(dest); err == nil {\n\t\t\tstate = \"differs\"\n\t\t\tif string(have) == string(data) {\n\t\t\t\tstate = \"unchanged\"\n\t\t\t}\n\t\t}\n\t\titems = append(items, installItem{Dest: dest, Data: data, Mode: mode, State: state})\n\t\treturn nil\n\t})\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\tsort.Slice(items, func(i, j int) bool { return items[i].Dest < items[j].Dest })\n\treturn items, nil\n}","file":"cmd/provenance/install_run.go","name":"installPlan"}}
func installPlan(r *repo, root, base string) ([]installItem, error) {
	var items []installItem
	err := fs.WalkDir(assets.Files, root, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		data, err := assets.Files.ReadFile(p)
		if err != nil {
			return err
		}
		rel := strings.TrimPrefix(p, root+"/")
		mode := fs.FileMode(0o644)
		if strings.HasSuffix(rel, ".sh") {
			mode = 0o755
		}
		dest := filepath.Join(r.Root, base, filepath.FromSlash(rel))
		state := "create"
		if have, err := os.ReadFile(dest); err == nil {
			state = "differs"
			if string(have) == string(data) {
				state = "unchanged"
			}
		}
		items = append(items, installItem{Dest: dest, Data: data, Mode: mode, State: state})
		return nil
	})
	if err != nil {
		return nil, err
	}
	sort.Slice(items, func(i, j int) bool { return items[i].Dest < items[j].Dest })
	return items, nil
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-d71d0066","gen":"sha256:cae7a2d3a080d43f1ba12537ce93365db84c39195d99ec904268e9c15daab6a9","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func runInstall(out io.Writer, r *repo, root, base, what string, force bool) error {\n\titems, err := installPlan(r, root, base)\n\tif err != nil {\n\t\treturn err\n\t}\n\tif len(items) == 0 {\n\t\treturn errors.New(\"nothing to install: the embedded \" + what + \" are empty\")\n\t}\n\tfmt.Fprintf(out, \"\\nWill write %d file(s) under %s:\\n\", len(items), filepath.ToSlash(base))\n\tconflicts := 0\n\tfor _, item := range items {\n\t\tlabel := item.State\n\t\tif item.State == \"differs\" {\n\t\t\tlabel = \"overwrite\"\n\t\t\tif !force {\n\t\t\t\tlabel = \"CONFLICT\"\n\t\t\t\tconflicts++\n\t\t\t}\n\t\t}\n\t\trel, _ := filepath.Rel(r.Root, item.Dest)\n\t\tfmt.Fprintf(out, \"    %-10s %s\\n\", label, filepath.ToSlash(rel))\n\t}\n\tif conflicts > 0 {\n\t\tfmt.Fprint(out, \"\\n  Existing files differ; nothing was written. Re-run with --force to overwrite them.\\n\\n\")\n\t\treturn fmt.Errorf(\"%d existing file(s) differ\", conflicts)\n\t}\n\twritten := 0\n\tfor _, item := range items {\n\t\tif item.State == \"unchanged\" {\n\t\t\tcontinue\n\t\t}\n\t\tif err := os.MkdirAll(filepath.Dir(item.Dest), 0o755); err != nil {\n\t\t\treturn err\n\t\t}\n\t\tif err := os.WriteFile(item.Dest, item.Data, item.Mode); err != nil {\n\t\t\treturn err\n\t\t}\n\t\tif err := os.Chmod(item.Dest, item.Mode); err != nil {\n\t\t\treturn err\n\t\t}\n\t\twritten++\n\t}\n\tfmt.Fprintf(out, \"\\n✓ Installed %s: %d written, %d unchanged\\n\\n\", what, written, len(items)-written)\n\treturn nil\n}","file":"cmd/provenance/install_run.go","name":"runInstall"}}
func runInstall(out io.Writer, r *repo, root, base, what string, force bool) error {
	items, err := installPlan(r, root, base)
	if err != nil {
		return err
	}
	if len(items) == 0 {
		return errors.New("nothing to install: the embedded " + what + " are empty")
	}
	fmt.Fprintf(out, "\nWill write %d file(s) under %s:\n", len(items), filepath.ToSlash(base))
	conflicts := 0
	for _, item := range items {
		label := item.State
		if item.State == "differs" {
			label = "overwrite"
			if !force {
				label = "CONFLICT"
				conflicts++
			}
		}
		rel, _ := filepath.Rel(r.Root, item.Dest)
		fmt.Fprintf(out, "    %-10s %s\n", label, filepath.ToSlash(rel))
	}
	if conflicts > 0 {
		fmt.Fprint(out, "\n  Existing files differ; nothing was written. Re-run with --force to overwrite them.\n\n")
		return fmt.Errorf("%d existing file(s) differ", conflicts)
	}
	written := 0
	for _, item := range items {
		if item.State == "unchanged" {
			continue
		}
		if err := os.MkdirAll(filepath.Dir(item.Dest), 0o755); err != nil {
			return err
		}
		if err := os.WriteFile(item.Dest, item.Data, item.Mode); err != nil {
			return err
		}
		if err := os.Chmod(item.Dest, item.Mode); err != nil {
			return err
		}
		written++
	}
	fmt.Fprintf(out, "\n✓ Installed %s: %d written, %d unchanged\n\n", what, written, len(items)-written)
	return nil
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-d71d0066","gen":"sha256:cae7a2d3a080d43f1ba12537ce93365db84c39195d99ec904268e9c15daab6a9","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func runInstallSkills(out io.Writer, cmd *cobra.Command, args []string) error {\n\tr, err := openRepo(cmd)\n\tif err != nil {\n\t\treturn err\n\t}\n\tpath, _ := cmd.Flags().GetString(\"path\")\n\tforce, _ := cmd.Flags().GetBool(\"force\")\n\tbase, err := installBase(path, \".claude/skills\")\n\tif err != nil {\n\t\treturn err\n\t}\n\treturn runInstall(out, r, \"skills\", base, \"skills\", force)\n}","file":"cmd/provenance/install_run.go","name":"runInstallSkills"}}
func runInstallSkills(out io.Writer, cmd *cobra.Command, args []string) error {
	r, err := openRepo(cmd)
	if err != nil {
		return err
	}
	path, _ := cmd.Flags().GetString("path")
	force, _ := cmd.Flags().GetBool("force")
	base, err := installBase(path, ".claude/skills")
	if err != nil {
		return err
	}
	return runInstall(out, r, "skills", base, "skills", force)
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-d71d0066","gen":"sha256:cae7a2d3a080d43f1ba12537ce93365db84c39195d99ec904268e9c15daab6a9","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func runInstallPlugin(out io.Writer, cmd *cobra.Command, args []string) error {\n\tr, err := openRepo(cmd)\n\tif err != nil {\n\t\treturn err\n\t}\n\tpath, _ := cmd.Flags().GetString(\"path\")\n\tforce, _ := cmd.Flags().GetBool(\"force\")\n\tbase, err := installBase(path, \".claude/plugins\")\n\tif err != nil {\n\t\treturn err\n\t}\n\treturn runInstall(out, r, \"plugins\", base, \"plugin\", force)\n}","file":"cmd/provenance/install_run.go","name":"runInstallPlugin"}}
func runInstallPlugin(out io.Writer, cmd *cobra.Command, args []string) error {
	r, err := openRepo(cmd)
	if err != nil {
		return err
	}
	path, _ := cmd.Flags().GetString("path")
	force, _ := cmd.Flags().GetBool("force")
	base, err := installBase(path, ".claude/plugins")
	if err != nil {
		return err
	}
	return runInstall(out, r, "plugins", base, "plugin", force)
}

// /sedum:setFunction
