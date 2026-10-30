package main

// sedum:anchor:imports
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-81333e45","gen":"sha256:d3d637e42ac358905da0d78c2d2a3cc3e324f2d89cca02317556039be91b66b1","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/session_run.go","path":"github.com/livecodelife/provenance/auth"},"kwargs":"sha256:d575fbc5f22be6d027b36f9d273c80a0286d90451797740adf49e4444e737056","body":"sha256:57cbe62cd356e3c713dcc35195098163d64975ccdb05ef9a2317cf089e95c547"}
import "github.com/livecodelife/provenance/auth"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-81333e45","gen":"sha256:d3d637e42ac358905da0d78c2d2a3cc3e324f2d89cca02317556039be91b66b1","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/session_run.go","path":"github.com/livecodelife/provenance/remote"},"kwargs":"sha256:d7f205a2cb2c361adbf248fe1338adf6b561127822dcab0e7a7ec5a534b6cec6","body":"sha256:8d05ee0496574b6404cf7e9fbb5899ea912303f93ead4d7ded94d22f09a229d2"}
import "github.com/livecodelife/provenance/remote"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-a0ad0a7f","gen":"sha256:dd29d8cc1590b0342455e97aa87320359bfbfee4affbf741dcc86c80a497061c","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/session_run.go","path":"context"},"kwargs":"sha256:5dd24f36591d31fd12a8247df02a3f3f3f3de3913e3627c97d12efb8abb09f48","body":"sha256:0f7443fb25d7ad6a970617d4d0df6123283e24801b94c1d84fbe431230f31e1e"}
import "context"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-a0ad0a7f","gen":"sha256:dd29d8cc1590b0342455e97aa87320359bfbfee4affbf741dcc86c80a497061c","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/session_run.go","path":"time"},"kwargs":"sha256:ed8a0fc7b10a2be8f32fb91588d229ba3774b52df5055a00a66dcec9ddb93dd8","body":"sha256:e99128ebf4e36998f40b33adf2eef57f4d6c315f1a4d1bdd3b7fc09a0d7bc82d"}
import "time"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-a51283a0","gen":"sha256:25743a0c40a8eff1905ae8c89926e35e8eb0ccddd6b35fc4be3bc508c591cfa4","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/session_run.go","path":"os"},"kwargs":"sha256:539bb031faaa606a7f51538cfd041f90654d90002ec839850891f3dc10db9a20","body":"sha256:7c6e309eb2575d4aa451ec2ab83161d13e9ae03bbca3a4d2ecf30ce12adf0905"}
import "os"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-ae5b9ab4","gen":"sha256:906ad1c7ebf364264b0d30e9a0453b3e9ec1897578d2c9a4f4f1e2c07cc1787a","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/session_run.go","path":"errors"},"kwargs":"sha256:218ad264b95d9f1bdfed3c4075de405e4af44a51015d1f2b56d15a6274d28df1","body":"sha256:2b718237ba0a56467c8fd494cc131e942ef410e6228c03b644d70476c71862dc"}
import "errors"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-da5d8cb8","gen":"sha256:1802c6e80df79a3253b8c10809eb08ec0bc9b3e9baaf85e48a0ef0902dc6ab1d","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/session_run.go","path":"path/filepath"},"kwargs":"sha256:d7ce4ba9138b9ac768eefe7cbba1f2f712cd31b0df41d83482ff6fcc89ae342a","body":"sha256:a6d196b558491e646be93f56a321f339068c64c76aa9d49bb18046e7f9c777a1"}
import "path/filepath"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-81333e45","gen":"sha256:d3d637e42ac358905da0d78c2d2a3cc3e324f2d89cca02317556039be91b66b1","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/session_run.go","name":"sessionError"},"kwargs":"sha256:d0532e7a767c9a50af87d37c0410f07f0da2b20d9441a2dc404c73a79e26521f","body":"sha256:fe4170eab448f148884775d4b18dc06d847fae82c0edded62549a86e29283d58"}
func sessionError(err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, auth.ErrNoSession) {
		return &exitError{code: 3, err: errors.New(msgNeedSession)}
	}
	if errors.Is(err, auth.ErrSessionExpired) {
		return &exitError{code: 3, err: errors.New(msgSessionExpire)}
	}
	if errors.Is(err, auth.ErrTokenRejected) {
		return &exitError{code: 3, err: errors.New(msgTokenRejected)}
	}
	var apiErr *remote.APIError
	if errors.As(err, &apiErr) && apiErr.Status == 403 && apiErr.Code == "insufficient_scope" {
		return &exitError{code: 3, err: errors.New(msgScope)}
	}
	return err
}

// /sedum:setFunction
// sedum:setMethod {"v":2,"tier":"owned","record":"prov-2026-15de5191","gen":"sha256:e6a1b0e7ffcfb8066b5bd1951009d4137af9f4681a055e66c29b2716aa92e6f6","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","id":{"file":"cmd/provenance/session_run.go","name":"ExitCode","receiver":"exitError"},"kwargs":"sha256:cf1874a76ca53c2edebbe4f6dd910917b75c4761c31fc52720dbf9f0ed07c9c1","body":"sha256:04cea1770978e5ff00fc0d2c61f1e09cb298e3bb7332b1473b2e4b0460849ba4"}
func (e *exitError) ExitCode() int {
	return e.code
}

// /sedum:setMethod
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-28d01dcf","gen":"sha256:e6659a3a24bb52c37733b2e89ed757564871e737a0e65d5606ac6f05b2fe6d77","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/session_run.go","name":"withSession"},"kwargs":"sha256:97640422e02351f9d8b0fabc6755113c93672ffa27d69504303a7cf9c65499ae","body":"sha256:0400851fd7789728fb09cd17a3756d0ec600fb388efce7047db2d4746bf20dc6"}
func withSession(ctx context.Context, fn func(string) error) error {
	m, err := sessionManager()
	if err != nil {
		return err
	}
	return sessionError(m.Do(ctx, fn))
}

// /sedum:setFunction
// sedum:setMethod {"v":2,"tier":"owned","record":"prov-2026-5a7e47c4","gen":"sha256:27b426390755927f3e455f7d4bdb1ff0cf4d4e70d94e004fd28d34526f6a9b5a","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","id":{"file":"cmd/provenance/session_run.go","name":"Unwrap","receiver":"exitError"},"kwargs":"sha256:c0da6fcbd05b085d86376eb141ece4a7ef5e729d47cf4de95e3369c11c5bfd5b","body":"sha256:de5373f318af6c9278d0c40a3475a9c25b8c5ddb80f2e0ce3ce173d20bf22746"}
func (e *exitError) Unwrap() error {
	return e.err
}

// /sedum:setMethod
// sedum:setMethod {"v":2,"tier":"owned","record":"prov-2026-790e16a4","gen":"sha256:a9bad494046cdb251f9dad0b3915a76553b5a74995da5e61ed10251a9e0c5620","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","id":{"file":"cmd/provenance/session_run.go","name":"Error","receiver":"exitError"},"kwargs":"sha256:3112339fed4b94461bb9603153dc8d65a46ead5ec4672a5f4dc6422ebb6e5b91","body":"sha256:2304464d1e907e3e6c08e5203214d632f86a40e34485401316cce6208d39a7d8"}
func (e *exitError) Error() string {
	return e.err.Error()
}

// /sedum:setMethod
// sedum:setConst {"v":2,"tier":"owned","record":"prov-2026-9895ce2c","gen":"sha256:e94a96baa86aacca7cd4be62ba3b7445cf3692cbc25e23d6cee2f6b0f5c4e5d2","pkg":"sha256:51be1c18998d63cc558c253e211a6bd188d96cb251dad68c6ea88e03c2d5cc4c","id":{"file":"cmd/provenance/session_run.go","name":"msgNeedSession"},"kwargs":"sha256:09e2b53f3eab74142972a70978b8b0296e86072c0e09cc44228f46fbceffffbb","body":"sha256:3a99dfcccec3210715e25de841141c60c479ce888ef6873db2c0007048bb4ee2"}
const (
	msgNeedSession   = "This command needs a Provisio session. Run provenance login (or set PROVISIO_TOKEN in CI)."
	msgSessionExpire = "Your Provisio session has expired. Run provenance login."
	msgTokenRejected = "The personal token was rejected (revoked or expired). Replace PROVISIO_TOKEN, or run provenance login."
	msgScope         = "Your Provisio sign-in lacks permission to write; sign in again with `provenance login`."
)

// /sedum:setConst
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-a0ad0a7f","gen":"sha256:dd29d8cc1590b0342455e97aa87320359bfbfee4affbf741dcc86c80a497061c","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/session_run.go","name":"withTimeout"},"kwargs":"sha256:afe9cb6ca087e33b0b35da90ef351c26bcb27c24b9a20ff420b803362da07fdf","body":"sha256:e5168ee87e450efb699cc5351bffa9e6bfa04702a08949d1b70e8bfe3d1e7e6e"}
func withTimeout(ctx context.Context, d time.Duration) (context.Context, func()) {
	return context.WithTimeout(ctx, d)
}

// /sedum:setFunction
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-a51283a0","gen":"sha256:25743a0c40a8eff1905ae8c89926e35e8eb0ccddd6b35fc4be3bc508c591cfa4","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/session_run.go","name":"serverURL"},"kwargs":"sha256:985e348359dd0c1862761629c44b7906cc34ed9363b4a31c01b37b15d735a1cb","body":"sha256:13b1cdcf9c25880a5afe9fc0b6fcef0b835d227336fff35a7ac7411d24f72e7b"}
func serverURL() string {
	if v := os.Getenv(auth.EnvServer); v != "" {
		return v
	}
	return auth.DefaultServer
}

// /sedum:setFunction
// sedum:setType {"v":2,"tier":"owned","record":"prov-2026-ac80218b","gen":"sha256:0d39be9d85fb64f164375608f5889fed40cf83945504b6e57fa1f8a5af3ef08a","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","id":{"file":"cmd/provenance/session_run.go","name":"exitError"},"kwargs":"sha256:772077e089c46578d5635a46c6cdd513a36b9edcc633dfd717aacc26f95afa27","body":"sha256:61883101c121ce3920aa9e6786147658f7807b3ac2096568954282d8caf6451b"}
type exitError struct {
	code int
	err  error
}

// /sedum:setType
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-ae5b9ab4","gen":"sha256:906ad1c7ebf364264b0d30e9a0453b3e9ec1897578d2c9a4f4f1e2c07cc1787a","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/session_run.go","name":"exitCode"},"kwargs":"sha256:c739cbcc28860a610087d2913184a881f285c75e68b8327a98895f1cfc14ebce","body":"sha256:c7554c9ef261ce351487459f030cfe56d3275e5c694d4f4c7d46c54887fcf5a4"}
func exitCode(err error) int {
	var ee *exitError
	if errors.As(err, &ee) {
		return ee.code
	}
	return 1
}

// /sedum:setFunction
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-da5d8cb8","gen":"sha256:1802c6e80df79a3253b8c10809eb08ec0bc9b3e9baaf85e48a0ef0902dc6ab1d","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/session_run.go","name":"sessionManager"},"kwargs":"sha256:e6d12941e9648931aae9660895861b7ad35f42868e0a388b899b9a73723161ed","body":"sha256:633d0ec30a10b5ef73f855bc14d1c26bb0dd271b8db53b32f9241bb4e9019a4f"}
func sessionManager() (*auth.Manager, error) {
	store, err := auth.DefaultStore()
	if err != nil {
		return nil, err
	}
	path, err := auth.DefaultFilePath()
	if err != nil {
		return nil, err
	}
	server := serverURL()
	if os.Getenv(auth.EnvServer) == "" && os.Getenv(auth.EnvToken) == "" {
		if cred, _, err := store.Load(); err == nil && cred != nil && cred.Server != "" {
			server = cred.Server
		}
	}
	return &auth.Manager{
		Store:    store,
		Client:   &auth.Client{Server: server},
		LockPath: filepath.Join(filepath.Dir(path), "credentials.lock"),
	}, nil
}

// /sedum:setFunction
