package remote

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// errorFrom serves one error answer and returns what the client reports.
func errorFrom(t *testing.T, status int, header, body string) *APIError {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if header != "" {
			w.Header().Set("X-Request-Id", header)
		}
		w.WriteHeader(status)
		w.Write([]byte(body))
	}))
	defer srv.Close()
	c := &Client{Base: srv.URL, Auth: func(_ context.Context, fn func(string) error) error { return fn("t") }}
	_, err := c.ListRecords(context.Background())
	ae, ok := err.(*APIError)
	if !ok {
		t.Fatalf("err = %T %v", err, err)
	}
	return ae
}

func TestAPIErrorCarriesTheRequestID(t *testing.T) {
	body := `{"error":{"code":"internal","message":"Internal error","requestId":"env-1"}}`
	ae := errorFrom(t, 500, "hdr-1", body)
	if ae.RequestID != "hdr-1" {
		t.Fatalf("header should win, got %q", ae.RequestID)
	}
	if got, want := ae.Error(), "internal (HTTP 500): Internal error (request hdr-1)"; got != want {
		t.Fatalf("Error() = %q, want %q", got, want)
	}
	if ae := errorFrom(t, 500, "", body); ae.RequestID != "env-1" {
		t.Fatalf("envelope fallback, got %q", ae.RequestID)
	}
	ae = errorFrom(t, 502, "", "")
	if ae.RequestID != "" || ae.Error() != "http_502 (HTTP 502)" {
		t.Fatalf("no id must change nothing: %q", ae.Error())
	}
}

func TestRequestIDFilter(t *testing.T) {
	long := strings.Repeat("a", 65)
	for _, bad := range []string{long, "a b", "a\\u001b[31m", "<x>", "a/b"} {
		body := `{"error":{"code":"internal","message":"m","requestId":"` + bad + `"}}`
		if ae := errorFrom(t, 500, "", body); ae.RequestID != "" || strings.Contains(ae.Error(), "request") {
			t.Fatalf("%q was printed: %q", bad, ae.Error())
		}
	}
	// A bad header falls back to a good envelope id.
	body := `{"error":{"code":"internal","message":"m","requestId":"ok.1_2-3"}}`
	if ae := errorFrom(t, 500, "bad id", body); ae.RequestID != "ok.1_2-3" {
		t.Fatalf("got %q", ae.RequestID)
	}
	if got := requestID(strings.Repeat("a", 64)); got == "" {
		t.Fatal("64 characters are allowed")
	}
}
