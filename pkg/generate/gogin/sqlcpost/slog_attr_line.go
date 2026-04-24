//ff:func feature=gen-gogin type=util control=selection
//ff:what slogAttrLine — 비민감 컬럼을 slog.Any 로 통일 emit (pgtype.* 호환)

package sqlcpost

import "fmt"

// slogAttrLine emits a slog.Any line for a non-sensitive column.
//
// We intentionally collapse every non-redacted field to slog.Any rather than
// dispatching on Go type (slog.String / slog.Time / slog.Int64 …). Reasons:
//
//  1. sqlc's `sql_package: pgx/v5` maps UUID / TIMESTAMP / NUMERIC / JSONB to
//     pgtype wrapper types (pgtype.UUID, pgtype.Timestamp, …). Those types
//     are NOT assignable to string / time.Time and blow up type-specific
//     constructors like slog.String / slog.Time at compile time (BUG-024).
//  2. slog.Any stores the value as KindAny and defers serialization to the
//     handler. yongol's default handler is JSON; pgtype wrapper types all
//     implement MarshalJSON, so output remains clean ("\"a1b2…\"",
//     "\"2026-04-24T…\"") without any codegen knowledge of sqlc's type
//     mapping.
//  3. Primitive types (string, int64, time.Time, bool, float64) are handled
//     transparently by slog.AnyValue — no regression for the database/sql
//     path.
//
// Trade-off: text-handler rendering of pgtype wrappers becomes uglier than
// with typed constructors, but JSON is the default output. Tracked for
// future refinement if users demand pretty text-handler output.
func slogAttrLine(colName, fieldName, goType string) string {
	_ = goType // intentionally unused — kept for signature stability
	return fmt.Sprintf("slog.Any(%q, r.%s)", colName, fieldName)
}
