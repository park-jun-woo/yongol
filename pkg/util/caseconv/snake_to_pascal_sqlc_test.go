//ff:func feature=util type=test control=iteration dimension=1 topic=string-convert
//ff:what SnakeToPascalSqlc 회귀 테이블 테스트

package caseconv

import "testing"

func TestSnakeToPascalSqlc(t *testing.T) {
	cases := []struct{ in, want string }{
		{"id", "ID"},
		{"ids", "IDS"},
		{"org_id", "OrgID"},
		{"user_ids", "UserIDS"},
		{"per_page", "PerPage"},
		{"org_name", "OrgName"},
		{"email", "Email"},
		{"", ""},
	}
	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			if got := SnakeToPascalSqlc(c.in); got != c.want {
				t.Errorf("SnakeToPascalSqlc(%q) = %q, want %q", c.in, got, c.want)
			}
		})
	}
}
