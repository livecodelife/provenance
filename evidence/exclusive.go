package evidence

// sedum:anchor:imports
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-b364fd75","gen":"sha256:04615758c2c5979274a7ade2b04e29a19ba7a1a8290dcf75eb9d80e4e7e77729","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"evidence/exclusive.go","path":"errors"},"kwargs":"sha256:1bea5ce872a254b1a81c584015417ebd244da0007d7309cf729373da43995709","body":"sha256:2b718237ba0a56467c8fd494cc131e942ef410e6228c03b644d70476c71862dc"}
import "errors"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-b364fd75","gen":"sha256:04615758c2c5979274a7ade2b04e29a19ba7a1a8290dcf75eb9d80e4e7e77729","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"evidence/exclusive.go","path":"fmt"},"kwargs":"sha256:4ec4c7fcf4582c69f53666e9d3b32de4618dcbc8fd0f6b4fc388a8a34579a7f0","body":"sha256:bcbac51883cbf8064e9196b371cb880cee429866652447ad557533d817857609"}
import "fmt"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-b364fd75","gen":"sha256:04615758c2c5979274a7ade2b04e29a19ba7a1a8290dcf75eb9d80e4e7e77729","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"evidence/exclusive.go","path":"io/fs"},"kwargs":"sha256:da70085ce9b723aab28221d7da7c9022c1d8cb0d1feca6db83a977ac4decae9b","body":"sha256:80f0597eb523a3ace402ac68fefdf64bf00074dda5ec249d04e8f1e525952547"}
import "io/fs"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-b364fd75","gen":"sha256:04615758c2c5979274a7ade2b04e29a19ba7a1a8290dcf75eb9d80e4e7e77729","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"evidence/exclusive.go","path":"os"},"kwargs":"sha256:cdb49f6e4b06217f96fda725237e21cf7ed019a1a77babf938b8a81b3712664b","body":"sha256:7c6e309eb2575d4aa451ec2ab83161d13e9ae03bbca3a4d2ecf30ce12adf0905"}
import "os"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-b364fd75","gen":"sha256:04615758c2c5979274a7ade2b04e29a19ba7a1a8290dcf75eb9d80e4e7e77729","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"evidence/exclusive.go","path":"path/filepath"},"kwargs":"sha256:e1f97fc2309cdc196a3855233ad8b7458fc20b9dc3abd2e6b95a26ea5bb43742","body":"sha256:a6d196b558491e646be93f56a321f339068c64c76aa9d49bb18046e7f9c777a1"}
import "path/filepath"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setConst {"v":2,"tier":"owned","record":"prov-2026-b364fd75","gen":"sha256:04615758c2c5979274a7ade2b04e29a19ba7a1a8290dcf75eb9d80e4e7e77729","pkg":"sha256:51be1c18998d63cc558c253e211a6bd188d96cb251dad68c6ea88e03c2d5cc4c","id":{"file":"evidence/exclusive.go","name":"maxAppendRetries"},"kwargs":"sha256:de506654caaaed64c21b6b56046b9e95f9db43010b4049ecd157e624040d4e30","body":"sha256:640f49f197015134c524cdedc6c831d8d1ee6bf076d8fc05617b5e20870267fd"}
// maxAppendRetries bounds how many successive names createExclusive tries
// when the first one is taken, before it gives up.
const maxAppendRetries = 1000

// /sedum:setConst
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-b364fd75","gen":"sha256:04615758c2c5979274a7ade2b04e29a19ba7a1a8290dcf75eb9d80e4e7e77729","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"evidence/exclusive.go","name":"createExclusive"},"kwargs":"sha256:7fe28617fff2f26f55628ec014b62089be3300df44cc689f1221eb1d4e6e8f6e","body":"sha256:8d603935cfef9e3d8c8742948dea5c2627a6358f9aadda4e32e1e23fb0e080fb"}
// createExclusive creates <dir>/<version>-<nanos>.json for writing, never
// overwriting: when the name exists it tries nanos+1, nanos+2, ... up to
// maxAppendRetries names. Any other error is returned at once.
func createExclusive(dir string, version int, nanos int64) (*os.File, error) {
	for i := int64(0); i < maxAppendRetries; i++ {
		path := filepath.Join(dir, fmt.Sprintf("%d-%d.json", version, nanos+i))
		f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o444)
		if err == nil {
			return f, nil
		}
		if !errors.Is(err, fs.ErrExist) {
			return nil, fmt.Errorf("evidence %s: %w", path, err)
		}
	}
	return nil, fmt.Errorf("evidence %s: no free file name after %d tries", dir, maxAppendRetries)
}

// /sedum:setFunction
