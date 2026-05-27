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

	// Build pagination suffix from PaginationArgs.
	var pagSuffix string
	for _, pa := range op.PaginationArgs {
		key := resolveArgKey(pa)
		val := renderArgValue(pa)
		switch key {
		case "per_page", "limit":
			pagSuffix += fmt.Sprintf(".limit(%s)", val)
		case "page_offset", "offset":
			pagSuffix += fmt.Sprintf(".offset(%s)", val)
		}
	}

	if op.IsList {
		b.WriteString(fmt.Sprintf("%sresult = await %s.execute(select(%s)%s%s)\n",
			indent, sessionRef, model, where, pagSuffix))
		b.WriteString(fmt.Sprintf("%s%s = result.scalars().all()\n",
			indent, op.VarName))
	} else {
		b.WriteString(fmt.Sprintf("%sresult = await %s.execute(select(%s)%s%s)\n",
			indent, sessionRef, model, where, pagSuffix))
		b.WriteString(fmt.Sprintf("%s%s = result.scalars().first()\n",
			indent, op.VarName))
	}
}
