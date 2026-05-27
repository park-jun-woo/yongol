//ff:func feature=gen-fastapi type=util control=sequence
//ff:what renderPutOp — PutOp → SQLAlchemy update Python 문 렌더링 (where/data 분리)

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// renderPutOp writes an SQLAlchemy async update statement. SSaC args are
// split into where (PK fields) and data (remaining fields) clauses.
func renderPutOp(b *strings.Builder, op *ir.PutOp, indent, sessionRef string) {
	if op == nil {
		return
	}
	model := pascalCase(op.Model)
	whereArgs, dataArgs := splitWhereData(op.Args)

	whereClause := renderSAWhere(op.Model, whereArgs)
	dataClause := renderSAData(dataArgs)

	b.WriteString(fmt.Sprintf("%sawait %s.execute(\n", indent, sessionRef))
	b.WriteString(fmt.Sprintf("%s    update(%s)%s.values(%s)\n",
		indent, model, whereClause, dataClause))
	b.WriteString(fmt.Sprintf("%s)\n", indent))
}
