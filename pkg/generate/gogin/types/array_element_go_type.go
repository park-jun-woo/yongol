//ff:func feature=gen-gogin type=util control=selection
//ff:what arrayElementGoType — 배열 element head → Go element 타입 (지원 4 family)

package types

// arrayElementGoType maps a PG array element head token to its Go
// element type. Returns ok=false for elements outside the four
// natively-supported families (Integer / Float / String / Boolean) so
// the caller can route to KindUnsupported.
func arrayElementGoType(elementHead string) (string, bool) {
	switch {
	case isIntegerHead(elementHead):
		return "int64", true
	case isFloatHead(elementHead):
		return "float64", true
	case isStringHead(elementHead):
		return "string", true
	case isBooleanHead(elementHead):
		return "bool", true
	}
	return "", false
}
