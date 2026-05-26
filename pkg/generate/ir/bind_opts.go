//ff:type feature=gen-ir type=model
//ff:what BindOpts — PGFamily 외에 타입 바인딩에 필요한 DDL 컬럼 메타 (NotNull/Default/Array/Enum)

package ir

// BindOpts carries the column-level options that influence type binding
// beyond the PGFamily classification. These come from DDL column metadata
// (NOT NULL, DEFAULT, CHECK IN, array marker).
type BindOpts struct {
	// NotNull is true when the column is effectively NOT NULL.
	NotNull bool

	// DefaultLiteral is the DDL column's DEFAULT value literal.
	DefaultLiteral string

	// IsArray is true when the column type ends with [].
	IsArray bool

	// ElementHead is the normalised head token of the array element
	// (e.g. "TEXT" for TEXT[], "BIGINT" for BIGINT[]). Only meaningful
	// when IsArray is true.
	ElementHead string

	// EnumValues holds the CHECK IN (...) enum values extracted from the
	// DDL column. Non-nil and non-empty signals enum-ness.
	EnumValues []string
}
