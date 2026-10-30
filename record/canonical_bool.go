package record

// sedum:anchor:imports
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-33344553","gen":"sha256:d217c82271a6e0db5ab28a53cbe12fb2e4db09aa260effd573674237e0cc1507","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"record/canonical_bool.go","path":"bytes"},"kwargs":"sha256:a54dc911548ffb59b5acb41aba77c033815f3d8ee31ff5b31a8df437f00a7e92","body":"sha256:468731f6dd513df0a66690d4649dc9e7244c96de227d8e6ead98cc5e9747c57a"}
import "bytes"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-33344553","gen":"sha256:d217c82271a6e0db5ab28a53cbe12fb2e4db09aa260effd573674237e0cc1507","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"record/canonical_bool.go","name":"writeBool"},"kwargs":"sha256:f84361d69da386b1499c4bcb7beea79285fa15b088bfb018bec1f559ccdaa108","body":"sha256:3f49fc9ed954241a503340aa79685fe4fdad51d497cab06014bc88b07c2d4752"}
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
