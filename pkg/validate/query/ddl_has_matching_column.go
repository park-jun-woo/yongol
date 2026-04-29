//ff:func feature=validate type=util control=iteration dimension=2 topic=query-structural
//ff:what ddlHasMatchingColumn — DDL 테이블 전체에서 filter 통과 컬럼이 1 개 이상인지 판정

package query

import "github.com/park-jun-woo/yongol/pkg/parser/ddl"

// ddlHasMatchingColumn returns true when at least one column across all
// parsed tables passes filter. Used by checkPgtypeOverride to gate
// per-type Q-NN rules — projects without matching columns skip the
// override check entirely.
func ddlHasMatchingColumn(tables []ddl.Table, filter func(col ddl.Column) bool) bool {
	if filter == nil {
		return false
	}
	for _, tbl := range tables {
		for _, col := range tbl.Columns {
			if filter(col) {
				return true
			}
		}
	}
	return false
}
