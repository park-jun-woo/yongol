//ff:func feature=gen-ir type=util control=iteration dimension=1
//ff:what findTablePK -- DDL 테이블명으로 첫 번째 primary key 컬럼명 조회

package ir

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// findTablePK returns the first primary key column name for the given DDL
// table name. Returns empty string if the table is not found or has no PK.
func findTablePK(fs *yongol.Fullstack, tableName string) string {
	for _, t := range fs.DDLTables {
		if !strings.EqualFold(t.Name, tableName) {
			continue
		}
		if len(t.PrimaryKey) > 0 {
			return t.PrimaryKey[0]
		}
		return ""
	}
	return ""
}
