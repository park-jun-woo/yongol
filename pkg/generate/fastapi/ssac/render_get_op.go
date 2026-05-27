//ff:func feature=gen-fastapi type=util control=sequence
//ff:what renderGetOp — GetOp → SQLAlchemy select 쿼리 Python 문 렌더링 (변수 섀도잉 방지)

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// renderGetOp writes an SQLAlchemy async select query. When the result
// variable name collides with a service method parameter, a "_result"
// suffix is appended to avoid shadowing.
func renderGetOp(b *strings.Builder, op *ir.GetOp, indent, sessionRef string) {
	if op == nil {
		return
	}
	model := pascalCase(op.Model)
	where := renderSAWhere(op.Model, op.Args)
	varName := safeVarName(op.VarName)
	if op.IsList {
		b.WriteString(fmt.Sprintf("%sresult = await %s.execute(select(%s)%s)\n",
			indent, sessionRef, model, where))
		b.WriteString(fmt.Sprintf("%s%s = result.scalars().all()\n",
			indent, varName))
	} else {
		b.WriteString(fmt.Sprintf("%sresult = await %s.execute(select(%s)%s)\n",
			indent, sessionRef, model, where))
		b.WriteString(fmt.Sprintf("%s%s = result.scalars().first()\n",
			indent, varName))
	}
}
