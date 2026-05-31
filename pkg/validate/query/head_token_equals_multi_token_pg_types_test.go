//ff:func feature=validate type=test control=iteration dimension=1 topic=query-structural
//ff:what headTokenEquals — head 토큰 추출 + 대소문자 무시 비교 (배열/파라미터/공백/정규화) 전분기 검증
package query

import (
	"testing"
)

func TestHeadTokenEquals_MultiTokenPGTypes(t *testing.T) {
	cases := []struct {
		name string
		raw  string
		want string
		ok   bool
	}{
		{name: "single_token_match", raw: "TIMESTAMPTZ", want: "TIMESTAMPTZ", ok: true},
		{name: "single_token_lowercase", raw: "timestamptz", want: "TIMESTAMPTZ", ok: true},
		{name: "multi_token_timestamp_with_tz", raw: "TIMESTAMP WITH TIME ZONE", want: "TIMESTAMPTZ", ok: true},
		{name: "multi_token_timestamp_without_tz", raw: "TIMESTAMP WITHOUT TIME ZONE", want: "TIMESTAMP", ok: true},
		{name: "multi_token_double_precision", raw: "DOUBLE PRECISION", want: "FLOAT8", ok: true},
		{name: "multi_token_character_varying_with_param", raw: "CHARACTER VARYING(255)", want: "VARCHAR", ok: true},
		{name: "multi_token_character_with_param", raw: "CHARACTER(10)", want: "CHAR", ok: true},
		{name: "no_match_unrelated", raw: "BIGINT", want: "TIMESTAMPTZ", ok: false},
		{name: "array_marker_stripped", raw: "TEXT[]", want: "TEXT", ok: true},
		{name: "param_stripped_uuid", raw: "UUID", want: "UUID", ok: true},
		{name: "lowercase_multi_token", raw: "timestamp with time zone", want: "TIMESTAMPTZ", ok: true},
		{name: "multi_token_timestamp_collision_no_tz", raw: "TIMESTAMP WITH TIME ZONE", want: "TIMESTAMP", ok: false},
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
