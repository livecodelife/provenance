package main

// sedum:anchor:imports
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-c31c9c59","gen":"sha256:9a18a4eb62cedc8b94c8ee92552042184ff66c2dcd6e7dd094c735af31f5de23","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/run_specs_run.go","path":"errors"},"kwargs":"sha256:0860696790834c8c40b76be4c7cfcc8fc079c2ba7ec16781d6a698b886b4a368","body":"sha256:2b718237ba0a56467c8fd494cc131e942ef410e6228c03b644d70476c71862dc"}
import "errors"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-c31c9c59","gen":"sha256:9a18a4eb62cedc8b94c8ee92552042184ff66c2dcd6e7dd094c735af31f5de23","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/run_specs_run.go","path":"fmt"},"kwargs":"sha256:ec505cee1b4dbd35210701d5b776a9a0f97f2d7387dfda42a0ced675b57239db","body":"sha256:bcbac51883cbf8064e9196b371cb880cee429866652447ad557533d817857609"}
import "fmt"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-c31c9c59","gen":"sha256:9a18a4eb62cedc8b94c8ee92552042184ff66c2dcd6e7dd094c735af31f5de23","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/run_specs_run.go","path":"io"},"kwargs":"sha256:339bee8e6d44392ae275f2db5a522da91c973168999ff8b53073984fdf5d69a0","body":"sha256:ef8e8abf4ef5ee40a19dc85b49bf09c5df1fe32abfe834e0ab4c2175baf80ddf"}
import "io"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-c31c9c59","gen":"sha256:9a18a4eb62cedc8b94c8ee92552042184ff66c2dcd6e7dd094c735af31f5de23","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/run_specs_run.go","path":"os"},"kwargs":"sha256:69589300325a8ee9dbb4ee9698ec32b431667156c345a3697bec6c7253d0b30b","body":"sha256:7c6e309eb2575d4aa451ec2ab83161d13e9ae03bbca3a4d2ecf30ce12adf0905"}
import "os"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-c31c9c59","gen":"sha256:9a18a4eb62cedc8b94c8ee92552042184ff66c2dcd6e7dd094c735af31f5de23","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/run_specs_run.go","path":"os/exec"},"kwargs":"sha256:51dcedffde2fc70ebe546c590a9443efecc3cf01f1c4fa79f00b0f8344c07f8e","body":"sha256:5970e9ca793f55e2beb08de623e98bdbc0ced178860fedd953eef7e15026ee19"}
import "os/exec"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-c31c9c59","gen":"sha256:9a18a4eb62cedc8b94c8ee92552042184ff66c2dcd6e7dd094c735af31f5de23","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/run_specs_run.go","path":"strings"},"kwargs":"sha256:d2acbd55b06a86e996bb072a3a63bfe820591bcf6749c00c6f4c3c9e34fa8d21","body":"sha256:a9af3c0bc4bf9342df9d7e89a324523fdba7941a9ca4190270e68ff3c4a25ea6"}
import "strings"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-c31c9c59","gen":"sha256:9a18a4eb62cedc8b94c8ee92552042184ff66c2dcd6e7dd094c735af31f5de23","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/run_specs_run.go","path":"github.com/spf13/cobra"},"kwargs":"sha256:4d2f5a1ad2ab56bfe65af520bde3ddf1e14eb568f13da5f5d2cf8ecad59f4639","body":"sha256:f361c02d0f8e1999cf3fb715393e16310eeec9debd5aaa13936af7617e2b46aa"}
import "github.com/spf13/cobra"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-c31c9c59","gen":"sha256:9a18a4eb62cedc8b94c8ee92552042184ff66c2dcd6e7dd094c735af31f5de23","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/run_specs_run.go","path":"github.com/livecodelife/provenance/record"},"kwargs":"sha256:9adc6fd048b573d09dc2fc2883718fc5059819a28bb238bf78a3408fa5afedcb","body":"sha256:acbf8628061291d11aa39e9ea64ecc91095a117057cfe6f1958ef2c437fac003"}
import "github.com/livecodelife/provenance/record"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setType {"v":2,"tier":"owned","record":"prov-2026-c31c9c59","gen":"sha256:9a18a4eb62cedc8b94c8ee92552042184ff66c2dcd6e7dd094c735af31f5de23","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","id":{"file":"cmd/provenance/run_specs_run.go","name":"specOutcome"},"kwargs":"sha256:2ad438aa7e88e7daeec94e585ae3f40cad887c635a6be76f4bce369f590ef230","body":"sha256:9e210623896492ac0c84c616583c8a9013a5754f9126b1a8e5dff84ca40a126f"}
type specOutcome struct {
	Path   string
	Status string
}

// /sedum:setType
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-c31c9c59","gen":"sha256:9a18a4eb62cedc8b94c8ee92552042184ff66c2dcd6e7dd094c735af31f5de23","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/run_specs_run.go","name":"specCommandBase"},"kwargs":"sha256:0eb19a5a7095522f10fd6d396cd0dc165d59dbd4a2d73ac6d6b5eb7b17c5fc65","body":"sha256:4b0c489ebf3d65005ded5916f249f05490c5340dfb4cfc042706e551f9327980"}
func specCommandBase(spec record.AssociatedSpec) (base string, templated bool, skip bool) {
	if spec.RunCommand != "" {
		if strings.Contains(spec.RunCommand, "{{path}}") {
			return "", true, false
		}
		return spec.RunCommand, false, false
	}
	switch spec.Type {
	case "linespec":
		if _, err := os.Stat("./linespec"); err == nil {
			return "./linespec test", false, false
		}
		return "linespec test", false, false
	case "rspec":
		return "bundle exec rspec", false, false
	case "pytest":
		return "pytest", false, false
	case "jest":
		return "npx jest", false, false
	}
	return "", false, true
}

// /sedum:setFunction
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-c31c9c59","gen":"sha256:9a18a4eb62cedc8b94c8ee92552042184ff66c2dcd6e7dd094c735af31f5de23","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/run_specs_run.go","name":"buildSpecCommand"},"kwargs":"sha256:48d5e374ad58ec7dd6c53abb31399ab277a24e6132ca68af28bccabbbbf3e0ec","body":"sha256:5f6832a231a055f30a3cdca6e0c85334080d3a6b7a977699bd20ee66408c06c5"}
func buildSpecCommand(spec record.AssociatedSpec) (string, bool) {
	base, templated, skip := specCommandBase(spec)
	if skip {
		return "", true
	}
	if templated {
		return strings.ReplaceAll(spec.RunCommand, "{{path}}", spec.Path), false
	}
	return base + " " + spec.Path, false
}

// /sedum:setFunction
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-c31c9c59","gen":"sha256:9a18a4eb62cedc8b94c8ee92552042184ff66c2dcd6e7dd094c735af31f5de23","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/run_specs_run.go","name":"specEnv"},"kwargs":"sha256:08fd5cbb5b174d5708482f335fc155e11a93452de850618692092d60b07f05ab","body":"sha256:0fe4b81699d640872fd1b54d8c639621f4fcc1dc827b863ad1dc0587776504a8"}
func specEnv() []string {
	var env []string
	for _, kv := range os.Environ() {
		if strings.HasPrefix(kv, "GIT_") {
			continue
		}
		env = append(env, kv)
	}
	return env
}

// /sedum:setFunction
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-c31c9c59","gen":"sha256:9a18a4eb62cedc8b94c8ee92552042184ff66c2dcd6e7dd094c735af31f5de23","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/run_specs_run.go","name":"runShell"},"kwargs":"sha256:304359fd7b2ca482eb4b167939863acaaa156b939fc034150fb46185b5540383","body":"sha256:e56e7b49dd7475b712e343e0062658c7d4133eb7ec5d2f81b8d34202f45472fd"}
func runShell(out io.Writer, dir, command string) error {
	cmd := exec.Command("sh", "-c", command)
	cmd.Dir = dir
	cmd.Env = specEnv()
	cmd.Stdout = out
	cmd.Stderr = out
	return cmd.Run()
}

// /sedum:setFunction
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-c31c9c59","gen":"sha256:9a18a4eb62cedc8b94c8ee92552042184ff66c2dcd6e7dd094c735af31f5de23","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/run_specs_run.go","name":"runRecordSpecs"},"kwargs":"sha256:3c4156b2c469e78c587a2afe3d43915f010577091be944dd48157e4d58d29c51","body":"sha256:6dbed110c559d2857e163da3ac985733a7c8adad5a2bc74ad031816e0410ceb6"}
func runRecordSpecs(out io.Writer, r *repo, rec *record.Record, seen map[string]bool) (ran bool, failed bool, outcomes []specOutcome) {
	specs := rec.AssociatedSpecs
	for i := 0; i < len(specs); {
		spec := specs[i]
		_, skip := buildSpecCommand(spec)
		if skip {
			fmt.Fprintf(out, "  · %s  (skipped — no type or run_command)\n", spec.Path)
			outcomes = append(outcomes, specOutcome{Path: spec.Path, Status: "skipped"})
			i++
			continue
		}
		base, templated, _ := specCommandBase(spec)
		group := []record.AssociatedSpec{spec}
		j := i + 1
		if !templated {
			for j < len(specs) {
				next := specs[j]
				nb, nt, ns := specCommandBase(next)
				if ns || nt || nb != base {
					break
				}
				group = append(group, next)
				j++
			}
		}
		ran = true
		var kept []record.AssociatedSpec
		var keptCmds []string
		for _, m := range group {
			c, _ := buildSpecCommand(m)
			if seen != nil && seen[c] {
				fmt.Fprintf(out, "  · %s  (already verified)\n", c)
				outcomes = append(outcomes, specOutcome{Path: m.Path, Status: "passed"})
				continue
			}
			kept = append(kept, m)
			keptCmds = append(keptCmds, c)
		}
		if seen != nil {
			for _, c := range keptCmds {
				seen[c] = true
			}
		}
		switch len(kept) {
		case 0:
		case 1:
			fmt.Fprintf(out, "  · %s\n    %s\n", kept[0].Path, keptCmds[0])
			if err := runShell(out, r.Root, keptCmds[0]); err != nil {
				fmt.Fprint(out, "    ✗ failed\n\n")
				failed = true
				outcomes = append(outcomes, specOutcome{Path: kept[0].Path, Status: "failed"})
			} else {
				fmt.Fprint(out, "    ✓ passed\n\n")
				outcomes = append(outcomes, specOutcome{Path: kept[0].Path, Status: "passed"})
			}
		default:
			paths := make([]string, len(kept))
			for k, m := range kept {
				paths[k] = m.Path
			}
			batch := base + " " + strings.Join(paths, " ")
			fmt.Fprintf(out, "  · %s\n    %s\n", strings.Join(paths, ", "), batch)
			if err := runShell(out, r.Root, batch); err == nil {
				for _, p := range paths {
					fmt.Fprintf(out, "    ✓ %s passed\n", p)
					outcomes = append(outcomes, specOutcome{Path: p, Status: "passed"})
				}
				fmt.Fprint(out, "\n")
			} else {
				failed = true
				fmt.Fprint(out, "    ✗ batch failed — localizing per spec:\n")
				for k, m := range kept {
					if err := runShell(out, r.Root, keptCmds[k]); err != nil {
						fmt.Fprintf(out, "    ✗ %s failed\n", m.Path)
						outcomes = append(outcomes, specOutcome{Path: m.Path, Status: "failed"})
					} else {
						fmt.Fprintf(out, "    ✓ %s passed\n", m.Path)
						outcomes = append(outcomes, specOutcome{Path: m.Path, Status: "passed"})
					}
				}
				fmt.Fprint(out, "\n")
			}
		}
		i = j
	}
	return ran, failed, outcomes
}

// /sedum:setFunction
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-c31c9c59","gen":"sha256:9a18a4eb62cedc8b94c8ee92552042184ff66c2dcd6e7dd094c735af31f5de23","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/run_specs_run.go","name":"runRunSpecs"},"kwargs":"sha256:554a961c85751af908335f947cdac8a3037d2271f8bfcaa5b7fc41d0490ea03a","body":"sha256:7d6148d047d5d112aae13913453f0ae8de91161c5534ebb42071916c2c9f1304"}
func runRunSpecs(out io.Writer, cmd *cobra.Command, args []string) error {
	r, err := openRepo(cmd)
	if err != nil {
		return err
	}
	id, _ := cmd.Flags().GetString("record")
	if id == "" {
		return errors.New(`required flag(s) "record" not set`)
	}
	rec, _, err := r.resolve(id)
	if err != nil {
		return err
	}
	if !r.Config.RunAssociatedSpecsOnComplete || len(rec.AssociatedSpecs) == 0 {
		return nil
	}
	fmt.Fprintf(out, "\nRunning %d associated spec(s) for %s...\n\n", len(rec.AssociatedSpecs), rec.ID)
	_, failed, _ := runRecordSpecs(out, r, rec, nil)
	if failed {
		return errors.New("one or more associated specs failed — commit blocked")
	}
	fmt.Fprint(out, "All associated specs passed.\n\n")
	return nil
}

// /sedum:setFunction
