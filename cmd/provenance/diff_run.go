package main

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-bfd55535","gen":"sha256:98fb8e2f8fd67f02e2cabc1153d5dfb84f1f27f5c19d0582f9236cd5fa8e4f10","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/diff_run.go","path":"errors"}}
import "errors"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-bfd55535","gen":"sha256:98fb8e2f8fd67f02e2cabc1153d5dfb84f1f27f5c19d0582f9236cd5fa8e4f10","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/diff_run.go","path":"fmt"}}
import "fmt"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-bfd55535","gen":"sha256:98fb8e2f8fd67f02e2cabc1153d5dfb84f1f27f5c19d0582f9236cd5fa8e4f10","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/diff_run.go","path":"io"}}
import "io"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-bfd55535","gen":"sha256:98fb8e2f8fd67f02e2cabc1153d5dfb84f1f27f5c19d0582f9236cd5fa8e4f10","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/diff_run.go","path":"strconv"}}
import "strconv"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-bfd55535","gen":"sha256:98fb8e2f8fd67f02e2cabc1153d5dfb84f1f27f5c19d0582f9236cd5fa8e4f10","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/diff_run.go","path":"github.com/spf13/cobra"}}
import "github.com/spf13/cobra"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-bfd55535","gen":"sha256:98fb8e2f8fd67f02e2cabc1153d5dfb84f1f27f5c19d0582f9236cd5fa8e4f10","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/diff_run.go","path":"github.com/livecodelife/provenance/record"}}
import "github.com/livecodelife/provenance/record"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"tier":"owned","record":"prov-2026-bfd55535","gen":"sha256:98fb8e2f8fd67f02e2cabc1153d5dfb84f1f27f5c19d0582f9236cd5fa8e4f10","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func diffScalar(out io.Writer, field, a, b string) bool {\n\tif a == b {\n\t\treturn false\n\t}\n\tfmt.Fprintf(out, \"  %s:\\n    - %s\\n    + %s\\n\", field, a, b)\n\treturn true\n}","file":"cmd/provenance/diff_run.go","name":"diffScalar"}}
func diffScalar(out io.Writer, field, a, b string) bool {
	if a == b {
		return false
	}
	fmt.Fprintf(out, "  %s:\n    - %s\n    + %s\n", field, a, b)
	return true
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-bfd55535","gen":"sha256:98fb8e2f8fd67f02e2cabc1153d5dfb84f1f27f5c19d0582f9236cd5fa8e4f10","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func diffList(out io.Writer, field string, a, b []string) bool {\n\tinA := make(map[string]bool, len(a))\n\tfor _, x := range a {\n\t\tinA[x] = true\n\t}\n\tinB := make(map[string]bool, len(b))\n\tfor _, x := range b {\n\t\tinB[x] = true\n\t}\n\tvar removed, added []string\n\tfor _, x := range a {\n\t\tif !inB[x] {\n\t\t\tremoved = append(removed, x)\n\t\t}\n\t}\n\tfor _, x := range b {\n\t\tif !inA[x] {\n\t\t\tadded = append(added, x)\n\t\t}\n\t}\n\tif len(removed) == 0 && len(added) == 0 {\n\t\treturn false\n\t}\n\tfmt.Fprintf(out, \"  %s:\\n\", field)\n\tfor _, x := range removed {\n\t\tfmt.Fprintf(out, \"    - %s\\n\", x)\n\t}\n\tfor _, x := range added {\n\t\tfmt.Fprintf(out, \"    + %s\\n\", x)\n\t}\n\treturn true\n}","file":"cmd/provenance/diff_run.go","name":"diffList"}}
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
// sedum:setFunction {"tier":"owned","record":"prov-2026-bfd55535","gen":"sha256:98fb8e2f8fd67f02e2cabc1153d5dfb84f1f27f5c19d0582f9236cd5fa8e4f10","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func specPaths(specs []record.AssociatedSpec) []string {\n\tpaths := make([]string, 0, len(specs))\n\tfor _, s := range specs {\n\t\tpaths = append(paths, s.Path)\n\t}\n\treturn paths\n}","file":"cmd/provenance/diff_run.go","name":"specPaths"}}
func specPaths(specs []record.AssociatedSpec) []string {
	paths := make([]string, 0, len(specs))
	for _, s := range specs {
		paths = append(paths, s.Path)
	}
	return paths
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-bfd55535","gen":"sha256:98fb8e2f8fd67f02e2cabc1153d5dfb84f1f27f5c19d0582f9236cd5fa8e4f10","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func runDiff(out io.Writer, cmd *cobra.Command, args []string) error {\n\tr, err := openRepo(cmd)\n\tif err != nil {\n\t\treturn err\n\t}\n\tflag, err := cmd.Flags().GetString(\"record\")\n\tif err != nil {\n\t\treturn err\n\t}\n\tvar id string\n\tvar rest []string\n\tif flag != \"\" {\n\t\tid = flag\n\t\trest = args\n\t} else {\n\t\tif len(args) == 0 {\n\t\t\treturn errors.New(\"a record id is required (as the first argument or --record)\")\n\t\t}\n\t\tid = args[0]\n\t\trest = args[1:]\n\t}\n\tif len(rest) != 0 && len(rest) != 2 {\n\t\treturn errors.New(\"give two versions to compare, or none\")\n\t}\n\tvar vA, vB int\n\tif len(rest) == 2 {\n\t\tvA, err = strconv.Atoi(rest[0])\n\t\tif err != nil {\n\t\t\treturn fmt.Errorf(\"version %q is not a number\", rest[0])\n\t\t}\n\t\tvB, err = strconv.Atoi(rest[1])\n\t\tif err != nil {\n\t\t\treturn fmt.Errorf(\"version %q is not a number\", rest[1])\n\t\t}\n\t}\n\trec, _, err := r.resolve(id)\n\tif err != nil {\n\t\treturn err\n\t}\n\tvs, err := r.Store.Versions(rec.ID)\n\tif err != nil {\n\t\treturn err\n\t}\n\tif len(vs) == 0 {\n\t\treturn fmt.Errorf(\"%s has no versions yet\", rec.ID)\n\t}\n\tvar oldRec, newRec *record.Record\n\tvar oldLabel, newLabel string\n\tif len(rest) == 2 {\n\t\tsnapA, err := r.Store.Snapshot(rec.ID, vA)\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n\t\tsnapB, err := r.Store.Snapshot(rec.ID, vB)\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n\t\toldRec, newRec = snapA.Content, snapB.Content\n\t\toldLabel = fmt.Sprintf(\"v%d\", vA)\n\t\tnewLabel = fmt.Sprintf(\"v%d\", vB)\n\t} else {\n\t\tlatestVer := vs[len(vs)-1]\n\t\tlatest, err := r.Store.Snapshot(rec.ID, latestVer)\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n\t\thash, _ := rec.ContentHash(r.Store.Specs())\n\t\tif hash != latest.ContentHash {\n\t\t\toldRec = latest.Content\n\t\t\tnewRec = rec\n\t\t\toldLabel = fmt.Sprintf(\"v%d\", latestVer)\n\t\t\tnewLabel = \"working copy\"\n\t\t} else {\n\t\t\tif len(vs) < 2 {\n\t\t\t\tfmt.Fprintf(out, \"\\n%s has only version 1; nothing to compare.\\n\\n\", rec.ID)\n\t\t\t\treturn nil\n\t\t\t}\n\t\t\tprevVer := vs[len(vs)-2]\n\t\t\tprev, err := r.Store.Snapshot(rec.ID, prevVer)\n\t\t\tif err != nil {\n\t\t\t\treturn err\n\t\t\t}\n\t\t\toldRec, newRec = prev.Content, latest.Content\n\t\t\toldLabel = fmt.Sprintf(\"v%d\", prevVer)\n\t\t\tnewLabel = fmt.Sprintf(\"v%d\", latestVer)\n\t\t}\n\t}\n\tfmt.Fprintf(out, \"\\n%s  %s → %s\\n\\n\", rec.ID, oldLabel, newLabel)\n\tchanged := diffScalar(out, \"title\", oldRec.Title, newRec.Title)\n\tchanged = diffScalar(out, \"intent\", oldRec.Intent, newRec.Intent) || changed\n\tchanged = diffList(out, \"constraints\", oldRec.Constraints, newRec.Constraints) || changed\n\tchanged = diffList(out, \"affected_scope\", oldRec.AffectedScope, newRec.AffectedScope) || changed\n\tchanged = diffList(out, \"forbidden_scope\", oldRec.ForbiddenScope, newRec.ForbiddenScope) || changed\n\tchanged = diffList(out, \"associated_specs\", specPaths(oldRec.AssociatedSpecs), specPaths(newRec.AssociatedSpecs)) || changed\n\tif !changed {\n\t\tfmt.Fprint(out, \"  no differences\\n\")\n\t}\n\tfmt.Fprintln(out)\n\treturn nil\n}","file":"cmd/provenance/diff_run.go","name":"runDiff"}}
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
