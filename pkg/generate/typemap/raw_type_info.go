//ff:type feature=gen-typemap type=model
//ff:what RawTypeInfo — DDL 원본 타입 문자열의 정규화 결과 (head/param/array/multi-token)

package typemap

// RawTypeInfo captures the normalised view of a DDL column type string.
//
// The DDL parser preserves the raw token verbatim (e.g. "BIGINT",
// "TEXT[]", "VARCHAR(255)", "NUMERIC(10,2)", "DOUBLE PRECISION",
// "TIMESTAMP WITH TIME ZONE"). ParseRawType splits off the array marker
// and the parameter list, upper-cases the head token, and reports
// whether the head is a multi-word phrase.
type RawTypeInfo struct {
	// Head is the upper-cased base type with array marker and parameter
	// list stripped (e.g. "BIGINT", "VARCHAR", "NUMERIC", "TIMESTAMPTZ").
	Head string

	// Param is the verbatim parameter list (e.g. "255", "10,2").
	// Empty when no parenthesised parameters were present.
	Param string

	// IsArray is true when the raw type ends with "[]".
	IsArray bool

	// MultiToken is true when the head contains whitespace after stripping
	// — i.e. a multi-word PG type name like "DOUBLE PRECISION".
	MultiToken bool
}
