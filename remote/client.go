package remote

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-0b19cb1a","gen":"sha256:b2197b95c6a31e21220ca35d7a537f5568704e1c9bf8642097d00bdfff11d9e1","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"remote/client.go","path":"context"}}
import "context"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-0b19cb1a","gen":"sha256:b2197b95c6a31e21220ca35d7a537f5568704e1c9bf8642097d00bdfff11d9e1","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"remote/client.go","path":"net/http"}}
import "net/http"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-0b19cb1a","gen":"sha256:b2197b95c6a31e21220ca35d7a537f5568704e1c9bf8642097d00bdfff11d9e1","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"remote/client.go","path":"net/url"}}
import "net/url"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-381337f1","gen":"sha256:9dd05c47c6bd93f33010b04f114b6a91439be52092c6f626c7fdd666c1a0a869","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"remote/client.go","path":"errors"}}
import "errors"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-587ecb9c","gen":"sha256:edaba77f9e89db7a15f2ee45022f534fee87072b51ead5efd2e6c8d7b1177ee8","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"remote/client.go","path":"time"}}
import "time"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-5d43c310","gen":"sha256:b1a323f25ef0e639107c67b5239058998f8dd0acb588ffa2057c982ad2ce61c0","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"remote/client.go","path":"github.com/livecodelife/provenance/record"}}
import "github.com/livecodelife/provenance/record"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-7987372d","gen":"sha256:419562ebad626cdaadfaf3572533bcc8bc8bc798e5fa1a1d4ec4e3fb890ee139","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"remote/client.go","path":"bytes"}}
import "bytes"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-7987372d","gen":"sha256:419562ebad626cdaadfaf3572533bcc8bc8bc798e5fa1a1d4ec4e3fb890ee139","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"remote/client.go","path":"encoding/json"}}
import "encoding/json"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-7987372d","gen":"sha256:419562ebad626cdaadfaf3572533bcc8bc8bc798e5fa1a1d4ec4e3fb890ee139","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"remote/client.go","path":"io"}}
import "io"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-7987372d","gen":"sha256:419562ebad626cdaadfaf3572533bcc8bc8bc798e5fa1a1d4ec4e3fb890ee139","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"remote/client.go","path":"strings"}}
import "strings"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-7987372d","gen":"sha256:419562ebad626cdaadfaf3572533bcc8bc8bc798e5fa1a1d4ec4e3fb890ee139","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"remote/client.go","path":"fmt"}}
import "fmt"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-7987372d","gen":"sha256:419562ebad626cdaadfaf3572533bcc8bc8bc798e5fa1a1d4ec4e3fb890ee139","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"remote/client.go","path":"github.com/livecodelife/provenance/auth"}}
import "github.com/livecodelife/provenance/auth"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setMethod {"tier":"owned","record":"prov-2026-0b19cb1a","gen":"sha256:b2197b95c6a31e21220ca35d7a537f5568704e1c9bf8642097d00bdfff11d9e1","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (c *Client) PutVersion(ctx context.Context, id string, req PutVersionRequest) (*Head, *Version, error) {\n\tvar out struct {\n\t\tHead    Head    `json:\"head\"`\n\t\tVersion Version `json:\"version\"`\n\t}\n\tif err := c.do(ctx, http.MethodPut, \"/records/\"+url.PathEscape(id)+\"/versions\", req, &out); err != nil {\n\t\treturn nil, nil, err\n\t}\n\tout.Head.ID = id\n\treturn &out.Head, &out.Version, nil\n}","file":"remote/client.go","name":"PutVersion","receiver":"Client"}}
func (c *Client) PutVersion(ctx context.Context, id string, req PutVersionRequest) (*Head, *Version, error) {
	var out struct {
		Head    Head    `json:"head"`
		Version Version `json:"version"`
	}
	if err := c.do(ctx, http.MethodPut, "/records/"+url.PathEscape(id)+"/versions", req, &out); err != nil {
		return nil, nil, err
	}
	out.Head.ID = id
	return &out.Head, &out.Version, nil
}

// /sedum:setMethod
// sedum:setConst {"tier":"owned","record":"prov-2026-112298ac","gen":"sha256:a12ffb8bb6f728524edc859488f862c96b11ff7e21a8d3aefc27d06e474098af","pkg":"sha256:51be1c18998d63cc558c253e211a6bd188d96cb251dad68c6ea88e03c2d5cc4c","kwargs":{"code":"const DefaultBase = \"https://api.provisio.dev/v1\"","file":"remote/client.go","name":"DefaultBase"}}
const DefaultBase = "https://api.provisio.dev/v1"

// /sedum:setConst
// sedum:setVar {"tier":"owned","record":"prov-2026-381337f1","gen":"sha256:9dd05c47c6bd93f33010b04f114b6a91439be52092c6f626c7fdd666c1a0a869","pkg":"sha256:c7b9d9ac1655e8211b8b2ca8c49cf34b8d510d8d823d4f5eb60ec7e1018aa2eb","kwargs":{"code":"var (\n\tErrNotPublished = errors.New(\"record is not published\")\n\tErrUnsupported  = errors.New(\"not supported by the Provisio store: a remote record changes only by a new version\")\n)","file":"remote/client.go","name":"ErrNotPublished"}}
var (
	ErrNotPublished = errors.New("record is not published")
	ErrUnsupported  = errors.New("not supported by the Provisio store: a remote record changes only by a new version")
)

// /sedum:setVar
// sedum:setMethod {"tier":"owned","record":"prov-2026-46529ebb","gen":"sha256:04280758ac05988721c1c14e54323f6030b1d97e31573b8d4ee8fb02b74d8305","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (e *ConflictError) Error() string {\n\treturn e.APIError.Error()\n}","file":"remote/client.go","name":"Error","receiver":"ConflictError"}}
func (e *ConflictError) Error() string {
	return e.APIError.Error()
}

// /sedum:setMethod
// sedum:setType {"tier":"owned","record":"prov-2026-481edfa2","gen":"sha256:1135fa1101e966ff94e3e13babc5c2aa59239954e42cc738bca1525451b5e5c2","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type Head struct {\n\trecord.PublishedHead\n\tID  string `json:\"id\"`\n\tURI string `json:\"uri\"`\n\tCID string `json:\"cid\"`\n}","file":"remote/client.go","name":"Head"}}
type Head struct {
	record.PublishedHead
	ID  string `json:"id"`
	URI string `json:"uri"`
	CID string `json:"cid"`
}

// /sedum:setType
// sedum:setMethod {"tier":"owned","record":"prov-2026-587ecb9c","gen":"sha256:edaba77f9e89db7a15f2ee45022f534fee87072b51ead5efd2e6c8d7b1177ee8","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (c *Client) httpClient() *http.Client {\n\tif c.HTTP != nil {\n\t\treturn c.HTTP\n\t}\n\treturn &http.Client{Timeout: 60 * time.Second}\n}","file":"remote/client.go","name":"httpClient","receiver":"Client"}}
func (c *Client) httpClient() *http.Client {
	if c.HTTP != nil {
		return c.HTTP
	}
	return &http.Client{Timeout: 60 * time.Second}
}

// /sedum:setMethod
// sedum:setType {"tier":"owned","record":"prov-2026-5d43c310","gen":"sha256:b1a323f25ef0e639107c67b5239058998f8dd0acb588ffa2057c982ad2ce61c0","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type PutVersionRequest struct {\n\tExpectedHeadCid *string                 `json:\"expectedHeadCid\"`\n\tVersion         int                     `json:\"version\"`\n\tContentHash     string                  `json:\"contentHash\"`\n\tGenHash         string                  `json:\"genHash,omitempty\"`\n\tSpecHashes      []string                `json:\"specHashes\"`\n\tContent         record.PublishedContent `json:\"content\"`\n\tStatus          string                  `json:\"status,omitempty\"`\n\tReason          string                  `json:\"reason,omitempty\"`\n\tReasonKind      string                  `json:\"reasonKind,omitempty\"`\n\tAuthor          string                  `json:\"author\"`\n\tCommit          string                  `json:\"commit,omitempty\"`\n\tSource          *record.PublishedSource `json:\"source,omitempty\"`\n\tProject         string                  `json:\"project,omitempty\"`\n\tOrigin          string                  `json:\"origin,omitempty\"`\n\tBaseVersion     int                     `json:\"baseVersion,omitempty\"`\n\tBaseContentHash string                  `json:\"baseContentHash,omitempty\"`\n}","file":"remote/client.go","name":"PutVersionRequest"}}
type PutVersionRequest struct {
	ExpectedHeadCid *string                 `json:"expectedHeadCid"`
	Version         int                     `json:"version"`
	ContentHash     string                  `json:"contentHash"`
	GenHash         string                  `json:"genHash,omitempty"`
	SpecHashes      []string                `json:"specHashes"`
	Content         record.PublishedContent `json:"content"`
	Status          string                  `json:"status,omitempty"`
	Reason          string                  `json:"reason,omitempty"`
	ReasonKind      string                  `json:"reasonKind,omitempty"`
	Author          string                  `json:"author"`
	Commit          string                  `json:"commit,omitempty"`
	Source          *record.PublishedSource `json:"source,omitempty"`
	Project         string                  `json:"project,omitempty"`
	Origin          string                  `json:"origin,omitempty"`
	BaseVersion     int                     `json:"baseVersion,omitempty"`
	BaseContentHash string                  `json:"baseContentHash,omitempty"`
}

// /sedum:setType
// sedum:setType {"tier":"owned","record":"prov-2026-63180fe6","gen":"sha256:ddbbbb739d8c2b6c7f3011db9958dc7da4465ada2eb7c69f0dd519d5033ba4f5","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type Client struct {\n\tBase string\n\tHTTP *http.Client\n\tAuth func(ctx context.Context, fn func(bearer string) error) error\n}","file":"remote/client.go","name":"Client"}}
type Client struct {
	Base string
	HTTP *http.Client
	Auth func(ctx context.Context, fn func(bearer string) error) error
}

// /sedum:setType
// sedum:setType {"tier":"owned","record":"prov-2026-6f7372d7","gen":"sha256:4dca28a6d5ddaba6c2bc33765117c9754310305c32543bdf2a48de6e0f49e4d5","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type ConflictError struct {\n\tAPIError\n\tHead *Head\n}","file":"remote/client.go","name":"ConflictError"}}
type ConflictError struct {
	APIError
	Head *Head
}

// /sedum:setType
// sedum:setFunction {"tier":"owned","record":"prov-2026-79215eb3","gen":"sha256:edf570b32560aaec05e0dd66fc26de5989da672396bd1d0419e39c2cc60de6b9","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func requestID(id string) string {\n\tif id == \"\" || len(id) > 64 {\n\t\treturn \"\"\n\t}\n\tfor _, c := range id {\n\t\tswitch {\n\t\tcase c >= 'a' && c <= 'z', c >= 'A' && c <= 'Z', c >= '0' && c <= '9', c == '.', c == '_', c == '-':\n\t\tdefault:\n\t\t\treturn \"\"\n\t\t}\n\t}\n\treturn id\n}","file":"remote/client.go","name":"requestID"}}
func requestID(id string) string {
	if id == "" || len(id) > 64 {
		return ""
	}
	for _, c := range id {
		switch {
		case c >= 'a' && c <= 'z', c >= 'A' && c <= 'Z', c >= '0' && c <= '9', c == '.', c == '_', c == '-':
		default:
			return ""
		}
	}
	return id
}

// /sedum:setFunction
// sedum:setMethod {"tier":"owned","record":"prov-2026-7987372d","gen":"sha256:419562ebad626cdaadfaf3572533bcc8bc8bc798e5fa1a1d4ec4e3fb890ee139","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (c *Client) once(ctx context.Context, method, path, bearer string, in, out any) error {\n\tvar body io.Reader\n\tif in != nil {\n\t\tdata, err := json.Marshal(in)\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n\t\tbody = bytes.NewReader(data)\n\t}\n\ttarget := strings.TrimRight(c.Base, \"/\") + path\n\treq, err := http.NewRequestWithContext(ctx, method, target, body)\n\tif err != nil {\n\t\treturn fmt.Errorf(\"%s %s: bad request\", method, path)\n\t}\n\treq.Header.Set(\"Accept\", \"application/json\")\n\treq.Header.Set(\"Authorization\", \"Bearer \"+bearer)\n\tif in != nil {\n\t\treq.Header.Set(\"Content-Type\", \"application/json\")\n\t}\n\tresp, err := c.httpClient().Do(req)\n\tif err != nil {\n\t\tif ctx.Err() != nil {\n\t\t\treturn ctx.Err()\n\t\t}\n\t\tcause := err\n\t\tvar uerr *url.Error\n\t\tif errors.As(err, &uerr) {\n\t\t\tcause = uerr.Err\n\t\t}\n\t\treturn fmt.Errorf(\"%s %s: %w\", method, path, cause)\n\t}\n\tdefer resp.Body.Close()\n\traw, _ := io.ReadAll(io.LimitReader(resp.Body, 8388608))\n\tif resp.StatusCode/100 == 2 {\n\t\tif out == nil {\n\t\t\treturn nil\n\t\t}\n\t\tif err := json.Unmarshal(raw, out); err != nil {\n\t\t\treturn fmt.Errorf(\"%s %s: unreadable response\", method, path)\n\t\t}\n\t\treturn nil\n\t}\n\tif resp.StatusCode == http.StatusUnauthorized {\n\t\treturn auth.ErrUnauthorized\n\t}\n\ttype errorBody struct {\n\t\tError struct {\n\t\t\tCode      string `json:\"code\"`\n\t\t\tMessage   string `json:\"message\"`\n\t\t\tRequestID string `json:\"requestId\"`\n\t\t\tDetails   struct {\n\t\t\t\tHead *Head `json:\"head\"`\n\t\t\t} `json:\"details\"`\n\t\t} `json:\"error\"`\n\t}\n\tvar payload errorBody\n\tif err := json.Unmarshal(raw, &payload); err != nil {\n\t\tpayload = errorBody{}\n\t}\n\trid := requestID(resp.Header.Get(\"X-Request-Id\"))\n\tif rid == \"\" {\n\t\trid = requestID(payload.Error.RequestID)\n\t}\n\tcode := payload.Error.Code\n\tif code == \"\" {\n\t\tcode = fmt.Sprintf(\"http_%d\", resp.StatusCode)\n\t}\n\tapiErr := APIError{\n\t\tStatus:    resp.StatusCode,\n\t\tCode:      code,\n\t\tMessage:   payload.Error.Message,\n\t\tRequestID: rid,\n\t}\n\tif resp.StatusCode == http.StatusConflict {\n\t\treturn &ConflictError{APIError: apiErr, Head: payload.Error.Details.Head}\n\t}\n\treturn &apiErr\n}","file":"remote/client.go","name":"once","receiver":"Client"}}
func (c *Client) once(ctx context.Context, method, path, bearer string, in, out any) error {
	var body io.Reader
	if in != nil {
		data, err := json.Marshal(in)
		if err != nil {
			return err
		}
		body = bytes.NewReader(data)
	}
	target := strings.TrimRight(c.Base, "/") + path
	req, err := http.NewRequestWithContext(ctx, method, target, body)
	if err != nil {
		return fmt.Errorf("%s %s: bad request", method, path)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+bearer)
	if in != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	resp, err := c.httpClient().Do(req)
	if err != nil {
		if ctx.Err() != nil {
			return ctx.Err()
		}
		cause := err
		var uerr *url.Error
		if errors.As(err, &uerr) {
			cause = uerr.Err
		}
		return fmt.Errorf("%s %s: %w", method, path, cause)
	}
	defer resp.Body.Close()
	raw, _ := io.ReadAll(io.LimitReader(resp.Body, 8388608))
	if resp.StatusCode/100 == 2 {
		if out == nil {
			return nil
		}
		if err := json.Unmarshal(raw, out); err != nil {
			return fmt.Errorf("%s %s: unreadable response", method, path)
		}
		return nil
	}
	if resp.StatusCode == http.StatusUnauthorized {
		return auth.ErrUnauthorized
	}
	type errorBody struct {
		Error struct {
			Code      string `json:"code"`
			Message   string `json:"message"`
			RequestID string `json:"requestId"`
			Details   struct {
				Head *Head `json:"head"`
			} `json:"details"`
		} `json:"error"`
	}
	var payload errorBody
	if err := json.Unmarshal(raw, &payload); err != nil {
		payload = errorBody{}
	}
	rid := requestID(resp.Header.Get("X-Request-Id"))
	if rid == "" {
		rid = requestID(payload.Error.RequestID)
	}
	code := payload.Error.Code
	if code == "" {
		code = fmt.Sprintf("http_%d", resp.StatusCode)
	}
	apiErr := APIError{
		Status:    resp.StatusCode,
		Code:      code,
		Message:   payload.Error.Message,
		RequestID: rid,
	}
	if resp.StatusCode == http.StatusConflict {
		return &ConflictError{APIError: apiErr, Head: payload.Error.Details.Head}
	}
	return &apiErr
}

// /sedum:setMethod
// sedum:setMethod {"tier":"owned","record":"prov-2026-89ea7edb","gen":"sha256:f823d98f7ecc5ab6702e28b833d8e26cdd9611fa7fb9e368a84fd43322b86a83","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (c *Client) ListVersions(ctx context.Context, id string) ([]Version, error) {\n\tvar resp struct {\n\t\tVersions []Version `json:\"versions\"`\n\t}\n\tif err := c.do(ctx, http.MethodGet, \"/records/\"+url.PathEscape(id)+\"/versions\", nil, &resp); err != nil {\n\t\tvar apiErr *APIError\n\t\tif errors.As(err, &apiErr) && apiErr.Status == http.StatusNotFound && apiErr.Code != \"http_404\" {\n\t\t\treturn nil, ErrNotPublished\n\t\t}\n\t\treturn nil, err\n\t}\n\treturn resp.Versions, nil\n}","file":"remote/client.go","name":"ListVersions","receiver":"Client"}}
func (c *Client) ListVersions(ctx context.Context, id string) ([]Version, error) {
	var resp struct {
		Versions []Version `json:"versions"`
	}
	if err := c.do(ctx, http.MethodGet, "/records/"+url.PathEscape(id)+"/versions", nil, &resp); err != nil {
		var apiErr *APIError
		if errors.As(err, &apiErr) && apiErr.Status == http.StatusNotFound && apiErr.Code != "http_404" {
			return nil, ErrNotPublished
		}
		return nil, err
	}
	return resp.Versions, nil
}

// /sedum:setMethod
// sedum:setConst {"tier":"owned","record":"prov-2026-944aad75","gen":"sha256:6838123b24ae8c0e9168cde940872eaa373b7a144ea20842ef30e615cce30324","pkg":"sha256:51be1c18998d63cc558c253e211a6bd188d96cb251dad68c6ea88e03c2d5cc4c","kwargs":{"code":"const EnvBase = \"PROVISIO_API_URL\"","file":"remote/client.go","name":"EnvBase"}}
const EnvBase = "PROVISIO_API_URL"

// /sedum:setConst
// sedum:setMethod {"tier":"owned","record":"prov-2026-a08d0566","gen":"sha256:fde9ef703ea50519df69b9625151aaa6507ebbabee9558ea405c4e6f33110f5b","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (e *APIError) Error() string {\n\ts := fmt.Sprintf(\"%s (HTTP %d)\", e.Code, e.Status)\n\tif e.Message != \"\" {\n\t\ts += \": \" + e.Message\n\t}\n\tif e.Code == \"http_404\" {\n\t\ts += \" (the answer was not from the record service: PROVISIO_API_URL must include the version path, for example, \" + DefaultBase + \")\"\n\t}\n\tif e.RequestID != \"\" {\n\t\ts += \" (request \" + e.RequestID + \")\"\n\t}\n\treturn s\n}","file":"remote/client.go","name":"Error","receiver":"APIError"}}
func (e *APIError) Error() string {
	s := fmt.Sprintf("%s (HTTP %d)", e.Code, e.Status)
	if e.Message != "" {
		s += ": " + e.Message
	}
	if e.Code == "http_404" {
		s += " (the answer was not from the record service: PROVISIO_API_URL must include the version path, for example, " + DefaultBase + ")"
	}
	if e.RequestID != "" {
		s += " (request " + e.RequestID + ")"
	}
	return s
}

// /sedum:setMethod
// sedum:setMethod {"tier":"owned","record":"prov-2026-ac76df08","gen":"sha256:1c162f07c6be12c559a20fa8b410859c102e20b5b4c16517156f19d25c1984c4","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (c *Client) do(ctx context.Context, method string, path string, in any, out any) error {\n\tif err := auth.CheckServer(c.Base); err != nil {\n\t\treturn err\n\t}\n\treturn c.Auth(ctx, func(bearer string) error {\n\t\treturn c.once(ctx, method, path, bearer, in, out)\n\t})\n}","file":"remote/client.go","name":"do","receiver":"Client"}}
func (c *Client) do(ctx context.Context, method string, path string, in any, out any) error {
	if err := auth.CheckServer(c.Base); err != nil {
		return err
	}
	return c.Auth(ctx, func(bearer string) error {
		return c.once(ctx, method, path, bearer, in, out)
	})
}

// /sedum:setMethod
// sedum:setType {"tier":"owned","record":"prov-2026-b31095ad","gen":"sha256:96f9d66d7c905f9efdcf01cf95144207704261cccf4b1922ac492040a5af7b7f","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type Version struct {\n\trecord.PublishedVersion\n\tURI string `json:\"uri\"`\n\tCID string `json:\"cid\"`\n}","file":"remote/client.go","name":"Version"}}
type Version struct {
	record.PublishedVersion
	URI string `json:"uri"`
	CID string `json:"cid"`
}

// /sedum:setType
// sedum:setMethod {"tier":"owned","record":"prov-2026-c8416d8d","gen":"sha256:cb455fec8c80fdfd1449edaebadc48f8e3f15aab15f0b5d4a54bdfb906c899ba","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (c *Client) ListRecords(ctx context.Context) ([]Head, error) {\n\tvar out struct {\n\t\tRecords []Head `json:\"records\"`\n\t}\n\tif err := c.do(ctx, http.MethodGet, \"/records\", nil, &out); err != nil {\n\t\treturn nil, err\n\t}\n\treturn out.Records, nil\n}","file":"remote/client.go","name":"ListRecords","receiver":"Client"}}
func (c *Client) ListRecords(ctx context.Context) ([]Head, error) {
	var out struct {
		Records []Head `json:"records"`
	}
	if err := c.do(ctx, http.MethodGet, "/records", nil, &out); err != nil {
		return nil, err
	}
	return out.Records, nil
}

// /sedum:setMethod
// sedum:setMethod {"tier":"owned","record":"prov-2026-e7e7a8ff","gen":"sha256:295b5e43ea2d3a4e378c61f7f64ee610cbb61d0109cb9477ffc6a6a6abe5bc15","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (c *Client) PutHead(ctx context.Context, id string, req PutHeadRequest) (*Head, error) {\n\tvar out struct {\n\t\tHead Head `json:\"head\"`\n\t}\n\tif err := c.do(ctx, http.MethodPut, \"/records/\"+url.PathEscape(id)+\"/head\", req, &out); err != nil {\n\t\treturn nil, err\n\t}\n\tout.Head.ID = id\n\treturn &out.Head, nil\n}","file":"remote/client.go","name":"PutHead","receiver":"Client"}}
func (c *Client) PutHead(ctx context.Context, id string, req PutHeadRequest) (*Head, error) {
	var out struct {
		Head Head `json:"head"`
	}
	if err := c.do(ctx, http.MethodPut, "/records/"+url.PathEscape(id)+"/head", req, &out); err != nil {
		return nil, err
	}
	out.Head.ID = id
	return &out.Head, nil
}

// /sedum:setMethod
// sedum:setType {"tier":"owned","record":"prov-2026-ec204b38","gen":"sha256:0180ebd45e4a160e23f2f8b092a13b1b10cc34c71b343075c2add25ca26b5338","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type PutHeadRequest struct {\n\tExpectedHeadCid string                  `json:\"expectedHeadCid\"`\n\tVersion         int                     `json:\"version\"`\n\tStatus          string                  `json:\"status,omitempty\"`\n\tSource          *record.PublishedSource `json:\"source,omitempty\"`\n\tProject         string                  `json:\"project,omitempty\"`\n}","file":"remote/client.go","name":"PutHeadRequest"}}
type PutHeadRequest struct {
	ExpectedHeadCid string                  `json:"expectedHeadCid"`
	Version         int                     `json:"version"`
	Status          string                  `json:"status,omitempty"`
	Source          *record.PublishedSource `json:"source,omitempty"`
	Project         string                  `json:"project,omitempty"`
}

// /sedum:setType
// sedum:setMethod {"tier":"owned","record":"prov-2026-f2df1516","gen":"sha256:294f346fe269be981a45336bb8b3cd303a097f871c246275cb6a253d6ac9c671","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (c *Client) GetRecord(ctx context.Context, id string) (*Head, error) {\n\tvar out struct {\n\t\tHead Head `json:\"head\"`\n\t}\n\tif err := c.do(ctx, http.MethodGet, \"/records/\"+url.PathEscape(id), nil, &out); err != nil {\n\t\tvar apiErr *APIError\n\t\tif errors.As(err, &apiErr) && apiErr.Status == http.StatusNotFound && apiErr.Code != \"http_404\" {\n\t\t\treturn nil, ErrNotPublished\n\t\t}\n\t\treturn nil, err\n\t}\n\tout.Head.ID = id\n\treturn &out.Head, nil\n}","file":"remote/client.go","name":"GetRecord","receiver":"Client"}}
func (c *Client) GetRecord(ctx context.Context, id string) (*Head, error) {
	var out struct {
		Head Head `json:"head"`
	}
	if err := c.do(ctx, http.MethodGet, "/records/"+url.PathEscape(id), nil, &out); err != nil {
		var apiErr *APIError
		if errors.As(err, &apiErr) && apiErr.Status == http.StatusNotFound && apiErr.Code != "http_404" {
			return nil, ErrNotPublished
		}
		return nil, err
	}
	out.Head.ID = id
	return &out.Head, nil
}

// /sedum:setMethod
// sedum:setType {"tier":"owned","record":"prov-2026-f52bba41","gen":"sha256:5e230d9e7a8ecd83f527683677fbd6e37dd43f575674655b5082f77d1b27868f","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type APIError struct {\n\tStatus    int\n\tCode      string\n\tMessage   string\n\tRequestID string\n}","file":"remote/client.go","name":"APIError"}}
type APIError struct {
	Status    int
	Code      string
	Message   string
	RequestID string
}

// /sedum:setType
