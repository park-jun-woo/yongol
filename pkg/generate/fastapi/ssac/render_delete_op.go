//ff:func feature=gen-fastapi type=util control=sequence
//ff:what renderDeleteOp — DeleteOp → SQLAlchemy delete Python 문 렌더링

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// renderDeleteOp writes an SQLAlchemy async delete statement.
func renderDeleteOp(b *strings.Builder, op *ir.DeleteOp, indent, sessionRef string) {
	if op == nil {
		return
	}
	model := pascalCase(op.Model)
	where := renderSAWhere(op.Model, op.Args)
	b.WriteString(fmt.Sprintf("%sawait %s.execute(delete(%s)%s)\n",
		indent, sessionRef, model, where))
}
