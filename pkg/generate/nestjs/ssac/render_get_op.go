//ff:func feature=gen-nestjs type=util control=sequence
//ff:what renderGetOp — GetOp → Prisma findUnique/findMany TypeScript 문 렌더링

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// renderGetOp writes a Prisma findUnique or findMany call.
func renderGetOp(b *strings.Builder, op *ir.GetOp, indent, prismaRef string) {
	if op == nil {
		return
	}
	model := lcFirst(op.Model)
	args := renderPrismaWhere(op.Args)
	method := "findUnique"
	if op.IsList {
		method = "findMany"
	}
	b.WriteString(fmt.Sprintf("%sconst %s = await %s.%s.%s(%s);\n",
		indent, op.VarName, prismaRef, model, method, args))
}
