package auth

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-2b4db381","gen":"sha256:45a2a4d28f38438bd03beb44de6ad9b5ee07bdad521a440fcf25bbc3ce684236","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"auth/store.go","path":"encoding/json"}}
import "encoding/json"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-2b4db381","gen":"sha256:45a2a4d28f38438bd03beb44de6ad9b5ee07bdad521a440fcf25bbc3ce684236","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"auth/store.go","path":"errors"}}
import "errors"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-2b4db381","gen":"sha256:45a2a4d28f38438bd03beb44de6ad9b5ee07bdad521a440fcf25bbc3ce684236","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"auth/store.go","path":"fmt"}}
import "fmt"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-2b4db381","gen":"sha256:45a2a4d28f38438bd03beb44de6ad9b5ee07bdad521a440fcf25bbc3ce684236","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"auth/store.go","path":"os"}}
import "os"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-2b4db381","gen":"sha256:45a2a4d28f38438bd03beb44de6ad9b5ee07bdad521a440fcf25bbc3ce684236","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"auth/store.go","path":"path/filepath"}}
import "path/filepath"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setType {"tier":"owned","record":"prov-2026-2b4db381","gen":"sha256:45a2a4d28f38438bd03beb44de6ad9b5ee07bdad521a440fcf25bbc3ce684236","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"// Store keeps the credential between runs.\ntype Store interface {\n\tLoad() (*Credential, Source, error)\n\tSave(c *Credential) error\n\tDelete() error\n}","file":"auth/store.go","name":"Store"}}
// Store keeps the credential between runs.
type Store interface {
	Load() (*Credential, Source, error)
	Save(c *Credential) error
	Delete() error
}

// /sedum:setType
// sedum:setType {"tier":"owned","record":"prov-2026-2b4db381","gen":"sha256:45a2a4d28f38438bd03beb44de6ad9b5ee07bdad521a440fcf25bbc3ce684236","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"// FileStore keeps the credential in a 0600 file inside a 0700 directory.\ntype FileStore struct{ Path string }","file":"auth/store.go","name":"FileStore"}}
// FileStore keeps the credential in a 0600 file inside a 0700 directory.
type FileStore struct{ Path string }

// /sedum:setType
// sedum:setFunction {"tier":"owned","record":"prov-2026-2b4db381","gen":"sha256:45a2a4d28f38438bd03beb44de6ad9b5ee07bdad521a440fcf25bbc3ce684236","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"// DefaultFilePath is ~/.provisio/credentials.\nfunc DefaultFilePath() (string, error) {\n\thome, err := os.UserHomeDir()\n\tif err != nil {\n\t\treturn \"\", fmt.Errorf(\"cannot find the home directory: %w\", err)\n\t}\n\treturn filepath.Join(home, \".provisio\", \"credentials\"), nil\n}","file":"auth/store.go","name":"DefaultFilePath"}}
// DefaultFilePath is ~/.provisio/credentials.
func DefaultFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("cannot find the home directory: %w", err)
	}
	return filepath.Join(home, ".provisio", "credentials"), nil
}

// /sedum:setFunction
// sedum:setMethod {"tier":"owned","record":"prov-2026-2b4db381","gen":"sha256:45a2a4d28f38438bd03beb44de6ad9b5ee07bdad521a440fcf25bbc3ce684236","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (s FileStore) Load() (*Credential, Source, error) {\n\tinfo, err := os.Stat(s.Path)\n\tif errors.Is(err, os.ErrNotExist) {\n\t\treturn nil, \"\", ErrNoSession\n\t}\n\tif err != nil {\n\t\treturn nil, \"\", err\n\t}\n\tif info.Mode().Perm()&0o077 != 0 {\n\t\treturn nil, \"\", fmt.Errorf(\"%s is readable by other users (mode %04o); run: chmod 600 %s\", s.Path, info.Mode().Perm(), s.Path)\n\t}\n\tdata, err := os.ReadFile(s.Path)\n\tif err != nil {\n\t\treturn nil, \"\", err\n\t}\n\tvar c Credential\n\tif err := json.Unmarshal(data, &c); err != nil {\n\t\treturn nil, \"\", fmt.Errorf(\"%s is not a valid credential file; run provenance login again\", s.Path)\n\t}\n\treturn &c, SourceFile, nil\n}","file":"auth/store.go","name":"Load","receiver":"FileStore"}}
func (s FileStore) Load() (*Credential, Source, error) {
	info, err := os.Stat(s.Path)
	if errors.Is(err, os.ErrNotExist) {
		return nil, "", ErrNoSession
	}
	if err != nil {
		return nil, "", err
	}
	if info.Mode().Perm()&0o077 != 0 {
		return nil, "", fmt.Errorf("%s is readable by other users (mode %04o); run: chmod 600 %s", s.Path, info.Mode().Perm(), s.Path)
	}
	data, err := os.ReadFile(s.Path)
	if err != nil {
		return nil, "", err
	}
	var c Credential
	if err := json.Unmarshal(data, &c); err != nil {
		return nil, "", fmt.Errorf("%s is not a valid credential file; run provenance login again", s.Path)
	}
	return &c, SourceFile, nil
}

// /sedum:setMethod
// sedum:setMethod {"tier":"owned","record":"prov-2026-2b4db381","gen":"sha256:45a2a4d28f38438bd03beb44de6ad9b5ee07bdad521a440fcf25bbc3ce684236","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"// Save writes atomically: a temp file in the same directory, then a rename.\nfunc (s FileStore) Save(c *Credential) error {\n\tdir := filepath.Dir(s.Path)\n\tif err := os.MkdirAll(dir, 0o700); err != nil {\n\t\treturn err\n\t}\n\tdata, err := json.Marshal(c)\n\tif err != nil {\n\t\treturn err\n\t}\n\ttmp, err := os.CreateTemp(dir, \".credentials-*\")\n\tif err != nil {\n\t\treturn err\n\t}\n\tdefer os.Remove(tmp.Name())\n\tif err := tmp.Chmod(0o600); err != nil {\n\t\ttmp.Close()\n\t\treturn err\n\t}\n\tif _, err := tmp.Write(data); err != nil {\n\t\ttmp.Close()\n\t\treturn err\n\t}\n\tif err := tmp.Close(); err != nil {\n\t\treturn err\n\t}\n\treturn os.Rename(tmp.Name(), s.Path)\n}","file":"auth/store.go","name":"Save","receiver":"FileStore"}}
// Save writes atomically: a temp file in the same directory, then a rename.
func (s FileStore) Save(c *Credential) error {
	dir := filepath.Dir(s.Path)
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return err
	}
	data, err := json.Marshal(c)
	if err != nil {
		return err
	}
	tmp, err := os.CreateTemp(dir, ".credentials-*")
	if err != nil {
		return err
	}
	defer os.Remove(tmp.Name())
	if err := tmp.Chmod(0o600); err != nil {
		tmp.Close()
		return err
	}
	if _, err := tmp.Write(data); err != nil {
		tmp.Close()
		return err
	}
	if err := tmp.Close(); err != nil {
		return err
	}
	return os.Rename(tmp.Name(), s.Path)
}

// /sedum:setMethod
// sedum:setMethod {"tier":"owned","record":"prov-2026-2b4db381","gen":"sha256:45a2a4d28f38438bd03beb44de6ad9b5ee07bdad521a440fcf25bbc3ce684236","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (s FileStore) Delete() error {\n\tif err := os.Remove(s.Path); err != nil && !errors.Is(err, os.ErrNotExist) {\n\t\treturn err\n\t}\n\treturn nil\n}","file":"auth/store.go","name":"Delete","receiver":"FileStore"}}
func (s FileStore) Delete() error {
	if err := os.Remove(s.Path); err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}
	return nil
}

// /sedum:setMethod
// sedum:setVar {"tier":"owned","record":"prov-2026-2b4db381","gen":"sha256:45a2a4d28f38438bd03beb44de6ad9b5ee07bdad521a440fcf25bbc3ce684236","pkg":"sha256:c7b9d9ac1655e8211b8b2ca8c49cf34b8d510d8d823d4f5eb60ec7e1018aa2eb","kwargs":{"code":"// ErrKeyNotFound is what a Keyring returns for a missing entry.\nvar ErrKeyNotFound = errors.New(\"keyring: not found\")","file":"auth/store.go","name":"ErrKeyNotFound"}}
// ErrKeyNotFound is what a Keyring returns for a missing entry.
var ErrKeyNotFound = errors.New("keyring: not found")

// /sedum:setVar
// sedum:setType {"tier":"owned","record":"prov-2026-2b4db381","gen":"sha256:45a2a4d28f38438bd03beb44de6ad9b5ee07bdad521a440fcf25bbc3ce684236","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"// Keyring is the OS keychain, behind an interface so it is optional at runtime.\ntype Keyring interface {\n\tGet(service, user string) (string, error)\n\tSet(service, user, value string) error\n\tDelete(service, user string) error\n}","file":"auth/store.go","name":"Keyring"}}
// Keyring is the OS keychain, behind an interface so it is optional at runtime.
type Keyring interface {
	Get(service, user string) (string, error)
	Set(service, user, value string) error
	Delete(service, user string) error
}

// /sedum:setType
// sedum:setConst {"tier":"owned","record":"prov-2026-2b4db381","gen":"sha256:45a2a4d28f38438bd03beb44de6ad9b5ee07bdad521a440fcf25bbc3ce684236","pkg":"sha256:51be1c18998d63cc558c253e211a6bd188d96cb251dad68c6ea88e03c2d5cc4c","kwargs":{"code":"const (\n\tkeyringService = \"provisio\"\n\tkeyringUser    = \"provenance-cli\"\n)","file":"auth/store.go","name":"keyringService"}}
const (
	keyringService = "provisio"
	keyringUser    = "provenance-cli"
)

// /sedum:setConst
// sedum:setType {"tier":"owned","record":"prov-2026-2b4db381","gen":"sha256:45a2a4d28f38438bd03beb44de6ad9b5ee07bdad521a440fcf25bbc3ce684236","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"// KeyringStore prefers the keychain and falls back to Fallback whenever the\n// keychain is missing or fails (a headless machine has none).\ntype KeyringStore struct {\n\tKeyring  Keyring\n\tFallback Store\n}","file":"auth/store.go","name":"KeyringStore"}}
// KeyringStore prefers the keychain and falls back to Fallback whenever the
// keychain is missing or fails (a headless machine has none).
type KeyringStore struct {
	Keyring  Keyring
	Fallback Store
}

// /sedum:setType
// sedum:setMethod {"tier":"owned","record":"prov-2026-2b4db381","gen":"sha256:45a2a4d28f38438bd03beb44de6ad9b5ee07bdad521a440fcf25bbc3ce684236","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (s KeyringStore) Load() (*Credential, Source, error) {\n\tif raw, err := s.Keyring.Get(keyringService, keyringUser); err == nil {\n\t\tvar c Credential\n\t\tif json.Unmarshal([]byte(raw), &c) == nil {\n\t\t\treturn &c, SourceKeychain, nil\n\t\t}\n\t}\n\treturn s.Fallback.Load()\n}","file":"auth/store.go","name":"Load","receiver":"KeyringStore"}}
func (s KeyringStore) Load() (*Credential, Source, error) {
	if raw, err := s.Keyring.Get(keyringService, keyringUser); err == nil {
		var c Credential
		if json.Unmarshal([]byte(raw), &c) == nil {
			return &c, SourceKeychain, nil
		}
	}
	return s.Fallback.Load()
}

// /sedum:setMethod
// sedum:setMethod {"tier":"owned","record":"prov-2026-2b4db381","gen":"sha256:45a2a4d28f38438bd03beb44de6ad9b5ee07bdad521a440fcf25bbc3ce684236","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (s KeyringStore) Save(c *Credential) error {\n\tdata, err := json.Marshal(c)\n\tif err != nil {\n\t\treturn err\n\t}\n\tif err := s.Keyring.Set(keyringService, keyringUser, string(data)); err != nil {\n\t\treturn s.Fallback.Save(c)\n\t}\n\treturn s.Fallback.Delete()\n}","file":"auth/store.go","name":"Save","receiver":"KeyringStore"}}
func (s KeyringStore) Save(c *Credential) error {
	data, err := json.Marshal(c)
	if err != nil {
		return err
	}
	if err := s.Keyring.Set(keyringService, keyringUser, string(data)); err != nil {
		return s.Fallback.Save(c)
	}
	return s.Fallback.Delete()
}

// /sedum:setMethod
// sedum:setMethod {"tier":"owned","record":"prov-2026-2b4db381","gen":"sha256:45a2a4d28f38438bd03beb44de6ad9b5ee07bdad521a440fcf25bbc3ce684236","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (s KeyringStore) Delete() error {\n\t_ = s.Keyring.Delete(keyringService, keyringUser)\n\treturn s.Fallback.Delete()\n}","file":"auth/store.go","name":"Delete","receiver":"KeyringStore"}}
func (s KeyringStore) Delete() error {
	_ = s.Keyring.Delete(keyringService, keyringUser)
	return s.Fallback.Delete()
}

// /sedum:setMethod
// sedum:setFunction {"tier":"owned","record":"prov-2026-2b4db381","gen":"sha256:45a2a4d28f38438bd03beb44de6ad9b5ee07bdad521a440fcf25bbc3ce684236","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"// DefaultStore is the keychain with the file as fallback, or just the file\n// when PROVISIO_NO_KEYCHAIN is set.\nfunc DefaultStore() (Store, error) {\n\tpath, err := DefaultFilePath()\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\tfile := FileStore{Path: path}\n\tif os.Getenv(EnvNoKeychain) != \"\" {\n\t\treturn file, nil\n\t}\n\treturn KeyringStore{Keyring: systemKeyring{}, Fallback: file}, nil\n}","file":"auth/store.go","name":"DefaultStore"}}
// DefaultStore is the keychain with the file as fallback, or just the file
// when PROVISIO_NO_KEYCHAIN is set.
func DefaultStore() (Store, error) {
	path, err := DefaultFilePath()
	if err != nil {
		return nil, err
	}
	file := FileStore{Path: path}
	if os.Getenv(EnvNoKeychain) != "" {
		return file, nil
	}
	return KeyringStore{Keyring: systemKeyring{}, Fallback: file}, nil
}

// /sedum:setFunction
