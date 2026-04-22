//ff:func feature=gen-hurl type=util control=iteration dimension=1
//ff:what lookupCheckEnum — DDL CHECK enum에서 컬럼의 첫 번째 허용값 조회
package hurl

import (
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// lookupCheckEnum finds the first CHECK enum value for a column name across DDL tables.
func lookupCheckEnum(fs *yongol.Fullstack, colName string) string {
	if fs == nil {
		return ""
	}
	for _, tbl := range fs.DDLTables {
		if vals, ok := tbl.CheckEnums[colName]; ok && len(vals) > 0 {
			return vals[0]
		}
	}
	return ""
}
