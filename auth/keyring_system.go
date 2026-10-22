package auth

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-1396d0cc","gen":"sha256:059dfb1a1030a88cbdb7f96836f4fd05ee5190e3ebeb43c80ca494d65d1030d7","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"auth/keyring_system.go","path":"errors"}}
import "errors"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-1396d0cc","gen":"sha256:059dfb1a1030a88cbdb7f96836f4fd05ee5190e3ebeb43c80ca494d65d1030d7","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"auth/keyring_system.go","path":"github.com/zalando/go-keyring"}}
import "github.com/zalando/go-keyring"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setMethod {"tier":"owned","record":"prov-2026-1396d0cc","gen":"sha256:059dfb1a1030a88cbdb7f96836f4fd05ee5190e3ebeb43c80ca494d65d1030d7","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (systemKeyring) Get(service, user string) (string, error) {\n\tsecret, err := keyring.Get(service, user)\n\tif errors.Is(err, keyring.ErrNotFound) {\n\t\treturn \"\", ErrKeyNotFound\n\t}\n\treturn secret, err\n}","file":"auth/keyring_system.go","name":"Get","receiver":"systemKeyring"}}
func (systemKeyring) Get(service, user string) (string, error) {
	secret, err := keyring.Get(service, user)
	if errors.Is(err, keyring.ErrNotFound) {
		return "", ErrKeyNotFound
	}
	return secret, err
}

// /sedum:setMethod
// sedum:setMethod {"tier":"owned","record":"prov-2026-274813c6","gen":"sha256:2297630466d7322dabdc5319b12a45a5b3bb573c5d07d627fc26b29cfad9ccee","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (s systemKeyring) Set(service, user, value string) error {\n\treturn keyring.Set(service, user, value)\n}","file":"auth/keyring_system.go","name":"Set","receiver":"systemKeyring"}}
func (s systemKeyring) Set(service, user, value string) error {
	return keyring.Set(service, user, value)
}

// /sedum:setMethod
// sedum:setMethod {"tier":"owned","record":"prov-2026-677ba83d","gen":"sha256:c3a590dff65597edb09f53bda0a33406bd6a274da33a3d44cdf8f68b9eb42db0","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (systemKeyring) Delete(service, user string) error {\n\terr := keyring.Delete(service, user)\n\tif errors.Is(err, keyring.ErrNotFound) {\n\t\treturn nil\n\t}\n\treturn err\n}","file":"auth/keyring_system.go","name":"Delete","receiver":"systemKeyring"}}
func (systemKeyring) Delete(service, user string) error {
	err := keyring.Delete(service, user)
	if errors.Is(err, keyring.ErrNotFound) {
		return nil
	}
	return err
}

// /sedum:setMethod
// sedum:setType {"tier":"owned","record":"prov-2026-8f5a178b","gen":"sha256:100328ab485ae7a0090feac05b5b11c1004070e46d26ab8d99a37c3123337d90","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type systemKeyring struct{}","file":"auth/keyring_system.go","name":"systemKeyring"}}
type systemKeyring struct{}

// /sedum:setType
