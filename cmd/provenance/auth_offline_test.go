package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/livecodelife/provenance/auth"
	"github.com/livecodelife/provenance/auth/authtest"
)

// Local commands never read the credential store and never touch the network:
// with a credential file nobody may read (mode 0644 is refused by the store),
// no PROVISIO_TOKEN and an auth server that is down, they behave as before.
func TestLocalCommandsNeedNoCredentialOrNetwork(t *testing.T) {
	root, cfg := checkRepo(t)
	f := authtest.New(t)
	t.Setenv(auth.EnvServer, f.URL())
	f.Close()
	home := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv(auth.EnvNoKeychain, "")
	t.Setenv(auth.EnvToken, "")
	if err := os.MkdirAll(filepath.Join(home, ".provisio"), 0o755); err != nil {
		t.Fatal(err)
	}
	bad := credPath(home)
	if err := os.WriteFile(bad, []byte("not json"), 0o644); err != nil {
		t.Fatal(err)
	}
	for _, args := range [][]string{
		{"status"}, {"lint"}, {"stale"}, {"history", chkID}, {"next"}, {"govern", "a.go"},
		{"graph"}, {"search", "record"}, {"context", "a.go"}, {"compile"}, {"reconcile"},
	} {
		full := append([]string{"-c", cfg}, args...)
		if _, _, err := runCLI(t, "", full...); err != nil && strings.Contains(err.Error(), "credential") {
			t.Errorf("provenance %s touched the credential store: %v", strings.Join(args, " "), err)
		}
	}
	data, err := os.ReadFile(bad)
	if err != nil || string(data) != "not json" {
		t.Fatalf("the credential file was changed: %v %q", err, data)
	}
	_ = root
}
