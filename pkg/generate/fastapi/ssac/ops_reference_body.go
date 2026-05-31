//ff:func feature=gen-fastapi type=util control=iteration dimension=1
//ff:what opsReferenceBody — Op 배열에 LocBody 참조 FieldArg 포함 여부 확인

package ssac

import "github.com/park-jun-woo/yongol/pkg/generate/ir"

// opsReferenceBody returns true if any FieldArg in ops has Location == LocBody.
func opsReferenceBody(ops []ir.Op) bool {
	for _, op := range ops {
		if opReferencesLocation(op, ir.LocBody) {
			return true
		}
	}
	return false
}
