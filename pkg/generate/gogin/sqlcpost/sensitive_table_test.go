//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestGenerate — 빈 DDL/민감컬럼 유무/mkdir 에러 분기 + 파일 산출 검증
package sqlcpost

import (
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func sensitiveTable() ddl.Table {
	return ddl.Table{
		Name: "users",
		Columns: map[string]ddl.Column{
			"id":            {Name: "id", RawType: "UUID"},
			"password_hash": {Name: "password_hash", RawType: "TEXT", Sensitive: true},
		},
		ColumnOrder: []string{"id", "password_hash"},
	}
}
