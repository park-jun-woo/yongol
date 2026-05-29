//ff:func feature=migration type=test control=iteration dimension=1
//ff:what TestFindLineCommentStart — -- 시작 인덱스, single-quote 내부는 무시
package migration

import "testing"

func TestFindLineCommentStart(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"SELECT 1; -- c", 10},
		{"SELECT 1;", -1},
		{"SELECT '-- x';", -1},
		{"a -- b -- c", 2},
	}
	for _, c := range cases {
		if got := findLineCommentStart(c.in); got != c.want {
			t.Errorf("findLineCommentStart(%q) = %d, want %d", c.in, got, c.want)
		}
	}
}
