#!/bin/sh
# Install the provenance CLI.
#
#   curl -fsSL https://raw.githubusercontent.com/livecodelife/provenance/main/install.sh | sh
#
# By default it downloads the prebuilt binary of the latest release for this
# platform, checks it against the release's checksums.txt and puts it in
# PROVENANCE_BIN_DIR (default $HOME/.local/bin). No sudo, and nothing else runs:
# the follow-up commands are only printed.
#
#   PROVENANCE_VERSION=2.0.0   install that release instead of the latest
#   PROVENANCE_BIN_DIR=DIR     install into DIR
#   PROVENANCE_REF=REF         build that commit, tag or branch with `go install`
#                              (needs Go; for unreleased code)
#
# Until the first release exists, or when no prebuilt binary matches this
# platform, it builds the commit baked in below with `go install` into GOBIN
# (else GOPATH/bin). Bump DEFAULT_REF to the current main tip when shipping.
set -eu

DEFAULT_REF=3773a30546addc0811cb9d547d5ee99c15c0f970
MODULE=github.com/livecodelife/provenance/cmd/provenance
RELEASES=${PROVENANCE_RELEASES_URL:-https://github.com/livecodelife/provenance/releases}

fail() {
    echo "provenance install: $*" >&2
    exit 1
}

# install_from_source REF builds REF with go install and sets bindir.
install_from_source() {
    ref=$1
    command -v go >/dev/null 2>&1 || fail "Go is required to build provenance, and 'go' was not found on PATH.
Install Go from https://go.dev/doc/install and run this script again."
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
    [ -x "$bindir/provenance" ] || fail "expected $bindir/provenance after go install, but it is not there."
}

sha256_of() {
    if command -v sha256sum >/dev/null 2>&1; then
        sha256sum "$1" | cut -d' ' -f1
    elif command -v shasum >/dev/null 2>&1; then
        shasum -a 256 "$1" | cut -d' ' -f1
    else
        return 1
    fi
}

# install_release VERSION downloads, verifies and installs a release binary
# and sets bindir. It returns 1 (after saying why) when the release or a
# binary for this platform is not available, and exits when a download fails
# its checksum.
install_release() {
    version=$1
    case $(uname -s) in
    Linux) os=linux ;;
    Darwin) os=darwin ;;
    *) echo "No prebuilt provenance for $(uname -s)." >&2; return 1 ;;
    esac
    case $(uname -m) in
    x86_64 | amd64) arch=amd64 ;;
    arm64 | aarch64) arch=arm64 ;;
    *) echo "No prebuilt provenance for $(uname -m)." >&2; return 1 ;;
    esac
    command -v curl >/dev/null 2>&1 || { echo "curl was not found on PATH." >&2; return 1; }

    if [ -z "$version" ]; then
        tag=$(curl -fsSLI -o /dev/null -w '%{url_effective}' "$RELEASES/latest" 2>/dev/null) || tag=
        version=${tag##*/v}
        case $version in
        [0-9]*.[0-9]*.[0-9]*) ;;
        *) echo "No release of provenance has been published yet." >&2; return 1 ;;
        esac
    fi
    name=provenance_${version}_${os}_${arch}.tar.gz
    base=$RELEASES/download/v$version

    work=$(mktemp -d)
    trap 'rm -rf "$work"' EXIT
    echo "Downloading provenance $version for $os/$arch"
    curl -fsSL -o "$work/$name" "$base/$name" || { echo "Could not download $base/$name." >&2; return 1; }
    curl -fsSL -o "$work/checksums.txt" "$base/checksums.txt" || { echo "Could not download $base/checksums.txt." >&2; return 1; }
    want=$(grep "  $name\$" "$work/checksums.txt" | cut -d' ' -f1)
    [ -n "$want" ] || fail "checksums.txt has no entry for $name."
    got=$(sha256_of "$work/$name") || fail "sha256sum or shasum is required to verify the download."
    [ "$want" = "$got" ] || fail "checksum mismatch for $name (expected $want, got $got); nothing was installed."
    tar -xzf "$work/$name" -C "$work" provenance || fail "could not unpack $name."

    bindir=${PROVENANCE_BIN_DIR:-$HOME/.local/bin}
    mkdir -p "$bindir"
    cp "$work/provenance" "$bindir/provenance.new"
    chmod 755 "$bindir/provenance.new"
    mv -f "$bindir/provenance.new" "$bindir/provenance"
}

bindir=
if [ -n "${PROVENANCE_REF:-}" ]; then
    install_from_source "$PROVENANCE_REF"
elif install_release "${PROVENANCE_VERSION:-}"; then
    :
elif [ -n "${PROVENANCE_VERSION:-}" ]; then
    fail "provenance $PROVENANCE_VERSION is not available."
else
    echo "Falling back to building the pinned commit from source." >&2
    install_from_source "$DEFAULT_REF"
fi
echo "Installed $bindir/provenance"

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
