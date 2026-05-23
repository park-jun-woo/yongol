//ff:func feature=validate type=test control=iteration dimension=1 topic=hurl-openapi
//ff:what sortedKeys — bool map의 key를 정렬된 slice로 반환 검증

package hurl_openapi

import "testing"

func TestSortedKeys(t *testing.T) {
	cases := []struct {
		name string
		m    map[string]bool
		want []string
	}{
		{name: "nil", m: nil, want: []string{}},
		{name: "empty", m: map[string]bool{}, want: []string{}},
		{name: "single", m: map[string]bool{"a": true}, want: []string{"a"}},
		{name: "sorted", m: map[string]bool{"c": true, "a": true, "b": true}, want: []string{"a", "b", "c"}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runStringSliceCase(t, sortedKeys(c.m), c.want)
		})
	}
}
