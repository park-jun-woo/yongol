//ff:func feature=gen-nestjs type=util control=sequence
//ff:what buildPrismaQueryOpts — GetOp → Prisma 쿼리 옵션(where + pagination) 목록 구성

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// buildPrismaQueryOpts assembles the Prisma query option list: the where clause
// from Args plus pagination options (take/skip/cursor) from PaginationArgs.
func buildPrismaQueryOpts(op *ir.GetOp) []string {
	var opts []string
	whereParts := buildWhereParts(op.Args)
	if len(whereParts) > 0 {
		opts = append(opts, fmt.Sprintf("where: { %s }", strings.Join(whereParts, ", ")))
	}
	opts = append(opts, buildPaginationOpts(op.PaginationArgs)...)
	return opts
}
