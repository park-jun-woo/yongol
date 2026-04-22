//ff:func feature=migration type=test control=iteration dimension=1
//ff:what NormalizeDefault — DEFAULT 표현 정규화 테이블 테스트
package migration

import "testing"

func TestNormalizeDefault(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"", ""},
		{"'foo'", "'foo'"},
		{"'foo'::text", "'foo'"},
		{"'foo'::character varying", "'foo'"},
		{"'foo'::varchar(255)", "'foo'::varchar(255)"}, // paren parameter keeps cast (not a bare ident)
		{"NOW()", "CURRENT_TIMESTAMP"},
		{"now()", "CURRENT_TIMESTAMP"},
		{"CURRENT_TIMESTAMP", "CURRENT_TIMESTAMP"},
		{"TRUE", "TRUE"},
		{"true", "TRUE"},
		{"FALSE", "FALSE"},
		{"false", "FALSE"},
		{"NULL", "NULL"},
		{"0", "0"},
		{"0::integer", "0"},
		{"42::bigint", "42"},
		{"'draft'::text", "'draft'"},
		{"nextval('users_id_seq'::regclass)", "nextval('users_id_seq')"},
		{"nextval('users_id_seq')", "nextval('users_id_seq')"},
	}
	for _, c := range cases {
		c := c
		t.Run(c.in, func(t *testing.T) {
			if got := NormalizeDefault(c.in); got != c.want {
				t.Errorf("got %q, want %q", got, c.want)
			}
		})
	}
}
