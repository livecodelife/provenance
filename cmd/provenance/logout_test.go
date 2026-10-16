package main

import (
	"strings"
	"testing"
)

func TestLogoutRevokesAndDeletes(t *testing.T) {
	f, home := authEnv(t)
	if _, _, err := runCLI(t, "", "login", "--no-browser"); err != nil {
		t.Fatal(err)
	}
	out, _, err := runCLI(t, "", "logout")
	if err != nil || !strings.Contains(out, "Signed out") {
		t.Fatalf("logout: %v\n%s", err, out)
	}
	if f.Logouts != 1 || fileExists(credPath(home)) {
		t.Fatalf("logouts=%d, credential still there: %v", f.Logouts, fileExists(credPath(home)))
	}
	if len(f.RefreshOK) != 0 {
		t.Fatal("the refresh token must be revoked at the server")
	}
	out, _, err = runCLI(t, "", "logout")
	if err != nil || !strings.Contains(out, "Not signed in") {
		t.Fatalf("second logout: %v\n%s", err, out)
	}
}

func TestLogoutDeletesEvenWhenServerIsDown(t *testing.T) {
	f, home := authEnv(t)
	if _, _, err := runCLI(t, "", "login", "--no-browser"); err != nil {
		t.Fatal(err)
	}
	f.Close()
	out, _, err := runCLI(t, "", "logout")
	if err != nil {
		t.Fatal(err)
	}
	if fileExists(credPath(home)) || !strings.Contains(out, "could not be reached") {
		t.Fatalf("credential must go and the output must say the revocation failed:\n%s", out)
	}
}
