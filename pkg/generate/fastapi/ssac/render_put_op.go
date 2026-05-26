//ff:func feature=gen-fastapi type=util control=sequence
//ff:what renderPutOp — PutOp → SQLAlchemy update Python 문 렌더링

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// renderPutOp writes an SQLAlchemy async update statement.
func renderPutOp(b *strings.Builder, op *ir.PutOp, indent, sessionRef string) {
	if op == nil {
		return
	}
	model := pascalCase(op.Model)
	b.WriteString(fmt.Sprintf("%sawait %s.execute(\n", indent, sessionRef))
	b.WriteString(fmt.Sprintf("%s    update(%s).where(%s.id == params[\"id\"]).values(**body)\n",
		indent, model, model))
	b.WriteString(fmt.Sprintf("%s)\n", indent))
}
