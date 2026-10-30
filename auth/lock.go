package auth

// sedum:anchor:imports
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-ad56d5ca","gen":"sha256:7497b7652ae0ee9bba2b9df7ee5f5c8ac160c31fc9647412815a8ac0c5239587","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"auth/lock.go","path":"errors"},"kwargs":"sha256:df6cb662090627add3284ebfbb5b6bcb3c56228fdc53540134ca8b5f77961c8f","body":"sha256:2b718237ba0a56467c8fd494cc131e942ef410e6228c03b644d70476c71862dc"}
import "errors"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-ad56d5ca","gen":"sha256:7497b7652ae0ee9bba2b9df7ee5f5c8ac160c31fc9647412815a8ac0c5239587","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"auth/lock.go","path":"os"},"kwargs":"sha256:64f21f8b9a796cfbefa23b2dd55f49af26c55044671b7586685b1684c1789c9f","body":"sha256:7c6e309eb2575d4aa451ec2ab83161d13e9ae03bbca3a4d2ecf30ce12adf0905"}
import "os"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-19b64d12","gen":"sha256:486730d04a37015b203b5b6f1aa1463b716ba289cd2b683d4ba5799d8d33fc4d","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"auth/lock.go","path":"time"},"kwargs":"sha256:f8ae776a3a0108e7849f854f06c7ec5b46a3d88ec9539d9895853744233c198f","body":"sha256:e99128ebf4e36998f40b33adf2eef57f4d6c315f1a4d1bdd3b7fc09a0d7bc82d"}
import "time"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-ad56d5ca","gen":"sha256:7497b7652ae0ee9bba2b9df7ee5f5c8ac160c31fc9647412815a8ac0c5239587","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"auth/lock.go","name":"lockFile"},"kwargs":"sha256:37bc4cd1dbcc49d646a155c228d5093e5809d7dd7416765916e6aee0c81307cc","body":"sha256:ee9b5e4c4da7756512429cfde545d5752680ea328c919b68010c6ce0b2aa8434"}
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
// sedum:setConst {"v":2,"tier":"owned","record":"prov-2026-19b64d12","gen":"sha256:486730d04a37015b203b5b6f1aa1463b716ba289cd2b683d4ba5799d8d33fc4d","pkg":"sha256:51be1c18998d63cc558c253e211a6bd188d96cb251dad68c6ea88e03c2d5cc4c","id":{"file":"auth/lock.go","name":"lockStale"},"kwargs":"sha256:102ed4999ca99d8a37cc03067d204ea4dd422fe65fe0c0123ec9bec9015d5755","body":"sha256:364c450f12bc40e2464414f54376ac1e8406b328bc2e4e4bbd2d71a126302a42"}
const (
	lockStale = time.Minute
	lockPoll  = 20 * time.Millisecond
)

// /sedum:setConst
