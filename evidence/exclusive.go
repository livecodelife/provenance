package evidence

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-b364fd75","gen":"sha256:04615758c2c5979274a7ade2b04e29a19ba7a1a8290dcf75eb9d80e4e7e77729","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"evidence/exclusive.go","path":"errors"}}
import "errors"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-b364fd75","gen":"sha256:04615758c2c5979274a7ade2b04e29a19ba7a1a8290dcf75eb9d80e4e7e77729","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"evidence/exclusive.go","path":"fmt"}}
import "fmt"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-b364fd75","gen":"sha256:04615758c2c5979274a7ade2b04e29a19ba7a1a8290dcf75eb9d80e4e7e77729","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"evidence/exclusive.go","path":"io/fs"}}
import "io/fs"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-b364fd75","gen":"sha256:04615758c2c5979274a7ade2b04e29a19ba7a1a8290dcf75eb9d80e4e7e77729","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"evidence/exclusive.go","path":"os"}}
import "os"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-b364fd75","gen":"sha256:04615758c2c5979274a7ade2b04e29a19ba7a1a8290dcf75eb9d80e4e7e77729","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"evidence/exclusive.go","path":"path/filepath"}}
import "path/filepath"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setConst {"tier":"owned","record":"prov-2026-b364fd75","gen":"sha256:04615758c2c5979274a7ade2b04e29a19ba7a1a8290dcf75eb9d80e4e7e77729","pkg":"sha256:51be1c18998d63cc558c253e211a6bd188d96cb251dad68c6ea88e03c2d5cc4c","kwargs":{"code":"// maxAppendRetries bounds how many successive names createExclusive tries\n// when the first one is taken, before it gives up.\nconst maxAppendRetries = 1000","file":"evidence/exclusive.go","name":"maxAppendRetries"}}
// maxAppendRetries bounds how many successive names createExclusive tries
// when the first one is taken, before it gives up.
const maxAppendRetries = 1000

// /sedum:setConst
// sedum:setFunction {"tier":"owned","record":"prov-2026-b364fd75","gen":"sha256:04615758c2c5979274a7ade2b04e29a19ba7a1a8290dcf75eb9d80e4e7e77729","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"// createExclusive creates <dir>/<version>-<nanos>.json for writing, never\n// overwriting: when the name exists it tries nanos+1, nanos+2, ... up to\n// maxAppendRetries names. Any other error is returned at once.\nfunc createExclusive(dir string, version int, nanos int64) (*os.File, error) {\n\tfor i := int64(0); i < maxAppendRetries; i++ {\n\t\tpath := filepath.Join(dir, fmt.Sprintf(\"%d-%d.json\", version, nanos+i))\n\t\tf, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o444)\n\t\tif err == nil {\n\t\t\treturn f, nil\n\t\t}\n\t\tif !errors.Is(err, fs.ErrExist) {\n\t\t\treturn nil, fmt.Errorf(\"evidence %s: %w\", path, err)\n\t\t}\n\t}\n\treturn nil, fmt.Errorf(\"evidence %s: no free file name after %d tries\", dir, maxAppendRetries)\n}","file":"evidence/exclusive.go","name":"createExclusive"}}
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
