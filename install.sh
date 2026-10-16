#!/bin/sh
# Install the provenance CLI by building it from source with `go install`.
#
#   curl -fsSL https://raw.githubusercontent.com/livecodelife/provenance/main/install.sh | sh
#
# PROVENANCE_REF overrides the commit, tag or branch to build (for example
# PROVENANCE_REF=main). By default it builds the commit baked in below; bump
# DEFAULT_REF to the current main tip each time we ship.
# It installs into GOBIN (else GOPATH/bin), needs no sudo, and runs nothing
# else: the follow-up commands are only printed.
set -eu

DEFAULT_REF=f3088a41a2846302d8d8c9f31fc480a92ec92948
MODULE=github.com/livecodelife/provenance/cmd/provenance

ref=${PROVENANCE_REF:-$DEFAULT_REF}

if ! command -v go >/dev/null 2>&1; then
    echo "provenance install: Go is required to build provenance, and 'go' was not found on PATH." >&2
    echo "Install Go from https://go.dev/doc/install and run this script again." >&2
    exit 1
fi

echo "Building provenance ($ref) with $(go version)"
if ! go install "$MODULE@$ref"; then
    echo "provenance install: 'go install $MODULE@$ref' failed." >&2
    echo "Check the ref and your network; the Go module proxy can lag a few minutes behind a fresh push." >&2
    exit 1
fi

bindir=$(go env GOBIN)
if [ -z "$bindir" ]; then
    bindir=$(go env GOPATH)/bin
fi
bin=$bindir/provenance
if [ ! -x "$bin" ]; then
    echo "provenance install: expected $bin after go install, but it is not there." >&2
    exit 1
fi
echo "Installed $bin"

case ":$PATH:" in
*":$bindir:"*) ;;
*)
    echo
    echo "$bindir is not on your PATH. Add it, for example:"
    echo "    export PATH=\"$bindir:\$PATH\""
    ;;
esac

cat <<'MSG'

Nothing else was changed. Next steps, when you want them:

  New repository:
    provenance init                  scaffold .provenance.yml and the records directory
  Repository with v1 (linespec) records:
    provenance migrate --plan        show what migrate would do; writes nothing
    provenance migrate               convert the records to v2
  Either way:
    provenance hooks install         install the git hooks (existing ones are kept as .bak)
MSG
