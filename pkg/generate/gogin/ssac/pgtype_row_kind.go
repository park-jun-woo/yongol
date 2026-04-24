//ff:type feature=gen-gogin type=model
//ff:what pgtypeRowKind — sqlc pgx/v5 row 필드의 pgtype 분류 enum

package ssac

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
