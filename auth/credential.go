package auth

// sedum:anchor:imports
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-077d6686","gen":"sha256:a3e15192f7789f5c4a29dea3699c79a210d8ba94e5620e05f40213943d45b670","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"auth/credential.go","path":"time"},"kwargs":"sha256:483693e25e8810478a9a9ad81f70dcf06fbca4478ef91305e2eb024103b1824f","body":"sha256:e99128ebf4e36998f40b33adf2eef57f4d6c315f1a4d1bdd3b7fc09a0d7bc82d"}
import "time"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-b2b743a0","gen":"sha256:3302c36b8bd09922f911abcc90515b6215a18643e297eaf7ab2e6a7377034120","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"auth/credential.go","path":"errors"},"kwargs":"sha256:8b5877a30320776886833e37f0ebc28b10e155079798513656260f55ae063814","body":"sha256:2b718237ba0a56467c8fd494cc131e942ef410e6228c03b644d70476c71862dc"}
import "errors"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setType {"v":2,"tier":"owned","record":"prov-2026-2b4db381","gen":"sha256:81965c9da75304ccc85943d2198b5b7523c50c94d4b14a69ce7d531dfb8027e7","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","id":{"file":"auth/credential.go","name":"Source"},"kwargs":"sha256:38ea59ae59d570ec4a42a53ecffab405ecc79befe2d90550980d4c25f64ddea0","body":"sha256:808fcee12e9829d14d848c353120641ff4835f8c81222d9fcdd56c036ca460a0"}
type Source string

// /sedum:setType
// sedum:setType {"v":2,"tier":"owned","record":"prov-2026-077d6686","gen":"sha256:a3e15192f7789f5c4a29dea3699c79a210d8ba94e5620e05f40213943d45b670","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","id":{"file":"auth/credential.go","name":"Credential"},"kwargs":"sha256:36d2c3bba3d11432a233bfb3718b75085608b2dc6195bea5023781dd04d8e7d6","body":"sha256:924a071ff89c25657bec0d4b8c0de83d18e41123c76dfb19423cb74b0f75c1d6"}
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
// sedum:setMethod {"v":2,"tier":"owned","record":"prov-2026-438a44bc","gen":"sha256:f90f0ef6518bf5a0e83e8dc3fbcb801a5727c331f3cc49fa4dd1ea66580454e9","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","id":{"file":"auth/credential.go","name":"Secrets","receiver":"Credential"},"kwargs":"sha256:3381553be69fc0455dce9ed47012b4d6ac3bab26606205953f607f3ee5c41a64","body":"sha256:c3c97e2a4b6d64203abfb8fea62d682e59ae043fe0105b1ffeed3f77afb9525d"}
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
// sedum:setConst {"v":2,"tier":"owned","record":"prov-2026-94eaab57","gen":"sha256:4bbc186fb3b9e9d15e2248ef7a26332d28bc89ec8451ccaa22d42ef9f04e419e","pkg":"sha256:51be1c18998d63cc558c253e211a6bd188d96cb251dad68c6ea88e03c2d5cc4c","id":{"file":"auth/credential.go","name":"SourceEnv"},"kwargs":"sha256:8750367822e1d27b22c0be997cf3094f4c7dfdc350da318199de065a527a87c9","body":"sha256:b4426ddc803e1fe2b601d03a6e59cc24af0993bff81130bc1043ed5de251f074"}
const (
	SourceEnv      Source = "env"
	SourceKeychain Source = "keychain"
	SourceFile     Source = "file"
)

// /sedum:setConst
// sedum:setVar {"v":2,"tier":"owned","record":"prov-2026-b2b743a0","gen":"sha256:3302c36b8bd09922f911abcc90515b6215a18643e297eaf7ab2e6a7377034120","pkg":"sha256:c7b9d9ac1655e8211b8b2ca8c49cf34b8d510d8d823d4f5eb60ec7e1018aa2eb","id":{"file":"auth/credential.go","name":"ErrNoSession"},"kwargs":"sha256:b33f252883638f79aa6dc2371409ffb87e6c9b034b2d4b32b133526f8d6cb494","body":"sha256:a384ccb4fc8f88e93e4d2e1c2786da34ac5970d498b06b2248555c31281e93d2"}
var (
	ErrNoSession      = errors.New("no Provisio session")
	ErrSessionExpired = errors.New("your Provisio session has expired")
	ErrUnauthorized   = errors.New("bearer token rejected")
	ErrTokenRejected  = errors.New("personal token rejected")
	ErrAccessDenied   = errors.New("sign-in was denied")
	ErrDeviceExpired  = errors.New("the sign-in code expired")
)

// /sedum:setVar
// sedum:setConst {"v":2,"tier":"owned","record":"prov-2026-cb6e6e2e","gen":"sha256:7896e8e8c36d0d66fe8d755ca5bcb76fc95e204b8251cb5ce5c7b79cdf56c843","pkg":"sha256:51be1c18998d63cc558c253e211a6bd188d96cb251dad68c6ea88e03c2d5cc4c","id":{"file":"auth/credential.go","name":"EnvToken"},"kwargs":"sha256:f57c884c7f6ad46b7e59dd856ad01ae87785b104d92c60613479966413bd8ddb","body":"sha256:53bbcc747c2410bee6722a410fdabf7b55e45afbc55c8719b18c261a0ad6d825"}
const (
	EnvToken      = "PROVISIO_TOKEN"
	EnvServer     = "PROVISIO_AUTH_URL"
	EnvNoKeychain = "PROVISIO_NO_KEYCHAIN"
	DefaultServer = "https://auth.provisio.dev"
	ClientID      = "provenance-cli"
)

// /sedum:setConst
