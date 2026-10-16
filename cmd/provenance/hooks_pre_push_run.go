package main

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-10bc3e89","gen":"sha256:e93ace3473e7af7dd8c92e4f660130b255ccb6bbefff31577a90dcb3982e21d4","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/hooks_pre_push_run.go","path":"bufio"}}
import "bufio"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-10bc3e89","gen":"sha256:e93ace3473e7af7dd8c92e4f660130b255ccb6bbefff31577a90dcb3982e21d4","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/hooks_pre_push_run.go","path":"fmt"}}
import "fmt"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-10bc3e89","gen":"sha256:e93ace3473e7af7dd8c92e4f660130b255ccb6bbefff31577a90dcb3982e21d4","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/hooks_pre_push_run.go","path":"io"}}
import "io"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-10bc3e89","gen":"sha256:e93ace3473e7af7dd8c92e4f660130b255ccb6bbefff31577a90dcb3982e21d4","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/hooks_pre_push_run.go","path":"os/exec"}}
import "os/exec"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-10bc3e89","gen":"sha256:e93ace3473e7af7dd8c92e4f660130b255ccb6bbefff31577a90dcb3982e21d4","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/hooks_pre_push_run.go","path":"strings"}}
import "strings"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-10bc3e89","gen":"sha256:e93ace3473e7af7dd8c92e4f660130b255ccb6bbefff31577a90dcb3982e21d4","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/hooks_pre_push_run.go","path":"github.com/spf13/cobra"}}
import "github.com/spf13/cobra"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setConst {"tier":"owned","record":"prov-2026-10bc3e89","gen":"sha256:e93ace3473e7af7dd8c92e4f660130b255ccb6bbefff31577a90dcb3982e21d4","pkg":"sha256:51be1c18998d63cc558c253e211a6bd188d96cb251dad68c6ea88e03c2d5cc4c","kwargs":{"code":"const prePushHook = `#!/bin/sh\n# provenance pre-push hook, installed by ` + \"`\" + `provenance hooks install` + \"`\" + `.\n# Runs the repository's pre_push_command (.provenance.yml) in a clean checkout\n# of every commit being pushed, so a half-landed commit cannot be pushed.\n# Without pre_push_command the hook does nothing.\n\n# $PROVENANCE, else provenance on PATH. Never a file in the repo: a\n# hook must not run whatever is checked in under that name.\nif [ -n \"$PROVENANCE\" ]; then\n    PROV=\"$PROVENANCE\"\nelif command -v provenance >/dev/null 2>&1; then\n    PROV=\"provenance\"\nelse\n    echo \"provenance hook: no provenance binary. Set PROVENANCE or put provenance on PATH.\" >&2\n    exit 1\nfi\n\nexec \"$PROV\" hooks pre-push \"$@\"\n`","file":"cmd/provenance/hooks_pre_push_run.go","name":"prePushHook"}}
const prePushHook = `#!/bin/sh
# provenance pre-push hook, installed by ` + "`" + `provenance hooks install` + "`" + `.
# Runs the repository's pre_push_command (.provenance.yml) in a clean checkout
# of every commit being pushed, so a half-landed commit cannot be pushed.
# Without pre_push_command the hook does nothing.

# $PROVENANCE, else provenance on PATH. Never a file in the repo: a
# hook must not run whatever is checked in under that name.
if [ -n "$PROVENANCE" ]; then
    PROV="$PROVENANCE"
elif command -v provenance >/dev/null 2>&1; then
    PROV="provenance"
else
    echo "provenance hook: no provenance binary. Set PROVENANCE or put provenance on PATH." >&2
    exit 1
fi

exec "$PROV" hooks pre-push "$@"
`

// /sedum:setConst
// sedum:setFunction {"tier":"owned","record":"prov-2026-10bc3e89","gen":"sha256:e93ace3473e7af7dd8c92e4f660130b255ccb6bbefff31577a90dcb3982e21d4","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func pushedCommits(r *repo, local, remote string) ([]string, error) {\n\tvar list string\n\tvar err error\n\tif strings.Trim(remote, \"0\") != \"\" {\n\t\tlist, err = r.git(\"rev-list\", \"--reverse\", remote+\"..\"+local)\n\t}\n\tif strings.Trim(remote, \"0\") == \"\" || err != nil {\n\t\tlist, err = r.git(\"rev-list\", \"--reverse\", local, \"--not\", \"--remotes\")\n\t}\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\treturn strings.Fields(list), nil\n}","file":"cmd/provenance/hooks_pre_push_run.go","name":"pushedCommits"}}
func pushedCommits(r *repo, local, remote string) ([]string, error) {
	var list string
	var err error
	if strings.Trim(remote, "0") != "" {
		list, err = r.git("rev-list", "--reverse", remote+".."+local)
	}
	if strings.Trim(remote, "0") == "" || err != nil {
		list, err = r.git("rev-list", "--reverse", local, "--not", "--remotes")
	}
	if err != nil {
		return nil, err
	}
	return strings.Fields(list), nil
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-10bc3e89","gen":"sha256:e93ace3473e7af7dd8c92e4f660130b255ccb6bbefff31577a90dcb3982e21d4","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func runHooksPrePush(out io.Writer, cmd *cobra.Command, args []string) error {\n\tr, err := openRepo(cmd)\n\tif err != nil {\n\t\treturn err\n\t}\n\tcommand := strings.TrimSpace(r.Config.PrePushCommand)\n\tif command == \"\" {\n\t\treturn nil\n\t}\n\tseen := map[string]bool{}\n\tvar shas []string\n\tscanner := bufio.NewScanner(cmd.InOrStdin())\n\tfor scanner.Scan() {\n\t\tf := strings.Fields(scanner.Text())\n\t\tif len(f) < 4 || strings.Trim(f[1], \"0\") == \"\" {\n\t\t\tcontinue\n\t\t}\n\t\tlist, err := pushedCommits(r, f[1], f[3])\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n\t\tfor _, sha := range list {\n\t\t\tif !seen[sha] {\n\t\t\t\tseen[sha] = true\n\t\t\t\tshas = append(shas, sha)\n\t\t\t}\n\t\t}\n\t}\n\tfailed := 0\n\tfor _, sha := range shas {\n\t\ttree, cleanup, err := commitTree(r, sha)\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n\t\tc := exec.Command(\"sh\", \"-c\", command)\n\t\tc.Dir = tree\n\t\tc.Env = cleanGitEnv()\n\t\tc.Stdout = out\n\t\tc.Stderr = out\n\t\trunErr := c.Run()\n\t\tcleanup()\n\t\tif runErr != nil {\n\t\t\tfailed++\n\t\t\tfmt.Fprintf(out, \"\\n✗ %s does not pass pre_push_command (%s): %v\\n\", sha[:7], command, runErr)\n\t\t}\n\t}\n\tif failed > 0 {\n\t\treturn fmt.Errorf(\"%d of %d commits fail pre_push_command\", failed, len(shas))\n\t}\n\treturn nil\n}","file":"cmd/provenance/hooks_pre_push_run.go","name":"runHooksPrePush"}}
func runHooksPrePush(out io.Writer, cmd *cobra.Command, args []string) error {
	r, err := openRepo(cmd)
	if err != nil {
		return err
	}
	command := strings.TrimSpace(r.Config.PrePushCommand)
	if command == "" {
		return nil
	}
	seen := map[string]bool{}
	var shas []string
	scanner := bufio.NewScanner(cmd.InOrStdin())
	for scanner.Scan() {
		f := strings.Fields(scanner.Text())
		if len(f) < 4 || strings.Trim(f[1], "0") == "" {
			continue
		}
		list, err := pushedCommits(r, f[1], f[3])
		if err != nil {
			return err
		}
		for _, sha := range list {
			if !seen[sha] {
				seen[sha] = true
				shas = append(shas, sha)
			}
		}
	}
	failed := 0
	for _, sha := range shas {
		tree, cleanup, err := commitTree(r, sha)
		if err != nil {
			return err
		}
		c := exec.Command("sh", "-c", command)
		c.Dir = tree
		c.Env = cleanGitEnv()
		c.Stdout = out
		c.Stderr = out
		runErr := c.Run()
		cleanup()
		if runErr != nil {
			failed++
			fmt.Fprintf(out, "\n✗ %s does not pass pre_push_command (%s): %v\n", sha[:7], command, runErr)
		}
	}
	if failed > 0 {
		return fmt.Errorf("%d of %d commits fail pre_push_command", failed, len(shas))
	}
	return nil
}

// /sedum:setFunction
