//ff:func feature=gen-fastapi type=util control=iteration dimension=1
//ff:what opReferencesLocation — 단일 Op 의 FieldArg 중 지정 Location 참조 포함 여부

package ssac

import "github.com/park-jun-woo/yongol/pkg/generate/ir"

// opReferencesLocation reports whether any FieldArg in op has the given
// Location.
func opReferencesLocation(op ir.Op, loc ir.ParamLocation) bool {
	for _, fa := range collectFieldArgs(op) {
		if fa.Location == loc {
			return true
		}
	}
	return false
}
