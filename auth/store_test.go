package auth

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestFileStoreModesAndRefusal(t *testing.T) {
	dir := t.TempDir()
	s := FileStore{Path: filepath.Join(dir, "d", "credentials")}
	if _, _, err := s.Load(); !errors.Is(err, ErrNoSession) {
		t.Fatalf("empty load = %v", err)
	}
	if err := s.Save(&Credential{Server: "https://x", RefreshToken: "R"}); err != nil {
		t.Fatal(err)
	}
	fi, _ := os.Stat(s.Path)
	di, _ := os.Stat(filepath.Dir(s.Path))
	if fi.Mode().Perm() != 0o600 || di.Mode().Perm() != 0o700 {
		t.Fatalf("modes file=%o dir=%o", fi.Mode().Perm(), di.Mode().Perm())
	}
	entries, _ := os.ReadDir(filepath.Dir(s.Path))
	if len(entries) != 1 {
		t.Fatalf("temp file left behind: %v", entries)
	}
	os.Chmod(s.Path, 0o644)
	if _, _, err := s.Load(); err == nil {
		t.Fatal("loose mode must be refused")
	}
	os.Chmod(s.Path, 0o600)
	if c, src, err := s.Load(); err != nil || c.RefreshToken != "R" || src != SourceFile {
		t.Fatalf("load = %+v %v %v", c, src, err)
	}
	if err := s.Delete(); err != nil || s.Delete() != nil {
		t.Fatal("delete must be idempotent")
	}
}

type memKeyring struct {
	m    map[string]string
	fail bool
}

func (k *memKeyring) Get(s, u string) (string, error) {
	if k.fail {
		return "", errors.New("no keychain")
	}
	v, ok := k.m[s+u]
	if !ok {
		return "", ErrKeyNotFound
	}
	return v, nil
}

func (k *memKeyring) Set(s, u, v string) error {
	if k.fail {
		return errors.New("no keychain")
	}
	k.m[s+u] = v
	return nil
}

func (k *memKeyring) Delete(s, u string) error { delete(k.m, s+u); return nil }

func TestKeyringStoreFallsBackToFile(t *testing.T) {
	file := FileStore{Path: filepath.Join(t.TempDir(), "credentials")}
	broken := KeyringStore{Keyring: &memKeyring{fail: true}, Fallback: file}
	if err := broken.Save(&Credential{RefreshToken: "R"}); err != nil {
		t.Fatal(err)
	}
	if c, src, err := broken.Load(); err != nil || c.RefreshToken != "R" || src != SourceFile {
		t.Fatalf("fallback load = %+v %v %v", c, src, err)
	}
	good := KeyringStore{Keyring: &memKeyring{m: map[string]string{}}, Fallback: file}
	if err := good.Save(&Credential{RefreshToken: "K"}); err != nil {
		t.Fatal(err)
	}
	if c, src, _ := good.Load(); c.RefreshToken != "K" || src != SourceKeychain {
		t.Fatalf("keychain load = %+v %v", c, src)
	}
	if _, _, err := file.Load(); !errors.Is(err, ErrNoSession) {
		t.Fatal("saving to the keychain must remove the file copy")
	}
}
