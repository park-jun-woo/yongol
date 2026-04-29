//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-sqlc
//ff:what returningCoversAllColumns — RETURNING 컬럼 set 이 DDL 컬럼 set 을 모두 덮는지

package ssac_sqlc

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// returningCoversAllColumns reports whether every DDL column key (lowercased)
// is present in the supplied returning-clause column set. Used by
// classifyReturningShape to distinguish ShapeFull (every column listed) from
// ShapePartial (strict subset).
func returningCoversAllColumns(cols map[string]bool, ddlCols map[string]ddl.Column) bool {
	for ddlCol := range ddlCols {
		if !cols[strings.ToLower(ddlCol)] {
			return false
		}
	}
	return true
}
