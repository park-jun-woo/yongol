//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what sortedKeys — map[string]int64 의 key 리스트를 문자열 순으로 정렬

package boot

import "testing"

func TestSortedKeys(t *testing.T) {
	cases := []struct {
		name string
		in   map[string]int64
		want []string
	}{
		{"empty", map[string]int64{}, []string{}},
		{"nil", nil, []string{}},
		{"sorted output", map[string]int64{"c": 3, "a": 1, "b": 2}, []string{"a", "b", "c"}},
		{"single", map[string]int64{"only": 9}, []string{"only"}},
	}
	for _, c := range cases {
		got := sortedKeys(c.in)
		if !equalStrings(got, c.want) {
			t.Errorf("%s: sortedKeys = %v, want %v", c.name, got, c.want)
		}
	}
}
