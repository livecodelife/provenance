package main

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/livecodelife/provenance/remote"
)

func TestInsufficientScopeAsksForANewSignIn(t *testing.T) {
	ae := &remote.APIError{Status: 403, Code: "insufficient_scope", Message: "token lacks repo write"}
	for _, err := range []error{ae, fmt.Errorf("publish: %w", ae)} {
		got := sessionError(err)
		if exitCode(got) != 3 {
			t.Fatalf("exit status = %d, want 3", exitCode(got))
		}
		if !strings.Contains(got.Error(), "sign in again with `provenance login`") {
			t.Fatalf("message = %q", got)
		}
	}
}

func TestOtherForbiddenAnswersPassThrough(t *testing.T) {
	ae := &remote.APIError{Status: 403, Code: "forbidden", Message: "not yours"}
	if got := sessionError(ae); !errors.Is(got, ae) || exitCode(got) != 1 {
		t.Fatalf("got %v (exit %d)", got, exitCode(got))
	}
}
