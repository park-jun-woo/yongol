//ff:func feature=gen-nestjs type=util control=sequence
//ff:what renderDeleteOp — DeleteOp → Prisma delete/deleteMany TypeScript 문 렌더링

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// renderDeleteOp writes a Prisma delete or deleteMany call. When the where
// condition is on a non-PK column (not "id"), deleteMany is used.
func renderDeleteOp(b *strings.Builder, op *ir.DeleteOp, indent, prismaRef string) {
	if op == nil {
		return
	}
	model := lcFirst(op.Model)
	args := renderPrismaWhere(op.Args)
	method := "delete"
	if !isDeleteByPK(op.Args) {
		method = "deleteMany"
	}
	b.WriteString(fmt.Sprintf("%sawait %s.%s.%s(%s);\n",
		indent, prismaRef, model, method, args))
}
