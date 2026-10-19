package main

import (
	"regexp"
	"testing"

	"github.com/livecodelife/provenance/version"
)

func TestVersionPrintsTheBareSemver(t *testing.T) {
	if !regexp.MustCompile(`^\d+\.\d+\.\d+$`).MatchString(version.Version) {
		t.Fatalf("version.Version = %q, want MAJOR.MINOR.PATCH", version.Version)
	}
	if got := mustRun(t, "version"); got != version.Version+"\n" {
		t.Fatalf("provenance version printed %q, want %q", got, version.Version+"\n")
	}
	if _, err := run(t, "version", "extra"); err == nil {
		t.Fatal("version takes no arguments")
	}
}
