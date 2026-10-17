#!/bin/sh
# M6 joint check: CLI login, publish, UI edit, CLI import, logout, against one
# backend. See docs/M6.md. Local variant: provisio-backend's dev server on :3100
# (test/dev/local-server.ts) and the UI checkout in M6_UI_REPO.
set -eu

here=$(cd "$(dirname "$0")/.." && pwd)
: "${PROVISIO_AUTH_URL:=http://127.0.0.1:3100}"
: "${PROVISIO_API_URL:=$PROVISIO_AUTH_URL/v1}"
export PROVISIO_AUTH_URL PROVISIO_API_URL
M6_REAL_HOME=$HOME
export M6_REAL_HOME

case "$PROVISIO_AUTH_URL" in
http://127.0.0.1:*|http://localhost:*)
	: "${M6_APPROVE_URL:=$PROVISIO_AUTH_URL/dev/approve}" ;;
esac
[ -z "${M6_APPROVE_URL:-}" ] || export M6_APPROVE_URL

vite_pid=
# vite starts node children and setsid may re-fork, so stop whatever listens on the UI port.
cleanup() {
	[ -z "$vite_pid" ] || { kill "$vite_pid" 2>/dev/null; fuser -k -TERM 5173/tcp >/dev/null 2>&1; } || true
}
trap cleanup EXIT INT TERM

if [ -z "${M6_UI_CMD:-}" ] && [ -n "${M6_UI_REPO:-}" ]; then
	ui=${M6_UI_URL:-http://127.0.0.1:5173}
	if curl -fsS "$ui" >/dev/null 2>&1; then
		echo "something already listens on $ui; stop it first (the UI must be started by this script with the right environment)" >&2
		exit 1
	fi
	# Own process group, so the cleanup stops vite and the node it starts.
	(cd "$M6_UI_REPO" && VITE_AUTH_MODE=provisio VITE_PROVISIO_AUTH_URL="$PROVISIO_AUTH_URL" VITE_PROVISIO_API_URL="$PROVISIO_AUTH_URL" \
		exec setsid npx vite --host 127.0.0.1 --port 5173 --strictPort) >"${TMPDIR:-/tmp}/m6-vite.log" 2>&1 &
	vite_pid=$!
	n=0
	until curl -fsS "$ui" >/dev/null 2>&1; do
		n=$((n + 1))
		[ "$n" -lt 60 ] || { echo "the UI dev server did not start; see ${TMPDIR:-/tmp}/m6-vite.log" >&2; exit 1; }
		sleep 1
	done
	M6_UI_CMD="cd '$M6_UI_REPO' && xvfb-run -a npx cypress run --config-file cypress.provisio.config.ts --spec cypress/e2e/provisio-m6-edit.cy.ts"
	export M6_UI_CMD
fi
if [ -z "${M6_UI_CMD:-}" ] && [ -z "${M6_UI_MANUAL:-}" ]; then
	echo "M6_UI_CMD, M6_UI_REPO and M6_UI_MANUAL are unset: only a loopback backend (or M6_STAND_IN=1) lets the API stand-in do the UI step (CLI half and wire contract only)." >&2
fi
export M6_UI_MANUAL M6_STAND_IN

cd "$here"
go test -tags m6 -count=1 -run TestM6Joint -v ./cmd/provenance/
