package assets

// sedum:anchor:imports
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-d71d0066","gen":"sha256:cae7a2d3a080d43f1ba12537ce93365db84c39195d99ec904268e9c15daab6a9","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"assets/embed.go","path":"embed"},"kwargs":"sha256:1b52e680b0bb47adea09e1e0837effdac08b5ee9cf9cef85c514b94b514517ac","body":"sha256:95d27c73c53a26e36b5a419e1a47a07501240af4cea018845b946ccac553aa1a"}
import "embed"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setVar {"v":2,"tier":"owned","record":"prov-2026-d71d0066","gen":"sha256:cae7a2d3a080d43f1ba12537ce93365db84c39195d99ec904268e9c15daab6a9","pkg":"sha256:c7b9d9ac1655e8211b8b2ca8c49cf34b8d510d8d823d4f5eb60ec7e1018aa2eb","id":{"file":"assets/embed.go","name":"Files"},"kwargs":"sha256:9b55587974cd87b9311104ef715a6655e48899c23c3274c43264f6df344c75e5","body":"sha256:7bd1c5aa2b8aa6d2ab714d3af675739e8c69cdea912a726940f621bad957c7c6"}
// Files is the embedded Claude Code skill and plugin tree that
// provenance install-skills and provenance install-plugin copy into a repo.
// The all: prefix includes the dot-prefixed .claude-plugin directory.
//
//go:embed all:skills all:plugins
var Files embed.FS

// /sedum:setVar
