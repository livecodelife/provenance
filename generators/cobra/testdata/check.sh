#!/bin/sh
# Proof for prov-2026-575d0238: the cobra package loads cleanly, grows the
# golden demo CLI byte for byte from example.json, regrows it without change,
# and the grown code is gofmt-clean, vets, builds and answers --help.
# SEDUM names the sedum binary. Without it, SEDUM_SRC names a sedum checkout
# to build one from; with neither, sedum on PATH is used. The package needs a
# sedum with action identity and kwarg defaults (grow-prep).
set -eu

here=$(cd "$(dirname "$0")" && pwd)
generators=$(cd "$here/../.." && pwd)
work=$(mktemp -d)
trap 'rm -rf "$work"' EXIT

if [ -z "${SEDUM:-}" ] && [ -n "${SEDUM_SRC:-}" ]; then
	(cd "$SEDUM_SRC" && go build -o "$work/sedum" ./cmd/sedum)
	SEDUM="$work/sedum"
fi
SEDUM=${SEDUM:-sedum}

"$SEDUM" validate --generators "$generators" --package cobra --strict

grow() {
	"$SEDUM" grow --generators "$generators" --records "$here/records" \
		--execute "$here/example.json" --output "$1" --log "$work/run.log" >/dev/null
}

mkdir "$work/out"
grow "$work/out"
diff -r "$here/golden" "$work/out"

# A second replay over grown output rewrites nothing.
grow "$work/out"
diff -r "$here/golden" "$work/out"

unformatted=$(gofmt -l "$work/out")
if [ -n "$unformatted" ]; then
	echo "not gofmt-clean: $unformatted" >&2
	exit 1
fi

mod="$work/mod"
mkdir "$mod"
cp "$work/out/cmd/demo/"*.go "$here/stubs/run_stubs.go" "$mod/"
cat > "$mod/go.mod" <<MOD
module demo

go 1.25

require github.com/spf13/cobra v1.10.2
MOD
export GOFLAGS=-mod=mod GOPROXY=off GOWORK=off
(cd "$mod" && go vet . && go build -o demo .)

for args in "" "greet" "echo" "list" "hooks" "hooks install"; do
	# shellcheck disable=SC2086
	"$mod/demo" $args --help >/dev/null
done
"$mod/demo" greet --help | grep -q -- '-g, --greeting string'
"$mod/demo" hooks install --help | grep -q -- '-c, --config string'

echo "cobra package: ok"
