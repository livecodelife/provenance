package main

// sedum:anchor:imports
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-d1d789de","gen":"sha256:6202ce5f10060b0fa61443fc52df17f80e57f28bdc0e09ce5ef4e680dc50ebff","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/hooks_install_run.go","path":"fmt"},"kwargs":"sha256:66115f10bd755c5592d735d7638e1b5cb9eb7b87dc655fae01ef4de1aacca5b3","body":"sha256:bcbac51883cbf8064e9196b371cb880cee429866652447ad557533d817857609"}
import "fmt"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-d1d789de","gen":"sha256:6202ce5f10060b0fa61443fc52df17f80e57f28bdc0e09ce5ef4e680dc50ebff","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/hooks_install_run.go","path":"io"},"kwargs":"sha256:55bcc2ca7c8286afbd6661705b40c466736de77dd874f0694c2424d1d3be6e79","body":"sha256:ef8e8abf4ef5ee40a19dc85b49bf09c5df1fe32abfe834e0ab4c2175baf80ddf"}
import "io"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-d1d789de","gen":"sha256:6202ce5f10060b0fa61443fc52df17f80e57f28bdc0e09ce5ef4e680dc50ebff","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/hooks_install_run.go","path":"os"},"kwargs":"sha256:30045fb9c7c87ef6664bceade3ba7ef69a2b4941d20cb6c2a5040e3802b49b6b","body":"sha256:7c6e309eb2575d4aa451ec2ab83161d13e9ae03bbca3a4d2ecf30ce12adf0905"}
import "os"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-d1d789de","gen":"sha256:6202ce5f10060b0fa61443fc52df17f80e57f28bdc0e09ce5ef4e680dc50ebff","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/hooks_install_run.go","path":"path/filepath"},"kwargs":"sha256:6c6f8bfa794a74687a278aac0f99c7d386fb76aa3091879562e65aa785821095","body":"sha256:a6d196b558491e646be93f56a321f339068c64c76aa9d49bb18046e7f9c777a1"}
import "path/filepath"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-d1d789de","gen":"sha256:6202ce5f10060b0fa61443fc52df17f80e57f28bdc0e09ce5ef4e680dc50ebff","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/hooks_install_run.go","path":"github.com/spf13/cobra"},"kwargs":"sha256:6ebc96ba5c8fbc2a4bb97320fe201b267e5f958ec8aafc01dd0b7559302e18b4","body":"sha256:f361c02d0f8e1999cf3fb715393e16310eeec9debd5aaa13936af7617e2b46aa"}
import "github.com/spf13/cobra"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setConst {"v":2,"tier":"owned","record":"prov-2026-d1d789de","gen":"sha256:6202ce5f10060b0fa61443fc52df17f80e57f28bdc0e09ce5ef4e680dc50ebff","pkg":"sha256:51be1c18998d63cc558c253e211a6bd188d96cb251dad68c6ea88e03c2d5cc4c","id":{"file":"cmd/provenance/hooks_install_run.go","name":"preCommitHook"},"kwargs":"sha256:882ca8efc259623ccdf13bf8e5cdfef22e1ed6bc6eabf38422094458c7b654ff","body":"sha256:d5fc3105ba6b7cea63c555f11ad22ed1bd36809825496d545963312f928ffb63"}
const preCommitHook = `#!/bin/sh
# provenance pre-commit hook, installed by ` + "`" + `provenance hooks install` + "`" + `.
# Lints each staged provenance record with the config nearest to it, and runs
# a record's associated specs when it moves from open to implemented.

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

staged=$(git diff --cached --name-only --diff-filter=ACM | grep -E "prov-[0-9]{4}-[^/]*\.ya?ml$" | grep -v "/\.versions/")
if [ -z "$staged" ]; then
    exit 0
fi

# Walk up from a file to the nearest .provenance.yml, else a .linespec.yml that
# has a provenance block (a LineSpec runner config is not a provenance root).
find_config() {
    dir=$(dirname "$1")
    while true; do
        for name in .provenance.yml .linespec.yml; do
            if [ -f "$dir/$name" ] && { [ "$name" = ".provenance.yml" ] || grep -q '^provenance:' "$dir/$name"; }; then
                echo "$dir/$name"
                return 0
            fi
        done
        if [ "$dir" = "." ] || [ "$dir" = "/" ]; then
            return 1
        fi
        dir=$(dirname "$dir")
    done
}

status_of() {
    git show "$1" 2>/dev/null | grep "^status:" | head -1 | awk '{print $2}'
}

exit_code=0
for file in $staged; do
    record=$(basename "$file" | sed -E 's/\.ya?ml$//')
    config=$(find_config "$file")
    if [ -n "$config" ]; then
        set -- -c "$config"
    else
        set --
    fi
    if ! "$PROV" "$@" lint --record "$record"; then
        echo "Commit blocked due to lint errors in $record"
        exit_code=1
    fi
    if [ "$(status_of "HEAD:$file")" = "open" ] && [ "$(status_of ":$file")" = "implemented" ]; then
        if ! "$PROV" "$@" run-specs --record "$record"; then
            exit_code=1
        fi
    fi
done

exit $exit_code
`

// /sedum:setConst
// sedum:setConst {"v":2,"tier":"owned","record":"prov-2026-d1d789de","gen":"sha256:6202ce5f10060b0fa61443fc52df17f80e57f28bdc0e09ce5ef4e680dc50ebff","pkg":"sha256:51be1c18998d63cc558c253e211a6bd188d96cb251dad68c6ea88e03c2d5cc4c","id":{"file":"cmd/provenance/hooks_install_run.go","name":"commitMsgHook"},"kwargs":"sha256:90dc861bbc50155d50904bbe781af56fe610658925e2623932d8d7848a6250bc","body":"sha256:d1b499132a82336c93c3c0c87b51b8a29c5bb9292c6ef5075eb933a5376212ef"}
const commitMsgHook = `#!/bin/sh
# provenance commit-msg hook, installed by ` + "`" + `provenance hooks install` + "`" + `.
# Checks the staged files against provenance scope for every config they fall
# under. A config outside the root applies only when files under its
# directory are staged.

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

COMMIT_MSG_FILE="$1"

# A .linespec.yml is a root only when it has a provenance block.
dirs=$(find . \( -name .provenance.yml -o -name .linespec.yml \) -not -path "./.git/*" 2>/dev/null | while read -r f; do
    case "$f" in
        */.provenance.yml) echo "$f" ;;
        *) grep -q '^provenance:' "$f" && echo "$f" ;;
    esac
done | xargs -n1 dirname 2>/dev/null | sort -u)

if [ -z "$dirs" ]; then
    if ! "$PROV" check --staged --message-file "$COMMIT_MSG_FILE"; then
        echo ""
        echo "Commit blocked due to provenance scope violations"
        exit 1
    fi
    exit 0
fi

exit_code=0
for dir in $dirs; do
    if [ -f "$dir/.provenance.yml" ]; then
        config="$dir/.provenance.yml"
    else
        config="$dir/.linespec.yml"
    fi
    rel=$(echo "$dir" | sed 's|^\./||')
    if [ "$rel" != "." ]; then
        if [ -z "$(git diff --cached --name-only | grep "^$rel/" | head -1)" ]; then
            continue
        fi
    fi
    if ! "$PROV" -c "$config" check --staged --message-file "$COMMIT_MSG_FILE"; then
        exit_code=1
    fi
done

if [ $exit_code -ne 0 ]; then
    echo ""
    echo "Commit blocked due to provenance scope violations"
fi

exit $exit_code
`

// /sedum:setConst
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-d1d789de","gen":"sha256:6202ce5f10060b0fa61443fc52df17f80e57f28bdc0e09ce5ef4e680dc50ebff","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/hooks_install_run.go","name":"installHook"},"kwargs":"sha256:77b05662a528809cde464d6f9d876d97f18fa90c062775db98b84208b549f62d","body":"sha256:f0904a07e9ae5c404a3e1b1367e28b500b4e0fd50c8f14c7a5cd8620a35e6fb4"}
func installHook(dir, name, body string) (string, error) {
	path := filepath.Join(dir, name)
	result := "installed"
	if existing, err := os.ReadFile(path); err == nil {
		if string(existing) == body {
			return "unchanged", nil
		}
		if _, err := os.Stat(path + ".bak"); os.IsNotExist(err) {
			if err := os.Rename(path, path+".bak"); err != nil {
				return "", err
			}
			result = "replaced (previous hook kept as " + name + ".bak)"
		} else {
			result = "replaced"
		}
	}
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return "", err
	}
	if err := os.WriteFile(path, []byte(body), 0o755); err != nil {
		return "", err
	}
	if err := os.Chmod(path, 0o755); err != nil {
		return "", err
	}
	return result, nil
}

// /sedum:setFunction
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-d1d789de","gen":"sha256:6202ce5f10060b0fa61443fc52df17f80e57f28bdc0e09ce5ef4e680dc50ebff","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/hooks_install_run.go","name":"runHooksInstall"},"kwargs":"sha256:6d925df3134cf27932cbb01f396fbadd148e868c5ab677a8beadcd3a97cb2864","body":"sha256:a339e50d2db19d034ef7ee700a652b781332bf17f3a2ab5e4814b988f409efa3"}
func runHooksInstall(out io.Writer, cmd *cobra.Command, args []string) error {
	r, err := openRepo(cmd)
	if err != nil {
		return err
	}
	dir, err := r.git("rev-parse", "--git-path", "hooks")
	if err != nil {
		return err
	}
	if !filepath.IsAbs(dir) {
		dir = filepath.Join(r.Root, dir)
	}
	preResult, err := installHook(dir, "pre-commit", preCommitHook)
	if err != nil {
		return err
	}
	msgResult, err := installHook(dir, "commit-msg", commitMsgHook)
	if err != nil {
		return err
	}
	pushResult, err := installHook(dir, "pre-push", prePushHook)
	if err != nil {
		return err
	}
	mergeResult, err := ensureMergeDriver(r.Root)
	if err != nil {
		return err
	}
	fmt.Fprintf(out, "\n✓ Installed git hooks to %s\n\n", dir)
	fmt.Fprintf(out, "  pre-commit (%s):\n    · Lints staged provenance records (nearest .provenance.yml, else .linespec.yml)\n    · Runs a record's associated specs when it moves open -> implemented\n\n", preResult)
	fmt.Fprintf(out, "  commit-msg (%s):\n    · Checks staged files against provenance scope (provenance check --staged)\n    · Validates provenance tags, including [id@version]\n\n", msgResult)
	fmt.Fprintf(out, "  pre-push (%s):\n    · Runs pre_push_command (when the repository sets one) in a clean checkout of every pushed commit\n\n", pushResult)
	fmt.Fprintf(out, "  merge driver (%s):\n    · Merges .provenance/versions.jsonl with provenance merge-manifest (.gitattributes and git config)\n\n", mergeResult)
	fmt.Fprint(out, "  The hooks run $PROVENANCE, else provenance on PATH.\n\n")
	if mergeResult != "unchanged" {
		fmt.Fprint(out, "  Commit .gitattributes so every clone gets the merge driver:\n    git add .gitattributes && git commit -m \"Merge the version manifest with provenance merge-manifest\"\n\n")
	}
	return nil
}

// /sedum:setFunction
