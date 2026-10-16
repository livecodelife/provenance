package record

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-69558a8c","gen":"sha256:4721a84e18ad254c6c184b2b7a9374526b466066a3f3a3756f205c20958c7040","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"record/canonical_int.go","path":"bytes"}}
import "bytes"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-69558a8c","gen":"sha256:4721a84e18ad254c6c184b2b7a9374526b466066a3f3a3756f205c20958c7040","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"record/canonical_int.go","path":"strconv"}}
import "strconv"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"tier":"owned","record":"prov-2026-69558a8c","gen":"sha256:4721a84e18ad254c6c184b2b7a9374526b466066a3f3a3756f205c20958c7040","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"// writeInt writes a signed integer as bare base-10 digits: no plus sign, no\n// leading zeros, negatives as -N (docs/HASHING.md section 5).\nfunc writeInt(b *bytes.Buffer, v int64) {\n\tb.WriteString(strconv.FormatInt(v, 10))\n}","file":"record/canonical_int.go","name":"writeInt"}}
// writeInt writes a signed integer as bare base-10 digits: no plus sign, no
// leading zeros, negatives as -N (docs/HASHING.md section 5).
func writeInt(b *bytes.Buffer, v int64) {
	b.WriteString(strconv.FormatInt(v, 10))
}

// /sedum:setFunction
// sedum:setFunction {"tier":"owned","record":"prov-2026-69558a8c","gen":"sha256:4721a84e18ad254c6c184b2b7a9374526b466066a3f3a3756f205c20958c7040","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"// writeUint writes an unsigned integer as bare base-10 digits.\nfunc writeUint(b *bytes.Buffer, v uint64) {\n\tb.WriteString(strconv.FormatUint(v, 10))\n}","file":"record/canonical_int.go","name":"writeUint"}}
// writeUint writes an unsigned integer as bare base-10 digits.
func writeUint(b *bytes.Buffer, v uint64) {
	b.WriteString(strconv.FormatUint(v, 10))
}

// /sedum:setFunction
