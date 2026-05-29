//ff:func feature=agent type=test control=iteration dimension=1
//ff:what TestAppendUnique — appendUnique 가 중복은 무시하고 신규만 추가하는지 검증

package agent

import (
	"reflect"
	"testing"
)

func TestAppendUnique(t *testing.T) {
	cases := []struct {
		name  string
		slice []string
		add   string
		want  []string
	}{
		{name: "empty append", slice: nil, add: "a", want: []string{"a"}},
		{name: "new element", slice: []string{"a", "b"}, add: "c", want: []string{"a", "b", "c"}},
		{name: "duplicate first", slice: []string{"a", "b"}, add: "a", want: []string{"a", "b"}},
		{name: "duplicate last", slice: []string{"a", "b"}, add: "b", want: []string{"a", "b"}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := appendUnique(tc.slice, tc.add)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("appendUnique(%v, %q) = %v, want %v", tc.slice, tc.add, got, tc.want)
			}
		})
	}
}
