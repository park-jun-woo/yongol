//ff:func feature=gen-react type=test control=iteration dimension=1
//ff:what tsSingleQuote — 따옴표/백슬래시/개행 이스케이프 및 일반 라벨 통과 검증

package react

import "testing"

func TestTSSingleQuote(t *testing.T) {
	cases := []struct{ in, want string }{
		{"건물 목록", "'건물 목록'"},
		{"it's", `'it\'s'`},
		{`a\b`, `'a\\b'`},
		{"a\nb", `'a\nb'`},
	}
	for _, c := range cases {
		if got := tsSingleQuote(c.in); got != c.want {
			t.Errorf("tsSingleQuote(%q) = %s, want %s", c.in, got, c.want)
		}
	}
}
