package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// A regeneration commit re-renders sedum regions of a file that an implemented
// record governs, without a new version. The fixtures below hold one such file
// (a.go) with two regions; the fake sedum stands in for `sedum grow --check`.

func regionHeader(action, gen, pkg, kwargs string) string {
	return `// sedum:` + action + ` {"tier":"owned","record":"` + chkID + `","gen":"` + gen + `","pkg":"` + pkg + `","kwargs":` + kwargs + `}` + "\n"
}

// regenFile renders a.go from its parts, so a test changes one part at a time.
type regenParts struct {
	action, gen, pkg, kwargs, body, outside string
	second                                  bool
}

func baseParts() regenParts {
	return regenParts{action: "addImport", gen: "sha256:g1", pkg: "sha256:p1", kwargs: `{"code":"fmt"}`, body: "import \"fmt\"\n", outside: "// helper\n", second: true}
}

func (p regenParts) render() string {
	s := "package a\n\n// sedum:anchor:imports\n" + regionHeader(p.action, p.gen, p.pkg, p.kwargs) + p.body + "// /sedum:" + p.action + "\n\n" + p.outside
	if p.second {
		s += regionHeader("setFunction", p.gen, p.pkg, `{"name":"f"}`) + "func f() {}\n" + "// /sedum:setFunction\n"
	}
	return s
}

// fakeSedum writes a sedum that answers `grow --check --only <id>` with the
// state held in its state file, and logs its arguments.
func fakeSedum(t *testing.T) (state, log string) {
	t.Helper()
	dir := t.TempDir()
	state, log = filepath.Join(dir, "state"), filepath.Join(dir, "args.log")
	writeFile(t, state, "current\n")
	script := "#!/bin/sh\necho \"$@\" >> '" + log + "'\nid=\"\"\nwhile [ $# -gt 0 ]; do\n  if [ \"$1\" = --only ]; then id=\"$2\"; fi\n  shift\ndone\n" +
		"if [ \"$(cat '" + state + "')\" = current ]; then echo \"current    $id\"; exit 0; fi\necho \"stale      $id\"\nexit 1\n"
	bin := filepath.Join(dir, "sedum")
	writeFile(t, bin, script)
	if err := os.Chmod(bin, 0o755); err != nil {
		t.Fatal(err)
	}
	t.Setenv("SEDUM_BIN", bin)
	return state, log
}

// regenRepo is checkRepo with a.go holding the base regions and the record
// completed (implemented, version 1).
func regenRepo(t *testing.T) (root, cfg, state, log string) {
	t.Helper()
	root, cfg = checkRepo(t)
	state, log = fakeSedum(t)
	stage(t, root, map[string]string{"a.go": baseParts().render()})
	gitIn(t, root, "commit", "-q", "-m", "Grow a ["+chkID+"@1]")
	path := filepath.Join("provenance", chkID+".yml")
	stage(t, root, map[string]string{path: strings.Replace(readFile(t, filepath.Join(root, path)), "status: open", "status: implemented", 1)})
	gitIn(t, root, "commit", "-q", "-m", "Complete provenance record ["+chkID+"@1]")
	return root, cfg, state, log
}

func regenMsg() string { return "Regenerate a [" + chkID + "@1]" }

func TestRegenerationCommitMovingOnlyThePackageAndBodyPasses(t *testing.T) {
	root, cfg, _, _ := regenRepo(t)
	p := baseParts()
	p.pkg, p.body = "sha256:p2", "import \"fmt\" // newer\n"
	stage(t, root, map[string]string{"a.go": p.render()})
	if out, ok := check(t, cfg, regenMsg()); !ok || !strings.Contains(out, "✓ No forbidden scope violations in staged") {
		t.Fatalf("a re-render must pass on an implemented record:\n%s", out)
	}
	gitIn(t, root, "commit", "-q", "-m", regenMsg())
	if out, err := run(t, "-c", cfg, "check", "--commit", "HEAD"); err != nil || !strings.Contains(out, "✓") {
		t.Fatalf("--commit must agree: %v\n%s", err, out)
	}
	if out, err := run(t, "-c", cfg, "check", "--range", "HEAD~1..HEAD"); err != nil || !strings.Contains(out, "✓") {
		t.Fatalf("--range must agree: %v\n%s", err, out)
	}
}

func TestRegenerationCommitMayChangeTheActionAndKwargsOfARegion(t *testing.T) {
	root, cfg, _, _ := regenRepo(t)
	p := baseParts()
	p.action, p.kwargs, p.pkg = "importFrom:named", `{"name":"fmt","from":"fmt"}`, "sha256:p2"
	stage(t, root, map[string]string{"a.go": p.render()})
	if out, ok := check(t, cfg, regenMsg()); !ok {
		t.Fatalf("converting a region to another action must pass:\n%s", out)
	}
}

func TestRegenerationCommitAsksSedumAboutTheTouchedRecordOnTheStagedTree(t *testing.T) {
	root, cfg, _, log := regenRepo(t)
	p := baseParts()
	p.pkg = "sha256:p2"
	stage(t, root, map[string]string{"a.go": p.render()})
	if _, ok := check(t, cfg, regenMsg()); !ok {
		t.Fatal("must pass")
	}
	args := readFile(t, log)
	for _, want := range []string{"grow", "--check", "--only " + chkID, "--output"} {
		if !strings.Contains(args, want) {
			t.Errorf("sedum was called with %q, want %q", args, want)
		}
	}
}

func TestRegenerationCommitIsRefusedWhenSedumSaysStale(t *testing.T) {
	root, cfg, state, _ := regenRepo(t)
	writeFile(t, state, "stale\n")
	p := baseParts()
	p.pkg = "sha256:p2"
	stage(t, root, map[string]string{"a.go": p.render()})
	if out, ok := check(t, cfg, regenMsg()); ok || !strings.Contains(out, "provenance revise") {
		t.Fatalf("a region sedum does not report current must be refused:\n%s", out)
	}
}

func TestRegenerationCommitIsRefusedWithoutSedum(t *testing.T) {
	root, cfg, _, _ := regenRepo(t)
	t.Setenv("SEDUM_BIN", filepath.Join(t.TempDir(), "no-such-sedum"))
	p := baseParts()
	p.pkg = "sha256:p2"
	stage(t, root, map[string]string{"a.go": p.render()})
	if out, ok := check(t, cfg, regenMsg()); ok || !strings.Contains(out, "provenance revise") {
		t.Fatalf("with no sedum the rule cannot be verified:\n%s", out)
	}
}

func TestRegenerationCommitIsRefusedWhenAMarkerGenOrRecordMoves(t *testing.T) {
	for name, edit := range map[string]func(string) string{
		"gen":    func(s string) string { return strings.Replace(s, `"gen":"sha256:g1"`, `"gen":"sha256:g2"`, 1) },
		"record": func(s string) string { return strings.Replace(s, chkID, "prov-2026-0000c999", 1) },
		"tier":   func(s string) string { return strings.Replace(s, `"tier":"owned"`, `"tier":"shared"`, 1) },
	} {
		t.Run(name, func(t *testing.T) {
			root, cfg, _, _ := regenRepo(t)
			stage(t, root, map[string]string{"a.go": edit(baseParts().render())})
			if out, ok := check(t, cfg, regenMsg()); ok || !strings.Contains(out, "provenance revise") {
				t.Fatalf("a changed marker %s must be refused:\n%s", name, out)
			}
		})
	}
}

func TestRegenerationCommitIsRefusedWhenTextOutsideARegionChanges(t *testing.T) {
	root, cfg, _, _ := regenRepo(t)
	p := baseParts()
	p.pkg, p.outside = "sha256:p2", "// helper, edited by hand\n"
	stage(t, root, map[string]string{"a.go": p.render()})
	if out, ok := check(t, cfg, regenMsg()); ok || !strings.Contains(out, "provenance revise") {
		t.Fatalf("a hand edit beside a re-render must be refused:\n%s", out)
	}
}

func TestRegenerationCommitIsRefusedWhenARegionIsAddedOrRemoved(t *testing.T) {
	for name, edit := range map[string]func(*regenParts){
		"removed": func(p *regenParts) { p.second = false },
		"added": func(p *regenParts) {
			p.body += "// /sedum:addImport\n" + regionHeader("addImport", "sha256:g1", "sha256:p1", `{"code":"os"}`) + "import \"os\"\n"
		},
	} {
		t.Run(name, func(t *testing.T) {
			root, cfg, _, _ := regenRepo(t)
			p := baseParts()
			edit(&p)
			stage(t, root, map[string]string{"a.go": p.render()})
			if out, ok := check(t, cfg, regenMsg()); ok || !strings.Contains(out, "provenance revise") {
				t.Fatalf("a %s region must be refused:\n%s", name, out)
			}
		})
	}
}

func TestRegenerationCommitIsRefusedOnAFileWithoutRegions(t *testing.T) {
	root, cfg, _, _ := regenRepo(t)
	stage(t, root, map[string]string{"spec.sh": "echo changed\n"})
	if out, ok := check(t, cfg, regenMsg()); ok || !strings.Contains(out, "provenance revise") {
		t.Fatalf("a file that holds no region is not a regeneration:\n%s", out)
	}
}

func TestRegenerationCommitIsRefusedWhenTheRecordOrItsSnapshotsMove(t *testing.T) {
	root, cfg, _, _ := regenRepo(t)
	p := baseParts()
	p.pkg = "sha256:p2"
	path := filepath.Join("provenance", chkID+".yml")
	stage(t, root, map[string]string{
		"a.go": p.render(),
		path:   readFile(t, filepath.Join(root, path)) + "# touched\n",
	})
	if out, ok := check(t, cfg, regenMsg()); ok {
		t.Fatalf("a record change beside a re-render must be refused:\n%s", out)
	}
	gitIn(t, root, "reset", "-q", "--hard")
	stage(t, root, map[string]string{
		"a.go": p.render(),
		"provenance/.versions/" + chkID + "/extra.yml": "x: 1\n",
	})
	if out, ok := check(t, cfg, regenMsg()); ok {
		t.Fatalf("a snapshot beside a re-render must be refused:\n%s", out)
	}
}

func TestRegenerationCommitStillNeedsTheCurrentVersionTag(t *testing.T) {
	root, cfg, _, _ := regenRepo(t)
	p := baseParts()
	p.pkg = "sha256:p2"
	stage(t, root, map[string]string{"a.go": p.render()})
	if out, ok := check(t, cfg, "Regenerate a"); ok || !strings.Contains(out, "Commit tag required") {
		t.Fatalf("an untagged regeneration must be refused:\n%s", out)
	}
	if out, ok := check(t, cfg, "Regenerate a ["+chkID+"@2]"); ok {
		t.Fatalf("a tag naming another version must be refused:\n%s", out)
	}
}

func TestRegenerationCommitStaysInsideTheRecordScope(t *testing.T) {
	root, cfg, _, _ := regenRepo(t)
	stage(t, root, map[string]string{"b.go": baseParts().render()})
	gitIn(t, root, "commit", "-q", "-m", "Add b out of scope")
	p := baseParts()
	p.pkg = "sha256:p2"
	stage(t, root, map[string]string{"b.go": p.render()})
	if out, ok := check(t, cfg, regenMsg()); ok || !strings.Contains(out, "b.go") {
		t.Fatalf("a re-render of a file outside the scope must be refused:\n%s", out)
	}
}
