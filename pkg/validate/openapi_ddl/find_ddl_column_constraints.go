//ff:func feature=validate type=util control=iteration dimension=2 topic=openapi-ddl
//ff:what findDDLColumnConstraints — 컬럼이 속한 DDL 테이블에서 VARCHAR/CHECK 제약 추출

package openapi_ddl

import "github.com/park-jun-woo/yongol/pkg/yongol"

// findDDLColumnConstraints scans all DDL tables for a column named col and
// returns its VARCHAR length and CHECK enum values from the first match.
func findDDLColumnConstraints(fs *yongol.Fullstack, col string) (varcharLen int, checkEnums []string, found bool) {
	for _, t := range fs.DDLTables {
		c, ok := t.Columns[col]
		if !ok {
			continue
		}
		found = true
		varcharLen = c.VarcharLen
		checkEnums = c.CheckEnum
		return
	}
	return
}
