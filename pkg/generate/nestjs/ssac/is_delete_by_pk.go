//ff:func feature=gen-nestjs type=util control=iteration dimension=1
//ff:what isDeleteByPK — delete 조건이 PK 컬럼인지 확인

package ssac

import "github.com/park-jun-woo/yongol/pkg/generate/ir"

// isDeleteByPK checks if the delete condition targets only PK columns.
func isDeleteByPK(args []ir.FieldArg) bool {
	if len(args) == 0 {
		return true
	}
	for _, a := range args {
		key := resolveArgKey(a)
		if key == "id" {
			return true
		}
	}
	return false
}
