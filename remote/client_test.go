package remote

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/livecodelife/provenance/auth"
)

func TestClientRefusesPlainHTTPAndNotPublished(t *testing.T) {
	c := &Client{Base: "http://api.example.com/v1", Auth: func(_ context.Context, fn func(string) error) error { return fn("t") }}
	if _, err := c.ListRecords(context.Background()); err == nil || !strings.Contains(err.Error(), "https") {
		t.Fatalf("err = %v", err)
	}
	c2, _ := newClient(t)
	if _, err := c2.GetRecord(context.Background(), "prov-2026-00000000"); !errors.Is(err, ErrNotPublished) {
		t.Fatalf("err = %v", err)
	}
}

func TestClient401IsUnauthorizedAndErrorsNeverCarryTheToken(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(401)
	}))
	defer srv.Close()
	c := &Client{Base: srv.URL, Auth: func(_ context.Context, fn func(string) error) error { return fn("SECRET-TOKEN") }}
	if _, err := c.ListRecords(context.Background()); !errors.Is(err, auth.ErrUnauthorized) {
		t.Fatalf("err = %v", err)
	}
	srv.Close()
	_, err := c.ListRecords(context.Background())
	if err == nil || strings.Contains(err.Error(), "SECRET-TOKEN") {
		t.Fatalf("err = %v", err)
	}
}

func TestClientAPIErrorShape(t *testing.T) {
	c, _ := newClient(t)
	_, _, err := c.PutVersion(context.Background(), testID, PutVersionRequest{Version: 1})
	var ae *APIError
	if !errors.As(err, &ae) || ae.Status != 400 || ae.Code != "invalid_record" {
		t.Fatalf("err = %v", err)
	}
}

func TestBodylessNotFoundIsAWrongURLNotAnUnpublishedRecord(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
	}))
	defer srv.Close()
	c := &Client{Base: srv.URL, Auth: func(_ context.Context, fn func(string) error) error { return fn("t") }}
	for name, call := range map[string]func() error{
		"GetRecord":    func() error { _, err := c.GetRecord(context.Background(), "prov-2026-00000000"); return err },
		"ListVersions": func() error { _, err := c.ListVersions(context.Background(), "prov-2026-00000000"); return err },
	} {
		err := call()
		if errors.Is(err, ErrNotPublished) {
			t.Fatalf("%s: a 404 without a JSON body must not read as 'not published'", name)
		}
		if err == nil || !strings.Contains(err.Error(), "http_404") || !strings.Contains(err.Error(), "PROVISIO_API_URL") || !strings.Contains(err.Error(), "/v1") {
			t.Fatalf("%s: err = %v", name, err)
		}
	}
	// A 404 the service itself sends still means 'not published'.
	c2, _ := newClient(t)
	if _, err := c2.GetRecord(context.Background(), "prov-2026-00000000"); !errors.Is(err, ErrNotPublished) {
		t.Fatalf("err = %v", err)
	}
}
