#!/bin/sh
# Proof for prov-2026-70c3163b: both packages load cleanly, cobra's own check
# passes, and on a .go path both claim, cobra keeps its wiring files while go
# takes the rest. SEDUM / SEDUM_SRC as in cobra/testdata/check.sh.
set -eu

here=$(cd "$(dirname "$0")" && pwd)
work=$(mktemp -d)
trap 'rm -rf "$work"' EXIT

if [ -z "${SEDUM:-}" ] && [ -n "${SEDUM_SRC:-}" ]; then
	(cd "$SEDUM_SRC" && go build -o "$work/sedum" ./cmd/sedum)
	SEDUM="$work/sedum"
fi
SEDUM=${SEDUM:-sedum}
export SEDUM

"$SEDUM" validate --generators "$here" --strict
"$here/cobra/testdata/check.sh"

resolved=$("$SEDUM" resolve --generators "$here" --records "$here/cobra/testdata/resolve")
# resolve prints each path on its own line, then "package <name>" under it.
pkgs=$(printf '%s\n' "$resolved" | awk '$1 ~ /^cmd\// {p=$1} $1 == "package" {print p, $2}')
want="cmd/demo/main.go cobra
cmd/demo/x_cmd.go cobra
cmd/demo/x_run.go go"
if [ "$pkgs" != "$want" ]; then
	printf 'resolution differs; want:\n%s\ngot:\n%s\n' "$want" "$pkgs" >&2
	exit 1
fi
echo "generators: ok"
