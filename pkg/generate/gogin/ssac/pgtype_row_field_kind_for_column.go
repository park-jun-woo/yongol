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
func pgtypeRowFieldKindForColumn(tables []ddl.Table, tableModelName, columnName string) pgtypeRowKind {
	tbl := findDDLTableByModelName(tables, tableModelName)
	if tbl == nil {
		return pgPrimitive
	}
	lower := strings.ToLower(columnName)
	goType, ok := tbl.Columns[lower]
	if !ok {
		return pgPrimitive
	}
	notNull := tbl.NotNullCols[lower]
	// ddl.Table.Columns values are Go types assigned by pg_type_to_go.go
	// (pre-pgx/v5 mapping). For pgx/v5 the canonical decision is on the
	// PostgreSQL type family, but the parser stores the Go-type projection
	// rather than the raw SQL type. Recover the SQL family via the Go-type
	// projection — it is lossless for the families we care about (time,
	// uuid, numeric) thanks to the one-to-one mapping in pgTypeToGo.
	return classifyGoTypeProjection(goType, notNull)
}
