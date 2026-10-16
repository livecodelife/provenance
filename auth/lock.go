package auth

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-ad56d5ca","gen":"sha256:b64ac5dd3402fa436945f27da8b32d59b0a8fc253e6d8ddb0e27f4d9b585565b","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"auth/lock.go","path":"errors"}}
import "errors"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-ad56d5ca","gen":"sha256:b64ac5dd3402fa436945f27da8b32d59b0a8fc253e6d8ddb0e27f4d9b585565b","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"auth/lock.go","path":"os"}}
import "os"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-ad56d5ca","gen":"sha256:b64ac5dd3402fa436945f27da8b32d59b0a8fc253e6d8ddb0e27f4d9b585565b","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"auth/lock.go","path":"time"}}
import "time"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setConst {"tier":"owned","record":"prov-2026-ad56d5ca","gen":"sha256:b64ac5dd3402fa436945f27da8b32d59b0a8fc253e6d8ddb0e27f4d9b585565b","pkg":"sha256:51be1c18998d63cc558c253e211a6bd188d96cb251dad68c6ea88e03c2d5cc4c","kwargs":{"code":"const (\n\tlockStale = time.Minute\n\tlockPoll  = 20 * time.Millisecond\n)","file":"auth/lock.go","name":"lockStale"}}
const (
	lockStale = time.Minute
	lockPoll  = 20 * time.Millisecond
)

// /sedum:setConst
// sedum:setFunction {"tier":"owned","record":"prov-2026-ad56d5ca","gen":"sha256:b64ac5dd3402fa436945f27da8b32d59b0a8fc253e6d8ddb0e27f4d9b585565b","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"// lockFile takes an exclusive lock by creating path with O_EXCL, polling until\n// it is free. A lock file older than lockStale was left by a crashed process\n// and is taken over. It returns the unlock func.\nfunc lockFile(path string) (func(), error) {\n\tfor {\n\t\tf, err := os.OpenFile(path, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0o600)\n\t\tif err == nil {\n\t\t\tf.Close()\n\t\t\treturn func() { os.Remove(path) }, nil\n\t\t}\n\t\tif !errors.Is(err, os.ErrExist) {\n\t\t\treturn nil, err\n\t\t}\n\t\tif info, err := os.Stat(path); err == nil && time.Since(info.ModTime()) > lockStale {\n\t\t\tos.Remove(path)\n\t\t\tcontinue\n\t\t}\n\t\ttime.Sleep(lockPoll)\n\t}\n}","file":"auth/lock.go","name":"lockFile"}}
// lockFile takes an exclusive lock by creating path with O_EXCL, polling until
// it is free. A lock file older than lockStale was left by a crashed process
// and is taken over. It returns the unlock func.
func lockFile(path string) (func(), error) {
	for {
		f, err := os.OpenFile(path, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0o600)
		if err == nil {
			f.Close()
			return func() { os.Remove(path) }, nil
		}
		if !errors.Is(err, os.ErrExist) {
			return nil, err
		}
		if info, err := os.Stat(path); err == nil && time.Since(info.ModTime()) > lockStale {
			os.Remove(path)
			continue
		}
		time.Sleep(lockPoll)
	}
}

// /sedum:setFunction
