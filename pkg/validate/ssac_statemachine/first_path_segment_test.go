//ff:func feature=validate type=test control=iteration dimension=1 topic=states
//ff:what TestFirstPathSegment — firstPathSegment 분기별 첫 세그먼트 추출 검증
package ssac_statemachine

import (
	"testing"
)

func TestFirstPathSegment(t *testing.T) {
	cases := []struct {
		name string
		path string
		want string
	}{
		{name: "simple leading slash", path: "/orders", want: "orders"},
		{name: "skip brace param then segment", path: "/{id}/items", want: "items"},
		{name: "nested segment", path: "/orders/{id}", want: "orders"},
		{name: "all empty or braces", path: "/{id}/", want: ""},
		{name: "empty string", path: "", want: ""},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := firstPathSegment(tc.path); got != tc.want {
				t.Errorf("firstPathSegment(%q) = %q, want %q", tc.path, got, tc.want)
			}
		})
	}
}
