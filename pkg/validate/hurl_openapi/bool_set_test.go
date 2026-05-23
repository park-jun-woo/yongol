//ff:func feature=validate type=test control=iteration dimension=1 topic=hurl-openapi
//ff:what boolSet — struct{} set을 bool map으로 변환 검증

package hurl_openapi

import "testing"

func TestBoolSet(t *testing.T) {
	cases := []struct {
		name string
		s    map[string]struct{}
		want map[string]bool
	}{
		{
			name: "nil_input",
			s:    nil,
			want: map[string]bool{},
		},
		{
			name: "empty_input",
			s:    map[string]struct{}{},
			want: map[string]bool{},
		},
		{
			name: "single_entry",
			s:    map[string]struct{}{"a": {}},
			want: map[string]bool{"a": true},
		},
		{
			name: "multiple_entries",
			s:    map[string]struct{}{"a": {}, "b": {}, "c": {}},
			want: map[string]bool{"a": true, "b": true, "c": true},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runBoolMapCase(t, boolSet(c.s), c.want)
		})
	}
}
