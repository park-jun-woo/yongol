//ff:func feature=gen-fastapi type=util control=iteration dimension=1
//ff:what opsReferencePath — Op 배열에 LocPath 참조 FieldArg 포함 여부 확인

package ssac

import "github.com/park-jun-woo/yongol/pkg/generate/ir"

// opsReferencePath returns true if any FieldArg in ops has Location == LocPath.
func opsReferencePath(ops []ir.Op) bool {
	for _, op := range ops {
		for _, fa := range collectFieldArgs(op) {
			if fa.Location == ir.LocPath {
				return true
			}
		}
	}
	return false
}
