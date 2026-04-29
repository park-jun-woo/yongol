//ff:func feature=validate type=test-helper control=sequence topic=ssac-sqlc
//ff:what workflowTable — XQS-20 테스트용 `workflows` DDL Table fixture (4 컬럼)

package ssac_sqlc

import "github.com/park-jun-woo/yongol/pkg/parser/ddl"

// workflowTable returns a minimal DDL Table for the `workflows` table
// (used by XQS-20 case 5). Four columns suffice for the full-RETURNING check.
func workflowTable() ddl.Table {
	return ddl.Table{
		Name: "workflows",
		Columns: map[string]ddl.Column{
			"id":     {Name: "id"},
			"name":   {Name: "name"},
			"status": {Name: "status"},
			"meta":   {Name: "meta"},
		},
		ColumnOrder: []string{"id", "name", "status", "meta"},
	}
}
