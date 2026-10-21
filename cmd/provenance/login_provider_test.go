package main

import (
	"os"
	"strings"
	"testing"
)

func TestLoginSendsTheProviderHint(t *testing.T) {
	for _, tc := range []struct {
		args []string
		want string
	}{
		{[]string{"login", "--no-browser"}, "provisio"},
		{[]string{"login", "--no-browser", "--provider", "provisio"}, "provisio"},
		{[]string{"login", "--no-browser", "--provider", "bluesky"}, "bluesky"},
	} {
		f, _ := authEnv(t)
		if _, _, err := runCLI(t, "", tc.args...); err != nil {
			t.Fatalf("%v: %v", tc.args, err)
		}
		if f.CodeRequests != 1 || f.LastProvider != tc.want {
			t.Errorf("%v: requests=%d provider=%q, want %q", tc.args, f.CodeRequests, f.LastProvider, tc.want)
		}
	}
}

func TestLoginRejectsAnUnknownProviderBeforeAnyRequest(t *testing.T) {
	for _, bad := range []string{"github", "Bluesky", "", " provisio"} {
		f, home := authEnv(t)
		_, _, err := runCLI(t, "", "login", "--no-browser", "--provider", bad)
		if err == nil || !strings.Contains(err.Error(), "--provider must be provisio or bluesky") {
			t.Errorf("%q: err = %v", bad, err)
		}
		if f.CodeRequests != 0 || fileExists(credPath(home)) {
			t.Errorf("%q: a bad provider must not reach the service or store anything (requests=%d)", bad, f.CodeRequests)
		}
	}
}

func TestLoginProviderIsOnlyAHintAndIsNotStored(t *testing.T) {
	f, home := authEnv(t)
	if _, _, err := runCLI(t, "", "login", "--no-browser", "--provider", "bluesky"); err != nil {
		t.Fatal(err)
	}
	raw, err := os.ReadFile(credPath(home))
	if err != nil || strings.Contains(strings.ToLower(string(raw)), "bluesky") || strings.Contains(strings.ToLower(string(raw)), "provider") {
		t.Fatalf("the credential must not record the provider: %v\n%s", err, raw)
	}
	out, _, err := runCLI(t, "", "whoami")
	if err != nil || strings.Contains(strings.ToLower(out), "bluesky") || strings.Contains(strings.ToLower(out), "provider") {
		t.Fatalf("whoami must not show the provider: %v\n%s", err, out)
	}
	_ = f
}

func TestLoginWithTokenIgnoresAValidProvider(t *testing.T) {
	f, _ := authEnv(t)
	out, _, err := runCLI(t, "pvt_PERSONAL\n", "login", "--with-token", "--provider", "bluesky")
	if err != nil || f.CodeRequests != 0 {
		t.Fatalf("err=%v requests=%d\n%s", err, f.CodeRequests, out)
	}
}
