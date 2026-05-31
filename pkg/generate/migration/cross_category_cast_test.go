//ff:func feature=migration type=test control=iteration dimension=1
//ff:what TestCrossCategoryCast — 숫자↔텍스트 카테고리 변환 판정
package migration

import (
	"testing"
)

func TestCrossCategoryCast(t *testing.T) {
	cases := []struct {
		from, to string
		want     bool
	}{
		{"INTEGER", "TEXT", true},
		{"TEXT", "BIGINT", true},
		{"INTEGER", "BIGINT", false},
		{"TEXT", "VARCHAR", false},
		{"UUID", "TEXT", false},
	}
	for _, c := range cases {
		if got := crossCategoryCast(c.from, c.to); got != c.want {
			t.Errorf("crossCategoryCast(%q,%q) = %v, want %v", c.from, c.to, got, c.want)
		}
	}
}
