//ff:func feature=gen-nestjs type=util control=iteration dimension=1
//ff:what buildWhereParts — FieldArg 목록 → Prisma where 절 "key: value" 조각 목록

package ssac

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// buildWhereParts renders each arg as a Prisma where clause entry "key: value".
func buildWhereParts(args []ir.FieldArg) []string {
	whereParts := make([]string, 0, len(args))
	for _, a := range args {
		whereParts = append(whereParts, fmt.Sprintf("%s: %s", resolveArgKey(a), renderArgValue(a)))
	}
	return whereParts
}
