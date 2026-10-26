#!/bin/sh
# Regenerates lexicons/gen/ts and lexicons/gen/go from lexicons/dev/**. Needs the
# network the first time (npm packages, Go modules). Versions are pinned:
#   @atproto/lex-cli 0.10.14           (gen/ts/package.json, package-lock.json)
#   indigo cmd/lexgen @ d8478c47f0ce6ba53f214e1294c219732795e03b
set -eu
here=$(cd "$(dirname "$0")" && pwd)
lex=$here/..
INDIGO=d8478c47f0ce6ba53f214e1294c219732795e03b

# TypeScript: types and validators. The client class lex-cli also writes needs
# the com.atproto.repo lexicons and is not used, so it is dropped.
cd "$here/ts"
npm ci
npx lex gen-api --yes src "$lex"/dev/provisio/provenance/*.json "$lex/dev/provisio/project.json" "$lex/vendor/com/atproto/repo/strongRef.json"
rm -f src/index.ts

# Go: lexgen, then cbor-gen. The generated types register themselves with
# lexutil, which needs the CBOR methods cbor-gen writes next, so registration
# is held back for that one step.
tmp=$(mktemp -d)
trap 'rm -rf "$tmp"' EXIT
mkdir -p "$tmp/lex" "$tmp/ext"
cp -R "$lex/dev" "$tmp/lex/dev"
cp "$lex/vendor/com/atproto/repo/strongRef.json" "$tmp/ext/"
cat > "$tmp/build.json" <<JSON
[
  {"package": "provenancelex", "prefix": "dev.provisio.provenance", "outdir": "$tmp/out",
   "import": "github.com/livecodelife/provenance/lexicons/gen/go"},
  {"package": "provenancelex", "prefix": "dev.provisio", "outdir": "$tmp/outproject",
   "import": "github.com/livecodelife/provenance/lexicons/gen/go"},
  {"package": "atproto", "prefix": "com.atproto", "outdir": "$tmp/extout",
   "import": "github.com/bluesky-social/indigo/api/atproto"}
]
JSON
cd "$tmp"
GOFLAGS=-mod=mod go run "github.com/bluesky-social/indigo/cmd/lexgen@$INDIGO" --build-file build.json --external-lexicons ext lex >/dev/null
cd "$here/go"
rm -f provenancedefs.go provenancerecord.go provenanceversion.go provisioproject.go cbor_gen.go
cp "$tmp"/out/*.go .
# The dev.provisio prefix also regenerates the provenance lexicons under other
# names in outproject; only the project type is kept.
cp "$tmp"/outproject/provisioproject.go .
for f in provenancerecord.go provenanceversion.go provisioproject.go; do cp "$f" "$tmp/$f"; done
python3 -I - <<'PY'
import re
for f in ["provenancerecord.go", "provenanceversion.go", "provisioproject.go"]:
    s = open(f).read()
    s = re.sub(r"func init\(\) \{\n\tlexutil\.RegisterType\([^\n]*\n\}\n\n", "", s)
    s = s.replace('\tlexutil "github.com/bluesky-social/indigo/lex/util"\n', "")
    open(f, "w").write(s)
PY
GOFLAGS=-mod=mod go run ./gen
for f in provenancerecord.go provenanceversion.go provisioproject.go; do cp "$tmp/$f" "$f"; done
gofmt -w provenance*.go provisioproject.go cbor_gen.go
