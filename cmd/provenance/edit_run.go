package main

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-57b8a95b","gen":"sha256:ecce14df038164301d4c4728b274460b7a56a8972e302aacf33c69ff352be339","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/edit_run.go","path":"errors"}}
import "errors"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-57b8a95b","gen":"sha256:ecce14df038164301d4c4728b274460b7a56a8972e302aacf33c69ff352be339","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/edit_run.go","path":"fmt"}}
import "fmt"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-57b8a95b","gen":"sha256:ecce14df038164301d4c4728b274460b7a56a8972e302aacf33c69ff352be339","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/edit_run.go","path":"sort"}}
import "sort"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-57b8a95b","gen":"sha256:ecce14df038164301d4c4728b274460b7a56a8972e302aacf33c69ff352be339","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/edit_run.go","path":"strconv"}}
import "strconv"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-57b8a95b","gen":"sha256:ecce14df038164301d4c4728b274460b7a56a8972e302aacf33c69ff352be339","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/edit_run.go","path":"github.com/spf13/cobra"}}
import "github.com/spf13/cobra"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-57b8a95b","gen":"sha256:ecce14df038164301d4c4728b274460b7a56a8972e302aacf33c69ff352be339","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/edit_run.go","path":"io"}}
import "io"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-57b8a95b","gen":"sha256:ecce14df038164301d4c4728b274460b7a56a8972e302aacf33c69ff352be339","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/edit_run.go","path":"strings"}}
import "strings"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-57b8a95b","gen":"sha256:ecce14df038164301d4c4728b274460b7a56a8972e302aacf33c69ff352be339","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/edit_run.go","path":"github.com/livecodelife/provenance/record"}}
import "github.com/livecodelife/provenance/record"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-57b8a95b","gen":"sha256:ecce14df038164301d4c4728b274460b7a56a8972e302aacf33c69ff352be339","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/edit_run.go","path":"os"}}
import "os"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-57b8a95b","gen":"sha256:ecce14df038164301d4c4728b274460b7a56a8972e302aacf33c69ff352be339","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/edit_run.go","path":"path/filepath"}}
import "path/filepath"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"tier":"owned","record":"prov-2026-57b8a95b","gen":"sha256:ecce14df038164301d4c4728b274460b7a56a8972e302aacf33c69ff352be339","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func removeValues(list, values []string, field string) ([]string, error) {\n\tout := append([]string{}, list...)\n\tfor _, value := range values {\n\t\tidx := -1\n\t\tfor i, item := range out {\n\t\t\tif item == value {\n\t\t\t\tidx = i\n\t\t\t\tbreak\n\t\t\t}\n\t\t}\n\t\tif idx < 0 {\n\t\t\treturn nil, fmt.Errorf(\"%s has no entry %q\", field, value)\n\t\t}\n\t\tout = append(out[:idx], out[idx+1:]...)\n\t}\n\treturn out, nil\n}","file":"cmd/provenance/edit_run.go","name":"removeValues"}}
func removeValues(list, values []string, field string) ([]string, error) {
	out := append([]string{}, list...)
	for _, value := range values {
		idx := -1
		for i, item := range out {
			if item == value {
				idx = i
				break
			}
		}
		if idx < 0 {
			return nil, fmt.Errorf("%s has no entry %q", field, value)
		}
		out = append(out[:idx], out[idx+1:]...)
	}
	return out, nil
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-57b8a95b","gen":"sha256:ecce14df038164301d4c4728b274460b7a56a8972e302aacf33c69ff352be339","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func removeIndexes(list []string, indexes []string) ([]string, error) {\n\tseen := map[int]bool{}\n\tvar positions []int\n\tfor _, raw := range indexes {\n\t\tn, err := strconv.Atoi(raw)\n\t\tif err != nil || n < 1 || n > len(list) {\n\t\t\treturn nil, fmt.Errorf(\"--remove-constraint %s: no constraint number %s (there are %d)\", raw, raw, len(list))\n\t\t}\n\t\tif !seen[n] {\n\t\t\tseen[n] = true\n\t\t\tpositions = append(positions, n)\n\t\t}\n\t}\n\tsort.Sort(sort.Reverse(sort.IntSlice(positions)))\n\tout := append([]string{}, list...)\n\tfor _, n := range positions {\n\t\tout = append(out[:n-1], out[n:]...)\n\t}\n\treturn out, nil\n}","file":"cmd/provenance/edit_run.go","name":"removeIndexes"}}
func removeIndexes(list []string, indexes []string) ([]string, error) {
	seen := map[int]bool{}
	var positions []int
	for _, raw := range indexes {
		n, err := strconv.Atoi(raw)
		if err != nil || n < 1 || n > len(list) {
			return nil, fmt.Errorf("--remove-constraint %s: no constraint number %s (there are %d)", raw, raw, len(list))
		}
		if !seen[n] {
			seen[n] = true
			positions = append(positions, n)
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(positions)))
	out := append([]string{}, list...)
	for _, n := range positions {
		out = append(out[:n-1], out[n:]...)
	}
	return out, nil
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-57b8a95b","gen":"sha256:ecce14df038164301d4c4728b274460b7a56a8972e302aacf33c69ff352be339","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func runEdit(out io.Writer, cmd *cobra.Command, args []string) error {\n\tr, err := openRepo(cmd)\n\tif err != nil {\n\t\treturn err\n\t}\n\tid, err := recordArg(cmd, args)\n\tif err != nil {\n\t\treturn err\n\t}\n\trec, doc, err := r.resolve(id)\n\tif err != nil {\n\t\treturn err\n\t}\n\tif err := versionable(r, rec); err != nil {\n\t\treturn err\n\t}\n\tflags := cmd.Flags()\n\tif !flags.Changed(\"title\") && !flags.Changed(\"intent\") && !flags.Changed(\"add-constraint\") && !flags.Changed(\"remove-constraint\") && !flags.Changed(\"add-scope\") && !flags.Changed(\"remove-scope\") && !flags.Changed(\"add-forbidden\") && !flags.Changed(\"remove-forbidden\") && !flags.Changed(\"refresh-specs\") && !flags.Changed(\"set-spec\") && !flags.Changed(\"add-spec\") && !flags.Changed(\"remove-spec\") {\n\t\treturn errors.New(\"nothing to edit: pass a field flag such as --intent or --add-constraint\")\n\t}\n\tif flags.Changed(\"title\") {\n\t\ttitle, _ := flags.GetString(\"title\")\n\t\trec.Title = title\n\t}\n\tif flags.Changed(\"intent\") {\n\t\tintent, _ := flags.GetString(\"intent\")\n\t\trec.Intent = intent\n\t}\n\tremoveC, _ := flags.GetStringArray(\"remove-constraint\")\n\taddC, _ := flags.GetStringArray(\"add-constraint\")\n\tconstraints, err := removeIndexes(rec.Constraints, removeC)\n\tif err != nil {\n\t\treturn err\n\t}\n\trec.Constraints = append(constraints, addC...)\n\tremoveS, _ := flags.GetStringArray(\"remove-scope\")\n\taddS, _ := flags.GetStringArray(\"add-scope\")\n\tscope, err := removeValues(rec.AffectedScope, removeS, \"affected_scope\")\n\tif err != nil {\n\t\treturn err\n\t}\n\trec.AffectedScope = append(scope, addS...)\n\tremoveF, _ := flags.GetStringArray(\"remove-forbidden\")\n\taddF, _ := flags.GetStringArray(\"add-forbidden\")\n\tforbidden, err := removeValues(rec.ForbiddenScope, removeF, \"forbidden_scope\")\n\tif err != nil {\n\t\treturn err\n\t}\n\trec.ForbiddenScope = append(forbidden, addF...)\n\tsetSpecs, _ := flags.GetStringArray(\"set-spec\")\n\tspecs, err := applySpecEdits(rec.AssociatedSpecs, setSpecs)\n\tif err != nil {\n\t\treturn err\n\t}\n\taddSpecs, _ := flags.GetStringArray(\"add-spec\")\n\tremoveSpecs, _ := flags.GetStringArray(\"remove-spec\")\n\tspecs, err = applySpecChanges(r, specs, addSpecs, removeSpecs)\n\tif err != nil {\n\t\treturn err\n\t}\n\trec.AssociatedSpecs = specs\n\treason, _ := flags.GetString(\"reason\")\n\tfix, _ := flags.GetBool(\"fix\")\n\treturn saveVersion(out, r, rec, doc, reason, fix)\n}","file":"cmd/provenance/edit_run.go","name":"runEdit"}}
func runEdit(out io.Writer, cmd *cobra.Command, args []string) error {
	r, err := openRepo(cmd)
	if err != nil {
		return err
	}
	id, err := recordArg(cmd, args)
	if err != nil {
		return err
	}
	rec, doc, err := r.resolve(id)
	if err != nil {
		return err
	}
	if err := versionable(r, rec); err != nil {
		return err
	}
	flags := cmd.Flags()
	if !flags.Changed("title") && !flags.Changed("intent") && !flags.Changed("add-constraint") && !flags.Changed("remove-constraint") && !flags.Changed("add-scope") && !flags.Changed("remove-scope") && !flags.Changed("add-forbidden") && !flags.Changed("remove-forbidden") && !flags.Changed("refresh-specs") && !flags.Changed("set-spec") && !flags.Changed("add-spec") && !flags.Changed("remove-spec") {
		return errors.New("nothing to edit: pass a field flag such as --intent or --add-constraint")
	}
	if flags.Changed("title") {
		title, _ := flags.GetString("title")
		rec.Title = title
	}
	if flags.Changed("intent") {
		intent, _ := flags.GetString("intent")
		rec.Intent = intent
	}
	removeC, _ := flags.GetStringArray("remove-constraint")
	addC, _ := flags.GetStringArray("add-constraint")
	constraints, err := removeIndexes(rec.Constraints, removeC)
	if err != nil {
		return err
	}
	rec.Constraints = append(constraints, addC...)
	removeS, _ := flags.GetStringArray("remove-scope")
	addS, _ := flags.GetStringArray("add-scope")
	scope, err := removeValues(rec.AffectedScope, removeS, "affected_scope")
	if err != nil {
		return err
	}
	rec.AffectedScope = append(scope, addS...)
	removeF, _ := flags.GetStringArray("remove-forbidden")
	addF, _ := flags.GetStringArray("add-forbidden")
	forbidden, err := removeValues(rec.ForbiddenScope, removeF, "forbidden_scope")
	if err != nil {
		return err
	}
	rec.ForbiddenScope = append(forbidden, addF...)
	setSpecs, _ := flags.GetStringArray("set-spec")
	specs, err := applySpecEdits(rec.AssociatedSpecs, setSpecs)
	if err != nil {
		return err
	}
	addSpecs, _ := flags.GetStringArray("add-spec")
	removeSpecs, _ := flags.GetStringArray("remove-spec")
	specs, err = applySpecChanges(r, specs, addSpecs, removeSpecs)
	if err != nil {
		return err
	}
	rec.AssociatedSpecs = specs
	reason, _ := flags.GetString("reason")
	fix, _ := flags.GetBool("fix")
	return saveVersion(out, r, rec, doc, reason, fix)
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-57b8a95b","gen":"sha256:ecce14df038164301d4c4728b274460b7a56a8972e302aacf33c69ff352be339","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func applySpecEdits(specs []record.AssociatedSpec, edits []string) ([]record.AssociatedSpec, error) {\n\tout := append([]record.AssociatedSpec(nil), specs...)\n\tfor _, e := range edits {\n\t\told, rest, ok := strings.Cut(e, \"=\")\n\t\tif !ok || old == \"\" || rest == \"\" {\n\t\t\treturn nil, fmt.Errorf(\"--set-spec %q: want <old>=<new> or <old>=<new>:<run_command>\", e)\n\t\t}\n\t\tnewPath, runCmd, hasCmd := strings.Cut(rest, \":\")\n\t\tat := -1\n\t\tfor i := range out {\n\t\t\tif out[i].Path == old {\n\t\t\t\tat = i\n\t\t\t\tbreak\n\t\t\t}\n\t\t}\n\t\tif at < 0 {\n\t\t\tvar have []string\n\t\t\tfor _, s := range out {\n\t\t\t\thave = append(have, s.Path)\n\t\t\t}\n\t\t\treturn nil, fmt.Errorf(\"--set-spec: no associated spec %q (has: %s)\", old, strings.Join(have, \", \"))\n\t\t}\n\t\tout[at].Path = newPath\n\t\tif hasCmd {\n\t\t\tout[at].RunCommand = runCmd\n\t\t}\n\t}\n\treturn out, nil\n}","file":"cmd/provenance/edit_run.go","name":"applySpecEdits"}}
func applySpecEdits(specs []record.AssociatedSpec, edits []string) ([]record.AssociatedSpec, error) {
	out := append([]record.AssociatedSpec(nil), specs...)
	for _, e := range edits {
		old, rest, ok := strings.Cut(e, "=")
		if !ok || old == "" || rest == "" {
			return nil, fmt.Errorf("--set-spec %q: want <old>=<new> or <old>=<new>:<run_command>", e)
		}
		newPath, runCmd, hasCmd := strings.Cut(rest, ":")
		at := -1
		for i := range out {
			if out[i].Path == old {
				at = i
				break
			}
		}
		if at < 0 {
			var have []string
			for _, s := range out {
				have = append(have, s.Path)
			}
			return nil, fmt.Errorf("--set-spec: no associated spec %q (has: %s)", old, strings.Join(have, ", "))
		}
		out[at].Path = newPath
		if hasCmd {
			out[at].RunCommand = runCmd
		}
	}
	return out, nil
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-57b8a95b","gen":"sha256:ecce14df038164301d4c4728b274460b7a56a8972e302aacf33c69ff352be339","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func applySpecChanges(r *repo, specs []record.AssociatedSpec, adds, removes []string) ([]record.AssociatedSpec, error) {\n\tout := append([]record.AssociatedSpec(nil), specs...)\n\tfor _, rm := range removes {\n\t\tat := -1\n\t\tfor i := range out {\n\t\t\tif out[i].Path == rm {\n\t\t\t\tat = i\n\t\t\t\tbreak\n\t\t\t}\n\t\t}\n\t\tif at < 0 {\n\t\t\tvar have []string\n\t\t\tfor _, s := range out {\n\t\t\t\thave = append(have, s.Path)\n\t\t\t}\n\t\t\treturn nil, fmt.Errorf(\"--remove-spec: no associated spec %q (has: %s)\", rm, strings.Join(have, \", \"))\n\t\t}\n\t\tout = append(out[:at], out[at+1:]...)\n\t}\n\tfor _, add := range adds {\n\t\tpath, runCmd, hasCmd := strings.Cut(add, \":\")\n\t\tif path == \"\" {\n\t\t\treturn nil, fmt.Errorf(\"--add-spec %q: want path or path:run_command\", add)\n\t\t}\n\t\tfor _, s := range out {\n\t\t\tif s.Path == path {\n\t\t\t\treturn nil, fmt.Errorf(\"--add-spec: %q is already an associated spec\", path)\n\t\t\t}\n\t\t}\n\t\tif _, err := os.Stat(filepath.Join(r.Root, path)); err != nil {\n\t\t\treturn nil, fmt.Errorf(\"--add-spec: %q does not exist: create the proof file first\", path)\n\t\t}\n\t\tspec := record.AssociatedSpec{Path: path}\n\t\tif hasCmd {\n\t\t\tspec.RunCommand = runCmd\n\t\t}\n\t\tout = append(out, spec)\n\t}\n\treturn out, nil\n}","file":"cmd/provenance/edit_run.go","name":"applySpecChanges"}}
func applySpecChanges(r *repo, specs []record.AssociatedSpec, adds, removes []string) ([]record.AssociatedSpec, error) {
	out := append([]record.AssociatedSpec(nil), specs...)
	for _, rm := range removes {
		at := -1
		for i := range out {
			if out[i].Path == rm {
				at = i
				break
			}
		}
		if at < 0 {
			var have []string
			for _, s := range out {
				have = append(have, s.Path)
			}
			return nil, fmt.Errorf("--remove-spec: no associated spec %q (has: %s)", rm, strings.Join(have, ", "))
		}
		out = append(out[:at], out[at+1:]...)
	}
	for _, add := range adds {
		path, runCmd, hasCmd := strings.Cut(add, ":")
		if path == "" {
			return nil, fmt.Errorf("--add-spec %q: want path or path:run_command", add)
		}
		for _, s := range out {
			if s.Path == path {
				return nil, fmt.Errorf("--add-spec: %q is already an associated spec", path)
			}
		}
		if _, err := os.Stat(filepath.Join(r.Root, path)); err != nil {
			return nil, fmt.Errorf("--add-spec: %q does not exist: create the proof file first", path)
		}
		spec := record.AssociatedSpec{Path: path}
		if hasCmd {
			spec.RunCommand = runCmd
		}
		out = append(out, spec)
	}
	return out, nil
}

// /sedum:setFunction
