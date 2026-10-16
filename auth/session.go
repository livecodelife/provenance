package auth

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-ad56d5ca","gen":"sha256:b64ac5dd3402fa436945f27da8b32d59b0a8fc253e6d8ddb0e27f4d9b585565b","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"auth/session.go","path":"context"}}
import "context"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-ad56d5ca","gen":"sha256:b64ac5dd3402fa436945f27da8b32d59b0a8fc253e6d8ddb0e27f4d9b585565b","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"auth/session.go","path":"errors"}}
import "errors"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-ad56d5ca","gen":"sha256:b64ac5dd3402fa436945f27da8b32d59b0a8fc253e6d8ddb0e27f4d9b585565b","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"auth/session.go","path":"os"}}
import "os"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-ad56d5ca","gen":"sha256:b64ac5dd3402fa436945f27da8b32d59b0a8fc253e6d8ddb0e27f4d9b585565b","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"auth/session.go","path":"path/filepath"}}
import "path/filepath"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-ad56d5ca","gen":"sha256:b64ac5dd3402fa436945f27da8b32d59b0a8fc253e6d8ddb0e27f4d9b585565b","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"auth/session.go","path":"time"}}
import "time"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setConst {"tier":"owned","record":"prov-2026-ad56d5ca","gen":"sha256:b64ac5dd3402fa436945f27da8b32d59b0a8fc253e6d8ddb0e27f4d9b585565b","pkg":"sha256:51be1c18998d63cc558c253e211a6bd188d96cb251dad68c6ea88e03c2d5cc4c","kwargs":{"code":"// refreshMargin: refresh when the access token expires sooner than this.\nconst refreshMargin = time.Minute","file":"auth/session.go","name":"refreshMargin"}}
// refreshMargin: refresh when the access token expires sooner than this.
const refreshMargin = time.Minute

// /sedum:setConst
// sedum:setType {"tier":"owned","record":"prov-2026-ad56d5ca","gen":"sha256:b64ac5dd3402fa436945f27da8b32d59b0a8fc253e6d8ddb0e27f4d9b585565b","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"// Token is a bearer token ready to use and where it came from.\ntype Token struct {\n\tBearer string\n\tSource Source\n\t// Personal is true for PROVISIO_TOKEN and stored personal tokens: never refreshed.\n\tPersonal bool\n}","file":"auth/session.go","name":"Token"}}
// Token is a bearer token ready to use and where it came from.
type Token struct {
	Bearer string
	Source Source
	// Personal is true for PROVISIO_TOKEN and stored personal tokens: never refreshed.
	Personal bool
}

// /sedum:setType
// sedum:setType {"tier":"owned","record":"prov-2026-ad56d5ca","gen":"sha256:b64ac5dd3402fa436945f27da8b32d59b0a8fc253e6d8ddb0e27f4d9b585565b","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"// Manager hands out valid bearer tokens, refreshing under a lock file.\ntype Manager struct {\n\tStore    Store\n\tClient   *Client\n\tLockPath string\n\tNow      func() time.Time\n}","file":"auth/session.go","name":"Manager"}}
// Manager hands out valid bearer tokens, refreshing under a lock file.
type Manager struct {
	Store    Store
	Client   *Client
	LockPath string
	Now      func() time.Time
}

// /sedum:setType
// sedum:setMethod {"tier":"owned","record":"prov-2026-ad56d5ca","gen":"sha256:b64ac5dd3402fa436945f27da8b32d59b0a8fc253e6d8ddb0e27f4d9b585565b","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (m *Manager) now() time.Time {\n\tif m.Now != nil {\n\t\treturn m.Now()\n\t}\n\treturn time.Now()\n}","file":"auth/session.go","name":"now","receiver":"Manager"}}
func (m *Manager) now() time.Time {
	if m.Now != nil {
		return m.Now()
	}
	return time.Now()
}

// /sedum:setMethod
// sedum:setMethod {"tier":"owned","record":"prov-2026-ad56d5ca","gen":"sha256:b64ac5dd3402fa436945f27da8b32d59b0a8fc253e6d8ddb0e27f4d9b585565b","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (m *Manager) fresh(c *Credential) bool {\n\treturn c.AccessToken != \"\" && m.now().Add(refreshMargin).Before(c.AccessExpiresAt)\n}","file":"auth/session.go","name":"fresh","receiver":"Manager"}}
func (m *Manager) fresh(c *Credential) bool {
	return c.AccessToken != "" && m.now().Add(refreshMargin).Before(c.AccessExpiresAt)
}

// /sedum:setMethod
// sedum:setMethod {"tier":"owned","record":"prov-2026-ad56d5ca","gen":"sha256:b64ac5dd3402fa436945f27da8b32d59b0a8fc253e6d8ddb0e27f4d9b585565b","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"// Token returns a bearer token. PROVISIO_TOKEN wins and is used as is. Local\n// commands must not call this; only commands that need a session do.\nfunc (m *Manager) Token(ctx context.Context) (*Token, error) {\n\tif v := os.Getenv(EnvToken); v != \"\" {\n\t\treturn &Token{Bearer: v, Source: SourceEnv, Personal: true}, nil\n\t}\n\tcred, src, err := m.Store.Load()\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\tif cred.PersonalToken != \"\" {\n\t\treturn &Token{Bearer: cred.PersonalToken, Source: src, Personal: true}, nil\n\t}\n\tif m.fresh(cred) {\n\t\treturn &Token{Bearer: cred.AccessToken, Source: src}, nil\n\t}\n\treturn m.refresh(ctx, cred.AccessToken)\n}","file":"auth/session.go","name":"Token","receiver":"Manager"}}
// Token returns a bearer token. PROVISIO_TOKEN wins and is used as is. Local
// commands must not call this; only commands that need a session do.
func (m *Manager) Token(ctx context.Context) (*Token, error) {
	if v := os.Getenv(EnvToken); v != "" {
		return &Token{Bearer: v, Source: SourceEnv, Personal: true}, nil
	}
	cred, src, err := m.Store.Load()
	if err != nil {
		return nil, err
	}
	if cred.PersonalToken != "" {
		return &Token{Bearer: cred.PersonalToken, Source: src, Personal: true}, nil
	}
	if m.fresh(cred) {
		return &Token{Bearer: cred.AccessToken, Source: src}, nil
	}
	return m.refresh(ctx, cred.AccessToken)
}

// /sedum:setMethod
// sedum:setMethod {"tier":"owned","record":"prov-2026-ad56d5ca","gen":"sha256:b64ac5dd3402fa436945f27da8b32d59b0a8fc253e6d8ddb0e27f4d9b585565b","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"// refresh renews the session. rejected is the access token the caller found\n// unusable; if another process already replaced it, that one is returned.\nfunc (m *Manager) refresh(ctx context.Context, rejected string) (*Token, error) {\n\tif m.LockPath != \"\" {\n\t\tif err := os.MkdirAll(filepath.Dir(m.LockPath), 0o700); err != nil {\n\t\t\treturn nil, err\n\t\t}\n\t\tunlock, err := lockFile(m.LockPath)\n\t\tif err != nil {\n\t\t\treturn nil, err\n\t\t}\n\t\tdefer unlock()\n\t}\n\tcred, src, err := m.Store.Load()\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\tif cred.AccessToken != rejected && m.fresh(cred) {\n\t\treturn &Token{Bearer: cred.AccessToken, Source: src}, nil\n\t}\n\tif cred.RefreshToken == \"\" || (!cred.RefreshExpiresAt.IsZero() && m.now().After(cred.RefreshExpiresAt)) {\n\t\treturn nil, ErrSessionExpired\n\t}\n\tclient := *m.Client\n\tif cred.Server != \"\" {\n\t\tclient.Server = cred.Server\n\t}\n\ttr, err := client.Refresh(ctx, cred.RefreshToken)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\tnow := m.now()\n\tnext := *cred\n\tnext.AccessToken = tr.AccessToken\n\tnext.AccessExpiresAt = now.Add(time.Duration(tr.ExpiresIn) * time.Second)\n\tnext.RefreshToken = tr.RefreshToken\n\tnext.RefreshExpiresAt = now.Add(time.Duration(tr.RefreshExpiresIn) * time.Second)\n\tif tr.DID != \"\" {\n\t\tnext.DID, next.Handle = tr.DID, tr.Handle\n\t}\n\tif err := m.Store.Save(&next); err != nil {\n\t\treturn nil, err\n\t}\n\treturn &Token{Bearer: next.AccessToken, Source: src}, nil\n}","file":"auth/session.go","name":"refresh","receiver":"Manager"}}
// refresh renews the session. rejected is the access token the caller found
// unusable; if another process already replaced it, that one is returned.
func (m *Manager) refresh(ctx context.Context, rejected string) (*Token, error) {
	if m.LockPath != "" {
		if err := os.MkdirAll(filepath.Dir(m.LockPath), 0o700); err != nil {
			return nil, err
		}
		unlock, err := lockFile(m.LockPath)
		if err != nil {
			return nil, err
		}
		defer unlock()
	}
	cred, src, err := m.Store.Load()
	if err != nil {
		return nil, err
	}
	if cred.AccessToken != rejected && m.fresh(cred) {
		return &Token{Bearer: cred.AccessToken, Source: src}, nil
	}
	if cred.RefreshToken == "" || (!cred.RefreshExpiresAt.IsZero() && m.now().After(cred.RefreshExpiresAt)) {
		return nil, ErrSessionExpired
	}
	client := *m.Client
	if cred.Server != "" {
		client.Server = cred.Server
	}
	tr, err := client.Refresh(ctx, cred.RefreshToken)
	if err != nil {
		return nil, err
	}
	now := m.now()
	next := *cred
	next.AccessToken = tr.AccessToken
	next.AccessExpiresAt = now.Add(time.Duration(tr.ExpiresIn) * time.Second)
	next.RefreshToken = tr.RefreshToken
	next.RefreshExpiresAt = now.Add(time.Duration(tr.RefreshExpiresIn) * time.Second)
	if tr.DID != "" {
		next.DID, next.Handle = tr.DID, tr.Handle
	}
	if err := m.Store.Save(&next); err != nil {
		return nil, err
	}
	return &Token{Bearer: next.AccessToken, Source: src}, nil
}

// /sedum:setMethod
// sedum:setMethod {"tier":"owned","record":"prov-2026-ad56d5ca","gen":"sha256:b64ac5dd3402fa436945f27da8b32d59b0a8fc253e6d8ddb0e27f4d9b585565b","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"// Do runs fn with a bearer token. fn returns ErrUnauthorized on a 401; Do then\n// refreshes once and runs fn once more. A rejected personal token is\n// ErrTokenRejected.\nfunc (m *Manager) Do(ctx context.Context, fn func(bearer string) error) error {\n\ttok, err := m.Token(ctx)\n\tif err != nil {\n\t\treturn err\n\t}\n\terr = fn(tok.Bearer)\n\tif !errors.Is(err, ErrUnauthorized) {\n\t\treturn err\n\t}\n\tif tok.Personal {\n\t\treturn ErrTokenRejected\n\t}\n\ttok, err = m.refresh(ctx, tok.Bearer)\n\tif err != nil {\n\t\treturn err\n\t}\n\treturn fn(tok.Bearer)\n}","file":"auth/session.go","name":"Do","receiver":"Manager"}}
// Do runs fn with a bearer token. fn returns ErrUnauthorized on a 401; Do then
// refreshes once and runs fn once more. A rejected personal token is
// ErrTokenRejected.
func (m *Manager) Do(ctx context.Context, fn func(bearer string) error) error {
	tok, err := m.Token(ctx)
	if err != nil {
		return err
	}
	err = fn(tok.Bearer)
	if !errors.Is(err, ErrUnauthorized) {
		return err
	}
	if tok.Personal {
		return ErrTokenRejected
	}
	tok, err = m.refresh(ctx, tok.Bearer)
	if err != nil {
		return err
	}
	return fn(tok.Bearer)
}

// /sedum:setMethod
