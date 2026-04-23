//ff:func feature=migration type=parser control=selection
//ff:what consumeMultiWordTypeTail — VARYING/PRECISION/WITH TIME ZONE 등 이어지는 토큰 소비
package migration

import "strings"

// consumeMultiWordTypeTail extends typePartsPtr with VARYING / PRECISION /
// "WITH TIME ZONE" / "WITHOUT TIME ZONE" tokens as applicable, and
// returns the new reader index.
func consumeMultiWordTypeTail(_ []string, toks []string, i int, typePartsPtr *[]string) int {
	next := strings.ToUpper(toks[i])
	switch {
	case next == "VARYING" || next == "PRECISION":
		*typePartsPtr = append(*typePartsPtr, toks[i])
		return i + 1
	case (next == "WITH" || next == "WITHOUT") &&
		i+2 < len(toks) &&
		strings.ToUpper(toks[i+1]) == "TIME" &&
		strings.ToUpper(toks[i+2]) == "ZONE":
		*typePartsPtr = append(*typePartsPtr, toks[i], toks[i+1], toks[i+2])
		return i + 3
	}
	return i
}
