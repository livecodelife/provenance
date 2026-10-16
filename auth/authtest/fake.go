// Package authtest is an in-memory Provisio auth service built from
// provisio-backend docs/DEVICE-FLOW.md, for tests of the auth package and of
// the commands that use it.
package authtest

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"sync"
	"testing"
)

// Fake implements the DEVICE-FLOW.md endpoints in memory.
type Fake struct {
	t   *testing.T
	srv *httptest.Server

	mu sync.Mutex

	Pending     int // authorization_pending answers before approval
	SlowDownAt  int // poll number that answers slow_down (0 = never)
	Deny        bool
	Expire      bool
	Polls       int
	Refreshes   int
	Logouts     int
	counter     int
	RefreshOK   map[string]bool // live refresh tokens
	AccessOK    map[string]bool // live access tokens
	lastRefresh string          // previous token, for the grace window
	lastPair    *Pair
	Secrets     []string // every token the fake ever issued
	LastHost    string
}

func New(t *testing.T) *Fake {
	f := &Fake{t: t, RefreshOK: map[string]bool{}, AccessOK: map[string]bool{}}
	mux := http.NewServeMux()
	mux.HandleFunc("/device/code", f.deviceCode)
	mux.HandleFunc("/device/token", f.deviceToken)
	mux.HandleFunc("/auth/refresh", f.refresh)
	mux.HandleFunc("/auth/logout", f.logout)
	mux.HandleFunc("/session", f.session)
	f.srv = httptest.NewServer(mux)
	t.Cleanup(f.srv.Close)
	return f
}

// PersonalToken is the personal token the fake accepts on GET /session.
const PersonalToken = "pvt_PERSONAL"

// Pair is a token pair the fake issued.
type Pair struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int    `json:"expires_in"`
	RefreshToken     string `json:"refresh_token"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
	DID              string `json:"did"`
	Handle           string `json:"handle"`
}

// URL is the fake's base URL.
func (f *Fake) URL() string { return f.srv.URL }

func reply(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func oauthErr(w http.ResponseWriter, status int, code string) {
	reply(w, status, map[string]string{"error": code})
}

// Mint issues a live token pair, as an approved sign-in would.
func (f *Fake) Mint() *Pair {
	f.counter++
	n := strconv.Itoa(f.counter)
	tr := &Pair{AccessToken: "ACCESS-" + n, ExpiresIn: 900, RefreshToken: "REFRESH-" + n, RefreshExpiresIn: 2592000, DID: "did:plc:test", Handle: "alice.test"}
	f.AccessOK[tr.AccessToken] = true
	f.RefreshOK[tr.RefreshToken] = true
	f.Secrets = append(f.Secrets, tr.AccessToken, tr.RefreshToken)
	return tr
}

func (f *Fake) deviceCode(w http.ResponseWriter, r *http.Request) {
	var in map[string]string
	json.NewDecoder(r.Body).Decode(&in)
	if in["client_id"] != "provenance-cli" {
		oauthErr(w, 400, "invalid_client")
		return
	}
	f.mu.Lock()
	f.LastHost = in["hostname"]
	f.Secrets = append(f.Secrets, "DEVICE-SECRET")
	f.mu.Unlock()
	reply(w, 200, map[string]any{"device_code": "DEVICE-SECRET", "user_code": "BDFG-HJKL", "verification_uri": f.srv.URL + "/device", "verification_uri_complete": f.srv.URL + "/device?user_code=BDFG-HJKL", "expires_in": 600, "interval": 1})
}

func (f *Fake) deviceToken(w http.ResponseWriter, r *http.Request) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.Polls++
	switch {
	case f.Expire:
		oauthErr(w, 400, "expired_token")
	case f.Deny:
		oauthErr(w, 400, "access_denied")
	case f.SlowDownAt == f.Polls:
		oauthErr(w, 400, "slow_down")
	case f.Polls <= f.Pending:
		oauthErr(w, 400, "authorization_pending")
	default:
		reply(w, 200, f.Mint())
	}
}

func (f *Fake) refresh(w http.ResponseWriter, r *http.Request) {
	var in map[string]string
	json.NewDecoder(r.Body).Decode(&in)
	f.mu.Lock()
	defer f.mu.Unlock()
	tok := in["refresh_token"]
	if tok == f.lastRefresh && f.lastPair != nil {
		reply(w, 200, f.lastPair)
		return
	}
	if !f.RefreshOK[tok] {
		reply(w, 401, map[string]string{"error": "invalid_grant", "error_description": "session expired"})
		return
	}
	f.Refreshes++
	delete(f.RefreshOK, tok)
	pair := f.Mint()
	f.lastRefresh, f.lastPair = tok, pair
	reply(w, 200, pair)
}

func (f *Fake) logout(w http.ResponseWriter, r *http.Request) {
	var in map[string]string
	json.NewDecoder(r.Body).Decode(&in)
	f.mu.Lock()
	defer f.mu.Unlock()
	f.Logouts++
	for tok := range f.RefreshOK {
		if tok == in["refresh_token"] {
			delete(f.RefreshOK, tok)
		}
	}
	f.lastPair, f.lastRefresh = nil, ""
	reply(w, 200, map[string]bool{"ok": true})
}

func (f *Fake) session(w http.ResponseWriter, r *http.Request) {
	f.mu.Lock()
	defer f.mu.Unlock()
	bearer := r.Header.Get("Authorization")
	if len(bearer) > 7 {
		bearer = bearer[7:]
	}
	if bearer == PersonalToken {
		reply(w, 200, map[string]any{"did": "did:plc:test", "handle": "alice.test", "expires_at": "2099-01-01T00:00:00Z", "token": map[string]any{"id": "t1", "name": "ci", "scopes": []string{"records:read"}}})
		return
	}
	if !f.AccessOK[bearer] {
		reply(w, 401, map[string]any{"error": map[string]string{"code": "unauthorized"}})
		return
	}
	reply(w, 200, map[string]any{"did": "did:plc:test", "handle": "alice.test", "expires_at": 4070908800})
}

// Close stops the server, as if the auth service went down.
func (f *Fake) Close() { f.srv.Close() }
