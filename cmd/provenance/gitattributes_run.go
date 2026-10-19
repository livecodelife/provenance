package main

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-86a13b1e","gen":"sha256:19b2d09bfe7b2e6e3363fb5af7869bf0ed912e65d9326c5255780ae198758265","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/gitattributes_run.go","path":"os"}}
import "os"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-86a13b1e","gen":"sha256:19b2d09bfe7b2e6e3363fb5af7869bf0ed912e65d9326c5255780ae198758265","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/gitattributes_run.go","path":"path/filepath"}}
import "path/filepath"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-86a13b1e","gen":"sha256:19b2d09bfe7b2e6e3363fb5af7869bf0ed912e65d9326c5255780ae198758265","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/gitattributes_run.go","path":"strings"}}
import "strings"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-86a13b1e","gen":"sha256:19b2d09bfe7b2e6e3363fb5af7869bf0ed912e65d9326c5255780ae198758265","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/gitattributes_run.go","path":"fmt"}}
import "fmt"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-86a13b1e","gen":"sha256:19b2d09bfe7b2e6e3363fb5af7869bf0ed912e65d9326c5255780ae198758265","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/gitattributes_run.go","path":"os/exec"}}
import "os/exec"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setConst {"tier":"owned","record":"prov-2026-86a13b1e","gen":"sha256:19b2d09bfe7b2e6e3363fb5af7869bf0ed912e65d9326c5255780ae198758265","pkg":"sha256:51be1c18998d63cc558c253e211a6bd188d96cb251dad68c6ea88e03c2d5cc4c","kwargs":{"code":"// gitattributesFile is the git attributes file at the repository root.\nconst gitattributesFile = \".gitattributes\"","file":"cmd/provenance/gitattributes_run.go","name":"gitattributesFile"}}
// gitattributesFile is the git attributes file at the repository root.
const gitattributesFile = ".gitattributes"

// /sedum:setConst
// sedum:setFunction {"tier":"owned","record":"prov-2026-86a13b1e","gen":"sha256:19b2d09bfe7b2e6e3363fb5af7869bf0ed912e65d9326c5255780ae198758265","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"// generatedAttrLines are the attribute lines that mark the files provenance\n// writes as generated: snapshots, evidence and the version manifest.\nfunc generatedAttrLines(dir string) []string {\n\tdir = strings.TrimSuffix(filepath.ToSlash(dir), \"/\")\n\treturn []string{\n\t\tdir + \"/.versions/** linguist-generated=true -diff\",\n\t\tdir + \"/.evidence/** linguist-generated=true -diff\",\n\t\tmanifestPath + \" linguist-generated=true -diff\",\n\t}\n}","file":"cmd/provenance/gitattributes_run.go","name":"generatedAttrLines"}}
// generatedAttrLines are the attribute lines that mark the files provenance
// writes as generated: snapshots, evidence and the version manifest.
func generatedAttrLines(dir string) []string {
	dir = strings.TrimSuffix(filepath.ToSlash(dir), "/")
	return []string{
		dir + "/.versions/** linguist-generated=true -diff",
		dir + "/.evidence/** linguist-generated=true -diff",
		manifestPath + " linguist-generated=true -diff",
	}
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-86a13b1e","gen":"sha256:19b2d09bfe7b2e6e3363fb5af7869bf0ed912e65d9326c5255780ae198758265","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"// missingAttrLines returns the generated lines that root's .gitattributes\n// lacks, comparing lines after trimming. A missing file lacks all of them.\nfunc missingAttrLines(root, dir string) ([]string, error) {\n\tdata, err := os.ReadFile(filepath.Join(root, gitattributesFile))\n\tif err != nil && !os.IsNotExist(err) {\n\t\treturn nil, err\n\t}\n\thave := map[string]bool{}\n\tfor _, l := range strings.Split(string(data), \"\\n\") {\n\t\thave[strings.TrimSpace(l)] = true\n\t}\n\tvar missing []string\n\tfor _, l := range generatedAttrLines(dir) {\n\t\tif !have[l] {\n\t\t\tmissing = append(missing, l)\n\t\t}\n\t}\n\treturn missing, nil\n}","file":"cmd/provenance/gitattributes_run.go","name":"missingAttrLines"}}
// missingAttrLines returns the generated lines that root's .gitattributes
// lacks, comparing lines after trimming. A missing file lacks all of them.
func missingAttrLines(root, dir string) ([]string, error) {
	data, err := os.ReadFile(filepath.Join(root, gitattributesFile))
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}
	have := map[string]bool{}
	for _, l := range strings.Split(string(data), "\n") {
		have[strings.TrimSpace(l)] = true
	}
	var missing []string
	for _, l := range generatedAttrLines(dir) {
		if !have[l] {
			missing = append(missing, l)
		}
	}
	return missing, nil
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-86a13b1e","gen":"sha256:19b2d09bfe7b2e6e3363fb5af7869bf0ed912e65d9326c5255780ae198758265","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"// ensureGitattributes appends the missing generated lines to root's\n// .gitattributes and returns \"created\", \"updated\" or \"unchanged\". It never\n// edits, reorders or removes an existing line.\nfunc ensureGitattributes(root, dir string) (string, error) {\n\tmissing, err := missingAttrLines(root, dir)\n\tif err != nil {\n\t\treturn \"\", err\n\t}\n\tif len(missing) == 0 {\n\t\treturn \"unchanged\", nil\n\t}\n\tpath := filepath.Join(root, gitattributesFile)\n\tstatus := \"updated\"\n\tdata, err := os.ReadFile(path)\n\tif os.IsNotExist(err) {\n\t\tstatus = \"created\"\n\t} else if err != nil {\n\t\treturn \"\", err\n\t}\n\ttext := string(data)\n\tif text != \"\" && !strings.HasSuffix(text, \"\\n\") {\n\t\ttext += \"\\n\"\n\t}\n\ttext += strings.Join(missing, \"\\n\") + \"\\n\"\n\tif err := os.WriteFile(path, []byte(text), 0o644); err != nil {\n\t\treturn \"\", err\n\t}\n\treturn status, nil\n}","file":"cmd/provenance/gitattributes_run.go","name":"ensureGitattributes"}}
// ensureGitattributes appends the missing generated lines to root's
// .gitattributes and returns "created", "updated" or "unchanged". It never
// edits, reorders or removes an existing line.
func ensureGitattributes(root, dir string) (string, error) {
	missing, err := missingAttrLines(root, dir)
	if err != nil {
		return "", err
	}
	if len(missing) == 0 {
		return "unchanged", nil
	}
	path := filepath.Join(root, gitattributesFile)
	status := "updated"
	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		status = "created"
	} else if err != nil {
		return "", err
	}
	text := string(data)
	if text != "" && !strings.HasSuffix(text, "\n") {
		text += "\n"
	}
	text += strings.Join(missing, "\n") + "\n"
	if err := os.WriteFile(path, []byte(text), 0o644); err != nil {
		return "", err
	}
	return status, nil
}

// /sedum:setFunction
// sedum:setConst {"tier":"owned","record":"prov-2026-86a13b1e","gen":"sha256:19b2d09bfe7b2e6e3363fb5af7869bf0ed912e65d9326c5255780ae198758265","pkg":"sha256:51be1c18998d63cc558c253e211a6bd188d96cb251dad68c6ea88e03c2d5cc4c","kwargs":{"code":"// manifestMergeAttr routes the version manifest through the provenance merge driver.\nconst manifestMergeAttr = manifestPath + \" merge=provenance\"","file":"cmd/provenance/gitattributes_run.go","name":"manifestMergeAttr"}}
// manifestMergeAttr routes the version manifest through the provenance merge driver.
const manifestMergeAttr = manifestPath + " merge=provenance"

// /sedum:setConst
// sedum:setFunction {"tier":"owned","record":"prov-2026-86a13b1e","gen":"sha256:19b2d09bfe7b2e6e3363fb5af7869bf0ed912e65d9326c5255780ae198758265","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"// ensureMergeDriver makes git merge the version manifest with `provenance\n// merge-manifest`: it appends the merge attribute to root's .gitattributes when\n// it is missing (never editing another line) and sets the merge.provenance\n// driver in the repository's git config. A root that is not a git repository\n// only gets the attribute. It returns \"created\", \"updated\" or \"unchanged\" for\n// .gitattributes.\nfunc ensureMergeDriver(root string) (string, error) {\n\tstatus := \"unchanged\"\n\tpath := filepath.Join(root, gitattributesFile)\n\tdata, err := os.ReadFile(path)\n\tif err != nil && !os.IsNotExist(err) {\n\t\treturn \"\", err\n\t}\n\thave := false\n\tfor _, l := range strings.Split(string(data), \"\\n\") {\n\t\tif strings.TrimSpace(l) == manifestMergeAttr {\n\t\t\thave = true\n\t\t}\n\t}\n\tif !have {\n\t\tstatus = \"updated\"\n\t\tif os.IsNotExist(err) {\n\t\t\tstatus = \"created\"\n\t\t}\n\t\ttext := string(data)\n\t\tif text != \"\" && !strings.HasSuffix(text, \"\\n\") {\n\t\t\ttext += \"\\n\"\n\t\t}\n\t\ttext += manifestMergeAttr + \"\\n\"\n\t\tif err := os.WriteFile(path, []byte(text), 0o644); err != nil {\n\t\t\treturn \"\", err\n\t\t}\n\t}\n\tif err := exec.Command(\"git\", \"-C\", root, \"rev-parse\", \"--git-dir\").Run(); err != nil {\n\t\treturn status, nil\n\t}\n\tfor _, kv := range [][2]string{{\"merge.provenance.name\", \"provenance version manifest\"}, {\"merge.provenance.driver\", \"provenance merge-manifest %O %A %B\"}} {\n\t\tif out, err := exec.Command(\"git\", \"-C\", root, \"config\", kv[0], kv[1]).CombinedOutput(); err != nil {\n\t\t\treturn \"\", fmt.Errorf(\"git config %s: %v: %s\", kv[0], err, out)\n\t\t}\n\t}\n\treturn status, nil\n}","file":"cmd/provenance/gitattributes_run.go","name":"ensureMergeDriver"}}
// ensureMergeDriver makes git merge the version manifest with `provenance
// merge-manifest`: it appends the merge attribute to root's .gitattributes when
// it is missing (never editing another line) and sets the merge.provenance
// driver in the repository's git config. A root that is not a git repository
// only gets the attribute. It returns "created", "updated" or "unchanged" for
// .gitattributes.
func ensureMergeDriver(root string) (string, error) {
	status := "unchanged"
	path := filepath.Join(root, gitattributesFile)
	data, err := os.ReadFile(path)
	if err != nil && !os.IsNotExist(err) {
		return "", err
	}
	have := false
	for _, l := range strings.Split(string(data), "\n") {
		if strings.TrimSpace(l) == manifestMergeAttr {
			have = true
		}
	}
	if !have {
		status = "updated"
		if os.IsNotExist(err) {
			status = "created"
		}
		text := string(data)
		if text != "" && !strings.HasSuffix(text, "\n") {
			text += "\n"
		}
		text += manifestMergeAttr + "\n"
		if err := os.WriteFile(path, []byte(text), 0o644); err != nil {
			return "", err
		}
	}
	if err := exec.Command("git", "-C", root, "rev-parse", "--git-dir").Run(); err != nil {
		return status, nil
	}
	for _, kv := range [][2]string{{"merge.provenance.name", "provenance version manifest"}, {"merge.provenance.driver", "provenance merge-manifest %O %A %B"}} {
		if out, err := exec.Command("git", "-C", root, "config", kv[0], kv[1]).CombinedOutput(); err != nil {
			return "", fmt.Errorf("git config %s: %v: %s", kv[0], err, out)
		}
	}
	return status, nil
}

// /sedum:setFunction
