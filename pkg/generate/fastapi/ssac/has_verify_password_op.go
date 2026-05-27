//ff:func feature=gen-fastapi type=util control=iteration dimension=1
//ff:what hasVerifyPasswordOp — Op 배열에 VerifyPassword 연산 포함 여부 확인

package ssac

import "github.com/park-jun-woo/yongol/pkg/generate/ir"

// hasVerifyPasswordOp returns true if any op is a verify-password operation.
// Login endpoints with @verify-password are pre-auth and must not require
// authentication dependencies.
func hasVerifyPasswordOp(ops []ir.Op) bool {
	for _, op := range ops {
		if op.Kind == ir.OpVerifyPassword {
			return true
		}
	}
	return false
}
