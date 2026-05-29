//ff:func feature=util type=test control=iteration dimension=1 topic=string-convert
//ff:what PascalToSnake 회귀 테이블 테스트 (ettle/strcase 거동 고정)

package caseconv

import "testing"

func TestPascalToSnake(t *testing.T) {
	cases := []struct{ in, want string }{
		{"UserID", "user_id"},
		{"OrgName", "org_name"},
		{"userName", "user_name"},
		{"HTTPServer", "http_server"},
		// ettle/strcase splits trailing lowercase letter after an uppercase run,
		// so "UserIDs" → "user_i_ds". Documented to pin behaviour for callers.
		{"UserIDs", "user_i_ds"},
	}
	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			if got := PascalToSnake(c.in); got != c.want {
				t.Errorf("PascalToSnake(%q) = %q, want %q", c.in, got, c.want)
			}
		})
	}
}
