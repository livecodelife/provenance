package main

// sedum:anchor:imports
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-89075633","gen":"sha256:2ba4ade1cd56f6970d70eae26bd4332bdbba4e9399976df4385be19fb6ff45a2","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/logout_run.go","path":"errors"},"kwargs":"sha256:3b18d07cb16ade01b99f4efcf82a9a3e6b31737d2e7d6ecf991d1431449a88b1","body":"sha256:2b718237ba0a56467c8fd494cc131e942ef410e6228c03b644d70476c71862dc"}
import "errors"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-89075633","gen":"sha256:2ba4ade1cd56f6970d70eae26bd4332bdbba4e9399976df4385be19fb6ff45a2","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/logout_run.go","path":"fmt"},"kwargs":"sha256:67232a347736453c1869d2ef9cf2a5912236a9815a569083c9535870d62ac339","body":"sha256:bcbac51883cbf8064e9196b371cb880cee429866652447ad557533d817857609"}
import "fmt"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-89075633","gen":"sha256:2ba4ade1cd56f6970d70eae26bd4332bdbba4e9399976df4385be19fb6ff45a2","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/logout_run.go","path":"io"},"kwargs":"sha256:81866f90e20c634ae5f5f008366594a2761f7e013fba792514a62c7824d03760","body":"sha256:ef8e8abf4ef5ee40a19dc85b49bf09c5df1fe32abfe834e0ab4c2175baf80ddf"}
import "io"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-89075633","gen":"sha256:2ba4ade1cd56f6970d70eae26bd4332bdbba4e9399976df4385be19fb6ff45a2","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/logout_run.go","path":"os"},"kwargs":"sha256:7006d9d9f947ca5f7cf6489e86e58a0f77bc2f28bebec8fc6689e86a8e20d952","body":"sha256:7c6e309eb2575d4aa451ec2ab83161d13e9ae03bbca3a4d2ecf30ce12adf0905"}
import "os"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-89075633","gen":"sha256:2ba4ade1cd56f6970d70eae26bd4332bdbba4e9399976df4385be19fb6ff45a2","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/logout_run.go","path":"time"},"kwargs":"sha256:7d48b55e3e062a491cf95a241e9973557889c9e5a4bb070f9dd44ee1912003c4","body":"sha256:e99128ebf4e36998f40b33adf2eef57f4d6c315f1a4d1bdd3b7fc09a0d7bc82d"}
import "time"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-89075633","gen":"sha256:2ba4ade1cd56f6970d70eae26bd4332bdbba4e9399976df4385be19fb6ff45a2","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/logout_run.go","path":"github.com/spf13/cobra"},"kwargs":"sha256:28ec9b83940ace2646044b05bba88facdd033e0c3b4fdb8eb3bc3e8a50235188","body":"sha256:f361c02d0f8e1999cf3fb715393e16310eeec9debd5aaa13936af7617e2b46aa"}
import "github.com/spf13/cobra"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-89075633","gen":"sha256:2ba4ade1cd56f6970d70eae26bd4332bdbba4e9399976df4385be19fb6ff45a2","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/logout_run.go","path":"github.com/livecodelife/provenance/auth"},"kwargs":"sha256:ca2510107b5681e7e605950a26bf02824bb040d7033564475e9f4673b77922b2","body":"sha256:57cbe62cd356e3c713dcc35195098163d64975ccdb05ef9a2317cf089e95c547"}
import "github.com/livecodelife/provenance/auth"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-89075633","gen":"sha256:2ba4ade1cd56f6970d70eae26bd4332bdbba4e9399976df4385be19fb6ff45a2","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/logout_run.go","name":"runLogout"},"kwargs":"sha256:457cd9857f629d16d2689d0151a799a42e92243f7eeaa5e078652c7042196e30","body":"sha256:17f9008b3b5dc6d5420981084e5e524c4f8304777e515716ad59aa1b4ecd3f19"}
func runLogout(out io.Writer, cmd *cobra.Command, args []string) error {
	store, err := auth.DefaultStore()
	if err != nil {
		return err
	}
	cred, _, err := store.Load()
	if err != nil {
		if errors.Is(err, auth.ErrNoSession) {
			fmt.Fprintln(out, "Not signed in.")
			return nil
		}
		return err
	}
	var revokeErr error
	if cred.RefreshToken != "" {
		ctx, cancel := withTimeout(cmd.Context(), 15*time.Second)
		client := &auth.Client{Server: cred.Server}
		revokeErr = client.Logout(ctx, cred.RefreshToken)
		cancel()
	}
	if err := store.Delete(); err != nil {
		return err
	}
	fmt.Fprintln(out, "✓ Signed out; the local credential is deleted.")
	if revokeErr != nil {
		fmt.Fprintf(out, "  The server could not be reached to revoke the session (%s); it expires on its own.\n", revokeErr.Error())
	}
	if cred.PersonalToken != "" {
		fmt.Fprintln(out, "  A personal token stays valid until you delete it in the Provisio UI.")
	}
	if v := os.Getenv(auth.EnvToken); v != "" {
		fmt.Fprintf(out, "  %s is still set in this environment.\n", auth.EnvToken)
	}
	return nil
}

// /sedum:setFunction
