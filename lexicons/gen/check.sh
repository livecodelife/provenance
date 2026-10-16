#!/bin/sh
# Runs the round-trip tests of the generated types (needs the network once for
# npm ci and the Go modules; both are cached afterwards).
set -eu
here=$(cd "$(dirname "$0")" && pwd)
(cd "$here/ts" && npm ci --silent && npm test)
(cd "$here/go" && GOFLAGS=-mod=mod go test ./...)
