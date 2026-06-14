//ff:func feature=validate type=util control=sequence topic=openapi-ddl
//ff:what isCollectionType — 타입 문자열이 슬라이스/래퍼(컬렉션)인지 판정

package openapi_ddl

import "strings"

// isCollectionType reports whether a raw type spec denotes a collection
// ("[]Gig", "Page[Gig]", "Cursor[Gig]") rather than a single resource. A
// pointer ("*User") or bare/package-qualified type is not a collection. Used to
// exclude list/paginated responses from canonical single-resource grouping.
func isCollectionType(raw string) bool {
	if strings.HasPrefix(raw, "[]") {
		return true
	}
	if i := strings.IndexByte(raw, '['); i > 0 && strings.HasSuffix(raw, "]") {
		return true
	}
	return false
}
