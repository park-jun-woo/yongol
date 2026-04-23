//ff:func feature=util type=test control=iteration dimension=1 topic=string-convert
//ff:what SnakeToPascal 회귀 테이블 테스트

package caseconv

import "testing"

func TestSnakeToPascal(t *testing.T) {
	cases := []struct{ in, want string }{
		{"user", "User"},
		{"user_id", "UserId"},
		{"created_at", "CreatedAt"},
		{"one_two_three", "OneTwoThree"},
		{"", ""},
		{"already", "Already"},
		{"__a__", "A"},
		{"per_page", "PerPage"},
	}
	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			if got := SnakeToPascal(c.in); got != c.want {
				t.Errorf("SnakeToPascal(%q) = %q, want %q", c.in, got, c.want)
			}
		})
	}
}
