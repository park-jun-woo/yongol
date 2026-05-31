//ff:func feature=gen-fastapi type=util control=sequence
//ff:what renderGetOp — GetOp → SQLAlchemy select + PaginationArgs → limit/offset 렌더링

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// renderGetOp writes an SQLAlchemy async select query. PaginationArgs
// from the Phase018 IR are rendered as SQLAlchemy .limit()/.offset()
// calls separate from the where clause. Variable shadowing is already
// resolved in the IR (Phase018), so VarName is used directly.
func renderGetOp(b *strings.Builder, op *ir.GetOp, indent, sessionRef string) {
	if op == nil {
		return
	}
	model := pascalCase(op.Model)
	where := renderSAWhere(op.Model, op.Args)
	pagSuffix := renderPaginationSuffix(op.PaginationArgs)
	scalar := "first"
	if op.IsList {
		scalar = "all"
	}
	b.WriteString(fmt.Sprintf("%sresult = await %s.execute(select(%s)%s%s)\n",
		indent, sessionRef, model, where, pagSuffix))
	b.WriteString(fmt.Sprintf("%s%s = result.scalars().%s()\n",
		indent, op.VarName, scalar))
}
