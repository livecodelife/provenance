package main

import (
	"strings"
	"testing"

	"github.com/livecodelife/provenance/auth"
	"github.com/livecodelife/provenance/auth/authtest"
)

// No token the service issues, and no token the user supplies, may reach
// stdout, stderr or an error message in any login, whoami or logout path.
func TestNoSecretReachesOutput(t *testing.T) {
	f, _ := authEnv(t)
	var seen []string
	capture := func(stdin string, args ...string) {
		out, errOut, err := runCLI(t, stdin, args...)
		seen = append(seen, out, errOut)
		if err != nil {
			seen = append(seen, err.Error())
		}
	}
	capture("", "login", "--no-browser")
	capture("", "whoami")
	capture("", "whoami", "--json")
	capture("", "logout")
	capture("", "logout")
	capture("pvt_BAD\n", "login", "--with-token")
	capture(authtest.PersonalToken+"\n", "login", "--with-token")
	capture("", "whoami", "--json")
	capture("", "logout")
	t.Setenv(auth.EnvToken, "pvt_ENVSECRET")
	capture("", "whoami")
	capture("", "whoami", "--json")
	f.Close()
	capture("", "whoami")
	capture("", "login", "--no-browser")
	secrets := append(f.Secrets, authtest.PersonalToken, "pvt_BAD", "pvt_ENVSECRET")
	all := strings.Join(seen, "\n")
	for _, s := range secrets {
		if strings.Contains(all, s) {
			t.Errorf("secret %q reached the output", s)
		}
	}
	if len(f.Secrets) < 3 {
		t.Fatalf("the fake issued too few tokens (%d) for this test to mean anything", len(f.Secrets))
	}
}
