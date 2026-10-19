package version

// sedum:anchor:imports

// sedum:anchor:decls
// sedum:setConst {"tier":"owned","record":"prov-2026-f5d9c614","gen":"sha256:64ae3642a3fed543e6a162cc532b31495fcc915f467bda1a7045e56f3b873756","pkg":"sha256:51be1c18998d63cc558c253e211a6bd188d96cb251dad68c6ea88e03c2d5cc4c","kwargs":{"code":"// Version is the release version of the provenance CLI, the bare semver that\n// `provenance version` prints. It is a constant, not a build-time stamp, so\n// go build and go run report it too; .github/workflows/release.yml refuses a\n// tag that differs from it.\nconst Version = \"2.0.0\"","file":"version/version.go","name":"Version"}}
// Version is the release version of the provenance CLI, the bare semver that
// `provenance version` prints. It is a constant, not a build-time stamp, so
// go build and go run report it too; .github/workflows/release.yml refuses a
// tag that differs from it.
const Version = "2.0.0"

// /sedum:setConst
