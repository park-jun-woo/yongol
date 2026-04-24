//ff:func feature=gen-gogin type=util control=selection
//ff:what classifyPgtypeRowField — DDL 타입/NOT NULL → pgtypeRowKind

package ssac

import (
	"strings"
)

// classifyPgtypeRowField inspects the DDL type string + nullability for a
// single column and returns the pgtypeRowKind sqlc pgx/v5 would emit.
// Inputs mirror ddl.Table.Columns semantics: ddlType is the uppercased
// PostgreSQL type token (e.g. "TIMESTAMP", "VARCHAR", "UUID"). Nullability
// is the caller's responsibility — passing notNull=false forces the
// Nullable-side mapping where sqlc pgx/v5 always emits a pgtype wrapper
// regardless of the base type.
func classifyPgtypeRowField(ddlType string, notNull bool) pgtypeRowKind {
	t := strings.ToUpper(strings.TrimSpace(ddlType))
	// Strip VARCHAR(n) / CHAR(n) lengths.
	if idx := strings.Index(t, "("); idx > 0 {
		t = t[:idx]
	}
	// NOT NULL path — sqlc emits pgtype wrappers only for timestamp / uuid
	// / date / numeric. Other NOT NULL columns get native Go primitives.
	if notNull {
		switch t {
		case "TIMESTAMP", "TIMESTAMPTZ", "DATE":
			return pgTimestamp
		case "UUID":
			return pgUUID
		case "NUMERIC", "DECIMAL":
			return pgNumeric
		case "BIGINT", "BIGSERIAL", "INTEGER", "INT", "INT4", "INT8",
			"SMALLINT", "INT2", "SERIAL", "SMALLSERIAL",
			"VARCHAR", "TEXT", "CHAR", "BPCHAR",
			"BOOLEAN", "BOOL",
			"BYTEA",
			"JSONB", "JSON",
			"REAL", "DOUBLE", "FLOAT", "FLOAT4", "FLOAT8":
			return pgPrimitive
		}
		return pgUnknown
	}
	// Nullable path — every nullable base type yields a pgtype wrapper.
	switch t {
	case "TIMESTAMP", "TIMESTAMPTZ", "DATE":
		return pgTimestamp
	case "UUID":
		return pgUUID
	case "NUMERIC", "DECIMAL":
		return pgNumeric
	case "VARCHAR", "TEXT", "CHAR", "BPCHAR":
		return pgTextWrapper
	// Nullable integers / booleans map to pgtype.Int4/8/2/Bool but the
	// convert emitter currently does not drive those cases on zenflow; add
	// when the first failing fixture appears.
	}
	return pgUnknown
}
