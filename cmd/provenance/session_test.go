package main

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/livecodelife/provenance/auth"
)

func TestExitCodeAndSessionError(t *testing.T) {
	if exitCode(errors.New("boom")) != 1 || exitCode(nil) != 1 {
		t.Fatal("ordinary errors exit 1")
	}
	for _, tc := range []struct {
		in   error
		want string
	}{
		{auth.ErrNoSession, "needs a Provisio session. Run provenance login (or set PROVISIO_TOKEN in CI)."},
		{auth.ErrSessionExpired, "Your Provisio session has expired. Run provenance login."},
		{auth.ErrTokenRejected, "PROVISIO_TOKEN"},
	} {
		err := sessionError(tc.in)
		if exitCode(err) != 3 || !strings.Contains(err.Error(), tc.want) {
			t.Errorf("sessionError(%v) = %v (exit %d)", tc.in, err, exitCode(err))
		}
	}
	boom := errors.New("boom")
	if sessionError(boom) != boom || sessionError(nil) != nil {
		t.Fatal("other errors pass through")
	}
}

func TestWithSessionNeedsLogin(t *testing.T) {
	authEnv(t)
	err := withSession(context.Background(), func(string) error { t.Fatal("must not run"); return nil })
	if exitCode(err) != 3 || !strings.Contains(err.Error(), "needs a Provisio session") {
		t.Fatalf("err = %v", err)
	}
}

func TestWithSessionRefreshesOnce(t *testing.T) {
	f, _ := authEnv(t)
	if _, _, err := runCLI(t, "", "login", "--no-browser"); err != nil {
		t.Fatal(err)
	}
	f.AccessOK = map[string]bool{} // the server forgot the access token
	calls := 0
	err := withSession(context.Background(), func(bearer string) error {
		calls++
		if calls == 1 {
			return auth.ErrUnauthorized
		}
		if !f.AccessOK[bearer] {
			t.Errorf("the retry must carry the refreshed token")
		}
		return nil
	})
	if err != nil || calls != 2 || f.Refreshes != 1 {
		t.Fatalf("err=%v calls=%d refreshes=%d", err, calls, f.Refreshes)
	}
}
