package remote

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-ba01059c","gen":"sha256:a2b910889c80430c436335f54c432153f3016ba3a3f69eb08da231a5cb0274ce","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"remote/client.go","path":"bytes"}}
import "bytes"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-ba01059c","gen":"sha256:a2b910889c80430c436335f54c432153f3016ba3a3f69eb08da231a5cb0274ce","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"remote/client.go","path":"context"}}
import "context"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-ba01059c","gen":"sha256:a2b910889c80430c436335f54c432153f3016ba3a3f69eb08da231a5cb0274ce","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"remote/client.go","path":"encoding/json"}}
import "encoding/json"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-ba01059c","gen":"sha256:a2b910889c80430c436335f54c432153f3016ba3a3f69eb08da231a5cb0274ce","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"remote/client.go","path":"errors"}}
import "errors"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-ba01059c","gen":"sha256:a2b910889c80430c436335f54c432153f3016ba3a3f69eb08da231a5cb0274ce","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"remote/client.go","path":"fmt"}}
import "fmt"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-ba01059c","gen":"sha256:a2b910889c80430c436335f54c432153f3016ba3a3f69eb08da231a5cb0274ce","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"remote/client.go","path":"io"}}
import "io"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-ba01059c","gen":"sha256:a2b910889c80430c436335f54c432153f3016ba3a3f69eb08da231a5cb0274ce","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"remote/client.go","path":"net/http"}}
import "net/http"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-ba01059c","gen":"sha256:a2b910889c80430c436335f54c432153f3016ba3a3f69eb08da231a5cb0274ce","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"remote/client.go","path":"net/url"}}
import "net/url"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-ba01059c","gen":"sha256:a2b910889c80430c436335f54c432153f3016ba3a3f69eb08da231a5cb0274ce","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"remote/client.go","path":"strings"}}
import "strings"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-ba01059c","gen":"sha256:a2b910889c80430c436335f54c432153f3016ba3a3f69eb08da231a5cb0274ce","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"remote/client.go","path":"time"}}
import "time"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-ba01059c","gen":"sha256:a2b910889c80430c436335f54c432153f3016ba3a3f69eb08da231a5cb0274ce","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"remote/client.go","path":"github.com/livecodelife/provenance/auth"}}
import "github.com/livecodelife/provenance/auth"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-ba01059c","gen":"sha256:a2b910889c80430c436335f54c432153f3016ba3a3f69eb08da231a5cb0274ce","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"remote/client.go","path":"github.com/livecodelife/provenance/record"}}
import "github.com/livecodelife/provenance/record"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setConst {"tier":"owned","record":"prov-2026-ba01059c","gen":"sha256:a2b910889c80430c436335f54c432153f3016ba3a3f69eb08da231a5cb0274ce","pkg":"sha256:51be1c18998d63cc558c253e211a6bd188d96cb251dad68c6ea88e03c2d5cc4c","kwargs":{"code":"// DefaultBase is the record server's API base (RECORDS-API.md).\nconst DefaultBase = \"https://api.provisio.dev/v1\"","file":"remote/client.go","name":"DefaultBase"}}
// DefaultBase is the record server's API base (RECORDS-API.md).
const DefaultBase = "https://api.provisio.dev/v1"

// /sedum:setConst
// sedum:setConst {"tier":"owned","record":"prov-2026-ba01059c","gen":"sha256:a2b910889c80430c436335f54c432153f3016ba3a3f69eb08da231a5cb0274ce","pkg":"sha256:51be1c18998d63cc558c253e211a6bd188d96cb251dad68c6ea88e03c2d5cc4c","kwargs":{"code":"// EnvBase overrides the API base.\nconst EnvBase = \"PROVISIO_API_URL\"","file":"remote/client.go","name":"EnvBase"}}
// EnvBase overrides the API base.
const EnvBase = "PROVISIO_API_URL"

// /sedum:setConst
// sedum:setVar {"tier":"owned","record":"prov-2026-ba01059c","gen":"sha256:a2b910889c80430c436335f54c432153f3016ba3a3f69eb08da231a5cb0274ce","pkg":"sha256:c7b9d9ac1655e8211b8b2ca8c49cf34b8d510d8d823d4f5eb60ec7e1018aa2eb","kwargs":{"code":"var (\n\t// ErrNotPublished means the server has no such record in the caller's repo.\n\tErrNotPublished = errors.New(\"record is not published\")\n\t// ErrUnsupported is returned by Store methods a remote record has no meaning for.\n\tErrUnsupported = errors.New(\"not supported by the Provisio store: a remote record changes only by a new version\")\n)","file":"remote/client.go","name":"ErrNotPublished"}}
var (
	// ErrNotPublished means the server has no such record in the caller's repo.
	ErrNotPublished = errors.New("record is not published")
	// ErrUnsupported is returned by Store methods a remote record has no meaning for.
	ErrUnsupported = errors.New("not supported by the Provisio store: a remote record changes only by a new version")
)

// /sedum:setVar
// sedum:setType {"tier":"owned","record":"prov-2026-ba01059c","gen":"sha256:a2b910889c80430c436335f54c432153f3016ba3a3f69eb08da231a5cb0274ce","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"// Head is a published record head plus the server's id, uri and cid.\ntype Head struct {\n\trecord.PublishedHead\n\tID  string `json:\"id\"`\n\tURI string `json:\"uri\"`\n\tCID string `json:\"cid\"`\n}","file":"remote/client.go","name":"Head"}}
// Head is a published record head plus the server's id, uri and cid.
type Head struct {
	record.PublishedHead
	ID  string `json:"id"`
	URI string `json:"uri"`
	CID string `json:"cid"`
}

// /sedum:setType
// sedum:setType {"tier":"owned","record":"prov-2026-ba01059c","gen":"sha256:a2b910889c80430c436335f54c432153f3016ba3a3f69eb08da231a5cb0274ce","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"// Version is a published version plus the server's uri and cid.\ntype Version struct {\n\trecord.PublishedVersion\n\tURI string `json:\"uri\"`\n\tCID string `json:\"cid\"`\n}","file":"remote/client.go","name":"Version"}}
// Version is a published version plus the server's uri and cid.
type Version struct {
	record.PublishedVersion
	URI string `json:"uri"`
	CID string `json:"cid"`
}

// /sedum:setType
// sedum:setType {"tier":"owned","record":"prov-2026-ba01059c","gen":"sha256:a2b910889c80430c436335f54c432153f3016ba3a3f69eb08da231a5cb0274ce","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"// PutVersionRequest is the body of PUT /records/:id/versions.\ntype PutVersionRequest struct {\n\tExpectedHeadCid *string                 `json:\"expectedHeadCid\"`\n\tVersion         int                     `json:\"version\"`\n\tContentHash     string                  `json:\"contentHash\"`\n\tGenHash         string                  `json:\"genHash,omitempty\"`\n\tSpecHashes      []string                `json:\"specHashes\"`\n\tContent         record.PublishedContent `json:\"content\"`\n\tStatus          string                  `json:\"status,omitempty\"`\n\tReason          string                  `json:\"reason,omitempty\"`\n\tReasonKind      string                  `json:\"reasonKind,omitempty\"`\n\tAuthor          string                  `json:\"author\"`\n\tCommit          string                  `json:\"commit,omitempty\"`\n\tSource          *record.PublishedSource `json:\"source,omitempty\"`\n\tProject         string                  `json:\"project,omitempty\"`\n\tOrigin          string                  `json:\"origin,omitempty\"`\n\tBaseVersion     int                     `json:\"baseVersion,omitempty\"`\n\tBaseContentHash string                  `json:\"baseContentHash,omitempty\"`\n}","file":"remote/client.go","name":"PutVersionRequest"}}
// PutVersionRequest is the body of PUT /records/:id/versions.
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
// sedum:setType {"tier":"owned","record":"prov-2026-ba01059c","gen":"sha256:a2b910889c80430c436335f54c432153f3016ba3a3f69eb08da231a5cb0274ce","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"// PutHeadRequest is the body of PUT /records/:id/head: a head-only update\n// (status moving forward, source, project) that writes no version.\ntype PutHeadRequest struct {\n\tExpectedHeadCid string                  `json:\"expectedHeadCid\"`\n\tVersion         int                     `json:\"version\"`\n\tStatus          string                  `json:\"status,omitempty\"`\n\tSource          *record.PublishedSource `json:\"source,omitempty\"`\n\tProject         string                  `json:\"project,omitempty\"`\n}","file":"remote/client.go","name":"PutHeadRequest"}}
// PutHeadRequest is the body of PUT /records/:id/head: a head-only update
// (status moving forward, source, project) that writes no version.
type PutHeadRequest struct {
	ExpectedHeadCid string                  `json:"expectedHeadCid"`
	Version         int                     `json:"version"`
	Status          string                  `json:"status,omitempty"`
	Source          *record.PublishedSource `json:"source,omitempty"`
	Project         string                  `json:"project,omitempty"`
}

// /sedum:setType
// sedum:setType {"tier":"owned","record":"prov-2026-ba01059c","gen":"sha256:a2b910889c80430c436335f54c432153f3016ba3a3f69eb08da231a5cb0274ce","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"// APIError is an error answer from the record server. It never carries a\n// request body or a token. RequestID is the id the server attached to the\n// answer, kept only if requestID accepts it.\ntype APIError struct {\n\tStatus    int\n\tCode      string\n\tMessage   string\n\tRequestID string\n}","file":"remote/client.go","name":"APIError"}}
// APIError is an error answer from the record server. It never carries a
// request body or a token. RequestID is the id the server attached to the
// answer, kept only if requestID accepts it.
type APIError struct {
	Status    int
	Code      string
	Message   string
	RequestID string
}

// /sedum:setType
// sedum:setMethod {"tier":"owned","record":"prov-2026-ba01059c","gen":"sha256:a2b910889c80430c436335f54c432153f3016ba3a3f69eb08da231a5cb0274ce","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (e *APIError) Error() string {\n\ts := fmt.Sprintf(\"%s (HTTP %d)\", e.Code, e.Status)\n\tif e.Message != \"\" {\n\t\ts = fmt.Sprintf(\"%s (HTTP %d): %s\", e.Code, e.Status, e.Message)\n\t}\n\tif e.Code == \"http_404\" {\n\t\ts += \" (the answer was not from the record service: PROVISIO_API_URL must include the version path, for example \" + DefaultBase + \")\"\n\t}\n\tif e.RequestID != \"\" {\n\t\ts += \" (request \" + e.RequestID + \")\"\n\t}\n\treturn s\n}","file":"remote/client.go","name":"Error","receiver":"APIError"}}
func (e *APIError) Error() string {
	s := fmt.Sprintf("%s (HTTP %d)", e.Code, e.Status)
	if e.Message != "" {
		s = fmt.Sprintf("%s (HTTP %d): %s", e.Code, e.Status, e.Message)
	}
	if e.Code == "http_404" {
		s += " (the answer was not from the record service: PROVISIO_API_URL must include the version path, for example " + DefaultBase + ")"
	}
	if e.RequestID != "" {
		s += " (request " + e.RequestID + ")"
	}
	return s
}

// /sedum:setMethod
// sedum:setType {"tier":"owned","record":"prov-2026-ba01059c","gen":"sha256:a2b910889c80430c436335f54c432153f3016ba3a3f69eb08da231a5cb0274ce","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"// ConflictError is a 409: the head moved. Head is the server's current head, or nil.\ntype ConflictError struct {\n\tAPIError\n\tHead *Head\n}","file":"remote/client.go","name":"ConflictError"}}
// ConflictError is a 409: the head moved. Head is the server's current head, or nil.
type ConflictError struct {
	APIError
	Head *Head
}

// /sedum:setType
// sedum:setType {"tier":"owned","record":"prov-2026-ba01059c","gen":"sha256:a2b910889c80430c436335f54c432153f3016ba3a3f69eb08da231a5cb0274ce","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"// Client talks to the record server. Auth runs a request function with a\n// bearer token and retries it once on auth.ErrUnauthorized (auth.Manager.Do).\ntype Client struct {\n\tBase string\n\tHTTP *http.Client\n\tAuth func(ctx context.Context, fn func(bearer string) error) error\n}","file":"remote/client.go","name":"Client"}}
// Client talks to the record server. Auth runs a request function with a
// bearer token and retries it once on auth.ErrUnauthorized (auth.Manager.Do).
type Client struct {
	Base string
	HTTP *http.Client
	Auth func(ctx context.Context, fn func(bearer string) error) error
}

// /sedum:setType
// sedum:setMethod {"tier":"owned","record":"prov-2026-ba01059c","gen":"sha256:a2b910889c80430c436335f54c432153f3016ba3a3f69eb08da231a5cb0274ce","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (c *Client) httpClient() *http.Client {\n\tif c.HTTP != nil {\n\t\treturn c.HTTP\n\t}\n\treturn &http.Client{Timeout: 60 * time.Second}\n}","file":"remote/client.go","name":"httpClient","receiver":"Client"}}
func (c *Client) httpClient() *http.Client {
	if c.HTTP != nil {
		return c.HTTP
	}
	return &http.Client{Timeout: 60 * time.Second}
}

// /sedum:setMethod
// sedum:setMethod {"tier":"owned","record":"prov-2026-ba01059c","gen":"sha256:a2b910889c80430c436335f54c432153f3016ba3a3f69eb08da231a5cb0274ce","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (c *Client) do(ctx context.Context, method, path string, in, out any) error {\n\tif err := auth.CheckServer(c.Base); err != nil {\n\t\treturn err\n\t}\n\treturn c.Auth(ctx, func(bearer string) error { return c.once(ctx, method, path, bearer, in, out) })\n}","file":"remote/client.go","name":"do","receiver":"Client"}}
func (c *Client) do(ctx context.Context, method, path string, in, out any) error {
	if err := auth.CheckServer(c.Base); err != nil {
		return err
	}
	return c.Auth(ctx, func(bearer string) error { return c.once(ctx, method, path, bearer, in, out) })
}

// /sedum:setMethod
// sedum:setMethod {"tier":"owned","record":"prov-2026-ba01059c","gen":"sha256:a2b910889c80430c436335f54c432153f3016ba3a3f69eb08da231a5cb0274ce","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (c *Client) once(ctx context.Context, method, path, bearer string, in, out any) error {\n\tvar body io.Reader\n\tif in != nil {\n\t\tdata, err := json.Marshal(in)\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n\t\tbody = bytes.NewReader(data)\n\t}\n\treq, err := http.NewRequestWithContext(ctx, method, strings.TrimRight(c.Base, \"/\")+path, body)\n\tif err != nil {\n\t\treturn fmt.Errorf(\"%s %s: bad request\", method, path)\n\t}\n\treq.Header.Set(\"Accept\", \"application/json\")\n\treq.Header.Set(\"Authorization\", \"Bearer \"+bearer)\n\tif in != nil {\n\t\treq.Header.Set(\"Content-Type\", \"application/json\")\n\t}\n\tresp, err := c.httpClient().Do(req)\n\tif err != nil {\n\t\tif ctx.Err() != nil {\n\t\t\treturn ctx.Err()\n\t\t}\n\t\tvar ue *url.Error\n\t\tif errors.As(err, &ue) {\n\t\t\terr = ue.Err\n\t\t}\n\t\treturn fmt.Errorf(\"%s %s: %w\", method, path, err)\n\t}\n\tdefer resp.Body.Close()\n\tdata, _ := io.ReadAll(io.LimitReader(resp.Body, 8<<20))\n\tif resp.StatusCode/100 == 2 {\n\t\tif out == nil {\n\t\t\treturn nil\n\t\t}\n\t\tif err := json.Unmarshal(data, out); err != nil {\n\t\t\treturn fmt.Errorf(\"%s %s: unreadable response\", method, path)\n\t\t}\n\t\treturn nil\n\t}\n\tif resp.StatusCode == http.StatusUnauthorized {\n\t\treturn auth.ErrUnauthorized\n\t}\n\tvar env struct {\n\t\tError struct {\n\t\t\tCode      string `json:\"code\"`\n\t\t\tMessage   string `json:\"message\"`\n\t\t\tRequestID string `json:\"requestId\"`\n\t\t\tDetails   struct {\n\t\t\t\tHead *Head `json:\"head\"`\n\t\t\t} `json:\"details\"`\n\t\t} `json:\"error\"`\n\t}\n\t_ = json.Unmarshal(data, &env)\n\tae := APIError{Status: resp.StatusCode, Code: env.Error.Code, Message: env.Error.Message, RequestID: requestID(resp.Header.Get(\"X-Request-Id\"))}\n\tif ae.RequestID == \"\" {\n\t\tae.RequestID = requestID(env.Error.RequestID)\n\t}\n\tif ae.Code == \"\" {\n\t\tae.Code = fmt.Sprintf(\"http_%d\", resp.StatusCode)\n\t}\n\tif resp.StatusCode == http.StatusConflict {\n\t\treturn &ConflictError{APIError: ae, Head: env.Error.Details.Head}\n\t}\n\treturn &ae\n}","file":"remote/client.go","name":"once","receiver":"Client"}}
func (c *Client) once(ctx context.Context, method, path, bearer string, in, out any) error {
	var body io.Reader
	if in != nil {
		data, err := json.Marshal(in)
		if err != nil {
			return err
		}
		body = bytes.NewReader(data)
	}
	req, err := http.NewRequestWithContext(ctx, method, strings.TrimRight(c.Base, "/")+path, body)
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
		var ue *url.Error
		if errors.As(err, &ue) {
			err = ue.Err
		}
		return fmt.Errorf("%s %s: %w", method, path, err)
	}
	defer resp.Body.Close()
	data, _ := io.ReadAll(io.LimitReader(resp.Body, 8<<20))
	if resp.StatusCode/100 == 2 {
		if out == nil {
			return nil
		}
		if err := json.Unmarshal(data, out); err != nil {
			return fmt.Errorf("%s %s: unreadable response", method, path)
		}
		return nil
	}
	if resp.StatusCode == http.StatusUnauthorized {
		return auth.ErrUnauthorized
	}
	var env struct {
		Error struct {
			Code      string `json:"code"`
			Message   string `json:"message"`
			RequestID string `json:"requestId"`
			Details   struct {
				Head *Head `json:"head"`
			} `json:"details"`
		} `json:"error"`
	}
	_ = json.Unmarshal(data, &env)
	ae := APIError{Status: resp.StatusCode, Code: env.Error.Code, Message: env.Error.Message, RequestID: requestID(resp.Header.Get("X-Request-Id"))}
	if ae.RequestID == "" {
		ae.RequestID = requestID(env.Error.RequestID)
	}
	if ae.Code == "" {
		ae.Code = fmt.Sprintf("http_%d", resp.StatusCode)
	}
	if resp.StatusCode == http.StatusConflict {
		return &ConflictError{APIError: ae, Head: env.Error.Details.Head}
	}
	return &ae
}

// /sedum:setMethod
// sedum:setMethod {"tier":"owned","record":"prov-2026-ba01059c","gen":"sha256:a2b910889c80430c436335f54c432153f3016ba3a3f69eb08da231a5cb0274ce","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (e *ConflictError) Error() string { return e.APIError.Error() }","file":"remote/client.go","name":"Error","receiver":"ConflictError"}}
func (e *ConflictError) Error() string { return e.APIError.Error() }

// /sedum:setMethod
// sedum:setMethod {"tier":"owned","record":"prov-2026-ba01059c","gen":"sha256:a2b910889c80430c436335f54c432153f3016ba3a3f69eb08da231a5cb0274ce","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"// ListRecords returns the caller's heads.\nfunc (c *Client) ListRecords(ctx context.Context) ([]Head, error) {\n\tvar out struct {\n\t\tRecords []Head `json:\"records\"`\n\t}\n\tif err := c.do(ctx, http.MethodGet, \"/records\", nil, &out); err != nil {\n\t\treturn nil, err\n\t}\n\treturn out.Records, nil\n}","file":"remote/client.go","name":"ListRecords","receiver":"Client"}}
// ListRecords returns the caller's heads.
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
// sedum:setMethod {"tier":"owned","record":"prov-2026-ba01059c","gen":"sha256:a2b910889c80430c436335f54c432153f3016ba3a3f69eb08da231a5cb0274ce","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"// GetRecord returns one head, or ErrNotPublished when the service says there\n// is no such record. A 404 without a JSON error body is not that answer.\nfunc (c *Client) GetRecord(ctx context.Context, id string) (*Head, error) {\n\tvar out struct {\n\t\tHead Head `json:\"head\"`\n\t}\n\terr := c.do(ctx, http.MethodGet, \"/records/\"+url.PathEscape(id), nil, &out)\n\tvar ae *APIError\n\tif errors.As(err, &ae) && ae.Status == http.StatusNotFound && ae.Code != \"http_404\" {\n\t\treturn nil, ErrNotPublished\n\t}\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\tout.Head.ID = id\n\treturn &out.Head, nil\n}","file":"remote/client.go","name":"GetRecord","receiver":"Client"}}
// GetRecord returns one head, or ErrNotPublished when the service says there
// is no such record. A 404 without a JSON error body is not that answer.
func (c *Client) GetRecord(ctx context.Context, id string) (*Head, error) {
	var out struct {
		Head Head `json:"head"`
	}
	err := c.do(ctx, http.MethodGet, "/records/"+url.PathEscape(id), nil, &out)
	var ae *APIError
	if errors.As(err, &ae) && ae.Status == http.StatusNotFound && ae.Code != "http_404" {
		return nil, ErrNotPublished
	}
	if err != nil {
		return nil, err
	}
	out.Head.ID = id
	return &out.Head, nil
}

// /sedum:setMethod
// sedum:setMethod {"tier":"owned","record":"prov-2026-ba01059c","gen":"sha256:a2b910889c80430c436335f54c432153f3016ba3a3f69eb08da231a5cb0274ce","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"// ListVersions returns a record's versions in order, or ErrNotPublished when\n// the service says there is no such record.\nfunc (c *Client) ListVersions(ctx context.Context, id string) ([]Version, error) {\n\tvar out struct {\n\t\tVersions []Version `json:\"versions\"`\n\t}\n\terr := c.do(ctx, http.MethodGet, \"/records/\"+url.PathEscape(id)+\"/versions\", nil, &out)\n\tvar ae *APIError\n\tif errors.As(err, &ae) && ae.Status == http.StatusNotFound && ae.Code != \"http_404\" {\n\t\treturn nil, ErrNotPublished\n\t}\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\treturn out.Versions, nil\n}","file":"remote/client.go","name":"ListVersions","receiver":"Client"}}
// ListVersions returns a record's versions in order, or ErrNotPublished when
// the service says there is no such record.
func (c *Client) ListVersions(ctx context.Context, id string) ([]Version, error) {
	var out struct {
		Versions []Version `json:"versions"`
	}
	err := c.do(ctx, http.MethodGet, "/records/"+url.PathEscape(id)+"/versions", nil, &out)
	var ae *APIError
	if errors.As(err, &ae) && ae.Status == http.StatusNotFound && ae.Code != "http_404" {
		return nil, ErrNotPublished
	}
	if err != nil {
		return nil, err
	}
	return out.Versions, nil
}

// /sedum:setMethod
// sedum:setMethod {"tier":"owned","record":"prov-2026-ba01059c","gen":"sha256:a2b910889c80430c436335f54c432153f3016ba3a3f69eb08da231a5cb0274ce","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"// PutVersion appends one version and returns the new head and version.\nfunc (c *Client) PutVersion(ctx context.Context, id string, req PutVersionRequest) (*Head, *Version, error) {\n\tvar out struct {\n\t\tHead    Head    `json:\"head\"`\n\t\tVersion Version `json:\"version\"`\n\t}\n\tif err := c.do(ctx, http.MethodPut, \"/records/\"+url.PathEscape(id)+\"/versions\", req, &out); err != nil {\n\t\treturn nil, nil, err\n\t}\n\tout.Head.ID = id\n\treturn &out.Head, &out.Version, nil\n}","file":"remote/client.go","name":"PutVersion","receiver":"Client"}}
// PutVersion appends one version and returns the new head and version.
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
// sedum:setMethod {"tier":"owned","record":"prov-2026-ba01059c","gen":"sha256:a2b910889c80430c436335f54c432153f3016ba3a3f69eb08da231a5cb0274ce","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"// PutHead updates the head only and returns it.\nfunc (c *Client) PutHead(ctx context.Context, id string, req PutHeadRequest) (*Head, error) {\n\tvar out struct {\n\t\tHead Head `json:\"head\"`\n\t}\n\tif err := c.do(ctx, http.MethodPut, \"/records/\"+url.PathEscape(id)+\"/head\", req, &out); err != nil {\n\t\treturn nil, err\n\t}\n\tout.Head.ID = id\n\treturn &out.Head, nil\n}","file":"remote/client.go","name":"PutHead","receiver":"Client"}}
// PutHead updates the head only and returns it.
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
// sedum:setFunction {"tier":"owned","record":"prov-2026-ba01059c","gen":"sha256:a2b910889c80430c436335f54c432153f3016ba3a3f69eb08da231a5cb0274ce","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"// requestID returns id when it is at most 64 characters of [A-Za-z0-9._-],\n// and \"\" otherwise, so nothing a server sends can reach the terminal as markup.\nfunc requestID(id string) string {\n\tif id == \"\" || len(id) > 64 {\n\t\treturn \"\"\n\t}\n\tfor _, r := range id {\n\t\tswitch {\n\t\tcase r >= 'a' && r <= 'z', r >= 'A' && r <= 'Z', r >= '0' && r <= '9', r == '.', r == '_', r == '-':\n\t\tdefault:\n\t\t\treturn \"\"\n\t\t}\n\t}\n\treturn id\n}","file":"remote/client.go","name":"requestID"}}
// requestID returns id when it is at most 64 characters of [A-Za-z0-9._-],
// and "" otherwise, so nothing a server sends can reach the terminal as markup.
func requestID(id string) string {
	if id == "" || len(id) > 64 {
		return ""
	}
	for _, r := range id {
		switch {
		case r >= 'a' && r <= 'z', r >= 'A' && r <= 'Z', r >= '0' && r <= '9', r == '.', r == '_', r == '-':
		default:
			return ""
		}
	}
	return id
}

// /sedum:setFunction
