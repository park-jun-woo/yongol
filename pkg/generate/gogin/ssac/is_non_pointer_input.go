//ff:func feature=gen-gogin type=util control=selection
//ff:what isNonPointerInput — SSaC 값이 Go value type 인지 판별 (Ptr/non-Ptr pgtypex 분기용)

package ssac

import "strings"

// isNonPointerInput returns true when the SSaC value resolves to a Go
// value type (not a pointer). Used to select between Ptr and non-Ptr
// pgtypex bridge variants.
func (g *methodGen) isNonPointerInput(ssacValue string) bool {
	if isLiteral(ssacValue) {
		return true
	}

	if !strings.HasPrefix(ssacValue, "request.") {
		return false
	}

	field := ssacValue[len("request."):]

	if g.PathParams[field] {
		return true
	}

	if qp, isQuery := g.QueryParams[field]; isQuery {
		return qp.IsRequired
	}

	return g.BodyRequiredFields[field]
}
