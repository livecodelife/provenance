package main

// sedum:anchor:imports
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-1a5d2549","gen":"sha256:95782fcfe653edd478163f4fae54335b865cf426a0e0fccd5f1fc4c3bb633ffe","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/init_run.go","path":"fmt"},"kwargs":"sha256:3d52c655aed3b89b0f825b5e68b774c9a8123acb688d231fb46890f4e5f2f73e","body":"sha256:bcbac51883cbf8064e9196b371cb880cee429866652447ad557533d817857609"}
import "fmt"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-1a5d2549","gen":"sha256:95782fcfe653edd478163f4fae54335b865cf426a0e0fccd5f1fc4c3bb633ffe","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/init_run.go","path":"io"},"kwargs":"sha256:74a9770a3079d0309e520f1dd08745cee2a0d2fcd7a3662310eb040c4bb058b6","body":"sha256:ef8e8abf4ef5ee40a19dc85b49bf09c5df1fe32abfe834e0ab4c2175baf80ddf"}
import "io"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-1a5d2549","gen":"sha256:95782fcfe653edd478163f4fae54335b865cf426a0e0fccd5f1fc4c3bb633ffe","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/init_run.go","path":"os"},"kwargs":"sha256:c2a51725adad00bbae9ed9b945d049e053433a6d5499b7fab1103c61d02b67db","body":"sha256:7c6e309eb2575d4aa451ec2ab83161d13e9ae03bbca3a4d2ecf30ce12adf0905"}
import "os"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-1a5d2549","gen":"sha256:95782fcfe653edd478163f4fae54335b865cf426a0e0fccd5f1fc4c3bb633ffe","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/init_run.go","path":"os/exec"},"kwargs":"sha256:4b621f911bc6fac5d496ff442dcc21a01863e9de3ff9fe6926d4d856d5d8ffd6","body":"sha256:5970e9ca793f55e2beb08de623e98bdbc0ced178860fedd953eef7e15026ee19"}
import "os/exec"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-1a5d2549","gen":"sha256:95782fcfe653edd478163f4fae54335b865cf426a0e0fccd5f1fc4c3bb633ffe","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/init_run.go","path":"path/filepath"},"kwargs":"sha256:7d8ed00926d521ac7c4e57903f78caa6c3ade03ccd62a9ebcf54b92378a435e1","body":"sha256:a6d196b558491e646be93f56a321f339068c64c76aa9d49bb18046e7f9c777a1"}
import "path/filepath"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-1a5d2549","gen":"sha256:95782fcfe653edd478163f4fae54335b865cf426a0e0fccd5f1fc4c3bb633ffe","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/init_run.go","path":"strings"},"kwargs":"sha256:5ea795571990735a7fda69625e2dfcae570df8c1fc70706035bbce44e2537cd6","body":"sha256:a9af3c0bc4bf9342df9d7e89a324523fdba7941a9ca4190270e68ff3c4a25ea6"}
import "strings"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-1a5d2549","gen":"sha256:95782fcfe653edd478163f4fae54335b865cf426a0e0fccd5f1fc4c3bb633ffe","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/init_run.go","path":"github.com/spf13/cobra"},"kwargs":"sha256:8aeff4777ad76922eba88f34040fa1da56dc7c478c6e53dea4b56a4af1371778","body":"sha256:f361c02d0f8e1999cf3fb715393e16310eeec9debd5aaa13936af7617e2b46aa"}
import "github.com/spf13/cobra"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setConst {"v":2,"tier":"owned","record":"prov-2026-1a5d2549","gen":"sha256:95782fcfe653edd478163f4fae54335b865cf426a0e0fccd5f1fc4c3bb633ffe","pkg":"sha256:51be1c18998d63cc558c253e211a6bd188d96cb251dad68c6ea88e03c2d5cc4c","id":{"file":"cmd/provenance/init_run.go","name":"initConfig"},"kwargs":"sha256:4eeaaccae3ca3945d6801512645f6dc2edba44f4fc6e95ff0d135577283c818c","body":"sha256:e50cbcbe2c47a1640c9b511ec88207cf7494dc226d6d07519e49eb69ff5e02f8"}
const initConfig = `# provenance configuration, written by ` + "`" + `provenance init` + "`" + `.

# How strictly records are enforced: none, warn or strict. Start with warn; move to
# strict when your records have specs and you want violations to fail lint and commits.
enforcement: warn

# Where the records live, relative to this file.
dir: provenance

# Paths that need no record: exact paths, directory prefixes, globs or /regex/.
exclude_paths: ['.gitignore', 'README.md', '.provenance.yml', '.gitattributes']

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
# A spec's run_command is run as ` + "`" + `<run_command> <path>` + "`" + `, or with {{path}} substituted
# wherever it appears (so ` + "`" + `go test ./... -run X {{path}}` + "`" + ` and ` + "`" + `test -f {{path}}` + "`" + ` both work).
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
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-1a5d2549","gen":"sha256:95782fcfe653edd478163f4fae54335b865cf426a0e0fccd5f1fc4c3bb633ffe","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/init_run.go","name":"initDirIsGit"},"kwargs":"sha256:a798a5435548c9f6b90f18f478c01a682921a66c01e382aa6d4abad3fb7696db","body":"sha256:47ccbbc9b61def6f75aa2f136065f8cd501f9e8f1335bb14e7e6d270081c17d2"}
func initDirIsGit(dir string) bool {
	c := exec.Command("git", "-C", dir, "rev-parse", "--is-inside-work-tree")
	c.Env = cleanGitEnv()
	return c.Run() == nil
}

// /sedum:setFunction
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-1a5d2549","gen":"sha256:95782fcfe653edd478163f4fae54335b865cf426a0e0fccd5f1fc4c3bb633ffe","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/init_run.go","name":"runInit"},"kwargs":"sha256:a5fb54dbb48c9a41521e1b27c66200218dd010f24e163bb8e929c6dd19d3faa2","body":"sha256:1fd7a9e4b03320c8ee1c68ca49c33d1f6956f454a579bc7d7bceefaec4f758fd"}
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
	cfgText := initConfigText()
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
	attrStatus, err := ensureGitattributes(dir, "provenance")
	if err != nil {
		return err
	}
	lines = append(lines, fmt.Sprintf("%-11s.gitattributes", attrStatus))
	fmt.Fprintf(out, "\n✓ Initialised provenance in %s\n\n", dir)
	for _, l := range lines {
		fmt.Fprintf(out, "    %s\n", l)
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
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-1a5d2549","gen":"sha256:95782fcfe653edd478163f4fae54335b865cf426a0e0fccd5f1fc4c3bb633ffe","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/init_run.go","name":"initConfigText"},"kwargs":"sha256:5de09df1ad430762485eb3b7964f0b7ff73919d44c1077f21ac9df47a656eb45","body":"sha256:60265a2b6c6a71caf4d6b78c32c0842da1f1c563be8ff74e86f2d0ab23cad463"}
func initConfigText() string {
	return initConfig
}

// /sedum:setFunction
