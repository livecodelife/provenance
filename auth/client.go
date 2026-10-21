package auth

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-03700026","gen":"sha256:cfe4d909b4ae2dcf117a0f65f0093e2fd8e532cec5c63cbe210a1a30769d022a","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"auth/client.go","path":"net/http"}}
import "net/http"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-2e8102ae","gen":"sha256:4118998e877b335b07f038fe5c3d67b41eccb52209c86d11105d7a077fc42ab4","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"auth/client.go","path":"fmt"}}
import "fmt"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-483067cd","gen":"sha256:a23085a101da665c67700ec1d6e4066919cd043916b62d84dca40dd47737356d","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"auth/client.go","path":"time"}}
import "time"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-6f0f74ec","gen":"sha256:2816476fbb50c5651fc5c82d92b5c98908bebac5977a96250030a539aba83090","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"auth/client.go","path":"net/url"}}
import "net/url"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-87b7ba7b","gen":"sha256:0cb602be28222f7017b0829cf50dc48c7ad99ed6cb36b0ae1b438c53ca3a9008","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"auth/client.go","path":"bytes"}}
import "bytes"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-87b7ba7b","gen":"sha256:0cb602be28222f7017b0829cf50dc48c7ad99ed6cb36b0ae1b438c53ca3a9008","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"auth/client.go","path":"errors"}}
import "errors"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-87b7ba7b","gen":"sha256:0cb602be28222f7017b0829cf50dc48c7ad99ed6cb36b0ae1b438c53ca3a9008","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"auth/client.go","path":"io"}}
import "io"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-87b7ba7b","gen":"sha256:0cb602be28222f7017b0829cf50dc48c7ad99ed6cb36b0ae1b438c53ca3a9008","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"auth/client.go","path":"strconv"}}
import "strconv"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-87b7ba7b","gen":"sha256:0cb602be28222f7017b0829cf50dc48c7ad99ed6cb36b0ae1b438c53ca3a9008","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"auth/client.go","path":"strings"}}
import "strings"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-87b7ba7b","gen":"sha256:0cb602be28222f7017b0829cf50dc48c7ad99ed6cb36b0ae1b438c53ca3a9008","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"auth/client.go","path":"context"}}
import "context"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-b4e29172","gen":"sha256:46b8f72efdcc18667736f3c3568613d3a30a772b02626de5d8ebb6e69f9722dc","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"auth/client.go","path":"encoding/json"}}
import "encoding/json"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setType {"tier":"owned","record":"prov-2026-03700026","gen":"sha256:cfe4d909b4ae2dcf117a0f65f0093e2fd8e532cec5c63cbe210a1a30769d022a","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type Client struct {\n\tServer   string\n\tHTTP     *http.Client\n\tHostname string\n\tProvider string\n}","file":"auth/client.go","name":"Client"}}
type Client struct {
	Server   string
	HTTP     *http.Client
	Hostname string
	Provider string
}

// /sedum:setType
// sedum:setMethod {"tier":"owned","record":"prov-2026-12513118","gen":"sha256:607db55180afa370158f2b155a123e906a1b344a8c5cca9eebc21a898c5bd54e","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (c *Client) Refresh(ctx context.Context, refreshToken string) (*TokenResponse, error) {\n\tbody := map[string]string{\"refresh_token\": refreshToken}\n\tvar out TokenResponse\n\tif err := c.do(ctx, http.MethodPost, \"/auth/refresh\", \"\", body, &out); err != nil {\n\t\tvar oauthErr *OAuthError\n\t\tif errors.As(err, &oauthErr) && (oauthErr.Code == \"invalid_grant\" || oauthErr.Status == http.StatusUnauthorized) {\n\t\t\treturn nil, ErrSessionExpired\n\t\t}\n\t\treturn nil, err\n\t}\n\treturn &out, nil\n}","file":"auth/client.go","name":"Refresh","receiver":"Client"}}
func (c *Client) Refresh(ctx context.Context, refreshToken string) (*TokenResponse, error) {
	body := map[string]string{"refresh_token": refreshToken}
	var out TokenResponse
	if err := c.do(ctx, http.MethodPost, "/auth/refresh", "", body, &out); err != nil {
		var oauthErr *OAuthError
		if errors.As(err, &oauthErr) && (oauthErr.Code == "invalid_grant" || oauthErr.Status == http.StatusUnauthorized) {
			return nil, ErrSessionExpired
		}
		return nil, err
	}
	return &out, nil
}

// /sedum:setMethod
// sedum:setMethod {"tier":"owned","record":"prov-2026-2e8102ae","gen":"sha256:4118998e877b335b07f038fe5c3d67b41eccb52209c86d11105d7a077fc42ab4","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (e *OAuthError) Error() string {\n\tif e.Description != \"\" {\n\t\treturn fmt.Sprintf(\"%s (HTTP %d): %s\", e.Code, e.Status, e.Description)\n\t}\n\treturn fmt.Sprintf(\"%s (HTTP %d)\", e.Code, e.Status)\n}","file":"auth/client.go","name":"Error","receiver":"OAuthError"}}
func (e *OAuthError) Error() string {
	if e.Description != "" {
		return fmt.Sprintf("%s (HTTP %d): %s", e.Code, e.Status, e.Description)
	}
	return fmt.Sprintf("%s (HTTP %d)", e.Code, e.Status)
}

// /sedum:setMethod
// sedum:setType {"tier":"owned","record":"prov-2026-3ab144df","gen":"sha256:20a404daa1fc8e6821660cbaf4df854a71f48f691f25dcbec028b78ff517a3c6","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type TokenRequest struct {\n\tName          string   `json:\"name\"`\n\tScopes        []string `json:\"scopes\"`\n\tExpiresInDays int      `json:\"expires_in_days,omitempty\"`\n}","file":"auth/client.go","name":"TokenRequest"}}
type TokenRequest struct {
	Name          string   `json:"name"`
	Scopes        []string `json:"scopes"`
	ExpiresInDays int      `json:"expires_in_days,omitempty"`
}

// /sedum:setType
// sedum:setFunction {"tier":"owned","record":"prov-2026-4289a60b","gen":"sha256:5b24a5b9a985141bbf7687e1ba8d6d3a78ce3586cd86c0041d663b1730b87e45","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func Wait(ctx context.Context, d time.Duration) error {\n\tt := time.NewTimer(d)\n\tdefer t.Stop()\n\tselect {\n\tcase <-t.C:\n\t\treturn nil\n\tcase <-ctx.Done():\n\t\treturn ctx.Err()\n\t}\n}","file":"auth/client.go","name":"Wait"}}
func Wait(ctx context.Context, d time.Duration) error {
	t := time.NewTimer(d)
	defer t.Stop()
	select {
	case <-t.C:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// /sedum:setFunction
// sedum:setType {"tier":"owned","record":"prov-2026-46d084ef","gen":"sha256:b65cae1456257e5cbdb9bcdfa745295f2208b766b4d7e4abb1cc20226ea90cc3","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type TokenView struct {\n\tID         string   `json:\"id\"`\n\tName       string   `json:\"name\"`\n\tScopes     []string `json:\"scopes\"`\n\tCreatedAt  string   `json:\"created_at\"`\n\tLastUsedAt string   `json:\"last_used_at\"`\n\tExpiresAt  string   `json:\"expires_at\"`\n\tToken      string   `json:\"token,omitempty\"`\n}","file":"auth/client.go","name":"TokenView"}}
type TokenView struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Scopes     []string `json:"scopes"`
	CreatedAt  string   `json:"created_at"`
	LastUsedAt string   `json:"last_used_at"`
	ExpiresAt  string   `json:"expires_at"`
	Token      string   `json:"token,omitempty"`
}

// /sedum:setType
// sedum:setMethod {"tier":"owned","record":"prov-2026-483067cd","gen":"sha256:a23085a101da665c67700ec1d6e4066919cd043916b62d84dca40dd47737356d","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (s *SessionInfo) Expiry() time.Time {\n\tif s == nil || len(s.ExpiresAt) == 0 {\n\t\treturn time.Time{}\n\t}\n\tvar v interface{}\n\tif err := json.Unmarshal(s.ExpiresAt, &v); err != nil {\n\t\treturn time.Time{}\n\t}\n\tswitch x := v.(type) {\n\tcase float64:\n\t\tif x <= 0 {\n\t\t\treturn time.Time{}\n\t\t}\n\t\treturn time.Unix(int64(x), 0)\n\tcase string:\n\t\tt, err := time.Parse(time.RFC3339, x)\n\t\tif err != nil {\n\t\t\treturn time.Time{}\n\t\t}\n\t\treturn t\n\t}\n\treturn time.Time{}\n}","file":"auth/client.go","name":"Expiry","receiver":"SessionInfo"}}
func (s *SessionInfo) Expiry() time.Time {
	if s == nil || len(s.ExpiresAt) == 0 {
		return time.Time{}
	}
	var v interface{}
	if err := json.Unmarshal(s.ExpiresAt, &v); err != nil {
		return time.Time{}
	}
	switch x := v.(type) {
	case float64:
		if x <= 0 {
			return time.Time{}
		}
		return time.Unix(int64(x), 0)
	case string:
		t, err := time.Parse(time.RFC3339, x)
		if err != nil {
			return time.Time{}
		}
		return t
	}
	return time.Time{}
}

// /sedum:setMethod
// sedum:setMethod {"tier":"owned","record":"prov-2026-4cf2c199","gen":"sha256:9a212ceb82034b3958d818d8c4e87570c1b11c6622df56636be393fc08737d76","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (c *Client) CreateToken(ctx context.Context, bearer string, req TokenRequest) (*TokenView, error) {\n\tvar out TokenView\n\tif err := c.do(ctx, http.MethodPost, \"/tokens\", bearer, req, &out); err != nil {\n\t\treturn nil, tokenErr(err)\n\t}\n\treturn &out, nil\n}","file":"auth/client.go","name":"CreateToken","receiver":"Client"}}
func (c *Client) CreateToken(ctx context.Context, bearer string, req TokenRequest) (*TokenView, error) {
	var out TokenView
	if err := c.do(ctx, http.MethodPost, "/tokens", bearer, req, &out); err != nil {
		return nil, tokenErr(err)
	}
	return &out, nil
}

// /sedum:setMethod
// sedum:setMethod {"tier":"owned","record":"prov-2026-67d8fbff","gen":"sha256:2952c103e8e45566548cb86f46443bb40c5240b1c1027f508a8725a9ff7f1238","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (c *Client) Logout(ctx context.Context, refreshToken string) error {\n\tbody := map[string]string{\"refresh_token\": refreshToken}\n\treturn c.do(ctx, http.MethodPost, \"/auth/logout\", \"\", body, nil)\n}","file":"auth/client.go","name":"Logout","receiver":"Client"}}
func (c *Client) Logout(ctx context.Context, refreshToken string) error {
	body := map[string]string{"refresh_token": refreshToken}
	return c.do(ctx, http.MethodPost, "/auth/logout", "", body, nil)
}

// /sedum:setMethod
// sedum:setMethod {"tier":"owned","record":"prov-2026-69f23c5b","gen":"sha256:cd855b1b497b83b5c00c838140fa4484f8d8fc899c5858fba5900d50bf379626","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (c *Client) RevokeToken(ctx context.Context, bearer string, id string) error {\n\terr := c.do(ctx, http.MethodDelete, \"/tokens/\"+url.PathEscape(id), bearer, nil, nil)\n\treturn tokenErr(err)\n}","file":"auth/client.go","name":"RevokeToken","receiver":"Client"}}
func (c *Client) RevokeToken(ctx context.Context, bearer string, id string) error {
	err := c.do(ctx, http.MethodDelete, "/tokens/"+url.PathEscape(id), bearer, nil, nil)
	return tokenErr(err)
}

// /sedum:setMethod
// sedum:setType {"tier":"owned","record":"prov-2026-6ec75972","gen":"sha256:81d162989d572a9dced946665f9462f186d3f06f283fceef1f86f3de6afd93d8","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type TokenResponse struct {\n\tAccessToken      string `json:\"access_token\"`\n\tExpiresIn        int    `json:\"expires_in\"`\n\tRefreshToken     string `json:\"refresh_token\"`\n\tRefreshExpiresIn int    `json:\"refresh_expires_in\"`\n\tDID              string `json:\"did\"`\n\tHandle           string `json:\"handle\"`\n}","file":"auth/client.go","name":"TokenResponse"}}
type TokenResponse struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int    `json:"expires_in"`
	RefreshToken     string `json:"refresh_token"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
	DID              string `json:"did"`
	Handle           string `json:"handle"`
}

// /sedum:setType
// sedum:setFunction {"tier":"owned","record":"prov-2026-6f0f74ec","gen":"sha256:2816476fbb50c5651fc5c82d92b5c98908bebac5977a96250030a539aba83090","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func CheckServer(raw string) error {\n\tu, err := url.Parse(raw)\n\tif err != nil {\n\t\treturn fmt.Errorf(\"%q is not a valid server URL\", raw)\n\t}\n\tif u.Host == \"\" {\n\t\treturn fmt.Errorf(\"%q is not a valid server URL\", raw)\n\t}\n\tswitch u.Scheme {\n\tcase \"https\":\n\t\treturn nil\n\tcase \"http\":\n\t\tswitch u.Hostname() {\n\t\tcase \"localhost\", \"127.0.0.1\", \"::1\":\n\t\t\treturn nil\n\t\t}\n\t}\n\treturn fmt.Errorf(\"server %s must use https (http is only allowed for localhost)\", u.Host)\n}","file":"auth/client.go","name":"CheckServer"}}
func CheckServer(raw string) error {
	u, err := url.Parse(raw)
	if err != nil {
		return fmt.Errorf("%q is not a valid server URL", raw)
	}
	if u.Host == "" {
		return fmt.Errorf("%q is not a valid server URL", raw)
	}
	switch u.Scheme {
	case "https":
		return nil
	case "http":
		switch u.Hostname() {
		case "localhost", "127.0.0.1", "::1":
			return nil
		}
	}
	return fmt.Errorf("server %s must use https (http is only allowed for localhost)", u.Host)
}

// /sedum:setFunction
// sedum:setMethod {"tier":"owned","record":"prov-2026-87b7ba7b","gen":"sha256:0cb602be28222f7017b0829cf50dc48c7ad99ed6cb36b0ae1b438c53ca3a9008","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (c *Client) do(ctx context.Context, method string, path string, bearer string, in any, out any) error {\n\tif err := CheckServer(c.Server); err != nil {\n\t\treturn err\n\t}\n\tvar reqBody io.Reader\n\tif in != nil {\n\t\tdata, err := json.Marshal(in)\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n\t\treqBody = bytes.NewReader(data)\n\t}\n\tbase := strings.TrimRight(c.Server, \"/\")\n\treq, err := http.NewRequestWithContext(ctx, method, base+path, reqBody)\n\tif err != nil {\n\t\treturn fmt.Errorf(\"%s %s: bad request\", method, path)\n\t}\n\treq.Header.Set(\"Accept\", \"application/json\")\n\tif in != nil {\n\t\treq.Header.Set(\"Content-Type\", \"application/json\")\n\t}\n\tif bearer != \"\" {\n\t\treq.Header.Set(\"Authorization\", \"Bearer \"+bearer)\n\t}\n\tresp, err := c.httpClient().Do(req)\n\tif err != nil {\n\t\tif ctxErr := ctx.Err(); ctxErr != nil {\n\t\t\treturn ctxErr\n\t\t}\n\t\tcause := err\n\t\tvar ue *url.Error\n\t\tif errors.As(err, &ue) {\n\t\t\tcause = ue.Err\n\t\t}\n\t\treturn fmt.Errorf(\"%s %s: %w\", method, path, cause)\n\t}\n\tdefer resp.Body.Close()\n\traw, _ := io.ReadAll(io.LimitReader(resp.Body, 1<<20))\n\tif resp.StatusCode >= 200 && resp.StatusCode <= 299 {\n\t\tif out == nil {\n\t\t\treturn nil\n\t\t}\n\t\tif err := json.Unmarshal(raw, out); err != nil {\n\t\t\treturn fmt.Errorf(\"%s %s: unreadable response\", method, path)\n\t\t}\n\t\treturn nil\n\t}\n\toe := &OAuthError{Status: resp.StatusCode, Code: \"http_\" + strconv.Itoa(resp.StatusCode)}\n\tvar eb struct {\n\t\tError            json.RawMessage `json:\"error\"`\n\t\tErrorDescription string          `json:\"error_description\"`\n\t}\n\tif json.Unmarshal(raw, &eb) == nil {\n\t\tvar s string\n\t\tif len(eb.Error) > 0 && json.Unmarshal(eb.Error, &s) == nil && s != \"\" {\n\t\t\toe.Code = s\n\t\t\toe.Description = eb.ErrorDescription\n\t\t} else {\n\t\t\tvar obj struct {\n\t\t\t\tCode    string `json:\"code\"`\n\t\t\t\tMessage string `json:\"message\"`\n\t\t\t}\n\t\t\tif json.Unmarshal(eb.Error, &obj) == nil && obj.Code != \"\" {\n\t\t\t\toe.Code = obj.Code\n\t\t\t\toe.Description = obj.Message\n\t\t\t}\n\t\t}\n\t}\n\tif n, err := strconv.Atoi(resp.Header.Get(\"Retry-After\")); err == nil {\n\t\toe.RetryAfter = time.Duration(n) * time.Second\n\t}\n\treturn oe\n}","file":"auth/client.go","name":"do","receiver":"Client"}}
func (c *Client) do(ctx context.Context, method string, path string, bearer string, in any, out any) error {
	if err := CheckServer(c.Server); err != nil {
		return err
	}
	var reqBody io.Reader
	if in != nil {
		data, err := json.Marshal(in)
		if err != nil {
			return err
		}
		reqBody = bytes.NewReader(data)
	}
	base := strings.TrimRight(c.Server, "/")
	req, err := http.NewRequestWithContext(ctx, method, base+path, reqBody)
	if err != nil {
		return fmt.Errorf("%s %s: bad request", method, path)
	}
	req.Header.Set("Accept", "application/json")
	if in != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if bearer != "" {
		req.Header.Set("Authorization", "Bearer "+bearer)
	}
	resp, err := c.httpClient().Do(req)
	if err != nil {
		if ctxErr := ctx.Err(); ctxErr != nil {
			return ctxErr
		}
		cause := err
		var ue *url.Error
		if errors.As(err, &ue) {
			cause = ue.Err
		}
		return fmt.Errorf("%s %s: %w", method, path, cause)
	}
	defer resp.Body.Close()
	raw, _ := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		if out == nil {
			return nil
		}
		if err := json.Unmarshal(raw, out); err != nil {
			return fmt.Errorf("%s %s: unreadable response", method, path)
		}
		return nil
	}
	oe := &OAuthError{Status: resp.StatusCode, Code: "http_" + strconv.Itoa(resp.StatusCode)}
	var eb struct {
		Error            json.RawMessage `json:"error"`
		ErrorDescription string          `json:"error_description"`
	}
	if json.Unmarshal(raw, &eb) == nil {
		var s string
		if len(eb.Error) > 0 && json.Unmarshal(eb.Error, &s) == nil && s != "" {
			oe.Code = s
			oe.Description = eb.ErrorDescription
		} else {
			var obj struct {
				Code    string `json:"code"`
				Message string `json:"message"`
			}
			if json.Unmarshal(eb.Error, &obj) == nil && obj.Code != "" {
				oe.Code = obj.Code
				oe.Description = obj.Message
			}
		}
	}
	if n, err := strconv.Atoi(resp.Header.Get("Retry-After")); err == nil {
		oe.RetryAfter = time.Duration(n) * time.Second
	}
	return oe
}

// /sedum:setMethod
// sedum:setMethod {"tier":"owned","record":"prov-2026-8913ed6d","gen":"sha256:147538d3bed4c7f1812d2782ea19cb683cb1e4aa87dddc7da0ad6abc7da8bf9d","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (c *Client) Session(ctx context.Context, bearer string) (*SessionInfo, error) {\n\tvar info SessionInfo\n\terr := c.do(ctx, http.MethodGet, \"/session\", bearer, nil, &info)\n\tif err != nil {\n\t\tvar oe *OAuthError\n\t\tif errors.As(err, &oe) && oe.Status == http.StatusUnauthorized {\n\t\t\treturn nil, ErrUnauthorized\n\t\t}\n\t\treturn nil, err\n\t}\n\treturn &info, nil\n}","file":"auth/client.go","name":"Session","receiver":"Client"}}
func (c *Client) Session(ctx context.Context, bearer string) (*SessionInfo, error) {
	var info SessionInfo
	err := c.do(ctx, http.MethodGet, "/session", bearer, nil, &info)
	if err != nil {
		var oe *OAuthError
		if errors.As(err, &oe) && oe.Status == http.StatusUnauthorized {
			return nil, ErrUnauthorized
		}
		return nil, err
	}
	return &info, nil
}

// /sedum:setMethod
// sedum:setFunction {"tier":"owned","record":"prov-2026-99d873a5","gen":"sha256:b424270ea121f845a3db0d11dafae172975f0e6ebf5fbcb57c534dd7942d625b","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func tokenErr(err error) error {\n\tvar oauthErr *OAuthError\n\tif errors.As(err, &oauthErr) && oauthErr.Status == http.StatusUnauthorized {\n\t\treturn ErrUnauthorized\n\t}\n\treturn err\n}","file":"auth/client.go","name":"tokenErr"}}
func tokenErr(err error) error {
	var oauthErr *OAuthError
	if errors.As(err, &oauthErr) && oauthErr.Status == http.StatusUnauthorized {
		return ErrUnauthorized
	}
	return err
}

// /sedum:setFunction
// sedum:setType {"tier":"owned","record":"prov-2026-b4e29172","gen":"sha256:46b8f72efdcc18667736f3c3568613d3a30a772b02626de5d8ebb6e69f9722dc","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type SessionInfo struct {\n\tDID       string          `json:\"did\"`\n\tHandle    string          `json:\"handle\"`\n\tExpiresAt json.RawMessage `json:\"expires_at\"`\n\tToken     *struct {\n\t\tID     string   `json:\"id\"`\n\t\tName   string   `json:\"name\"`\n\t\tScopes []string `json:\"scopes\"`\n\t} `json:\"token,omitempty\"`\n}","file":"auth/client.go","name":"SessionInfo"}}
type SessionInfo struct {
	DID       string          `json:"did"`
	Handle    string          `json:"handle"`
	ExpiresAt json.RawMessage `json:"expires_at"`
	Token     *struct {
		ID     string   `json:"id"`
		Name   string   `json:"name"`
		Scopes []string `json:"scopes"`
	} `json:"token,omitempty"`
}

// /sedum:setType
// sedum:setMethod {"tier":"owned","record":"prov-2026-c21a212d","gen":"sha256:69f514a66933df2821f4f43695bd9a15566f965c909962022b4d046fa7f3d147","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (c *Client) httpClient() *http.Client {\n\tif c.HTTP != nil {\n\t\treturn c.HTTP\n\t}\n\treturn &http.Client{Timeout: 30 * time.Second}\n}","file":"auth/client.go","name":"httpClient","receiver":"Client"}}
func (c *Client) httpClient() *http.Client {
	if c.HTTP != nil {
		return c.HTTP
	}
	return &http.Client{Timeout: 30 * time.Second}
}

// /sedum:setMethod
// sedum:setType {"tier":"owned","record":"prov-2026-c90585a6","gen":"sha256:2e603ee9399325e28edbee5a4f8781fce4353171e4d63409388cc399377aa810","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type OAuthError struct {\n\tCode        string\n\tDescription string\n\tStatus      int\n\tRetryAfter  time.Duration\n}","file":"auth/client.go","name":"OAuthError"}}
type OAuthError struct {
	Code        string
	Description string
	Status      int
	RetryAfter  time.Duration
}

// /sedum:setType
// sedum:setType {"tier":"owned","record":"prov-2026-e0b27562","gen":"sha256:7c22420b7e10891c3998814fbaacad91039f01265160b16c9d79bd1336e5d65e","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type DeviceCode struct {\n\tDeviceCode              string `json:\"device_code\"`\n\tUserCode                string `json:\"user_code\"`\n\tVerificationURI         string `json:\"verification_uri\"`\n\tVerificationURIComplete string `json:\"verification_uri_complete\"`\n\tExpiresIn               int    `json:\"expires_in\"`\n\tInterval                int    `json:\"interval\"`\n}","file":"auth/client.go","name":"DeviceCode"}}
type DeviceCode struct {
	DeviceCode              string `json:"device_code"`
	UserCode                string `json:"user_code"`
	VerificationURI         string `json:"verification_uri"`
	VerificationURIComplete string `json:"verification_uri_complete"`
	ExpiresIn               int    `json:"expires_in"`
	Interval                int    `json:"interval"`
}

// /sedum:setType
// sedum:setMethod {"tier":"owned","record":"prov-2026-e2440e40","gen":"sha256:22c6468ef066a140fdc09726bd5d5e1d59bae10ea00cf4a24b7d0e268dd2e744","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (c *Client) ListTokens(ctx context.Context, bearer string) ([]TokenView, error) {\n\tvar resp struct {\n\t\tTokens []TokenView `json:\"tokens\"`\n\t}\n\tif err := c.do(ctx, http.MethodGet, \"/tokens\", bearer, nil, &resp); err != nil {\n\t\treturn nil, tokenErr(err)\n\t}\n\treturn resp.Tokens, nil\n}","file":"auth/client.go","name":"ListTokens","receiver":"Client"}}
func (c *Client) ListTokens(ctx context.Context, bearer string) ([]TokenView, error) {
	var resp struct {
		Tokens []TokenView `json:"tokens"`
	}
	if err := c.do(ctx, http.MethodGet, "/tokens", bearer, nil, &resp); err != nil {
		return nil, tokenErr(err)
	}
	return resp.Tokens, nil
}

// /sedum:setMethod
// sedum:setMethod {"tier":"owned","record":"prov-2026-e86f5342","gen":"sha256:6f720ad91ccd1443be390975a8e4dfa40fad41c88c2b355b7e7f409055819f6b","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (c *Client) PollToken(ctx context.Context, dc *DeviceCode, sleep func(context.Context, time.Duration) error) (*TokenResponse, error) {\n\tif sleep == nil {\n\t\tsleep = Wait\n\t}\n\tinterval := time.Duration(dc.Interval) * time.Second\n\tbody := map[string]string{\n\t\t\"grant_type\":  \"urn:ietf:params:oauth:grant-type:device_code\",\n\t\t\"device_code\": dc.DeviceCode,\n\t}\n\tfor {\n\t\tif err := sleep(ctx, interval); err != nil {\n\t\t\treturn nil, err\n\t\t}\n\t\tvar tok TokenResponse\n\t\terr := c.do(ctx, http.MethodPost, \"/device/token\", \"\", body, &tok)\n\t\tif err == nil {\n\t\t\treturn &tok, nil\n\t\t}\n\t\tvar oe *OAuthError\n\t\tif !errors.As(err, &oe) {\n\t\t\treturn nil, err\n\t\t}\n\t\tswitch oe.Code {\n\t\tcase \"authorization_pending\":\n\t\tcase \"slow_down\":\n\t\t\tinterval += 5 * time.Second\n\t\tcase \"access_denied\":\n\t\t\treturn nil, ErrAccessDenied\n\t\tcase \"expired_token\":\n\t\t\treturn nil, ErrDeviceExpired\n\t\tdefault:\n\t\t\treturn nil, err\n\t\t}\n\t}\n}","file":"auth/client.go","name":"PollToken","receiver":"Client"}}
func (c *Client) PollToken(ctx context.Context, dc *DeviceCode, sleep func(context.Context, time.Duration) error) (*TokenResponse, error) {
	if sleep == nil {
		sleep = Wait
	}
	interval := time.Duration(dc.Interval) * time.Second
	body := map[string]string{
		"grant_type":  "urn:ietf:params:oauth:grant-type:device_code",
		"device_code": dc.DeviceCode,
	}
	for {
		if err := sleep(ctx, interval); err != nil {
			return nil, err
		}
		var tok TokenResponse
		err := c.do(ctx, http.MethodPost, "/device/token", "", body, &tok)
		if err == nil {
			return &tok, nil
		}
		var oe *OAuthError
		if !errors.As(err, &oe) {
			return nil, err
		}
		switch oe.Code {
		case "authorization_pending":
		case "slow_down":
			interval += 5 * time.Second
		case "access_denied":
			return nil, ErrAccessDenied
		case "expired_token":
			return nil, ErrDeviceExpired
		default:
			return nil, err
		}
	}
}

// /sedum:setMethod
// sedum:setMethod {"tier":"owned","record":"prov-2026-fead189c","gen":"sha256:a1e20305978db4c0431b367d10de50a24d0888c72492483aa3b05a97b5805fa9","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (c *Client) RequestDeviceCode(ctx context.Context) (*DeviceCode, error) {\n\tbody := map[string]string{\"client_id\": ClientID}\n\tif c.Hostname != \"\" {\n\t\th := c.Hostname\n\t\tif len(h) > 80 {\n\t\t\th = h[:80]\n\t\t}\n\t\tbody[\"hostname\"] = h\n\t}\n\tif c.Provider != \"\" {\n\t\tbody[\"provider\"] = c.Provider\n\t}\n\tvar dc DeviceCode\n\tif err := c.do(ctx, http.MethodPost, \"/device/code\", \"\", body, &dc); err != nil {\n\t\treturn nil, err\n\t}\n\tif dc.DeviceCode == \"\" || dc.UserCode == \"\" {\n\t\treturn nil, errors.New(\"POST /device/code: incomplete response\")\n\t}\n\tif dc.Interval <= 0 {\n\t\tdc.Interval = 5\n\t}\n\treturn &dc, nil\n}","file":"auth/client.go","name":"RequestDeviceCode","receiver":"Client"}}
func (c *Client) RequestDeviceCode(ctx context.Context) (*DeviceCode, error) {
	body := map[string]string{"client_id": ClientID}
	if c.Hostname != "" {
		h := c.Hostname
		if len(h) > 80 {
			h = h[:80]
		}
		body["hostname"] = h
	}
	if c.Provider != "" {
		body["provider"] = c.Provider
	}
	var dc DeviceCode
	if err := c.do(ctx, http.MethodPost, "/device/code", "", body, &dc); err != nil {
		return nil, err
	}
	if dc.DeviceCode == "" || dc.UserCode == "" {
		return nil, errors.New("POST /device/code: incomplete response")
	}
	if dc.Interval <= 0 {
		dc.Interval = 5
	}
	return &dc, nil
}

// /sedum:setMethod
