//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what collectSensitiveKeysFromTable — 한 테이블의 @sensitive 컬럼명을 seen 맵에 누적

package boot

import "github.com/park-jun-woo/yongol/pkg/parser/ddl"

// collectSensitiveKeysFromTable accumulates names of `-- @sensitive`
// columns from one DDL table's column map into the seen set. Hoisted
// out of collectSensitiveKeys so the outer loop stays at depth 2.
func collectSensitiveKeysFromTable(cols map[string]ddl.Column, seen map[string]bool) {
	for col, c := range cols {
		if c.Sensitive {
			seen[col] = true
		}
	}
}
