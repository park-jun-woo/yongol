//ff:func feature=gen-nestjs type=util control=sequence
//ff:what renderGetOp — GetOp → Prisma findUnique/findMany + PaginationArgs 렌더링

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// renderGetOp writes a Prisma findUnique or findMany call. PaginationArgs
// from the Phase018 IR are rendered as Prisma take/skip/cursor options
// separate from the where clause. Variable shadowing is already resolved
// in the IR (Phase018), so VarName is used directly.
func renderGetOp(b *strings.Builder, op *ir.GetOp, indent, prismaRef string) {
	if op == nil {
		return
	}
	model := lcFirst(op.Model)

	if op.IsCount {
		renderCountQuery(b, op, indent, prismaRef, model)
		return
	}

	opts := buildPrismaQueryOpts(op)
	argsStr := "{}"
	if len(opts) > 0 {
		argsStr = "{ " + strings.Join(opts, ", ") + " }"
	}
	method := "findUnique"
	if op.IsList {
		method = "findMany"
	}
	b.WriteString(fmt.Sprintf("%sconst %s = await %s.%s.%s(%s);\n",
		indent, op.VarName, prismaRef, model, method, argsStr))
}
