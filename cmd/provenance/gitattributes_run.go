package main

// sedum:anchor:imports
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-86a13b1e","gen":"sha256:19b2d09bfe7b2e6e3363fb5af7869bf0ed912e65d9326c5255780ae198758265","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/gitattributes_run.go","path":"os"},"kwargs":"sha256:1ad7f4e792f91ec8b2be9dbc2e0bef62c33095879e0985330e93e1d57a31dcdd","body":"sha256:7c6e309eb2575d4aa451ec2ab83161d13e9ae03bbca3a4d2ecf30ce12adf0905"}
import "os"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-86a13b1e","gen":"sha256:19b2d09bfe7b2e6e3363fb5af7869bf0ed912e65d9326c5255780ae198758265","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/gitattributes_run.go","path":"path/filepath"},"kwargs":"sha256:f0dc9a53d20551e19af7937740b0105defa277c7cb6dfd1976498ff228fb3e10","body":"sha256:a6d196b558491e646be93f56a321f339068c64c76aa9d49bb18046e7f9c777a1"}
import "path/filepath"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-86a13b1e","gen":"sha256:19b2d09bfe7b2e6e3363fb5af7869bf0ed912e65d9326c5255780ae198758265","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/gitattributes_run.go","path":"strings"},"kwargs":"sha256:a2804f924e53a6f71ff858536d36741ae271fa7851c5f8ffc225f0a9cbacffdb","body":"sha256:a9af3c0bc4bf9342df9d7e89a324523fdba7941a9ca4190270e68ff3c4a25ea6"}
import "strings"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-86a13b1e","gen":"sha256:19b2d09bfe7b2e6e3363fb5af7869bf0ed912e65d9326c5255780ae198758265","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/gitattributes_run.go","path":"fmt"},"kwargs":"sha256:8f096b91856967851787276ac67e1d6c3432043a2d8a8383b93cc1b598b80e2a","body":"sha256:bcbac51883cbf8064e9196b371cb880cee429866652447ad557533d817857609"}
import "fmt"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-86a13b1e","gen":"sha256:19b2d09bfe7b2e6e3363fb5af7869bf0ed912e65d9326c5255780ae198758265","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/gitattributes_run.go","path":"os/exec"},"kwargs":"sha256:24c32c9b4cb7949c5c0296b064ea8fde832723a6d7431d0e705bc564a65f78c6","body":"sha256:5970e9ca793f55e2beb08de623e98bdbc0ced178860fedd953eef7e15026ee19"}
import "os/exec"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setConst {"v":2,"tier":"owned","record":"prov-2026-86a13b1e","gen":"sha256:19b2d09bfe7b2e6e3363fb5af7869bf0ed912e65d9326c5255780ae198758265","pkg":"sha256:51be1c18998d63cc558c253e211a6bd188d96cb251dad68c6ea88e03c2d5cc4c","id":{"file":"cmd/provenance/gitattributes_run.go","name":"gitattributesFile"},"kwargs":"sha256:ef183f2f7a00bad3c7d31fabb66bbfe138153a6f2bd664bc0ad51a3f05435da3","body":"sha256:b450b430e44953b46a4ff2d7b0856ca1be9ad9d6128061244f5952544dbd78d0"}
// gitattributesFile is the git attributes file at the repository root.
const gitattributesFile = ".gitattributes"

// /sedum:setConst
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-86a13b1e","gen":"sha256:19b2d09bfe7b2e6e3363fb5af7869bf0ed912e65d9326c5255780ae198758265","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/gitattributes_run.go","name":"generatedAttrLines"},"kwargs":"sha256:bd778754f1a0c3b5eb6ab2f049a2dbbeffc786eb3cf2d37c35268cfa969e6227","body":"sha256:5246f96a03eeba8fa2df21020dc0b20d2cd7c8bac7589226ee4af2af972d195f"}
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
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-86a13b1e","gen":"sha256:19b2d09bfe7b2e6e3363fb5af7869bf0ed912e65d9326c5255780ae198758265","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/gitattributes_run.go","name":"missingAttrLines"},"kwargs":"sha256:e060950bd787866b9f86cb2a25c3c9dde997564b7917ae0596102aef31c0276c","body":"sha256:2af066c508175d495e99bbcf303a4860d896a6395f3b08bdc574dc00b64a2443"}
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
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-86a13b1e","gen":"sha256:19b2d09bfe7b2e6e3363fb5af7869bf0ed912e65d9326c5255780ae198758265","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/gitattributes_run.go","name":"ensureGitattributes"},"kwargs":"sha256:4046fc6fd3527a73ad4a0c104d2f85733094c9592c6554ca8f5628f076708915","body":"sha256:e60af2d1ee84838b28199da5cec5afa39d62215fc28ec53754edd361935631ba"}
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
// sedum:setConst {"v":2,"tier":"owned","record":"prov-2026-86a13b1e","gen":"sha256:19b2d09bfe7b2e6e3363fb5af7869bf0ed912e65d9326c5255780ae198758265","pkg":"sha256:51be1c18998d63cc558c253e211a6bd188d96cb251dad68c6ea88e03c2d5cc4c","id":{"file":"cmd/provenance/gitattributes_run.go","name":"manifestMergeAttr"},"kwargs":"sha256:0c5de825c8b9b6ffc938bfcdf46f3700e4383ef22a17cea644b982c70f0c1d1f","body":"sha256:f20bde20473917e18b4f6e0cc26f95d88efe857152217c0440630d9d4c655446"}
// manifestMergeAttr routes the version manifest through the provenance merge driver.
const manifestMergeAttr = manifestPath + " merge=provenance"

// /sedum:setConst
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-86a13b1e","gen":"sha256:19b2d09bfe7b2e6e3363fb5af7869bf0ed912e65d9326c5255780ae198758265","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/gitattributes_run.go","name":"ensureMergeDriver"},"kwargs":"sha256:b1661b7c496b09a5f3e1a436f4a1694a02215a1da891846996277cdc8d140c44","body":"sha256:dc710aa0964d2a2337d8189c12458d9dc892b33a8b96e728b6db246accec7f23"}
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
