//ff:func feature=validate type=test control=iteration dimension=1 topic=hurl-openapi
//ff:what segmentsMatch — 두 정규화 세그먼트 slice의 element-wise 비교 검증

package hurl_openapi

import "testing"

func TestSegmentsMatch(t *testing.T) {
	cases := []struct {
		name string
		a, b []string
		want bool
	}{
		{name: "both_nil", a: nil, b: nil, want: true},
		{name: "both_empty", a: []string{}, b: []string{}, want: true},
		{name: "equal", a: []string{"users", ":param"}, b: []string{"users", ":param"}, want: true},
		{name: "different_len", a: []string{"users"}, b: []string{"users", ":param"}, want: false},
		{name: "different_content", a: []string{"users"}, b: []string{"orders"}, want: false},
		{name: "nil_vs_empty", a: nil, b: []string{}, want: true},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := segmentsMatch(c.a, c.b)
			if got != c.want {
				t.Errorf("segmentsMatch(%v, %v) = %v, want %v", c.a, c.b, got, c.want)
			}
		})
	}
}
