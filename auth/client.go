package auth

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-03700026","gen":"sha256:a8575efa29a0db9a616cb59f6a86a89236daee7e3c99cfbdad7610ca2d61b2a6","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"auth/client.go","path":"bytes"}}
import "bytes"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-03700026","gen":"sha256:a8575efa29a0db9a616cb59f6a86a89236daee7e3c99cfbdad7610ca2d61b2a6","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"auth/client.go","path":"context"}}
import "context"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-03700026","gen":"sha256:a8575efa29a0db9a616cb59f6a86a89236daee7e3c99cfbdad7610ca2d61b2a6","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"auth/client.go","path":"encoding/json"}}
import "encoding/json"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-03700026","gen":"sha256:a8575efa29a0db9a616cb59f6a86a89236daee7e3c99cfbdad7610ca2d61b2a6","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"auth/client.go","path":"errors"}}
import "errors"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-03700026","gen":"sha256:a8575efa29a0db9a616cb59f6a86a89236daee7e3c99cfbdad7610ca2d61b2a6","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"auth/client.go","path":"fmt"}}
import "fmt"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-03700026","gen":"sha256:a8575efa29a0db9a616cb59f6a86a89236daee7e3c99cfbdad7610ca2d61b2a6","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"auth/client.go","path":"io"}}
import "io"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-03700026","gen":"sha256:a8575efa29a0db9a616cb59f6a86a89236daee7e3c99cfbdad7610ca2d61b2a6","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"auth/client.go","path":"net/http"}}
import "net/http"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-03700026","gen":"sha256:a8575efa29a0db9a616cb59f6a86a89236daee7e3c99cfbdad7610ca2d61b2a6","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"auth/client.go","path":"net/url"}}
import "net/url"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-03700026","gen":"sha256:a8575efa29a0db9a616cb59f6a86a89236daee7e3c99cfbdad7610ca2d61b2a6","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"auth/client.go","path":"strconv"}}
import "strconv"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-03700026","gen":"sha256:a8575efa29a0db9a616cb59f6a86a89236daee7e3c99cfbdad7610ca2d61b2a6","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"auth/client.go","path":"strings"}}
import "strings"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-03700026","gen":"sha256:a8575efa29a0db9a616cb59f6a86a89236daee7e3c99cfbdad7610ca2d61b2a6","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"auth/client.go","path":"time"}}
import "time"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setType {"tier":"owned","record":"prov-2026-03700026","gen":"sha256:a8575efa29a0db9a616cb59f6a86a89236daee7e3c99cfbdad7610ca2d61b2a6","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"// Client talks to the Provisio auth service (provisio-backend docs/DEVICE-FLOW.md).\ntype Client struct {\n\tServer   string\n\tHTTP     *http.Client\n\tHostname string\n\t// Provider is the sign-in hint sent with the device code (\"provisio\" or\n\t// \"bluesky\"); empty sends none.\n\tProvider string\n}","file":"auth/client.go","name":"Client"}}
// Client talks to the Provisio auth service (provisio-backend docs/DEVICE-FLOW.md).
type Client struct {
	Server   string
	HTTP     *http.Client
	Hostname string
	// Provider is the sign-in hint sent with the device code ("provisio" or
	// "bluesky"); empty sends none.
	Provider string
}

// /sedum:setType
// sedum:setType {"tier":"owned","record":"prov-2026-03700026","gen":"sha256:a8575efa29a0db9a616cb59f6a86a89236daee7e3c99cfbdad7610ca2d61b2a6","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"// DeviceCode is the response of POST /device/code.\ntype DeviceCode struct {\n\tDeviceCode              string `json:\"device_code\"`\n\tUserCode                string `json:\"user_code\"`\n\tVerificationURI         string `json:\"verification_uri\"`\n\tVerificationURIComplete string `json:\"verification_uri_complete\"`\n\tExpiresIn               int    `json:\"expires_in\"`\n\tInterval                int    `json:\"interval\"`\n}","file":"auth/client.go","name":"DeviceCode"}}
// DeviceCode is the response of POST /device/code.
type DeviceCode struct {
	DeviceCode              string `json:"device_code"`
	UserCode                string `json:"user_code"`
	VerificationURI         string `json:"verification_uri"`
	VerificationURIComplete string `json:"verification_uri_complete"`
	ExpiresIn               int    `json:"expires_in"`
	Interval                int    `json:"interval"`
}

// /sedum:setType
// sedum:setType {"tier":"owned","record":"prov-2026-03700026","gen":"sha256:a8575efa29a0db9a616cb59f6a86a89236daee7e3c99cfbdad7610ca2d61b2a6","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"// TokenResponse is the response of POST /device/token and POST /auth/refresh.\ntype TokenResponse struct {\n\tAccessToken      string `json:\"access_token\"`\n\tExpiresIn        int    `json:\"expires_in\"`\n\tRefreshToken     string `json:\"refresh_token\"`\n\tRefreshExpiresIn int    `json:\"refresh_expires_in\"`\n\tDID              string `json:\"did\"`\n\tHandle           string `json:\"handle\"`\n}","file":"auth/client.go","name":"TokenResponse"}}
// TokenResponse is the response of POST /device/token and POST /auth/refresh.
type TokenResponse struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int    `json:"expires_in"`
	RefreshToken     string `json:"refresh_token"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
	DID              string `json:"did"`
	Handle           string `json:"handle"`
}

// /sedum:setType
// sedum:setType {"tier":"owned","record":"prov-2026-03700026","gen":"sha256:a8575efa29a0db9a616cb59f6a86a89236daee7e3c99cfbdad7610ca2d61b2a6","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"// SessionInfo is the response of GET /session. expires_at is unix seconds (an\n// RFC 3339 string is accepted too) and absent for a token that never expires.\ntype SessionInfo struct {\n\tDID       string          `json:\"did\"`\n\tHandle    string          `json:\"handle\"`\n\tExpiresAt json.RawMessage `json:\"expires_at\"`\n\tToken     *struct {\n\t\tID     string   `json:\"id\"`\n\t\tName   string   `json:\"name\"`\n\t\tScopes []string `json:\"scopes\"`\n\t} `json:\"token,omitempty\"`\n}","file":"auth/client.go","name":"SessionInfo"}}
// SessionInfo is the response of GET /session. expires_at is unix seconds (an
// RFC 3339 string is accepted too) and absent for a token that never expires.
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
// sedum:setMethod {"tier":"owned","record":"prov-2026-03700026","gen":"sha256:a8575efa29a0db9a616cb59f6a86a89236daee7e3c99cfbdad7610ca2d61b2a6","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"// Expiry parses ExpiresAt; the zero time when absent, null or unparsable.\nfunc (s *SessionInfo) Expiry() time.Time {\n\tvar secs float64\n\tif json.Unmarshal(s.ExpiresAt, &secs) == nil && secs > 0 {\n\t\treturn time.Unix(int64(secs), 0)\n\t}\n\tvar str string\n\tif json.Unmarshal(s.ExpiresAt, &str) == nil {\n\t\tt, _ := time.Parse(time.RFC3339, str)\n\t\treturn t\n\t}\n\treturn time.Time{}\n}","file":"auth/client.go","name":"Expiry","receiver":"SessionInfo"}}
// Expiry parses ExpiresAt; the zero time when absent, null or unparsable.
func (s *SessionInfo) Expiry() time.Time {
	var secs float64
	if json.Unmarshal(s.ExpiresAt, &secs) == nil && secs > 0 {
		return time.Unix(int64(secs), 0)
	}
	var str string
	if json.Unmarshal(s.ExpiresAt, &str) == nil {
		t, _ := time.Parse(time.RFC3339, str)
		return t
	}
	return time.Time{}
}

// /sedum:setMethod
// sedum:setType {"tier":"owned","record":"prov-2026-03700026","gen":"sha256:a8575efa29a0db9a616cb59f6a86a89236daee7e3c99cfbdad7610ca2d61b2a6","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"// OAuthError is an OAuth-style failure from /device/* or /auth/*. It never\n// carries a request body, so it cannot leak a token.\ntype OAuthError struct {\n\tCode        string\n\tDescription string\n\tStatus      int\n\tRetryAfter  time.Duration\n}","file":"auth/client.go","name":"OAuthError"}}
// OAuthError is an OAuth-style failure from /device/* or /auth/*. It never
// carries a request body, so it cannot leak a token.
type OAuthError struct {
	Code        string
	Description string
	Status      int
	RetryAfter  time.Duration
}

// /sedum:setType
// sedum:setMethod {"tier":"owned","record":"prov-2026-03700026","gen":"sha256:a8575efa29a0db9a616cb59f6a86a89236daee7e3c99cfbdad7610ca2d61b2a6","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (e *OAuthError) Error() string {\n\tif e.Description != \"\" {\n\t\treturn fmt.Sprintf(\"%s (HTTP %d): %s\", e.Code, e.Status, e.Description)\n\t}\n\treturn fmt.Sprintf(\"%s (HTTP %d)\", e.Code, e.Status)\n}","file":"auth/client.go","name":"Error","receiver":"OAuthError"}}
func (e *OAuthError) Error() string {
	if e.Description != "" {
		return fmt.Sprintf("%s (HTTP %d): %s", e.Code, e.Status, e.Description)
	}
	return fmt.Sprintf("%s (HTTP %d)", e.Code, e.Status)
}

// /sedum:setMethod
// sedum:setFunction {"tier":"owned","record":"prov-2026-03700026","gen":"sha256:a8575efa29a0db9a616cb59f6a86a89236daee7e3c99cfbdad7610ca2d61b2a6","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"// CheckServer accepts https URLs, and http only for a loopback host.\nfunc CheckServer(raw string) error {\n\tu, err := url.Parse(raw)\n\tif err != nil || u.Host == \"\" {\n\t\treturn fmt.Errorf(\"%q is not a valid server URL\", raw)\n\t}\n\tswitch u.Scheme {\n\tcase \"https\":\n\t\treturn nil\n\tcase \"http\":\n\t\th := u.Hostname()\n\t\tif h == \"localhost\" || h == \"127.0.0.1\" || h == \"::1\" {\n\t\t\treturn nil\n\t\t}\n\t}\n\treturn fmt.Errorf(\"server %s must use https (http is only allowed for localhost)\", u.Host)\n}","file":"auth/client.go","name":"CheckServer"}}
// CheckServer accepts https URLs, and http only for a loopback host.
func CheckServer(raw string) error {
	u, err := url.Parse(raw)
	if err != nil || u.Host == "" {
		return fmt.Errorf("%q is not a valid server URL", raw)
	}
	switch u.Scheme {
	case "https":
		return nil
	case "http":
		h := u.Hostname()
		if h == "localhost" || h == "127.0.0.1" || h == "::1" {
			return nil
		}
	}
	return fmt.Errorf("server %s must use https (http is only allowed for localhost)", u.Host)
}

// /sedum:setFunction
// sedum:setMethod {"tier":"owned","record":"prov-2026-03700026","gen":"sha256:a8575efa29a0db9a616cb59f6a86a89236daee7e3c99cfbdad7610ca2d61b2a6","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (c *Client) httpClient() *http.Client {\n\tif c.HTTP != nil {\n\t\treturn c.HTTP\n\t}\n\treturn &http.Client{Timeout: 30 * time.Second}\n}","file":"auth/client.go","name":"httpClient","receiver":"Client"}}
func (c *Client) httpClient() *http.Client {
	if c.HTTP != nil {
		return c.HTTP
	}
	return &http.Client{Timeout: 30 * time.Second}
}

// /sedum:setMethod
// sedum:setMethod {"tier":"owned","record":"prov-2026-03700026","gen":"sha256:a8575efa29a0db9a616cb59f6a86a89236daee7e3c99cfbdad7610ca2d61b2a6","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"// do sends one request. Errors name only the method, the path and the status.\nfunc (c *Client) do(ctx context.Context, method, path, bearer string, in, out any) error {\n\tif err := CheckServer(c.Server); err != nil {\n\t\treturn err\n\t}\n\tvar body io.Reader\n\tif in != nil {\n\t\tdata, err := json.Marshal(in)\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n\t\tbody = bytes.NewReader(data)\n\t}\n\treq, err := http.NewRequestWithContext(ctx, method, strings.TrimRight(c.Server, \"/\")+path, body)\n\tif err != nil {\n\t\treturn fmt.Errorf(\"%s %s: bad request\", method, path)\n\t}\n\treq.Header.Set(\"Accept\", \"application/json\")\n\tif in != nil {\n\t\treq.Header.Set(\"Content-Type\", \"application/json\")\n\t}\n\tif bearer != \"\" {\n\t\treq.Header.Set(\"Authorization\", \"Bearer \"+bearer)\n\t}\n\tresp, err := c.httpClient().Do(req)\n\tif err != nil {\n\t\tif ctx.Err() != nil {\n\t\t\treturn ctx.Err()\n\t\t}\n\t\tvar ue *url.Error\n\t\tif errors.As(err, &ue) {\n\t\t\terr = ue.Err\n\t\t}\n\t\treturn fmt.Errorf(\"%s %s: %w\", method, path, err)\n\t}\n\tdefer resp.Body.Close()\n\tdata, _ := io.ReadAll(io.LimitReader(resp.Body, 1<<20))\n\tif resp.StatusCode >= 200 && resp.StatusCode < 300 {\n\t\tif out == nil {\n\t\t\treturn nil\n\t\t}\n\t\tif err := json.Unmarshal(data, out); err != nil {\n\t\t\treturn fmt.Errorf(\"%s %s: unreadable response\", method, path)\n\t\t}\n\t\treturn nil\n\t}\n\toe := &OAuthError{Status: resp.StatusCode, Code: \"http_\" + strconv.Itoa(resp.StatusCode)}\n\tvar eb struct {\n\t\tError       json.RawMessage `json:\"error\"`\n\t\tDescription string          `json:\"error_description\"`\n\t}\n\tif json.Unmarshal(data, &eb) == nil {\n\t\tvar code string\n\t\tvar nested struct {\n\t\t\tCode    string `json:\"code\"`\n\t\t\tMessage string `json:\"message\"`\n\t\t}\n\t\tif json.Unmarshal(eb.Error, &code) == nil && code != \"\" {\n\t\t\toe.Code, oe.Description = code, eb.Description\n\t\t} else if json.Unmarshal(eb.Error, &nested) == nil && nested.Code != \"\" {\n\t\t\toe.Code, oe.Description = nested.Code, nested.Message\n\t\t}\n\t}\n\tif secs, err := strconv.Atoi(resp.Header.Get(\"Retry-After\")); err == nil {\n\t\toe.RetryAfter = time.Duration(secs) * time.Second\n\t}\n\treturn oe\n}","file":"auth/client.go","name":"do","receiver":"Client"}}
// do sends one request. Errors name only the method, the path and the status.
func (c *Client) do(ctx context.Context, method, path, bearer string, in, out any) error {
	if err := CheckServer(c.Server); err != nil {
		return err
	}
	var body io.Reader
	if in != nil {
		data, err := json.Marshal(in)
		if err != nil {
			return err
		}
		body = bytes.NewReader(data)
	}
	req, err := http.NewRequestWithContext(ctx, method, strings.TrimRight(c.Server, "/")+path, body)
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
		if ctx.Err() != nil {
			return ctx.Err()
		}
		var ue *url.Error
		if errors.As(err, &ue) {
			err = ue.Err
		}
		return fmt.Errorf("%s %s: %w", method, path, err)
	}
	defer resp.Body.Close()
	data, _ := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		if out == nil {
			return nil
		}
		if err := json.Unmarshal(data, out); err != nil {
			return fmt.Errorf("%s %s: unreadable response", method, path)
		}
		return nil
	}
	oe := &OAuthError{Status: resp.StatusCode, Code: "http_" + strconv.Itoa(resp.StatusCode)}
	var eb struct {
		Error       json.RawMessage `json:"error"`
		Description string          `json:"error_description"`
	}
	if json.Unmarshal(data, &eb) == nil {
		var code string
		var nested struct {
			Code    string `json:"code"`
			Message string `json:"message"`
		}
		if json.Unmarshal(eb.Error, &code) == nil && code != "" {
			oe.Code, oe.Description = code, eb.Description
		} else if json.Unmarshal(eb.Error, &nested) == nil && nested.Code != "" {
			oe.Code, oe.Description = nested.Code, nested.Message
		}
	}
	if secs, err := strconv.Atoi(resp.Header.Get("Retry-After")); err == nil {
		oe.RetryAfter = time.Duration(secs) * time.Second
	}
	return oe
}

// /sedum:setMethod
// sedum:setMethod {"tier":"owned","record":"prov-2026-03700026","gen":"sha256:a8575efa29a0db9a616cb59f6a86a89236daee7e3c99cfbdad7610ca2d61b2a6","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"// RequestDeviceCode starts a device sign-in.\nfunc (c *Client) RequestDeviceCode(ctx context.Context) (*DeviceCode, error) {\n\tin := map[string]string{\"client_id\": ClientID}\n\tif h := c.Hostname; h != \"\" {\n\t\tif len(h) > 80 {\n\t\t\th = h[:80]\n\t\t}\n\t\tin[\"hostname\"] = h\n\t}\n\tif c.Provider != \"\" {\n\t\tin[\"provider\"] = c.Provider\n\t}\n\tvar dc DeviceCode\n\tif err := c.do(ctx, http.MethodPost, \"/device/code\", \"\", in, &dc); err != nil {\n\t\treturn nil, err\n\t}\n\tif dc.DeviceCode == \"\" || dc.UserCode == \"\" {\n\t\treturn nil, errors.New(\"POST /device/code: incomplete response\")\n\t}\n\tif dc.Interval <= 0 {\n\t\tdc.Interval = 5\n\t}\n\treturn &dc, nil\n}","file":"auth/client.go","name":"RequestDeviceCode","receiver":"Client"}}
// RequestDeviceCode starts a device sign-in.
func (c *Client) RequestDeviceCode(ctx context.Context) (*DeviceCode, error) {
	in := map[string]string{"client_id": ClientID}
	if h := c.Hostname; h != "" {
		if len(h) > 80 {
			h = h[:80]
		}
		in["hostname"] = h
	}
	if c.Provider != "" {
		in["provider"] = c.Provider
	}
	var dc DeviceCode
	if err := c.do(ctx, http.MethodPost, "/device/code", "", in, &dc); err != nil {
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
// sedum:setFunction {"tier":"owned","record":"prov-2026-03700026","gen":"sha256:a8575efa29a0db9a616cb59f6a86a89236daee7e3c99cfbdad7610ca2d61b2a6","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"// Wait sleeps for d or until ctx ends.\nfunc Wait(ctx context.Context, d time.Duration) error {\n\tt := time.NewTimer(d)\n\tdefer t.Stop()\n\tselect {\n\tcase <-ctx.Done():\n\t\treturn ctx.Err()\n\tcase <-t.C:\n\t\treturn nil\n\t}\n}","file":"auth/client.go","name":"Wait"}}
// Wait sleeps for d or until ctx ends.
func Wait(ctx context.Context, d time.Duration) error {
	t := time.NewTimer(d)
	defer t.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-t.C:
		return nil
	}
}

// /sedum:setFunction
// sedum:setMethod {"tier":"owned","record":"prov-2026-03700026","gen":"sha256:a8575efa29a0db9a616cb59f6a86a89236daee7e3c99cfbdad7610ca2d61b2a6","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"// PollToken polls POST /device/token at the server's interval until the user\n// approves, denies, or the code expires. sleep is Wait unless a test replaces it.\nfunc (c *Client) PollToken(ctx context.Context, dc *DeviceCode, sleep func(context.Context, time.Duration) error) (*TokenResponse, error) {\n\tif sleep == nil {\n\t\tsleep = Wait\n\t}\n\tinterval := time.Duration(dc.Interval) * time.Second\n\tin := map[string]string{\"grant_type\": \"urn:ietf:params:oauth:grant-type:device_code\", \"device_code\": dc.DeviceCode}\n\tfor {\n\t\tif err := sleep(ctx, interval); err != nil {\n\t\t\treturn nil, err\n\t\t}\n\t\tvar tr TokenResponse\n\t\terr := c.do(ctx, http.MethodPost, \"/device/token\", \"\", in, &tr)\n\t\tif err == nil {\n\t\t\treturn &tr, nil\n\t\t}\n\t\tvar oe *OAuthError\n\t\tif !errors.As(err, &oe) {\n\t\t\treturn nil, err\n\t\t}\n\t\tswitch oe.Code {\n\t\tcase \"authorization_pending\":\n\t\tcase \"slow_down\":\n\t\t\tinterval += 5 * time.Second\n\t\tcase \"access_denied\":\n\t\t\treturn nil, ErrAccessDenied\n\t\tcase \"expired_token\":\n\t\t\treturn nil, ErrDeviceExpired\n\t\tdefault:\n\t\t\treturn nil, err\n\t\t}\n\t}\n}","file":"auth/client.go","name":"PollToken","receiver":"Client"}}
// PollToken polls POST /device/token at the server's interval until the user
// approves, denies, or the code expires. sleep is Wait unless a test replaces it.
func (c *Client) PollToken(ctx context.Context, dc *DeviceCode, sleep func(context.Context, time.Duration) error) (*TokenResponse, error) {
	if sleep == nil {
		sleep = Wait
	}
	interval := time.Duration(dc.Interval) * time.Second
	in := map[string]string{"grant_type": "urn:ietf:params:oauth:grant-type:device_code", "device_code": dc.DeviceCode}
	for {
		if err := sleep(ctx, interval); err != nil {
			return nil, err
		}
		var tr TokenResponse
		err := c.do(ctx, http.MethodPost, "/device/token", "", in, &tr)
		if err == nil {
			return &tr, nil
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
// sedum:setMethod {"tier":"owned","record":"prov-2026-03700026","gen":"sha256:a8575efa29a0db9a616cb59f6a86a89236daee7e3c99cfbdad7610ca2d61b2a6","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"// Refresh trades a refresh token for a new pair. Any invalid_grant is ErrSessionExpired.\nfunc (c *Client) Refresh(ctx context.Context, refreshToken string) (*TokenResponse, error) {\n\tvar tr TokenResponse\n\terr := c.do(ctx, http.MethodPost, \"/auth/refresh\", \"\", map[string]string{\"refresh_token\": refreshToken}, &tr)\n\tvar oe *OAuthError\n\tif errors.As(err, &oe) && (oe.Code == \"invalid_grant\" || oe.Status == http.StatusUnauthorized) {\n\t\treturn nil, ErrSessionExpired\n\t}\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\treturn &tr, nil\n}","file":"auth/client.go","name":"Refresh","receiver":"Client"}}
// Refresh trades a refresh token for a new pair. Any invalid_grant is ErrSessionExpired.
func (c *Client) Refresh(ctx context.Context, refreshToken string) (*TokenResponse, error) {
	var tr TokenResponse
	err := c.do(ctx, http.MethodPost, "/auth/refresh", "", map[string]string{"refresh_token": refreshToken}, &tr)
	var oe *OAuthError
	if errors.As(err, &oe) && (oe.Code == "invalid_grant" || oe.Status == http.StatusUnauthorized) {
		return nil, ErrSessionExpired
	}
	if err != nil {
		return nil, err
	}
	return &tr, nil
}

// /sedum:setMethod
// sedum:setMethod {"tier":"owned","record":"prov-2026-03700026","gen":"sha256:a8575efa29a0db9a616cb59f6a86a89236daee7e3c99cfbdad7610ca2d61b2a6","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"// Logout revokes the refresh family; the server answers 200 for unknown tokens too.\nfunc (c *Client) Logout(ctx context.Context, refreshToken string) error {\n\treturn c.do(ctx, http.MethodPost, \"/auth/logout\", \"\", map[string]string{\"refresh_token\": refreshToken}, nil)\n}","file":"auth/client.go","name":"Logout","receiver":"Client"}}
// Logout revokes the refresh family; the server answers 200 for unknown tokens too.
func (c *Client) Logout(ctx context.Context, refreshToken string) error {
	return c.do(ctx, http.MethodPost, "/auth/logout", "", map[string]string{"refresh_token": refreshToken}, nil)
}

// /sedum:setMethod
// sedum:setMethod {"tier":"owned","record":"prov-2026-03700026","gen":"sha256:a8575efa29a0db9a616cb59f6a86a89236daee7e3c99cfbdad7610ca2d61b2a6","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"// Session is GET /session. A 401 is ErrUnauthorized.\nfunc (c *Client) Session(ctx context.Context, bearer string) (*SessionInfo, error) {\n\tvar si SessionInfo\n\terr := c.do(ctx, http.MethodGet, \"/session\", bearer, nil, &si)\n\tvar oe *OAuthError\n\tif errors.As(err, &oe) && oe.Status == http.StatusUnauthorized {\n\t\treturn nil, ErrUnauthorized\n\t}\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\treturn &si, nil\n}","file":"auth/client.go","name":"Session","receiver":"Client"}}
// Session is GET /session. A 401 is ErrUnauthorized.
func (c *Client) Session(ctx context.Context, bearer string) (*SessionInfo, error) {
	var si SessionInfo
	err := c.do(ctx, http.MethodGet, "/session", bearer, nil, &si)
	var oe *OAuthError
	if errors.As(err, &oe) && oe.Status == http.StatusUnauthorized {
		return nil, ErrUnauthorized
	}
	if err != nil {
		return nil, err
	}
	return &si, nil
}

// /sedum:setMethod
// sedum:setType {"tier":"owned","record":"prov-2026-03700026","gen":"sha256:a8575efa29a0db9a616cb59f6a86a89236daee7e3c99cfbdad7610ca2d61b2a6","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"// TokenRequest is the body of POST /tokens.\ntype TokenRequest struct {\n\tName          string   `json:\"name\"`\n\tScopes        []string `json:\"scopes\"`\n\tExpiresInDays int      `json:\"expires_in_days,omitempty\"`\n}","file":"auth/client.go","name":"TokenRequest"}}
// TokenRequest is the body of POST /tokens.
type TokenRequest struct {
	Name          string   `json:"name"`
	Scopes        []string `json:"scopes"`
	ExpiresInDays int      `json:"expires_in_days,omitempty"`
}

// /sedum:setType
// sedum:setType {"tier":"owned","record":"prov-2026-03700026","gen":"sha256:a8575efa29a0db9a616cb59f6a86a89236daee7e3c99cfbdad7610ca2d61b2a6","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"// TokenView is a personal token as the service shows it. Token, the secret, is\n// set only in the answer to CreateToken; the service never returns it again.\ntype TokenView struct {\n\tID         string   `json:\"id\"`\n\tName       string   `json:\"name\"`\n\tScopes     []string `json:\"scopes\"`\n\tCreatedAt  string   `json:\"created_at\"`\n\tLastUsedAt string   `json:\"last_used_at\"`\n\tExpiresAt  string   `json:\"expires_at\"`\n\tToken      string   `json:\"token,omitempty\"`\n}","file":"auth/client.go","name":"TokenView"}}
// TokenView is a personal token as the service shows it. Token, the secret, is
// set only in the answer to CreateToken; the service never returns it again.
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
// sedum:setFunction {"tier":"owned","record":"prov-2026-03700026","gen":"sha256:a8575efa29a0db9a616cb59f6a86a89236daee7e3c99cfbdad7610ca2d61b2a6","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"// tokenErr maps a 401 to ErrUnauthorized, as Session does, so the caller can\n// refresh once; other errors pass through.\nfunc tokenErr(err error) error {\n\tvar oe *OAuthError\n\tif errors.As(err, &oe) && oe.Status == http.StatusUnauthorized {\n\t\treturn ErrUnauthorized\n\t}\n\treturn err\n}","file":"auth/client.go","name":"tokenErr"}}
// tokenErr maps a 401 to ErrUnauthorized, as Session does, so the caller can
// refresh once; other errors pass through.
func tokenErr(err error) error {
	var oe *OAuthError
	if errors.As(err, &oe) && oe.Status == http.StatusUnauthorized {
		return ErrUnauthorized
	}
	return err
}

// /sedum:setFunction
// sedum:setMethod {"tier":"owned","record":"prov-2026-03700026","gen":"sha256:a8575efa29a0db9a616cb59f6a86a89236daee7e3c99cfbdad7610ca2d61b2a6","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"// CreateToken is POST /tokens. The secret is in the answer only.\nfunc (c *Client) CreateToken(ctx context.Context, bearer string, req TokenRequest) (*TokenView, error) {\n\tvar tv TokenView\n\tif err := c.do(ctx, http.MethodPost, \"/tokens\", bearer, req, &tv); err != nil {\n\t\treturn nil, tokenErr(err)\n\t}\n\treturn &tv, nil\n}","file":"auth/client.go","name":"CreateToken","receiver":"Client"}}
// CreateToken is POST /tokens. The secret is in the answer only.
func (c *Client) CreateToken(ctx context.Context, bearer string, req TokenRequest) (*TokenView, error) {
	var tv TokenView
	if err := c.do(ctx, http.MethodPost, "/tokens", bearer, req, &tv); err != nil {
		return nil, tokenErr(err)
	}
	return &tv, nil
}

// /sedum:setMethod
// sedum:setMethod {"tier":"owned","record":"prov-2026-03700026","gen":"sha256:a8575efa29a0db9a616cb59f6a86a89236daee7e3c99cfbdad7610ca2d61b2a6","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"// ListTokens is GET /tokens: metadata only, never a secret.\nfunc (c *Client) ListTokens(ctx context.Context, bearer string) ([]TokenView, error) {\n\tvar out struct {\n\t\tTokens []TokenView `json:\"tokens\"`\n\t}\n\tif err := c.do(ctx, http.MethodGet, \"/tokens\", bearer, nil, &out); err != nil {\n\t\treturn nil, tokenErr(err)\n\t}\n\treturn out.Tokens, nil\n}","file":"auth/client.go","name":"ListTokens","receiver":"Client"}}
// ListTokens is GET /tokens: metadata only, never a secret.
func (c *Client) ListTokens(ctx context.Context, bearer string) ([]TokenView, error) {
	var out struct {
		Tokens []TokenView `json:"tokens"`
	}
	if err := c.do(ctx, http.MethodGet, "/tokens", bearer, nil, &out); err != nil {
		return nil, tokenErr(err)
	}
	return out.Tokens, nil
}

// /sedum:setMethod
// sedum:setMethod {"tier":"owned","record":"prov-2026-03700026","gen":"sha256:a8575efa29a0db9a616cb59f6a86a89236daee7e3c99cfbdad7610ca2d61b2a6","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"// RevokeToken is DELETE /tokens/:id; an unknown id is a 404 OAuthError.\nfunc (c *Client) RevokeToken(ctx context.Context, bearer, id string) error {\n\treturn tokenErr(c.do(ctx, http.MethodDelete, \"/tokens/\"+url.PathEscape(id), bearer, nil, nil))\n}","file":"auth/client.go","name":"RevokeToken","receiver":"Client"}}
// RevokeToken is DELETE /tokens/:id; an unknown id is a 404 OAuthError.
func (c *Client) RevokeToken(ctx context.Context, bearer, id string) error {
	return tokenErr(c.do(ctx, http.MethodDelete, "/tokens/"+url.PathEscape(id), bearer, nil, nil))
}

// /sedum:setMethod
