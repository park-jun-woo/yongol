//ff:func feature=validate type=util control=iteration dimension=2 topic=openapi-ddl
//ff:what findDDLColumnConstraints — 컬럼이 속한 DDL 테이블에서 VARCHAR/CHECK 제약 추출

package openapi_ddl

import "github.com/park-jun-woo/yongol/pkg/yongol"

// findDDLColumnConstraints scans all DDL tables for a column named col and
// returns its VARCHAR length and CHECK enum values from the first match.
func findDDLColumnConstraints(fs *yongol.Fullstack, col string) (varcharLen int, checkEnums []string, found bool) {
	for _, t := range fs.DDLTables {
		if _, ok := t.Columns[col]; !ok {
			continue
		}
		found = true
		if t.VarcharLen != nil {
			if n, ok := t.VarcharLen[col]; ok {
				varcharLen = n
			}
		}
		if t.CheckEnums != nil {
			if vals, ok := t.CheckEnums[col]; ok {
				checkEnums = vals
			}
		}
		return
	}
	return
}
