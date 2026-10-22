package auth

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-494aca4f","gen":"sha256:358c44140aa8dc21a423adc09a5e69530354c12ba57e7d99bf6c8ff18621a4fa","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"auth/session.go","path":"time"}}
import "time"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-5c3ad712","gen":"sha256:a78a005c60e30de8e657049b37039760864a35eb82809e77012051f0e597fc5e","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"auth/session.go","path":"context"}}
import "context"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-5c3ad712","gen":"sha256:a78a005c60e30de8e657049b37039760864a35eb82809e77012051f0e597fc5e","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"auth/session.go","path":"os"}}
import "os"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-67038588","gen":"sha256:e7e337184a45eb6c41fdc7870ab8f62863196ebc13b36543f09e1ad1ebbd5179","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"auth/session.go","path":"errors"}}
import "errors"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-8df1eced","gen":"sha256:f9ac5c5321e44bf2d79ad6c064b67e6938cde160fdeb8db361b6f701cccc1ef4","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"auth/session.go","path":"path/filepath"}}
import "path/filepath"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setType {"tier":"owned","record":"prov-2026-210dc5fb","gen":"sha256:4b4901fd78535117a445b644e4d8e252d6d68dcba9bf698bfde7c8f6d8820981","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type Token struct {\n\tBearer   string\n\tSource   Source\n\tPersonal bool\n}","file":"auth/session.go","name":"Token"}}
type Token struct {
	Bearer   string
	Source   Source
	Personal bool
}

// /sedum:setType
// sedum:setConst {"tier":"owned","record":"prov-2026-494aca4f","gen":"sha256:358c44140aa8dc21a423adc09a5e69530354c12ba57e7d99bf6c8ff18621a4fa","pkg":"sha256:51be1c18998d63cc558c253e211a6bd188d96cb251dad68c6ea88e03c2d5cc4c","kwargs":{"code":"const refreshMargin = time.Minute","file":"auth/session.go","name":"refreshMargin"}}
const refreshMargin = time.Minute

// /sedum:setConst
// sedum:setMethod {"tier":"owned","record":"prov-2026-4a667d7d","gen":"sha256:aa5b1f0cf1c5bc42c291f8e28d19f6411b9a292ab486d902feb9d7ec3beb5bd6","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (m *Manager) now() time.Time {\n\tif m.Now != nil {\n\t\treturn m.Now()\n\t}\n\treturn time.Now()\n}","file":"auth/session.go","name":"now","receiver":"Manager"}}
func (m *Manager) now() time.Time {
	if m.Now != nil {
		return m.Now()
	}
	return time.Now()
}

// /sedum:setMethod
// sedum:setMethod {"tier":"owned","record":"prov-2026-5c3ad712","gen":"sha256:a78a005c60e30de8e657049b37039760864a35eb82809e77012051f0e597fc5e","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (m *Manager) Token(ctx context.Context) (*Token, error) {\n\tif v := os.Getenv(EnvToken); v != \"\" {\n\t\treturn &Token{Bearer: v, Source: SourceEnv, Personal: true}, nil\n\t}\n\tcred, src, err := m.Store.Load()\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\tif cred.PersonalToken != \"\" {\n\t\treturn &Token{Bearer: cred.PersonalToken, Source: src, Personal: true}, nil\n\t}\n\tif m.fresh(cred) {\n\t\treturn &Token{Bearer: cred.AccessToken, Source: src, Personal: false}, nil\n\t}\n\treturn m.refresh(ctx, cred.AccessToken)\n}","file":"auth/session.go","name":"Token","receiver":"Manager"}}
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
		return &Token{Bearer: cred.AccessToken, Source: src, Personal: false}, nil
	}
	return m.refresh(ctx, cred.AccessToken)
}

// /sedum:setMethod
// sedum:setType {"tier":"owned","record":"prov-2026-6102bb27","gen":"sha256:2bb1d937322abb869626e1fb040bd11b124570c1f850ee73c1908977e900dbbb","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type Manager struct {\n\tStore    Store\n\tClient   *Client\n\tLockPath string\n\tNow      func() time.Time\n}","file":"auth/session.go","name":"Manager"}}
type Manager struct {
	Store    Store
	Client   *Client
	LockPath string
	Now      func() time.Time
}

// /sedum:setType
// sedum:setMethod {"tier":"owned","record":"prov-2026-67038588","gen":"sha256:e7e337184a45eb6c41fdc7870ab8f62863196ebc13b36543f09e1ad1ebbd5179","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (m *Manager) Do(ctx context.Context, fn func(bearer string) error) error {\n\ttok, err := m.Token(ctx)\n\tif err != nil {\n\t\treturn err\n\t}\n\terr = fn(tok.Bearer)\n\tif !errors.Is(err, ErrUnauthorized) {\n\t\treturn err\n\t}\n\tif tok.Personal {\n\t\treturn ErrTokenRejected\n\t}\n\tfresh, err := m.refresh(ctx, tok.Bearer)\n\tif err != nil {\n\t\treturn err\n\t}\n\treturn fn(fresh.Bearer)\n}","file":"auth/session.go","name":"Do","receiver":"Manager"}}
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
	fresh, err := m.refresh(ctx, tok.Bearer)
	if err != nil {
		return err
	}
	return fn(fresh.Bearer)
}

// /sedum:setMethod
// sedum:setMethod {"tier":"owned","record":"prov-2026-8df1eced","gen":"sha256:f9ac5c5321e44bf2d79ad6c064b67e6938cde160fdeb8db361b6f701cccc1ef4","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (m *Manager) refresh(ctx context.Context, rejected string) (*Token, error) {\n\tif m.LockPath != \"\" {\n\t\tif err := os.MkdirAll(filepath.Dir(m.LockPath), 0700); err != nil {\n\t\t\treturn nil, err\n\t\t}\n\t\tunlock, err := lockFile(m.LockPath)\n\t\tif err != nil {\n\t\t\treturn nil, err\n\t\t}\n\t\tdefer unlock()\n\t}\n\n\tcred, source, err := m.Store.Load()\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\n\tif cred.AccessToken != rejected && m.fresh(cred) {\n\t\treturn &Token{Bearer: cred.AccessToken, Source: source, Personal: false}, nil\n\t}\n\n\tif cred.RefreshToken == \"\" {\n\t\treturn nil, ErrSessionExpired\n\t}\n\tif !cred.RefreshExpiresAt.IsZero() && m.now().After(cred.RefreshExpiresAt) {\n\t\treturn nil, ErrSessionExpired\n\t}\n\n\tclient := *m.Client\n\tif cred.Server != \"\" {\n\t\tclient.Server = cred.Server\n\t}\n\n\tresp, err := client.Refresh(ctx, cred.RefreshToken)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\n\tnow := m.now()\n\tupdated := *cred\n\tupdated.AccessToken = resp.AccessToken\n\tupdated.AccessExpiresAt = now.Add(time.Duration(resp.ExpiresIn) * time.Second)\n\tupdated.RefreshToken = resp.RefreshToken\n\tupdated.RefreshExpiresAt = now.Add(time.Duration(resp.RefreshExpiresIn) * time.Second)\n\tif resp.DID != \"\" {\n\t\tupdated.DID = resp.DID\n\t\tupdated.Handle = resp.Handle\n\t}\n\n\tif err := m.Store.Save(&updated); err != nil {\n\t\treturn nil, err\n\t}\n\n\treturn &Token{Bearer: updated.AccessToken, Source: source, Personal: false}, nil\n}","file":"auth/session.go","name":"refresh","receiver":"Manager"}}
func (m *Manager) refresh(ctx context.Context, rejected string) (*Token, error) {
	if m.LockPath != "" {
		if err := os.MkdirAll(filepath.Dir(m.LockPath), 0700); err != nil {
			return nil, err
		}
		unlock, err := lockFile(m.LockPath)
		if err != nil {
			return nil, err
		}
		defer unlock()
	}

	cred, source, err := m.Store.Load()
	if err != nil {
		return nil, err
	}

	if cred.AccessToken != rejected && m.fresh(cred) {
		return &Token{Bearer: cred.AccessToken, Source: source, Personal: false}, nil
	}

	if cred.RefreshToken == "" {
		return nil, ErrSessionExpired
	}
	if !cred.RefreshExpiresAt.IsZero() && m.now().After(cred.RefreshExpiresAt) {
		return nil, ErrSessionExpired
	}

	client := *m.Client
	if cred.Server != "" {
		client.Server = cred.Server
	}

	resp, err := client.Refresh(ctx, cred.RefreshToken)
	if err != nil {
		return nil, err
	}

	now := m.now()
	updated := *cred
	updated.AccessToken = resp.AccessToken
	updated.AccessExpiresAt = now.Add(time.Duration(resp.ExpiresIn) * time.Second)
	updated.RefreshToken = resp.RefreshToken
	updated.RefreshExpiresAt = now.Add(time.Duration(resp.RefreshExpiresIn) * time.Second)
	if resp.DID != "" {
		updated.DID = resp.DID
		updated.Handle = resp.Handle
	}

	if err := m.Store.Save(&updated); err != nil {
		return nil, err
	}

	return &Token{Bearer: updated.AccessToken, Source: source, Personal: false}, nil
}

// /sedum:setMethod
// sedum:setMethod {"tier":"owned","record":"prov-2026-eaab7d24","gen":"sha256:058545addd1b57effdf613a49ef3130740cbe70c19263620a7b650bc5ffe72bd","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (m *Manager) fresh(c *Credential) bool {\n\treturn c.AccessToken != \"\" && m.now().Add(refreshMargin).Before(c.AccessExpiresAt)\n}","file":"auth/session.go","name":"fresh","receiver":"Manager"}}
func (m *Manager) fresh(c *Credential) bool {
	return c.AccessToken != "" && m.now().Add(refreshMargin).Before(c.AccessExpiresAt)
}

// /sedum:setMethod
