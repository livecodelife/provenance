// Package remotetest is an in-memory record server built from
// provisio-backend docs/RECORDS-API.md, for tests of the remote package and of
// the commands that use it. It recomputes every content hash with the record
// package, enforces compare-and-swap on the head and the version chain, and
// answers 409 with the current head.
package remotetest

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sort"
	"strconv"
	"strings"
	"sync"
	"testing"

	"github.com/livecodelife/provenance/record"
)

type version struct {
	record.PublishedVersion
	URI string `json:"uri"`
	CID string `json:"cid"`
}

type head struct {
	record.PublishedHead
	ID  string `json:"id"`
	URI string `json:"uri"`
	CID string `json:"cid"`
}

type entry struct {
	head     head
	versions []version
}

// Server is the fake record server.
type Server struct {
	srv *httptest.Server

	mu      sync.Mutex
	records map[string]*entry
	counter int

	// Accept decides whether a bearer token is valid; nil accepts any non-empty one.
	Accept func(bearer string) bool
	// Writes counts accepted PUTs to versions, HeadWrites accepted PUTs to heads.
	Writes, HeadWrites int
}

// New starts the server; the test closes it.
func New(t *testing.T) *Server {
	s := &Server{records: map[string]*entry{}}
	s.srv = httptest.NewServer(http.HandlerFunc(s.serve))
	t.Cleanup(s.srv.Close)
	return s
}

// URL is the API base (it ends in /v1, like the real one).
func (s *Server) URL() string { return s.srv.URL + "/v1" }

// Close stops the server.
func (s *Server) Close() { s.srv.Close() }

func reply(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func fail(w http.ResponseWriter, status int, code, msg string, details any) {
	e := map[string]any{"code": code, "message": msg}
	if details != nil {
		e["details"] = details
	}
	reply(w, status, map[string]any{"error": e})
}

func (s *Server) conflict(w http.ResponseWriter, e *entry) {
	var h any
	if e != nil {
		h = e.head
	}
	fail(w, 409, "conflict", "The record head moved", map[string]any{"head": h})
}

func (s *Server) serve(w http.ResponseWriter, r *http.Request) {
	bearer := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	if bearer == "" || (s.Accept != nil && !s.Accept(bearer)) {
		fail(w, 401, "invalid_token", "bad token", nil)
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	path := strings.TrimPrefix(r.URL.Path, "/v1")
	parts := strings.Split(strings.Trim(path, "/"), "/")
	switch {
	case r.Method == "GET" && path == "/records":
		heads := []head{}
		ids := make([]string, 0, len(s.records))
		for id := range s.records {
			ids = append(ids, id)
		}
		sort.Strings(ids)
		for _, id := range ids {
			heads = append(heads, s.records[id].head)
		}
		reply(w, 200, map[string]any{"records": heads})
	case len(parts) == 2 && parts[0] == "records" && r.Method == "GET":
		e := s.records[parts[1]]
		if e == nil {
			fail(w, 404, "not_found", "no such record", nil)
			return
		}
		reply(w, 200, map[string]any{"head": e.head, "latest": e.versions[len(e.versions)-1]})
	case len(parts) == 3 && parts[0] == "records" && parts[2] == "versions" && r.Method == "GET":
		e := s.records[parts[1]]
		if e == nil {
			fail(w, 404, "not_found", "no such record", nil)
			return
		}
		reply(w, 200, map[string]any{"versions": e.versions})
	case len(parts) == 3 && parts[0] == "records" && parts[2] == "versions" && r.Method == "PUT":
		s.putVersion(w, r, parts[1])
	case len(parts) == 3 && parts[0] == "records" && parts[2] == "head" && r.Method == "PUT":
		s.putHead(w, r, parts[1])
	default:
		fail(w, 404, "not_found", "no such route", nil)
	}
}

func rank(s string) int {
	return map[string]int{"draft": 0, "open": 1, "implemented": 2, "deprecated": 3}[s]
}

func (s *Server) nextCID() string {
	s.counter++
	return "bafy" + strconv.Itoa(s.counter)
}

func (s *Server) putVersion(w http.ResponseWriter, r *http.Request, id string) {
	var raw map[string]json.RawMessage
	var body struct {
		ExpectedHeadCid *string                 `json:"expectedHeadCid"`
		Version         int                     `json:"version"`
		ContentHash     string                  `json:"contentHash"`
		GenHash         string                  `json:"genHash"`
		SpecHashes      []string                `json:"specHashes"`
		Content         record.PublishedContent `json:"content"`
		Status          string                  `json:"status"`
		Reason          string                  `json:"reason"`
		ReasonKind      string                  `json:"reasonKind"`
		Author          string                  `json:"author"`
		Commit          string                  `json:"commit"`
		Project         string                  `json:"project"`
		Source          *record.PublishedSource `json:"source"`
		Origin          string                  `json:"origin"`
		BaseVersion     int                     `json:"baseVersion"`
		BaseHash        string                  `json:"baseContentHash"`
	}
	buf := new(strings.Builder)
	if _, err := bufCopy(buf, r); err != nil {
		fail(w, 400, "invalid_record", "unreadable body", nil)
		return
	}
	if json.Unmarshal([]byte(buf.String()), &raw) != nil || json.Unmarshal([]byte(buf.String()), &body) != nil {
		fail(w, 400, "invalid_record", "malformed body", nil)
		return
	}
	if body.Content.Title == "" {
		fail(w, 400, "invalid_record", "content.title is required", nil)
		return
	}
	if len(body.SpecHashes) != len(body.Content.AssociatedSpecs) {
		fail(w, 400, "hash_mismatch", "specHashes has the wrong length", nil)
		return
	}
	rec := body.Content.Record()
	rec.ID = id
	ptrs := make([]*string, len(body.SpecHashes))
	for i := range body.SpecHashes {
		if body.SpecHashes[i] != "" {
			h := body.SpecHashes[i]
			ptrs[i] = &h
		}
	}
	if got := record.ContentHashWithSpecs(rec, ptrs); got != body.ContentHash {
		fail(w, 400, "hash_mismatch", "contentHash differs from the server's recomputation", nil)
		return
	}
	if body.GenHash != "" && body.GenHash != rec.GenHash() {
		fail(w, 400, "hash_mismatch", "genHash differs from the server's recomputation", nil)
		return
	}
	e := s.records[id]
	if _, present := raw["expectedHeadCid"]; present {
		if body.ExpectedHeadCid == nil && e != nil || body.ExpectedHeadCid != nil && (e == nil || e.head.CID != *body.ExpectedHeadCid) {
			s.conflict(w, e)
			return
		}
	}
	want := 1
	if e != nil {
		want = e.head.Version + 1
	}
	if body.Version != want {
		fail(w, 400, "bad_version", "version is not head + 1", map[string]any{"expected": want})
		return
	}
	if want > 1 && body.Reason == "" {
		fail(w, 400, "invalid_record", "reason is required after version 1", nil)
		return
	}
	if body.Origin == "ui" && body.BaseVersion != want-1 {
		s.conflict(w, e)
		return
	}
	status := body.Status
	if status == "" {
		status = "draft"
		if e != nil {
			status = e.head.Status
		}
	}
	uri := "at://did:plc:test/dev.provisio.provenance.record/" + id
	v := version{PublishedVersion: record.PublishedVersion{
		Record: uri, RecordID: id, Version: want, ContentHash: body.ContentHash, Reason: body.Reason,
		ReasonKind: body.ReasonKind, Author: body.Author, CreatedAt: "2026-10-16T00:00:00Z", Commit: body.Commit,
		SpecHashes: body.SpecHashes, Content: body.Content, Origin: body.Origin,
		BaseVersion: body.BaseVersion, BaseContentHash: body.BaseHash,
	}, URI: uri + "/v" + strconv.Itoa(want), CID: s.nextCID()}
	if e != nil {
		v.ParentHash = e.versions[len(e.versions)-1].ContentHash
		v.Parent = &record.StrongRef{URI: e.versions[len(e.versions)-1].URI, CID: e.versions[len(e.versions)-1].CID}
	} else {
		e = &entry{}
		s.records[id] = e
	}
	e.versions = append(e.versions, v)
	e.head = head{PublishedHead: record.PublishedHead{
		Status: status, Version: want, ContentHash: body.ContentHash, GenHash: rec.GenHash(),
		Latest: record.StrongRef{URI: v.URI, CID: v.CID}, Content: body.Content, Project: body.Project, Source: body.Source,
		UpdatedAt: "2026-10-16T00:00:00Z",
	}, ID: id, URI: uri, CID: s.nextCID()}
	s.Writes++
	reply(w, 201, map[string]any{"head": e.head, "version": v})
}

func (s *Server) putHead(w http.ResponseWriter, r *http.Request, id string) {
	var body struct {
		ExpectedHeadCid string                  `json:"expectedHeadCid"`
		Version         int                     `json:"version"`
		Status          string                  `json:"status"`
		Source          *record.PublishedSource `json:"source"`
		Project         string                  `json:"project"`
	}
	if json.NewDecoder(r.Body).Decode(&body) != nil {
		fail(w, 400, "invalid_record", "malformed body", nil)
		return
	}
	e := s.records[id]
	if e == nil {
		fail(w, 404, "not_found", "no such record", nil)
		return
	}
	if body.ExpectedHeadCid != e.head.CID {
		s.conflict(w, e)
		return
	}
	if body.Version != e.head.Version {
		fail(w, 400, "bad_version", "version must equal the head's", map[string]any{"expected": e.head.Version})
		return
	}
	if body.Status != "" {
		if rank(body.Status) <= rank(e.head.Status) {
			fail(w, 400, "bad_status", "status may only move forward", nil)
			return
		}
		e.head.Status = body.Status
	}
	if body.Source != nil {
		e.head.Source = body.Source
	}
	if body.Project != "" {
		e.head.Project = body.Project
	}
	e.head.CID = s.nextCID()
	s.HeadWrites++
	reply(w, 200, map[string]any{"head": e.head})
}

// Source returns the source the head of id carries, nil when there is none.
func (s *Server) Source(id string) *record.PublishedSource {
	s.mu.Lock()
	defer s.mu.Unlock()
	if e := s.records[id]; e != nil {
		return e.head.Source
	}
	return nil
}

// Project returns the project the head of id carries, "" when there is none.
func (s *Server) Project(id string) string {
	s.mu.Lock()
	defer s.mu.Unlock()
	if e := s.records[id]; e != nil {
		return e.head.Project
	}
	return ""
}

// Versions returns the stored versions of id, for assertions.
func (s *Server) Versions(id string) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	if e := s.records[id]; e != nil {
		return len(e.versions)
	}
	return 0
}

// Status returns the head status of id, "" when unknown.
func (s *Server) Status(id string) string {
	s.mu.Lock()
	defer s.mu.Unlock()
	if e := s.records[id]; e != nil {
		return e.head.Status
	}
	return ""
}

func bufCopy(b *strings.Builder, r *http.Request) (int64, error) {
	buf := make([]byte, 32<<10)
	var n int64
	for {
		m, err := r.Body.Read(buf)
		b.Write(buf[:m])
		n += int64(m)
		if err != nil {
			if err.Error() == "EOF" {
				return n, nil
			}
			return n, err
		}
	}
}
