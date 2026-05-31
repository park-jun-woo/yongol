//ff:func feature=report type=test control=iteration dimension=1 topic=sarif
//ff:what TestSortStrings — in-place 오름차순 정렬 (정상/이미정렬/빈슬라이스) 검증
package sarif

import (
	"testing"
)

func TestSortStrings(t *testing.T) {
	cases := []struct {
		name string
		in   []string
		want []string
	}{
		{"reversed", []string{"c", "b", "a"}, []string{"a", "b", "c"}},
		{"already sorted", []string{"a", "b", "c"}, []string{"a", "b", "c"}},
		{"duplicates", []string{"b", "a", "b"}, []string{"a", "b", "b"}},
		{"single", []string{"x"}, []string{"x"}},
		{"empty", []string{}, []string{}},
		{"rule ids", []string{"XOS-15", "S-27", "M-2"}, []string{"M-2", "S-27", "XOS-15"}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assertSortStrings(t, c.in, c.want)
		})
	}
}
