//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what tableHasSensitiveColumn — DDL 테이블에 @sensitive 컬럼이 하나라도 있는지

package sqlcpost

import "github.com/park-jun-woo/yongol/pkg/parser/ddl"

// tableHasSensitiveColumn reports whether any column on the given DDL
// table is annotated with `-- @sensitive`. Used by Generate to decide
// whether to emit a <table>_log.go file.
func tableHasSensitiveColumn(t ddl.Table) bool {
	for _, c := range t.Columns {
		if c.Sensitive {
			return true
		}
	}
	return false
}
