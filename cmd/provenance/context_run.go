package main

// sedum:anchor:imports
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:57a095967ea4cbd86875cea0f96e504d2e1d138c4f31fba23d7b01ba011caa69","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/context_run.go","path":"encoding/json"},"kwargs":"sha256:a8ed679eb55d9d50923d8893529d0b3d3cc013c68859e6568f88e75435918f67","body":"sha256:872d7c3e9dee608e79a7f8ef57930d24b5e30f2a57d891a52549179ace45946a"}
import "encoding/json"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:57a095967ea4cbd86875cea0f96e504d2e1d138c4f31fba23d7b01ba011caa69","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/context_run.go","path":"fmt"},"kwargs":"sha256:7ef25b959b2a3d1ce42d8621700520be1c4f95f066b43fb5d977dca18ec3d2cc","body":"sha256:bcbac51883cbf8064e9196b371cb880cee429866652447ad557533d817857609"}
import "fmt"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:57a095967ea4cbd86875cea0f96e504d2e1d138c4f31fba23d7b01ba011caa69","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/context_run.go","path":"io"},"kwargs":"sha256:15ef99b279eac89d2007b56a85be8cf43421cf1374d23dc51a0579029b1c0258","body":"sha256:ef8e8abf4ef5ee40a19dc85b49bf09c5df1fe32abfe834e0ab4c2175baf80ddf"}
import "io"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:57a095967ea4cbd86875cea0f96e504d2e1d138c4f31fba23d7b01ba011caa69","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/context_run.go","path":"sort"},"kwargs":"sha256:dcc974dc8e17498a5e7e989fd2ea42aea8101b3f28f851eec72eac439b4f35e4","body":"sha256:f197172617ae21ef23b1855b09aafb14571020ae81bff877d63d3ff81f4bf63c"}
import "sort"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:57a095967ea4cbd86875cea0f96e504d2e1d138c4f31fba23d7b01ba011caa69","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/context_run.go","path":"strings"},"kwargs":"sha256:a0ab05e7da739f273b05c1fceeab55cf5dafdda0acb05e012e015bfe82b8c1ec","body":"sha256:a9af3c0bc4bf9342df9d7e89a324523fdba7941a9ca4190270e68ff3c4a25ea6"}
import "strings"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:57a095967ea4cbd86875cea0f96e504d2e1d138c4f31fba23d7b01ba011caa69","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/context_run.go","path":"github.com/spf13/cobra"},"kwargs":"sha256:20634da1e0801659a0694d9e8471b2baae2af400a6581d158dc561ce00fdb2c0","body":"sha256:f361c02d0f8e1999cf3fb715393e16310eeec9debd5aaa13936af7617e2b46aa"}
import "github.com/spf13/cobra"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:57a095967ea4cbd86875cea0f96e504d2e1d138c4f31fba23d7b01ba011caa69","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/context_run.go","path":"github.com/livecodelife/provenance/record"},"kwargs":"sha256:7a8637e2111b5fde843fc15c8a92fab60a270a7543dece73b5e9fbd20c88c64d","body":"sha256:acbf8628061291d11aa39e9ea64ecc91095a117057cfe6f1958ef2c437fac003"}
import "github.com/livecodelife/provenance/record"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:57a095967ea4cbd86875cea0f96e504d2e1d138c4f31fba23d7b01ba011caa69","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/context_run.go","path":"github.com/livecodelife/provenance/stale"},"kwargs":"sha256:bca242cd3ee294bf653f5f5095eda4bfcaa9af2c8577b9e94d655bbfe4bc9118","body":"sha256:9d7308eb26aa328e6ffd9c6c6d9334e95ead3e237675b3953a96d84096c468d0"}
import "github.com/livecodelife/provenance/stale"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setType {"v":2,"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:57a095967ea4cbd86875cea0f96e504d2e1d138c4f31fba23d7b01ba011caa69","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","id":{"file":"cmd/provenance/context_run.go","name":"ctxRecord"},"kwargs":"sha256:b6aba1e3a98df905661b7ccb157bfd696e55b641d4ac45db7738936dfb5087e6","body":"sha256:a4e422a20a8174a07f766826e54a5e08796b7950c4d4314115fab50ec5c155c1"}
type ctxRecord struct {
	ID                string   `json:"id"`
	Type              string   `json:"type"`
	Status            string   `json:"status"`
	Version           int      `json:"version"`
	Title             string   `json:"title"`
	LatestReason      string   `json:"latest_reason,omitempty"`
	Implements        string   `json:"implements,omitempty"`
	ImplementsVersion int      `json:"implements_version,omitempty"`
	Stale             string   `json:"stale,omitempty"`
	Evidence          string   `json:"evidence,omitempty"`
	Specs             []string `json:"specs"`
}

// /sedum:setType
// sedum:setType {"v":2,"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:57a095967ea4cbd86875cea0f96e504d2e1d138c4f31fba23d7b01ba011caa69","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","id":{"file":"cmd/provenance/context_run.go","name":"ctxAncestor"},"kwargs":"sha256:cf17298df2d8724cad32949844af915aba4725f8ec7d9261ea9e460158a91eb4","body":"sha256:50e2aafbec5181dcf9372a52727a4c9a8b64cd03f960fa8a354228daa33ab33f"}
type ctxAncestor struct {
	ID      string `json:"id"`
	Version int    `json:"version"`
	Status  string `json:"status"`
	Title   string `json:"title"`
	Shared  bool   `json:"shared,omitempty"`
}

// /sedum:setType
// sedum:setType {"v":2,"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:57a095967ea4cbd86875cea0f96e504d2e1d138c4f31fba23d7b01ba011caa69","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","id":{"file":"cmd/provenance/context_run.go","name":"ctxVersion"},"kwargs":"sha256:139cebf787a606b927b5e2131d72c6135e0aa6a87f700747ecac78d19eed6f0d","body":"sha256:d65ff95fefdb720477b58c6982deb79c903cbdccf89ee382f1916a245b1e60e0"}
type ctxVersion struct {
	Version int    `json:"version"`
	Reason  string `json:"reason,omitempty"`
}

// /sedum:setType
// sedum:setType {"v":2,"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:57a095967ea4cbd86875cea0f96e504d2e1d138c4f31fba23d7b01ba011caa69","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","id":{"file":"cmd/provenance/context_run.go","name":"ctxChild"},"kwargs":"sha256:47f64c279b7551e72ac699dcf34837f037b0c17634eaa2a12ecd0c6bccbeb9fd","body":"sha256:424569722ec7bf7bf38242d28cd20fed112c82005f6209df66d52ca4ba68966b"}
type ctxChild struct {
	ID      string `json:"id"`
	Version int    `json:"version"`
	Pinned  int    `json:"pinned"`
	Stale   string `json:"stale,omitempty"`
}

// /sedum:setType
// sedum:setType {"v":2,"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:57a095967ea4cbd86875cea0f96e504d2e1d138c4f31fba23d7b01ba011caa69","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","id":{"file":"cmd/provenance/context_run.go","name":"ctxRecordDetail"},"kwargs":"sha256:6f89873d0c35986de9324641b16beb39536e00c54f518eafb436bc35200cb8e8","body":"sha256:1bc96d062f0bf4f551aaf9e01a5dcf05f4b5cf81ac04c3c7b3a6fff34c7e139a"}
type ctxRecordDetail struct {
	ctxRecord
	Intent         string        `json:"intent"`
	Constraints    []string      `json:"constraints"`
	AffectedScope  []string      `json:"affected_scope"`
	ForbiddenScope []string      `json:"forbidden_scope"`
	Versions       []ctxVersion  `json:"versions"`
	Parents        []ctxAncestor `json:"parents"`
	Children       []ctxChild    `json:"children"`
	Commits        []string      `json:"commits"`
}

// /sedum:setType
// sedum:setType {"v":2,"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:57a095967ea4cbd86875cea0f96e504d2e1d138c4f31fba23d7b01ba011caa69","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","id":{"file":"cmd/provenance/context_run.go","name":"ctxTarget"},"kwargs":"sha256:d64c2b8163d76585375dedaf2dcbf9cb85dc81d8f5a0aaf9223df299b78b8ce3","body":"sha256:da4f1ed69722cac07de3f9770283befc75c3ba7b60937d3ce73a75160ad2f9cc"}
type ctxTarget struct {
	Target    string                   `json:"target"`
	Kind      string                   `json:"kind"`
	Exempt    bool                     `json:"exempt,omitempty"`
	Governing []ctxRecord              `json:"governing,omitempty"`
	Ancestry  map[string][]ctxAncestor `json:"ancestry,omitempty"`
	Conflicts []string                 `json:"conflicts,omitempty"`
	Next      *nextAction              `json:"next,omitempty"`
	Record    *ctxRecordDetail         `json:"record,omitempty"`
}

// /sedum:setType
// sedum:setType {"v":2,"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:57a095967ea4cbd86875cea0f96e504d2e1d138c4f31fba23d7b01ba011caa69","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","id":{"file":"cmd/provenance/context_run.go","name":"ctxDoc"},"kwargs":"sha256:f7e32bbdf04e9734f46cdad6f2e42a9985e4a4204b1f713aebdca6843d38a33e","body":"sha256:3a4041345e92cc0792d267ae7927c0ed23349a03def90ba20fdcb2530127c427"}
type ctxDoc struct {
	Targets []ctxTarget `json:"targets"`
}

// /sedum:setType
// sedum:setType {"v":2,"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:57a095967ea4cbd86875cea0f96e504d2e1d138c4f31fba23d7b01ba011caa69","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","id":{"file":"cmd/provenance/context_run.go","name":"ctxWorld"},"kwargs":"sha256:c015a2ab09939695666a265052af3d806f86f403e9101ca3aa83a76ab06050e3","body":"sha256:efbd059ffd32bd9c7e42df799a2ecfb58edd72ae78d5f02eb14b83f9e0d87f25"}
type ctxWorld struct {
	r        *repo
	local    map[string]*record.Record
	shared   map[string]*record.Record
	stale    map[string]string
	evidence map[string]string
}

// /sedum:setType
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:57a095967ea4cbd86875cea0f96e504d2e1d138c4f31fba23d7b01ba011caa69","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/context_run.go","name":"ctxClean"},"kwargs":"sha256:86c17a04ed701d2f9959fe66fee0c4e75e76afd3680b7eda8d4d98f20521e81f","body":"sha256:98b53041a16a7095ab4a42605ef3cefe05db2c0efed7fbe57e47f2c718b10ba2"}
func ctxClean(s string, max int) string {
	var b strings.Builder
	for _, c := range s {
		switch {
		case c == '\n' || c == '\t':
			b.WriteRune(' ')
		case c < 32 || (c >= 127 && c <= 159):
		default:
			b.WriteRune(c)
		}
	}
	out := b.String()
	if max > 0 && len([]rune(out)) > max {
		return string([]rune(out)[:max-3]) + "..."
	}
	return out
}

// /sedum:setFunction
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:57a095967ea4cbd86875cea0f96e504d2e1d138c4f31fba23d7b01ba011caa69","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/context_run.go","name":"loadCtxWorld"},"kwargs":"sha256:4d86ae7a87e4557058770d6bf2f094638d5a213949f3f508abdc15d6d1d19047","body":"sha256:400e1b377e994ff76bd3d77a9a998ca629d12e22654bbe8d7a06cd119c139dea"}
func loadCtxWorld(r *repo) (*ctxWorld, error) {
	w := &ctxWorld{r: r, local: map[string]*record.Record{}, stale: map[string]string{}, evidence: map[string]string{}}
	ids, err := r.Store.List()
	if err != nil {
		return nil, err
	}
	for _, id := range ids {
		rec, _, err := r.Store.Load(id)
		if err != nil {
			return nil, err
		}
		w.local[id] = rec
	}
	w.shared, _ = loadShared(r)
	found, err := stale.FindShared(r.Store, w.shared)
	if err != nil {
		return nil, err
	}
	for _, f := range found {
		w.stale[f.ID] = string(f.Kind)
	}
	rows, err := checkEvidence(r)
	if err != nil {
		return nil, err
	}
	for _, row := range rows {
		w.evidence[row.ID] = row.State
	}
	return w, nil
}

// /sedum:setFunction
// sedum:setMethod {"v":2,"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:57a095967ea4cbd86875cea0f96e504d2e1d138c4f31fba23d7b01ba011caa69","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","id":{"file":"cmd/provenance/context_run.go","name":"summary","receiver":"ctxWorld"},"kwargs":"sha256:ba457177df50d11bf29f7c19614783f1729d34f1775a64c8acc2eef41bf39824","body":"sha256:8621bc412fd32f1b713fc61dd89b2b047d5a6c7b60a6b6531bb025d0ca7362b9"}
func (w *ctxWorld) summary(rec *record.Record) ctxRecord {
	c := ctxRecord{ID: rec.ID, Type: string(rec.EffectiveType()), Status: string(rec.Status), Version: rec.Version, Title: rec.Title, Implements: rec.Implements, ImplementsVersion: rec.ImplementsVersion, Stale: w.stale[rec.ID], Evidence: w.evidence[rec.ID], Specs: []string{}}
	if versions, err := w.r.Store.Versions(rec.ID); err == nil && len(versions) > 0 {
		if snap, err := w.r.Store.Snapshot(rec.ID, versions[len(versions)-1]); err == nil {
			c.LatestReason = snap.Reason
		}
	}
	for _, s := range rec.AssociatedSpecs {
		c.Specs = append(c.Specs, s.Path)
	}
	return c
}

// /sedum:setMethod
// sedum:setMethod {"v":2,"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:57a095967ea4cbd86875cea0f96e504d2e1d138c4f31fba23d7b01ba011caa69","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","id":{"file":"cmd/provenance/context_run.go","name":"ancestors","receiver":"ctxWorld"},"kwargs":"sha256:23c898e6beba098aa635898f08105b687f2026d029c1342607eb7404bc7faa51","body":"sha256:c4b47a2c6dab3a3b8f745d6b125bdf540e73da1d266a2f3fde37dfbb4a97a0d4"}
func (w *ctxWorld) ancestors(rec *record.Record) []ctxAncestor {
	out := []ctxAncestor{}
	seen := map[string]bool{rec.ID: true}
	for next := rec.Implements; next != "" && !seen[next] && len(out) < 10; {
		seen[next] = true
		if p, ok := w.local[next]; ok {
			out = append(out, ctxAncestor{ID: p.ID, Version: p.Version, Status: string(p.Status), Title: p.Title})
			next = p.Implements
		} else if p, ok := w.shared[next]; ok {
			out = append(out, ctxAncestor{ID: p.ID, Version: p.Version, Status: string(p.Status), Title: p.Title, Shared: true})
			next = p.Implements
		} else {
			out = append(out, ctxAncestor{ID: next, Status: "unresolved"})
			break
		}
	}
	return out
}

// /sedum:setMethod
// sedum:setMethod {"v":2,"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:57a095967ea4cbd86875cea0f96e504d2e1d138c4f31fba23d7b01ba011caa69","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","id":{"file":"cmd/provenance/context_run.go","name":"pathTarget","receiver":"ctxWorld"},"kwargs":"sha256:e784e1f061d4a71cf87d91431ddaf415fd580ef43b5fdcd6629b05c7002754cf","body":"sha256:d622d15481a8afac1a59786559a99b2f590576112aba886ff771186a8b7b3d10"}
func (w *ctxWorld) pathTarget(path string) (ctxTarget, error) {
	t := ctxTarget{Target: path, Kind: "path"}
	if !nextRelevant(w.r, path) {
		t.Exempt = true
		return t, nil
	}
	var ids []string
	for id := range w.local {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	var open []string
	t.Ancestry = map[string][]ctxAncestor{}
	for _, id := range ids {
		rec := w.local[id]
		if rec.Status == record.StatusDeprecated || !nextGoverns(rec, path) {
			continue
		}
		t.Governing = append(t.Governing, w.summary(rec))
		t.Ancestry[id] = w.ancestors(rec)
		if rec.Status == record.StatusOpen {
			open = append(open, id)
		}
	}
	if len(open) > 1 {
		t.Conflicts = open
	}
	st, err := gatherNext(w.r, []string{path})
	if err != nil {
		return t, err
	}
	if actions := adviseNext(st); len(actions) > 0 {
		t.Next = &actions[0]
	}
	return t, nil
}

// /sedum:setMethod
// sedum:setMethod {"v":2,"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:57a095967ea4cbd86875cea0f96e504d2e1d138c4f31fba23d7b01ba011caa69","pkg":"sha256:8ea07fc52bbc260836983438ef08ad02746b86a1028c8b1a51b60eeb7b83cc2b","id":{"file":"cmd/provenance/context_run.go","name":"recordTarget","receiver":"ctxWorld"},"kwargs":"sha256:61552e95b08b2b972f9b2b576f1cd5ccd11b60eeb2f16588607a855b4230a819","body":"sha256:0f2a9b8710c74603a1fa21446f228e0628c624e2a0086f92bf578907fa30be85"}
func (w *ctxWorld) recordTarget(rec *record.Record) ctxTarget {
	d := &ctxRecordDetail{ctxRecord: w.summary(rec), Intent: rec.Intent, Constraints: append([]string{}, rec.Constraints...), AffectedScope: append([]string{}, rec.AffectedScope...), ForbiddenScope: append([]string{}, rec.ForbiddenScope...), Versions: []ctxVersion{}, Parents: w.ancestors(rec), Children: []ctxChild{}, Commits: []string{}}
	if versions, err := w.r.Store.Versions(rec.ID); err == nil {
		for _, v := range versions {
			if snap, err := w.r.Store.Snapshot(rec.ID, v); err == nil {
				d.Versions = append(d.Versions, ctxVersion{Version: v, Reason: snap.Reason})
			}
		}
	}
	var kids []string
	for id, c := range w.local {
		if c.Implements == rec.ID {
			kids = append(kids, id)
		}
	}
	sort.Strings(kids)
	for _, id := range kids {
		c := w.local[id]
		d.Children = append(d.Children, ctxChild{ID: id, Version: c.Version, Pinned: c.ImplementsVersion, Stale: w.stale[id]})
	}
	for _, sha := range commitsForRecord(w.r, rec) {
		if len(d.Commits) >= 20 {
			break
		}
		if line, err := w.r.git("log", "-1", "--format=%h %s", sha); err == nil {
			d.Commits = append(d.Commits, line)
		}
	}
	return ctxTarget{Target: rec.ID, Kind: "record", Record: d}
}

// /sedum:setMethod
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:57a095967ea4cbd86875cea0f96e504d2e1d138c4f31fba23d7b01ba011caa69","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/context_run.go","name":"ctxLine"},"kwargs":"sha256:d7e1fa83c511b677374ad12a3dcf51d85c76ef8a2f0c67b93a14331caec56609","body":"sha256:38037fa478a973beadd0404ac0b111b8612b16f8d2a19089d1ede2fad8705f6d"}
func ctxLine(c ctxRecord, max int) string {
	line := fmt.Sprintf("%s@%d  %s  %s  %s", c.ID, c.Version, c.Status, c.Type, ctxClean(c.Title, max))
	if c.Stale != "" {
		line += "  stale (" + c.Stale + ")"
	}
	if c.Evidence != "" {
		line += "  evidence: " + c.Evidence
	}
	return line
}

// /sedum:setFunction
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:57a095967ea4cbd86875cea0f96e504d2e1d138c4f31fba23d7b01ba011caa69","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/context_run.go","name":"printCtxHuman"},"kwargs":"sha256:b6fab5a5d392c3f128ad2359bdbe4b2886d6623beb661f9d30da4072b722932a","body":"sha256:894f5f040a3ed632e3611d696c1e8c1f0811e793db91a5e3b700b47bf0efcce8"}
func printCtxHuman(out io.Writer, t ctxTarget) {
	fmt.Fprintf(out, "\nCONTEXT: %s\n", t.Target)
	if t.Kind == "path" {
		if t.Exempt {
			fmt.Fprint(out, "  Exempt: no provenance rules apply to this path.\n")
			return
		}
		if len(t.Governing) == 0 {
			fmt.Fprint(out, "  No record governs this path.\n")
		}
		for _, g := range t.Governing {
			fmt.Fprintf(out, "  %s\n", ctxLine(g, 120))
			if g.LatestReason != "" {
				fmt.Fprintf(out, "      latest: %s\n", ctxClean(g.LatestReason, 160))
			}
			if len(g.Specs) > 0 {
				fmt.Fprintf(out, "      specs: %s\n", ctxClean(strings.Join(g.Specs, ", "), 200))
			}
			for _, a := range t.Ancestry[g.ID] {
				mark := ""
				if a.Shared {
					mark = " (shared)"
				}
				fmt.Fprintf(out, "      from: %s@%d  %s  %s%s\n", a.ID, a.Version, a.Status, ctxClean(a.Title, 120), mark)
			}
		}
		if len(t.Conflicts) > 0 {
			fmt.Fprintf(out, "  Conflict: %d open records cover this path: %s\n", len(t.Conflicts), strings.Join(t.Conflicts, ", "))
		}
		if t.Next != nil {
			fmt.Fprintf(out, "  Next: %s: %s\n", t.Next.Kind, ctxClean(t.Next.Reason, 300))
			if t.Next.Command != "" {
				fmt.Fprintf(out, "    Run: %s\n", ctxClean(t.Next.Command, 300))
			}
		}
		return
	}
	d := t.Record
	fmt.Fprintf(out, "  %s\n", ctxLine(d.ctxRecord, 120))
	fmt.Fprint(out, "  Intent:\n")
	for _, l := range strings.Split(d.Intent, "\n") {
		fmt.Fprintf(out, "    %s\n", ctxClean(l, 400))
	}
	section := func(name string, items []string) {
		if len(items) == 0 {
			return
		}
		fmt.Fprintf(out, "  %s:\n", name)
		for _, it := range items {
			fmt.Fprintf(out, "    · %s\n", ctxClean(it, 400))
		}
	}
	section("Constraints", d.Constraints)
	section("Affected scope", d.AffectedScope)
	section("Forbidden scope", d.ForbiddenScope)
	section("Specs", d.Specs)
	fmt.Fprint(out, "  Versions:\n")
	for _, v := range d.Versions {
		fmt.Fprintf(out, "    · %d  %s\n", v.Version, ctxClean(v.Reason, 200))
	}
	for _, p := range d.Parents {
		fmt.Fprintf(out, "  Implements: %s@%d  %s  %s\n", p.ID, p.Version, p.Status, ctxClean(p.Title, 120))
	}
	for _, c := range d.Children {
		fmt.Fprintf(out, "  Implemented by: %s@%d -> pinned @%d %s\n", c.ID, c.Version, c.Pinned, c.Stale)
	}
	section("Commits", d.Commits)
}

// /sedum:setFunction
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:57a095967ea4cbd86875cea0f96e504d2e1d138c4f31fba23d7b01ba011caa69","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/context_run.go","name":"printCtxCompact"},"kwargs":"sha256:6ac6bb88afaa5b4aabd386ae834f4c938a7bf7a3b11d7e25106af7584e2231f6","body":"sha256:55401296e9c9db3a56af8468a4ae3598993a82a56b675ead9229a5a6ff0a2eb8"}
func printCtxCompact(out io.Writer, t ctxTarget) {
	switch {
	case t.Kind == "record":
		fmt.Fprintln(out, ctxLine(t.Record.ctxRecord, 80))
	case t.Exempt:
		fmt.Fprintf(out, "%s: exempt\n", t.Target)
	case len(t.Governing) == 0:
		fmt.Fprintf(out, "%s: no governing record\n", t.Target)
	default:
		for _, g := range t.Governing {
			fmt.Fprintf(out, "%s: %s\n", t.Target, ctxLine(g, 80))
		}
	}
}

// /sedum:setFunction
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-54ce06a5","gen":"sha256:57a095967ea4cbd86875cea0f96e504d2e1d138c4f31fba23d7b01ba011caa69","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/context_run.go","name":"runContext"},"kwargs":"sha256:85b7cf468363ea2d0a8db9f403b92766501aaac250df8f76b3d3cbbff0ce70d2","body":"sha256:49fdaabd234e2b3c7adf3ba0fffc653d550783fb808909cd585d4800be19ab73"}
func runContext(out io.Writer, cmd *cobra.Command, args []string) error {
	r, err := openRepo(cmd)
	if err != nil {
		return err
	}
	files := splitFileList(mustStringArray(cmd, "files"))
	format, _ := cmd.Flags().GetString("format")
	switch format {
	case "", "human", "compact", "json":
	default:
		return fmt.Errorf("unknown format %q (human|compact|json)", format)
	}
	targets := append(append([]string{}, args...), files...)
	if len(targets) == 0 {
		return fmt.Errorf("no target: give file paths or record ids as arguments, or use --files")
	}
	w, err := loadCtxWorld(r)
	if err != nil {
		return err
	}
	doc := ctxDoc{Targets: []ctxTarget{}}
	for _, target := range targets {
		if rec, ok := w.local[target]; ok {
			doc.Targets = append(doc.Targets, w.recordTarget(rec))
			continue
		}
		t, err := w.pathTarget(strings.TrimPrefix(target, "./"))
		if err != nil {
			return err
		}
		doc.Targets = append(doc.Targets, t)
	}
	switch format {
	case "json":
		enc := json.NewEncoder(out)
		enc.SetIndent("", "  ")
		return enc.Encode(doc)
	case "compact":
		for _, t := range doc.Targets {
			printCtxCompact(out, t)
		}
	default:
		for _, t := range doc.Targets {
			printCtxHuman(out, t)
		}
		fmt.Fprint(out, "\n")
	}
	return nil
}

// /sedum:setFunction
