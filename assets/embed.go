package assets

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-d71d0066","gen":"sha256:cae7a2d3a080d43f1ba12537ce93365db84c39195d99ec904268e9c15daab6a9","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"assets/embed.go","path":"embed"}}
import "embed"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setVar {"tier":"owned","record":"prov-2026-d71d0066","gen":"sha256:cae7a2d3a080d43f1ba12537ce93365db84c39195d99ec904268e9c15daab6a9","pkg":"sha256:c7b9d9ac1655e8211b8b2ca8c49cf34b8d510d8d823d4f5eb60ec7e1018aa2eb","kwargs":{"code":"// Files is the embedded Claude Code skill and plugin tree that\n// provenance install-skills and provenance install-plugin copy into a repo.\n// The all: prefix includes the dot-prefixed .claude-plugin directory.\n//\n//go:embed all:skills all:plugins\nvar Files embed.FS","file":"assets/embed.go","name":"Files"}}
// Files is the embedded Claude Code skill and plugin tree that
// provenance install-skills and provenance install-plugin copy into a repo.
// The all: prefix includes the dot-prefixed .claude-plugin directory.
//
//go:embed all:skills all:plugins
var Files embed.FS

// /sedum:setVar
