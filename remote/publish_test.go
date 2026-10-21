package remote

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/livecodelife/provenance/record"
	"github.com/livecodelife/provenance/remote/remotetest"
)

func TestPublishNewRecordThenNothing(t *testing.T) {
	c, srv := newClient(t)
	st := localStore(t, 3)
	ctx := context.Background()
	plan, err := PlanPublish(ctx, st, []string{testID}, c)
	if err != nil || plan.Empty() || plan.Versions() != 3 || plan.Items[0].Snaps[0].Version != 1 {
		t.Fatalf("plan = %+v %v", plan, err)
	}
	var seen []int
	if err := ApplyPublish(ctx, plan, c, "", nil, func(_ string, v int) { seen = append(seen, v) }); err != nil {
		t.Fatal(err)
	}
	if srv.Versions(testID) != 3 || len(seen) != 3 || seen[0] != 1 || seen[2] != 3 || srv.Status(testID) != "open" {
		t.Fatalf("versions=%d seen=%v status=%s", srv.Versions(testID), seen, srv.Status(testID))
	}
	writes := srv.Writes
	again, err := PlanPublish(ctx, st, []string{testID}, c)
	if err != nil || !again.Empty() {
		t.Fatalf("second plan must be empty: %+v %v", again, err)
	}
	if err := ApplyPublish(ctx, again, c, "", nil, nil); err != nil || srv.Writes != writes || srv.HeadWrites != 0 {
		t.Fatalf("an empty plan wrote something: %v", err)
	}
}

func TestPublishOnlyNewVersions(t *testing.T) {
	c, srv := newClient(t)
	st := localStore(t, 2)
	ctx := context.Background()
	plan, _ := PlanPublish(ctx, st, []string{testID}, c)
	if err := ApplyPublish(ctx, plan, c, "", nil, nil); err != nil {
		t.Fatal(err)
	}
	// A third version locally.
	r, doc, _ := st.Load(testID)
	r.Intent = "Do the thing, take 9."
	if _, err := commitMore(st, r, doc); err != nil {
		t.Fatal(err)
	}
	plan, err := PlanPublish(ctx, st, []string{testID}, c)
	if err != nil || plan.Versions() != 1 || plan.Items[0].Snaps[0].Version != 3 {
		t.Fatalf("plan = %+v %v", plan, err)
	}
	if err := ApplyPublish(ctx, plan, c, "", func(id string, v int) string { return "abc123" }, nil); err != nil || srv.Versions(testID) != 3 {
		t.Fatalf("apply: %v versions=%d", err, srv.Versions(testID))
	}
}

func TestPublishHeadStatusOnly(t *testing.T) {
	c, srv := newClient(t)
	st := localStore(t, 1)
	ctx := context.Background()
	plan, _ := PlanPublish(ctx, st, []string{testID}, c)
	if err := ApplyPublish(ctx, plan, c, "", nil, nil); err != nil {
		t.Fatal(err)
	}
	r, doc, _ := st.Load(testID)
	r.Status = "implemented"
	if err := st.Save(r, doc); err != nil {
		t.Fatal(err)
	}
	plan, err := PlanPublish(ctx, st, []string{testID}, c)
	if err != nil || plan.Versions() != 0 || len(plan.Items) != 1 || plan.Items[0].HeadStatus != "implemented" {
		t.Fatalf("plan = %+v %v", plan, err)
	}
	if err := ApplyPublish(ctx, plan, c, "", nil, nil); err != nil || srv.Status(testID) != "implemented" || srv.Versions(testID) != 1 || srv.HeadWrites != 1 {
		t.Fatalf("apply: %v status=%s versions=%d headWrites=%d", err, srv.Status(testID), srv.Versions(testID), srv.HeadWrites)
	}
	if again, _ := PlanPublish(ctx, st, []string{testID}, c); !again.Empty() {
		t.Fatal("status in sync must plan nothing")
	}
}

func TestPublishDraftIsSkipped(t *testing.T) {
	c, _ := newClient(t)
	st := localStore(t, 0)
	plan, err := PlanPublish(context.Background(), st, []string{testID}, c)
	if err != nil || !plan.Empty() || len(plan.Skipped) != 1 {
		t.Fatalf("plan = %+v %v", plan, err)
	}
}

func TestPublishReportsDivergenceAndKeepsGoing(t *testing.T) {
	c, _ := newClient(t)
	ctx := context.Background()
	a := localStore(t, 2)
	plan, _ := PlanPublish(ctx, a, []string{testID}, c)
	if err := ApplyPublish(ctx, plan, c, "", nil, nil); err != nil {
		t.Fatal(err)
	}
	// Another clone with a different version 2.
	b := localStore(t, 1)
	r, doc, _ := b.Load(testID)
	r.Intent = "A different second version."
	if _, err := commitMore(b, r, doc); err != nil {
		t.Fatal(err)
	}
	plan, err := PlanPublish(ctx, b, []string{testID}, c)
	if err != nil || len(plan.Problems) != 1 || plan.Problems[0].Kind != "diverged" || !strings.Contains(plan.Problems[0].Reason, "differs") || !plan.Empty() {
		t.Fatalf("plan = %+v err = %v", plan, err)
	}
	// And one that is behind the server.
	short := localStore(t, 1)
	plan, err = PlanPublish(ctx, short, []string{testID}, c)
	if err != nil || len(plan.Problems) != 1 || plan.Problems[0].Kind != "server ahead" || !strings.Contains(plan.Problems[0].Reason, "ahead") {
		t.Fatalf("plan = %+v err = %v", plan, err)
	}
	// A record the server already has in full is up to date, not a problem.
	plan, err = PlanPublish(ctx, a, []string{testID}, c)
	if err != nil || len(plan.UpToDate) != 1 || plan.UpToDate[0] != testID || len(plan.Problems) != 0 {
		t.Fatalf("plan = %+v err = %v", plan, err)
	}
}

func TestPublishSendsTheItemSource(t *testing.T) {
	c, srv := newClient(t)
	ctx := context.Background()
	st := localStore(t, 1)
	src := &record.PublishedSource{Repo: "https://github.com/acme/widgets", Path: "provenance/" + testID + ".yml", Commit: "abc123"}
	plan, _ := PlanPublish(ctx, st, []string{testID}, c)
	plan.Items[0].Source = src
	if err := ApplyPublish(ctx, plan, c, "", nil, nil); err != nil {
		t.Fatal(err)
	}
	if got := srv.Source(testID); got == nil || *got != *src {
		t.Fatalf("head source = %+v", got)
	}
}

func TestPublishConflictReportsHead(t *testing.T) {
	c, srv := newClient(t)
	ctx := context.Background()
	st := localStore(t, 1)
	plan, _ := PlanPublish(ctx, st, []string{testID}, c)
	// Someone else creates the record between plan and apply.
	other := localStore(t, 1)
	otherPlan, _ := PlanPublish(ctx, other, []string{testID}, c)
	if err := ApplyPublish(ctx, otherPlan, c, "", nil, nil); err != nil {
		t.Fatal(err)
	}
	err := ApplyPublish(ctx, plan, c, "", nil, nil)
	var ce *ConflictError
	if !errors.As(err, &ce) || ce.Head == nil || ce.Head.Version != 1 || srv.Versions(testID) != 1 {
		t.Fatalf("err = %v", err)
	}
}

const wantProject = "at://did:plc:test/dev.provisio.project/3kabc"

func publishedOnce(t *testing.T) (*Client, *remotetest.Server, *record.FilesStore) {
	t.Helper()
	c, srv := newClient(t)
	st := localStore(t, 2)
	plan, err := PlanPublish(context.Background(), st, []string{testID}, c)
	if err != nil {
		t.Fatal(err)
	}
	if err := ApplyPublish(context.Background(), plan, c, "", nil, nil); err != nil {
		t.Fatal(err)
	}
	return c, srv, st
}

func TestPublishSetsTheProjectOnAnUnchangedRecord(t *testing.T) {
	c, srv, st := publishedOnce(t)
	ctx := context.Background()
	writes := srv.Writes
	want := HeadWant{Project: wantProject}
	plan, err := PlanPublishWith(ctx, st, []string{testID}, c, want)
	if err != nil || plan.Empty() || len(plan.UpToDate) != 0 || len(plan.Items) != 1 {
		t.Fatalf("plan = %+v %v", plan, err)
	}
	it := plan.Items[0]
	if it.SetProject != wantProject || len(it.Snaps) != 0 || it.HeadStatus != "" || it.SetSource != nil {
		t.Fatalf("item = %+v", it)
	}
	if err := ApplyPublish(ctx, plan, c, wantProject, nil, nil); err != nil {
		t.Fatal(err)
	}
	if srv.Project(testID) != wantProject || srv.HeadWrites != 1 || srv.Writes != writes || srv.Versions(testID) != 2 {
		t.Fatalf("project=%q headWrites=%d writes %d→%d versions=%d", srv.Project(testID), srv.HeadWrites, writes, srv.Writes, srv.Versions(testID))
	}
	again, err := PlanPublishWith(ctx, st, []string{testID}, c, want)
	if err != nil || !again.Empty() || len(again.UpToDate) != 1 {
		t.Fatalf("a matching project must be up to date: %+v %v", again, err)
	}
	// Without a wanted project the head's is left alone.
	none, err := PlanPublishWith(ctx, st, []string{testID}, c, HeadWant{})
	if err != nil || !none.Empty() || len(none.UpToDate) != 1 {
		t.Fatalf("no wanted project: %+v %v", none, err)
	}
}

func TestPublishChangesTheProject(t *testing.T) {
	c, srv, st := publishedOnce(t)
	ctx := context.Background()
	for _, p := range []string{wantProject, wantProject + "2"} {
		plan, err := PlanPublishWith(ctx, st, []string{testID}, c, HeadWant{Project: p})
		if err != nil || len(plan.Items) != 1 || plan.Items[0].SetProject != p {
			t.Fatalf("plan for %s = %+v %v", p, plan, err)
		}
		if err := ApplyPublish(ctx, plan, c, p, nil, nil); err != nil || srv.Project(testID) != p {
			t.Fatalf("apply %s: %v project=%q", p, err, srv.Project(testID))
		}
	}
}

func TestPublishSetsTheSourceWhenRepoOrPathDiffer(t *testing.T) {
	c, srv, st := publishedOnce(t)
	ctx := context.Background()
	src := &record.PublishedSource{Repo: "https://github.com/acme/widgets", Path: "provenance/" + testID + ".yml", Commit: "aaa111"}
	want := HeadWant{Source: func(string) *record.PublishedSource { return src }}
	plan, err := PlanPublishWith(ctx, st, []string{testID}, c, want)
	if err != nil || len(plan.Items) != 1 || plan.Items[0].SetSource == nil || plan.Items[0].SetProject != "" {
		t.Fatalf("plan = %+v %v", plan, err)
	}
	if plan.Items[0].Source == nil || *plan.Items[0].Source != *src {
		t.Fatalf("the item must carry the source it writes: %+v", plan.Items[0].Source)
	}
	if err := ApplyPublish(ctx, plan, c, "", nil, nil); err != nil || srv.Source(testID) == nil || *srv.Source(testID) != *src {
		t.Fatalf("apply: %v source=%+v", err, srv.Source(testID))
	}
	// Only the commit moved: nothing to write.
	moved := *src
	moved.Commit = "bbb222"
	again, err := PlanPublishWith(ctx, st, []string{testID}, c, HeadWant{Source: func(string) *record.PublishedSource { return &moved }})
	if err != nil || !again.Empty() || len(again.UpToDate) != 1 {
		t.Fatalf("a new commit alone must stay up to date: %+v %v", again, err)
	}
	// Another path is a change.
	moved.Path = "elsewhere/" + testID + ".yml"
	again, err = PlanPublishWith(ctx, st, []string{testID}, c, HeadWant{Source: func(string) *record.PublishedSource { return &moved }})
	if err != nil || len(again.Items) != 1 || again.Items[0].SetSource == nil {
		t.Fatalf("a changed path must be planned: %+v %v", again, err)
	}
}

func TestPublishProjectRidesAlongWithAStatusMove(t *testing.T) {
	c, srv, st := publishedOnce(t)
	ctx := context.Background()
	r, doc, _ := st.Load(testID)
	r.Status = "implemented"
	if err := st.Save(r, doc); err != nil {
		t.Fatal(err)
	}
	plan, err := PlanPublishWith(ctx, st, []string{testID}, c, HeadWant{Project: wantProject})
	if err != nil || len(plan.Items) != 1 || plan.Items[0].HeadStatus != "implemented" || plan.Items[0].SetProject != wantProject {
		t.Fatalf("plan = %+v %v", plan, err)
	}
	if err := ApplyPublish(ctx, plan, c, wantProject, nil, nil); err != nil || srv.HeadWrites != 1 || srv.Status(testID) != "implemented" || srv.Project(testID) != wantProject {
		t.Fatalf("apply: %v headWrites=%d status=%s project=%q", err, srv.HeadWrites, srv.Status(testID), srv.Project(testID))
	}
}

func TestPublishNewVersionsNeedNoHeadUpdateForTheProject(t *testing.T) {
	c, srv := newClient(t)
	st := localStore(t, 2)
	ctx := context.Background()
	plan, err := PlanPublishWith(ctx, st, []string{testID}, c, HeadWant{Project: wantProject})
	if err != nil || len(plan.Items) != 1 || plan.Items[0].SetProject != "" {
		t.Fatalf("plan = %+v %v", plan, err)
	}
	if err := ApplyPublish(ctx, plan, c, wantProject, nil, nil); err != nil || srv.Project(testID) != wantProject || srv.HeadWrites != 0 {
		t.Fatalf("apply: %v project=%q headWrites=%d", err, srv.Project(testID), srv.HeadWrites)
	}
}

func TestServerIDPatternMatchesOnlyWhatTheServerAccepts(t *testing.T) {
	for id, want := range map[string]bool{
		"prov-2026-1a2b3c4d":   true,
		"prov-2026-000264e5":   true,
		"prov-2026-001":        false,
		"prov-2026-037":        false,
		"prov-2026-1A2B3C4D":   false,
		"prov-2026-1a2b3c4":    false,
		"prov-2026-1a2b3c4d5":  false,
		"prov-26-1a2b3c4d":     false,
		"prov-2026-1a2b3c4d-x": false,
		"":                     false,
	} {
		if got := ValidServerID(id); got != want {
			t.Errorf("ValidServerID(%q) = %v, want %v", id, got, want)
		}
	}
}

func TestPublishFlagsIDsTheServerRejectsBeforeAnyRequest(t *testing.T) {
	c, srv := newClient(t)
	requests := 0
	srv.Accept = func(string) bool { requests++; return true }
	st := localStore(t, 1)
	ctx := context.Background()
	// The legacy ids are not in the store at all: they must not even be loaded.
	plan, err := PlanPublish(ctx, st, []string{"prov-2026-001", "prov-2026-037"}, c)
	if err != nil || len(plan.Problems) != 2 || !plan.Empty() || requests != 0 {
		t.Fatalf("plan = %+v %v requests=%d", plan, err, requests)
	}
	for i, id := range []string{"prov-2026-001", "prov-2026-037"} {
		p := plan.Problems[i]
		if p.ID != id || p.Kind != "invalid id" || !strings.Contains(p.Reason, "prov-2026-1a2b3c4d") {
			t.Errorf("problem %d = %+v", i, p)
		}
	}
}

func TestPublishPlansTheValidRecordsAroundTheBadIDs(t *testing.T) {
	c, srv := newClient(t)
	st := localStore(t, 2)
	ctx := context.Background()
	plan, err := PlanPublish(ctx, st, []string{"prov-2026-001", testID, "prov-2026-002"}, c)
	if err != nil || len(plan.Problems) != 2 || len(plan.Items) != 1 || plan.Items[0].ID != testID {
		t.Fatalf("plan = %+v %v", plan, err)
	}
	if err := ApplyPublish(ctx, plan, c, "", nil, nil); err != nil || srv.Versions(testID) != 2 {
		t.Fatalf("apply: %v versions=%d", err, srv.Versions(testID))
	}
}
