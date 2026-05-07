//ff:func feature=gen-gogin type=util control=sequence
//ff:what lookupDDLColumn — (table, column) → *ddl.Column 조회 (없으면 nil)

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/util/caseconv"
)

// lookupDDLColumn resolves (table, column) → *ddl.Column using DDLTables
// as the source of truth. Returns nil when the table or column cannot be
// found — convert sites then fall back to direct row.<Field> assignment
// (matches the historic behaviour for api wrapper schemas with no
// backing DDL table).
//
// Phase001 — replaces pgtypeRowFieldKindForColumn. Emit sites previously
// translated the column to a small kind enum (pgPrimitive / pgTimestamp
// / pgUUID / pgNumeric / pgTextWrapper); they now consume the parsed
// column directly and route through pkg/generate/gogin/types for the
// canonical Go-side expressions.
func lookupDDLColumn(tables []ddl.Table, tableModelName, columnName string) *ddl.Column {
	tbl := findDDLTableByModelName(tables, tableModelName)
	if tbl == nil {
		return nil
	}
	lower := caseconv.PascalToSnake(columnName)
	c, ok := tbl.Columns[lower]
	if !ok {
		return nil
	}
	return &c
}
