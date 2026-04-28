//ff:func feature=ssacmeta type=util control=sequence
//ff:what evaluateWhenTruthy — handle the truthy `manifest.<path>` case of EvaluateWhen

package ssacmeta

import "strings"

// evaluateWhenTruthy handles the bare `manifest.<path>` (truthy) form of the
// when: DSL.
func evaluateWhenTruthy(expr string, manifest map[string]any) bool {
	v, ok := lookupPath(manifest, strings.TrimPrefix(expr, "manifest."))
	if !ok {
		return false
	}
	return truthy(v)
}
