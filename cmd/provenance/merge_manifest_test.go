package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"testing"

	"github.com/livecodelife/provenance/record"
)

func mline(id string, version int, hash string) record.ManifestLine {
	return record.ManifestLine{ID: id, Version: version, ContentHash: "sha256:" + strings.Repeat(hash, 64)}
}

func mtext(lines ...record.ManifestLine) string {
	return string(record.MarshalManifest(lines))
}

func TestMergeManifestsPerRecord(t *testing.T) {
	a1, b1, c1 := mline("prov-2026-0000000a", 1, "a"), mline("prov-2026-0000000b", 1, "b"), mline("prov-2026-0000000c", 1, "c")
	a2, c2 := mline("prov-2026-0000000a", 2, "d"), mline("prov-2026-0000000c", 2, "e")
	a3 := mline("prov-2026-0000000a", 3, "f")
	d1, e1 := mline("prov-2026-0000000d", 1, "1"), mline("prov-2026-0000000e", 1, "2")
	base := []record.ManifestLine{a1, b1, c1}

	cases := []struct {
		name         string
		ours, theirs []record.ManifestLine
		want         []record.ManifestLine
		conflict     string
	}{
		{"different records revised", []record.ManifestLine{a2, b1, c1}, []record.ManifestLine{a1, b1, c2}, []record.ManifestLine{a2, b1, c2}, ""},
		{"same record, higher version wins", []record.ManifestLine{a2, b1, c1}, []record.ManifestLine{a3, b1, c1}, []record.ManifestLine{a3, b1, c1}, ""},
		{"same record, same version, same content", []record.ManifestLine{a2, b1, c1}, []record.ManifestLine{a2, b1, c1}, []record.ManifestLine{a2, b1, c1}, ""},
		{"both add different records", []record.ManifestLine{a1, b1, c1, d1}, []record.ManifestLine{a1, b1, c1, e1}, []record.ManifestLine{a1, b1, c1, d1, e1}, ""},
		{"one side removes an untouched record", []record.ManifestLine{a1, c1}, []record.ManifestLine{a1, b1, c1}, []record.ManifestLine{a1, c1}, ""},
		{"same record, same version, different hash", []record.ManifestLine{mline("prov-2026-0000000a", 2, "d"), b1, c1}, []record.ManifestLine{mline("prov-2026-0000000a", 2, "9"), b1, c1}, nil, "prov-2026-0000000a"},
		{"one side removes, the other revises", []record.ManifestLine{b1, c1}, []record.ManifestLine{a2, b1, c1}, nil, "prov-2026-0000000a"},
	}
	for _, c := range cases {
		got, err := mergeManifests(base, c.ours, c.theirs)
		if c.conflict != "" {
			if err == nil || !strings.Contains(err.Error(), c.conflict) {
				t.Errorf("%s: want a conflict naming %s, got %v", c.name, c.conflict, err)
			}
			continue
		}
		if err != nil {
			t.Errorf("%s: %v", c.name, err)
			continue
		}
		if mtext(got...) != mtext(c.want...) {
			t.Errorf("%s:\n%s\nwant\n%s", c.name, mtext(got...), mtext(c.want...))
		}
	}
}

var (
	mergeBinOnce sync.Once
	mergeBin     string
	mergeBinErr  error
)

func mergeDriverBin(t *testing.T) string {
	t.Helper()
	mergeBinOnce.Do(func() {
		mergeBin = filepath.Join(testTemp, "bin", "provenance")
		build := exec.Command("go", "build", "-o", mergeBin, ".")
		build.Env = append(os.Environ(), "GOFLAGS=-modcacherw")
		if out, err := build.CombinedOutput(); err != nil {
			mergeBinErr = fmt.Errorf("build: %v\n%s", err, out)
		}
	})
	if mergeBinErr != nil {
		t.Fatal(mergeBinErr)
	}
	return mergeBin
}

// gitTry runs git in dir and returns its output and error, for merges that
// are expected to stop.
func gitTry(t *testing.T, dir string, args ...string) (string, error) {
	t.Helper()
	c := exec.Command("git", append([]string{"-C", dir}, args...)...)
	out, err := c.CombinedOutput()
	return string(out), err
}

func gitMust(t *testing.T, dir string, args ...string) string {
	t.Helper()
	out, err := gitTry(t, dir, args...)
	if err != nil {
		t.Fatalf("git %s: %v\n%s", strings.Join(args, " "), err, out)
	}
	return out
}

// mergeRepo is a repository whose manifest holds base on main, with the
// provenance merge driver configured to the built binary.
func mergeRepo(t *testing.T, base string) string {
	t.Helper()
	root := t.TempDir()
	gitMust(t, root, "init", "-q")
	gitMust(t, root, "config", "merge.provenance.driver", mergeDriverBin(t)+" merge-manifest %O %A %B")
	writeFile(t, filepath.Join(root, ".gitattributes"), manifestMergeAttr+"\n")
	writeFile(t, filepath.Join(root, manifestPath), base)
	gitMust(t, root, "add", "-A")
	gitMust(t, root, "commit", "-q", "-m", "base")
	return root
}

// branchCommit writes the manifest on a new branch off main and commits it.
func branchCommit(t *testing.T, root, branch, manifest string) {
	t.Helper()
	gitMust(t, root, "checkout", "-q", "-b", branch, "main")
	writeFile(t, filepath.Join(root, manifestPath), manifest)
	gitMust(t, root, "commit", "-q", "-am", branch)
}

func TestManifestMergeDriverMergesRealBranches(t *testing.T) {
	a1, b1, c1 := mline("prov-2026-0000000a", 1, "a"), mline("prov-2026-0000000b", 1, "b"), mline("prov-2026-0000000c", 1, "c")
	a2, b2 := mline("prov-2026-0000000a", 2, "d"), mline("prov-2026-0000000b", 2, "e")
	a3 := mline("prov-2026-0000000a", 3, "f")
	base := mtext(a1, b1, c1)
	manifest := func(root string) string { return readFile(t, filepath.Join(root, manifestPath)) }

	t.Run("different records", func(t *testing.T) {
		root := mergeRepo(t, base)
		branchCommit(t, root, "ours", mtext(a2, b1, c1))
		// Adjacent lines: a plain text merge would conflict here.
		branchCommit(t, root, "theirs", mtext(a1, b2, c1))
		gitMust(t, root, "checkout", "-q", "ours")
		gitMust(t, root, "merge", "--no-edit", "theirs")
		if got, want := manifest(root), mtext(a2, b2, c1); got != want {
			t.Fatalf("merged manifest:\n%s\nwant\n%s", got, want)
		}
	})

	t.Run("same record, different versions", func(t *testing.T) {
		root := mergeRepo(t, base)
		branchCommit(t, root, "ours", mtext(a2, b1, c1))
		branchCommit(t, root, "theirs", mtext(a3, b1, c1))
		gitMust(t, root, "checkout", "-q", "ours")
		gitMust(t, root, "merge", "--no-edit", "theirs")
		if got, want := manifest(root), mtext(a3, b1, c1); got != want {
			t.Fatalf("the higher version should win:\n%s\nwant\n%s", got, want)
		}
	})

	t.Run("same record, same version, different content", func(t *testing.T) {
		root := mergeRepo(t, base)
		branchCommit(t, root, "ours", mtext(a2, b1, c1))
		branchCommit(t, root, "theirs", mtext(mline("prov-2026-0000000a", 2, "9"), b1, c1))
		gitMust(t, root, "checkout", "-q", "ours")
		if out, err := gitTry(t, root, "merge", "--no-edit", "theirs"); err == nil {
			t.Fatalf("a real conflict merged cleanly:\n%s", out)
		}
		if status := gitMust(t, root, "status", "--porcelain"); !strings.Contains(status, "UU "+manifestPath) {
			t.Fatalf("the manifest should be left conflicted:\n%s", status)
		}
	})
}
