package main

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-0cd78a69","gen":"sha256:68031ac3f96a3a3f0e9006298fecd34aa48479da7e6fea55e9896b5381e3c089","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/merge_exempt_run.go","path":"errors"}}
import "errors"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-0cd78a69","gen":"sha256:68031ac3f96a3a3f0e9006298fecd34aa48479da7e6fea55e9896b5381e3c089","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/merge_exempt_run.go","path":"os/exec"}}
import "os/exec"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-0cd78a69","gen":"sha256:68031ac3f96a3a3f0e9006298fecd34aa48479da7e6fea55e9896b5381e3c089","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/merge_exempt_run.go","path":"strings"}}
import "strings"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-25b92cac","gen":"sha256:d4c3e0911f8f61f87015834a62392fade2aeb0be841e49369942f84280739e5b","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/merge_exempt_run.go","path":"os"}}
import "os"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-25b92cac","gen":"sha256:d4c3e0911f8f61f87015834a62392fade2aeb0be841e49369942f84280739e5b","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/merge_exempt_run.go","path":"path/filepath"}}
import "path/filepath"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"tier":"owned","record":"prov-2026-06e21d9f","gen":"sha256:7e62ff0f5ab2fb79a08c2d0b59160c7c2bf93b74c2e8681b6f4a7eda9e1a7d70","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func commitParents(r *repo, rev string) ([]string, error) {\n\tout, err := r.git(\"rev-list\", \"--parents\", \"-n\", \"1\", rev)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\tfields := strings.Fields(out)\n\tif len(fields) == 0 {\n\t\treturn nil, nil\n\t}\n\treturn fields[1:], nil\n}","file":"cmd/provenance/merge_exempt_run.go","name":"commitParents"}}
func commitParents(r *repo, rev string) ([]string, error) {
	out, err := r.git("rev-list", "--parents", "-n", "1", rev)
	if err != nil {
		return nil, err
	}
	fields := strings.Fields(out)
	if len(fields) == 0 {
		return nil, nil
	}
	return fields[1:], nil
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-0cd78a69","gen":"sha256:68031ac3f96a3a3f0e9006298fecd34aa48479da7e6fea55e9896b5381e3c089","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func autoMergeTree(r *repo, first, second string) (string, []string, error) {\n\tcmd := exec.Command(\"git\", \"merge-tree\", \"--write-tree\", \"--name-only\", first, second)\n\tcmd.Dir = r.Root\n\tcmd.Env = r.gitEnv()\n\tout, err := cmd.Output()\n\tif err != nil {\n\t\tvar exitErr *exec.ExitError\n\t\tif !errors.As(err, &exitErr) || exitErr.ExitCode() != 1 {\n\t\t\treturn \"\", nil, err\n\t\t}\n\t}\n\tlines := strings.Split(string(out), \"\\n\")\n\ttreeID := strings.TrimSpace(lines[0])\n\tif treeID == \"\" {\n\t\treturn \"\", nil, errors.New(\"merge-tree produced no tree id\")\n\t}\n\tvar paths []string\n\tfor _, line := range lines[1:] {\n\t\tif strings.TrimSpace(line) == \"\" {\n\t\t\tbreak\n\t\t}\n\t\tpaths = append(paths, strings.TrimSpace(line))\n\t}\n\treturn treeID, paths, nil\n}","file":"cmd/provenance/merge_exempt_run.go","name":"autoMergeTree"}}
func autoMergeTree(r *repo, first, second string) (string, []string, error) {
	cmd := exec.Command("git", "merge-tree", "--write-tree", "--name-only", first, second)
	cmd.Dir = r.Root
	cmd.Env = r.gitEnv()
	out, err := cmd.Output()
	if err != nil {
		var exitErr *exec.ExitError
		if !errors.As(err, &exitErr) || exitErr.ExitCode() != 1 {
			return "", nil, err
		}
	}
	lines := strings.Split(string(out), "\n")
	treeID := strings.TrimSpace(lines[0])
	if treeID == "" {
		return "", nil, errors.New("merge-tree produced no tree id")
	}
	var paths []string
	for _, line := range lines[1:] {
		if strings.TrimSpace(line) == "" {
			break
		}
		paths = append(paths, strings.TrimSpace(line))
	}
	return treeID, paths, nil
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-25b92cac","gen":"sha256:d4c3e0911f8f61f87015834a62392fade2aeb0be841e49369942f84280739e5b","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func pendingMergeParents(r *repo) []string {\n\thead, err := r.git(\"rev-parse\", \"HEAD\")\n\tif err != nil {\n\t\treturn nil\n\t}\n\trel, err := r.git(\"rev-parse\", \"--git-path\", \"MERGE_HEAD\")\n\tif err != nil {\n\t\treturn nil\n\t}\n\tpath := rel\n\tif !filepath.IsAbs(path) {\n\t\tpath = filepath.Join(r.Root, path)\n\t}\n\tdata, err := os.ReadFile(path)\n\tif err != nil {\n\t\treturn nil\n\t}\n\tparents := []string{head}\n\tfor _, line := range strings.Split(string(data), \"\\n\") {\n\t\tline = strings.TrimSpace(line)\n\t\tif line == \"\" {\n\t\t\tcontinue\n\t\t}\n\t\tparents = append(parents, line)\n\t}\n\treturn parents\n}","file":"cmd/provenance/merge_exempt_run.go","name":"pendingMergeParents"}}
func pendingMergeParents(r *repo) []string {
	head, err := r.git("rev-parse", "HEAD")
	if err != nil {
		return nil
	}
	rel, err := r.git("rev-parse", "--git-path", "MERGE_HEAD")
	if err != nil {
		return nil
	}
	path := rel
	if !filepath.IsAbs(path) {
		path = filepath.Join(r.Root, path)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	parents := []string{head}
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parents = append(parents, line)
	}
	return parents
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-55145cec","gen":"sha256:677ee693abaed0acc1e4a59dfd064387957b5e32ee76bc5bcc3661315e9807b7","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func mergeIsAutomatic(r *repo, treeID string, parents []string) bool {\n\tif len(parents) != 2 {\n\t\treturn false\n\t}\n\tautoTree, conflicts, err := autoMergeTree(r, parents[0], parents[1])\n\tif err != nil {\n\t\treturn false\n\t}\n\tout, err := r.git(\"diff-tree\", \"-r\", \"--name-only\", \"--no-renames\", autoTree, treeID)\n\tif err != nil {\n\t\treturn false\n\t}\n\tvar diffPaths []string\n\tfor _, line := range strings.Split(out, \"\\n\") {\n\t\tif line != \"\" {\n\t\t\tdiffPaths = append(diffPaths, line)\n\t\t}\n\t}\n\tif len(diffPaths) == 0 {\n\t\treturn true\n\t}\n\tfor _, p := range diffPaths {\n\t\tif p != manifestPath {\n\t\t\treturn false\n\t\t}\n\t}\n\tfor _, c := range conflicts {\n\t\tif c == manifestPath {\n\t\t\treturn true\n\t\t}\n\t}\n\treturn false\n}","file":"cmd/provenance/merge_exempt_run.go","name":"mergeIsAutomatic"}}
func mergeIsAutomatic(r *repo, treeID string, parents []string) bool {
	if len(parents) != 2 {
		return false
	}
	autoTree, conflicts, err := autoMergeTree(r, parents[0], parents[1])
	if err != nil {
		return false
	}
	out, err := r.git("diff-tree", "-r", "--name-only", "--no-renames", autoTree, treeID)
	if err != nil {
		return false
	}
	var diffPaths []string
	for _, line := range strings.Split(out, "\n") {
		if line != "" {
			diffPaths = append(diffPaths, line)
		}
	}
	if len(diffPaths) == 0 {
		return true
	}
	for _, p := range diffPaths {
		if p != manifestPath {
			return false
		}
	}
	for _, c := range conflicts {
		if c == manifestPath {
			return true
		}
	}
	return false
}

// /sedum:setFunction
