//ff:func feature=gen-nestjs type=util control=iteration dimension=1
//ff:what hasAuthOp — Op 배열에 Auth 연산 포함 여부 확인

package ssac

import "github.com/park-jun-woo/yongol/pkg/generate/ir"

// hasAuthOp returns true if any op is an auth operation.
func hasAuthOp(ops []ir.Op) bool {
	for _, op := range ops {
		if op.Kind == ir.OpAuth {
			return true
		}
	}
	return false
}
