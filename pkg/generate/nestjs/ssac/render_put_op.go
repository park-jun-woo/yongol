//ff:func feature=gen-nestjs type=util control=sequence
//ff:what renderPutOp — PutOp → Prisma update TypeScript 문 렌더링

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// renderPutOp writes a Prisma update call. SSaC args are split into where
// (PK fields) and data (remaining fields) clauses.
func renderPutOp(b *strings.Builder, op *ir.PutOp, indent, prismaRef string) {
	if op == nil {
		return
	}
	model := lcFirst(op.Model)
	where, data := splitWhereData(op.Args)
	b.WriteString(fmt.Sprintf("%sawait %s.%s.update({ where: { %s }, data: { %s } });\n",
		indent, prismaRef, model, where, data))
}
