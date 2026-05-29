//ff:func feature=contract type=test control=iteration dimension=1
//ff:what test: TestToSortedSlice — 정렬·중복없음·nil(빈맵) 분기 검증

package contract

import (
	"reflect"
	"testing"
)

func TestToSortedSlice(t *testing.T) {
	cases := []struct {
		name string
		in   map[string]struct{}
		want []string
	}{
		{"nil_map", nil, nil},
		{"empty_map", map[string]struct{}{}, nil},
		{"single", map[string]struct{}{"a": {}}, []string{"a"}},
		{"sorted", map[string]struct{}{"c": {}, "a": {}, "b": {}}, []string{"a", "b", "c"}},
	}
	for _, c := range cases {
		got := toSortedSlice(c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("%s: got %v want %v", c.name, got, c.want)
		}
	}
}
