//ff:func feature=ssacmeta type=util control=sequence
//ff:what evaluateWhenEquality — handle `manifest.<path> == "<value>"` case of EvaluateWhen

package ssacmeta

import "strings"

// evaluateWhenEquality handles the `manifest.<path> == "<value>"` form of
// the when: DSL. `eqIdx` is the index of the `==` separator in expr.
func evaluateWhenEquality(expr string, eqIdx int, manifest map[string]any) bool {
	lhs := strings.TrimSpace(expr[:eqIdx])
	rhs := strings.TrimSpace(expr[eqIdx+2:])
	rhs = strings.Trim(rhs, `"`)
	v, ok := lookupPath(manifest, strings.TrimPrefix(lhs, "manifest."))
	if !ok {
		return false
	}
	return fmtLiteral(v) == rhs
}
