//ff:func feature=gen-fastapi type=util control=sequence
//ff:what renderPutOp — PutOp → IsPK 기반 where/data 분리 SQLAlchemy update 렌더링

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// renderPutOp writes an SQLAlchemy async update statement. Args with IsPK ==
// true go to the where clause; the rest go to the data clause. This uses the
// Phase018 IR enrichment instead of heuristic PK detection.
func renderPutOp(b *strings.Builder, op *ir.PutOp, indent, sessionRef string) {
	if op == nil {
		return
	}
	model := pascalCase(op.Model)
	whereArgs, dataArgs := splitPKArgs(op.Args)
	whereClause := renderSAWhere(op.Model, whereArgs)
	dataClause := renderSAData(dataArgs)

	b.WriteString(fmt.Sprintf("%sawait %s.execute(\n", indent, sessionRef))
	b.WriteString(fmt.Sprintf("%s    update(%s)%s.values(%s)\n",
		indent, model, whereClause, dataClause))
	b.WriteString(fmt.Sprintf("%s)\n", indent))
}
