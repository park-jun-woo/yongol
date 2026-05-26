//ff:func feature=gen-fastapi type=util control=sequence
//ff:what renderPostOp — PostOp → SQLAlchemy session.add Python 문 렌더링

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// renderPostOp writes an SQLAlchemy model instantiation and session add.
func renderPostOp(b *strings.Builder, op *ir.PostOp, indent, sessionRef string) {
	if op == nil {
		return
	}
	model := pascalCase(op.Model)
	data := renderSAData(op.Args)
	b.WriteString(fmt.Sprintf("%s%s = %s(%s)\n",
		indent, op.VarName, model, data))
	b.WriteString(fmt.Sprintf("%s%s.add(%s)\n",
		indent, sessionRef, op.VarName))
	b.WriteString(fmt.Sprintf("%sawait %s.flush()\n",
		indent, sessionRef))
}
