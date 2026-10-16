package main

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/livecodelife/provenance/auth"
)

func TestWhoamiLoggedOut(t *testing.T) {
	authEnv(t)
	out, _, err := runCLI(t, "", "whoami")
	if err == nil || exitCode(err) != 1 || !strings.Contains(out, "Not signed in") {
		t.Fatalf("err=%v code=%d out=%q", err, exitCode(err), out)
	}
}

func TestWhoamiShowsAccountAndSource(t *testing.T) {
	authEnv(t)
	if _, _, err := runCLI(t, "", "login", "--no-browser"); err != nil {
		t.Fatal(err)
	}
	out, _, err := runCLI(t, "", "whoami")
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{"@alice.test", "did:plc:test", "Source:  file", "Expires: 2099-01-01T00:00:00Z"} {
		if !strings.Contains(out, want) {
			t.Errorf("whoami lacks %q:\n%s", want, out)
		}
	}
	js, _, err := runCLI(t, "", "whoami", "--json")
	if err != nil {
		t.Fatal(err)
	}
	var v map[string]any
	if err := json.Unmarshal([]byte(js), &v); err != nil || v["handle"] != "alice.test" || v["did"] != "did:plc:test" || v["source"] != "file" {
		t.Fatalf("json: %v\n%s", err, js)
	}
	if strings.Contains(strings.ToLower(js), "token\":") && v["token"] != nil {
		t.Fatalf("a session login must not report a personal token: %s", js)
	}
}

func TestWhoamiExpiredSessionExitsThree(t *testing.T) {
	f, _ := authEnv(t)
	if _, _, err := runCLI(t, "", "login", "--no-browser"); err != nil {
		t.Fatal(err)
	}
	for tok := range f.RefreshOK {
		delete(f.RefreshOK, tok)
	}
	f.AccessOK = map[string]bool{}
	_, _, err := runCLI(t, "", "whoami")
	if err == nil || exitCode(err) != 3 || !strings.Contains(err.Error(), "Run provenance login") {
		t.Fatalf("err=%v code=%d", err, exitCode(err))
	}
}

func TestWhoamiEnvToken(t *testing.T) {
	authEnv(t)
	t.Setenv(auth.EnvToken, "pvt_PERSONAL")
	out, _, err := runCLI(t, "", "whoami")
	if err != nil || !strings.Contains(out, "Source:  env") {
		t.Fatalf("whoami with %s: %v\n%s", auth.EnvToken, err, out)
	}
	t.Setenv(auth.EnvToken, "pvt_REVOKED")
	_, _, err = runCLI(t, "", "whoami")
	if err == nil || exitCode(err) != 3 || !strings.Contains(err.Error(), auth.EnvToken) {
		t.Fatalf("a rejected %s must exit 3 and name the variable: %v", auth.EnvToken, err)
	}
}
