package auth

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-2b4db381","gen":"sha256:45a2a4d28f38438bd03beb44de6ad9b5ee07bdad521a440fcf25bbc3ce684236","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"auth/credential.go","path":"errors"}}
import "errors"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-2b4db381","gen":"sha256:45a2a4d28f38438bd03beb44de6ad9b5ee07bdad521a440fcf25bbc3ce684236","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"auth/credential.go","path":"time"}}
import "time"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setType {"tier":"owned","record":"prov-2026-2b4db381","gen":"sha256:45a2a4d28f38438bd03beb44de6ad9b5ee07bdad521a440fcf25bbc3ce684236","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"// Source says where a credential came from.\ntype Source string","file":"auth/credential.go","name":"Source"}}
// Source says where a credential came from.
type Source string

// /sedum:setType
// sedum:setConst {"tier":"owned","record":"prov-2026-2b4db381","gen":"sha256:45a2a4d28f38438bd03beb44de6ad9b5ee07bdad521a440fcf25bbc3ce684236","pkg":"sha256:51be1c18998d63cc558c253e211a6bd188d96cb251dad68c6ea88e03c2d5cc4c","kwargs":{"code":"const (\n\tSourceEnv      Source = \"env\"\n\tSourceKeychain Source = \"keychain\"\n\tSourceFile     Source = \"file\"\n)","file":"auth/credential.go","name":"SourceEnv"}}
const (
	SourceEnv      Source = "env"
	SourceKeychain Source = "keychain"
	SourceFile     Source = "file"
)

// /sedum:setConst
// sedum:setConst {"tier":"owned","record":"prov-2026-2b4db381","gen":"sha256:45a2a4d28f38438bd03beb44de6ad9b5ee07bdad521a440fcf25bbc3ce684236","pkg":"sha256:51be1c18998d63cc558c253e211a6bd188d96cb251dad68c6ea88e03c2d5cc4c","kwargs":{"code":"const (\n\t// EnvToken holds a personal token (CI). It wins over any stored credential.\n\tEnvToken = \"PROVISIO_TOKEN\"\n\t// EnvServer overrides the auth service base URL at login.\n\tEnvServer = \"PROVISIO_AUTH_URL\"\n\t// EnvNoKeychain set to a non-empty value keeps credentials out of the OS keychain.\n\tEnvNoKeychain = \"PROVISIO_NO_KEYCHAIN\"\n\t// DefaultServer is the auth service used when nothing else is configured.\n\tDefaultServer = \"https://auth.provisio.dev\"\n\t// ClientID is what the CLI sends as client_id.\n\tClientID = \"provenance-cli\"\n)","file":"auth/credential.go","name":"EnvToken"}}
const (
	// EnvToken holds a personal token (CI). It wins over any stored credential.
	EnvToken = "PROVISIO_TOKEN"
	// EnvServer overrides the auth service base URL at login.
	EnvServer = "PROVISIO_AUTH_URL"
	// EnvNoKeychain set to a non-empty value keeps credentials out of the OS keychain.
	EnvNoKeychain = "PROVISIO_NO_KEYCHAIN"
	// DefaultServer is the auth service used when nothing else is configured.
	DefaultServer = "https://auth.provisio.dev"
	// ClientID is what the CLI sends as client_id.
	ClientID = "provenance-cli"
)

// /sedum:setConst
// sedum:setVar {"tier":"owned","record":"prov-2026-2b4db381","gen":"sha256:45a2a4d28f38438bd03beb44de6ad9b5ee07bdad521a440fcf25bbc3ce684236","pkg":"sha256:c7b9d9ac1655e8211b8b2ca8c49cf34b8d510d8d823d4f5eb60ec7e1018aa2eb","kwargs":{"code":"var (\n\t// ErrNoSession means no credential is stored and PROVISIO_TOKEN is unset.\n\tErrNoSession = errors.New(\"no Provisio session\")\n\t// ErrSessionExpired means the refresh token was rejected or has expired.\n\tErrSessionExpired = errors.New(\"your Provisio session has expired\")\n\t// ErrUnauthorized means a bearer token was rejected with 401.\n\tErrUnauthorized = errors.New(\"bearer token rejected\")\n\t// ErrTokenRejected means a personal token (PROVISIO_TOKEN or stored) was rejected.\n\tErrTokenRejected = errors.New(\"personal token rejected\")\n\t// ErrAccessDenied means the user denied the device approval.\n\tErrAccessDenied = errors.New(\"sign-in was denied\")\n\t// ErrDeviceExpired means the device code expired or was already used.\n\tErrDeviceExpired = errors.New(\"the sign-in code expired\")\n)","file":"auth/credential.go","name":"ErrNoSession"}}
var (
	// ErrNoSession means no credential is stored and PROVISIO_TOKEN is unset.
	ErrNoSession = errors.New("no Provisio session")
	// ErrSessionExpired means the refresh token was rejected or has expired.
	ErrSessionExpired = errors.New("your Provisio session has expired")
	// ErrUnauthorized means a bearer token was rejected with 401.
	ErrUnauthorized = errors.New("bearer token rejected")
	// ErrTokenRejected means a personal token (PROVISIO_TOKEN or stored) was rejected.
	ErrTokenRejected = errors.New("personal token rejected")
	// ErrAccessDenied means the user denied the device approval.
	ErrAccessDenied = errors.New("sign-in was denied")
	// ErrDeviceExpired means the device code expired or was already used.
	ErrDeviceExpired = errors.New("the sign-in code expired")
)

// /sedum:setVar
// sedum:setType {"tier":"owned","record":"prov-2026-2b4db381","gen":"sha256:45a2a4d28f38438bd03beb44de6ad9b5ee07bdad521a440fcf25bbc3ce684236","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"// Credential is what the CLI keeps between runs. A personal token credential\n// (login --with-token) has only PersonalToken set.\ntype Credential struct {\n\tServer           string    `json:\"server\"`\n\tAccessToken      string    `json:\"access_token,omitempty\"`\n\tAccessExpiresAt  time.Time `json:\"access_expires_at,omitzero\"`\n\tRefreshToken     string    `json:\"refresh_token,omitempty\"`\n\tRefreshExpiresAt time.Time `json:\"refresh_expires_at,omitzero\"`\n\tPersonalToken    string    `json:\"personal_token,omitempty\"`\n\tDID              string    `json:\"did,omitempty\"`\n\tHandle           string    `json:\"handle,omitempty\"`\n}","file":"auth/credential.go","name":"Credential"}}
// Credential is what the CLI keeps between runs. A personal token credential
// (login --with-token) has only PersonalToken set.
type Credential struct {
	Server           string    `json:"server"`
	AccessToken      string    `json:"access_token,omitempty"`
	AccessExpiresAt  time.Time `json:"access_expires_at,omitzero"`
	RefreshToken     string    `json:"refresh_token,omitempty"`
	RefreshExpiresAt time.Time `json:"refresh_expires_at,omitzero"`
	PersonalToken    string    `json:"personal_token,omitempty"`
	DID              string    `json:"did,omitempty"`
	Handle           string    `json:"handle,omitempty"`
}

// /sedum:setType
// sedum:setMethod {"tier":"owned","record":"prov-2026-2b4db381","gen":"sha256:45a2a4d28f38438bd03beb44de6ad9b5ee07bdad521a440fcf25bbc3ce684236","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"// Secrets lists every secret string in the credential, for redaction tests.\nfunc (c *Credential) Secrets() []string {\n\tvar out []string\n\tfor _, s := range []string{c.AccessToken, c.RefreshToken, c.PersonalToken} {\n\t\tif s != \"\" {\n\t\t\tout = append(out, s)\n\t\t}\n\t}\n\treturn out\n}","file":"auth/credential.go","name":"Secrets","receiver":"Credential"}}
// Secrets lists every secret string in the credential, for redaction tests.
func (c *Credential) Secrets() []string {
	var out []string
	for _, s := range []string{c.AccessToken, c.RefreshToken, c.PersonalToken} {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

// /sedum:setMethod
