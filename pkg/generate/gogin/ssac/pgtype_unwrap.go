//ff:func feature=gen-gogin type=util control=selection topic=pgtype-unwrap
//ff:what pgtypeUnwrapExpr — sqlc pgx/v5 row 필드의 언래핑 표현식 선택 (convert 전용)

package ssac

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// pgtypeRowKind classifies the sqlc pgx/v5 Go type emitted for a DDL column.
// The convert emitter uses this to decide whether a row.<Field> read needs
// unwrapping before assignment to the api struct (which expects primitive
// Go types matching oapi-codegen output).
type pgtypeRowKind int

const (
	// pgPrimitive — sqlc emits a Go primitive (int64, int32, int16, string,
	// bool, []byte). Assignment is direct.
	pgPrimitive pgtypeRowKind = iota
	// pgTimestamp — sqlc emits pgtype.Timestamp / .Timestamptz / .Date. The
	// wrapper exposes the value via `.Time`.
	pgTimestamp
	// pgUUID — sqlc emits pgtype.UUID. Convert goes through the
	// pgUUIDToString helper emitted into internal/service/pgtype_helpers.go
	// (NOT NULL) or pgUUIDToStringPtr (nullable).
	pgUUID
	// pgTextWrapper — sqlc emits pgtype.Text for nullable VARCHAR / TEXT.
	// Unwrap via `.String`.
	pgTextWrapper
	// pgNumeric — sqlc emits pgtype.Numeric. Deferred: we emit a helper
	// that calls `.Value()` and stringifies. For NOT NULL numeric columns
	// the api surface is usually string/float anyway so the helper returns
	// "" on invalid.
	pgNumeric
	// pgUnknown — column resolved but none of the known pgtype wrappers
	// applied. Falls back to direct assignment; if sqlc actually emitted a
	// wrapper the build fails (same as before the refit) and the author
	// adds the mapping here.
	pgUnknown
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

// classifyGoTypeProjection reverses pkg/parser/ddl/pg_type_to_go.go. The
// parser projected SQL types into a small Go-type alphabet; this function
// reads that projection + nullability and returns the pgx/v5 row-field
// classification. Kept separate from classifyPgtypeRowField so the two
// inference paths (raw SQL type vs. projected Go type) stay testable.
func classifyGoTypeProjection(goType string, notNull bool) pgtypeRowKind {
	switch strings.TrimSpace(goType) {
	case "time.Time":
		return pgTimestamp
	case "json.RawMessage":
		return pgPrimitive
	case "int64", "int32", "int16", "bool", "float64":
		if notNull {
			return pgPrimitive
		}
		return pgUnknown
	case "string":
		// UUID and VARCHAR / TEXT both project to "string" in pgTypeToGo,
		// but only UUID gets a pgtype wrapper in pgx/v5 NOT NULL mode.
		// Without the raw SQL type we conservatively treat NOT NULL
		// string as primitive (VARCHAR / TEXT case). UUID columns are
		// distinguished via the fallback path in pickConvertRHS using
		// apiCastFor(openapi_types.UUID).
		if notNull {
			return pgPrimitive
		}
		return pgTextWrapper
	}
	return pgPrimitive
}

// findDDLTableByModelName looks up a DDL table by its sqlc model name. The
// parser stores the raw table name ("workflows"); the convert emitter
// receives the PascalCase sqlc model name ("Workflow"). We normalise both
// sides to the same lower-snake singular form ("workflow") for comparison.
func findDDLTableByModelName(tables []ddl.Table, modelName string) *ddl.Table {
	target := ddlTableSingular(pascalToSnake(modelName))
	for i := range tables {
		t := &tables[i]
		lower := strings.ToLower(t.Name)
		if ddlTableSingular(lower) == target {
			return t
		}
	}
	return nil
}

// pascalToSnake converts "ExecutionLog" → "execution_log" for comparison
// with DDL table names. Rolls over the minimal alphabet we actually see in
// sqlc model names (ASCII, no digits leading) — for anything richer we
// defer to dedicated case converters elsewhere in the codebase.
func pascalToSnake(s string) string {
	var out strings.Builder
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			out.WriteByte('_')
		}
		if r >= 'A' && r <= 'Z' {
			out.WriteRune(r + ('a' - 'A'))
		} else {
			out.WriteRune(r)
		}
	}
	return out.String()
}

// ddlTableSingular desingularises a lower-snake table name to the sqlc
// model name lower form. Kept simple — matches inflection.Singular on the
// zenflow fixture (users / organizations / workflows / actions /
// execution_logs).
func ddlTableSingular(name string) string {
	switch {
	case strings.HasSuffix(name, "ies"):
		return name[:len(name)-3] + "y"
	case strings.HasSuffix(name, "sses"):
		return name[:len(name)-2]
	case strings.HasSuffix(name, "xes"):
		return name[:len(name)-2]
	case strings.HasSuffix(name, "s") && !strings.HasSuffix(name, "ss"):
		return name[:len(name)-1]
	default:
		return name
	}
}

// pgtypeRowUnwrap returns the Go expression that extracts the primitive
// value from a sqlc pgx/v5 row field. rowAccess is the full selector (e.g.
// "row.CreatedAt"). Returns ("", false) when no unwrap is required (the
// primitive path — caller emits rowAccess verbatim).
func pgtypeRowUnwrap(kind pgtypeRowKind, rowAccess string, apiCast string) (string, bool) {
	switch kind {
	case pgTimestamp:
		return rowAccess + ".Time", true
	case pgTextWrapper:
		return rowAccess + ".String", true
	case pgUUID:
		// pgUUIDToString helper centralises the Valid + [16]byte → canonical
		// UUID string conversion. apiCast (e.g. openapi_types.UUID) is applied
		// by the caller on top.
		_ = apiCast
		return "pgUUIDToString(" + rowAccess + ")", true
	case pgNumeric:
		return "pgNumericToString(" + rowAccess + ")", true
	}
	return "", false
}

// pgtypeHelpersEmit returns the Go source for the pgtype helper file. The
// file is emitted into internal/service/pgtype_helpers.go by
// emitPgtypeHelpers. Kept as a single function (no template indirection) so
// the output is deterministic across runs.
func pgtypeHelpersEmit() string {
	return `//ff:func feature=service type=util control=sequence topic=pgtype-unwrap
//ff:what pgtype helpers — pgtype.UUID / pgtype.Numeric 를 primitive string 으로 변환

package service

import (
	"encoding/hex"

	"github.com/jackc/pgx/v5/pgtype"
)

// pgUUIDToString returns the canonical 8-4-4-4-12 UUID string for a valid
// pgtype.UUID, or "" when the value is SQL NULL. The generated convert
// functions call this when mapping a sqlc UUID column (pgtype.UUID) onto
// an api struct field typed as string (openapi_types.UUID / plain string).
func pgUUIDToString(u pgtype.UUID) string {
	if !u.Valid {
		return ""
	}
	b := u.Bytes
	dst := make([]byte, 36)
	hex.Encode(dst[0:8], b[0:4])
	dst[8] = '-'
	hex.Encode(dst[9:13], b[4:6])
	dst[13] = '-'
	hex.Encode(dst[14:18], b[6:8])
	dst[18] = '-'
	hex.Encode(dst[19:23], b[8:10])
	dst[23] = '-'
	hex.Encode(dst[24:36], b[10:16])
	return string(dst)
}

// pgNumericToString renders a pgtype.Numeric as its textual form. Returns
// "" on NULL or marshal failure. Callers that need arithmetic precision
// should use pgtype.Numeric directly instead of going through convert.
func pgNumericToString(n pgtype.Numeric) string {
	if !n.Valid {
		return ""
	}
	buf, err := n.MarshalJSON()
	if err != nil {
		return ""
	}
	// MarshalJSON wraps numeric values in quotes only when the driver chose
	// string serialisation; strip them for direct string consumption.
	s := string(buf)
	if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
		return s[1 : len(s)-1]
	}
	return s
}
`
}
