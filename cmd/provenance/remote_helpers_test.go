package main

import (
	"path/filepath"
	"testing"

	"github.com/livecodelife/provenance/auth/authtest"
	"github.com/livecodelife/provenance/remote"
	"github.com/livecodelife/provenance/remote/remotetest"
)

// remoteEnv is authEnv plus a fake record server that accepts the access
// tokens the fake auth service issued, with PROVISIO_API_URL pointing at it.
func remoteEnv(t *testing.T) (*authtest.Fake, *remotetest.Server) {
	t.Helper()
	f, _ := authEnv(t)
	srv := remotetest.New(t)
	srv.Accept = func(bearer string) bool { return f.AccessOK[bearer] }
	t.Setenv(remote.EnvBase, srv.URL())
	return f, srv
}

// signedIn runs login against the fake.
func signedIn(t *testing.T) {
	t.Helper()
	if _, _, err := runCLI(t, "", "login", "--no-browser"); err != nil {
		t.Fatal(err)
	}
}

// manifestLines reads a repo's version manifest.
func manifestLines(t *testing.T, root string) string {
	t.Helper()
	return readFile(t, filepath.Join(root, ".provenance", "versions.jsonl"))
}
