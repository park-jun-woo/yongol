//ff:func feature=validate type=util control=iteration dimension=1 topic=manifest-infra
//ff:what findUserTable — fs.DDLTables 에서 이름 일치하는 ddl.Table 을 반환 (없으면 nil)

package manifest_ddl

import (
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// findUserTable returns the parsed DDL Table whose Name equals userTable, or
// nil when not found. Empty userTable always yields nil — callers that need
// the "missing user_table" diagnostic should gate on XDN-01 first.
func findUserTable(fs *yongol.Fullstack, userTable string) *ddl.Table {
	if fs == nil || userTable == "" {
		return nil
	}
	for i := range fs.DDLTables {
		if fs.DDLTables[i].Name == userTable {
			return &fs.DDLTables[i]
		}
	}
	return nil
}
