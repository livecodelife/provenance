package main

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/livecodelife/provenance/auth"
	"github.com/livecodelife/provenance/auth/authtest"
)

func TestTokenCreateUsesTheDefaultsAndPrintsTheSecretOnce(t *testing.T) {
	f, _ := authEnv(t)
	root, cfg := checkRepo(t)
	gitIn(t, root, "remote", "add", "origin", "https://github.com/o/myrepo.git")
	signedIn(t)
	out, errOut, err := runCLI(t, "", "-c", cfg, "token", "create")
	rows := f.TokenRows()
	if err != nil || len(rows) != 1 {
		t.Fatalf("err=%v rows=%+v\n%s", err, rows, out)
	}
	r := rows[0]
	if r.Name != "ci-myrepo" || len(r.Scopes) != 1 || r.Scopes[0] != "records:write" || r.ExpiresInDays != 90 {
		t.Fatalf("defaults: %+v", r)
	}
	secret := f.LastTokenSecret()
	if !strings.HasPrefix(secret, "pvt_") || strings.Count(out, secret) != 1 || strings.Contains(errOut, secret) {
		t.Fatalf("the secret must be on stdout exactly once and never on stderr:\n%s\n--\n%s", out, errOut)
	}
	if !strings.Contains(out, "ci-myrepo") || !strings.Contains(out, "not be shown again") {
		t.Fatalf("output:\n%s", out)
	}
}

func TestTokenCreateSecretOnlyPrintsJustTheSecret(t *testing.T) {
	f, _ := authEnv(t)
	signedIn(t)
	out, _, err := runCLI(t, "", "token", "create", "--name", "ci-x", "--secret-only")
	if err != nil || out != f.LastTokenSecret()+"\n" {
		t.Fatalf("err=%v out=%q", err, out)
	}
}

func TestTokenCreateJSONAndFlags(t *testing.T) {
	f, _ := authEnv(t)
	signedIn(t)
	out, _, err := runCLI(t, "", "token", "create", "--name", "ci-y", "--scope", "records:read", "--scope", "records:write", "--expires-in-days", "30", "--json")
	if err != nil {
		t.Fatal(err)
	}
	var v struct {
		ID     string   `json:"id"`
		Name   string   `json:"name"`
		Scopes []string `json:"scopes"`
		Token  string   `json:"token"`
	}
	if json.Unmarshal([]byte(out), &v) != nil || v.Token != f.LastTokenSecret() || v.ID == "" || v.Name != "ci-y" || len(v.Scopes) != 2 {
		t.Fatalf("json: %s", out)
	}
	if r := f.TokenRows()[0]; r.ExpiresInDays != 30 || len(r.Scopes) != 2 {
		t.Fatalf("row: %+v", r)
	}
}

func TestTokenCreateRejectsBadFlagsBeforeAnyRequest(t *testing.T) {
	f, _ := authEnv(t)
	signedIn(t)
	for _, args := range [][]string{
		{"token", "create", "--name", "x", "--scope", "admin"},
		{"token", "create", "--name", "x", "--expires-in-days", "0"},
		{"token", "create", "--name", "x", "--expires-in-days", "366"},
		{"token", "create", "--name", " "},
		{"token", "create", "--name", strings.Repeat("a", 61)},
	} {
		if _, _, err := runCLI(t, "", args...); err == nil {
			t.Errorf("%v: want an error", args)
		}
	}
	if len(f.TokenRows()) != 0 {
		t.Fatalf("a bad flag must not reach the service: %+v", f.TokenRows())
	}
	// Outside a repository there is no default name.
	t.Chdir(t.TempDir())
	if _, _, err := runCLI(t, "", "token", "create"); err == nil || !strings.Contains(err.Error(), "--name") {
		t.Fatalf("want --name required: %v", err)
	}
}

func TestTokenListShowsMetadataNeverSecrets(t *testing.T) {
	f, _ := authEnv(t)
	signedIn(t)
	runCLI(t, "", "token", "create", "--name", "ci-a")
	runCLI(t, "", "token", "create", "--name", "ci-b", "--expires-in-days", "7")
	out, _, err := runCLI(t, "", "token", "list")
	if err != nil || !strings.Contains(out, "ci-a") || !strings.Contains(out, "ci-b") || !strings.Contains(out, "records:write") {
		t.Fatalf("err=%v\n%s", err, out)
	}
	for _, s := range f.Secrets {
		if strings.Contains(out, s) {
			t.Fatalf("list leaked a secret:\n%s", out)
		}
	}
	out, _, err = runCLI(t, "", "token", "list", "--json")
	var rows []struct {
		ID, Name string
	}
	if err != nil || json.Unmarshal([]byte(out), &rows) != nil || len(rows) != 2 || strings.Contains(out, "pvt_") {
		t.Fatalf("json: %v\n%s", err, out)
	}
}

func TestTokenRevoke(t *testing.T) {
	f, _ := authEnv(t)
	signedIn(t)
	runCLI(t, "", "token", "create", "--name", "ci-a")
	id := f.TokenRows()[0].ID
	out, _, err := runCLI(t, "", "token", "revoke", id)
	if err != nil || !f.TokenRows()[0].Revoked || !strings.Contains(out, "Revoked") {
		t.Fatalf("err=%v out=%s", err, out)
	}
	if _, _, err := runCLI(t, "", "token", "revoke", "nope"); err == nil || !strings.Contains(err.Error(), "no token") {
		t.Fatalf("unknown id: %v", err)
	}
	if _, _, err := runCLI(t, "", "token", "revoke"); err == nil {
		t.Fatal("revoke needs an id")
	}
}

func TestTokenCommandsNeedASignedInSession(t *testing.T) {
	authEnv(t)
	_, _, err := runCLI(t, "", "token", "list")
	if err == nil || exitCode(err) != 3 || !strings.Contains(err.Error(), "needs a Provisio session") {
		t.Fatalf("not signed in: %v", err)
	}
}

func TestTokenCommandsRefuseAPersonalToken(t *testing.T) {
	f, _ := authEnv(t)
	t.Setenv(auth.EnvToken, authtest.PersonalToken)
	for _, args := range [][]string{{"token", "list"}, {"token", "create", "--name", "x"}, {"token", "revoke", "t1"}} {
		_, _, err := runCLI(t, "", args...)
		if err == nil || exitCode(err) != 3 || !strings.Contains(err.Error(), "provenance login") || strings.Contains(err.Error(), authtest.PersonalToken) {
			t.Errorf("%v: %v", args, err)
		}
	}
	if len(f.TokenRows()) != 0 {
		t.Fatal("a personal token must not mint tokens")
	}
}
