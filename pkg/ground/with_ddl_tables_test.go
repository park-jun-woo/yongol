//ff:func feature=rule type=test-helper control=sequence
//ff:what withDDLTables — DDL 테이블 슬라이스를 Fullstack.DDLTables 에 append 하는 option

package ground

import (
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// withDDLTables attaches DDL tables.
func withDDLTables(tables ...ddl.Table) func(*yongol.Fullstack) {
	return func(fs *yongol.Fullstack) { fs.DDLTables = append(fs.DDLTables, tables...) }
}
