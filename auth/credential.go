package auth

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-077d6686","gen":"sha256:a3e15192f7789f5c4a29dea3699c79a210d8ba94e5620e05f40213943d45b670","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"auth/credential.go","path":"time"}}
import "time"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-b2b743a0","gen":"sha256:3302c36b8bd09922f911abcc90515b6215a18643e297eaf7ab2e6a7377034120","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"auth/credential.go","path":"errors"}}
import "errors"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setType {"tier":"owned","record":"prov-2026-2b4db381","gen":"sha256:81965c9da75304ccc85943d2198b5b7523c50c94d4b14a69ce7d531dfb8027e7","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type Source string","file":"auth/credential.go","name":"Source"}}
type Source string

// /sedum:setType
// sedum:setType {"tier":"owned","record":"prov-2026-077d6686","gen":"sha256:a3e15192f7789f5c4a29dea3699c79a210d8ba94e5620e05f40213943d45b670","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type Credential struct {\n\tServer           string    `json:\"server\"`\n\tAccessToken      string    `json:\"access_token,omitempty\"`\n\tAccessExpiresAt  time.Time `json:\"access_expires_at,omitzero\"`\n\tRefreshToken     string    `json:\"refresh_token,omitempty\"`\n\tRefreshExpiresAt time.Time `json:\"refresh_expires_at,omitzero\"`\n\tPersonalToken    string    `json:\"personal_token,omitempty\"`\n\tDID              string    `json:\"did,omitempty\"`\n\tHandle           string    `json:\"handle,omitempty\"`\n}","file":"auth/credential.go","name":"Credential"}}
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
// sedum:setMethod {"tier":"owned","record":"prov-2026-438a44bc","gen":"sha256:f90f0ef6518bf5a0e83e8dc3fbcb801a5727c331f3cc49fa4dd1ea66580454e9","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (c *Credential) Secrets() []string {\n\tvar out []string\n\tif c.AccessToken != \"\" {\n\t\tout = append(out, c.AccessToken)\n\t}\n\tif c.RefreshToken != \"\" {\n\t\tout = append(out, c.RefreshToken)\n\t}\n\tif c.PersonalToken != \"\" {\n\t\tout = append(out, c.PersonalToken)\n\t}\n\treturn out\n}","file":"auth/credential.go","name":"Secrets","receiver":"Credential"}}
func (c *Credential) Secrets() []string {
	var out []string
	if c.AccessToken != "" {
		out = append(out, c.AccessToken)
	}
	if c.RefreshToken != "" {
		out = append(out, c.RefreshToken)
	}
	if c.PersonalToken != "" {
		out = append(out, c.PersonalToken)
	}
	return out
}

// /sedum:setMethod
// sedum:setConst {"tier":"owned","record":"prov-2026-94eaab57","gen":"sha256:4bbc186fb3b9e9d15e2248ef7a26332d28bc89ec8451ccaa22d42ef9f04e419e","pkg":"sha256:51be1c18998d63cc558c253e211a6bd188d96cb251dad68c6ea88e03c2d5cc4c","kwargs":{"code":"const (\n\tSourceEnv      Source = \"env\"\n\tSourceKeychain Source = \"keychain\"\n\tSourceFile     Source = \"file\"\n)","file":"auth/credential.go","name":"SourceEnv"}}
const (
	SourceEnv      Source = "env"
	SourceKeychain Source = "keychain"
	SourceFile     Source = "file"
)

// /sedum:setConst
// sedum:setVar {"tier":"owned","record":"prov-2026-b2b743a0","gen":"sha256:3302c36b8bd09922f911abcc90515b6215a18643e297eaf7ab2e6a7377034120","pkg":"sha256:c7b9d9ac1655e8211b8b2ca8c49cf34b8d510d8d823d4f5eb60ec7e1018aa2eb","kwargs":{"code":"var (\n\tErrNoSession      = errors.New(\"no Provisio session\")\n\tErrSessionExpired = errors.New(\"your Provisio session has expired\")\n\tErrUnauthorized   = errors.New(\"bearer token rejected\")\n\tErrTokenRejected  = errors.New(\"personal token rejected\")\n\tErrAccessDenied   = errors.New(\"sign-in was denied\")\n\tErrDeviceExpired  = errors.New(\"the sign-in code expired\")\n)","file":"auth/credential.go","name":"ErrNoSession"}}
var (
	ErrNoSession      = errors.New("no Provisio session")
	ErrSessionExpired = errors.New("your Provisio session has expired")
	ErrUnauthorized   = errors.New("bearer token rejected")
	ErrTokenRejected  = errors.New("personal token rejected")
	ErrAccessDenied   = errors.New("sign-in was denied")
	ErrDeviceExpired  = errors.New("the sign-in code expired")
)

// /sedum:setVar
// sedum:setConst {"tier":"owned","record":"prov-2026-cb6e6e2e","gen":"sha256:7896e8e8c36d0d66fe8d755ca5bcb76fc95e204b8251cb5ce5c7b79cdf56c843","pkg":"sha256:51be1c18998d63cc558c253e211a6bd188d96cb251dad68c6ea88e03c2d5cc4c","kwargs":{"code":"const (\n\tEnvToken      = \"PROVISIO_TOKEN\"\n\tEnvServer     = \"PROVISIO_AUTH_URL\"\n\tEnvNoKeychain = \"PROVISIO_NO_KEYCHAIN\"\n\tDefaultServer = \"https://auth.provisio.dev\"\n\tClientID      = \"provenance-cli\"\n)","file":"auth/credential.go","name":"EnvToken"}}
const (
	EnvToken      = "PROVISIO_TOKEN"
	EnvServer     = "PROVISIO_AUTH_URL"
	EnvNoKeychain = "PROVISIO_NO_KEYCHAIN"
	DefaultServer = "https://auth.provisio.dev"
	ClientID      = "provenance-cli"
)

// /sedum:setConst
