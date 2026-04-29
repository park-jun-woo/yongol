//ff:func feature=validate type=util control=selection topic=ssac-sqlc
//ff:what classifyReturningShape — RETURNING 컬럼 set ↔ DDL Column set 비교 → "full" / "partial" / "none"

package ssac_sqlc

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// classifyReturningShape inspects a RETURNING column-list string and the DDL
// table it targets, then reports the resulting sqlc emission shape.
//
//   - clause == ""          → ShapeNone
//   - clause == "*"         → ShapeFull
//   - clause column set ⊇ DDL columns → ShapeFull
//   - otherwise             → ShapePartial
//
// table may be nil (e.g. when the model name does not resolve to a known DDL
// table); in that case any non-empty / non-"*" clause is treated as partial,
// because we cannot prove fullness without a column set to compare against.
func classifyReturningShape(clause string, table *ddl.Table) ReturningShape {
	clause = strings.TrimSpace(clause)
	switch {
	case clause == "":
		return ShapeNone
	case clause == "*":
		return ShapeFull
	case table == nil || len(table.Columns) == 0:
		return ShapePartial
	case returningCoversAllColumns(splitReturningColumns(clause), table.Columns):
		return ShapeFull
	default:
		return ShapePartial
	}
}
