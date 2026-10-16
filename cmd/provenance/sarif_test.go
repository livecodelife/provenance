package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/santhosh-tekuri/jsonschema/v5"
)

// validateSarif checks doc against the vendored OASIS SARIF 2.1.0 schema, so
// the test does not only check this package's own idea of the shape.
func validateSarif(t *testing.T, doc string) error {
	t.Helper()
	schema, err := os.ReadFile(filepath.Join("testdata", "sarif-schema-2.1.0.json"))
	if err != nil {
		t.Fatal(err)
	}
	c := jsonschema.NewCompiler()
	if err := c.AddResource("sarif-schema-2.1.0.json", bytes.NewReader(schema)); err != nil {
		t.Fatal(err)
	}
	sch, err := c.Compile("sarif-schema-2.1.0.json")
	if err != nil {
		t.Fatal(err)
	}
	var v any
	if err := json.Unmarshal([]byte(doc), &v); err != nil {
		t.Fatalf("not JSON: %v\n%s", err, doc)
	}
	return sch.Validate(v)
}

func TestLintSarif(t *testing.T) {
	root, cfg := v2Repo(t, false, map[string]string{draftID: v2Draft(draftID, "bogus")})
	out, err := run(t, "-c", cfg, "lint", "--format", "sarif")
	if err == nil {
		t.Fatal("lint errors must still fail the command with --format sarif")
	}
	if verr := validateSarif(t, out); verr != nil {
		t.Fatalf("not valid SARIF 2.1.0: %v\n%s", verr, out)
	}
	var doc sarifDoc
	if err := json.Unmarshal([]byte(out), &doc); err != nil {
		t.Fatal(err)
	}
	run0 := doc.Runs[0]
	if run0.Tool.Driver.Name != "provenance" || len(run0.Tool.Driver.Rules) < 30 {
		t.Fatalf("driver %+v", run0.Tool.Driver.Name)
	}
	var found bool
	for _, res := range run0.Results {
		if res.RuleID == "PROV003" {
			found = true
			uri := res.Locations[0].PhysicalLocation.ArtifactLocation.URI
			if res.Level != "error" || uri != "provenance/"+draftID+".yml" || res.Properties["recordId"] != draftID || run0.Tool.Driver.Rules[res.RuleIndex].ID != "PROV003" {
				t.Fatalf("result %+v", res)
			}
		}
	}
	if !found {
		t.Fatalf("no PROV003 result:\n%s", out)
	}
	if len(run0.Artifacts) != 1 || run0.Artifacts[0].Hashes["sha-256"] == "" || run0.Artifacts[0].Location.URI != "provenance/"+draftID+".yml" {
		t.Fatalf("artifacts %+v", run0.Artifacts)
	}
	if again, _ := run(t, "-c", cfg, "lint", "--format", "sarif"); again != out {
		t.Fatal("the output is not deterministic")
	}
	if !strings.Contains(out, "\n  \"version\": \"2.1.0\"") {
		t.Fatalf("output:\n%s", out)
	}
	_ = root
}

func TestLintSarifCleanAndSchemaTeeth(t *testing.T) {
	_, cfg := openedRepo(t, draftID)
	out := mustRun(t, "-c", cfg, "lint", "--format", "sarif")
	if err := validateSarif(t, out); err != nil {
		t.Fatalf("clean repo SARIF invalid: %v\n%s", err, out)
	}
	var doc map[string]any
	if err := json.Unmarshal([]byte(out), &doc); err != nil {
		t.Fatal(err)
	}
	delete(doc, "version")
	broken, _ := json.Marshal(doc)
	if validateSarif(t, string(broken)) == nil {
		t.Fatal("the schema validator accepted a document without a version: the check has no teeth")
	}
}
