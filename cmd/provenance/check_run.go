package main

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-28d0e0c4","gen":"sha256:542fd262b3a00ce98cd1a43195b06bc567f4e3321ab4bf58c5a46d21ac272b5e","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/check_run.go","path":"os"}}
import "os"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-7e91a9e6","gen":"sha256:7f05f4f908c228baecf7a239495fa24a2dafaccb50d0ef44ac12ec80b3f332b8","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/check_run.go","path":"github.com/spf13/cobra"}}
import "github.com/spf13/cobra"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-7e91a9e6","gen":"sha256:7f05f4f908c228baecf7a239495fa24a2dafaccb50d0ef44ac12ec80b3f332b8","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/check_run.go","path":"errors"}}
import "errors"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-7e91a9e6","gen":"sha256:7f05f4f908c228baecf7a239495fa24a2dafaccb50d0ef44ac12ec80b3f332b8","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/check_run.go","path":"fmt"}}
import "fmt"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-7e91a9e6","gen":"sha256:7f05f4f908c228baecf7a239495fa24a2dafaccb50d0ef44ac12ec80b3f332b8","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/check_run.go","path":"strings"}}
import "strings"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-8a4ab260","gen":"sha256:f6076b76826a874919295cb20d7498bdeb9676aa6f7bed91764ba1b518fd65b0","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/check_run.go","path":"github.com/livecodelife/provenance/lint"},"kwargs":"sha256:40fd44ebf11e33ee82f78ac85e0c08b4734c16aa57af4776a8728d7ca1875369","body":"sha256:9abff5a2d82dd8a4177704699227186fed48d7e0d093d2d2d5c565b09ff787d7"}
import "github.com/livecodelife/provenance/lint"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-8a4ab260","gen":"sha256:f6076b76826a874919295cb20d7498bdeb9676aa6f7bed91764ba1b518fd65b0","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/check_run.go","path":"github.com/livecodelife/provenance/scope"},"kwargs":"sha256:3e6627102947a22f3bfb35e2e327e2d1b27a3a12dc196a29243b0a535683506d","body":"sha256:4ec0116f2b0d829a4ed0e1f2a40f17d465f6f981c35c8fd5427721c94a450e53"}
import "github.com/livecodelife/provenance/scope"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-95bdf404","gen":"sha256:8d6145bb0d0ef8ad84822d04701b947361cfb553f59d3482afab08e9250748ed","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/check_run.go","path":"io"}}
import "io"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-aa14b7e3","gen":"sha256:41c53627e33bfcdc825412036fee77bab5731d6cebd63c913144b21bdd5ebca4","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/check_run.go","path":"github.com/livecodelife/provenance/record"}}
import "github.com/livecodelife/provenance/record"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-da957222","gen":"sha256:a93432c5bd7b5a061c922dee9bfaf239d5322893e14c789bcd3d69137db71e63","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/check_run.go","path":"os/exec"}}
import "os/exec"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"tier":"owned","record":"prov-2026-15cd7313","gen":"sha256:abf80be8de5d49481529dcbf5413d23ae0097d4d9bcfa92b5aa10ec51388f395","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func headHas(r *repo, base, path string) bool {\n\t_, err := r.git(\"cat-file\", \"-e\", base+\":./\"+path)\n\treturn err == nil\n}","file":"cmd/provenance/check_run.go","name":"headHas"}}
func headHas(r *repo, base, path string) bool {
	_, err := r.git("cat-file", "-e", base+":./"+path)
	return err == nil
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-28d0e0c4","gen":"sha256:542fd262b3a00ce98cd1a43195b06bc567f4e3321ab4bf58c5a46d21ac272b5e","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func stagedTree(r *repo) (string, func(), error) {\n\tdir, err := os.MkdirTemp(\"\", \"provenance-staged-\")\n\tif err != nil {\n\t\treturn \"\", nil, err\n\t}\n\tcleanup := func() { os.RemoveAll(dir) }\n\tif _, err := r.git(\"checkout-index\", \"-a\", \"--prefix=\"+dir+\"/\"); err != nil {\n\t\tcleanup()\n\t\treturn \"\", nil, err\n\t}\n\tprefix, err := r.git(\"rev-parse\", \"--show-prefix\")\n\tif err != nil {\n\t\tcleanup()\n\t\treturn \"\", nil, err\n\t}\n\tpath := dir\n\tif prefix != \"\" {\n\t\tpath = dir + string(os.PathSeparator) + strings.ReplaceAll(prefix, \"/\", string(os.PathSeparator))\n\t}\n\tif err := os.MkdirAll(path, 0o755); err != nil {\n\t\tcleanup()\n\t\treturn \"\", nil, err\n\t}\n\treturn path, cleanup, nil\n}","file":"cmd/provenance/check_run.go","name":"stagedTree"}}
func stagedTree(r *repo) (string, func(), error) {
	dir, err := os.MkdirTemp("", "provenance-staged-")
	if err != nil {
		return "", nil, err
	}
	cleanup := func() { os.RemoveAll(dir) }
	if _, err := r.git("checkout-index", "-a", "--prefix="+dir+"/"); err != nil {
		cleanup()
		return "", nil, err
	}
	prefix, err := r.git("rev-parse", "--show-prefix")
	if err != nil {
		cleanup()
		return "", nil, err
	}
	path := dir
	if prefix != "" {
		path = dir + string(os.PathSeparator) + strings.ReplaceAll(prefix, "/", string(os.PathSeparator))
	}
	if err := os.MkdirAll(path, 0o755); err != nil {
		cleanup()
		return "", nil, err
	}
	return path, cleanup, nil
}

// /sedum:setFunction
// sedum:setType {"tier":"owned","record":"prov-2026-32011bf9","gen":"sha256:5d981c25af362d009de7b95154b78d216ba8d1ca1231013af78704e6b0d373cb","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type violation struct {\n\tRecordID string\n\tFile     string\n\tMessage  string\n}","file":"cmd/provenance/check_run.go","name":"violation"}}
type violation struct {
	RecordID string
	File     string
	Message  string
}

// /sedum:setType
// sedum:setFunction {"tier":"owned","record":"prov-2026-363fac5a","gen":"sha256:79d5f1c4c94e3907300f9bd87cfd19a05edb39a56d7c43d0ce4cb8df98fc728e","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func stateExempt(r *repo, path string) bool {\n\tif strings.HasPrefix(path, \".provenance/\") {\n\t\treturn true\n\t}\n\tdir := strings.ReplaceAll(r.Config.Dir, \"\\\\\", \"/\")\n\treturn strings.HasPrefix(path, dir+\"/\") && recordFileID(r, path) == \"\"\n}","file":"cmd/provenance/check_run.go","name":"stateExempt"}}
func stateExempt(r *repo, path string) bool {
	if strings.HasPrefix(path, ".provenance/") {
		return true
	}
	dir := strings.ReplaceAll(r.Config.Dir, "\\", "/")
	return strings.HasPrefix(path, dir+"/") && recordFileID(r, path) == ""
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-3b156d2b","gen":"sha256:4a8ac05a2bf5cc249fdf39f3ea55bab9688879778ced8bf9ab2e7731d0104f98","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func snapshotOwner(r *repo, path string) string {\n\tdir := []byte(r.Config.Dir)\n\tfor i, c := range dir {\n\t\tif c == '\\\\' {\n\t\t\tdir[i] = '/'\n\t\t}\n\t}\n\tprefix := string(dir) + \"/.versions/\"\n\tif len(path) < len(prefix) || path[:len(prefix)] != prefix {\n\t\treturn \"\"\n\t}\n\trest := path[len(prefix):]\n\tfor i := 0; i < len(rest); i++ {\n\t\tif rest[i] == '/' {\n\t\t\tif i == 0 {\n\t\t\t\treturn \"\"\n\t\t\t}\n\t\t\treturn rest[:i]\n\t\t}\n\t}\n\treturn \"\"\n}","file":"cmd/provenance/check_run.go","name":"snapshotOwner"}}
func snapshotOwner(r *repo, path string) string {
	dir := []byte(r.Config.Dir)
	for i, c := range dir {
		if c == '\\' {
			dir[i] = '/'
		}
	}
	prefix := string(dir) + "/.versions/"
	if len(path) < len(prefix) || path[:len(prefix)] != prefix {
		return ""
	}
	rest := path[len(prefix):]
	for i := 0; i < len(rest); i++ {
		if rest[i] == '/' {
			if i == 0 {
				return ""
			}
			return rest[:i]
		}
	}
	return ""
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-3b34988b","gen":"sha256:b7facd033ca12e01abb89a1a6b1e2b4a7b1c20023a9e60a7a1ffa9797d0d1a12","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func deletedFromPreviousScope(r *repo, store *record.FilesStore, base, tree string, rec *record.Record, path string) bool {\n\tif rec.Version < 2 {\n\t\treturn false\n\t}\n\tfull := tree + string(os.PathSeparator) + strings.ReplaceAll(path, \"/\", string(os.PathSeparator))\n\tif _, err := os.Stat(full); err == nil || !os.IsNotExist(err) {\n\t\treturn false\n\t}\n\tif !headHas(r, base, path) {\n\t\treturn false\n\t}\n\tsnap, err := store.Snapshot(rec.ID, rec.Version-1)\n\tif err != nil || snap.Content == nil {\n\t\treturn false\n\t}\n\tfor _, pat := range rec.ForbiddenScope {\n\t\tif ok, err := scope.Match(path, pat); err == nil && ok {\n\t\t\treturn false\n\t\t}\n\t}\n\tfor _, pat := range snap.Content.AffectedScope {\n\t\tif ok, err := scope.Match(path, pat); err == nil && ok {\n\t\t\treturn true\n\t\t}\n\t}\n\treturn false\n}","file":"cmd/provenance/check_run.go","name":"deletedFromPreviousScope"}}
func deletedFromPreviousScope(r *repo, store *record.FilesStore, base, tree string, rec *record.Record, path string) bool {
	if rec.Version < 2 {
		return false
	}
	full := tree + string(os.PathSeparator) + strings.ReplaceAll(path, "/", string(os.PathSeparator))
	if _, err := os.Stat(full); err == nil || !os.IsNotExist(err) {
		return false
	}
	if !headHas(r, base, path) {
		return false
	}
	snap, err := store.Snapshot(rec.ID, rec.Version-1)
	if err != nil || snap.Content == nil {
		return false
	}
	for _, pat := range rec.ForbiddenScope {
		if ok, err := scope.Match(path, pat); err == nil && ok {
			return false
		}
	}
	for _, pat := range snap.Content.AffectedScope {
		if ok, err := scope.Match(path, pat); err == nil && ok {
			return true
		}
	}
	return false
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-3c33725b","gen":"sha256:83c43f77c39045cb01ef7272862e2c513841423c0ad8350ae30c4857c5e0681f","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func recordFileID(r *repo, path string) string {\n\tprefix := r.Config.Dir + \"/\"\n\tconst suffix = \".yml\"\n\tif len(path) <= len(prefix)+len(suffix) || path[:len(prefix)] != prefix || path[len(path)-len(suffix):] != suffix {\n\t\treturn \"\"\n\t}\n\tid := path[len(prefix) : len(path)-len(suffix)]\n\tfor i := 0; i < len(id); i++ {\n\t\tif id[i] == '/' {\n\t\t\treturn \"\"\n\t\t}\n\t}\n\treturn id\n}","file":"cmd/provenance/check_run.go","name":"recordFileID"}}
func recordFileID(r *repo, path string) string {
	prefix := r.Config.Dir + "/"
	const suffix = ".yml"
	if len(path) <= len(prefix)+len(suffix) || path[:len(prefix)] != prefix || path[len(path)-len(suffix):] != suffix {
		return ""
	}
	id := path[len(prefix) : len(path)-len(suffix)]
	for i := 0; i < len(id); i++ {
		if id[i] == '/' {
			return ""
		}
	}
	return id
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-5f2b3d00","gen":"sha256:38e8ecb9d7fe5018089836499a52cad0a70d6a0744e6e0d58aeaf9adfd81409d","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func blobOf(r *repo, spec string) string {\n\tout, err := r.git(\"rev-parse\", \"-q\", \"--verify\", spec)\n\tif err != nil {\n\t\treturn \"\"\n\t}\n\treturn out\n}","file":"cmd/provenance/check_run.go","name":"blobOf"}}
func blobOf(r *repo, spec string) string {
	out, err := r.git("rev-parse", "-q", "--verify", spec)
	if err != nil {
		return ""
	}
	return out
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-7e91a9e6","gen":"sha256:7f05f4f908c228baecf7a239495fa24a2dafaccb50d0ef44ac12ec80b3f332b8","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func runCheck(out io.Writer, cmd *cobra.Command, args []string) error {\n\tr, err := openRepo(cmd)\n\tif err != nil {\n\t\treturn err\n\t}\n\n\tstaged, _ := cmd.Flags().GetBool(\"staged\")\n\tcommit, _ := cmd.Flags().GetString(\"commit\")\n\trangeSpec, _ := cmd.Flags().GetString(\"range\")\n\tevidence, _ := cmd.Flags().GetBool(\"evidence\")\n\tformat, _ := cmd.Flags().GetString(\"format\")\n\tmessageFile, _ := cmd.Flags().GetString(\"message-file\")\n\n\tif !staged && commit == \"\" && rangeSpec == \"\" && !evidence {\n\t\treturn errors.New(\"pass --staged, --commit <sha>, --range <a..b> or --evidence\")\n\t}\n\tif format != \"\" && format != \"human\" && format != \"sarif\" {\n\t\treturn fmt.Errorf(\"unknown format %q (human|sarif)\", format)\n\t}\n\tif format == \"sarif\" {\n\t\treturn checkSarif(out, cmd, r)\n\t}\n\tif commit != \"\" || rangeSpec != \"\" {\n\t\treturn runCheckCommits(out, r, commit, rangeSpec)\n\t}\n\tif evidence {\n\t\treturn runCheckEvidence(out, cmd, r)\n\t}\n\n\tvar message string\n\tif messageFile != \"\" {\n\t\tdata, err := os.ReadFile(messageFile)\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n\t\tmessage = strings.TrimSpace(string(data))\n\t} else {\n\t\tmessage, err = r.git(\"log\", \"-1\", \"--format=%B\")\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n\t}\n\n\tviolations, err := checkStaged(r, message)\n\tif err != nil {\n\t\tio.WriteString(out, \"\\n✗ Check failed: \"+err.Error()+\"\\n\\n\")\n\t\treturn err\n\t}\n\tprintViolations(out, violations)\n\tif len(violations) > 0 {\n\t\treturn errors.New(\"forbidden scope violations found\")\n\t}\n\treturn nil\n}","file":"cmd/provenance/check_run.go","name":"runCheck"}}
func runCheck(out io.Writer, cmd *cobra.Command, args []string) error {
	r, err := openRepo(cmd)
	if err != nil {
		return err
	}

	staged, _ := cmd.Flags().GetBool("staged")
	commit, _ := cmd.Flags().GetString("commit")
	rangeSpec, _ := cmd.Flags().GetString("range")
	evidence, _ := cmd.Flags().GetBool("evidence")
	format, _ := cmd.Flags().GetString("format")
	messageFile, _ := cmd.Flags().GetString("message-file")

	if !staged && commit == "" && rangeSpec == "" && !evidence {
		return errors.New("pass --staged, --commit <sha>, --range <a..b> or --evidence")
	}
	if format != "" && format != "human" && format != "sarif" {
		return fmt.Errorf("unknown format %q (human|sarif)", format)
	}
	if format == "sarif" {
		return checkSarif(out, cmd, r)
	}
	if commit != "" || rangeSpec != "" {
		return runCheckCommits(out, r, commit, rangeSpec)
	}
	if evidence {
		return runCheckEvidence(out, cmd, r)
	}

	var message string
	if messageFile != "" {
		data, err := os.ReadFile(messageFile)
		if err != nil {
			return err
		}
		message = strings.TrimSpace(string(data))
	} else {
		message, err = r.git("log", "-1", "--format=%B")
		if err != nil {
			return err
		}
	}

	violations, err := checkStaged(r, message)
	if err != nil {
		io.WriteString(out, "\n✗ Check failed: "+err.Error()+"\n\n")
		return err
	}
	printViolations(out, violations)
	if len(violations) > 0 {
		return errors.New("forbidden scope violations found")
	}
	return nil
}

// /sedum:setFunction
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-8a4ab260","gen":"sha256:f6076b76826a874919295cb20d7498bdeb9676aa6f7bed91764ba1b518fd65b0","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/check_run.go","name":"checkChange"},"kwargs":"sha256:464556dc48d918495d0d62aa7301533fc5887e6bf4af0f7ce0ec9b67ab04b7f6","body":"sha256:8cd15f76f80fef59063f601855569aac554ec1b4df5ea916078777bfe5cda3db"}
func checkChange(r *repo, base, tree string, changed []string, message string) ([]violation, error) {
	store := &record.FilesStore{Root: tree, Dir: r.Config.Dir, SpecRoot: tree}
	shared, _ := loadShared(r)
	enforcement := r.Config.Enforcement
	if enforcement == "" {
		enforcement = "none"
	}

	var out []violation
	for _, p := range changed {
		owner := snapshotOwner(r, p)
		if owner == "" || !headHas(r, base, p) || pinLogAppendOnly(r, base, tree, p) {
			continue
		}
		out = append(out, violation{RecordID: owner, Message: "snapshot files are append-only: " + p + " already exists in " + base})
	}

	for _, p := range changed {
		id := recordFileID(r, p)
		if id == "" {
			continue
		}
		staged, _, err := store.Load(id)
		if err != nil {
			continue
		}
		old, ok := headRecord(r, base, p)
		if !ok || old == nil || old.Schema != 2 || old.Status == staged.Status {
			continue
		}
		if err := lint.CheckTransition(old.Status, staged.Status, staged.Version > old.Version); err != nil {
			out = append(out, violation{RecordID: id, Message: id + ": " + err.Error()})
		}
	}

	res, err := lint.Lint(store, lint.Options{
		Root:         tree,
		Enforcement:  enforcement,
		ExcludePaths: r.Config.ExcludePaths,
		ManifestPath: manifestPath,
		Shared:       shared,
		SharedRepos:  sharedNames(r),
	})
	if err != nil {
		return nil, err
	}

	tags := lint.ParseCommitTags(message)
	tagIDs := map[string]bool{}
	for _, t := range tags {
		tagIDs[t.ID] = true
	}
	changedIDs := map[string]bool{}
	for _, p := range changed {
		if id := recordFileID(r, p); id != "" {
			changedIDs[id] = true
		}
	}
	for _, is := range res.Issues {
		if is.Severity != lint.SeverityError {
			continue
		}
		if is.RecordID != "" && !tagIDs[is.RecordID] && !changedIDs[is.RecordID] {
			continue
		}
		out = append(out, violation{RecordID: is.RecordID, Message: "lint " + is.Rule + ": " + is.Message})
	}

	var nonExcluded []string
	for _, p := range changed {
		if scope.Excluded(p, r.Config.ExcludePaths) || stateExempt(r, p) {
			continue
		}
		nonExcluded = append(nonExcluded, p)
	}

	if len(tags) == 0 {
		if r.Config.CommitTagRequired && len(nonExcluded) > 0 {
			out = append(out, violation{Message: "Commit tag required but no provenance ID found in message"})
		}
		return out, nil
	}
	if len(nonExcluded) == 0 {
		return out, nil
	}

	for _, tag := range tags {
		resolved, err := lint.ResolveTag(store, tag)
		if err != nil {
			if tag.Version == 0 {
				continue
			}
			out = append(out, violation{RecordID: tag.ID, Message: err.Error()})
			continue
		}
		rec, _, err := store.Load(resolved.ID)
		if err != nil {
			continue
		}
		if tag.Version != 0 && tag.Version != rec.Version {
			out = append(out, violation{RecordID: rec.ID, Message: fmt.Sprintf("%s is at version %d in this commit: tag it [%s@%d], not %s", rec.ID, rec.Version, rec.ID, rec.Version, tag.String())})
			continue
		}

		own := r.Config.Dir + "/" + rec.ID + ".yml"
		baseOwn, hasBaseOwn := headRecord(r, base, own)
		var enforce []string
		for _, p := range nonExcluded {
			if p == manifestPath {
				continue
			}
			if owner := snapshotOwner(r, p); owner == rec.ID {
				continue
			}
			if p == own && hasBaseOwn && baseOwn != nil && baseOwn.Status == rec.Status && baseOwn.Version == rec.Version && baseOwn.ImplementsVersion != rec.ImplementsVersion {
				continue
			}
			if id := recordFileID(r, p); id != "" {
				if other, _, err := store.Load(id); err == nil && other.Implements == rec.ID {
					continue
				}
			}
			enforce = append(enforce, p)
		}
		if len(enforce) == 0 {
			continue
		}

		stagesOwn := false
		for _, p := range changed {
			if p == own {
				stagesOwn = true
			}
		}
		completion := stagesOwn && hasBaseOwn && baseOwn != nil && baseOwn.Status != record.StatusImplemented
		if rec.Status == record.StatusImplemented && !completion {
			out = append(out, violation{RecordID: rec.ID, Message: rec.ID + " is implemented: run 'provenance revise " + rec.ID + "' to start a new version, then tag the commit with it."})
			continue
		}

		for _, p := range enforce {
			if p == own {
				for _, pat := range rec.ForbiddenScope {
					matched, err := scope.Match(p, pat)
					if err != nil {
						return nil, err
					}
					if matched {
						out = append(out, violation{RecordID: rec.ID, File: p, Message: rec.ID + " forbids changes to " + p})
						break
					}
				}
				continue
			}
			in, err := scope.InScope(rec, p)
			if err != nil {
				return nil, err
			}
			if in {
				continue
			}
			if deletedFromPreviousScope(r, store, base, tree, rec, p) {
				continue
			}
			out = append(out, violation{RecordID: rec.ID, File: p, Message: rec.ID + " forbids changes to " + p})
		}
	}

	return out, nil
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-95bdf404","gen":"sha256:8d6145bb0d0ef8ad84822d04701b947361cfb553f59d3482afab08e9250748ed","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func printViolations(out io.Writer, violations []violation) {\n\twrite := func(s string) { _, _ = out.Write([]byte(s)) }\n\tif len(violations) == 0 {\n\t\twrite(\"\\n✓ No forbidden scope violations in staged\\n\\n\")\n\t\treturn\n\t}\n\twrite(\"\\n✗ Forbidden scope violation in staged\\n\\n\")\n\tvar order []string\n\tgroups := map[string][]violation{}\n\tfor _, v := range violations {\n\t\tif v.RecordID == \"\" {\n\t\t\twrite(\"  \" + v.Message + \"\\n\")\n\t\t\tcontinue\n\t\t}\n\t\tif _, ok := groups[v.RecordID]; !ok {\n\t\t\torder = append(order, v.RecordID)\n\t\t}\n\t\tgroups[v.RecordID] = append(groups[v.RecordID], v)\n\t}\n\tfor _, id := range order {\n\t\twrite(\"  \" + id + \" forbids changes to:\\n\")\n\t\tfor _, v := range groups[id] {\n\t\t\tif v.File != \"\" {\n\t\t\t\twrite(\"    · \" + v.File + \"\\n\")\n\t\t\t} else {\n\t\t\t\twrite(\"    \" + v.Message + \"\\n\")\n\t\t\t}\n\t\t}\n\t\twrite(\"\\n\")\n\t}\n\twrite(\"  Hint: If this change is intentional, revise the record to widen its scope (provenance edit --add-scope) or create a record that covers it.\\n\\n\")\n}\n","file":"cmd/provenance/check_run.go","name":"printViolations"}}
func printViolations(out io.Writer, violations []violation) {
	write := func(s string) { _, _ = out.Write([]byte(s)) }
	if len(violations) == 0 {
		write("\n✓ No forbidden scope violations in staged\n\n")
		return
	}
	write("\n✗ Forbidden scope violation in staged\n\n")
	var order []string
	groups := map[string][]violation{}
	for _, v := range violations {
		if v.RecordID == "" {
			write("  " + v.Message + "\n")
			continue
		}
		if _, ok := groups[v.RecordID]; !ok {
			order = append(order, v.RecordID)
		}
		groups[v.RecordID] = append(groups[v.RecordID], v)
	}
	for _, id := range order {
		write("  " + id + " forbids changes to:\n")
		for _, v := range groups[id] {
			if v.File != "" {
				write("    · " + v.File + "\n")
			} else {
				write("    " + v.Message + "\n")
			}
		}
		write("\n")
	}
	write("  Hint: If this change is intentional, revise the record to widen its scope (provenance edit --add-scope) or create a record that covers it.\n\n")
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-9ad86fdf","gen":"sha256:c474cc49e6baca3817c91d3f96b6ec26ddc66be203e1bf57bda97310c10b1b22","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func checkStaged(r *repo, message string) ([]violation, error) {\n\tparents := pendingMergeParents(r)\n\tif len(parents) == 2 {\n\t\tif tree, err := r.git(\"write-tree\"); err == nil {\n\t\t\tif mergeIsAutomatic(r, tree, parents) {\n\t\t\t\treturn nil, nil\n\t\t\t}\n\t\t}\n\t}\n\n\tpaths, err := stagedPaths(r)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\n\ttreePath, cleanup, err := stagedTree(r)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\tdefer cleanup()\n\n\treturn checkChange(r, \"HEAD\", treePath, paths, message)\n}","file":"cmd/provenance/check_run.go","name":"checkStaged"}}
func checkStaged(r *repo, message string) ([]violation, error) {
	parents := pendingMergeParents(r)
	if len(parents) == 2 {
		if tree, err := r.git("write-tree"); err == nil {
			if mergeIsAutomatic(r, tree, parents) {
				return nil, nil
			}
		}
	}

	paths, err := stagedPaths(r)
	if err != nil {
		return nil, err
	}

	treePath, cleanup, err := stagedTree(r)
	if err != nil {
		return nil, err
	}
	defer cleanup()

	return checkChange(r, "HEAD", treePath, paths, message)
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-a62e977f","gen":"sha256:5af633d197cbfc51185b3cce06a702684cad207396d6252dd5d6c6525c57cbfa","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func mergeFiltered(r *repo, paths []string) []string {\n\tif _, err := r.git(\"rev-parse\", \"-q\", \"--verify\", \"MERGE_HEAD\"); err != nil {\n\t\treturn paths\n\t}\n\tvar kept []string\n\tfor _, p := range paths {\n\t\tidx := blobOf(r, \":./\"+p)\n\t\thead := blobOf(r, \"HEAD:./\"+p)\n\t\tmergeHead := blobOf(r, \"MERGE_HEAD:./\"+p)\n\t\tif idx != head && idx != mergeHead {\n\t\t\tkept = append(kept, p)\n\t\t}\n\t}\n\treturn kept\n}","file":"cmd/provenance/check_run.go","name":"mergeFiltered"}}
func mergeFiltered(r *repo, paths []string) []string {
	if _, err := r.git("rev-parse", "-q", "--verify", "MERGE_HEAD"); err != nil {
		return paths
	}
	var kept []string
	for _, p := range paths {
		idx := blobOf(r, ":./"+p)
		head := blobOf(r, "HEAD:./"+p)
		mergeHead := blobOf(r, "MERGE_HEAD:./"+p)
		if idx != head && idx != mergeHead {
			kept = append(kept, p)
		}
	}
	return kept
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-aa14b7e3","gen":"sha256:41c53627e33bfcdc825412036fee77bab5731d6cebd63c913144b21bdd5ebca4","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func headRecord(r *repo, base, path string) (*record.Record, bool) {\n\ttext, err := r.git(\"show\", base+\":./\"+path)\n\tif err != nil {\n\t\treturn nil, false\n\t}\n\tdoc, err := record.Parse([]byte(text + \"\\n\"))\n\tif err != nil {\n\t\treturn nil, false\n\t}\n\treturn doc.Record(), true\n}","file":"cmd/provenance/check_run.go","name":"headRecord"}}
func headRecord(r *repo, base, path string) (*record.Record, bool) {
	text, err := r.git("show", base+":./"+path)
	if err != nil {
		return nil, false
	}
	doc, err := record.Parse([]byte(text + "\n"))
	if err != nil {
		return nil, false
	}
	return doc.Record(), true
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-cbd32da5","gen":"sha256:a23dc1074da7f75593b77e9a59e742fda7d74113108f795922b4f72fe124f518","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func stagedPaths(r *repo) ([]string, error) {\n\tout, err := r.git(\"diff\", \"--cached\", \"--name-only\", \"--no-renames\", \"--relative\")\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\tvar paths []string\n\tstart := 0\n\tfor i := 0; i <= len(out); i++ {\n\t\tif i == len(out) || out[i] == '\\n' {\n\t\t\tif line := out[start:i]; line != \"\" {\n\t\t\t\tpaths = append(paths, line)\n\t\t\t}\n\t\t\tstart = i + 1\n\t\t}\n\t}\n\treturn mergeFiltered(r, paths), nil\n}","file":"cmd/provenance/check_run.go","name":"stagedPaths"}}
func stagedPaths(r *repo) ([]string, error) {
	out, err := r.git("diff", "--cached", "--name-only", "--no-renames", "--relative")
	if err != nil {
		return nil, err
	}
	var paths []string
	start := 0
	for i := 0; i <= len(out); i++ {
		if i == len(out) || out[i] == '\n' {
			if line := out[start:i]; line != "" {
				paths = append(paths, line)
			}
			start = i + 1
		}
	}
	return mergeFiltered(r, paths), nil
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-da957222","gen":"sha256:a93432c5bd7b5a061c922dee9bfaf239d5322893e14c789bcd3d69137db71e63","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func pinLogAppendOnly(r *repo, base, tree, path string) bool {\n\towner := snapshotOwner(r, path)\n\tif owner == \"\" {\n\t\treturn false\n\t}\n\tif path != r.Config.Dir+\"/.versions/\"+owner+\"/pins.jsonl\" {\n\t\treturn false\n\t}\n\tcmd := exec.Command(\"git\", \"cat-file\", \"blob\", base+\":./\"+path)\n\tcmd.Dir = r.Root\n\tcmd.Env = r.gitEnv()\n\tbaseOut, err := cmd.Output()\n\tif err != nil {\n\t\treturn false\n\t}\n\ttreeData, err := os.ReadFile(tree + \"/\" + path)\n\tif err != nil {\n\t\treturn false\n\t}\n\treturn strings.HasPrefix(string(treeData), string(baseOut))\n}\n","file":"cmd/provenance/check_run.go","name":"pinLogAppendOnly"}}
func pinLogAppendOnly(r *repo, base, tree, path string) bool {
	owner := snapshotOwner(r, path)
	if owner == "" {
		return false
	}
	if path != r.Config.Dir+"/.versions/"+owner+"/pins.jsonl" {
		return false
	}
	cmd := exec.Command("git", "cat-file", "blob", base+":./"+path)
	cmd.Dir = r.Root
	cmd.Env = r.gitEnv()
	baseOut, err := cmd.Output()
	if err != nil {
		return false
	}
	treeData, err := os.ReadFile(tree + "/" + path)
	if err != nil {
		return false
	}
	return strings.HasPrefix(string(treeData), string(baseOut))
}

// /sedum:setFunction
