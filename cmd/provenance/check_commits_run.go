package main

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-5577cfa8","gen":"sha256:23d91dd2a4ee8beef8cfb151d8370aae9a25e27df2af3110ce3dfa77ec51de22","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/check_commits_run.go","path":"bytes"}}
import "bytes"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-5577cfa8","gen":"sha256:23d91dd2a4ee8beef8cfb151d8370aae9a25e27df2af3110ce3dfa77ec51de22","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/check_commits_run.go","path":"fmt"}}
import "fmt"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-5577cfa8","gen":"sha256:23d91dd2a4ee8beef8cfb151d8370aae9a25e27df2af3110ce3dfa77ec51de22","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/check_commits_run.go","path":"io"}}
import "io"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-5577cfa8","gen":"sha256:23d91dd2a4ee8beef8cfb151d8370aae9a25e27df2af3110ce3dfa77ec51de22","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/check_commits_run.go","path":"os"}}
import "os"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-5577cfa8","gen":"sha256:23d91dd2a4ee8beef8cfb151d8370aae9a25e27df2af3110ce3dfa77ec51de22","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/check_commits_run.go","path":"os/exec"}}
import "os/exec"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-5577cfa8","gen":"sha256:23d91dd2a4ee8beef8cfb151d8370aae9a25e27df2af3110ce3dfa77ec51de22","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/check_commits_run.go","path":"sort"}}
import "sort"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-5577cfa8","gen":"sha256:23d91dd2a4ee8beef8cfb151d8370aae9a25e27df2af3110ce3dfa77ec51de22","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/check_commits_run.go","path":"strings"}}
import "strings"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"tier":"owned","record":"prov-2026-5577cfa8","gen":"sha256:23d91dd2a4ee8beef8cfb151d8370aae9a25e27df2af3110ce3dfa77ec51de22","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func commitFiles(r *repo, sha string) ([]string, error) {\n\tline, err := r.git(\"rev-list\", \"--parents\", \"-n\", \"1\", sha)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\tparents := strings.Fields(line)[1:]\n\tvar sets []map[string]bool\n\tif len(parents) == 0 {\n\t\tout, err := r.git(\"diff-tree\", \"--root\", \"--no-commit-id\", \"--name-only\", \"-r\", \"--no-renames\", sha)\n\t\tif err != nil {\n\t\t\treturn nil, err\n\t\t}\n\t\tset := map[string]bool{}\n\t\tfor _, f := range strings.Fields(out) {\n\t\t\tset[f] = true\n\t\t}\n\t\tsets = append(sets, set)\n\t}\n\tfor _, p := range parents {\n\t\tout, err := r.git(\"diff-tree\", \"--no-commit-id\", \"--name-only\", \"-r\", \"--no-renames\", p, sha)\n\t\tif err != nil {\n\t\t\treturn nil, err\n\t\t}\n\t\tset := map[string]bool{}\n\t\tfor _, f := range strings.Split(out, \"\\n\") {\n\t\t\tif f != \"\" {\n\t\t\t\tset[f] = true\n\t\t\t}\n\t\t}\n\t\tsets = append(sets, set)\n\t}\n\tvar files []string\n\tfor f := range sets[0] {\n\t\tinAll := true\n\t\tfor _, s := range sets[1:] {\n\t\t\tif !s[f] {\n\t\t\t\tinAll = false\n\t\t\t}\n\t\t}\n\t\tif inAll {\n\t\t\tfiles = append(files, f)\n\t\t}\n\t}\n\tsort.Strings(files)\n\treturn files, nil\n}","file":"cmd/provenance/check_commits_run.go","name":"commitFiles"}}
func commitFiles(r *repo, sha string) ([]string, error) {
	line, err := r.git("rev-list", "--parents", "-n", "1", sha)
	if err != nil {
		return nil, err
	}
	parents := strings.Fields(line)[1:]
	var sets []map[string]bool
	if len(parents) == 0 {
		out, err := r.git("diff-tree", "--root", "--no-commit-id", "--name-only", "-r", "--no-renames", sha)
		if err != nil {
			return nil, err
		}
		set := map[string]bool{}
		for _, f := range strings.Fields(out) {
			set[f] = true
		}
		sets = append(sets, set)
	}
	for _, p := range parents {
		out, err := r.git("diff-tree", "--no-commit-id", "--name-only", "-r", "--no-renames", p, sha)
		if err != nil {
			return nil, err
		}
		set := map[string]bool{}
		for _, f := range strings.Split(out, "\n") {
			if f != "" {
				set[f] = true
			}
		}
		sets = append(sets, set)
	}
	var files []string
	for f := range sets[0] {
		inAll := true
		for _, s := range sets[1:] {
			if !s[f] {
				inAll = false
			}
		}
		if inAll {
			files = append(files, f)
		}
	}
	sort.Strings(files)
	return files, nil
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-5577cfa8","gen":"sha256:23d91dd2a4ee8beef8cfb151d8370aae9a25e27df2af3110ce3dfa77ec51de22","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func commitTree(r *repo, sha string) (string, func(), error) {\n\tdir, err := os.MkdirTemp(\"\", \"provenance-commit-\")\n\tif err != nil {\n\t\treturn \"\", nil, err\n\t}\n\tc := exec.Command(\"sh\", \"-c\", `git -C \"$0\" archive \"$1\" | tar -x -C \"$2\"`, r.Root, sha, dir)\n\tc.Env = r.gitEnv()\n\tvar stderr bytes.Buffer\n\tc.Stderr = &stderr\n\tif err := c.Run(); err != nil {\n\t\tos.RemoveAll(dir)\n\t\treturn \"\", nil, fmt.Errorf(\"git archive %s: %w\\n%s\", sha, err, stderr.String())\n\t}\n\treturn dir, func() { os.RemoveAll(dir) }, nil\n}","file":"cmd/provenance/check_commits_run.go","name":"commitTree"}}
func commitTree(r *repo, sha string) (string, func(), error) {
	dir, err := os.MkdirTemp("", "provenance-commit-")
	if err != nil {
		return "", nil, err
	}
	c := exec.Command("sh", "-c", `git -C "$0" archive "$1" | tar -x -C "$2"`, r.Root, sha, dir)
	c.Env = r.gitEnv()
	var stderr bytes.Buffer
	c.Stderr = &stderr
	if err := c.Run(); err != nil {
		os.RemoveAll(dir)
		return "", nil, fmt.Errorf("git archive %s: %w\n%s", sha, err, stderr.String())
	}
	return dir, func() { os.RemoveAll(dir) }, nil
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-5577cfa8","gen":"sha256:23d91dd2a4ee8beef8cfb151d8370aae9a25e27df2af3110ce3dfa77ec51de22","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func checkCommit(r *repo, sha string) ([]violation, error) {\n\tfiles, err := commitFiles(r, sha)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\tmessage, err := r.git(\"log\", \"-1\", \"--format=%B\", sha)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\ttree, cleanup, err := commitTree(r, sha)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\tdefer cleanup()\n\tbase := sha + \"^\"\n\tif _, err := r.git(\"rev-parse\", \"-q\", \"--verify\", base); err != nil {\n\t\tbase, err = r.git(\"hash-object\", \"-t\", \"tree\", \"/dev/null\")\n\t\tif err != nil {\n\t\t\treturn nil, err\n\t\t}\n\t}\n\treturn checkChange(r, base, tree, files, message)\n}","file":"cmd/provenance/check_commits_run.go","name":"checkCommit"}}
func checkCommit(r *repo, sha string) ([]violation, error) {
	files, err := commitFiles(r, sha)
	if err != nil {
		return nil, err
	}
	message, err := r.git("log", "-1", "--format=%B", sha)
	if err != nil {
		return nil, err
	}
	tree, cleanup, err := commitTree(r, sha)
	if err != nil {
		return nil, err
	}
	defer cleanup()
	base := sha + "^"
	if _, err := r.git("rev-parse", "-q", "--verify", base); err != nil {
		base, err = r.git("hash-object", "-t", "tree", "/dev/null")
		if err != nil {
			return nil, err
		}
	}
	return checkChange(r, base, tree, files, message)
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-5577cfa8","gen":"sha256:23d91dd2a4ee8beef8cfb151d8370aae9a25e27df2af3110ce3dfa77ec51de22","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func runCheckCommits(out io.Writer, r *repo, commit, rng string) error {\n\tvar shas []string\n\tif commit != \"\" {\n\t\tsha, err := r.git(\"rev-parse\", \"--verify\", commit+\"^{commit}\")\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n\t\tshas = append(shas, sha)\n\t}\n\tif rng != \"\" {\n\t\tlist, err := r.git(\"rev-list\", \"--reverse\", rng)\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n\t\tshas = append(shas, strings.Fields(list)...)\n\t}\n\tfailed := 0\n\tfor _, sha := range shas {\n\t\tvs, err := checkCommit(r, sha)\n\t\tif err != nil {\n\t\t\tfmt.Fprintf(out, \"\\n✗ Check failed for %s: %v\\n\\n\", sha[:7], err)\n\t\t\treturn err\n\t\t}\n\t\tif len(vs) == 0 {\n\t\t\tcontinue\n\t\t}\n\t\tfailed++\n\t\tfmt.Fprintf(out, \"\\n✗ %s\\n\", sha[:7])\n\t\tfor _, v := range vs {\n\t\t\tfmt.Fprintf(out, \"    · %s\\n\", v.Message)\n\t\t}\n\t}\n\tif failed > 0 {\n\t\tfmt.Fprint(out, \"\\n\")\n\t\treturn fmt.Errorf(\"%d of %d commits violate their records\", failed, len(shas))\n\t}\n\tfmt.Fprintf(out, \"\\n✓ No forbidden scope violations in %d commit(s)\\n\\n\", len(shas))\n\treturn nil\n}","file":"cmd/provenance/check_commits_run.go","name":"runCheckCommits"}}
func runCheckCommits(out io.Writer, r *repo, commit, rng string) error {
	var shas []string
	if commit != "" {
		sha, err := r.git("rev-parse", "--verify", commit+"^{commit}")
		if err != nil {
			return err
		}
		shas = append(shas, sha)
	}
	if rng != "" {
		list, err := r.git("rev-list", "--reverse", rng)
		if err != nil {
			return err
		}
		shas = append(shas, strings.Fields(list)...)
	}
	failed := 0
	for _, sha := range shas {
		vs, err := checkCommit(r, sha)
		if err != nil {
			fmt.Fprintf(out, "\n✗ Check failed for %s: %v\n\n", sha[:7], err)
			return err
		}
		if len(vs) == 0 {
			continue
		}
		failed++
		fmt.Fprintf(out, "\n✗ %s\n", sha[:7])
		for _, v := range vs {
			fmt.Fprintf(out, "    · %s\n", v.Message)
		}
	}
	if failed > 0 {
		fmt.Fprint(out, "\n")
		return fmt.Errorf("%d of %d commits violate their records", failed, len(shas))
	}
	fmt.Fprintf(out, "\n✓ No forbidden scope violations in %d commit(s)\n\n", len(shas))
	return nil
}

// /sedum:setFunction
