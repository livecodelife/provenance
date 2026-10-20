package record

import "regexp"

// implementsURIPattern matches the at-uri of a parent's head record; group 1
// is the record id.
var implementsURIPattern = regexp.MustCompile(`^at://did:[a-z]+:[A-Za-z0-9._:%-]+/dev\.provisio\.provenance\.record/([A-Za-z0-9._:~-]+)$`)

// ParseImplementsURI returns the record id named by an implements_uri, and
// false when uri is not at://<did>/dev.provisio.provenance.record/<id>.
func ParseImplementsURI(uri string) (string, bool) {
	m := implementsURIPattern.FindStringSubmatch(uri)
	if m == nil {
		return "", false
	}
	return m[1], true
}
