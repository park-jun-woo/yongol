//ff:type feature=gen-gogin type=model
//ff:what rawTypeInfo — Column.RawType 토큰을 family/길이/배열 마커로 분해한 정규화 결과

package types

// rawTypeInfo captures the normalised view of a Column.RawType string.
//
// The DDL parser preserves the raw token verbatim (Phase002) so callers
// observe values like "BIGINT", "TEXT[]", "VARCHAR(255)", "NUMERIC(10,2)",
// "DOUBLE PRECISION", and "TIMESTAMP WITH TIME ZONE". parseRawType
// splits off the array marker and the parameter list, upper-cases the
// head token, and reports whether the head is a multi-word phrase the
// parser was unable to fold into a single token.
type rawTypeInfo struct {
	// Head is the upper-cased base type with array marker and parameter
	// list stripped (e.g. "BIGINT", "VARCHAR", "NUMERIC", "TIMESTAMPTZ").
	Head string

	// Param is the verbatim parameter list (e.g. "(255)" → "255",
	// "(10,2)" → "10,2"). Empty when no parens were present.
	Param string

	// IsArray is true when the raw type ends with "[]".
	IsArray bool

	// MultiToken is true when the head contains whitespace after stripping
	// — i.e. a multi-word PG type name like "DOUBLE PRECISION" that the
	// parser preserved as-is. Such columns are unsupported in v1.
	MultiToken bool
}
