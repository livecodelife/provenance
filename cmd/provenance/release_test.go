package main

import (
	"archive/tar"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/livecodelife/provenance/version"
)

func repoFile(t *testing.T, rel string) string {
	t.Helper()
	return readFile(t, filepath.Join("..", "..", rel))
}

func TestGoreleaserBuildsStaticArchivesAndNeverStampsTheVersion(t *testing.T) {
	cfg := repoFile(t, ".goreleaser.yml")
	for _, want := range []string{
		"main: ./cmd/provenance", "CGO_ENABLED=0", "-trimpath",
		"provenance_{{ .Version }}_{{ .Os }}_{{ .Arch }}", "checksums.txt", "darwin", "linux", "arm64", "amd64",
	} {
		if !strings.Contains(cfg, want) {
			t.Errorf(".goreleaser.yml lacks %q", want)
		}
	}
	for _, line := range strings.Split(cfg, "\n") {
		if strings.HasPrefix(strings.TrimSpace(line), "#") {
			continue
		}
		if strings.Contains(line, "-X") || strings.Contains(line, "main.version") {
			t.Errorf(".goreleaser.yml stamps the version: %q", line)
		}
	}
}

func TestWorkflows(t *testing.T) {
	rel := repoFile(t, ".github/workflows/release.yml")
	for _, want := range []string{
		`- "v*"`, "version/version.go", "go test ./...", "provenance check", "provenance lint", "goreleaser/goreleaser-action", "needs: [version, test]",
	} {
		if !strings.Contains(rel, want) {
			t.Errorf("release.yml lacks %q", want)
		}
	}
	ci := repoFile(t, ".github/workflows/ci.yml")
	for _, want := range []string{"pull_request", "go test ./...", "check --range", "provenance lint"} {
		if !strings.Contains(ci, want) {
			t.Errorf("ci.yml lacks %q", want)
		}
	}
}

// fakeRelease lays out dir the way GitHub serves a release: download/vX/NAME
// and download/vX/checksums.txt, as a file:// releases URL.
func fakeRelease(t *testing.T, bin []byte, sum string) (url string) {
	t.Helper()
	dir := t.TempDir()
	name := fmt.Sprintf("provenance_9.9.9_%s_%s.tar.gz", runtime.GOOS, runtime.GOARCH)
	rel := filepath.Join(dir, "download", "v9.9.9")
	if err := os.MkdirAll(rel, 0o755); err != nil {
		t.Fatal(err)
	}
	f, err := os.Create(filepath.Join(rel, name))
	if err != nil {
		t.Fatal(err)
	}
	gz := gzip.NewWriter(f)
	tw := tar.NewWriter(gz)
	tw.WriteHeader(&tar.Header{Name: "provenance", Mode: 0o755, Size: int64(len(bin))})
	tw.Write(bin)
	tw.Close()
	gz.Close()
	f.Close()
	data, _ := os.ReadFile(f.Name())
	if sum == "" {
		sum = fmt.Sprintf("%x", sha256.Sum256(data))
	}
	writeFile(t, filepath.Join(rel, "checksums.txt"), sum+"  "+name+"\n")
	return "file://" + dir
}

func runInstaller(t *testing.T, releases, version string) (string, string, error) {
	t.Helper()
	bindir := filepath.Join(t.TempDir(), "bin")
	c := exec.Command("sh", filepath.Join("..", "..", "install.sh"))
	c.Env = append(os.Environ(), "PROVENANCE_RELEASES_URL="+releases, "PROVENANCE_BIN_DIR="+bindir, "PROVENANCE_VERSION="+version)
	out, err := c.CombinedOutput()
	return bindir, string(out), err
}

func TestInstallScriptInstallsAVerifiedRelease(t *testing.T) {
	for _, tool := range []string{"sh", "curl", "tar"} {
		if _, err := exec.LookPath(tool); err != nil {
			t.Skipf("%s not found", tool)
		}
	}
	if (runtime.GOOS != "linux" && runtime.GOOS != "darwin") || (runtime.GOARCH != "amd64" && runtime.GOARCH != "arm64") {
		t.Skip("no prebuilt platform here")
	}
	bin := filepath.Join(t.TempDir(), "provenance")
	build := exec.Command("go", "build", "-o", bin, ".")
	build.Env = append(os.Environ(), "GOFLAGS=-modcacherw")
	if out, err := build.CombinedOutput(); err != nil {
		t.Fatalf("go build: %v\n%s", err, out)
	}
	data, _ := os.ReadFile(bin)

	bindir, out, err := runInstaller(t, fakeRelease(t, data, ""), "9.9.9")
	if err != nil {
		t.Fatalf("install failed: %v\n%s", err, out)
	}
	got, err := exec.Command(filepath.Join(bindir, "provenance"), "version").Output()
	if err != nil || string(got) != version.Version+"\n" {
		t.Fatalf("installed binary printed %q (%v)", got, err)
	}

	// A wrong checksum installs nothing.
	bindir, out, err = runInstaller(t, fakeRelease(t, data, strings.Repeat("0", 64)), "9.9.9")
	if err == nil || !strings.Contains(out, "checksum mismatch") {
		t.Fatalf("a bad checksum must fail: %v\n%s", err, out)
	}
	if _, statErr := os.Stat(filepath.Join(bindir, "provenance")); statErr == nil {
		t.Fatal("a binary was installed despite the checksum mismatch")
	}

	// A version that was never released fails instead of building something else.
	_, out, err = runInstaller(t, fakeRelease(t, data, ""), "8.8.8")
	if err == nil || !strings.Contains(out, "8.8.8 is not available") {
		t.Fatalf("an unknown version must fail: %v\n%s", err, out)
	}
}
