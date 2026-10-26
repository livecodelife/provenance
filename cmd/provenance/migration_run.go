package main

// sedum:anchor:imports
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-6b30af71","gen":"sha256:2afcc86f270aa4eb1ff2f6a5d32a09219e45161c03f8df1ea5ca8964f6a86872","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/migration_run.go","path":"os"},"kwargs":"sha256:ed72b54100880d5a8bb77fb74d560af46135765bbcbd8e0c111c9a65571b3c13","body":"sha256:7c6e309eb2575d4aa451ec2ab83161d13e9ae03bbca3a4d2ecf30ce12adf0905"}
import "os"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-6b30af71","gen":"sha256:2afcc86f270aa4eb1ff2f6a5d32a09219e45161c03f8df1ea5ca8964f6a86872","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/migration_run.go","path":"os/exec"},"kwargs":"sha256:07febf2323cc38ba876caf99488e74e768a5e993f06654b2467bc820b305e555","body":"sha256:5970e9ca793f55e2beb08de623e98bdbc0ced178860fedd953eef7e15026ee19"}
import "os/exec"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-6b30af71","gen":"sha256:2afcc86f270aa4eb1ff2f6a5d32a09219e45161c03f8df1ea5ca8964f6a86872","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/migration_run.go","path":"path/filepath"},"kwargs":"sha256:e5a9bb671511f4dd5de7f8120977204922468db08c6f807cf6b3ea2b22a71e66","body":"sha256:a6d196b558491e646be93f56a321f339068c64c76aa9d49bb18046e7f9c777a1"}
import "path/filepath"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-6b30af71","gen":"sha256:2afcc86f270aa4eb1ff2f6a5d32a09219e45161c03f8df1ea5ca8964f6a86872","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/migration_run.go","path":"strings"},"kwargs":"sha256:1169ac1e9a7940c3b33981417ff62389bc28639133b2013b37d348f5d2ab8f22","body":"sha256:a9af3c0bc4bf9342df9d7e89a324523fdba7941a9ca4190270e68ff3c4a25ea6"}
import "strings"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-1b5e1e6e","gen":"sha256:f49b3583629862cc2b5732062082755818d184203dca18a2e2c24dc87a88fd1f","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/migration_run.go","name":"exportBase"},"kwargs":"sha256:9c187693c5b20b28b62a036a4f2c6587c6f7174c1aebc96ad28a445bfa545281","body":"sha256:1a48ae2cd03b16cd1b98be38128d4ae37647484479c52acde2a6b0fb292c5f27"}
func exportBase(r *repo, base string, dest string) bool {
	gitCmd := exec.Command("git", "archive", "--format=tar", base+":./")
	gitCmd.Dir = r.Root
	gitCmd.Env = r.gitEnv()
	gitCmd.Stderr = nil
	archive, err := gitCmd.Output()
	if err != nil {
		return false
	}
	tarCmd := exec.Command("tar", "-x", "-C", dest)
	tarCmd.Stdin = strings.NewReader(string(archive))
	if err := tarCmd.Run(); err != nil {
		return false
	}
	return true
}

// /sedum:setFunction
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-3cc00a7b","gen":"sha256:e7bcae9394b96a736ebc1a80890c60fbf6c53b4cba593c2acec70526b76c952f","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/migration_run.go","name":"classesUnchanged"},"kwargs":"sha256:fef23d539c5a518787e7a5e7c56fdb711c4642cc484f8fe3a378b669fd0b711d","body":"sha256:8b64d38e7a86fe561b4039c3905d6b3deaeb6767d533abb13237c3c97263ed7d"}
func classesUnchanged(r *repo, base string, treePath string, ids map[string]bool) bool {
	tmp, err := os.MkdirTemp("", "sedum-base-")
	if err != nil {
		return false
	}
	defer os.RemoveAll(tmp)

	if !exportBase(r, base, tmp) {
		return false
	}

	baseClasses, ok := recordClasses(r, tmp)
	if !ok {
		return false
	}
	stagedClasses, ok := recordClasses(r, treePath)
	if !ok {
		return false
	}

	for id := range ids {
		baseClass, ok := baseClasses[id]
		if !ok {
			return false
		}
		stagedClass, ok := stagedClasses[id]
		if !ok {
			return false
		}
		if baseClass != stagedClass {
			return false
		}
	}

	return true
}

// /sedum:setFunction
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-6b30af71","gen":"sha256:2afcc86f270aa4eb1ff2f6a5d32a09219e45161c03f8df1ea5ca8964f6a86872","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/migration_run.go","name":"recordClasses"},"kwargs":"sha256:f043717698d265e2c9660b3f0eba12af791aaad7462d6ce79fea9f160c3143f9","body":"sha256:d166f010744a6dd0c3f7a4b5becb4f1988b2339e0e5aade87101a6ccd58a59ca"}
func recordClasses(r *repo, treePath string) (map[string]string, bool) {
	bin := os.Getenv("SEDUM_BIN")
	if bin == "" {
		bin = "sedum"
	}
	generators := filepath.Join(treePath, "generators")
	if info, err := os.Stat(generators); err != nil || !info.IsDir() {
		empty, err := os.MkdirTemp("", "sedum-empty-")
		if err != nil {
			return nil, false
		}
		defer os.RemoveAll(empty)
		generators = empty
	}
	records := filepath.Join(treePath, r.Config.Dir)
	cmd := exec.Command(bin, "grow", "--check", "--generators", generators, "--records", records, "--output", treePath)
	cmd.Dir = treePath
	out, err := cmd.Output()
	if err != nil {
		if _, ok := err.(*exec.ExitError); !ok {
			return nil, false
		}
	}
	classes := make(map[string]string)
	for _, line := range strings.Split(string(out), "\n") {
		if line == "" || line[0] == ' ' || line[0] == '\t' {
			continue
		}
		fields := strings.FieldsFunc(line, func(c rune) bool { return c == ' ' || c == '\t' })
		if len(fields) < 2 {
			continue
		}
		class := fields[0]
		lower := class != ""
		for _, c := range class {
			if c < 'a' || c > 'z' {
				lower = false
				break
			}
		}
		if !lower {
			continue
		}
		id := fields[1]
		if _, seen := classes[id]; !seen {
			classes[id] = class
		}
	}
	return classes, true
}

// /sedum:setFunction
