package auth

// sedum:anchor:imports
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-1396d0cc","gen":"sha256:059dfb1a1030a88cbdb7f96836f4fd05ee5190e3ebeb43c80ca494d65d1030d7","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"auth/keyring_system.go","path":"errors"},"kwargs":"sha256:d6c76d3e2e032eb211553bc1841258104b79200afa239f390d935ab75d6ce551","body":"sha256:2b718237ba0a56467c8fd494cc131e942ef410e6228c03b644d70476c71862dc"}
import "errors"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-1396d0cc","gen":"sha256:059dfb1a1030a88cbdb7f96836f4fd05ee5190e3ebeb43c80ca494d65d1030d7","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"auth/keyring_system.go","path":"github.com/zalando/go-keyring"},"kwargs":"sha256:c719191787f34fe4ee037e083a1c1fa3f2124b647587487407348eaed7761728","body":"sha256:c8fd639ddf392c0e818d5f230079b2531b82fef86160646db6de2fc077f420bd"}
import "github.com/zalando/go-keyring"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setMethod {"v":2,"tier":"owned","record":"prov-2026-1396d0cc","gen":"sha256:059dfb1a1030a88cbdb7f96836f4fd05ee5190e3ebeb43c80ca494d65d1030d7","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","id":{"file":"auth/keyring_system.go","name":"Get","receiver":"systemKeyring"},"kwargs":"sha256:06863103b39a3c2d0a0e313e25ca30b255575f8e4aa45754bd3c38636d4b3dcd","body":"sha256:895c4113ffab7e9f34a802f428cf2eee3ee6f523868120c6500dcddea73b9965"}
func (systemKeyring) Get(service, user string) (string, error) {
	secret, err := keyring.Get(service, user)
	if errors.Is(err, keyring.ErrNotFound) {
		return "", ErrKeyNotFound
	}
	return secret, err
}

// /sedum:setMethod
// sedum:setMethod {"v":2,"tier":"owned","record":"prov-2026-274813c6","gen":"sha256:2297630466d7322dabdc5319b12a45a5b3bb573c5d07d627fc26b29cfad9ccee","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","id":{"file":"auth/keyring_system.go","name":"Set","receiver":"systemKeyring"},"kwargs":"sha256:0cb6f66c386f176b07cf99cd6d8c8479ff534737d0ab234649a7c2fed43cc144","body":"sha256:1fed6198004da7baa4793fe9058e4febb21857deb1bf0bc4113084d5b57b9df8"}
func (s systemKeyring) Set(service, user, value string) error {
	return keyring.Set(service, user, value)
}

// /sedum:setMethod
// sedum:setMethod {"v":2,"tier":"owned","record":"prov-2026-677ba83d","gen":"sha256:c3a590dff65597edb09f53bda0a33406bd6a274da33a3d44cdf8f68b9eb42db0","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","id":{"file":"auth/keyring_system.go","name":"Delete","receiver":"systemKeyring"},"kwargs":"sha256:29b33c3a608bd10fdd5327d631e6ef1552d72d7d97d554ae30520ae4cdc53d31","body":"sha256:93f345b97a4a6ac67774dd83ffd4c1f7cae6e8e9dc92e76b5be481af6aed8d8c"}
func (systemKeyring) Delete(service, user string) error {
	err := keyring.Delete(service, user)
	if errors.Is(err, keyring.ErrNotFound) {
		return nil
	}
	return err
}

// /sedum:setMethod
// sedum:setType {"v":2,"tier":"owned","record":"prov-2026-8f5a178b","gen":"sha256:100328ab485ae7a0090feac05b5b11c1004070e46d26ab8d99a37c3123337d90","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","id":{"file":"auth/keyring_system.go","name":"systemKeyring"},"kwargs":"sha256:93a96768049030bad9b5984089768c82a34e304217a7d6b00a7a3f09dd71ca5a","body":"sha256:0e95562df369f99f00cd00228fea3af9639a6ec653a4b687b3b88933ad8be96c"}
type systemKeyring struct{}

// /sedum:setType
