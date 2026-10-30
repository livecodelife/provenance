package main

// sedum:anchor:imports
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-adabd65f","gen":"sha256:4ae46f3089d333bd12c54eda92a2d20ae90a3ad1ebf7c65d3eda625c0356ee02","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/version_run.go","path":"fmt"},"kwargs":"sha256:94fd049e07c446ed6f6630b760f417bbd7172111502a1bea744dc199202e0abc","body":"sha256:bcbac51883cbf8064e9196b371cb880cee429866652447ad557533d817857609"}
import "fmt"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-adabd65f","gen":"sha256:4ae46f3089d333bd12c54eda92a2d20ae90a3ad1ebf7c65d3eda625c0356ee02","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/version_run.go","path":"io"},"kwargs":"sha256:062d68a39042102834fdc6381e5619befe2565bdb1bd2daf8e0fdd13d21cf8ec","body":"sha256:ef8e8abf4ef5ee40a19dc85b49bf09c5df1fe32abfe834e0ab4c2175baf80ddf"}
import "io"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-adabd65f","gen":"sha256:4ae46f3089d333bd12c54eda92a2d20ae90a3ad1ebf7c65d3eda625c0356ee02","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/version_run.go","path":"github.com/spf13/cobra"},"kwargs":"sha256:ea20b171c56342d942c58dd5f27810186e802979b0f2453489a0bb0bd962fbc9","body":"sha256:f361c02d0f8e1999cf3fb715393e16310eeec9debd5aaa13936af7617e2b46aa"}
import "github.com/spf13/cobra"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-adabd65f","gen":"sha256:4ae46f3089d333bd12c54eda92a2d20ae90a3ad1ebf7c65d3eda625c0356ee02","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/version_run.go","path":"github.com/livecodelife/provenance/version"},"kwargs":"sha256:179f2e0ab05d406f5f245870033b8ced33455222644f1d6ffc39aba45b93fe37","body":"sha256:4e7c63ee372e1974951764cd1d475c3b4a32cabd87143982ff3cc1812195ce5e"}
import "github.com/livecodelife/provenance/version"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-adabd65f","gen":"sha256:4ae46f3089d333bd12c54eda92a2d20ae90a3ad1ebf7c65d3eda625c0356ee02","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/version_run.go","name":"runVersion"},"kwargs":"sha256:78d724a9b49e74633166f7298937355362647efa913f3fe752eb34d63a66a61e","body":"sha256:2cd7860393a1ece841b7082e528619dc82ace5a31c57df7bbc789ebb198a3ebb"}
func runVersion(out io.Writer, cmd *cobra.Command, args []string) error {
	_, err := fmt.Fprintln(out, version.Version)
	return err
}

// /sedum:setFunction
