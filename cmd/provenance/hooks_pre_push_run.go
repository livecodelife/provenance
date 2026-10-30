package main

// sedum:anchor:imports
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-10bc3e89","gen":"sha256:e93ace3473e7af7dd8c92e4f660130b255ccb6bbefff31577a90dcb3982e21d4","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/hooks_pre_push_run.go","path":"bufio"},"kwargs":"sha256:c627fe60d2c39a8dfc867c4c385ed89f40d50c3664598eb050d0a5da7c025ff9","body":"sha256:f392686aa91926837d96af56e8fce4f794d441bdf829d8efba66160f33dfd873"}
import "bufio"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-10bc3e89","gen":"sha256:e93ace3473e7af7dd8c92e4f660130b255ccb6bbefff31577a90dcb3982e21d4","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/hooks_pre_push_run.go","path":"fmt"},"kwargs":"sha256:c397d2dd293b21c36bc32128e0eee749066e9ad686cbbd36459b4ca6c5512ce0","body":"sha256:bcbac51883cbf8064e9196b371cb880cee429866652447ad557533d817857609"}
import "fmt"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-10bc3e89","gen":"sha256:e93ace3473e7af7dd8c92e4f660130b255ccb6bbefff31577a90dcb3982e21d4","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/hooks_pre_push_run.go","path":"io"},"kwargs":"sha256:0f7818f491334eea8b804a402868e5f3ad1fd56ff589a955249ec0f3da392720","body":"sha256:ef8e8abf4ef5ee40a19dc85b49bf09c5df1fe32abfe834e0ab4c2175baf80ddf"}
import "io"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-10bc3e89","gen":"sha256:e93ace3473e7af7dd8c92e4f660130b255ccb6bbefff31577a90dcb3982e21d4","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/hooks_pre_push_run.go","path":"os/exec"},"kwargs":"sha256:1ae29f1ac74ff5148c03a151546e5dbec71efe7298f05ff874976e5a5a10c4bc","body":"sha256:5970e9ca793f55e2beb08de623e98bdbc0ced178860fedd953eef7e15026ee19"}
import "os/exec"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-10bc3e89","gen":"sha256:e93ace3473e7af7dd8c92e4f660130b255ccb6bbefff31577a90dcb3982e21d4","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/hooks_pre_push_run.go","path":"strings"},"kwargs":"sha256:fd23020199cd0b216fe199a277d0765eab7798dfdee5c039efadcd613506154a","body":"sha256:a9af3c0bc4bf9342df9d7e89a324523fdba7941a9ca4190270e68ff3c4a25ea6"}
import "strings"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-10bc3e89","gen":"sha256:e93ace3473e7af7dd8c92e4f660130b255ccb6bbefff31577a90dcb3982e21d4","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/hooks_pre_push_run.go","path":"github.com/spf13/cobra"},"kwargs":"sha256:40f2edee26dae8d4152d9aba2a3e977536d8fc888cc2a0b7aab11e3081e64490","body":"sha256:f361c02d0f8e1999cf3fb715393e16310eeec9debd5aaa13936af7617e2b46aa"}
import "github.com/spf13/cobra"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setConst {"v":2,"tier":"owned","record":"prov-2026-10bc3e89","gen":"sha256:e93ace3473e7af7dd8c92e4f660130b255ccb6bbefff31577a90dcb3982e21d4","pkg":"sha256:51be1c18998d63cc558c253e211a6bd188d96cb251dad68c6ea88e03c2d5cc4c","id":{"file":"cmd/provenance/hooks_pre_push_run.go","name":"prePushHook"},"kwargs":"sha256:782bcac858d79ab7921e46cfd67238b65633e5484116de4977a48d7b4c8a9d38","body":"sha256:5897072a1318814bf13b2b526c4ba34659c90ae1357efebfc62dbcc244722ee1"}
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
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-10bc3e89","gen":"sha256:e93ace3473e7af7dd8c92e4f660130b255ccb6bbefff31577a90dcb3982e21d4","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/hooks_pre_push_run.go","name":"pushedCommits"},"kwargs":"sha256:c8c76b7669c6a4d48f85a95135e440c466a5dfb426a75df767c7952cc2f4b54d","body":"sha256:7b305304a0c8183c35db8f2623aa7558921039656fd10ea4d0a087b4b99e5965"}
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
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-10bc3e89","gen":"sha256:e93ace3473e7af7dd8c92e4f660130b255ccb6bbefff31577a90dcb3982e21d4","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/hooks_pre_push_run.go","name":"runHooksPrePush"},"kwargs":"sha256:7c5234fed7a0033a202c5143cb25a1d887c52ef7d34897f917a00be6a7f446cc","body":"sha256:d3e2bd5f555c2c02cb5a5a429fe68a5fee829786151874c10d6e18d55b317cc8"}
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
