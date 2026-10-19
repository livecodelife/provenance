package remote

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/livecodelife/provenance/record"
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
