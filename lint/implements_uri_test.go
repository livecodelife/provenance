package lint

import (
	"testing"

	"github.com/livecodelife/provenance/record"
)

func TestImplementsURIRule(t *testing.T) {
	const good = "at://did:plc:ragnj6bgwbfkhz4luybnthat/dev.provisio.provenance.record/prov-2026-0000b001"
	cases := []struct {
		name string
		uri  string
		want int
	}{
		{"valid", good, 0},
		{"not an at-uri", "https://example.com/prov-2026-0000b001", 1},
		{"wrong collection", "at://did:plc:ragnj6bgwbfkhz4luybnthat/dev.provisio.provenance.version/prov-2026-0000b001", 1},
		{"names another record", "at://did:plc:ragnj6bgwbfkhz4luybnthat/dev.provisio.provenance.record/prov-2026-0000b999", 1},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			s := &record.FilesStore{Root: t.TempDir()}
			bp := newRecord("prov-2026-0000b001", record.StatusDraft)
			im := newRecord("prov-2026-0000c001", record.StatusDraft)
			im.Type, im.Implements, im.ImplementsURI = record.TypeImprint, bp.ID, c.uri
			for _, r := range []*record.Record{bp, im} {
				_ = s.Save(r, nil)
			}
			if got := rules(lintStore(t, s), im.ID)[RuleImplementsPin]; got != c.want {
				t.Fatalf("PROV027 x%d, want %d", got, c.want)
			}
		})
	}
}

func TestImplementsURIWithoutImplementsIsFlagged(t *testing.T) {
	s := &record.FilesStore{Root: t.TempDir()}
	r := newRecord("prov-2026-0000c001", record.StatusDraft)
	r.ImplementsURI = "at://did:plc:ragnj6bgwbfkhz4luybnthat/dev.provisio.provenance.record/prov-2026-0000b001"
	_ = s.Save(r, nil)
	if got := rules(lintStore(t, s), r.ID)[RuleImplementsPin]; got != 1 {
		t.Fatalf("PROV027 x%d, want 1", got)
	}
}
