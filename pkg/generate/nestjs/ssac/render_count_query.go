//ff:func feature=gen-nestjs type=util control=sequence
//ff:what renderCountQuery — GetOp(IsCount) → Prisma count({ where }) 호출 렌더링

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// renderCountQuery writes a Prisma count() call. The where clause is omitted
// when the op has no args.
func renderCountQuery(b *strings.Builder, op *ir.GetOp, indent, prismaRef, model string) {
	whereParts := buildWhereParts(op.Args)
	if len(whereParts) > 0 {
		b.WriteString(fmt.Sprintf("%sconst %s = await %s.%s.count({ where: { %s } });\n",
			indent, op.VarName, prismaRef, model, strings.Join(whereParts, ", ")))
		return
	}
	b.WriteString(fmt.Sprintf("%sconst %s = await %s.%s.count();\n",
		indent, op.VarName, prismaRef, model))
}
