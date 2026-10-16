package auth

import (
	"encoding/json"
	"testing"
	"time"
)

func TestSessionExpiryForms(t *testing.T) {
	want := time.Date(2099, 1, 1, 0, 0, 0, 0, time.UTC)
	for name, body := range map[string]string{
		"unix seconds": `{"expires_at":4070908800}`,
		"rfc3339":      `{"expires_at":"2099-01-01T00:00:00Z"}`,
	} {
		var si SessionInfo
		if err := json.Unmarshal([]byte(body), &si); err != nil {
			t.Fatalf("%s: %v", name, err)
		}
		if got := si.Expiry(); !got.Equal(want) {
			t.Errorf("%s: Expiry = %v, want %v", name, got, want)
		}
	}
	for name, body := range map[string]string{"absent": `{}`, "null": `{"expires_at":null}`, "garbage": `{"expires_at":"soon"}`} {
		var si SessionInfo
		if err := json.Unmarshal([]byte(body), &si); err != nil || !si.Expiry().IsZero() {
			t.Errorf("%s: err=%v expiry=%v, want zero", name, err, si.Expiry())
		}
	}
}
