package auth

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-ad56d5ca","gen":"sha256:7497b7652ae0ee9bba2b9df7ee5f5c8ac160c31fc9647412815a8ac0c5239587","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"auth/lock.go","path":"errors"}}
import "errors"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-ad56d5ca","gen":"sha256:7497b7652ae0ee9bba2b9df7ee5f5c8ac160c31fc9647412815a8ac0c5239587","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"auth/lock.go","path":"os"}}
import "os"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-19b64d12","gen":"sha256:486730d04a37015b203b5b6f1aa1463b716ba289cd2b683d4ba5799d8d33fc4d","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"auth/lock.go","path":"time"}}
import "time"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"tier":"owned","record":"prov-2026-ad56d5ca","gen":"sha256:7497b7652ae0ee9bba2b9df7ee5f5c8ac160c31fc9647412815a8ac0c5239587","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func lockFile(path string) (func(), error) {\n\tfor {\n\t\tf, err := os.OpenFile(path, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0600)\n\t\tif err == nil {\n\t\t\t_ = f.Close()\n\t\t\treturn func() { _ = os.Remove(path) }, nil\n\t\t}\n\t\tif !errors.Is(err, os.ErrExist) {\n\t\t\treturn nil, err\n\t\t}\n\t\tinfo, statErr := os.Stat(path)\n\t\tif statErr != nil {\n\t\t\ttime.Sleep(lockPoll)\n\t\t\tcontinue\n\t\t}\n\t\tif time.Since(info.ModTime()) > lockStale {\n\t\t\t_ = os.Remove(path)\n\t\t\tcontinue\n\t\t}\n\t\ttime.Sleep(lockPoll)\n\t}\n}","file":"auth/lock.go","name":"lockFile"}}
func lockFile(path string) (func(), error) {
	for {
		f, err := os.OpenFile(path, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0600)
		if err == nil {
			_ = f.Close()
			return func() { _ = os.Remove(path) }, nil
		}
		if !errors.Is(err, os.ErrExist) {
			return nil, err
		}
		info, statErr := os.Stat(path)
		if statErr != nil {
			time.Sleep(lockPoll)
			continue
		}
		if time.Since(info.ModTime()) > lockStale {
			_ = os.Remove(path)
			continue
		}
		time.Sleep(lockPoll)
	}
}

// /sedum:setFunction
// sedum:setConst {"tier":"owned","record":"prov-2026-19b64d12","gen":"sha256:486730d04a37015b203b5b6f1aa1463b716ba289cd2b683d4ba5799d8d33fc4d","pkg":"sha256:51be1c18998d63cc558c253e211a6bd188d96cb251dad68c6ea88e03c2d5cc4c","kwargs":{"code":"const (\n\tlockStale = time.Minute\n\tlockPoll  = 20 * time.Millisecond\n)","file":"auth/lock.go","name":"lockStale"}}
const (
	lockStale = time.Minute
	lockPoll  = 20 * time.Millisecond
)

// /sedum:setConst
