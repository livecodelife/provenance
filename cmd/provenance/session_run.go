package main

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-81333e45","gen":"sha256:f2644a283573d45641e67e9a276d6c4d80b979c23d681b83ce11bd5078e4a850","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/session_run.go","path":"context"}}
import "context"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-81333e45","gen":"sha256:f2644a283573d45641e67e9a276d6c4d80b979c23d681b83ce11bd5078e4a850","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/session_run.go","path":"errors"}}
import "errors"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-81333e45","gen":"sha256:f2644a283573d45641e67e9a276d6c4d80b979c23d681b83ce11bd5078e4a850","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/session_run.go","path":"os"}}
import "os"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-81333e45","gen":"sha256:f2644a283573d45641e67e9a276d6c4d80b979c23d681b83ce11bd5078e4a850","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/session_run.go","path":"path/filepath"}}
import "path/filepath"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-81333e45","gen":"sha256:f2644a283573d45641e67e9a276d6c4d80b979c23d681b83ce11bd5078e4a850","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/session_run.go","path":"time"}}
import "time"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-81333e45","gen":"sha256:f2644a283573d45641e67e9a276d6c4d80b979c23d681b83ce11bd5078e4a850","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/session_run.go","path":"github.com/livecodelife/provenance/auth"}}
import "github.com/livecodelife/provenance/auth"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setType {"tier":"owned","record":"prov-2026-81333e45","gen":"sha256:f2644a283573d45641e67e9a276d6c4d80b979c23d681b83ce11bd5078e4a850","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"// exitError carries a process exit status for main to use.\ntype exitError struct {\n\tcode int\n\terr  error\n}","file":"cmd/provenance/session_run.go","name":"exitError"}}
// exitError carries a process exit status for main to use.
type exitError struct {
	code int
	err  error
}

// /sedum:setType
// sedum:setMethod {"tier":"owned","record":"prov-2026-81333e45","gen":"sha256:f2644a283573d45641e67e9a276d6c4d80b979c23d681b83ce11bd5078e4a850","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (e *exitError) Error() string { return e.err.Error() }","file":"cmd/provenance/session_run.go","name":"Error","receiver":"exitError"}}
func (e *exitError) Error() string { return e.err.Error() }

// /sedum:setMethod
// sedum:setMethod {"tier":"owned","record":"prov-2026-81333e45","gen":"sha256:f2644a283573d45641e67e9a276d6c4d80b979c23d681b83ce11bd5078e4a850","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (e *exitError) Unwrap() error { return e.err }","file":"cmd/provenance/session_run.go","name":"Unwrap","receiver":"exitError"}}
func (e *exitError) Unwrap() error { return e.err }

// /sedum:setMethod
// sedum:setFunction {"tier":"owned","record":"prov-2026-81333e45","gen":"sha256:f2644a283573d45641e67e9a276d6c4d80b979c23d681b83ce11bd5078e4a850","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"// exitCode is the process exit status for err: 3 when a Provisio session is\n// needed or has expired, 1 for anything else.\nfunc exitCode(err error) int {\n\tvar ee *exitError\n\tif errors.As(err, &ee) {\n\t\treturn ee.code\n\t}\n\treturn 1\n}","file":"cmd/provenance/session_run.go","name":"exitCode"}}
// exitCode is the process exit status for err: 3 when a Provisio session is
// needed or has expired, 1 for anything else.
func exitCode(err error) int {
	var ee *exitError
	if errors.As(err, &ee) {
		return ee.code
	}
	return 1
}

// /sedum:setFunction
// sedum:setConst {"tier":"owned","record":"prov-2026-81333e45","gen":"sha256:f2644a283573d45641e67e9a276d6c4d80b979c23d681b83ce11bd5078e4a850","pkg":"sha256:51be1c18998d63cc558c253e211a6bd188d96cb251dad68c6ea88e03c2d5cc4c","kwargs":{"code":"const (\n\tmsgNeedSession   = \"This command needs a Provisio session. Run provenance login (or set PROVISIO_TOKEN in CI).\"\n\tmsgSessionExpire = \"Your Provisio session has expired. Run provenance login.\"\n\tmsgTokenRejected = \"The personal token was rejected (revoked or expired). Replace PROVISIO_TOKEN, or run provenance login.\"\n)","file":"cmd/provenance/session_run.go","name":"msgNeedSession"}}
const (
	msgNeedSession   = "This command needs a Provisio session. Run provenance login (or set PROVISIO_TOKEN in CI)."
	msgSessionExpire = "Your Provisio session has expired. Run provenance login."
	msgTokenRejected = "The personal token was rejected (revoked or expired). Replace PROVISIO_TOKEN, or run provenance login."
)

// /sedum:setConst
// sedum:setFunction {"tier":"owned","record":"prov-2026-81333e45","gen":"sha256:f2644a283573d45641e67e9a276d6c4d80b979c23d681b83ce11bd5078e4a850","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"// sessionError turns an auth failure into the user-facing message with exit\n// status 3; other errors pass through unchanged.\nfunc sessionError(err error) error {\n\tswitch {\n\tcase err == nil:\n\t\treturn nil\n\tcase errors.Is(err, auth.ErrNoSession):\n\t\treturn &exitError{3, errors.New(msgNeedSession)}\n\tcase errors.Is(err, auth.ErrSessionExpired):\n\t\treturn &exitError{3, errors.New(msgSessionExpire)}\n\tcase errors.Is(err, auth.ErrTokenRejected):\n\t\treturn &exitError{3, errors.New(msgTokenRejected)}\n\t}\n\treturn err\n}","file":"cmd/provenance/session_run.go","name":"sessionError"}}
// sessionError turns an auth failure into the user-facing message with exit
// status 3; other errors pass through unchanged.
func sessionError(err error) error {
	switch {
	case err == nil:
		return nil
	case errors.Is(err, auth.ErrNoSession):
		return &exitError{3, errors.New(msgNeedSession)}
	case errors.Is(err, auth.ErrSessionExpired):
		return &exitError{3, errors.New(msgSessionExpire)}
	case errors.Is(err, auth.ErrTokenRejected):
		return &exitError{3, errors.New(msgTokenRejected)}
	}
	return err
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-81333e45","gen":"sha256:f2644a283573d45641e67e9a276d6c4d80b979c23d681b83ce11bd5078e4a850","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"// serverURL is the auth service for a new sign-in: PROVISIO_AUTH_URL, else the\n// default. Commands that use a stored sign-in prefer the server it was made on.\nfunc serverURL() string {\n\tif s := os.Getenv(auth.EnvServer); s != \"\" {\n\t\treturn s\n\t}\n\treturn auth.DefaultServer\n}","file":"cmd/provenance/session_run.go","name":"serverURL"}}
// serverURL is the auth service for a new sign-in: PROVISIO_AUTH_URL, else the
// default. Commands that use a stored sign-in prefer the server it was made on.
func serverURL() string {
	if s := os.Getenv(auth.EnvServer); s != "" {
		return s
	}
	return auth.DefaultServer
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-81333e45","gen":"sha256:f2644a283573d45641e67e9a276d6c4d80b979c23d681b83ce11bd5078e4a850","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"// sessionManager builds the token manager over the default credential store.\nfunc sessionManager() (*auth.Manager, error) {\n\tstore, err := auth.DefaultStore()\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\tpath, err := auth.DefaultFilePath()\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\tserver := serverURL()\n\tif os.Getenv(auth.EnvServer) == \"\" && os.Getenv(auth.EnvToken) == \"\" {\n\t\tif cred, _, err := store.Load(); err == nil && cred.Server != \"\" {\n\t\t\tserver = cred.Server\n\t\t}\n\t}\n\treturn &auth.Manager{\n\t\tStore:    store,\n\t\tClient:   &auth.Client{Server: server},\n\t\tLockPath: filepath.Join(filepath.Dir(path), \"credentials.lock\"),\n\t}, nil\n}","file":"cmd/provenance/session_run.go","name":"sessionManager"}}
// sessionManager builds the token manager over the default credential store.
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
		if cred, _, err := store.Load(); err == nil && cred.Server != "" {
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
// sedum:setFunction {"tier":"owned","record":"prov-2026-81333e45","gen":"sha256:f2644a283573d45641e67e9a276d6c4d80b979c23d681b83ce11bd5078e4a850","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"// withTimeout is context.WithTimeout with a cancel func that needs no import\n// at the call site.\nfunc withTimeout(ctx context.Context, d time.Duration) (context.Context, func()) {\n\treturn context.WithTimeout(ctx, d)\n}","file":"cmd/provenance/session_run.go","name":"withTimeout"}}
// withTimeout is context.WithTimeout with a cancel func that needs no import
// at the call site.
func withTimeout(ctx context.Context, d time.Duration) (context.Context, func()) {
	return context.WithTimeout(ctx, d)
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-81333e45","gen":"sha256:f2644a283573d45641e67e9a276d6c4d80b979c23d681b83ce11bd5078e4a850","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"// withSession runs fn with a bearer token, refreshing once on a 401. Commands\n// that need a Provisio session (publish, clone, import) go through it; local\n// commands never call it.\nfunc withSession(ctx context.Context, fn func(bearer string) error) error {\n\tm, err := sessionManager()\n\tif err != nil {\n\t\treturn err\n\t}\n\treturn sessionError(m.Do(ctx, fn))\n}","file":"cmd/provenance/session_run.go","name":"withSession"}}
// withSession runs fn with a bearer token, refreshing once on a 401. Commands
// that need a Provisio session (publish, clone, import) go through it; local
// commands never call it.
func withSession(ctx context.Context, fn func(bearer string) error) error {
	m, err := sessionManager()
	if err != nil {
		return err
	}
	return sessionError(m.Do(ctx, fn))
}

// /sedum:setFunction
