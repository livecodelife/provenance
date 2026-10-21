package auth

import (
	"context"
	"errors"
	"testing"

	"github.com/livecodelife/provenance/auth/authtest"
)

func tokenClient(t *testing.T) (*Client, *authtest.Fake, string) {
	t.Helper()
	f := authtest.New(t)
	bearer := f.Mint().AccessToken
	return &Client{Server: f.URL()}, f, bearer
}

func TestCreateTokenSendsTheRequestAndReturnsTheSecretOnce(t *testing.T) {
	c, f, bearer := tokenClient(t)
	ctx := context.Background()
	tok, err := c.CreateToken(ctx, bearer, TokenRequest{Name: "ci-demo", Scopes: []string{"records:write"}, ExpiresInDays: 90})
	if err != nil {
		t.Fatal(err)
	}
	if tok.ID == "" || tok.Name != "ci-demo" || tok.Token == "" || tok.ExpiresAt == "" || len(tok.Scopes) != 1 || tok.Scopes[0] != "records:write" {
		t.Fatalf("token = %+v", tok)
	}
	rows := f.TokenRows()
	if len(rows) != 1 || rows[0].Name != "ci-demo" || rows[0].ExpiresInDays != 90 || rows[0].Revoked {
		t.Fatalf("rows = %+v", rows)
	}
	list, err := c.ListTokens(ctx, bearer)
	if err != nil || len(list) != 1 || list[0].ID != tok.ID || list[0].Token != "" {
		t.Fatalf("list = %+v %v: the secret must not come back from a list", list, err)
	}
}

func TestRevokeTokenAndUnknownID(t *testing.T) {
	c, f, bearer := tokenClient(t)
	ctx := context.Background()
	tok, _ := c.CreateToken(ctx, bearer, TokenRequest{Name: "x", Scopes: []string{"records:read"}})
	if err := c.RevokeToken(ctx, bearer, tok.ID); err != nil || !f.TokenRows()[0].Revoked {
		t.Fatalf("revoke: %v rows=%+v", err, f.TokenRows())
	}
	var oe *OAuthError
	if err := c.RevokeToken(ctx, bearer, "nope"); !errors.As(err, &oe) || oe.Status != 404 {
		t.Fatalf("unknown id: %v", err)
	}
}

func TestTokensRefusePersonalTokensAndBadSessions(t *testing.T) {
	c, _, _ := tokenClient(t)
	ctx := context.Background()
	var oe *OAuthError
	_, err := c.ListTokens(ctx, authtest.PersonalToken)
	if !errors.As(err, &oe) || oe.Status != 403 || oe.Code != "forbidden" {
		t.Fatalf("personal token: %v", err)
	}
	_, err = c.CreateToken(ctx, "ACCESS-unknown", TokenRequest{Name: "x", Scopes: []string{"records:read"}})
	if !errors.Is(err, ErrUnauthorized) {
		t.Fatalf("bad session: %v", err)
	}
}

func TestCreateTokenSurfacesTheServiceValidationMessage(t *testing.T) {
	c, _, bearer := tokenClient(t)
	var oe *OAuthError
	_, err := c.CreateToken(context.Background(), bearer, TokenRequest{Name: "", Scopes: []string{"records:read"}})
	if !errors.As(err, &oe) || oe.Status != 400 || oe.Code != "invalid_request" || oe.Description == "" {
		t.Fatalf("err = %v", err)
	}
}
