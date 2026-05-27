//ff:func feature=gen-nestjs type=util control=iteration dimension=1
//ff:what isDeleteByPK — FieldArg.IsPK 기반 delete 조건이 PK 컬럼인지 확인

package ssac

import "github.com/park-jun-woo/yongol/pkg/generate/ir"

// isDeleteByPK checks if the delete condition targets only PK columns.
// Uses the Phase018 IR IsPK flag for accurate detection, with fallback
// to key name heuristic when IsPK is not set.
func isDeleteByPK(args []ir.FieldArg) bool {
	if len(args) == 0 {
		return true
	}
	for _, a := range args {
		if a.IsPK {
			return true
		}
	}
	// Fallback: check key name heuristic.
	for _, a := range args {
		key := resolveArgKey(a)
		if key == "id" {
			return true
		}
	}
	return false
}
