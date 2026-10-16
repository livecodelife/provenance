package main

import (
	"os"
	"strings"
	"testing"

	"github.com/livecodelife/provenance/auth/authtest"
)

func TestLoginDeviceFlow(t *testing.T) {
	f, home := authEnv(t)
	f.Pending = 1
	out, _, err := runCLI(t, "", "login", "--no-browser")
	if err != nil {
		t.Fatal(err, out)
	}
	for _, want := range []string{f.URL() + "/device", "BDFG-HJKL", "Signed in as @alice.test (did:plc:test)"} {
		if !strings.Contains(out, want) {
			t.Errorf("login output lacks %q:\n%s", want, out)
		}
	}
	info, err := os.Stat(credPath(home))
	if err != nil || info.Mode().Perm() != 0o600 {
		t.Fatalf("credential file: %v %v", info, err)
	}
	if f.LastHost == "" {
		t.Error("the hostname was not sent")
	}
}

func TestLoginDeniedStoresNothing(t *testing.T) {
	f, home := authEnv(t)
	f.Deny = true
	if _, _, err := runCLI(t, "", "login", "--no-browser"); err == nil {
		t.Fatal("a denied sign-in must fail")
	}
	if fileExists(credPath(home)) {
		t.Fatal("nothing may be stored after a denial")
	}
}

func TestLoginWithPersonalToken(t *testing.T) {
	_, home := authEnv(t)
	if _, _, err := runCLI(t, "pvt_BAD\n", "login", "--with-token"); err == nil {
		t.Fatal("a token the server rejects must fail")
	}
	if fileExists(credPath(home)) {
		t.Fatal("a rejected token must not be stored")
	}
	if _, _, err := runCLI(t, "not-a-token\n", "login", "--with-token"); err == nil {
		t.Fatal("a string without the pvt_ prefix must fail")
	}
	out, _, err := runCLI(t, authtest.PersonalToken+"\n", "login", "--with-token")
	if err != nil || !strings.Contains(out, "Signed in as @alice.test") {
		t.Fatalf("login --with-token: %v\n%s", err, out)
	}
	who, _, err := runCLI(t, "", "whoami", "--json")
	if err != nil || !strings.Contains(who, `"name": "ci"`) {
		t.Fatalf("whoami after --with-token: %v\n%s", err, who)
	}
}

func TestLoginRefusesPlainHTTPServer(t *testing.T) {
	_, home := authEnv(t)
	if _, _, err := runCLI(t, "", "login", "--no-browser", "--server", "http://auth.example.com"); err == nil {
		t.Fatal("http to a non-loopback host must be refused")
	}
	if fileExists(credPath(home)) {
		t.Fatal("nothing may be stored")
	}
}
