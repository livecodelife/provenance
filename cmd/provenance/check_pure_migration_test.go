package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// A marker migration (`sedum markers migrate`) rewrites the opening marker line
// of each sedum region and writes one recording per record under
// .sedum/recordings/. The fixtures hold two implemented records, chkID owning
// a.go and migID owning b.go, each file with two regions. The fake sedum lists
// the class of every record from classes.txt in the tree it is pointed at, so
// a test sets the class before and after the commit.

const migID = "prov-2026-0000c002"

const migClasses = ".sedum/recordings/classes.txt"

func migOpening(id string, v2 bool) string {
	head := `// sedum:addImport {"tier":"owned","record":"` + id + `","gen":"sha256:g1","pkg":"sha256:p1",`
	if v2 {
		return head + `"v":2,"kwargs":"sha256:k1","id":{"code":"fmt"},"body":"sha256:b1"}` + "\n"
	}
	return head + `"kwargs":{"code":"fmt"}}` + "\n"
}

func migSecond(id string, v2 bool) string {
	head := `// sedum:setFunction {"tier":"owned","record":"` + id + `","gen":"sha256:g1","pkg":"sha256:p1",`
	if v2 {
		return head + `"v":2,"kwargs":"sha256:k2","id":{"name":"f"},"body":"sha256:b2"}` + "\n"
	}
	return head + `"kwargs":{"name":"f"}}` + "\n"
}

// migFile renders a file of the package with two regions, in v1 or v2 markers.
func migFile(pkg, id string, v2 bool) string {
	return "package " + pkg + "\n\n// sedum:anchor:imports\n" + migOpening(id, v2) + "import \"fmt\"\n// /sedum:addImport\n\n// helper\n" +
		migSecond(id, v2) + "func f() {}\n// /sedum:setFunction\n"
}

func migClassFile(a, b string) string { return a + " " + chkID + "\n" + b + " " + migID + "\n" }

func migRecording(id string) string { return `{"record":"` + id + `"}` + "\n" }

// migSedum writes a sedum that answers `grow --check` over the whole records
// directory and exits 1 when some record is not current, as the real one does.
func migSedum(t *testing.T) (log string) {
	t.Helper()
	dir := t.TempDir()
	log = filepath.Join(dir, "args.log")
	script := "#!/bin/sh\necho \"$@\" >> '" + log + "'\nout=\"\"\nrecs=\"\"\n" +
		"while [ $# -gt 0 ]; do\n  case \"$1\" in --output) out=\"$2\";; --records) recs=\"$2\";; esac\n  shift\ndone\nrc=0\n" +
		"for f in \"$recs\"/prov-*.yml; do\n  id=$(basename \"$f\" .yml)\n  cls=current\n" +
		"  if [ -f \"$out/" + migClasses + "\" ]; then\n    c=$(grep \" $id\\$\" \"$out/" + migClasses + "\" | cut -d' ' -f1)\n    if [ -n \"$c\" ]; then cls=$c; fi\n  fi\n" +
		"  printf '%-10s %s\\n' \"$cls\" \"$id\"\n  if [ \"$cls\" != current ]; then rc=1; fi\ndone\nprintf '\\n2 current, 0 stale\\n'\nexit $rc\n"
	bin := filepath.Join(dir, "sedum")
	writeFile(t, bin, script)
	if err := os.Chmod(bin, 0o755); err != nil {
		t.Fatal(err)
	}
	t.Setenv("SEDUM_BIN", bin)
	return log
}

// migRepo is a git repo whose a.go and b.go hold v1 regions, each owned by an
// implemented record at version 1, with baseClasses as the sedum classes.
func migRepo(t *testing.T, baseClasses string) (root, cfg, log string) {
	t.Helper()
	root = t.TempDir()
	cfg = filepath.Join(root, ".provenance.yml")
	writeFile(t, cfg, "dir: provenance\ncommit_tag_required: true\nexclude_paths: ['README.md']\n")
	for _, r := range []struct{ id, file, spec string }{{chkID, "a.go", "spec.sh"}, {migID, "b.go", "spec2.sh"}} {
		rec := strings.Replace(v2Draft(r.id, "draft"), "affected_scope:\n  - a.go\n", "affected_scope:\n  - "+r.file+"\n  - "+r.spec+"\n", 1)
		rec = strings.Replace(rec, "associated_specs: []\n", "associated_specs:\n  - path: "+r.spec+"\n    run_command: sh\n", 1)
		writeFile(t, filepath.Join(root, "provenance", r.id+".yml"), rec)
		writeFile(t, filepath.Join(root, r.spec), "true\n")
	}
	writeFile(t, filepath.Join(root, "README.md"), "hi\n")
	mustRun(t, "-c", cfg, "open", "--record", chkID)
	mustRun(t, "-c", cfg, "open", "--record", migID)
	gitRepo(t, root)
	log = migSedum(t)
	stage(t, root, map[string]string{"a.go": migFile("a", chkID, false), "b.go": migFile("b", migID, false), migClasses: baseClasses})
	gitIn(t, root, "commit", "-q", "-m", "Grow ["+chkID+"@1]")
	for _, id := range []string{chkID, migID} {
		path := filepath.Join("provenance", id+".yml")
		stage(t, root, map[string]string{path: strings.Replace(readFile(t, filepath.Join(root, path)), "status: open", "status: implemented", 1)})
		gitIn(t, root, "commit", "-q", "-m", "Complete provenance record ["+id+"@1]")
	}
	return root, cfg, log
}

// migChange is the content of a migration commit: both files in v2 markers,
// both recordings and the classes after the commit.
type migChange map[string]string

func migDefault(classes string) migChange {
	return migChange{
		"a.go": migFile("a", chkID, true), "b.go": migFile("b", migID, true),
		".sedum/recordings/" + chkID + ".json": migRecording(chkID),
		".sedum/recordings/" + migID + ".json": migRecording(migID),
		migClasses:                             classes,
	}
}

func migMsg() string { return "Migrate markers [" + chkID + "@1]" }

func allCurrent() string { return migClassFile("current", "current") }

func TestMigrationCommitAcrossTwoOwnersPassesWithOneTagAndNoVersionBump(t *testing.T) {
	root, cfg, _ := migRepo(t, allCurrent())
	stage(t, root, migDefault(allCurrent()))
	if out, ok := check(t, cfg, migMsg()); !ok || !strings.Contains(out, "✓ No forbidden scope violations in staged") {
		t.Fatalf("a pure migration touching two records' files must pass with one tag:\n%s", out)
	}
	gitIn(t, root, "commit", "-q", "-m", migMsg())
	if out, err := run(t, "-c", cfg, "check", "--commit", "HEAD"); err != nil || !strings.Contains(out, "✓") {
		t.Fatalf("--commit must agree: %v\n%s", err, out)
	}
	if out, err := run(t, "-c", cfg, "check", "--range", "HEAD~1..HEAD"); err != nil || !strings.Contains(out, "✓") {
		t.Fatalf("--range must agree: %v\n%s", err, out)
	}
}

func TestMigrationCommitAsksSedumOnTheBaseAndOnTheCommitTreeAndRemovesTheBaseExport(t *testing.T) {
	root, cfg, log := migRepo(t, allCurrent())
	stage(t, root, migDefault(allCurrent()))
	if out, ok := check(t, cfg, migMsg()); !ok {
		t.Fatalf("must pass:\n%s", out)
	}
	var outputs []string
	for _, line := range strings.Split(readFile(t, log), "\n") {
		fields := strings.Fields(line)
		for i, f := range fields {
			if f == "--output" && i+1 < len(fields) {
				outputs = append(outputs, fields[i+1])
			}
		}
		if strings.HasPrefix(line, "grow --check") && strings.Contains(line, "--only") {
			t.Errorf("the class listing covers the whole records directory, not one record: %q", line)
		}
	}
	if len(outputs) < 2 || outputs[0] == outputs[1] {
		t.Fatalf("sedum must be asked about two different trees (base and commit), got %v", outputs)
	}
	for _, o := range outputs {
		if filepath.Clean(o) == filepath.Clean(root) {
			continue
		}
		if _, err := os.Stat(o); err == nil {
			t.Errorf("the exported tree %s was not removed", o)
		}
	}
}

func TestMigrationCommitMayKeepAStaleRecordStale(t *testing.T) {
	stale := migClassFile("stale", "current")
	root, cfg, _ := migRepo(t, stale)
	stage(t, root, migDefault(stale))
	if out, ok := check(t, cfg, migMsg()); !ok {
		t.Fatalf("a record that was stale before and is stale after must not block the migration (sedum exits 1 then):\n%s", out)
	}
}

func TestMigrationCommitIsRefusedWhenARecordChangesClass(t *testing.T) {
	for name, after := range map[string]string{
		"current to stale":  migClassFile("stale", "current"),
		"current to edited": migClassFile("current", "edited"),
	} {
		t.Run(name, func(t *testing.T) {
			root, cfg, _ := migRepo(t, allCurrent())
			stage(t, root, migDefault(after))
			if out, ok := check(t, cfg, migMsg()); ok || !strings.Contains(out, "provenance revise") {
				t.Fatalf("a record whose class changes must be refused:\n%s", out)
			}
		})
	}
	t.Run("stale to current", func(t *testing.T) {
		root, cfg, _ := migRepo(t, migClassFile("stale", "current"))
		stage(t, root, migDefault(allCurrent()))
		if out, ok := check(t, cfg, migMsg()); ok {
			t.Fatalf("a record whose class changes in any direction is not a pure migration:\n%s", out)
		}
	})
}

func TestMigrationCommitIsRefusedWithoutSedum(t *testing.T) {
	root, cfg, _ := migRepo(t, allCurrent())
	stage(t, root, migDefault(allCurrent()))
	t.Setenv("SEDUM_BIN", filepath.Join(t.TempDir(), "no-such-sedum"))
	if out, ok := check(t, cfg, migMsg()); ok || !strings.Contains(out, "provenance revise") {
		t.Fatalf("with no sedum the classes cannot be compared:\n%s", out)
	}
}

func TestMigrationCommitIsRefusedWhenATextInsideARegionChanges(t *testing.T) {
	root, cfg, _ := migRepo(t, allCurrent())
	c := migDefault(allCurrent())
	c["a.go"] = strings.Replace(c["a.go"], "import \"fmt\"\n", "import \"fmt\" // sneaked in\n", 1)
	stage(t, root, c)
	if out, ok := check(t, cfg, migMsg()); ok || !strings.Contains(out, "provenance revise") {
		t.Fatalf("a hand edit inside a region must not ride along with a migration:\n%s", out)
	}
}

func TestMigrationCommitIsRefusedWhenAnyOtherPartOfAFileChanges(t *testing.T) {
	for name, edit := range map[string]func(string) string{
		"outside text": func(s string) string { return strings.Replace(s, "// helper\n", "// helper, edited\n", 1) },
		"text after a region": func(s string) string {
			return strings.Replace(s, "// /sedum:addImport\n", "// /sedum:addImport\n// extra\n", 1)
		},
		"tier":           func(s string) string { return strings.Replace(s, `"tier":"owned"`, `"tier":"shared"`, 1) },
		"record":         func(s string) string { return strings.Replace(s, `"record":"`+chkID+`"`, `"record":"`+migID+`"`, 1) },
		"gen":            func(s string) string { return strings.Replace(s, `"gen":"sha256:g1"`, `"gen":"sha256:g2"`, 1) },
		"pkg":            func(s string) string { return strings.Replace(s, `"pkg":"sha256:p1"`, `"pkg":"sha256:p2"`, 1) },
		"action":         func(s string) string { return strings.Replace(s, "sedum:addImport ", "sedum:setConst ", 1) },
		"indentation":    func(s string) string { return strings.Replace(s, "// sedum:addImport ", "  // sedum:addImport ", 1) },
		"region removed": func(s string) string { return s[:strings.Index(s, "// helper")] },
	} {
		t.Run(name, func(t *testing.T) {
			root, cfg, _ := migRepo(t, allCurrent())
			c := migDefault(allCurrent())
			c["a.go"] = edit(c["a.go"])
			stage(t, root, c)
			if out, ok := check(t, cfg, migMsg()); ok || !strings.Contains(out, "provenance revise") {
				t.Fatalf("a changed %s is not a pure migration:\n%s", name, out)
			}
		})
	}
}

func TestMigrationCommitIsRefusedWhenAnotherFileChangesBesideIt(t *testing.T) {
	for name, extra := range map[string]map[string]string{
		"a file with no markers":       {"spec.sh": "echo changed\n"},
		"a new file with no record":    {"c.go": "package c\n"},
		"a record file":                {filepath.Join("provenance", chkID+".yml"): "touched\n"},
		"a snapshot file":              {"provenance/.versions/" + chkID + "/extra.yml": "x: 1\n"},
		"a file outside recordings":    {".sedum/other.json": "{}\n"},
		"a recording outside its path": {".sedum/recordings/sub/x.json": "{}\n"},
	} {
		t.Run(name, func(t *testing.T) {
			root, cfg, _ := migRepo(t, allCurrent())
			c := migDefault(allCurrent())
			for k, v := range extra {
				c[k] = v
			}
			stage(t, root, c)
			if out, ok := check(t, cfg, migMsg()); ok {
				t.Fatalf("%s beside a migration gets the normal rules:\n%s", name, out)
			}
		})
	}
}

func TestMigrationCommitIsRefusedWhenARecordingIsDeleted(t *testing.T) {
	root, cfg, _ := migRepo(t, allCurrent())
	stage(t, root, map[string]string{".sedum/recordings/" + chkID + ".json": migRecording(chkID)})
	gitIn(t, root, "commit", "-q", "-m", "Add a recording ["+chkID+"@1]")
	c := migDefault(allCurrent())
	delete(c, ".sedum/recordings/"+chkID+".json")
	stage(t, root, c)
	gitIn(t, root, "rm", "-q", "--", ".sedum/recordings/"+chkID+".json")
	if out, ok := check(t, cfg, migMsg()); ok {
		t.Fatalf("a deleted recording is not a migration:\n%s", out)
	}
}

func TestMigrationCommitOfRecordingsAloneGetsTheNormalRules(t *testing.T) {
	root, cfg, _ := migRepo(t, allCurrent())
	stage(t, root, map[string]string{".sedum/recordings/" + chkID + ".json": migRecording(chkID)})
	if out, ok := check(t, cfg, migMsg()); ok {
		t.Fatalf("recordings without a marker change are not a migration:\n%s", out)
	}
}

func TestMigrationCommitStillNeedsATagOfARecordThatOwnsAChangedRegion(t *testing.T) {
	root, cfg, _ := migRepo(t, allCurrent())
	stage(t, root, migDefault(allCurrent()))
	if out, ok := check(t, cfg, "Migrate markers"); ok || !strings.Contains(out, "Commit tag required") {
		t.Fatalf("an untagged migration must be refused:\n%s", out)
	}
	if out, ok := check(t, cfg, "Migrate markers ["+chkID+"@2]"); ok {
		t.Fatalf("a tag naming another version must be refused:\n%s", out)
	}
	if out, ok := check(t, cfg, "Migrate markers [prov-2026-0000c999@1]"); ok {
		t.Fatalf("a tag naming an unknown record must be refused:\n%s", out)
	}
	if out, ok := check(t, cfg, "Migrate markers ["+migID+"@1]"); !ok {
		t.Fatalf("a tag naming either owner at its current version passes:\n%s", out)
	}
}

func TestMigrationCommitIsRefusedWhenTheTaggedRecordOwnsNoChangedRegion(t *testing.T) {
	root, cfg, _ := migRepo(t, allCurrent())
	c := migDefault(allCurrent())
	delete(c, "b.go")
	stage(t, root, c)
	if out, ok := check(t, cfg, "Migrate markers ["+migID+"@1]"); ok || !strings.Contains(out, "provenance revise") {
		t.Fatalf("a tag whose record owns none of the migrated regions gets the normal rules:\n%s", out)
	}
	if out, ok := check(t, cfg, migMsg()); !ok {
		t.Fatalf("the owner of the migrated regions passes:\n%s", out)
	}
}

func TestMigrationCommitPassesWhenTheTaggedRecordIsOpen(t *testing.T) {
	root, cfg, _ := migRepo(t, allCurrent())
	path := filepath.Join("provenance", chkID+".yml")
	stage(t, root, map[string]string{path: strings.Replace(readFile(t, filepath.Join(root, path)), "status: implemented", "status: open", 1)})
	gitIn(t, root, "commit", "-q", "-m", "Reopen ["+chkID+"@1]")
	stage(t, root, migDefault(allCurrent()))
	if out, ok := check(t, cfg, migMsg()); !ok {
		t.Fatalf("the migration passes whatever the status of the tagged record, files of other records included:\n%s", out)
	}
}
