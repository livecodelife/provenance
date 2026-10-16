package main

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-03078504","gen":"sha256:aaa940c9ff8508cc84857988bd8cb585939fdb57f5223d6e874b86715d0eb765","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/sarif_run.go","path":"crypto/sha256"}}
import "crypto/sha256"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-03078504","gen":"sha256:aaa940c9ff8508cc84857988bd8cb585939fdb57f5223d6e874b86715d0eb765","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/sarif_run.go","path":"encoding/hex"}}
import "encoding/hex"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-03078504","gen":"sha256:aaa940c9ff8508cc84857988bd8cb585939fdb57f5223d6e874b86715d0eb765","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/sarif_run.go","path":"encoding/json"}}
import "encoding/json"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-03078504","gen":"sha256:aaa940c9ff8508cc84857988bd8cb585939fdb57f5223d6e874b86715d0eb765","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/sarif_run.go","path":"io"}}
import "io"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-03078504","gen":"sha256:aaa940c9ff8508cc84857988bd8cb585939fdb57f5223d6e874b86715d0eb765","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/sarif_run.go","path":"os"}}
import "os"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-03078504","gen":"sha256:aaa940c9ff8508cc84857988bd8cb585939fdb57f5223d6e874b86715d0eb765","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/sarif_run.go","path":"path/filepath"}}
import "path/filepath"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-03078504","gen":"sha256:aaa940c9ff8508cc84857988bd8cb585939fdb57f5223d6e874b86715d0eb765","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/sarif_run.go","path":"runtime/debug"}}
import "runtime/debug"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-03078504","gen":"sha256:aaa940c9ff8508cc84857988bd8cb585939fdb57f5223d6e874b86715d0eb765","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/sarif_run.go","path":"sort"}}
import "sort"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-03078504","gen":"sha256:aaa940c9ff8508cc84857988bd8cb585939fdb57f5223d6e874b86715d0eb765","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/sarif_run.go","path":"strings"}}
import "strings"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-03078504","gen":"sha256:aaa940c9ff8508cc84857988bd8cb585939fdb57f5223d6e874b86715d0eb765","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/sarif_run.go","path":"github.com/livecodelife/provenance/lint"}}
import "github.com/livecodelife/provenance/lint"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setType {"tier":"owned","record":"prov-2026-03078504","gen":"sha256:aaa940c9ff8508cc84857988bd8cb585939fdb57f5223d6e874b86715d0eb765","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type sarifMessage struct {\n\tText string `json:\"text\"`\n}","file":"cmd/provenance/sarif_run.go","name":"sarifMessage"}}
type sarifMessage struct {
	Text string `json:"text"`
}

// /sedum:setType
// sedum:setType {"tier":"owned","record":"prov-2026-03078504","gen":"sha256:aaa940c9ff8508cc84857988bd8cb585939fdb57f5223d6e874b86715d0eb765","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type sarifRuleConfig struct {\n\tLevel string `json:\"level\"`\n}","file":"cmd/provenance/sarif_run.go","name":"sarifRuleConfig"}}
type sarifRuleConfig struct {
	Level string `json:"level"`
}

// /sedum:setType
// sedum:setType {"tier":"owned","record":"prov-2026-03078504","gen":"sha256:aaa940c9ff8508cc84857988bd8cb585939fdb57f5223d6e874b86715d0eb765","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type sarifRule struct {\n\tID                   string          `json:\"id\"`\n\tName                 string          `json:\"name\"`\n\tShortDescription     sarifMessage    `json:\"shortDescription\"`\n\tDefaultConfiguration sarifRuleConfig `json:\"defaultConfiguration\"`\n}","file":"cmd/provenance/sarif_run.go","name":"sarifRule"}}
type sarifRule struct {
	ID                   string          `json:"id"`
	Name                 string          `json:"name"`
	ShortDescription     sarifMessage    `json:"shortDescription"`
	DefaultConfiguration sarifRuleConfig `json:"defaultConfiguration"`
}

// /sedum:setType
// sedum:setType {"tier":"owned","record":"prov-2026-03078504","gen":"sha256:aaa940c9ff8508cc84857988bd8cb585939fdb57f5223d6e874b86715d0eb765","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type sarifArtifactLocation struct {\n\tURI       string `json:\"uri\"`\n\tURIBaseID string `json:\"uriBaseId,omitempty\"`\n}","file":"cmd/provenance/sarif_run.go","name":"sarifArtifactLocation"}}
type sarifArtifactLocation struct {
	URI       string `json:"uri"`
	URIBaseID string `json:"uriBaseId,omitempty"`
}

// /sedum:setType
// sedum:setType {"tier":"owned","record":"prov-2026-03078504","gen":"sha256:aaa940c9ff8508cc84857988bd8cb585939fdb57f5223d6e874b86715d0eb765","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type sarifRegion struct {\n\tStartLine   int `json:\"startLine\"`\n\tStartColumn int `json:\"startColumn\"`\n}","file":"cmd/provenance/sarif_run.go","name":"sarifRegion"}}
type sarifRegion struct {
	StartLine   int `json:"startLine"`
	StartColumn int `json:"startColumn"`
}

// /sedum:setType
// sedum:setType {"tier":"owned","record":"prov-2026-03078504","gen":"sha256:aaa940c9ff8508cc84857988bd8cb585939fdb57f5223d6e874b86715d0eb765","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type sarifPhysicalLocation struct {\n\tArtifactLocation sarifArtifactLocation `json:\"artifactLocation\"`\n\tRegion           sarifRegion           `json:\"region\"`\n}","file":"cmd/provenance/sarif_run.go","name":"sarifPhysicalLocation"}}
type sarifPhysicalLocation struct {
	ArtifactLocation sarifArtifactLocation `json:"artifactLocation"`
	Region           sarifRegion           `json:"region"`
}

// /sedum:setType
// sedum:setType {"tier":"owned","record":"prov-2026-03078504","gen":"sha256:aaa940c9ff8508cc84857988bd8cb585939fdb57f5223d6e874b86715d0eb765","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type sarifLocation struct {\n\tPhysicalLocation sarifPhysicalLocation `json:\"physicalLocation\"`\n}","file":"cmd/provenance/sarif_run.go","name":"sarifLocation"}}
type sarifLocation struct {
	PhysicalLocation sarifPhysicalLocation `json:"physicalLocation"`
}

// /sedum:setType
// sedum:setType {"tier":"owned","record":"prov-2026-03078504","gen":"sha256:aaa940c9ff8508cc84857988bd8cb585939fdb57f5223d6e874b86715d0eb765","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type sarifResult struct {\n\tRuleID     string            `json:\"ruleId\"`\n\tRuleIndex  int               `json:\"ruleIndex\"`\n\tLevel      string            `json:\"level\"`\n\tMessage    sarifMessage      `json:\"message\"`\n\tLocations  []sarifLocation   `json:\"locations,omitempty\"`\n\tProperties map[string]string `json:\"properties,omitempty\"`\n}","file":"cmd/provenance/sarif_run.go","name":"sarifResult"}}
type sarifResult struct {
	RuleID     string            `json:"ruleId"`
	RuleIndex  int               `json:"ruleIndex"`
	Level      string            `json:"level"`
	Message    sarifMessage      `json:"message"`
	Locations  []sarifLocation   `json:"locations,omitempty"`
	Properties map[string]string `json:"properties,omitempty"`
}

// /sedum:setType
// sedum:setType {"tier":"owned","record":"prov-2026-03078504","gen":"sha256:aaa940c9ff8508cc84857988bd8cb585939fdb57f5223d6e874b86715d0eb765","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type sarifArtifact struct {\n\tLocation sarifArtifactLocation `json:\"location\"`\n\tMimeType string                `json:\"mimeType,omitempty\"`\n\tHashes   map[string]string     `json:\"hashes,omitempty\"`\n}","file":"cmd/provenance/sarif_run.go","name":"sarifArtifact"}}
type sarifArtifact struct {
	Location sarifArtifactLocation `json:"location"`
	MimeType string                `json:"mimeType,omitempty"`
	Hashes   map[string]string     `json:"hashes,omitempty"`
}

// /sedum:setType
// sedum:setType {"tier":"owned","record":"prov-2026-03078504","gen":"sha256:aaa940c9ff8508cc84857988bd8cb585939fdb57f5223d6e874b86715d0eb765","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type sarifDriver struct {\n\tName           string      `json:\"name\"`\n\tVersion        string      `json:\"version\"`\n\tInformationURI string      `json:\"informationUri\"`\n\tRules          []sarifRule `json:\"rules\"`\n}","file":"cmd/provenance/sarif_run.go","name":"sarifDriver"}}
type sarifDriver struct {
	Name           string      `json:"name"`
	Version        string      `json:"version"`
	InformationURI string      `json:"informationUri"`
	Rules          []sarifRule `json:"rules"`
}

// /sedum:setType
// sedum:setType {"tier":"owned","record":"prov-2026-03078504","gen":"sha256:aaa940c9ff8508cc84857988bd8cb585939fdb57f5223d6e874b86715d0eb765","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type sarifTool struct {\n\tDriver sarifDriver `json:\"driver\"`\n}","file":"cmd/provenance/sarif_run.go","name":"sarifTool"}}
type sarifTool struct {
	Driver sarifDriver `json:"driver"`
}

// /sedum:setType
// sedum:setType {"tier":"owned","record":"prov-2026-03078504","gen":"sha256:aaa940c9ff8508cc84857988bd8cb585939fdb57f5223d6e874b86715d0eb765","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type sarifRun struct {\n\tTool      sarifTool       `json:\"tool\"`\n\tResults   []sarifResult   `json:\"results\"`\n\tArtifacts []sarifArtifact `json:\"artifacts\"`\n}","file":"cmd/provenance/sarif_run.go","name":"sarifRun"}}
type sarifRun struct {
	Tool      sarifTool       `json:"tool"`
	Results   []sarifResult   `json:"results"`
	Artifacts []sarifArtifact `json:"artifacts"`
}

// /sedum:setType
// sedum:setType {"tier":"owned","record":"prov-2026-03078504","gen":"sha256:aaa940c9ff8508cc84857988bd8cb585939fdb57f5223d6e874b86715d0eb765","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type sarifDoc struct {\n\tSchema  string     `json:\"$schema\"`\n\tVersion string     `json:\"version\"`\n\tRuns    []sarifRun `json:\"runs\"`\n}","file":"cmd/provenance/sarif_run.go","name":"sarifDoc"}}
type sarifDoc struct {
	Schema  string     `json:"$schema"`
	Version string     `json:"version"`
	Runs    []sarifRun `json:"runs"`
}

// /sedum:setType
// sedum:setType {"tier":"owned","record":"prov-2026-03078504","gen":"sha256:aaa940c9ff8508cc84857988bd8cb585939fdb57f5223d6e874b86715d0eb765","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type sarifFinding struct {\n\tRuleID   string\n\tLevel    string\n\tMessage  string\n\tFile     string\n\tRecordID string\n\tField    string\n}","file":"cmd/provenance/sarif_run.go","name":"sarifFinding"}}
type sarifFinding struct {
	RuleID   string
	Level    string
	Message  string
	File     string
	RecordID string
	Field    string
}

// /sedum:setType
// sedum:setVar {"tier":"owned","record":"prov-2026-03078504","gen":"sha256:aaa940c9ff8508cc84857988bd8cb585939fdb57f5223d6e874b86715d0eb765","pkg":"sha256:c7b9d9ac1655e8211b8b2ca8c49cf34b8d510d8d823d4f5eb60ec7e1018aa2eb","kwargs":{"code":"var sarifCatalog = []sarifRule{\n\t{\"PROV002\", \"RequiredField\", sarifMessage{\"A required record field is missing\"}, sarifRuleConfig{\"error\"}},\n\t{\"PROV003\", \"Status\", sarifMessage{\"The record status is not valid\"}, sarifRuleConfig{\"error\"}},\n\t{\"PROV004\", \"ID\", sarifMessage{\"The record id is not in the prov-YYYY-XXXXXXXX form\"}, sarifRuleConfig{\"error\"}},\n\t{\"PROV005\", \"Date\", sarifMessage{\"A date field is not in the expected format\"}, sarifRuleConfig{\"error\"}},\n\t{\"PROV008\", \"Related\", sarifMessage{\"related names a record that does not exist\"}, sarifRuleConfig{\"warning\"}},\n\t{\"PROV009\", \"ScopeConflict\", sarifMessage{\"A pattern is in both affected_scope and forbidden_scope\"}, sarifRuleConfig{\"error\"}},\n\t{\"PROV010\", \"SpecsMissing\", sarifMessage{\"An open record has no associated specs\"}, sarifRuleConfig{\"error\"}},\n\t{\"PROV011\", \"ConstraintsNoSpecs\", sarifMessage{\"A record has constraints but no associated specs\"}, sarifRuleConfig{\"warning\"}},\n\t{\"PROV012\", \"IntentNoConstraints\", sarifMessage{\"A record has an intent but no constraints\"}, sarifRuleConfig{\"warning\"}},\n\t{\"PROV013\", \"TitleLength\", sarifMessage{\"The title is too long\"}, sarifRuleConfig{\"warning\"}},\n\t{\"PROV015\", \"ScopeOverlap\", sarifMessage{\"Open records share scope\"}, sarifRuleConfig{\"warning\"}},\n\t{\"PROV016\", \"DeadRecord\", sarifMessage{\"Every file a record governs has been deleted\"}, sarifRuleConfig{\"warning\"}},\n\t{\"PROV018\", \"SpecFile\", sarifMessage{\"An associated spec file does not exist\"}, sarifRuleConfig{\"error\"}},\n\t{\"PROV019\", \"InvalidPattern\", sarifMessage{\"A scope pattern is not a valid pattern\"}, sarifRuleConfig{\"error\"}},\n\t{\"PROV021\", \"ImplementsType\", sarifMessage{\"implements points at the wrong tier\"}, sarifRuleConfig{\"error\"}},\n\t{\"PROV022\", \"ImplementsUnknown\", sarifMessage{\"implements names a record that cannot be found\"}, sarifRuleConfig{\"error\"}},\n\t{\"PROV023\", \"ImprintContainment\", sarifMessage{\"An imprint's scope exceeds its blueprint's scope\"}, sarifRuleConfig{\"error\"}},\n\t{\"PROV024\", \"UnbumpedChange\", sarifMessage{\"Content changed without a version bump\"}, sarifRuleConfig{\"error\"}},\n\t{\"PROV025\", \"BumpWithoutReason\", sarifMessage{\"A version was written without a reason\"}, sarifRuleConfig{\"error\"}},\n\t{\"PROV026\", \"VersionSnapshots\", sarifMessage{\"The version manifest and the snapshots disagree\"}, sarifRuleConfig{\"error\"}},\n\t{\"PROV027\", \"ImplementsPin\", sarifMessage{\"The implements_version pin is not valid\"}, sarifRuleConfig{\"error\"}},\n\t{\"PROV028\", \"NotMigrated\", sarifMessage{\"The record is still on the v1 schema\"}, sarifRuleConfig{\"warning\"}},\n\t{\"PROV029\", \"ScopePath\", sarifMessage{\"A scope path does not exist\"}, sarifRuleConfig{\"warning\"}},\n\t{\"PROV030\", \"Type\", sarifMessage{\"The record type is missing or not valid\"}, sarifRuleConfig{\"note\"}},\n\t{\"PROV031\", \"NotApplicable\", sarifMessage{\"A field is not applicable to this tier\"}, sarifRuleConfig{\"warning\"}},\n\t{\"PROV032\", \"SealedAtSHA\", sarifMessage{\"sealed_at_sha does not match the record status\"}, sarifRuleConfig{\"warning\"}},\n\t{\"PROV033\", \"LockedScope\", sarifMessage{\"An open record's scope conflicts with a locked record\"}, sarifRuleConfig{\"error\"}},\n\t{\"SCOPE001\", \"ScopeViolation\", sarifMessage{\"A changed file is outside, or forbidden by, the tagged record's scope\"}, sarifRuleConfig{\"error\"}},\n\t{\"TAG001\", \"CommitTag\", sarifMessage{\"The commit tag is missing or names the wrong record version\"}, sarifRuleConfig{\"error\"}},\n\t{\"STATE001\", \"RecordState\", sarifMessage{\"A snapshot was rewritten, a status transition is not allowed, or an implemented record changed without a revise\"}, sarifRuleConfig{\"error\"}},\n\t{\"EVID001\", \"EvidenceStale\", sarifMessage{\"An implemented record's evidence is stale\"}, sarifRuleConfig{\"error\"}},\n\t{\"EVID002\", \"EvidenceFailing\", sarifMessage{\"An implemented record's latest evidence is a failing run\"}, sarifRuleConfig{\"error\"}},\n}","file":"cmd/provenance/sarif_run.go","name":"sarifCatalog"}}
var sarifCatalog = []sarifRule{
	{"PROV002", "RequiredField", sarifMessage{"A required record field is missing"}, sarifRuleConfig{"error"}},
	{"PROV003", "Status", sarifMessage{"The record status is not valid"}, sarifRuleConfig{"error"}},
	{"PROV004", "ID", sarifMessage{"The record id is not in the prov-YYYY-XXXXXXXX form"}, sarifRuleConfig{"error"}},
	{"PROV005", "Date", sarifMessage{"A date field is not in the expected format"}, sarifRuleConfig{"error"}},
	{"PROV008", "Related", sarifMessage{"related names a record that does not exist"}, sarifRuleConfig{"warning"}},
	{"PROV009", "ScopeConflict", sarifMessage{"A pattern is in both affected_scope and forbidden_scope"}, sarifRuleConfig{"error"}},
	{"PROV010", "SpecsMissing", sarifMessage{"An open record has no associated specs"}, sarifRuleConfig{"error"}},
	{"PROV011", "ConstraintsNoSpecs", sarifMessage{"A record has constraints but no associated specs"}, sarifRuleConfig{"warning"}},
	{"PROV012", "IntentNoConstraints", sarifMessage{"A record has an intent but no constraints"}, sarifRuleConfig{"warning"}},
	{"PROV013", "TitleLength", sarifMessage{"The title is too long"}, sarifRuleConfig{"warning"}},
	{"PROV015", "ScopeOverlap", sarifMessage{"Open records share scope"}, sarifRuleConfig{"warning"}},
	{"PROV016", "DeadRecord", sarifMessage{"Every file a record governs has been deleted"}, sarifRuleConfig{"warning"}},
	{"PROV018", "SpecFile", sarifMessage{"An associated spec file does not exist"}, sarifRuleConfig{"error"}},
	{"PROV019", "InvalidPattern", sarifMessage{"A scope pattern is not a valid pattern"}, sarifRuleConfig{"error"}},
	{"PROV021", "ImplementsType", sarifMessage{"implements points at the wrong tier"}, sarifRuleConfig{"error"}},
	{"PROV022", "ImplementsUnknown", sarifMessage{"implements names a record that cannot be found"}, sarifRuleConfig{"error"}},
	{"PROV023", "ImprintContainment", sarifMessage{"An imprint's scope exceeds its blueprint's scope"}, sarifRuleConfig{"error"}},
	{"PROV024", "UnbumpedChange", sarifMessage{"Content changed without a version bump"}, sarifRuleConfig{"error"}},
	{"PROV025", "BumpWithoutReason", sarifMessage{"A version was written without a reason"}, sarifRuleConfig{"error"}},
	{"PROV026", "VersionSnapshots", sarifMessage{"The version manifest and the snapshots disagree"}, sarifRuleConfig{"error"}},
	{"PROV027", "ImplementsPin", sarifMessage{"The implements_version pin is not valid"}, sarifRuleConfig{"error"}},
	{"PROV028", "NotMigrated", sarifMessage{"The record is still on the v1 schema"}, sarifRuleConfig{"warning"}},
	{"PROV029", "ScopePath", sarifMessage{"A scope path does not exist"}, sarifRuleConfig{"warning"}},
	{"PROV030", "Type", sarifMessage{"The record type is missing or not valid"}, sarifRuleConfig{"note"}},
	{"PROV031", "NotApplicable", sarifMessage{"A field is not applicable to this tier"}, sarifRuleConfig{"warning"}},
	{"PROV032", "SealedAtSHA", sarifMessage{"sealed_at_sha does not match the record status"}, sarifRuleConfig{"warning"}},
	{"PROV033", "LockedScope", sarifMessage{"An open record's scope conflicts with a locked record"}, sarifRuleConfig{"error"}},
	{"SCOPE001", "ScopeViolation", sarifMessage{"A changed file is outside, or forbidden by, the tagged record's scope"}, sarifRuleConfig{"error"}},
	{"TAG001", "CommitTag", sarifMessage{"The commit tag is missing or names the wrong record version"}, sarifRuleConfig{"error"}},
	{"STATE001", "RecordState", sarifMessage{"A snapshot was rewritten, a status transition is not allowed, or an implemented record changed without a revise"}, sarifRuleConfig{"error"}},
	{"EVID001", "EvidenceStale", sarifMessage{"An implemented record's evidence is stale"}, sarifRuleConfig{"error"}},
	{"EVID002", "EvidenceFailing", sarifMessage{"An implemented record's latest evidence is a failing run"}, sarifRuleConfig{"error"}},
}

// /sedum:setVar
// sedum:setFunction {"tier":"owned","record":"prov-2026-03078504","gen":"sha256:aaa940c9ff8508cc84857988bd8cb585939fdb57f5223d6e874b86715d0eb765","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func sarifVersion() string {\n\tif info, ok := debug.ReadBuildInfo(); ok && info.Main.Version != \"\" && info.Main.Version != \"(devel)\" {\n\t\treturn info.Main.Version\n\t}\n\treturn \"dev\"\n}","file":"cmd/provenance/sarif_run.go","name":"sarifVersion"}}
func sarifVersion() string {
	if info, ok := debug.ReadBuildInfo(); ok && info.Main.Version != "" && info.Main.Version != "(devel)" {
		return info.Main.Version
	}
	return "dev"
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-03078504","gen":"sha256:aaa940c9ff8508cc84857988bd8cb585939fdb57f5223d6e874b86715d0eb765","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func sarifMime(file string) string {\n\tif strings.HasSuffix(file, \".yml\") || strings.HasSuffix(file, \".yaml\") {\n\t\treturn \"application/yaml\"\n\t}\n\treturn \"text/plain\"\n}","file":"cmd/provenance/sarif_run.go","name":"sarifMime"}}
func sarifMime(file string) string {
	if strings.HasSuffix(file, ".yml") || strings.HasSuffix(file, ".yaml") {
		return "application/yaml"
	}
	return "text/plain"
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-03078504","gen":"sha256:aaa940c9ff8508cc84857988bd8cb585939fdb57f5223d6e874b86715d0eb765","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func buildSarif(r *repo, findings []sarifFinding) sarifDoc {\n\trules := append([]sarifRule{}, sarifCatalog...)\n\tindex := map[string]int{}\n\tfor i, rule := range rules {\n\t\tindex[rule.ID] = i\n\t}\n\tresults := []sarifResult{}\n\tfiles := map[string]bool{}\n\tfor _, f := range findings {\n\t\tat, known := index[f.RuleID]\n\t\tif !known {\n\t\t\tat = len(rules)\n\t\t\tindex[f.RuleID] = at\n\t\t\trules = append(rules, sarifRule{ID: f.RuleID, Name: f.RuleID, ShortDescription: sarifMessage{f.RuleID}, DefaultConfiguration: sarifRuleConfig{\"warning\"}})\n\t\t}\n\t\tres := sarifResult{RuleID: f.RuleID, RuleIndex: at, Level: f.Level, Message: sarifMessage{f.Message}}\n\t\tif f.File != \"\" {\n\t\t\tfiles[f.File] = true\n\t\t\tres.Locations = []sarifLocation{{PhysicalLocation: sarifPhysicalLocation{\n\t\t\t\tArtifactLocation: sarifArtifactLocation{URI: f.File, URIBaseID: \"%SRCROOT%\"},\n\t\t\t\tRegion:           sarifRegion{StartLine: 1, StartColumn: 1},\n\t\t\t}}}\n\t\t}\n\t\tprops := map[string]string{}\n\t\tif f.RecordID != \"\" {\n\t\t\tprops[\"recordId\"] = f.RecordID\n\t\t}\n\t\tif f.Field != \"\" {\n\t\t\tprops[\"field\"] = f.Field\n\t\t}\n\t\tif len(props) > 0 {\n\t\t\tres.Properties = props\n\t\t}\n\t\tresults = append(results, res)\n\t}\n\tvar uris []string\n\tfor f := range files {\n\t\turis = append(uris, f)\n\t}\n\tsort.Strings(uris)\n\tartifacts := []sarifArtifact{}\n\tfor _, uri := range uris {\n\t\tart := sarifArtifact{Location: sarifArtifactLocation{URI: uri, URIBaseID: \"%SRCROOT%\"}, MimeType: sarifMime(uri)}\n\t\tif data, err := os.ReadFile(filepath.Join(r.Root, filepath.FromSlash(uri))); err == nil {\n\t\t\tsum := sha256.Sum256(data)\n\t\t\tart.Hashes = map[string]string{\"sha-256\": hex.EncodeToString(sum[:])}\n\t\t}\n\t\tartifacts = append(artifacts, art)\n\t}\n\treturn sarifDoc{\n\t\tSchema:  \"https://json.schemastore.org/sarif-2.1.0.json\",\n\t\tVersion: \"2.1.0\",\n\t\tRuns: []sarifRun{{\n\t\t\tTool:      sarifTool{Driver: sarifDriver{Name: \"provenance\", Version: sarifVersion(), InformationURI: \"https://github.com/livecodelife/provenance\", Rules: rules}},\n\t\t\tResults:   results,\n\t\t\tArtifacts: artifacts,\n\t\t}},\n\t}\n}","file":"cmd/provenance/sarif_run.go","name":"buildSarif"}}
func buildSarif(r *repo, findings []sarifFinding) sarifDoc {
	rules := append([]sarifRule{}, sarifCatalog...)
	index := map[string]int{}
	for i, rule := range rules {
		index[rule.ID] = i
	}
	results := []sarifResult{}
	files := map[string]bool{}
	for _, f := range findings {
		at, known := index[f.RuleID]
		if !known {
			at = len(rules)
			index[f.RuleID] = at
			rules = append(rules, sarifRule{ID: f.RuleID, Name: f.RuleID, ShortDescription: sarifMessage{f.RuleID}, DefaultConfiguration: sarifRuleConfig{"warning"}})
		}
		res := sarifResult{RuleID: f.RuleID, RuleIndex: at, Level: f.Level, Message: sarifMessage{f.Message}}
		if f.File != "" {
			files[f.File] = true
			res.Locations = []sarifLocation{{PhysicalLocation: sarifPhysicalLocation{
				ArtifactLocation: sarifArtifactLocation{URI: f.File, URIBaseID: "%SRCROOT%"},
				Region:           sarifRegion{StartLine: 1, StartColumn: 1},
			}}}
		}
		props := map[string]string{}
		if f.RecordID != "" {
			props["recordId"] = f.RecordID
		}
		if f.Field != "" {
			props["field"] = f.Field
		}
		if len(props) > 0 {
			res.Properties = props
		}
		results = append(results, res)
	}
	var uris []string
	for f := range files {
		uris = append(uris, f)
	}
	sort.Strings(uris)
	artifacts := []sarifArtifact{}
	for _, uri := range uris {
		art := sarifArtifact{Location: sarifArtifactLocation{URI: uri, URIBaseID: "%SRCROOT%"}, MimeType: sarifMime(uri)}
		if data, err := os.ReadFile(filepath.Join(r.Root, filepath.FromSlash(uri))); err == nil {
			sum := sha256.Sum256(data)
			art.Hashes = map[string]string{"sha-256": hex.EncodeToString(sum[:])}
		}
		artifacts = append(artifacts, art)
	}
	return sarifDoc{
		Schema:  "https://json.schemastore.org/sarif-2.1.0.json",
		Version: "2.1.0",
		Runs: []sarifRun{{
			Tool:      sarifTool{Driver: sarifDriver{Name: "provenance", Version: sarifVersion(), InformationURI: "https://github.com/livecodelife/provenance", Rules: rules}},
			Results:   results,
			Artifacts: artifacts,
		}},
	}
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-03078504","gen":"sha256:aaa940c9ff8508cc84857988bd8cb585939fdb57f5223d6e874b86715d0eb765","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func printSarif(out io.Writer, doc sarifDoc) error {\n\tenc := json.NewEncoder(out)\n\tenc.SetIndent(\"\", \"  \")\n\treturn enc.Encode(doc)\n}","file":"cmd/provenance/sarif_run.go","name":"printSarif"}}
func printSarif(out io.Writer, doc sarifDoc) error {
	enc := json.NewEncoder(out)
	enc.SetIndent("", "  ")
	return enc.Encode(doc)
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-03078504","gen":"sha256:aaa940c9ff8508cc84857988bd8cb585939fdb57f5223d6e874b86715d0eb765","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func sarifLevel(sev lint.Severity) string {\n\tswitch sev {\n\tcase lint.SeverityError:\n\t\treturn \"error\"\n\tcase lint.SeverityWarning:\n\t\treturn \"warning\"\n\tdefault:\n\t\treturn \"note\"\n\t}\n}","file":"cmd/provenance/sarif_run.go","name":"sarifLevel"}}
func sarifLevel(sev lint.Severity) string {
	switch sev {
	case lint.SeverityError:
		return "error"
	case lint.SeverityWarning:
		return "warning"
	default:
		return "note"
	}
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-03078504","gen":"sha256:aaa940c9ff8508cc84857988bd8cb585939fdb57f5223d6e874b86715d0eb765","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func lintFindings(r *repo, res *lint.Result) []sarifFinding {\n\tvar out []sarifFinding\n\tfor _, issue := range res.Issues {\n\t\tf := sarifFinding{RuleID: issue.Rule, Level: sarifLevel(issue.Severity), Message: issue.Message, RecordID: issue.RecordID, Field: issue.Field}\n\t\tif issue.RecordID != \"\" {\n\t\t\tf.File = filepath.ToSlash(r.recordPath(issue.RecordID))\n\t\t}\n\t\tout = append(out, f)\n\t}\n\tsort.SliceStable(out, func(i, j int) bool {\n\t\ta, b := out[i], out[j]\n\t\tif a.File != b.File {\n\t\t\treturn a.File < b.File\n\t\t}\n\t\tif a.RuleID != b.RuleID {\n\t\t\treturn a.RuleID < b.RuleID\n\t\t}\n\t\treturn a.Message < b.Message\n\t})\n\treturn out\n}","file":"cmd/provenance/sarif_run.go","name":"lintFindings"}}
func lintFindings(r *repo, res *lint.Result) []sarifFinding {
	var out []sarifFinding
	for _, issue := range res.Issues {
		f := sarifFinding{RuleID: issue.Rule, Level: sarifLevel(issue.Severity), Message: issue.Message, RecordID: issue.RecordID, Field: issue.Field}
		if issue.RecordID != "" {
			f.File = filepath.ToSlash(r.recordPath(issue.RecordID))
		}
		out = append(out, f)
	}
	sort.SliceStable(out, func(i, j int) bool {
		a, b := out[i], out[j]
		if a.File != b.File {
			return a.File < b.File
		}
		if a.RuleID != b.RuleID {
			return a.RuleID < b.RuleID
		}
		return a.Message < b.Message
	})
	return out
}

// /sedum:setFunction
