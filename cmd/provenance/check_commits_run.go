package main

// sedum:anchor:imports
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-5577cfa8","gen":"sha256:16246bada207a17685e16a70cef71c918ba64d4fc0e5d2c505f7d449b095ee2d","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/check_commits_run.go","path":"io"},"kwargs":"sha256:8f92205dffd5e134eb936a8ac759617957abc467d97e2ab704fc628fbdcb27bd","body":"sha256:ef8e8abf4ef5ee40a19dc85b49bf09c5df1fe32abfe834e0ab4c2175baf80ddf"}
import "io"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-6af220a3","gen":"sha256:7ef5b2e9e97885ddd3d2ad89dd63ba859f60570f731c30142c4bfb3e4b123a4d","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/check_commits_run.go","path":"sort"},"kwargs":"sha256:a44655bdd2ad3e1e1918ed1651d956f5eab3401920f63f4e82dc30fd8b1436f1","body":"sha256:f197172617ae21ef23b1855b09aafb14571020ae81bff877d63d3ff81f4bf63c"}
import "sort"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-6af220a3","gen":"sha256:7ef5b2e9e97885ddd3d2ad89dd63ba859f60570f731c30142c4bfb3e4b123a4d","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/check_commits_run.go","path":"strings"},"kwargs":"sha256:99f6f52656ffdd7ca4afb27c0021957f046538e8dd2c1433af0e5773c1a4e2e5","body":"sha256:a9af3c0bc4bf9342df9d7e89a324523fdba7941a9ca4190270e68ff3c4a25ea6"}
import "strings"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-d741e072","gen":"sha256:e8fbdb78d72ad43032428f8218b40ae3126aee67c1483d2c3f8e7790203d8516","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/check_commits_run.go","path":"bytes"},"kwargs":"sha256:02e5f2c6ab1019a3791da3c155c861d9f3124e6fb0cb2dff5b8f7b173b246526","body":"sha256:468731f6dd513df0a66690d4649dc9e7244c96de227d8e6ead98cc5e9747c57a"}
import "bytes"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-d741e072","gen":"sha256:e8fbdb78d72ad43032428f8218b40ae3126aee67c1483d2c3f8e7790203d8516","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/check_commits_run.go","path":"fmt"},"kwargs":"sha256:5a8c75efa8ad082d8e50a5f0bbde560fff648b4f99fdacd9ac1daf1dddb8d960","body":"sha256:bcbac51883cbf8064e9196b371cb880cee429866652447ad557533d817857609"}
import "fmt"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-d741e072","gen":"sha256:e8fbdb78d72ad43032428f8218b40ae3126aee67c1483d2c3f8e7790203d8516","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/check_commits_run.go","path":"os"},"kwargs":"sha256:604443dd628fca3b722becae8281c0f2e0846cf1636fe275b4a9b905f4ab1d5c","body":"sha256:7c6e309eb2575d4aa451ec2ab83161d13e9ae03bbca3a4d2ecf30ce12adf0905"}
import "os"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-d741e072","gen":"sha256:e8fbdb78d72ad43032428f8218b40ae3126aee67c1483d2c3f8e7790203d8516","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/check_commits_run.go","path":"os/exec"},"kwargs":"sha256:8591e3ae263392d1ba3584a725f9b08443c3bf2b50fea2979afcdaa2c4ac9afc","body":"sha256:5970e9ca793f55e2beb08de623e98bdbc0ced178860fedd953eef7e15026ee19"}
import "os/exec"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-5577cfa8","gen":"sha256:16246bada207a17685e16a70cef71c918ba64d4fc0e5d2c505f7d449b095ee2d","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/check_commits_run.go","name":"runCheckCommits"},"kwargs":"sha256:2fc508b824127c7b1062eeee7a7a132366eb95a751dc5e19991162ca19d7b6b9","body":"sha256:d9b4c2af0a7599202688c93ef02bf40fc93a98479db4f944455f19fd973405a9"}
func runCheckCommits(out io.Writer, r *repo, commit string, rng string) error {
	var commits []string
	if commit != "" {
		id, err := r.git("rev-parse", "--verify", commit+"^{commit}")
		if err != nil {
			return err
		}
		commits = append(commits, id)
	}
	if rng != "" {
		ids, err := r.git("rev-list", "--reverse", rng)
		if err != nil {
			return err
		}
		commits = append(commits, strings.Fields(ids)...)
	}
	failed := 0
	for _, id := range commits {
		violations, err := checkCommit(r, id)
		if err != nil {
			fmt.Fprintf(out, "\n✗ Check failed for %s: %v\n\n", id[:7], err)
			return err
		}
		if len(violations) == 0 {
			continue
		}
		failed++
		fmt.Fprintf(out, "\n✗ %s\n", id[:7])
		for _, v := range violations {
			fmt.Fprintf(out, "    · %s\n", v.Message)
		}
	}
	if failed > 0 {
		fmt.Fprintln(out)
		return fmt.Errorf("%d of %d commits violate their records", failed, len(commits))
	}
	fmt.Fprintf(out, "\n✓ No forbidden scope violations in %d commit(s)\n\n", len(commits))
	return nil
}

// /sedum:setFunction
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-1d77499a","gen":"sha256:daef8833ac4464f6f4734a3db1568eaea8ad7ece7816204f3faab65224357a6a","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/check_commits_run.go","name":"checkCommit"},"kwargs":"sha256:8d98638a9f01e5e4df30a7b5bfb271183ef1b44f0f9c5a7ed9460dfb2e4c2a52","body":"sha256:5f8a3bff4ff5b11d840899df06498c10055f057d465820eaf4dcf74b61f37bc9"}
func checkCommit(r *repo, rev string) ([]violation, error) {
	if parents, err := commitParents(r, rev); err == nil && len(parents) == 2 {
		if tree, err := r.git("rev-parse", rev+"^{tree}"); err == nil {
			if mergeIsAutomatic(r, tree, parents) {
				return nil, nil
			}
		}
	}
	files, err := commitFiles(r, rev)
	if err != nil {
		return nil, err
	}
	message, err := r.git("log", "-1", "--format=%B", rev)
	if err != nil {
		return nil, err
	}
	dir, cleanup, err := commitTree(r, rev)
	if err != nil {
		return nil, err
	}
	defer cleanup()
	base, err := r.git("rev-parse", "-q", "--verify", rev+"^")
	if err != nil {
		base, err = r.git("hash-object", "-t", "tree", "/dev/null")
		if err != nil {
			return nil, err
		}
	}
	return checkChange(r, base, dir, files, message)
}

// /sedum:setFunction
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-6af220a3","gen":"sha256:7ef5b2e9e97885ddd3d2ad89dd63ba859f60570f731c30142c4bfb3e4b123a4d","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/check_commits_run.go","name":"commitFiles"},"kwargs":"sha256:51ea9251c5cfeab5aa8a564416c33f837307a5774685e8818ce727693326067e","body":"sha256:8db52b50e2b168a836984c88b9b1132e747729c314420aaa11829fe7257be178"}
func commitFiles(r *repo, rev string) ([]string, error) {
	out, err := r.git("rev-list", "--parents", "-n", "1", rev)
	if err != nil {
		return nil, err
	}
	fields := strings.Fields(out)
	var sets []map[string]bool
	if len(fields) <= 1 {
		out, err := r.git("diff-tree", "--root", "--no-commit-id", "--name-only", "-r", "--no-renames", "--relative", rev)
		if err != nil {
			return nil, err
		}
		set := map[string]bool{}
		for _, p := range strings.Fields(out) {
			set[p] = true
		}
		sets = append(sets, set)
	} else {
		for _, parent := range fields[1:] {
			out, err := r.git("diff-tree", "--no-commit-id", "--name-only", "-r", "--no-renames", "--relative", parent, rev)
			if err != nil {
				return nil, err
			}
			set := map[string]bool{}
			for _, p := range strings.Split(out, "\n") {
				if p != "" {
					set[p] = true
				}
			}
			sets = append(sets, set)
		}
	}
	counts := map[string]int{}
	for _, set := range sets {
		for p := range set {
			counts[p]++
		}
	}
	var result []string
	for p, n := range counts {
		if n == len(sets) {
			result = append(result, p)
		}
	}
	sort.Strings(result)
	return result, nil
}

// /sedum:setFunction
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-d741e072","gen":"sha256:e8fbdb78d72ad43032428f8218b40ae3126aee67c1483d2c3f8e7790203d8516","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/check_commits_run.go","name":"commitTree"},"kwargs":"sha256:de619a6e8d17e6ec8ef4af059e1bcde759f46475104037d0e948129db2fbf6c4","body":"sha256:e2efc80051a8288217339a10423d395679abb1207cc23c838d44280dac3720f4"}
func commitTree(r *repo, rev string) (string, func(), error) {
	dir, err := os.MkdirTemp("", "provenance-commit-")
	if err != nil {
		return "", nil, err
	}
	cleanup := func() { os.RemoveAll(dir) }
	top, err := r.git("rev-parse", "--show-toplevel")
	if err != nil {
		cleanup()
		return "", nil, err
	}
	prefix, err := r.git("rev-parse", "--show-prefix")
	if err != nil {
		cleanup()
		return "", nil, err
	}
	script := `git -C "$1" archive "$2" | tar -x -C "$3"`
	cmd := exec.Command("sh", "-c", script, "sh", top, rev, dir)
	cmd.Env = r.gitEnv()
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		cleanup()
		return "", nil, fmt.Errorf("git archive %s: %w\n%s", rev, err, stderr.String())
	}
	path := dir
	if prefix != "" {
		path = dir + string(os.PathSeparator) + strings.ReplaceAll(prefix, "/", string(os.PathSeparator))
	}
	return path, cleanup, nil
}

// /sedum:setFunction
