//ff:func feature=validate type=util control=sequence topic=query-structural
//ff:what isNumericColumn — Column.RawType 의 head 가 NUMERIC/DECIMAL 인지 판정 (Q-13 filter)

package query

import "github.com/park-jun-woo/yongol/pkg/parser/ddl"

// isNumericColumn matches NUMERIC / DECIMAL columns regardless of
// precision/scale parameters or nullability.
func isNumericColumn(col ddl.Column) bool {
	return headTokenEquals(col.RawType, "NUMERIC") || headTokenEquals(col.RawType, "DECIMAL")
}
