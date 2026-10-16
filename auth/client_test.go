package auth

import (
	"context"
	"errors"
	"github.com/livecodelife/provenance/auth/authtest"
	"strings"
	"testing"
	"time"
)

func noSleep(context.Context, time.Duration) error { return nil }

func TestDeviceFlowPendingThenApproved(t *testing.T) {
	f := authtest.New(t)
	f.Pending = 2
	c := &Client{Server: f.URL(), Hostname: "box"}
	dc, err := c.RequestDeviceCode(context.Background())
	if err != nil || dc.UserCode != "BDFG-HJKL" || f.LastHost != "box" {
		t.Fatalf("device code: %+v %v host=%q", dc, err, f.LastHost)
	}
	tr, err := c.PollToken(context.Background(), dc, noSleep)
	if err != nil || tr.Handle != "alice.test" || f.Polls != 3 {
		t.Fatalf("poll: %+v %v polls=%d", tr, err, f.Polls)
	}
}

func TestPollSlowDownAddsFiveSeconds(t *testing.T) {
	f := authtest.New(t)
	f.SlowDownAt, f.Pending = 1, 2
	c := &Client{Server: f.URL()}
	dc, _ := c.RequestDeviceCode(context.Background())
	var waits []time.Duration
	_, err := c.PollToken(context.Background(), dc, func(_ context.Context, d time.Duration) error {
		waits = append(waits, d)
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	want := []time.Duration{time.Second, 6 * time.Second, 6 * time.Second}
	if len(waits) != 3 || waits[0] != want[0] || waits[1] != want[1] || waits[2] != want[2] {
		t.Fatalf("waits = %v, want %v", waits, want)
	}
}

func TestPollDeniedAndExpired(t *testing.T) {
	for _, tc := range []struct {
		deny, expire bool
		want         error
	}{{true, false, ErrAccessDenied}, {false, true, ErrDeviceExpired}} {
		f := authtest.New(t)
		f.Deny, f.Expire = tc.deny, tc.expire
		c := &Client{Server: f.URL()}
		dc, _ := c.RequestDeviceCode(context.Background())
		if _, err := c.PollToken(context.Background(), dc, noSleep); !errors.Is(err, tc.want) {
			t.Fatalf("err = %v, want %v", err, tc.want)
		}
	}
}

func TestPollStopsOnContext(t *testing.T) {
	f := authtest.New(t)
	f.Pending = 1000
	c := &Client{Server: f.URL()}
	dc, _ := c.RequestDeviceCode(context.Background())
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := c.PollToken(ctx, dc, nil); !errors.Is(err, context.Canceled) {
		t.Fatalf("err = %v", err)
	}
}

func TestCheckServer(t *testing.T) {
	for _, ok := range []string{"https://auth.provisio.dev", "http://127.0.0.1:8080", "http://localhost:1"} {
		if err := CheckServer(ok); err != nil {
			t.Errorf("%s: %v", ok, err)
		}
	}
	for _, bad := range []string{"http://auth.provisio.dev", "ftp://x", "nonsense", ""} {
		if CheckServer(bad) == nil {
			t.Errorf("%q accepted", bad)
		}
	}
}

func TestRefreshRotationAndGrace(t *testing.T) {
	f := authtest.New(t)
	first := f.Mint()
	c := &Client{Server: f.URL()}
	a, err := c.Refresh(context.Background(), first.RefreshToken)
	if err != nil {
		t.Fatal(err)
	}
	b, err := c.Refresh(context.Background(), first.RefreshToken) // retry inside the grace window
	if err != nil || b.RefreshToken != a.RefreshToken || f.Refreshes != 1 {
		t.Fatalf("grace retry: %+v %v refreshes=%d", b, err, f.Refreshes)
	}
	if _, err := c.Refresh(context.Background(), "REFRESH-unknown"); !errors.Is(err, ErrSessionExpired) {
		t.Fatalf("unknown refresh = %v", err)
	}
}

func TestErrorsNeverCarryTokens(t *testing.T) {
	f := authtest.New(t)
	c := &Client{Server: "http://127.0.0.1:1"} // nothing listens
	_, err := c.Refresh(context.Background(), "REFRESH-SECRET")
	if err == nil || strings.Contains(err.Error(), "REFRESH-SECRET") {
		t.Fatalf("err = %v", err)
	}
	_, err = c.Session(context.Background(), "ACCESS-SECRET")
	if err == nil || strings.Contains(err.Error(), "ACCESS-SECRET") {
		t.Fatalf("err = %v", err)
	}
	_ = f
}
