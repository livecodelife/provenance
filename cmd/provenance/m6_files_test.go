package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// The M6 driver is behind the m6 build tag so go test ./... never runs it.
// This test keeps the pieces around it in place: the script, the runbook, and a
// driver that still compiles.
func TestM6PiecesExist(t *testing.T) {
	root, err := filepath.Abs("../..")
	if err != nil {
		t.Fatal(err)
	}
	info, err := os.Stat(filepath.Join(root, "scripts", "m6.sh"))
	if err != nil || info.Mode()&0o111 == 0 {
		t.Fatalf("scripts/m6.sh must exist and be executable: %v", err)
	}
	runbook, err := os.ReadFile(filepath.Join(root, "docs", "M6.md"))
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{"M6_UI_CMD", "M6_APPROVE_URL", "throwaway", "provisio-m6-edit.cy.ts"} {
		if !strings.Contains(string(runbook), want) {
			t.Errorf("docs/M6.md does not mention %s", want)
		}
	}
	untagged := exec.Command("go", "test", "-count=1", "-run", "TestM6Joint", "-v", ".")
	out, err := untagged.CombinedOutput()
	if err != nil || strings.Contains(string(out), "--- RUN") {
		t.Fatalf("the driver must not run without the m6 tag: %v\n%s", err, out)
	}
	vet := exec.Command("go", "vet", "-tags", "m6", ".")
	if out, err := vet.CombinedOutput(); err != nil {
		t.Fatalf("the m6 driver does not compile: %v\n%s", err, out)
	}
}
