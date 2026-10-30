package auth

// sedum:anchor:imports
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-3055e578","gen":"sha256:ac13be62e6d437ecbcf5c083c75b4fc254bf2a0b93097bb7f4893aa535c02e90","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"auth/store.go","path":"errors"},"kwargs":"sha256:31458d5c5473ca49f4042411e9b730d3d79ade576ecf00571a96d66b80f58686","body":"sha256:2b718237ba0a56467c8fd494cc131e942ef410e6228c03b644d70476c71862dc"}
import "errors"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-3055e578","gen":"sha256:ac13be62e6d437ecbcf5c083c75b4fc254bf2a0b93097bb7f4893aa535c02e90","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"auth/store.go","path":"encoding/json"},"kwargs":"sha256:cd4f5e55f61b071ac002d0dc1aa30f86a05436a01abcbab8038612891699743c","body":"sha256:872d7c3e9dee608e79a7f8ef57930d24b5e30f2a57d891a52549179ace45946a"}
import "encoding/json"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-d69bf77a","gen":"sha256:ee023230754d05a5997ae0cf33a00bbac05b8438ed302fb745cc14ad4f5fdd7f","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"auth/store.go","path":"fmt"},"kwargs":"sha256:21d4df7168961ce10d168efea28949dcd82d5293857ab3b2c9a2c6c3d5570bf0","body":"sha256:bcbac51883cbf8064e9196b371cb880cee429866652447ad557533d817857609"}
import "fmt"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-d69bf77a","gen":"sha256:ee023230754d05a5997ae0cf33a00bbac05b8438ed302fb745cc14ad4f5fdd7f","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"auth/store.go","path":"os"},"kwargs":"sha256:66b546b68aab8762f0de5893019f1359702872c60dbd047bea27a514d8dabdde","body":"sha256:7c6e309eb2575d4aa451ec2ab83161d13e9ae03bbca3a4d2ecf30ce12adf0905"}
import "os"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-d69bf77a","gen":"sha256:ee023230754d05a5997ae0cf33a00bbac05b8438ed302fb745cc14ad4f5fdd7f","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"auth/store.go","path":"path/filepath"},"kwargs":"sha256:83803e7fef394fc9f965241ec33350b5dddbde54807d891e7adc583f6de398e4","body":"sha256:a6d196b558491e646be93f56a321f339068c64c76aa9d49bb18046e7f9c777a1"}
import "path/filepath"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setMethod {"v":2,"tier":"owned","record":"prov-2026-3055e578","gen":"sha256:ac13be62e6d437ecbcf5c083c75b4fc254bf2a0b93097bb7f4893aa535c02e90","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","id":{"file":"auth/store.go","name":"Load","receiver":"FileStore"},"kwargs":"sha256:e467312f8365465e66ab4e899abd64a072233441206c711b64fb69b45c712f71","body":"sha256:fe41a4f0c4bd74f4a0a74aca5388a29c8e66dc66ed8d888dfbb59c145db4b1da"}
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
// sedum:setType {"v":2,"tier":"owned","record":"prov-2026-309499de","gen":"sha256:4fb2138d83a91ffd2fd9ecc88d96535a74106f7bfa79a2875509809b994398c0","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","id":{"file":"auth/store.go","name":"Keyring"},"kwargs":"sha256:0efd7227c3f84932f0eb577f5b7e4a8eb3e7f048d1f5046962dfa9e4771d5e96","body":"sha256:ce4fa00e6778c3e53472d4a624dda0534aed1c4ed3132662e544af21e6d1a259"}
type Keyring interface {
	Get(service, user string) (string, error)
	Set(service, user, value string) error
	Delete(service, user string) error
}

// /sedum:setType
// sedum:setVar {"v":2,"tier":"owned","record":"prov-2026-3419a57b","gen":"sha256:13903aa302916d0808a867732981c717aa1c89b9dc084a6b2c27d6b25d26e10a","pkg":"sha256:c7b9d9ac1655e8211b8b2ca8c49cf34b8d510d8d823d4f5eb60ec7e1018aa2eb","id":{"file":"auth/store.go","name":"ErrKeyNotFound"},"kwargs":"sha256:533efcc6be9b649a2b0a4d380544fff588eaba977315b4111f529b94670eced3","body":"sha256:382468217d98fd7acf66dfc0d2106c80c910f15f9e98da078e4c2eefac4c8c8e"}
var ErrKeyNotFound = errors.New("keyring: not found")

// /sedum:setVar
// sedum:setMethod {"v":2,"tier":"owned","record":"prov-2026-3b6a607f","gen":"sha256:1f260ca3b5b86a51abce055a71ae8dee641829534e4f551a75c1a2f5a646fb55","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","id":{"file":"auth/store.go","name":"Delete","receiver":"KeyringStore"},"kwargs":"sha256:14e4bef67997132330517f52e8bcf3ed37fbb1551ef112c1619440a0ef15c60d","body":"sha256:1d84a0b33d4a69166d334e04bf4fb7daa149ba662b15beca11c99a6653e4d98e"}
func (s KeyringStore) Delete() error {
	_ = s.Keyring.Delete(keyringService, keyringUser)
	return s.Fallback.Delete()
}

// /sedum:setMethod
// sedum:setType {"v":2,"tier":"owned","record":"prov-2026-47a53069","gen":"sha256:bf49404badae550c4d37cc91d01098aed7543ce05b441a4e9520a54804e48559","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","id":{"file":"auth/store.go","name":"KeyringStore"},"kwargs":"sha256:3806b66ce61231df9e7fa5fee86ff0b87e00e0afafaf85aa9962c5743907104f","body":"sha256:0df8ab3a680871dbc90b07c0696b1de320e938837945928cf0b9ff910eeca382"}
type KeyringStore struct {
	Keyring  Keyring
	Fallback Store
}

// /sedum:setType
// sedum:setMethod {"v":2,"tier":"owned","record":"prov-2026-535555b5","gen":"sha256:9284b0e87c8dc6ceb232329e6fa2e11a0bce963b684d134569b6e61fe3cc7e53","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","id":{"file":"auth/store.go","name":"Delete","receiver":"FileStore"},"kwargs":"sha256:cc74b9da786d6d82ac18897a8af07baeac4312583538564165370141b9ba8919","body":"sha256:8ee13230a6010d228651832952f06252a6bacd49586d3e38e766c38edb883e46"}
func (s FileStore) Delete() error {
	err := os.Remove(s.Path)
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	return err
}

// /sedum:setMethod
// sedum:setMethod {"v":2,"tier":"owned","record":"prov-2026-83941993","gen":"sha256:a6cdf2f9d87c9b75dddfcdd7baabd8ce619ddf6e04741fbe58f3851681ad8c9b","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","id":{"file":"auth/store.go","name":"Save","receiver":"KeyringStore"},"kwargs":"sha256:a06cd7644a24b9e71581ff591541be5ff14211c55f2224211aaa061f35eef1b2","body":"sha256:dd98570ef8f606875ba6494ed350c16e6de3dc8c8fa9140284be0292bd0fbdf7"}
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
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-8c290540","gen":"sha256:eefe6ed94ee80932048eba9289ba4b65b1418c97c0f4636bd8cd98b07525d7a0","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"auth/store.go","name":"DefaultStore"},"kwargs":"sha256:58abf196a6074c7575e2e71c26c7ed0638f2a34331489e04dc97fb16fd3776b3","body":"sha256:5c4b6aff77aa7e5266f712e8979d34261b67bd53f33910ac1d8cdff5431fffde"}
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
// sedum:setConst {"v":2,"tier":"owned","record":"prov-2026-ba2fe91f","gen":"sha256:4e4fa845b8c36b60bd817ceb779309ac9350bcba85f6126d26ce440e0621eda0","pkg":"sha256:51be1c18998d63cc558c253e211a6bd188d96cb251dad68c6ea88e03c2d5cc4c","id":{"file":"auth/store.go","name":"keyringService"},"kwargs":"sha256:6cf77c5f27a2206b8aa25615a26c1e4c0f9e9cb0f7decdecbb5bb7193836bd32","body":"sha256:baeed8c87bdbfb5e2391ca095770975d850b46aadd19472bd77b8b1875186b02"}
const (
	keyringService = "provisio"
	keyringUser    = "provenance-cli"
)

// /sedum:setConst
// sedum:setType {"v":2,"tier":"owned","record":"prov-2026-cac00d74","gen":"sha256:2a51bd29dbae4fb3316d65c8179a3967adb234c4813dca5304308ab7879ca791","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","id":{"file":"auth/store.go","name":"FileStore"},"kwargs":"sha256:812cab3a0bd0e6b2d98ed46fcabc143d8cd77c4121fb825f043f93aa1127eb42","body":"sha256:64d1b31a2406794689c1fb432a27f11e0b090abce009bdddfbd86309c00374f5"}
type FileStore struct {
	Path string
}

// /sedum:setType
// sedum:setMethod {"v":2,"tier":"owned","record":"prov-2026-ce5c322b","gen":"sha256:0020a6666c64e5f4edd88a9cbf915c57105dac4e1c90fe97dc2cad2719d5b7a0","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","id":{"file":"auth/store.go","name":"Load","receiver":"KeyringStore"},"kwargs":"sha256:9c87194932a8633aca6c60d149042454729810d3169e62ca623d5c32a8db27ce","body":"sha256:49e60184c0ebc4bf41f47c8f0492d85fc4c2873ea53bf7fe4ce8954f85301d02"}
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
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-d69bf77a","gen":"sha256:ee023230754d05a5997ae0cf33a00bbac05b8438ed302fb745cc14ad4f5fdd7f","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"auth/store.go","name":"DefaultFilePath"},"kwargs":"sha256:b2e94cfc720e1007ebb8070ca9e2f1dcd3e8ba96f5a30d19e121496140e10211","body":"sha256:6465e8277d2f03f7e79bd154ec100286b051ea4be025dc8bc99557d5cc1e00a4"}
func DefaultFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("cannot find the home directory: %w", err)
	}
	return filepath.Join(home, ".provisio", "credentials"), nil
}

// /sedum:setFunction
// sedum:setType {"v":2,"tier":"owned","record":"prov-2026-ee489fd6","gen":"sha256:6222ed3c8cf3985f49c881df8f85cf53458d279745f9d2a7885950294e6ef790","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","id":{"file":"auth/store.go","name":"Store"},"kwargs":"sha256:da2266e10de81afb7deccd16630d960edc9004bf2345b07cd99089e76af31024","body":"sha256:2296e99c0f1e3af3c4fac98d251fe053d86ff8f6bacdf7366fe4b5d58e0a678c"}
type Store interface {
	Load() (*Credential, Source, error)
	Save(c *Credential) error
	Delete() error
}

// /sedum:setType
// sedum:setMethod {"v":2,"tier":"owned","record":"prov-2026-ef1c214e","gen":"sha256:0b1a771e73121250c9df3048cf896354b367f346693b389c16d73f3c09a57699","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","id":{"file":"auth/store.go","name":"Save","receiver":"FileStore"},"kwargs":"sha256:d784b189e0c12f0709704af3fbb42c5f8ebe568adaebc0c5ef43b0539544d850","body":"sha256:2b33bd842d5c2e587ddeebb2927a0c76162a2acb87e72b9fbeaf6b044dd84b8f"}
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
