package main

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-d17b9e44","gen":"sha256:315c66a50f0a199a267746ec3105d78a1b8a619606b138f294efc1ae2c242dba","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/graph_run.go","path":"encoding/json"}}
import "encoding/json"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-d17b9e44","gen":"sha256:315c66a50f0a199a267746ec3105d78a1b8a619606b138f294efc1ae2c242dba","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/graph_run.go","path":"fmt"}}
import "fmt"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-d17b9e44","gen":"sha256:315c66a50f0a199a267746ec3105d78a1b8a619606b138f294efc1ae2c242dba","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/graph_run.go","path":"io"}}
import "io"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-d17b9e44","gen":"sha256:315c66a50f0a199a267746ec3105d78a1b8a619606b138f294efc1ae2c242dba","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/graph_run.go","path":"sort"}}
import "sort"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-d17b9e44","gen":"sha256:315c66a50f0a199a267746ec3105d78a1b8a619606b138f294efc1ae2c242dba","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/graph_run.go","path":"strings"}}
import "strings"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-d17b9e44","gen":"sha256:315c66a50f0a199a267746ec3105d78a1b8a619606b138f294efc1ae2c242dba","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/graph_run.go","path":"github.com/spf13/cobra"}}
import "github.com/spf13/cobra"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-d17b9e44","gen":"sha256:315c66a50f0a199a267746ec3105d78a1b8a619606b138f294efc1ae2c242dba","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/graph_run.go","path":"github.com/livecodelife/provenance/record"}}
import "github.com/livecodelife/provenance/record"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-d17b9e44","gen":"sha256:315c66a50f0a199a267746ec3105d78a1b8a619606b138f294efc1ae2c242dba","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/graph_run.go","path":"github.com/livecodelife/provenance/stale"}}
import "github.com/livecodelife/provenance/stale"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setType {"tier":"owned","record":"prov-2026-d17b9e44","gen":"sha256:315c66a50f0a199a267746ec3105d78a1b8a619606b138f294efc1ae2c242dba","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type graphVersion struct {\n\tVersion    int    `json:\"version\"`\n\tReason     string `json:\"reason,omitempty\"`\n\tReasonKind string `json:\"reason_kind,omitempty\"`\n}","file":"cmd/provenance/graph_run.go","name":"graphVersion"}}
type graphVersion struct {
	Version    int    `json:"version"`
	Reason     string `json:"reason,omitempty"`
	ReasonKind string `json:"reason_kind,omitempty"`
}

// /sedum:setType
// sedum:setType {"tier":"owned","record":"prov-2026-d17b9e44","gen":"sha256:315c66a50f0a199a267746ec3105d78a1b8a619606b138f294efc1ae2c242dba","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type graphNode struct {\n\tID                string         `json:\"id\"`\n\tTitle             string         `json:\"title\"`\n\tType              string         `json:\"type\"`\n\tStatus            string         `json:\"status\"`\n\tVersion           int            `json:\"version\"`\n\tVersions          []graphVersion `json:\"versions\"`\n\tImplements        string         `json:\"implements,omitempty\"`\n\tImplementsVersion int            `json:\"implements_version,omitempty\"`\n\tShared            bool           `json:\"shared,omitempty\"`\n}","file":"cmd/provenance/graph_run.go","name":"graphNode"}}
type graphNode struct {
	ID                string         `json:"id"`
	Title             string         `json:"title"`
	Type              string         `json:"type"`
	Status            string         `json:"status"`
	Version           int            `json:"version"`
	Versions          []graphVersion `json:"versions"`
	Implements        string         `json:"implements,omitempty"`
	ImplementsVersion int            `json:"implements_version,omitempty"`
	Shared            bool           `json:"shared,omitempty"`
}

// /sedum:setType
// sedum:setType {"tier":"owned","record":"prov-2026-d17b9e44","gen":"sha256:315c66a50f0a199a267746ec3105d78a1b8a619606b138f294efc1ae2c242dba","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type graphEdge struct {\n\tFrom          string `json:\"from\"`\n\tTo            string `json:\"to\"`\n\tEdge          string `json:\"edge\"`\n\tPinned        int    `json:\"pinned,omitempty\"`\n\tParentVersion int    `json:\"parent_version,omitempty\"`\n\tStale         string `json:\"stale,omitempty\"`\n}","file":"cmd/provenance/graph_run.go","name":"graphEdge"}}
type graphEdge struct {
	From          string `json:"from"`
	To            string `json:"to"`
	Edge          string `json:"edge"`
	Pinned        int    `json:"pinned,omitempty"`
	ParentVersion int    `json:"parent_version,omitempty"`
	Stale         string `json:"stale,omitempty"`
}

// /sedum:setType
// sedum:setType {"tier":"owned","record":"prov-2026-d17b9e44","gen":"sha256:315c66a50f0a199a267746ec3105d78a1b8a619606b138f294efc1ae2c242dba","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","kwargs":{"code":"type graphDoc struct {\n\tNodes []graphNode `json:\"nodes\"`\n\tEdges []graphEdge `json:\"edges\"`\n}","file":"cmd/provenance/graph_run.go","name":"graphDoc"}}
type graphDoc struct {
	Nodes []graphNode `json:"nodes"`
	Edges []graphEdge `json:"edges"`
}

// /sedum:setType
// sedum:setFunction {"tier":"owned","record":"prov-2026-d17b9e44","gen":"sha256:315c66a50f0a199a267746ec3105d78a1b8a619606b138f294efc1ae2c242dba","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func graphNodeOf(r *repo, rec *record.Record, shared bool) graphNode {\n\tn := graphNode{ID: rec.ID, Title: rec.Title, Type: string(rec.EffectiveType()), Status: string(rec.Status), Version: rec.Version, Versions: []graphVersion{}, Implements: rec.Implements, ImplementsVersion: rec.ImplementsVersion, Shared: shared}\n\tif shared {\n\t\treturn n\n\t}\n\tversions, _ := r.Store.Versions(rec.ID)\n\tfor _, v := range versions {\n\t\tif snap, err := r.Store.Snapshot(rec.ID, v); err == nil {\n\t\t\tn.Versions = append(n.Versions, graphVersion{Version: v, Reason: snap.Reason, ReasonKind: snap.ReasonKind})\n\t\t}\n\t}\n\treturn n\n}","file":"cmd/provenance/graph_run.go","name":"graphNodeOf"}}
func graphNodeOf(r *repo, rec *record.Record, shared bool) graphNode {
	n := graphNode{ID: rec.ID, Title: rec.Title, Type: string(rec.EffectiveType()), Status: string(rec.Status), Version: rec.Version, Versions: []graphVersion{}, Implements: rec.Implements, ImplementsVersion: rec.ImplementsVersion, Shared: shared}
	if shared {
		return n
	}
	versions, _ := r.Store.Versions(rec.ID)
	for _, v := range versions {
		if snap, err := r.Store.Snapshot(rec.ID, v); err == nil {
			n.Versions = append(n.Versions, graphVersion{Version: v, Reason: snap.Reason, ReasonKind: snap.ReasonKind})
		}
	}
	return n
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-d17b9e44","gen":"sha256:315c66a50f0a199a267746ec3105d78a1b8a619606b138f294efc1ae2c242dba","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func buildGraph(r *repo, root, filter string) (graphDoc, error) {\n\tids, err := r.Store.List()\n\tif err != nil {\n\t\treturn graphDoc{}, err\n\t}\n\tlocal := map[string]*record.Record{}\n\tfor _, id := range ids {\n\t\trec, _, err := r.Store.Load(id)\n\t\tif err != nil {\n\t\t\treturn graphDoc{}, err\n\t\t}\n\t\tlocal[id] = rec\n\t}\n\tshared, _ := loadShared(r)\n\tfound, err := stale.Find(r.Store)\n\tif err != nil {\n\t\treturn graphDoc{}, err\n\t}\n\tstate := map[string]string{}\n\tfor _, f := range found {\n\t\tstate[f.ID] = string(f.Kind)\n\t}\n\tkeep := map[string]bool{}\n\tif root == \"\" {\n\t\tfor id := range local {\n\t\t\tkeep[id] = true\n\t\t}\n\t} else {\n\t\tif _, ok := local[root]; !ok {\n\t\t\treturn graphDoc{}, fmt.Errorf(\"record not found: %s\", root)\n\t\t}\n\t\tkeep[root] = true\n\t\tif p := local[root].Implements; p != \"\" {\n\t\t\tkeep[p] = true\n\t\t}\n\t\tfor grew := true; grew; {\n\t\t\tgrew = false\n\t\t\tfor id, rec := range local {\n\t\t\t\tif !keep[id] && keep[rec.Implements] && rec.Implements != local[root].Implements {\n\t\t\t\t\tkeep[id] = true\n\t\t\t\t\tgrew = true\n\t\t\t\t}\n\t\t\t}\n\t\t}\n\t}\n\tif filter != \"\" {\n\t\tfor id := range keep {\n\t\t\tif rec, ok := local[id]; ok && string(rec.Status) != filter {\n\t\t\t\tdelete(keep, id)\n\t\t\t}\n\t\t}\n\t}\n\tvar order []string\n\tfor id := range keep {\n\t\tif _, ok := local[id]; ok {\n\t\t\torder = append(order, id)\n\t\t}\n\t}\n\tsort.Strings(order)\n\tdoc := graphDoc{Nodes: []graphNode{}, Edges: []graphEdge{}}\n\tsharedSeen := map[string]bool{}\n\tfor _, id := range order {\n\t\trec := local[id]\n\t\tdoc.Nodes = append(doc.Nodes, graphNodeOf(r, rec, false))\n\t\tif rec.Implements != \"\" {\n\t\t\te := graphEdge{From: id, To: rec.Implements, Edge: \"implements\", Pinned: rec.ImplementsVersion, Stale: state[id]}\n\t\t\tif p, ok := local[rec.Implements]; ok {\n\t\t\t\te.ParentVersion = p.Version\n\t\t\t} else if p, ok := shared[rec.Implements]; ok {\n\t\t\t\te.ParentVersion = p.Version\n\t\t\t\tif !sharedSeen[p.ID] {\n\t\t\t\t\tsharedSeen[p.ID] = true\n\t\t\t\t\tdoc.Nodes = append(doc.Nodes, graphNodeOf(r, p, true))\n\t\t\t\t}\n\t\t\t}\n\t\t\tdoc.Edges = append(doc.Edges, e)\n\t\t}\n\t\tfor _, rel := range rec.Related {\n\t\t\tdoc.Edges = append(doc.Edges, graphEdge{From: id, To: rel, Edge: \"related\"})\n\t\t}\n\t}\n\tsort.SliceStable(doc.Nodes, func(i, j int) bool { return doc.Nodes[i].ID < doc.Nodes[j].ID })\n\tsort.SliceStable(doc.Edges, func(i, j int) bool {\n\t\ta, b := doc.Edges[i], doc.Edges[j]\n\t\tif a.From != b.From {\n\t\t\treturn a.From < b.From\n\t\t}\n\t\tif a.Edge != b.Edge {\n\t\t\treturn a.Edge < b.Edge\n\t\t}\n\t\treturn a.To < b.To\n\t})\n\treturn doc, nil\n}","file":"cmd/provenance/graph_run.go","name":"buildGraph"}}
func buildGraph(r *repo, root, filter string) (graphDoc, error) {
	ids, err := r.Store.List()
	if err != nil {
		return graphDoc{}, err
	}
	local := map[string]*record.Record{}
	for _, id := range ids {
		rec, _, err := r.Store.Load(id)
		if err != nil {
			return graphDoc{}, err
		}
		local[id] = rec
	}
	shared, _ := loadShared(r)
	found, err := stale.Find(r.Store)
	if err != nil {
		return graphDoc{}, err
	}
	state := map[string]string{}
	for _, f := range found {
		state[f.ID] = string(f.Kind)
	}
	keep := map[string]bool{}
	if root == "" {
		for id := range local {
			keep[id] = true
		}
	} else {
		if _, ok := local[root]; !ok {
			return graphDoc{}, fmt.Errorf("record not found: %s", root)
		}
		keep[root] = true
		if p := local[root].Implements; p != "" {
			keep[p] = true
		}
		for grew := true; grew; {
			grew = false
			for id, rec := range local {
				if !keep[id] && keep[rec.Implements] && rec.Implements != local[root].Implements {
					keep[id] = true
					grew = true
				}
			}
		}
	}
	if filter != "" {
		for id := range keep {
			if rec, ok := local[id]; ok && string(rec.Status) != filter {
				delete(keep, id)
			}
		}
	}
	var order []string
	for id := range keep {
		if _, ok := local[id]; ok {
			order = append(order, id)
		}
	}
	sort.Strings(order)
	doc := graphDoc{Nodes: []graphNode{}, Edges: []graphEdge{}}
	sharedSeen := map[string]bool{}
	for _, id := range order {
		rec := local[id]
		doc.Nodes = append(doc.Nodes, graphNodeOf(r, rec, false))
		if rec.Implements != "" {
			e := graphEdge{From: id, To: rec.Implements, Edge: "implements", Pinned: rec.ImplementsVersion, Stale: state[id]}
			if p, ok := local[rec.Implements]; ok {
				e.ParentVersion = p.Version
			} else if p, ok := shared[rec.Implements]; ok {
				e.ParentVersion = p.Version
				if !sharedSeen[p.ID] {
					sharedSeen[p.ID] = true
					doc.Nodes = append(doc.Nodes, graphNodeOf(r, p, true))
				}
			}
			doc.Edges = append(doc.Edges, e)
		}
		for _, rel := range rec.Related {
			doc.Edges = append(doc.Edges, graphEdge{From: id, To: rel, Edge: "related"})
		}
	}
	sort.SliceStable(doc.Nodes, func(i, j int) bool { return doc.Nodes[i].ID < doc.Nodes[j].ID })
	sort.SliceStable(doc.Edges, func(i, j int) bool {
		a, b := doc.Edges[i], doc.Edges[j]
		if a.From != b.From {
			return a.From < b.From
		}
		if a.Edge != b.Edge {
			return a.Edge < b.Edge
		}
		return a.To < b.To
	})
	return doc, nil
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-d17b9e44","gen":"sha256:315c66a50f0a199a267746ec3105d78a1b8a619606b138f294efc1ae2c242dba","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func graphTitle(title string) string {\n\tif len(title) > 60 {\n\t\treturn title[:57] + \"...\"\n\t}\n\treturn title\n}","file":"cmd/provenance/graph_run.go","name":"graphTitle"}}
func graphTitle(title string) string {
	if len(title) > 60 {
		return title[:57] + "..."
	}
	return title
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-d17b9e44","gen":"sha256:315c66a50f0a199a267746ec3105d78a1b8a619606b138f294efc1ae2c242dba","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func printGraphTree(out io.Writer, doc graphDoc) {\n\tbyID := map[string]graphNode{}\n\tfor _, n := range doc.Nodes {\n\t\tbyID[n.ID] = n\n\t}\n\tpin := map[string]graphEdge{}\n\tchildren := map[string][]string{}\n\tfor _, e := range doc.Edges {\n\t\tif e.Edge == \"implements\" {\n\t\t\tpin[e.From] = e\n\t\t\tif _, ok := byID[e.To]; ok {\n\t\t\t\tchildren[e.To] = append(children[e.To], e.From)\n\t\t\t}\n\t\t}\n\t}\n\tvar walk func(id string, depth int)\n\twalk = func(id string, depth int) {\n\t\tn := byID[id]\n\t\tline := fmt.Sprintf(\"%s%s@%d  %s  %s  %s\", strings.Repeat(\"  \", depth+1), n.ID, n.Version, n.Status, n.Type, graphTitle(n.Title))\n\t\tif n.Shared {\n\t\t\tline += \"  (shared)\"\n\t\t}\n\t\tif e, ok := pin[id]; ok {\n\t\t\tline += fmt.Sprintf(\"  -> %s@%d\", e.To, e.Pinned)\n\t\t\tswitch e.Stale {\n\t\t\tcase \"direct\", \"transitive\":\n\t\t\t\tline += fmt.Sprintf(\"  stale (%s, parent at %d)\", e.Stale, e.ParentVersion)\n\t\t\tcase \"unknown\":\n\t\t\t\tline += \"  unknown pin\"\n\t\t\t}\n\t\t}\n\t\tfmt.Fprintln(out, line)\n\t\tfor _, c := range children[id] {\n\t\t\twalk(c, depth+1)\n\t\t}\n\t}\n\tfmt.Fprint(out, \"\\nPROVENANCE GRAPH\\n\\n\")\n\tfor _, n := range doc.Nodes {\n\t\tif e, ok := pin[n.ID]; ok {\n\t\t\tif _, parentShown := byID[e.To]; parentShown {\n\t\t\t\tcontinue\n\t\t\t}\n\t\t}\n\t\twalk(n.ID, 0)\n\t}\n\tfmt.Fprint(out, \"\\n\")\n}","file":"cmd/provenance/graph_run.go","name":"printGraphTree"}}
func printGraphTree(out io.Writer, doc graphDoc) {
	byID := map[string]graphNode{}
	for _, n := range doc.Nodes {
		byID[n.ID] = n
	}
	pin := map[string]graphEdge{}
	children := map[string][]string{}
	for _, e := range doc.Edges {
		if e.Edge == "implements" {
			pin[e.From] = e
			if _, ok := byID[e.To]; ok {
				children[e.To] = append(children[e.To], e.From)
			}
		}
	}
	var walk func(id string, depth int)
	walk = func(id string, depth int) {
		n := byID[id]
		line := fmt.Sprintf("%s%s@%d  %s  %s  %s", strings.Repeat("  ", depth+1), n.ID, n.Version, n.Status, n.Type, graphTitle(n.Title))
		if n.Shared {
			line += "  (shared)"
		}
		if e, ok := pin[id]; ok {
			line += fmt.Sprintf("  -> %s@%d", e.To, e.Pinned)
			switch e.Stale {
			case "direct", "transitive":
				line += fmt.Sprintf("  stale (%s, parent at %d)", e.Stale, e.ParentVersion)
			case "unknown":
				line += "  unknown pin"
			}
		}
		fmt.Fprintln(out, line)
		for _, c := range children[id] {
			walk(c, depth+1)
		}
	}
	fmt.Fprint(out, "\nPROVENANCE GRAPH\n\n")
	for _, n := range doc.Nodes {
		if e, ok := pin[n.ID]; ok {
			if _, parentShown := byID[e.To]; parentShown {
				continue
			}
		}
		walk(n.ID, 0)
	}
	fmt.Fprint(out, "\n")
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-d17b9e44","gen":"sha256:315c66a50f0a199a267746ec3105d78a1b8a619606b138f294efc1ae2c242dba","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func printGraphDot(out io.Writer, doc graphDoc) {\n\tcolors := map[string]string{\"draft\": \"gray\", \"open\": \"orange\", \"implemented\": \"green\", \"deprecated\": \"red\"}\n\tfmt.Fprintln(out, \"digraph provenance {\")\n\tfmt.Fprintln(out, \"  rankdir=TB;\")\n\tfmt.Fprintln(out, \"  node [shape=box];\")\n\tfor _, n := range doc.Nodes {\n\t\tcolor := colors[n.Status]\n\t\tif color == \"\" {\n\t\t\tcolor = \"black\"\n\t\t}\n\t\tfmt.Fprintf(out, \"  %q [label=%q, color=%s];\\n\", n.ID, fmt.Sprintf(\"%s@%d\\n%s\", n.ID, n.Version, graphTitle(n.Title)), color)\n\t}\n\tfor _, e := range doc.Edges {\n\t\tswitch {\n\t\tcase e.Edge == \"related\":\n\t\t\tfmt.Fprintf(out, \"  %q -> %q [style=dotted];\\n\", e.From, e.To)\n\t\tcase e.Stale == \"direct\" || e.Stale == \"transitive\":\n\t\t\tfmt.Fprintf(out, \"  %q -> %q [label=%q, style=dashed, color=red];\\n\", e.From, e.To, fmt.Sprintf(\"@%d (parent at %d)\", e.Pinned, e.ParentVersion))\n\t\tdefault:\n\t\t\tfmt.Fprintf(out, \"  %q -> %q [label=%q];\\n\", e.From, e.To, fmt.Sprintf(\"@%d\", e.Pinned))\n\t\t}\n\t}\n\tfmt.Fprintln(out, \"}\")\n}","file":"cmd/provenance/graph_run.go","name":"printGraphDot"}}
func printGraphDot(out io.Writer, doc graphDoc) {
	colors := map[string]string{"draft": "gray", "open": "orange", "implemented": "green", "deprecated": "red"}
	fmt.Fprintln(out, "digraph provenance {")
	fmt.Fprintln(out, "  rankdir=TB;")
	fmt.Fprintln(out, "  node [shape=box];")
	for _, n := range doc.Nodes {
		color := colors[n.Status]
		if color == "" {
			color = "black"
		}
		fmt.Fprintf(out, "  %q [label=%q, color=%s];\n", n.ID, fmt.Sprintf("%s@%d\n%s", n.ID, n.Version, graphTitle(n.Title)), color)
	}
	for _, e := range doc.Edges {
		switch {
		case e.Edge == "related":
			fmt.Fprintf(out, "  %q -> %q [style=dotted];\n", e.From, e.To)
		case e.Stale == "direct" || e.Stale == "transitive":
			fmt.Fprintf(out, "  %q -> %q [label=%q, style=dashed, color=red];\n", e.From, e.To, fmt.Sprintf("@%d (parent at %d)", e.Pinned, e.ParentVersion))
		default:
			fmt.Fprintf(out, "  %q -> %q [label=%q];\n", e.From, e.To, fmt.Sprintf("@%d", e.Pinned))
		}
	}
	fmt.Fprintln(out, "}")
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-d17b9e44","gen":"sha256:315c66a50f0a199a267746ec3105d78a1b8a619606b138f294efc1ae2c242dba","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func runGraph(out io.Writer, cmd *cobra.Command, args []string) error {\n\tr, err := openRepo(cmd)\n\tif err != nil {\n\t\treturn err\n\t}\n\troot, _ := cmd.Flags().GetString(\"root\")\n\tfilter, _ := cmd.Flags().GetString(\"filter\")\n\tformat, _ := cmd.Flags().GetString(\"format\")\n\tswitch format {\n\tcase \"\", \"human\", \"json\", \"dot\":\n\tdefault:\n\t\treturn fmt.Errorf(\"unknown format %q (human|json|dot)\", format)\n\t}\n\tswitch filter {\n\tcase \"\", \"draft\", \"open\", \"implemented\", \"deprecated\":\n\tdefault:\n\t\treturn fmt.Errorf(\"unknown filter %q (draft|open|implemented|deprecated)\", filter)\n\t}\n\tdoc, err := buildGraph(r, root, filter)\n\tif err != nil {\n\t\treturn err\n\t}\n\tswitch format {\n\tcase \"json\":\n\t\tenc := json.NewEncoder(out)\n\t\tenc.SetIndent(\"\", \"  \")\n\t\treturn enc.Encode(doc)\n\tcase \"dot\":\n\t\tprintGraphDot(out, doc)\n\tdefault:\n\t\tprintGraphTree(out, doc)\n\t}\n\treturn nil\n}","file":"cmd/provenance/graph_run.go","name":"runGraph"}}
func runGraph(out io.Writer, cmd *cobra.Command, args []string) error {
	r, err := openRepo(cmd)
	if err != nil {
		return err
	}
	root, _ := cmd.Flags().GetString("root")
	filter, _ := cmd.Flags().GetString("filter")
	format, _ := cmd.Flags().GetString("format")
	switch format {
	case "", "human", "json", "dot":
	default:
		return fmt.Errorf("unknown format %q (human|json|dot)", format)
	}
	switch filter {
	case "", "draft", "open", "implemented", "deprecated":
	default:
		return fmt.Errorf("unknown filter %q (draft|open|implemented|deprecated)", filter)
	}
	doc, err := buildGraph(r, root, filter)
	if err != nil {
		return err
	}
	switch format {
	case "json":
		enc := json.NewEncoder(out)
		enc.SetIndent("", "  ")
		return enc.Encode(doc)
	case "dot":
		printGraphDot(out, doc)
	default:
		printGraphTree(out, doc)
	}
	return nil
}

// /sedum:setFunction
