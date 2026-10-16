#!/bin/sh
# End-to-end done-check for prov-2026-409e803e (CLI part 2): on a temp clone
# of a v1 provenance repository, `provenance migrate`, `provenance hooks
# install`, then the commit lifecycle through the installed hooks.
#
# Usage: e2e.sh [source-repo]   (default: this repository)
#
# Git safety (COMMON.md "Git in tests and scripts"): every GIT_* variable is
# cleared, HOME and the global git config live in the temp dir, every git
# command names its repository with -C, and the clone's toplevel is checked
# before anything writes to it. The only directory ever removed is the one
# mktemp made.
set -eu

for v in $(env | sed -n 's/^\(GIT_[A-Za-z0-9_]*\)=.*/\1/p'); do
	unset "$v"
done

here=$(cd "$(dirname "$0")" && pwd)
root=$(cd "$here/../../.." && pwd)
src=${1:-$root}

work=$(mktemp -d "${TMPDIR:-/tmp}/provenance-e2e.XXXXXX")
cleanup() {
	case "$work" in
	*/provenance-e2e.*) rm -rf -- "$work" ;;
	esac
}
trap cleanup EXIT

# Build first, with the caller's Go caches: a temp HOME would make go fetch a
# read-only module cache into the work dir.
(cd "$root" && go build -o "$work/provenance" ./cmd/provenance)

export HOME="$work/home"
mkdir -p "$HOME/no-hooks"
export GIT_CONFIG_GLOBAL="$HOME/.gitconfig"
export GIT_CONFIG_NOSYSTEM=1
printf '[user]\n\temail = e2e@example.com\n\tname = E2E\n[init]\n\tdefaultBranch = main\n' > "$GIT_CONFIG_GLOBAL"
# The hooks find the binary on PATH: PROVENANCE stays unset, so a repo's own
# provenance/ directory can never be mistaken for it.
mkdir -p "$work/bin"
mv "$work/provenance" "$work/bin/provenance"
unset PROVENANCE
export PATH="$work/bin:$PATH"
P="$work/bin/provenance"

repo="$work/repo"
git clone -q "$src" "$repo"
top=$(git -C "$repo" rev-parse --show-toplevel)
if [ "$(cd "$top" && pwd -P)" != "$(cd "$repo" && pwd -P)" ]; then
	echo "refusing: $repo resolves to $top" >&2
	exit 1
fi
git -C "$repo" config core.hooksPath "$repo/.git/hooks"

pass() { echo "PASS  $1"; }
fail() { echo "FAIL  $1" >&2; exit 1; }
commit() { git -C "$repo" commit -q -m "$1" >"$work/commit.log" 2>&1; }

cd "$repo"

# 1. Migrate and install the hooks.
"$P" migrate >"$work/migrate.log" || fail "migrate"
[ -f .provenance.yml ] && [ -f .provenance/versions.jsonl ] || fail "migrate wrote config and manifest"
pass "migrate"
"$P" hooks install >/dev/null || fail "hooks install"
pass "hooks install"

# 2. A record governing the migration, then the migration commit through the hooks.
id=$("$P" create --no-edit --type blueprint --title "Migrate to v2" \
	--intent "Records move to v2: versions, snapshots, manifest and .provenance.yml." \
	--scope "provenance/**" --scope ".provenance.yml" --scope ".provenance/**" \
	--scope ".linespec.yml" --scope ".linespec/**" | sed -n 's|.*/\(prov-[^/]*\)\.yml$|\1|p')
[ -n "$id" ] || fail "create"
# Drafts are freely editable: give it a constraint and a spec before it opens.
rec="provenance/$id.yml"
sed -i.bak "s|^constraints: \[\]|constraints:\n  - Migration leaves lint clean|; s|^associated_specs: \[\]|associated_specs:\n  - path: .provenance.yml\n    run_command: test -f|" "$rec" && rm -f "$rec.bak"
git -C "$repo" add -A
commit "Migrate records to v2 [$id]" || { cat "$work/commit.log" >&2; fail "tagged migration commit"; }
pass "tagged in-scope migration commit"

# 3. Untagged and out-of-scope commits are refused.
f=go.mod
echo "// e2e untagged" >> "$f"
git -C "$repo" add "$f"
if commit "No tag"; then fail "untagged commit refused"; fi
pass "untagged commit refused"
if commit "Out of scope [$id]"; then fail "out-of-scope commit refused"; fi
pass "out-of-scope commit refused"
git -C "$repo" reset -q --hard

# 4. open (auto-commit, v1 snapshot), edit -> version 2, commit [id@2].
"$P" open --record "$id" >/dev/null || fail "open"
"$P" edit "$id" --add-constraint "Hooks are provenance's" --reason "E2E revise" >/dev/null || fail "edit"
git -C "$repo" add -A
if commit "Revise [$id@1]"; then fail "stale version tag refused"; fi
pass "stale version tag refused"
commit "Revise [$id@2]" || { cat "$work/commit.log" >&2; fail "revise commit [id@2]"; }
pass "revise commit [$id@2]"

# 5. A staged edit of an existing snapshot is refused.
snap=$(ls "provenance/.versions/$id"/1-*.yml)
chmod u+w "$snap"
echo "# tampered" >> "$snap"
git -C "$repo" add "$snap"
if commit "Tamper [$id]"; then fail "snapshot edit refused"; fi
pass "snapshot edit refused"
git -C "$repo" reset -q --hard

# 6. complete runs the record's spec and auto-commits through the hooks.
"$P" complete --record "$id" >"$work/complete.log" 2>&1 || { cat "$work/complete.log" >&2; fail "complete"; }
grep -q "Running 1 associated spec(s)" "$work/complete.log" || fail "complete ran the spec"
[ "$(git -C "$repo" log -1 --format=%s)" = "Complete provenance record [$id@2]" ] || fail "complete auto-commit"
pass "complete (spec run, auto-commit [$id@2])"

# 7. Lint is clean and nothing is left behind.
"$P" lint >"$work/lint.log" || { cat "$work/lint.log" >&2; fail "lint"; }
grep -q "✗ 0 errors" "$work/lint.log" || fail "lint reports 0 errors"
[ -z "$(git -C "$repo" status --porcelain)" ] || fail "clean tree"
pass "lint 0 errors, clean tree"

echo "e2e: ok"
