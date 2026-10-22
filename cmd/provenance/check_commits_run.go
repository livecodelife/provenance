package main

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-5577cfa8","gen":"sha256:16246bada207a17685e16a70cef71c918ba64d4fc0e5d2c505f7d449b095ee2d","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/check_commits_run.go","path":"io"}}
import "io"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-6af220a3","gen":"sha256:76a5f5d143723b1856f3918485de44194b8f07a12cc6bb254c51d9859d7cc62c","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/check_commits_run.go","path":"sort"}}
import "sort"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-6af220a3","gen":"sha256:76a5f5d143723b1856f3918485de44194b8f07a12cc6bb254c51d9859d7cc62c","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/check_commits_run.go","path":"strings"}}
import "strings"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-d741e072","gen":"sha256:0d47fcd216b28b0d88c85aaa3acd14243174d60420d6357bc6bf67d1390075d2","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/check_commits_run.go","path":"bytes"}}
import "bytes"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-d741e072","gen":"sha256:0d47fcd216b28b0d88c85aaa3acd14243174d60420d6357bc6bf67d1390075d2","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/check_commits_run.go","path":"fmt"}}
import "fmt"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-d741e072","gen":"sha256:0d47fcd216b28b0d88c85aaa3acd14243174d60420d6357bc6bf67d1390075d2","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/check_commits_run.go","path":"os"}}
import "os"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-d741e072","gen":"sha256:0d47fcd216b28b0d88c85aaa3acd14243174d60420d6357bc6bf67d1390075d2","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/check_commits_run.go","path":"os/exec"}}
import "os/exec"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"tier":"owned","record":"prov-2026-5577cfa8","gen":"sha256:16246bada207a17685e16a70cef71c918ba64d4fc0e5d2c505f7d449b095ee2d","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func runCheckCommits(out io.Writer, r *repo, commit string, rng string) error {\n\tvar commits []string\n\tif commit != \"\" {\n\t\tid, err := r.git(\"rev-parse\", \"--verify\", commit+\"^{commit}\")\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n\t\tcommits = append(commits, id)\n\t}\n\tif rng != \"\" {\n\t\tids, err := r.git(\"rev-list\", \"--reverse\", rng)\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n\t\tcommits = append(commits, strings.Fields(ids)...)\n\t}\n\tfailed := 0\n\tfor _, id := range commits {\n\t\tviolations, err := checkCommit(r, id)\n\t\tif err != nil {\n\t\t\tfmt.Fprintf(out, \"\\n✗ Check failed for %s: %v\\n\\n\", id[:7], err)\n\t\t\treturn err\n\t\t}\n\t\tif len(violations) == 0 {\n\t\t\tcontinue\n\t\t}\n\t\tfailed++\n\t\tfmt.Fprintf(out, \"\\n✗ %s\\n\", id[:7])\n\t\tfor _, v := range violations {\n\t\t\tfmt.Fprintf(out, \"    · %s\\n\", v.Message)\n\t\t}\n\t}\n\tif failed > 0 {\n\t\tfmt.Fprintln(out)\n\t\treturn fmt.Errorf(\"%d of %d commits violate their records\", failed, len(commits))\n\t}\n\tfmt.Fprintf(out, \"\\n✓ No forbidden scope violations in %d commit(s)\\n\\n\", len(commits))\n\treturn nil\n}","file":"cmd/provenance/check_commits_run.go","name":"runCheckCommits"}}
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
// sedum:setFunction {"tier":"owned","record":"prov-2026-1d77499a","gen":"sha256:daef8833ac4464f6f4734a3db1568eaea8ad7ece7816204f3faab65224357a6a","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func checkCommit(r *repo, rev string) ([]violation, error) {\n\tif parents, err := commitParents(r, rev); err == nil && len(parents) == 2 {\n\t\tif tree, err := r.git(\"rev-parse\", rev+\"^{tree}\"); err == nil {\n\t\t\tif mergeIsAutomatic(r, tree, parents) {\n\t\t\t\treturn nil, nil\n\t\t\t}\n\t\t}\n\t}\n\tfiles, err := commitFiles(r, rev)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\tmessage, err := r.git(\"log\", \"-1\", \"--format=%B\", rev)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\tdir, cleanup, err := commitTree(r, rev)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\tdefer cleanup()\n\tbase, err := r.git(\"rev-parse\", \"-q\", \"--verify\", rev+\"^\")\n\tif err != nil {\n\t\tbase, err = r.git(\"hash-object\", \"-t\", \"tree\", \"/dev/null\")\n\t\tif err != nil {\n\t\t\treturn nil, err\n\t\t}\n\t}\n\treturn checkChange(r, base, dir, files, message)\n}","file":"cmd/provenance/check_commits_run.go","name":"checkCommit"}}
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
// sedum:setFunction {"tier":"owned","record":"prov-2026-6af220a3","gen":"sha256:76a5f5d143723b1856f3918485de44194b8f07a12cc6bb254c51d9859d7cc62c","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func commitFiles(r *repo, rev string) ([]string, error) {\n\tout, err := r.git(\"rev-list\", \"--parents\", \"-n\", \"1\", rev)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\tfields := strings.Fields(out)\n\tvar sets []map[string]bool\n\tif len(fields) <= 1 {\n\t\tout, err := r.git(\"diff-tree\", \"--root\", \"--no-commit-id\", \"--name-only\", \"-r\", \"--no-renames\", rev)\n\t\tif err != nil {\n\t\t\treturn nil, err\n\t\t}\n\t\tset := map[string]bool{}\n\t\tfor _, p := range strings.Fields(out) {\n\t\t\tset[p] = true\n\t\t}\n\t\tsets = append(sets, set)\n\t} else {\n\t\tfor _, parent := range fields[1:] {\n\t\t\tout, err := r.git(\"diff-tree\", \"--no-commit-id\", \"--name-only\", \"-r\", \"--no-renames\", parent, rev)\n\t\t\tif err != nil {\n\t\t\t\treturn nil, err\n\t\t\t}\n\t\t\tset := map[string]bool{}\n\t\t\tfor _, p := range strings.Split(out, \"\\n\") {\n\t\t\t\tif p != \"\" {\n\t\t\t\t\tset[p] = true\n\t\t\t\t}\n\t\t\t}\n\t\t\tsets = append(sets, set)\n\t\t}\n\t}\n\tcounts := map[string]int{}\n\tfor _, set := range sets {\n\t\tfor p := range set {\n\t\t\tcounts[p]++\n\t\t}\n\t}\n\tvar result []string\n\tfor p, n := range counts {\n\t\tif n == len(sets) {\n\t\t\tresult = append(result, p)\n\t\t}\n\t}\n\tsort.Strings(result)\n\treturn result, nil\n}","file":"cmd/provenance/check_commits_run.go","name":"commitFiles"}}
func commitFiles(r *repo, rev string) ([]string, error) {
	out, err := r.git("rev-list", "--parents", "-n", "1", rev)
	if err != nil {
		return nil, err
	}
	fields := strings.Fields(out)
	var sets []map[string]bool
	if len(fields) <= 1 {
		out, err := r.git("diff-tree", "--root", "--no-commit-id", "--name-only", "-r", "--no-renames", rev)
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
			out, err := r.git("diff-tree", "--no-commit-id", "--name-only", "-r", "--no-renames", parent, rev)
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
// sedum:setFunction {"tier":"owned","record":"prov-2026-d741e072","gen":"sha256:0d47fcd216b28b0d88c85aaa3acd14243174d60420d6357bc6bf67d1390075d2","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func commitTree(r *repo, rev string) (string, func(), error) {\n\tdir, err := os.MkdirTemp(\"\", \"provenance-commit-\")\n\tif err != nil {\n\t\treturn \"\", nil, err\n\t}\n\tscript := `git -C \"$1\" archive \"$2\" | tar -x -C \"$3\"`\n\tcmd := exec.Command(\"sh\", \"-c\", script, \"sh\", r.Root, rev, dir)\n\tcmd.Env = r.gitEnv()\n\tvar stderr bytes.Buffer\n\tcmd.Stderr = &stderr\n\tif err := cmd.Run(); err != nil {\n\t\tos.RemoveAll(dir)\n\t\treturn \"\", nil, fmt.Errorf(\"git archive %s: %w\\n%s\", rev, err, stderr.String())\n\t}\n\treturn dir, func() { os.RemoveAll(dir) }, nil\n}","file":"cmd/provenance/check_commits_run.go","name":"commitTree"}}
func commitTree(r *repo, rev string) (string, func(), error) {
	dir, err := os.MkdirTemp("", "provenance-commit-")
	if err != nil {
		return "", nil, err
	}
	script := `git -C "$1" archive "$2" | tar -x -C "$3"`
	cmd := exec.Command("sh", "-c", script, "sh", r.Root, rev, dir)
	cmd.Env = r.gitEnv()
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		os.RemoveAll(dir)
		return "", nil, fmt.Errorf("git archive %s: %w\n%s", rev, err, stderr.String())
	}
	return dir, func() { os.RemoveAll(dir) }, nil
}

// /sedum:setFunction
