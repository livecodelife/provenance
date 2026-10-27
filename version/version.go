package version

// sedum:anchor:imports

// sedum:anchor:decls
// sedum:setConst {"v":2,"tier":"owned","record":"prov-2026-f5d9c614","gen":"sha256:6befbee5939305e145b494bf03fc8e6d7e5ffbba6a86328642ca5f94e9ebc4cc","pkg":"sha256:51be1c18998d63cc558c253e211a6bd188d96cb251dad68c6ea88e03c2d5cc4c","id":{"file":"version/version.go","name":"Version"},"kwargs":"sha256:3ff6f7b472ea652e09911779b4b2b1f84a03c3055c0b6f63bbcc04aa825c265f","body":"sha256:aefaa15c3c56c192511f16a9b66c6be1c8a2a07fab605ef3bcba52b5a13fd8dc"}
// Version is the release version of the provenance CLI, the bare semver that
// `provenance version` prints. It is a constant, not a build-time stamp, so
// go build and go run report it too; .github/workflows/release.yml refuses a
// tag that differs from it.
const Version = "2.1.0"

// /sedum:setConst
