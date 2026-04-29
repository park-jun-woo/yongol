//ff:type feature=gen-gogin type=model
//ff:what BindingKind — PG 컬럼이 매핑되는 Go 표현 분류 (8 종)

package types

// BindingKind classifies the Go-side representation of a parsed PostgreSQL
// column. The dispatcher in MapPGType picks one of these kinds; downstream
// emitters use Kind only for branching that cannot be expressed via the
// string templates on GoTypeBinding (e.g. import wiring).
type BindingKind int

const (
	// KindNative — sqlc emits a Go primitive directly: int64, float64,
	// string, bool. Applies to NOT NULL integer / float / string / boolean
	// columns whose pg type is natively supported by sqlc pgx/v5.
	KindNative BindingKind = iota

	// KindPointer — nullable native column. sqlc pgx/v5 still emits the
	// primitive Go type, but yongol wraps it as *T to mirror the OpenAPI
	// optional-field convention (`*string`, `*int64`).
	KindPointer

	// KindPgtype — sqlc emits a pgtype wrapper (pgtype.UUID, pgtype.Numeric,
	// pgtype.Timestamptz, pgtype.Inet, pgtype.Interval). Convert sites must
	// inspect `.Valid` and extract the inner field (`.Bytes`, `.Time`, …).
	KindPgtype

	// KindJSONB — column stores JSON. NOT NULL → map[string]any, nullable →
	// *map[string]any. INSERT path serialises via json.Marshal to []byte;
	// response path unmarshals before emitting.
	KindJSONB

	// KindBytea — sqlc emits []byte natively (pgx/v5). Slices are themselves
	// nullable so no pointer wrapping is required. JSON responses emit
	// base64 via the standard json encoder.
	KindBytea

	// KindArray — PG array column (TEXT[], BIGINT[], …). Element binding is
	// derived from the base type; the slice itself is nullable so no extra
	// pointer wrapping is required.
	KindArray

	// KindEnum — VARCHAR(N) plus CHECK IN (...) constraint. Stored as Go
	// string + apiCast; nullable variant is *string + apiCast.
	KindEnum

	// KindUnsupported — column rejected at codegen time. validate D-11 emits
	// an ERROR before generate ever runs; this kind exists so emit sites
	// fail fast instead of silently fall back to a wrong representation.
	KindUnsupported
)

// String renders the kind for diagnostic / debug output.
func (k BindingKind) String() string {
	switch k {
	case KindNative:
		return "Native"
	case KindPointer:
		return "Pointer"
	case KindPgtype:
		return "Pgtype"
	case KindJSONB:
		return "JSONB"
	case KindBytea:
		return "Bytea"
	case KindArray:
		return "Array"
	case KindEnum:
		return "Enum"
	case KindUnsupported:
		return "Unsupported"
	}
	return "Unknown"
}
