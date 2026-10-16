package record

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-33344553","gen":"sha256:d217c82271a6e0db5ab28a53cbe12fb2e4db09aa260effd573674237e0cc1507","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"record/canonical_bool.go","path":"bytes"}}
import "bytes"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"tier":"owned","record":"prov-2026-33344553","gen":"sha256:d217c82271a6e0db5ab28a53cbe12fb2e4db09aa260effd573674237e0cc1507","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"// writeBool writes a boolean as the bare JSON literal true or false\n// (docs/HASHING.md section 5). It is the one value type writeCanonical\n// accepts beyond strings, null, lists and maps.\nfunc writeBool(b *bytes.Buffer, v bool) {\n\tif v {\n\t\tb.WriteString(\"true\")\n\t\treturn\n\t}\n\tb.WriteString(\"false\")\n}","file":"record/canonical_bool.go","name":"writeBool"}}
// writeBool writes a boolean as the bare JSON literal true or false
// (docs/HASHING.md section 5). It is the one value type writeCanonical
// accepts beyond strings, null, lists and maps.
func writeBool(b *bytes.Buffer, v bool) {
	if v {
		b.WriteString("true")
		return
	}
	b.WriteString("false")
}

// /sedum:setFunction
