package main

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-4eb70ef7","gen":"sha256:8dd0cf07c10ca2de12963c41bee8f3709b76949e3c89a1471e2831d9ab3e96fe","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/search_run.go","path":"encoding/json"}}
import "encoding/json"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-4eb70ef7","gen":"sha256:8dd0cf07c10ca2de12963c41bee8f3709b76949e3c89a1471e2831d9ab3e96fe","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/search_run.go","path":"fmt"}}
import "fmt"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-4eb70ef7","gen":"sha256:8dd0cf07c10ca2de12963c41bee8f3709b76949e3c89a1471e2831d9ab3e96fe","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/search_run.go","path":"io"}}
import "io"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-4eb70ef7","gen":"sha256:8dd0cf07c10ca2de12963c41bee8f3709b76949e3c89a1471e2831d9ab3e96fe","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/search_run.go","path":"math"}}
import "math"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-4eb70ef7","gen":"sha256:8dd0cf07c10ca2de12963c41bee8f3709b76949e3c89a1471e2831d9ab3e96fe","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/search_run.go","path":"sort"}}
import "sort"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-4eb70ef7","gen":"sha256:8dd0cf07c10ca2de12963c41bee8f3709b76949e3c89a1471e2831d9ab3e96fe","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/search_run.go","path":"strings"}}
import "strings"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-4eb70ef7","gen":"sha256:8dd0cf07c10ca2de12963c41bee8f3709b76949e3c89a1471e2831d9ab3e96fe","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/search_run.go","path":"unicode"}}
import "unicode"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-4eb70ef7","gen":"sha256:8dd0cf07c10ca2de12963c41bee8f3709b76949e3c89a1471e2831d9ab3e96fe","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/search_run.go","path":"github.com/spf13/cobra"}}
import "github.com/spf13/cobra"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-4eb70ef7","gen":"sha256:8dd0cf07c10ca2de12963c41bee8f3709b76949e3c89a1471e2831d9ab3e96fe","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/search_run.go","path":"github.com/livecodelife/provenance/record"}}
import "github.com/livecodelife/provenance/record"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setType {"tier":"owned","record":"prov-2026-4eb70ef7","gen":"sha256:8dd0cf07c10ca2de12963c41bee8f3709b76949e3c89a1471e2831d9ab3e96fe","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type searchHit struct {\n\tID      string   `json:\"id\"`\n\tVersion int      `json:\"version\"`\n\tStatus  string   `json:\"status\"`\n\tType    string   `json:\"type\"`\n\tTitle   string   `json:\"title\"`\n\tScore   float64  `json:\"score\"`\n\tShared  bool     `json:\"shared,omitempty\"`\n\tMatched []string `json:\"matched\"`\n\tExcerpt string   `json:\"excerpt\"`\n}","file":"cmd/provenance/search_run.go","name":"searchHit"}}
type searchHit struct {
	ID      string   `json:"id"`
	Version int      `json:"version"`
	Status  string   `json:"status"`
	Type    string   `json:"type"`
	Title   string   `json:"title"`
	Score   float64  `json:"score"`
	Shared  bool     `json:"shared,omitempty"`
	Matched []string `json:"matched"`
	Excerpt string   `json:"excerpt"`
}

// /sedum:setType
// sedum:setType {"tier":"owned","record":"prov-2026-4eb70ef7","gen":"sha256:8dd0cf07c10ca2de12963c41bee8f3709b76949e3c89a1471e2831d9ab3e96fe","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type searchResult struct {\n\tQuery   string      `json:\"query\"`\n\tBackend string      `json:\"backend\"`\n\tResults []searchHit `json:\"results\"`\n}","file":"cmd/provenance/search_run.go","name":"searchResult"}}
type searchResult struct {
	Query   string      `json:"query"`
	Backend string      `json:"backend"`
	Results []searchHit `json:"results"`
}

// /sedum:setType
// sedum:setType {"tier":"owned","record":"prov-2026-4eb70ef7","gen":"sha256:8dd0cf07c10ca2de12963c41bee8f3709b76949e3c89a1471e2831d9ab3e96fe","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type searchField struct {\n\tName   string\n\tWeight float64\n\tText   string\n}","file":"cmd/provenance/search_run.go","name":"searchField"}}
type searchField struct {
	Name   string
	Weight float64
	Text   string
}

// /sedum:setType
// sedum:setType {"tier":"owned","record":"prov-2026-4eb70ef7","gen":"sha256:8dd0cf07c10ca2de12963c41bee8f3709b76949e3c89a1471e2831d9ab3e96fe","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type searchEntry struct {\n\tRec    *record.Record\n\tShared bool\n\tFields []searchField\n\tTF     map[string]float64\n\tLen    float64\n}","file":"cmd/provenance/search_run.go","name":"searchEntry"}}
type searchEntry struct {
	Rec    *record.Record
	Shared bool
	Fields []searchField
	TF     map[string]float64
	Len    float64
}

// /sedum:setType
// sedum:setVar {"tier":"owned","record":"prov-2026-4eb70ef7","gen":"sha256:8dd0cf07c10ca2de12963c41bee8f3709b76949e3c89a1471e2831d9ab3e96fe","pkg":"sha256:c7b9d9ac1655e8211b8b2ca8c49cf34b8d510d8d823d4f5eb60ec7e1018aa2eb","kwargs":{"code":"var searchStop = map[string]bool{\"a\": true, \"an\": true, \"and\": true, \"are\": true, \"as\": true, \"at\": true, \"be\": true, \"by\": true, \"for\": true, \"from\": true, \"has\": true, \"in\": true, \"is\": true, \"it\": true, \"its\": true, \"of\": true, \"on\": true, \"or\": true, \"that\": true, \"the\": true, \"this\": true, \"to\": true, \"was\": true, \"with\": true}","file":"cmd/provenance/search_run.go","name":"searchStop"}}
var searchStop = map[string]bool{"a": true, "an": true, "and": true, "are": true, "as": true, "at": true, "be": true, "by": true, "for": true, "from": true, "has": true, "in": true, "is": true, "it": true, "its": true, "of": true, "on": true, "or": true, "that": true, "the": true, "this": true, "to": true, "was": true, "with": true}

// /sedum:setVar
// sedum:setFunction {"tier":"owned","record":"prov-2026-4eb70ef7","gen":"sha256:8dd0cf07c10ca2de12963c41bee8f3709b76949e3c89a1471e2831d9ab3e96fe","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func searchTerms(s string) []string {\n\tvar terms []string\n\tfor _, w := range strings.FieldsFunc(strings.ToLower(s), func(c rune) bool { return !unicode.IsLetter(c) && !unicode.IsDigit(c) }) {\n\t\tif len([]rune(w)) >= 2 && !searchStop[w] {\n\t\t\tterms = append(terms, w)\n\t\t}\n\t}\n\treturn terms\n}","file":"cmd/provenance/search_run.go","name":"searchTerms"}}
func searchTerms(s string) []string {
	var terms []string
	for _, w := range strings.FieldsFunc(strings.ToLower(s), func(c rune) bool { return !unicode.IsLetter(c) && !unicode.IsDigit(c) }) {
		if len([]rune(w)) >= 2 && !searchStop[w] {
			terms = append(terms, w)
		}
	}
	return terms
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-4eb70ef7","gen":"sha256:8dd0cf07c10ca2de12963c41bee8f3709b76949e3c89a1471e2831d9ab3e96fe","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func searchEntryOf(r *repo, rec *record.Record, shared bool) *searchEntry {\n\te := &searchEntry{Rec: rec, Shared: shared, TF: map[string]float64{}}\n\treasons := \"\"\n\tif !shared {\n\t\tif versions, err := r.Store.Versions(rec.ID); err == nil {\n\t\t\tfor _, v := range versions {\n\t\t\t\tif snap, err := r.Store.Snapshot(rec.ID, v); err == nil {\n\t\t\t\t\treasons += snap.Reason + \"\\n\"\n\t\t\t\t}\n\t\t\t}\n\t\t}\n\t}\n\te.Fields = []searchField{\n\t\t{\"id\", 5, rec.ID + \" \" + strings.Join(rec.Aliases, \" \")},\n\t\t{\"title\", 4, rec.Title},\n\t\t{\"tags\", 3, strings.Join(rec.Tags, \" \")},\n\t\t{\"scope\", 2, strings.Join(rec.AffectedScope, \" \")},\n\t\t{\"intent\", 1, rec.Intent},\n\t\t{\"constraints\", 1, strings.Join(rec.Constraints, \"\\n\")},\n\t\t{\"reasons\", 1, reasons},\n\t}\n\tfor _, f := range e.Fields {\n\t\tfor _, t := range searchTerms(f.Text) {\n\t\t\te.TF[t] += f.Weight\n\t\t\te.Len += f.Weight\n\t\t}\n\t}\n\treturn e\n}","file":"cmd/provenance/search_run.go","name":"searchEntryOf"}}
func searchEntryOf(r *repo, rec *record.Record, shared bool) *searchEntry {
	e := &searchEntry{Rec: rec, Shared: shared, TF: map[string]float64{}}
	reasons := ""
	if !shared {
		if versions, err := r.Store.Versions(rec.ID); err == nil {
			for _, v := range versions {
				if snap, err := r.Store.Snapshot(rec.ID, v); err == nil {
					reasons += snap.Reason + "\n"
				}
			}
		}
	}
	e.Fields = []searchField{
		{"id", 5, rec.ID + " " + strings.Join(rec.Aliases, " ")},
		{"title", 4, rec.Title},
		{"tags", 3, strings.Join(rec.Tags, " ")},
		{"scope", 2, strings.Join(rec.AffectedScope, " ")},
		{"intent", 1, rec.Intent},
		{"constraints", 1, strings.Join(rec.Constraints, "\n")},
		{"reasons", 1, reasons},
	}
	for _, f := range e.Fields {
		for _, t := range searchTerms(f.Text) {
			e.TF[t] += f.Weight
			e.Len += f.Weight
		}
	}
	return e
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-4eb70ef7","gen":"sha256:8dd0cf07c10ca2de12963c41bee8f3709b76949e3c89a1471e2831d9ab3e96fe","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func searchExcerpt(e *searchEntry, query []string) (string, []string) {\n\twant := map[string]bool{}\n\tfor _, q := range query {\n\t\twant[q] = true\n\t}\n\tvar matched []string\n\tbestWeight, best := 0.0, \"\"\n\tfor _, f := range e.Fields {\n\t\thits := 0\n\t\tfor _, t := range searchTerms(f.Text) {\n\t\t\tif want[t] {\n\t\t\t\thits++\n\t\t\t}\n\t\t}\n\t\tif hits == 0 {\n\t\t\tcontinue\n\t\t}\n\t\tmatched = append(matched, f.Name)\n\t\tif w := f.Weight * float64(hits); w > bestWeight {\n\t\t\tbestWeight = w\n\t\t\tbest = f.Text\n\t\t\tfor _, line := range strings.Split(f.Text, \"\\n\") {\n\t\t\t\tfor _, t := range searchTerms(line) {\n\t\t\t\t\tif want[t] {\n\t\t\t\t\t\tbest = line\n\t\t\t\t\t\tbreak\n\t\t\t\t\t}\n\t\t\t\t}\n\t\t\t\tif best == line {\n\t\t\t\t\tbreak\n\t\t\t\t}\n\t\t\t}\n\t\t}\n\t}\n\treturn ctxClean(strings.TrimSpace(best), 120), matched\n}","file":"cmd/provenance/search_run.go","name":"searchExcerpt"}}
func searchExcerpt(e *searchEntry, query []string) (string, []string) {
	want := map[string]bool{}
	for _, q := range query {
		want[q] = true
	}
	var matched []string
	bestWeight, best := 0.0, ""
	for _, f := range e.Fields {
		hits := 0
		for _, t := range searchTerms(f.Text) {
			if want[t] {
				hits++
			}
		}
		if hits == 0 {
			continue
		}
		matched = append(matched, f.Name)
		if w := f.Weight * float64(hits); w > bestWeight {
			bestWeight = w
			best = f.Text
			for _, line := range strings.Split(f.Text, "\n") {
				for _, t := range searchTerms(line) {
					if want[t] {
						best = line
						break
					}
				}
				if best == line {
					break
				}
			}
		}
	}
	return ctxClean(strings.TrimSpace(best), 120), matched
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-4eb70ef7","gen":"sha256:8dd0cf07c10ca2de12963c41bee8f3709b76949e3c89a1471e2831d9ab3e96fe","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func searchRecords(r *repo, query []string, status, typ string) ([]searchHit, error) {\n\tids, err := r.Store.List()\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\tvar entries []*searchEntry\n\tlocalIDs := map[string]bool{}\n\tfor _, id := range ids {\n\t\tlocalIDs[id] = true\n\t\trec, _, err := r.Store.Load(id)\n\t\tif err != nil {\n\t\t\treturn nil, err\n\t\t}\n\t\tentries = append(entries, searchEntryOf(r, rec, false))\n\t}\n\tshared, _ := loadShared(r)\n\tfor id, rec := range shared {\n\t\tif !localIDs[id] {\n\t\t\tentries = append(entries, searchEntryOf(r, rec, true))\n\t\t}\n\t}\n\tvar pool []*searchEntry\n\tfor _, e := range entries {\n\t\tif status != \"\" && string(e.Rec.Status) != status {\n\t\t\tcontinue\n\t\t}\n\t\tif status == \"\" && e.Rec.Status == record.StatusDeprecated {\n\t\t\tcontinue\n\t\t}\n\t\tif typ != \"\" && string(e.Rec.EffectiveType()) != typ {\n\t\t\tcontinue\n\t\t}\n\t\tpool = append(pool, e)\n\t}\n\tn := float64(len(pool))\n\ttotal := 0.0\n\tdf := map[string]float64{}\n\tfor _, e := range pool {\n\t\ttotal += e.Len\n\t\tfor t := range e.TF {\n\t\t\tdf[t]++\n\t\t}\n\t}\n\tavg := 1.0\n\tif n > 0 && total > 0 {\n\t\tavg = total / n\n\t}\n\tconst k1, b = 1.2, 0.75\n\tvar hits []searchHit\n\tfor _, e := range pool {\n\t\tscore := 0.0\n\t\tfor _, q := range query {\n\t\t\ttf := e.TF[q]\n\t\t\tif tf == 0 {\n\t\t\t\tcontinue\n\t\t\t}\n\t\t\tidf := math.Log(1 + (n-df[q]+0.5)/(df[q]+0.5))\n\t\t\tscore += idf * tf * (k1 + 1) / (tf + k1*(1-b+b*e.Len/avg))\n\t\t}\n\t\tif score <= 0 {\n\t\t\tcontinue\n\t\t}\n\t\texcerpt, matched := searchExcerpt(e, query)\n\t\thits = append(hits, searchHit{ID: e.Rec.ID, Version: e.Rec.Version, Status: string(e.Rec.Status), Type: string(e.Rec.EffectiveType()), Title: e.Rec.Title, Score: math.Round(score*1000) / 1000, Shared: e.Shared, Matched: matched, Excerpt: excerpt})\n\t}\n\tsort.SliceStable(hits, func(i, j int) bool {\n\t\tif hits[i].Score != hits[j].Score {\n\t\t\treturn hits[i].Score > hits[j].Score\n\t\t}\n\t\treturn hits[i].ID < hits[j].ID\n\t})\n\treturn hits, nil\n}","file":"cmd/provenance/search_run.go","name":"searchRecords"}}
func searchRecords(r *repo, query []string, status, typ string) ([]searchHit, error) {
	ids, err := r.Store.List()
	if err != nil {
		return nil, err
	}
	var entries []*searchEntry
	localIDs := map[string]bool{}
	for _, id := range ids {
		localIDs[id] = true
		rec, _, err := r.Store.Load(id)
		if err != nil {
			return nil, err
		}
		entries = append(entries, searchEntryOf(r, rec, false))
	}
	shared, _ := loadShared(r)
	for id, rec := range shared {
		if !localIDs[id] {
			entries = append(entries, searchEntryOf(r, rec, true))
		}
	}
	var pool []*searchEntry
	for _, e := range entries {
		if status != "" && string(e.Rec.Status) != status {
			continue
		}
		if status == "" && e.Rec.Status == record.StatusDeprecated {
			continue
		}
		if typ != "" && string(e.Rec.EffectiveType()) != typ {
			continue
		}
		pool = append(pool, e)
	}
	n := float64(len(pool))
	total := 0.0
	df := map[string]float64{}
	for _, e := range pool {
		total += e.Len
		for t := range e.TF {
			df[t]++
		}
	}
	avg := 1.0
	if n > 0 && total > 0 {
		avg = total / n
	}
	const k1, b = 1.2, 0.75
	var hits []searchHit
	for _, e := range pool {
		score := 0.0
		for _, q := range query {
			tf := e.TF[q]
			if tf == 0 {
				continue
			}
			idf := math.Log(1 + (n-df[q]+0.5)/(df[q]+0.5))
			score += idf * tf * (k1 + 1) / (tf + k1*(1-b+b*e.Len/avg))
		}
		if score <= 0 {
			continue
		}
		excerpt, matched := searchExcerpt(e, query)
		hits = append(hits, searchHit{ID: e.Rec.ID, Version: e.Rec.Version, Status: string(e.Rec.Status), Type: string(e.Rec.EffectiveType()), Title: e.Rec.Title, Score: math.Round(score*1000) / 1000, Shared: e.Shared, Matched: matched, Excerpt: excerpt})
	}
	sort.SliceStable(hits, func(i, j int) bool {
		if hits[i].Score != hits[j].Score {
			return hits[i].Score > hits[j].Score
		}
		return hits[i].ID < hits[j].ID
	})
	return hits, nil
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-4eb70ef7","gen":"sha256:8dd0cf07c10ca2de12963c41bee8f3709b76949e3c89a1471e2831d9ab3e96fe","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func runSearch(out io.Writer, cmd *cobra.Command, args []string) error {\n\tr, err := openRepo(cmd)\n\tif err != nil {\n\t\treturn err\n\t}\n\ttext, _ := cmd.Flags().GetString(\"query\")\n\tif text == \"\" && len(args) > 0 {\n\t\ttext = args[0]\n\t}\n\tlimit, _ := cmd.Flags().GetInt(\"limit\")\n\tstatus, _ := cmd.Flags().GetString(\"status\")\n\ttyp, _ := cmd.Flags().GetString(\"type\")\n\tformat, _ := cmd.Flags().GetString(\"format\")\n\tif format != \"\" && format != \"human\" && format != \"json\" {\n\t\treturn fmt.Errorf(\"unknown format %q (human|json)\", format)\n\t}\n\tswitch status {\n\tcase \"\", \"draft\", \"open\", \"implemented\", \"deprecated\":\n\tdefault:\n\t\treturn fmt.Errorf(\"unknown status %q (draft|open|implemented|deprecated)\", status)\n\t}\n\tswitch typ {\n\tcase \"\", \"brief\", \"blueprint\", \"imprint\":\n\tdefault:\n\t\treturn fmt.Errorf(\"unknown type %q (brief|blueprint|imprint)\", typ)\n\t}\n\tif limit == 0 {\n\t\tlimit = 5\n\t}\n\tif limit < 1 || limit > 50 {\n\t\treturn fmt.Errorf(\"--limit must be between 1 and 50\")\n\t}\n\tquery := searchTerms(text)\n\tif len(query) == 0 {\n\t\treturn fmt.Errorf(\"empty query: give search words (letters and digits, at least 2 characters, not only common words)\")\n\t}\n\thits, err := searchRecords(r, nameQuery(query), status, typ)\n\tif err != nil {\n\t\treturn err\n\t}\n\tif len(hits) > limit {\n\t\thits = hits[:limit]\n\t}\n\tif hits == nil {\n\t\thits = []searchHit{}\n\t}\n\tif format == \"json\" {\n\t\tenc := json.NewEncoder(out)\n\t\tenc.SetIndent(\"\", \"  \")\n\t\treturn enc.Encode(searchResult{Query: text, Backend: \"lexical\", Results: hits})\n\t}\n\tfmt.Fprintf(out, \"\\nSEARCH: %s (lexical)\\n\\n\", ctxClean(text, 200))\n\tif len(hits) == 0 {\n\t\tfmt.Fprint(out, \"  no matching records\\n\\n\")\n\t\treturn nil\n\t}\n\tfor _, h := range hits {\n\t\tmark := \"\"\n\t\tif h.Shared {\n\t\t\tmark = \"  (shared)\"\n\t\t}\n\t\tfmt.Fprintf(out, \"  %s@%d  %s  %s  %s  (score %.3f)%s\\n\", h.ID, h.Version, h.Status, h.Type, ctxClean(h.Title, 100), h.Score, mark)\n\t\tif h.Excerpt != \"\" {\n\t\t\tfmt.Fprintf(out, \"      %s\\n\", h.Excerpt)\n\t\t}\n\t}\n\tfmt.Fprint(out, \"\\n\")\n\treturn nil\n}","file":"cmd/provenance/search_run.go","name":"runSearch"}}
func runSearch(out io.Writer, cmd *cobra.Command, args []string) error {
	r, err := openRepo(cmd)
	if err != nil {
		return err
	}
	text, _ := cmd.Flags().GetString("query")
	if text == "" && len(args) > 0 {
		text = args[0]
	}
	limit, _ := cmd.Flags().GetInt("limit")
	status, _ := cmd.Flags().GetString("status")
	typ, _ := cmd.Flags().GetString("type")
	format, _ := cmd.Flags().GetString("format")
	if format != "" && format != "human" && format != "json" {
		return fmt.Errorf("unknown format %q (human|json)", format)
	}
	switch status {
	case "", "draft", "open", "implemented", "deprecated":
	default:
		return fmt.Errorf("unknown status %q (draft|open|implemented|deprecated)", status)
	}
	switch typ {
	case "", "brief", "blueprint", "imprint":
	default:
		return fmt.Errorf("unknown type %q (brief|blueprint|imprint)", typ)
	}
	if limit == 0 {
		limit = 5
	}
	if limit < 1 || limit > 50 {
		return fmt.Errorf("--limit must be between 1 and 50")
	}
	query := searchTerms(text)
	if len(query) == 0 {
		return fmt.Errorf("empty query: give search words (letters and digits, at least 2 characters, not only common words)")
	}
	hits, err := searchRecords(r, nameQuery(query), status, typ)
	if err != nil {
		return err
	}
	if len(hits) > limit {
		hits = hits[:limit]
	}
	if hits == nil {
		hits = []searchHit{}
	}
	if format == "json" {
		enc := json.NewEncoder(out)
		enc.SetIndent("", "  ")
		return enc.Encode(searchResult{Query: text, Backend: "lexical", Results: hits})
	}
	fmt.Fprintf(out, "\nSEARCH: %s (lexical)\n\n", ctxClean(text, 200))
	if len(hits) == 0 {
		fmt.Fprint(out, "  no matching records\n\n")
		return nil
	}
	for _, h := range hits {
		mark := ""
		if h.Shared {
			mark = "  (shared)"
		}
		fmt.Fprintf(out, "  %s@%d  %s  %s  %s  (score %.3f)%s\n", h.ID, h.Version, h.Status, h.Type, ctxClean(h.Title, 100), h.Score, mark)
		if h.Excerpt != "" {
			fmt.Fprintf(out, "      %s\n", h.Excerpt)
		}
	}
	fmt.Fprint(out, "\n")
	return nil
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-4eb70ef7","gen":"sha256:8dd0cf07c10ca2de12963c41bee8f3709b76949e3c89a1471e2831d9ab3e96fe","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func nameQuery(terms []string) []string {\n\tseen := map[string]bool{}\n\tvar out []string\n\tfor _, t := range terms {\n\t\tif !seen[t] {\n\t\t\tseen[t] = true\n\t\t\tout = append(out, t)\n\t\t}\n\t}\n\treturn out\n}","file":"cmd/provenance/search_run.go","name":"nameQuery"}}
func nameQuery(terms []string) []string {
	seen := map[string]bool{}
	var out []string
	for _, t := range terms {
		if !seen[t] {
			seen[t] = true
			out = append(out, t)
		}
	}
	return out
}

// /sedum:setFunction
