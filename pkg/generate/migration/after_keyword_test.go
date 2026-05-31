//ff:func feature=migration type=test control=iteration dimension=1
//ff:what TestAfterKeyword — 키워드 이후(case-insensitive) 추출, 없으면 원본
package migration

import (
	"testing"
)

func TestAfterKeyword(t *testing.T) {
	cases := []struct {
		s, kw, want string
	}{
		{"ALTER TABLE users RENAME TO acc", "RENAME TO", "acc"},
		{"alter table users rename to acc", "RENAME TO", "acc"},
		{"no match here", "MISSING", "no match here"},
	}
	for _, c := range cases {
		if got := afterKeyword(c.s, c.kw); got != c.want {
			t.Errorf("afterKeyword(%q, %q) = %q, want %q", c.s, c.kw, got, c.want)
		}
	}
}
