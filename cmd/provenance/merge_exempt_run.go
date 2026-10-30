package main

// sedum:anchor:imports
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-0cd78a69","gen":"sha256:1dc1772259692a10c337c2e4de26790175e7c3babe17ec512bad7f63bf72e2a7","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/merge_exempt_run.go","path":"errors"},"kwargs":"sha256:60119789fd13550af88a387b1e1fb326d9caf879ca80595d050390c734672766","body":"sha256:2b718237ba0a56467c8fd494cc131e942ef410e6228c03b644d70476c71862dc"}
import "errors"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-0cd78a69","gen":"sha256:1dc1772259692a10c337c2e4de26790175e7c3babe17ec512bad7f63bf72e2a7","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/merge_exempt_run.go","path":"os/exec"},"kwargs":"sha256:517785a01b9789c6f111f455392b737f8375f471b90d61ac5bfcf3797562ac01","body":"sha256:5970e9ca793f55e2beb08de623e98bdbc0ced178860fedd953eef7e15026ee19"}
import "os/exec"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-0cd78a69","gen":"sha256:1dc1772259692a10c337c2e4de26790175e7c3babe17ec512bad7f63bf72e2a7","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/merge_exempt_run.go","path":"strings"},"kwargs":"sha256:aab1596cab9b166947b77c0056c40afb4f15ad5ecf93998ef4b95a98d3273331","body":"sha256:a9af3c0bc4bf9342df9d7e89a324523fdba7941a9ca4190270e68ff3c4a25ea6"}
import "strings"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-25b92cac","gen":"sha256:d4c3e0911f8f61f87015834a62392fade2aeb0be841e49369942f84280739e5b","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/merge_exempt_run.go","path":"os"},"kwargs":"sha256:79c29c54f0d6d9d5ed6c913f2266404a69a244ea53a2d97119df2570110b7ed3","body":"sha256:7c6e309eb2575d4aa451ec2ab83161d13e9ae03bbca3a4d2ecf30ce12adf0905"}
import "os"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-25b92cac","gen":"sha256:d4c3e0911f8f61f87015834a62392fade2aeb0be841e49369942f84280739e5b","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/merge_exempt_run.go","path":"path/filepath"},"kwargs":"sha256:12b7b1aabdb75dd1e6c612e01f0213bb4a4d95a4ddb3b547f3b56567bbfa7a75","body":"sha256:a6d196b558491e646be93f56a321f339068c64c76aa9d49bb18046e7f9c777a1"}
import "path/filepath"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-06e21d9f","gen":"sha256:7e62ff0f5ab2fb79a08c2d0b59160c7c2bf93b74c2e8681b6f4a7eda9e1a7d70","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/merge_exempt_run.go","name":"commitParents"},"kwargs":"sha256:36185f27303e8bccf6f532ef75376d81ae75a8f14f671e8d94b7cb5088161bd0","body":"sha256:de47e3d6c9cbba61d5ba69ccc414fc759e3b101f3c30887759d26059489c31c3"}
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
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-0cd78a69","gen":"sha256:1dc1772259692a10c337c2e4de26790175e7c3babe17ec512bad7f63bf72e2a7","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/merge_exempt_run.go","name":"autoMergeTree"},"kwargs":"sha256:c6bf6765db6f12a5fa8ed599e2f4c78adcdfe98a021c46b49e9539ca56a7d988","body":"sha256:df3a2e98aa0419140e0661c84ce550c5b7b8f90af029ef7075e9bf08e730df00"}
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
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-25b92cac","gen":"sha256:d4c3e0911f8f61f87015834a62392fade2aeb0be841e49369942f84280739e5b","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/merge_exempt_run.go","name":"pendingMergeParents"},"kwargs":"sha256:dfebbd82c17effdb5ce00e3e67829210d50c0965b7d760417940483e1696c7b5","body":"sha256:7c8645e9de30b59962fdac7cf304c84f99ac922e18df40c822fff9c805032fef"}
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
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-55145cec","gen":"sha256:b4aaede5ead83caf69b2d9a6038d1486ed3afb82606b0ed932a18070c7617ee2","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/merge_exempt_run.go","name":"mergeIsAutomatic"},"kwargs":"sha256:97bd1bc26ddb9d251889a6fcd2d227cbf4ea5547c03e708477f61569e8702bc2","body":"sha256:791f1b49063d0491b5b9616410063e6fa9b1f2b2a4347db598cd3eb3c8c5d2f0"}
func mergeIsAutomatic(r *repo, treeID string, parents []string) bool {
	if len(parents) != 2 {
		return false
	}
	autoTree, conflicts, err := autoMergeTree(r, parents[0], parents[1])
	if err != nil {
		return false
	}
	out, err := r.git("diff-tree", "-r", "--name-only", "--no-renames", "--relative", autoTree, treeID)
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
