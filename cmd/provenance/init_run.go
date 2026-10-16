package main

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-1a5d2549","gen":"sha256:591388b47574a3d570ed1b3a39ebceed0735baf1a2aca6d2bcf67109a0e9879a","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/init_run.go","path":"fmt"}}
import "fmt"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-1a5d2549","gen":"sha256:591388b47574a3d570ed1b3a39ebceed0735baf1a2aca6d2bcf67109a0e9879a","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/init_run.go","path":"io"}}
import "io"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-1a5d2549","gen":"sha256:591388b47574a3d570ed1b3a39ebceed0735baf1a2aca6d2bcf67109a0e9879a","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/init_run.go","path":"os"}}
import "os"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-1a5d2549","gen":"sha256:591388b47574a3d570ed1b3a39ebceed0735baf1a2aca6d2bcf67109a0e9879a","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/init_run.go","path":"os/exec"}}
import "os/exec"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-1a5d2549","gen":"sha256:591388b47574a3d570ed1b3a39ebceed0735baf1a2aca6d2bcf67109a0e9879a","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/init_run.go","path":"path/filepath"}}
import "path/filepath"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-1a5d2549","gen":"sha256:591388b47574a3d570ed1b3a39ebceed0735baf1a2aca6d2bcf67109a0e9879a","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/init_run.go","path":"strings"}}
import "strings"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-1a5d2549","gen":"sha256:591388b47574a3d570ed1b3a39ebceed0735baf1a2aca6d2bcf67109a0e9879a","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/init_run.go","path":"github.com/spf13/cobra"}}
import "github.com/spf13/cobra"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setConst {"tier":"owned","record":"prov-2026-1a5d2549","gen":"sha256:591388b47574a3d570ed1b3a39ebceed0735baf1a2aca6d2bcf67109a0e9879a","pkg":"sha256:51be1c18998d63cc558c253e211a6bd188d96cb251dad68c6ea88e03c2d5cc4c","kwargs":{"code":"const initConfig = `# provenance configuration, written by ` + \"`\" + `provenance init` + \"`\" + `.\n\n# How strictly records are enforced: none, warn or strict. Start with warn; move to\n# strict when your records have specs and you want violations to fail lint and commits.\nenforcement: warn\n\n# Where the records live, relative to this file.\ndir: provenance\n\n# Paths that need no record: exact paths, directory prefixes, globs or /regex/.\nexclude_paths: ['.gitignore', 'README.md', '.provenance.yml']\n\n# Records of other repositories that yours may implement (read-only; ` + \"`\" + `provenance sync` + \"`\" + `\n# fetches them). Example:\n#   shared_repos:\n#     - name: shared\n#       url: git@github.com:your-org/shared.git\n#       ref: main\nshared_repos: []\n\n# Minutes a shared-repo cache stays fresh (0 means the default, 60).\ncache_ttl_minutes: 0\n\n# true: a commit that changes governed files must carry a tag such as\n# [prov-2026-1a2b3c4d] or [prov-2026-1a2b3c4d@2] (checked by the commit-msg hook).\ncommit_tag_required: false\n\n# true: ` + \"`\" + `provenance complete` + \"`\" + ` runs the record's associated specs and records evidence\n# (see ` + \"`\" + `provenance check --evidence` + \"`\" + `).\nrun_associated_specs_on_complete: true\n\n# What ` + \"`\" + `complete` + \"`\" + ` does when it touches files of other implemented records whose specs\n# now fail: block, warn or off.\noverlap_specs_on_complete: warn\n\n# true: open, complete and deprecate commit the record change themselves.\ncommit_on_status_change: false\n\n# false: no file is locked. true: every tracked file that no open record's scope\n# covers is made read-only (` + \"`\" + `provenance reconcile` + \"`\" + `), which suits a project once its\n# first records exist. Omit the key for the default, which is true.\nwrite_restriction: false\n\n# A command run in a clean checkout of every commit you push (the pre-push hook;\n# install it with ` + \"`\" + `provenance hooks install` + \"`\" + `). Leave it out to skip the check.\n# pre_push_command: go build ./...\n`","file":"cmd/provenance/init_run.go","name":"initConfig"}}
const initConfig = `# provenance configuration, written by ` + "`" + `provenance init` + "`" + `.

# How strictly records are enforced: none, warn or strict. Start with warn; move to
# strict when your records have specs and you want violations to fail lint and commits.
enforcement: warn

# Where the records live, relative to this file.
dir: provenance

# Paths that need no record: exact paths, directory prefixes, globs or /regex/.
exclude_paths: ['.gitignore', 'README.md', '.provenance.yml']

# Records of other repositories that yours may implement (read-only; ` + "`" + `provenance sync` + "`" + `
# fetches them). Example:
#   shared_repos:
#     - name: shared
#       url: git@github.com:your-org/shared.git
#       ref: main
shared_repos: []

# Minutes a shared-repo cache stays fresh (0 means the default, 60).
cache_ttl_minutes: 0

# true: a commit that changes governed files must carry a tag such as
# [prov-2026-1a2b3c4d] or [prov-2026-1a2b3c4d@2] (checked by the commit-msg hook).
commit_tag_required: false

# true: ` + "`" + `provenance complete` + "`" + ` runs the record's associated specs and records evidence
# (see ` + "`" + `provenance check --evidence` + "`" + `).
run_associated_specs_on_complete: true

# What ` + "`" + `complete` + "`" + ` does when it touches files of other implemented records whose specs
# now fail: block, warn or off.
overlap_specs_on_complete: warn

# true: open, complete and deprecate commit the record change themselves.
commit_on_status_change: false

# false: no file is locked. true: every tracked file that no open record's scope
# covers is made read-only (` + "`" + `provenance reconcile` + "`" + `), which suits a project once its
# first records exist. Omit the key for the default, which is true.
write_restriction: false

# A command run in a clean checkout of every commit you push (the pre-push hook;
# install it with ` + "`" + `provenance hooks install` + "`" + `). Leave it out to skip the check.
# pre_push_command: go build ./...
`

// /sedum:setConst
// sedum:setConst {"tier":"owned","record":"prov-2026-1a5d2549","gen":"sha256:591388b47574a3d570ed1b3a39ebceed0735baf1a2aca6d2bcf67109a0e9879a","pkg":"sha256:51be1c18998d63cc558c253e211a6bd188d96cb251dad68c6ea88e03c2d5cc4c","kwargs":{"code":"const initSedumContext = `# Context for sedum grow\n\nPass with: sedum grow --generators <generators dir> --records provenance --prompt-file .sedum/context.md --only <record id>\nAdvisory only: behaviour belongs in the record, and editing this file regenerates nothing.\nWrite here what the model needs but the records should not carry.\n\n## Conventions\n- (package conventions, output rules, error handling)\n\n## Helpers that exist (call them; do not redefine)\n- (signatures the generated code may call)\n`","file":"cmd/provenance/init_run.go","name":"initSedumContext"}}
const initSedumContext = `# Context for sedum grow

Pass with: sedum grow --generators <generators dir> --records provenance --prompt-file .sedum/context.md --only <record id>
Advisory only: behaviour belongs in the record, and editing this file regenerates nothing.
Write here what the model needs but the records should not carry.

## Conventions
- (package conventions, output rules, error handling)

## Helpers that exist (call them; do not redefine)
- (signatures the generated code may call)
`

// /sedum:setConst
// sedum:setFunction {"tier":"owned","record":"prov-2026-1a5d2549","gen":"sha256:591388b47574a3d570ed1b3a39ebceed0735baf1a2aca6d2bcf67109a0e9879a","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func initDirIsGit(dir string) bool {\n\tc := exec.Command(\"git\", \"-C\", dir, \"rev-parse\", \"--is-inside-work-tree\")\n\tc.Env = cleanGitEnv()\n\treturn c.Run() == nil\n}","file":"cmd/provenance/init_run.go","name":"initDirIsGit"}}
func initDirIsGit(dir string) bool {
	c := exec.Command("git", "-C", dir, "rev-parse", "--is-inside-work-tree")
	c.Env = cleanGitEnv()
	return c.Run() == nil
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-1a5d2549","gen":"sha256:591388b47574a3d570ed1b3a39ebceed0735baf1a2aca6d2bcf67109a0e9879a","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func runInit(out io.Writer, cmd *cobra.Command, args []string) error {\n\tdirFlag, _ := cmd.Flags().GetString(\"dir\")\n\tforce, _ := cmd.Flags().GetBool(\"force\")\n\tif dirFlag == \"\" {\n\t\tdirFlag = \".\"\n\t}\n\tdir, err := filepath.Abs(dirFlag)\n\tif err != nil {\n\t\treturn err\n\t}\n\tif info, err := os.Stat(dir); err != nil || !info.IsDir() {\n\t\treturn fmt.Errorf(\"--dir %q is not an existing directory\", dirFlag)\n\t}\n\tbin := os.Getenv(\"SEDUM_BIN\")\n\tif bin == \"\" {\n\t\tbin = \"sedum\"\n\t}\n\t_, lookErr := exec.LookPath(bin)\n\t_, genErr := os.Stat(filepath.Join(dir, \"generators\"))\n\twithSedum := lookErr == nil || genErr == nil\n\tcfgText := initConfigText(withSedum)\n\tcfgPath := filepath.Join(dir, configFile)\n\texisting, readErr := os.ReadFile(cfgPath)\n\thasConfig := readErr == nil\n\tvar lines []string\n\tswitch {\n\tcase hasConfig && string(existing) == cfgText:\n\t\tlines = append(lines, \"unchanged  \"+configFile)\n\tcase hasConfig && !force:\n\t\treturn fmt.Errorf(\"%s already has a %s: init does not overwrite it (use --force; the old file is saved as %s.bak)\", dir, configFile, configFile)\n\tdefault:\n\t\tif legacy, err := os.ReadFile(filepath.Join(dir, legacyConfigFile)); err == nil && !hasConfig && !force && strings.Contains(\"\\n\"+string(legacy), \"\\nprovenance:\") {\n\t\t\treturn fmt.Errorf(\"%s has a provenance block in %s: run provenance migrate to convert it (or use --force to write a fresh %s)\", dir, legacyConfigFile, configFile)\n\t\t}\n\t\tlabel := \"created    \"\n\t\tif hasConfig {\n\t\t\tif err := os.WriteFile(cfgPath+\".bak\", existing, 0o644); err != nil {\n\t\t\t\treturn err\n\t\t\t}\n\t\t\tlabel = \"overwrote  \"\n\t\t}\n\t\tif err := os.WriteFile(cfgPath, []byte(cfgText), 0o644); err != nil {\n\t\t\treturn err\n\t\t}\n\t\tlines = append(lines, label+configFile)\n\t}\n\trecDir := filepath.Join(dir, \"provenance\")\n\tif _, err := os.Stat(recDir); os.IsNotExist(err) {\n\t\tif err := os.MkdirAll(recDir, 0o755); err != nil {\n\t\t\treturn err\n\t\t}\n\t\tif err := os.WriteFile(filepath.Join(recDir, \".gitkeep\"), nil, 0o644); err != nil {\n\t\t\treturn err\n\t\t}\n\t\tlines = append(lines, \"created    provenance/.gitkeep\")\n\t}\n\tctxPath := filepath.Join(dir, \".sedum\", \"context.md\")\n\tif withSedum {\n\t\tif _, err := os.Stat(ctxPath); os.IsNotExist(err) {\n\t\t\tif err := os.MkdirAll(filepath.Dir(ctxPath), 0o755); err != nil {\n\t\t\t\treturn err\n\t\t\t}\n\t\t\tif err := os.WriteFile(ctxPath, []byte(initSedumContext), 0o644); err != nil {\n\t\t\t\treturn err\n\t\t\t}\n\t\t\tlines = append(lines, \"created    .sedum/context.md\")\n\t\t}\n\t}\n\tfmt.Fprintf(out, \"\\n✓ Initialised provenance in %s\\n\\n\", dir)\n\tfor _, l := range lines {\n\t\tfmt.Fprintf(out, \"    %s\\n\", l)\n\t}\n\tif withSedum {\n\t\tfmt.Fprint(out, \"\\n  sedum writes .sedum/run.log: add .sedum/run.log to your .gitignore.\\n\")\n\t}\n\tfmt.Fprint(out, \"\\nNext steps (run them yourself; init installs nothing):\\n\")\n\tfmt.Fprint(out, \"  1. provenance hooks install         install the pre-commit, commit-msg and pre-push git hooks\\n\")\n\tfmt.Fprint(out, \"  2. provenance install-skills        install the Claude Code provenance skill\\n\")\n\tfmt.Fprint(out, \"  3. provenance install-plugin        install the Claude Code provenance plugin\\n\")\n\tfmt.Fprint(out, \"  4. provenance create --type blueprint --no-edit --title \\\"...\\\"   start your first record\\n\")\n\tfmt.Fprint(out, \"  5. provenance next                  ask what to do next\\n\")\n\tif !initDirIsGit(dir) {\n\t\tfmt.Fprint(out, \"\\n  Note: this directory is not a git repository; the hooks and commit tags need git (git init).\\n\")\n\t}\n\tfmt.Fprint(out, \"\\n\")\n\treturn nil\n}","file":"cmd/provenance/init_run.go","name":"runInit"}}
func runInit(out io.Writer, cmd *cobra.Command, args []string) error {
	dirFlag, _ := cmd.Flags().GetString("dir")
	force, _ := cmd.Flags().GetBool("force")
	if dirFlag == "" {
		dirFlag = "."
	}
	dir, err := filepath.Abs(dirFlag)
	if err != nil {
		return err
	}
	if info, err := os.Stat(dir); err != nil || !info.IsDir() {
		return fmt.Errorf("--dir %q is not an existing directory", dirFlag)
	}
	bin := os.Getenv("SEDUM_BIN")
	if bin == "" {
		bin = "sedum"
	}
	_, lookErr := exec.LookPath(bin)
	_, genErr := os.Stat(filepath.Join(dir, "generators"))
	withSedum := lookErr == nil || genErr == nil
	cfgText := initConfigText(withSedum)
	cfgPath := filepath.Join(dir, configFile)
	existing, readErr := os.ReadFile(cfgPath)
	hasConfig := readErr == nil
	var lines []string
	switch {
	case hasConfig && string(existing) == cfgText:
		lines = append(lines, "unchanged  "+configFile)
	case hasConfig && !force:
		return fmt.Errorf("%s already has a %s: init does not overwrite it (use --force; the old file is saved as %s.bak)", dir, configFile, configFile)
	default:
		if legacy, err := os.ReadFile(filepath.Join(dir, legacyConfigFile)); err == nil && !hasConfig && !force && strings.Contains("\n"+string(legacy), "\nprovenance:") {
			return fmt.Errorf("%s has a provenance block in %s: run provenance migrate to convert it (or use --force to write a fresh %s)", dir, legacyConfigFile, configFile)
		}
		label := "created    "
		if hasConfig {
			if err := os.WriteFile(cfgPath+".bak", existing, 0o644); err != nil {
				return err
			}
			label = "overwrote  "
		}
		if err := os.WriteFile(cfgPath, []byte(cfgText), 0o644); err != nil {
			return err
		}
		lines = append(lines, label+configFile)
	}
	recDir := filepath.Join(dir, "provenance")
	if _, err := os.Stat(recDir); os.IsNotExist(err) {
		if err := os.MkdirAll(recDir, 0o755); err != nil {
			return err
		}
		if err := os.WriteFile(filepath.Join(recDir, ".gitkeep"), nil, 0o644); err != nil {
			return err
		}
		lines = append(lines, "created    provenance/.gitkeep")
	}
	ctxPath := filepath.Join(dir, ".sedum", "context.md")
	if withSedum {
		if _, err := os.Stat(ctxPath); os.IsNotExist(err) {
			if err := os.MkdirAll(filepath.Dir(ctxPath), 0o755); err != nil {
				return err
			}
			if err := os.WriteFile(ctxPath, []byte(initSedumContext), 0o644); err != nil {
				return err
			}
			lines = append(lines, "created    .sedum/context.md")
		}
	}
	fmt.Fprintf(out, "\n✓ Initialised provenance in %s\n\n", dir)
	for _, l := range lines {
		fmt.Fprintf(out, "    %s\n", l)
	}
	if withSedum {
		fmt.Fprint(out, "\n  sedum writes .sedum/run.log: add .sedum/run.log to your .gitignore.\n")
	}
	fmt.Fprint(out, "\nNext steps (run them yourself; init installs nothing):\n")
	fmt.Fprint(out, "  1. provenance hooks install         install the pre-commit, commit-msg and pre-push git hooks\n")
	fmt.Fprint(out, "  2. provenance install-skills        install the Claude Code provenance skill\n")
	fmt.Fprint(out, "  3. provenance install-plugin        install the Claude Code provenance plugin\n")
	fmt.Fprint(out, "  4. provenance create --type blueprint --no-edit --title \"...\"   start your first record\n")
	fmt.Fprint(out, "  5. provenance next                  ask what to do next\n")
	if !initDirIsGit(dir) {
		fmt.Fprint(out, "\n  Note: this directory is not a git repository; the hooks and commit tags need git (git init).\n")
	}
	fmt.Fprint(out, "\n")
	return nil
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-1a5d2549","gen":"sha256:591388b47574a3d570ed1b3a39ebceed0735baf1a2aca6d2bcf67109a0e9879a","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func initConfigText(withSedum bool) string {\n\tif !withSedum {\n\t\treturn initConfig\n\t}\n\tconst plain = \"exclude_paths: ['.gitignore', 'README.md', '.provenance.yml']\\n\"\n\tconst withGenerators = \"# generators/** and .sedum/** are sedum's tooling (its generator packages and prompt\\n# file), added here because sedum is present: they need no record.\\nexclude_paths: ['.gitignore', 'README.md', '.provenance.yml', 'generators/**', '.sedum/**']\\n\"\n\treturn strings.Replace(initConfig, plain, withGenerators, 1)\n}","file":"cmd/provenance/init_run.go","name":"initConfigText"}}
func initConfigText(withSedum bool) string {
	if !withSedum {
		return initConfig
	}
	const plain = "exclude_paths: ['.gitignore', 'README.md', '.provenance.yml']\n"
	const withGenerators = "# generators/** and .sedum/** are sedum's tooling (its generator packages and prompt\n# file), added here because sedum is present: they need no record.\nexclude_paths: ['.gitignore', 'README.md', '.provenance.yml', 'generators/**', '.sedum/**']\n"
	return strings.Replace(initConfig, plain, withGenerators, 1)
}

// /sedum:setFunction
