package main

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-81333e45","gen":"sha256:d3d637e42ac358905da0d78c2d2a3cc3e324f2d89cca02317556039be91b66b1","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/session_run.go","path":"github.com/livecodelife/provenance/auth"}}
import "github.com/livecodelife/provenance/auth"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-81333e45","gen":"sha256:d3d637e42ac358905da0d78c2d2a3cc3e324f2d89cca02317556039be91b66b1","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/session_run.go","path":"github.com/livecodelife/provenance/remote"}}
import "github.com/livecodelife/provenance/remote"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-a0ad0a7f","gen":"sha256:dd29d8cc1590b0342455e97aa87320359bfbfee4affbf741dcc86c80a497061c","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/session_run.go","path":"context"}}
import "context"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-a0ad0a7f","gen":"sha256:dd29d8cc1590b0342455e97aa87320359bfbfee4affbf741dcc86c80a497061c","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/session_run.go","path":"time"}}
import "time"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-a51283a0","gen":"sha256:25743a0c40a8eff1905ae8c89926e35e8eb0ccddd6b35fc4be3bc508c591cfa4","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/session_run.go","path":"os"}}
import "os"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-ae5b9ab4","gen":"sha256:906ad1c7ebf364264b0d30e9a0453b3e9ec1897578d2c9a4f4f1e2c07cc1787a","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/session_run.go","path":"errors"}}
import "errors"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-da5d8cb8","gen":"sha256:1802c6e80df79a3253b8c10809eb08ec0bc9b3e9baaf85e48a0ef0902dc6ab1d","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/session_run.go","path":"path/filepath"}}
import "path/filepath"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"tier":"owned","record":"prov-2026-81333e45","gen":"sha256:d3d637e42ac358905da0d78c2d2a3cc3e324f2d89cca02317556039be91b66b1","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func sessionError(err error) error {\n\tif err == nil {\n\t\treturn nil\n\t}\n\tif errors.Is(err, auth.ErrNoSession) {\n\t\treturn &exitError{code: 3, err: errors.New(msgNeedSession)}\n\t}\n\tif errors.Is(err, auth.ErrSessionExpired) {\n\t\treturn &exitError{code: 3, err: errors.New(msgSessionExpire)}\n\t}\n\tif errors.Is(err, auth.ErrTokenRejected) {\n\t\treturn &exitError{code: 3, err: errors.New(msgTokenRejected)}\n\t}\n\tvar apiErr *remote.APIError\n\tif errors.As(err, &apiErr) && apiErr.Status == 403 && apiErr.Code == \"insufficient_scope\" {\n\t\treturn &exitError{code: 3, err: errors.New(msgScope)}\n\t}\n\treturn err\n}","file":"cmd/provenance/session_run.go","name":"sessionError"}}
func sessionError(err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, auth.ErrNoSession) {
		return &exitError{code: 3, err: errors.New(msgNeedSession)}
	}
	if errors.Is(err, auth.ErrSessionExpired) {
		return &exitError{code: 3, err: errors.New(msgSessionExpire)}
	}
	if errors.Is(err, auth.ErrTokenRejected) {
		return &exitError{code: 3, err: errors.New(msgTokenRejected)}
	}
	var apiErr *remote.APIError
	if errors.As(err, &apiErr) && apiErr.Status == 403 && apiErr.Code == "insufficient_scope" {
		return &exitError{code: 3, err: errors.New(msgScope)}
	}
	return err
}

// /sedum:setFunction
// sedum:setMethod {"tier":"owned","record":"prov-2026-15de5191","gen":"sha256:e6a1b0e7ffcfb8066b5bd1951009d4137af9f4681a055e66c29b2716aa92e6f6","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (e *exitError) ExitCode() int {\n\treturn e.code\n}","file":"cmd/provenance/session_run.go","name":"ExitCode","receiver":"exitError"}}
func (e *exitError) ExitCode() int {
	return e.code
}

// /sedum:setMethod
// sedum:setFunction {"tier":"owned","record":"prov-2026-28d01dcf","gen":"sha256:e6659a3a24bb52c37733b2e89ed757564871e737a0e65d5606ac6f05b2fe6d77","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func withSession(ctx context.Context, fn func(string) error) error {\n\tm, err := sessionManager()\n\tif err != nil {\n\t\treturn err\n\t}\n\treturn sessionError(m.Do(ctx, fn))\n}","file":"cmd/provenance/session_run.go","name":"withSession"}}
func withSession(ctx context.Context, fn func(string) error) error {
	m, err := sessionManager()
	if err != nil {
		return err
	}
	return sessionError(m.Do(ctx, fn))
}

// /sedum:setFunction
// sedum:setMethod {"tier":"owned","record":"prov-2026-5a7e47c4","gen":"sha256:27b426390755927f3e455f7d4bdb1ff0cf4d4e70d94e004fd28d34526f6a9b5a","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (e *exitError) Unwrap() error {\n\treturn e.err\n}","file":"cmd/provenance/session_run.go","name":"Unwrap","receiver":"exitError"}}
func (e *exitError) Unwrap() error {
	return e.err
}

// /sedum:setMethod
// sedum:setMethod {"tier":"owned","record":"prov-2026-790e16a4","gen":"sha256:a9bad494046cdb251f9dad0b3915a76553b5a74995da5e61ed10251a9e0c5620","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (e *exitError) Error() string {\n\treturn e.err.Error()\n}","file":"cmd/provenance/session_run.go","name":"Error","receiver":"exitError"}}
func (e *exitError) Error() string {
	return e.err.Error()
}

// /sedum:setMethod
// sedum:setConst {"tier":"owned","record":"prov-2026-9895ce2c","gen":"sha256:e94a96baa86aacca7cd4be62ba3b7445cf3692cbc25e23d6cee2f6b0f5c4e5d2","pkg":"sha256:51be1c18998d63cc558c253e211a6bd188d96cb251dad68c6ea88e03c2d5cc4c","kwargs":{"code":"const (\n\tmsgNeedSession   = \"This command needs a Provisio session. Run provenance login (or set PROVISIO_TOKEN in CI).\"\n\tmsgSessionExpire = \"Your Provisio session has expired. Run provenance login.\"\n\tmsgTokenRejected = \"The personal token was rejected (revoked or expired). Replace PROVISIO_TOKEN, or run provenance login.\"\n\tmsgScope         = \"Your Provisio sign-in lacks permission to write; sign in again with `provenance login`.\"\n)","file":"cmd/provenance/session_run.go","name":"msgNeedSession"}}
const (
	msgNeedSession   = "This command needs a Provisio session. Run provenance login (or set PROVISIO_TOKEN in CI)."
	msgSessionExpire = "Your Provisio session has expired. Run provenance login."
	msgTokenRejected = "The personal token was rejected (revoked or expired). Replace PROVISIO_TOKEN, or run provenance login."
	msgScope         = "Your Provisio sign-in lacks permission to write; sign in again with `provenance login`."
)

// /sedum:setConst
// sedum:setFunction {"tier":"owned","record":"prov-2026-a0ad0a7f","gen":"sha256:dd29d8cc1590b0342455e97aa87320359bfbfee4affbf741dcc86c80a497061c","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func withTimeout(ctx context.Context, d time.Duration) (context.Context, func()) {\n\treturn context.WithTimeout(ctx, d)\n}","file":"cmd/provenance/session_run.go","name":"withTimeout"}}
func withTimeout(ctx context.Context, d time.Duration) (context.Context, func()) {
	return context.WithTimeout(ctx, d)
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-a51283a0","gen":"sha256:25743a0c40a8eff1905ae8c89926e35e8eb0ccddd6b35fc4be3bc508c591cfa4","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func serverURL() string {\n\tif v := os.Getenv(auth.EnvServer); v != \"\" {\n\t\treturn v\n\t}\n\treturn auth.DefaultServer\n}","file":"cmd/provenance/session_run.go","name":"serverURL"}}
func serverURL() string {
	if v := os.Getenv(auth.EnvServer); v != "" {
		return v
	}
	return auth.DefaultServer
}

// /sedum:setFunction
// sedum:setType {"tier":"owned","record":"prov-2026-ac80218b","gen":"sha256:0d39be9d85fb64f164375608f5889fed40cf83945504b6e57fa1f8a5af3ef08a","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type exitError struct {\n\tcode int\n\terr  error\n}","file":"cmd/provenance/session_run.go","name":"exitError"}}
type exitError struct {
	code int
	err  error
}

// /sedum:setType
// sedum:setFunction {"tier":"owned","record":"prov-2026-ae5b9ab4","gen":"sha256:906ad1c7ebf364264b0d30e9a0453b3e9ec1897578d2c9a4f4f1e2c07cc1787a","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func exitCode(err error) int {\n\tvar ee *exitError\n\tif errors.As(err, &ee) {\n\t\treturn ee.code\n\t}\n\treturn 1\n}","file":"cmd/provenance/session_run.go","name":"exitCode"}}
func exitCode(err error) int {
	var ee *exitError
	if errors.As(err, &ee) {
		return ee.code
	}
	return 1
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-da5d8cb8","gen":"sha256:1802c6e80df79a3253b8c10809eb08ec0bc9b3e9baaf85e48a0ef0902dc6ab1d","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func sessionManager() (*auth.Manager, error) {\n\tstore, err := auth.DefaultStore()\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\tpath, err := auth.DefaultFilePath()\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\tserver := serverURL()\n\tif os.Getenv(auth.EnvServer) == \"\" && os.Getenv(auth.EnvToken) == \"\" {\n\t\tif cred, _, err := store.Load(); err == nil && cred != nil && cred.Server != \"\" {\n\t\t\tserver = cred.Server\n\t\t}\n\t}\n\treturn &auth.Manager{\n\t\tStore:    store,\n\t\tClient:   &auth.Client{Server: server},\n\t\tLockPath: filepath.Join(filepath.Dir(path), \"credentials.lock\"),\n\t}, nil\n}","file":"cmd/provenance/session_run.go","name":"sessionManager"}}
func sessionManager() (*auth.Manager, error) {
	store, err := auth.DefaultStore()
	if err != nil {
		return nil, err
	}
	path, err := auth.DefaultFilePath()
	if err != nil {
		return nil, err
	}
	server := serverURL()
	if os.Getenv(auth.EnvServer) == "" && os.Getenv(auth.EnvToken) == "" {
		if cred, _, err := store.Load(); err == nil && cred != nil && cred.Server != "" {
			server = cred.Server
		}
	}
	return &auth.Manager{
		Store:    store,
		Client:   &auth.Client{Server: server},
		LockPath: filepath.Join(filepath.Dir(path), "credentials.lock"),
	}, nil
}

// /sedum:setFunction
