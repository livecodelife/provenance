# Context for sedum grow (record package)

Pass with: sedum-haiku grow --generators generators --records provenance --prompt-file .sedum/context.md --only <id>
Advisory only: behaviour belongs in the record. Editing this file does not regenerate anything.

Conventions
- Package record (module github.com/livecodelife/provenance). Generated Go is gofmt-clean and calls only helpers listed here or in the record.
- Canonical encoding is docs/HASHING.md section 5: bytes are written to a *bytes.Buffer; no encoding/json.
- Each import is its own addImport. Write exactly `_, skip := ...` when a binding would otherwise be unused.

Helpers that exist in package record (call, do not redefine)
- writeCanonical(b *bytes.Buffer, v any): the encoder switch in hash.go. Its default case panics "record: non-canonical value type".
- writeString(b *bytes.Buffer, s string): a JSON string with only the section 5 escapes.
- Canonical(v any) string; hashBytes(data []byte) string.
