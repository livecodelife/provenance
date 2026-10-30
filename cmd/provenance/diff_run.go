package main

// sedum:anchor:imports
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-bfd55535","gen":"sha256:98fb8e2f8fd67f02e2cabc1153d5dfb84f1f27f5c19d0582f9236cd5fa8e4f10","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/diff_run.go","path":"errors"},"kwargs":"sha256:5eea977897684cc049bfe4ea10f9ec3d6241fbbe34c1e0f2e53fd4f5a6364820","body":"sha256:2b718237ba0a56467c8fd494cc131e942ef410e6228c03b644d70476c71862dc"}
import "errors"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-bfd55535","gen":"sha256:98fb8e2f8fd67f02e2cabc1153d5dfb84f1f27f5c19d0582f9236cd5fa8e4f10","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/diff_run.go","path":"fmt"},"kwargs":"sha256:b069523ac99801b457155518590890e57ace2ea9a4bbd179b56ce1a47997a548","body":"sha256:bcbac51883cbf8064e9196b371cb880cee429866652447ad557533d817857609"}
import "fmt"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-bfd55535","gen":"sha256:98fb8e2f8fd67f02e2cabc1153d5dfb84f1f27f5c19d0582f9236cd5fa8e4f10","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/diff_run.go","path":"io"},"kwargs":"sha256:c4c0b19975a40d1dc38ceff565c69bf09c06826baf96c65e5c266c52f0e7e535","body":"sha256:ef8e8abf4ef5ee40a19dc85b49bf09c5df1fe32abfe834e0ab4c2175baf80ddf"}
import "io"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-bfd55535","gen":"sha256:98fb8e2f8fd67f02e2cabc1153d5dfb84f1f27f5c19d0582f9236cd5fa8e4f10","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/diff_run.go","path":"strconv"},"kwargs":"sha256:c2b7b2431e3fc6f1754bcb297c9134f1b710b47a4ecde8fbf0ca2cca4a5265f0","body":"sha256:f33f981c64a42d78a5daa7ff3fca15b10045080cccd282fef544b59eaf7f6dc8"}
import "strconv"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-bfd55535","gen":"sha256:98fb8e2f8fd67f02e2cabc1153d5dfb84f1f27f5c19d0582f9236cd5fa8e4f10","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/diff_run.go","path":"github.com/spf13/cobra"},"kwargs":"sha256:38e18910c12322c5b512dd56e44ce4259e9126ca25df723e20caf1e71484fdbb","body":"sha256:f361c02d0f8e1999cf3fb715393e16310eeec9debd5aaa13936af7617e2b46aa"}
import "github.com/spf13/cobra"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-bfd55535","gen":"sha256:98fb8e2f8fd67f02e2cabc1153d5dfb84f1f27f5c19d0582f9236cd5fa8e4f10","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/diff_run.go","path":"github.com/livecodelife/provenance/record"},"kwargs":"sha256:937e41d337098b56fb71f71a123cdc282a5e1e85e1678752ae97e12139a6b28d","body":"sha256:acbf8628061291d11aa39e9ea64ecc91095a117057cfe6f1958ef2c437fac003"}
import "github.com/livecodelife/provenance/record"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-bfd55535","gen":"sha256:98fb8e2f8fd67f02e2cabc1153d5dfb84f1f27f5c19d0582f9236cd5fa8e4f10","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/diff_run.go","name":"diffScalar"},"kwargs":"sha256:a85cd306adc56f16470e4955a5cbb360dec3bb2e4e8d8499bda52bb932a75325","body":"sha256:a6bb32431eef726eb9239fe7c37fb7378ad23282508f71996e8b45b310329da2"}
func diffScalar(out io.Writer, field, a, b string) bool {
	if a == b {
		return false
	}
	fmt.Fprintf(out, "  %s:\n    - %s\n    + %s\n", field, a, b)
	return true
}

// /sedum:setFunction
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-bfd55535","gen":"sha256:98fb8e2f8fd67f02e2cabc1153d5dfb84f1f27f5c19d0582f9236cd5fa8e4f10","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/diff_run.go","name":"diffList"},"kwargs":"sha256:86f36b5b2d5f92c7581a803486bbbb8859af7946206321243b843e7312776244","body":"sha256:2d5c9f6253a8a6725e5188b4c67d9d32a81a8b102647ff3f95256c26538d3399"}
func diffList(out io.Writer, field string, a, b []string) bool {
	inA := make(map[string]bool, len(a))
	for _, x := range a {
		inA[x] = true
	}
	inB := make(map[string]bool, len(b))
	for _, x := range b {
		inB[x] = true
	}
	var removed, added []string
	for _, x := range a {
		if !inB[x] {
			removed = append(removed, x)
		}
	}
	for _, x := range b {
		if !inA[x] {
			added = append(added, x)
		}
	}
	if len(removed) == 0 && len(added) == 0 {
		return false
	}
	fmt.Fprintf(out, "  %s:\n", field)
	for _, x := range removed {
		fmt.Fprintf(out, "    - %s\n", x)
	}
	for _, x := range added {
		fmt.Fprintf(out, "    + %s\n", x)
	}
	return true
}

// /sedum:setFunction
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-bfd55535","gen":"sha256:98fb8e2f8fd67f02e2cabc1153d5dfb84f1f27f5c19d0582f9236cd5fa8e4f10","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/diff_run.go","name":"specPaths"},"kwargs":"sha256:3e7836aa160617f10f37c34379c1c47e2482e177dfa60a0e46ade350e0f3890a","body":"sha256:a5fea3c1f3f5b3ae5ff08592608ef72f61fa7514a94988b1c26a17304006d6d3"}
func specPaths(specs []record.AssociatedSpec) []string {
	paths := make([]string, 0, len(specs))
	for _, s := range specs {
		paths = append(paths, s.Path)
	}
	return paths
}

// /sedum:setFunction
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-bfd55535","gen":"sha256:98fb8e2f8fd67f02e2cabc1153d5dfb84f1f27f5c19d0582f9236cd5fa8e4f10","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/diff_run.go","name":"runDiff"},"kwargs":"sha256:d7f844d8bd83e658a4778a9e46bbfef69afc704abf500594e3d7ac9ab8fc44e6","body":"sha256:0e1bdae74ca3a646bc76b2bfaf55e077bb234e789f6ed4cf12ff10f5fc4ec9a9"}
func runDiff(out io.Writer, cmd *cobra.Command, args []string) error {
	r, err := openRepo(cmd)
	if err != nil {
		return err
	}
	flag, err := cmd.Flags().GetString("record")
	if err != nil {
		return err
	}
	var id string
	var rest []string
	if flag != "" {
		id = flag
		rest = args
	} else {
		if len(args) == 0 {
			return errors.New("a record id is required (as the first argument or --record)")
		}
		id = args[0]
		rest = args[1:]
	}
	if len(rest) != 0 && len(rest) != 2 {
		return errors.New("give two versions to compare, or none")
	}
	var vA, vB int
	if len(rest) == 2 {
		vA, err = strconv.Atoi(rest[0])
		if err != nil {
			return fmt.Errorf("version %q is not a number", rest[0])
		}
		vB, err = strconv.Atoi(rest[1])
		if err != nil {
			return fmt.Errorf("version %q is not a number", rest[1])
		}
	}
	rec, _, err := r.resolve(id)
	if err != nil {
		return err
	}
	vs, err := r.Store.Versions(rec.ID)
	if err != nil {
		return err
	}
	if len(vs) == 0 {
		return fmt.Errorf("%s has no versions yet", rec.ID)
	}
	var oldRec, newRec *record.Record
	var oldLabel, newLabel string
	if len(rest) == 2 {
		snapA, err := r.Store.Snapshot(rec.ID, vA)
		if err != nil {
			return err
		}
		snapB, err := r.Store.Snapshot(rec.ID, vB)
		if err != nil {
			return err
		}
		oldRec, newRec = snapA.Content, snapB.Content
		oldLabel = fmt.Sprintf("v%d", vA)
		newLabel = fmt.Sprintf("v%d", vB)
	} else {
		latestVer := vs[len(vs)-1]
		latest, err := r.Store.Snapshot(rec.ID, latestVer)
		if err != nil {
			return err
		}
		hash, _ := rec.ContentHash(r.Store.Specs())
		if hash != latest.ContentHash {
			oldRec = latest.Content
			newRec = rec
			oldLabel = fmt.Sprintf("v%d", latestVer)
			newLabel = "working copy"
		} else {
			if len(vs) < 2 {
				fmt.Fprintf(out, "\n%s has only version 1; nothing to compare.\n\n", rec.ID)
				return nil
			}
			prevVer := vs[len(vs)-2]
			prev, err := r.Store.Snapshot(rec.ID, prevVer)
			if err != nil {
				return err
			}
			oldRec, newRec = prev.Content, latest.Content
			oldLabel = fmt.Sprintf("v%d", prevVer)
			newLabel = fmt.Sprintf("v%d", latestVer)
		}
	}
	fmt.Fprintf(out, "\n%s  %s → %s\n\n", rec.ID, oldLabel, newLabel)
	changed := diffScalar(out, "title", oldRec.Title, newRec.Title)
	changed = diffScalar(out, "intent", oldRec.Intent, newRec.Intent) || changed
	changed = diffList(out, "constraints", oldRec.Constraints, newRec.Constraints) || changed
	changed = diffList(out, "affected_scope", oldRec.AffectedScope, newRec.AffectedScope) || changed
	changed = diffList(out, "forbidden_scope", oldRec.ForbiddenScope, newRec.ForbiddenScope) || changed
	changed = diffList(out, "associated_specs", specPaths(oldRec.AssociatedSpecs), specPaths(newRec.AssociatedSpecs)) || changed
	if !changed {
		fmt.Fprint(out, "  no differences\n")
	}
	fmt.Fprintln(out)
	return nil
}

// /sedum:setFunction
