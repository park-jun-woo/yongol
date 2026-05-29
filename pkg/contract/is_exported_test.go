//ff:func feature=contract type=test control=iteration dimension=1
//ff:what test: TestIsExported — 대문자 시작/소문자/빈 문자열/유니코드 분기 테이블 검증

package contract

import "testing"

func TestIsExported(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want bool
	}{
		{"upper", "Foo", true},
		{"lower", "foo", false},
		{"empty", "", false},
		{"single_upper", "X", true},
		{"single_lower", "x", false},
		{"underscore", "_priv", false},
		{"digit", "1abc", false},
	}
	for _, c := range cases {
		if got := isExported(c.in); got != c.want {
			t.Errorf("%s: isExported(%q) = %v want %v", c.name, c.in, got, c.want)
		}
	}
}
