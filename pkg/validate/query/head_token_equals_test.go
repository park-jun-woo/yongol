//ff:func feature=validate type=test control=iteration dimension=1 topic=query-structural
//ff:what headTokenEquals — head 토큰 추출 + 대소문자 무시 비교 (배열/파라미터/공백/정규화) 전분기 검증

package query

import "testing"

func TestHeadTokenEquals(t *testing.T) {
	cases := []struct {
		name string
		raw  string
		want string
		ok   bool
	}{
		// exact match
		{name: "exact_match", raw: "UUID", want: "UUID", ok: true},
		// case insensitive
		{name: "case_insensitive", raw: "uuid", want: "UUID", ok: true},
		// array suffix stripped
		{name: "array_suffix", raw: "UUID[]", want: "UUID", ok: true},
		// array suffix with spaces
		{name: "array_suffix_spaces", raw: "  UUID[]  ", want: "UUID", ok: true},
		// parenthesized parameter stripped
		{name: "paren_stripped", raw: "NUMERIC(10,2)", want: "NUMERIC", ok: true},
		// multi-word normalization: TIMESTAMP WITH TIME ZONE -> TIMESTAMPTZ
		{name: "multi_word_timestamptz", raw: "TIMESTAMP WITH TIME ZONE", want: "TIMESTAMPTZ", ok: true},
		// multi-word normalization: TIMESTAMP WITHOUT TIME ZONE -> TIMESTAMP
		{name: "multi_word_timestamp", raw: "TIMESTAMP WITHOUT TIME ZONE", want: "TIMESTAMP", ok: true},
		// multi-word normalization: DOUBLE PRECISION -> FLOAT8
		{name: "multi_word_float8", raw: "DOUBLE PRECISION", want: "FLOAT8", ok: true},
		// multi-word with parameter: CHARACTER VARYING(255) -> VARCHAR
		{name: "multi_word_varchar", raw: "CHARACTER VARYING(255)", want: "VARCHAR", ok: true},
		// multi-word: CHARACTER(10) -> CHAR
		{name: "multi_word_char", raw: "CHARACTER(10)", want: "CHAR", ok: true},
		// no match
		{name: "no_match", raw: "TEXT", want: "UUID", ok: false},
		// empty string
		{name: "empty_raw", raw: "", want: "UUID", ok: false},
		// whitespace only
		{name: "whitespace_only", raw: "   ", want: "UUID", ok: false},
		// array + param combo: [] stripped first, then paren stripped
		{name: "array_and_paren", raw: "NUMERIC(18)[]", want: "NUMERIC", ok: true},
		// leading/trailing whitespace
		{name: "leading_trailing_space", raw: "  INET  ", want: "INET", ok: true},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := headTokenEquals(c.raw, c.want)
			if got != c.ok {
				t.Errorf("headTokenEquals(%q, %q) = %v, want %v", c.raw, c.want, got, c.ok)
			}
		})
	}
}
