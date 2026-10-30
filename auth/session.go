package auth

// sedum:anchor:imports
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-494aca4f","gen":"sha256:358c44140aa8dc21a423adc09a5e69530354c12ba57e7d99bf6c8ff18621a4fa","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"auth/session.go","path":"time"},"kwargs":"sha256:973954a9e33c1bfb2a3d2cbf04331459a2f8fe0526ef53d107d3662f7b18cdc2","body":"sha256:e99128ebf4e36998f40b33adf2eef57f4d6c315f1a4d1bdd3b7fc09a0d7bc82d"}
import "time"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-5c3ad712","gen":"sha256:a78a005c60e30de8e657049b37039760864a35eb82809e77012051f0e597fc5e","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"auth/session.go","path":"context"},"kwargs":"sha256:e7ec63c9e439cc5b32e0e2aa431ceb603a4bb69d61c950255c6840a77246552a","body":"sha256:0f7443fb25d7ad6a970617d4d0df6123283e24801b94c1d84fbe431230f31e1e"}
import "context"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-5c3ad712","gen":"sha256:a78a005c60e30de8e657049b37039760864a35eb82809e77012051f0e597fc5e","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"auth/session.go","path":"os"},"kwargs":"sha256:4305b03bc9ffcc86257da655122865d7920addeaaffd9978a1cd65c97b19b215","body":"sha256:7c6e309eb2575d4aa451ec2ab83161d13e9ae03bbca3a4d2ecf30ce12adf0905"}
import "os"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-67038588","gen":"sha256:e7e337184a45eb6c41fdc7870ab8f62863196ebc13b36543f09e1ad1ebbd5179","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"auth/session.go","path":"errors"},"kwargs":"sha256:00e7704e2b184f5b343c1d19a5e35e24af97200c544564557996570d0e2d1c09","body":"sha256:2b718237ba0a56467c8fd494cc131e942ef410e6228c03b644d70476c71862dc"}
import "errors"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-8df1eced","gen":"sha256:f9ac5c5321e44bf2d79ad6c064b67e6938cde160fdeb8db361b6f701cccc1ef4","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"auth/session.go","path":"path/filepath"},"kwargs":"sha256:3e90cd4944322376778ba35986c06b752493a92232b5390fee95e87b3321e8d4","body":"sha256:a6d196b558491e646be93f56a321f339068c64c76aa9d49bb18046e7f9c777a1"}
import "path/filepath"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setType {"v":2,"tier":"owned","record":"prov-2026-210dc5fb","gen":"sha256:4b4901fd78535117a445b644e4d8e252d6d68dcba9bf698bfde7c8f6d8820981","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","id":{"file":"auth/session.go","name":"Token"},"kwargs":"sha256:ed1afbfa10473cc4ea84f2a461e76c4d18030a23c9be9236803d0c1778704efd","body":"sha256:b5ce0da95c11c3d8c1b68b27fa061b168c4eed3b7fbf431eb406000c0d6a537b"}
type Token struct {
	Bearer   string
	Source   Source
	Personal bool
}

// /sedum:setType
// sedum:setConst {"v":2,"tier":"owned","record":"prov-2026-494aca4f","gen":"sha256:358c44140aa8dc21a423adc09a5e69530354c12ba57e7d99bf6c8ff18621a4fa","pkg":"sha256:51be1c18998d63cc558c253e211a6bd188d96cb251dad68c6ea88e03c2d5cc4c","id":{"file":"auth/session.go","name":"refreshMargin"},"kwargs":"sha256:3ab7eb208da54a1ca0d82c6a49f75756552f88220b1d2dcec4e2ab6dc33ca554","body":"sha256:d1b4b6bc0009dbcb54a53433387fe256bb94c28e2d68d354cbb4a9aa7e5e4c8d"}
const refreshMargin = time.Minute

// /sedum:setConst
// sedum:setMethod {"v":2,"tier":"owned","record":"prov-2026-4a667d7d","gen":"sha256:aa5b1f0cf1c5bc42c291f8e28d19f6411b9a292ab486d902feb9d7ec3beb5bd6","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","id":{"file":"auth/session.go","name":"now","receiver":"Manager"},"kwargs":"sha256:c96bcbb5e69550190d64e7bf4105ce42eaafe86669a5eda46d93a68330fb8ea1","body":"sha256:0d0d8379a5e3ccc18416ce18cebb7a4372e4916f0549bb29ed2c2fe5c56175f7"}
func (m *Manager) now() time.Time {
	if m.Now != nil {
		return m.Now()
	}
	return time.Now()
}

// /sedum:setMethod
// sedum:setMethod {"v":2,"tier":"owned","record":"prov-2026-5c3ad712","gen":"sha256:a78a005c60e30de8e657049b37039760864a35eb82809e77012051f0e597fc5e","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","id":{"file":"auth/session.go","name":"Token","receiver":"Manager"},"kwargs":"sha256:1da71487285f666bf2b862b688cce2836e72bebeb779ce74a019bbef0f089184","body":"sha256:32a68a45d4c0c5678ed05e690ec6bec285ec4be863e3c612113cf3adb52c3f88"}
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
// sedum:setType {"v":2,"tier":"owned","record":"prov-2026-6102bb27","gen":"sha256:2bb1d937322abb869626e1fb040bd11b124570c1f850ee73c1908977e900dbbb","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","id":{"file":"auth/session.go","name":"Manager"},"kwargs":"sha256:b7db3d7ab208c2155b66636cc10c95f8ee67775b869451c696f81b49ef2f6f5e","body":"sha256:a518075d51cfc4cc29b81f42515c10c88e9b5164864675e7ed56e92be682f285"}
type Manager struct {
	Store    Store
	Client   *Client
	LockPath string
	Now      func() time.Time
}

// /sedum:setType
// sedum:setMethod {"v":2,"tier":"owned","record":"prov-2026-67038588","gen":"sha256:e7e337184a45eb6c41fdc7870ab8f62863196ebc13b36543f09e1ad1ebbd5179","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","id":{"file":"auth/session.go","name":"Do","receiver":"Manager"},"kwargs":"sha256:cd6e5f949cd5d600aac7537ecf0af53c5a8420585572fefb19610684d544fe36","body":"sha256:8cb36b69bae682fd38b688c9c97b993db08e0d19990bc5b91970236379e08c39"}
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
// sedum:setMethod {"v":2,"tier":"owned","record":"prov-2026-8df1eced","gen":"sha256:f9ac5c5321e44bf2d79ad6c064b67e6938cde160fdeb8db361b6f701cccc1ef4","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","id":{"file":"auth/session.go","name":"refresh","receiver":"Manager"},"kwargs":"sha256:e8c0d6c3456a664237f9bb7214154a824ad3e6747f72f82af3afd784a8e1be89","body":"sha256:3414c979f8602492180ce84363b72ac273e7a06b1a93c6c9eab6a61d348dd67e"}
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
// sedum:setMethod {"v":2,"tier":"owned","record":"prov-2026-eaab7d24","gen":"sha256:058545addd1b57effdf613a49ef3130740cbe70c19263620a7b650bc5ffe72bd","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","id":{"file":"auth/session.go","name":"fresh","receiver":"Manager"},"kwargs":"sha256:c93b9e6a93c79c6aaca46f7cff7fbc63e7814f4fbe8d98e8778d1bfe169794a2","body":"sha256:1fe33e545e7d4a2db2459a839fba86a1efb6069352e0bbc6f4b31fec6604024e"}
func (m *Manager) fresh(c *Credential) bool {
	return c.AccessToken != "" && m.now().Add(refreshMargin).Before(c.AccessExpiresAt)
}

// /sedum:setMethod
