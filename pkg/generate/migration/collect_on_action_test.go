//ff:func feature=migration type=test control=iteration dimension=1
//ff:what parse_statements_unit_test — parseCreateTable/parseCreateIndex/parseAlterTable + 하위 parse 헬퍼 통합 단위 테스트
package migration

import (
	"testing"
)

func TestCollectOnAction(t *testing.T) {
	cases := []struct {
		toks []string
		want string
		n    int
	}{
		{[]string{"CASCADE"}, "CASCADE", 1},
		{[]string{"set", "null"}, "SET NULL", 2},
		{[]string{"no", "action"}, "NO ACTION", 2},
		{[]string{"restrict", "X"}, "RESTRICT", 1},
		{nil, "", 0},
	}
	for _, c := range cases {
		got, n := collectOnAction(c.toks)
		if got != c.want || n != c.n {
			t.Errorf("collectOnAction(%v) = (%q,%d), want (%q,%d)", c.toks, got, n, c.want, c.n)
		}
	}
}
