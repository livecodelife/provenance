package auth

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-3055e578","gen":"sha256:ac13be62e6d437ecbcf5c083c75b4fc254bf2a0b93097bb7f4893aa535c02e90","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"auth/store.go","path":"errors"}}
import "errors"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-3055e578","gen":"sha256:ac13be62e6d437ecbcf5c083c75b4fc254bf2a0b93097bb7f4893aa535c02e90","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"auth/store.go","path":"encoding/json"}}
import "encoding/json"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-d69bf77a","gen":"sha256:ee023230754d05a5997ae0cf33a00bbac05b8438ed302fb745cc14ad4f5fdd7f","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"auth/store.go","path":"fmt"}}
import "fmt"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-d69bf77a","gen":"sha256:ee023230754d05a5997ae0cf33a00bbac05b8438ed302fb745cc14ad4f5fdd7f","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"auth/store.go","path":"os"}}
import "os"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-d69bf77a","gen":"sha256:ee023230754d05a5997ae0cf33a00bbac05b8438ed302fb745cc14ad4f5fdd7f","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"auth/store.go","path":"path/filepath"}}
import "path/filepath"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setMethod {"tier":"owned","record":"prov-2026-3055e578","gen":"sha256:ac13be62e6d437ecbcf5c083c75b4fc254bf2a0b93097bb7f4893aa535c02e90","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (s FileStore) Load() (cred *Credential, src Source, err error) {\n\tinfo, err := os.Stat(s.Path)\n\tif err != nil {\n\t\tif errors.Is(err, os.ErrNotExist) {\n\t\t\treturn nil, src, ErrNoSession\n\t\t}\n\t\treturn nil, src, err\n\t}\n\tif perm := info.Mode().Perm(); perm&0o077 != 0 {\n\t\treturn nil, src, fmt.Errorf(\"%s is readable by other users (mode %04o); run: chmod 600 %s\", s.Path, perm, s.Path)\n\t}\n\tdata, err := os.ReadFile(s.Path)\n\tif err != nil {\n\t\treturn nil, src, err\n\t}\n\tvar c Credential\n\tif err := json.Unmarshal(data, &c); err != nil {\n\t\treturn nil, src, fmt.Errorf(\"%s is not a valid credential file; run provenance login again\", s.Path)\n\t}\n\treturn &c, SourceFile, nil\n}","file":"auth/store.go","name":"Load","receiver":"FileStore"}}
func (s FileStore) Load() (cred *Credential, src Source, err error) {
	info, err := os.Stat(s.Path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, src, ErrNoSession
		}
		return nil, src, err
	}
	if perm := info.Mode().Perm(); perm&0o077 != 0 {
		return nil, src, fmt.Errorf("%s is readable by other users (mode %04o); run: chmod 600 %s", s.Path, perm, s.Path)
	}
	data, err := os.ReadFile(s.Path)
	if err != nil {
		return nil, src, err
	}
	var c Credential
	if err := json.Unmarshal(data, &c); err != nil {
		return nil, src, fmt.Errorf("%s is not a valid credential file; run provenance login again", s.Path)
	}
	return &c, SourceFile, nil
}

// /sedum:setMethod
// sedum:setType {"tier":"owned","record":"prov-2026-309499de","gen":"sha256:4fb2138d83a91ffd2fd9ecc88d96535a74106f7bfa79a2875509809b994398c0","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type Keyring interface {\n\tGet(service, user string) (string, error)\n\tSet(service, user, value string) error\n\tDelete(service, user string) error\n}","file":"auth/store.go","name":"Keyring"}}
type Keyring interface {
	Get(service, user string) (string, error)
	Set(service, user, value string) error
	Delete(service, user string) error
}

// /sedum:setType
// sedum:setVar {"tier":"owned","record":"prov-2026-3419a57b","gen":"sha256:13903aa302916d0808a867732981c717aa1c89b9dc084a6b2c27d6b25d26e10a","pkg":"sha256:c7b9d9ac1655e8211b8b2ca8c49cf34b8d510d8d823d4f5eb60ec7e1018aa2eb","kwargs":{"code":"var ErrKeyNotFound = errors.New(\"keyring: not found\")","file":"auth/store.go","name":"ErrKeyNotFound"}}
var ErrKeyNotFound = errors.New("keyring: not found")

// /sedum:setVar
// sedum:setMethod {"tier":"owned","record":"prov-2026-3b6a607f","gen":"sha256:1f260ca3b5b86a51abce055a71ae8dee641829534e4f551a75c1a2f5a646fb55","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (s KeyringStore) Delete() error {\n\t_ = s.Keyring.Delete(keyringService, keyringUser)\n\treturn s.Fallback.Delete()\n}","file":"auth/store.go","name":"Delete","receiver":"KeyringStore"}}
func (s KeyringStore) Delete() error {
	_ = s.Keyring.Delete(keyringService, keyringUser)
	return s.Fallback.Delete()
}

// /sedum:setMethod
// sedum:setType {"tier":"owned","record":"prov-2026-47a53069","gen":"sha256:bf49404badae550c4d37cc91d01098aed7543ce05b441a4e9520a54804e48559","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type KeyringStore struct {\n\tKeyring  Keyring\n\tFallback Store\n}","file":"auth/store.go","name":"KeyringStore"}}
type KeyringStore struct {
	Keyring  Keyring
	Fallback Store
}

// /sedum:setType
// sedum:setMethod {"tier":"owned","record":"prov-2026-535555b5","gen":"sha256:9284b0e87c8dc6ceb232329e6fa2e11a0bce963b684d134569b6e61fe3cc7e53","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (s FileStore) Delete() error {\n\terr := os.Remove(s.Path)\n\tif errors.Is(err, os.ErrNotExist) {\n\t\treturn nil\n\t}\n\treturn err\n}","file":"auth/store.go","name":"Delete","receiver":"FileStore"}}
func (s FileStore) Delete() error {
	err := os.Remove(s.Path)
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	return err
}

// /sedum:setMethod
// sedum:setMethod {"tier":"owned","record":"prov-2026-83941993","gen":"sha256:a6cdf2f9d87c9b75dddfcdd7baabd8ce619ddf6e04741fbe58f3851681ad8c9b","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (k KeyringStore) Save(c *Credential) error {\n\tdata, err := json.Marshal(c)\n\tif err != nil {\n\t\treturn err\n\t}\n\tif err := k.Keyring.Set(keyringService, keyringUser, string(data)); err != nil {\n\t\treturn k.Fallback.Save(c)\n\t}\n\treturn k.Fallback.Delete()\n}","file":"auth/store.go","name":"Save","receiver":"KeyringStore"}}
func (k KeyringStore) Save(c *Credential) error {
	data, err := json.Marshal(c)
	if err != nil {
		return err
	}
	if err := k.Keyring.Set(keyringService, keyringUser, string(data)); err != nil {
		return k.Fallback.Save(c)
	}
	return k.Fallback.Delete()
}

// /sedum:setMethod
// sedum:setFunction {"tier":"owned","record":"prov-2026-8c290540","gen":"sha256:eefe6ed94ee80932048eba9289ba4b65b1418c97c0f4636bd8cd98b07525d7a0","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func DefaultStore() (Store, error) {\n\tpath, err := DefaultFilePath()\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\tfile := FileStore{Path: path}\n\tif os.Getenv(EnvNoKeychain) != \"\" {\n\t\treturn file, nil\n\t}\n\treturn KeyringStore{Keyring: systemKeyring{}, Fallback: file}, nil\n}","file":"auth/store.go","name":"DefaultStore"}}
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
// sedum:setConst {"tier":"owned","record":"prov-2026-ba2fe91f","gen":"sha256:4e4fa845b8c36b60bd817ceb779309ac9350bcba85f6126d26ce440e0621eda0","pkg":"sha256:51be1c18998d63cc558c253e211a6bd188d96cb251dad68c6ea88e03c2d5cc4c","kwargs":{"code":"const (\n\tkeyringService = \"provisio\"\n\tkeyringUser    = \"provenance-cli\"\n)","file":"auth/store.go","name":"keyringService"}}
const (
	keyringService = "provisio"
	keyringUser    = "provenance-cli"
)

// /sedum:setConst
// sedum:setType {"tier":"owned","record":"prov-2026-cac00d74","gen":"sha256:2a51bd29dbae4fb3316d65c8179a3967adb234c4813dca5304308ab7879ca791","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type FileStore struct {\n\tPath string\n}","file":"auth/store.go","name":"FileStore"}}
type FileStore struct {
	Path string
}

// /sedum:setType
// sedum:setMethod {"tier":"owned","record":"prov-2026-ce5c322b","gen":"sha256:0020a6666c64e5f4edd88a9cbf915c57105dac4e1c90fe97dc2cad2719d5b7a0","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (k KeyringStore) Load() (*Credential, Source, error) {\n\traw, err := k.Keyring.Get(keyringService, keyringUser)\n\tif err != nil {\n\t\treturn k.Fallback.Load()\n\t}\n\tvar c Credential\n\tif err := json.Unmarshal([]byte(raw), &c); err != nil {\n\t\treturn k.Fallback.Load()\n\t}\n\treturn &c, SourceKeychain, nil\n}","file":"auth/store.go","name":"Load","receiver":"KeyringStore"}}
func (k KeyringStore) Load() (*Credential, Source, error) {
	raw, err := k.Keyring.Get(keyringService, keyringUser)
	if err != nil {
		return k.Fallback.Load()
	}
	var c Credential
	if err := json.Unmarshal([]byte(raw), &c); err != nil {
		return k.Fallback.Load()
	}
	return &c, SourceKeychain, nil
}

// /sedum:setMethod
// sedum:setFunction {"tier":"owned","record":"prov-2026-d69bf77a","gen":"sha256:ee023230754d05a5997ae0cf33a00bbac05b8438ed302fb745cc14ad4f5fdd7f","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func DefaultFilePath() (string, error) {\n\thome, err := os.UserHomeDir()\n\tif err != nil {\n\t\treturn \"\", fmt.Errorf(\"cannot find the home directory: %w\", err)\n\t}\n\treturn filepath.Join(home, \".provisio\", \"credentials\"), nil\n}","file":"auth/store.go","name":"DefaultFilePath"}}
func DefaultFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("cannot find the home directory: %w", err)
	}
	return filepath.Join(home, ".provisio", "credentials"), nil
}

// /sedum:setFunction
// sedum:setType {"tier":"owned","record":"prov-2026-ee489fd6","gen":"sha256:6222ed3c8cf3985f49c881df8f85cf53458d279745f9d2a7885950294e6ef790","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type Store interface {\n\tLoad() (*Credential, Source, error)\n\tSave(c *Credential) error\n\tDelete() error\n}","file":"auth/store.go","name":"Store"}}
type Store interface {
	Load() (*Credential, Source, error)
	Save(c *Credential) error
	Delete() error
}

// /sedum:setType
// sedum:setMethod {"tier":"owned","record":"prov-2026-ef1c214e","gen":"sha256:0b1a771e73121250c9df3048cf896354b367f346693b389c16d73f3c09a57699","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","kwargs":{"code":"func (s FileStore) Save(c *Credential) error {\n\tdir := filepath.Dir(s.Path)\n\tif err := os.MkdirAll(dir, 0700); err != nil {\n\t\treturn err\n\t}\n\tdata, err := json.Marshal(c)\n\tif err != nil {\n\t\treturn err\n\t}\n\ttmp, err := os.CreateTemp(dir, \".credentials-*\")\n\tif err != nil {\n\t\treturn err\n\t}\n\tdefer os.Remove(tmp.Name())\n\tif err := tmp.Chmod(0600); err != nil {\n\t\ttmp.Close()\n\t\treturn err\n\t}\n\tif _, err := tmp.Write(data); err != nil {\n\t\ttmp.Close()\n\t\treturn err\n\t}\n\tif err := tmp.Close(); err != nil {\n\t\treturn err\n\t}\n\treturn os.Rename(tmp.Name(), s.Path)\n}","file":"auth/store.go","name":"Save","receiver":"FileStore"}}
func (s FileStore) Save(c *Credential) error {
	dir := filepath.Dir(s.Path)
	if err := os.MkdirAll(dir, 0700); err != nil {
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
	if err := tmp.Chmod(0600); err != nil {
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
