//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-ddl
//ff:what isSqlcRowType — fs.SQLcQueries[*].RowType 조회 (sqlc 합성 row struct 매칭)

package ssac_ddl

import "github.com/park-jun-woo/yongol/pkg/yongol"

// isSqlcRowType reports whether typeName is a sqlc-synthesized row struct
// exposed via fs.SQLcQueries[*].RowType. Enables JOIN queries to supply
// their own result type without a manual DTO declaration.
func isSqlcRowType(fs *yongol.Fullstack, typeName string) bool {
	if fs == nil || typeName == "" {
		return false
	}
	for _, q := range fs.SQLcQueries {
		if q.RowType == typeName {
			return true
		}
	}
	return false
}
