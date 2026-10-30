package record

// sedum:anchor:imports
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-69558a8c","gen":"sha256:4721a84e18ad254c6c184b2b7a9374526b466066a3f3a3756f205c20958c7040","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"record/canonical_int.go","path":"bytes"},"kwargs":"sha256:a3ddad3bcf9e1fe230e64970f95ab6a0997ddd3a5c1e60651e1b1c80c3ac611e","body":"sha256:468731f6dd513df0a66690d4649dc9e7244c96de227d8e6ead98cc5e9747c57a"}
import "bytes"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-69558a8c","gen":"sha256:4721a84e18ad254c6c184b2b7a9374526b466066a3f3a3756f205c20958c7040","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"record/canonical_int.go","path":"strconv"},"kwargs":"sha256:b9517bd844a84a09f0022b1f1ff4fa37721ae7fb6230a012a066a02835187d2f","body":"sha256:f33f981c64a42d78a5daa7ff3fca15b10045080cccd282fef544b59eaf7f6dc8"}
import "strconv"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-69558a8c","gen":"sha256:4721a84e18ad254c6c184b2b7a9374526b466066a3f3a3756f205c20958c7040","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"record/canonical_int.go","name":"writeInt"},"kwargs":"sha256:fd43c86a337d4c0ffa10bf2a232fae99ef32d408db0f23a19ebb97897802f41b","body":"sha256:32cd315dbd5124fc584a0c0ff664f506b4fb3ca42f327fe42568b81073a90125"}
// writeInt writes a signed integer as bare base-10 digits: no plus sign, no
// leading zeros, negatives as -N (docs/HASHING.md section 5).
func writeInt(b *bytes.Buffer, v int64) {
	b.WriteString(strconv.FormatInt(v, 10))
}

// /sedum:setFunction
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-69558a8c","gen":"sha256:4721a84e18ad254c6c184b2b7a9374526b466066a3f3a3756f205c20958c7040","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"record/canonical_int.go","name":"writeUint"},"kwargs":"sha256:85e5bb274aca29d77b17f72ca5fb7ccec7d63c493a407afbb5b87be4fa552195","body":"sha256:fb045481e6ed59f84381bc8f5cf03a9450b5064198e92703d7689721be7398a9"}
// writeUint writes an unsigned integer as bare base-10 digits.
func writeUint(b *bytes.Buffer, v uint64) {
	b.WriteString(strconv.FormatUint(v, 10))
}

// /sedum:setFunction
