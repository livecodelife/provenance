package auth

import (
	"context"
	"errors"
	"github.com/livecodelife/provenance/auth/authtest"
	"os"
	"path/filepath"
	"sync"
	"testing"
	"time"
)

func newManager(t *testing.T, f *authtest.Fake) (*Manager, FileStore) {
	dir := t.TempDir()
	store := FileStore{Path: filepath.Join(dir, ".provisio", "credentials")}
	return &Manager{Store: store, Client: &Client{Server: f.URL()}, LockPath: filepath.Join(dir, ".provisio", "credentials.lock")}, store
}

func saveCred(t *testing.T, store FileStore, f *authtest.Fake, access time.Duration) *authtest.Pair {
	tr := f.Mint()
	err := store.Save(&Credential{Server: f.URL(), AccessToken: tr.AccessToken, AccessExpiresAt: time.Now().Add(access), RefreshToken: tr.RefreshToken, RefreshExpiresAt: time.Now().Add(time.Hour), DID: tr.DID, Handle: tr.Handle})
	if err != nil {
		t.Fatal(err)
	}
	return tr
}

func TestTokenRefreshesWhenNearExpiry(t *testing.T) {
	t.Setenv(EnvToken, "")
	f := authtest.New(t)
	m, store := newManager(t, f)
	old := saveCred(t, store, f, 30*time.Second)
	tok, err := m.Token(context.Background())
	if err != nil || tok.Bearer == old.AccessToken || f.Refreshes != 1 {
		t.Fatalf("token = %+v %v refreshes=%d", tok, err, f.Refreshes)
	}
	again, _ := m.Token(context.Background())
	if again.Bearer != tok.Bearer || f.Refreshes != 1 {
		t.Fatal("a fresh token must be reused without a refresh")
	}
}

func TestDoRefreshesOnceOn401(t *testing.T) {
	t.Setenv(EnvToken, "")
	f := authtest.New(t)
	m, store := newManager(t, f)
	old := saveCred(t, store, f, time.Hour)
	delete(f.AccessOK, old.AccessToken) // server forgot the access token
	calls := 0
	err := m.Do(context.Background(), func(bearer string) error {
		calls++
		_, err := m.Client.Session(context.Background(), bearer)
		return err
	})
	if err != nil || calls != 2 || f.Refreshes != 1 {
		t.Fatalf("err=%v calls=%d refreshes=%d", err, calls, f.Refreshes)
	}
}

func TestSessionExpiredKeepsCredentialAndReports(t *testing.T) {
	t.Setenv(EnvToken, "")
	f := authtest.New(t)
	m, store := newManager(t, f)
	old := saveCred(t, store, f, time.Second)
	delete(f.RefreshOK, old.RefreshToken)
	if _, err := m.Token(context.Background()); !errors.Is(err, ErrSessionExpired) {
		t.Fatalf("err = %v", err)
	}
	if _, _, err := store.Load(); err != nil {
		t.Fatalf("credential must stay for logout to revoke: %v", err)
	}
}

func TestEnvTokenWinsAndTouchesNothing(t *testing.T) {
	t.Setenv(EnvToken, authtest.PersonalToken)
	f := authtest.New(t)
	m, store := newManager(t, f)
	tok, err := m.Token(context.Background())
	if err != nil || tok.Bearer != authtest.PersonalToken || tok.Source != SourceEnv || !tok.Personal {
		t.Fatalf("token = %+v %v", tok, err)
	}
	if _, err := os.Stat(store.Path); !errors.Is(err, os.ErrNotExist) {
		t.Fatal("credential file must not be created")
	}
	err = m.Do(context.Background(), func(string) error { return ErrUnauthorized })
	if !errors.Is(err, ErrTokenRejected) || f.Refreshes != 0 {
		t.Fatalf("err = %v refreshes=%d", err, f.Refreshes)
	}
}

func TestConcurrentRefreshRefreshesOnce(t *testing.T) {
	t.Setenv(EnvToken, "")
	f := authtest.New(t)
	m, store := newManager(t, f)
	saveCred(t, store, f, time.Second)
	var wg sync.WaitGroup
	for i := 0; i < 6; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if _, err := m.Token(context.Background()); err != nil {
				t.Error(err)
			}
		}()
	}
	wg.Wait()
	if f.Refreshes != 1 {
		t.Fatalf("refreshes = %d, want 1", f.Refreshes)
	}
}
