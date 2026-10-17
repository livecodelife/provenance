//go:build m6

package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/livecodelife/provenance/auth"
	"github.com/livecodelife/provenance/record"
	"github.com/livecodelife/provenance/remote"
)

// syncBuf is an output buffer the test can read while the command still runs.
type syncBuf struct {
	mu sync.Mutex
	b  bytes.Buffer
}

func (s *syncBuf) Write(p []byte) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.b.Write(p)
}

func (s *syncBuf) String() string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.b.String()
}

// TestM6Joint is the scripted joint check of docs/PLAN-IDENTITY-UI-FIRST.md M6:
// sign in on the CLI, publish a record, edit it in the UI as the same account,
// import the edit as a new version, sign out. scripts/m6.sh sets it up; see
// docs/M6.md. It publishes a record to a real account: use a throwaway one.
func TestM6Joint(t *testing.T) {
	authURL, apiURL := os.Getenv(auth.EnvServer), os.Getenv("PROVISIO_API_URL")
	if authURL == "" || apiURL == "" {
		t.Fatal("set PROVISIO_AUTH_URL and PROVISIO_API_URL (scripts/m6.sh does)")
	}
	// The package's TestMain already points HOME at a temp dir: scripts/m6.sh passes the real one for the UI tooling.
	realHome := os.Getenv("M6_REAL_HOME")
	home := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv(auth.EnvNoKeychain, "1")
	t.Setenv(auth.EnvToken, "")
	var transcript strings.Builder
	record := func(s ...string) { transcript.WriteString(strings.Join(s, "\n")) }
	step := func(name string) { t.Logf("--- %s", name) }
	cli := func(stdin string, args ...string) string {
		t.Helper()
		out, errOut, err := runCLI(t, stdin, args...)
		record(out, errOut)
		if err != nil {
			t.Fatalf("provenance %s: %v\n%s%s", strings.Join(args, " "), err, out, errOut)
		}
		return out
	}

	step("CLI: login --no-browser, approval")
	buf := &syncBuf{}
	login := newRootCmd()
	login.SetOut(buf)
	login.SetErr(buf)
	login.SetArgs([]string{"login", "--no-browser"})
	done := make(chan error, 1)
	go func() { done <- login.Execute() }()
	codeRE := regexp.MustCompile(`[BCDFGHJKLMNPQRSTVWXZ]{4}-[BCDFGHJKLMNPQRSTVWXZ]{4}`)
	var code string
	for deadline := time.Now().Add(30 * time.Second); time.Now().Before(deadline) && code == ""; time.Sleep(200 * time.Millisecond) {
		code = codeRE.FindString(buf.String())
	}
	if code == "" {
		t.Fatalf("no user code in the login output:\n%s", buf.String())
	}
	if approve := os.Getenv("M6_APPROVE_URL"); approve != "" {
		body := map[string]string{"user_code": code}
		if h := os.Getenv("M6_HANDLE"); h != "" {
			body["handle"] = h
		}
		data, _ := json.Marshal(body)
		resp, err := http.Post(approve, "application/json", bytes.NewReader(data))
		if err != nil || resp.StatusCode != 200 {
			t.Fatalf("approving the code: %v %v", err, resp)
		}
		resp.Body.Close()
	} else {
		t.Logf("APPROVE THE SIGN-IN YOURSELF now:\n%s", buf.String())
	}
	select {
	case err := <-done:
		record(buf.String())
		if err != nil {
			t.Fatalf("login: %v\n%s", err, buf.String())
		}
	case <-time.After(11 * time.Minute):
		t.Fatal("the sign-in was not approved in time")
	}
	var who struct{ Handle, DID string }
	if err := json.Unmarshal([]byte(cli("", "whoami", "--json")), &who); err != nil || who.DID == "" {
		t.Fatalf("whoami: %v %+v", err, who)
	}
	t.Logf("signed in as @%s (%s)", who.Handle, who.DID)

	step("CLI: create, open and publish a record")
	root := t.TempDir()
	cfg := filepath.Join(root, ".provenance.yml")
	writeFile(t, cfg, "dir: provenance\nexclude_paths: ['.provenance.yml']\n")
	writeFile(t, filepath.Join(root, "a.go"), "package a\n")
	writeFile(t, filepath.Join(root, "spec.sh"), "true\n")
	gitRepo(t, root)
	created := cli("", "-c", cfg, "create", "--type", "blueprint", "--no-edit", "--title", "M6 joint check "+time.Now().UTC().Format("2006-01-02T15:04:05Z"),
		"--intent", "Throwaway record published by the M6 joint check.", "--scope", "a.go", "--scope", "spec.sh")
	id := regexp.MustCompile(`prov-\d{4}-[0-9a-f]{8}`).FindString(created)
	if id == "" {
		t.Fatalf("no record id in:\n%s", created)
	}
	cli("", "-c", cfg, "edit", id, "--add-spec", "spec.sh:sh")
	cli("", "-c", cfg, "open", "--record", id)
	published := cli("", "-c", cfg, "publish", id, "--yes")
	if !strings.Contains(published, "published  "+id+" v1") {
		t.Fatalf("publish:\n%s", published)
	}
	if again := cli("", "-c", cfg, "publish", id, "--yes"); !strings.Contains(again, "Nothing to publish") {
		t.Fatalf("a second publish must write nothing:\n%s", again)
	}

	step("UI: edit the record as the same account")
	const newIntent, reason = "Intent edited in the UI by the M6 joint check.", "Edited in the UI (M6)"
	if ui := os.Getenv("M6_UI_CMD"); ui != "" {
		if realHome == "" {
			t.Fatal("set M6_REAL_HOME to your real home directory (scripts/m6.sh does)")
		}
		c := exec.CommandContext(context.Background(), "sh", "-c", ui)
		c.Env = append(os.Environ(),
			"HOME="+realHome, "XDG_CONFIG_HOME="+realHome+"/.config", // the UI tooling keeps its own caches; only the CLI gets the empty HOME
			"CYPRESS_M6_RECORD_ID="+id, "CYPRESS_M6_HANDLE="+who.Handle, "CYPRESS_M6_SERVICE_URL="+strings.TrimSuffix(apiURL, "/v1"),
			"CYPRESS_M6_NEW_INTENT="+newIntent, "CYPRESS_M6_REASON="+reason)
		out, err := c.CombinedOutput()
		t.Logf("UI step output:\n%s", out)
		if err != nil {
			t.Fatalf("the UI step failed: %v", err)
		}
	} else {
		t.Log("M6_UI_CMD is not set: the API STAND-IN does the UI step over the wire. This run proves the CLI half and the wire contract, NOT the UI.")
		standInEdit(t, root, id, newIntent, reason, who.DID)
	}

	step("CLI: import brings the UI edit in as version 2")
	imported := cli("", "-c", cfg, "import", "--all")
	if !strings.Contains(imported, id+" v2") || !strings.Contains(imported, "ui)") {
		t.Fatalf("import:\n%s", imported)
	}
	hist := cli("", "-c", cfg, "history", id)
	if !strings.Contains(hist, "v2") || !strings.Contains(hist, who.DID) {
		t.Fatalf("history lacks the UI version by %s:\n%s", who.DID, hist)
	}
	if rec := readFile(t, filepath.Join(root, "provenance", id+".yml")); !strings.Contains(rec, "version: 2") {
		t.Fatalf("the record file is not at version 2:\n%s", rec)
	}
	if second := cli("", "-c", cfg, "import", "--all"); !strings.Contains(second, "Nothing to import") {
		t.Fatalf("a second import must import nothing:\n%s", second)
	}

	step("CLI: logout, the stolen refresh token is refused")
	var cred auth.Credential
	data, err := os.ReadFile(filepath.Join(home, ".provisio", "credentials"))
	if err != nil || json.Unmarshal(data, &cred) != nil || cred.RefreshToken == "" {
		t.Fatalf("credential file: %v", err)
	}
	cli("", "logout")
	body, _ := json.Marshal(map[string]string{"refresh_token": cred.RefreshToken})
	resp, err := http.Post(strings.TrimRight(authURL, "/")+"/auth/refresh", "application/json", bytes.NewReader(body))
	if err != nil || resp.StatusCode != 401 {
		t.Fatalf("a refresh token must be refused after logout: %v %v", err, resp)
	}
	resp.Body.Close()

	step("no secret reached any output")
	for _, s := range cred.Secrets() {
		if strings.Contains(transcript.String(), s) {
			t.Errorf("a token reached the output (%d characters)", len(s))
		}
	}
	fmt.Fprintf(os.Stderr, "M6 joint check passed for %s (%s)\n", id, who.Handle)
}

// standInEdit does what the UI does on the wire when it saves an edit: PUT the
// next version with origin ui, the base version and hash taken from the head,
// a reason, authored by the signed-in account.
func standInEdit(t *testing.T, root, id, intent, reason, did string) {
	t.Helper()
	c, err := apiClient()
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	head, err := c.GetRecord(ctx, id)
	if err != nil {
		t.Fatal(err)
	}
	st := &record.FilesStore{Root: root, Dir: "provenance"}
	r, _, _ := st.Load(id)
	r.Intent = intent
	snap := record.NewSnapshot(r, head.Version+1, head.ContentHash, reason, "", did, time.Now().UTC().Format(time.RFC3339), st.Specs())
	req := remote.VersionRequest(snap, &head.CID, "", "", "")
	req.Origin, req.BaseVersion, req.BaseContentHash = "ui", head.Version, head.ContentHash
	if _, _, err := c.PutVersion(ctx, id, req); err != nil {
		t.Fatal(err)
	}
}
