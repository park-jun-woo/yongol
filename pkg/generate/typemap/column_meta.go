//ff:type feature=gen-typemap type=model
//ff:what ColumnMeta — DDL 컬럼에서 family 분류에 필요한 최소 인터페이스

package typemap

// ColumnMeta abstracts the minimum information needed from a DDL column
// to classify its PGFamily. This decouples the typemap package from the
// concrete ddl.Column struct, allowing any DDL parser to provide the
// necessary fields.
type ColumnMeta interface {
	// RawType returns the verbatim DDL type string (e.g. "BIGINT",
	// "VARCHAR(255)", "TEXT[]").
	RawType() string

	// CheckEnum returns the CHECK constraint enum values extracted from
	// the column definition. Returns nil or empty when no CHECK IN (...)
	// is present.
	CheckEnum() []string
}
