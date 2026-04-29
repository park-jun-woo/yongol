//ff:func feature=validate type=test-helper control=sequence topic=ssac-sqlc
//ff:what userTable — XQS-20 테스트용 `users` DDL Table fixture (6 컬럼)

package ssac_sqlc

import "github.com/park-jun-woo/yongol/pkg/parser/ddl"

// userTable returns a minimal DDL Table for the `users` table used across
// the XQS-20 test cases. Six columns mirror BUG-037's reproducer.
func userTable() ddl.Table {
	return ddl.Table{
		Name: "users",
		Columns: map[string]ddl.Column{
			"id":            {Name: "id"},
			"org_id":        {Name: "org_id"},
			"email":         {Name: "email"},
			"password_hash": {Name: "password_hash"},
			"role":          {Name: "role"},
			"claims":        {Name: "claims"},
		},
		ColumnOrder: []string{"id", "org_id", "email", "password_hash", "role", "claims"},
	}
}
