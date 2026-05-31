//ff:func feature=gen-nestjs type=util control=iteration dimension=1
//ff:what buildPaginationOpts — PaginationArgs → Prisma take/skip/cursor 옵션 목록

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// buildPaginationOpts maps each pagination arg to its Prisma option string
// (take/skip/cursor, or a passthrough key: value).
func buildPaginationOpts(pagArgs []ir.FieldArg) []string {
	var opts []string
	for _, pa := range pagArgs {
		opts = append(opts, paginationOpt(resolveArgKey(pa), renderArgValue(pa)))
	}
	return opts
}
