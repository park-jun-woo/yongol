//ff:func feature=gen-nestjs type=util control=iteration dimension=1
//ff:what opsReferenceQuery — Op 배열에 LocQuery 참조 FieldArg 포함 여부 확인

package ssac

import "github.com/park-jun-woo/yongol/pkg/generate/ir"

// opsReferenceQuery returns true if any FieldArg in ops has Location == LocQuery.
func opsReferenceQuery(ops []ir.Op) bool {
	for _, op := range ops {
		if opReferencesLocation(op, ir.LocQuery) {
			return true
		}
	}
	return false
}
