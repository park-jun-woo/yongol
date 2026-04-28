//ff:func feature=gen-gogin type=util control=sequence
//ff:what pgtypeRowFieldKindForColumn — (table, column) → pgtypeRowKind 조회

package ssac

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// pgtypeRowFieldKindForColumn resolves (table, column) → pgtypeRowKind using
// DDLTables as the source of truth. Falls back to pgPrimitive when the
// table or column is not found — this matches the pre-refit behaviour where
// convert emission assumes direct assignment.
//
// With Phase002 the parser preserves the raw PostgreSQL type token on
// Column.RawType, so we no longer need to round-trip through the Go-type
// projection (`classifyGoTypeProjection`). Direct dispatch on RawType +
// NotNull yields the same kind classification with strictly more fidelity
// (UUID vs VARCHAR, NUMERIC vs TEXT, etc).
func pgtypeRowFieldKindForColumn(tables []ddl.Table, tableModelName, columnName string) pgtypeRowKind {
	tbl := findDDLTableByModelName(tables, tableModelName)
	if tbl == nil {
		return pgPrimitive
	}
	lower := strings.ToLower(columnName)
	c, ok := tbl.Columns[lower]
	if !ok {
		return pgPrimitive
	}
	return classifyPgtypeRowField(c.RawType, c.NotNull)
}
