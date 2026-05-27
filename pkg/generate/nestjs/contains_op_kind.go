//ff:func feature=gen-nestjs type=util control=iteration dimension=1
//ff:what containsOpKind — Op 배열에서 특정 OpKind 포함 여부 확인

package nestjs

import "github.com/park-jun-woo/yongol/pkg/generate/ir"

// containsOpKind returns true if any op matches the given kind.
func containsOpKind(ops []ir.Op, kind ir.OpKind) bool {
	for _, op := range ops {
		if op.Kind == kind {
			return true
		}
	}
	return false
}
