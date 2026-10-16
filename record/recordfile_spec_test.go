package record

import "testing"

// Proof for prov-2026-4234f0bf: a spec that is a record file contributes the
// constant marker record-file/v1 to content_hash (docs/HASHING.md section 4).

func recordFileHash(t *testing.T, specPath string, files map[string]string) (string, []string) {
	t.Helper()
	r := &Record{Intent: "Does the thing.", AssociatedSpecs: []AssociatedSpec{{Path: specPath}}}
	resolve := func(path string) (string, bool) {
		c, ok := files[path]
		if !ok {
			return "", false
		}
		return SpecHash([]byte(c)), true
	}
	return r.ContentHash(resolve)
}

func TestRecordFileSpecIgnoresBytes(t *testing.T) {
	const p = "provenance/prov-2026-0000b001.yml"
	a, missA := recordFileHash(t, p, map[string]string{p: "version: 1\n"})
	b, missB := recordFileHash(t, p, map[string]string{p: "version: 2\ncontent_hash: x\n"})
	c, missC := recordFileHash(t, p, nil)
	if a != b || a != c {
		t.Errorf("record-file spec hash depends on the file: %s %s %s", a, b, c)
	}
	if len(missA)+len(missB)+len(missC) != 0 {
		t.Errorf("record-file spec reported missing: %v %v %v", missA, missB, missC)
	}
}

func TestRecordFileSpecAnyDirectoryAndSuffix(t *testing.T) {
	want, _ := recordFileHash(t, "provenance/prov-2026-0000b001.yml", nil)
	for _, p := range []string{"prov-2026-0000b001.yml", "other/dir/prov-2026-0000b001.yml", "provenance/prov-2026-0000b001-api.yml"} {
		if got, _ := recordFileHash(t, p, nil); got != want {
			t.Errorf("%s: hash %s, want %s", p, got, want)
		}
	}
}

func TestRecordFileSpecDiffersFromMissing(t *testing.T) {
	rec, _ := recordFileHash(t, "provenance/prov-2026-0000b001.yml", nil)
	other, miss := recordFileHash(t, "specs/x.linespec", nil)
	if rec == other {
		t.Error("marker hashes like a missing spec")
	}
	if len(miss) != 1 {
		t.Errorf("missing = %v, want the one unreadable spec", miss)
	}
}

func TestLookalikeIsAnOrdinarySpec(t *testing.T) {
	for _, p := range []string{"specs/prov-2026-0000b001.txt", "specs/prov-2026-xyz.yml", "specs/prov-26-0000b001.yml", "specs/my-prov-2026-0000b001.yml"} {
		got, _ := recordFileHash(t, p, map[string]string{p: "a\n"})
		again, _ := recordFileHash(t, p, map[string]string{p: "b\n"})
		if got == again {
			t.Errorf("%s: treated as a record file", p)
		}
	}
}

func TestRecordFileSpecKeepsSlot(t *testing.T) {
	r := &Record{AssociatedSpecs: []AssociatedSpec{{Path: "provenance/prov-2026-0000b001.yml"}, {Path: "specs/a"}}}
	s := &Record{AssociatedSpecs: []AssociatedSpec{{Path: "specs/a"}, {Path: "provenance/prov-2026-0000b001.yml"}}}
	res := func(string) (string, bool) { return SpecHash([]byte("a")), true }
	a, _ := r.ContentHash(res)
	b, _ := s.ContentHash(res)
	if a == b {
		t.Error("reordering specs did not change the hash")
	}
}
