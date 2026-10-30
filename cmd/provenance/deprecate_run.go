package main

// sedum:anchor:imports
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-35439261","gen":"sha256:bb73fb8a27febfb6636afc91c6f46ef087ceb45d41e95db14a66c7469f2256d4","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/deprecate_run.go","path":"errors"},"kwargs":"sha256:71bdb8734404b0c3604069aee2d1df95e3ae119785a97d34678c5f725ea959a1","body":"sha256:2b718237ba0a56467c8fd494cc131e942ef410e6228c03b644d70476c71862dc"}
import "errors"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-35439261","gen":"sha256:bb73fb8a27febfb6636afc91c6f46ef087ceb45d41e95db14a66c7469f2256d4","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/deprecate_run.go","path":"fmt"},"kwargs":"sha256:94b3b2827a7559d6001741f689cfd522fc61752c6e8fe0396168c09f3928dcf0","body":"sha256:bcbac51883cbf8064e9196b371cb880cee429866652447ad557533d817857609"}
import "fmt"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-35439261","gen":"sha256:bb73fb8a27febfb6636afc91c6f46ef087ceb45d41e95db14a66c7469f2256d4","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/deprecate_run.go","path":"io"},"kwargs":"sha256:d3138e88b7181d96eaf492344d3295290e938c1e049e05ee406ca10ac7b63e19","body":"sha256:ef8e8abf4ef5ee40a19dc85b49bf09c5df1fe32abfe834e0ab4c2175baf80ddf"}
import "io"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-35439261","gen":"sha256:bb73fb8a27febfb6636afc91c6f46ef087ceb45d41e95db14a66c7469f2256d4","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/deprecate_run.go","path":"path/filepath"},"kwargs":"sha256:c895c30a3ec9955ac8942b30ac027e0223c9bdff4d9fbd2efad8b238b0f63988","body":"sha256:a6d196b558491e646be93f56a321f339068c64c76aa9d49bb18046e7f9c777a1"}
import "path/filepath"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-35439261","gen":"sha256:bb73fb8a27febfb6636afc91c6f46ef087ceb45d41e95db14a66c7469f2256d4","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/deprecate_run.go","path":"github.com/spf13/cobra"},"kwargs":"sha256:af98fd6943c15f0ab7b292d0ee0cebb249b35940de44072df5a61b3e32787856","body":"sha256:f361c02d0f8e1999cf3fb715393e16310eeec9debd5aaa13936af7617e2b46aa"}
import "github.com/spf13/cobra"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-35439261","gen":"sha256:bb73fb8a27febfb6636afc91c6f46ef087ceb45d41e95db14a66c7469f2256d4","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/deprecate_run.go","path":"github.com/livecodelife/provenance/lint"},"kwargs":"sha256:962619f491aa09a4fa03eacc383e92c1d986ae5cfe8a1a3a2f1b44c1eda56c0c","body":"sha256:9abff5a2d82dd8a4177704699227186fed48d7e0d093d2d2d5c565b09ff787d7"}
import "github.com/livecodelife/provenance/lint"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-35439261","gen":"sha256:bb73fb8a27febfb6636afc91c6f46ef087ceb45d41e95db14a66c7469f2256d4","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/deprecate_run.go","path":"github.com/livecodelife/provenance/record"},"kwargs":"sha256:2ad585e592154f479a072f8fe27d58f8a53e12146ea4a174d76a2712e1695e16","body":"sha256:acbf8628061291d11aa39e9ea64ecc91095a117057cfe6f1958ef2c437fac003"}
import "github.com/livecodelife/provenance/record"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-35439261","gen":"sha256:bb73fb8a27febfb6636afc91c6f46ef087ceb45d41e95db14a66c7469f2256d4","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/deprecate_run.go","name":"runDeprecate"},"kwargs":"sha256:0eb82f9ea1f128d9368211be1dea5d43ffbeff5c598ad043a4cfdedf16f20262","body":"sha256:4a42cab5a2cec5352ef008976347dee76ca6b35b477712d1125e8ab8ef78d834"}
func runDeprecate(out io.Writer, cmd *cobra.Command, args []string) error {
	r, err := openRepo(cmd)
	if err != nil {
		return err
	}
	id, _ := cmd.Flags().GetString("record")
	if id == "" {
		return errors.New(`required flag(s) "record" not set`)
	}
	reason, _ := cmd.Flags().GetString("reason")
	if reason == "" {
		return errors.New(`required flag(s) "reason" not set`)
	}
	rec, doc, err := r.resolve(id)
	if err != nil {
		return err
	}
	if rec.Schema != 2 {
		r.fail("Cannot deprecate %s: it is a v1 record; run 'provenance migrate' first.", rec.ID)
		return errors.New("not a v2 record")
	}
	if rec.Status == record.StatusDeprecated {
		r.fail("Record %s is already deprecated", rec.ID)
		return errors.New("already deprecated")
	}
	if err := lint.CheckTransition(rec.Status, record.StatusDeprecated, false); err != nil {
		r.fail("Cannot deprecate %s: %v", rec.ID, err)
		return err
	}
	rel := r.recordPath(rec.ID)
	states := snapshotFiles(filepath.Join(r.Root, rel))
	orig := rec.Status
	rec.Status = record.StatusDeprecated
	if err := r.Store.Save(rec, doc); err != nil {
		restoreFiles(states)
		r.fail("Failed to save record: %v", err)
		return err
	}
	if r.Config.CommitOnStatusChange {
		msg := "Deprecate provenance record " + tag(rec.ID, rec.Version) + ": " + reason
		if err := r.commit(msg, rel); err != nil {
			restoreFiles(states)
			r.unstage(rel)
			r.fail("Cannot deprecate %s: the auto-commit was rejected:\n\n    %v\n\n  The transition was rolled back — %s is unchanged (still %s). Nothing was staged.\n  Resolve the issue above (e.g. a failing pre-commit hook)\n  and run 'provenance deprecate' again.", rec.ID, err, rec.ID, orig)
			return err
		}
	}
	reconcileQuietly(r)
	fmt.Fprintf(out, "\n✓ %s marked as deprecated\n\n", rec.ID)
	return nil
}

// /sedum:setFunction
