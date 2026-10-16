package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/livecodelife/provenance/auth"
	"github.com/livecodelife/provenance/auth/authtest"
)

// authEnv points the CLI at a fake auth service with an empty HOME, the file
// store (no keychain) and no PROVISIO_TOKEN, and returns the fake and HOME.
func authEnv(t *testing.T) (*authtest.Fake, string) {
	t.Helper()
	f := authtest.New(t)
	home := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv(auth.EnvNoKeychain, "1")
	t.Setenv(auth.EnvToken, "")
	t.Setenv(auth.EnvServer, f.URL())
	return f, home
}

func credPath(home string) string { return filepath.Join(home, ".provisio", "credentials") }

// runCLI runs the command tree with stdin and returns stdout and stderr apart.
func runCLI(t *testing.T, stdin string, args ...string) (stdout, stderr string, err error) {
	t.Helper()
	root := newRootCmd()
	var out, errb bytes.Buffer
	root.SetOut(&out)
	root.SetErr(&errb)
	root.SetIn(strings.NewReader(stdin))
	root.SetArgs(args)
	err = root.Execute()
	return out.String(), errb.String(), err
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
