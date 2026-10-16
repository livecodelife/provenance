package auth

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-2b4db381","gen":"sha256:45a2a4d28f38438bd03beb44de6ad9b5ee07bdad521a440fcf25bbc3ce684236","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"auth/keyring_system.go","path":"errors"}}
import "errors"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-2b4db381","gen":"sha256:45a2a4d28f38438bd03beb44de6ad9b5ee07bdad521a440fcf25bbc3ce684236","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"auth/keyring_system.go","path":"github.com/zalando/go-keyring"}}
import "github.com/zalando/go-keyring"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setType {"tier":"owned","record":"prov-2026-2b4db381","gen":"sha256:45a2a4d28f38438bd03beb44de6ad9b5ee07bdad521a440fcf25bbc3ce684236","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"// systemKeyring is the OS keychain through go-keyring.\ntype systemKeyring struct{}","file":"auth/keyring_system.go","name":"systemKeyring"}}
// systemKeyring is the OS keychain through go-keyring.
type systemKeyring struct{}

// /sedum:setType
// sedum:setMethod {"tier":"owned","record":"prov-2026-2b4db381","gen":"sha256:45a2a4d28f38438bd03beb44de6ad9b5ee07bdad521a440fcf25bbc3ce684236","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (systemKeyring) Get(service, user string) (string, error) {\n\tv, err := keyring.Get(service, user)\n\tif errors.Is(err, keyring.ErrNotFound) {\n\t\treturn \"\", ErrKeyNotFound\n\t}\n\treturn v, err\n}","file":"auth/keyring_system.go","name":"Get","receiver":"systemKeyring"}}
func (systemKeyring) Get(service, user string) (string, error) {
	v, err := keyring.Get(service, user)
	if errors.Is(err, keyring.ErrNotFound) {
		return "", ErrKeyNotFound
	}
	return v, err
}

// /sedum:setMethod
// sedum:setMethod {"tier":"owned","record":"prov-2026-2b4db381","gen":"sha256:45a2a4d28f38438bd03beb44de6ad9b5ee07bdad521a440fcf25bbc3ce684236","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (systemKeyring) Set(service, user, value string) error { return keyring.Set(service, user, value) }","file":"auth/keyring_system.go","name":"Set","receiver":"systemKeyring"}}
func (systemKeyring) Set(service, user, value string) error { return keyring.Set(service, user, value) }

// /sedum:setMethod
// sedum:setMethod {"tier":"owned","record":"prov-2026-2b4db381","gen":"sha256:45a2a4d28f38438bd03beb44de6ad9b5ee07bdad521a440fcf25bbc3ce684236","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (systemKeyring) Delete(service, user string) error {\n\terr := keyring.Delete(service, user)\n\tif errors.Is(err, keyring.ErrNotFound) {\n\t\treturn nil\n\t}\n\treturn err\n}","file":"auth/keyring_system.go","name":"Delete","receiver":"systemKeyring"}}
func (systemKeyring) Delete(service, user string) error {
	err := keyring.Delete(service, user)
	if errors.Is(err, keyring.ErrNotFound) {
		return nil
	}
	return err
}

// /sedum:setMethod
